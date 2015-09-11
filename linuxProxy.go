// ----------------------------------------------------------------------------
// Linux proxy.
// Fernando Visca. August 2013.
// ----------------------------------------------------------------------------
package main

import (
	"ethos/edbtypes"
	"ethos/nsg"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const DEBUG = true // Toggles debug logging information
func DBG_PRINT(logger *log.Logger, format string, v ...interface{}) {
	if DEBUG {
		logger.Printf(format+"\n", v)
	}
}

const (
	Attach     = iota // Attach to the running process
	qSupported        // Declares supported features.
	// We only indicate the packet maximum supported size
	qSymbol   // Query packet
	qTStatus  // Query packet
	qAttached // Did we attach or spawn a new process?

	qOffsets // The qOffsets query was originally added for NetWare systems.
	// It is used to handle cases where the image being debugged has been relocated by the target.
	// It returns the new addresses of the text, data, and bss sections.

	Hg  // "Hg0". Requests that all operations apply to all of the threads
	why // "?". Why did we stop execution? We reply signal 5...
	Hc  // "Hc-1". Future step/continue operations should apply to all of the threads
	qC  // Is current thread -1?
	g   // Read general registers
	m   // mAA..,LL.. read LL.. bytes at address AA..
	unexpected
)

var hexChars []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}

// Read next packet from gdb and send acknowledgement
func nextPacket(gdbConn net.Conn) ([]byte, int) {
	ret := make([]byte, 1024)
	var length int

	// Ignore packets of length 1 (ack/nak) TODO
	for {
		length, err := gdbConn.Read(ret)
		if err != nil {
			panic(err)
		}
		if length != 1 {
			break
		}
	}

	// Perform packet checksum test TODO

	// Send acknowledgement
	_, err := gdbConn.Write([]byte("+"))
	if err != nil {
		panic(err)
	}

	// Return the package portion only - $...# - without any ack or checksum
	var i, j int
	for i = 0; i < length && ret[i] != '$'; i++ {
	}
	for j = i + 1; j < length && ret[j] != '#'; j++ {
	}

	return ret[i : j+1], length
}

// Send reply to gdb, appending the required checksum
func gdbReply(gdbConn net.Conn, reply string) {
	replyBytes := appendChecksum([]byte(reply))
	_, err := gdbConn.Write(replyBytes)
	if err != nil {
		panic(err)
	}
}

// Appends checksum to a given packet (assume the packet is in the form $...#)
func appendChecksum(packet []byte) []byte {
	_, checksum := checkPacket(packet)
	newPacket := make([]byte, len(packet)+2)

	copy(newPacket, packet)
	newPacket[len(packet)] = hexChars[checksum>>4]
	newPacket[len(packet)+1] = hexChars[checksum%16]

	return newPacket
}

// Packet parsing (assume packet is $...#)
func packetParse(pkt []byte, length int) (ret edbtypes.GdbProxyCall) {
	Id := identify(pkt)
	switch Id {
	default:
		ret.Id = uint8(Id)
	case m:
		ret = mPacket(pkt)
	case unexpected:
		panic("Unexpected packet encountered")
	}
	return
}

func edbRequest(edbPkt edbtypes.GdbProxyCall, en *edbtypes.Encoder, de *edbtypes.Decoder) (ret edbtypes.GdbProxyReply) {
	(*en).EdbRpcDebug(0, &edbPkt)
	(*de).HandleEdbRpc(en)
	ret = edbtypes.ProxyReply
	return
}

// Generate string reply for gdb and determine whether debugging session is to be terminated
func generateGdbReply(edbReply edbtypes.GdbProxyReply) (ret1 string, ret2 bool) {
	ret2 = false // TODO
	switch edbReply.Id {
	default:
		panic("Invalid reply received from GdbProxy")
	case qSupported:
		ret1 = "$PacketSize=" + fmt.Sprintf("%d", edbReply.PacketSize) + "#"
	case Hg, Hc, qSymbol:
		ret1 = "$OK#"
	case why:
		ret1 = "$S05#"
	case qC:
		ret1 = "$#"
	case qAttached:
		ret1 = "$1#"
	case qOffsets:
		ret1 = "$Text=0;Data=0;Bss=0#"
	case g:
		ret1 = gReply(edbReply)
	case m:
		ret1 = mReply(edbReply)
	case qTStatus:
		ret1 = "$tnotrun:0#"
	}
	return
}

// pktmgmt *********************************************************************************

func identify(packet []byte) (index int) {
	p := string(packet)

	if strings.HasPrefix(p, "$qSupported") {
		return qSupported
	} else if strings.HasPrefix(p, "$Hg0") {
		return Hg
	} else if strings.HasPrefix(p, "$?") {
		return why
	} else if strings.HasPrefix(p, "$Hc-1") {
		return Hc
	} else if strings.HasPrefix(p, "$qC") {
		return qC
	} else if strings.HasPrefix(p, "$qAttached") {
		return qAttached
	} else if strings.HasPrefix(p, "$qOffsets") {
		return qOffsets
	} else if strings.HasPrefix(p, "$g") {
		return g
	} else if strings.HasPrefix(p, "$m") {
		return m
	} else if strings.HasPrefix(p, "$qSymbol") {
		return qSymbol
	} else if strings.HasPrefix(p, "$qTStatus") {
		return qTStatus
	}

	return unexpected
}

func mPacket(pkt []byte) (ret edbtypes.GdbProxyCall) {
	ret.Id = m
	var i, j int

	// Parse the address
	for j = 0; true; j++ {
		if pkt[j] == 'm' {
			i = j + 1
		}
		if pkt[j] == ',' {
			ret.Addr = stringToUint32(string(pkt[i:j]), 16)
			break
		}
	}

	// Parse the size
	for i = j + 1; true; j++ {
		if pkt[j] == '#' {
			ret.Size = stringToUint32(string(pkt[i:j]), 16)
			break
		}
	}

	return
}

func uint32ToString(num uint32) (str string) {
	var i uint8
	for i = 0; i < 32; i += 8 {
		b := uint8(num >> i)
		str += string(hexChars[b>>4]) + string(hexChars[b%16])
	}
	return
}

func gReply(edbReply edbtypes.GdbProxyReply) string {
	Regs := edbReply.GReg

	// The order of the registers is defined by gdb
	replyStr := "$"
	replyStr += uint32ToString(Regs.Eax)
	replyStr += uint32ToString(Regs.Ecx)
	replyStr += uint32ToString(Regs.Edx)
	replyStr += uint32ToString(Regs.Ebx)
	replyStr += uint32ToString(Regs.Xesp)
	replyStr += uint32ToString(Regs.Ebp)
	replyStr += uint32ToString(Regs.Esi)
	replyStr += uint32ToString(Regs.Edi)
	replyStr += uint32ToString(Regs.Eip)
	replyStr += uint32ToString(Regs.Eflags)
	replyStr += uint32ToString(Regs.Xcs)
	replyStr += uint32ToString(Regs.Xss)
	replyStr += uint32ToString(Regs.Xds)
	replyStr += uint32ToString(Regs.Xes)
	replyStr += uint32ToString(Regs.Xfs)
	replyStr += uint32ToString(Regs.Xgs)
	replyStr += "#"

	return replyStr
}

func mReply(edbReply edbtypes.GdbProxyReply) string {
	replyStr := "$"
	for c := 0; c < int(edbReply.MemorySize); c++ {
		replyStr += uint8ToString(edbReply.Memory[c])
	}
	replyStr += "#"
	return replyStr
}

func uint8ToString(num uint8) (str string) {
	str = string(hexChars[num>>4]) + string(hexChars[num%16])
	return
}

func stringToUint32(str string, base int) (num uint32) {
	strBytes := []byte(str)
	num = 0

	if base == 10 {
		for i := 0; i < len(strBytes); i++ {
			num *= 10
			if strBytes[i] < '0' || strBytes[i] > '9' {
				panic("Bad parameters in stringToUint32")
			}
			num += uint32(strBytes[i] - '0')
		}
	} else if base == 16 {
		for i := 0; i < len(str); i++ {
			num *= 16
			num += hex(byte(str[i]))
		}
	} else {
		panic("Base not supported in stringToUint32")
	}

	return
}

func hex(ch byte) uint32 {
	if ch >= 'a' && ch <= 'f' {
		return uint32(ch - 'a' + 10)
	}
	if ch >= '0' && ch <= '9' {
		return uint32(ch - '0')
	}
	if ch >= 'A' && ch <= 'F' {
		return uint32(ch - 'A' + 10)
	}
	panic("Unexpected parameter in hex")
	return 0
}

// Check a packet we received or compute the checksum of a packet we are sending
// Returns true if the check succeded and returns the checksum anyway
func checkPacket(packet []byte) (passed bool, sum byte) {
	var i int

	for i = 0; i < len(packet); i++ {
		if packet[i] == '$' {
			continue
		}
		if packet[i] == '#' {
			i++
			break
		}
		sum += packet[i]
	}

	if i < len(packet)-2 { // If no checksum is received, return false
		checksum := stringToUint32(string(packet[i:i+2]), 16)

		if uint32(sum) == checksum {
			return true, sum
		}
	}

	return false, sum
}

// *********************************************************************************************************

func main() {
	var hostname, username string
	var pid uint64
	var port int

	// Create default logger
	logger := log.New(os.Stderr, "", 0)

	// Parse command line parameters
	flag.StringVar(&hostname, "hostname", "test", "Ethos hostname configured with NetStackGo")
	flag.StringVar(&username, "username", "mike", "User who owns the debugging process")
	flag.Uint64Var(&pid, "pid", 10, "Process ID of the debugging process")
	flag.IntVar(&port, "port", 1111, "Port on which to listen for gdb RSP packets")
	flag.Parse()

	DBG_PRINT(logger, "Hostname = %s, Username = %s, PID = %d, Port = %d", hostname, username, pid, port)

	// Debugging service for selected user
	service := "debug/" + username

	// Ipc + Block and retire on debugging service on Ethos
	event, err := nsg.Ipc(hostname, service, service)
	if err != nil {
		panic(err)
	}
	_, c, err := event.BlockAndRetire()
	if err != nil {
		panic(err)
	}
	DBG_PRINT(logger, "Ipc and BlockAndRetire completed")

	en := edbtypes.NewEncoder(c)
	de := edbtypes.NewDecoder(c)

	// Request attach
	en.EdbRpcAttach(0, pid)
	DBG_PRINT(logger, "Attach operation requested")

	// If the request is denied, terminate program
	de.HandleEdbRpc(en)
	if !edbtypes.AttachReply {
		panic("Attach operation refused")
	}
	DBG_PRINT(logger, "Attach operation granted")

	// If the request is granted, wait for gdb to start and begin debugging session
	ln, err := net.Listen("tcp", ":"+fmt.Sprintf("%d", port))
	if err != nil {
		panic(err)
	}
	gdbConn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	DBG_PRINT(logger, "Entering debugging session...")

	// Debugging session using gdbConn, en and de (nsg encoder and decoder)
	for {
		gdbpkt, length := nextPacket(gdbConn) // Get packet $...# and its length
		DBG_PRINT(logger, "Received packet %s of length %d", string(gdbpkt), length)

		edbPkt := packetParse(gdbpkt, length) // Parse the packet
		DBG_PRINT(logger, "Packet parsed. Id = %d; Addr = %d; Size = %d", edbPkt.Id, edbPkt.Addr, edbPkt.Size)

		edbReply := edbRequest(edbPkt, en, de) // Send the packet to GdbProxy on Ethos and receive reply
		DBG_PRINT(logger, "Received reply from GdbProxy")

		gdbReplyString, terminate := generateGdbReply(edbReply) // Read answer from GdbProxy, transform it for gdb and
		// determine whether we have to terminate the debugging session
		DBG_PRINT(logger, "Reply string generated: %s", gdbReplyString)

		if terminate {
			if gdbReplyString != "" {
				gdbReply(gdbConn, gdbReplyString) // Reply to gdb only if required
			}
			break // Terminate debugging session
		} else {
			gdbReply(gdbConn, gdbReplyString) // Send GdbProxy reply to gdb
		}
	}

	defer gdbConn.Close() // Socket
	defer c.Close()       // netStackGo IPC
}

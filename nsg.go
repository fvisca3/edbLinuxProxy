package nsg
import(
	"io"
	"math"
	"reflect"
	"unsafe"
	"llrb"
//	"net"
	"os"
	"fmt"
	"bytes"
	"encoding/hex"
	
)

// FIXME: this is a place holder for bytes. bytes is only used when there is an any type
var xxx = bytes.MinRead
var yyy hex.InvalidHexCharError



type Status uint32


type UdpPort uint16


type Authenticator [64]byte


type IpAddress uint32


type PublicEncryptionKey [32]byte


type HashValue [64]byte


type DirectoryServiceRecord struct {
     
     IpAddress IpAddress
     
     UdpPort UdpPort
     
     PublicKey PublicEncryptionKey
     
}

const(
	
	
	methodIdNetStackGoAdvertise = iota
	
	
	
	methodIdNetStackGoAdvertiseReply
	
	
	
	methodIdNetStackGoBlockAndRetire
	
	
	
	methodIdNetStackGoBlockAndRetireReply
	
	
	
	methodIdNetStackGoDirectoryServiceLookup
	
	
	
	methodIdNetStackGoDirectoryServiceLookupReply
	
	
	
	methodIdNetStackGoDsIpc
	
	
	
	methodIdNetStackGoDsIpcReply
	
	
	
	methodIdNetStackGoImport
	
	
	
	methodIdNetStackGoImportReply
	
	
	
	methodIdNetStackGoIpc
	
	
	
	methodIdNetStackGoIpcReply
	
	
	
	methodIdNetStackGoRead
	
	
	
	methodIdNetStackGoReadReply
	
	
	
	methodIdNetStackGoTerminate
	
	
	
	methodIdNetStackGoWrite
	
	
	
	methodIdNetStackGoWriteReply
	
	
)

func (e *Encoder) NetStackGoAdvertise(cid uint64  ,servicePath string) (err os.Error){
     err = e.uint64(methodIdNetStackGoAdvertise)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.string(servicePath)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) netStackGoAdvertiseReply(cid uint64  ,status *Status) (err os.Error){
     err = e.uint64(methodIdNetStackGoAdvertiseReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) NetStackGoBlockAndRetire(cid uint64  ,eventId uint64) (err os.Error){
     err = e.uint64(methodIdNetStackGoBlockAndRetire)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.uint64(eventId)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) netStackGoBlockAndRetireReply(cid uint64  ,status *Status ,value []byte ,tunnelId uint64 ,connectionId uint32) (err os.Error){
     err = e.uint64(methodIdNetStackGoBlockAndRetireReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.byteSlice(value)
     if err != nil {
     	return err
     }
     
     err = e.uint64(tunnelId)
     if err != nil {
     	return err
     }
     
     err = e.uint32(connectionId)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) NetStackGoDirectoryServiceLookup(cid uint64  ,hostname string ,eventId uint64) (err os.Error){
     err = e.uint64(methodIdNetStackGoDirectoryServiceLookup)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.string(hostname)
     if err != nil {
     	return err
     }
     
     err = e.uint64(eventId)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) netStackGoDirectoryServiceLookupReply(cid uint64  ,toMachine string ,ip *IpAddress ,port *UdpPort ,publicKey *PublicEncryptionKey ,reventId uint64 ,status *Status) (err os.Error){
     err = e.uint64(methodIdNetStackGoDirectoryServiceLookupReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.string(toMachine)
     if err != nil {
     	return err
     }
     
     err = e.ipAddress(ip)
     if err != nil {
     	return err
     }
     
     err = e.udpPort(port)
     if err != nil {
     	return err
     }
     
     err = e.publicEncryptionKey(publicKey)
     if err != nil {
     	return err
     }
     
     err = e.uint64(reventId)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) NetStackGoDsIpc(cid uint64 ) (err os.Error){
     err = e.uint64(methodIdNetStackGoDsIpc)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) netStackGoDsIpcReply(cid uint64  ,status *Status ,eventId uint64) (err os.Error){
     err = e.uint64(methodIdNetStackGoDsIpcReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.uint64(eventId)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) NetStackGoImport(cid uint64  ,servicePath string) (err os.Error){
     err = e.uint64(methodIdNetStackGoImport)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.string(servicePath)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) netStackGoImportReply(cid uint64  ,status *Status ,eventId uint64) (err os.Error){
     err = e.uint64(methodIdNetStackGoImportReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.uint64(eventId)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) NetStackGoIpc(cid uint64  ,toMachine string ,servicePath string) (err os.Error){
     err = e.uint64(methodIdNetStackGoIpc)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.string(toMachine)
     if err != nil {
     	return err
     }
     
     err = e.string(servicePath)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) netStackGoIpcReply(cid uint64  ,status *Status ,eventId uint64) (err os.Error){
     err = e.uint64(methodIdNetStackGoIpcReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.uint64(eventId)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) NetStackGoRead(cid uint64  ,tunnelId uint64 ,connectionId uint32) (err os.Error){
     err = e.uint64(methodIdNetStackGoRead)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.uint64(tunnelId)
     if err != nil {
     	return err
     }
     
     err = e.uint32(connectionId)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) netStackGoReadReply(cid uint64  ,status *Status ,eventId uint64) (err os.Error){
     err = e.uint64(methodIdNetStackGoReadReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.uint64(eventId)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) NetStackGoTerminate(cid uint64 ) (err os.Error){
     err = e.uint64(methodIdNetStackGoTerminate)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) NetStackGoWrite(cid uint64  ,tunnelId uint64 ,connectionId uint32 ,value []byte) (err os.Error){
     err = e.uint64(methodIdNetStackGoWrite)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.uint64(tunnelId)
     if err != nil {
     	return err
     }
     
     err = e.uint32(connectionId)
     if err != nil {
     	return err
     }
     
     err = e.byteSlice(value)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) netStackGoWriteReply(cid uint64  ,status *Status ,eventId uint64) (err os.Error){
     err = e.uint64(methodIdNetStackGoWriteReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.uint64(eventId)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (d *Decoder) HandleNetStackGo(e *Encoder) (error os.Error){
     d.ReadAll()
     methodId, err := d.uint64()
     if err != nil {
     	return err		
     }
//     methodId := *m
     switch *methodId{
     	    
	    case methodIdNetStackGoAdvertise:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		servicePath, err0 := d.string()
		if err0 != nil {
		   return err0
		}
		
		
//		fReturnValue0   := netStackGoAdvertise(*cid, servicePath)
		netStackGoAdvertise(e, *cid, servicePath)
//		e.netStackGoAdvertiseReply(cid , &fReturnValue0)
		
		
	    
	    case methodIdNetStackGoAdvertiseReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		status, err0 := d.status()
		if err0 != nil {
		   return err0
		}
		
		
		netStackGoAdvertiseReply(e, *cid, status)
		
		
	    
	    case methodIdNetStackGoBlockAndRetire:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		eventId, err0 := d.uint64()
		if err0 != nil {
		   return err0
		}
		
		
//		fReturnValue0   , fReturnValue1 , fReturnValue2 , fReturnValue3 := netStackGoBlockAndRetire(*cid, eventId)
		netStackGoBlockAndRetire(e, *cid, eventId)
//		e.netStackGoBlockAndRetireReply(cid , &fReturnValue0, fReturnValue1, fReturnValue2, fReturnValue3)
		
		
	    
	    case methodIdNetStackGoBlockAndRetireReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		status, err0 := d.status()
		if err0 != nil {
		   return err0
		}
		
		value, err1 := d.byteSlice()
		if err1 != nil {
		   return err1
		}
		
		tunnelId, err2 := d.uint64()
		if err2 != nil {
		   return err2
		}
		
		connectionId, err3 := d.uint32()
		if err3 != nil {
		   return err3
		}
		
		
		netStackGoBlockAndRetireReply(e, *cid, status,value,tunnelId,connectionId)
		
		
	    
	    case methodIdNetStackGoDirectoryServiceLookup:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		hostname, err0 := d.string()
		if err0 != nil {
		   return err0
		}
		
		eventId, err1 := d.uint64()
		if err1 != nil {
		   return err1
		}
		
		
//		fReturnValue0   , fReturnValue1 , fReturnValue2 , fReturnValue3 , fReturnValue4 , fReturnValue5 := netStackGoDirectoryServiceLookup(*cid, hostname,eventId)
		netStackGoDirectoryServiceLookup(e, *cid, hostname,eventId)
//		e.netStackGoDirectoryServiceLookupReply(cid , fReturnValue0, &fReturnValue1, &fReturnValue2, &fReturnValue3, fReturnValue4, &fReturnValue5)
		
		
	    
	    case methodIdNetStackGoDirectoryServiceLookupReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		toMachine, err0 := d.string()
		if err0 != nil {
		   return err0
		}
		
		ip, err1 := d.ipAddress()
		if err1 != nil {
		   return err1
		}
		
		port, err2 := d.udpPort()
		if err2 != nil {
		   return err2
		}
		
		publicKey, err3 := d.publicEncryptionKey()
		if err3 != nil {
		   return err3
		}
		
		reventId, err4 := d.uint64()
		if err4 != nil {
		   return err4
		}
		
		status, err5 := d.status()
		if err5 != nil {
		   return err5
		}
		
		
		netStackGoDirectoryServiceLookupReply(e, *cid, toMachine,ip,port,publicKey,reventId,status)
		
		
	    
	    case methodIdNetStackGoDsIpc:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		
//		fReturnValue0   , fReturnValue1 := netStackGoDsIpc(*cid, )
		netStackGoDsIpc(e, *cid, )
//		e.netStackGoDsIpcReply(cid , &fReturnValue0, fReturnValue1)
		
		
	    
	    case methodIdNetStackGoDsIpcReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		status, err0 := d.status()
		if err0 != nil {
		   return err0
		}
		
		eventId, err1 := d.uint64()
		if err1 != nil {
		   return err1
		}
		
		
		netStackGoDsIpcReply(e, *cid, status,eventId)
		
		
	    
	    case methodIdNetStackGoImport:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		servicePath, err0 := d.string()
		if err0 != nil {
		   return err0
		}
		
		
//		fReturnValue0   , fReturnValue1 := netStackGoImport(*cid, servicePath)
		netStackGoImport(e, *cid, servicePath)
//		e.netStackGoImportReply(cid , &fReturnValue0, fReturnValue1)
		
		
	    
	    case methodIdNetStackGoImportReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		status, err0 := d.status()
		if err0 != nil {
		   return err0
		}
		
		eventId, err1 := d.uint64()
		if err1 != nil {
		   return err1
		}
		
		
		netStackGoImportReply(e, *cid, status,eventId)
		
		
	    
	    case methodIdNetStackGoIpc:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		toMachine, err0 := d.string()
		if err0 != nil {
		   return err0
		}
		
		servicePath, err1 := d.string()
		if err1 != nil {
		   return err1
		}
		
		
//		fReturnValue0   , fReturnValue1 := netStackGoIpc(*cid, toMachine,servicePath)
		netStackGoIpc(e, *cid, toMachine,servicePath)
//		e.netStackGoIpcReply(cid , &fReturnValue0, fReturnValue1)
		
		
	    
	    case methodIdNetStackGoIpcReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		status, err0 := d.status()
		if err0 != nil {
		   return err0
		}
		
		eventId, err1 := d.uint64()
		if err1 != nil {
		   return err1
		}
		
		
		netStackGoIpcReply(e, *cid, status,eventId)
		
		
	    
	    case methodIdNetStackGoRead:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		tunnelId, err0 := d.uint64()
		if err0 != nil {
		   return err0
		}
		
		connectionId, err1 := d.uint32()
		if err1 != nil {
		   return err1
		}
		
		
//		fReturnValue0   , fReturnValue1 := netStackGoRead(*cid, tunnelId,connectionId)
		netStackGoRead(e, *cid, tunnelId,connectionId)
//		e.netStackGoReadReply(cid , &fReturnValue0, fReturnValue1)
		
		
	    
	    case methodIdNetStackGoReadReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		status, err0 := d.status()
		if err0 != nil {
		   return err0
		}
		
		eventId, err1 := d.uint64()
		if err1 != nil {
		   return err1
		}
		
		
		netStackGoReadReply(e, *cid, status,eventId)
		
		
	    
	    case methodIdNetStackGoTerminate:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		
		netStackGoTerminate(e, *cid, )
		
		
	    
	    case methodIdNetStackGoWrite:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		tunnelId, err0 := d.uint64()
		if err0 != nil {
		   return err0
		}
		
		connectionId, err1 := d.uint32()
		if err1 != nil {
		   return err1
		}
		
		value, err2 := d.byteSlice()
		if err2 != nil {
		   return err2
		}
		
		
//		fReturnValue0   , fReturnValue1 := netStackGoWrite(*cid, tunnelId,connectionId,value)
		netStackGoWrite(e, *cid, tunnelId,connectionId,value)
//		e.netStackGoWriteReply(cid , &fReturnValue0, fReturnValue1)
		
		
	    
	    case methodIdNetStackGoWriteReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		status, err0 := d.status()
		if err0 != nil {
		   return err0
		}
		
		eventId, err1 := d.uint64()
		if err1 != nil {
		   return err1
		}
		
		
		netStackGoWriteReply(e, *cid, status,eventId)
		
		
	    
	    default:
	    	 e := NewSayIError("Wrong MethodID")
		 return e
     }
     return nil
}

const(
	
	
	methodIdRpcKernelDirectoryServiceLookup = iota
	
	
	
	methodIdRpcKernelDirectoryServiceLookupReply
	
	
	
	methodIdRpcKernelDirectoryServiceRegister
	
	
	
	methodIdRpcKernelIpcWrite
	
	
	
	methodIdRpcKernelIpcWriteReply
	
	
	
	methodIdRpcKernelRemoteConnectionCreate
	
	
	
	methodIdRpcKernelRemoteConnectionCreateAuthenticateUser
	
	
	
	methodIdRpcKernelRemoteConnectionCreateAuthenticateUserReply
	
	
	
	methodIdRpcKernelRemoteConnectionCreateReply
	
	
)

func (e *Encoder) RpcKernelDirectoryServiceLookup(cid uint64  ,hostname string) (err os.Error){
     err = e.uint64(methodIdRpcKernelDirectoryServiceLookup)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.string(hostname)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) rpcKernelDirectoryServiceLookupReply(cid uint64  ,record *DirectoryServiceRecord ,status *Status) (err os.Error){
     err = e.uint64(methodIdRpcKernelDirectoryServiceLookupReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.directoryServiceRecord(record)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) RpcKernelDirectoryServiceRegister(cid uint64  ,hostname string ,record *DirectoryServiceRecord) (err os.Error){
     err = e.uint64(methodIdRpcKernelDirectoryServiceRegister)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.string(hostname)
     if err != nil {
     	return err
     }
     
     err = e.directoryServiceRecord(record)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) RpcKernelIpcWrite(cid uint64  ,contents string) (err os.Error){
     err = e.uint64(methodIdRpcKernelIpcWrite)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.string(contents)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) rpcKernelIpcWriteReply(cid uint64  ,status *Status) (err os.Error){
     err = e.uint64(methodIdRpcKernelIpcWriteReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) RpcKernelRemoteConnectionCreate(cid uint64  ,service string ,typeHash *HashValue ,connectionId uint32 ,writeValue string) (err os.Error){
     err = e.uint64(methodIdRpcKernelRemoteConnectionCreate)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.string(service)
     if err != nil {
     	return err
     }
     
     err = e.hashValue(typeHash)
     if err != nil {
     	return err
     }
     
     err = e.uint32(connectionId)
     if err != nil {
     	return err
     }
     
     err = e.string(writeValue)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) RpcKernelRemoteConnectionCreateAuthenticateUser(cid uint64  ,service string ,typeHash *HashValue ,connectionId uint32 ,userPublicKey *PublicEncryptionKey ,userAuthenticator *Authenticator ,writeValue string) (err os.Error){
     err = e.uint64(methodIdRpcKernelRemoteConnectionCreateAuthenticateUser)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.string(service)
     if err != nil {
     	return err
     }
     
     err = e.hashValue(typeHash)
     if err != nil {
     	return err
     }
     
     err = e.uint32(connectionId)
     if err != nil {
     	return err
     }
     
     err = e.publicEncryptionKey(userPublicKey)
     if err != nil {
     	return err
     }
     
     err = e.authenticator(userAuthenticator)
     if err != nil {
     	return err
     }
     
     err = e.string(writeValue)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) rpcKernelRemoteConnectionCreateAuthenticateUserReply(cid uint64  ,status *Status) (err os.Error){
     err = e.uint64(methodIdRpcKernelRemoteConnectionCreateAuthenticateUserReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) rpcKernelRemoteConnectionCreateReply(cid uint64  ,status *Status) (err os.Error){
     err = e.uint64(methodIdRpcKernelRemoteConnectionCreateReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.status(status)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (d *Decoder) HandleRpcKernel(e *Encoder) (error os.Error){
     d.ReadAll()
     methodId, err := d.uint64()
     if err != nil {
     	return err		
     }
//     methodId := *m
     switch *methodId{
     	    
	    case methodIdRpcKernelDirectoryServiceLookup:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		hostname, err0 := d.string()
		if err0 != nil {
		   return err0
		}
		
		
//		fReturnValue0   , fReturnValue1 := rpcKernelDirectoryServiceLookup(*cid, hostname)
		rpcKernelDirectoryServiceLookup(e, *cid, hostname)
//		e.rpcKernelDirectoryServiceLookupReply(cid , &fReturnValue0, &fReturnValue1)
		
		
	    
	    case methodIdRpcKernelDirectoryServiceLookupReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		record, err0 := d.directoryServiceRecord()
		if err0 != nil {
		   return err0
		}
		
		status, err1 := d.status()
		if err1 != nil {
		   return err1
		}
		
		
		rpcKernelDirectoryServiceLookupReply(e, *cid, record,status)
		
		
	    
	    case methodIdRpcKernelDirectoryServiceRegister:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		hostname, err0 := d.string()
		if err0 != nil {
		   return err0
		}
		
		record, err1 := d.directoryServiceRecord()
		if err1 != nil {
		   return err1
		}
		
		
		rpcKernelDirectoryServiceRegister(e, *cid, hostname,record)
		
		
	    
	    case methodIdRpcKernelIpcWrite:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		contents, err0 := d.string()
		if err0 != nil {
		   return err0
		}
		
		
//		fReturnValue0   := rpcKernelIpcWrite(*cid, contents)
		rpcKernelIpcWrite(e, *cid, contents)
//		e.rpcKernelIpcWriteReply(cid , &fReturnValue0)
		
		
	    
	    case methodIdRpcKernelIpcWriteReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		status, err0 := d.status()
		if err0 != nil {
		   return err0
		}
		
		
		rpcKernelIpcWriteReply(e, *cid, status)
		
		
	    
	    case methodIdRpcKernelRemoteConnectionCreate:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		service, err0 := d.string()
		if err0 != nil {
		   return err0
		}
		
		typeHash, err1 := d.hashValue()
		if err1 != nil {
		   return err1
		}
		
		connectionId, err2 := d.uint32()
		if err2 != nil {
		   return err2
		}
		
		writeValue, err3 := d.string()
		if err3 != nil {
		   return err3
		}
		
		
//		fReturnValue0   := rpcKernelRemoteConnectionCreate(*cid, service,typeHash,connectionId,writeValue)
		rpcKernelRemoteConnectionCreate(e, *cid, service,typeHash,connectionId,writeValue)
//		e.rpcKernelRemoteConnectionCreateReply(cid , &fReturnValue0)
		
		
	    
	    case methodIdRpcKernelRemoteConnectionCreateAuthenticateUser:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		service, err0 := d.string()
		if err0 != nil {
		   return err0
		}
		
		typeHash, err1 := d.hashValue()
		if err1 != nil {
		   return err1
		}
		
		connectionId, err2 := d.uint32()
		if err2 != nil {
		   return err2
		}
		
		userPublicKey, err3 := d.publicEncryptionKey()
		if err3 != nil {
		   return err3
		}
		
		userAuthenticator, err4 := d.authenticator()
		if err4 != nil {
		   return err4
		}
		
		writeValue, err5 := d.string()
		if err5 != nil {
		   return err5
		}
		
		
//		fReturnValue0   := rpcKernelRemoteConnectionCreateAuthenticateUser(*cid, service,typeHash,connectionId,userPublicKey,userAuthenticator,writeValue)
		rpcKernelRemoteConnectionCreateAuthenticateUser(e, *cid, service,typeHash,connectionId,userPublicKey,userAuthenticator,writeValue)
//		e.rpcKernelRemoteConnectionCreateAuthenticateUserReply(cid , &fReturnValue0)
		
		
	    
	    case methodIdRpcKernelRemoteConnectionCreateAuthenticateUserReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		status, err0 := d.status()
		if err0 != nil {
		   return err0
		}
		
		
		rpcKernelRemoteConnectionCreateAuthenticateUserReply(e, *cid, status)
		
		
	    
	    case methodIdRpcKernelRemoteConnectionCreateReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		 cid, err := d.uint64()
     		 if err != nil {
		      	return err		
		 }
		 // 
		 // cid := *m
		 // 
		
		status, err0 := d.status()
		if err0 != nil {
		   return err0
		}
		
		
		rpcKernelRemoteConnectionCreateReply(e, *cid, status)
		
		
	    
	    default:
	    	 e := NewSayIError("Wrong MethodID")
		 return e
     }
     return nil
}

func (e *Encoder) directoryServiceRecord(v *DirectoryServiceRecord) (err os.Error){
     
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(v), "f96c20bf07fde529982e2b24768fc5b04ec3a9d8dc50077e9d54ca95187f6ac26707b5373363867778aa6eb0ba4b721e8903cfeac689b8f4d02c913a49f8d0f2", uint64(unsafe.Sizeof(*v))); !overlap{
     	if error != nil {
	   return error
	}
     	
       err = e.ipAddressInternal(&v.IpAddress)
       if err != nil {
	   return err
	}
       
       err = e.udpPortInternal(&v.UdpPort)
       if err != nil {
	   return err
	}
       
       err = e.publicEncryptionKeyInternal(&v.PublicKey)
       if err != nil {
	   return err
	}
        
     }   
     return nil 
}

func (e *Encoder) DirectoryServiceRecord (v *DirectoryServiceRecord) (err os.Error){
     return e.directoryServiceRecord(v)
}
func (d *Decoder) directoryServiceRecord() (v *DirectoryServiceRecord, error os.Error){
     
     
     var valv DirectoryServiceRecord
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, err0 := d.ipAddressInternal()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.IpAddress = *p0
     
     
     		
     p1, err1 := d.udpPortInternal()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.UdpPort = *p1
     
     
     		
     p2, err2 := d.publicEncryptionKeyInternal()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.PublicKey = *p2
     
     
     v = &valv
     return v, nil
}
func (d *Decoder) DirectoryServiceRecord() (v *DirectoryServiceRecord, error os.Error){
     return d.directoryServiceRecord()
}
func (e *Encoder) status(v *Status) (err os.Error){
     
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(v), "8570d2cdf4597ef2800a5d8839fe5635cde28784abe0385d5a40c79117e14046178b07d6dd67366cbcba790a4c01f010050b13786ff84e047235de32eb31d14d", uint64(unsafe.Sizeof(*v))); !overlap{
     	if error != nil {
	   return error
	}
     	
       err = e.uint32(uint32(*v))
       if err != nil {
	   return err
	}
        
     }   
     return nil 
}

func (e *Encoder) Status (v *Status) (err os.Error){
     return e.status(v)
}
func (d *Decoder) status() (v *Status, error os.Error){
     
     
     var valv Status
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, err0 := d.uint32()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = Status(*p0)
     
     
     v = &valv
     return v, nil
}
func (d *Decoder) Status() (v *Status, error os.Error){
     return d.status()
}
func (e *Encoder) publicEncryptionKey(v *PublicEncryptionKey) (err os.Error){
     
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(v), "d2272db510879377b56191da150b208fcb77221c31a432b98407e05805b984d62505ef7f976421593c806b4a620fd4fea820a22e32b8584c08b585cf02d0aaf9", uint64(unsafe.Sizeof(*v))); !overlap{
     	if error != nil {
	   return error
	}
     	
       err = e.byte_32Slice([32]byte(*v))
       if err != nil {
	   return err
	}
        
     }   
     return nil 
}

func (e *Encoder) PublicEncryptionKey (v *PublicEncryptionKey) (err os.Error){
     return e.publicEncryptionKey(v)
}
func (d *Decoder) publicEncryptionKey() (v *PublicEncryptionKey, error os.Error){
     
     
     var valv PublicEncryptionKey
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, err0 := d.byte_32Slice()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = PublicEncryptionKey(*p0)
     
     
     v = &valv
     return v, nil
}
func (d *Decoder) PublicEncryptionKey() (v *PublicEncryptionKey, error os.Error){
     return d.publicEncryptionKey()
}
func (e *Encoder) byteSlice(v []byte) (err os.Error){
     
     if len(v) == 0 {
     	
	err = e.SliceOfBytes(v[:])
	if err != nil {
	   return err
	}
	
	return nil
     }
     //vValue := reflect.ValueOf(&v).Elem()
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(&v[0]), "81bd4454423be6fa612b61ddca1d2621742888e293d25b55deb14cdeac37dad1c0d4dfbcd49891dfebde8de850fe458f412b397298df87e5d588b6c168ca3655", uint64(len(v))); !overlap{
     	if error != nil {
	   return error
	}
     	
	err = e.SliceOfBytes(v[:])
	if err != nil {
	   return err
	}
	
     }    
     return nil
}

func (d *Decoder) byteSlice() (v *[]byte, error os.Error){
     
     var valv []byte
     
     l, err := d.uint32()
     if err != nil {
     	return &valv, err
     }
     length := *l
     valv = make([]byte, length)
     
     
     error = d.SliceOfBytes(valv[:], length)
     if error != nil {
     	return &valv, error
     }
     
     d.indexToValue = append(d.indexToValue, v)
     return &valv, nil
}

func (e *Encoder) byte_64Slice(v [64]byte) (err os.Error){
     
     if len(v) == 0 {
     	
	
     	for _, ele := range v {
	    
       	    err = e.byte(ele)
	    if err != nil {
	       return err
	    }
       	     
	}
	
	return nil
     }
     //vValue := reflect.ValueOf(&v).Elem()
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(&v[0]), "487c4cf7536d19fbd620cfc0339e04a9eb85eac426fdf82ea889461fedb11e18de226490d8790e347eceef9db7d07b2aca46f4abed324dc4288803aa06800c25", uint64(len(v))); !overlap{
     	if error != nil {
	   return error
	}
     	
	
     	for _, ele := range v {
	    
       	    err = e.byte(ele)
	    if err != nil {
	       return err
	    }
       	     
	}
	
     }    
     return nil
}

func (d *Decoder) byte_64Slice() (v *[64]byte, error os.Error){
     
     var valv [64]byte
     
     length := 64
     
     
     i := 0
     for ; length > 0 ; length-- {
     	 
	 
	 p, err := d.byte()
	 if err != nil {
	    return &valv, err
	 }
	 valv[i] = *p
     	 
	 i++ 
     }
     
     d.indexToValue = append(d.indexToValue, v)
     return &valv, nil
}

func (e *Encoder) authenticator(v *Authenticator) (err os.Error){
     
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(v), "a4e28df1e32067b09bead002023904205a8a8f5186d07a5d236c73ecec188b661a194a0f38f6ff6da994dd033bb970efc9e25b48728deb5bcf2df39fa8006996", uint64(unsafe.Sizeof(*v))); !overlap{
     	if error != nil {
	   return error
	}
     	
       err = e.byte_64Slice([64]byte(*v))
       if err != nil {
	   return err
	}
        
     }   
     return nil 
}

func (e *Encoder) Authenticator (v *Authenticator) (err os.Error){
     return e.authenticator(v)
}
func (d *Decoder) authenticator() (v *Authenticator, error os.Error){
     
     
     var valv Authenticator
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, err0 := d.byte_64Slice()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = Authenticator(*p0)
     
     
     v = &valv
     return v, nil
}
func (d *Decoder) Authenticator() (v *Authenticator, error os.Error){
     return d.authenticator()
}
func (e *Encoder) ipAddress(v *IpAddress) (err os.Error){
     
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(v), "c23f712cbb2992216390af71585511c75d3938bcd384b3c9070ed1afb81c618bfe1593dc1afb40fb7fde4e0a301997859bf9ed957648b46e43c48c53185d4678", uint64(unsafe.Sizeof(*v))); !overlap{
     	if error != nil {
	   return error
	}
     	
       err = e.uint32(uint32(*v))
       if err != nil {
	   return err
	}
        
     }   
     return nil 
}

func (e *Encoder) IpAddress (v *IpAddress) (err os.Error){
     return e.ipAddress(v)
}
func (d *Decoder) ipAddress() (v *IpAddress, error os.Error){
     
     
     var valv IpAddress
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, err0 := d.uint32()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = IpAddress(*p0)
     
     
     v = &valv
     return v, nil
}
func (d *Decoder) IpAddress() (v *IpAddress, error os.Error){
     return d.ipAddress()
}
func (e *Encoder) udpPort(v *UdpPort) (err os.Error){
     
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(v), "a0bbbd31541c1d47ef8e332619cb6bcc97d09a4b580ea6c644cf415fce148eceb24c6fa58f19c193604944317d950c22e6533164ff69213dcc4bc8bbcb6e7e62", uint64(unsafe.Sizeof(*v))); !overlap{
     	if error != nil {
	   return error
	}
     	
       err = e.uint16(uint16(*v))
       if err != nil {
	   return err
	}
        
     }   
     return nil 
}

func (e *Encoder) UdpPort (v *UdpPort) (err os.Error){
     return e.udpPort(v)
}
func (d *Decoder) udpPort() (v *UdpPort, error os.Error){
     
     
     var valv UdpPort
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, err0 := d.uint16()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = UdpPort(*p0)
     
     
     v = &valv
     return v, nil
}
func (d *Decoder) UdpPort() (v *UdpPort, error os.Error){
     return d.udpPort()
}
func (e *Encoder) byte_32Slice(v [32]byte) (err os.Error){
     
     if len(v) == 0 {
     	
	
     	for _, ele := range v {
	    
       	    err = e.byte(ele)
	    if err != nil {
	       return err
	    }
       	     
	}
	
	return nil
     }
     //vValue := reflect.ValueOf(&v).Elem()
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(&v[0]), "76b222b85eca6936e97ba840bc2e682264213905c1477bdf220a86bcdc61b31bed6205365d2c57d7aa75698c3e0335727670a77a8c8599c5f93b9cc3772b4789", uint64(len(v))); !overlap{
     	if error != nil {
	   return error
	}
     	
	
     	for _, ele := range v {
	    
       	    err = e.byte(ele)
	    if err != nil {
	       return err
	    }
       	     
	}
	
     }    
     return nil
}

func (d *Decoder) byte_32Slice() (v *[32]byte, error os.Error){
     
     var valv [32]byte
     
     length := 32
     
     
     i := 0
     for ; length > 0 ; length-- {
     	 
	 
	 p, err := d.byte()
	 if err != nil {
	    return &valv, err
	 }
	 valv[i] = *p
     	 
	 i++ 
     }
     
     d.indexToValue = append(d.indexToValue, v)
     return &valv, nil
}

func (e *Encoder) hashValue(v *HashValue) (err os.Error){
     
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(v), "f637a654c5e2c5a71568247b9be447df52abab489ba366a38c10a4db80d4f236698d8b0549209578291b2425c8229200bbc159c683df6e41deb80092415f25a2", uint64(unsafe.Sizeof(*v))); !overlap{
     	if error != nil {
	   return error
	}
     	
       err = e.byte_64Slice([64]byte(*v))
       if err != nil {
	   return err
	}
        
     }   
     return nil 
}

func (e *Encoder) HashValue (v *HashValue) (err os.Error){
     return e.hashValue(v)
}
func (d *Decoder) hashValue() (v *HashValue, error os.Error){
     
     
     var valv HashValue
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, err0 := d.byte_64Slice()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = HashValue(*p0)
     
     
     v = &valv
     return v, nil
}
func (d *Decoder) HashValue() (v *HashValue, error os.Error){
     return d.hashValue()
}
func (e *Encoder) directoryServiceRecordInternal(v *DirectoryServiceRecord) (err os.Error){
     
     
     	     err = e.ipAddressInternal(&v.IpAddress)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.udpPortInternal(&v.UdpPort)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.publicEncryptionKeyInternal(&v.PublicKey)
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) directoryServiceRecordInternal() (v *DirectoryServiceRecord, error os.Error){
     
      
     var valv DirectoryServiceRecord
     
          
     p0, err0 := d.ipAddressInternal()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.IpAddress = *p0
     
     
          
     p1, err1 := d.udpPortInternal()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.UdpPort = *p1
     
     
          
     p2, err2 := d.publicEncryptionKeyInternal()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.PublicKey = *p2
     
     
     v = &valv     
     return v, nil
}

func (e *Encoder) statusInternal(v *Status) (err os.Error){
     
     
     	     err = e.uint32(uint32(*v))
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) statusInternal() (v *Status, error os.Error){
     
      
     var valv Status
     
          
     p0, err0 := d.uint32()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = Status(*p0)
     
     
     v = &valv     
     return v, nil
}

func (e *Encoder) publicEncryptionKeyInternal(v *PublicEncryptionKey) (err os.Error){
     
     
     	     err = e.byte_32Slice([32]byte(*v))
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) publicEncryptionKeyInternal() (v *PublicEncryptionKey, error os.Error){
     
      
     var valv PublicEncryptionKey
     
          
     p0, err0 := d.byte_32Slice()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = PublicEncryptionKey(*p0)
     
     
     v = &valv     
     return v, nil
}

func (e *Encoder) byteSliceInternal(v []byte) (err os.Error){
     
     
     err = e.SliceOfBytes(v[:])
     if err != nil {
	   return err
     }
      
     return nil
}

func (d *Decoder) byteSliceInternal() (v *[]byte, error os.Error){
     
     var valv []byte	
     
     l, err := d.uint32()
		 if err != nil {
		    return &valv, err
		 }
     length := *l
     valv = make([]byte, length)
     
     
     error = d.SliceOfBytes(valv[:], length)
     if error != nil {
     	return &valv, error
     }
     
     return &valv, nil
}

func (e *Encoder) byte_64SliceInternal(v [64]byte) (err os.Error){
     
     
	
     for _, ele := range v {
	 
       	 err = e.byte(ele)
	 if err != nil {
	   return err
	 }
       	  
     }
      
     return nil
}

func (d *Decoder) byte_64SliceInternal() (v *[64]byte, error os.Error){
     
     var valv [64]byte	
     
     length := 64
     
     
     i := 0
      for ; length > 0 ; length-- {
     	 
	 
	 p, err := d.byte()
	 if err != nil {	
	    return &valv, err
	 }
	 valv[i] = *p
     	 
	 i++ 
     }
     
     return &valv, nil
}

func (e *Encoder) authenticatorInternal(v *Authenticator) (err os.Error){
     
     
     	     err = e.byte_64Slice([64]byte(*v))
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) authenticatorInternal() (v *Authenticator, error os.Error){
     
      
     var valv Authenticator
     
          
     p0, err0 := d.byte_64Slice()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = Authenticator(*p0)
     
     
     v = &valv     
     return v, nil
}

func (e *Encoder) ipAddressInternal(v *IpAddress) (err os.Error){
     
     
     	     err = e.uint32(uint32(*v))
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) ipAddressInternal() (v *IpAddress, error os.Error){
     
      
     var valv IpAddress
     
          
     p0, err0 := d.uint32()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = IpAddress(*p0)
     
     
     v = &valv     
     return v, nil
}

func (e *Encoder) udpPortInternal(v *UdpPort) (err os.Error){
     
     
     	     err = e.uint16(uint16(*v))
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) udpPortInternal() (v *UdpPort, error os.Error){
     
      
     var valv UdpPort
     
          
     p0, err0 := d.uint16()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = UdpPort(*p0)
     
     
     v = &valv     
     return v, nil
}

func (e *Encoder) byte_32SliceInternal(v [32]byte) (err os.Error){
     
     
	
     for _, ele := range v {
	 
       	 err = e.byte(ele)
	 if err != nil {
	   return err
	 }
       	  
     }
      
     return nil
}

func (d *Decoder) byte_32SliceInternal() (v *[32]byte, error os.Error){
     
     var valv [32]byte	
     
     length := 32
     
     
     i := 0
      for ; length > 0 ; length-- {
     	 
	 
	 p, err := d.byte()
	 if err != nil {	
	    return &valv, err
	 }
	 valv[i] = *p
     	 
	 i++ 
     }
     
     return &valv, nil
}

func (e *Encoder) hashValueInternal(v *HashValue) (err os.Error){
     
     
     	     err = e.byte_64Slice([64]byte(*v))
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) hashValueInternal() (v *HashValue, error os.Error){
     
      
     var valv HashValue
     
          
     p0, err0 := d.byte_64Slice()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv = HashValue(*p0)
     
     
     v = &valv     
     return v, nil
}

const(
	pNIL = 0
	pIDX = 1
	pVAL = 2
	BufSize = 1024 * 50
)
type Encoder struct {
	w io.Writer
	buf []byte
	t *TypeTree
	m []interface{} // this is for storing maps
	curPos int
	bufSpace uint32
	bufStart uint32
	bufLen uint32
	count uint32
}
type Decoder struct {
	r io.Reader
	buf []byte
	indexToValue []interface{}
	m []interface{} // this is for storing maps
	curPos uint32
	bufStart uint32
	bufLen uint32
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w:w, buf:make([]byte, BufSize), t: NewTypeTree(), curPos:4, bufSpace:BufSize, bufStart:0, bufLen:0, count:0}
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r:r, buf:make([]byte, BufSize), curPos:0, bufStart:0, bufLen:0}
}

func (e *Encoder) MapCheck(t interface{}) (int, bool) {
	for index, entry := range e.m {
		if entry == t {
			return index, true
		}
	}
	return -1, false
}

func (e *Encoder) isEnoughSpace(length uint32) bool {
     if length <= e.bufSpace {
     	return true
     }    
     return false
}

func (d *Decoder) isEnoughData(length uint32) bool {
     if length <= d.bufLen {
     	return true
     }
     return false
}


func (d *Decoder) readAtLeast(length uint32) os.Error{
     if d.bufLen > 0 {
     	copy(d.buf, d.buf[d.bufStart:(d.bufStart + d.bufLen)])
     }
     n, err := io.ReadAtLeast(d.r, d.buf[d.bufLen:], int(length))
     if err != nil {
	return err
     }
     d.bufLen += uint32(n)
     d.bufStart = 0
     return nil
}

func (e *Encoder) byte(b byte) (os.Error){
     return e.uint8(uint8(b))
     
}

func (e *Encoder) Byte(b byte) (os.Error){
     return e.byte(b)
}

func (d *Decoder) byte() (b *byte, error os.Error) {
     v, err := d.uint8()
     value := byte(*v)
     return &value, err
}

func (d *Decoder) Byte() (b *byte, err os.Error) {
     return d.byte()
}

func (e *Encoder) uint8(u uint8) os.Error{
     	// e.write([]byte{byte(u)})
	// e.buf[e.curPos] = byte(u)
	// e.curPos++
	if !e.isEnoughSpace(1) {
	   err := e.Flush() 
	   if err != nil {
	      return err
	   }
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen++
	e.bufSpace--
	return nil
}

func (e *Encoder) Uint8(u uint8) (os.Error){
     return e.uint8(u)
}

func (e *Encoder) int8(u int8) (os.Error){
	return e.uint8(uint8(u))
}

func (e *Encoder) Int8(u int8) (os.Error){
     return e.int8(u)
}

func (d *Decoder) uint8() (w *uint8, err os.Error) {
	if !d.isEnoughData(1) {
	   err = d.readAtLeast(1)	   
	   if err != nil {
	      return nil, err
	   }
	}
	v := uint8(d.buf[d.bufStart])
	d.bufStart = (d.bufStart + 1) % BufSize
	d.bufLen -= 1
	return &v, nil
}

func (d *Decoder) Uint8() (w *uint8, err os.Error) {
     return d.uint8()
}

func (d *Decoder) int8() (w *int8, error os.Error) {
     	v, err := d.uint8()
	r := int8(*v)
	return &r, err
}

func (d *Decoder) Int8() (w *int8, err os.Error) {
     return d.int8()
}

func (e *Encoder) uint16(u uint16) os.Error{
	if !e.isEnoughSpace(2) {
	   err := e.Flush() 
	   if err != nil {
	      return err
	   }
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 8)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen += 2
	e.bufSpace -= 2
	return nil
}

func (e *Encoder) Uint16(u uint16) (os.Error){
     return e.uint16(u)
}

func (e *Encoder) int16(u int16) (os.Error){

	return e.uint16(uint16(u))
}

func (e *Encoder) Int16(u int16) (os.Error){
     return e.int16(u)
}

func (d *Decoder) uint16() (w *uint16, err os.Error) {
	if !d.isEnoughData(2) {
	   err = d.readAtLeast(2)	   
	   if err != nil {
	      return nil, err
	   }
	}
	v := uint16(d.buf[d.bufStart]) | uint16(d.buf[d.bufStart + 1]) << 8
	d.bufStart = (d.bufStart + 2) % BufSize
	d.bufLen -= 2
	return &v, nil
}

func (d *Decoder) Uint16() (w *uint16, err os.Error) {
     return d.uint16()
}

func (d *Decoder) int16() (w *int16, error os.Error) {
     	v, err := d.uint16()
	r := int16(*v)
	return &r, err
}

func (d *Decoder) Int16() (w *int16, err os.Error) {
     return d.int16()
}

func (e *Encoder) uint32(u uint32) os.Error{
	if !e.isEnoughSpace(4) {
	   err := e.Flush()   		      
	   if err != nil {
	      return err
	   }
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 8)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 16)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 24)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen += 4
	e.bufSpace -= 4
	return nil
}

func (e *Encoder) Uint32(u uint32) (os.Error){
     return e.uint32(u)
}

func (e *Encoder) int32(u int32) (os.Error){
	return e.uint32(uint32(u))
}

func (e *Encoder) Int32(u int32) (os.Error){
     return e.int32(u)
}

func (d *Decoder) uint32() (w *uint32, err os.Error) {
	if !d.isEnoughData(4) {
	   err = d.readAtLeast(4)	   
	   if err != nil {
	      return nil, err
	   }
	}
	v := uint32(d.buf[d.bufStart]) | uint32(d.buf[d.bufStart + 1]) << 8 | uint32(d.buf[d.bufStart + 2]) << 16 | uint32(d.buf[d.bufStart + 3]) << 24
	d.bufStart = (d.bufStart + 4) % BufSize
	d.bufLen -= 4
	return &v, nil
}

func (d *Decoder) Uint32() (w *uint32, err os.Error) {
     return d.uint32()
}

func (d *Decoder) int32() (w *int32, error os.Error) {
     	v, err := d.uint32()
	r := int32(*v)
	return &r, err
}

func (d *Decoder) Int32() (w *int32, err os.Error) {
     return d.int32()
}

func (e *Encoder) uint64(u uint64) os.Error{
	if !e.isEnoughSpace(8) {
	   	err := e.Flush()    
		if err != nil {
	      	   return err
	      	}  
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 8)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 16)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 24)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 32)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 40)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 48)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 56)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen += 8
	e.bufSpace -= 8
	return nil
}

func (e *Encoder) Uint64(u uint64) (os.Error){
     return e.uint64(u)
}

func (e *Encoder) int64(u int64) (os.Error){
	return e.uint64(uint64(u))
}

func (e *Encoder) Int64(u int64) (os.Error){
     return e.int64(u)
}

func (d *Decoder) uint64() (w *uint64, err os.Error) {
	if !d.isEnoughData(8) {
	   err = d.readAtLeast(8)	   
	   if err != nil {
	      return nil, err
	   }
	}
	v := uint64(d.buf[d.bufStart]) | uint64(d.buf[d.bufStart + 1]) << 8 | uint64(d.buf[d.bufStart + 2]) << 16 | uint64(d.buf[d.bufStart + 3]) << 24 | uint64(d.buf[d.bufStart + 4]) << 32 | uint64(d.buf[d.bufStart + 5]) << 40 | uint64(d.buf[d.bufStart + 6]) << 48 | uint64(d.buf[d.bufStart + 7]) << 56
	d.bufStart = (d.bufStart + 8) % BufSize
	d.bufLen -= 8
	return &v, nil
}

func (d *Decoder) Uint64() (w *uint64, err os.Error) {
     return d.uint64()
}

func (d *Decoder) int64() (w *int64, error os.Error) {
     	v, err := d.uint64()
	r := int64(*v)
	return &r, err
}

func (d *Decoder) Int64() (w *int64, err os.Error) {
     return d.int64()
}

func (e *Encoder) float32(u float32) (os.Error){
	return e.uint32(math.Float32bits(u))
}

func (e *Encoder) Float32(u float32) (os.Error){
     return e.float32(u)
}

func (d *Decoder) float32() (w *float32, error os.Error) {
     	v, err := d.uint32()
	r := math.Float32frombits(*v)
	return &r, err
}

func (d *Decoder) Float32() (w *float32, err os.Error) {
     return d.float32()
}

func (e *Encoder) float64(u float64) (os.Error){
	return e.uint64(math.Float64bits(u))
}

func (e *Encoder) Float64(u float64) (os.Error){
     return e.float64(u)
}

func (d *Decoder) float64() (w *float64, error os.Error) {
     	v, err := d.uint64()   
	r := math.Float64frombits(*v)
	return &r, err
}

func (d *Decoder) Float64() (w *float64, err os.Error) {
     return d.float64()
}

func (e *Encoder) bool(u bool) (err os.Error){
	if u {
		err = e.uint8(1)
	} else {
		err = e.uint8(0)
	}
	return err
}

func (e *Encoder) Bool(u bool) (os.Error){
     return e.bool(u)
}

func (d *Decoder) bool() (w *bool, error os.Error) {
	v, err := d.uint8()
	var u bool
	if *v == 1 {
		u = true
	} else {
		u = false
	}
	return &u, err
}

func (d *Decoder) Bool() (w *bool, err os.Error) {
     return d.bool()
}

func (e *Encoder) SliceOfBytes(u []byte) (err os.Error){
     err = e.length(uint32(len(u)))
     if err != nil {
     	return err
     }
     sliceStartPos := uint32(0)
     for ;!e.isEnoughSpace(uint32(len(u[sliceStartPos:]))); {
	copy(e.buf[e.bufStart:], u[sliceStartPos:(sliceStartPos + e.bufSpace)])
	sliceStartPos += e.bufSpace
	e.bufLen += e.bufSpace
	e.bufSpace = 0
	if e.bufLen != 0 {
	   err := e.Flush()
	   if err != nil {
	      return err
	   }
	}
     } 
     if len(u[sliceStartPos:]) > 0 {
       copy(e.buf[e.bufStart:], u[sliceStartPos:])
       e.bufStart += uint32(len(u[sliceStartPos:]))	
       e.bufLen += uint32(len(u[sliceStartPos:]))
       e.bufSpace -= uint32(len(u[sliceStartPos:]))
     }
     return nil
     
}

func (d *Decoder) SliceOfBytes(v []byte, length uint32) (err os.Error){
     if length > d.bufLen {
     	copy(v, d.buf[d.bufStart:(d.bufStart + d.bufLen)])
	io.ReadFull(d.r, v[d.bufLen:])
	d.bufStart = 0
	d.bufLen = 0
	return
     }
     if !d.isEnoughData(length) {
	   err = d.readAtLeast(length)
	   if err != nil {
	      return err
	   }
	}
     copy(v, d.buf[d.bufStart: (d.bufStart + length)])
     d.bufStart = (d.bufStart + length) % BufSize
     d.bufLen -= length
     return nil
}

func (e *Encoder) string(u string) (err os.Error){
	err = e.length(uint32(len(u)))
	if err != nil {
	   return err
	}
	stringStartPos := uint32(0)
	for ;!e.isEnoughSpace(uint32(len(u[stringStartPos:]))); {
	   copy(e.buf[e.bufStart:], u[stringStartPos:(stringStartPos + e.bufSpace)])
	   stringStartPos += e.bufSpace
	   e.bufLen += e.bufSpace
	   e.bufSpace = 0
	   if e.bufSpace == 0 {
	      err = e.Flush()
	      if err != nil {
	      	 return err
	      }
	   }	  
	} 
	if len(u[stringStartPos:]) > 0 {
		copy(e.buf[e.bufStart:], u[stringStartPos:])
		e.bufStart += uint32(len(u[stringStartPos:]))
		e.bufLen += uint32(len(u[stringStartPos:]))
		e.bufSpace -= uint32(len(u[stringStartPos:]))
	}
	return nil
	
}

func (e *Encoder) String(u string) (os.Error){
     return e.string(u)
}

func (d *Decoder) string() (w *string, err os.Error) {
	len, err := d.length()
	if err != nil {
	   return nil, err
	}

	if len > d.bufLen {
 	   b := make([]byte, len) 
	   copy(b[0:], d.buf[d.bufStart:(d.bufStart + d.bufLen)])
	   _, err = io.ReadFull(d.r, b[d.bufLen:])
	   if err != nil {
	      fmt.Println("Read error:", err)
	   }
	   d.bufStart = 0
	   d.bufLen = 0	   
	   str := string(b)
	   return &str, nil
	}
	if !d.isEnoughData(uint32(len)) {
	  	   err = d.readAtLeast(uint32(len))
		   if err != nil {
	      	      return nil, err
		   }  
	}
	b := d.buf[d.bufStart:(d.bufStart + len)]
	d.bufStart = (d.bufStart + len) % BufSize
	d.bufLen -= len
	str := string(b)
	return &str, nil
}

func (d *Decoder) String() (w *string, err os.Error) {
     return d.string()
}

func (e *Encoder) length(l uint32) (os.Error){
	// e.buf[e.curPos] = byte(l)
	// e.curPos++
	// e.buf[e.curPos] = byte(l >> 8)
	// e.curPos++
	// e.buf[e.curPos] = byte(l >> 16)
	// e.curPos++
	// e.buf[e.curPos] = byte(l >> 24)
	// e.curPos++
	// return
	return e.uint32(l)
}

func (d *Decoder) length() (l uint32, error os.Error) {
     	v, err := d.uint32()
	return *v, err
}

func (e *Encoder) Flush() os.Error{
	// e.bufSize = uint32(e.curPos) - 4
	// e.buf[0] = byte(e.bufSize)
	// e.buf[1] = byte(e.bufSize >> 8)
	// e.buf[2] = byte(e.bufSize >> 16)
	// e.buf[3] = byte(e.bufSize >> 24)
	// e.write(e.buf[:e.curPos])
	// e.curPos = 4
	
        // e.count++
	// if e.count %40 != 0 {
	//     return 
	// }
	if e.bufLen == 0 {
	   return nil
	}
	if _, err := e.w.Write(e.buf[:e.bufLen]); err != nil {
	   return err
	}
	e.bufStart = 0
	e.bufLen = 0
	e.bufSpace = BufSize
	e.reset()
	return nil
}

func (d *Decoder) ReadAll() {
     /* FIXME: Temporarily place reset here. We also reset in RPC calls. This 
      *        should probably go in one place: each type decoder, but we need
      *        to be able to differentiate between an application-called decode
      *        and a decode of internal elements. The latter case should not 
      *        call reset. */
     d.reset() 
}

func (e *Encoder) reset() {
	//e.t = NewTypeTree()
	e.t.Reset()
	e.m = make([]interface{}, 0)
}

func (d *Decoder) reset() {
	d.indexToValue = make([]interface{}, 0)
	d.m = make([]interface{}, 0)
	d.curPos = 0
}


func Hash(v interface{}) reflect.Type{
     return reflect.ValueOf(v).Type()
}

func Sizeof(v interface{}) uint64 {
     return uint64(unsafe.Sizeof(v))
}

// type Node struct {
// 	ptr unsafe.Pointer
// 	eleType string
// 	eleSize uint64
// 	index uint32
// }

var emptyPtrNode *llrb.Item

type TypeTree struct{
	tree *llrb.Tree
	Index uint32
	reusedItem llrb.Item
	min *llrb.Item
	max *llrb.Item
}

func lessPtr(a, b llrb.Item) bool {
	return uintptr(a.Ptr) < uintptr(b.Ptr)
}

func NewTypeTree() *TypeTree{
	t := TypeTree{}
	t.tree = llrb.New(lessPtr)
	t.Index = 0
	emptyPtrNode = &llrb.Item{0, "", 0, 0}
	return &t
}

func (t *TypeTree) Reset() {
     t.tree.Reset()
}

func (t *TypeTree) closestPtr(ptr uint64, typ string, size uint64, index uint32) (*llrb.Item, *llrb.Item){
	t.reusedItem.Ptr = ptr 
	t.reusedItem.EleType = typ
	t.reusedItem.EleSize = size
	t.reusedItem.Index = index
	minItem, maxItem := t.tree.FindAdjacentNodes(t.reusedItem)
	return minItem, maxItem
	
}

func (t *TypeTree) addToTree(elePtr uint64, eleType string, eleSize uint64, eleIndex uint32) {
	t.reusedItem.Ptr = elePtr 
	t.reusedItem.EleType = eleType
	t.reusedItem.EleSize = eleSize
	t.reusedItem.Index = eleIndex
	t.tree.InsertNoReplace(&t.reusedItem)
}

func (t *TypeTree) PointerCheck(ptr_unsafe unsafe.Pointer, typ string, size uint64) (index uint32, encoded bool, err os.Error) {
       	
	ptr := uint64(uintptr(ptr_unsafe))
	t.reusedItem.Ptr = ptr 
	t.reusedItem.EleType = typ
	t.reusedItem.EleSize = size
	t.reusedItem.Index = 0
	sameItem := t.tree.Get(t.reusedItem)
	t.min, t.max = t.closestPtr(ptr, sameItem.EleType, sameItem.EleSize, sameItem.Index) 
	switch {
	case !sameItem.Equal(emptyPtrNode)  && sameItem.EleType == typ:
		// already in the tree
		return sameItem.Index, true, nil
	case !sameItem.Equal(emptyPtrNode)  && sameItem.EleType != typ:
		t.addToTree(ptr, typ, size, t.Index)
		t.Index++
		return t.Index, false, nil
	case (t.min.Equal(emptyPtrNode) && !t.max.Equal(emptyPtrNode) && (ptr + size) <= t.max.Ptr) ||
	     (!t.min.Equal(emptyPtrNode) && t.max.Equal(emptyPtrNode) && ptr >= (t.min.Ptr + t.min.EleSize)) ||
	     (!t.min.Equal(emptyPtrNode) && !t.max.Equal(emptyPtrNode) && (ptr + size) <= t.max.Ptr && ptr >= (t.min.Ptr + t.min.EleSize)) ||
	     (t.min.Equal(emptyPtrNode) && t.max.Equal(emptyPtrNode)):
		t.addToTree(ptr, typ, size, t.Index)
		t.Index++
		return t.Index, false, nil
	default:
		e := NewSayIError("Illegal pointer")		
		return 0, false, e
	}
	e := NewSayIError("Illegal pointer")		
	return 0, false, e
}

type sayIError struct {
     errMsg string
}

func NewSayIError(msg string) *sayIError {
     var e sayIError
     e.errMsg = msg
     return &e
}

func (e *sayIError) String() string {
     return e.errMsg;
}

/*type BufferedIO struct{
	writeBuf []byte
	readBuf []byte
	readBufCount int
	head int
	hasReadOnce bool
	conn net.Conn
	writeBufStart int
	writeBufLen int
}
var ReadBufSize = 1024*40

func NewBufferedIO(conn net.Conn) (bufIO *BufferedIO, error os.Error){
	//buff := make([]byte, 4096*1024)
	bufIO = &BufferedIO{}
	if conn == nil{
		fmt.Println("Connection Failed")
		return nil, os.EINVAL
	}
	bufIO.conn = conn
	bufIO.readBuf = make([]byte, ReadBufSize)
	bufIO.readBufCount = 0
	bufIO.head = 0
	bufIO.hasReadOnce = false
	bufIO.writeBuf = make([]byte, ReadBufSize)
	bufIO.writeBufStart = 0
	bufIO.writeBufLen = 0
	return bufIO, error
}

func (bo *BufferedIO) GetConn() net.Conn{
	return bo.conn
}

func (bo *BufferedIO) Write(b []byte) (n int, err os.Error){
	copy(bo.writeBuf[bo.writeBufStart:], b)
	bo.writeBufLen += len(b)
	bo.writeBufStart = bo.writeBufLen
	return bo.writeBufLen, nil
}

func (bo *BufferedIO) Flush() os.Error{

	if bo.writeBuf != nil{
		_, error := bo.conn.Write(bo.writeBuf[:bo.writeBufLen])
		bo.writeBufStart = 0
		bo.writeBufLen = 0
		return error
	}

	return nil
}

func DoMove (dest []byte, src *BufferedIO, size int) {
	copy(dest, src.readBuf[src.head:src.head+size])
	src.readBufCount -= size
	src.head += size
}

func (bo *BufferedIO) Read(b []byte) (n int, err os.Error){
	// read the smaller of the size of b, or what's left in bo
	//MyPrint("Read is called")
	copied :=  len(b)
	if bo.readBufCount < len(b) {
		copied = bo.readBufCount
	}
	remainder := len (b) - copied
	DoMove (b, bo, copied)
	//MyPrint ("READ", bo.readBuf[:copied])

	if remainder > ReadBufSize { // big read, do it directly
		n, err = bo.conn.Read (b[copied:])
		if err != nil {
			return n, err
		}
		//MyPrint("REAL READ", b[copied:])
	} else if remainder != 0 { // small read, buffer it
		n, err = io.ReadAtLeast (bo.conn, bo.readBuf, remainder)
		bo.readBufCount = n
		bo.head = 0
		//MyPrint("REAL READ", bo.readBuf)
		if err != nil {
			return n, err
		}
		DoMove (b[copied:], bo, remainder)
		//MyPrint("READ", b[copied:])
		//b = append (b, bo.readBuf[:remainder])
	}
	return len(b), err
}

func (bo *BufferedIO) Close() os.Error{
	var error os.Error
	fmt.Println("BO is=", bo)
	if bo.writeBuf != nil{
		//MyPrint("Flush buffer before close", bo, bo.conn, bo.writeBuf)
		_, error = bo.conn.Write(bo.writeBuf[:bo.writeBufLen])
		bo.writeBufStart = 0
		bo.writeBufLen = 0
	}else{
		//MyPrint("No data to flush before close")
	}
	//MyPrint("CLOSE", bo.buf[:len(bo.buf)])
	if error != nil{
		bo.conn.Close()
		return error
	}
	if bo.conn != nil{
		return bo.conn.Close()
	}
	return nil
	//return bo.conn.Close()
}*/

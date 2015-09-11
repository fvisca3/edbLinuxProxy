//_____________________________________________________________________________
// This is the package interface that wraps the RPC between netStackgo and 
// its client (Linux app). 
// Application should use the package interface instead of directly making
// RPC calls.
// Operations are async. sync interface is also provided.
// type *Conn implements io.ReaderWriter, and can be used to make encoder/decoder
//
// Dec 2012, Yaohua Li
//_____________________________________________________________________________
package nsg

import (
	"net"
	"os"
	"flag"
)

var Rootfs string

// XXX change name
type Tunnel struct {
	c net.Conn
	e *Encoder
	d *Decoder
}

type Advertisement struct {
	servicePath string
	tunnel *Tunnel
}

type Event struct {
	eventType uint64
	eventId uint64
	buffer []byte
	tunnelId uint64
	connectionId uint32
	tunnel *Tunnel
}

type Conn struct {
	tunnel *Tunnel
	tunnelId uint64
	connectionId uint32
}

const (
	EventTypeImport = 1
	EventTypeIpc = 2
	EventTypeRead = 3
	EventTypeWrite = 4
)

func Advertise (service string, serviceInstance string) (advertise *Advertisement, error os.Error) {
	var t Tunnel
	t.c, error = net.Dial ("unix", Rootfs + "/import/services/" + service + "/" + serviceInstance)
	if error != nil {
		t.c.Close ()
		return nil, error
	}

	t.e = NewEncoder (t.c)
	t.d = NewDecoder (t.c)

	servicePath := "/services/" + service + "/" + serviceInstance
	t.e.NetStackGoAdvertise (uint64(101), servicePath)
	t.d.HandleNetStackGo (t.e)

	advertise = &Advertisement{servicePath, &t}

	if resultStatus != statusOk {
		error = os.NewError ("Failure on Advertise")
	}

	return
}

func (a *Advertisement) Import () (event *Event, error os.Error) {
	a.tunnel.e.NetStackGoImport (uint64(101), a.servicePath)
	a.tunnel.d.HandleNetStackGo (a.tunnel.e)

	event = &Event{}
	event.eventType = EventTypeImport
	event.eventId = resultEventId
	event.tunnel = a.tunnel

	if resultStatus != statusOk {
		error = os.NewError ("Failure on Import")
	}

	return 
}

func Ipc (toMachine string, service string, serviceInstance string) (event *Event, err os.Error) {
	// path example: rootfs/ipc/services/ping/ping
	c, err := net.Dial ("unix", Rootfs + "/ipc/services/" + service + "/" + serviceInstance)
	if err != nil {
		c.Close()
		return nil, err
	}

	var tunnel Tunnel
	tunnel.c = c
	tunnel.e = NewEncoder (c)
	tunnel.d = NewDecoder (c)

	tunnel.e.NetStackGoIpc (uint64(101), toMachine, "/services/" + service + "/" + serviceInstance)
	tunnel.d.HandleNetStackGo (tunnel.e)

	if resultStatus != statusOk {
		return nil, os.NewError ("Failure on Ipc")
	}

	event = &Event{}
	event.eventType = EventTypeIpc
	event.eventId = resultEventId
	event.tunnel = &tunnel

	return event, nil
}

// only to be used by dsAgent
func DsIpc () (event *Event, err os.Error) {
	c, err := net.Dial ("unix", Rootfs + "/ipc/services/directoryService/directoryService")
	if err != nil {
		c.Close()
		return nil, err
	}

	var tunnel Tunnel
	tunnel.c = c
	tunnel.e = NewEncoder (c)
	tunnel.d = NewDecoder (c)

	tunnel.e.NetStackGoDsIpc (uint64(101))
	tunnel.d.HandleNetStackGo (tunnel.e)

	if resultStatus != statusOk {
		return nil, os.NewError ("Failure on DsIpc")
	}

	event = &Event{}
	event.eventType = EventTypeIpc
	event.eventId = resultEventId
	event.tunnel = &tunnel

	return event, nil
}

func (c *Conn) IpcRead () (event *Event, error os.Error) {
	c.tunnel.e.NetStackGoRead (uint64(101), c.tunnelId, c.connectionId)
	c.tunnel.d.HandleNetStackGo (c.tunnel.e)

	event = &Event{}
	event.eventType = EventTypeRead
	event.eventId = resultEventId
	event.tunnel = c.tunnel

	if resultStatus != statusOk {
		error = os.NewError ("Failure on Read")
	}

	return
}

func (c *Conn) IpcWrite (b []byte) (event *Event, error os.Error) {
	c.tunnel.e.NetStackGoWrite (uint64(101), c.tunnelId, c.connectionId, b)
	c.tunnel.d.HandleNetStackGo (c.tunnel.e)

	event = &Event{}
	event.eventType = EventTypeWrite
	event.eventId = resultEventId
	event.tunnel = c.tunnel

	if resultStatus != statusOk {
		error = os.NewError ("Failure on Write")
	}

	return
}

func (e *Event) BlockAndRetire() (data []byte, c *Conn, error os.Error) {
	e.tunnel.e.NetStackGoBlockAndRetire (uint64(101), e.eventId)
	e.tunnel.d.HandleNetStackGo (e.tunnel.e)

	data = *resultValue
	c = &Conn{tunnel:e.tunnel, tunnelId:resultTunnelId, connectionId:resultConnectionId}

	if resultStatus != statusOk {
		error = os.NewError ("Failure on BlockAndRetire")
	}

	return
}

func (c *Conn) Read (p []byte) (n int, err os.Error) {
	event, err := c.IpcRead ()
	if err != nil {
		return 0, err
	}

	data, _, err := event.BlockAndRetire ()
	if err != nil {
		return 0, err
	}
	n = copy(p, data)

	return
}

func (c *Conn) Write (p []byte) (n int, err os.Error) {
	event, err := c.IpcWrite (p)
	if err != nil {
		return 0, err
	}

	_, _, err = event.BlockAndRetire ()
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

// close imported connection will close all connections imported from that service path
func (conn *Conn) Close() (err os.Error) {
	conn.tunnel.e.NetStackGoTerminate (uint64(101))
	err = conn.tunnel.c.Close ()
	return
}

func init () {
	flag.StringVar (&Rootfs, "r", "/var/lib/ethos/ethos-netStackGo-x86_32/rootfs", "rootfs for netStackGo Daemon")
	flag.Parse ()
}

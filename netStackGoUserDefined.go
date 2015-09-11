package nsg

var Buffer []byte

var resultStatus Status
var resultEventId uint64
var resultTunnelId uint64
var resultConnectionId uint32
var resultValue *[]byte
var resultRecord *DirectoryServiceRecord

const statusOk = Status(0)

func netStackGoAdvertise (e *Encoder, cid uint64, servicePath *string) (status Status) {
	panic("BUG")
	return Status(0)
}

func netStackGoAdvertiseReply (e *Encoder, rcid uint64, status *Status) {
	resultStatus = *status
}

func netStackGoImport (e *Encoder, cid uint64, servicePath *string) (status Status, eventId uint64)  {
	panic("BUG")
	return Status(0), uint64(0)
}

func netStackGoImportReply (e *Encoder, rcid uint64, status *Status, eventId *uint64) {
	resultStatus = *status
	resultEventId = *eventId
}

func netStackGoIpc (e *Encoder, cid uint64, toMachine *string, servicePath *string) (status Status, eventId uint64) {
	panic("BUG")
	return Status(0), uint64(0)
}

func netStackGoIpcReply (e *Encoder, rcid uint64, status *Status, eventId *uint64) {
	resultStatus = *status
	resultEventId = *eventId
}

func netStackGoDsIpc (e *Encoder, cid uint64) (status Status, eventId uint64) {
	panic ("BUG")
	return Status(0), uint64(0)
}

func netStackGoDsIpcReply (e *Encoder, rcid uint64, status *Status, eventId *uint64) {
	resultStatus = *status
	resultEventId = *eventId
}

func netStackGoWrite (e *Encoder, cid uint64, tunnelId *uint64, connectionId *uint32, value *[]byte) (status Status, eventId uint64) {
	panic("BUG")
	return Status(0), uint64(0)
}

func netStackGoWriteReply (e *Encoder, rcid uint64, status *Status, eventId *uint64) {
	resultStatus = *status
	resultEventId = *eventId
}

func netStackGoRead (e *Encoder, cid uint64, tunnelId *uint64, connectionId *uint32) (status Status, eventId uint64) {
	panic("BUG")
	return Status(0), uint64(0)
}

func netStackGoReadReply (e *Encoder, rcid uint64, status *Status, eventId *uint64) {
	resultStatus = *status
	resultEventId = *eventId
}

func netStackGoBlockAndRetire (e *Encoder, cid uint64, eventId *uint64) (status Status, value []byte, tunnelId uint64, connectionId uint32) {
	panic("BUG")
	return Status(0), nil, 0, 0
}

func netStackGoBlockAndRetireReply (e *Encoder, rcid uint64, status *Status, value *[]byte, tunnelId *uint64, connectionId *uint32) {
	resultStatus = *status
	resultValue = value
	resultTunnelId = *tunnelId
	resultConnectionId = *connectionId
}

func netStackGoTerminate (e *Encoder, cid uint64) {
	panic ("BUG")
}

func netStackGoDirectoryServiceLookup (e *Encoder, cid uint64, hostname *string, eventId *uint64) (toMachine string, ip IpAddress, port UdpPort, publicKey PublicEncryptionKey, reventId uint64) {
	event, err := DsIpc ()
	if err != nil {
		// return NULL values in RPC?
		panic (err)
		return
	}

	_, c, err := event.BlockAndRetire ()
	if err != nil {
		panic (err)
		return
	}

	en := NewEncoder (c)
	de := NewDecoder (c)

	en.RpcKernelDirectoryServiceLookup (uint64(101), *hostname)
	de.HandleRpcKernel(en)

	c.Close()

	e.netStackGoDirectoryServiceLookupReply (uint64(101), *hostname, &resultRecord.IpAddress, &resultRecord.UdpPort, &resultRecord.PublicKey, *eventId, &resultStatus)
	return
}

func netStackGoDirectoryServiceLookupReply (e *Encoder, rcid uint64, toMachine *string, ip *IpAddress, port *UdpPort, publicKey *PublicEncryptionKey, reventId *uint64, status *Status) {
	panic ("BUG")
}

func rpcKernelDirectoryServiceLookup (e *Encoder, cid uint64, hostname *string) (rcid uint64, record DirectoryServiceRecord, status Status) {
	panic ("BUG")
}

func rpcKernelDirectoryServiceLookupReply (e *Encoder, rcid uint64, record *DirectoryServiceRecord, status *Status) {
	resultRecord = record
	resultStatus = *status
}

func rpcKernelRemoteConnectionCreateAuthenticateUser (e *Encoder, cid uint64, service *string, typeHash *HashValue, connectionId *uint32, userPublicKey *PublicEncryptionKey, userAuthenticator *Authenticator, writeValue *string) (rcid uint64, status Status) {
	panic ("BUG")
}

func rpcKernelRemoteConnectionCreateAuthenticateUserReply (e *Encoder, rcid uint64, status *Status) {
	panic ("BUG")
}

func rpcKernelRemoteConnectionCreate (e *Encoder, cid uint64, service *string, typeHash *HashValue, connectionId *uint32, writeValue *string) (rcid uint64, status Status) {
	panic ("BUG")
}

func rpcKernelRemoteConnectionCreateReply (e *Encoder, rcid uint64, status *Status) {
	panic ("BUG")
}

func rpcKernelIpcWrite (e *Encoder, cid uint64, contents *string) (rcid uint64, status Status) {
	panic ("BUG")
}

func rpcKernelIpcWriteReply (e *Encoder, rcid uint64, status *Status) {
	panic ("BUG")
}

func rpcKernelDirectoryServiceRegister (e *Encoder, cid uint64, hostname *string, record *DirectoryServiceRecord) (status Status) {
	panic ("BUG")
}

func rpcKernelDirectoryServiceRegisterReply (e *Encoder, status *Status) {
	panic ("BUG")
}


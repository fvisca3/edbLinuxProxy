Status uint32
IpAddress uint32
UdpPort uint16
PublicEncryptionKey [32]byte
DirectoryServiceRecord struct {
	IpAddress    IpAddress
	UdpPort      UdpPort
	PublicKey    PublicEncryptionKey
}

HashValue [64]byte
Authenticator [64]byte

NetStackGo interface {
	Advertise(cid uint64, servicePath string) (rcid uint64, status Status)
	Ipc(cid uint64, toMachine string, servicePath string) (rcid uint64, status Status, eventId uint64)
	DsIpc(cid uint64) (rcid uint64, status Status, eventId uint64)
	Import(cid uint64, servicePath string) (rcid uint64, status Status, eventId uint64)
	Write(cid uint64, tunnelId uint64, connectionId uint32, value []byte) (rcid uint64, status Status, eventId uint64)
	Read(cid uint64, tunnelId uint64, connectionId uint32) (rcid uint64, status Status, eventId uint64)
	BlockAndRetire(cid uint64, eventId uint64) (rcid uint64, status Status, value []byte, tunnelId uint64, connectionId uint32)
	Terminate(cid uint64) (rcid uint64)
	DirectoryServiceLookup(cid uint64, hostname string, eventId uint64) (rcid uint64, toMachine string, ip IpAddress, port UdpPort, publicKey PublicEncryptionKey, reventId uint64, status Status)
}

// only DirectoryServiceLookup is used 
// others are to populate RPC ID
RpcKernel interface {
	// Kernel-to-kernel RPCs.
	RemoteConnectionCreateAuthenticateUser (cid uint64, service string, typeHash HashValue, connectionId uint32, userPublicKey PublicEncryptionKey, userAuthenticator Authenticator, writeValue string) (rcid uint64, status Status)
	RemoteConnectionCreate (cid uint64, service string, typeHash HashValue, connectionId uint32, writeValue string) (rcid uint64, status Status)
	IpcWrite (cid uint64, contents string) (rcid uint64, status Status)

	// Kernel-to-directory service RPCs.
	DirectoryServiceLookup (cid uint64, hostname string) (rcid uint64, record DirectoryServiceRecord, status Status)
	DirectoryServiceRegister (cid uint64, hostname string, record DirectoryServiceRecord) (status Status)
}

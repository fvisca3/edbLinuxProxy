GRegisters struct {
   Ebx uint32
   Ecx uint32
   Edx uint32
   Esi uint32
   Edi uint32
   Ebp uint32
   Eax uint32
   Xds uint32
   Xgs uint32
   Xfs uint32
   Xes uint32
   Eip uint32
   Xcs uint32
   Eflags uint32
   Xesp uint32
   Xss uint32
   Error_code uint32
}

GdbProxyCall struct {
   Id uint8
   Addr uint32
   Size uint32
}

GdbProxyReply struct {
   Id uint8
   PacketSize uint64
   GReg GRegisters
   Memory [512]uint8
   MemorySize uint16
}

EdbRpc interface {
   Debug(cid uint64, p GdbProxyCall) (rcid uint64, r GdbProxyReply)
   Attach(cid uint64, pid uint64) (rcid uint64, r uint64)
}

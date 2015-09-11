package edbtypes

import (
   "log"
)

var ProxyReply GdbProxyReply
var AttachReply bool

func edbRpcDebug(e *Encoder, cid uint64, p *GdbProxyCall) {
   log.Fatalf("")
}

func edbRpcDebugReply(e *Encoder, cid uint64, p *GdbProxyReply) {
   log.Fatalf("")
}

func edbRpcAttach(e *Encoder, cid uint64, pid *uint64) {
   log.Fatalf("")
}

func edbRpcAttachReply(e *Encoder, cid uint64, r *uint64) {
   log.Fatalf("")
}

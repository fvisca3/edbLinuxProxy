package edbtypes
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



type GdbProxyReply struct {
     
     Id uint8
     
     PacketSize uint64
     
     GReg GRegisters
     
     Memory [512]uint8
     
     MemorySize uint16
     
}


type GdbProxyCall struct {
     
     Id uint8
     
     Addr uint32
     
     Size uint32
     
}


type GRegisters struct {
     
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

const(
	
	
	methodIdEdbRpcAttach = iota
	
	
	
	methodIdEdbRpcAttachReply
	
	
	
	methodIdEdbRpcDebug
	
	
	
	methodIdEdbRpcDebugReply
	
	
)

func (e *Encoder) EdbRpcAttach(cid uint64  ,pid uint64) (err os.Error){
     err = e.uint64(methodIdEdbRpcAttach)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.uint64(pid)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) edbRpcAttachReply(cid uint64  ,r uint64) (err os.Error){
     err = e.uint64(methodIdEdbRpcAttachReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.uint64(r)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) EdbRpcDebug(cid uint64  ,p *GdbProxyCall) (err os.Error){
     err = e.uint64(methodIdEdbRpcDebug)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.gdbProxyCall(p)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (e *Encoder) edbRpcDebugReply(cid uint64  ,r *GdbProxyReply) (err os.Error){
     err = e.uint64(methodIdEdbRpcDebugReply)
     if err != nil {
     	return err
     }
     err = e.uint64(cid)
     if err != nil {
     	return err
     }
     
     err = e.gdbProxyReply(r)
     if err != nil {
     	return err
     }
     
     err = e.Flush()
     if err != nil {
     	return err
     }
     return err
}

func (d *Decoder) HandleEdbRpc(e *Encoder) (error os.Error){
     d.ReadAll()
     methodId, err := d.uint64()
     if err != nil {
     	return err		
     }
//     methodId := *m
     switch *methodId{
     	    
	    case methodIdEdbRpcAttach:
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
		
		pid, err0 := d.uint64()
		if err0 != nil {
		   return err0
		}
		
		
//		fReturnValue0   := edbRpcAttach(*cid, pid)
		edbRpcAttach(e, *cid, pid)
//		e.edbRpcAttachReply(cid , fReturnValue0)
		
		
	    
	    case methodIdEdbRpcAttachReply:
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
		
		r, err0 := d.uint64()
		if err0 != nil {
		   return err0
		}
		
		
		edbRpcAttachReply(e, *cid, r)
		
		
	    
	    case methodIdEdbRpcDebug:
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
		
		p, err0 := d.gdbProxyCall()
		if err0 != nil {
		   return err0
		}
		
		
//		fReturnValue0   := edbRpcDebug(*cid, p)
		edbRpcDebug(e, *cid, p)
//		e.edbRpcDebugReply(cid , &fReturnValue0)
		
		
	    
	    case methodIdEdbRpcDebugReply:
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
		
		r, err0 := d.gdbProxyReply()
		if err0 != nil {
		   return err0
		}
		
		
		edbRpcDebugReply(e, *cid, r)
		
		
	    
	    default:
	    	 e := NewSayIError("Wrong MethodID")
		 return e
     }
     return nil
}

func (e *Encoder) gRegisters(v *GRegisters) (err os.Error){
     
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(v), "da58fd477490becd7b36afd89c57727eaea601798d3ff5f16623ea8e2bd886c6e0804179073a3084e656092dd9e59ddc246b1c0e78307ebbcf9aad30a62c7c81", uint64(unsafe.Sizeof(*v))); !overlap{
     	if error != nil {
	   return error
	}
     	
       err = e.uint32(v.Ebx)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Ecx)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Edx)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Esi)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Edi)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Ebp)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Eax)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Xds)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Xgs)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Xfs)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Xes)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Eip)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Xcs)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Eflags)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Xesp)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Xss)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Error_code)
       if err != nil {
	   return err
	}
        
     }   
     return nil 
}

func (e *Encoder) GRegisters (v *GRegisters) (err os.Error){
     return e.gRegisters(v)
}
func (d *Decoder) gRegisters() (v *GRegisters, error os.Error){
     
     
     var valv GRegisters
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, err0 := d.uint32()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.Ebx = *p0
     
     
     		
     p1, err1 := d.uint32()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.Ecx = *p1
     
     
     		
     p2, err2 := d.uint32()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.Edx = *p2
     
     
     		
     p3, err3 := d.uint32()
     if err3 != nil {
     	return &valv, err3
     }
     
     valv.Esi = *p3
     
     
     		
     p4, err4 := d.uint32()
     if err4 != nil {
     	return &valv, err4
     }
     
     valv.Edi = *p4
     
     
     		
     p5, err5 := d.uint32()
     if err5 != nil {
     	return &valv, err5
     }
     
     valv.Ebp = *p5
     
     
     		
     p6, err6 := d.uint32()
     if err6 != nil {
     	return &valv, err6
     }
     
     valv.Eax = *p6
     
     
     		
     p7, err7 := d.uint32()
     if err7 != nil {
     	return &valv, err7
     }
     
     valv.Xds = *p7
     
     
     		
     p8, err8 := d.uint32()
     if err8 != nil {
     	return &valv, err8
     }
     
     valv.Xgs = *p8
     
     
     		
     p9, err9 := d.uint32()
     if err9 != nil {
     	return &valv, err9
     }
     
     valv.Xfs = *p9
     
     
     		
     p10, err10 := d.uint32()
     if err10 != nil {
     	return &valv, err10
     }
     
     valv.Xes = *p10
     
     
     		
     p11, err11 := d.uint32()
     if err11 != nil {
     	return &valv, err11
     }
     
     valv.Eip = *p11
     
     
     		
     p12, err12 := d.uint32()
     if err12 != nil {
     	return &valv, err12
     }
     
     valv.Xcs = *p12
     
     
     		
     p13, err13 := d.uint32()
     if err13 != nil {
     	return &valv, err13
     }
     
     valv.Eflags = *p13
     
     
     		
     p14, err14 := d.uint32()
     if err14 != nil {
     	return &valv, err14
     }
     
     valv.Xesp = *p14
     
     
     		
     p15, err15 := d.uint32()
     if err15 != nil {
     	return &valv, err15
     }
     
     valv.Xss = *p15
     
     
     		
     p16, err16 := d.uint32()
     if err16 != nil {
     	return &valv, err16
     }
     
     valv.Error_code = *p16
     
     
     v = &valv
     return v, nil
}
func (d *Decoder) GRegisters() (v *GRegisters, error os.Error){
     return d.gRegisters()
}
func (e *Encoder) uint8_512Slice(v [512]uint8) (err os.Error){
     
     if len(v) == 0 {
     	
	
     	for _, ele := range v {
	    
       	    err = e.uint8(ele)
	    if err != nil {
	       return err
	    }
       	     
	}
	
	return nil
     }
     //vValue := reflect.ValueOf(&v).Elem()
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(&v[0]), "4660b084f55f6b2259a7572917199b634b7370dcfc30a631d4befdbc689b311ae1238537e2bc9b0fe7575910410f428a9a5bf2f4985c6b813c39a957a32494f8", uint64(len(v))); !overlap{
     	if error != nil {
	   return error
	}
     	
	
     	for _, ele := range v {
	    
       	    err = e.uint8(ele)
	    if err != nil {
	       return err
	    }
       	     
	}
	
     }    
     return nil
}

func (d *Decoder) uint8_512Slice() (v *[512]uint8, error os.Error){
     
     var valv [512]uint8
     
     length := 512
     
     
     i := 0
     for ; length > 0 ; length-- {
     	 
	 
	 p, err := d.uint8()
	 if err != nil {
	    return &valv, err
	 }
	 valv[i] = *p
     	 
	 i++ 
     }
     
     d.indexToValue = append(d.indexToValue, v)
     return &valv, nil
}

func (e *Encoder) gdbProxyReply(v *GdbProxyReply) (err os.Error){
     
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(v), "010ddf7b94f7f0db51edf6ad2320e43b5f1a3f99ae17d99b3daeed16815804aaf1b906bc4e8ee6a4ee030d2bf74ea6ec49c016807ab13812c7120b95cbfed78a", uint64(unsafe.Sizeof(*v))); !overlap{
     	if error != nil {
	   return error
	}
     	
       err = e.uint8(v.Id)
       if err != nil {
	   return err
	}
       
       err = e.uint64(v.PacketSize)
       if err != nil {
	   return err
	}
       
       err = e.gRegistersInternal(&v.GReg)
       if err != nil {
	   return err
	}
       
       err = e.uint8_512SliceInternal(v.Memory)
       if err != nil {
	   return err
	}
       
       err = e.uint16(v.MemorySize)
       if err != nil {
	   return err
	}
        
     }   
     return nil 
}

func (e *Encoder) GdbProxyReply (v *GdbProxyReply) (err os.Error){
     return e.gdbProxyReply(v)
}
func (d *Decoder) gdbProxyReply() (v *GdbProxyReply, error os.Error){
     
     
     var valv GdbProxyReply
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, err0 := d.uint8()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.Id = *p0
     
     
     		
     p1, err1 := d.uint64()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.PacketSize = *p1
     
     
     		
     p2, err2 := d.gRegistersInternal()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.GReg = *p2
     
     
     		
     p3, err3 := d.uint8_512SliceInternal()
     if err3 != nil {
     	return &valv, err3
     }
     
     valv.Memory = *p3
     
     
     		
     p4, err4 := d.uint16()
     if err4 != nil {
     	return &valv, err4
     }
     
     valv.MemorySize = *p4
     
     
     v = &valv
     return v, nil
}
func (d *Decoder) GdbProxyReply() (v *GdbProxyReply, error os.Error){
     return d.gdbProxyReply()
}
func (e *Encoder) gdbProxyCall(v *GdbProxyCall) (err os.Error){
     
     if _, overlap, error := e.t.PointerCheck(unsafe.Pointer(v), "4db8b1bc9d57bbd42d8007e1d002539df59e9eeace8211c2b1ddf33ae300bc9a6c44813b178e391242c6f543739bf2ff86d0c7a96418546879fefd0e53bbda3a", uint64(unsafe.Sizeof(*v))); !overlap{
     	if error != nil {
	   return error
	}
     	
       err = e.uint8(v.Id)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Addr)
       if err != nil {
	   return err
	}
       
       err = e.uint32(v.Size)
       if err != nil {
	   return err
	}
        
     }   
     return nil 
}

func (e *Encoder) GdbProxyCall (v *GdbProxyCall) (err os.Error){
     return e.gdbProxyCall(v)
}
func (d *Decoder) gdbProxyCall() (v *GdbProxyCall, error os.Error){
     
     
     var valv GdbProxyCall
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, err0 := d.uint8()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.Id = *p0
     
     
     		
     p1, err1 := d.uint32()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.Addr = *p1
     
     
     		
     p2, err2 := d.uint32()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.Size = *p2
     
     
     v = &valv
     return v, nil
}
func (d *Decoder) GdbProxyCall() (v *GdbProxyCall, error os.Error){
     return d.gdbProxyCall()
}
func (e *Encoder) gRegistersInternal(v *GRegisters) (err os.Error){
     
     
     	     err = e.uint32(v.Ebx)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Ecx)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Edx)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Esi)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Edi)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Ebp)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Eax)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Xds)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Xgs)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Xfs)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Xes)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Eip)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Xcs)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Eflags)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Xesp)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Xss)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Error_code)
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) gRegistersInternal() (v *GRegisters, error os.Error){
     
      
     var valv GRegisters
     
          
     p0, err0 := d.uint32()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.Ebx = *p0
     
     
          
     p1, err1 := d.uint32()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.Ecx = *p1
     
     
          
     p2, err2 := d.uint32()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.Edx = *p2
     
     
          
     p3, err3 := d.uint32()
     if err3 != nil {
     	return &valv, err3
     }
     
     valv.Esi = *p3
     
     
          
     p4, err4 := d.uint32()
     if err4 != nil {
     	return &valv, err4
     }
     
     valv.Edi = *p4
     
     
          
     p5, err5 := d.uint32()
     if err5 != nil {
     	return &valv, err5
     }
     
     valv.Ebp = *p5
     
     
          
     p6, err6 := d.uint32()
     if err6 != nil {
     	return &valv, err6
     }
     
     valv.Eax = *p6
     
     
          
     p7, err7 := d.uint32()
     if err7 != nil {
     	return &valv, err7
     }
     
     valv.Xds = *p7
     
     
          
     p8, err8 := d.uint32()
     if err8 != nil {
     	return &valv, err8
     }
     
     valv.Xgs = *p8
     
     
          
     p9, err9 := d.uint32()
     if err9 != nil {
     	return &valv, err9
     }
     
     valv.Xfs = *p9
     
     
          
     p10, err10 := d.uint32()
     if err10 != nil {
     	return &valv, err10
     }
     
     valv.Xes = *p10
     
     
          
     p11, err11 := d.uint32()
     if err11 != nil {
     	return &valv, err11
     }
     
     valv.Eip = *p11
     
     
          
     p12, err12 := d.uint32()
     if err12 != nil {
     	return &valv, err12
     }
     
     valv.Xcs = *p12
     
     
          
     p13, err13 := d.uint32()
     if err13 != nil {
     	return &valv, err13
     }
     
     valv.Eflags = *p13
     
     
          
     p14, err14 := d.uint32()
     if err14 != nil {
     	return &valv, err14
     }
     
     valv.Xesp = *p14
     
     
          
     p15, err15 := d.uint32()
     if err15 != nil {
     	return &valv, err15
     }
     
     valv.Xss = *p15
     
     
          
     p16, err16 := d.uint32()
     if err16 != nil {
     	return &valv, err16
     }
     
     valv.Error_code = *p16
     
     
     v = &valv     
     return v, nil
}

func (e *Encoder) uint8_512SliceInternal(v [512]uint8) (err os.Error){
     
     
	
     for _, ele := range v {
	 
       	 err = e.uint8(ele)
	 if err != nil {
	   return err
	 }
       	  
     }
      
     return nil
}

func (d *Decoder) uint8_512SliceInternal() (v *[512]uint8, error os.Error){
     
     var valv [512]uint8	
     
     length := 512
     
     
     i := 0
      for ; length > 0 ; length-- {
     	 
	 
	 p, err := d.uint8()
	 if err != nil {	
	    return &valv, err
	 }
	 valv[i] = *p
     	 
	 i++ 
     }
     
     return &valv, nil
}

func (e *Encoder) gdbProxyReplyInternal(v *GdbProxyReply) (err os.Error){
     
     
     	     err = e.uint8(v.Id)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint64(v.PacketSize)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.gRegistersInternal(&v.GReg)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint8_512SliceInternal(v.Memory)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint16(v.MemorySize)
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) gdbProxyReplyInternal() (v *GdbProxyReply, error os.Error){
     
      
     var valv GdbProxyReply
     
          
     p0, err0 := d.uint8()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.Id = *p0
     
     
          
     p1, err1 := d.uint64()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.PacketSize = *p1
     
     
          
     p2, err2 := d.gRegistersInternal()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.GReg = *p2
     
     
          
     p3, err3 := d.uint8_512SliceInternal()
     if err3 != nil {
     	return &valv, err3
     }
     
     valv.Memory = *p3
     
     
          
     p4, err4 := d.uint16()
     if err4 != nil {
     	return &valv, err4
     }
     
     valv.MemorySize = *p4
     
     
     v = &valv     
     return v, nil
}

func (e *Encoder) gdbProxyCallInternal(v *GdbProxyCall) (err os.Error){
     
     
     	     err = e.uint8(v.Id)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Addr)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.uint32(v.Size)
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) gdbProxyCallInternal() (v *GdbProxyCall, error os.Error){
     
      
     var valv GdbProxyCall
     
          
     p0, err0 := d.uint8()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.Id = *p0
     
     
          
     p1, err1 := d.uint32()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.Addr = *p1
     
     
          
     p2, err2 := d.uint32()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.Size = *p2
     
     
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

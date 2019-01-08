package common

import (
	"bytes"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// TKpiMessage struct for collector
// url https://elbIp:8923/{project_id}/kpi/istio

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - AgentContext
//  - TenantName
//  - Messages
type TAgentMessage struct {
	AgentContext string                        `thrift:"AgentContext,1"  json:"agentContext"`
	TenantName   string                        `thrift:"TenantName,2"  json:"tenantName"`
	Messages     map[string]map[int64][][]byte `thrift:"Messages,3" json:"messages"`
}

func NewTAgentMessage() *TAgentMessage {
	return &TAgentMessage{}
}

func (p *TAgentMessage) GetAgentContext() string {
	return p.AgentContext
}

func (p *TAgentMessage) GetTenantName() string {
	return p.TenantName
}

func (p *TAgentMessage) GetMessages() map[string]map[int64][][]byte {
	return p.Messages
}
func (p *TAgentMessage) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TAgentMessage) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AgentContext = v
	}
	return nil
}

func (p *TAgentMessage) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.TenantName = v
	}
	return nil
}

func (p *TAgentMessage) ReadField3(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]map[int64][][]byte, size)
	p.Messages = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key0 = v
		}
		_, _, size, err := iprot.ReadMapBegin()
		if err != nil {
			return thrift.PrependError("error reading map begin: ", err)
		}
		tMap := make(map[int64][][]byte, size)
		_val1 := tMap
		for i := 0; i < size; i++ {
			var _key2 int64
			if v, err := iprot.ReadI64(); err != nil {
				return thrift.PrependError("error reading field 0: ", err)
			} else {
				_key2 = v
			}
			_, size, err := iprot.ReadListBegin()
			if err != nil {
				return thrift.PrependError("error reading list begin: ", err)
			}
			tSlice := make([][]byte, 0, size)
			_val3 := tSlice
			for i := 0; i < size; i++ {
				var _elem4 []byte
				if v, err := iprot.ReadBinary(); err != nil {
					return thrift.PrependError("error reading field 0: ", err)
				} else {
					_elem4 = v
				}
				_val3 = append(_val3, _elem4)
			}
			if err := iprot.ReadListEnd(); err != nil {
				return thrift.PrependError("error reading list end: ", err)
			}
			_val1[_key2] = _val3
		}
		if err := iprot.ReadMapEnd(); err != nil {
			return thrift.PrependError("error reading map end: ", err)
		}
		p.Messages[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TAgentMessage) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TAgentMessage"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TAgentMessage) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("AgentContext", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:AgentContext: ", p), err)
	}
	if err := oprot.WriteString(string(p.AgentContext)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.AgentContext (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:AgentContext: ", p), err)
	}
	return err
}

func (p *TAgentMessage) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("TenantName", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:TenantName: ", p), err)
	}
	if err := oprot.WriteString(string(p.TenantName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.TenantName (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:TenantName: ", p), err)
	}
	return err
}

func (p *TAgentMessage) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Messages", thrift.MAP, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:Messages: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.MAP, len(p.Messages)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Messages {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.I64, thrift.LIST, len(v)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range v {
			if err := oprot.WriteI64(int64(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := oprot.WriteListBegin(thrift.STRING, len(v)); err != nil {
				return thrift.PrependError("error writing list begin: ", err)
			}
			for _, v := range v {
				if err := oprot.WriteBinary(v); err != nil {
					return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
				}
			}
			if err := oprot.WriteListEnd(); err != nil {
				return thrift.PrependError("error writing list end: ", err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:Messages: ", p), err)
	}
	return err
}

func (p *TAgentMessage) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TAgentMessage(%+v)", *p)
}

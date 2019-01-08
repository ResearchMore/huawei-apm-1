package common

import (
	"bytes"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - SourceResouceId
//  - DestResouceId
//  - TransactionType
//  - AppId
//  - SelfErrorLatency
//  - Throughput
//  - SelfLatency
//  - TotalLatency
//  - TotalLatencyList
//  - TotalErrorIndicatorList
//  - TotalErrorLatency
//  - NamespaceName
type TKpiMessage struct {
	// unused fields # 1 to 9
	SourceResouceId string `thrift:"SourceResouceId,10"  json:"sourceResouceId"`
	// unused fields # 11 to 19
	DestResouceId string `thrift:"DestResouceId,20"  json:"destResouceId"`
	// unused fields # 21 to 29
	TransactionType string `thrift:"TransactionType,30" json:"transactionType"`
	// unused fields # 31 to 39
	AppId string `thrift:"AppId,40" json:"appId"`
	// unused fields # 41 to 49
	SelfErrorLatency []byte `thrift:"SelfErrorLatency,50"  json:"selfErrorLatency"`
	// unused fields # 51 to 59
	Throughput int32 `thrift:"Throughput,60"  json:"throughput"`
	// unused fields # 61 to 69
	SelfLatency []byte `thrift:"SelfLatency,70"  json:"selfLatency"`
	// unused fields # 71 to 79
	TotalLatency []byte `thrift:"TotalLatency,80"  json:"totalLatency"`
	// unused fields # 81 to 89
	TotalLatencyList []int32 `thrift:"TotalLatencyList,90"   json:"totalLatencyList"`
	// unused fields # 91 to 99
	TotalErrorIndicatorList []bool `thrift:"TotalErrorIndicatorList,100"  json:"totalErrorIndicatorList"`
	// unused fields # 101 to 109
	TotalErrorLatency []byte `thrift:"TotalErrorLatency,110"   json:"totalErrorLatency"`
	// unused fields # 111 to 149
	NamespaceName string `thrift:"NamespaceName,150"   json:"namespaceName"`
}

func NewTKpiMessage() *TKpiMessage {
	return &TKpiMessage{}
}

func (p *TKpiMessage) GetSourceResouceId() string {
	return p.SourceResouceId
}

func (p *TKpiMessage) GetDestResouceId() string {
	return p.DestResouceId
}

func (p *TKpiMessage) GetTransactionType() string {
	return p.TransactionType
}

func (p *TKpiMessage) GetAppId() string {
	return p.AppId
}

func (p *TKpiMessage) GetSelfErrorLatency() []byte {
	return p.SelfErrorLatency
}

func (p *TKpiMessage) GetThroughput() int32 {
	return p.Throughput
}

func (p *TKpiMessage) GetSelfLatency() []byte {
	return p.SelfLatency
}

func (p *TKpiMessage) GetTotalLatency() []byte {
	return p.TotalLatency
}

func (p *TKpiMessage) GetTotalLatencyList() []int32 {
	return p.TotalLatencyList
}

func (p *TKpiMessage) GetTotalErrorIndicatorList() []bool {
	return p.TotalErrorIndicatorList
}

func (p *TKpiMessage) GetTotalErrorLatency() []byte {
	return p.TotalErrorLatency
}

func (p *TKpiMessage) GetNamespaceName() string {
	return p.NamespaceName
}
func (p *TKpiMessage) Read(iprot thrift.TProtocol) error {
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
		case 10:
			if err := p.ReadField10(iprot); err != nil {
				return err
			}
		case 20:
			if err := p.ReadField20(iprot); err != nil {
				return err
			}
		case 30:
			if err := p.ReadField30(iprot); err != nil {
				return err
			}
		case 40:
			if err := p.ReadField40(iprot); err != nil {
				return err
			}
		case 50:
			if err := p.ReadField50(iprot); err != nil {
				return err
			}
		case 60:
			if err := p.ReadField60(iprot); err != nil {
				return err
			}
		case 70:
			if err := p.ReadField70(iprot); err != nil {
				return err
			}
		case 80:
			if err := p.ReadField80(iprot); err != nil {
				return err
			}
		case 90:
			if err := p.ReadField90(iprot); err != nil {
				return err
			}
		case 100:
			if err := p.ReadField100(iprot); err != nil {
				return err
			}
		case 110:
			if err := p.ReadField110(iprot); err != nil {
				return err
			}
		case 150:
			if err := p.ReadField150(iprot); err != nil {
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

func (p *TKpiMessage) ReadField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.SourceResouceId = v
	}
	return nil
}

func (p *TKpiMessage) ReadField20(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 20: ", err)
	} else {
		p.DestResouceId = v
	}
	return nil
}

func (p *TKpiMessage) ReadField30(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 30: ", err)
	} else {
		p.TransactionType = v
	}
	return nil
}

func (p *TKpiMessage) ReadField40(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 40: ", err)
	} else {
		p.AppId = v
	}
	return nil
}

func (p *TKpiMessage) ReadField50(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 50: ", err)
	} else {
		p.SelfErrorLatency = v
	}
	return nil
}

func (p *TKpiMessage) ReadField60(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 60: ", err)
	} else {
		p.Throughput = v
	}
	return nil
}

func (p *TKpiMessage) ReadField70(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 70: ", err)
	} else {
		p.SelfLatency = v
	}
	return nil
}

func (p *TKpiMessage) ReadField80(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 80: ", err)
	} else {
		p.TotalLatency = v
	}
	return nil
}

func (p *TKpiMessage) ReadField90(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]int32, 0, size)
	p.TotalLatencyList = tSlice
	for i := 0; i < size; i++ {
		var _elem0 int32
		if v, err := iprot.ReadI32(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.TotalLatencyList = append(p.TotalLatencyList, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TKpiMessage) ReadField100(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]bool, 0, size)
	p.TotalErrorIndicatorList = tSlice
	for i := 0; i < size; i++ {
		var _elem1 bool
		if v, err := iprot.ReadBool(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem1 = v
		}
		p.TotalErrorIndicatorList = append(p.TotalErrorIndicatorList, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TKpiMessage) ReadField110(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 110: ", err)
	} else {
		p.TotalErrorLatency = v
	}
	return nil
}

func (p *TKpiMessage) ReadField150(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 150: ", err)
	} else {
		p.NamespaceName = v
	}
	return nil
}

func (p *TKpiMessage) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TKpiMessage"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField10(oprot); err != nil {
			return err
		}
		if err := p.writeField20(oprot); err != nil {
			return err
		}
		if err := p.writeField30(oprot); err != nil {
			return err
		}
		if err := p.writeField40(oprot); err != nil {
			return err
		}
		if err := p.writeField50(oprot); err != nil {
			return err
		}
		if err := p.writeField60(oprot); err != nil {
			return err
		}
		if err := p.writeField70(oprot); err != nil {
			return err
		}
		if err := p.writeField80(oprot); err != nil {
			return err
		}
		if err := p.writeField90(oprot); err != nil {
			return err
		}
		if err := p.writeField100(oprot); err != nil {
			return err
		}
		if err := p.writeField110(oprot); err != nil {
			return err
		}
		if err := p.writeField150(oprot); err != nil {
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

func (p *TKpiMessage) writeField10(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("SourceResouceId", thrift.STRING, 10); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:SourceResouceId: ", p), err)
	}
	if err := oprot.WriteString(string(p.SourceResouceId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.SourceResouceId (10) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 10:SourceResouceId: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField20(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("DestResouceId", thrift.STRING, 20); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 20:DestResouceId: ", p), err)
	}
	if err := oprot.WriteString(string(p.DestResouceId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.DestResouceId (20) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 20:DestResouceId: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField30(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("TransactionType", thrift.STRING, 30); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 30:TransactionType: ", p), err)
	}
	if err := oprot.WriteString(string(p.TransactionType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.TransactionType (30) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 30:TransactionType: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField40(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("AppId", thrift.STRING, 40); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 40:AppId: ", p), err)
	}
	if err := oprot.WriteString(string(p.AppId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.AppId (40) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 40:AppId: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField50(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("SelfErrorLatency", thrift.STRING, 50); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 50:SelfErrorLatency: ", p), err)
	}
	if err := oprot.WriteBinary(p.SelfErrorLatency); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.SelfErrorLatency (50) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 50:SelfErrorLatency: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField60(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Throughput", thrift.I32, 60); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 60:Throughput: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Throughput)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Throughput (60) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 60:Throughput: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField70(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("SelfLatency", thrift.STRING, 70); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 70:SelfLatency: ", p), err)
	}
	if err := oprot.WriteBinary(p.SelfLatency); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.SelfLatency (70) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 70:SelfLatency: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField80(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("TotalLatency", thrift.STRING, 80); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 80:TotalLatency: ", p), err)
	}
	if err := oprot.WriteBinary(p.TotalLatency); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.TotalLatency (80) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 80:TotalLatency: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField90(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("TotalLatencyList", thrift.LIST, 90); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 90:TotalLatencyList: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.I32, len(p.TotalLatencyList)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.TotalLatencyList {
		if err := oprot.WriteI32(int32(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 90:TotalLatencyList: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField100(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("TotalErrorIndicatorList", thrift.LIST, 100); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 100:TotalErrorIndicatorList: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.BOOL, len(p.TotalErrorIndicatorList)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.TotalErrorIndicatorList {
		if err := oprot.WriteBool(bool(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 100:TotalErrorIndicatorList: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField110(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("TotalErrorLatency", thrift.STRING, 110); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 110:TotalErrorLatency: ", p), err)
	}
	if err := oprot.WriteBinary(p.TotalErrorLatency); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.TotalErrorLatency (110) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 110:TotalErrorLatency: ", p), err)
	}
	return err
}

func (p *TKpiMessage) writeField150(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("NamespaceName", thrift.STRING, 150); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 150:NamespaceName: ", p), err)
	}
	if err := oprot.WriteString(string(p.NamespaceName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.NamespaceName (150) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 150:NamespaceName: ", p), err)
	}
	return err
}

func (p *TKpiMessage) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TKpiMessage(%+v)", *p)
}

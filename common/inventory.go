package common

import (
	"bytes"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// url https://elbIp:8923/{project_id}/inventory/istio

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - Hostname
//  - IP
//  - AgentId
//  - AppName
//  - ClusterKey
//  - ServiceType
//  - DisplayName
//  - InstanceName
//  - ContainerId
//  - Pid
//  - ProjectId
//  - CollectorId
//  - AppId
//  - Props
//  - Ports
//  - Ips
//  - Tier
//  - NamespaceName
//  - Created
//  - Updated
//  - Deleted
type TDiscoveryInfo struct {
	Hostname     string `thrift:"Hostname,1"  json:"hostname"`
	IP           string `thrift:"IP,2" json:"ip"`
	AgentId      string `thrift:"AgentId,3"   json:"agentId"`
	AppName      string `thrift:"AppName,4"   json:"appName"`
	ClusterKey   string `thrift:"ClusterKey,5"  json:"clusterKey"`
	ServiceType  string `thrift:"ServiceType,6"   json:"serviceType"`
	DisplayName  string `thrift:"DisplayName,7"   json:"displayName"`
	InstanceName string `thrift:"InstanceName,8"  json:"instanceName"`
	ContainerId  string `thrift:"ContainerId,9"  json:"containerId"`
	Pid          int32  `thrift:"Pid,10"   json:"pid"`
	// unused field # 11
	ProjectId string `thrift:"ProjectId,12" json:"projectId"`
	// unused field # 13
	CollectorId string `thrift:"CollectorId,14" json:"collectorId"`
	AppId       string `thrift:"AppId,15" json:"appId"`
	// unused fields # 16 to 19
	Props map[string]string `thrift:"Props,20" json:"props"`
	// unused fields # 21 to 29
	Ports []int32 `thrift:"Ports,30" json:"ports"`
	// unused fields # 31 to 39
	Ips []string `thrift:"Ips,40" json:"ips"`
	// unused fields # 41 to 49
	Tier string `thrift:"Tier,50"  json:"tier"`
	// unused fields # 51 to 59
	NamespaceName string `thrift:"NamespaceName,60"  json:"namespaceName"`
	Created       int64  `thrift:"Created,61"  json:"created"`
	Updated       int64  `thrift:"Updated,62"  json:"updated"`
	Deleted       int64  `thrift:"Deleted,63" json:"deleted"`
}

func NewTDiscoveryInfo() *TDiscoveryInfo {
	return &TDiscoveryInfo{}
}

func (p *TDiscoveryInfo) GetHostname() string {
	return p.Hostname
}

func (p *TDiscoveryInfo) GetIP() string {
	return p.IP
}

func (p *TDiscoveryInfo) GetAgentId() string {
	return p.AgentId
}

func (p *TDiscoveryInfo) GetAppName() string {
	return p.AppName
}

func (p *TDiscoveryInfo) GetClusterKey() string {
	return p.ClusterKey
}

func (p *TDiscoveryInfo) GetServiceType() string {
	return p.ServiceType
}

func (p *TDiscoveryInfo) GetDisplayName() string {
	return p.DisplayName
}

func (p *TDiscoveryInfo) GetInstanceName() string {
	return p.InstanceName
}

func (p *TDiscoveryInfo) GetContainerId() string {
	return p.ContainerId
}

func (p *TDiscoveryInfo) GetPid() int32 {
	return p.Pid
}

func (p *TDiscoveryInfo) GetProjectId() string {
	return p.ProjectId
}

func (p *TDiscoveryInfo) GetCollectorId() string {
	return p.CollectorId
}

func (p *TDiscoveryInfo) GetAppId() string {
	return p.AppId
}

func (p *TDiscoveryInfo) GetProps() map[string]string {
	return p.Props
}

func (p *TDiscoveryInfo) GetPorts() []int32 {
	return p.Ports
}

func (p *TDiscoveryInfo) GetIps() []string {
	return p.Ips
}

func (p *TDiscoveryInfo) GetTier() string {
	return p.Tier
}

func (p *TDiscoveryInfo) GetNamespaceName() string {
	return p.NamespaceName
}

func (p *TDiscoveryInfo) GetCreated() int64 {
	return p.Created
}

func (p *TDiscoveryInfo) GetUpdated() int64 {
	return p.Updated
}

func (p *TDiscoveryInfo) GetDeleted() int64 {
	return p.Deleted
}
func (p *TDiscoveryInfo) Read(iprot thrift.TProtocol) error {
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
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.ReadField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.ReadField8(iprot); err != nil {
				return err
			}
		case 9:
			if err := p.ReadField9(iprot); err != nil {
				return err
			}
		case 10:
			if err := p.ReadField10(iprot); err != nil {
				return err
			}
		case 12:
			if err := p.ReadField12(iprot); err != nil {
				return err
			}
		case 14:
			if err := p.ReadField14(iprot); err != nil {
				return err
			}
		case 15:
			if err := p.ReadField15(iprot); err != nil {
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
		case 61:
			if err := p.ReadField61(iprot); err != nil {
				return err
			}
		case 62:
			if err := p.ReadField62(iprot); err != nil {
				return err
			}
		case 63:
			if err := p.ReadField63(iprot); err != nil {
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

func (p *TDiscoveryInfo) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Hostname = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.IP = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.AgentId = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.AppName = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.ClusterKey = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.ServiceType = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.DisplayName = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.InstanceName = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.ContainerId = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.Pid = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField12(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 12: ", err)
	} else {
		p.ProjectId = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField14(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 14: ", err)
	} else {
		p.CollectorId = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField15(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 15: ", err)
	} else {
		p.AppId = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField20(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Props = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key0 = v
		}
		var _val1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val1 = v
		}
		p.Props[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField30(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]int32, 0, size)
	p.Ports = tSlice
	for i := 0; i < size; i++ {
		var _elem2 int32
		if v, err := iprot.ReadI32(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem2 = v
		}
		p.Ports = append(p.Ports, _elem2)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField40(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Ips = tSlice
	for i := 0; i < size; i++ {
		var _elem3 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem3 = v
		}
		p.Ips = append(p.Ips, _elem3)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField50(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 50: ", err)
	} else {
		p.Tier = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField60(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 60: ", err)
	} else {
		p.NamespaceName = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField61(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 61: ", err)
	} else {
		p.Created = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField62(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 62: ", err)
	} else {
		p.Updated = v
	}
	return nil
}

func (p *TDiscoveryInfo) ReadField63(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 63: ", err)
	} else {
		p.Deleted = v
	}
	return nil
}

func (p *TDiscoveryInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TDiscoveryInfo"); err != nil {
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
		if err := p.writeField4(oprot); err != nil {
			return err
		}
		if err := p.writeField5(oprot); err != nil {
			return err
		}
		if err := p.writeField6(oprot); err != nil {
			return err
		}
		if err := p.writeField7(oprot); err != nil {
			return err
		}
		if err := p.writeField8(oprot); err != nil {
			return err
		}
		if err := p.writeField9(oprot); err != nil {
			return err
		}
		if err := p.writeField10(oprot); err != nil {
			return err
		}
		if err := p.writeField12(oprot); err != nil {
			return err
		}
		if err := p.writeField14(oprot); err != nil {
			return err
		}
		if err := p.writeField15(oprot); err != nil {
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
		if err := p.writeField61(oprot); err != nil {
			return err
		}
		if err := p.writeField62(oprot); err != nil {
			return err
		}
		if err := p.writeField63(oprot); err != nil {
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

func (p *TDiscoveryInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Hostname", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Hostname: ", p), err)
	}
	if err := oprot.WriteString(string(p.Hostname)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Hostname (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Hostname: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("IP", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:IP: ", p), err)
	}
	if err := oprot.WriteString(string(p.IP)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.IP (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:IP: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("AgentId", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:AgentId: ", p), err)
	}
	if err := oprot.WriteString(string(p.AgentId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.AgentId (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:AgentId: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("AppName", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:AppName: ", p), err)
	}
	if err := oprot.WriteString(string(p.AppName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.AppName (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:AppName: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ClusterKey", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:ClusterKey: ", p), err)
	}
	if err := oprot.WriteString(string(p.ClusterKey)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ClusterKey (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:ClusterKey: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ServiceType", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:ServiceType: ", p), err)
	}
	if err := oprot.WriteString(string(p.ServiceType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ServiceType (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:ServiceType: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("DisplayName", thrift.STRING, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:DisplayName: ", p), err)
	}
	if err := oprot.WriteString(string(p.DisplayName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.DisplayName (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:DisplayName: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("InstanceName", thrift.STRING, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:InstanceName: ", p), err)
	}
	if err := oprot.WriteString(string(p.InstanceName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.InstanceName (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:InstanceName: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField9(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ContainerId", thrift.STRING, 9); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:ContainerId: ", p), err)
	}
	if err := oprot.WriteString(string(p.ContainerId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ContainerId (9) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 9:ContainerId: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField10(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Pid", thrift.I32, 10); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:Pid: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Pid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Pid (10) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 10:Pid: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField12(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ProjectId", thrift.STRING, 12); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 12:ProjectId: ", p), err)
	}
	if err := oprot.WriteString(string(p.ProjectId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ProjectId (12) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 12:ProjectId: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField14(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("CollectorId", thrift.STRING, 14); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 14:CollectorId: ", p), err)
	}
	if err := oprot.WriteString(string(p.CollectorId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.CollectorId (14) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 14:CollectorId: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField15(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("AppId", thrift.STRING, 15); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 15:AppId: ", p), err)
	}
	if err := oprot.WriteString(string(p.AppId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.AppId (15) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 15:AppId: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField20(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Props", thrift.MAP, 20); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 20:Props: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Props)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Props {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 20:Props: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField30(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Ports", thrift.LIST, 30); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 30:Ports: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.I32, len(p.Ports)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Ports {
		if err := oprot.WriteI32(int32(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 30:Ports: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField40(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Ips", thrift.LIST, 40); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 40:Ips: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.Ips)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Ips {
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 40:Ips: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField50(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Tier", thrift.STRING, 50); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 50:Tier: ", p), err)
	}
	if err := oprot.WriteString(string(p.Tier)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Tier (50) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 50:Tier: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField60(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("NamespaceName", thrift.STRING, 60); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 60:NamespaceName: ", p), err)
	}
	if err := oprot.WriteString(string(p.NamespaceName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.NamespaceName (60) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 60:NamespaceName: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField61(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Created", thrift.I64, 61); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 61:Created: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Created)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Created (61) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 61:Created: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField62(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Updated", thrift.I64, 62); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 62:Updated: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Updated)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Updated (62) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 62:Updated: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) writeField63(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Deleted", thrift.I64, 63); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 63:Deleted: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Deleted)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Deleted (63) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 63:Deleted: ", p), err)
	}
	return err
}

func (p *TDiscoveryInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDiscoveryInfo(%+v)", *p)
}

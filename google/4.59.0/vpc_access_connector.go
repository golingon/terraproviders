// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	vpcaccessconnector "github.com/golingon/terraproviders/google/4.59.0/vpcaccessconnector"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewVpcAccessConnector(name string, args VpcAccessConnectorArgs) *VpcAccessConnector {
	return &VpcAccessConnector{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcAccessConnector)(nil)

type VpcAccessConnector struct {
	Name  string
	Args  VpcAccessConnectorArgs
	state *vpcAccessConnectorState
}

func (vac *VpcAccessConnector) Type() string {
	return "google_vpc_access_connector"
}

func (vac *VpcAccessConnector) LocalName() string {
	return vac.Name
}

func (vac *VpcAccessConnector) Configuration() interface{} {
	return vac.Args
}

func (vac *VpcAccessConnector) Attributes() vpcAccessConnectorAttributes {
	return vpcAccessConnectorAttributes{ref: terra.ReferenceResource(vac)}
}

func (vac *VpcAccessConnector) ImportState(av io.Reader) error {
	vac.state = &vpcAccessConnectorState{}
	if err := json.NewDecoder(av).Decode(vac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vac.Type(), vac.LocalName(), err)
	}
	return nil
}

func (vac *VpcAccessConnector) State() (*vpcAccessConnectorState, bool) {
	return vac.state, vac.state != nil
}

func (vac *VpcAccessConnector) StateMust() *vpcAccessConnectorState {
	if vac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vac.Type(), vac.LocalName()))
	}
	return vac.state
}

func (vac *VpcAccessConnector) DependOn() terra.Reference {
	return terra.ReferenceResource(vac)
}

type VpcAccessConnectorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpCidrRange: string, optional
	IpCidrRange terra.StringValue `hcl:"ip_cidr_range,attr"`
	// MachineType: string, optional
	MachineType terra.StringValue `hcl:"machine_type,attr"`
	// MaxInstances: number, optional
	MaxInstances terra.NumberValue `hcl:"max_instances,attr"`
	// MaxThroughput: number, optional
	MaxThroughput terra.NumberValue `hcl:"max_throughput,attr"`
	// MinInstances: number, optional
	MinInstances terra.NumberValue `hcl:"min_instances,attr"`
	// MinThroughput: number, optional
	MinThroughput terra.NumberValue `hcl:"min_throughput,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Subnet: optional
	Subnet *vpcaccessconnector.Subnet `hcl:"subnet,block"`
	// Timeouts: optional
	Timeouts *vpcaccessconnector.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that VpcAccessConnector depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type vpcAccessConnectorAttributes struct {
	ref terra.Reference
}

func (vac vpcAccessConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("id"))
}

func (vac vpcAccessConnectorAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("ip_cidr_range"))
}

func (vac vpcAccessConnectorAttributes) MachineType() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("machine_type"))
}

func (vac vpcAccessConnectorAttributes) MaxInstances() terra.NumberValue {
	return terra.ReferenceNumber(vac.ref.Append("max_instances"))
}

func (vac vpcAccessConnectorAttributes) MaxThroughput() terra.NumberValue {
	return terra.ReferenceNumber(vac.ref.Append("max_throughput"))
}

func (vac vpcAccessConnectorAttributes) MinInstances() terra.NumberValue {
	return terra.ReferenceNumber(vac.ref.Append("min_instances"))
}

func (vac vpcAccessConnectorAttributes) MinThroughput() terra.NumberValue {
	return terra.ReferenceNumber(vac.ref.Append("min_throughput"))
}

func (vac vpcAccessConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("name"))
}

func (vac vpcAccessConnectorAttributes) Network() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("network"))
}

func (vac vpcAccessConnectorAttributes) Project() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("project"))
}

func (vac vpcAccessConnectorAttributes) Region() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("region"))
}

func (vac vpcAccessConnectorAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("self_link"))
}

func (vac vpcAccessConnectorAttributes) State() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("state"))
}

func (vac vpcAccessConnectorAttributes) Subnet() terra.ListValue[vpcaccessconnector.SubnetAttributes] {
	return terra.ReferenceList[vpcaccessconnector.SubnetAttributes](vac.ref.Append("subnet"))
}

func (vac vpcAccessConnectorAttributes) Timeouts() vpcaccessconnector.TimeoutsAttributes {
	return terra.ReferenceSingle[vpcaccessconnector.TimeoutsAttributes](vac.ref.Append("timeouts"))
}

type vpcAccessConnectorState struct {
	Id            string                            `json:"id"`
	IpCidrRange   string                            `json:"ip_cidr_range"`
	MachineType   string                            `json:"machine_type"`
	MaxInstances  float64                           `json:"max_instances"`
	MaxThroughput float64                           `json:"max_throughput"`
	MinInstances  float64                           `json:"min_instances"`
	MinThroughput float64                           `json:"min_throughput"`
	Name          string                            `json:"name"`
	Network       string                            `json:"network"`
	Project       string                            `json:"project"`
	Region        string                            `json:"region"`
	SelfLink      string                            `json:"self_link"`
	State         string                            `json:"state"`
	Subnet        []vpcaccessconnector.SubnetState  `json:"subnet"`
	Timeouts      *vpcaccessconnector.TimeoutsState `json:"timeouts"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vpcaccessconnector "github.com/golingon/terraproviders/googlebeta/4.75.1/vpcaccessconnector"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcAccessConnector creates a new instance of [VpcAccessConnector].
func NewVpcAccessConnector(name string, args VpcAccessConnectorArgs) *VpcAccessConnector {
	return &VpcAccessConnector{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcAccessConnector)(nil)

// VpcAccessConnector represents the Terraform resource google_vpc_access_connector.
type VpcAccessConnector struct {
	Name      string
	Args      VpcAccessConnectorArgs
	state     *vpcAccessConnectorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcAccessConnector].
func (vac *VpcAccessConnector) Type() string {
	return "google_vpc_access_connector"
}

// LocalName returns the local name for [VpcAccessConnector].
func (vac *VpcAccessConnector) LocalName() string {
	return vac.Name
}

// Configuration returns the configuration (args) for [VpcAccessConnector].
func (vac *VpcAccessConnector) Configuration() interface{} {
	return vac.Args
}

// DependOn is used for other resources to depend on [VpcAccessConnector].
func (vac *VpcAccessConnector) DependOn() terra.Reference {
	return terra.ReferenceResource(vac)
}

// Dependencies returns the list of resources [VpcAccessConnector] depends_on.
func (vac *VpcAccessConnector) Dependencies() terra.Dependencies {
	return vac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcAccessConnector].
func (vac *VpcAccessConnector) LifecycleManagement() *terra.Lifecycle {
	return vac.Lifecycle
}

// Attributes returns the attributes for [VpcAccessConnector].
func (vac *VpcAccessConnector) Attributes() vpcAccessConnectorAttributes {
	return vpcAccessConnectorAttributes{ref: terra.ReferenceResource(vac)}
}

// ImportState imports the given attribute values into [VpcAccessConnector]'s state.
func (vac *VpcAccessConnector) ImportState(av io.Reader) error {
	vac.state = &vpcAccessConnectorState{}
	if err := json.NewDecoder(av).Decode(vac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vac.Type(), vac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcAccessConnector] has state.
func (vac *VpcAccessConnector) State() (*vpcAccessConnectorState, bool) {
	return vac.state, vac.state != nil
}

// StateMust returns the state for [VpcAccessConnector]. Panics if the state is nil.
func (vac *VpcAccessConnector) StateMust() *vpcAccessConnectorState {
	if vac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vac.Type(), vac.LocalName()))
	}
	return vac.state
}

// VpcAccessConnectorArgs contains the configurations for google_vpc_access_connector.
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
}
type vpcAccessConnectorAttributes struct {
	ref terra.Reference
}

// ConnectedProjects returns a reference to field connected_projects of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) ConnectedProjects() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vac.ref.Append("connected_projects"))
}

// Id returns a reference to field id of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("id"))
}

// IpCidrRange returns a reference to field ip_cidr_range of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("ip_cidr_range"))
}

// MachineType returns a reference to field machine_type of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("machine_type"))
}

// MaxInstances returns a reference to field max_instances of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) MaxInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(vac.ref.Append("max_instances"))
}

// MaxThroughput returns a reference to field max_throughput of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) MaxThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(vac.ref.Append("max_throughput"))
}

// MinInstances returns a reference to field min_instances of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) MinInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(vac.ref.Append("min_instances"))
}

// MinThroughput returns a reference to field min_throughput of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) MinThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(vac.ref.Append("min_throughput"))
}

// Name returns a reference to field name of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("name"))
}

// Network returns a reference to field network of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("network"))
}

// Project returns a reference to field project of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("project"))
}

// Region returns a reference to field region of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("self_link"))
}

// State returns a reference to field state of google_vpc_access_connector.
func (vac vpcAccessConnectorAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("state"))
}

func (vac vpcAccessConnectorAttributes) Subnet() terra.ListValue[vpcaccessconnector.SubnetAttributes] {
	return terra.ReferenceAsList[vpcaccessconnector.SubnetAttributes](vac.ref.Append("subnet"))
}

func (vac vpcAccessConnectorAttributes) Timeouts() vpcaccessconnector.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcaccessconnector.TimeoutsAttributes](vac.ref.Append("timeouts"))
}

type vpcAccessConnectorState struct {
	ConnectedProjects []string                          `json:"connected_projects"`
	Id                string                            `json:"id"`
	IpCidrRange       string                            `json:"ip_cidr_range"`
	MachineType       string                            `json:"machine_type"`
	MaxInstances      float64                           `json:"max_instances"`
	MaxThroughput     float64                           `json:"max_throughput"`
	MinInstances      float64                           `json:"min_instances"`
	MinThroughput     float64                           `json:"min_throughput"`
	Name              string                            `json:"name"`
	Network           string                            `json:"network"`
	Project           string                            `json:"project"`
	Region            string                            `json:"region"`
	SelfLink          string                            `json:"self_link"`
	State             string                            `json:"state"`
	Subnet            []vpcaccessconnector.SubnetState  `json:"subnet"`
	Timeouts          *vpcaccessconnector.TimeoutsState `json:"timeouts"`
}

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_vpc_access_connector

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_vpc_access_connector.
type Resource struct {
	Name      string
	Args      Args
	state     *googleVpcAccessConnectorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gvac *Resource) Type() string {
	return "google_vpc_access_connector"
}

// LocalName returns the local name for [Resource].
func (gvac *Resource) LocalName() string {
	return gvac.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gvac *Resource) Configuration() interface{} {
	return gvac.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gvac *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gvac)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gvac *Resource) Dependencies() terra.Dependencies {
	return gvac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gvac *Resource) LifecycleManagement() *terra.Lifecycle {
	return gvac.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gvac *Resource) Attributes() googleVpcAccessConnectorAttributes {
	return googleVpcAccessConnectorAttributes{ref: terra.ReferenceResource(gvac)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gvac *Resource) ImportState(state io.Reader) error {
	gvac.state = &googleVpcAccessConnectorState{}
	if err := json.NewDecoder(state).Decode(gvac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gvac.Type(), gvac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gvac *Resource) State() (*googleVpcAccessConnectorState, bool) {
	return gvac.state, gvac.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gvac *Resource) StateMust() *googleVpcAccessConnectorState {
	if gvac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gvac.Type(), gvac.LocalName()))
	}
	return gvac.state
}

// Args contains the configurations for google_vpc_access_connector.
type Args struct {
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
	Subnet *Subnet `hcl:"subnet,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleVpcAccessConnectorAttributes struct {
	ref terra.Reference
}

// ConnectedProjects returns a reference to field connected_projects of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) ConnectedProjects() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gvac.ref.Append("connected_projects"))
}

// Id returns a reference to field id of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gvac.ref.Append("id"))
}

// IpCidrRange returns a reference to field ip_cidr_range of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(gvac.ref.Append("ip_cidr_range"))
}

// MachineType returns a reference to field machine_type of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(gvac.ref.Append("machine_type"))
}

// MaxInstances returns a reference to field max_instances of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) MaxInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(gvac.ref.Append("max_instances"))
}

// MaxThroughput returns a reference to field max_throughput of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) MaxThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(gvac.ref.Append("max_throughput"))
}

// MinInstances returns a reference to field min_instances of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) MinInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(gvac.ref.Append("min_instances"))
}

// MinThroughput returns a reference to field min_throughput of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) MinThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(gvac.ref.Append("min_throughput"))
}

// Name returns a reference to field name of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gvac.ref.Append("name"))
}

// Network returns a reference to field network of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(gvac.ref.Append("network"))
}

// Project returns a reference to field project of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gvac.ref.Append("project"))
}

// Region returns a reference to field region of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gvac.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gvac.ref.Append("self_link"))
}

// State returns a reference to field state of google_vpc_access_connector.
func (gvac googleVpcAccessConnectorAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gvac.ref.Append("state"))
}

func (gvac googleVpcAccessConnectorAttributes) Subnet() terra.ListValue[SubnetAttributes] {
	return terra.ReferenceAsList[SubnetAttributes](gvac.ref.Append("subnet"))
}

func (gvac googleVpcAccessConnectorAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gvac.ref.Append("timeouts"))
}

type googleVpcAccessConnectorState struct {
	ConnectedProjects []string       `json:"connected_projects"`
	Id                string         `json:"id"`
	IpCidrRange       string         `json:"ip_cidr_range"`
	MachineType       string         `json:"machine_type"`
	MaxInstances      float64        `json:"max_instances"`
	MaxThroughput     float64        `json:"max_throughput"`
	MinInstances      float64        `json:"min_instances"`
	MinThroughput     float64        `json:"min_throughput"`
	Name              string         `json:"name"`
	Network           string         `json:"network"`
	Project           string         `json:"project"`
	Region            string         `json:"region"`
	SelfLink          string         `json:"self_link"`
	State             string         `json:"state"`
	Subnet            []SubnetState  `json:"subnet"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	edgenetworksubnet "github.com/golingon/terraproviders/googlebeta/5.24.0/edgenetworksubnet"
	"io"
)

// NewEdgenetworkSubnet creates a new instance of [EdgenetworkSubnet].
func NewEdgenetworkSubnet(name string, args EdgenetworkSubnetArgs) *EdgenetworkSubnet {
	return &EdgenetworkSubnet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EdgenetworkSubnet)(nil)

// EdgenetworkSubnet represents the Terraform resource google_edgenetwork_subnet.
type EdgenetworkSubnet struct {
	Name      string
	Args      EdgenetworkSubnetArgs
	state     *edgenetworkSubnetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EdgenetworkSubnet].
func (es *EdgenetworkSubnet) Type() string {
	return "google_edgenetwork_subnet"
}

// LocalName returns the local name for [EdgenetworkSubnet].
func (es *EdgenetworkSubnet) LocalName() string {
	return es.Name
}

// Configuration returns the configuration (args) for [EdgenetworkSubnet].
func (es *EdgenetworkSubnet) Configuration() interface{} {
	return es.Args
}

// DependOn is used for other resources to depend on [EdgenetworkSubnet].
func (es *EdgenetworkSubnet) DependOn() terra.Reference {
	return terra.ReferenceResource(es)
}

// Dependencies returns the list of resources [EdgenetworkSubnet] depends_on.
func (es *EdgenetworkSubnet) Dependencies() terra.Dependencies {
	return es.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EdgenetworkSubnet].
func (es *EdgenetworkSubnet) LifecycleManagement() *terra.Lifecycle {
	return es.Lifecycle
}

// Attributes returns the attributes for [EdgenetworkSubnet].
func (es *EdgenetworkSubnet) Attributes() edgenetworkSubnetAttributes {
	return edgenetworkSubnetAttributes{ref: terra.ReferenceResource(es)}
}

// ImportState imports the given attribute values into [EdgenetworkSubnet]'s state.
func (es *EdgenetworkSubnet) ImportState(av io.Reader) error {
	es.state = &edgenetworkSubnetState{}
	if err := json.NewDecoder(av).Decode(es.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", es.Type(), es.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EdgenetworkSubnet] has state.
func (es *EdgenetworkSubnet) State() (*edgenetworkSubnetState, bool) {
	return es.state, es.state != nil
}

// StateMust returns the state for [EdgenetworkSubnet]. Panics if the state is nil.
func (es *EdgenetworkSubnet) StateMust() *edgenetworkSubnetState {
	if es.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", es.Type(), es.LocalName()))
	}
	return es.state
}

// EdgenetworkSubnetArgs contains the configurations for google_edgenetwork_subnet.
type EdgenetworkSubnetArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Ipv4Cidr: list of string, optional
	Ipv4Cidr terra.ListValue[terra.StringValue] `hcl:"ipv4_cidr,attr"`
	// Ipv6Cidr: list of string, optional
	Ipv6Cidr terra.ListValue[terra.StringValue] `hcl:"ipv6_cidr,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// VlanId: number, optional
	VlanId terra.NumberValue `hcl:"vlan_id,attr"`
	// Zone: string, required
	Zone terra.StringValue `hcl:"zone,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *edgenetworksubnet.Timeouts `hcl:"timeouts,block"`
}
type edgenetworkSubnetAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("create_time"))
}

// Description returns a reference to field description of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("description"))
}

// Id returns a reference to field id of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("id"))
}

// Ipv4Cidr returns a reference to field ipv4_cidr of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) Ipv4Cidr() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](es.ref.Append("ipv4_cidr"))
}

// Ipv6Cidr returns a reference to field ipv6_cidr of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) Ipv6Cidr() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](es.ref.Append("ipv6_cidr"))
}

// Labels returns a reference to field labels of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](es.ref.Append("labels"))
}

// Location returns a reference to field location of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("location"))
}

// Name returns a reference to field name of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("name"))
}

// Network returns a reference to field network of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("network"))
}

// Project returns a reference to field project of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("project"))
}

// State returns a reference to field state of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("state"))
}

// SubnetId returns a reference to field subnet_id of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("subnet_id"))
}

// UpdateTime returns a reference to field update_time of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("update_time"))
}

// VlanId returns a reference to field vlan_id of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) VlanId() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("vlan_id"))
}

// Zone returns a reference to field zone of google_edgenetwork_subnet.
func (es edgenetworkSubnetAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("zone"))
}

func (es edgenetworkSubnetAttributes) Timeouts() edgenetworksubnet.TimeoutsAttributes {
	return terra.ReferenceAsSingle[edgenetworksubnet.TimeoutsAttributes](es.ref.Append("timeouts"))
}

type edgenetworkSubnetState struct {
	CreateTime  string                           `json:"create_time"`
	Description string                           `json:"description"`
	Id          string                           `json:"id"`
	Ipv4Cidr    []string                         `json:"ipv4_cidr"`
	Ipv6Cidr    []string                         `json:"ipv6_cidr"`
	Labels      map[string]string                `json:"labels"`
	Location    string                           `json:"location"`
	Name        string                           `json:"name"`
	Network     string                           `json:"network"`
	Project     string                           `json:"project"`
	State       string                           `json:"state"`
	SubnetId    string                           `json:"subnet_id"`
	UpdateTime  string                           `json:"update_time"`
	VlanId      float64                          `json:"vlan_id"`
	Zone        string                           `json:"zone"`
	Timeouts    *edgenetworksubnet.TimeoutsState `json:"timeouts"`
}

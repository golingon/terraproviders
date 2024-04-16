// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_route

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

// Resource represents the Terraform resource azurestack_route.
type Resource struct {
	Name      string
	Args      Args
	state     *azurestackRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ar *Resource) Type() string {
	return "azurestack_route"
}

// LocalName returns the local name for [Resource].
func (ar *Resource) LocalName() string {
	return ar.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ar *Resource) Configuration() interface{} {
	return ar.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ar *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ar)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ar *Resource) Dependencies() terra.Dependencies {
	return ar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ar *Resource) LifecycleManagement() *terra.Lifecycle {
	return ar.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ar *Resource) Attributes() azurestackRouteAttributes {
	return azurestackRouteAttributes{ref: terra.ReferenceResource(ar)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ar *Resource) ImportState(state io.Reader) error {
	ar.state = &azurestackRouteState{}
	if err := json.NewDecoder(state).Decode(ar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ar.Type(), ar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ar *Resource) State() (*azurestackRouteState, bool) {
	return ar.state, ar.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ar *Resource) StateMust() *azurestackRouteState {
	if ar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ar.Type(), ar.LocalName()))
	}
	return ar.state
}

// Args contains the configurations for azurestack_route.
type Args struct {
	// AddressPrefix: string, required
	AddressPrefix terra.StringValue `hcl:"address_prefix,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NextHopInIpAddress: string, optional
	NextHopInIpAddress terra.StringValue `hcl:"next_hop_in_ip_address,attr"`
	// NextHopType: string, required
	NextHopType terra.StringValue `hcl:"next_hop_type,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RouteTableName: string, required
	RouteTableName terra.StringValue `hcl:"route_table_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurestackRouteAttributes struct {
	ref terra.Reference
}

// AddressPrefix returns a reference to field address_prefix of azurestack_route.
func (ar azurestackRouteAttributes) AddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("address_prefix"))
}

// Id returns a reference to field id of azurestack_route.
func (ar azurestackRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("id"))
}

// Name returns a reference to field name of azurestack_route.
func (ar azurestackRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("name"))
}

// NextHopInIpAddress returns a reference to field next_hop_in_ip_address of azurestack_route.
func (ar azurestackRouteAttributes) NextHopInIpAddress() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("next_hop_in_ip_address"))
}

// NextHopType returns a reference to field next_hop_type of azurestack_route.
func (ar azurestackRouteAttributes) NextHopType() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("next_hop_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_route.
func (ar azurestackRouteAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("resource_group_name"))
}

// RouteTableName returns a reference to field route_table_name of azurestack_route.
func (ar azurestackRouteAttributes) RouteTableName() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("route_table_name"))
}

func (ar azurestackRouteAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ar.ref.Append("timeouts"))
}

type azurestackRouteState struct {
	AddressPrefix      string         `json:"address_prefix"`
	Id                 string         `json:"id"`
	Name               string         `json:"name"`
	NextHopInIpAddress string         `json:"next_hop_in_ip_address"`
	NextHopType        string         `json:"next_hop_type"`
	ResourceGroupName  string         `json:"resource_group_name"`
	RouteTableName     string         `json:"route_table_name"`
	Timeouts           *TimeoutsState `json:"timeouts"`
}

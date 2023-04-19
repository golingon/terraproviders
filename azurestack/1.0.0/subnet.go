// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurestack

import (
	"encoding/json"
	"fmt"
	subnet "github.com/golingon/terraproviders/azurestack/1.0.0/subnet"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubnet creates a new instance of [Subnet].
func NewSubnet(name string, args SubnetArgs) *Subnet {
	return &Subnet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Subnet)(nil)

// Subnet represents the Terraform resource azurestack_subnet.
type Subnet struct {
	Name      string
	Args      SubnetArgs
	state     *subnetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Subnet].
func (s *Subnet) Type() string {
	return "azurestack_subnet"
}

// LocalName returns the local name for [Subnet].
func (s *Subnet) LocalName() string {
	return s.Name
}

// Configuration returns the configuration (args) for [Subnet].
func (s *Subnet) Configuration() interface{} {
	return s.Args
}

// DependOn is used for other resources to depend on [Subnet].
func (s *Subnet) DependOn() terra.Reference {
	return terra.ReferenceResource(s)
}

// Dependencies returns the list of resources [Subnet] depends_on.
func (s *Subnet) Dependencies() terra.Dependencies {
	return s.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Subnet].
func (s *Subnet) LifecycleManagement() *terra.Lifecycle {
	return s.Lifecycle
}

// Attributes returns the attributes for [Subnet].
func (s *Subnet) Attributes() subnetAttributes {
	return subnetAttributes{ref: terra.ReferenceResource(s)}
}

// ImportState imports the given attribute values into [Subnet]'s state.
func (s *Subnet) ImportState(av io.Reader) error {
	s.state = &subnetState{}
	if err := json.NewDecoder(av).Decode(s.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", s.Type(), s.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Subnet] has state.
func (s *Subnet) State() (*subnetState, bool) {
	return s.state, s.state != nil
}

// StateMust returns the state for [Subnet]. Panics if the state is nil.
func (s *Subnet) StateMust() *subnetState {
	if s.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", s.Type(), s.LocalName()))
	}
	return s.state
}

// SubnetArgs contains the configurations for azurestack_subnet.
type SubnetArgs struct {
	// AddressPrefix: string, required
	AddressPrefix terra.StringValue `hcl:"address_prefix,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// VirtualNetworkName: string, required
	VirtualNetworkName terra.StringValue `hcl:"virtual_network_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *subnet.Timeouts `hcl:"timeouts,block"`
}
type subnetAttributes struct {
	ref terra.Reference
}

// AddressPrefix returns a reference to field address_prefix of azurestack_subnet.
func (s subnetAttributes) AddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("address_prefix"))
}

// Id returns a reference to field id of azurestack_subnet.
func (s subnetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

// Name returns a reference to field name of azurestack_subnet.
func (s subnetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_subnet.
func (s subnetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("resource_group_name"))
}

// VirtualNetworkName returns a reference to field virtual_network_name of azurestack_subnet.
func (s subnetAttributes) VirtualNetworkName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("virtual_network_name"))
}

func (s subnetAttributes) Timeouts() subnet.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subnet.TimeoutsAttributes](s.ref.Append("timeouts"))
}

type subnetState struct {
	AddressPrefix      string                `json:"address_prefix"`
	Id                 string                `json:"id"`
	Name               string                `json:"name"`
	ResourceGroupName  string                `json:"resource_group_name"`
	VirtualNetworkName string                `json:"virtual_network_name"`
	Timeouts           *subnet.TimeoutsState `json:"timeouts"`
}
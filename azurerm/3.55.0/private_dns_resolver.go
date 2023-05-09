// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednsresolver "github.com/golingon/terraproviders/azurerm/3.55.0/privatednsresolver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsResolver creates a new instance of [PrivateDnsResolver].
func NewPrivateDnsResolver(name string, args PrivateDnsResolverArgs) *PrivateDnsResolver {
	return &PrivateDnsResolver{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsResolver)(nil)

// PrivateDnsResolver represents the Terraform resource azurerm_private_dns_resolver.
type PrivateDnsResolver struct {
	Name      string
	Args      PrivateDnsResolverArgs
	state     *privateDnsResolverState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsResolver].
func (pdr *PrivateDnsResolver) Type() string {
	return "azurerm_private_dns_resolver"
}

// LocalName returns the local name for [PrivateDnsResolver].
func (pdr *PrivateDnsResolver) LocalName() string {
	return pdr.Name
}

// Configuration returns the configuration (args) for [PrivateDnsResolver].
func (pdr *PrivateDnsResolver) Configuration() interface{} {
	return pdr.Args
}

// DependOn is used for other resources to depend on [PrivateDnsResolver].
func (pdr *PrivateDnsResolver) DependOn() terra.Reference {
	return terra.ReferenceResource(pdr)
}

// Dependencies returns the list of resources [PrivateDnsResolver] depends_on.
func (pdr *PrivateDnsResolver) Dependencies() terra.Dependencies {
	return pdr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsResolver].
func (pdr *PrivateDnsResolver) LifecycleManagement() *terra.Lifecycle {
	return pdr.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsResolver].
func (pdr *PrivateDnsResolver) Attributes() privateDnsResolverAttributes {
	return privateDnsResolverAttributes{ref: terra.ReferenceResource(pdr)}
}

// ImportState imports the given attribute values into [PrivateDnsResolver]'s state.
func (pdr *PrivateDnsResolver) ImportState(av io.Reader) error {
	pdr.state = &privateDnsResolverState{}
	if err := json.NewDecoder(av).Decode(pdr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdr.Type(), pdr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsResolver] has state.
func (pdr *PrivateDnsResolver) State() (*privateDnsResolverState, bool) {
	return pdr.state, pdr.state != nil
}

// StateMust returns the state for [PrivateDnsResolver]. Panics if the state is nil.
func (pdr *PrivateDnsResolver) StateMust() *privateDnsResolverState {
	if pdr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdr.Type(), pdr.LocalName()))
	}
	return pdr.state
}

// PrivateDnsResolverArgs contains the configurations for azurerm_private_dns_resolver.
type PrivateDnsResolverArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualNetworkId: string, required
	VirtualNetworkId terra.StringValue `hcl:"virtual_network_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *privatednsresolver.Timeouts `hcl:"timeouts,block"`
}
type privateDnsResolverAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_resolver.
func (pdr privateDnsResolverAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_private_dns_resolver.
func (pdr privateDnsResolverAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver.
func (pdr privateDnsResolverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_resolver.
func (pdr privateDnsResolverAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_resolver.
func (pdr privateDnsResolverAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdr.ref.Append("tags"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_private_dns_resolver.
func (pdr privateDnsResolverAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("virtual_network_id"))
}

func (pdr privateDnsResolverAttributes) Timeouts() privatednsresolver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednsresolver.TimeoutsAttributes](pdr.ref.Append("timeouts"))
}

type privateDnsResolverState struct {
	Id                string                            `json:"id"`
	Location          string                            `json:"location"`
	Name              string                            `json:"name"`
	ResourceGroupName string                            `json:"resource_group_name"`
	Tags              map[string]string                 `json:"tags"`
	VirtualNetworkId  string                            `json:"virtual_network_id"`
	Timeouts          *privatednsresolver.TimeoutsState `json:"timeouts"`
}

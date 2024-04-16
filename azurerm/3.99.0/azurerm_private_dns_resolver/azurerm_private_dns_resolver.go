// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_private_dns_resolver

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

// Resource represents the Terraform resource azurerm_private_dns_resolver.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermPrivateDnsResolverState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (apdr *Resource) Type() string {
	return "azurerm_private_dns_resolver"
}

// LocalName returns the local name for [Resource].
func (apdr *Resource) LocalName() string {
	return apdr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (apdr *Resource) Configuration() interface{} {
	return apdr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (apdr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(apdr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (apdr *Resource) Dependencies() terra.Dependencies {
	return apdr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (apdr *Resource) LifecycleManagement() *terra.Lifecycle {
	return apdr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (apdr *Resource) Attributes() azurermPrivateDnsResolverAttributes {
	return azurermPrivateDnsResolverAttributes{ref: terra.ReferenceResource(apdr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (apdr *Resource) ImportState(state io.Reader) error {
	apdr.state = &azurermPrivateDnsResolverState{}
	if err := json.NewDecoder(state).Decode(apdr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", apdr.Type(), apdr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (apdr *Resource) State() (*azurermPrivateDnsResolverState, bool) {
	return apdr.state, apdr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (apdr *Resource) StateMust() *azurermPrivateDnsResolverState {
	if apdr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", apdr.Type(), apdr.LocalName()))
	}
	return apdr.state
}

// Args contains the configurations for azurerm_private_dns_resolver.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermPrivateDnsResolverAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_resolver.
func (apdr azurermPrivateDnsResolverAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apdr.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_private_dns_resolver.
func (apdr azurermPrivateDnsResolverAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(apdr.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver.
func (apdr azurermPrivateDnsResolverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(apdr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_resolver.
func (apdr azurermPrivateDnsResolverAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(apdr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_resolver.
func (apdr azurermPrivateDnsResolverAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](apdr.ref.Append("tags"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_private_dns_resolver.
func (apdr azurermPrivateDnsResolverAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(apdr.ref.Append("virtual_network_id"))
}

func (apdr azurermPrivateDnsResolverAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](apdr.ref.Append("timeouts"))
}

type azurermPrivateDnsResolverState struct {
	Id                string            `json:"id"`
	Location          string            `json:"location"`
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	Tags              map[string]string `json:"tags"`
	VirtualNetworkId  string            `json:"virtual_network_id"`
	Timeouts          *TimeoutsState    `json:"timeouts"`
}

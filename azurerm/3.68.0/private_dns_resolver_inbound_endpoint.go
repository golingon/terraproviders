// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednsresolverinboundendpoint "github.com/golingon/terraproviders/azurerm/3.68.0/privatednsresolverinboundendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsResolverInboundEndpoint creates a new instance of [PrivateDnsResolverInboundEndpoint].
func NewPrivateDnsResolverInboundEndpoint(name string, args PrivateDnsResolverInboundEndpointArgs) *PrivateDnsResolverInboundEndpoint {
	return &PrivateDnsResolverInboundEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsResolverInboundEndpoint)(nil)

// PrivateDnsResolverInboundEndpoint represents the Terraform resource azurerm_private_dns_resolver_inbound_endpoint.
type PrivateDnsResolverInboundEndpoint struct {
	Name      string
	Args      PrivateDnsResolverInboundEndpointArgs
	state     *privateDnsResolverInboundEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsResolverInboundEndpoint].
func (pdrie *PrivateDnsResolverInboundEndpoint) Type() string {
	return "azurerm_private_dns_resolver_inbound_endpoint"
}

// LocalName returns the local name for [PrivateDnsResolverInboundEndpoint].
func (pdrie *PrivateDnsResolverInboundEndpoint) LocalName() string {
	return pdrie.Name
}

// Configuration returns the configuration (args) for [PrivateDnsResolverInboundEndpoint].
func (pdrie *PrivateDnsResolverInboundEndpoint) Configuration() interface{} {
	return pdrie.Args
}

// DependOn is used for other resources to depend on [PrivateDnsResolverInboundEndpoint].
func (pdrie *PrivateDnsResolverInboundEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(pdrie)
}

// Dependencies returns the list of resources [PrivateDnsResolverInboundEndpoint] depends_on.
func (pdrie *PrivateDnsResolverInboundEndpoint) Dependencies() terra.Dependencies {
	return pdrie.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsResolverInboundEndpoint].
func (pdrie *PrivateDnsResolverInboundEndpoint) LifecycleManagement() *terra.Lifecycle {
	return pdrie.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsResolverInboundEndpoint].
func (pdrie *PrivateDnsResolverInboundEndpoint) Attributes() privateDnsResolverInboundEndpointAttributes {
	return privateDnsResolverInboundEndpointAttributes{ref: terra.ReferenceResource(pdrie)}
}

// ImportState imports the given attribute values into [PrivateDnsResolverInboundEndpoint]'s state.
func (pdrie *PrivateDnsResolverInboundEndpoint) ImportState(av io.Reader) error {
	pdrie.state = &privateDnsResolverInboundEndpointState{}
	if err := json.NewDecoder(av).Decode(pdrie.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdrie.Type(), pdrie.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsResolverInboundEndpoint] has state.
func (pdrie *PrivateDnsResolverInboundEndpoint) State() (*privateDnsResolverInboundEndpointState, bool) {
	return pdrie.state, pdrie.state != nil
}

// StateMust returns the state for [PrivateDnsResolverInboundEndpoint]. Panics if the state is nil.
func (pdrie *PrivateDnsResolverInboundEndpoint) StateMust() *privateDnsResolverInboundEndpointState {
	if pdrie.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdrie.Type(), pdrie.LocalName()))
	}
	return pdrie.state
}

// PrivateDnsResolverInboundEndpointArgs contains the configurations for azurerm_private_dns_resolver_inbound_endpoint.
type PrivateDnsResolverInboundEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateDnsResolverId: string, required
	PrivateDnsResolverId terra.StringValue `hcl:"private_dns_resolver_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// IpConfigurations: min=1
	IpConfigurations []privatednsresolverinboundendpoint.IpConfigurations `hcl:"ip_configurations,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *privatednsresolverinboundendpoint.Timeouts `hcl:"timeouts,block"`
}
type privateDnsResolverInboundEndpointAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_resolver_inbound_endpoint.
func (pdrie privateDnsResolverInboundEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdrie.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_private_dns_resolver_inbound_endpoint.
func (pdrie privateDnsResolverInboundEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pdrie.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_inbound_endpoint.
func (pdrie privateDnsResolverInboundEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdrie.ref.Append("name"))
}

// PrivateDnsResolverId returns a reference to field private_dns_resolver_id of azurerm_private_dns_resolver_inbound_endpoint.
func (pdrie privateDnsResolverInboundEndpointAttributes) PrivateDnsResolverId() terra.StringValue {
	return terra.ReferenceAsString(pdrie.ref.Append("private_dns_resolver_id"))
}

// Tags returns a reference to field tags of azurerm_private_dns_resolver_inbound_endpoint.
func (pdrie privateDnsResolverInboundEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdrie.ref.Append("tags"))
}

func (pdrie privateDnsResolverInboundEndpointAttributes) IpConfigurations() terra.ListValue[privatednsresolverinboundendpoint.IpConfigurationsAttributes] {
	return terra.ReferenceAsList[privatednsresolverinboundendpoint.IpConfigurationsAttributes](pdrie.ref.Append("ip_configurations"))
}

func (pdrie privateDnsResolverInboundEndpointAttributes) Timeouts() privatednsresolverinboundendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednsresolverinboundendpoint.TimeoutsAttributes](pdrie.ref.Append("timeouts"))
}

type privateDnsResolverInboundEndpointState struct {
	Id                   string                                                    `json:"id"`
	Location             string                                                    `json:"location"`
	Name                 string                                                    `json:"name"`
	PrivateDnsResolverId string                                                    `json:"private_dns_resolver_id"`
	Tags                 map[string]string                                         `json:"tags"`
	IpConfigurations     []privatednsresolverinboundendpoint.IpConfigurationsState `json:"ip_configurations"`
	Timeouts             *privatednsresolverinboundendpoint.TimeoutsState          `json:"timeouts"`
}

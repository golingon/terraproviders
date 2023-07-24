// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednsresolveroutboundendpoint "github.com/golingon/terraproviders/azurerm/3.66.0/privatednsresolveroutboundendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsResolverOutboundEndpoint creates a new instance of [PrivateDnsResolverOutboundEndpoint].
func NewPrivateDnsResolverOutboundEndpoint(name string, args PrivateDnsResolverOutboundEndpointArgs) *PrivateDnsResolverOutboundEndpoint {
	return &PrivateDnsResolverOutboundEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsResolverOutboundEndpoint)(nil)

// PrivateDnsResolverOutboundEndpoint represents the Terraform resource azurerm_private_dns_resolver_outbound_endpoint.
type PrivateDnsResolverOutboundEndpoint struct {
	Name      string
	Args      PrivateDnsResolverOutboundEndpointArgs
	state     *privateDnsResolverOutboundEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsResolverOutboundEndpoint].
func (pdroe *PrivateDnsResolverOutboundEndpoint) Type() string {
	return "azurerm_private_dns_resolver_outbound_endpoint"
}

// LocalName returns the local name for [PrivateDnsResolverOutboundEndpoint].
func (pdroe *PrivateDnsResolverOutboundEndpoint) LocalName() string {
	return pdroe.Name
}

// Configuration returns the configuration (args) for [PrivateDnsResolverOutboundEndpoint].
func (pdroe *PrivateDnsResolverOutboundEndpoint) Configuration() interface{} {
	return pdroe.Args
}

// DependOn is used for other resources to depend on [PrivateDnsResolverOutboundEndpoint].
func (pdroe *PrivateDnsResolverOutboundEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(pdroe)
}

// Dependencies returns the list of resources [PrivateDnsResolverOutboundEndpoint] depends_on.
func (pdroe *PrivateDnsResolverOutboundEndpoint) Dependencies() terra.Dependencies {
	return pdroe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsResolverOutboundEndpoint].
func (pdroe *PrivateDnsResolverOutboundEndpoint) LifecycleManagement() *terra.Lifecycle {
	return pdroe.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsResolverOutboundEndpoint].
func (pdroe *PrivateDnsResolverOutboundEndpoint) Attributes() privateDnsResolverOutboundEndpointAttributes {
	return privateDnsResolverOutboundEndpointAttributes{ref: terra.ReferenceResource(pdroe)}
}

// ImportState imports the given attribute values into [PrivateDnsResolverOutboundEndpoint]'s state.
func (pdroe *PrivateDnsResolverOutboundEndpoint) ImportState(av io.Reader) error {
	pdroe.state = &privateDnsResolverOutboundEndpointState{}
	if err := json.NewDecoder(av).Decode(pdroe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdroe.Type(), pdroe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsResolverOutboundEndpoint] has state.
func (pdroe *PrivateDnsResolverOutboundEndpoint) State() (*privateDnsResolverOutboundEndpointState, bool) {
	return pdroe.state, pdroe.state != nil
}

// StateMust returns the state for [PrivateDnsResolverOutboundEndpoint]. Panics if the state is nil.
func (pdroe *PrivateDnsResolverOutboundEndpoint) StateMust() *privateDnsResolverOutboundEndpointState {
	if pdroe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdroe.Type(), pdroe.LocalName()))
	}
	return pdroe.state
}

// PrivateDnsResolverOutboundEndpointArgs contains the configurations for azurerm_private_dns_resolver_outbound_endpoint.
type PrivateDnsResolverOutboundEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateDnsResolverId: string, required
	PrivateDnsResolverId terra.StringValue `hcl:"private_dns_resolver_id,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *privatednsresolveroutboundendpoint.Timeouts `hcl:"timeouts,block"`
}
type privateDnsResolverOutboundEndpointAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe privateDnsResolverOutboundEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdroe.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe privateDnsResolverOutboundEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pdroe.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe privateDnsResolverOutboundEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdroe.ref.Append("name"))
}

// PrivateDnsResolverId returns a reference to field private_dns_resolver_id of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe privateDnsResolverOutboundEndpointAttributes) PrivateDnsResolverId() terra.StringValue {
	return terra.ReferenceAsString(pdroe.ref.Append("private_dns_resolver_id"))
}

// SubnetId returns a reference to field subnet_id of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe privateDnsResolverOutboundEndpointAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(pdroe.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe privateDnsResolverOutboundEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdroe.ref.Append("tags"))
}

func (pdroe privateDnsResolverOutboundEndpointAttributes) Timeouts() privatednsresolveroutboundendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednsresolveroutboundendpoint.TimeoutsAttributes](pdroe.ref.Append("timeouts"))
}

type privateDnsResolverOutboundEndpointState struct {
	Id                   string                                            `json:"id"`
	Location             string                                            `json:"location"`
	Name                 string                                            `json:"name"`
	PrivateDnsResolverId string                                            `json:"private_dns_resolver_id"`
	SubnetId             string                                            `json:"subnet_id"`
	Tags                 map[string]string                                 `json:"tags"`
	Timeouts             *privatednsresolveroutboundendpoint.TimeoutsState `json:"timeouts"`
}

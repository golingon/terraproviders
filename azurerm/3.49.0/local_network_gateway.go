// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	localnetworkgateway "github.com/golingon/terraproviders/azurerm/3.49.0/localnetworkgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLocalNetworkGateway creates a new instance of [LocalNetworkGateway].
func NewLocalNetworkGateway(name string, args LocalNetworkGatewayArgs) *LocalNetworkGateway {
	return &LocalNetworkGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LocalNetworkGateway)(nil)

// LocalNetworkGateway represents the Terraform resource azurerm_local_network_gateway.
type LocalNetworkGateway struct {
	Name      string
	Args      LocalNetworkGatewayArgs
	state     *localNetworkGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LocalNetworkGateway].
func (lng *LocalNetworkGateway) Type() string {
	return "azurerm_local_network_gateway"
}

// LocalName returns the local name for [LocalNetworkGateway].
func (lng *LocalNetworkGateway) LocalName() string {
	return lng.Name
}

// Configuration returns the configuration (args) for [LocalNetworkGateway].
func (lng *LocalNetworkGateway) Configuration() interface{} {
	return lng.Args
}

// DependOn is used for other resources to depend on [LocalNetworkGateway].
func (lng *LocalNetworkGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(lng)
}

// Dependencies returns the list of resources [LocalNetworkGateway] depends_on.
func (lng *LocalNetworkGateway) Dependencies() terra.Dependencies {
	return lng.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LocalNetworkGateway].
func (lng *LocalNetworkGateway) LifecycleManagement() *terra.Lifecycle {
	return lng.Lifecycle
}

// Attributes returns the attributes for [LocalNetworkGateway].
func (lng *LocalNetworkGateway) Attributes() localNetworkGatewayAttributes {
	return localNetworkGatewayAttributes{ref: terra.ReferenceResource(lng)}
}

// ImportState imports the given attribute values into [LocalNetworkGateway]'s state.
func (lng *LocalNetworkGateway) ImportState(av io.Reader) error {
	lng.state = &localNetworkGatewayState{}
	if err := json.NewDecoder(av).Decode(lng.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lng.Type(), lng.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LocalNetworkGateway] has state.
func (lng *LocalNetworkGateway) State() (*localNetworkGatewayState, bool) {
	return lng.state, lng.state != nil
}

// StateMust returns the state for [LocalNetworkGateway]. Panics if the state is nil.
func (lng *LocalNetworkGateway) StateMust() *localNetworkGatewayState {
	if lng.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lng.Type(), lng.LocalName()))
	}
	return lng.state
}

// LocalNetworkGatewayArgs contains the configurations for azurerm_local_network_gateway.
type LocalNetworkGatewayArgs struct {
	// AddressSpace: list of string, optional
	AddressSpace terra.ListValue[terra.StringValue] `hcl:"address_space,attr"`
	// GatewayAddress: string, optional
	GatewayAddress terra.StringValue `hcl:"gateway_address,attr"`
	// GatewayFqdn: string, optional
	GatewayFqdn terra.StringValue `hcl:"gateway_fqdn,attr"`
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
	// BgpSettings: optional
	BgpSettings *localnetworkgateway.BgpSettings `hcl:"bgp_settings,block"`
	// Timeouts: optional
	Timeouts *localnetworkgateway.Timeouts `hcl:"timeouts,block"`
}
type localNetworkGatewayAttributes struct {
	ref terra.Reference
}

// AddressSpace returns a reference to field address_space of azurerm_local_network_gateway.
func (lng localNetworkGatewayAttributes) AddressSpace() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lng.ref.Append("address_space"))
}

// GatewayAddress returns a reference to field gateway_address of azurerm_local_network_gateway.
func (lng localNetworkGatewayAttributes) GatewayAddress() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("gateway_address"))
}

// GatewayFqdn returns a reference to field gateway_fqdn of azurerm_local_network_gateway.
func (lng localNetworkGatewayAttributes) GatewayFqdn() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("gateway_fqdn"))
}

// Id returns a reference to field id of azurerm_local_network_gateway.
func (lng localNetworkGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_local_network_gateway.
func (lng localNetworkGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_local_network_gateway.
func (lng localNetworkGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_local_network_gateway.
func (lng localNetworkGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_local_network_gateway.
func (lng localNetworkGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lng.ref.Append("tags"))
}

func (lng localNetworkGatewayAttributes) BgpSettings() terra.ListValue[localnetworkgateway.BgpSettingsAttributes] {
	return terra.ReferenceAsList[localnetworkgateway.BgpSettingsAttributes](lng.ref.Append("bgp_settings"))
}

func (lng localNetworkGatewayAttributes) Timeouts() localnetworkgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[localnetworkgateway.TimeoutsAttributes](lng.ref.Append("timeouts"))
}

type localNetworkGatewayState struct {
	AddressSpace      []string                               `json:"address_space"`
	GatewayAddress    string                                 `json:"gateway_address"`
	GatewayFqdn       string                                 `json:"gateway_fqdn"`
	Id                string                                 `json:"id"`
	Location          string                                 `json:"location"`
	Name              string                                 `json:"name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	Tags              map[string]string                      `json:"tags"`
	BgpSettings       []localnetworkgateway.BgpSettingsState `json:"bgp_settings"`
	Timeouts          *localnetworkgateway.TimeoutsState     `json:"timeouts"`
}

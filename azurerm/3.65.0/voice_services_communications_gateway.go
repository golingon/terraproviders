// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	voiceservicescommunicationsgateway "github.com/golingon/terraproviders/azurerm/3.65.0/voiceservicescommunicationsgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVoiceServicesCommunicationsGateway creates a new instance of [VoiceServicesCommunicationsGateway].
func NewVoiceServicesCommunicationsGateway(name string, args VoiceServicesCommunicationsGatewayArgs) *VoiceServicesCommunicationsGateway {
	return &VoiceServicesCommunicationsGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VoiceServicesCommunicationsGateway)(nil)

// VoiceServicesCommunicationsGateway represents the Terraform resource azurerm_voice_services_communications_gateway.
type VoiceServicesCommunicationsGateway struct {
	Name      string
	Args      VoiceServicesCommunicationsGatewayArgs
	state     *voiceServicesCommunicationsGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VoiceServicesCommunicationsGateway].
func (vscg *VoiceServicesCommunicationsGateway) Type() string {
	return "azurerm_voice_services_communications_gateway"
}

// LocalName returns the local name for [VoiceServicesCommunicationsGateway].
func (vscg *VoiceServicesCommunicationsGateway) LocalName() string {
	return vscg.Name
}

// Configuration returns the configuration (args) for [VoiceServicesCommunicationsGateway].
func (vscg *VoiceServicesCommunicationsGateway) Configuration() interface{} {
	return vscg.Args
}

// DependOn is used for other resources to depend on [VoiceServicesCommunicationsGateway].
func (vscg *VoiceServicesCommunicationsGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(vscg)
}

// Dependencies returns the list of resources [VoiceServicesCommunicationsGateway] depends_on.
func (vscg *VoiceServicesCommunicationsGateway) Dependencies() terra.Dependencies {
	return vscg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VoiceServicesCommunicationsGateway].
func (vscg *VoiceServicesCommunicationsGateway) LifecycleManagement() *terra.Lifecycle {
	return vscg.Lifecycle
}

// Attributes returns the attributes for [VoiceServicesCommunicationsGateway].
func (vscg *VoiceServicesCommunicationsGateway) Attributes() voiceServicesCommunicationsGatewayAttributes {
	return voiceServicesCommunicationsGatewayAttributes{ref: terra.ReferenceResource(vscg)}
}

// ImportState imports the given attribute values into [VoiceServicesCommunicationsGateway]'s state.
func (vscg *VoiceServicesCommunicationsGateway) ImportState(av io.Reader) error {
	vscg.state = &voiceServicesCommunicationsGatewayState{}
	if err := json.NewDecoder(av).Decode(vscg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vscg.Type(), vscg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VoiceServicesCommunicationsGateway] has state.
func (vscg *VoiceServicesCommunicationsGateway) State() (*voiceServicesCommunicationsGatewayState, bool) {
	return vscg.state, vscg.state != nil
}

// StateMust returns the state for [VoiceServicesCommunicationsGateway]. Panics if the state is nil.
func (vscg *VoiceServicesCommunicationsGateway) StateMust() *voiceServicesCommunicationsGatewayState {
	if vscg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vscg.Type(), vscg.LocalName()))
	}
	return vscg.state
}

// VoiceServicesCommunicationsGatewayArgs contains the configurations for azurerm_voice_services_communications_gateway.
type VoiceServicesCommunicationsGatewayArgs struct {
	// ApiBridge: string, optional
	ApiBridge terra.StringValue `hcl:"api_bridge,attr"`
	// AutoGeneratedDomainNameLabelScope: string, optional
	AutoGeneratedDomainNameLabelScope terra.StringValue `hcl:"auto_generated_domain_name_label_scope,attr"`
	// Codecs: string, required
	Codecs terra.StringValue `hcl:"codecs,attr" validate:"required"`
	// Connectivity: string, required
	Connectivity terra.StringValue `hcl:"connectivity,attr" validate:"required"`
	// E911Type: string, required
	E911Type terra.StringValue `hcl:"e911_type,attr" validate:"required"`
	// EmergencyDialStrings: list of string, optional
	EmergencyDialStrings terra.ListValue[terra.StringValue] `hcl:"emergency_dial_strings,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MicrosoftTeamsVoicemailPilotNumber: string, optional
	MicrosoftTeamsVoicemailPilotNumber terra.StringValue `hcl:"microsoft_teams_voicemail_pilot_number,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OnPremMcpEnabled: bool, optional
	OnPremMcpEnabled terra.BoolValue `hcl:"on_prem_mcp_enabled,attr"`
	// Platforms: list of string, required
	Platforms terra.ListValue[terra.StringValue] `hcl:"platforms,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ServiceLocation: min=1
	ServiceLocation []voiceservicescommunicationsgateway.ServiceLocation `hcl:"service_location,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *voiceservicescommunicationsgateway.Timeouts `hcl:"timeouts,block"`
}
type voiceServicesCommunicationsGatewayAttributes struct {
	ref terra.Reference
}

// ApiBridge returns a reference to field api_bridge of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) ApiBridge() terra.StringValue {
	return terra.ReferenceAsString(vscg.ref.Append("api_bridge"))
}

// AutoGeneratedDomainNameLabelScope returns a reference to field auto_generated_domain_name_label_scope of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) AutoGeneratedDomainNameLabelScope() terra.StringValue {
	return terra.ReferenceAsString(vscg.ref.Append("auto_generated_domain_name_label_scope"))
}

// Codecs returns a reference to field codecs of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) Codecs() terra.StringValue {
	return terra.ReferenceAsString(vscg.ref.Append("codecs"))
}

// Connectivity returns a reference to field connectivity of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) Connectivity() terra.StringValue {
	return terra.ReferenceAsString(vscg.ref.Append("connectivity"))
}

// E911Type returns a reference to field e911_type of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) E911Type() terra.StringValue {
	return terra.ReferenceAsString(vscg.ref.Append("e911_type"))
}

// EmergencyDialStrings returns a reference to field emergency_dial_strings of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) EmergencyDialStrings() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vscg.ref.Append("emergency_dial_strings"))
}

// Id returns a reference to field id of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vscg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vscg.ref.Append("location"))
}

// MicrosoftTeamsVoicemailPilotNumber returns a reference to field microsoft_teams_voicemail_pilot_number of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) MicrosoftTeamsVoicemailPilotNumber() terra.StringValue {
	return terra.ReferenceAsString(vscg.ref.Append("microsoft_teams_voicemail_pilot_number"))
}

// Name returns a reference to field name of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vscg.ref.Append("name"))
}

// OnPremMcpEnabled returns a reference to field on_prem_mcp_enabled of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) OnPremMcpEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vscg.ref.Append("on_prem_mcp_enabled"))
}

// Platforms returns a reference to field platforms of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) Platforms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vscg.ref.Append("platforms"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vscg.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_voice_services_communications_gateway.
func (vscg voiceServicesCommunicationsGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vscg.ref.Append("tags"))
}

func (vscg voiceServicesCommunicationsGatewayAttributes) ServiceLocation() terra.SetValue[voiceservicescommunicationsgateway.ServiceLocationAttributes] {
	return terra.ReferenceAsSet[voiceservicescommunicationsgateway.ServiceLocationAttributes](vscg.ref.Append("service_location"))
}

func (vscg voiceServicesCommunicationsGatewayAttributes) Timeouts() voiceservicescommunicationsgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[voiceservicescommunicationsgateway.TimeoutsAttributes](vscg.ref.Append("timeouts"))
}

type voiceServicesCommunicationsGatewayState struct {
	ApiBridge                          string                                                    `json:"api_bridge"`
	AutoGeneratedDomainNameLabelScope  string                                                    `json:"auto_generated_domain_name_label_scope"`
	Codecs                             string                                                    `json:"codecs"`
	Connectivity                       string                                                    `json:"connectivity"`
	E911Type                           string                                                    `json:"e911_type"`
	EmergencyDialStrings               []string                                                  `json:"emergency_dial_strings"`
	Id                                 string                                                    `json:"id"`
	Location                           string                                                    `json:"location"`
	MicrosoftTeamsVoicemailPilotNumber string                                                    `json:"microsoft_teams_voicemail_pilot_number"`
	Name                               string                                                    `json:"name"`
	OnPremMcpEnabled                   bool                                                      `json:"on_prem_mcp_enabled"`
	Platforms                          []string                                                  `json:"platforms"`
	ResourceGroupName                  string                                                    `json:"resource_group_name"`
	Tags                               map[string]string                                         `json:"tags"`
	ServiceLocation                    []voiceservicescommunicationsgateway.ServiceLocationState `json:"service_location"`
	Timeouts                           *voiceservicescommunicationsgateway.TimeoutsState         `json:"timeouts"`
}

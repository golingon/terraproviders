// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	voiceservicescommunicationsgatewaytestline "github.com/golingon/terraproviders/azurerm/3.58.0/voiceservicescommunicationsgatewaytestline"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVoiceServicesCommunicationsGatewayTestLine creates a new instance of [VoiceServicesCommunicationsGatewayTestLine].
func NewVoiceServicesCommunicationsGatewayTestLine(name string, args VoiceServicesCommunicationsGatewayTestLineArgs) *VoiceServicesCommunicationsGatewayTestLine {
	return &VoiceServicesCommunicationsGatewayTestLine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VoiceServicesCommunicationsGatewayTestLine)(nil)

// VoiceServicesCommunicationsGatewayTestLine represents the Terraform resource azurerm_voice_services_communications_gateway_test_line.
type VoiceServicesCommunicationsGatewayTestLine struct {
	Name      string
	Args      VoiceServicesCommunicationsGatewayTestLineArgs
	state     *voiceServicesCommunicationsGatewayTestLineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VoiceServicesCommunicationsGatewayTestLine].
func (vscgtl *VoiceServicesCommunicationsGatewayTestLine) Type() string {
	return "azurerm_voice_services_communications_gateway_test_line"
}

// LocalName returns the local name for [VoiceServicesCommunicationsGatewayTestLine].
func (vscgtl *VoiceServicesCommunicationsGatewayTestLine) LocalName() string {
	return vscgtl.Name
}

// Configuration returns the configuration (args) for [VoiceServicesCommunicationsGatewayTestLine].
func (vscgtl *VoiceServicesCommunicationsGatewayTestLine) Configuration() interface{} {
	return vscgtl.Args
}

// DependOn is used for other resources to depend on [VoiceServicesCommunicationsGatewayTestLine].
func (vscgtl *VoiceServicesCommunicationsGatewayTestLine) DependOn() terra.Reference {
	return terra.ReferenceResource(vscgtl)
}

// Dependencies returns the list of resources [VoiceServicesCommunicationsGatewayTestLine] depends_on.
func (vscgtl *VoiceServicesCommunicationsGatewayTestLine) Dependencies() terra.Dependencies {
	return vscgtl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VoiceServicesCommunicationsGatewayTestLine].
func (vscgtl *VoiceServicesCommunicationsGatewayTestLine) LifecycleManagement() *terra.Lifecycle {
	return vscgtl.Lifecycle
}

// Attributes returns the attributes for [VoiceServicesCommunicationsGatewayTestLine].
func (vscgtl *VoiceServicesCommunicationsGatewayTestLine) Attributes() voiceServicesCommunicationsGatewayTestLineAttributes {
	return voiceServicesCommunicationsGatewayTestLineAttributes{ref: terra.ReferenceResource(vscgtl)}
}

// ImportState imports the given attribute values into [VoiceServicesCommunicationsGatewayTestLine]'s state.
func (vscgtl *VoiceServicesCommunicationsGatewayTestLine) ImportState(av io.Reader) error {
	vscgtl.state = &voiceServicesCommunicationsGatewayTestLineState{}
	if err := json.NewDecoder(av).Decode(vscgtl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vscgtl.Type(), vscgtl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VoiceServicesCommunicationsGatewayTestLine] has state.
func (vscgtl *VoiceServicesCommunicationsGatewayTestLine) State() (*voiceServicesCommunicationsGatewayTestLineState, bool) {
	return vscgtl.state, vscgtl.state != nil
}

// StateMust returns the state for [VoiceServicesCommunicationsGatewayTestLine]. Panics if the state is nil.
func (vscgtl *VoiceServicesCommunicationsGatewayTestLine) StateMust() *voiceServicesCommunicationsGatewayTestLineState {
	if vscgtl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vscgtl.Type(), vscgtl.LocalName()))
	}
	return vscgtl.state
}

// VoiceServicesCommunicationsGatewayTestLineArgs contains the configurations for azurerm_voice_services_communications_gateway_test_line.
type VoiceServicesCommunicationsGatewayTestLineArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PhoneNumber: string, required
	PhoneNumber terra.StringValue `hcl:"phone_number,attr" validate:"required"`
	// Purpose: string, required
	Purpose terra.StringValue `hcl:"purpose,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VoiceServicesCommunicationsGatewayId: string, required
	VoiceServicesCommunicationsGatewayId terra.StringValue `hcl:"voice_services_communications_gateway_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *voiceservicescommunicationsgatewaytestline.Timeouts `hcl:"timeouts,block"`
}
type voiceServicesCommunicationsGatewayTestLineAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_voice_services_communications_gateway_test_line.
func (vscgtl voiceServicesCommunicationsGatewayTestLineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vscgtl.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_voice_services_communications_gateway_test_line.
func (vscgtl voiceServicesCommunicationsGatewayTestLineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vscgtl.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_voice_services_communications_gateway_test_line.
func (vscgtl voiceServicesCommunicationsGatewayTestLineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vscgtl.ref.Append("name"))
}

// PhoneNumber returns a reference to field phone_number of azurerm_voice_services_communications_gateway_test_line.
func (vscgtl voiceServicesCommunicationsGatewayTestLineAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceAsString(vscgtl.ref.Append("phone_number"))
}

// Purpose returns a reference to field purpose of azurerm_voice_services_communications_gateway_test_line.
func (vscgtl voiceServicesCommunicationsGatewayTestLineAttributes) Purpose() terra.StringValue {
	return terra.ReferenceAsString(vscgtl.ref.Append("purpose"))
}

// Tags returns a reference to field tags of azurerm_voice_services_communications_gateway_test_line.
func (vscgtl voiceServicesCommunicationsGatewayTestLineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vscgtl.ref.Append("tags"))
}

// VoiceServicesCommunicationsGatewayId returns a reference to field voice_services_communications_gateway_id of azurerm_voice_services_communications_gateway_test_line.
func (vscgtl voiceServicesCommunicationsGatewayTestLineAttributes) VoiceServicesCommunicationsGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vscgtl.ref.Append("voice_services_communications_gateway_id"))
}

func (vscgtl voiceServicesCommunicationsGatewayTestLineAttributes) Timeouts() voiceservicescommunicationsgatewaytestline.TimeoutsAttributes {
	return terra.ReferenceAsSingle[voiceservicescommunicationsgatewaytestline.TimeoutsAttributes](vscgtl.ref.Append("timeouts"))
}

type voiceServicesCommunicationsGatewayTestLineState struct {
	Id                                   string                                                    `json:"id"`
	Location                             string                                                    `json:"location"`
	Name                                 string                                                    `json:"name"`
	PhoneNumber                          string                                                    `json:"phone_number"`
	Purpose                              string                                                    `json:"purpose"`
	Tags                                 map[string]string                                         `json:"tags"`
	VoiceServicesCommunicationsGatewayId string                                                    `json:"voice_services_communications_gateway_id"`
	Timeouts                             *voiceservicescommunicationsgatewaytestline.TimeoutsState `json:"timeouts"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	chimesdkvoiceglobalsettings "github.com/golingon/terraproviders/aws/5.6.2/chimesdkvoiceglobalsettings"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChimesdkvoiceGlobalSettings creates a new instance of [ChimesdkvoiceGlobalSettings].
func NewChimesdkvoiceGlobalSettings(name string, args ChimesdkvoiceGlobalSettingsArgs) *ChimesdkvoiceGlobalSettings {
	return &ChimesdkvoiceGlobalSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChimesdkvoiceGlobalSettings)(nil)

// ChimesdkvoiceGlobalSettings represents the Terraform resource aws_chimesdkvoice_global_settings.
type ChimesdkvoiceGlobalSettings struct {
	Name      string
	Args      ChimesdkvoiceGlobalSettingsArgs
	state     *chimesdkvoiceGlobalSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChimesdkvoiceGlobalSettings].
func (cgs *ChimesdkvoiceGlobalSettings) Type() string {
	return "aws_chimesdkvoice_global_settings"
}

// LocalName returns the local name for [ChimesdkvoiceGlobalSettings].
func (cgs *ChimesdkvoiceGlobalSettings) LocalName() string {
	return cgs.Name
}

// Configuration returns the configuration (args) for [ChimesdkvoiceGlobalSettings].
func (cgs *ChimesdkvoiceGlobalSettings) Configuration() interface{} {
	return cgs.Args
}

// DependOn is used for other resources to depend on [ChimesdkvoiceGlobalSettings].
func (cgs *ChimesdkvoiceGlobalSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(cgs)
}

// Dependencies returns the list of resources [ChimesdkvoiceGlobalSettings] depends_on.
func (cgs *ChimesdkvoiceGlobalSettings) Dependencies() terra.Dependencies {
	return cgs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChimesdkvoiceGlobalSettings].
func (cgs *ChimesdkvoiceGlobalSettings) LifecycleManagement() *terra.Lifecycle {
	return cgs.Lifecycle
}

// Attributes returns the attributes for [ChimesdkvoiceGlobalSettings].
func (cgs *ChimesdkvoiceGlobalSettings) Attributes() chimesdkvoiceGlobalSettingsAttributes {
	return chimesdkvoiceGlobalSettingsAttributes{ref: terra.ReferenceResource(cgs)}
}

// ImportState imports the given attribute values into [ChimesdkvoiceGlobalSettings]'s state.
func (cgs *ChimesdkvoiceGlobalSettings) ImportState(av io.Reader) error {
	cgs.state = &chimesdkvoiceGlobalSettingsState{}
	if err := json.NewDecoder(av).Decode(cgs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cgs.Type(), cgs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChimesdkvoiceGlobalSettings] has state.
func (cgs *ChimesdkvoiceGlobalSettings) State() (*chimesdkvoiceGlobalSettingsState, bool) {
	return cgs.state, cgs.state != nil
}

// StateMust returns the state for [ChimesdkvoiceGlobalSettings]. Panics if the state is nil.
func (cgs *ChimesdkvoiceGlobalSettings) StateMust() *chimesdkvoiceGlobalSettingsState {
	if cgs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cgs.Type(), cgs.LocalName()))
	}
	return cgs.state
}

// ChimesdkvoiceGlobalSettingsArgs contains the configurations for aws_chimesdkvoice_global_settings.
type ChimesdkvoiceGlobalSettingsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VoiceConnector: required
	VoiceConnector *chimesdkvoiceglobalsettings.VoiceConnector `hcl:"voice_connector,block" validate:"required"`
}
type chimesdkvoiceGlobalSettingsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_chimesdkvoice_global_settings.
func (cgs chimesdkvoiceGlobalSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cgs.ref.Append("id"))
}

func (cgs chimesdkvoiceGlobalSettingsAttributes) VoiceConnector() terra.ListValue[chimesdkvoiceglobalsettings.VoiceConnectorAttributes] {
	return terra.ReferenceAsList[chimesdkvoiceglobalsettings.VoiceConnectorAttributes](cgs.ref.Append("voice_connector"))
}

type chimesdkvoiceGlobalSettingsState struct {
	Id             string                                            `json:"id"`
	VoiceConnector []chimesdkvoiceglobalsettings.VoiceConnectorState `json:"voice_connector"`
}

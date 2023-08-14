// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	chimevoiceconnectororigination "github.com/golingon/terraproviders/aws/5.12.0/chimevoiceconnectororigination"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChimeVoiceConnectorOrigination creates a new instance of [ChimeVoiceConnectorOrigination].
func NewChimeVoiceConnectorOrigination(name string, args ChimeVoiceConnectorOriginationArgs) *ChimeVoiceConnectorOrigination {
	return &ChimeVoiceConnectorOrigination{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChimeVoiceConnectorOrigination)(nil)

// ChimeVoiceConnectorOrigination represents the Terraform resource aws_chime_voice_connector_origination.
type ChimeVoiceConnectorOrigination struct {
	Name      string
	Args      ChimeVoiceConnectorOriginationArgs
	state     *chimeVoiceConnectorOriginationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChimeVoiceConnectorOrigination].
func (cvco *ChimeVoiceConnectorOrigination) Type() string {
	return "aws_chime_voice_connector_origination"
}

// LocalName returns the local name for [ChimeVoiceConnectorOrigination].
func (cvco *ChimeVoiceConnectorOrigination) LocalName() string {
	return cvco.Name
}

// Configuration returns the configuration (args) for [ChimeVoiceConnectorOrigination].
func (cvco *ChimeVoiceConnectorOrigination) Configuration() interface{} {
	return cvco.Args
}

// DependOn is used for other resources to depend on [ChimeVoiceConnectorOrigination].
func (cvco *ChimeVoiceConnectorOrigination) DependOn() terra.Reference {
	return terra.ReferenceResource(cvco)
}

// Dependencies returns the list of resources [ChimeVoiceConnectorOrigination] depends_on.
func (cvco *ChimeVoiceConnectorOrigination) Dependencies() terra.Dependencies {
	return cvco.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChimeVoiceConnectorOrigination].
func (cvco *ChimeVoiceConnectorOrigination) LifecycleManagement() *terra.Lifecycle {
	return cvco.Lifecycle
}

// Attributes returns the attributes for [ChimeVoiceConnectorOrigination].
func (cvco *ChimeVoiceConnectorOrigination) Attributes() chimeVoiceConnectorOriginationAttributes {
	return chimeVoiceConnectorOriginationAttributes{ref: terra.ReferenceResource(cvco)}
}

// ImportState imports the given attribute values into [ChimeVoiceConnectorOrigination]'s state.
func (cvco *ChimeVoiceConnectorOrigination) ImportState(av io.Reader) error {
	cvco.state = &chimeVoiceConnectorOriginationState{}
	if err := json.NewDecoder(av).Decode(cvco.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cvco.Type(), cvco.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChimeVoiceConnectorOrigination] has state.
func (cvco *ChimeVoiceConnectorOrigination) State() (*chimeVoiceConnectorOriginationState, bool) {
	return cvco.state, cvco.state != nil
}

// StateMust returns the state for [ChimeVoiceConnectorOrigination]. Panics if the state is nil.
func (cvco *ChimeVoiceConnectorOrigination) StateMust() *chimeVoiceConnectorOriginationState {
	if cvco.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cvco.Type(), cvco.LocalName()))
	}
	return cvco.state
}

// ChimeVoiceConnectorOriginationArgs contains the configurations for aws_chime_voice_connector_origination.
type ChimeVoiceConnectorOriginationArgs struct {
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VoiceConnectorId: string, required
	VoiceConnectorId terra.StringValue `hcl:"voice_connector_id,attr" validate:"required"`
	// Route: min=1,max=20
	Route []chimevoiceconnectororigination.Route `hcl:"route,block" validate:"min=1,max=20"`
}
type chimeVoiceConnectorOriginationAttributes struct {
	ref terra.Reference
}

// Disabled returns a reference to field disabled of aws_chime_voice_connector_origination.
func (cvco chimeVoiceConnectorOriginationAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(cvco.ref.Append("disabled"))
}

// Id returns a reference to field id of aws_chime_voice_connector_origination.
func (cvco chimeVoiceConnectorOriginationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvco.ref.Append("id"))
}

// VoiceConnectorId returns a reference to field voice_connector_id of aws_chime_voice_connector_origination.
func (cvco chimeVoiceConnectorOriginationAttributes) VoiceConnectorId() terra.StringValue {
	return terra.ReferenceAsString(cvco.ref.Append("voice_connector_id"))
}

func (cvco chimeVoiceConnectorOriginationAttributes) Route() terra.SetValue[chimevoiceconnectororigination.RouteAttributes] {
	return terra.ReferenceAsSet[chimevoiceconnectororigination.RouteAttributes](cvco.ref.Append("route"))
}

type chimeVoiceConnectorOriginationState struct {
	Disabled         bool                                        `json:"disabled"`
	Id               string                                      `json:"id"`
	VoiceConnectorId string                                      `json:"voice_connector_id"`
	Route            []chimevoiceconnectororigination.RouteState `json:"route"`
}

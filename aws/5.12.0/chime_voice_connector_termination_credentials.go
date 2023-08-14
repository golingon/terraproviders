// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	chimevoiceconnectorterminationcredentials "github.com/golingon/terraproviders/aws/5.12.0/chimevoiceconnectorterminationcredentials"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChimeVoiceConnectorTerminationCredentials creates a new instance of [ChimeVoiceConnectorTerminationCredentials].
func NewChimeVoiceConnectorTerminationCredentials(name string, args ChimeVoiceConnectorTerminationCredentialsArgs) *ChimeVoiceConnectorTerminationCredentials {
	return &ChimeVoiceConnectorTerminationCredentials{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChimeVoiceConnectorTerminationCredentials)(nil)

// ChimeVoiceConnectorTerminationCredentials represents the Terraform resource aws_chime_voice_connector_termination_credentials.
type ChimeVoiceConnectorTerminationCredentials struct {
	Name      string
	Args      ChimeVoiceConnectorTerminationCredentialsArgs
	state     *chimeVoiceConnectorTerminationCredentialsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChimeVoiceConnectorTerminationCredentials].
func (cvctc *ChimeVoiceConnectorTerminationCredentials) Type() string {
	return "aws_chime_voice_connector_termination_credentials"
}

// LocalName returns the local name for [ChimeVoiceConnectorTerminationCredentials].
func (cvctc *ChimeVoiceConnectorTerminationCredentials) LocalName() string {
	return cvctc.Name
}

// Configuration returns the configuration (args) for [ChimeVoiceConnectorTerminationCredentials].
func (cvctc *ChimeVoiceConnectorTerminationCredentials) Configuration() interface{} {
	return cvctc.Args
}

// DependOn is used for other resources to depend on [ChimeVoiceConnectorTerminationCredentials].
func (cvctc *ChimeVoiceConnectorTerminationCredentials) DependOn() terra.Reference {
	return terra.ReferenceResource(cvctc)
}

// Dependencies returns the list of resources [ChimeVoiceConnectorTerminationCredentials] depends_on.
func (cvctc *ChimeVoiceConnectorTerminationCredentials) Dependencies() terra.Dependencies {
	return cvctc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChimeVoiceConnectorTerminationCredentials].
func (cvctc *ChimeVoiceConnectorTerminationCredentials) LifecycleManagement() *terra.Lifecycle {
	return cvctc.Lifecycle
}

// Attributes returns the attributes for [ChimeVoiceConnectorTerminationCredentials].
func (cvctc *ChimeVoiceConnectorTerminationCredentials) Attributes() chimeVoiceConnectorTerminationCredentialsAttributes {
	return chimeVoiceConnectorTerminationCredentialsAttributes{ref: terra.ReferenceResource(cvctc)}
}

// ImportState imports the given attribute values into [ChimeVoiceConnectorTerminationCredentials]'s state.
func (cvctc *ChimeVoiceConnectorTerminationCredentials) ImportState(av io.Reader) error {
	cvctc.state = &chimeVoiceConnectorTerminationCredentialsState{}
	if err := json.NewDecoder(av).Decode(cvctc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cvctc.Type(), cvctc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChimeVoiceConnectorTerminationCredentials] has state.
func (cvctc *ChimeVoiceConnectorTerminationCredentials) State() (*chimeVoiceConnectorTerminationCredentialsState, bool) {
	return cvctc.state, cvctc.state != nil
}

// StateMust returns the state for [ChimeVoiceConnectorTerminationCredentials]. Panics if the state is nil.
func (cvctc *ChimeVoiceConnectorTerminationCredentials) StateMust() *chimeVoiceConnectorTerminationCredentialsState {
	if cvctc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cvctc.Type(), cvctc.LocalName()))
	}
	return cvctc.state
}

// ChimeVoiceConnectorTerminationCredentialsArgs contains the configurations for aws_chime_voice_connector_termination_credentials.
type ChimeVoiceConnectorTerminationCredentialsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VoiceConnectorId: string, required
	VoiceConnectorId terra.StringValue `hcl:"voice_connector_id,attr" validate:"required"`
	// Credentials: min=1,max=10
	Credentials []chimevoiceconnectorterminationcredentials.Credentials `hcl:"credentials,block" validate:"min=1,max=10"`
}
type chimeVoiceConnectorTerminationCredentialsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_chime_voice_connector_termination_credentials.
func (cvctc chimeVoiceConnectorTerminationCredentialsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvctc.ref.Append("id"))
}

// VoiceConnectorId returns a reference to field voice_connector_id of aws_chime_voice_connector_termination_credentials.
func (cvctc chimeVoiceConnectorTerminationCredentialsAttributes) VoiceConnectorId() terra.StringValue {
	return terra.ReferenceAsString(cvctc.ref.Append("voice_connector_id"))
}

func (cvctc chimeVoiceConnectorTerminationCredentialsAttributes) Credentials() terra.SetValue[chimevoiceconnectorterminationcredentials.CredentialsAttributes] {
	return terra.ReferenceAsSet[chimevoiceconnectorterminationcredentials.CredentialsAttributes](cvctc.ref.Append("credentials"))
}

type chimeVoiceConnectorTerminationCredentialsState struct {
	Id               string                                                       `json:"id"`
	VoiceConnectorId string                                                       `json:"voice_connector_id"`
	Credentials      []chimevoiceconnectorterminationcredentials.CredentialsState `json:"credentials"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	chimevoiceconnectorgroup "github.com/golingon/terraproviders/aws/5.0.1/chimevoiceconnectorgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChimeVoiceConnectorGroup creates a new instance of [ChimeVoiceConnectorGroup].
func NewChimeVoiceConnectorGroup(name string, args ChimeVoiceConnectorGroupArgs) *ChimeVoiceConnectorGroup {
	return &ChimeVoiceConnectorGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChimeVoiceConnectorGroup)(nil)

// ChimeVoiceConnectorGroup represents the Terraform resource aws_chime_voice_connector_group.
type ChimeVoiceConnectorGroup struct {
	Name      string
	Args      ChimeVoiceConnectorGroupArgs
	state     *chimeVoiceConnectorGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChimeVoiceConnectorGroup].
func (cvcg *ChimeVoiceConnectorGroup) Type() string {
	return "aws_chime_voice_connector_group"
}

// LocalName returns the local name for [ChimeVoiceConnectorGroup].
func (cvcg *ChimeVoiceConnectorGroup) LocalName() string {
	return cvcg.Name
}

// Configuration returns the configuration (args) for [ChimeVoiceConnectorGroup].
func (cvcg *ChimeVoiceConnectorGroup) Configuration() interface{} {
	return cvcg.Args
}

// DependOn is used for other resources to depend on [ChimeVoiceConnectorGroup].
func (cvcg *ChimeVoiceConnectorGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cvcg)
}

// Dependencies returns the list of resources [ChimeVoiceConnectorGroup] depends_on.
func (cvcg *ChimeVoiceConnectorGroup) Dependencies() terra.Dependencies {
	return cvcg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChimeVoiceConnectorGroup].
func (cvcg *ChimeVoiceConnectorGroup) LifecycleManagement() *terra.Lifecycle {
	return cvcg.Lifecycle
}

// Attributes returns the attributes for [ChimeVoiceConnectorGroup].
func (cvcg *ChimeVoiceConnectorGroup) Attributes() chimeVoiceConnectorGroupAttributes {
	return chimeVoiceConnectorGroupAttributes{ref: terra.ReferenceResource(cvcg)}
}

// ImportState imports the given attribute values into [ChimeVoiceConnectorGroup]'s state.
func (cvcg *ChimeVoiceConnectorGroup) ImportState(av io.Reader) error {
	cvcg.state = &chimeVoiceConnectorGroupState{}
	if err := json.NewDecoder(av).Decode(cvcg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cvcg.Type(), cvcg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChimeVoiceConnectorGroup] has state.
func (cvcg *ChimeVoiceConnectorGroup) State() (*chimeVoiceConnectorGroupState, bool) {
	return cvcg.state, cvcg.state != nil
}

// StateMust returns the state for [ChimeVoiceConnectorGroup]. Panics if the state is nil.
func (cvcg *ChimeVoiceConnectorGroup) StateMust() *chimeVoiceConnectorGroupState {
	if cvcg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cvcg.Type(), cvcg.LocalName()))
	}
	return cvcg.state
}

// ChimeVoiceConnectorGroupArgs contains the configurations for aws_chime_voice_connector_group.
type ChimeVoiceConnectorGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Connector: min=0,max=3
	Connector []chimevoiceconnectorgroup.Connector `hcl:"connector,block" validate:"min=0,max=3"`
}
type chimeVoiceConnectorGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_chime_voice_connector_group.
func (cvcg chimeVoiceConnectorGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvcg.ref.Append("id"))
}

// Name returns a reference to field name of aws_chime_voice_connector_group.
func (cvcg chimeVoiceConnectorGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cvcg.ref.Append("name"))
}

func (cvcg chimeVoiceConnectorGroupAttributes) Connector() terra.SetValue[chimevoiceconnectorgroup.ConnectorAttributes] {
	return terra.ReferenceAsSet[chimevoiceconnectorgroup.ConnectorAttributes](cvcg.ref.Append("connector"))
}

type chimeVoiceConnectorGroupState struct {
	Id        string                                    `json:"id"`
	Name      string                                    `json:"name"`
	Connector []chimevoiceconnectorgroup.ConnectorState `json:"connector"`
}

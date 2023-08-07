// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChimeVoiceConnector creates a new instance of [ChimeVoiceConnector].
func NewChimeVoiceConnector(name string, args ChimeVoiceConnectorArgs) *ChimeVoiceConnector {
	return &ChimeVoiceConnector{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChimeVoiceConnector)(nil)

// ChimeVoiceConnector represents the Terraform resource aws_chime_voice_connector.
type ChimeVoiceConnector struct {
	Name      string
	Args      ChimeVoiceConnectorArgs
	state     *chimeVoiceConnectorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChimeVoiceConnector].
func (cvc *ChimeVoiceConnector) Type() string {
	return "aws_chime_voice_connector"
}

// LocalName returns the local name for [ChimeVoiceConnector].
func (cvc *ChimeVoiceConnector) LocalName() string {
	return cvc.Name
}

// Configuration returns the configuration (args) for [ChimeVoiceConnector].
func (cvc *ChimeVoiceConnector) Configuration() interface{} {
	return cvc.Args
}

// DependOn is used for other resources to depend on [ChimeVoiceConnector].
func (cvc *ChimeVoiceConnector) DependOn() terra.Reference {
	return terra.ReferenceResource(cvc)
}

// Dependencies returns the list of resources [ChimeVoiceConnector] depends_on.
func (cvc *ChimeVoiceConnector) Dependencies() terra.Dependencies {
	return cvc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChimeVoiceConnector].
func (cvc *ChimeVoiceConnector) LifecycleManagement() *terra.Lifecycle {
	return cvc.Lifecycle
}

// Attributes returns the attributes for [ChimeVoiceConnector].
func (cvc *ChimeVoiceConnector) Attributes() chimeVoiceConnectorAttributes {
	return chimeVoiceConnectorAttributes{ref: terra.ReferenceResource(cvc)}
}

// ImportState imports the given attribute values into [ChimeVoiceConnector]'s state.
func (cvc *ChimeVoiceConnector) ImportState(av io.Reader) error {
	cvc.state = &chimeVoiceConnectorState{}
	if err := json.NewDecoder(av).Decode(cvc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cvc.Type(), cvc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChimeVoiceConnector] has state.
func (cvc *ChimeVoiceConnector) State() (*chimeVoiceConnectorState, bool) {
	return cvc.state, cvc.state != nil
}

// StateMust returns the state for [ChimeVoiceConnector]. Panics if the state is nil.
func (cvc *ChimeVoiceConnector) StateMust() *chimeVoiceConnectorState {
	if cvc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cvc.Type(), cvc.LocalName()))
	}
	return cvc.state
}

// ChimeVoiceConnectorArgs contains the configurations for aws_chime_voice_connector.
type ChimeVoiceConnectorArgs struct {
	// AwsRegion: string, optional
	AwsRegion terra.StringValue `hcl:"aws_region,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RequireEncryption: bool, required
	RequireEncryption terra.BoolValue `hcl:"require_encryption,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type chimeVoiceConnectorAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_chime_voice_connector.
func (cvc chimeVoiceConnectorAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cvc.ref.Append("arn"))
}

// AwsRegion returns a reference to field aws_region of aws_chime_voice_connector.
func (cvc chimeVoiceConnectorAttributes) AwsRegion() terra.StringValue {
	return terra.ReferenceAsString(cvc.ref.Append("aws_region"))
}

// Id returns a reference to field id of aws_chime_voice_connector.
func (cvc chimeVoiceConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvc.ref.Append("id"))
}

// Name returns a reference to field name of aws_chime_voice_connector.
func (cvc chimeVoiceConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cvc.ref.Append("name"))
}

// OutboundHostName returns a reference to field outbound_host_name of aws_chime_voice_connector.
func (cvc chimeVoiceConnectorAttributes) OutboundHostName() terra.StringValue {
	return terra.ReferenceAsString(cvc.ref.Append("outbound_host_name"))
}

// RequireEncryption returns a reference to field require_encryption of aws_chime_voice_connector.
func (cvc chimeVoiceConnectorAttributes) RequireEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(cvc.ref.Append("require_encryption"))
}

// Tags returns a reference to field tags of aws_chime_voice_connector.
func (cvc chimeVoiceConnectorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cvc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_chime_voice_connector.
func (cvc chimeVoiceConnectorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cvc.ref.Append("tags_all"))
}

type chimeVoiceConnectorState struct {
	Arn               string            `json:"arn"`
	AwsRegion         string            `json:"aws_region"`
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	OutboundHostName  string            `json:"outbound_host_name"`
	RequireEncryption bool              `json:"require_encryption"`
	Tags              map[string]string `json:"tags"`
	TagsAll           map[string]string `json:"tags_all"`
}

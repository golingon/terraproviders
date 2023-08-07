// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMqConfiguration creates a new instance of [MqConfiguration].
func NewMqConfiguration(name string, args MqConfigurationArgs) *MqConfiguration {
	return &MqConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MqConfiguration)(nil)

// MqConfiguration represents the Terraform resource aws_mq_configuration.
type MqConfiguration struct {
	Name      string
	Args      MqConfigurationArgs
	state     *mqConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MqConfiguration].
func (mc *MqConfiguration) Type() string {
	return "aws_mq_configuration"
}

// LocalName returns the local name for [MqConfiguration].
func (mc *MqConfiguration) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [MqConfiguration].
func (mc *MqConfiguration) Configuration() interface{} {
	return mc.Args
}

// DependOn is used for other resources to depend on [MqConfiguration].
func (mc *MqConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(mc)
}

// Dependencies returns the list of resources [MqConfiguration] depends_on.
func (mc *MqConfiguration) Dependencies() terra.Dependencies {
	return mc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MqConfiguration].
func (mc *MqConfiguration) LifecycleManagement() *terra.Lifecycle {
	return mc.Lifecycle
}

// Attributes returns the attributes for [MqConfiguration].
func (mc *MqConfiguration) Attributes() mqConfigurationAttributes {
	return mqConfigurationAttributes{ref: terra.ReferenceResource(mc)}
}

// ImportState imports the given attribute values into [MqConfiguration]'s state.
func (mc *MqConfiguration) ImportState(av io.Reader) error {
	mc.state = &mqConfigurationState{}
	if err := json.NewDecoder(av).Decode(mc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mc.Type(), mc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MqConfiguration] has state.
func (mc *MqConfiguration) State() (*mqConfigurationState, bool) {
	return mc.state, mc.state != nil
}

// StateMust returns the state for [MqConfiguration]. Panics if the state is nil.
func (mc *MqConfiguration) StateMust() *mqConfigurationState {
	if mc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mc.Type(), mc.LocalName()))
	}
	return mc.state
}

// MqConfigurationArgs contains the configurations for aws_mq_configuration.
type MqConfigurationArgs struct {
	// AuthenticationStrategy: string, optional
	AuthenticationStrategy terra.StringValue `hcl:"authentication_strategy,attr"`
	// Data: string, required
	Data terra.StringValue `hcl:"data,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EngineType: string, required
	EngineType terra.StringValue `hcl:"engine_type,attr" validate:"required"`
	// EngineVersion: string, required
	EngineVersion terra.StringValue `hcl:"engine_version,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type mqConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_mq_configuration.
func (mc mqConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("arn"))
}

// AuthenticationStrategy returns a reference to field authentication_strategy of aws_mq_configuration.
func (mc mqConfigurationAttributes) AuthenticationStrategy() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("authentication_strategy"))
}

// Data returns a reference to field data of aws_mq_configuration.
func (mc mqConfigurationAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("data"))
}

// Description returns a reference to field description of aws_mq_configuration.
func (mc mqConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("description"))
}

// EngineType returns a reference to field engine_type of aws_mq_configuration.
func (mc mqConfigurationAttributes) EngineType() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("engine_type"))
}

// EngineVersion returns a reference to field engine_version of aws_mq_configuration.
func (mc mqConfigurationAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_mq_configuration.
func (mc mqConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// LatestRevision returns a reference to field latest_revision of aws_mq_configuration.
func (mc mqConfigurationAttributes) LatestRevision() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("latest_revision"))
}

// Name returns a reference to field name of aws_mq_configuration.
func (mc mqConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_mq_configuration.
func (mc mqConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_mq_configuration.
func (mc mqConfigurationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags_all"))
}

type mqConfigurationState struct {
	Arn                    string            `json:"arn"`
	AuthenticationStrategy string            `json:"authentication_strategy"`
	Data                   string            `json:"data"`
	Description            string            `json:"description"`
	EngineType             string            `json:"engine_type"`
	EngineVersion          string            `json:"engine_version"`
	Id                     string            `json:"id"`
	LatestRevision         float64           `json:"latest_revision"`
	Name                   string            `json:"name"`
	Tags                   map[string]string `json:"tags"`
	TagsAll                map[string]string `json:"tags_all"`
}

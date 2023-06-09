// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMskConfiguration creates a new instance of [MskConfiguration].
func NewMskConfiguration(name string, args MskConfigurationArgs) *MskConfiguration {
	return &MskConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MskConfiguration)(nil)

// MskConfiguration represents the Terraform resource aws_msk_configuration.
type MskConfiguration struct {
	Name      string
	Args      MskConfigurationArgs
	state     *mskConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MskConfiguration].
func (mc *MskConfiguration) Type() string {
	return "aws_msk_configuration"
}

// LocalName returns the local name for [MskConfiguration].
func (mc *MskConfiguration) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [MskConfiguration].
func (mc *MskConfiguration) Configuration() interface{} {
	return mc.Args
}

// DependOn is used for other resources to depend on [MskConfiguration].
func (mc *MskConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(mc)
}

// Dependencies returns the list of resources [MskConfiguration] depends_on.
func (mc *MskConfiguration) Dependencies() terra.Dependencies {
	return mc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MskConfiguration].
func (mc *MskConfiguration) LifecycleManagement() *terra.Lifecycle {
	return mc.Lifecycle
}

// Attributes returns the attributes for [MskConfiguration].
func (mc *MskConfiguration) Attributes() mskConfigurationAttributes {
	return mskConfigurationAttributes{ref: terra.ReferenceResource(mc)}
}

// ImportState imports the given attribute values into [MskConfiguration]'s state.
func (mc *MskConfiguration) ImportState(av io.Reader) error {
	mc.state = &mskConfigurationState{}
	if err := json.NewDecoder(av).Decode(mc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mc.Type(), mc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MskConfiguration] has state.
func (mc *MskConfiguration) State() (*mskConfigurationState, bool) {
	return mc.state, mc.state != nil
}

// StateMust returns the state for [MskConfiguration]. Panics if the state is nil.
func (mc *MskConfiguration) StateMust() *mskConfigurationState {
	if mc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mc.Type(), mc.LocalName()))
	}
	return mc.state
}

// MskConfigurationArgs contains the configurations for aws_msk_configuration.
type MskConfigurationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KafkaVersions: set of string, optional
	KafkaVersions terra.SetValue[terra.StringValue] `hcl:"kafka_versions,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServerProperties: string, required
	ServerProperties terra.StringValue `hcl:"server_properties,attr" validate:"required"`
}
type mskConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_msk_configuration.
func (mc mskConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("arn"))
}

// Description returns a reference to field description of aws_msk_configuration.
func (mc mskConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("description"))
}

// Id returns a reference to field id of aws_msk_configuration.
func (mc mskConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// KafkaVersions returns a reference to field kafka_versions of aws_msk_configuration.
func (mc mskConfigurationAttributes) KafkaVersions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mc.ref.Append("kafka_versions"))
}

// LatestRevision returns a reference to field latest_revision of aws_msk_configuration.
func (mc mskConfigurationAttributes) LatestRevision() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("latest_revision"))
}

// Name returns a reference to field name of aws_msk_configuration.
func (mc mskConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("name"))
}

// ServerProperties returns a reference to field server_properties of aws_msk_configuration.
func (mc mskConfigurationAttributes) ServerProperties() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("server_properties"))
}

type mskConfigurationState struct {
	Arn              string   `json:"arn"`
	Description      string   `json:"description"`
	Id               string   `json:"id"`
	KafkaVersions    []string `json:"kafka_versions"`
	LatestRevision   float64  `json:"latest_revision"`
	Name             string   `json:"name"`
	ServerProperties string   `json:"server_properties"`
}

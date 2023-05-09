// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	imagebuilderinfrastructureconfiguration "github.com/golingon/terraproviders/aws/4.66.1/imagebuilderinfrastructureconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewImagebuilderInfrastructureConfiguration creates a new instance of [ImagebuilderInfrastructureConfiguration].
func NewImagebuilderInfrastructureConfiguration(name string, args ImagebuilderInfrastructureConfigurationArgs) *ImagebuilderInfrastructureConfiguration {
	return &ImagebuilderInfrastructureConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ImagebuilderInfrastructureConfiguration)(nil)

// ImagebuilderInfrastructureConfiguration represents the Terraform resource aws_imagebuilder_infrastructure_configuration.
type ImagebuilderInfrastructureConfiguration struct {
	Name      string
	Args      ImagebuilderInfrastructureConfigurationArgs
	state     *imagebuilderInfrastructureConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ImagebuilderInfrastructureConfiguration].
func (iic *ImagebuilderInfrastructureConfiguration) Type() string {
	return "aws_imagebuilder_infrastructure_configuration"
}

// LocalName returns the local name for [ImagebuilderInfrastructureConfiguration].
func (iic *ImagebuilderInfrastructureConfiguration) LocalName() string {
	return iic.Name
}

// Configuration returns the configuration (args) for [ImagebuilderInfrastructureConfiguration].
func (iic *ImagebuilderInfrastructureConfiguration) Configuration() interface{} {
	return iic.Args
}

// DependOn is used for other resources to depend on [ImagebuilderInfrastructureConfiguration].
func (iic *ImagebuilderInfrastructureConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(iic)
}

// Dependencies returns the list of resources [ImagebuilderInfrastructureConfiguration] depends_on.
func (iic *ImagebuilderInfrastructureConfiguration) Dependencies() terra.Dependencies {
	return iic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ImagebuilderInfrastructureConfiguration].
func (iic *ImagebuilderInfrastructureConfiguration) LifecycleManagement() *terra.Lifecycle {
	return iic.Lifecycle
}

// Attributes returns the attributes for [ImagebuilderInfrastructureConfiguration].
func (iic *ImagebuilderInfrastructureConfiguration) Attributes() imagebuilderInfrastructureConfigurationAttributes {
	return imagebuilderInfrastructureConfigurationAttributes{ref: terra.ReferenceResource(iic)}
}

// ImportState imports the given attribute values into [ImagebuilderInfrastructureConfiguration]'s state.
func (iic *ImagebuilderInfrastructureConfiguration) ImportState(av io.Reader) error {
	iic.state = &imagebuilderInfrastructureConfigurationState{}
	if err := json.NewDecoder(av).Decode(iic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iic.Type(), iic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ImagebuilderInfrastructureConfiguration] has state.
func (iic *ImagebuilderInfrastructureConfiguration) State() (*imagebuilderInfrastructureConfigurationState, bool) {
	return iic.state, iic.state != nil
}

// StateMust returns the state for [ImagebuilderInfrastructureConfiguration]. Panics if the state is nil.
func (iic *ImagebuilderInfrastructureConfiguration) StateMust() *imagebuilderInfrastructureConfigurationState {
	if iic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iic.Type(), iic.LocalName()))
	}
	return iic.state
}

// ImagebuilderInfrastructureConfigurationArgs contains the configurations for aws_imagebuilder_infrastructure_configuration.
type ImagebuilderInfrastructureConfigurationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceProfileName: string, required
	InstanceProfileName terra.StringValue `hcl:"instance_profile_name,attr" validate:"required"`
	// InstanceTypes: set of string, optional
	InstanceTypes terra.SetValue[terra.StringValue] `hcl:"instance_types,attr"`
	// KeyPair: string, optional
	KeyPair terra.StringValue `hcl:"key_pair,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceTags: map of string, optional
	ResourceTags terra.MapValue[terra.StringValue] `hcl:"resource_tags,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SnsTopicArn: string, optional
	SnsTopicArn terra.StringValue `hcl:"sns_topic_arn,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TerminateInstanceOnFailure: bool, optional
	TerminateInstanceOnFailure terra.BoolValue `hcl:"terminate_instance_on_failure,attr"`
	// InstanceMetadataOptions: optional
	InstanceMetadataOptions *imagebuilderinfrastructureconfiguration.InstanceMetadataOptions `hcl:"instance_metadata_options,block"`
	// Logging: optional
	Logging *imagebuilderinfrastructureconfiguration.Logging `hcl:"logging,block"`
}
type imagebuilderInfrastructureConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("arn"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("date_created"))
}

// DateUpdated returns a reference to field date_updated of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) DateUpdated() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("date_updated"))
}

// Description returns a reference to field description of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("description"))
}

// Id returns a reference to field id of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("id"))
}

// InstanceProfileName returns a reference to field instance_profile_name of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) InstanceProfileName() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("instance_profile_name"))
}

// InstanceTypes returns a reference to field instance_types of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) InstanceTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iic.ref.Append("instance_types"))
}

// KeyPair returns a reference to field key_pair of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) KeyPair() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("key_pair"))
}

// Name returns a reference to field name of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("name"))
}

// ResourceTags returns a reference to field resource_tags of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) ResourceTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iic.ref.Append("resource_tags"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iic.ref.Append("security_group_ids"))
}

// SnsTopicArn returns a reference to field sns_topic_arn of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) SnsTopicArn() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("sns_topic_arn"))
}

// SubnetId returns a reference to field subnet_id of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iic.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iic.ref.Append("tags_all"))
}

// TerminateInstanceOnFailure returns a reference to field terminate_instance_on_failure of aws_imagebuilder_infrastructure_configuration.
func (iic imagebuilderInfrastructureConfigurationAttributes) TerminateInstanceOnFailure() terra.BoolValue {
	return terra.ReferenceAsBool(iic.ref.Append("terminate_instance_on_failure"))
}

func (iic imagebuilderInfrastructureConfigurationAttributes) InstanceMetadataOptions() terra.ListValue[imagebuilderinfrastructureconfiguration.InstanceMetadataOptionsAttributes] {
	return terra.ReferenceAsList[imagebuilderinfrastructureconfiguration.InstanceMetadataOptionsAttributes](iic.ref.Append("instance_metadata_options"))
}

func (iic imagebuilderInfrastructureConfigurationAttributes) Logging() terra.ListValue[imagebuilderinfrastructureconfiguration.LoggingAttributes] {
	return terra.ReferenceAsList[imagebuilderinfrastructureconfiguration.LoggingAttributes](iic.ref.Append("logging"))
}

type imagebuilderInfrastructureConfigurationState struct {
	Arn                        string                                                                 `json:"arn"`
	DateCreated                string                                                                 `json:"date_created"`
	DateUpdated                string                                                                 `json:"date_updated"`
	Description                string                                                                 `json:"description"`
	Id                         string                                                                 `json:"id"`
	InstanceProfileName        string                                                                 `json:"instance_profile_name"`
	InstanceTypes              []string                                                               `json:"instance_types"`
	KeyPair                    string                                                                 `json:"key_pair"`
	Name                       string                                                                 `json:"name"`
	ResourceTags               map[string]string                                                      `json:"resource_tags"`
	SecurityGroupIds           []string                                                               `json:"security_group_ids"`
	SnsTopicArn                string                                                                 `json:"sns_topic_arn"`
	SubnetId                   string                                                                 `json:"subnet_id"`
	Tags                       map[string]string                                                      `json:"tags"`
	TagsAll                    map[string]string                                                      `json:"tags_all"`
	TerminateInstanceOnFailure bool                                                                   `json:"terminate_instance_on_failure"`
	InstanceMetadataOptions    []imagebuilderinfrastructureconfiguration.InstanceMetadataOptionsState `json:"instance_metadata_options"`
	Logging                    []imagebuilderinfrastructureconfiguration.LoggingState                 `json:"logging"`
}

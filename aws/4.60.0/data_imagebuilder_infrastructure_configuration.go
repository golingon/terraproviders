// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataimagebuilderinfrastructureconfiguration "github.com/golingon/terraproviders/aws/4.60.0/dataimagebuilderinfrastructureconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataImagebuilderInfrastructureConfiguration creates a new instance of [DataImagebuilderInfrastructureConfiguration].
func NewDataImagebuilderInfrastructureConfiguration(name string, args DataImagebuilderInfrastructureConfigurationArgs) *DataImagebuilderInfrastructureConfiguration {
	return &DataImagebuilderInfrastructureConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataImagebuilderInfrastructureConfiguration)(nil)

// DataImagebuilderInfrastructureConfiguration represents the Terraform data resource aws_imagebuilder_infrastructure_configuration.
type DataImagebuilderInfrastructureConfiguration struct {
	Name string
	Args DataImagebuilderInfrastructureConfigurationArgs
}

// DataSource returns the Terraform object type for [DataImagebuilderInfrastructureConfiguration].
func (iic *DataImagebuilderInfrastructureConfiguration) DataSource() string {
	return "aws_imagebuilder_infrastructure_configuration"
}

// LocalName returns the local name for [DataImagebuilderInfrastructureConfiguration].
func (iic *DataImagebuilderInfrastructureConfiguration) LocalName() string {
	return iic.Name
}

// Configuration returns the configuration (args) for [DataImagebuilderInfrastructureConfiguration].
func (iic *DataImagebuilderInfrastructureConfiguration) Configuration() interface{} {
	return iic.Args
}

// Attributes returns the attributes for [DataImagebuilderInfrastructureConfiguration].
func (iic *DataImagebuilderInfrastructureConfiguration) Attributes() dataImagebuilderInfrastructureConfigurationAttributes {
	return dataImagebuilderInfrastructureConfigurationAttributes{ref: terra.ReferenceDataResource(iic)}
}

// DataImagebuilderInfrastructureConfigurationArgs contains the configurations for aws_imagebuilder_infrastructure_configuration.
type DataImagebuilderInfrastructureConfigurationArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceTags: map of string, optional
	ResourceTags terra.MapValue[terra.StringValue] `hcl:"resource_tags,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// InstanceMetadataOptions: min=0
	InstanceMetadataOptions []dataimagebuilderinfrastructureconfiguration.InstanceMetadataOptions `hcl:"instance_metadata_options,block" validate:"min=0"`
	// Logging: min=0
	Logging []dataimagebuilderinfrastructureconfiguration.Logging `hcl:"logging,block" validate:"min=0"`
}
type dataImagebuilderInfrastructureConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("arn"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("date_created"))
}

// DateUpdated returns a reference to field date_updated of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) DateUpdated() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("date_updated"))
}

// Description returns a reference to field description of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("description"))
}

// Id returns a reference to field id of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("id"))
}

// InstanceProfileName returns a reference to field instance_profile_name of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) InstanceProfileName() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("instance_profile_name"))
}

// InstanceTypes returns a reference to field instance_types of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) InstanceTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iic.ref.Append("instance_types"))
}

// KeyPair returns a reference to field key_pair of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) KeyPair() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("key_pair"))
}

// Name returns a reference to field name of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("name"))
}

// ResourceTags returns a reference to field resource_tags of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) ResourceTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iic.ref.Append("resource_tags"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iic.ref.Append("security_group_ids"))
}

// SnsTopicArn returns a reference to field sns_topic_arn of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) SnsTopicArn() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("sns_topic_arn"))
}

// SubnetId returns a reference to field subnet_id of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iic.ref.Append("tags"))
}

// TerminateInstanceOnFailure returns a reference to field terminate_instance_on_failure of aws_imagebuilder_infrastructure_configuration.
func (iic dataImagebuilderInfrastructureConfigurationAttributes) TerminateInstanceOnFailure() terra.BoolValue {
	return terra.ReferenceAsBool(iic.ref.Append("terminate_instance_on_failure"))
}

func (iic dataImagebuilderInfrastructureConfigurationAttributes) InstanceMetadataOptions() terra.ListValue[dataimagebuilderinfrastructureconfiguration.InstanceMetadataOptionsAttributes] {
	return terra.ReferenceAsList[dataimagebuilderinfrastructureconfiguration.InstanceMetadataOptionsAttributes](iic.ref.Append("instance_metadata_options"))
}

func (iic dataImagebuilderInfrastructureConfigurationAttributes) Logging() terra.ListValue[dataimagebuilderinfrastructureconfiguration.LoggingAttributes] {
	return terra.ReferenceAsList[dataimagebuilderinfrastructureconfiguration.LoggingAttributes](iic.ref.Append("logging"))
}

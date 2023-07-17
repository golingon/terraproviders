// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataappconfigconfigurationprofile "github.com/golingon/terraproviders/aws/5.7.0/dataappconfigconfigurationprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppconfigConfigurationProfile creates a new instance of [DataAppconfigConfigurationProfile].
func NewDataAppconfigConfigurationProfile(name string, args DataAppconfigConfigurationProfileArgs) *DataAppconfigConfigurationProfile {
	return &DataAppconfigConfigurationProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppconfigConfigurationProfile)(nil)

// DataAppconfigConfigurationProfile represents the Terraform data resource aws_appconfig_configuration_profile.
type DataAppconfigConfigurationProfile struct {
	Name string
	Args DataAppconfigConfigurationProfileArgs
}

// DataSource returns the Terraform object type for [DataAppconfigConfigurationProfile].
func (acp *DataAppconfigConfigurationProfile) DataSource() string {
	return "aws_appconfig_configuration_profile"
}

// LocalName returns the local name for [DataAppconfigConfigurationProfile].
func (acp *DataAppconfigConfigurationProfile) LocalName() string {
	return acp.Name
}

// Configuration returns the configuration (args) for [DataAppconfigConfigurationProfile].
func (acp *DataAppconfigConfigurationProfile) Configuration() interface{} {
	return acp.Args
}

// Attributes returns the attributes for [DataAppconfigConfigurationProfile].
func (acp *DataAppconfigConfigurationProfile) Attributes() dataAppconfigConfigurationProfileAttributes {
	return dataAppconfigConfigurationProfileAttributes{ref: terra.ReferenceDataResource(acp)}
}

// DataAppconfigConfigurationProfileArgs contains the configurations for aws_appconfig_configuration_profile.
type DataAppconfigConfigurationProfileArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// ConfigurationProfileId: string, required
	ConfigurationProfileId terra.StringValue `hcl:"configuration_profile_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Validator: min=0
	Validator []dataappconfigconfigurationprofile.Validator `hcl:"validator,block" validate:"min=0"`
}
type dataAppconfigConfigurationProfileAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of aws_appconfig_configuration_profile.
func (acp dataAppconfigConfigurationProfileAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("application_id"))
}

// Arn returns a reference to field arn of aws_appconfig_configuration_profile.
func (acp dataAppconfigConfigurationProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("arn"))
}

// ConfigurationProfileId returns a reference to field configuration_profile_id of aws_appconfig_configuration_profile.
func (acp dataAppconfigConfigurationProfileAttributes) ConfigurationProfileId() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("configuration_profile_id"))
}

// Description returns a reference to field description of aws_appconfig_configuration_profile.
func (acp dataAppconfigConfigurationProfileAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("description"))
}

// Id returns a reference to field id of aws_appconfig_configuration_profile.
func (acp dataAppconfigConfigurationProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("id"))
}

// LocationUri returns a reference to field location_uri of aws_appconfig_configuration_profile.
func (acp dataAppconfigConfigurationProfileAttributes) LocationUri() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("location_uri"))
}

// Name returns a reference to field name of aws_appconfig_configuration_profile.
func (acp dataAppconfigConfigurationProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("name"))
}

// RetrievalRoleArn returns a reference to field retrieval_role_arn of aws_appconfig_configuration_profile.
func (acp dataAppconfigConfigurationProfileAttributes) RetrievalRoleArn() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("retrieval_role_arn"))
}

// Tags returns a reference to field tags of aws_appconfig_configuration_profile.
func (acp dataAppconfigConfigurationProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acp.ref.Append("tags"))
}

// Type returns a reference to field type of aws_appconfig_configuration_profile.
func (acp dataAppconfigConfigurationProfileAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("type"))
}

func (acp dataAppconfigConfigurationProfileAttributes) Validator() terra.SetValue[dataappconfigconfigurationprofile.ValidatorAttributes] {
	return terra.ReferenceAsSet[dataappconfigconfigurationprofile.ValidatorAttributes](acp.ref.Append("validator"))
}
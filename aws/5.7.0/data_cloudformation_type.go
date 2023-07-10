// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datacloudformationtype "github.com/golingon/terraproviders/aws/5.7.0/datacloudformationtype"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudformationType creates a new instance of [DataCloudformationType].
func NewDataCloudformationType(name string, args DataCloudformationTypeArgs) *DataCloudformationType {
	return &DataCloudformationType{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudformationType)(nil)

// DataCloudformationType represents the Terraform data resource aws_cloudformation_type.
type DataCloudformationType struct {
	Name string
	Args DataCloudformationTypeArgs
}

// DataSource returns the Terraform object type for [DataCloudformationType].
func (ct *DataCloudformationType) DataSource() string {
	return "aws_cloudformation_type"
}

// LocalName returns the local name for [DataCloudformationType].
func (ct *DataCloudformationType) LocalName() string {
	return ct.Name
}

// Configuration returns the configuration (args) for [DataCloudformationType].
func (ct *DataCloudformationType) Configuration() interface{} {
	return ct.Args
}

// Attributes returns the attributes for [DataCloudformationType].
func (ct *DataCloudformationType) Attributes() dataCloudformationTypeAttributes {
	return dataCloudformationTypeAttributes{ref: terra.ReferenceDataResource(ct)}
}

// DataCloudformationTypeArgs contains the configurations for aws_cloudformation_type.
type DataCloudformationTypeArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// TypeName: string, optional
	TypeName terra.StringValue `hcl:"type_name,attr"`
	// VersionId: string, optional
	VersionId terra.StringValue `hcl:"version_id,attr"`
	// LoggingConfig: min=0
	LoggingConfig []datacloudformationtype.LoggingConfig `hcl:"logging_config,block" validate:"min=0"`
}
type dataCloudformationTypeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("arn"))
}

// DefaultVersionId returns a reference to field default_version_id of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) DefaultVersionId() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("default_version_id"))
}

// DeprecatedStatus returns a reference to field deprecated_status of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) DeprecatedStatus() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("deprecated_status"))
}

// Description returns a reference to field description of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("description"))
}

// DocumentationUrl returns a reference to field documentation_url of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) DocumentationUrl() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("documentation_url"))
}

// ExecutionRoleArn returns a reference to field execution_role_arn of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) ExecutionRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("execution_role_arn"))
}

// Id returns a reference to field id of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("id"))
}

// IsDefaultVersion returns a reference to field is_default_version of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) IsDefaultVersion() terra.BoolValue {
	return terra.ReferenceAsBool(ct.ref.Append("is_default_version"))
}

// ProvisioningType returns a reference to field provisioning_type of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) ProvisioningType() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("provisioning_type"))
}

// Schema returns a reference to field schema of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) Schema() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("schema"))
}

// SourceUrl returns a reference to field source_url of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) SourceUrl() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("source_url"))
}

// Type returns a reference to field type of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("type"))
}

// TypeArn returns a reference to field type_arn of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) TypeArn() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("type_arn"))
}

// TypeName returns a reference to field type_name of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) TypeName() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("type_name"))
}

// VersionId returns a reference to field version_id of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("version_id"))
}

// Visibility returns a reference to field visibility of aws_cloudformation_type.
func (ct dataCloudformationTypeAttributes) Visibility() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("visibility"))
}

func (ct dataCloudformationTypeAttributes) LoggingConfig() terra.ListValue[datacloudformationtype.LoggingConfigAttributes] {
	return terra.ReferenceAsList[datacloudformationtype.LoggingConfigAttributes](ct.ref.Append("logging_config"))
}

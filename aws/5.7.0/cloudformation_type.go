// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudformationtype "github.com/golingon/terraproviders/aws/5.7.0/cloudformationtype"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudformationType creates a new instance of [CloudformationType].
func NewCloudformationType(name string, args CloudformationTypeArgs) *CloudformationType {
	return &CloudformationType{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudformationType)(nil)

// CloudformationType represents the Terraform resource aws_cloudformation_type.
type CloudformationType struct {
	Name      string
	Args      CloudformationTypeArgs
	state     *cloudformationTypeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudformationType].
func (ct *CloudformationType) Type() string {
	return "aws_cloudformation_type"
}

// LocalName returns the local name for [CloudformationType].
func (ct *CloudformationType) LocalName() string {
	return ct.Name
}

// Configuration returns the configuration (args) for [CloudformationType].
func (ct *CloudformationType) Configuration() interface{} {
	return ct.Args
}

// DependOn is used for other resources to depend on [CloudformationType].
func (ct *CloudformationType) DependOn() terra.Reference {
	return terra.ReferenceResource(ct)
}

// Dependencies returns the list of resources [CloudformationType] depends_on.
func (ct *CloudformationType) Dependencies() terra.Dependencies {
	return ct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudformationType].
func (ct *CloudformationType) LifecycleManagement() *terra.Lifecycle {
	return ct.Lifecycle
}

// Attributes returns the attributes for [CloudformationType].
func (ct *CloudformationType) Attributes() cloudformationTypeAttributes {
	return cloudformationTypeAttributes{ref: terra.ReferenceResource(ct)}
}

// ImportState imports the given attribute values into [CloudformationType]'s state.
func (ct *CloudformationType) ImportState(av io.Reader) error {
	ct.state = &cloudformationTypeState{}
	if err := json.NewDecoder(av).Decode(ct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ct.Type(), ct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudformationType] has state.
func (ct *CloudformationType) State() (*cloudformationTypeState, bool) {
	return ct.state, ct.state != nil
}

// StateMust returns the state for [CloudformationType]. Panics if the state is nil.
func (ct *CloudformationType) StateMust() *cloudformationTypeState {
	if ct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ct.Type(), ct.LocalName()))
	}
	return ct.state
}

// CloudformationTypeArgs contains the configurations for aws_cloudformation_type.
type CloudformationTypeArgs struct {
	// ExecutionRoleArn: string, optional
	ExecutionRoleArn terra.StringValue `hcl:"execution_role_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SchemaHandlerPackage: string, required
	SchemaHandlerPackage terra.StringValue `hcl:"schema_handler_package,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// TypeName: string, required
	TypeName terra.StringValue `hcl:"type_name,attr" validate:"required"`
	// LoggingConfig: optional
	LoggingConfig *cloudformationtype.LoggingConfig `hcl:"logging_config,block"`
}
type cloudformationTypeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("arn"))
}

// DefaultVersionId returns a reference to field default_version_id of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) DefaultVersionId() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("default_version_id"))
}

// DeprecatedStatus returns a reference to field deprecated_status of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) DeprecatedStatus() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("deprecated_status"))
}

// Description returns a reference to field description of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("description"))
}

// DocumentationUrl returns a reference to field documentation_url of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) DocumentationUrl() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("documentation_url"))
}

// ExecutionRoleArn returns a reference to field execution_role_arn of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) ExecutionRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("execution_role_arn"))
}

// Id returns a reference to field id of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("id"))
}

// IsDefaultVersion returns a reference to field is_default_version of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) IsDefaultVersion() terra.BoolValue {
	return terra.ReferenceAsBool(ct.ref.Append("is_default_version"))
}

// ProvisioningType returns a reference to field provisioning_type of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) ProvisioningType() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("provisioning_type"))
}

// Schema returns a reference to field schema of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) Schema() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("schema"))
}

// SchemaHandlerPackage returns a reference to field schema_handler_package of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) SchemaHandlerPackage() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("schema_handler_package"))
}

// SourceUrl returns a reference to field source_url of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) SourceUrl() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("source_url"))
}

// Type returns a reference to field type of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("type"))
}

// TypeArn returns a reference to field type_arn of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) TypeArn() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("type_arn"))
}

// TypeName returns a reference to field type_name of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) TypeName() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("type_name"))
}

// VersionId returns a reference to field version_id of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("version_id"))
}

// Visibility returns a reference to field visibility of aws_cloudformation_type.
func (ct cloudformationTypeAttributes) Visibility() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("visibility"))
}

func (ct cloudformationTypeAttributes) LoggingConfig() terra.ListValue[cloudformationtype.LoggingConfigAttributes] {
	return terra.ReferenceAsList[cloudformationtype.LoggingConfigAttributes](ct.ref.Append("logging_config"))
}

type cloudformationTypeState struct {
	Arn                  string                                  `json:"arn"`
	DefaultVersionId     string                                  `json:"default_version_id"`
	DeprecatedStatus     string                                  `json:"deprecated_status"`
	Description          string                                  `json:"description"`
	DocumentationUrl     string                                  `json:"documentation_url"`
	ExecutionRoleArn     string                                  `json:"execution_role_arn"`
	Id                   string                                  `json:"id"`
	IsDefaultVersion     bool                                    `json:"is_default_version"`
	ProvisioningType     string                                  `json:"provisioning_type"`
	Schema               string                                  `json:"schema"`
	SchemaHandlerPackage string                                  `json:"schema_handler_package"`
	SourceUrl            string                                  `json:"source_url"`
	Type                 string                                  `json:"type"`
	TypeArn              string                                  `json:"type_arn"`
	TypeName             string                                  `json:"type_name"`
	VersionId            string                                  `json:"version_id"`
	Visibility           string                                  `json:"visibility"`
	LoggingConfig        []cloudformationtype.LoggingConfigState `json:"logging_config"`
}

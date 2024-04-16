// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_sagemaker_domain

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_sagemaker_domain.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSagemakerDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asd *Resource) Type() string {
	return "aws_sagemaker_domain"
}

// LocalName returns the local name for [Resource].
func (asd *Resource) LocalName() string {
	return asd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asd *Resource) Configuration() interface{} {
	return asd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asd *Resource) Dependencies() terra.Dependencies {
	return asd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asd *Resource) LifecycleManagement() *terra.Lifecycle {
	return asd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asd *Resource) Attributes() awsSagemakerDomainAttributes {
	return awsSagemakerDomainAttributes{ref: terra.ReferenceResource(asd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asd *Resource) ImportState(state io.Reader) error {
	asd.state = &awsSagemakerDomainState{}
	if err := json.NewDecoder(state).Decode(asd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asd.Type(), asd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asd *Resource) State() (*awsSagemakerDomainState, bool) {
	return asd.state, asd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asd *Resource) StateMust() *awsSagemakerDomainState {
	if asd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asd.Type(), asd.LocalName()))
	}
	return asd.state
}

// Args contains the configurations for aws_sagemaker_domain.
type Args struct {
	// AppNetworkAccessType: string, optional
	AppNetworkAccessType terra.StringValue `hcl:"app_network_access_type,attr"`
	// AppSecurityGroupManagement: string, optional
	AppSecurityGroupManagement terra.StringValue `hcl:"app_security_group_management,attr"`
	// AuthMode: string, required
	AuthMode terra.StringValue `hcl:"auth_mode,attr" validate:"required"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// SubnetIds: set of string, required
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// DefaultSpaceSettings: optional
	DefaultSpaceSettings *DefaultSpaceSettings `hcl:"default_space_settings,block"`
	// DefaultUserSettings: required
	DefaultUserSettings *DefaultUserSettings `hcl:"default_user_settings,block" validate:"required"`
	// DomainSettings: optional
	DomainSettings *DomainSettings `hcl:"domain_settings,block"`
	// RetentionPolicy: optional
	RetentionPolicy *RetentionPolicy `hcl:"retention_policy,block"`
}

type awsSagemakerDomainAttributes struct {
	ref terra.Reference
}

// AppNetworkAccessType returns a reference to field app_network_access_type of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) AppNetworkAccessType() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("app_network_access_type"))
}

// AppSecurityGroupManagement returns a reference to field app_security_group_management of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) AppSecurityGroupManagement() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("app_security_group_management"))
}

// Arn returns a reference to field arn of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("arn"))
}

// AuthMode returns a reference to field auth_mode of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) AuthMode() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("auth_mode"))
}

// DomainName returns a reference to field domain_name of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("domain_name"))
}

// HomeEfsFileSystemId returns a reference to field home_efs_file_system_id of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) HomeEfsFileSystemId() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("home_efs_file_system_id"))
}

// Id returns a reference to field id of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("kms_key_id"))
}

// SecurityGroupIdForDomainBoundary returns a reference to field security_group_id_for_domain_boundary of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) SecurityGroupIdForDomainBoundary() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("security_group_id_for_domain_boundary"))
}

// SingleSignOnApplicationArn returns a reference to field single_sign_on_application_arn of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) SingleSignOnApplicationArn() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("single_sign_on_application_arn"))
}

// SingleSignOnManagedApplicationInstanceId returns a reference to field single_sign_on_managed_application_instance_id of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) SingleSignOnManagedApplicationInstanceId() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("single_sign_on_managed_application_instance_id"))
}

// SubnetIds returns a reference to field subnet_ids of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](asd.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asd.ref.Append("tags_all"))
}

// Url returns a reference to field url of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("url"))
}

// VpcId returns a reference to field vpc_id of aws_sagemaker_domain.
func (asd awsSagemakerDomainAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("vpc_id"))
}

func (asd awsSagemakerDomainAttributes) DefaultSpaceSettings() terra.ListValue[DefaultSpaceSettingsAttributes] {
	return terra.ReferenceAsList[DefaultSpaceSettingsAttributes](asd.ref.Append("default_space_settings"))
}

func (asd awsSagemakerDomainAttributes) DefaultUserSettings() terra.ListValue[DefaultUserSettingsAttributes] {
	return terra.ReferenceAsList[DefaultUserSettingsAttributes](asd.ref.Append("default_user_settings"))
}

func (asd awsSagemakerDomainAttributes) DomainSettings() terra.ListValue[DomainSettingsAttributes] {
	return terra.ReferenceAsList[DomainSettingsAttributes](asd.ref.Append("domain_settings"))
}

func (asd awsSagemakerDomainAttributes) RetentionPolicy() terra.ListValue[RetentionPolicyAttributes] {
	return terra.ReferenceAsList[RetentionPolicyAttributes](asd.ref.Append("retention_policy"))
}

type awsSagemakerDomainState struct {
	AppNetworkAccessType                     string                      `json:"app_network_access_type"`
	AppSecurityGroupManagement               string                      `json:"app_security_group_management"`
	Arn                                      string                      `json:"arn"`
	AuthMode                                 string                      `json:"auth_mode"`
	DomainName                               string                      `json:"domain_name"`
	HomeEfsFileSystemId                      string                      `json:"home_efs_file_system_id"`
	Id                                       string                      `json:"id"`
	KmsKeyId                                 string                      `json:"kms_key_id"`
	SecurityGroupIdForDomainBoundary         string                      `json:"security_group_id_for_domain_boundary"`
	SingleSignOnApplicationArn               string                      `json:"single_sign_on_application_arn"`
	SingleSignOnManagedApplicationInstanceId string                      `json:"single_sign_on_managed_application_instance_id"`
	SubnetIds                                []string                    `json:"subnet_ids"`
	Tags                                     map[string]string           `json:"tags"`
	TagsAll                                  map[string]string           `json:"tags_all"`
	Url                                      string                      `json:"url"`
	VpcId                                    string                      `json:"vpc_id"`
	DefaultSpaceSettings                     []DefaultSpaceSettingsState `json:"default_space_settings"`
	DefaultUserSettings                      []DefaultUserSettingsState  `json:"default_user_settings"`
	DomainSettings                           []DomainSettingsState       `json:"domain_settings"`
	RetentionPolicy                          []RetentionPolicyState      `json:"retention_policy"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerdomain "github.com/golingon/terraproviders/aws/5.7.0/sagemakerdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerDomain creates a new instance of [SagemakerDomain].
func NewSagemakerDomain(name string, args SagemakerDomainArgs) *SagemakerDomain {
	return &SagemakerDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerDomain)(nil)

// SagemakerDomain represents the Terraform resource aws_sagemaker_domain.
type SagemakerDomain struct {
	Name      string
	Args      SagemakerDomainArgs
	state     *sagemakerDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerDomain].
func (sd *SagemakerDomain) Type() string {
	return "aws_sagemaker_domain"
}

// LocalName returns the local name for [SagemakerDomain].
func (sd *SagemakerDomain) LocalName() string {
	return sd.Name
}

// Configuration returns the configuration (args) for [SagemakerDomain].
func (sd *SagemakerDomain) Configuration() interface{} {
	return sd.Args
}

// DependOn is used for other resources to depend on [SagemakerDomain].
func (sd *SagemakerDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(sd)
}

// Dependencies returns the list of resources [SagemakerDomain] depends_on.
func (sd *SagemakerDomain) Dependencies() terra.Dependencies {
	return sd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerDomain].
func (sd *SagemakerDomain) LifecycleManagement() *terra.Lifecycle {
	return sd.Lifecycle
}

// Attributes returns the attributes for [SagemakerDomain].
func (sd *SagemakerDomain) Attributes() sagemakerDomainAttributes {
	return sagemakerDomainAttributes{ref: terra.ReferenceResource(sd)}
}

// ImportState imports the given attribute values into [SagemakerDomain]'s state.
func (sd *SagemakerDomain) ImportState(av io.Reader) error {
	sd.state = &sagemakerDomainState{}
	if err := json.NewDecoder(av).Decode(sd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sd.Type(), sd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerDomain] has state.
func (sd *SagemakerDomain) State() (*sagemakerDomainState, bool) {
	return sd.state, sd.state != nil
}

// StateMust returns the state for [SagemakerDomain]. Panics if the state is nil.
func (sd *SagemakerDomain) StateMust() *sagemakerDomainState {
	if sd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sd.Type(), sd.LocalName()))
	}
	return sd.state
}

// SagemakerDomainArgs contains the configurations for aws_sagemaker_domain.
type SagemakerDomainArgs struct {
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
	DefaultSpaceSettings *sagemakerdomain.DefaultSpaceSettings `hcl:"default_space_settings,block"`
	// DefaultUserSettings: required
	DefaultUserSettings *sagemakerdomain.DefaultUserSettings `hcl:"default_user_settings,block" validate:"required"`
	// DomainSettings: optional
	DomainSettings *sagemakerdomain.DomainSettings `hcl:"domain_settings,block"`
	// RetentionPolicy: optional
	RetentionPolicy *sagemakerdomain.RetentionPolicy `hcl:"retention_policy,block"`
}
type sagemakerDomainAttributes struct {
	ref terra.Reference
}

// AppNetworkAccessType returns a reference to field app_network_access_type of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) AppNetworkAccessType() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("app_network_access_type"))
}

// AppSecurityGroupManagement returns a reference to field app_security_group_management of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) AppSecurityGroupManagement() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("app_security_group_management"))
}

// Arn returns a reference to field arn of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("arn"))
}

// AuthMode returns a reference to field auth_mode of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) AuthMode() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("auth_mode"))
}

// DomainName returns a reference to field domain_name of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("domain_name"))
}

// HomeEfsFileSystemId returns a reference to field home_efs_file_system_id of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) HomeEfsFileSystemId() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("home_efs_file_system_id"))
}

// Id returns a reference to field id of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("kms_key_id"))
}

// SecurityGroupIdForDomainBoundary returns a reference to field security_group_id_for_domain_boundary of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) SecurityGroupIdForDomainBoundary() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("security_group_id_for_domain_boundary"))
}

// SingleSignOnManagedApplicationInstanceId returns a reference to field single_sign_on_managed_application_instance_id of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) SingleSignOnManagedApplicationInstanceId() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("single_sign_on_managed_application_instance_id"))
}

// SubnetIds returns a reference to field subnet_ids of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sd.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sd.ref.Append("tags_all"))
}

// Url returns a reference to field url of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("url"))
}

// VpcId returns a reference to field vpc_id of aws_sagemaker_domain.
func (sd sagemakerDomainAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("vpc_id"))
}

func (sd sagemakerDomainAttributes) DefaultSpaceSettings() terra.ListValue[sagemakerdomain.DefaultSpaceSettingsAttributes] {
	return terra.ReferenceAsList[sagemakerdomain.DefaultSpaceSettingsAttributes](sd.ref.Append("default_space_settings"))
}

func (sd sagemakerDomainAttributes) DefaultUserSettings() terra.ListValue[sagemakerdomain.DefaultUserSettingsAttributes] {
	return terra.ReferenceAsList[sagemakerdomain.DefaultUserSettingsAttributes](sd.ref.Append("default_user_settings"))
}

func (sd sagemakerDomainAttributes) DomainSettings() terra.ListValue[sagemakerdomain.DomainSettingsAttributes] {
	return terra.ReferenceAsList[sagemakerdomain.DomainSettingsAttributes](sd.ref.Append("domain_settings"))
}

func (sd sagemakerDomainAttributes) RetentionPolicy() terra.ListValue[sagemakerdomain.RetentionPolicyAttributes] {
	return terra.ReferenceAsList[sagemakerdomain.RetentionPolicyAttributes](sd.ref.Append("retention_policy"))
}

type sagemakerDomainState struct {
	AppNetworkAccessType                     string                                      `json:"app_network_access_type"`
	AppSecurityGroupManagement               string                                      `json:"app_security_group_management"`
	Arn                                      string                                      `json:"arn"`
	AuthMode                                 string                                      `json:"auth_mode"`
	DomainName                               string                                      `json:"domain_name"`
	HomeEfsFileSystemId                      string                                      `json:"home_efs_file_system_id"`
	Id                                       string                                      `json:"id"`
	KmsKeyId                                 string                                      `json:"kms_key_id"`
	SecurityGroupIdForDomainBoundary         string                                      `json:"security_group_id_for_domain_boundary"`
	SingleSignOnManagedApplicationInstanceId string                                      `json:"single_sign_on_managed_application_instance_id"`
	SubnetIds                                []string                                    `json:"subnet_ids"`
	Tags                                     map[string]string                           `json:"tags"`
	TagsAll                                  map[string]string                           `json:"tags_all"`
	Url                                      string                                      `json:"url"`
	VpcId                                    string                                      `json:"vpc_id"`
	DefaultSpaceSettings                     []sagemakerdomain.DefaultSpaceSettingsState `json:"default_space_settings"`
	DefaultUserSettings                      []sagemakerdomain.DefaultUserSettingsState  `json:"default_user_settings"`
	DomainSettings                           []sagemakerdomain.DomainSettingsState       `json:"domain_settings"`
	RetentionPolicy                          []sagemakerdomain.RetentionPolicyState      `json:"retention_policy"`
}

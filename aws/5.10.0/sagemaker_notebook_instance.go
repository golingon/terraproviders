// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakernotebookinstance "github.com/golingon/terraproviders/aws/5.10.0/sagemakernotebookinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerNotebookInstance creates a new instance of [SagemakerNotebookInstance].
func NewSagemakerNotebookInstance(name string, args SagemakerNotebookInstanceArgs) *SagemakerNotebookInstance {
	return &SagemakerNotebookInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerNotebookInstance)(nil)

// SagemakerNotebookInstance represents the Terraform resource aws_sagemaker_notebook_instance.
type SagemakerNotebookInstance struct {
	Name      string
	Args      SagemakerNotebookInstanceArgs
	state     *sagemakerNotebookInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerNotebookInstance].
func (sni *SagemakerNotebookInstance) Type() string {
	return "aws_sagemaker_notebook_instance"
}

// LocalName returns the local name for [SagemakerNotebookInstance].
func (sni *SagemakerNotebookInstance) LocalName() string {
	return sni.Name
}

// Configuration returns the configuration (args) for [SagemakerNotebookInstance].
func (sni *SagemakerNotebookInstance) Configuration() interface{} {
	return sni.Args
}

// DependOn is used for other resources to depend on [SagemakerNotebookInstance].
func (sni *SagemakerNotebookInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(sni)
}

// Dependencies returns the list of resources [SagemakerNotebookInstance] depends_on.
func (sni *SagemakerNotebookInstance) Dependencies() terra.Dependencies {
	return sni.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerNotebookInstance].
func (sni *SagemakerNotebookInstance) LifecycleManagement() *terra.Lifecycle {
	return sni.Lifecycle
}

// Attributes returns the attributes for [SagemakerNotebookInstance].
func (sni *SagemakerNotebookInstance) Attributes() sagemakerNotebookInstanceAttributes {
	return sagemakerNotebookInstanceAttributes{ref: terra.ReferenceResource(sni)}
}

// ImportState imports the given attribute values into [SagemakerNotebookInstance]'s state.
func (sni *SagemakerNotebookInstance) ImportState(av io.Reader) error {
	sni.state = &sagemakerNotebookInstanceState{}
	if err := json.NewDecoder(av).Decode(sni.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sni.Type(), sni.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerNotebookInstance] has state.
func (sni *SagemakerNotebookInstance) State() (*sagemakerNotebookInstanceState, bool) {
	return sni.state, sni.state != nil
}

// StateMust returns the state for [SagemakerNotebookInstance]. Panics if the state is nil.
func (sni *SagemakerNotebookInstance) StateMust() *sagemakerNotebookInstanceState {
	if sni.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sni.Type(), sni.LocalName()))
	}
	return sni.state
}

// SagemakerNotebookInstanceArgs contains the configurations for aws_sagemaker_notebook_instance.
type SagemakerNotebookInstanceArgs struct {
	// AcceleratorTypes: set of string, optional
	AcceleratorTypes terra.SetValue[terra.StringValue] `hcl:"accelerator_types,attr"`
	// AdditionalCodeRepositories: set of string, optional
	AdditionalCodeRepositories terra.SetValue[terra.StringValue] `hcl:"additional_code_repositories,attr"`
	// DefaultCodeRepository: string, optional
	DefaultCodeRepository terra.StringValue `hcl:"default_code_repository,attr"`
	// DirectInternetAccess: string, optional
	DirectInternetAccess terra.StringValue `hcl:"direct_internet_access,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// LifecycleConfigName: string, optional
	LifecycleConfigName terra.StringValue `hcl:"lifecycle_config_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PlatformIdentifier: string, optional
	PlatformIdentifier terra.StringValue `hcl:"platform_identifier,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// RootAccess: string, optional
	RootAccess terra.StringValue `hcl:"root_access,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VolumeSize: number, optional
	VolumeSize terra.NumberValue `hcl:"volume_size,attr"`
	// InstanceMetadataServiceConfiguration: optional
	InstanceMetadataServiceConfiguration *sagemakernotebookinstance.InstanceMetadataServiceConfiguration `hcl:"instance_metadata_service_configuration,block"`
}
type sagemakerNotebookInstanceAttributes struct {
	ref terra.Reference
}

// AcceleratorTypes returns a reference to field accelerator_types of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) AcceleratorTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sni.ref.Append("accelerator_types"))
}

// AdditionalCodeRepositories returns a reference to field additional_code_repositories of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) AdditionalCodeRepositories() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sni.ref.Append("additional_code_repositories"))
}

// Arn returns a reference to field arn of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("arn"))
}

// DefaultCodeRepository returns a reference to field default_code_repository of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) DefaultCodeRepository() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("default_code_repository"))
}

// DirectInternetAccess returns a reference to field direct_internet_access of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) DirectInternetAccess() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("direct_internet_access"))
}

// Id returns a reference to field id of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("id"))
}

// InstanceType returns a reference to field instance_type of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("instance_type"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("kms_key_id"))
}

// LifecycleConfigName returns a reference to field lifecycle_config_name of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) LifecycleConfigName() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("lifecycle_config_name"))
}

// Name returns a reference to field name of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("name"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("network_interface_id"))
}

// PlatformIdentifier returns a reference to field platform_identifier of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) PlatformIdentifier() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("platform_identifier"))
}

// RoleArn returns a reference to field role_arn of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("role_arn"))
}

// RootAccess returns a reference to field root_access of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) RootAccess() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("root_access"))
}

// SecurityGroups returns a reference to field security_groups of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sni.ref.Append("security_groups"))
}

// SubnetId returns a reference to field subnet_id of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sni.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sni.ref.Append("tags_all"))
}

// Url returns a reference to field url of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sni.ref.Append("url"))
}

// VolumeSize returns a reference to field volume_size of aws_sagemaker_notebook_instance.
func (sni sagemakerNotebookInstanceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(sni.ref.Append("volume_size"))
}

func (sni sagemakerNotebookInstanceAttributes) InstanceMetadataServiceConfiguration() terra.ListValue[sagemakernotebookinstance.InstanceMetadataServiceConfigurationAttributes] {
	return terra.ReferenceAsList[sagemakernotebookinstance.InstanceMetadataServiceConfigurationAttributes](sni.ref.Append("instance_metadata_service_configuration"))
}

type sagemakerNotebookInstanceState struct {
	AcceleratorTypes                     []string                                                              `json:"accelerator_types"`
	AdditionalCodeRepositories           []string                                                              `json:"additional_code_repositories"`
	Arn                                  string                                                                `json:"arn"`
	DefaultCodeRepository                string                                                                `json:"default_code_repository"`
	DirectInternetAccess                 string                                                                `json:"direct_internet_access"`
	Id                                   string                                                                `json:"id"`
	InstanceType                         string                                                                `json:"instance_type"`
	KmsKeyId                             string                                                                `json:"kms_key_id"`
	LifecycleConfigName                  string                                                                `json:"lifecycle_config_name"`
	Name                                 string                                                                `json:"name"`
	NetworkInterfaceId                   string                                                                `json:"network_interface_id"`
	PlatformIdentifier                   string                                                                `json:"platform_identifier"`
	RoleArn                              string                                                                `json:"role_arn"`
	RootAccess                           string                                                                `json:"root_access"`
	SecurityGroups                       []string                                                              `json:"security_groups"`
	SubnetId                             string                                                                `json:"subnet_id"`
	Tags                                 map[string]string                                                     `json:"tags"`
	TagsAll                              map[string]string                                                     `json:"tags_all"`
	Url                                  string                                                                `json:"url"`
	VolumeSize                           float64                                                               `json:"volume_size"`
	InstanceMetadataServiceConfiguration []sagemakernotebookinstance.InstanceMetadataServiceConfigurationState `json:"instance_metadata_service_configuration"`
}

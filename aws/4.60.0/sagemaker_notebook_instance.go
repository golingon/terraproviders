// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakernotebookinstance "github.com/golingon/terraproviders/aws/4.60.0/sagemakernotebookinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSagemakerNotebookInstance(name string, args SagemakerNotebookInstanceArgs) *SagemakerNotebookInstance {
	return &SagemakerNotebookInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerNotebookInstance)(nil)

type SagemakerNotebookInstance struct {
	Name  string
	Args  SagemakerNotebookInstanceArgs
	state *sagemakerNotebookInstanceState
}

func (sni *SagemakerNotebookInstance) Type() string {
	return "aws_sagemaker_notebook_instance"
}

func (sni *SagemakerNotebookInstance) LocalName() string {
	return sni.Name
}

func (sni *SagemakerNotebookInstance) Configuration() interface{} {
	return sni.Args
}

func (sni *SagemakerNotebookInstance) Attributes() sagemakerNotebookInstanceAttributes {
	return sagemakerNotebookInstanceAttributes{ref: terra.ReferenceResource(sni)}
}

func (sni *SagemakerNotebookInstance) ImportState(av io.Reader) error {
	sni.state = &sagemakerNotebookInstanceState{}
	if err := json.NewDecoder(av).Decode(sni.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sni.Type(), sni.LocalName(), err)
	}
	return nil
}

func (sni *SagemakerNotebookInstance) State() (*sagemakerNotebookInstanceState, bool) {
	return sni.state, sni.state != nil
}

func (sni *SagemakerNotebookInstance) StateMust() *sagemakerNotebookInstanceState {
	if sni.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sni.Type(), sni.LocalName()))
	}
	return sni.state
}

func (sni *SagemakerNotebookInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(sni)
}

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
	// DependsOn contains resources that SagemakerNotebookInstance depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sagemakerNotebookInstanceAttributes struct {
	ref terra.Reference
}

func (sni sagemakerNotebookInstanceAttributes) AcceleratorTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](sni.ref.Append("accelerator_types"))
}

func (sni sagemakerNotebookInstanceAttributes) AdditionalCodeRepositories() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](sni.ref.Append("additional_code_repositories"))
}

func (sni sagemakerNotebookInstanceAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("arn"))
}

func (sni sagemakerNotebookInstanceAttributes) DefaultCodeRepository() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("default_code_repository"))
}

func (sni sagemakerNotebookInstanceAttributes) DirectInternetAccess() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("direct_internet_access"))
}

func (sni sagemakerNotebookInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("id"))
}

func (sni sagemakerNotebookInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("instance_type"))
}

func (sni sagemakerNotebookInstanceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("kms_key_id"))
}

func (sni sagemakerNotebookInstanceAttributes) LifecycleConfigName() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("lifecycle_config_name"))
}

func (sni sagemakerNotebookInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("name"))
}

func (sni sagemakerNotebookInstanceAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("network_interface_id"))
}

func (sni sagemakerNotebookInstanceAttributes) PlatformIdentifier() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("platform_identifier"))
}

func (sni sagemakerNotebookInstanceAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("role_arn"))
}

func (sni sagemakerNotebookInstanceAttributes) RootAccess() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("root_access"))
}

func (sni sagemakerNotebookInstanceAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](sni.ref.Append("security_groups"))
}

func (sni sagemakerNotebookInstanceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("subnet_id"))
}

func (sni sagemakerNotebookInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sni.ref.Append("tags"))
}

func (sni sagemakerNotebookInstanceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sni.ref.Append("tags_all"))
}

func (sni sagemakerNotebookInstanceAttributes) Url() terra.StringValue {
	return terra.ReferenceString(sni.ref.Append("url"))
}

func (sni sagemakerNotebookInstanceAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceNumber(sni.ref.Append("volume_size"))
}

func (sni sagemakerNotebookInstanceAttributes) InstanceMetadataServiceConfiguration() terra.ListValue[sagemakernotebookinstance.InstanceMetadataServiceConfigurationAttributes] {
	return terra.ReferenceList[sagemakernotebookinstance.InstanceMetadataServiceConfigurationAttributes](sni.ref.Append("instance_metadata_service_configuration"))
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

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_launch_template

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

// Resource represents the Terraform resource aws_launch_template.
type Resource struct {
	Name      string
	Args      Args
	state     *awsLaunchTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (alt *Resource) Type() string {
	return "aws_launch_template"
}

// LocalName returns the local name for [Resource].
func (alt *Resource) LocalName() string {
	return alt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (alt *Resource) Configuration() interface{} {
	return alt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (alt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(alt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (alt *Resource) Dependencies() terra.Dependencies {
	return alt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (alt *Resource) LifecycleManagement() *terra.Lifecycle {
	return alt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (alt *Resource) Attributes() awsLaunchTemplateAttributes {
	return awsLaunchTemplateAttributes{ref: terra.ReferenceResource(alt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (alt *Resource) ImportState(state io.Reader) error {
	alt.state = &awsLaunchTemplateState{}
	if err := json.NewDecoder(state).Decode(alt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", alt.Type(), alt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (alt *Resource) State() (*awsLaunchTemplateState, bool) {
	return alt.state, alt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (alt *Resource) StateMust() *awsLaunchTemplateState {
	if alt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", alt.Type(), alt.LocalName()))
	}
	return alt.state
}

// Args contains the configurations for aws_launch_template.
type Args struct {
	// DefaultVersion: number, optional
	DefaultVersion terra.NumberValue `hcl:"default_version,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableApiStop: bool, optional
	DisableApiStop terra.BoolValue `hcl:"disable_api_stop,attr"`
	// DisableApiTermination: bool, optional
	DisableApiTermination terra.BoolValue `hcl:"disable_api_termination,attr"`
	// EbsOptimized: string, optional
	EbsOptimized terra.StringValue `hcl:"ebs_optimized,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageId: string, optional
	ImageId terra.StringValue `hcl:"image_id,attr"`
	// InstanceInitiatedShutdownBehavior: string, optional
	InstanceInitiatedShutdownBehavior terra.StringValue `hcl:"instance_initiated_shutdown_behavior,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// KernelId: string, optional
	KernelId terra.StringValue `hcl:"kernel_id,attr"`
	// KeyName: string, optional
	KeyName terra.StringValue `hcl:"key_name,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// RamDiskId: string, optional
	RamDiskId terra.StringValue `hcl:"ram_disk_id,attr"`
	// SecurityGroupNames: set of string, optional
	SecurityGroupNames terra.SetValue[terra.StringValue] `hcl:"security_group_names,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UpdateDefaultVersion: bool, optional
	UpdateDefaultVersion terra.BoolValue `hcl:"update_default_version,attr"`
	// UserData: string, optional
	UserData terra.StringValue `hcl:"user_data,attr"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// BlockDeviceMappings: min=0
	BlockDeviceMappings []BlockDeviceMappings `hcl:"block_device_mappings,block" validate:"min=0"`
	// CapacityReservationSpecification: optional
	CapacityReservationSpecification *CapacityReservationSpecification `hcl:"capacity_reservation_specification,block"`
	// CpuOptions: optional
	CpuOptions *CpuOptions `hcl:"cpu_options,block"`
	// CreditSpecification: optional
	CreditSpecification *CreditSpecification `hcl:"credit_specification,block"`
	// ElasticGpuSpecifications: min=0
	ElasticGpuSpecifications []ElasticGpuSpecifications `hcl:"elastic_gpu_specifications,block" validate:"min=0"`
	// ElasticInferenceAccelerator: optional
	ElasticInferenceAccelerator *ElasticInferenceAccelerator `hcl:"elastic_inference_accelerator,block"`
	// EnclaveOptions: optional
	EnclaveOptions *EnclaveOptions `hcl:"enclave_options,block"`
	// HibernationOptions: optional
	HibernationOptions *HibernationOptions `hcl:"hibernation_options,block"`
	// IamInstanceProfile: optional
	IamInstanceProfile *IamInstanceProfile `hcl:"iam_instance_profile,block"`
	// InstanceMarketOptions: optional
	InstanceMarketOptions *InstanceMarketOptions `hcl:"instance_market_options,block"`
	// InstanceRequirements: optional
	InstanceRequirements *InstanceRequirements `hcl:"instance_requirements,block"`
	// LicenseSpecification: min=0
	LicenseSpecification []LicenseSpecification `hcl:"license_specification,block" validate:"min=0"`
	// MaintenanceOptions: optional
	MaintenanceOptions *MaintenanceOptions `hcl:"maintenance_options,block"`
	// MetadataOptions: optional
	MetadataOptions *MetadataOptions `hcl:"metadata_options,block"`
	// Monitoring: optional
	Monitoring *Monitoring `hcl:"monitoring,block"`
	// NetworkInterfaces: min=0
	NetworkInterfaces []NetworkInterfaces `hcl:"network_interfaces,block" validate:"min=0"`
	// Placement: optional
	Placement *Placement `hcl:"placement,block"`
	// PrivateDnsNameOptions: optional
	PrivateDnsNameOptions *PrivateDnsNameOptions `hcl:"private_dns_name_options,block"`
	// TagSpecifications: min=0
	TagSpecifications []TagSpecifications `hcl:"tag_specifications,block" validate:"min=0"`
}

type awsLaunchTemplateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_launch_template.
func (alt awsLaunchTemplateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("arn"))
}

// DefaultVersion returns a reference to field default_version of aws_launch_template.
func (alt awsLaunchTemplateAttributes) DefaultVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(alt.ref.Append("default_version"))
}

// Description returns a reference to field description of aws_launch_template.
func (alt awsLaunchTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("description"))
}

// DisableApiStop returns a reference to field disable_api_stop of aws_launch_template.
func (alt awsLaunchTemplateAttributes) DisableApiStop() terra.BoolValue {
	return terra.ReferenceAsBool(alt.ref.Append("disable_api_stop"))
}

// DisableApiTermination returns a reference to field disable_api_termination of aws_launch_template.
func (alt awsLaunchTemplateAttributes) DisableApiTermination() terra.BoolValue {
	return terra.ReferenceAsBool(alt.ref.Append("disable_api_termination"))
}

// EbsOptimized returns a reference to field ebs_optimized of aws_launch_template.
func (alt awsLaunchTemplateAttributes) EbsOptimized() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("ebs_optimized"))
}

// Id returns a reference to field id of aws_launch_template.
func (alt awsLaunchTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("id"))
}

// ImageId returns a reference to field image_id of aws_launch_template.
func (alt awsLaunchTemplateAttributes) ImageId() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("image_id"))
}

// InstanceInitiatedShutdownBehavior returns a reference to field instance_initiated_shutdown_behavior of aws_launch_template.
func (alt awsLaunchTemplateAttributes) InstanceInitiatedShutdownBehavior() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("instance_initiated_shutdown_behavior"))
}

// InstanceType returns a reference to field instance_type of aws_launch_template.
func (alt awsLaunchTemplateAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("instance_type"))
}

// KernelId returns a reference to field kernel_id of aws_launch_template.
func (alt awsLaunchTemplateAttributes) KernelId() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("kernel_id"))
}

// KeyName returns a reference to field key_name of aws_launch_template.
func (alt awsLaunchTemplateAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("key_name"))
}

// LatestVersion returns a reference to field latest_version of aws_launch_template.
func (alt awsLaunchTemplateAttributes) LatestVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(alt.ref.Append("latest_version"))
}

// Name returns a reference to field name of aws_launch_template.
func (alt awsLaunchTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_launch_template.
func (alt awsLaunchTemplateAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("name_prefix"))
}

// RamDiskId returns a reference to field ram_disk_id of aws_launch_template.
func (alt awsLaunchTemplateAttributes) RamDiskId() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("ram_disk_id"))
}

// SecurityGroupNames returns a reference to field security_group_names of aws_launch_template.
func (alt awsLaunchTemplateAttributes) SecurityGroupNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](alt.ref.Append("security_group_names"))
}

// Tags returns a reference to field tags of aws_launch_template.
func (alt awsLaunchTemplateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_launch_template.
func (alt awsLaunchTemplateAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alt.ref.Append("tags_all"))
}

// UpdateDefaultVersion returns a reference to field update_default_version of aws_launch_template.
func (alt awsLaunchTemplateAttributes) UpdateDefaultVersion() terra.BoolValue {
	return terra.ReferenceAsBool(alt.ref.Append("update_default_version"))
}

// UserData returns a reference to field user_data of aws_launch_template.
func (alt awsLaunchTemplateAttributes) UserData() terra.StringValue {
	return terra.ReferenceAsString(alt.ref.Append("user_data"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_launch_template.
func (alt awsLaunchTemplateAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](alt.ref.Append("vpc_security_group_ids"))
}

func (alt awsLaunchTemplateAttributes) BlockDeviceMappings() terra.ListValue[BlockDeviceMappingsAttributes] {
	return terra.ReferenceAsList[BlockDeviceMappingsAttributes](alt.ref.Append("block_device_mappings"))
}

func (alt awsLaunchTemplateAttributes) CapacityReservationSpecification() terra.ListValue[CapacityReservationSpecificationAttributes] {
	return terra.ReferenceAsList[CapacityReservationSpecificationAttributes](alt.ref.Append("capacity_reservation_specification"))
}

func (alt awsLaunchTemplateAttributes) CpuOptions() terra.ListValue[CpuOptionsAttributes] {
	return terra.ReferenceAsList[CpuOptionsAttributes](alt.ref.Append("cpu_options"))
}

func (alt awsLaunchTemplateAttributes) CreditSpecification() terra.ListValue[CreditSpecificationAttributes] {
	return terra.ReferenceAsList[CreditSpecificationAttributes](alt.ref.Append("credit_specification"))
}

func (alt awsLaunchTemplateAttributes) ElasticGpuSpecifications() terra.ListValue[ElasticGpuSpecificationsAttributes] {
	return terra.ReferenceAsList[ElasticGpuSpecificationsAttributes](alt.ref.Append("elastic_gpu_specifications"))
}

func (alt awsLaunchTemplateAttributes) ElasticInferenceAccelerator() terra.ListValue[ElasticInferenceAcceleratorAttributes] {
	return terra.ReferenceAsList[ElasticInferenceAcceleratorAttributes](alt.ref.Append("elastic_inference_accelerator"))
}

func (alt awsLaunchTemplateAttributes) EnclaveOptions() terra.ListValue[EnclaveOptionsAttributes] {
	return terra.ReferenceAsList[EnclaveOptionsAttributes](alt.ref.Append("enclave_options"))
}

func (alt awsLaunchTemplateAttributes) HibernationOptions() terra.ListValue[HibernationOptionsAttributes] {
	return terra.ReferenceAsList[HibernationOptionsAttributes](alt.ref.Append("hibernation_options"))
}

func (alt awsLaunchTemplateAttributes) IamInstanceProfile() terra.ListValue[IamInstanceProfileAttributes] {
	return terra.ReferenceAsList[IamInstanceProfileAttributes](alt.ref.Append("iam_instance_profile"))
}

func (alt awsLaunchTemplateAttributes) InstanceMarketOptions() terra.ListValue[InstanceMarketOptionsAttributes] {
	return terra.ReferenceAsList[InstanceMarketOptionsAttributes](alt.ref.Append("instance_market_options"))
}

func (alt awsLaunchTemplateAttributes) InstanceRequirements() terra.ListValue[InstanceRequirementsAttributes] {
	return terra.ReferenceAsList[InstanceRequirementsAttributes](alt.ref.Append("instance_requirements"))
}

func (alt awsLaunchTemplateAttributes) LicenseSpecification() terra.SetValue[LicenseSpecificationAttributes] {
	return terra.ReferenceAsSet[LicenseSpecificationAttributes](alt.ref.Append("license_specification"))
}

func (alt awsLaunchTemplateAttributes) MaintenanceOptions() terra.ListValue[MaintenanceOptionsAttributes] {
	return terra.ReferenceAsList[MaintenanceOptionsAttributes](alt.ref.Append("maintenance_options"))
}

func (alt awsLaunchTemplateAttributes) MetadataOptions() terra.ListValue[MetadataOptionsAttributes] {
	return terra.ReferenceAsList[MetadataOptionsAttributes](alt.ref.Append("metadata_options"))
}

func (alt awsLaunchTemplateAttributes) Monitoring() terra.ListValue[MonitoringAttributes] {
	return terra.ReferenceAsList[MonitoringAttributes](alt.ref.Append("monitoring"))
}

func (alt awsLaunchTemplateAttributes) NetworkInterfaces() terra.ListValue[NetworkInterfacesAttributes] {
	return terra.ReferenceAsList[NetworkInterfacesAttributes](alt.ref.Append("network_interfaces"))
}

func (alt awsLaunchTemplateAttributes) Placement() terra.ListValue[PlacementAttributes] {
	return terra.ReferenceAsList[PlacementAttributes](alt.ref.Append("placement"))
}

func (alt awsLaunchTemplateAttributes) PrivateDnsNameOptions() terra.ListValue[PrivateDnsNameOptionsAttributes] {
	return terra.ReferenceAsList[PrivateDnsNameOptionsAttributes](alt.ref.Append("private_dns_name_options"))
}

func (alt awsLaunchTemplateAttributes) TagSpecifications() terra.ListValue[TagSpecificationsAttributes] {
	return terra.ReferenceAsList[TagSpecificationsAttributes](alt.ref.Append("tag_specifications"))
}

type awsLaunchTemplateState struct {
	Arn                               string                                  `json:"arn"`
	DefaultVersion                    float64                                 `json:"default_version"`
	Description                       string                                  `json:"description"`
	DisableApiStop                    bool                                    `json:"disable_api_stop"`
	DisableApiTermination             bool                                    `json:"disable_api_termination"`
	EbsOptimized                      string                                  `json:"ebs_optimized"`
	Id                                string                                  `json:"id"`
	ImageId                           string                                  `json:"image_id"`
	InstanceInitiatedShutdownBehavior string                                  `json:"instance_initiated_shutdown_behavior"`
	InstanceType                      string                                  `json:"instance_type"`
	KernelId                          string                                  `json:"kernel_id"`
	KeyName                           string                                  `json:"key_name"`
	LatestVersion                     float64                                 `json:"latest_version"`
	Name                              string                                  `json:"name"`
	NamePrefix                        string                                  `json:"name_prefix"`
	RamDiskId                         string                                  `json:"ram_disk_id"`
	SecurityGroupNames                []string                                `json:"security_group_names"`
	Tags                              map[string]string                       `json:"tags"`
	TagsAll                           map[string]string                       `json:"tags_all"`
	UpdateDefaultVersion              bool                                    `json:"update_default_version"`
	UserData                          string                                  `json:"user_data"`
	VpcSecurityGroupIds               []string                                `json:"vpc_security_group_ids"`
	BlockDeviceMappings               []BlockDeviceMappingsState              `json:"block_device_mappings"`
	CapacityReservationSpecification  []CapacityReservationSpecificationState `json:"capacity_reservation_specification"`
	CpuOptions                        []CpuOptionsState                       `json:"cpu_options"`
	CreditSpecification               []CreditSpecificationState              `json:"credit_specification"`
	ElasticGpuSpecifications          []ElasticGpuSpecificationsState         `json:"elastic_gpu_specifications"`
	ElasticInferenceAccelerator       []ElasticInferenceAcceleratorState      `json:"elastic_inference_accelerator"`
	EnclaveOptions                    []EnclaveOptionsState                   `json:"enclave_options"`
	HibernationOptions                []HibernationOptionsState               `json:"hibernation_options"`
	IamInstanceProfile                []IamInstanceProfileState               `json:"iam_instance_profile"`
	InstanceMarketOptions             []InstanceMarketOptionsState            `json:"instance_market_options"`
	InstanceRequirements              []InstanceRequirementsState             `json:"instance_requirements"`
	LicenseSpecification              []LicenseSpecificationState             `json:"license_specification"`
	MaintenanceOptions                []MaintenanceOptionsState               `json:"maintenance_options"`
	MetadataOptions                   []MetadataOptionsState                  `json:"metadata_options"`
	Monitoring                        []MonitoringState                       `json:"monitoring"`
	NetworkInterfaces                 []NetworkInterfacesState                `json:"network_interfaces"`
	Placement                         []PlacementState                        `json:"placement"`
	PrivateDnsNameOptions             []PrivateDnsNameOptionsState            `json:"private_dns_name_options"`
	TagSpecifications                 []TagSpecificationsState                `json:"tag_specifications"`
}

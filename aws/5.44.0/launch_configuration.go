// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	launchconfiguration "github.com/golingon/terraproviders/aws/5.44.0/launchconfiguration"
	"io"
)

// NewLaunchConfiguration creates a new instance of [LaunchConfiguration].
func NewLaunchConfiguration(name string, args LaunchConfigurationArgs) *LaunchConfiguration {
	return &LaunchConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LaunchConfiguration)(nil)

// LaunchConfiguration represents the Terraform resource aws_launch_configuration.
type LaunchConfiguration struct {
	Name      string
	Args      LaunchConfigurationArgs
	state     *launchConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LaunchConfiguration].
func (lc *LaunchConfiguration) Type() string {
	return "aws_launch_configuration"
}

// LocalName returns the local name for [LaunchConfiguration].
func (lc *LaunchConfiguration) LocalName() string {
	return lc.Name
}

// Configuration returns the configuration (args) for [LaunchConfiguration].
func (lc *LaunchConfiguration) Configuration() interface{} {
	return lc.Args
}

// DependOn is used for other resources to depend on [LaunchConfiguration].
func (lc *LaunchConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(lc)
}

// Dependencies returns the list of resources [LaunchConfiguration] depends_on.
func (lc *LaunchConfiguration) Dependencies() terra.Dependencies {
	return lc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LaunchConfiguration].
func (lc *LaunchConfiguration) LifecycleManagement() *terra.Lifecycle {
	return lc.Lifecycle
}

// Attributes returns the attributes for [LaunchConfiguration].
func (lc *LaunchConfiguration) Attributes() launchConfigurationAttributes {
	return launchConfigurationAttributes{ref: terra.ReferenceResource(lc)}
}

// ImportState imports the given attribute values into [LaunchConfiguration]'s state.
func (lc *LaunchConfiguration) ImportState(av io.Reader) error {
	lc.state = &launchConfigurationState{}
	if err := json.NewDecoder(av).Decode(lc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lc.Type(), lc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LaunchConfiguration] has state.
func (lc *LaunchConfiguration) State() (*launchConfigurationState, bool) {
	return lc.state, lc.state != nil
}

// StateMust returns the state for [LaunchConfiguration]. Panics if the state is nil.
func (lc *LaunchConfiguration) StateMust() *launchConfigurationState {
	if lc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lc.Type(), lc.LocalName()))
	}
	return lc.state
}

// LaunchConfigurationArgs contains the configurations for aws_launch_configuration.
type LaunchConfigurationArgs struct {
	// AssociatePublicIpAddress: bool, optional
	AssociatePublicIpAddress terra.BoolValue `hcl:"associate_public_ip_address,attr"`
	// EbsOptimized: bool, optional
	EbsOptimized terra.BoolValue `hcl:"ebs_optimized,attr"`
	// EnableMonitoring: bool, optional
	EnableMonitoring terra.BoolValue `hcl:"enable_monitoring,attr"`
	// IamInstanceProfile: string, optional
	IamInstanceProfile terra.StringValue `hcl:"iam_instance_profile,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageId: string, required
	ImageId terra.StringValue `hcl:"image_id,attr" validate:"required"`
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// KeyName: string, optional
	KeyName terra.StringValue `hcl:"key_name,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// PlacementTenancy: string, optional
	PlacementTenancy terra.StringValue `hcl:"placement_tenancy,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// SpotPrice: string, optional
	SpotPrice terra.StringValue `hcl:"spot_price,attr"`
	// UserData: string, optional
	UserData terra.StringValue `hcl:"user_data,attr"`
	// UserDataBase64: string, optional
	UserDataBase64 terra.StringValue `hcl:"user_data_base64,attr"`
	// EbsBlockDevice: min=0
	EbsBlockDevice []launchconfiguration.EbsBlockDevice `hcl:"ebs_block_device,block" validate:"min=0"`
	// EphemeralBlockDevice: min=0
	EphemeralBlockDevice []launchconfiguration.EphemeralBlockDevice `hcl:"ephemeral_block_device,block" validate:"min=0"`
	// MetadataOptions: optional
	MetadataOptions *launchconfiguration.MetadataOptions `hcl:"metadata_options,block"`
	// RootBlockDevice: optional
	RootBlockDevice *launchconfiguration.RootBlockDevice `hcl:"root_block_device,block"`
}
type launchConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_launch_configuration.
func (lc launchConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("arn"))
}

// AssociatePublicIpAddress returns a reference to field associate_public_ip_address of aws_launch_configuration.
func (lc launchConfigurationAttributes) AssociatePublicIpAddress() terra.BoolValue {
	return terra.ReferenceAsBool(lc.ref.Append("associate_public_ip_address"))
}

// EbsOptimized returns a reference to field ebs_optimized of aws_launch_configuration.
func (lc launchConfigurationAttributes) EbsOptimized() terra.BoolValue {
	return terra.ReferenceAsBool(lc.ref.Append("ebs_optimized"))
}

// EnableMonitoring returns a reference to field enable_monitoring of aws_launch_configuration.
func (lc launchConfigurationAttributes) EnableMonitoring() terra.BoolValue {
	return terra.ReferenceAsBool(lc.ref.Append("enable_monitoring"))
}

// IamInstanceProfile returns a reference to field iam_instance_profile of aws_launch_configuration.
func (lc launchConfigurationAttributes) IamInstanceProfile() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("iam_instance_profile"))
}

// Id returns a reference to field id of aws_launch_configuration.
func (lc launchConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("id"))
}

// ImageId returns a reference to field image_id of aws_launch_configuration.
func (lc launchConfigurationAttributes) ImageId() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("image_id"))
}

// InstanceType returns a reference to field instance_type of aws_launch_configuration.
func (lc launchConfigurationAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("instance_type"))
}

// KeyName returns a reference to field key_name of aws_launch_configuration.
func (lc launchConfigurationAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("key_name"))
}

// Name returns a reference to field name of aws_launch_configuration.
func (lc launchConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_launch_configuration.
func (lc launchConfigurationAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("name_prefix"))
}

// PlacementTenancy returns a reference to field placement_tenancy of aws_launch_configuration.
func (lc launchConfigurationAttributes) PlacementTenancy() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("placement_tenancy"))
}

// SecurityGroups returns a reference to field security_groups of aws_launch_configuration.
func (lc launchConfigurationAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lc.ref.Append("security_groups"))
}

// SpotPrice returns a reference to field spot_price of aws_launch_configuration.
func (lc launchConfigurationAttributes) SpotPrice() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("spot_price"))
}

// UserData returns a reference to field user_data of aws_launch_configuration.
func (lc launchConfigurationAttributes) UserData() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("user_data"))
}

// UserDataBase64 returns a reference to field user_data_base64 of aws_launch_configuration.
func (lc launchConfigurationAttributes) UserDataBase64() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("user_data_base64"))
}

func (lc launchConfigurationAttributes) EbsBlockDevice() terra.SetValue[launchconfiguration.EbsBlockDeviceAttributes] {
	return terra.ReferenceAsSet[launchconfiguration.EbsBlockDeviceAttributes](lc.ref.Append("ebs_block_device"))
}

func (lc launchConfigurationAttributes) EphemeralBlockDevice() terra.SetValue[launchconfiguration.EphemeralBlockDeviceAttributes] {
	return terra.ReferenceAsSet[launchconfiguration.EphemeralBlockDeviceAttributes](lc.ref.Append("ephemeral_block_device"))
}

func (lc launchConfigurationAttributes) MetadataOptions() terra.ListValue[launchconfiguration.MetadataOptionsAttributes] {
	return terra.ReferenceAsList[launchconfiguration.MetadataOptionsAttributes](lc.ref.Append("metadata_options"))
}

func (lc launchConfigurationAttributes) RootBlockDevice() terra.ListValue[launchconfiguration.RootBlockDeviceAttributes] {
	return terra.ReferenceAsList[launchconfiguration.RootBlockDeviceAttributes](lc.ref.Append("root_block_device"))
}

type launchConfigurationState struct {
	Arn                      string                                          `json:"arn"`
	AssociatePublicIpAddress bool                                            `json:"associate_public_ip_address"`
	EbsOptimized             bool                                            `json:"ebs_optimized"`
	EnableMonitoring         bool                                            `json:"enable_monitoring"`
	IamInstanceProfile       string                                          `json:"iam_instance_profile"`
	Id                       string                                          `json:"id"`
	ImageId                  string                                          `json:"image_id"`
	InstanceType             string                                          `json:"instance_type"`
	KeyName                  string                                          `json:"key_name"`
	Name                     string                                          `json:"name"`
	NamePrefix               string                                          `json:"name_prefix"`
	PlacementTenancy         string                                          `json:"placement_tenancy"`
	SecurityGroups           []string                                        `json:"security_groups"`
	SpotPrice                string                                          `json:"spot_price"`
	UserData                 string                                          `json:"user_data"`
	UserDataBase64           string                                          `json:"user_data_base64"`
	EbsBlockDevice           []launchconfiguration.EbsBlockDeviceState       `json:"ebs_block_device"`
	EphemeralBlockDevice     []launchconfiguration.EphemeralBlockDeviceState `json:"ephemeral_block_device"`
	MetadataOptions          []launchconfiguration.MetadataOptionsState      `json:"metadata_options"`
	RootBlockDevice          []launchconfiguration.RootBlockDeviceState      `json:"root_block_device"`
}

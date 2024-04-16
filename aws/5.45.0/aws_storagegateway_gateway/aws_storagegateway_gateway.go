// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_storagegateway_gateway

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

// Resource represents the Terraform resource aws_storagegateway_gateway.
type Resource struct {
	Name      string
	Args      Args
	state     *awsStoragegatewayGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asg *Resource) Type() string {
	return "aws_storagegateway_gateway"
}

// LocalName returns the local name for [Resource].
func (asg *Resource) LocalName() string {
	return asg.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asg *Resource) Configuration() interface{} {
	return asg.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asg *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asg)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asg *Resource) Dependencies() terra.Dependencies {
	return asg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asg *Resource) LifecycleManagement() *terra.Lifecycle {
	return asg.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asg *Resource) Attributes() awsStoragegatewayGatewayAttributes {
	return awsStoragegatewayGatewayAttributes{ref: terra.ReferenceResource(asg)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asg *Resource) ImportState(state io.Reader) error {
	asg.state = &awsStoragegatewayGatewayState{}
	if err := json.NewDecoder(state).Decode(asg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asg.Type(), asg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asg *Resource) State() (*awsStoragegatewayGatewayState, bool) {
	return asg.state, asg.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asg *Resource) StateMust() *awsStoragegatewayGatewayState {
	if asg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asg.Type(), asg.LocalName()))
	}
	return asg.state
}

// Args contains the configurations for aws_storagegateway_gateway.
type Args struct {
	// ActivationKey: string, optional
	ActivationKey terra.StringValue `hcl:"activation_key,attr"`
	// AverageDownloadRateLimitInBitsPerSec: number, optional
	AverageDownloadRateLimitInBitsPerSec terra.NumberValue `hcl:"average_download_rate_limit_in_bits_per_sec,attr"`
	// AverageUploadRateLimitInBitsPerSec: number, optional
	AverageUploadRateLimitInBitsPerSec terra.NumberValue `hcl:"average_upload_rate_limit_in_bits_per_sec,attr"`
	// CloudwatchLogGroupArn: string, optional
	CloudwatchLogGroupArn terra.StringValue `hcl:"cloudwatch_log_group_arn,attr"`
	// GatewayIpAddress: string, optional
	GatewayIpAddress terra.StringValue `hcl:"gateway_ip_address,attr"`
	// GatewayName: string, required
	GatewayName terra.StringValue `hcl:"gateway_name,attr" validate:"required"`
	// GatewayTimezone: string, required
	GatewayTimezone terra.StringValue `hcl:"gateway_timezone,attr" validate:"required"`
	// GatewayType: string, optional
	GatewayType terra.StringValue `hcl:"gateway_type,attr"`
	// GatewayVpcEndpoint: string, optional
	GatewayVpcEndpoint terra.StringValue `hcl:"gateway_vpc_endpoint,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MediumChangerType: string, optional
	MediumChangerType terra.StringValue `hcl:"medium_changer_type,attr"`
	// SmbFileShareVisibility: bool, optional
	SmbFileShareVisibility terra.BoolValue `hcl:"smb_file_share_visibility,attr"`
	// SmbGuestPassword: string, optional
	SmbGuestPassword terra.StringValue `hcl:"smb_guest_password,attr"`
	// SmbSecurityStrategy: string, optional
	SmbSecurityStrategy terra.StringValue `hcl:"smb_security_strategy,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TapeDriveType: string, optional
	TapeDriveType terra.StringValue `hcl:"tape_drive_type,attr"`
	// MaintenanceStartTime: optional
	MaintenanceStartTime *MaintenanceStartTime `hcl:"maintenance_start_time,block"`
	// SmbActiveDirectorySettings: optional
	SmbActiveDirectorySettings *SmbActiveDirectorySettings `hcl:"smb_active_directory_settings,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsStoragegatewayGatewayAttributes struct {
	ref terra.Reference
}

// ActivationKey returns a reference to field activation_key of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) ActivationKey() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("activation_key"))
}

// Arn returns a reference to field arn of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("arn"))
}

// AverageDownloadRateLimitInBitsPerSec returns a reference to field average_download_rate_limit_in_bits_per_sec of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) AverageDownloadRateLimitInBitsPerSec() terra.NumberValue {
	return terra.ReferenceAsNumber(asg.ref.Append("average_download_rate_limit_in_bits_per_sec"))
}

// AverageUploadRateLimitInBitsPerSec returns a reference to field average_upload_rate_limit_in_bits_per_sec of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) AverageUploadRateLimitInBitsPerSec() terra.NumberValue {
	return terra.ReferenceAsNumber(asg.ref.Append("average_upload_rate_limit_in_bits_per_sec"))
}

// CloudwatchLogGroupArn returns a reference to field cloudwatch_log_group_arn of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) CloudwatchLogGroupArn() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("cloudwatch_log_group_arn"))
}

// Ec2InstanceId returns a reference to field ec2_instance_id of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) Ec2InstanceId() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("ec2_instance_id"))
}

// EndpointType returns a reference to field endpoint_type of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("endpoint_type"))
}

// GatewayId returns a reference to field gateway_id of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) GatewayId() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("gateway_id"))
}

// GatewayIpAddress returns a reference to field gateway_ip_address of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) GatewayIpAddress() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("gateway_ip_address"))
}

// GatewayName returns a reference to field gateway_name of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) GatewayName() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("gateway_name"))
}

// GatewayTimezone returns a reference to field gateway_timezone of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) GatewayTimezone() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("gateway_timezone"))
}

// GatewayType returns a reference to field gateway_type of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) GatewayType() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("gateway_type"))
}

// GatewayVpcEndpoint returns a reference to field gateway_vpc_endpoint of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) GatewayVpcEndpoint() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("gateway_vpc_endpoint"))
}

// HostEnvironment returns a reference to field host_environment of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) HostEnvironment() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("host_environment"))
}

// Id returns a reference to field id of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("id"))
}

// MediumChangerType returns a reference to field medium_changer_type of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) MediumChangerType() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("medium_changer_type"))
}

// SmbFileShareVisibility returns a reference to field smb_file_share_visibility of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) SmbFileShareVisibility() terra.BoolValue {
	return terra.ReferenceAsBool(asg.ref.Append("smb_file_share_visibility"))
}

// SmbGuestPassword returns a reference to field smb_guest_password of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) SmbGuestPassword() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("smb_guest_password"))
}

// SmbSecurityStrategy returns a reference to field smb_security_strategy of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) SmbSecurityStrategy() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("smb_security_strategy"))
}

// Tags returns a reference to field tags of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asg.ref.Append("tags_all"))
}

// TapeDriveType returns a reference to field tape_drive_type of aws_storagegateway_gateway.
func (asg awsStoragegatewayGatewayAttributes) TapeDriveType() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("tape_drive_type"))
}

func (asg awsStoragegatewayGatewayAttributes) GatewayNetworkInterface() terra.ListValue[GatewayNetworkInterfaceAttributes] {
	return terra.ReferenceAsList[GatewayNetworkInterfaceAttributes](asg.ref.Append("gateway_network_interface"))
}

func (asg awsStoragegatewayGatewayAttributes) MaintenanceStartTime() terra.ListValue[MaintenanceStartTimeAttributes] {
	return terra.ReferenceAsList[MaintenanceStartTimeAttributes](asg.ref.Append("maintenance_start_time"))
}

func (asg awsStoragegatewayGatewayAttributes) SmbActiveDirectorySettings() terra.ListValue[SmbActiveDirectorySettingsAttributes] {
	return terra.ReferenceAsList[SmbActiveDirectorySettingsAttributes](asg.ref.Append("smb_active_directory_settings"))
}

func (asg awsStoragegatewayGatewayAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asg.ref.Append("timeouts"))
}

type awsStoragegatewayGatewayState struct {
	ActivationKey                        string                            `json:"activation_key"`
	Arn                                  string                            `json:"arn"`
	AverageDownloadRateLimitInBitsPerSec float64                           `json:"average_download_rate_limit_in_bits_per_sec"`
	AverageUploadRateLimitInBitsPerSec   float64                           `json:"average_upload_rate_limit_in_bits_per_sec"`
	CloudwatchLogGroupArn                string                            `json:"cloudwatch_log_group_arn"`
	Ec2InstanceId                        string                            `json:"ec2_instance_id"`
	EndpointType                         string                            `json:"endpoint_type"`
	GatewayId                            string                            `json:"gateway_id"`
	GatewayIpAddress                     string                            `json:"gateway_ip_address"`
	GatewayName                          string                            `json:"gateway_name"`
	GatewayTimezone                      string                            `json:"gateway_timezone"`
	GatewayType                          string                            `json:"gateway_type"`
	GatewayVpcEndpoint                   string                            `json:"gateway_vpc_endpoint"`
	HostEnvironment                      string                            `json:"host_environment"`
	Id                                   string                            `json:"id"`
	MediumChangerType                    string                            `json:"medium_changer_type"`
	SmbFileShareVisibility               bool                              `json:"smb_file_share_visibility"`
	SmbGuestPassword                     string                            `json:"smb_guest_password"`
	SmbSecurityStrategy                  string                            `json:"smb_security_strategy"`
	Tags                                 map[string]string                 `json:"tags"`
	TagsAll                              map[string]string                 `json:"tags_all"`
	TapeDriveType                        string                            `json:"tape_drive_type"`
	GatewayNetworkInterface              []GatewayNetworkInterfaceState    `json:"gateway_network_interface"`
	MaintenanceStartTime                 []MaintenanceStartTimeState       `json:"maintenance_start_time"`
	SmbActiveDirectorySettings           []SmbActiveDirectorySettingsState `json:"smb_active_directory_settings"`
	Timeouts                             *TimeoutsState                    `json:"timeouts"`
}

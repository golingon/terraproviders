// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	storagegatewaygateway "github.com/golingon/terraproviders/aws/5.9.0/storagegatewaygateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStoragegatewayGateway creates a new instance of [StoragegatewayGateway].
func NewStoragegatewayGateway(name string, args StoragegatewayGatewayArgs) *StoragegatewayGateway {
	return &StoragegatewayGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StoragegatewayGateway)(nil)

// StoragegatewayGateway represents the Terraform resource aws_storagegateway_gateway.
type StoragegatewayGateway struct {
	Name      string
	Args      StoragegatewayGatewayArgs
	state     *storagegatewayGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StoragegatewayGateway].
func (sg *StoragegatewayGateway) Type() string {
	return "aws_storagegateway_gateway"
}

// LocalName returns the local name for [StoragegatewayGateway].
func (sg *StoragegatewayGateway) LocalName() string {
	return sg.Name
}

// Configuration returns the configuration (args) for [StoragegatewayGateway].
func (sg *StoragegatewayGateway) Configuration() interface{} {
	return sg.Args
}

// DependOn is used for other resources to depend on [StoragegatewayGateway].
func (sg *StoragegatewayGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(sg)
}

// Dependencies returns the list of resources [StoragegatewayGateway] depends_on.
func (sg *StoragegatewayGateway) Dependencies() terra.Dependencies {
	return sg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StoragegatewayGateway].
func (sg *StoragegatewayGateway) LifecycleManagement() *terra.Lifecycle {
	return sg.Lifecycle
}

// Attributes returns the attributes for [StoragegatewayGateway].
func (sg *StoragegatewayGateway) Attributes() storagegatewayGatewayAttributes {
	return storagegatewayGatewayAttributes{ref: terra.ReferenceResource(sg)}
}

// ImportState imports the given attribute values into [StoragegatewayGateway]'s state.
func (sg *StoragegatewayGateway) ImportState(av io.Reader) error {
	sg.state = &storagegatewayGatewayState{}
	if err := json.NewDecoder(av).Decode(sg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sg.Type(), sg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StoragegatewayGateway] has state.
func (sg *StoragegatewayGateway) State() (*storagegatewayGatewayState, bool) {
	return sg.state, sg.state != nil
}

// StateMust returns the state for [StoragegatewayGateway]. Panics if the state is nil.
func (sg *StoragegatewayGateway) StateMust() *storagegatewayGatewayState {
	if sg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sg.Type(), sg.LocalName()))
	}
	return sg.state
}

// StoragegatewayGatewayArgs contains the configurations for aws_storagegateway_gateway.
type StoragegatewayGatewayArgs struct {
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
	// GatewayNetworkInterface: min=0
	GatewayNetworkInterface []storagegatewaygateway.GatewayNetworkInterface `hcl:"gateway_network_interface,block" validate:"min=0"`
	// MaintenanceStartTime: optional
	MaintenanceStartTime *storagegatewaygateway.MaintenanceStartTime `hcl:"maintenance_start_time,block"`
	// SmbActiveDirectorySettings: optional
	SmbActiveDirectorySettings *storagegatewaygateway.SmbActiveDirectorySettings `hcl:"smb_active_directory_settings,block"`
	// Timeouts: optional
	Timeouts *storagegatewaygateway.Timeouts `hcl:"timeouts,block"`
}
type storagegatewayGatewayAttributes struct {
	ref terra.Reference
}

// ActivationKey returns a reference to field activation_key of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) ActivationKey() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("activation_key"))
}

// Arn returns a reference to field arn of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("arn"))
}

// AverageDownloadRateLimitInBitsPerSec returns a reference to field average_download_rate_limit_in_bits_per_sec of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) AverageDownloadRateLimitInBitsPerSec() terra.NumberValue {
	return terra.ReferenceAsNumber(sg.ref.Append("average_download_rate_limit_in_bits_per_sec"))
}

// AverageUploadRateLimitInBitsPerSec returns a reference to field average_upload_rate_limit_in_bits_per_sec of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) AverageUploadRateLimitInBitsPerSec() terra.NumberValue {
	return terra.ReferenceAsNumber(sg.ref.Append("average_upload_rate_limit_in_bits_per_sec"))
}

// CloudwatchLogGroupArn returns a reference to field cloudwatch_log_group_arn of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) CloudwatchLogGroupArn() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("cloudwatch_log_group_arn"))
}

// Ec2InstanceId returns a reference to field ec2_instance_id of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) Ec2InstanceId() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("ec2_instance_id"))
}

// EndpointType returns a reference to field endpoint_type of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("endpoint_type"))
}

// GatewayId returns a reference to field gateway_id of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) GatewayId() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("gateway_id"))
}

// GatewayIpAddress returns a reference to field gateway_ip_address of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) GatewayIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("gateway_ip_address"))
}

// GatewayName returns a reference to field gateway_name of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) GatewayName() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("gateway_name"))
}

// GatewayTimezone returns a reference to field gateway_timezone of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) GatewayTimezone() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("gateway_timezone"))
}

// GatewayType returns a reference to field gateway_type of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) GatewayType() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("gateway_type"))
}

// GatewayVpcEndpoint returns a reference to field gateway_vpc_endpoint of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) GatewayVpcEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("gateway_vpc_endpoint"))
}

// HostEnvironment returns a reference to field host_environment of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) HostEnvironment() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("host_environment"))
}

// Id returns a reference to field id of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("id"))
}

// MediumChangerType returns a reference to field medium_changer_type of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) MediumChangerType() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("medium_changer_type"))
}

// SmbFileShareVisibility returns a reference to field smb_file_share_visibility of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) SmbFileShareVisibility() terra.BoolValue {
	return terra.ReferenceAsBool(sg.ref.Append("smb_file_share_visibility"))
}

// SmbGuestPassword returns a reference to field smb_guest_password of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) SmbGuestPassword() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("smb_guest_password"))
}

// SmbSecurityStrategy returns a reference to field smb_security_strategy of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) SmbSecurityStrategy() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("smb_security_strategy"))
}

// Tags returns a reference to field tags of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sg.ref.Append("tags_all"))
}

// TapeDriveType returns a reference to field tape_drive_type of aws_storagegateway_gateway.
func (sg storagegatewayGatewayAttributes) TapeDriveType() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("tape_drive_type"))
}

func (sg storagegatewayGatewayAttributes) GatewayNetworkInterface() terra.ListValue[storagegatewaygateway.GatewayNetworkInterfaceAttributes] {
	return terra.ReferenceAsList[storagegatewaygateway.GatewayNetworkInterfaceAttributes](sg.ref.Append("gateway_network_interface"))
}

func (sg storagegatewayGatewayAttributes) MaintenanceStartTime() terra.ListValue[storagegatewaygateway.MaintenanceStartTimeAttributes] {
	return terra.ReferenceAsList[storagegatewaygateway.MaintenanceStartTimeAttributes](sg.ref.Append("maintenance_start_time"))
}

func (sg storagegatewayGatewayAttributes) SmbActiveDirectorySettings() terra.ListValue[storagegatewaygateway.SmbActiveDirectorySettingsAttributes] {
	return terra.ReferenceAsList[storagegatewaygateway.SmbActiveDirectorySettingsAttributes](sg.ref.Append("smb_active_directory_settings"))
}

func (sg storagegatewayGatewayAttributes) Timeouts() storagegatewaygateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagegatewaygateway.TimeoutsAttributes](sg.ref.Append("timeouts"))
}

type storagegatewayGatewayState struct {
	ActivationKey                        string                                                  `json:"activation_key"`
	Arn                                  string                                                  `json:"arn"`
	AverageDownloadRateLimitInBitsPerSec float64                                                 `json:"average_download_rate_limit_in_bits_per_sec"`
	AverageUploadRateLimitInBitsPerSec   float64                                                 `json:"average_upload_rate_limit_in_bits_per_sec"`
	CloudwatchLogGroupArn                string                                                  `json:"cloudwatch_log_group_arn"`
	Ec2InstanceId                        string                                                  `json:"ec2_instance_id"`
	EndpointType                         string                                                  `json:"endpoint_type"`
	GatewayId                            string                                                  `json:"gateway_id"`
	GatewayIpAddress                     string                                                  `json:"gateway_ip_address"`
	GatewayName                          string                                                  `json:"gateway_name"`
	GatewayTimezone                      string                                                  `json:"gateway_timezone"`
	GatewayType                          string                                                  `json:"gateway_type"`
	GatewayVpcEndpoint                   string                                                  `json:"gateway_vpc_endpoint"`
	HostEnvironment                      string                                                  `json:"host_environment"`
	Id                                   string                                                  `json:"id"`
	MediumChangerType                    string                                                  `json:"medium_changer_type"`
	SmbFileShareVisibility               bool                                                    `json:"smb_file_share_visibility"`
	SmbGuestPassword                     string                                                  `json:"smb_guest_password"`
	SmbSecurityStrategy                  string                                                  `json:"smb_security_strategy"`
	Tags                                 map[string]string                                       `json:"tags"`
	TagsAll                              map[string]string                                       `json:"tags_all"`
	TapeDriveType                        string                                                  `json:"tape_drive_type"`
	GatewayNetworkInterface              []storagegatewaygateway.GatewayNetworkInterfaceState    `json:"gateway_network_interface"`
	MaintenanceStartTime                 []storagegatewaygateway.MaintenanceStartTimeState       `json:"maintenance_start_time"`
	SmbActiveDirectorySettings           []storagegatewaygateway.SmbActiveDirectorySettingsState `json:"smb_active_directory_settings"`
	Timeouts                             *storagegatewaygateway.TimeoutsState                    `json:"timeouts"`
}

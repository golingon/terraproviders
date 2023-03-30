// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opsworksinstance "github.com/golingon/terraproviders/aws/4.60.0/opsworksinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewOpsworksInstance(name string, args OpsworksInstanceArgs) *OpsworksInstance {
	return &OpsworksInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksInstance)(nil)

type OpsworksInstance struct {
	Name  string
	Args  OpsworksInstanceArgs
	state *opsworksInstanceState
}

func (oi *OpsworksInstance) Type() string {
	return "aws_opsworks_instance"
}

func (oi *OpsworksInstance) LocalName() string {
	return oi.Name
}

func (oi *OpsworksInstance) Configuration() interface{} {
	return oi.Args
}

func (oi *OpsworksInstance) Attributes() opsworksInstanceAttributes {
	return opsworksInstanceAttributes{ref: terra.ReferenceResource(oi)}
}

func (oi *OpsworksInstance) ImportState(av io.Reader) error {
	oi.state = &opsworksInstanceState{}
	if err := json.NewDecoder(av).Decode(oi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oi.Type(), oi.LocalName(), err)
	}
	return nil
}

func (oi *OpsworksInstance) State() (*opsworksInstanceState, bool) {
	return oi.state, oi.state != nil
}

func (oi *OpsworksInstance) StateMust() *opsworksInstanceState {
	if oi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oi.Type(), oi.LocalName()))
	}
	return oi.state
}

func (oi *OpsworksInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(oi)
}

type OpsworksInstanceArgs struct {
	// AgentVersion: string, optional
	AgentVersion terra.StringValue `hcl:"agent_version,attr"`
	// AmiId: string, optional
	AmiId terra.StringValue `hcl:"ami_id,attr"`
	// Architecture: string, optional
	Architecture terra.StringValue `hcl:"architecture,attr"`
	// AutoScalingType: string, optional
	AutoScalingType terra.StringValue `hcl:"auto_scaling_type,attr"`
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// CreatedAt: string, optional
	CreatedAt terra.StringValue `hcl:"created_at,attr"`
	// DeleteEbs: bool, optional
	DeleteEbs terra.BoolValue `hcl:"delete_ebs,attr"`
	// DeleteEip: bool, optional
	DeleteEip terra.BoolValue `hcl:"delete_eip,attr"`
	// EbsOptimized: bool, optional
	EbsOptimized terra.BoolValue `hcl:"ebs_optimized,attr"`
	// EcsClusterArn: string, optional
	EcsClusterArn terra.StringValue `hcl:"ecs_cluster_arn,attr"`
	// ElasticIp: string, optional
	ElasticIp terra.StringValue `hcl:"elastic_ip,attr"`
	// Hostname: string, optional
	Hostname terra.StringValue `hcl:"hostname,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InfrastructureClass: string, optional
	InfrastructureClass terra.StringValue `hcl:"infrastructure_class,attr"`
	// InstallUpdatesOnBoot: bool, optional
	InstallUpdatesOnBoot terra.BoolValue `hcl:"install_updates_on_boot,attr"`
	// InstanceProfileArn: string, optional
	InstanceProfileArn terra.StringValue `hcl:"instance_profile_arn,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// LayerIds: list of string, required
	LayerIds terra.ListValue[terra.StringValue] `hcl:"layer_ids,attr" validate:"required"`
	// Os: string, optional
	Os terra.StringValue `hcl:"os,attr"`
	// RootDeviceType: string, optional
	RootDeviceType terra.StringValue `hcl:"root_device_type,attr"`
	// SecurityGroupIds: list of string, optional
	SecurityGroupIds terra.ListValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SshKeyName: string, optional
	SshKeyName terra.StringValue `hcl:"ssh_key_name,attr"`
	// StackId: string, required
	StackId terra.StringValue `hcl:"stack_id,attr" validate:"required"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Tenancy: string, optional
	Tenancy terra.StringValue `hcl:"tenancy,attr"`
	// VirtualizationType: string, optional
	VirtualizationType terra.StringValue `hcl:"virtualization_type,attr"`
	// EbsBlockDevice: min=0
	EbsBlockDevice []opsworksinstance.EbsBlockDevice `hcl:"ebs_block_device,block" validate:"min=0"`
	// EphemeralBlockDevice: min=0
	EphemeralBlockDevice []opsworksinstance.EphemeralBlockDevice `hcl:"ephemeral_block_device,block" validate:"min=0"`
	// RootBlockDevice: min=0
	RootBlockDevice []opsworksinstance.RootBlockDevice `hcl:"root_block_device,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *opsworksinstance.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that OpsworksInstance depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type opsworksInstanceAttributes struct {
	ref terra.Reference
}

func (oi opsworksInstanceAttributes) AgentVersion() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("agent_version"))
}

func (oi opsworksInstanceAttributes) AmiId() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("ami_id"))
}

func (oi opsworksInstanceAttributes) Architecture() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("architecture"))
}

func (oi opsworksInstanceAttributes) AutoScalingType() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("auto_scaling_type"))
}

func (oi opsworksInstanceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("availability_zone"))
}

func (oi opsworksInstanceAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("created_at"))
}

func (oi opsworksInstanceAttributes) DeleteEbs() terra.BoolValue {
	return terra.ReferenceBool(oi.ref.Append("delete_ebs"))
}

func (oi opsworksInstanceAttributes) DeleteEip() terra.BoolValue {
	return terra.ReferenceBool(oi.ref.Append("delete_eip"))
}

func (oi opsworksInstanceAttributes) EbsOptimized() terra.BoolValue {
	return terra.ReferenceBool(oi.ref.Append("ebs_optimized"))
}

func (oi opsworksInstanceAttributes) Ec2InstanceId() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("ec2_instance_id"))
}

func (oi opsworksInstanceAttributes) EcsClusterArn() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("ecs_cluster_arn"))
}

func (oi opsworksInstanceAttributes) ElasticIp() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("elastic_ip"))
}

func (oi opsworksInstanceAttributes) Hostname() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("hostname"))
}

func (oi opsworksInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("id"))
}

func (oi opsworksInstanceAttributes) InfrastructureClass() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("infrastructure_class"))
}

func (oi opsworksInstanceAttributes) InstallUpdatesOnBoot() terra.BoolValue {
	return terra.ReferenceBool(oi.ref.Append("install_updates_on_boot"))
}

func (oi opsworksInstanceAttributes) InstanceProfileArn() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("instance_profile_arn"))
}

func (oi opsworksInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("instance_type"))
}

func (oi opsworksInstanceAttributes) LastServiceErrorId() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("last_service_error_id"))
}

func (oi opsworksInstanceAttributes) LayerIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](oi.ref.Append("layer_ids"))
}

func (oi opsworksInstanceAttributes) Os() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("os"))
}

func (oi opsworksInstanceAttributes) Platform() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("platform"))
}

func (oi opsworksInstanceAttributes) PrivateDns() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("private_dns"))
}

func (oi opsworksInstanceAttributes) PrivateIp() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("private_ip"))
}

func (oi opsworksInstanceAttributes) PublicDns() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("public_dns"))
}

func (oi opsworksInstanceAttributes) PublicIp() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("public_ip"))
}

func (oi opsworksInstanceAttributes) RegisteredBy() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("registered_by"))
}

func (oi opsworksInstanceAttributes) ReportedAgentVersion() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("reported_agent_version"))
}

func (oi opsworksInstanceAttributes) ReportedOsFamily() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("reported_os_family"))
}

func (oi opsworksInstanceAttributes) ReportedOsName() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("reported_os_name"))
}

func (oi opsworksInstanceAttributes) ReportedOsVersion() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("reported_os_version"))
}

func (oi opsworksInstanceAttributes) RootDeviceType() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("root_device_type"))
}

func (oi opsworksInstanceAttributes) RootDeviceVolumeId() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("root_device_volume_id"))
}

func (oi opsworksInstanceAttributes) SecurityGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](oi.ref.Append("security_group_ids"))
}

func (oi opsworksInstanceAttributes) SshHostDsaKeyFingerprint() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("ssh_host_dsa_key_fingerprint"))
}

func (oi opsworksInstanceAttributes) SshHostRsaKeyFingerprint() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("ssh_host_rsa_key_fingerprint"))
}

func (oi opsworksInstanceAttributes) SshKeyName() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("ssh_key_name"))
}

func (oi opsworksInstanceAttributes) StackId() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("stack_id"))
}

func (oi opsworksInstanceAttributes) State() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("state"))
}

func (oi opsworksInstanceAttributes) Status() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("status"))
}

func (oi opsworksInstanceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("subnet_id"))
}

func (oi opsworksInstanceAttributes) Tenancy() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("tenancy"))
}

func (oi opsworksInstanceAttributes) VirtualizationType() terra.StringValue {
	return terra.ReferenceString(oi.ref.Append("virtualization_type"))
}

func (oi opsworksInstanceAttributes) EbsBlockDevice() terra.SetValue[opsworksinstance.EbsBlockDeviceAttributes] {
	return terra.ReferenceSet[opsworksinstance.EbsBlockDeviceAttributes](oi.ref.Append("ebs_block_device"))
}

func (oi opsworksInstanceAttributes) EphemeralBlockDevice() terra.SetValue[opsworksinstance.EphemeralBlockDeviceAttributes] {
	return terra.ReferenceSet[opsworksinstance.EphemeralBlockDeviceAttributes](oi.ref.Append("ephemeral_block_device"))
}

func (oi opsworksInstanceAttributes) RootBlockDevice() terra.SetValue[opsworksinstance.RootBlockDeviceAttributes] {
	return terra.ReferenceSet[opsworksinstance.RootBlockDeviceAttributes](oi.ref.Append("root_block_device"))
}

func (oi opsworksInstanceAttributes) Timeouts() opsworksinstance.TimeoutsAttributes {
	return terra.ReferenceSingle[opsworksinstance.TimeoutsAttributes](oi.ref.Append("timeouts"))
}

type opsworksInstanceState struct {
	AgentVersion             string                                       `json:"agent_version"`
	AmiId                    string                                       `json:"ami_id"`
	Architecture             string                                       `json:"architecture"`
	AutoScalingType          string                                       `json:"auto_scaling_type"`
	AvailabilityZone         string                                       `json:"availability_zone"`
	CreatedAt                string                                       `json:"created_at"`
	DeleteEbs                bool                                         `json:"delete_ebs"`
	DeleteEip                bool                                         `json:"delete_eip"`
	EbsOptimized             bool                                         `json:"ebs_optimized"`
	Ec2InstanceId            string                                       `json:"ec2_instance_id"`
	EcsClusterArn            string                                       `json:"ecs_cluster_arn"`
	ElasticIp                string                                       `json:"elastic_ip"`
	Hostname                 string                                       `json:"hostname"`
	Id                       string                                       `json:"id"`
	InfrastructureClass      string                                       `json:"infrastructure_class"`
	InstallUpdatesOnBoot     bool                                         `json:"install_updates_on_boot"`
	InstanceProfileArn       string                                       `json:"instance_profile_arn"`
	InstanceType             string                                       `json:"instance_type"`
	LastServiceErrorId       string                                       `json:"last_service_error_id"`
	LayerIds                 []string                                     `json:"layer_ids"`
	Os                       string                                       `json:"os"`
	Platform                 string                                       `json:"platform"`
	PrivateDns               string                                       `json:"private_dns"`
	PrivateIp                string                                       `json:"private_ip"`
	PublicDns                string                                       `json:"public_dns"`
	PublicIp                 string                                       `json:"public_ip"`
	RegisteredBy             string                                       `json:"registered_by"`
	ReportedAgentVersion     string                                       `json:"reported_agent_version"`
	ReportedOsFamily         string                                       `json:"reported_os_family"`
	ReportedOsName           string                                       `json:"reported_os_name"`
	ReportedOsVersion        string                                       `json:"reported_os_version"`
	RootDeviceType           string                                       `json:"root_device_type"`
	RootDeviceVolumeId       string                                       `json:"root_device_volume_id"`
	SecurityGroupIds         []string                                     `json:"security_group_ids"`
	SshHostDsaKeyFingerprint string                                       `json:"ssh_host_dsa_key_fingerprint"`
	SshHostRsaKeyFingerprint string                                       `json:"ssh_host_rsa_key_fingerprint"`
	SshKeyName               string                                       `json:"ssh_key_name"`
	StackId                  string                                       `json:"stack_id"`
	State                    string                                       `json:"state"`
	Status                   string                                       `json:"status"`
	SubnetId                 string                                       `json:"subnet_id"`
	Tenancy                  string                                       `json:"tenancy"`
	VirtualizationType       string                                       `json:"virtualization_type"`
	EbsBlockDevice           []opsworksinstance.EbsBlockDeviceState       `json:"ebs_block_device"`
	EphemeralBlockDevice     []opsworksinstance.EphemeralBlockDeviceState `json:"ephemeral_block_device"`
	RootBlockDevice          []opsworksinstance.RootBlockDeviceState      `json:"root_block_device"`
	Timeouts                 *opsworksinstance.TimeoutsState              `json:"timeouts"`
}

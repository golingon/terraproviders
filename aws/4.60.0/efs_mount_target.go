// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	efsmounttarget "github.com/golingon/terraproviders/aws/4.60.0/efsmounttarget"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEfsMountTarget creates a new instance of [EfsMountTarget].
func NewEfsMountTarget(name string, args EfsMountTargetArgs) *EfsMountTarget {
	return &EfsMountTarget{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EfsMountTarget)(nil)

// EfsMountTarget represents the Terraform resource aws_efs_mount_target.
type EfsMountTarget struct {
	Name      string
	Args      EfsMountTargetArgs
	state     *efsMountTargetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EfsMountTarget].
func (emt *EfsMountTarget) Type() string {
	return "aws_efs_mount_target"
}

// LocalName returns the local name for [EfsMountTarget].
func (emt *EfsMountTarget) LocalName() string {
	return emt.Name
}

// Configuration returns the configuration (args) for [EfsMountTarget].
func (emt *EfsMountTarget) Configuration() interface{} {
	return emt.Args
}

// DependOn is used for other resources to depend on [EfsMountTarget].
func (emt *EfsMountTarget) DependOn() terra.Reference {
	return terra.ReferenceResource(emt)
}

// Dependencies returns the list of resources [EfsMountTarget] depends_on.
func (emt *EfsMountTarget) Dependencies() terra.Dependencies {
	return emt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EfsMountTarget].
func (emt *EfsMountTarget) LifecycleManagement() *terra.Lifecycle {
	return emt.Lifecycle
}

// Attributes returns the attributes for [EfsMountTarget].
func (emt *EfsMountTarget) Attributes() efsMountTargetAttributes {
	return efsMountTargetAttributes{ref: terra.ReferenceResource(emt)}
}

// ImportState imports the given attribute values into [EfsMountTarget]'s state.
func (emt *EfsMountTarget) ImportState(av io.Reader) error {
	emt.state = &efsMountTargetState{}
	if err := json.NewDecoder(av).Decode(emt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", emt.Type(), emt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EfsMountTarget] has state.
func (emt *EfsMountTarget) State() (*efsMountTargetState, bool) {
	return emt.state, emt.state != nil
}

// StateMust returns the state for [EfsMountTarget]. Panics if the state is nil.
func (emt *EfsMountTarget) StateMust() *efsMountTargetState {
	if emt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", emt.Type(), emt.LocalName()))
	}
	return emt.state
}

// EfsMountTargetArgs contains the configurations for aws_efs_mount_target.
type EfsMountTargetArgs struct {
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *efsmounttarget.Timeouts `hcl:"timeouts,block"`
}
type efsMountTargetAttributes struct {
	ref terra.Reference
}

// AvailabilityZoneId returns a reference to field availability_zone_id of aws_efs_mount_target.
func (emt efsMountTargetAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("availability_zone_id"))
}

// AvailabilityZoneName returns a reference to field availability_zone_name of aws_efs_mount_target.
func (emt efsMountTargetAttributes) AvailabilityZoneName() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("availability_zone_name"))
}

// DnsName returns a reference to field dns_name of aws_efs_mount_target.
func (emt efsMountTargetAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("dns_name"))
}

// FileSystemArn returns a reference to field file_system_arn of aws_efs_mount_target.
func (emt efsMountTargetAttributes) FileSystemArn() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("file_system_arn"))
}

// FileSystemId returns a reference to field file_system_id of aws_efs_mount_target.
func (emt efsMountTargetAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("file_system_id"))
}

// Id returns a reference to field id of aws_efs_mount_target.
func (emt efsMountTargetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of aws_efs_mount_target.
func (emt efsMountTargetAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("ip_address"))
}

// MountTargetDnsName returns a reference to field mount_target_dns_name of aws_efs_mount_target.
func (emt efsMountTargetAttributes) MountTargetDnsName() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("mount_target_dns_name"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_efs_mount_target.
func (emt efsMountTargetAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("network_interface_id"))
}

// OwnerId returns a reference to field owner_id of aws_efs_mount_target.
func (emt efsMountTargetAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("owner_id"))
}

// SecurityGroups returns a reference to field security_groups of aws_efs_mount_target.
func (emt efsMountTargetAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](emt.ref.Append("security_groups"))
}

// SubnetId returns a reference to field subnet_id of aws_efs_mount_target.
func (emt efsMountTargetAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(emt.ref.Append("subnet_id"))
}

func (emt efsMountTargetAttributes) Timeouts() efsmounttarget.TimeoutsAttributes {
	return terra.ReferenceAsSingle[efsmounttarget.TimeoutsAttributes](emt.ref.Append("timeouts"))
}

type efsMountTargetState struct {
	AvailabilityZoneId   string                        `json:"availability_zone_id"`
	AvailabilityZoneName string                        `json:"availability_zone_name"`
	DnsName              string                        `json:"dns_name"`
	FileSystemArn        string                        `json:"file_system_arn"`
	FileSystemId         string                        `json:"file_system_id"`
	Id                   string                        `json:"id"`
	IpAddress            string                        `json:"ip_address"`
	MountTargetDnsName   string                        `json:"mount_target_dns_name"`
	NetworkInterfaceId   string                        `json:"network_interface_id"`
	OwnerId              string                        `json:"owner_id"`
	SecurityGroups       []string                      `json:"security_groups"`
	SubnetId             string                        `json:"subnet_id"`
	Timeouts             *efsmounttarget.TimeoutsState `json:"timeouts"`
}

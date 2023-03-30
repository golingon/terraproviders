// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	efsmounttarget "github.com/golingon/terraproviders/aws/4.60.0/efsmounttarget"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEfsMountTarget(name string, args EfsMountTargetArgs) *EfsMountTarget {
	return &EfsMountTarget{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EfsMountTarget)(nil)

type EfsMountTarget struct {
	Name  string
	Args  EfsMountTargetArgs
	state *efsMountTargetState
}

func (emt *EfsMountTarget) Type() string {
	return "aws_efs_mount_target"
}

func (emt *EfsMountTarget) LocalName() string {
	return emt.Name
}

func (emt *EfsMountTarget) Configuration() interface{} {
	return emt.Args
}

func (emt *EfsMountTarget) Attributes() efsMountTargetAttributes {
	return efsMountTargetAttributes{ref: terra.ReferenceResource(emt)}
}

func (emt *EfsMountTarget) ImportState(av io.Reader) error {
	emt.state = &efsMountTargetState{}
	if err := json.NewDecoder(av).Decode(emt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", emt.Type(), emt.LocalName(), err)
	}
	return nil
}

func (emt *EfsMountTarget) State() (*efsMountTargetState, bool) {
	return emt.state, emt.state != nil
}

func (emt *EfsMountTarget) StateMust() *efsMountTargetState {
	if emt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", emt.Type(), emt.LocalName()))
	}
	return emt.state
}

func (emt *EfsMountTarget) DependOn() terra.Reference {
	return terra.ReferenceResource(emt)
}

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
	// DependsOn contains resources that EfsMountTarget depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type efsMountTargetAttributes struct {
	ref terra.Reference
}

func (emt efsMountTargetAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("availability_zone_id"))
}

func (emt efsMountTargetAttributes) AvailabilityZoneName() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("availability_zone_name"))
}

func (emt efsMountTargetAttributes) DnsName() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("dns_name"))
}

func (emt efsMountTargetAttributes) FileSystemArn() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("file_system_arn"))
}

func (emt efsMountTargetAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("file_system_id"))
}

func (emt efsMountTargetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("id"))
}

func (emt efsMountTargetAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("ip_address"))
}

func (emt efsMountTargetAttributes) MountTargetDnsName() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("mount_target_dns_name"))
}

func (emt efsMountTargetAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("network_interface_id"))
}

func (emt efsMountTargetAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("owner_id"))
}

func (emt efsMountTargetAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](emt.ref.Append("security_groups"))
}

func (emt efsMountTargetAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(emt.ref.Append("subnet_id"))
}

func (emt efsMountTargetAttributes) Timeouts() efsmounttarget.TimeoutsAttributes {
	return terra.ReferenceSingle[efsmounttarget.TimeoutsAttributes](emt.ref.Append("timeouts"))
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

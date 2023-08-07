// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2host "github.com/golingon/terraproviders/aws/5.11.0/dataec2host"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2Host creates a new instance of [DataEc2Host].
func NewDataEc2Host(name string, args DataEc2HostArgs) *DataEc2Host {
	return &DataEc2Host{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2Host)(nil)

// DataEc2Host represents the Terraform data resource aws_ec2_host.
type DataEc2Host struct {
	Name string
	Args DataEc2HostArgs
}

// DataSource returns the Terraform object type for [DataEc2Host].
func (eh *DataEc2Host) DataSource() string {
	return "aws_ec2_host"
}

// LocalName returns the local name for [DataEc2Host].
func (eh *DataEc2Host) LocalName() string {
	return eh.Name
}

// Configuration returns the configuration (args) for [DataEc2Host].
func (eh *DataEc2Host) Configuration() interface{} {
	return eh.Args
}

// Attributes returns the attributes for [DataEc2Host].
func (eh *DataEc2Host) Attributes() dataEc2HostAttributes {
	return dataEc2HostAttributes{ref: terra.ReferenceDataResource(eh)}
}

// DataEc2HostArgs contains the configurations for aws_ec2_host.
type DataEc2HostArgs struct {
	// HostId: string, optional
	HostId terra.StringValue `hcl:"host_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2host.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2host.Timeouts `hcl:"timeouts,block"`
}
type dataEc2HostAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_host.
func (eh dataEc2HostAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("arn"))
}

// AssetId returns a reference to field asset_id of aws_ec2_host.
func (eh dataEc2HostAttributes) AssetId() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("asset_id"))
}

// AutoPlacement returns a reference to field auto_placement of aws_ec2_host.
func (eh dataEc2HostAttributes) AutoPlacement() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("auto_placement"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_ec2_host.
func (eh dataEc2HostAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("availability_zone"))
}

// Cores returns a reference to field cores of aws_ec2_host.
func (eh dataEc2HostAttributes) Cores() terra.NumberValue {
	return terra.ReferenceAsNumber(eh.ref.Append("cores"))
}

// HostId returns a reference to field host_id of aws_ec2_host.
func (eh dataEc2HostAttributes) HostId() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("host_id"))
}

// HostRecovery returns a reference to field host_recovery of aws_ec2_host.
func (eh dataEc2HostAttributes) HostRecovery() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("host_recovery"))
}

// Id returns a reference to field id of aws_ec2_host.
func (eh dataEc2HostAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("id"))
}

// InstanceFamily returns a reference to field instance_family of aws_ec2_host.
func (eh dataEc2HostAttributes) InstanceFamily() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("instance_family"))
}

// InstanceType returns a reference to field instance_type of aws_ec2_host.
func (eh dataEc2HostAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("instance_type"))
}

// OutpostArn returns a reference to field outpost_arn of aws_ec2_host.
func (eh dataEc2HostAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("outpost_arn"))
}

// OwnerId returns a reference to field owner_id of aws_ec2_host.
func (eh dataEc2HostAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("owner_id"))
}

// Sockets returns a reference to field sockets of aws_ec2_host.
func (eh dataEc2HostAttributes) Sockets() terra.NumberValue {
	return terra.ReferenceAsNumber(eh.ref.Append("sockets"))
}

// Tags returns a reference to field tags of aws_ec2_host.
func (eh dataEc2HostAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eh.ref.Append("tags"))
}

// TotalVcpus returns a reference to field total_vcpus of aws_ec2_host.
func (eh dataEc2HostAttributes) TotalVcpus() terra.NumberValue {
	return terra.ReferenceAsNumber(eh.ref.Append("total_vcpus"))
}

func (eh dataEc2HostAttributes) Filter() terra.SetValue[dataec2host.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2host.FilterAttributes](eh.ref.Append("filter"))
}

func (eh dataEc2HostAttributes) Timeouts() dataec2host.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2host.TimeoutsAttributes](eh.ref.Append("timeouts"))
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcipampool "github.com/golingon/terraproviders/aws/5.7.0/datavpcipampool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcIpamPool creates a new instance of [DataVpcIpamPool].
func NewDataVpcIpamPool(name string, args DataVpcIpamPoolArgs) *DataVpcIpamPool {
	return &DataVpcIpamPool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcIpamPool)(nil)

// DataVpcIpamPool represents the Terraform data resource aws_vpc_ipam_pool.
type DataVpcIpamPool struct {
	Name string
	Args DataVpcIpamPoolArgs
}

// DataSource returns the Terraform object type for [DataVpcIpamPool].
func (vip *DataVpcIpamPool) DataSource() string {
	return "aws_vpc_ipam_pool"
}

// LocalName returns the local name for [DataVpcIpamPool].
func (vip *DataVpcIpamPool) LocalName() string {
	return vip.Name
}

// Configuration returns the configuration (args) for [DataVpcIpamPool].
func (vip *DataVpcIpamPool) Configuration() interface{} {
	return vip.Args
}

// Attributes returns the attributes for [DataVpcIpamPool].
func (vip *DataVpcIpamPool) Attributes() dataVpcIpamPoolAttributes {
	return dataVpcIpamPoolAttributes{ref: terra.ReferenceDataResource(vip)}
}

// DataVpcIpamPoolArgs contains the configurations for aws_vpc_ipam_pool.
type DataVpcIpamPoolArgs struct {
	// AllocationResourceTags: map of string, optional
	AllocationResourceTags terra.MapValue[terra.StringValue] `hcl:"allocation_resource_tags,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpamPoolId: string, optional
	IpamPoolId terra.StringValue `hcl:"ipam_pool_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []datavpcipampool.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpcipampool.Timeouts `hcl:"timeouts,block"`
}
type dataVpcIpamPoolAttributes struct {
	ref terra.Reference
}

// AddressFamily returns a reference to field address_family of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) AddressFamily() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("address_family"))
}

// AllocationDefaultNetmaskLength returns a reference to field allocation_default_netmask_length of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) AllocationDefaultNetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vip.ref.Append("allocation_default_netmask_length"))
}

// AllocationMaxNetmaskLength returns a reference to field allocation_max_netmask_length of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) AllocationMaxNetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vip.ref.Append("allocation_max_netmask_length"))
}

// AllocationMinNetmaskLength returns a reference to field allocation_min_netmask_length of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) AllocationMinNetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vip.ref.Append("allocation_min_netmask_length"))
}

// AllocationResourceTags returns a reference to field allocation_resource_tags of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) AllocationResourceTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vip.ref.Append("allocation_resource_tags"))
}

// Arn returns a reference to field arn of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("arn"))
}

// AutoImport returns a reference to field auto_import of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) AutoImport() terra.BoolValue {
	return terra.ReferenceAsBool(vip.ref.Append("auto_import"))
}

// AwsService returns a reference to field aws_service of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) AwsService() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("aws_service"))
}

// Description returns a reference to field description of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("description"))
}

// Id returns a reference to field id of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("id"))
}

// IpamPoolId returns a reference to field ipam_pool_id of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("ipam_pool_id"))
}

// IpamScopeId returns a reference to field ipam_scope_id of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) IpamScopeId() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("ipam_scope_id"))
}

// IpamScopeType returns a reference to field ipam_scope_type of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) IpamScopeType() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("ipam_scope_type"))
}

// Locale returns a reference to field locale of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) Locale() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("locale"))
}

// PoolDepth returns a reference to field pool_depth of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) PoolDepth() terra.NumberValue {
	return terra.ReferenceAsNumber(vip.ref.Append("pool_depth"))
}

// PubliclyAdvertisable returns a reference to field publicly_advertisable of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) PubliclyAdvertisable() terra.BoolValue {
	return terra.ReferenceAsBool(vip.ref.Append("publicly_advertisable"))
}

// SourceIpamPoolId returns a reference to field source_ipam_pool_id of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) SourceIpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("source_ipam_pool_id"))
}

// State returns a reference to field state of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_vpc_ipam_pool.
func (vip dataVpcIpamPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vip.ref.Append("tags"))
}

func (vip dataVpcIpamPoolAttributes) Filter() terra.SetValue[datavpcipampool.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcipampool.FilterAttributes](vip.ref.Append("filter"))
}

func (vip dataVpcIpamPoolAttributes) Timeouts() datavpcipampool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpcipampool.TimeoutsAttributes](vip.ref.Append("timeouts"))
}

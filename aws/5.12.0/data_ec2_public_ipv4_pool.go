// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2publicipv4pool "github.com/golingon/terraproviders/aws/5.12.0/dataec2publicipv4pool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2PublicIpv4Pool creates a new instance of [DataEc2PublicIpv4Pool].
func NewDataEc2PublicIpv4Pool(name string, args DataEc2PublicIpv4PoolArgs) *DataEc2PublicIpv4Pool {
	return &DataEc2PublicIpv4Pool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2PublicIpv4Pool)(nil)

// DataEc2PublicIpv4Pool represents the Terraform data resource aws_ec2_public_ipv4_pool.
type DataEc2PublicIpv4Pool struct {
	Name string
	Args DataEc2PublicIpv4PoolArgs
}

// DataSource returns the Terraform object type for [DataEc2PublicIpv4Pool].
func (epip *DataEc2PublicIpv4Pool) DataSource() string {
	return "aws_ec2_public_ipv4_pool"
}

// LocalName returns the local name for [DataEc2PublicIpv4Pool].
func (epip *DataEc2PublicIpv4Pool) LocalName() string {
	return epip.Name
}

// Configuration returns the configuration (args) for [DataEc2PublicIpv4Pool].
func (epip *DataEc2PublicIpv4Pool) Configuration() interface{} {
	return epip.Args
}

// Attributes returns the attributes for [DataEc2PublicIpv4Pool].
func (epip *DataEc2PublicIpv4Pool) Attributes() dataEc2PublicIpv4PoolAttributes {
	return dataEc2PublicIpv4PoolAttributes{ref: terra.ReferenceDataResource(epip)}
}

// DataEc2PublicIpv4PoolArgs contains the configurations for aws_ec2_public_ipv4_pool.
type DataEc2PublicIpv4PoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PoolId: string, required
	PoolId terra.StringValue `hcl:"pool_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// PoolAddressRanges: min=0
	PoolAddressRanges []dataec2publicipv4pool.PoolAddressRanges `hcl:"pool_address_ranges,block" validate:"min=0"`
}
type dataEc2PublicIpv4PoolAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_ec2_public_ipv4_pool.
func (epip dataEc2PublicIpv4PoolAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(epip.ref.Append("description"))
}

// Id returns a reference to field id of aws_ec2_public_ipv4_pool.
func (epip dataEc2PublicIpv4PoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(epip.ref.Append("id"))
}

// NetworkBorderGroup returns a reference to field network_border_group of aws_ec2_public_ipv4_pool.
func (epip dataEc2PublicIpv4PoolAttributes) NetworkBorderGroup() terra.StringValue {
	return terra.ReferenceAsString(epip.ref.Append("network_border_group"))
}

// PoolId returns a reference to field pool_id of aws_ec2_public_ipv4_pool.
func (epip dataEc2PublicIpv4PoolAttributes) PoolId() terra.StringValue {
	return terra.ReferenceAsString(epip.ref.Append("pool_id"))
}

// Tags returns a reference to field tags of aws_ec2_public_ipv4_pool.
func (epip dataEc2PublicIpv4PoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](epip.ref.Append("tags"))
}

// TotalAddressCount returns a reference to field total_address_count of aws_ec2_public_ipv4_pool.
func (epip dataEc2PublicIpv4PoolAttributes) TotalAddressCount() terra.NumberValue {
	return terra.ReferenceAsNumber(epip.ref.Append("total_address_count"))
}

// TotalAvailableAddressCount returns a reference to field total_available_address_count of aws_ec2_public_ipv4_pool.
func (epip dataEc2PublicIpv4PoolAttributes) TotalAvailableAddressCount() terra.NumberValue {
	return terra.ReferenceAsNumber(epip.ref.Append("total_available_address_count"))
}

func (epip dataEc2PublicIpv4PoolAttributes) PoolAddressRanges() terra.ListValue[dataec2publicipv4pool.PoolAddressRangesAttributes] {
	return terra.ReferenceAsList[dataec2publicipv4pool.PoolAddressRangesAttributes](epip.ref.Append("pool_address_ranges"))
}

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataNetblockIpRanges creates a new instance of [DataNetblockIpRanges].
func NewDataNetblockIpRanges(name string, args DataNetblockIpRangesArgs) *DataNetblockIpRanges {
	return &DataNetblockIpRanges{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetblockIpRanges)(nil)

// DataNetblockIpRanges represents the Terraform data resource google_netblock_ip_ranges.
type DataNetblockIpRanges struct {
	Name string
	Args DataNetblockIpRangesArgs
}

// DataSource returns the Terraform object type for [DataNetblockIpRanges].
func (nir *DataNetblockIpRanges) DataSource() string {
	return "google_netblock_ip_ranges"
}

// LocalName returns the local name for [DataNetblockIpRanges].
func (nir *DataNetblockIpRanges) LocalName() string {
	return nir.Name
}

// Configuration returns the configuration (args) for [DataNetblockIpRanges].
func (nir *DataNetblockIpRanges) Configuration() interface{} {
	return nir.Args
}

// Attributes returns the attributes for [DataNetblockIpRanges].
func (nir *DataNetblockIpRanges) Attributes() dataNetblockIpRangesAttributes {
	return dataNetblockIpRangesAttributes{ref: terra.ReferenceDataResource(nir)}
}

// DataNetblockIpRangesArgs contains the configurations for google_netblock_ip_ranges.
type DataNetblockIpRangesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RangeType: string, optional
	RangeType terra.StringValue `hcl:"range_type,attr"`
}
type dataNetblockIpRangesAttributes struct {
	ref terra.Reference
}

// CidrBlocks returns a reference to field cidr_blocks of google_netblock_ip_ranges.
func (nir dataNetblockIpRangesAttributes) CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nir.ref.Append("cidr_blocks"))
}

// CidrBlocksIpv4 returns a reference to field cidr_blocks_ipv4 of google_netblock_ip_ranges.
func (nir dataNetblockIpRangesAttributes) CidrBlocksIpv4() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nir.ref.Append("cidr_blocks_ipv4"))
}

// CidrBlocksIpv6 returns a reference to field cidr_blocks_ipv6 of google_netblock_ip_ranges.
func (nir dataNetblockIpRangesAttributes) CidrBlocksIpv6() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nir.ref.Append("cidr_blocks_ipv6"))
}

// Id returns a reference to field id of google_netblock_ip_ranges.
func (nir dataNetblockIpRangesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nir.ref.Append("id"))
}

// RangeType returns a reference to field range_type of google_netblock_ip_ranges.
func (nir dataNetblockIpRangesAttributes) RangeType() terra.StringValue {
	return terra.ReferenceAsString(nir.ref.Append("range_type"))
}

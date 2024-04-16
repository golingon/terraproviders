// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_outposts_asset

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_outposts_asset.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aoa *DataSource) DataSource() string {
	return "aws_outposts_asset"
}

// LocalName returns the local name for [DataSource].
func (aoa *DataSource) LocalName() string {
	return aoa.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aoa *DataSource) Configuration() interface{} {
	return aoa.Args
}

// Attributes returns the attributes for [DataSource].
func (aoa *DataSource) Attributes() dataAwsOutpostsAssetAttributes {
	return dataAwsOutpostsAssetAttributes{ref: terra.ReferenceDataSource(aoa)}
}

// DataArgs contains the configurations for aws_outposts_asset.
type DataArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// AssetId: string, required
	AssetId terra.StringValue `hcl:"asset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type dataAwsOutpostsAssetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_outposts_asset.
func (aoa dataAwsOutpostsAssetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aoa.ref.Append("arn"))
}

// AssetId returns a reference to field asset_id of aws_outposts_asset.
func (aoa dataAwsOutpostsAssetAttributes) AssetId() terra.StringValue {
	return terra.ReferenceAsString(aoa.ref.Append("asset_id"))
}

// AssetType returns a reference to field asset_type of aws_outposts_asset.
func (aoa dataAwsOutpostsAssetAttributes) AssetType() terra.StringValue {
	return terra.ReferenceAsString(aoa.ref.Append("asset_type"))
}

// HostId returns a reference to field host_id of aws_outposts_asset.
func (aoa dataAwsOutpostsAssetAttributes) HostId() terra.StringValue {
	return terra.ReferenceAsString(aoa.ref.Append("host_id"))
}

// Id returns a reference to field id of aws_outposts_asset.
func (aoa dataAwsOutpostsAssetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aoa.ref.Append("id"))
}

// RackElevation returns a reference to field rack_elevation of aws_outposts_asset.
func (aoa dataAwsOutpostsAssetAttributes) RackElevation() terra.NumberValue {
	return terra.ReferenceAsNumber(aoa.ref.Append("rack_elevation"))
}

// RackId returns a reference to field rack_id of aws_outposts_asset.
func (aoa dataAwsOutpostsAssetAttributes) RackId() terra.StringValue {
	return terra.ReferenceAsString(aoa.ref.Append("rack_id"))
}

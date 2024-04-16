// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataplex_asset_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_dataplex_asset_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gdaip *DataSource) DataSource() string {
	return "google_dataplex_asset_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gdaip *DataSource) LocalName() string {
	return gdaip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gdaip *DataSource) Configuration() interface{} {
	return gdaip.Args
}

// Attributes returns the attributes for [DataSource].
func (gdaip *DataSource) Attributes() dataGoogleDataplexAssetIamPolicyAttributes {
	return dataGoogleDataplexAssetIamPolicyAttributes{ref: terra.ReferenceDataSource(gdaip)}
}

// DataArgs contains the configurations for google_dataplex_asset_iam_policy.
type DataArgs struct {
	// Asset: string, required
	Asset terra.StringValue `hcl:"asset,attr" validate:"required"`
	// DataplexZone: string, required
	DataplexZone terra.StringValue `hcl:"dataplex_zone,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type dataGoogleDataplexAssetIamPolicyAttributes struct {
	ref terra.Reference
}

// Asset returns a reference to field asset of google_dataplex_asset_iam_policy.
func (gdaip dataGoogleDataplexAssetIamPolicyAttributes) Asset() terra.StringValue {
	return terra.ReferenceAsString(gdaip.ref.Append("asset"))
}

// DataplexZone returns a reference to field dataplex_zone of google_dataplex_asset_iam_policy.
func (gdaip dataGoogleDataplexAssetIamPolicyAttributes) DataplexZone() terra.StringValue {
	return terra.ReferenceAsString(gdaip.ref.Append("dataplex_zone"))
}

// Etag returns a reference to field etag of google_dataplex_asset_iam_policy.
func (gdaip dataGoogleDataplexAssetIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gdaip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_asset_iam_policy.
func (gdaip dataGoogleDataplexAssetIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdaip.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_asset_iam_policy.
func (gdaip dataGoogleDataplexAssetIamPolicyAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(gdaip.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_asset_iam_policy.
func (gdaip dataGoogleDataplexAssetIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gdaip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataplex_asset_iam_policy.
func (gdaip dataGoogleDataplexAssetIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gdaip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataplex_asset_iam_policy.
func (gdaip dataGoogleDataplexAssetIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdaip.ref.Append("project"))
}

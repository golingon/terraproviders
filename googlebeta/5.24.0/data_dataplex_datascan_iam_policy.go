// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataDataplexDatascanIamPolicy creates a new instance of [DataDataplexDatascanIamPolicy].
func NewDataDataplexDatascanIamPolicy(name string, args DataDataplexDatascanIamPolicyArgs) *DataDataplexDatascanIamPolicy {
	return &DataDataplexDatascanIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataplexDatascanIamPolicy)(nil)

// DataDataplexDatascanIamPolicy represents the Terraform data resource google_dataplex_datascan_iam_policy.
type DataDataplexDatascanIamPolicy struct {
	Name string
	Args DataDataplexDatascanIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataDataplexDatascanIamPolicy].
func (ddip *DataDataplexDatascanIamPolicy) DataSource() string {
	return "google_dataplex_datascan_iam_policy"
}

// LocalName returns the local name for [DataDataplexDatascanIamPolicy].
func (ddip *DataDataplexDatascanIamPolicy) LocalName() string {
	return ddip.Name
}

// Configuration returns the configuration (args) for [DataDataplexDatascanIamPolicy].
func (ddip *DataDataplexDatascanIamPolicy) Configuration() interface{} {
	return ddip.Args
}

// Attributes returns the attributes for [DataDataplexDatascanIamPolicy].
func (ddip *DataDataplexDatascanIamPolicy) Attributes() dataDataplexDatascanIamPolicyAttributes {
	return dataDataplexDatascanIamPolicyAttributes{ref: terra.ReferenceDataResource(ddip)}
}

// DataDataplexDatascanIamPolicyArgs contains the configurations for google_dataplex_datascan_iam_policy.
type DataDataplexDatascanIamPolicyArgs struct {
	// DataScanId: string, required
	DataScanId terra.StringValue `hcl:"data_scan_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataDataplexDatascanIamPolicyAttributes struct {
	ref terra.Reference
}

// DataScanId returns a reference to field data_scan_id of google_dataplex_datascan_iam_policy.
func (ddip dataDataplexDatascanIamPolicyAttributes) DataScanId() terra.StringValue {
	return terra.ReferenceAsString(ddip.ref.Append("data_scan_id"))
}

// Etag returns a reference to field etag of google_dataplex_datascan_iam_policy.
func (ddip dataDataplexDatascanIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ddip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_datascan_iam_policy.
func (ddip dataDataplexDatascanIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ddip.ref.Append("id"))
}

// Location returns a reference to field location of google_dataplex_datascan_iam_policy.
func (ddip dataDataplexDatascanIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ddip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataplex_datascan_iam_policy.
func (ddip dataDataplexDatascanIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ddip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataplex_datascan_iam_policy.
func (ddip dataDataplexDatascanIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ddip.ref.Append("project"))
}

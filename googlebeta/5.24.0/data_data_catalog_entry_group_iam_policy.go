// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataDataCatalogEntryGroupIamPolicy creates a new instance of [DataDataCatalogEntryGroupIamPolicy].
func NewDataDataCatalogEntryGroupIamPolicy(name string, args DataDataCatalogEntryGroupIamPolicyArgs) *DataDataCatalogEntryGroupIamPolicy {
	return &DataDataCatalogEntryGroupIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataCatalogEntryGroupIamPolicy)(nil)

// DataDataCatalogEntryGroupIamPolicy represents the Terraform data resource google_data_catalog_entry_group_iam_policy.
type DataDataCatalogEntryGroupIamPolicy struct {
	Name string
	Args DataDataCatalogEntryGroupIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataDataCatalogEntryGroupIamPolicy].
func (dcegip *DataDataCatalogEntryGroupIamPolicy) DataSource() string {
	return "google_data_catalog_entry_group_iam_policy"
}

// LocalName returns the local name for [DataDataCatalogEntryGroupIamPolicy].
func (dcegip *DataDataCatalogEntryGroupIamPolicy) LocalName() string {
	return dcegip.Name
}

// Configuration returns the configuration (args) for [DataDataCatalogEntryGroupIamPolicy].
func (dcegip *DataDataCatalogEntryGroupIamPolicy) Configuration() interface{} {
	return dcegip.Args
}

// Attributes returns the attributes for [DataDataCatalogEntryGroupIamPolicy].
func (dcegip *DataDataCatalogEntryGroupIamPolicy) Attributes() dataDataCatalogEntryGroupIamPolicyAttributes {
	return dataDataCatalogEntryGroupIamPolicyAttributes{ref: terra.ReferenceDataResource(dcegip)}
}

// DataDataCatalogEntryGroupIamPolicyArgs contains the configurations for google_data_catalog_entry_group_iam_policy.
type DataDataCatalogEntryGroupIamPolicyArgs struct {
	// EntryGroup: string, required
	EntryGroup terra.StringValue `hcl:"entry_group,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataDataCatalogEntryGroupIamPolicyAttributes struct {
	ref terra.Reference
}

// EntryGroup returns a reference to field entry_group of google_data_catalog_entry_group_iam_policy.
func (dcegip dataDataCatalogEntryGroupIamPolicyAttributes) EntryGroup() terra.StringValue {
	return terra.ReferenceAsString(dcegip.ref.Append("entry_group"))
}

// Etag returns a reference to field etag of google_data_catalog_entry_group_iam_policy.
func (dcegip dataDataCatalogEntryGroupIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcegip.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_entry_group_iam_policy.
func (dcegip dataDataCatalogEntryGroupIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcegip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_data_catalog_entry_group_iam_policy.
func (dcegip dataDataCatalogEntryGroupIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dcegip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_data_catalog_entry_group_iam_policy.
func (dcegip dataDataCatalogEntryGroupIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dcegip.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_entry_group_iam_policy.
func (dcegip dataDataCatalogEntryGroupIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dcegip.ref.Append("region"))
}

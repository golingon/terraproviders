// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDataCatalogTaxonomyIamPolicy creates a new instance of [DataDataCatalogTaxonomyIamPolicy].
func NewDataDataCatalogTaxonomyIamPolicy(name string, args DataDataCatalogTaxonomyIamPolicyArgs) *DataDataCatalogTaxonomyIamPolicy {
	return &DataDataCatalogTaxonomyIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataCatalogTaxonomyIamPolicy)(nil)

// DataDataCatalogTaxonomyIamPolicy represents the Terraform data resource google_data_catalog_taxonomy_iam_policy.
type DataDataCatalogTaxonomyIamPolicy struct {
	Name string
	Args DataDataCatalogTaxonomyIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataDataCatalogTaxonomyIamPolicy].
func (dctip *DataDataCatalogTaxonomyIamPolicy) DataSource() string {
	return "google_data_catalog_taxonomy_iam_policy"
}

// LocalName returns the local name for [DataDataCatalogTaxonomyIamPolicy].
func (dctip *DataDataCatalogTaxonomyIamPolicy) LocalName() string {
	return dctip.Name
}

// Configuration returns the configuration (args) for [DataDataCatalogTaxonomyIamPolicy].
func (dctip *DataDataCatalogTaxonomyIamPolicy) Configuration() interface{} {
	return dctip.Args
}

// Attributes returns the attributes for [DataDataCatalogTaxonomyIamPolicy].
func (dctip *DataDataCatalogTaxonomyIamPolicy) Attributes() dataDataCatalogTaxonomyIamPolicyAttributes {
	return dataDataCatalogTaxonomyIamPolicyAttributes{ref: terra.ReferenceDataResource(dctip)}
}

// DataDataCatalogTaxonomyIamPolicyArgs contains the configurations for google_data_catalog_taxonomy_iam_policy.
type DataDataCatalogTaxonomyIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Taxonomy: string, required
	Taxonomy terra.StringValue `hcl:"taxonomy,attr" validate:"required"`
}
type dataDataCatalogTaxonomyIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_catalog_taxonomy_iam_policy.
func (dctip dataDataCatalogTaxonomyIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dctip.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_taxonomy_iam_policy.
func (dctip dataDataCatalogTaxonomyIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dctip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_data_catalog_taxonomy_iam_policy.
func (dctip dataDataCatalogTaxonomyIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dctip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_data_catalog_taxonomy_iam_policy.
func (dctip dataDataCatalogTaxonomyIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dctip.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_taxonomy_iam_policy.
func (dctip dataDataCatalogTaxonomyIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dctip.ref.Append("region"))
}

// Taxonomy returns a reference to field taxonomy of google_data_catalog_taxonomy_iam_policy.
func (dctip dataDataCatalogTaxonomyIamPolicyAttributes) Taxonomy() terra.StringValue {
	return terra.ReferenceAsString(dctip.ref.Append("taxonomy"))
}

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataBigtableTableIamPolicy creates a new instance of [DataBigtableTableIamPolicy].
func NewDataBigtableTableIamPolicy(name string, args DataBigtableTableIamPolicyArgs) *DataBigtableTableIamPolicy {
	return &DataBigtableTableIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBigtableTableIamPolicy)(nil)

// DataBigtableTableIamPolicy represents the Terraform data resource google_bigtable_table_iam_policy.
type DataBigtableTableIamPolicy struct {
	Name string
	Args DataBigtableTableIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataBigtableTableIamPolicy].
func (btip *DataBigtableTableIamPolicy) DataSource() string {
	return "google_bigtable_table_iam_policy"
}

// LocalName returns the local name for [DataBigtableTableIamPolicy].
func (btip *DataBigtableTableIamPolicy) LocalName() string {
	return btip.Name
}

// Configuration returns the configuration (args) for [DataBigtableTableIamPolicy].
func (btip *DataBigtableTableIamPolicy) Configuration() interface{} {
	return btip.Args
}

// Attributes returns the attributes for [DataBigtableTableIamPolicy].
func (btip *DataBigtableTableIamPolicy) Attributes() dataBigtableTableIamPolicyAttributes {
	return dataBigtableTableIamPolicyAttributes{ref: terra.ReferenceDataResource(btip)}
}

// DataBigtableTableIamPolicyArgs contains the configurations for google_bigtable_table_iam_policy.
type DataBigtableTableIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Table: string, required
	Table terra.StringValue `hcl:"table,attr" validate:"required"`
}
type dataBigtableTableIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_bigtable_table_iam_policy.
func (btip dataBigtableTableIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigtable_table_iam_policy.
func (btip dataBigtableTableIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("id"))
}

// Instance returns a reference to field instance of google_bigtable_table_iam_policy.
func (btip dataBigtableTableIamPolicyAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("instance"))
}

// PolicyData returns a reference to field policy_data of google_bigtable_table_iam_policy.
func (btip dataBigtableTableIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_bigtable_table_iam_policy.
func (btip dataBigtableTableIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("project"))
}

// Table returns a reference to field table of google_bigtable_table_iam_policy.
func (btip dataBigtableTableIamPolicyAttributes) Table() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("table"))
}

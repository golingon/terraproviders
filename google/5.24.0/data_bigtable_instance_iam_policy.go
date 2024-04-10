// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataBigtableInstanceIamPolicy creates a new instance of [DataBigtableInstanceIamPolicy].
func NewDataBigtableInstanceIamPolicy(name string, args DataBigtableInstanceIamPolicyArgs) *DataBigtableInstanceIamPolicy {
	return &DataBigtableInstanceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBigtableInstanceIamPolicy)(nil)

// DataBigtableInstanceIamPolicy represents the Terraform data resource google_bigtable_instance_iam_policy.
type DataBigtableInstanceIamPolicy struct {
	Name string
	Args DataBigtableInstanceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataBigtableInstanceIamPolicy].
func (biip *DataBigtableInstanceIamPolicy) DataSource() string {
	return "google_bigtable_instance_iam_policy"
}

// LocalName returns the local name for [DataBigtableInstanceIamPolicy].
func (biip *DataBigtableInstanceIamPolicy) LocalName() string {
	return biip.Name
}

// Configuration returns the configuration (args) for [DataBigtableInstanceIamPolicy].
func (biip *DataBigtableInstanceIamPolicy) Configuration() interface{} {
	return biip.Args
}

// Attributes returns the attributes for [DataBigtableInstanceIamPolicy].
func (biip *DataBigtableInstanceIamPolicy) Attributes() dataBigtableInstanceIamPolicyAttributes {
	return dataBigtableInstanceIamPolicyAttributes{ref: terra.ReferenceDataResource(biip)}
}

// DataBigtableInstanceIamPolicyArgs contains the configurations for google_bigtable_instance_iam_policy.
type DataBigtableInstanceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataBigtableInstanceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_bigtable_instance_iam_policy.
func (biip dataBigtableInstanceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(biip.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigtable_instance_iam_policy.
func (biip dataBigtableInstanceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(biip.ref.Append("id"))
}

// Instance returns a reference to field instance of google_bigtable_instance_iam_policy.
func (biip dataBigtableInstanceIamPolicyAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(biip.ref.Append("instance"))
}

// PolicyData returns a reference to field policy_data of google_bigtable_instance_iam_policy.
func (biip dataBigtableInstanceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(biip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_bigtable_instance_iam_policy.
func (biip dataBigtableInstanceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(biip.ref.Append("project"))
}

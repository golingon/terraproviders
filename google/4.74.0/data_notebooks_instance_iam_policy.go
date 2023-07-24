// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataNotebooksInstanceIamPolicy creates a new instance of [DataNotebooksInstanceIamPolicy].
func NewDataNotebooksInstanceIamPolicy(name string, args DataNotebooksInstanceIamPolicyArgs) *DataNotebooksInstanceIamPolicy {
	return &DataNotebooksInstanceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNotebooksInstanceIamPolicy)(nil)

// DataNotebooksInstanceIamPolicy represents the Terraform data resource google_notebooks_instance_iam_policy.
type DataNotebooksInstanceIamPolicy struct {
	Name string
	Args DataNotebooksInstanceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataNotebooksInstanceIamPolicy].
func (niip *DataNotebooksInstanceIamPolicy) DataSource() string {
	return "google_notebooks_instance_iam_policy"
}

// LocalName returns the local name for [DataNotebooksInstanceIamPolicy].
func (niip *DataNotebooksInstanceIamPolicy) LocalName() string {
	return niip.Name
}

// Configuration returns the configuration (args) for [DataNotebooksInstanceIamPolicy].
func (niip *DataNotebooksInstanceIamPolicy) Configuration() interface{} {
	return niip.Args
}

// Attributes returns the attributes for [DataNotebooksInstanceIamPolicy].
func (niip *DataNotebooksInstanceIamPolicy) Attributes() dataNotebooksInstanceIamPolicyAttributes {
	return dataNotebooksInstanceIamPolicyAttributes{ref: terra.ReferenceDataResource(niip)}
}

// DataNotebooksInstanceIamPolicyArgs contains the configurations for google_notebooks_instance_iam_policy.
type DataNotebooksInstanceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataNotebooksInstanceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_notebooks_instance_iam_policy.
func (niip dataNotebooksInstanceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(niip.ref.Append("etag"))
}

// Id returns a reference to field id of google_notebooks_instance_iam_policy.
func (niip dataNotebooksInstanceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(niip.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of google_notebooks_instance_iam_policy.
func (niip dataNotebooksInstanceIamPolicyAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(niip.ref.Append("instance_name"))
}

// Location returns a reference to field location of google_notebooks_instance_iam_policy.
func (niip dataNotebooksInstanceIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(niip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_notebooks_instance_iam_policy.
func (niip dataNotebooksInstanceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(niip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_notebooks_instance_iam_policy.
func (niip dataNotebooksInstanceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(niip.ref.Append("project"))
}

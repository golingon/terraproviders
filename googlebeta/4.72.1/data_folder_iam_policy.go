// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataFolderIamPolicy creates a new instance of [DataFolderIamPolicy].
func NewDataFolderIamPolicy(name string, args DataFolderIamPolicyArgs) *DataFolderIamPolicy {
	return &DataFolderIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataFolderIamPolicy)(nil)

// DataFolderIamPolicy represents the Terraform data resource google_folder_iam_policy.
type DataFolderIamPolicy struct {
	Name string
	Args DataFolderIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataFolderIamPolicy].
func (fip *DataFolderIamPolicy) DataSource() string {
	return "google_folder_iam_policy"
}

// LocalName returns the local name for [DataFolderIamPolicy].
func (fip *DataFolderIamPolicy) LocalName() string {
	return fip.Name
}

// Configuration returns the configuration (args) for [DataFolderIamPolicy].
func (fip *DataFolderIamPolicy) Configuration() interface{} {
	return fip.Args
}

// Attributes returns the attributes for [DataFolderIamPolicy].
func (fip *DataFolderIamPolicy) Attributes() dataFolderIamPolicyAttributes {
	return dataFolderIamPolicyAttributes{ref: terra.ReferenceDataResource(fip)}
}

// DataFolderIamPolicyArgs contains the configurations for google_folder_iam_policy.
type DataFolderIamPolicyArgs struct {
	// Folder: string, required
	Folder terra.StringValue `hcl:"folder,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataFolderIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_folder_iam_policy.
func (fip dataFolderIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(fip.ref.Append("etag"))
}

// Folder returns a reference to field folder of google_folder_iam_policy.
func (fip dataFolderIamPolicyAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(fip.ref.Append("folder"))
}

// Id returns a reference to field id of google_folder_iam_policy.
func (fip dataFolderIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_folder_iam_policy.
func (fip dataFolderIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(fip.ref.Append("policy_data"))
}

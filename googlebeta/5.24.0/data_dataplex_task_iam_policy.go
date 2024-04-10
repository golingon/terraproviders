// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataDataplexTaskIamPolicy creates a new instance of [DataDataplexTaskIamPolicy].
func NewDataDataplexTaskIamPolicy(name string, args DataDataplexTaskIamPolicyArgs) *DataDataplexTaskIamPolicy {
	return &DataDataplexTaskIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataplexTaskIamPolicy)(nil)

// DataDataplexTaskIamPolicy represents the Terraform data resource google_dataplex_task_iam_policy.
type DataDataplexTaskIamPolicy struct {
	Name string
	Args DataDataplexTaskIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataDataplexTaskIamPolicy].
func (dtip *DataDataplexTaskIamPolicy) DataSource() string {
	return "google_dataplex_task_iam_policy"
}

// LocalName returns the local name for [DataDataplexTaskIamPolicy].
func (dtip *DataDataplexTaskIamPolicy) LocalName() string {
	return dtip.Name
}

// Configuration returns the configuration (args) for [DataDataplexTaskIamPolicy].
func (dtip *DataDataplexTaskIamPolicy) Configuration() interface{} {
	return dtip.Args
}

// Attributes returns the attributes for [DataDataplexTaskIamPolicy].
func (dtip *DataDataplexTaskIamPolicy) Attributes() dataDataplexTaskIamPolicyAttributes {
	return dataDataplexTaskIamPolicyAttributes{ref: terra.ReferenceDataResource(dtip)}
}

// DataDataplexTaskIamPolicyArgs contains the configurations for google_dataplex_task_iam_policy.
type DataDataplexTaskIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TaskId: string, required
	TaskId terra.StringValue `hcl:"task_id,attr" validate:"required"`
}
type dataDataplexTaskIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataplex_task_iam_policy.
func (dtip dataDataplexTaskIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_task_iam_policy.
func (dtip dataDataplexTaskIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_task_iam_policy.
func (dtip dataDataplexTaskIamPolicyAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_task_iam_policy.
func (dtip dataDataplexTaskIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataplex_task_iam_policy.
func (dtip dataDataplexTaskIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataplex_task_iam_policy.
func (dtip dataDataplexTaskIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("project"))
}

// TaskId returns a reference to field task_id of google_dataplex_task_iam_policy.
func (dtip dataDataplexTaskIamPolicyAttributes) TaskId() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("task_id"))
}

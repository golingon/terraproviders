// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataCloudTasksQueueIamPolicy creates a new instance of [DataCloudTasksQueueIamPolicy].
func NewDataCloudTasksQueueIamPolicy(name string, args DataCloudTasksQueueIamPolicyArgs) *DataCloudTasksQueueIamPolicy {
	return &DataCloudTasksQueueIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudTasksQueueIamPolicy)(nil)

// DataCloudTasksQueueIamPolicy represents the Terraform data resource google_cloud_tasks_queue_iam_policy.
type DataCloudTasksQueueIamPolicy struct {
	Name string
	Args DataCloudTasksQueueIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataCloudTasksQueueIamPolicy].
func (ctqip *DataCloudTasksQueueIamPolicy) DataSource() string {
	return "google_cloud_tasks_queue_iam_policy"
}

// LocalName returns the local name for [DataCloudTasksQueueIamPolicy].
func (ctqip *DataCloudTasksQueueIamPolicy) LocalName() string {
	return ctqip.Name
}

// Configuration returns the configuration (args) for [DataCloudTasksQueueIamPolicy].
func (ctqip *DataCloudTasksQueueIamPolicy) Configuration() interface{} {
	return ctqip.Args
}

// Attributes returns the attributes for [DataCloudTasksQueueIamPolicy].
func (ctqip *DataCloudTasksQueueIamPolicy) Attributes() dataCloudTasksQueueIamPolicyAttributes {
	return dataCloudTasksQueueIamPolicyAttributes{ref: terra.ReferenceDataResource(ctqip)}
}

// DataCloudTasksQueueIamPolicyArgs contains the configurations for google_cloud_tasks_queue_iam_policy.
type DataCloudTasksQueueIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataCloudTasksQueueIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_tasks_queue_iam_policy.
func (ctqip dataCloudTasksQueueIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ctqip.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_tasks_queue_iam_policy.
func (ctqip dataCloudTasksQueueIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ctqip.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_tasks_queue_iam_policy.
func (ctqip dataCloudTasksQueueIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ctqip.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_tasks_queue_iam_policy.
func (ctqip dataCloudTasksQueueIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ctqip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_cloud_tasks_queue_iam_policy.
func (ctqip dataCloudTasksQueueIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ctqip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_cloud_tasks_queue_iam_policy.
func (ctqip dataCloudTasksQueueIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ctqip.ref.Append("project"))
}

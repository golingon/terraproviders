// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataproc_job_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_dataproc_job_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gdjip *DataSource) DataSource() string {
	return "google_dataproc_job_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gdjip *DataSource) LocalName() string {
	return gdjip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gdjip *DataSource) Configuration() interface{} {
	return gdjip.Args
}

// Attributes returns the attributes for [DataSource].
func (gdjip *DataSource) Attributes() dataGoogleDataprocJobIamPolicyAttributes {
	return dataGoogleDataprocJobIamPolicyAttributes{ref: terra.ReferenceDataSource(gdjip)}
}

// DataArgs contains the configurations for google_dataproc_job_iam_policy.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JobId: string, required
	JobId terra.StringValue `hcl:"job_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}

type dataGoogleDataprocJobIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_job_iam_policy.
func (gdjip dataGoogleDataprocJobIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gdjip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_job_iam_policy.
func (gdjip dataGoogleDataprocJobIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdjip.ref.Append("id"))
}

// JobId returns a reference to field job_id of google_dataproc_job_iam_policy.
func (gdjip dataGoogleDataprocJobIamPolicyAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(gdjip.ref.Append("job_id"))
}

// PolicyData returns a reference to field policy_data of google_dataproc_job_iam_policy.
func (gdjip dataGoogleDataprocJobIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gdjip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataproc_job_iam_policy.
func (gdjip dataGoogleDataprocJobIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdjip.ref.Append("project"))
}

// Region returns a reference to field region of google_dataproc_job_iam_policy.
func (gdjip dataGoogleDataprocJobIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gdjip.ref.Append("region"))
}

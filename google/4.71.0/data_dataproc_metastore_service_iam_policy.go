// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDataprocMetastoreServiceIamPolicy creates a new instance of [DataDataprocMetastoreServiceIamPolicy].
func NewDataDataprocMetastoreServiceIamPolicy(name string, args DataDataprocMetastoreServiceIamPolicyArgs) *DataDataprocMetastoreServiceIamPolicy {
	return &DataDataprocMetastoreServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataprocMetastoreServiceIamPolicy)(nil)

// DataDataprocMetastoreServiceIamPolicy represents the Terraform data resource google_dataproc_metastore_service_iam_policy.
type DataDataprocMetastoreServiceIamPolicy struct {
	Name string
	Args DataDataprocMetastoreServiceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataDataprocMetastoreServiceIamPolicy].
func (dmsip *DataDataprocMetastoreServiceIamPolicy) DataSource() string {
	return "google_dataproc_metastore_service_iam_policy"
}

// LocalName returns the local name for [DataDataprocMetastoreServiceIamPolicy].
func (dmsip *DataDataprocMetastoreServiceIamPolicy) LocalName() string {
	return dmsip.Name
}

// Configuration returns the configuration (args) for [DataDataprocMetastoreServiceIamPolicy].
func (dmsip *DataDataprocMetastoreServiceIamPolicy) Configuration() interface{} {
	return dmsip.Args
}

// Attributes returns the attributes for [DataDataprocMetastoreServiceIamPolicy].
func (dmsip *DataDataprocMetastoreServiceIamPolicy) Attributes() dataDataprocMetastoreServiceIamPolicyAttributes {
	return dataDataprocMetastoreServiceIamPolicyAttributes{ref: terra.ReferenceDataResource(dmsip)}
}

// DataDataprocMetastoreServiceIamPolicyArgs contains the configurations for google_dataproc_metastore_service_iam_policy.
type DataDataprocMetastoreServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceId: string, required
	ServiceId terra.StringValue `hcl:"service_id,attr" validate:"required"`
}
type dataDataprocMetastoreServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_metastore_service_iam_policy.
func (dmsip dataDataprocMetastoreServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_metastore_service_iam_policy.
func (dmsip dataDataprocMetastoreServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_metastore_service_iam_policy.
func (dmsip dataDataprocMetastoreServiceIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataproc_metastore_service_iam_policy.
func (dmsip dataDataprocMetastoreServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataproc_metastore_service_iam_policy.
func (dmsip dataDataprocMetastoreServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("project"))
}

// ServiceId returns a reference to field service_id of google_dataproc_metastore_service_iam_policy.
func (dmsip dataDataprocMetastoreServiceIamPolicyAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("service_id"))
}

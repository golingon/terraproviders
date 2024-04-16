// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_secure_source_manager_instance_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_secure_source_manager_instance_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gssmiip *DataSource) DataSource() string {
	return "google_secure_source_manager_instance_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gssmiip *DataSource) LocalName() string {
	return gssmiip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gssmiip *DataSource) Configuration() interface{} {
	return gssmiip.Args
}

// Attributes returns the attributes for [DataSource].
func (gssmiip *DataSource) Attributes() dataGoogleSecureSourceManagerInstanceIamPolicyAttributes {
	return dataGoogleSecureSourceManagerInstanceIamPolicyAttributes{ref: terra.ReferenceDataSource(gssmiip)}
}

// DataArgs contains the configurations for google_secure_source_manager_instance_iam_policy.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type dataGoogleSecureSourceManagerInstanceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_secure_source_manager_instance_iam_policy.
func (gssmiip dataGoogleSecureSourceManagerInstanceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gssmiip.ref.Append("etag"))
}

// Id returns a reference to field id of google_secure_source_manager_instance_iam_policy.
func (gssmiip dataGoogleSecureSourceManagerInstanceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gssmiip.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_secure_source_manager_instance_iam_policy.
func (gssmiip dataGoogleSecureSourceManagerInstanceIamPolicyAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(gssmiip.ref.Append("instance_id"))
}

// Location returns a reference to field location of google_secure_source_manager_instance_iam_policy.
func (gssmiip dataGoogleSecureSourceManagerInstanceIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gssmiip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_secure_source_manager_instance_iam_policy.
func (gssmiip dataGoogleSecureSourceManagerInstanceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gssmiip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_secure_source_manager_instance_iam_policy.
func (gssmiip dataGoogleSecureSourceManagerInstanceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gssmiip.ref.Append("project"))
}

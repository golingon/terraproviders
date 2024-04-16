// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_spanner_database_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_spanner_database_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gsdip *DataSource) DataSource() string {
	return "google_spanner_database_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gsdip *DataSource) LocalName() string {
	return gsdip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gsdip *DataSource) Configuration() interface{} {
	return gsdip.Args
}

// Attributes returns the attributes for [DataSource].
func (gsdip *DataSource) Attributes() dataGoogleSpannerDatabaseIamPolicyAttributes {
	return dataGoogleSpannerDatabaseIamPolicyAttributes{ref: terra.ReferenceDataSource(gsdip)}
}

// DataArgs contains the configurations for google_spanner_database_iam_policy.
type DataArgs struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type dataGoogleSpannerDatabaseIamPolicyAttributes struct {
	ref terra.Reference
}

// Database returns a reference to field database of google_spanner_database_iam_policy.
func (gsdip dataGoogleSpannerDatabaseIamPolicyAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(gsdip.ref.Append("database"))
}

// Etag returns a reference to field etag of google_spanner_database_iam_policy.
func (gsdip dataGoogleSpannerDatabaseIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gsdip.ref.Append("etag"))
}

// Id returns a reference to field id of google_spanner_database_iam_policy.
func (gsdip dataGoogleSpannerDatabaseIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsdip.ref.Append("id"))
}

// Instance returns a reference to field instance of google_spanner_database_iam_policy.
func (gsdip dataGoogleSpannerDatabaseIamPolicyAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(gsdip.ref.Append("instance"))
}

// PolicyData returns a reference to field policy_data of google_spanner_database_iam_policy.
func (gsdip dataGoogleSpannerDatabaseIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gsdip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_spanner_database_iam_policy.
func (gsdip dataGoogleSpannerDatabaseIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gsdip.ref.Append("project"))
}

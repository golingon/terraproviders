// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_storage_project_service_account

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_storage_project_service_account.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gspsa *DataSource) DataSource() string {
	return "google_storage_project_service_account"
}

// LocalName returns the local name for [DataSource].
func (gspsa *DataSource) LocalName() string {
	return gspsa.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gspsa *DataSource) Configuration() interface{} {
	return gspsa.Args
}

// Attributes returns the attributes for [DataSource].
func (gspsa *DataSource) Attributes() dataGoogleStorageProjectServiceAccountAttributes {
	return dataGoogleStorageProjectServiceAccountAttributes{ref: terra.ReferenceDataSource(gspsa)}
}

// DataArgs contains the configurations for google_storage_project_service_account.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// UserProject: string, optional
	UserProject terra.StringValue `hcl:"user_project,attr"`
}

type dataGoogleStorageProjectServiceAccountAttributes struct {
	ref terra.Reference
}

// EmailAddress returns a reference to field email_address of google_storage_project_service_account.
func (gspsa dataGoogleStorageProjectServiceAccountAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceAsString(gspsa.ref.Append("email_address"))
}

// Id returns a reference to field id of google_storage_project_service_account.
func (gspsa dataGoogleStorageProjectServiceAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gspsa.ref.Append("id"))
}

// Member returns a reference to field member of google_storage_project_service_account.
func (gspsa dataGoogleStorageProjectServiceAccountAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(gspsa.ref.Append("member"))
}

// Project returns a reference to field project of google_storage_project_service_account.
func (gspsa dataGoogleStorageProjectServiceAccountAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gspsa.ref.Append("project"))
}

// UserProject returns a reference to field user_project of google_storage_project_service_account.
func (gspsa dataGoogleStorageProjectServiceAccountAttributes) UserProject() terra.StringValue {
	return terra.ReferenceAsString(gspsa.ref.Append("user_project"))
}

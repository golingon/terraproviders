// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_storage_transfer_project_service_account

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_storage_transfer_project_service_account.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gstpsa *DataSource) DataSource() string {
	return "google_storage_transfer_project_service_account"
}

// LocalName returns the local name for [DataSource].
func (gstpsa *DataSource) LocalName() string {
	return gstpsa.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gstpsa *DataSource) Configuration() interface{} {
	return gstpsa.Args
}

// Attributes returns the attributes for [DataSource].
func (gstpsa *DataSource) Attributes() dataGoogleStorageTransferProjectServiceAccountAttributes {
	return dataGoogleStorageTransferProjectServiceAccountAttributes{ref: terra.ReferenceDataSource(gstpsa)}
}

// DataArgs contains the configurations for google_storage_transfer_project_service_account.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type dataGoogleStorageTransferProjectServiceAccountAttributes struct {
	ref terra.Reference
}

// Email returns a reference to field email of google_storage_transfer_project_service_account.
func (gstpsa dataGoogleStorageTransferProjectServiceAccountAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(gstpsa.ref.Append("email"))
}

// Id returns a reference to field id of google_storage_transfer_project_service_account.
func (gstpsa dataGoogleStorageTransferProjectServiceAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gstpsa.ref.Append("id"))
}

// Member returns a reference to field member of google_storage_transfer_project_service_account.
func (gstpsa dataGoogleStorageTransferProjectServiceAccountAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(gstpsa.ref.Append("member"))
}

// Project returns a reference to field project of google_storage_transfer_project_service_account.
func (gstpsa dataGoogleStorageTransferProjectServiceAccountAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gstpsa.ref.Append("project"))
}

// SubjectId returns a reference to field subject_id of google_storage_transfer_project_service_account.
func (gstpsa dataGoogleStorageTransferProjectServiceAccountAttributes) SubjectId() terra.StringValue {
	return terra.ReferenceAsString(gstpsa.ref.Append("subject_id"))
}

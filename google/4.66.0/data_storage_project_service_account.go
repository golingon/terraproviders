// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataStorageProjectServiceAccount creates a new instance of [DataStorageProjectServiceAccount].
func NewDataStorageProjectServiceAccount(name string, args DataStorageProjectServiceAccountArgs) *DataStorageProjectServiceAccount {
	return &DataStorageProjectServiceAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageProjectServiceAccount)(nil)

// DataStorageProjectServiceAccount represents the Terraform data resource google_storage_project_service_account.
type DataStorageProjectServiceAccount struct {
	Name string
	Args DataStorageProjectServiceAccountArgs
}

// DataSource returns the Terraform object type for [DataStorageProjectServiceAccount].
func (spsa *DataStorageProjectServiceAccount) DataSource() string {
	return "google_storage_project_service_account"
}

// LocalName returns the local name for [DataStorageProjectServiceAccount].
func (spsa *DataStorageProjectServiceAccount) LocalName() string {
	return spsa.Name
}

// Configuration returns the configuration (args) for [DataStorageProjectServiceAccount].
func (spsa *DataStorageProjectServiceAccount) Configuration() interface{} {
	return spsa.Args
}

// Attributes returns the attributes for [DataStorageProjectServiceAccount].
func (spsa *DataStorageProjectServiceAccount) Attributes() dataStorageProjectServiceAccountAttributes {
	return dataStorageProjectServiceAccountAttributes{ref: terra.ReferenceDataResource(spsa)}
}

// DataStorageProjectServiceAccountArgs contains the configurations for google_storage_project_service_account.
type DataStorageProjectServiceAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// UserProject: string, optional
	UserProject terra.StringValue `hcl:"user_project,attr"`
}
type dataStorageProjectServiceAccountAttributes struct {
	ref terra.Reference
}

// EmailAddress returns a reference to field email_address of google_storage_project_service_account.
func (spsa dataStorageProjectServiceAccountAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceAsString(spsa.ref.Append("email_address"))
}

// Id returns a reference to field id of google_storage_project_service_account.
func (spsa dataStorageProjectServiceAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spsa.ref.Append("id"))
}

// Member returns a reference to field member of google_storage_project_service_account.
func (spsa dataStorageProjectServiceAccountAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(spsa.ref.Append("member"))
}

// Project returns a reference to field project of google_storage_project_service_account.
func (spsa dataStorageProjectServiceAccountAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(spsa.ref.Append("project"))
}

// UserProject returns a reference to field user_project of google_storage_project_service_account.
func (spsa dataStorageProjectServiceAccountAttributes) UserProject() terra.StringValue {
	return terra.ReferenceAsString(spsa.ref.Append("user_project"))
}

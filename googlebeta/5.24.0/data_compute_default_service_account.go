// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataComputeDefaultServiceAccount creates a new instance of [DataComputeDefaultServiceAccount].
func NewDataComputeDefaultServiceAccount(name string, args DataComputeDefaultServiceAccountArgs) *DataComputeDefaultServiceAccount {
	return &DataComputeDefaultServiceAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeDefaultServiceAccount)(nil)

// DataComputeDefaultServiceAccount represents the Terraform data resource google_compute_default_service_account.
type DataComputeDefaultServiceAccount struct {
	Name string
	Args DataComputeDefaultServiceAccountArgs
}

// DataSource returns the Terraform object type for [DataComputeDefaultServiceAccount].
func (cdsa *DataComputeDefaultServiceAccount) DataSource() string {
	return "google_compute_default_service_account"
}

// LocalName returns the local name for [DataComputeDefaultServiceAccount].
func (cdsa *DataComputeDefaultServiceAccount) LocalName() string {
	return cdsa.Name
}

// Configuration returns the configuration (args) for [DataComputeDefaultServiceAccount].
func (cdsa *DataComputeDefaultServiceAccount) Configuration() interface{} {
	return cdsa.Args
}

// Attributes returns the attributes for [DataComputeDefaultServiceAccount].
func (cdsa *DataComputeDefaultServiceAccount) Attributes() dataComputeDefaultServiceAccountAttributes {
	return dataComputeDefaultServiceAccountAttributes{ref: terra.ReferenceDataResource(cdsa)}
}

// DataComputeDefaultServiceAccountArgs contains the configurations for google_compute_default_service_account.
type DataComputeDefaultServiceAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataComputeDefaultServiceAccountAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_compute_default_service_account.
func (cdsa dataComputeDefaultServiceAccountAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(cdsa.ref.Append("display_name"))
}

// Email returns a reference to field email of google_compute_default_service_account.
func (cdsa dataComputeDefaultServiceAccountAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(cdsa.ref.Append("email"))
}

// Id returns a reference to field id of google_compute_default_service_account.
func (cdsa dataComputeDefaultServiceAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdsa.ref.Append("id"))
}

// Member returns a reference to field member of google_compute_default_service_account.
func (cdsa dataComputeDefaultServiceAccountAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(cdsa.ref.Append("member"))
}

// Name returns a reference to field name of google_compute_default_service_account.
func (cdsa dataComputeDefaultServiceAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cdsa.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_default_service_account.
func (cdsa dataComputeDefaultServiceAccountAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cdsa.ref.Append("project"))
}

// UniqueId returns a reference to field unique_id of google_compute_default_service_account.
func (cdsa dataComputeDefaultServiceAccountAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceAsString(cdsa.ref.Append("unique_id"))
}

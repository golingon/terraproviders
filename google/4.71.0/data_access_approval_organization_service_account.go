// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataAccessApprovalOrganizationServiceAccount creates a new instance of [DataAccessApprovalOrganizationServiceAccount].
func NewDataAccessApprovalOrganizationServiceAccount(name string, args DataAccessApprovalOrganizationServiceAccountArgs) *DataAccessApprovalOrganizationServiceAccount {
	return &DataAccessApprovalOrganizationServiceAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAccessApprovalOrganizationServiceAccount)(nil)

// DataAccessApprovalOrganizationServiceAccount represents the Terraform data resource google_access_approval_organization_service_account.
type DataAccessApprovalOrganizationServiceAccount struct {
	Name string
	Args DataAccessApprovalOrganizationServiceAccountArgs
}

// DataSource returns the Terraform object type for [DataAccessApprovalOrganizationServiceAccount].
func (aaosa *DataAccessApprovalOrganizationServiceAccount) DataSource() string {
	return "google_access_approval_organization_service_account"
}

// LocalName returns the local name for [DataAccessApprovalOrganizationServiceAccount].
func (aaosa *DataAccessApprovalOrganizationServiceAccount) LocalName() string {
	return aaosa.Name
}

// Configuration returns the configuration (args) for [DataAccessApprovalOrganizationServiceAccount].
func (aaosa *DataAccessApprovalOrganizationServiceAccount) Configuration() interface{} {
	return aaosa.Args
}

// Attributes returns the attributes for [DataAccessApprovalOrganizationServiceAccount].
func (aaosa *DataAccessApprovalOrganizationServiceAccount) Attributes() dataAccessApprovalOrganizationServiceAccountAttributes {
	return dataAccessApprovalOrganizationServiceAccountAttributes{ref: terra.ReferenceDataResource(aaosa)}
}

// DataAccessApprovalOrganizationServiceAccountArgs contains the configurations for google_access_approval_organization_service_account.
type DataAccessApprovalOrganizationServiceAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrganizationId: string, required
	OrganizationId terra.StringValue `hcl:"organization_id,attr" validate:"required"`
}
type dataAccessApprovalOrganizationServiceAccountAttributes struct {
	ref terra.Reference
}

// AccountEmail returns a reference to field account_email of google_access_approval_organization_service_account.
func (aaosa dataAccessApprovalOrganizationServiceAccountAttributes) AccountEmail() terra.StringValue {
	return terra.ReferenceAsString(aaosa.ref.Append("account_email"))
}

// Id returns a reference to field id of google_access_approval_organization_service_account.
func (aaosa dataAccessApprovalOrganizationServiceAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aaosa.ref.Append("id"))
}

// Name returns a reference to field name of google_access_approval_organization_service_account.
func (aaosa dataAccessApprovalOrganizationServiceAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aaosa.ref.Append("name"))
}

// OrganizationId returns a reference to field organization_id of google_access_approval_organization_service_account.
func (aaosa dataAccessApprovalOrganizationServiceAccountAttributes) OrganizationId() terra.StringValue {
	return terra.ReferenceAsString(aaosa.ref.Append("organization_id"))
}
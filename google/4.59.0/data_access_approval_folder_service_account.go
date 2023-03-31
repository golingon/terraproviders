// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataAccessApprovalFolderServiceAccount(name string, args DataAccessApprovalFolderServiceAccountArgs) *DataAccessApprovalFolderServiceAccount {
	return &DataAccessApprovalFolderServiceAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAccessApprovalFolderServiceAccount)(nil)

type DataAccessApprovalFolderServiceAccount struct {
	Name string
	Args DataAccessApprovalFolderServiceAccountArgs
}

func (aafsa *DataAccessApprovalFolderServiceAccount) DataSource() string {
	return "google_access_approval_folder_service_account"
}

func (aafsa *DataAccessApprovalFolderServiceAccount) LocalName() string {
	return aafsa.Name
}

func (aafsa *DataAccessApprovalFolderServiceAccount) Configuration() interface{} {
	return aafsa.Args
}

func (aafsa *DataAccessApprovalFolderServiceAccount) Attributes() dataAccessApprovalFolderServiceAccountAttributes {
	return dataAccessApprovalFolderServiceAccountAttributes{ref: terra.ReferenceDataResource(aafsa)}
}

type DataAccessApprovalFolderServiceAccountArgs struct {
	// FolderId: string, required
	FolderId terra.StringValue `hcl:"folder_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataAccessApprovalFolderServiceAccountAttributes struct {
	ref terra.Reference
}

func (aafsa dataAccessApprovalFolderServiceAccountAttributes) AccountEmail() terra.StringValue {
	return terra.ReferenceString(aafsa.ref.Append("account_email"))
}

func (aafsa dataAccessApprovalFolderServiceAccountAttributes) FolderId() terra.StringValue {
	return terra.ReferenceString(aafsa.ref.Append("folder_id"))
}

func (aafsa dataAccessApprovalFolderServiceAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceString(aafsa.ref.Append("id"))
}

func (aafsa dataAccessApprovalFolderServiceAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceString(aafsa.ref.Append("name"))
}

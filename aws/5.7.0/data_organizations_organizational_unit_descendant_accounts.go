// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataorganizationsorganizationalunitdescendantaccounts "github.com/golingon/terraproviders/aws/5.7.0/dataorganizationsorganizationalunitdescendantaccounts"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataOrganizationsOrganizationalUnitDescendantAccounts creates a new instance of [DataOrganizationsOrganizationalUnitDescendantAccounts].
func NewDataOrganizationsOrganizationalUnitDescendantAccounts(name string, args DataOrganizationsOrganizationalUnitDescendantAccountsArgs) *DataOrganizationsOrganizationalUnitDescendantAccounts {
	return &DataOrganizationsOrganizationalUnitDescendantAccounts{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOrganizationsOrganizationalUnitDescendantAccounts)(nil)

// DataOrganizationsOrganizationalUnitDescendantAccounts represents the Terraform data resource aws_organizations_organizational_unit_descendant_accounts.
type DataOrganizationsOrganizationalUnitDescendantAccounts struct {
	Name string
	Args DataOrganizationsOrganizationalUnitDescendantAccountsArgs
}

// DataSource returns the Terraform object type for [DataOrganizationsOrganizationalUnitDescendantAccounts].
func (oouda *DataOrganizationsOrganizationalUnitDescendantAccounts) DataSource() string {
	return "aws_organizations_organizational_unit_descendant_accounts"
}

// LocalName returns the local name for [DataOrganizationsOrganizationalUnitDescendantAccounts].
func (oouda *DataOrganizationsOrganizationalUnitDescendantAccounts) LocalName() string {
	return oouda.Name
}

// Configuration returns the configuration (args) for [DataOrganizationsOrganizationalUnitDescendantAccounts].
func (oouda *DataOrganizationsOrganizationalUnitDescendantAccounts) Configuration() interface{} {
	return oouda.Args
}

// Attributes returns the attributes for [DataOrganizationsOrganizationalUnitDescendantAccounts].
func (oouda *DataOrganizationsOrganizationalUnitDescendantAccounts) Attributes() dataOrganizationsOrganizationalUnitDescendantAccountsAttributes {
	return dataOrganizationsOrganizationalUnitDescendantAccountsAttributes{ref: terra.ReferenceDataResource(oouda)}
}

// DataOrganizationsOrganizationalUnitDescendantAccountsArgs contains the configurations for aws_organizations_organizational_unit_descendant_accounts.
type DataOrganizationsOrganizationalUnitDescendantAccountsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ParentId: string, required
	ParentId terra.StringValue `hcl:"parent_id,attr" validate:"required"`
	// Accounts: min=0
	Accounts []dataorganizationsorganizationalunitdescendantaccounts.Accounts `hcl:"accounts,block" validate:"min=0"`
}
type dataOrganizationsOrganizationalUnitDescendantAccountsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_organizations_organizational_unit_descendant_accounts.
func (oouda dataOrganizationsOrganizationalUnitDescendantAccountsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oouda.ref.Append("id"))
}

// ParentId returns a reference to field parent_id of aws_organizations_organizational_unit_descendant_accounts.
func (oouda dataOrganizationsOrganizationalUnitDescendantAccountsAttributes) ParentId() terra.StringValue {
	return terra.ReferenceAsString(oouda.ref.Append("parent_id"))
}

func (oouda dataOrganizationsOrganizationalUnitDescendantAccountsAttributes) Accounts() terra.ListValue[dataorganizationsorganizationalunitdescendantaccounts.AccountsAttributes] {
	return terra.ReferenceAsList[dataorganizationsorganizationalunitdescendantaccounts.AccountsAttributes](oouda.ref.Append("accounts"))
}

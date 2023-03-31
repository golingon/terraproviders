// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacosmosdbrestorabledatabaseaccounts "github.com/golingon/terraproviders/azurerm/3.49.0/datacosmosdbrestorabledatabaseaccounts"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataCosmosdbRestorableDatabaseAccounts(name string, args DataCosmosdbRestorableDatabaseAccountsArgs) *DataCosmosdbRestorableDatabaseAccounts {
	return &DataCosmosdbRestorableDatabaseAccounts{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCosmosdbRestorableDatabaseAccounts)(nil)

type DataCosmosdbRestorableDatabaseAccounts struct {
	Name string
	Args DataCosmosdbRestorableDatabaseAccountsArgs
}

func (crda *DataCosmosdbRestorableDatabaseAccounts) DataSource() string {
	return "azurerm_cosmosdb_restorable_database_accounts"
}

func (crda *DataCosmosdbRestorableDatabaseAccounts) LocalName() string {
	return crda.Name
}

func (crda *DataCosmosdbRestorableDatabaseAccounts) Configuration() interface{} {
	return crda.Args
}

func (crda *DataCosmosdbRestorableDatabaseAccounts) Attributes() dataCosmosdbRestorableDatabaseAccountsAttributes {
	return dataCosmosdbRestorableDatabaseAccountsAttributes{ref: terra.ReferenceDataResource(crda)}
}

type DataCosmosdbRestorableDatabaseAccountsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Accounts: min=0
	Accounts []datacosmosdbrestorabledatabaseaccounts.Accounts `hcl:"accounts,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datacosmosdbrestorabledatabaseaccounts.Timeouts `hcl:"timeouts,block"`
}
type dataCosmosdbRestorableDatabaseAccountsAttributes struct {
	ref terra.Reference
}

func (crda dataCosmosdbRestorableDatabaseAccountsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(crda.ref.Append("id"))
}

func (crda dataCosmosdbRestorableDatabaseAccountsAttributes) Location() terra.StringValue {
	return terra.ReferenceString(crda.ref.Append("location"))
}

func (crda dataCosmosdbRestorableDatabaseAccountsAttributes) Name() terra.StringValue {
	return terra.ReferenceString(crda.ref.Append("name"))
}

func (crda dataCosmosdbRestorableDatabaseAccountsAttributes) Accounts() terra.ListValue[datacosmosdbrestorabledatabaseaccounts.AccountsAttributes] {
	return terra.ReferenceList[datacosmosdbrestorabledatabaseaccounts.AccountsAttributes](crda.ref.Append("accounts"))
}

func (crda dataCosmosdbRestorableDatabaseAccountsAttributes) Timeouts() datacosmosdbrestorabledatabaseaccounts.TimeoutsAttributes {
	return terra.ReferenceSingle[datacosmosdbrestorabledatabaseaccounts.TimeoutsAttributes](crda.ref.Append("timeouts"))
}

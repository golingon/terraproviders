// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadatashareaccount "github.com/golingon/terraproviders/azurerm/3.67.0/datadatashareaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDataShareAccount creates a new instance of [DataDataShareAccount].
func NewDataDataShareAccount(name string, args DataDataShareAccountArgs) *DataDataShareAccount {
	return &DataDataShareAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataShareAccount)(nil)

// DataDataShareAccount represents the Terraform data resource azurerm_data_share_account.
type DataDataShareAccount struct {
	Name string
	Args DataDataShareAccountArgs
}

// DataSource returns the Terraform object type for [DataDataShareAccount].
func (dsa *DataDataShareAccount) DataSource() string {
	return "azurerm_data_share_account"
}

// LocalName returns the local name for [DataDataShareAccount].
func (dsa *DataDataShareAccount) LocalName() string {
	return dsa.Name
}

// Configuration returns the configuration (args) for [DataDataShareAccount].
func (dsa *DataDataShareAccount) Configuration() interface{} {
	return dsa.Args
}

// Attributes returns the attributes for [DataDataShareAccount].
func (dsa *DataDataShareAccount) Attributes() dataDataShareAccountAttributes {
	return dataDataShareAccountAttributes{ref: terra.ReferenceDataResource(dsa)}
}

// DataDataShareAccountArgs contains the configurations for azurerm_data_share_account.
type DataDataShareAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datadatashareaccount.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadatashareaccount.Timeouts `hcl:"timeouts,block"`
}
type dataDataShareAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_data_share_account.
func (dsa dataDataShareAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_share_account.
func (dsa dataDataShareAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_data_share_account.
func (dsa dataDataShareAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_data_share_account.
func (dsa dataDataShareAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsa.ref.Append("tags"))
}

func (dsa dataDataShareAccountAttributes) Identity() terra.ListValue[datadatashareaccount.IdentityAttributes] {
	return terra.ReferenceAsList[datadatashareaccount.IdentityAttributes](dsa.ref.Append("identity"))
}

func (dsa dataDataShareAccountAttributes) Timeouts() datadatashareaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadatashareaccount.TimeoutsAttributes](dsa.ref.Append("timeouts"))
}

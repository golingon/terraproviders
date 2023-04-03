// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalogicappintegrationaccount "github.com/golingon/terraproviders/azurerm/3.49.0/datalogicappintegrationaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLogicAppIntegrationAccount creates a new instance of [DataLogicAppIntegrationAccount].
func NewDataLogicAppIntegrationAccount(name string, args DataLogicAppIntegrationAccountArgs) *DataLogicAppIntegrationAccount {
	return &DataLogicAppIntegrationAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLogicAppIntegrationAccount)(nil)

// DataLogicAppIntegrationAccount represents the Terraform data resource azurerm_logic_app_integration_account.
type DataLogicAppIntegrationAccount struct {
	Name string
	Args DataLogicAppIntegrationAccountArgs
}

// DataSource returns the Terraform object type for [DataLogicAppIntegrationAccount].
func (laia *DataLogicAppIntegrationAccount) DataSource() string {
	return "azurerm_logic_app_integration_account"
}

// LocalName returns the local name for [DataLogicAppIntegrationAccount].
func (laia *DataLogicAppIntegrationAccount) LocalName() string {
	return laia.Name
}

// Configuration returns the configuration (args) for [DataLogicAppIntegrationAccount].
func (laia *DataLogicAppIntegrationAccount) Configuration() interface{} {
	return laia.Args
}

// Attributes returns the attributes for [DataLogicAppIntegrationAccount].
func (laia *DataLogicAppIntegrationAccount) Attributes() dataLogicAppIntegrationAccountAttributes {
	return dataLogicAppIntegrationAccountAttributes{ref: terra.ReferenceDataResource(laia)}
}

// DataLogicAppIntegrationAccountArgs contains the configurations for azurerm_logic_app_integration_account.
type DataLogicAppIntegrationAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datalogicappintegrationaccount.Timeouts `hcl:"timeouts,block"`
}
type dataLogicAppIntegrationAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_logic_app_integration_account.
func (laia dataLogicAppIntegrationAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_logic_app_integration_account.
func (laia dataLogicAppIntegrationAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_logic_app_integration_account.
func (laia dataLogicAppIntegrationAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_integration_account.
func (laia dataLogicAppIntegrationAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_logic_app_integration_account.
func (laia dataLogicAppIntegrationAccountAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_logic_app_integration_account.
func (laia dataLogicAppIntegrationAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](laia.ref.Append("tags"))
}

func (laia dataLogicAppIntegrationAccountAttributes) Timeouts() datalogicappintegrationaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalogicappintegrationaccount.TimeoutsAttributes](laia.ref.Append("timeouts"))
}

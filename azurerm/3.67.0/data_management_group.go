// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamanagementgroup "github.com/golingon/terraproviders/azurerm/3.67.0/datamanagementgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataManagementGroup creates a new instance of [DataManagementGroup].
func NewDataManagementGroup(name string, args DataManagementGroupArgs) *DataManagementGroup {
	return &DataManagementGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataManagementGroup)(nil)

// DataManagementGroup represents the Terraform data resource azurerm_management_group.
type DataManagementGroup struct {
	Name string
	Args DataManagementGroupArgs
}

// DataSource returns the Terraform object type for [DataManagementGroup].
func (mg *DataManagementGroup) DataSource() string {
	return "azurerm_management_group"
}

// LocalName returns the local name for [DataManagementGroup].
func (mg *DataManagementGroup) LocalName() string {
	return mg.Name
}

// Configuration returns the configuration (args) for [DataManagementGroup].
func (mg *DataManagementGroup) Configuration() interface{} {
	return mg.Args
}

// Attributes returns the attributes for [DataManagementGroup].
func (mg *DataManagementGroup) Attributes() dataManagementGroupAttributes {
	return dataManagementGroupAttributes{ref: terra.ReferenceDataResource(mg)}
}

// DataManagementGroupArgs contains the configurations for azurerm_management_group.
type DataManagementGroupArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Timeouts: optional
	Timeouts *datamanagementgroup.Timeouts `hcl:"timeouts,block"`
}
type dataManagementGroupAttributes struct {
	ref terra.Reference
}

// AllManagementGroupIds returns a reference to field all_management_group_ids of azurerm_management_group.
func (mg dataManagementGroupAttributes) AllManagementGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mg.ref.Append("all_management_group_ids"))
}

// AllSubscriptionIds returns a reference to field all_subscription_ids of azurerm_management_group.
func (mg dataManagementGroupAttributes) AllSubscriptionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mg.ref.Append("all_subscription_ids"))
}

// DisplayName returns a reference to field display_name of azurerm_management_group.
func (mg dataManagementGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_management_group.
func (mg dataManagementGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("id"))
}

// ManagementGroupIds returns a reference to field management_group_ids of azurerm_management_group.
func (mg dataManagementGroupAttributes) ManagementGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mg.ref.Append("management_group_ids"))
}

// Name returns a reference to field name of azurerm_management_group.
func (mg dataManagementGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("name"))
}

// ParentManagementGroupId returns a reference to field parent_management_group_id of azurerm_management_group.
func (mg dataManagementGroupAttributes) ParentManagementGroupId() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("parent_management_group_id"))
}

// SubscriptionIds returns a reference to field subscription_ids of azurerm_management_group.
func (mg dataManagementGroupAttributes) SubscriptionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mg.ref.Append("subscription_ids"))
}

func (mg dataManagementGroupAttributes) Timeouts() datamanagementgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamanagementgroup.TimeoutsAttributes](mg.ref.Append("timeouts"))
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datanetappaccount "github.com/golingon/terraproviders/azurerm/3.52.0/datanetappaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetappAccount creates a new instance of [DataNetappAccount].
func NewDataNetappAccount(name string, args DataNetappAccountArgs) *DataNetappAccount {
	return &DataNetappAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetappAccount)(nil)

// DataNetappAccount represents the Terraform data resource azurerm_netapp_account.
type DataNetappAccount struct {
	Name string
	Args DataNetappAccountArgs
}

// DataSource returns the Terraform object type for [DataNetappAccount].
func (na *DataNetappAccount) DataSource() string {
	return "azurerm_netapp_account"
}

// LocalName returns the local name for [DataNetappAccount].
func (na *DataNetappAccount) LocalName() string {
	return na.Name
}

// Configuration returns the configuration (args) for [DataNetappAccount].
func (na *DataNetappAccount) Configuration() interface{} {
	return na.Args
}

// Attributes returns the attributes for [DataNetappAccount].
func (na *DataNetappAccount) Attributes() dataNetappAccountAttributes {
	return dataNetappAccountAttributes{ref: terra.ReferenceDataResource(na)}
}

// DataNetappAccountArgs contains the configurations for azurerm_netapp_account.
type DataNetappAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datanetappaccount.Timeouts `hcl:"timeouts,block"`
}
type dataNetappAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_netapp_account.
func (na dataNetappAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_netapp_account.
func (na dataNetappAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_netapp_account.
func (na dataNetappAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_netapp_account.
func (na dataNetappAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("resource_group_name"))
}

func (na dataNetappAccountAttributes) Timeouts() datanetappaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetappaccount.TimeoutsAttributes](na.ref.Append("timeouts"))
}

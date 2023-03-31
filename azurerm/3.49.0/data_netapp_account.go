// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datanetappaccount "github.com/golingon/terraproviders/azurerm/3.49.0/datanetappaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataNetappAccount(name string, args DataNetappAccountArgs) *DataNetappAccount {
	return &DataNetappAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetappAccount)(nil)

type DataNetappAccount struct {
	Name string
	Args DataNetappAccountArgs
}

func (na *DataNetappAccount) DataSource() string {
	return "azurerm_netapp_account"
}

func (na *DataNetappAccount) LocalName() string {
	return na.Name
}

func (na *DataNetappAccount) Configuration() interface{} {
	return na.Args
}

func (na *DataNetappAccount) Attributes() dataNetappAccountAttributes {
	return dataNetappAccountAttributes{ref: terra.ReferenceDataResource(na)}
}

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

func (na dataNetappAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceString(na.ref.Append("id"))
}

func (na dataNetappAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceString(na.ref.Append("location"))
}

func (na dataNetappAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceString(na.ref.Append("name"))
}

func (na dataNetappAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(na.ref.Append("resource_group_name"))
}

func (na dataNetappAccountAttributes) Timeouts() datanetappaccount.TimeoutsAttributes {
	return terra.ReferenceSingle[datanetappaccount.TimeoutsAttributes](na.ref.Append("timeouts"))
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataappconfigurationkeys "github.com/golingon/terraproviders/azurerm/3.49.0/dataappconfigurationkeys"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataAppConfigurationKeys(name string, args DataAppConfigurationKeysArgs) *DataAppConfigurationKeys {
	return &DataAppConfigurationKeys{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppConfigurationKeys)(nil)

type DataAppConfigurationKeys struct {
	Name string
	Args DataAppConfigurationKeysArgs
}

func (ack *DataAppConfigurationKeys) DataSource() string {
	return "azurerm_app_configuration_keys"
}

func (ack *DataAppConfigurationKeys) LocalName() string {
	return ack.Name
}

func (ack *DataAppConfigurationKeys) Configuration() interface{} {
	return ack.Args
}

func (ack *DataAppConfigurationKeys) Attributes() dataAppConfigurationKeysAttributes {
	return dataAppConfigurationKeysAttributes{ref: terra.ReferenceDataResource(ack)}
}

type DataAppConfigurationKeysArgs struct {
	// ConfigurationStoreId: string, required
	ConfigurationStoreId terra.StringValue `hcl:"configuration_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// Items: min=0
	Items []dataappconfigurationkeys.Items `hcl:"items,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataappconfigurationkeys.Timeouts `hcl:"timeouts,block"`
}
type dataAppConfigurationKeysAttributes struct {
	ref terra.Reference
}

func (ack dataAppConfigurationKeysAttributes) ConfigurationStoreId() terra.StringValue {
	return terra.ReferenceString(ack.ref.Append("configuration_store_id"))
}

func (ack dataAppConfigurationKeysAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ack.ref.Append("id"))
}

func (ack dataAppConfigurationKeysAttributes) Key() terra.StringValue {
	return terra.ReferenceString(ack.ref.Append("key"))
}

func (ack dataAppConfigurationKeysAttributes) Label() terra.StringValue {
	return terra.ReferenceString(ack.ref.Append("label"))
}

func (ack dataAppConfigurationKeysAttributes) Items() terra.ListValue[dataappconfigurationkeys.ItemsAttributes] {
	return terra.ReferenceList[dataappconfigurationkeys.ItemsAttributes](ack.ref.Append("items"))
}

func (ack dataAppConfigurationKeysAttributes) Timeouts() dataappconfigurationkeys.TimeoutsAttributes {
	return terra.ReferenceSingle[dataappconfigurationkeys.TimeoutsAttributes](ack.ref.Append("timeouts"))
}

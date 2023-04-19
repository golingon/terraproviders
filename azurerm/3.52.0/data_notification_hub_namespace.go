// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datanotificationhubnamespace "github.com/golingon/terraproviders/azurerm/3.52.0/datanotificationhubnamespace"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNotificationHubNamespace creates a new instance of [DataNotificationHubNamespace].
func NewDataNotificationHubNamespace(name string, args DataNotificationHubNamespaceArgs) *DataNotificationHubNamespace {
	return &DataNotificationHubNamespace{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNotificationHubNamespace)(nil)

// DataNotificationHubNamespace represents the Terraform data resource azurerm_notification_hub_namespace.
type DataNotificationHubNamespace struct {
	Name string
	Args DataNotificationHubNamespaceArgs
}

// DataSource returns the Terraform object type for [DataNotificationHubNamespace].
func (nhn *DataNotificationHubNamespace) DataSource() string {
	return "azurerm_notification_hub_namespace"
}

// LocalName returns the local name for [DataNotificationHubNamespace].
func (nhn *DataNotificationHubNamespace) LocalName() string {
	return nhn.Name
}

// Configuration returns the configuration (args) for [DataNotificationHubNamespace].
func (nhn *DataNotificationHubNamespace) Configuration() interface{} {
	return nhn.Args
}

// Attributes returns the attributes for [DataNotificationHubNamespace].
func (nhn *DataNotificationHubNamespace) Attributes() dataNotificationHubNamespaceAttributes {
	return dataNotificationHubNamespaceAttributes{ref: terra.ReferenceDataResource(nhn)}
}

// DataNotificationHubNamespaceArgs contains the configurations for azurerm_notification_hub_namespace.
type DataNotificationHubNamespaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: min=0
	Sku []datanotificationhubnamespace.Sku `hcl:"sku,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanotificationhubnamespace.Timeouts `hcl:"timeouts,block"`
}
type dataNotificationHubNamespaceAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_notification_hub_namespace.
func (nhn dataNotificationHubNamespaceAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(nhn.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_notification_hub_namespace.
func (nhn dataNotificationHubNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nhn.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_notification_hub_namespace.
func (nhn dataNotificationHubNamespaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nhn.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_notification_hub_namespace.
func (nhn dataNotificationHubNamespaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nhn.ref.Append("name"))
}

// NamespaceType returns a reference to field namespace_type of azurerm_notification_hub_namespace.
func (nhn dataNotificationHubNamespaceAttributes) NamespaceType() terra.StringValue {
	return terra.ReferenceAsString(nhn.ref.Append("namespace_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_notification_hub_namespace.
func (nhn dataNotificationHubNamespaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nhn.ref.Append("resource_group_name"))
}

// ServicebusEndpoint returns a reference to field servicebus_endpoint of azurerm_notification_hub_namespace.
func (nhn dataNotificationHubNamespaceAttributes) ServicebusEndpoint() terra.StringValue {
	return terra.ReferenceAsString(nhn.ref.Append("servicebus_endpoint"))
}

// Tags returns a reference to field tags of azurerm_notification_hub_namespace.
func (nhn dataNotificationHubNamespaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nhn.ref.Append("tags"))
}

func (nhn dataNotificationHubNamespaceAttributes) Sku() terra.ListValue[datanotificationhubnamespace.SkuAttributes] {
	return terra.ReferenceAsList[datanotificationhubnamespace.SkuAttributes](nhn.ref.Append("sku"))
}

func (nhn dataNotificationHubNamespaceAttributes) Timeouts() datanotificationhubnamespace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanotificationhubnamespace.TimeoutsAttributes](nhn.ref.Append("timeouts"))
}

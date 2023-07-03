// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datafunctionapphostkeys "github.com/golingon/terraproviders/azurerm/3.63.0/datafunctionapphostkeys"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataFunctionAppHostKeys creates a new instance of [DataFunctionAppHostKeys].
func NewDataFunctionAppHostKeys(name string, args DataFunctionAppHostKeysArgs) *DataFunctionAppHostKeys {
	return &DataFunctionAppHostKeys{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataFunctionAppHostKeys)(nil)

// DataFunctionAppHostKeys represents the Terraform data resource azurerm_function_app_host_keys.
type DataFunctionAppHostKeys struct {
	Name string
	Args DataFunctionAppHostKeysArgs
}

// DataSource returns the Terraform object type for [DataFunctionAppHostKeys].
func (fahk *DataFunctionAppHostKeys) DataSource() string {
	return "azurerm_function_app_host_keys"
}

// LocalName returns the local name for [DataFunctionAppHostKeys].
func (fahk *DataFunctionAppHostKeys) LocalName() string {
	return fahk.Name
}

// Configuration returns the configuration (args) for [DataFunctionAppHostKeys].
func (fahk *DataFunctionAppHostKeys) Configuration() interface{} {
	return fahk.Args
}

// Attributes returns the attributes for [DataFunctionAppHostKeys].
func (fahk *DataFunctionAppHostKeys) Attributes() dataFunctionAppHostKeysAttributes {
	return dataFunctionAppHostKeysAttributes{ref: terra.ReferenceDataResource(fahk)}
}

// DataFunctionAppHostKeysArgs contains the configurations for azurerm_function_app_host_keys.
type DataFunctionAppHostKeysArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datafunctionapphostkeys.Timeouts `hcl:"timeouts,block"`
}
type dataFunctionAppHostKeysAttributes struct {
	ref terra.Reference
}

// BlobsExtensionKey returns a reference to field blobs_extension_key of azurerm_function_app_host_keys.
func (fahk dataFunctionAppHostKeysAttributes) BlobsExtensionKey() terra.StringValue {
	return terra.ReferenceAsString(fahk.ref.Append("blobs_extension_key"))
}

// DefaultFunctionKey returns a reference to field default_function_key of azurerm_function_app_host_keys.
func (fahk dataFunctionAppHostKeysAttributes) DefaultFunctionKey() terra.StringValue {
	return terra.ReferenceAsString(fahk.ref.Append("default_function_key"))
}

// DurabletaskExtensionKey returns a reference to field durabletask_extension_key of azurerm_function_app_host_keys.
func (fahk dataFunctionAppHostKeysAttributes) DurabletaskExtensionKey() terra.StringValue {
	return terra.ReferenceAsString(fahk.ref.Append("durabletask_extension_key"))
}

// EventGridExtensionConfigKey returns a reference to field event_grid_extension_config_key of azurerm_function_app_host_keys.
func (fahk dataFunctionAppHostKeysAttributes) EventGridExtensionConfigKey() terra.StringValue {
	return terra.ReferenceAsString(fahk.ref.Append("event_grid_extension_config_key"))
}

// Id returns a reference to field id of azurerm_function_app_host_keys.
func (fahk dataFunctionAppHostKeysAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fahk.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_function_app_host_keys.
func (fahk dataFunctionAppHostKeysAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fahk.ref.Append("name"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_function_app_host_keys.
func (fahk dataFunctionAppHostKeysAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(fahk.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_function_app_host_keys.
func (fahk dataFunctionAppHostKeysAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(fahk.ref.Append("resource_group_name"))
}

// SignalrExtensionKey returns a reference to field signalr_extension_key of azurerm_function_app_host_keys.
func (fahk dataFunctionAppHostKeysAttributes) SignalrExtensionKey() terra.StringValue {
	return terra.ReferenceAsString(fahk.ref.Append("signalr_extension_key"))
}

// WebpubsubExtensionKey returns a reference to field webpubsub_extension_key of azurerm_function_app_host_keys.
func (fahk dataFunctionAppHostKeysAttributes) WebpubsubExtensionKey() terra.StringValue {
	return terra.ReferenceAsString(fahk.ref.Append("webpubsub_extension_key"))
}

func (fahk dataFunctionAppHostKeysAttributes) Timeouts() datafunctionapphostkeys.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafunctionapphostkeys.TimeoutsAttributes](fahk.ref.Append("timeouts"))
}

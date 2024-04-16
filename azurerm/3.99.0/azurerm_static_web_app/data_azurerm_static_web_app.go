// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_static_web_app

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_static_web_app.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aswa *DataSource) DataSource() string {
	return "azurerm_static_web_app"
}

// LocalName returns the local name for [DataSource].
func (aswa *DataSource) LocalName() string {
	return aswa.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aswa *DataSource) Configuration() interface{} {
	return aswa.Args
}

// Attributes returns the attributes for [DataSource].
func (aswa *DataSource) Attributes() dataAzurermStaticWebAppAttributes {
	return dataAzurermStaticWebAppAttributes{ref: terra.ReferenceDataSource(aswa)}
}

// DataArgs contains the configurations for azurerm_static_web_app.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermStaticWebAppAttributes struct {
	ref terra.Reference
}

// ApiKey returns a reference to field api_key of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) ApiKey() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("api_key"))
}

// AppSettings returns a reference to field app_settings of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aswa.ref.Append("app_settings"))
}

// ConfigurationFileChangesEnabled returns a reference to field configuration_file_changes_enabled of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) ConfigurationFileChangesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aswa.ref.Append("configuration_file_changes_enabled"))
}

// DefaultHostName returns a reference to field default_host_name of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) DefaultHostName() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("default_host_name"))
}

// Id returns a reference to field id of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("name"))
}

// PreviewEnvironmentsEnabled returns a reference to field preview_environments_enabled of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) PreviewEnvironmentsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aswa.ref.Append("preview_environments_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("resource_group_name"))
}

// SkuSize returns a reference to field sku_size of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) SkuSize() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("sku_size"))
}

// SkuTier returns a reference to field sku_tier of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) SkuTier() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("sku_tier"))
}

// Tags returns a reference to field tags of azurerm_static_web_app.
func (aswa dataAzurermStaticWebAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aswa.ref.Append("tags"))
}

func (aswa dataAzurermStaticWebAppAttributes) BasicAuth() terra.ListValue[DataBasicAuthAttributes] {
	return terra.ReferenceAsList[DataBasicAuthAttributes](aswa.ref.Append("basic_auth"))
}

func (aswa dataAzurermStaticWebAppAttributes) Identity() terra.ListValue[DataIdentityAttributes] {
	return terra.ReferenceAsList[DataIdentityAttributes](aswa.ref.Append("identity"))
}

func (aswa dataAzurermStaticWebAppAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](aswa.ref.Append("timeouts"))
}

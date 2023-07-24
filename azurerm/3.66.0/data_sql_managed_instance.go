// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasqlmanagedinstance "github.com/golingon/terraproviders/azurerm/3.66.0/datasqlmanagedinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSqlManagedInstance creates a new instance of [DataSqlManagedInstance].
func NewDataSqlManagedInstance(name string, args DataSqlManagedInstanceArgs) *DataSqlManagedInstance {
	return &DataSqlManagedInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSqlManagedInstance)(nil)

// DataSqlManagedInstance represents the Terraform data resource azurerm_sql_managed_instance.
type DataSqlManagedInstance struct {
	Name string
	Args DataSqlManagedInstanceArgs
}

// DataSource returns the Terraform object type for [DataSqlManagedInstance].
func (smi *DataSqlManagedInstance) DataSource() string {
	return "azurerm_sql_managed_instance"
}

// LocalName returns the local name for [DataSqlManagedInstance].
func (smi *DataSqlManagedInstance) LocalName() string {
	return smi.Name
}

// Configuration returns the configuration (args) for [DataSqlManagedInstance].
func (smi *DataSqlManagedInstance) Configuration() interface{} {
	return smi.Args
}

// Attributes returns the attributes for [DataSqlManagedInstance].
func (smi *DataSqlManagedInstance) Attributes() dataSqlManagedInstanceAttributes {
	return dataSqlManagedInstanceAttributes{ref: terra.ReferenceDataResource(smi)}
}

// DataSqlManagedInstanceArgs contains the configurations for azurerm_sql_managed_instance.
type DataSqlManagedInstanceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: min=0
	Identity []datasqlmanagedinstance.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasqlmanagedinstance.Timeouts `hcl:"timeouts,block"`
}
type dataSqlManagedInstanceAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("administrator_login"))
}

// Collation returns a reference to field collation of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("collation"))
}

// DnsZonePartnerId returns a reference to field dns_zone_partner_id of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) DnsZonePartnerId() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("dns_zone_partner_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("location"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("name"))
}

// ProxyOverride returns a reference to field proxy_override of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) ProxyOverride() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("proxy_override"))
}

// PublicDataEndpointEnabled returns a reference to field public_data_endpoint_enabled of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) PublicDataEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(smi.ref.Append("public_data_endpoint_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("sku_name"))
}

// StorageAccountType returns a reference to field storage_account_type of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("storage_account_type"))
}

// StorageSizeInGb returns a reference to field storage_size_in_gb of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) StorageSizeInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(smi.ref.Append("storage_size_in_gb"))
}

// SubnetId returns a reference to field subnet_id of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](smi.ref.Append("tags"))
}

// TimezoneId returns a reference to field timezone_id of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) TimezoneId() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("timezone_id"))
}

// Vcores returns a reference to field vcores of azurerm_sql_managed_instance.
func (smi dataSqlManagedInstanceAttributes) Vcores() terra.NumberValue {
	return terra.ReferenceAsNumber(smi.ref.Append("vcores"))
}

func (smi dataSqlManagedInstanceAttributes) Identity() terra.ListValue[datasqlmanagedinstance.IdentityAttributes] {
	return terra.ReferenceAsList[datasqlmanagedinstance.IdentityAttributes](smi.ref.Append("identity"))
}

func (smi dataSqlManagedInstanceAttributes) Timeouts() datasqlmanagedinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasqlmanagedinstance.TimeoutsAttributes](smi.ref.Append("timeouts"))
}

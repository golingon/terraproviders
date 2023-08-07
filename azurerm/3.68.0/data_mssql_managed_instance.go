// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamssqlmanagedinstance "github.com/golingon/terraproviders/azurerm/3.68.0/datamssqlmanagedinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMssqlManagedInstance creates a new instance of [DataMssqlManagedInstance].
func NewDataMssqlManagedInstance(name string, args DataMssqlManagedInstanceArgs) *DataMssqlManagedInstance {
	return &DataMssqlManagedInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMssqlManagedInstance)(nil)

// DataMssqlManagedInstance represents the Terraform data resource azurerm_mssql_managed_instance.
type DataMssqlManagedInstance struct {
	Name string
	Args DataMssqlManagedInstanceArgs
}

// DataSource returns the Terraform object type for [DataMssqlManagedInstance].
func (mmi *DataMssqlManagedInstance) DataSource() string {
	return "azurerm_mssql_managed_instance"
}

// LocalName returns the local name for [DataMssqlManagedInstance].
func (mmi *DataMssqlManagedInstance) LocalName() string {
	return mmi.Name
}

// Configuration returns the configuration (args) for [DataMssqlManagedInstance].
func (mmi *DataMssqlManagedInstance) Configuration() interface{} {
	return mmi.Args
}

// Attributes returns the attributes for [DataMssqlManagedInstance].
func (mmi *DataMssqlManagedInstance) Attributes() dataMssqlManagedInstanceAttributes {
	return dataMssqlManagedInstanceAttributes{ref: terra.ReferenceDataResource(mmi)}
}

// DataMssqlManagedInstanceArgs contains the configurations for azurerm_mssql_managed_instance.
type DataMssqlManagedInstanceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datamssqlmanagedinstance.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamssqlmanagedinstance.Timeouts `hcl:"timeouts,block"`
}
type dataMssqlManagedInstanceAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("administrator_login"))
}

// Collation returns a reference to field collation of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("collation"))
}

// CustomerManagedKeyId returns a reference to field customer_managed_key_id of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) CustomerManagedKeyId() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("customer_managed_key_id"))
}

// DnsZonePartnerId returns a reference to field dns_zone_partner_id of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) DnsZonePartnerId() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("dns_zone_partner_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("location"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("name"))
}

// ProxyOverride returns a reference to field proxy_override of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) ProxyOverride() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("proxy_override"))
}

// PublicDataEndpointEnabled returns a reference to field public_data_endpoint_enabled of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) PublicDataEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mmi.ref.Append("public_data_endpoint_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("sku_name"))
}

// StorageAccountType returns a reference to field storage_account_type of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("storage_account_type"))
}

// StorageSizeInGb returns a reference to field storage_size_in_gb of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) StorageSizeInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(mmi.ref.Append("storage_size_in_gb"))
}

// SubnetId returns a reference to field subnet_id of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mmi.ref.Append("tags"))
}

// TimezoneId returns a reference to field timezone_id of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) TimezoneId() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("timezone_id"))
}

// Vcores returns a reference to field vcores of azurerm_mssql_managed_instance.
func (mmi dataMssqlManagedInstanceAttributes) Vcores() terra.NumberValue {
	return terra.ReferenceAsNumber(mmi.ref.Append("vcores"))
}

func (mmi dataMssqlManagedInstanceAttributes) Identity() terra.ListValue[datamssqlmanagedinstance.IdentityAttributes] {
	return terra.ReferenceAsList[datamssqlmanagedinstance.IdentityAttributes](mmi.ref.Append("identity"))
}

func (mmi dataMssqlManagedInstanceAttributes) Timeouts() datamssqlmanagedinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamssqlmanagedinstance.TimeoutsAttributes](mmi.ref.Append("timeouts"))
}

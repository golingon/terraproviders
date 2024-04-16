// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mssql_managed_instance

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_mssql_managed_instance.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermMssqlManagedInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ammi *Resource) Type() string {
	return "azurerm_mssql_managed_instance"
}

// LocalName returns the local name for [Resource].
func (ammi *Resource) LocalName() string {
	return ammi.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ammi *Resource) Configuration() interface{} {
	return ammi.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ammi *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ammi)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ammi *Resource) Dependencies() terra.Dependencies {
	return ammi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ammi *Resource) LifecycleManagement() *terra.Lifecycle {
	return ammi.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ammi *Resource) Attributes() azurermMssqlManagedInstanceAttributes {
	return azurermMssqlManagedInstanceAttributes{ref: terra.ReferenceResource(ammi)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ammi *Resource) ImportState(state io.Reader) error {
	ammi.state = &azurermMssqlManagedInstanceState{}
	if err := json.NewDecoder(state).Decode(ammi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ammi.Type(), ammi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ammi *Resource) State() (*azurermMssqlManagedInstanceState, bool) {
	return ammi.state, ammi.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ammi *Resource) StateMust() *azurermMssqlManagedInstanceState {
	if ammi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ammi.Type(), ammi.LocalName()))
	}
	return ammi.state
}

// Args contains the configurations for azurerm_mssql_managed_instance.
type Args struct {
	// AdministratorLogin: string, required
	AdministratorLogin terra.StringValue `hcl:"administrator_login,attr" validate:"required"`
	// AdministratorLoginPassword: string, required
	AdministratorLoginPassword terra.StringValue `hcl:"administrator_login_password,attr" validate:"required"`
	// Collation: string, optional
	Collation terra.StringValue `hcl:"collation,attr"`
	// DnsZonePartnerId: string, optional
	DnsZonePartnerId terra.StringValue `hcl:"dns_zone_partner_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseType: string, required
	LicenseType terra.StringValue `hcl:"license_type,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaintenanceConfigurationName: string, optional
	MaintenanceConfigurationName terra.StringValue `hcl:"maintenance_configuration_name,attr"`
	// MinimumTlsVersion: string, optional
	MinimumTlsVersion terra.StringValue `hcl:"minimum_tls_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProxyOverride: string, optional
	ProxyOverride terra.StringValue `hcl:"proxy_override,attr"`
	// PublicDataEndpointEnabled: bool, optional
	PublicDataEndpointEnabled terra.BoolValue `hcl:"public_data_endpoint_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// StorageAccountType: string, optional
	StorageAccountType terra.StringValue `hcl:"storage_account_type,attr"`
	// StorageSizeInGb: number, required
	StorageSizeInGb terra.NumberValue `hcl:"storage_size_in_gb,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TimezoneId: string, optional
	TimezoneId terra.StringValue `hcl:"timezone_id,attr"`
	// Vcores: number, required
	Vcores terra.NumberValue `hcl:"vcores,attr" validate:"required"`
	// ZoneRedundantEnabled: bool, optional
	ZoneRedundantEnabled terra.BoolValue `hcl:"zone_redundant_enabled,attr"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermMssqlManagedInstanceAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("administrator_login"))
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("administrator_login_password"))
}

// Collation returns a reference to field collation of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("collation"))
}

// DnsZone returns a reference to field dns_zone of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) DnsZone() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("dns_zone"))
}

// DnsZonePartnerId returns a reference to field dns_zone_partner_id of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) DnsZonePartnerId() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("dns_zone_partner_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("location"))
}

// MaintenanceConfigurationName returns a reference to field maintenance_configuration_name of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) MaintenanceConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("maintenance_configuration_name"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("name"))
}

// ProxyOverride returns a reference to field proxy_override of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) ProxyOverride() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("proxy_override"))
}

// PublicDataEndpointEnabled returns a reference to field public_data_endpoint_enabled of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) PublicDataEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ammi.ref.Append("public_data_endpoint_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("sku_name"))
}

// StorageAccountType returns a reference to field storage_account_type of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("storage_account_type"))
}

// StorageSizeInGb returns a reference to field storage_size_in_gb of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) StorageSizeInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ammi.ref.Append("storage_size_in_gb"))
}

// SubnetId returns a reference to field subnet_id of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ammi.ref.Append("tags"))
}

// TimezoneId returns a reference to field timezone_id of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) TimezoneId() terra.StringValue {
	return terra.ReferenceAsString(ammi.ref.Append("timezone_id"))
}

// Vcores returns a reference to field vcores of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) Vcores() terra.NumberValue {
	return terra.ReferenceAsNumber(ammi.ref.Append("vcores"))
}

// ZoneRedundantEnabled returns a reference to field zone_redundant_enabled of azurerm_mssql_managed_instance.
func (ammi azurermMssqlManagedInstanceAttributes) ZoneRedundantEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ammi.ref.Append("zone_redundant_enabled"))
}

func (ammi azurermMssqlManagedInstanceAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](ammi.ref.Append("identity"))
}

func (ammi azurermMssqlManagedInstanceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ammi.ref.Append("timeouts"))
}

type azurermMssqlManagedInstanceState struct {
	AdministratorLogin           string            `json:"administrator_login"`
	AdministratorLoginPassword   string            `json:"administrator_login_password"`
	Collation                    string            `json:"collation"`
	DnsZone                      string            `json:"dns_zone"`
	DnsZonePartnerId             string            `json:"dns_zone_partner_id"`
	Fqdn                         string            `json:"fqdn"`
	Id                           string            `json:"id"`
	LicenseType                  string            `json:"license_type"`
	Location                     string            `json:"location"`
	MaintenanceConfigurationName string            `json:"maintenance_configuration_name"`
	MinimumTlsVersion            string            `json:"minimum_tls_version"`
	Name                         string            `json:"name"`
	ProxyOverride                string            `json:"proxy_override"`
	PublicDataEndpointEnabled    bool              `json:"public_data_endpoint_enabled"`
	ResourceGroupName            string            `json:"resource_group_name"`
	SkuName                      string            `json:"sku_name"`
	StorageAccountType           string            `json:"storage_account_type"`
	StorageSizeInGb              float64           `json:"storage_size_in_gb"`
	SubnetId                     string            `json:"subnet_id"`
	Tags                         map[string]string `json:"tags"`
	TimezoneId                   string            `json:"timezone_id"`
	Vcores                       float64           `json:"vcores"`
	ZoneRedundantEnabled         bool              `json:"zone_redundant_enabled"`
	Identity                     []IdentityState   `json:"identity"`
	Timeouts                     *TimeoutsState    `json:"timeouts"`
}

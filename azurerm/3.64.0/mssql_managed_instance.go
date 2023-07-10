// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlmanagedinstance "github.com/golingon/terraproviders/azurerm/3.64.0/mssqlmanagedinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlManagedInstance creates a new instance of [MssqlManagedInstance].
func NewMssqlManagedInstance(name string, args MssqlManagedInstanceArgs) *MssqlManagedInstance {
	return &MssqlManagedInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlManagedInstance)(nil)

// MssqlManagedInstance represents the Terraform resource azurerm_mssql_managed_instance.
type MssqlManagedInstance struct {
	Name      string
	Args      MssqlManagedInstanceArgs
	state     *mssqlManagedInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlManagedInstance].
func (mmi *MssqlManagedInstance) Type() string {
	return "azurerm_mssql_managed_instance"
}

// LocalName returns the local name for [MssqlManagedInstance].
func (mmi *MssqlManagedInstance) LocalName() string {
	return mmi.Name
}

// Configuration returns the configuration (args) for [MssqlManagedInstance].
func (mmi *MssqlManagedInstance) Configuration() interface{} {
	return mmi.Args
}

// DependOn is used for other resources to depend on [MssqlManagedInstance].
func (mmi *MssqlManagedInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(mmi)
}

// Dependencies returns the list of resources [MssqlManagedInstance] depends_on.
func (mmi *MssqlManagedInstance) Dependencies() terra.Dependencies {
	return mmi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlManagedInstance].
func (mmi *MssqlManagedInstance) LifecycleManagement() *terra.Lifecycle {
	return mmi.Lifecycle
}

// Attributes returns the attributes for [MssqlManagedInstance].
func (mmi *MssqlManagedInstance) Attributes() mssqlManagedInstanceAttributes {
	return mssqlManagedInstanceAttributes{ref: terra.ReferenceResource(mmi)}
}

// ImportState imports the given attribute values into [MssqlManagedInstance]'s state.
func (mmi *MssqlManagedInstance) ImportState(av io.Reader) error {
	mmi.state = &mssqlManagedInstanceState{}
	if err := json.NewDecoder(av).Decode(mmi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mmi.Type(), mmi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlManagedInstance] has state.
func (mmi *MssqlManagedInstance) State() (*mssqlManagedInstanceState, bool) {
	return mmi.state, mmi.state != nil
}

// StateMust returns the state for [MssqlManagedInstance]. Panics if the state is nil.
func (mmi *MssqlManagedInstance) StateMust() *mssqlManagedInstanceState {
	if mmi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mmi.Type(), mmi.LocalName()))
	}
	return mmi.state
}

// MssqlManagedInstanceArgs contains the configurations for azurerm_mssql_managed_instance.
type MssqlManagedInstanceArgs struct {
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
	// Identity: optional
	Identity *mssqlmanagedinstance.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *mssqlmanagedinstance.Timeouts `hcl:"timeouts,block"`
}
type mssqlManagedInstanceAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("administrator_login"))
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("administrator_login_password"))
}

// Collation returns a reference to field collation of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("collation"))
}

// DnsZonePartnerId returns a reference to field dns_zone_partner_id of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) DnsZonePartnerId() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("dns_zone_partner_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("location"))
}

// MaintenanceConfigurationName returns a reference to field maintenance_configuration_name of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) MaintenanceConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("maintenance_configuration_name"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("name"))
}

// ProxyOverride returns a reference to field proxy_override of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) ProxyOverride() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("proxy_override"))
}

// PublicDataEndpointEnabled returns a reference to field public_data_endpoint_enabled of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) PublicDataEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mmi.ref.Append("public_data_endpoint_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("sku_name"))
}

// StorageAccountType returns a reference to field storage_account_type of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("storage_account_type"))
}

// StorageSizeInGb returns a reference to field storage_size_in_gb of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) StorageSizeInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(mmi.ref.Append("storage_size_in_gb"))
}

// SubnetId returns a reference to field subnet_id of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mmi.ref.Append("tags"))
}

// TimezoneId returns a reference to field timezone_id of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) TimezoneId() terra.StringValue {
	return terra.ReferenceAsString(mmi.ref.Append("timezone_id"))
}

// Vcores returns a reference to field vcores of azurerm_mssql_managed_instance.
func (mmi mssqlManagedInstanceAttributes) Vcores() terra.NumberValue {
	return terra.ReferenceAsNumber(mmi.ref.Append("vcores"))
}

func (mmi mssqlManagedInstanceAttributes) Identity() terra.ListValue[mssqlmanagedinstance.IdentityAttributes] {
	return terra.ReferenceAsList[mssqlmanagedinstance.IdentityAttributes](mmi.ref.Append("identity"))
}

func (mmi mssqlManagedInstanceAttributes) Timeouts() mssqlmanagedinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlmanagedinstance.TimeoutsAttributes](mmi.ref.Append("timeouts"))
}

type mssqlManagedInstanceState struct {
	AdministratorLogin           string                               `json:"administrator_login"`
	AdministratorLoginPassword   string                               `json:"administrator_login_password"`
	Collation                    string                               `json:"collation"`
	DnsZonePartnerId             string                               `json:"dns_zone_partner_id"`
	Fqdn                         string                               `json:"fqdn"`
	Id                           string                               `json:"id"`
	LicenseType                  string                               `json:"license_type"`
	Location                     string                               `json:"location"`
	MaintenanceConfigurationName string                               `json:"maintenance_configuration_name"`
	MinimumTlsVersion            string                               `json:"minimum_tls_version"`
	Name                         string                               `json:"name"`
	ProxyOverride                string                               `json:"proxy_override"`
	PublicDataEndpointEnabled    bool                                 `json:"public_data_endpoint_enabled"`
	ResourceGroupName            string                               `json:"resource_group_name"`
	SkuName                      string                               `json:"sku_name"`
	StorageAccountType           string                               `json:"storage_account_type"`
	StorageSizeInGb              float64                              `json:"storage_size_in_gb"`
	SubnetId                     string                               `json:"subnet_id"`
	Tags                         map[string]string                    `json:"tags"`
	TimezoneId                   string                               `json:"timezone_id"`
	Vcores                       float64                              `json:"vcores"`
	Identity                     []mssqlmanagedinstance.IdentityState `json:"identity"`
	Timeouts                     *mssqlmanagedinstance.TimeoutsState  `json:"timeouts"`
}

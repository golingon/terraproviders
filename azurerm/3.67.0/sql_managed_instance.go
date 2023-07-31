// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sqlmanagedinstance "github.com/golingon/terraproviders/azurerm/3.67.0/sqlmanagedinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlManagedInstance creates a new instance of [SqlManagedInstance].
func NewSqlManagedInstance(name string, args SqlManagedInstanceArgs) *SqlManagedInstance {
	return &SqlManagedInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlManagedInstance)(nil)

// SqlManagedInstance represents the Terraform resource azurerm_sql_managed_instance.
type SqlManagedInstance struct {
	Name      string
	Args      SqlManagedInstanceArgs
	state     *sqlManagedInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlManagedInstance].
func (smi *SqlManagedInstance) Type() string {
	return "azurerm_sql_managed_instance"
}

// LocalName returns the local name for [SqlManagedInstance].
func (smi *SqlManagedInstance) LocalName() string {
	return smi.Name
}

// Configuration returns the configuration (args) for [SqlManagedInstance].
func (smi *SqlManagedInstance) Configuration() interface{} {
	return smi.Args
}

// DependOn is used for other resources to depend on [SqlManagedInstance].
func (smi *SqlManagedInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(smi)
}

// Dependencies returns the list of resources [SqlManagedInstance] depends_on.
func (smi *SqlManagedInstance) Dependencies() terra.Dependencies {
	return smi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlManagedInstance].
func (smi *SqlManagedInstance) LifecycleManagement() *terra.Lifecycle {
	return smi.Lifecycle
}

// Attributes returns the attributes for [SqlManagedInstance].
func (smi *SqlManagedInstance) Attributes() sqlManagedInstanceAttributes {
	return sqlManagedInstanceAttributes{ref: terra.ReferenceResource(smi)}
}

// ImportState imports the given attribute values into [SqlManagedInstance]'s state.
func (smi *SqlManagedInstance) ImportState(av io.Reader) error {
	smi.state = &sqlManagedInstanceState{}
	if err := json.NewDecoder(av).Decode(smi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smi.Type(), smi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlManagedInstance] has state.
func (smi *SqlManagedInstance) State() (*sqlManagedInstanceState, bool) {
	return smi.state, smi.state != nil
}

// StateMust returns the state for [SqlManagedInstance]. Panics if the state is nil.
func (smi *SqlManagedInstance) StateMust() *sqlManagedInstanceState {
	if smi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smi.Type(), smi.LocalName()))
	}
	return smi.state
}

// SqlManagedInstanceArgs contains the configurations for azurerm_sql_managed_instance.
type SqlManagedInstanceArgs struct {
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
	Identity *sqlmanagedinstance.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *sqlmanagedinstance.Timeouts `hcl:"timeouts,block"`
}
type sqlManagedInstanceAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("administrator_login"))
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("administrator_login_password"))
}

// Collation returns a reference to field collation of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("collation"))
}

// DnsZonePartnerId returns a reference to field dns_zone_partner_id of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) DnsZonePartnerId() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("dns_zone_partner_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("location"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("name"))
}

// ProxyOverride returns a reference to field proxy_override of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) ProxyOverride() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("proxy_override"))
}

// PublicDataEndpointEnabled returns a reference to field public_data_endpoint_enabled of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) PublicDataEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(smi.ref.Append("public_data_endpoint_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("sku_name"))
}

// StorageAccountType returns a reference to field storage_account_type of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("storage_account_type"))
}

// StorageSizeInGb returns a reference to field storage_size_in_gb of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) StorageSizeInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(smi.ref.Append("storage_size_in_gb"))
}

// SubnetId returns a reference to field subnet_id of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](smi.ref.Append("tags"))
}

// TimezoneId returns a reference to field timezone_id of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) TimezoneId() terra.StringValue {
	return terra.ReferenceAsString(smi.ref.Append("timezone_id"))
}

// Vcores returns a reference to field vcores of azurerm_sql_managed_instance.
func (smi sqlManagedInstanceAttributes) Vcores() terra.NumberValue {
	return terra.ReferenceAsNumber(smi.ref.Append("vcores"))
}

func (smi sqlManagedInstanceAttributes) Identity() terra.ListValue[sqlmanagedinstance.IdentityAttributes] {
	return terra.ReferenceAsList[sqlmanagedinstance.IdentityAttributes](smi.ref.Append("identity"))
}

func (smi sqlManagedInstanceAttributes) Timeouts() sqlmanagedinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlmanagedinstance.TimeoutsAttributes](smi.ref.Append("timeouts"))
}

type sqlManagedInstanceState struct {
	AdministratorLogin         string                             `json:"administrator_login"`
	AdministratorLoginPassword string                             `json:"administrator_login_password"`
	Collation                  string                             `json:"collation"`
	DnsZonePartnerId           string                             `json:"dns_zone_partner_id"`
	Fqdn                       string                             `json:"fqdn"`
	Id                         string                             `json:"id"`
	LicenseType                string                             `json:"license_type"`
	Location                   string                             `json:"location"`
	MinimumTlsVersion          string                             `json:"minimum_tls_version"`
	Name                       string                             `json:"name"`
	ProxyOverride              string                             `json:"proxy_override"`
	PublicDataEndpointEnabled  bool                               `json:"public_data_endpoint_enabled"`
	ResourceGroupName          string                             `json:"resource_group_name"`
	SkuName                    string                             `json:"sku_name"`
	StorageAccountType         string                             `json:"storage_account_type"`
	StorageSizeInGb            float64                            `json:"storage_size_in_gb"`
	SubnetId                   string                             `json:"subnet_id"`
	Tags                       map[string]string                  `json:"tags"`
	TimezoneId                 string                             `json:"timezone_id"`
	Vcores                     float64                            `json:"vcores"`
	Identity                   []sqlmanagedinstance.IdentityState `json:"identity"`
	Timeouts                   *sqlmanagedinstance.TimeoutsState  `json:"timeouts"`
}

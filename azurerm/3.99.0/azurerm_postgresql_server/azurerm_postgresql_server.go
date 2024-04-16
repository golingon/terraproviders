// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_postgresql_server

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

// Resource represents the Terraform resource azurerm_postgresql_server.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermPostgresqlServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aps *Resource) Type() string {
	return "azurerm_postgresql_server"
}

// LocalName returns the local name for [Resource].
func (aps *Resource) LocalName() string {
	return aps.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aps *Resource) Configuration() interface{} {
	return aps.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aps *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aps)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aps *Resource) Dependencies() terra.Dependencies {
	return aps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aps *Resource) LifecycleManagement() *terra.Lifecycle {
	return aps.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aps *Resource) Attributes() azurermPostgresqlServerAttributes {
	return azurermPostgresqlServerAttributes{ref: terra.ReferenceResource(aps)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aps *Resource) ImportState(state io.Reader) error {
	aps.state = &azurermPostgresqlServerState{}
	if err := json.NewDecoder(state).Decode(aps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aps.Type(), aps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aps *Resource) State() (*azurermPostgresqlServerState, bool) {
	return aps.state, aps.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aps *Resource) StateMust() *azurermPostgresqlServerState {
	if aps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aps.Type(), aps.LocalName()))
	}
	return aps.state
}

// Args contains the configurations for azurerm_postgresql_server.
type Args struct {
	// AdministratorLogin: string, optional
	AdministratorLogin terra.StringValue `hcl:"administrator_login,attr"`
	// AdministratorLoginPassword: string, optional
	AdministratorLoginPassword terra.StringValue `hcl:"administrator_login_password,attr"`
	// AutoGrowEnabled: bool, optional
	AutoGrowEnabled terra.BoolValue `hcl:"auto_grow_enabled,attr"`
	// BackupRetentionDays: number, optional
	BackupRetentionDays terra.NumberValue `hcl:"backup_retention_days,attr"`
	// CreateMode: string, optional
	CreateMode terra.StringValue `hcl:"create_mode,attr"`
	// CreationSourceServerId: string, optional
	CreationSourceServerId terra.StringValue `hcl:"creation_source_server_id,attr"`
	// GeoRedundantBackupEnabled: bool, optional
	GeoRedundantBackupEnabled terra.BoolValue `hcl:"geo_redundant_backup_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InfrastructureEncryptionEnabled: bool, optional
	InfrastructureEncryptionEnabled terra.BoolValue `hcl:"infrastructure_encryption_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RestorePointInTime: string, optional
	RestorePointInTime terra.StringValue `hcl:"restore_point_in_time,attr"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// SslEnforcementEnabled: bool, required
	SslEnforcementEnabled terra.BoolValue `hcl:"ssl_enforcement_enabled,attr" validate:"required"`
	// SslMinimalTlsVersionEnforced: string, optional
	SslMinimalTlsVersionEnforced terra.StringValue `hcl:"ssl_minimal_tls_version_enforced,attr"`
	// StorageMb: number, optional
	StorageMb terra.NumberValue `hcl:"storage_mb,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// ThreatDetectionPolicy: optional
	ThreatDetectionPolicy *ThreatDetectionPolicy `hcl:"threat_detection_policy,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermPostgresqlServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("administrator_login"))
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("administrator_login_password"))
}

// AutoGrowEnabled returns a reference to field auto_grow_enabled of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) AutoGrowEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aps.ref.Append("auto_grow_enabled"))
}

// BackupRetentionDays returns a reference to field backup_retention_days of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) BackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(aps.ref.Append("backup_retention_days"))
}

// CreateMode returns a reference to field create_mode of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) CreateMode() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("create_mode"))
}

// CreationSourceServerId returns a reference to field creation_source_server_id of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) CreationSourceServerId() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("creation_source_server_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("fqdn"))
}

// GeoRedundantBackupEnabled returns a reference to field geo_redundant_backup_enabled of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) GeoRedundantBackupEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aps.ref.Append("geo_redundant_backup_enabled"))
}

// Id returns a reference to field id of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("id"))
}

// InfrastructureEncryptionEnabled returns a reference to field infrastructure_encryption_enabled of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) InfrastructureEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aps.ref.Append("infrastructure_encryption_enabled"))
}

// Location returns a reference to field location of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aps.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("resource_group_name"))
}

// RestorePointInTime returns a reference to field restore_point_in_time of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) RestorePointInTime() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("restore_point_in_time"))
}

// SkuName returns a reference to field sku_name of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("sku_name"))
}

// SslEnforcementEnabled returns a reference to field ssl_enforcement_enabled of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) SslEnforcementEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aps.ref.Append("ssl_enforcement_enabled"))
}

// SslMinimalTlsVersionEnforced returns a reference to field ssl_minimal_tls_version_enforced of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) SslMinimalTlsVersionEnforced() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("ssl_minimal_tls_version_enforced"))
}

// StorageMb returns a reference to field storage_mb of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) StorageMb() terra.NumberValue {
	return terra.ReferenceAsNumber(aps.ref.Append("storage_mb"))
}

// Tags returns a reference to field tags of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aps.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_postgresql_server.
func (aps azurermPostgresqlServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aps.ref.Append("version"))
}

func (aps azurermPostgresqlServerAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](aps.ref.Append("identity"))
}

func (aps azurermPostgresqlServerAttributes) ThreatDetectionPolicy() terra.ListValue[ThreatDetectionPolicyAttributes] {
	return terra.ReferenceAsList[ThreatDetectionPolicyAttributes](aps.ref.Append("threat_detection_policy"))
}

func (aps azurermPostgresqlServerAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aps.ref.Append("timeouts"))
}

type azurermPostgresqlServerState struct {
	AdministratorLogin              string                       `json:"administrator_login"`
	AdministratorLoginPassword      string                       `json:"administrator_login_password"`
	AutoGrowEnabled                 bool                         `json:"auto_grow_enabled"`
	BackupRetentionDays             float64                      `json:"backup_retention_days"`
	CreateMode                      string                       `json:"create_mode"`
	CreationSourceServerId          string                       `json:"creation_source_server_id"`
	Fqdn                            string                       `json:"fqdn"`
	GeoRedundantBackupEnabled       bool                         `json:"geo_redundant_backup_enabled"`
	Id                              string                       `json:"id"`
	InfrastructureEncryptionEnabled bool                         `json:"infrastructure_encryption_enabled"`
	Location                        string                       `json:"location"`
	Name                            string                       `json:"name"`
	PublicNetworkAccessEnabled      bool                         `json:"public_network_access_enabled"`
	ResourceGroupName               string                       `json:"resource_group_name"`
	RestorePointInTime              string                       `json:"restore_point_in_time"`
	SkuName                         string                       `json:"sku_name"`
	SslEnforcementEnabled           bool                         `json:"ssl_enforcement_enabled"`
	SslMinimalTlsVersionEnforced    string                       `json:"ssl_minimal_tls_version_enforced"`
	StorageMb                       float64                      `json:"storage_mb"`
	Tags                            map[string]string            `json:"tags"`
	Version                         string                       `json:"version"`
	Identity                        []IdentityState              `json:"identity"`
	ThreatDetectionPolicy           []ThreatDetectionPolicyState `json:"threat_detection_policy"`
	Timeouts                        *TimeoutsState               `json:"timeouts"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	postgresqlserver "github.com/golingon/terraproviders/azurerm/3.68.0/postgresqlserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPostgresqlServer creates a new instance of [PostgresqlServer].
func NewPostgresqlServer(name string, args PostgresqlServerArgs) *PostgresqlServer {
	return &PostgresqlServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PostgresqlServer)(nil)

// PostgresqlServer represents the Terraform resource azurerm_postgresql_server.
type PostgresqlServer struct {
	Name      string
	Args      PostgresqlServerArgs
	state     *postgresqlServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PostgresqlServer].
func (ps *PostgresqlServer) Type() string {
	return "azurerm_postgresql_server"
}

// LocalName returns the local name for [PostgresqlServer].
func (ps *PostgresqlServer) LocalName() string {
	return ps.Name
}

// Configuration returns the configuration (args) for [PostgresqlServer].
func (ps *PostgresqlServer) Configuration() interface{} {
	return ps.Args
}

// DependOn is used for other resources to depend on [PostgresqlServer].
func (ps *PostgresqlServer) DependOn() terra.Reference {
	return terra.ReferenceResource(ps)
}

// Dependencies returns the list of resources [PostgresqlServer] depends_on.
func (ps *PostgresqlServer) Dependencies() terra.Dependencies {
	return ps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PostgresqlServer].
func (ps *PostgresqlServer) LifecycleManagement() *terra.Lifecycle {
	return ps.Lifecycle
}

// Attributes returns the attributes for [PostgresqlServer].
func (ps *PostgresqlServer) Attributes() postgresqlServerAttributes {
	return postgresqlServerAttributes{ref: terra.ReferenceResource(ps)}
}

// ImportState imports the given attribute values into [PostgresqlServer]'s state.
func (ps *PostgresqlServer) ImportState(av io.Reader) error {
	ps.state = &postgresqlServerState{}
	if err := json.NewDecoder(av).Decode(ps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ps.Type(), ps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PostgresqlServer] has state.
func (ps *PostgresqlServer) State() (*postgresqlServerState, bool) {
	return ps.state, ps.state != nil
}

// StateMust returns the state for [PostgresqlServer]. Panics if the state is nil.
func (ps *PostgresqlServer) StateMust() *postgresqlServerState {
	if ps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ps.Type(), ps.LocalName()))
	}
	return ps.state
}

// PostgresqlServerArgs contains the configurations for azurerm_postgresql_server.
type PostgresqlServerArgs struct {
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
	Identity *postgresqlserver.Identity `hcl:"identity,block"`
	// ThreatDetectionPolicy: optional
	ThreatDetectionPolicy *postgresqlserver.ThreatDetectionPolicy `hcl:"threat_detection_policy,block"`
	// Timeouts: optional
	Timeouts *postgresqlserver.Timeouts `hcl:"timeouts,block"`
}
type postgresqlServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("administrator_login"))
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("administrator_login_password"))
}

// AutoGrowEnabled returns a reference to field auto_grow_enabled of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) AutoGrowEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("auto_grow_enabled"))
}

// BackupRetentionDays returns a reference to field backup_retention_days of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) BackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ps.ref.Append("backup_retention_days"))
}

// CreateMode returns a reference to field create_mode of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) CreateMode() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("create_mode"))
}

// CreationSourceServerId returns a reference to field creation_source_server_id of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) CreationSourceServerId() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("creation_source_server_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("fqdn"))
}

// GeoRedundantBackupEnabled returns a reference to field geo_redundant_backup_enabled of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) GeoRedundantBackupEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("geo_redundant_backup_enabled"))
}

// Id returns a reference to field id of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("id"))
}

// InfrastructureEncryptionEnabled returns a reference to field infrastructure_encryption_enabled of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) InfrastructureEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("infrastructure_encryption_enabled"))
}

// Location returns a reference to field location of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("resource_group_name"))
}

// RestorePointInTime returns a reference to field restore_point_in_time of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) RestorePointInTime() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("restore_point_in_time"))
}

// SkuName returns a reference to field sku_name of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("sku_name"))
}

// SslEnforcementEnabled returns a reference to field ssl_enforcement_enabled of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) SslEnforcementEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("ssl_enforcement_enabled"))
}

// SslMinimalTlsVersionEnforced returns a reference to field ssl_minimal_tls_version_enforced of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) SslMinimalTlsVersionEnforced() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("ssl_minimal_tls_version_enforced"))
}

// StorageMb returns a reference to field storage_mb of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) StorageMb() terra.NumberValue {
	return terra.ReferenceAsNumber(ps.ref.Append("storage_mb"))
}

// Tags returns a reference to field tags of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ps.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_postgresql_server.
func (ps postgresqlServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("version"))
}

func (ps postgresqlServerAttributes) Identity() terra.ListValue[postgresqlserver.IdentityAttributes] {
	return terra.ReferenceAsList[postgresqlserver.IdentityAttributes](ps.ref.Append("identity"))
}

func (ps postgresqlServerAttributes) ThreatDetectionPolicy() terra.ListValue[postgresqlserver.ThreatDetectionPolicyAttributes] {
	return terra.ReferenceAsList[postgresqlserver.ThreatDetectionPolicyAttributes](ps.ref.Append("threat_detection_policy"))
}

func (ps postgresqlServerAttributes) Timeouts() postgresqlserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[postgresqlserver.TimeoutsAttributes](ps.ref.Append("timeouts"))
}

type postgresqlServerState struct {
	AdministratorLogin              string                                        `json:"administrator_login"`
	AdministratorLoginPassword      string                                        `json:"administrator_login_password"`
	AutoGrowEnabled                 bool                                          `json:"auto_grow_enabled"`
	BackupRetentionDays             float64                                       `json:"backup_retention_days"`
	CreateMode                      string                                        `json:"create_mode"`
	CreationSourceServerId          string                                        `json:"creation_source_server_id"`
	Fqdn                            string                                        `json:"fqdn"`
	GeoRedundantBackupEnabled       bool                                          `json:"geo_redundant_backup_enabled"`
	Id                              string                                        `json:"id"`
	InfrastructureEncryptionEnabled bool                                          `json:"infrastructure_encryption_enabled"`
	Location                        string                                        `json:"location"`
	Name                            string                                        `json:"name"`
	PublicNetworkAccessEnabled      bool                                          `json:"public_network_access_enabled"`
	ResourceGroupName               string                                        `json:"resource_group_name"`
	RestorePointInTime              string                                        `json:"restore_point_in_time"`
	SkuName                         string                                        `json:"sku_name"`
	SslEnforcementEnabled           bool                                          `json:"ssl_enforcement_enabled"`
	SslMinimalTlsVersionEnforced    string                                        `json:"ssl_minimal_tls_version_enforced"`
	StorageMb                       float64                                       `json:"storage_mb"`
	Tags                            map[string]string                             `json:"tags"`
	Version                         string                                        `json:"version"`
	Identity                        []postgresqlserver.IdentityState              `json:"identity"`
	ThreatDetectionPolicy           []postgresqlserver.ThreatDetectionPolicyState `json:"threat_detection_policy"`
	Timeouts                        *postgresqlserver.TimeoutsState               `json:"timeouts"`
}
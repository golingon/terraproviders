// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqlserver "github.com/golingon/terraproviders/azurerm/3.67.0/mysqlserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMysqlServer creates a new instance of [MysqlServer].
func NewMysqlServer(name string, args MysqlServerArgs) *MysqlServer {
	return &MysqlServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlServer)(nil)

// MysqlServer represents the Terraform resource azurerm_mysql_server.
type MysqlServer struct {
	Name      string
	Args      MysqlServerArgs
	state     *mysqlServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlServer].
func (ms *MysqlServer) Type() string {
	return "azurerm_mysql_server"
}

// LocalName returns the local name for [MysqlServer].
func (ms *MysqlServer) LocalName() string {
	return ms.Name
}

// Configuration returns the configuration (args) for [MysqlServer].
func (ms *MysqlServer) Configuration() interface{} {
	return ms.Args
}

// DependOn is used for other resources to depend on [MysqlServer].
func (ms *MysqlServer) DependOn() terra.Reference {
	return terra.ReferenceResource(ms)
}

// Dependencies returns the list of resources [MysqlServer] depends_on.
func (ms *MysqlServer) Dependencies() terra.Dependencies {
	return ms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlServer].
func (ms *MysqlServer) LifecycleManagement() *terra.Lifecycle {
	return ms.Lifecycle
}

// Attributes returns the attributes for [MysqlServer].
func (ms *MysqlServer) Attributes() mysqlServerAttributes {
	return mysqlServerAttributes{ref: terra.ReferenceResource(ms)}
}

// ImportState imports the given attribute values into [MysqlServer]'s state.
func (ms *MysqlServer) ImportState(av io.Reader) error {
	ms.state = &mysqlServerState{}
	if err := json.NewDecoder(av).Decode(ms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ms.Type(), ms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlServer] has state.
func (ms *MysqlServer) State() (*mysqlServerState, bool) {
	return ms.state, ms.state != nil
}

// StateMust returns the state for [MysqlServer]. Panics if the state is nil.
func (ms *MysqlServer) StateMust() *mysqlServerState {
	if ms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ms.Type(), ms.LocalName()))
	}
	return ms.state
}

// MysqlServerArgs contains the configurations for azurerm_mysql_server.
type MysqlServerArgs struct {
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
	Identity *mysqlserver.Identity `hcl:"identity,block"`
	// ThreatDetectionPolicy: optional
	ThreatDetectionPolicy *mysqlserver.ThreatDetectionPolicy `hcl:"threat_detection_policy,block"`
	// Timeouts: optional
	Timeouts *mysqlserver.Timeouts `hcl:"timeouts,block"`
}
type mysqlServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_mysql_server.
func (ms mysqlServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("administrator_login"))
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_mysql_server.
func (ms mysqlServerAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("administrator_login_password"))
}

// AutoGrowEnabled returns a reference to field auto_grow_enabled of azurerm_mysql_server.
func (ms mysqlServerAttributes) AutoGrowEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("auto_grow_enabled"))
}

// BackupRetentionDays returns a reference to field backup_retention_days of azurerm_mysql_server.
func (ms mysqlServerAttributes) BackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("backup_retention_days"))
}

// CreateMode returns a reference to field create_mode of azurerm_mysql_server.
func (ms mysqlServerAttributes) CreateMode() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("create_mode"))
}

// CreationSourceServerId returns a reference to field creation_source_server_id of azurerm_mysql_server.
func (ms mysqlServerAttributes) CreationSourceServerId() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("creation_source_server_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_mysql_server.
func (ms mysqlServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("fqdn"))
}

// GeoRedundantBackupEnabled returns a reference to field geo_redundant_backup_enabled of azurerm_mysql_server.
func (ms mysqlServerAttributes) GeoRedundantBackupEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("geo_redundant_backup_enabled"))
}

// Id returns a reference to field id of azurerm_mysql_server.
func (ms mysqlServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("id"))
}

// InfrastructureEncryptionEnabled returns a reference to field infrastructure_encryption_enabled of azurerm_mysql_server.
func (ms mysqlServerAttributes) InfrastructureEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("infrastructure_encryption_enabled"))
}

// Location returns a reference to field location of azurerm_mysql_server.
func (ms mysqlServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_mysql_server.
func (ms mysqlServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_mysql_server.
func (ms mysqlServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_server.
func (ms mysqlServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("resource_group_name"))
}

// RestorePointInTime returns a reference to field restore_point_in_time of azurerm_mysql_server.
func (ms mysqlServerAttributes) RestorePointInTime() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("restore_point_in_time"))
}

// SkuName returns a reference to field sku_name of azurerm_mysql_server.
func (ms mysqlServerAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("sku_name"))
}

// SslEnforcementEnabled returns a reference to field ssl_enforcement_enabled of azurerm_mysql_server.
func (ms mysqlServerAttributes) SslEnforcementEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("ssl_enforcement_enabled"))
}

// SslMinimalTlsVersionEnforced returns a reference to field ssl_minimal_tls_version_enforced of azurerm_mysql_server.
func (ms mysqlServerAttributes) SslMinimalTlsVersionEnforced() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("ssl_minimal_tls_version_enforced"))
}

// StorageMb returns a reference to field storage_mb of azurerm_mysql_server.
func (ms mysqlServerAttributes) StorageMb() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("storage_mb"))
}

// Tags returns a reference to field tags of azurerm_mysql_server.
func (ms mysqlServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ms.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_mysql_server.
func (ms mysqlServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("version"))
}

func (ms mysqlServerAttributes) Identity() terra.ListValue[mysqlserver.IdentityAttributes] {
	return terra.ReferenceAsList[mysqlserver.IdentityAttributes](ms.ref.Append("identity"))
}

func (ms mysqlServerAttributes) ThreatDetectionPolicy() terra.ListValue[mysqlserver.ThreatDetectionPolicyAttributes] {
	return terra.ReferenceAsList[mysqlserver.ThreatDetectionPolicyAttributes](ms.ref.Append("threat_detection_policy"))
}

func (ms mysqlServerAttributes) Timeouts() mysqlserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqlserver.TimeoutsAttributes](ms.ref.Append("timeouts"))
}

type mysqlServerState struct {
	AdministratorLogin              string                                   `json:"administrator_login"`
	AdministratorLoginPassword      string                                   `json:"administrator_login_password"`
	AutoGrowEnabled                 bool                                     `json:"auto_grow_enabled"`
	BackupRetentionDays             float64                                  `json:"backup_retention_days"`
	CreateMode                      string                                   `json:"create_mode"`
	CreationSourceServerId          string                                   `json:"creation_source_server_id"`
	Fqdn                            string                                   `json:"fqdn"`
	GeoRedundantBackupEnabled       bool                                     `json:"geo_redundant_backup_enabled"`
	Id                              string                                   `json:"id"`
	InfrastructureEncryptionEnabled bool                                     `json:"infrastructure_encryption_enabled"`
	Location                        string                                   `json:"location"`
	Name                            string                                   `json:"name"`
	PublicNetworkAccessEnabled      bool                                     `json:"public_network_access_enabled"`
	ResourceGroupName               string                                   `json:"resource_group_name"`
	RestorePointInTime              string                                   `json:"restore_point_in_time"`
	SkuName                         string                                   `json:"sku_name"`
	SslEnforcementEnabled           bool                                     `json:"ssl_enforcement_enabled"`
	SslMinimalTlsVersionEnforced    string                                   `json:"ssl_minimal_tls_version_enforced"`
	StorageMb                       float64                                  `json:"storage_mb"`
	Tags                            map[string]string                        `json:"tags"`
	Version                         string                                   `json:"version"`
	Identity                        []mysqlserver.IdentityState              `json:"identity"`
	ThreatDetectionPolicy           []mysqlserver.ThreatDetectionPolicyState `json:"threat_detection_policy"`
	Timeouts                        *mysqlserver.TimeoutsState               `json:"timeouts"`
}

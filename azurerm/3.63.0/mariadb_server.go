// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mariadbserver "github.com/golingon/terraproviders/azurerm/3.63.0/mariadbserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMariadbServer creates a new instance of [MariadbServer].
func NewMariadbServer(name string, args MariadbServerArgs) *MariadbServer {
	return &MariadbServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MariadbServer)(nil)

// MariadbServer represents the Terraform resource azurerm_mariadb_server.
type MariadbServer struct {
	Name      string
	Args      MariadbServerArgs
	state     *mariadbServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MariadbServer].
func (ms *MariadbServer) Type() string {
	return "azurerm_mariadb_server"
}

// LocalName returns the local name for [MariadbServer].
func (ms *MariadbServer) LocalName() string {
	return ms.Name
}

// Configuration returns the configuration (args) for [MariadbServer].
func (ms *MariadbServer) Configuration() interface{} {
	return ms.Args
}

// DependOn is used for other resources to depend on [MariadbServer].
func (ms *MariadbServer) DependOn() terra.Reference {
	return terra.ReferenceResource(ms)
}

// Dependencies returns the list of resources [MariadbServer] depends_on.
func (ms *MariadbServer) Dependencies() terra.Dependencies {
	return ms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MariadbServer].
func (ms *MariadbServer) LifecycleManagement() *terra.Lifecycle {
	return ms.Lifecycle
}

// Attributes returns the attributes for [MariadbServer].
func (ms *MariadbServer) Attributes() mariadbServerAttributes {
	return mariadbServerAttributes{ref: terra.ReferenceResource(ms)}
}

// ImportState imports the given attribute values into [MariadbServer]'s state.
func (ms *MariadbServer) ImportState(av io.Reader) error {
	ms.state = &mariadbServerState{}
	if err := json.NewDecoder(av).Decode(ms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ms.Type(), ms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MariadbServer] has state.
func (ms *MariadbServer) State() (*mariadbServerState, bool) {
	return ms.state, ms.state != nil
}

// StateMust returns the state for [MariadbServer]. Panics if the state is nil.
func (ms *MariadbServer) StateMust() *mariadbServerState {
	if ms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ms.Type(), ms.LocalName()))
	}
	return ms.state
}

// MariadbServerArgs contains the configurations for azurerm_mariadb_server.
type MariadbServerArgs struct {
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
	// Timeouts: optional
	Timeouts *mariadbserver.Timeouts `hcl:"timeouts,block"`
}
type mariadbServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_mariadb_server.
func (ms mariadbServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("administrator_login"))
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_mariadb_server.
func (ms mariadbServerAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("administrator_login_password"))
}

// AutoGrowEnabled returns a reference to field auto_grow_enabled of azurerm_mariadb_server.
func (ms mariadbServerAttributes) AutoGrowEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("auto_grow_enabled"))
}

// BackupRetentionDays returns a reference to field backup_retention_days of azurerm_mariadb_server.
func (ms mariadbServerAttributes) BackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("backup_retention_days"))
}

// CreateMode returns a reference to field create_mode of azurerm_mariadb_server.
func (ms mariadbServerAttributes) CreateMode() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("create_mode"))
}

// CreationSourceServerId returns a reference to field creation_source_server_id of azurerm_mariadb_server.
func (ms mariadbServerAttributes) CreationSourceServerId() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("creation_source_server_id"))
}

// Fqdn returns a reference to field fqdn of azurerm_mariadb_server.
func (ms mariadbServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("fqdn"))
}

// GeoRedundantBackupEnabled returns a reference to field geo_redundant_backup_enabled of azurerm_mariadb_server.
func (ms mariadbServerAttributes) GeoRedundantBackupEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("geo_redundant_backup_enabled"))
}

// Id returns a reference to field id of azurerm_mariadb_server.
func (ms mariadbServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mariadb_server.
func (ms mariadbServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_mariadb_server.
func (ms mariadbServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_mariadb_server.
func (ms mariadbServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mariadb_server.
func (ms mariadbServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("resource_group_name"))
}

// RestorePointInTime returns a reference to field restore_point_in_time of azurerm_mariadb_server.
func (ms mariadbServerAttributes) RestorePointInTime() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("restore_point_in_time"))
}

// SkuName returns a reference to field sku_name of azurerm_mariadb_server.
func (ms mariadbServerAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("sku_name"))
}

// SslEnforcementEnabled returns a reference to field ssl_enforcement_enabled of azurerm_mariadb_server.
func (ms mariadbServerAttributes) SslEnforcementEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("ssl_enforcement_enabled"))
}

// SslMinimalTlsVersionEnforced returns a reference to field ssl_minimal_tls_version_enforced of azurerm_mariadb_server.
func (ms mariadbServerAttributes) SslMinimalTlsVersionEnforced() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("ssl_minimal_tls_version_enforced"))
}

// StorageMb returns a reference to field storage_mb of azurerm_mariadb_server.
func (ms mariadbServerAttributes) StorageMb() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("storage_mb"))
}

// Tags returns a reference to field tags of azurerm_mariadb_server.
func (ms mariadbServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ms.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_mariadb_server.
func (ms mariadbServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("version"))
}

func (ms mariadbServerAttributes) Timeouts() mariadbserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mariadbserver.TimeoutsAttributes](ms.ref.Append("timeouts"))
}

type mariadbServerState struct {
	AdministratorLogin           string                       `json:"administrator_login"`
	AdministratorLoginPassword   string                       `json:"administrator_login_password"`
	AutoGrowEnabled              bool                         `json:"auto_grow_enabled"`
	BackupRetentionDays          float64                      `json:"backup_retention_days"`
	CreateMode                   string                       `json:"create_mode"`
	CreationSourceServerId       string                       `json:"creation_source_server_id"`
	Fqdn                         string                       `json:"fqdn"`
	GeoRedundantBackupEnabled    bool                         `json:"geo_redundant_backup_enabled"`
	Id                           string                       `json:"id"`
	Location                     string                       `json:"location"`
	Name                         string                       `json:"name"`
	PublicNetworkAccessEnabled   bool                         `json:"public_network_access_enabled"`
	ResourceGroupName            string                       `json:"resource_group_name"`
	RestorePointInTime           string                       `json:"restore_point_in_time"`
	SkuName                      string                       `json:"sku_name"`
	SslEnforcementEnabled        bool                         `json:"ssl_enforcement_enabled"`
	SslMinimalTlsVersionEnforced string                       `json:"ssl_minimal_tls_version_enforced"`
	StorageMb                    float64                      `json:"storage_mb"`
	Tags                         map[string]string            `json:"tags"`
	Version                      string                       `json:"version"`
	Timeouts                     *mariadbserver.TimeoutsState `json:"timeouts"`
}

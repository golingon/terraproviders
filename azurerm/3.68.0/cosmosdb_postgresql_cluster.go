// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbpostgresqlcluster "github.com/golingon/terraproviders/azurerm/3.68.0/cosmosdbpostgresqlcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbPostgresqlCluster creates a new instance of [CosmosdbPostgresqlCluster].
func NewCosmosdbPostgresqlCluster(name string, args CosmosdbPostgresqlClusterArgs) *CosmosdbPostgresqlCluster {
	return &CosmosdbPostgresqlCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbPostgresqlCluster)(nil)

// CosmosdbPostgresqlCluster represents the Terraform resource azurerm_cosmosdb_postgresql_cluster.
type CosmosdbPostgresqlCluster struct {
	Name      string
	Args      CosmosdbPostgresqlClusterArgs
	state     *cosmosdbPostgresqlClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbPostgresqlCluster].
func (cpc *CosmosdbPostgresqlCluster) Type() string {
	return "azurerm_cosmosdb_postgresql_cluster"
}

// LocalName returns the local name for [CosmosdbPostgresqlCluster].
func (cpc *CosmosdbPostgresqlCluster) LocalName() string {
	return cpc.Name
}

// Configuration returns the configuration (args) for [CosmosdbPostgresqlCluster].
func (cpc *CosmosdbPostgresqlCluster) Configuration() interface{} {
	return cpc.Args
}

// DependOn is used for other resources to depend on [CosmosdbPostgresqlCluster].
func (cpc *CosmosdbPostgresqlCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(cpc)
}

// Dependencies returns the list of resources [CosmosdbPostgresqlCluster] depends_on.
func (cpc *CosmosdbPostgresqlCluster) Dependencies() terra.Dependencies {
	return cpc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbPostgresqlCluster].
func (cpc *CosmosdbPostgresqlCluster) LifecycleManagement() *terra.Lifecycle {
	return cpc.Lifecycle
}

// Attributes returns the attributes for [CosmosdbPostgresqlCluster].
func (cpc *CosmosdbPostgresqlCluster) Attributes() cosmosdbPostgresqlClusterAttributes {
	return cosmosdbPostgresqlClusterAttributes{ref: terra.ReferenceResource(cpc)}
}

// ImportState imports the given attribute values into [CosmosdbPostgresqlCluster]'s state.
func (cpc *CosmosdbPostgresqlCluster) ImportState(av io.Reader) error {
	cpc.state = &cosmosdbPostgresqlClusterState{}
	if err := json.NewDecoder(av).Decode(cpc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpc.Type(), cpc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbPostgresqlCluster] has state.
func (cpc *CosmosdbPostgresqlCluster) State() (*cosmosdbPostgresqlClusterState, bool) {
	return cpc.state, cpc.state != nil
}

// StateMust returns the state for [CosmosdbPostgresqlCluster]. Panics if the state is nil.
func (cpc *CosmosdbPostgresqlCluster) StateMust() *cosmosdbPostgresqlClusterState {
	if cpc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpc.Type(), cpc.LocalName()))
	}
	return cpc.state
}

// CosmosdbPostgresqlClusterArgs contains the configurations for azurerm_cosmosdb_postgresql_cluster.
type CosmosdbPostgresqlClusterArgs struct {
	// AdministratorLoginPassword: string, required
	AdministratorLoginPassword terra.StringValue `hcl:"administrator_login_password,attr" validate:"required"`
	// CitusVersion: string, optional
	CitusVersion terra.StringValue `hcl:"citus_version,attr"`
	// CoordinatorPublicIpAccessEnabled: bool, optional
	CoordinatorPublicIpAccessEnabled terra.BoolValue `hcl:"coordinator_public_ip_access_enabled,attr"`
	// CoordinatorServerEdition: string, optional
	CoordinatorServerEdition terra.StringValue `hcl:"coordinator_server_edition,attr"`
	// CoordinatorStorageQuotaInMb: number, required
	CoordinatorStorageQuotaInMb terra.NumberValue `hcl:"coordinator_storage_quota_in_mb,attr" validate:"required"`
	// CoordinatorVcoreCount: number, required
	CoordinatorVcoreCount terra.NumberValue `hcl:"coordinator_vcore_count,attr" validate:"required"`
	// HaEnabled: bool, optional
	HaEnabled terra.BoolValue `hcl:"ha_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NodeCount: number, required
	NodeCount terra.NumberValue `hcl:"node_count,attr" validate:"required"`
	// NodePublicIpAccessEnabled: bool, optional
	NodePublicIpAccessEnabled terra.BoolValue `hcl:"node_public_ip_access_enabled,attr"`
	// NodeServerEdition: string, optional
	NodeServerEdition terra.StringValue `hcl:"node_server_edition,attr"`
	// NodeStorageQuotaInMb: number, optional
	NodeStorageQuotaInMb terra.NumberValue `hcl:"node_storage_quota_in_mb,attr"`
	// NodeVcores: number, optional
	NodeVcores terra.NumberValue `hcl:"node_vcores,attr"`
	// PointInTimeInUtc: string, optional
	PointInTimeInUtc terra.StringValue `hcl:"point_in_time_in_utc,attr"`
	// PreferredPrimaryZone: string, optional
	PreferredPrimaryZone terra.StringValue `hcl:"preferred_primary_zone,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ShardsOnCoordinatorEnabled: bool, optional
	ShardsOnCoordinatorEnabled terra.BoolValue `hcl:"shards_on_coordinator_enabled,attr"`
	// SourceLocation: string, optional
	SourceLocation terra.StringValue `hcl:"source_location,attr"`
	// SourceResourceId: string, optional
	SourceResourceId terra.StringValue `hcl:"source_resource_id,attr"`
	// SqlVersion: string, optional
	SqlVersion terra.StringValue `hcl:"sql_version,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// MaintenanceWindow: optional
	MaintenanceWindow *cosmosdbpostgresqlcluster.MaintenanceWindow `hcl:"maintenance_window,block"`
	// Timeouts: optional
	Timeouts *cosmosdbpostgresqlcluster.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbPostgresqlClusterAttributes struct {
	ref terra.Reference
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("administrator_login_password"))
}

// CitusVersion returns a reference to field citus_version of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) CitusVersion() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("citus_version"))
}

// CoordinatorPublicIpAccessEnabled returns a reference to field coordinator_public_ip_access_enabled of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) CoordinatorPublicIpAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cpc.ref.Append("coordinator_public_ip_access_enabled"))
}

// CoordinatorServerEdition returns a reference to field coordinator_server_edition of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) CoordinatorServerEdition() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("coordinator_server_edition"))
}

// CoordinatorStorageQuotaInMb returns a reference to field coordinator_storage_quota_in_mb of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) CoordinatorStorageQuotaInMb() terra.NumberValue {
	return terra.ReferenceAsNumber(cpc.ref.Append("coordinator_storage_quota_in_mb"))
}

// CoordinatorVcoreCount returns a reference to field coordinator_vcore_count of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) CoordinatorVcoreCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cpc.ref.Append("coordinator_vcore_count"))
}

// EarliestRestoreTime returns a reference to field earliest_restore_time of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) EarliestRestoreTime() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("earliest_restore_time"))
}

// HaEnabled returns a reference to field ha_enabled of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) HaEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cpc.ref.Append("ha_enabled"))
}

// Id returns a reference to field id of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("name"))
}

// NodeCount returns a reference to field node_count of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cpc.ref.Append("node_count"))
}

// NodePublicIpAccessEnabled returns a reference to field node_public_ip_access_enabled of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) NodePublicIpAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cpc.ref.Append("node_public_ip_access_enabled"))
}

// NodeServerEdition returns a reference to field node_server_edition of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) NodeServerEdition() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("node_server_edition"))
}

// NodeStorageQuotaInMb returns a reference to field node_storage_quota_in_mb of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) NodeStorageQuotaInMb() terra.NumberValue {
	return terra.ReferenceAsNumber(cpc.ref.Append("node_storage_quota_in_mb"))
}

// NodeVcores returns a reference to field node_vcores of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) NodeVcores() terra.NumberValue {
	return terra.ReferenceAsNumber(cpc.ref.Append("node_vcores"))
}

// PointInTimeInUtc returns a reference to field point_in_time_in_utc of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) PointInTimeInUtc() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("point_in_time_in_utc"))
}

// PreferredPrimaryZone returns a reference to field preferred_primary_zone of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) PreferredPrimaryZone() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("preferred_primary_zone"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("resource_group_name"))
}

// ShardsOnCoordinatorEnabled returns a reference to field shards_on_coordinator_enabled of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) ShardsOnCoordinatorEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cpc.ref.Append("shards_on_coordinator_enabled"))
}

// SourceLocation returns a reference to field source_location of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) SourceLocation() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("source_location"))
}

// SourceResourceId returns a reference to field source_resource_id of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) SourceResourceId() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("source_resource_id"))
}

// SqlVersion returns a reference to field sql_version of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) SqlVersion() terra.StringValue {
	return terra.ReferenceAsString(cpc.ref.Append("sql_version"))
}

// Tags returns a reference to field tags of azurerm_cosmosdb_postgresql_cluster.
func (cpc cosmosdbPostgresqlClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cpc.ref.Append("tags"))
}

func (cpc cosmosdbPostgresqlClusterAttributes) MaintenanceWindow() terra.ListValue[cosmosdbpostgresqlcluster.MaintenanceWindowAttributes] {
	return terra.ReferenceAsList[cosmosdbpostgresqlcluster.MaintenanceWindowAttributes](cpc.ref.Append("maintenance_window"))
}

func (cpc cosmosdbPostgresqlClusterAttributes) Timeouts() cosmosdbpostgresqlcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbpostgresqlcluster.TimeoutsAttributes](cpc.ref.Append("timeouts"))
}

type cosmosdbPostgresqlClusterState struct {
	AdministratorLoginPassword       string                                             `json:"administrator_login_password"`
	CitusVersion                     string                                             `json:"citus_version"`
	CoordinatorPublicIpAccessEnabled bool                                               `json:"coordinator_public_ip_access_enabled"`
	CoordinatorServerEdition         string                                             `json:"coordinator_server_edition"`
	CoordinatorStorageQuotaInMb      float64                                            `json:"coordinator_storage_quota_in_mb"`
	CoordinatorVcoreCount            float64                                            `json:"coordinator_vcore_count"`
	EarliestRestoreTime              string                                             `json:"earliest_restore_time"`
	HaEnabled                        bool                                               `json:"ha_enabled"`
	Id                               string                                             `json:"id"`
	Location                         string                                             `json:"location"`
	Name                             string                                             `json:"name"`
	NodeCount                        float64                                            `json:"node_count"`
	NodePublicIpAccessEnabled        bool                                               `json:"node_public_ip_access_enabled"`
	NodeServerEdition                string                                             `json:"node_server_edition"`
	NodeStorageQuotaInMb             float64                                            `json:"node_storage_quota_in_mb"`
	NodeVcores                       float64                                            `json:"node_vcores"`
	PointInTimeInUtc                 string                                             `json:"point_in_time_in_utc"`
	PreferredPrimaryZone             string                                             `json:"preferred_primary_zone"`
	ResourceGroupName                string                                             `json:"resource_group_name"`
	ShardsOnCoordinatorEnabled       bool                                               `json:"shards_on_coordinator_enabled"`
	SourceLocation                   string                                             `json:"source_location"`
	SourceResourceId                 string                                             `json:"source_resource_id"`
	SqlVersion                       string                                             `json:"sql_version"`
	Tags                             map[string]string                                  `json:"tags"`
	MaintenanceWindow                []cosmosdbpostgresqlcluster.MaintenanceWindowState `json:"maintenance_window"`
	Timeouts                         *cosmosdbpostgresqlcluster.TimeoutsState           `json:"timeouts"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqldatabase "github.com/golingon/terraproviders/azurerm/3.64.0/mssqldatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlDatabase creates a new instance of [MssqlDatabase].
func NewMssqlDatabase(name string, args MssqlDatabaseArgs) *MssqlDatabase {
	return &MssqlDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlDatabase)(nil)

// MssqlDatabase represents the Terraform resource azurerm_mssql_database.
type MssqlDatabase struct {
	Name      string
	Args      MssqlDatabaseArgs
	state     *mssqlDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlDatabase].
func (md *MssqlDatabase) Type() string {
	return "azurerm_mssql_database"
}

// LocalName returns the local name for [MssqlDatabase].
func (md *MssqlDatabase) LocalName() string {
	return md.Name
}

// Configuration returns the configuration (args) for [MssqlDatabase].
func (md *MssqlDatabase) Configuration() interface{} {
	return md.Args
}

// DependOn is used for other resources to depend on [MssqlDatabase].
func (md *MssqlDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(md)
}

// Dependencies returns the list of resources [MssqlDatabase] depends_on.
func (md *MssqlDatabase) Dependencies() terra.Dependencies {
	return md.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlDatabase].
func (md *MssqlDatabase) LifecycleManagement() *terra.Lifecycle {
	return md.Lifecycle
}

// Attributes returns the attributes for [MssqlDatabase].
func (md *MssqlDatabase) Attributes() mssqlDatabaseAttributes {
	return mssqlDatabaseAttributes{ref: terra.ReferenceResource(md)}
}

// ImportState imports the given attribute values into [MssqlDatabase]'s state.
func (md *MssqlDatabase) ImportState(av io.Reader) error {
	md.state = &mssqlDatabaseState{}
	if err := json.NewDecoder(av).Decode(md.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", md.Type(), md.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlDatabase] has state.
func (md *MssqlDatabase) State() (*mssqlDatabaseState, bool) {
	return md.state, md.state != nil
}

// StateMust returns the state for [MssqlDatabase]. Panics if the state is nil.
func (md *MssqlDatabase) StateMust() *mssqlDatabaseState {
	if md.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", md.Type(), md.LocalName()))
	}
	return md.state
}

// MssqlDatabaseArgs contains the configurations for azurerm_mssql_database.
type MssqlDatabaseArgs struct {
	// AutoPauseDelayInMinutes: number, optional
	AutoPauseDelayInMinutes terra.NumberValue `hcl:"auto_pause_delay_in_minutes,attr"`
	// Collation: string, optional
	Collation terra.StringValue `hcl:"collation,attr"`
	// CreateMode: string, optional
	CreateMode terra.StringValue `hcl:"create_mode,attr"`
	// CreationSourceDatabaseId: string, optional
	CreationSourceDatabaseId terra.StringValue `hcl:"creation_source_database_id,attr"`
	// ElasticPoolId: string, optional
	ElasticPoolId terra.StringValue `hcl:"elastic_pool_id,attr"`
	// GeoBackupEnabled: bool, optional
	GeoBackupEnabled terra.BoolValue `hcl:"geo_backup_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LedgerEnabled: bool, optional
	LedgerEnabled terra.BoolValue `hcl:"ledger_enabled,attr"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// MaintenanceConfigurationName: string, optional
	MaintenanceConfigurationName terra.StringValue `hcl:"maintenance_configuration_name,attr"`
	// MaxSizeGb: number, optional
	MaxSizeGb terra.NumberValue `hcl:"max_size_gb,attr"`
	// MinCapacity: number, optional
	MinCapacity terra.NumberValue `hcl:"min_capacity,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ReadReplicaCount: number, optional
	ReadReplicaCount terra.NumberValue `hcl:"read_replica_count,attr"`
	// ReadScale: bool, optional
	ReadScale terra.BoolValue `hcl:"read_scale,attr"`
	// RecoverDatabaseId: string, optional
	RecoverDatabaseId terra.StringValue `hcl:"recover_database_id,attr"`
	// RestoreDroppedDatabaseId: string, optional
	RestoreDroppedDatabaseId terra.StringValue `hcl:"restore_dropped_database_id,attr"`
	// RestorePointInTime: string, optional
	RestorePointInTime terra.StringValue `hcl:"restore_point_in_time,attr"`
	// SampleName: string, optional
	SampleName terra.StringValue `hcl:"sample_name,attr"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// SkuName: string, optional
	SkuName terra.StringValue `hcl:"sku_name,attr"`
	// StorageAccountType: string, optional
	StorageAccountType terra.StringValue `hcl:"storage_account_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TransparentDataEncryptionEnabled: bool, optional
	TransparentDataEncryptionEnabled terra.BoolValue `hcl:"transparent_data_encryption_enabled,attr"`
	// ZoneRedundant: bool, optional
	ZoneRedundant terra.BoolValue `hcl:"zone_redundant,attr"`
	// Import: optional
	Import *mssqldatabase.Import `hcl:"import,block"`
	// LongTermRetentionPolicy: optional
	LongTermRetentionPolicy *mssqldatabase.LongTermRetentionPolicy `hcl:"long_term_retention_policy,block"`
	// ShortTermRetentionPolicy: optional
	ShortTermRetentionPolicy *mssqldatabase.ShortTermRetentionPolicy `hcl:"short_term_retention_policy,block"`
	// ThreatDetectionPolicy: optional
	ThreatDetectionPolicy *mssqldatabase.ThreatDetectionPolicy `hcl:"threat_detection_policy,block"`
	// Timeouts: optional
	Timeouts *mssqldatabase.Timeouts `hcl:"timeouts,block"`
}
type mssqlDatabaseAttributes struct {
	ref terra.Reference
}

// AutoPauseDelayInMinutes returns a reference to field auto_pause_delay_in_minutes of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) AutoPauseDelayInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(md.ref.Append("auto_pause_delay_in_minutes"))
}

// Collation returns a reference to field collation of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("collation"))
}

// CreateMode returns a reference to field create_mode of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) CreateMode() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("create_mode"))
}

// CreationSourceDatabaseId returns a reference to field creation_source_database_id of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) CreationSourceDatabaseId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("creation_source_database_id"))
}

// ElasticPoolId returns a reference to field elastic_pool_id of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) ElasticPoolId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("elastic_pool_id"))
}

// GeoBackupEnabled returns a reference to field geo_backup_enabled of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) GeoBackupEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(md.ref.Append("geo_backup_enabled"))
}

// Id returns a reference to field id of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("id"))
}

// LedgerEnabled returns a reference to field ledger_enabled of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) LedgerEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(md.ref.Append("ledger_enabled"))
}

// LicenseType returns a reference to field license_type of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("license_type"))
}

// MaintenanceConfigurationName returns a reference to field maintenance_configuration_name of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) MaintenanceConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("maintenance_configuration_name"))
}

// MaxSizeGb returns a reference to field max_size_gb of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) MaxSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(md.ref.Append("max_size_gb"))
}

// MinCapacity returns a reference to field min_capacity of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) MinCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(md.ref.Append("min_capacity"))
}

// Name returns a reference to field name of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("name"))
}

// ReadReplicaCount returns a reference to field read_replica_count of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) ReadReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(md.ref.Append("read_replica_count"))
}

// ReadScale returns a reference to field read_scale of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) ReadScale() terra.BoolValue {
	return terra.ReferenceAsBool(md.ref.Append("read_scale"))
}

// RecoverDatabaseId returns a reference to field recover_database_id of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) RecoverDatabaseId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("recover_database_id"))
}

// RestoreDroppedDatabaseId returns a reference to field restore_dropped_database_id of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) RestoreDroppedDatabaseId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("restore_dropped_database_id"))
}

// RestorePointInTime returns a reference to field restore_point_in_time of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) RestorePointInTime() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("restore_point_in_time"))
}

// SampleName returns a reference to field sample_name of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) SampleName() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("sample_name"))
}

// ServerId returns a reference to field server_id of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("server_id"))
}

// SkuName returns a reference to field sku_name of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("sku_name"))
}

// StorageAccountType returns a reference to field storage_account_type of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("storage_account_type"))
}

// Tags returns a reference to field tags of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](md.ref.Append("tags"))
}

// TransparentDataEncryptionEnabled returns a reference to field transparent_data_encryption_enabled of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) TransparentDataEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(md.ref.Append("transparent_data_encryption_enabled"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_mssql_database.
func (md mssqlDatabaseAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(md.ref.Append("zone_redundant"))
}

func (md mssqlDatabaseAttributes) Import() terra.ListValue[mssqldatabase.ImportAttributes] {
	return terra.ReferenceAsList[mssqldatabase.ImportAttributes](md.ref.Append("import"))
}

func (md mssqlDatabaseAttributes) LongTermRetentionPolicy() terra.ListValue[mssqldatabase.LongTermRetentionPolicyAttributes] {
	return terra.ReferenceAsList[mssqldatabase.LongTermRetentionPolicyAttributes](md.ref.Append("long_term_retention_policy"))
}

func (md mssqlDatabaseAttributes) ShortTermRetentionPolicy() terra.ListValue[mssqldatabase.ShortTermRetentionPolicyAttributes] {
	return terra.ReferenceAsList[mssqldatabase.ShortTermRetentionPolicyAttributes](md.ref.Append("short_term_retention_policy"))
}

func (md mssqlDatabaseAttributes) ThreatDetectionPolicy() terra.ListValue[mssqldatabase.ThreatDetectionPolicyAttributes] {
	return terra.ReferenceAsList[mssqldatabase.ThreatDetectionPolicyAttributes](md.ref.Append("threat_detection_policy"))
}

func (md mssqlDatabaseAttributes) Timeouts() mssqldatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqldatabase.TimeoutsAttributes](md.ref.Append("timeouts"))
}

type mssqlDatabaseState struct {
	AutoPauseDelayInMinutes          float64                                       `json:"auto_pause_delay_in_minutes"`
	Collation                        string                                        `json:"collation"`
	CreateMode                       string                                        `json:"create_mode"`
	CreationSourceDatabaseId         string                                        `json:"creation_source_database_id"`
	ElasticPoolId                    string                                        `json:"elastic_pool_id"`
	GeoBackupEnabled                 bool                                          `json:"geo_backup_enabled"`
	Id                               string                                        `json:"id"`
	LedgerEnabled                    bool                                          `json:"ledger_enabled"`
	LicenseType                      string                                        `json:"license_type"`
	MaintenanceConfigurationName     string                                        `json:"maintenance_configuration_name"`
	MaxSizeGb                        float64                                       `json:"max_size_gb"`
	MinCapacity                      float64                                       `json:"min_capacity"`
	Name                             string                                        `json:"name"`
	ReadReplicaCount                 float64                                       `json:"read_replica_count"`
	ReadScale                        bool                                          `json:"read_scale"`
	RecoverDatabaseId                string                                        `json:"recover_database_id"`
	RestoreDroppedDatabaseId         string                                        `json:"restore_dropped_database_id"`
	RestorePointInTime               string                                        `json:"restore_point_in_time"`
	SampleName                       string                                        `json:"sample_name"`
	ServerId                         string                                        `json:"server_id"`
	SkuName                          string                                        `json:"sku_name"`
	StorageAccountType               string                                        `json:"storage_account_type"`
	Tags                             map[string]string                             `json:"tags"`
	TransparentDataEncryptionEnabled bool                                          `json:"transparent_data_encryption_enabled"`
	ZoneRedundant                    bool                                          `json:"zone_redundant"`
	Import                           []mssqldatabase.ImportState                   `json:"import"`
	LongTermRetentionPolicy          []mssqldatabase.LongTermRetentionPolicyState  `json:"long_term_retention_policy"`
	ShortTermRetentionPolicy         []mssqldatabase.ShortTermRetentionPolicyState `json:"short_term_retention_policy"`
	ThreatDetectionPolicy            []mssqldatabase.ThreatDetectionPolicyState    `json:"threat_detection_policy"`
	Timeouts                         *mssqldatabase.TimeoutsState                  `json:"timeouts"`
}

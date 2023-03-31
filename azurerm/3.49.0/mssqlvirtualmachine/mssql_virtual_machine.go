// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mssqlvirtualmachine

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Assessment struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// RunImmediately: bool, optional
	RunImmediately terra.BoolValue `hcl:"run_immediately,attr"`
	// Schedule: optional
	Schedule *Schedule `hcl:"schedule,block"`
}

type Schedule struct {
	// DayOfWeek: string, required
	DayOfWeek terra.StringValue `hcl:"day_of_week,attr" validate:"required"`
	// MonthlyOccurrence: number, optional
	MonthlyOccurrence terra.NumberValue `hcl:"monthly_occurrence,attr"`
	// StartTime: string, required
	StartTime terra.StringValue `hcl:"start_time,attr" validate:"required"`
	// WeeklyInterval: number, optional
	WeeklyInterval terra.NumberValue `hcl:"weekly_interval,attr"`
}

type AutoBackup struct {
	// EncryptionEnabled: bool, optional
	EncryptionEnabled terra.BoolValue `hcl:"encryption_enabled,attr"`
	// EncryptionPassword: string, optional
	EncryptionPassword terra.StringValue `hcl:"encryption_password,attr"`
	// RetentionPeriodInDays: number, required
	RetentionPeriodInDays terra.NumberValue `hcl:"retention_period_in_days,attr" validate:"required"`
	// StorageAccountAccessKey: string, required
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr" validate:"required"`
	// StorageBlobEndpoint: string, required
	StorageBlobEndpoint terra.StringValue `hcl:"storage_blob_endpoint,attr" validate:"required"`
	// SystemDatabasesBackupEnabled: bool, optional
	SystemDatabasesBackupEnabled terra.BoolValue `hcl:"system_databases_backup_enabled,attr"`
	// ManualSchedule: optional
	ManualSchedule *ManualSchedule `hcl:"manual_schedule,block"`
}

type ManualSchedule struct {
	// DaysOfWeek: set of string, optional
	DaysOfWeek terra.SetValue[terra.StringValue] `hcl:"days_of_week,attr"`
	// FullBackupFrequency: string, required
	FullBackupFrequency terra.StringValue `hcl:"full_backup_frequency,attr" validate:"required"`
	// FullBackupStartHour: number, required
	FullBackupStartHour terra.NumberValue `hcl:"full_backup_start_hour,attr" validate:"required"`
	// FullBackupWindowInHours: number, required
	FullBackupWindowInHours terra.NumberValue `hcl:"full_backup_window_in_hours,attr" validate:"required"`
	// LogBackupFrequencyInMinutes: number, required
	LogBackupFrequencyInMinutes terra.NumberValue `hcl:"log_backup_frequency_in_minutes,attr" validate:"required"`
}

type AutoPatching struct {
	// DayOfWeek: string, required
	DayOfWeek terra.StringValue `hcl:"day_of_week,attr" validate:"required"`
	// MaintenanceWindowDurationInMinutes: number, required
	MaintenanceWindowDurationInMinutes terra.NumberValue `hcl:"maintenance_window_duration_in_minutes,attr" validate:"required"`
	// MaintenanceWindowStartingHour: number, required
	MaintenanceWindowStartingHour terra.NumberValue `hcl:"maintenance_window_starting_hour,attr" validate:"required"`
}

type KeyVaultCredential struct {
	// KeyVaultUrl: string, required
	KeyVaultUrl terra.StringValue `hcl:"key_vault_url,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServicePrincipalName: string, required
	ServicePrincipalName terra.StringValue `hcl:"service_principal_name,attr" validate:"required"`
	// ServicePrincipalSecret: string, required
	ServicePrincipalSecret terra.StringValue `hcl:"service_principal_secret,attr" validate:"required"`
}

type SqlInstance struct {
	// AdhocWorkloadsOptimizationEnabled: bool, optional
	AdhocWorkloadsOptimizationEnabled terra.BoolValue `hcl:"adhoc_workloads_optimization_enabled,attr"`
	// Collation: string, optional
	Collation terra.StringValue `hcl:"collation,attr"`
	// InstantFileInitializationEnabled: bool, optional
	InstantFileInitializationEnabled terra.BoolValue `hcl:"instant_file_initialization_enabled,attr"`
	// LockPagesInMemoryEnabled: bool, optional
	LockPagesInMemoryEnabled terra.BoolValue `hcl:"lock_pages_in_memory_enabled,attr"`
	// MaxDop: number, optional
	MaxDop terra.NumberValue `hcl:"max_dop,attr"`
	// MaxServerMemoryMb: number, optional
	MaxServerMemoryMb terra.NumberValue `hcl:"max_server_memory_mb,attr"`
	// MinServerMemoryMb: number, optional
	MinServerMemoryMb terra.NumberValue `hcl:"min_server_memory_mb,attr"`
}

type StorageConfiguration struct {
	// DiskType: string, required
	DiskType terra.StringValue `hcl:"disk_type,attr" validate:"required"`
	// StorageWorkloadType: string, required
	StorageWorkloadType terra.StringValue `hcl:"storage_workload_type,attr" validate:"required"`
	// SystemDbOnDataDiskEnabled: bool, optional
	SystemDbOnDataDiskEnabled terra.BoolValue `hcl:"system_db_on_data_disk_enabled,attr"`
	// DataSettings: optional
	DataSettings *DataSettings `hcl:"data_settings,block"`
	// LogSettings: optional
	LogSettings *LogSettings `hcl:"log_settings,block"`
	// TempDbSettings: optional
	TempDbSettings *TempDbSettings `hcl:"temp_db_settings,block"`
}

type DataSettings struct {
	// DefaultFilePath: string, required
	DefaultFilePath terra.StringValue `hcl:"default_file_path,attr" validate:"required"`
	// Luns: list of number, required
	Luns terra.ListValue[terra.NumberValue] `hcl:"luns,attr" validate:"required"`
}

type LogSettings struct {
	// DefaultFilePath: string, required
	DefaultFilePath terra.StringValue `hcl:"default_file_path,attr" validate:"required"`
	// Luns: list of number, required
	Luns terra.ListValue[terra.NumberValue] `hcl:"luns,attr" validate:"required"`
}

type TempDbSettings struct {
	// DataFileCount: number, optional
	DataFileCount terra.NumberValue `hcl:"data_file_count,attr"`
	// DataFileGrowthInMb: number, optional
	DataFileGrowthInMb terra.NumberValue `hcl:"data_file_growth_in_mb,attr"`
	// DataFileSizeMb: number, optional
	DataFileSizeMb terra.NumberValue `hcl:"data_file_size_mb,attr"`
	// DefaultFilePath: string, required
	DefaultFilePath terra.StringValue `hcl:"default_file_path,attr" validate:"required"`
	// LogFileGrowthMb: number, optional
	LogFileGrowthMb terra.NumberValue `hcl:"log_file_growth_mb,attr"`
	// LogFileSizeMb: number, optional
	LogFileSizeMb terra.NumberValue `hcl:"log_file_size_mb,attr"`
	// Luns: list of number, required
	Luns terra.ListValue[terra.NumberValue] `hcl:"luns,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AssessmentAttributes struct {
	ref terra.Reference
}

func (a AssessmentAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AssessmentAttributes) InternalWithRef(ref terra.Reference) AssessmentAttributes {
	return AssessmentAttributes{ref: ref}
}

func (a AssessmentAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AssessmentAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(a.ref.Append("enabled"))
}

func (a AssessmentAttributes) RunImmediately() terra.BoolValue {
	return terra.ReferenceBool(a.ref.Append("run_immediately"))
}

func (a AssessmentAttributes) Schedule() terra.ListValue[ScheduleAttributes] {
	return terra.ReferenceList[ScheduleAttributes](a.ref.Append("schedule"))
}

type ScheduleAttributes struct {
	ref terra.Reference
}

func (s ScheduleAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s ScheduleAttributes) InternalWithRef(ref terra.Reference) ScheduleAttributes {
	return ScheduleAttributes{ref: ref}
}

func (s ScheduleAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s ScheduleAttributes) DayOfWeek() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("day_of_week"))
}

func (s ScheduleAttributes) MonthlyOccurrence() terra.NumberValue {
	return terra.ReferenceNumber(s.ref.Append("monthly_occurrence"))
}

func (s ScheduleAttributes) StartTime() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("start_time"))
}

func (s ScheduleAttributes) WeeklyInterval() terra.NumberValue {
	return terra.ReferenceNumber(s.ref.Append("weekly_interval"))
}

type AutoBackupAttributes struct {
	ref terra.Reference
}

func (ab AutoBackupAttributes) InternalRef() terra.Reference {
	return ab.ref
}

func (ab AutoBackupAttributes) InternalWithRef(ref terra.Reference) AutoBackupAttributes {
	return AutoBackupAttributes{ref: ref}
}

func (ab AutoBackupAttributes) InternalTokens() hclwrite.Tokens {
	return ab.ref.InternalTokens()
}

func (ab AutoBackupAttributes) EncryptionEnabled() terra.BoolValue {
	return terra.ReferenceBool(ab.ref.Append("encryption_enabled"))
}

func (ab AutoBackupAttributes) EncryptionPassword() terra.StringValue {
	return terra.ReferenceString(ab.ref.Append("encryption_password"))
}

func (ab AutoBackupAttributes) RetentionPeriodInDays() terra.NumberValue {
	return terra.ReferenceNumber(ab.ref.Append("retention_period_in_days"))
}

func (ab AutoBackupAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceString(ab.ref.Append("storage_account_access_key"))
}

func (ab AutoBackupAttributes) StorageBlobEndpoint() terra.StringValue {
	return terra.ReferenceString(ab.ref.Append("storage_blob_endpoint"))
}

func (ab AutoBackupAttributes) SystemDatabasesBackupEnabled() terra.BoolValue {
	return terra.ReferenceBool(ab.ref.Append("system_databases_backup_enabled"))
}

func (ab AutoBackupAttributes) ManualSchedule() terra.ListValue[ManualScheduleAttributes] {
	return terra.ReferenceList[ManualScheduleAttributes](ab.ref.Append("manual_schedule"))
}

type ManualScheduleAttributes struct {
	ref terra.Reference
}

func (ms ManualScheduleAttributes) InternalRef() terra.Reference {
	return ms.ref
}

func (ms ManualScheduleAttributes) InternalWithRef(ref terra.Reference) ManualScheduleAttributes {
	return ManualScheduleAttributes{ref: ref}
}

func (ms ManualScheduleAttributes) InternalTokens() hclwrite.Tokens {
	return ms.ref.InternalTokens()
}

func (ms ManualScheduleAttributes) DaysOfWeek() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ms.ref.Append("days_of_week"))
}

func (ms ManualScheduleAttributes) FullBackupFrequency() terra.StringValue {
	return terra.ReferenceString(ms.ref.Append("full_backup_frequency"))
}

func (ms ManualScheduleAttributes) FullBackupStartHour() terra.NumberValue {
	return terra.ReferenceNumber(ms.ref.Append("full_backup_start_hour"))
}

func (ms ManualScheduleAttributes) FullBackupWindowInHours() terra.NumberValue {
	return terra.ReferenceNumber(ms.ref.Append("full_backup_window_in_hours"))
}

func (ms ManualScheduleAttributes) LogBackupFrequencyInMinutes() terra.NumberValue {
	return terra.ReferenceNumber(ms.ref.Append("log_backup_frequency_in_minutes"))
}

type AutoPatchingAttributes struct {
	ref terra.Reference
}

func (ap AutoPatchingAttributes) InternalRef() terra.Reference {
	return ap.ref
}

func (ap AutoPatchingAttributes) InternalWithRef(ref terra.Reference) AutoPatchingAttributes {
	return AutoPatchingAttributes{ref: ref}
}

func (ap AutoPatchingAttributes) InternalTokens() hclwrite.Tokens {
	return ap.ref.InternalTokens()
}

func (ap AutoPatchingAttributes) DayOfWeek() terra.StringValue {
	return terra.ReferenceString(ap.ref.Append("day_of_week"))
}

func (ap AutoPatchingAttributes) MaintenanceWindowDurationInMinutes() terra.NumberValue {
	return terra.ReferenceNumber(ap.ref.Append("maintenance_window_duration_in_minutes"))
}

func (ap AutoPatchingAttributes) MaintenanceWindowStartingHour() terra.NumberValue {
	return terra.ReferenceNumber(ap.ref.Append("maintenance_window_starting_hour"))
}

type KeyVaultCredentialAttributes struct {
	ref terra.Reference
}

func (kvc KeyVaultCredentialAttributes) InternalRef() terra.Reference {
	return kvc.ref
}

func (kvc KeyVaultCredentialAttributes) InternalWithRef(ref terra.Reference) KeyVaultCredentialAttributes {
	return KeyVaultCredentialAttributes{ref: ref}
}

func (kvc KeyVaultCredentialAttributes) InternalTokens() hclwrite.Tokens {
	return kvc.ref.InternalTokens()
}

func (kvc KeyVaultCredentialAttributes) KeyVaultUrl() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("key_vault_url"))
}

func (kvc KeyVaultCredentialAttributes) Name() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("name"))
}

func (kvc KeyVaultCredentialAttributes) ServicePrincipalName() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("service_principal_name"))
}

func (kvc KeyVaultCredentialAttributes) ServicePrincipalSecret() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("service_principal_secret"))
}

type SqlInstanceAttributes struct {
	ref terra.Reference
}

func (si SqlInstanceAttributes) InternalRef() terra.Reference {
	return si.ref
}

func (si SqlInstanceAttributes) InternalWithRef(ref terra.Reference) SqlInstanceAttributes {
	return SqlInstanceAttributes{ref: ref}
}

func (si SqlInstanceAttributes) InternalTokens() hclwrite.Tokens {
	return si.ref.InternalTokens()
}

func (si SqlInstanceAttributes) AdhocWorkloadsOptimizationEnabled() terra.BoolValue {
	return terra.ReferenceBool(si.ref.Append("adhoc_workloads_optimization_enabled"))
}

func (si SqlInstanceAttributes) Collation() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("collation"))
}

func (si SqlInstanceAttributes) InstantFileInitializationEnabled() terra.BoolValue {
	return terra.ReferenceBool(si.ref.Append("instant_file_initialization_enabled"))
}

func (si SqlInstanceAttributes) LockPagesInMemoryEnabled() terra.BoolValue {
	return terra.ReferenceBool(si.ref.Append("lock_pages_in_memory_enabled"))
}

func (si SqlInstanceAttributes) MaxDop() terra.NumberValue {
	return terra.ReferenceNumber(si.ref.Append("max_dop"))
}

func (si SqlInstanceAttributes) MaxServerMemoryMb() terra.NumberValue {
	return terra.ReferenceNumber(si.ref.Append("max_server_memory_mb"))
}

func (si SqlInstanceAttributes) MinServerMemoryMb() terra.NumberValue {
	return terra.ReferenceNumber(si.ref.Append("min_server_memory_mb"))
}

type StorageConfigurationAttributes struct {
	ref terra.Reference
}

func (sc StorageConfigurationAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc StorageConfigurationAttributes) InternalWithRef(ref terra.Reference) StorageConfigurationAttributes {
	return StorageConfigurationAttributes{ref: ref}
}

func (sc StorageConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc StorageConfigurationAttributes) DiskType() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("disk_type"))
}

func (sc StorageConfigurationAttributes) StorageWorkloadType() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("storage_workload_type"))
}

func (sc StorageConfigurationAttributes) SystemDbOnDataDiskEnabled() terra.BoolValue {
	return terra.ReferenceBool(sc.ref.Append("system_db_on_data_disk_enabled"))
}

func (sc StorageConfigurationAttributes) DataSettings() terra.ListValue[DataSettingsAttributes] {
	return terra.ReferenceList[DataSettingsAttributes](sc.ref.Append("data_settings"))
}

func (sc StorageConfigurationAttributes) LogSettings() terra.ListValue[LogSettingsAttributes] {
	return terra.ReferenceList[LogSettingsAttributes](sc.ref.Append("log_settings"))
}

func (sc StorageConfigurationAttributes) TempDbSettings() terra.ListValue[TempDbSettingsAttributes] {
	return terra.ReferenceList[TempDbSettingsAttributes](sc.ref.Append("temp_db_settings"))
}

type DataSettingsAttributes struct {
	ref terra.Reference
}

func (ds DataSettingsAttributes) InternalRef() terra.Reference {
	return ds.ref
}

func (ds DataSettingsAttributes) InternalWithRef(ref terra.Reference) DataSettingsAttributes {
	return DataSettingsAttributes{ref: ref}
}

func (ds DataSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ds.ref.InternalTokens()
}

func (ds DataSettingsAttributes) DefaultFilePath() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("default_file_path"))
}

func (ds DataSettingsAttributes) Luns() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceList[terra.NumberValue](ds.ref.Append("luns"))
}

type LogSettingsAttributes struct {
	ref terra.Reference
}

func (ls LogSettingsAttributes) InternalRef() terra.Reference {
	return ls.ref
}

func (ls LogSettingsAttributes) InternalWithRef(ref terra.Reference) LogSettingsAttributes {
	return LogSettingsAttributes{ref: ref}
}

func (ls LogSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ls.ref.InternalTokens()
}

func (ls LogSettingsAttributes) DefaultFilePath() terra.StringValue {
	return terra.ReferenceString(ls.ref.Append("default_file_path"))
}

func (ls LogSettingsAttributes) Luns() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceList[terra.NumberValue](ls.ref.Append("luns"))
}

type TempDbSettingsAttributes struct {
	ref terra.Reference
}

func (tds TempDbSettingsAttributes) InternalRef() terra.Reference {
	return tds.ref
}

func (tds TempDbSettingsAttributes) InternalWithRef(ref terra.Reference) TempDbSettingsAttributes {
	return TempDbSettingsAttributes{ref: ref}
}

func (tds TempDbSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return tds.ref.InternalTokens()
}

func (tds TempDbSettingsAttributes) DataFileCount() terra.NumberValue {
	return terra.ReferenceNumber(tds.ref.Append("data_file_count"))
}

func (tds TempDbSettingsAttributes) DataFileGrowthInMb() terra.NumberValue {
	return terra.ReferenceNumber(tds.ref.Append("data_file_growth_in_mb"))
}

func (tds TempDbSettingsAttributes) DataFileSizeMb() terra.NumberValue {
	return terra.ReferenceNumber(tds.ref.Append("data_file_size_mb"))
}

func (tds TempDbSettingsAttributes) DefaultFilePath() terra.StringValue {
	return terra.ReferenceString(tds.ref.Append("default_file_path"))
}

func (tds TempDbSettingsAttributes) LogFileGrowthMb() terra.NumberValue {
	return terra.ReferenceNumber(tds.ref.Append("log_file_growth_mb"))
}

func (tds TempDbSettingsAttributes) LogFileSizeMb() terra.NumberValue {
	return terra.ReferenceNumber(tds.ref.Append("log_file_size_mb"))
}

func (tds TempDbSettingsAttributes) Luns() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceList[terra.NumberValue](tds.ref.Append("luns"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type AssessmentState struct {
	Enabled        bool            `json:"enabled"`
	RunImmediately bool            `json:"run_immediately"`
	Schedule       []ScheduleState `json:"schedule"`
}

type ScheduleState struct {
	DayOfWeek         string  `json:"day_of_week"`
	MonthlyOccurrence float64 `json:"monthly_occurrence"`
	StartTime         string  `json:"start_time"`
	WeeklyInterval    float64 `json:"weekly_interval"`
}

type AutoBackupState struct {
	EncryptionEnabled            bool                  `json:"encryption_enabled"`
	EncryptionPassword           string                `json:"encryption_password"`
	RetentionPeriodInDays        float64               `json:"retention_period_in_days"`
	StorageAccountAccessKey      string                `json:"storage_account_access_key"`
	StorageBlobEndpoint          string                `json:"storage_blob_endpoint"`
	SystemDatabasesBackupEnabled bool                  `json:"system_databases_backup_enabled"`
	ManualSchedule               []ManualScheduleState `json:"manual_schedule"`
}

type ManualScheduleState struct {
	DaysOfWeek                  []string `json:"days_of_week"`
	FullBackupFrequency         string   `json:"full_backup_frequency"`
	FullBackupStartHour         float64  `json:"full_backup_start_hour"`
	FullBackupWindowInHours     float64  `json:"full_backup_window_in_hours"`
	LogBackupFrequencyInMinutes float64  `json:"log_backup_frequency_in_minutes"`
}

type AutoPatchingState struct {
	DayOfWeek                          string  `json:"day_of_week"`
	MaintenanceWindowDurationInMinutes float64 `json:"maintenance_window_duration_in_minutes"`
	MaintenanceWindowStartingHour      float64 `json:"maintenance_window_starting_hour"`
}

type KeyVaultCredentialState struct {
	KeyVaultUrl            string `json:"key_vault_url"`
	Name                   string `json:"name"`
	ServicePrincipalName   string `json:"service_principal_name"`
	ServicePrincipalSecret string `json:"service_principal_secret"`
}

type SqlInstanceState struct {
	AdhocWorkloadsOptimizationEnabled bool    `json:"adhoc_workloads_optimization_enabled"`
	Collation                         string  `json:"collation"`
	InstantFileInitializationEnabled  bool    `json:"instant_file_initialization_enabled"`
	LockPagesInMemoryEnabled          bool    `json:"lock_pages_in_memory_enabled"`
	MaxDop                            float64 `json:"max_dop"`
	MaxServerMemoryMb                 float64 `json:"max_server_memory_mb"`
	MinServerMemoryMb                 float64 `json:"min_server_memory_mb"`
}

type StorageConfigurationState struct {
	DiskType                  string                `json:"disk_type"`
	StorageWorkloadType       string                `json:"storage_workload_type"`
	SystemDbOnDataDiskEnabled bool                  `json:"system_db_on_data_disk_enabled"`
	DataSettings              []DataSettingsState   `json:"data_settings"`
	LogSettings               []LogSettingsState    `json:"log_settings"`
	TempDbSettings            []TempDbSettingsState `json:"temp_db_settings"`
}

type DataSettingsState struct {
	DefaultFilePath string    `json:"default_file_path"`
	Luns            []float64 `json:"luns"`
}

type LogSettingsState struct {
	DefaultFilePath string    `json:"default_file_path"`
	Luns            []float64 `json:"luns"`
}

type TempDbSettingsState struct {
	DataFileCount      float64   `json:"data_file_count"`
	DataFileGrowthInMb float64   `json:"data_file_growth_in_mb"`
	DataFileSizeMb     float64   `json:"data_file_size_mb"`
	DefaultFilePath    string    `json:"default_file_path"`
	LogFileGrowthMb    float64   `json:"log_file_growth_mb"`
	LogFileSizeMb      float64   `json:"log_file_size_mb"`
	Luns               []float64 `json:"luns"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

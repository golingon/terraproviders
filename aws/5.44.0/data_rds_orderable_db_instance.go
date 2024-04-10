// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataRdsOrderableDbInstance creates a new instance of [DataRdsOrderableDbInstance].
func NewDataRdsOrderableDbInstance(name string, args DataRdsOrderableDbInstanceArgs) *DataRdsOrderableDbInstance {
	return &DataRdsOrderableDbInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRdsOrderableDbInstance)(nil)

// DataRdsOrderableDbInstance represents the Terraform data resource aws_rds_orderable_db_instance.
type DataRdsOrderableDbInstance struct {
	Name string
	Args DataRdsOrderableDbInstanceArgs
}

// DataSource returns the Terraform object type for [DataRdsOrderableDbInstance].
func (rodi *DataRdsOrderableDbInstance) DataSource() string {
	return "aws_rds_orderable_db_instance"
}

// LocalName returns the local name for [DataRdsOrderableDbInstance].
func (rodi *DataRdsOrderableDbInstance) LocalName() string {
	return rodi.Name
}

// Configuration returns the configuration (args) for [DataRdsOrderableDbInstance].
func (rodi *DataRdsOrderableDbInstance) Configuration() interface{} {
	return rodi.Args
}

// Attributes returns the attributes for [DataRdsOrderableDbInstance].
func (rodi *DataRdsOrderableDbInstance) Attributes() dataRdsOrderableDbInstanceAttributes {
	return dataRdsOrderableDbInstanceAttributes{ref: terra.ReferenceDataResource(rodi)}
}

// DataRdsOrderableDbInstanceArgs contains the configurations for aws_rds_orderable_db_instance.
type DataRdsOrderableDbInstanceArgs struct {
	// AvailabilityZoneGroup: string, optional
	AvailabilityZoneGroup terra.StringValue `hcl:"availability_zone_group,attr"`
	// Engine: string, required
	Engine terra.StringValue `hcl:"engine,attr" validate:"required"`
	// EngineLatestVersion: bool, optional
	EngineLatestVersion terra.BoolValue `hcl:"engine_latest_version,attr"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceClass: string, optional
	InstanceClass terra.StringValue `hcl:"instance_class,attr"`
	// LicenseModel: string, optional
	LicenseModel terra.StringValue `hcl:"license_model,attr"`
	// PreferredEngineVersions: list of string, optional
	PreferredEngineVersions terra.ListValue[terra.StringValue] `hcl:"preferred_engine_versions,attr"`
	// PreferredInstanceClasses: list of string, optional
	PreferredInstanceClasses terra.ListValue[terra.StringValue] `hcl:"preferred_instance_classes,attr"`
	// ReadReplicaCapable: bool, optional
	ReadReplicaCapable terra.BoolValue `hcl:"read_replica_capable,attr"`
	// StorageType: string, optional
	StorageType terra.StringValue `hcl:"storage_type,attr"`
	// SupportedEngineModes: list of string, optional
	SupportedEngineModes terra.ListValue[terra.StringValue] `hcl:"supported_engine_modes,attr"`
	// SupportedNetworkTypes: list of string, optional
	SupportedNetworkTypes terra.ListValue[terra.StringValue] `hcl:"supported_network_types,attr"`
	// SupportsClusters: bool, optional
	SupportsClusters terra.BoolValue `hcl:"supports_clusters,attr"`
	// SupportsEnhancedMonitoring: bool, optional
	SupportsEnhancedMonitoring terra.BoolValue `hcl:"supports_enhanced_monitoring,attr"`
	// SupportsGlobalDatabases: bool, optional
	SupportsGlobalDatabases terra.BoolValue `hcl:"supports_global_databases,attr"`
	// SupportsIamDatabaseAuthentication: bool, optional
	SupportsIamDatabaseAuthentication terra.BoolValue `hcl:"supports_iam_database_authentication,attr"`
	// SupportsIops: bool, optional
	SupportsIops terra.BoolValue `hcl:"supports_iops,attr"`
	// SupportsKerberosAuthentication: bool, optional
	SupportsKerberosAuthentication terra.BoolValue `hcl:"supports_kerberos_authentication,attr"`
	// SupportsMultiAz: bool, optional
	SupportsMultiAz terra.BoolValue `hcl:"supports_multi_az,attr"`
	// SupportsPerformanceInsights: bool, optional
	SupportsPerformanceInsights terra.BoolValue `hcl:"supports_performance_insights,attr"`
	// SupportsStorageAutoscaling: bool, optional
	SupportsStorageAutoscaling terra.BoolValue `hcl:"supports_storage_autoscaling,attr"`
	// SupportsStorageEncryption: bool, optional
	SupportsStorageEncryption terra.BoolValue `hcl:"supports_storage_encryption,attr"`
	// Vpc: bool, optional
	Vpc terra.BoolValue `hcl:"vpc,attr"`
}
type dataRdsOrderableDbInstanceAttributes struct {
	ref terra.Reference
}

// AvailabilityZoneGroup returns a reference to field availability_zone_group of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) AvailabilityZoneGroup() terra.StringValue {
	return terra.ReferenceAsString(rodi.ref.Append("availability_zone_group"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) AvailabilityZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rodi.ref.Append("availability_zones"))
}

// Engine returns a reference to field engine of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(rodi.ref.Append("engine"))
}

// EngineLatestVersion returns a reference to field engine_latest_version of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) EngineLatestVersion() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("engine_latest_version"))
}

// EngineVersion returns a reference to field engine_version of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(rodi.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rodi.ref.Append("id"))
}

// InstanceClass returns a reference to field instance_class of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) InstanceClass() terra.StringValue {
	return terra.ReferenceAsString(rodi.ref.Append("instance_class"))
}

// LicenseModel returns a reference to field license_model of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceAsString(rodi.ref.Append("license_model"))
}

// MaxIopsPerDbInstance returns a reference to field max_iops_per_db_instance of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) MaxIopsPerDbInstance() terra.NumberValue {
	return terra.ReferenceAsNumber(rodi.ref.Append("max_iops_per_db_instance"))
}

// MaxIopsPerGib returns a reference to field max_iops_per_gib of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) MaxIopsPerGib() terra.NumberValue {
	return terra.ReferenceAsNumber(rodi.ref.Append("max_iops_per_gib"))
}

// MaxStorageSize returns a reference to field max_storage_size of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) MaxStorageSize() terra.NumberValue {
	return terra.ReferenceAsNumber(rodi.ref.Append("max_storage_size"))
}

// MinIopsPerDbInstance returns a reference to field min_iops_per_db_instance of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) MinIopsPerDbInstance() terra.NumberValue {
	return terra.ReferenceAsNumber(rodi.ref.Append("min_iops_per_db_instance"))
}

// MinIopsPerGib returns a reference to field min_iops_per_gib of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) MinIopsPerGib() terra.NumberValue {
	return terra.ReferenceAsNumber(rodi.ref.Append("min_iops_per_gib"))
}

// MinStorageSize returns a reference to field min_storage_size of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) MinStorageSize() terra.NumberValue {
	return terra.ReferenceAsNumber(rodi.ref.Append("min_storage_size"))
}

// MultiAzCapable returns a reference to field multi_az_capable of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) MultiAzCapable() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("multi_az_capable"))
}

// OutpostCapable returns a reference to field outpost_capable of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) OutpostCapable() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("outpost_capable"))
}

// PreferredEngineVersions returns a reference to field preferred_engine_versions of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) PreferredEngineVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rodi.ref.Append("preferred_engine_versions"))
}

// PreferredInstanceClasses returns a reference to field preferred_instance_classes of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) PreferredInstanceClasses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rodi.ref.Append("preferred_instance_classes"))
}

// ReadReplicaCapable returns a reference to field read_replica_capable of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) ReadReplicaCapable() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("read_replica_capable"))
}

// StorageType returns a reference to field storage_type of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(rodi.ref.Append("storage_type"))
}

// SupportedEngineModes returns a reference to field supported_engine_modes of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportedEngineModes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rodi.ref.Append("supported_engine_modes"))
}

// SupportedNetworkTypes returns a reference to field supported_network_types of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportedNetworkTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rodi.ref.Append("supported_network_types"))
}

// SupportsClusters returns a reference to field supports_clusters of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportsClusters() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("supports_clusters"))
}

// SupportsEnhancedMonitoring returns a reference to field supports_enhanced_monitoring of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportsEnhancedMonitoring() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("supports_enhanced_monitoring"))
}

// SupportsGlobalDatabases returns a reference to field supports_global_databases of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportsGlobalDatabases() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("supports_global_databases"))
}

// SupportsIamDatabaseAuthentication returns a reference to field supports_iam_database_authentication of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportsIamDatabaseAuthentication() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("supports_iam_database_authentication"))
}

// SupportsIops returns a reference to field supports_iops of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportsIops() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("supports_iops"))
}

// SupportsKerberosAuthentication returns a reference to field supports_kerberos_authentication of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportsKerberosAuthentication() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("supports_kerberos_authentication"))
}

// SupportsMultiAz returns a reference to field supports_multi_az of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportsMultiAz() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("supports_multi_az"))
}

// SupportsPerformanceInsights returns a reference to field supports_performance_insights of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportsPerformanceInsights() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("supports_performance_insights"))
}

// SupportsStorageAutoscaling returns a reference to field supports_storage_autoscaling of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportsStorageAutoscaling() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("supports_storage_autoscaling"))
}

// SupportsStorageEncryption returns a reference to field supports_storage_encryption of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) SupportsStorageEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("supports_storage_encryption"))
}

// Vpc returns a reference to field vpc of aws_rds_orderable_db_instance.
func (rodi dataRdsOrderableDbInstanceAttributes) Vpc() terra.BoolValue {
	return terra.ReferenceAsBool(rodi.ref.Append("vpc"))
}

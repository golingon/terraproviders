// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataNeptuneOrderableDbInstance creates a new instance of [DataNeptuneOrderableDbInstance].
func NewDataNeptuneOrderableDbInstance(name string, args DataNeptuneOrderableDbInstanceArgs) *DataNeptuneOrderableDbInstance {
	return &DataNeptuneOrderableDbInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNeptuneOrderableDbInstance)(nil)

// DataNeptuneOrderableDbInstance represents the Terraform data resource aws_neptune_orderable_db_instance.
type DataNeptuneOrderableDbInstance struct {
	Name string
	Args DataNeptuneOrderableDbInstanceArgs
}

// DataSource returns the Terraform object type for [DataNeptuneOrderableDbInstance].
func (nodi *DataNeptuneOrderableDbInstance) DataSource() string {
	return "aws_neptune_orderable_db_instance"
}

// LocalName returns the local name for [DataNeptuneOrderableDbInstance].
func (nodi *DataNeptuneOrderableDbInstance) LocalName() string {
	return nodi.Name
}

// Configuration returns the configuration (args) for [DataNeptuneOrderableDbInstance].
func (nodi *DataNeptuneOrderableDbInstance) Configuration() interface{} {
	return nodi.Args
}

// Attributes returns the attributes for [DataNeptuneOrderableDbInstance].
func (nodi *DataNeptuneOrderableDbInstance) Attributes() dataNeptuneOrderableDbInstanceAttributes {
	return dataNeptuneOrderableDbInstanceAttributes{ref: terra.ReferenceDataResource(nodi)}
}

// DataNeptuneOrderableDbInstanceArgs contains the configurations for aws_neptune_orderable_db_instance.
type DataNeptuneOrderableDbInstanceArgs struct {
	// Engine: string, optional
	Engine terra.StringValue `hcl:"engine,attr"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceClass: string, optional
	InstanceClass terra.StringValue `hcl:"instance_class,attr"`
	// LicenseModel: string, optional
	LicenseModel terra.StringValue `hcl:"license_model,attr"`
	// PreferredInstanceClasses: list of string, optional
	PreferredInstanceClasses terra.ListValue[terra.StringValue] `hcl:"preferred_instance_classes,attr"`
	// Vpc: bool, optional
	Vpc terra.BoolValue `hcl:"vpc,attr"`
}
type dataNeptuneOrderableDbInstanceAttributes struct {
	ref terra.Reference
}

// AvailabilityZones returns a reference to field availability_zones of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) AvailabilityZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nodi.ref.Append("availability_zones"))
}

// Engine returns a reference to field engine of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(nodi.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(nodi.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nodi.ref.Append("id"))
}

// InstanceClass returns a reference to field instance_class of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) InstanceClass() terra.StringValue {
	return terra.ReferenceAsString(nodi.ref.Append("instance_class"))
}

// LicenseModel returns a reference to field license_model of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceAsString(nodi.ref.Append("license_model"))
}

// MaxIopsPerDbInstance returns a reference to field max_iops_per_db_instance of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) MaxIopsPerDbInstance() terra.NumberValue {
	return terra.ReferenceAsNumber(nodi.ref.Append("max_iops_per_db_instance"))
}

// MaxIopsPerGib returns a reference to field max_iops_per_gib of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) MaxIopsPerGib() terra.NumberValue {
	return terra.ReferenceAsNumber(nodi.ref.Append("max_iops_per_gib"))
}

// MaxStorageSize returns a reference to field max_storage_size of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) MaxStorageSize() terra.NumberValue {
	return terra.ReferenceAsNumber(nodi.ref.Append("max_storage_size"))
}

// MinIopsPerDbInstance returns a reference to field min_iops_per_db_instance of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) MinIopsPerDbInstance() terra.NumberValue {
	return terra.ReferenceAsNumber(nodi.ref.Append("min_iops_per_db_instance"))
}

// MinIopsPerGib returns a reference to field min_iops_per_gib of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) MinIopsPerGib() terra.NumberValue {
	return terra.ReferenceAsNumber(nodi.ref.Append("min_iops_per_gib"))
}

// MinStorageSize returns a reference to field min_storage_size of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) MinStorageSize() terra.NumberValue {
	return terra.ReferenceAsNumber(nodi.ref.Append("min_storage_size"))
}

// MultiAzCapable returns a reference to field multi_az_capable of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) MultiAzCapable() terra.BoolValue {
	return terra.ReferenceAsBool(nodi.ref.Append("multi_az_capable"))
}

// PreferredInstanceClasses returns a reference to field preferred_instance_classes of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) PreferredInstanceClasses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nodi.ref.Append("preferred_instance_classes"))
}

// ReadReplicaCapable returns a reference to field read_replica_capable of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) ReadReplicaCapable() terra.BoolValue {
	return terra.ReferenceAsBool(nodi.ref.Append("read_replica_capable"))
}

// StorageType returns a reference to field storage_type of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(nodi.ref.Append("storage_type"))
}

// SupportsEnhancedMonitoring returns a reference to field supports_enhanced_monitoring of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) SupportsEnhancedMonitoring() terra.BoolValue {
	return terra.ReferenceAsBool(nodi.ref.Append("supports_enhanced_monitoring"))
}

// SupportsIamDatabaseAuthentication returns a reference to field supports_iam_database_authentication of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) SupportsIamDatabaseAuthentication() terra.BoolValue {
	return terra.ReferenceAsBool(nodi.ref.Append("supports_iam_database_authentication"))
}

// SupportsIops returns a reference to field supports_iops of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) SupportsIops() terra.BoolValue {
	return terra.ReferenceAsBool(nodi.ref.Append("supports_iops"))
}

// SupportsPerformanceInsights returns a reference to field supports_performance_insights of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) SupportsPerformanceInsights() terra.BoolValue {
	return terra.ReferenceAsBool(nodi.ref.Append("supports_performance_insights"))
}

// SupportsStorageEncryption returns a reference to field supports_storage_encryption of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) SupportsStorageEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(nodi.ref.Append("supports_storage_encryption"))
}

// Vpc returns a reference to field vpc of aws_neptune_orderable_db_instance.
func (nodi dataNeptuneOrderableDbInstanceAttributes) Vpc() terra.BoolValue {
	return terra.ReferenceAsBool(nodi.ref.Append("vpc"))
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	rdsclusterinstance "github.com/golingon/terraproviders/aws/5.10.0/rdsclusterinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRdsClusterInstance creates a new instance of [RdsClusterInstance].
func NewRdsClusterInstance(name string, args RdsClusterInstanceArgs) *RdsClusterInstance {
	return &RdsClusterInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RdsClusterInstance)(nil)

// RdsClusterInstance represents the Terraform resource aws_rds_cluster_instance.
type RdsClusterInstance struct {
	Name      string
	Args      RdsClusterInstanceArgs
	state     *rdsClusterInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RdsClusterInstance].
func (rci *RdsClusterInstance) Type() string {
	return "aws_rds_cluster_instance"
}

// LocalName returns the local name for [RdsClusterInstance].
func (rci *RdsClusterInstance) LocalName() string {
	return rci.Name
}

// Configuration returns the configuration (args) for [RdsClusterInstance].
func (rci *RdsClusterInstance) Configuration() interface{} {
	return rci.Args
}

// DependOn is used for other resources to depend on [RdsClusterInstance].
func (rci *RdsClusterInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(rci)
}

// Dependencies returns the list of resources [RdsClusterInstance] depends_on.
func (rci *RdsClusterInstance) Dependencies() terra.Dependencies {
	return rci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RdsClusterInstance].
func (rci *RdsClusterInstance) LifecycleManagement() *terra.Lifecycle {
	return rci.Lifecycle
}

// Attributes returns the attributes for [RdsClusterInstance].
func (rci *RdsClusterInstance) Attributes() rdsClusterInstanceAttributes {
	return rdsClusterInstanceAttributes{ref: terra.ReferenceResource(rci)}
}

// ImportState imports the given attribute values into [RdsClusterInstance]'s state.
func (rci *RdsClusterInstance) ImportState(av io.Reader) error {
	rci.state = &rdsClusterInstanceState{}
	if err := json.NewDecoder(av).Decode(rci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rci.Type(), rci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RdsClusterInstance] has state.
func (rci *RdsClusterInstance) State() (*rdsClusterInstanceState, bool) {
	return rci.state, rci.state != nil
}

// StateMust returns the state for [RdsClusterInstance]. Panics if the state is nil.
func (rci *RdsClusterInstance) StateMust() *rdsClusterInstanceState {
	if rci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rci.Type(), rci.LocalName()))
	}
	return rci.state
}

// RdsClusterInstanceArgs contains the configurations for aws_rds_cluster_instance.
type RdsClusterInstanceArgs struct {
	// ApplyImmediately: bool, optional
	ApplyImmediately terra.BoolValue `hcl:"apply_immediately,attr"`
	// AutoMinorVersionUpgrade: bool, optional
	AutoMinorVersionUpgrade terra.BoolValue `hcl:"auto_minor_version_upgrade,attr"`
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// CaCertIdentifier: string, optional
	CaCertIdentifier terra.StringValue `hcl:"ca_cert_identifier,attr"`
	// ClusterIdentifier: string, required
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr" validate:"required"`
	// CopyTagsToSnapshot: bool, optional
	CopyTagsToSnapshot terra.BoolValue `hcl:"copy_tags_to_snapshot,attr"`
	// DbParameterGroupName: string, optional
	DbParameterGroupName terra.StringValue `hcl:"db_parameter_group_name,attr"`
	// DbSubnetGroupName: string, optional
	DbSubnetGroupName terra.StringValue `hcl:"db_subnet_group_name,attr"`
	// Engine: string, required
	Engine terra.StringValue `hcl:"engine,attr" validate:"required"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Identifier: string, optional
	Identifier terra.StringValue `hcl:"identifier,attr"`
	// IdentifierPrefix: string, optional
	IdentifierPrefix terra.StringValue `hcl:"identifier_prefix,attr"`
	// InstanceClass: string, required
	InstanceClass terra.StringValue `hcl:"instance_class,attr" validate:"required"`
	// MonitoringInterval: number, optional
	MonitoringInterval terra.NumberValue `hcl:"monitoring_interval,attr"`
	// MonitoringRoleArn: string, optional
	MonitoringRoleArn terra.StringValue `hcl:"monitoring_role_arn,attr"`
	// PerformanceInsightsEnabled: bool, optional
	PerformanceInsightsEnabled terra.BoolValue `hcl:"performance_insights_enabled,attr"`
	// PerformanceInsightsKmsKeyId: string, optional
	PerformanceInsightsKmsKeyId terra.StringValue `hcl:"performance_insights_kms_key_id,attr"`
	// PerformanceInsightsRetentionPeriod: number, optional
	PerformanceInsightsRetentionPeriod terra.NumberValue `hcl:"performance_insights_retention_period,attr"`
	// PreferredBackupWindow: string, optional
	PreferredBackupWindow terra.StringValue `hcl:"preferred_backup_window,attr"`
	// PreferredMaintenanceWindow: string, optional
	PreferredMaintenanceWindow terra.StringValue `hcl:"preferred_maintenance_window,attr"`
	// PromotionTier: number, optional
	PromotionTier terra.NumberValue `hcl:"promotion_tier,attr"`
	// PubliclyAccessible: bool, optional
	PubliclyAccessible terra.BoolValue `hcl:"publicly_accessible,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *rdsclusterinstance.Timeouts `hcl:"timeouts,block"`
}
type rdsClusterInstanceAttributes struct {
	ref terra.Reference
}

// ApplyImmediately returns a reference to field apply_immediately of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceAsBool(rci.ref.Append("apply_immediately"))
}

// Arn returns a reference to field arn of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("arn"))
}

// AutoMinorVersionUpgrade returns a reference to field auto_minor_version_upgrade of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) AutoMinorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(rci.ref.Append("auto_minor_version_upgrade"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("availability_zone"))
}

// CaCertIdentifier returns a reference to field ca_cert_identifier of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) CaCertIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("ca_cert_identifier"))
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("cluster_identifier"))
}

// CopyTagsToSnapshot returns a reference to field copy_tags_to_snapshot of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) CopyTagsToSnapshot() terra.BoolValue {
	return terra.ReferenceAsBool(rci.ref.Append("copy_tags_to_snapshot"))
}

// DbParameterGroupName returns a reference to field db_parameter_group_name of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) DbParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("db_parameter_group_name"))
}

// DbSubnetGroupName returns a reference to field db_subnet_group_name of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) DbSubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("db_subnet_group_name"))
}

// DbiResourceId returns a reference to field dbi_resource_id of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) DbiResourceId() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("dbi_resource_id"))
}

// Endpoint returns a reference to field endpoint of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("endpoint"))
}

// Engine returns a reference to field engine of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("engine_version"))
}

// EngineVersionActual returns a reference to field engine_version_actual of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) EngineVersionActual() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("engine_version_actual"))
}

// Id returns a reference to field id of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("id"))
}

// Identifier returns a reference to field identifier of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) Identifier() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("identifier"))
}

// IdentifierPrefix returns a reference to field identifier_prefix of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) IdentifierPrefix() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("identifier_prefix"))
}

// InstanceClass returns a reference to field instance_class of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) InstanceClass() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("instance_class"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("kms_key_id"))
}

// MonitoringInterval returns a reference to field monitoring_interval of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) MonitoringInterval() terra.NumberValue {
	return terra.ReferenceAsNumber(rci.ref.Append("monitoring_interval"))
}

// MonitoringRoleArn returns a reference to field monitoring_role_arn of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) MonitoringRoleArn() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("monitoring_role_arn"))
}

// NetworkType returns a reference to field network_type of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) NetworkType() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("network_type"))
}

// PerformanceInsightsEnabled returns a reference to field performance_insights_enabled of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) PerformanceInsightsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rci.ref.Append("performance_insights_enabled"))
}

// PerformanceInsightsKmsKeyId returns a reference to field performance_insights_kms_key_id of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) PerformanceInsightsKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("performance_insights_kms_key_id"))
}

// PerformanceInsightsRetentionPeriod returns a reference to field performance_insights_retention_period of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) PerformanceInsightsRetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(rci.ref.Append("performance_insights_retention_period"))
}

// Port returns a reference to field port of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(rci.ref.Append("port"))
}

// PreferredBackupWindow returns a reference to field preferred_backup_window of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) PreferredBackupWindow() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("preferred_backup_window"))
}

// PreferredMaintenanceWindow returns a reference to field preferred_maintenance_window of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(rci.ref.Append("preferred_maintenance_window"))
}

// PromotionTier returns a reference to field promotion_tier of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) PromotionTier() terra.NumberValue {
	return terra.ReferenceAsNumber(rci.ref.Append("promotion_tier"))
}

// PubliclyAccessible returns a reference to field publicly_accessible of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) PubliclyAccessible() terra.BoolValue {
	return terra.ReferenceAsBool(rci.ref.Append("publicly_accessible"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(rci.ref.Append("storage_encrypted"))
}

// Tags returns a reference to field tags of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rci.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rci.ref.Append("tags_all"))
}

// Writer returns a reference to field writer of aws_rds_cluster_instance.
func (rci rdsClusterInstanceAttributes) Writer() terra.BoolValue {
	return terra.ReferenceAsBool(rci.ref.Append("writer"))
}

func (rci rdsClusterInstanceAttributes) Timeouts() rdsclusterinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[rdsclusterinstance.TimeoutsAttributes](rci.ref.Append("timeouts"))
}

type rdsClusterInstanceState struct {
	ApplyImmediately                   bool                              `json:"apply_immediately"`
	Arn                                string                            `json:"arn"`
	AutoMinorVersionUpgrade            bool                              `json:"auto_minor_version_upgrade"`
	AvailabilityZone                   string                            `json:"availability_zone"`
	CaCertIdentifier                   string                            `json:"ca_cert_identifier"`
	ClusterIdentifier                  string                            `json:"cluster_identifier"`
	CopyTagsToSnapshot                 bool                              `json:"copy_tags_to_snapshot"`
	DbParameterGroupName               string                            `json:"db_parameter_group_name"`
	DbSubnetGroupName                  string                            `json:"db_subnet_group_name"`
	DbiResourceId                      string                            `json:"dbi_resource_id"`
	Endpoint                           string                            `json:"endpoint"`
	Engine                             string                            `json:"engine"`
	EngineVersion                      string                            `json:"engine_version"`
	EngineVersionActual                string                            `json:"engine_version_actual"`
	Id                                 string                            `json:"id"`
	Identifier                         string                            `json:"identifier"`
	IdentifierPrefix                   string                            `json:"identifier_prefix"`
	InstanceClass                      string                            `json:"instance_class"`
	KmsKeyId                           string                            `json:"kms_key_id"`
	MonitoringInterval                 float64                           `json:"monitoring_interval"`
	MonitoringRoleArn                  string                            `json:"monitoring_role_arn"`
	NetworkType                        string                            `json:"network_type"`
	PerformanceInsightsEnabled         bool                              `json:"performance_insights_enabled"`
	PerformanceInsightsKmsKeyId        string                            `json:"performance_insights_kms_key_id"`
	PerformanceInsightsRetentionPeriod float64                           `json:"performance_insights_retention_period"`
	Port                               float64                           `json:"port"`
	PreferredBackupWindow              string                            `json:"preferred_backup_window"`
	PreferredMaintenanceWindow         string                            `json:"preferred_maintenance_window"`
	PromotionTier                      float64                           `json:"promotion_tier"`
	PubliclyAccessible                 bool                              `json:"publicly_accessible"`
	StorageEncrypted                   bool                              `json:"storage_encrypted"`
	Tags                               map[string]string                 `json:"tags"`
	TagsAll                            map[string]string                 `json:"tags_all"`
	Writer                             bool                              `json:"writer"`
	Timeouts                           *rdsclusterinstance.TimeoutsState `json:"timeouts"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	docdbclusterinstance "github.com/golingon/terraproviders/aws/5.9.0/docdbclusterinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDocdbClusterInstance creates a new instance of [DocdbClusterInstance].
func NewDocdbClusterInstance(name string, args DocdbClusterInstanceArgs) *DocdbClusterInstance {
	return &DocdbClusterInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DocdbClusterInstance)(nil)

// DocdbClusterInstance represents the Terraform resource aws_docdb_cluster_instance.
type DocdbClusterInstance struct {
	Name      string
	Args      DocdbClusterInstanceArgs
	state     *docdbClusterInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DocdbClusterInstance].
func (dci *DocdbClusterInstance) Type() string {
	return "aws_docdb_cluster_instance"
}

// LocalName returns the local name for [DocdbClusterInstance].
func (dci *DocdbClusterInstance) LocalName() string {
	return dci.Name
}

// Configuration returns the configuration (args) for [DocdbClusterInstance].
func (dci *DocdbClusterInstance) Configuration() interface{} {
	return dci.Args
}

// DependOn is used for other resources to depend on [DocdbClusterInstance].
func (dci *DocdbClusterInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(dci)
}

// Dependencies returns the list of resources [DocdbClusterInstance] depends_on.
func (dci *DocdbClusterInstance) Dependencies() terra.Dependencies {
	return dci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DocdbClusterInstance].
func (dci *DocdbClusterInstance) LifecycleManagement() *terra.Lifecycle {
	return dci.Lifecycle
}

// Attributes returns the attributes for [DocdbClusterInstance].
func (dci *DocdbClusterInstance) Attributes() docdbClusterInstanceAttributes {
	return docdbClusterInstanceAttributes{ref: terra.ReferenceResource(dci)}
}

// ImportState imports the given attribute values into [DocdbClusterInstance]'s state.
func (dci *DocdbClusterInstance) ImportState(av io.Reader) error {
	dci.state = &docdbClusterInstanceState{}
	if err := json.NewDecoder(av).Decode(dci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dci.Type(), dci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DocdbClusterInstance] has state.
func (dci *DocdbClusterInstance) State() (*docdbClusterInstanceState, bool) {
	return dci.state, dci.state != nil
}

// StateMust returns the state for [DocdbClusterInstance]. Panics if the state is nil.
func (dci *DocdbClusterInstance) StateMust() *docdbClusterInstanceState {
	if dci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dci.Type(), dci.LocalName()))
	}
	return dci.state
}

// DocdbClusterInstanceArgs contains the configurations for aws_docdb_cluster_instance.
type DocdbClusterInstanceArgs struct {
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
	// EnablePerformanceInsights: bool, optional
	EnablePerformanceInsights terra.BoolValue `hcl:"enable_performance_insights,attr"`
	// Engine: string, optional
	Engine terra.StringValue `hcl:"engine,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Identifier: string, optional
	Identifier terra.StringValue `hcl:"identifier,attr"`
	// IdentifierPrefix: string, optional
	IdentifierPrefix terra.StringValue `hcl:"identifier_prefix,attr"`
	// InstanceClass: string, required
	InstanceClass terra.StringValue `hcl:"instance_class,attr" validate:"required"`
	// PerformanceInsightsKmsKeyId: string, optional
	PerformanceInsightsKmsKeyId terra.StringValue `hcl:"performance_insights_kms_key_id,attr"`
	// PreferredMaintenanceWindow: string, optional
	PreferredMaintenanceWindow terra.StringValue `hcl:"preferred_maintenance_window,attr"`
	// PromotionTier: number, optional
	PromotionTier terra.NumberValue `hcl:"promotion_tier,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *docdbclusterinstance.Timeouts `hcl:"timeouts,block"`
}
type docdbClusterInstanceAttributes struct {
	ref terra.Reference
}

// ApplyImmediately returns a reference to field apply_immediately of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceAsBool(dci.ref.Append("apply_immediately"))
}

// Arn returns a reference to field arn of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("arn"))
}

// AutoMinorVersionUpgrade returns a reference to field auto_minor_version_upgrade of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) AutoMinorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(dci.ref.Append("auto_minor_version_upgrade"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("availability_zone"))
}

// CaCertIdentifier returns a reference to field ca_cert_identifier of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) CaCertIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("ca_cert_identifier"))
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("cluster_identifier"))
}

// DbSubnetGroupName returns a reference to field db_subnet_group_name of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) DbSubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("db_subnet_group_name"))
}

// DbiResourceId returns a reference to field dbi_resource_id of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) DbiResourceId() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("dbi_resource_id"))
}

// EnablePerformanceInsights returns a reference to field enable_performance_insights of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) EnablePerformanceInsights() terra.BoolValue {
	return terra.ReferenceAsBool(dci.ref.Append("enable_performance_insights"))
}

// Endpoint returns a reference to field endpoint of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("endpoint"))
}

// Engine returns a reference to field engine of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("id"))
}

// Identifier returns a reference to field identifier of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) Identifier() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("identifier"))
}

// IdentifierPrefix returns a reference to field identifier_prefix of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) IdentifierPrefix() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("identifier_prefix"))
}

// InstanceClass returns a reference to field instance_class of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) InstanceClass() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("instance_class"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("kms_key_id"))
}

// PerformanceInsightsKmsKeyId returns a reference to field performance_insights_kms_key_id of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) PerformanceInsightsKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("performance_insights_kms_key_id"))
}

// Port returns a reference to field port of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(dci.ref.Append("port"))
}

// PreferredBackupWindow returns a reference to field preferred_backup_window of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) PreferredBackupWindow() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("preferred_backup_window"))
}

// PreferredMaintenanceWindow returns a reference to field preferred_maintenance_window of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(dci.ref.Append("preferred_maintenance_window"))
}

// PromotionTier returns a reference to field promotion_tier of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) PromotionTier() terra.NumberValue {
	return terra.ReferenceAsNumber(dci.ref.Append("promotion_tier"))
}

// PubliclyAccessible returns a reference to field publicly_accessible of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) PubliclyAccessible() terra.BoolValue {
	return terra.ReferenceAsBool(dci.ref.Append("publicly_accessible"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(dci.ref.Append("storage_encrypted"))
}

// Tags returns a reference to field tags of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dci.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dci.ref.Append("tags_all"))
}

// Writer returns a reference to field writer of aws_docdb_cluster_instance.
func (dci docdbClusterInstanceAttributes) Writer() terra.BoolValue {
	return terra.ReferenceAsBool(dci.ref.Append("writer"))
}

func (dci docdbClusterInstanceAttributes) Timeouts() docdbclusterinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[docdbclusterinstance.TimeoutsAttributes](dci.ref.Append("timeouts"))
}

type docdbClusterInstanceState struct {
	ApplyImmediately            bool                                `json:"apply_immediately"`
	Arn                         string                              `json:"arn"`
	AutoMinorVersionUpgrade     bool                                `json:"auto_minor_version_upgrade"`
	AvailabilityZone            string                              `json:"availability_zone"`
	CaCertIdentifier            string                              `json:"ca_cert_identifier"`
	ClusterIdentifier           string                              `json:"cluster_identifier"`
	DbSubnetGroupName           string                              `json:"db_subnet_group_name"`
	DbiResourceId               string                              `json:"dbi_resource_id"`
	EnablePerformanceInsights   bool                                `json:"enable_performance_insights"`
	Endpoint                    string                              `json:"endpoint"`
	Engine                      string                              `json:"engine"`
	EngineVersion               string                              `json:"engine_version"`
	Id                          string                              `json:"id"`
	Identifier                  string                              `json:"identifier"`
	IdentifierPrefix            string                              `json:"identifier_prefix"`
	InstanceClass               string                              `json:"instance_class"`
	KmsKeyId                    string                              `json:"kms_key_id"`
	PerformanceInsightsKmsKeyId string                              `json:"performance_insights_kms_key_id"`
	Port                        float64                             `json:"port"`
	PreferredBackupWindow       string                              `json:"preferred_backup_window"`
	PreferredMaintenanceWindow  string                              `json:"preferred_maintenance_window"`
	PromotionTier               float64                             `json:"promotion_tier"`
	PubliclyAccessible          bool                                `json:"publicly_accessible"`
	StorageEncrypted            bool                                `json:"storage_encrypted"`
	Tags                        map[string]string                   `json:"tags"`
	TagsAll                     map[string]string                   `json:"tags_all"`
	Writer                      bool                                `json:"writer"`
	Timeouts                    *docdbclusterinstance.TimeoutsState `json:"timeouts"`
}
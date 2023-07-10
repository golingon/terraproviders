// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	neptuneclusterinstance "github.com/golingon/terraproviders/aws/5.7.0/neptuneclusterinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNeptuneClusterInstance creates a new instance of [NeptuneClusterInstance].
func NewNeptuneClusterInstance(name string, args NeptuneClusterInstanceArgs) *NeptuneClusterInstance {
	return &NeptuneClusterInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NeptuneClusterInstance)(nil)

// NeptuneClusterInstance represents the Terraform resource aws_neptune_cluster_instance.
type NeptuneClusterInstance struct {
	Name      string
	Args      NeptuneClusterInstanceArgs
	state     *neptuneClusterInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NeptuneClusterInstance].
func (nci *NeptuneClusterInstance) Type() string {
	return "aws_neptune_cluster_instance"
}

// LocalName returns the local name for [NeptuneClusterInstance].
func (nci *NeptuneClusterInstance) LocalName() string {
	return nci.Name
}

// Configuration returns the configuration (args) for [NeptuneClusterInstance].
func (nci *NeptuneClusterInstance) Configuration() interface{} {
	return nci.Args
}

// DependOn is used for other resources to depend on [NeptuneClusterInstance].
func (nci *NeptuneClusterInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(nci)
}

// Dependencies returns the list of resources [NeptuneClusterInstance] depends_on.
func (nci *NeptuneClusterInstance) Dependencies() terra.Dependencies {
	return nci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NeptuneClusterInstance].
func (nci *NeptuneClusterInstance) LifecycleManagement() *terra.Lifecycle {
	return nci.Lifecycle
}

// Attributes returns the attributes for [NeptuneClusterInstance].
func (nci *NeptuneClusterInstance) Attributes() neptuneClusterInstanceAttributes {
	return neptuneClusterInstanceAttributes{ref: terra.ReferenceResource(nci)}
}

// ImportState imports the given attribute values into [NeptuneClusterInstance]'s state.
func (nci *NeptuneClusterInstance) ImportState(av io.Reader) error {
	nci.state = &neptuneClusterInstanceState{}
	if err := json.NewDecoder(av).Decode(nci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nci.Type(), nci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NeptuneClusterInstance] has state.
func (nci *NeptuneClusterInstance) State() (*neptuneClusterInstanceState, bool) {
	return nci.state, nci.state != nil
}

// StateMust returns the state for [NeptuneClusterInstance]. Panics if the state is nil.
func (nci *NeptuneClusterInstance) StateMust() *neptuneClusterInstanceState {
	if nci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nci.Type(), nci.LocalName()))
	}
	return nci.state
}

// NeptuneClusterInstanceArgs contains the configurations for aws_neptune_cluster_instance.
type NeptuneClusterInstanceArgs struct {
	// ApplyImmediately: bool, optional
	ApplyImmediately terra.BoolValue `hcl:"apply_immediately,attr"`
	// AutoMinorVersionUpgrade: bool, optional
	AutoMinorVersionUpgrade terra.BoolValue `hcl:"auto_minor_version_upgrade,attr"`
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// ClusterIdentifier: string, required
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr" validate:"required"`
	// Engine: string, optional
	Engine terra.StringValue `hcl:"engine,attr"`
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
	// NeptuneParameterGroupName: string, optional
	NeptuneParameterGroupName terra.StringValue `hcl:"neptune_parameter_group_name,attr"`
	// NeptuneSubnetGroupName: string, optional
	NeptuneSubnetGroupName terra.StringValue `hcl:"neptune_subnet_group_name,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
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
	Timeouts *neptuneclusterinstance.Timeouts `hcl:"timeouts,block"`
}
type neptuneClusterInstanceAttributes struct {
	ref terra.Reference
}

// Address returns a reference to field address of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("address"))
}

// ApplyImmediately returns a reference to field apply_immediately of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceAsBool(nci.ref.Append("apply_immediately"))
}

// Arn returns a reference to field arn of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("arn"))
}

// AutoMinorVersionUpgrade returns a reference to field auto_minor_version_upgrade of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) AutoMinorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(nci.ref.Append("auto_minor_version_upgrade"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("availability_zone"))
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("cluster_identifier"))
}

// DbiResourceId returns a reference to field dbi_resource_id of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) DbiResourceId() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("dbi_resource_id"))
}

// Endpoint returns a reference to field endpoint of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("endpoint"))
}

// Engine returns a reference to field engine of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("id"))
}

// Identifier returns a reference to field identifier of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) Identifier() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("identifier"))
}

// IdentifierPrefix returns a reference to field identifier_prefix of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) IdentifierPrefix() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("identifier_prefix"))
}

// InstanceClass returns a reference to field instance_class of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) InstanceClass() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("instance_class"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("kms_key_arn"))
}

// NeptuneParameterGroupName returns a reference to field neptune_parameter_group_name of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) NeptuneParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("neptune_parameter_group_name"))
}

// NeptuneSubnetGroupName returns a reference to field neptune_subnet_group_name of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) NeptuneSubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("neptune_subnet_group_name"))
}

// Port returns a reference to field port of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(nci.ref.Append("port"))
}

// PreferredBackupWindow returns a reference to field preferred_backup_window of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) PreferredBackupWindow() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("preferred_backup_window"))
}

// PreferredMaintenanceWindow returns a reference to field preferred_maintenance_window of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(nci.ref.Append("preferred_maintenance_window"))
}

// PromotionTier returns a reference to field promotion_tier of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) PromotionTier() terra.NumberValue {
	return terra.ReferenceAsNumber(nci.ref.Append("promotion_tier"))
}

// PubliclyAccessible returns a reference to field publicly_accessible of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) PubliclyAccessible() terra.BoolValue {
	return terra.ReferenceAsBool(nci.ref.Append("publicly_accessible"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(nci.ref.Append("storage_encrypted"))
}

// Tags returns a reference to field tags of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nci.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nci.ref.Append("tags_all"))
}

// Writer returns a reference to field writer of aws_neptune_cluster_instance.
func (nci neptuneClusterInstanceAttributes) Writer() terra.BoolValue {
	return terra.ReferenceAsBool(nci.ref.Append("writer"))
}

func (nci neptuneClusterInstanceAttributes) Timeouts() neptuneclusterinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[neptuneclusterinstance.TimeoutsAttributes](nci.ref.Append("timeouts"))
}

type neptuneClusterInstanceState struct {
	Address                    string                                `json:"address"`
	ApplyImmediately           bool                                  `json:"apply_immediately"`
	Arn                        string                                `json:"arn"`
	AutoMinorVersionUpgrade    bool                                  `json:"auto_minor_version_upgrade"`
	AvailabilityZone           string                                `json:"availability_zone"`
	ClusterIdentifier          string                                `json:"cluster_identifier"`
	DbiResourceId              string                                `json:"dbi_resource_id"`
	Endpoint                   string                                `json:"endpoint"`
	Engine                     string                                `json:"engine"`
	EngineVersion              string                                `json:"engine_version"`
	Id                         string                                `json:"id"`
	Identifier                 string                                `json:"identifier"`
	IdentifierPrefix           string                                `json:"identifier_prefix"`
	InstanceClass              string                                `json:"instance_class"`
	KmsKeyArn                  string                                `json:"kms_key_arn"`
	NeptuneParameterGroupName  string                                `json:"neptune_parameter_group_name"`
	NeptuneSubnetGroupName     string                                `json:"neptune_subnet_group_name"`
	Port                       float64                               `json:"port"`
	PreferredBackupWindow      string                                `json:"preferred_backup_window"`
	PreferredMaintenanceWindow string                                `json:"preferred_maintenance_window"`
	PromotionTier              float64                               `json:"promotion_tier"`
	PubliclyAccessible         bool                                  `json:"publicly_accessible"`
	StorageEncrypted           bool                                  `json:"storage_encrypted"`
	Tags                       map[string]string                     `json:"tags"`
	TagsAll                    map[string]string                     `json:"tags_all"`
	Writer                     bool                                  `json:"writer"`
	Timeouts                   *neptuneclusterinstance.TimeoutsState `json:"timeouts"`
}

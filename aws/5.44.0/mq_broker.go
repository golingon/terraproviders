// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	mqbroker "github.com/golingon/terraproviders/aws/5.44.0/mqbroker"
	"io"
)

// NewMqBroker creates a new instance of [MqBroker].
func NewMqBroker(name string, args MqBrokerArgs) *MqBroker {
	return &MqBroker{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MqBroker)(nil)

// MqBroker represents the Terraform resource aws_mq_broker.
type MqBroker struct {
	Name      string
	Args      MqBrokerArgs
	state     *mqBrokerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MqBroker].
func (mb *MqBroker) Type() string {
	return "aws_mq_broker"
}

// LocalName returns the local name for [MqBroker].
func (mb *MqBroker) LocalName() string {
	return mb.Name
}

// Configuration returns the configuration (args) for [MqBroker].
func (mb *MqBroker) Configuration() interface{} {
	return mb.Args
}

// DependOn is used for other resources to depend on [MqBroker].
func (mb *MqBroker) DependOn() terra.Reference {
	return terra.ReferenceResource(mb)
}

// Dependencies returns the list of resources [MqBroker] depends_on.
func (mb *MqBroker) Dependencies() terra.Dependencies {
	return mb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MqBroker].
func (mb *MqBroker) LifecycleManagement() *terra.Lifecycle {
	return mb.Lifecycle
}

// Attributes returns the attributes for [MqBroker].
func (mb *MqBroker) Attributes() mqBrokerAttributes {
	return mqBrokerAttributes{ref: terra.ReferenceResource(mb)}
}

// ImportState imports the given attribute values into [MqBroker]'s state.
func (mb *MqBroker) ImportState(av io.Reader) error {
	mb.state = &mqBrokerState{}
	if err := json.NewDecoder(av).Decode(mb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mb.Type(), mb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MqBroker] has state.
func (mb *MqBroker) State() (*mqBrokerState, bool) {
	return mb.state, mb.state != nil
}

// StateMust returns the state for [MqBroker]. Panics if the state is nil.
func (mb *MqBroker) StateMust() *mqBrokerState {
	if mb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mb.Type(), mb.LocalName()))
	}
	return mb.state
}

// MqBrokerArgs contains the configurations for aws_mq_broker.
type MqBrokerArgs struct {
	// ApplyImmediately: bool, optional
	ApplyImmediately terra.BoolValue `hcl:"apply_immediately,attr"`
	// AuthenticationStrategy: string, optional
	AuthenticationStrategy terra.StringValue `hcl:"authentication_strategy,attr"`
	// AutoMinorVersionUpgrade: bool, optional
	AutoMinorVersionUpgrade terra.BoolValue `hcl:"auto_minor_version_upgrade,attr"`
	// BrokerName: string, required
	BrokerName terra.StringValue `hcl:"broker_name,attr" validate:"required"`
	// DataReplicationMode: string, optional
	DataReplicationMode terra.StringValue `hcl:"data_replication_mode,attr"`
	// DataReplicationPrimaryBrokerArn: string, optional
	DataReplicationPrimaryBrokerArn terra.StringValue `hcl:"data_replication_primary_broker_arn,attr"`
	// DeploymentMode: string, optional
	DeploymentMode terra.StringValue `hcl:"deployment_mode,attr"`
	// EngineType: string, required
	EngineType terra.StringValue `hcl:"engine_type,attr" validate:"required"`
	// EngineVersion: string, required
	EngineVersion terra.StringValue `hcl:"engine_version,attr" validate:"required"`
	// HostInstanceType: string, required
	HostInstanceType terra.StringValue `hcl:"host_instance_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PubliclyAccessible: bool, optional
	PubliclyAccessible terra.BoolValue `hcl:"publicly_accessible,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// StorageType: string, optional
	StorageType terra.StringValue `hcl:"storage_type,attr"`
	// SubnetIds: set of string, optional
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Instances: min=0
	Instances []mqbroker.Instances `hcl:"instances,block" validate:"min=0"`
	// Configuration: optional
	Configuration *mqbroker.Configuration `hcl:"configuration,block"`
	// EncryptionOptions: optional
	EncryptionOptions *mqbroker.EncryptionOptions `hcl:"encryption_options,block"`
	// LdapServerMetadata: optional
	LdapServerMetadata *mqbroker.LdapServerMetadata `hcl:"ldap_server_metadata,block"`
	// Logs: optional
	Logs *mqbroker.Logs `hcl:"logs,block"`
	// MaintenanceWindowStartTime: optional
	MaintenanceWindowStartTime *mqbroker.MaintenanceWindowStartTime `hcl:"maintenance_window_start_time,block"`
	// Timeouts: optional
	Timeouts *mqbroker.Timeouts `hcl:"timeouts,block"`
	// User: min=1
	User []mqbroker.User `hcl:"user,block" validate:"min=1"`
}
type mqBrokerAttributes struct {
	ref terra.Reference
}

// ApplyImmediately returns a reference to field apply_immediately of aws_mq_broker.
func (mb mqBrokerAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceAsBool(mb.ref.Append("apply_immediately"))
}

// Arn returns a reference to field arn of aws_mq_broker.
func (mb mqBrokerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("arn"))
}

// AuthenticationStrategy returns a reference to field authentication_strategy of aws_mq_broker.
func (mb mqBrokerAttributes) AuthenticationStrategy() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("authentication_strategy"))
}

// AutoMinorVersionUpgrade returns a reference to field auto_minor_version_upgrade of aws_mq_broker.
func (mb mqBrokerAttributes) AutoMinorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(mb.ref.Append("auto_minor_version_upgrade"))
}

// BrokerName returns a reference to field broker_name of aws_mq_broker.
func (mb mqBrokerAttributes) BrokerName() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("broker_name"))
}

// DataReplicationMode returns a reference to field data_replication_mode of aws_mq_broker.
func (mb mqBrokerAttributes) DataReplicationMode() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("data_replication_mode"))
}

// DataReplicationPrimaryBrokerArn returns a reference to field data_replication_primary_broker_arn of aws_mq_broker.
func (mb mqBrokerAttributes) DataReplicationPrimaryBrokerArn() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("data_replication_primary_broker_arn"))
}

// DeploymentMode returns a reference to field deployment_mode of aws_mq_broker.
func (mb mqBrokerAttributes) DeploymentMode() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("deployment_mode"))
}

// EngineType returns a reference to field engine_type of aws_mq_broker.
func (mb mqBrokerAttributes) EngineType() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("engine_type"))
}

// EngineVersion returns a reference to field engine_version of aws_mq_broker.
func (mb mqBrokerAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("engine_version"))
}

// HostInstanceType returns a reference to field host_instance_type of aws_mq_broker.
func (mb mqBrokerAttributes) HostInstanceType() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("host_instance_type"))
}

// Id returns a reference to field id of aws_mq_broker.
func (mb mqBrokerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("id"))
}

// PendingDataReplicationMode returns a reference to field pending_data_replication_mode of aws_mq_broker.
func (mb mqBrokerAttributes) PendingDataReplicationMode() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("pending_data_replication_mode"))
}

// PubliclyAccessible returns a reference to field publicly_accessible of aws_mq_broker.
func (mb mqBrokerAttributes) PubliclyAccessible() terra.BoolValue {
	return terra.ReferenceAsBool(mb.ref.Append("publicly_accessible"))
}

// SecurityGroups returns a reference to field security_groups of aws_mq_broker.
func (mb mqBrokerAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mb.ref.Append("security_groups"))
}

// StorageType returns a reference to field storage_type of aws_mq_broker.
func (mb mqBrokerAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(mb.ref.Append("storage_type"))
}

// SubnetIds returns a reference to field subnet_ids of aws_mq_broker.
func (mb mqBrokerAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mb.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_mq_broker.
func (mb mqBrokerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mb.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_mq_broker.
func (mb mqBrokerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mb.ref.Append("tags_all"))
}

func (mb mqBrokerAttributes) Instances() terra.ListValue[mqbroker.InstancesAttributes] {
	return terra.ReferenceAsList[mqbroker.InstancesAttributes](mb.ref.Append("instances"))
}

func (mb mqBrokerAttributes) Configuration() terra.ListValue[mqbroker.ConfigurationAttributes] {
	return terra.ReferenceAsList[mqbroker.ConfigurationAttributes](mb.ref.Append("configuration"))
}

func (mb mqBrokerAttributes) EncryptionOptions() terra.ListValue[mqbroker.EncryptionOptionsAttributes] {
	return terra.ReferenceAsList[mqbroker.EncryptionOptionsAttributes](mb.ref.Append("encryption_options"))
}

func (mb mqBrokerAttributes) LdapServerMetadata() terra.ListValue[mqbroker.LdapServerMetadataAttributes] {
	return terra.ReferenceAsList[mqbroker.LdapServerMetadataAttributes](mb.ref.Append("ldap_server_metadata"))
}

func (mb mqBrokerAttributes) Logs() terra.ListValue[mqbroker.LogsAttributes] {
	return terra.ReferenceAsList[mqbroker.LogsAttributes](mb.ref.Append("logs"))
}

func (mb mqBrokerAttributes) MaintenanceWindowStartTime() terra.ListValue[mqbroker.MaintenanceWindowStartTimeAttributes] {
	return terra.ReferenceAsList[mqbroker.MaintenanceWindowStartTimeAttributes](mb.ref.Append("maintenance_window_start_time"))
}

func (mb mqBrokerAttributes) Timeouts() mqbroker.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mqbroker.TimeoutsAttributes](mb.ref.Append("timeouts"))
}

func (mb mqBrokerAttributes) User() terra.SetValue[mqbroker.UserAttributes] {
	return terra.ReferenceAsSet[mqbroker.UserAttributes](mb.ref.Append("user"))
}

type mqBrokerState struct {
	ApplyImmediately                bool                                       `json:"apply_immediately"`
	Arn                             string                                     `json:"arn"`
	AuthenticationStrategy          string                                     `json:"authentication_strategy"`
	AutoMinorVersionUpgrade         bool                                       `json:"auto_minor_version_upgrade"`
	BrokerName                      string                                     `json:"broker_name"`
	DataReplicationMode             string                                     `json:"data_replication_mode"`
	DataReplicationPrimaryBrokerArn string                                     `json:"data_replication_primary_broker_arn"`
	DeploymentMode                  string                                     `json:"deployment_mode"`
	EngineType                      string                                     `json:"engine_type"`
	EngineVersion                   string                                     `json:"engine_version"`
	HostInstanceType                string                                     `json:"host_instance_type"`
	Id                              string                                     `json:"id"`
	PendingDataReplicationMode      string                                     `json:"pending_data_replication_mode"`
	PubliclyAccessible              bool                                       `json:"publicly_accessible"`
	SecurityGroups                  []string                                   `json:"security_groups"`
	StorageType                     string                                     `json:"storage_type"`
	SubnetIds                       []string                                   `json:"subnet_ids"`
	Tags                            map[string]string                          `json:"tags"`
	TagsAll                         map[string]string                          `json:"tags_all"`
	Instances                       []mqbroker.InstancesState                  `json:"instances"`
	Configuration                   []mqbroker.ConfigurationState              `json:"configuration"`
	EncryptionOptions               []mqbroker.EncryptionOptionsState          `json:"encryption_options"`
	LdapServerMetadata              []mqbroker.LdapServerMetadataState         `json:"ldap_server_metadata"`
	Logs                            []mqbroker.LogsState                       `json:"logs"`
	MaintenanceWindowStartTime      []mqbroker.MaintenanceWindowStartTimeState `json:"maintenance_window_start_time"`
	Timeouts                        *mqbroker.TimeoutsState                    `json:"timeouts"`
	User                            []mqbroker.UserState                       `json:"user"`
}

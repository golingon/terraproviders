// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	finspacekxcluster "github.com/golingon/terraproviders/aws/5.44.0/finspacekxcluster"
	"io"
)

// NewFinspaceKxCluster creates a new instance of [FinspaceKxCluster].
func NewFinspaceKxCluster(name string, args FinspaceKxClusterArgs) *FinspaceKxCluster {
	return &FinspaceKxCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FinspaceKxCluster)(nil)

// FinspaceKxCluster represents the Terraform resource aws_finspace_kx_cluster.
type FinspaceKxCluster struct {
	Name      string
	Args      FinspaceKxClusterArgs
	state     *finspaceKxClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FinspaceKxCluster].
func (fkc *FinspaceKxCluster) Type() string {
	return "aws_finspace_kx_cluster"
}

// LocalName returns the local name for [FinspaceKxCluster].
func (fkc *FinspaceKxCluster) LocalName() string {
	return fkc.Name
}

// Configuration returns the configuration (args) for [FinspaceKxCluster].
func (fkc *FinspaceKxCluster) Configuration() interface{} {
	return fkc.Args
}

// DependOn is used for other resources to depend on [FinspaceKxCluster].
func (fkc *FinspaceKxCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(fkc)
}

// Dependencies returns the list of resources [FinspaceKxCluster] depends_on.
func (fkc *FinspaceKxCluster) Dependencies() terra.Dependencies {
	return fkc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FinspaceKxCluster].
func (fkc *FinspaceKxCluster) LifecycleManagement() *terra.Lifecycle {
	return fkc.Lifecycle
}

// Attributes returns the attributes for [FinspaceKxCluster].
func (fkc *FinspaceKxCluster) Attributes() finspaceKxClusterAttributes {
	return finspaceKxClusterAttributes{ref: terra.ReferenceResource(fkc)}
}

// ImportState imports the given attribute values into [FinspaceKxCluster]'s state.
func (fkc *FinspaceKxCluster) ImportState(av io.Reader) error {
	fkc.state = &finspaceKxClusterState{}
	if err := json.NewDecoder(av).Decode(fkc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fkc.Type(), fkc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FinspaceKxCluster] has state.
func (fkc *FinspaceKxCluster) State() (*finspaceKxClusterState, bool) {
	return fkc.state, fkc.state != nil
}

// StateMust returns the state for [FinspaceKxCluster]. Panics if the state is nil.
func (fkc *FinspaceKxCluster) StateMust() *finspaceKxClusterState {
	if fkc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fkc.Type(), fkc.LocalName()))
	}
	return fkc.state
}

// FinspaceKxClusterArgs contains the configurations for aws_finspace_kx_cluster.
type FinspaceKxClusterArgs struct {
	// AvailabilityZoneId: string, optional
	AvailabilityZoneId terra.StringValue `hcl:"availability_zone_id,attr"`
	// AzMode: string, required
	AzMode terra.StringValue `hcl:"az_mode,attr" validate:"required"`
	// CommandLineArguments: map of string, optional
	CommandLineArguments terra.MapValue[terra.StringValue] `hcl:"command_line_arguments,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnvironmentId: string, required
	EnvironmentId terra.StringValue `hcl:"environment_id,attr" validate:"required"`
	// ExecutionRole: string, optional
	ExecutionRole terra.StringValue `hcl:"execution_role,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InitializationScript: string, optional
	InitializationScript terra.StringValue `hcl:"initialization_script,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ReleaseLabel: string, required
	ReleaseLabel terra.StringValue `hcl:"release_label,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// AutoScalingConfiguration: optional
	AutoScalingConfiguration *finspacekxcluster.AutoScalingConfiguration `hcl:"auto_scaling_configuration,block"`
	// CacheStorageConfigurations: min=0
	CacheStorageConfigurations []finspacekxcluster.CacheStorageConfigurations `hcl:"cache_storage_configurations,block" validate:"min=0"`
	// CapacityConfiguration: optional
	CapacityConfiguration *finspacekxcluster.CapacityConfiguration `hcl:"capacity_configuration,block"`
	// Code: optional
	Code *finspacekxcluster.Code `hcl:"code,block"`
	// Database: min=0
	Database []finspacekxcluster.Database `hcl:"database,block" validate:"min=0"`
	// SavedownStorageConfiguration: optional
	SavedownStorageConfiguration *finspacekxcluster.SavedownStorageConfiguration `hcl:"savedown_storage_configuration,block"`
	// ScalingGroupConfiguration: optional
	ScalingGroupConfiguration *finspacekxcluster.ScalingGroupConfiguration `hcl:"scaling_group_configuration,block"`
	// TickerplantLogConfiguration: min=0
	TickerplantLogConfiguration []finspacekxcluster.TickerplantLogConfiguration `hcl:"tickerplant_log_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *finspacekxcluster.Timeouts `hcl:"timeouts,block"`
	// VpcConfiguration: required
	VpcConfiguration *finspacekxcluster.VpcConfiguration `hcl:"vpc_configuration,block" validate:"required"`
}
type finspaceKxClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("arn"))
}

// AvailabilityZoneId returns a reference to field availability_zone_id of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("availability_zone_id"))
}

// AzMode returns a reference to field az_mode of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) AzMode() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("az_mode"))
}

// CommandLineArguments returns a reference to field command_line_arguments of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) CommandLineArguments() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fkc.ref.Append("command_line_arguments"))
}

// CreatedTimestamp returns a reference to field created_timestamp of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) CreatedTimestamp() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("created_timestamp"))
}

// Description returns a reference to field description of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("description"))
}

// EnvironmentId returns a reference to field environment_id of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) EnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("environment_id"))
}

// ExecutionRole returns a reference to field execution_role of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) ExecutionRole() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("execution_role"))
}

// Id returns a reference to field id of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("id"))
}

// InitializationScript returns a reference to field initialization_script of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) InitializationScript() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("initialization_script"))
}

// LastModifiedTimestamp returns a reference to field last_modified_timestamp of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) LastModifiedTimestamp() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("last_modified_timestamp"))
}

// Name returns a reference to field name of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("name"))
}

// ReleaseLabel returns a reference to field release_label of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) ReleaseLabel() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("release_label"))
}

// Status returns a reference to field status of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("status"))
}

// StatusReason returns a reference to field status_reason of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) StatusReason() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("status_reason"))
}

// Tags returns a reference to field tags of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fkc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fkc.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_finspace_kx_cluster.
func (fkc finspaceKxClusterAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(fkc.ref.Append("type"))
}

func (fkc finspaceKxClusterAttributes) AutoScalingConfiguration() terra.ListValue[finspacekxcluster.AutoScalingConfigurationAttributes] {
	return terra.ReferenceAsList[finspacekxcluster.AutoScalingConfigurationAttributes](fkc.ref.Append("auto_scaling_configuration"))
}

func (fkc finspaceKxClusterAttributes) CacheStorageConfigurations() terra.ListValue[finspacekxcluster.CacheStorageConfigurationsAttributes] {
	return terra.ReferenceAsList[finspacekxcluster.CacheStorageConfigurationsAttributes](fkc.ref.Append("cache_storage_configurations"))
}

func (fkc finspaceKxClusterAttributes) CapacityConfiguration() terra.ListValue[finspacekxcluster.CapacityConfigurationAttributes] {
	return terra.ReferenceAsList[finspacekxcluster.CapacityConfigurationAttributes](fkc.ref.Append("capacity_configuration"))
}

func (fkc finspaceKxClusterAttributes) Code() terra.ListValue[finspacekxcluster.CodeAttributes] {
	return terra.ReferenceAsList[finspacekxcluster.CodeAttributes](fkc.ref.Append("code"))
}

func (fkc finspaceKxClusterAttributes) Database() terra.ListValue[finspacekxcluster.DatabaseAttributes] {
	return terra.ReferenceAsList[finspacekxcluster.DatabaseAttributes](fkc.ref.Append("database"))
}

func (fkc finspaceKxClusterAttributes) SavedownStorageConfiguration() terra.ListValue[finspacekxcluster.SavedownStorageConfigurationAttributes] {
	return terra.ReferenceAsList[finspacekxcluster.SavedownStorageConfigurationAttributes](fkc.ref.Append("savedown_storage_configuration"))
}

func (fkc finspaceKxClusterAttributes) ScalingGroupConfiguration() terra.ListValue[finspacekxcluster.ScalingGroupConfigurationAttributes] {
	return terra.ReferenceAsList[finspacekxcluster.ScalingGroupConfigurationAttributes](fkc.ref.Append("scaling_group_configuration"))
}

func (fkc finspaceKxClusterAttributes) TickerplantLogConfiguration() terra.ListValue[finspacekxcluster.TickerplantLogConfigurationAttributes] {
	return terra.ReferenceAsList[finspacekxcluster.TickerplantLogConfigurationAttributes](fkc.ref.Append("tickerplant_log_configuration"))
}

func (fkc finspaceKxClusterAttributes) Timeouts() finspacekxcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[finspacekxcluster.TimeoutsAttributes](fkc.ref.Append("timeouts"))
}

func (fkc finspaceKxClusterAttributes) VpcConfiguration() terra.ListValue[finspacekxcluster.VpcConfigurationAttributes] {
	return terra.ReferenceAsList[finspacekxcluster.VpcConfigurationAttributes](fkc.ref.Append("vpc_configuration"))
}

type finspaceKxClusterState struct {
	Arn                          string                                                `json:"arn"`
	AvailabilityZoneId           string                                                `json:"availability_zone_id"`
	AzMode                       string                                                `json:"az_mode"`
	CommandLineArguments         map[string]string                                     `json:"command_line_arguments"`
	CreatedTimestamp             string                                                `json:"created_timestamp"`
	Description                  string                                                `json:"description"`
	EnvironmentId                string                                                `json:"environment_id"`
	ExecutionRole                string                                                `json:"execution_role"`
	Id                           string                                                `json:"id"`
	InitializationScript         string                                                `json:"initialization_script"`
	LastModifiedTimestamp        string                                                `json:"last_modified_timestamp"`
	Name                         string                                                `json:"name"`
	ReleaseLabel                 string                                                `json:"release_label"`
	Status                       string                                                `json:"status"`
	StatusReason                 string                                                `json:"status_reason"`
	Tags                         map[string]string                                     `json:"tags"`
	TagsAll                      map[string]string                                     `json:"tags_all"`
	Type                         string                                                `json:"type"`
	AutoScalingConfiguration     []finspacekxcluster.AutoScalingConfigurationState     `json:"auto_scaling_configuration"`
	CacheStorageConfigurations   []finspacekxcluster.CacheStorageConfigurationsState   `json:"cache_storage_configurations"`
	CapacityConfiguration        []finspacekxcluster.CapacityConfigurationState        `json:"capacity_configuration"`
	Code                         []finspacekxcluster.CodeState                         `json:"code"`
	Database                     []finspacekxcluster.DatabaseState                     `json:"database"`
	SavedownStorageConfiguration []finspacekxcluster.SavedownStorageConfigurationState `json:"savedown_storage_configuration"`
	ScalingGroupConfiguration    []finspacekxcluster.ScalingGroupConfigurationState    `json:"scaling_group_configuration"`
	TickerplantLogConfiguration  []finspacekxcluster.TickerplantLogConfigurationState  `json:"tickerplant_log_configuration"`
	Timeouts                     *finspacekxcluster.TimeoutsState                      `json:"timeouts"`
	VpcConfiguration             []finspacekxcluster.VpcConfigurationState             `json:"vpc_configuration"`
}

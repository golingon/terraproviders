// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	mwaaenvironment "github.com/golingon/terraproviders/aws/5.11.0/mwaaenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMwaaEnvironment creates a new instance of [MwaaEnvironment].
func NewMwaaEnvironment(name string, args MwaaEnvironmentArgs) *MwaaEnvironment {
	return &MwaaEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MwaaEnvironment)(nil)

// MwaaEnvironment represents the Terraform resource aws_mwaa_environment.
type MwaaEnvironment struct {
	Name      string
	Args      MwaaEnvironmentArgs
	state     *mwaaEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MwaaEnvironment].
func (me *MwaaEnvironment) Type() string {
	return "aws_mwaa_environment"
}

// LocalName returns the local name for [MwaaEnvironment].
func (me *MwaaEnvironment) LocalName() string {
	return me.Name
}

// Configuration returns the configuration (args) for [MwaaEnvironment].
func (me *MwaaEnvironment) Configuration() interface{} {
	return me.Args
}

// DependOn is used for other resources to depend on [MwaaEnvironment].
func (me *MwaaEnvironment) DependOn() terra.Reference {
	return terra.ReferenceResource(me)
}

// Dependencies returns the list of resources [MwaaEnvironment] depends_on.
func (me *MwaaEnvironment) Dependencies() terra.Dependencies {
	return me.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MwaaEnvironment].
func (me *MwaaEnvironment) LifecycleManagement() *terra.Lifecycle {
	return me.Lifecycle
}

// Attributes returns the attributes for [MwaaEnvironment].
func (me *MwaaEnvironment) Attributes() mwaaEnvironmentAttributes {
	return mwaaEnvironmentAttributes{ref: terra.ReferenceResource(me)}
}

// ImportState imports the given attribute values into [MwaaEnvironment]'s state.
func (me *MwaaEnvironment) ImportState(av io.Reader) error {
	me.state = &mwaaEnvironmentState{}
	if err := json.NewDecoder(av).Decode(me.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", me.Type(), me.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MwaaEnvironment] has state.
func (me *MwaaEnvironment) State() (*mwaaEnvironmentState, bool) {
	return me.state, me.state != nil
}

// StateMust returns the state for [MwaaEnvironment]. Panics if the state is nil.
func (me *MwaaEnvironment) StateMust() *mwaaEnvironmentState {
	if me.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", me.Type(), me.LocalName()))
	}
	return me.state
}

// MwaaEnvironmentArgs contains the configurations for aws_mwaa_environment.
type MwaaEnvironmentArgs struct {
	// AirflowConfigurationOptions: map of string, optional
	AirflowConfigurationOptions terra.MapValue[terra.StringValue] `hcl:"airflow_configuration_options,attr"`
	// AirflowVersion: string, optional
	AirflowVersion terra.StringValue `hcl:"airflow_version,attr"`
	// DagS3Path: string, required
	DagS3Path terra.StringValue `hcl:"dag_s3_path,attr" validate:"required"`
	// EnvironmentClass: string, optional
	EnvironmentClass terra.StringValue `hcl:"environment_class,attr"`
	// ExecutionRoleArn: string, required
	ExecutionRoleArn terra.StringValue `hcl:"execution_role_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKey: string, optional
	KmsKey terra.StringValue `hcl:"kms_key,attr"`
	// MaxWorkers: number, optional
	MaxWorkers terra.NumberValue `hcl:"max_workers,attr"`
	// MinWorkers: number, optional
	MinWorkers terra.NumberValue `hcl:"min_workers,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PluginsS3ObjectVersion: string, optional
	PluginsS3ObjectVersion terra.StringValue `hcl:"plugins_s3_object_version,attr"`
	// PluginsS3Path: string, optional
	PluginsS3Path terra.StringValue `hcl:"plugins_s3_path,attr"`
	// RequirementsS3ObjectVersion: string, optional
	RequirementsS3ObjectVersion terra.StringValue `hcl:"requirements_s3_object_version,attr"`
	// RequirementsS3Path: string, optional
	RequirementsS3Path terra.StringValue `hcl:"requirements_s3_path,attr"`
	// Schedulers: number, optional
	Schedulers terra.NumberValue `hcl:"schedulers,attr"`
	// SourceBucketArn: string, required
	SourceBucketArn terra.StringValue `hcl:"source_bucket_arn,attr" validate:"required"`
	// StartupScriptS3ObjectVersion: string, optional
	StartupScriptS3ObjectVersion terra.StringValue `hcl:"startup_script_s3_object_version,attr"`
	// StartupScriptS3Path: string, optional
	StartupScriptS3Path terra.StringValue `hcl:"startup_script_s3_path,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// WebserverAccessMode: string, optional
	WebserverAccessMode terra.StringValue `hcl:"webserver_access_mode,attr"`
	// WeeklyMaintenanceWindowStart: string, optional
	WeeklyMaintenanceWindowStart terra.StringValue `hcl:"weekly_maintenance_window_start,attr"`
	// LastUpdated: min=0
	LastUpdated []mwaaenvironment.LastUpdated `hcl:"last_updated,block" validate:"min=0"`
	// LoggingConfiguration: optional
	LoggingConfiguration *mwaaenvironment.LoggingConfiguration `hcl:"logging_configuration,block"`
	// NetworkConfiguration: required
	NetworkConfiguration *mwaaenvironment.NetworkConfiguration `hcl:"network_configuration,block" validate:"required"`
	// Timeouts: optional
	Timeouts *mwaaenvironment.Timeouts `hcl:"timeouts,block"`
}
type mwaaEnvironmentAttributes struct {
	ref terra.Reference
}

// AirflowConfigurationOptions returns a reference to field airflow_configuration_options of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) AirflowConfigurationOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](me.ref.Append("airflow_configuration_options"))
}

// AirflowVersion returns a reference to field airflow_version of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) AirflowVersion() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("airflow_version"))
}

// Arn returns a reference to field arn of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("created_at"))
}

// DagS3Path returns a reference to field dag_s3_path of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) DagS3Path() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("dag_s3_path"))
}

// EnvironmentClass returns a reference to field environment_class of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) EnvironmentClass() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("environment_class"))
}

// ExecutionRoleArn returns a reference to field execution_role_arn of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) ExecutionRoleArn() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("execution_role_arn"))
}

// Id returns a reference to field id of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("id"))
}

// KmsKey returns a reference to field kms_key of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("kms_key"))
}

// MaxWorkers returns a reference to field max_workers of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) MaxWorkers() terra.NumberValue {
	return terra.ReferenceAsNumber(me.ref.Append("max_workers"))
}

// MinWorkers returns a reference to field min_workers of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) MinWorkers() terra.NumberValue {
	return terra.ReferenceAsNumber(me.ref.Append("min_workers"))
}

// Name returns a reference to field name of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("name"))
}

// PluginsS3ObjectVersion returns a reference to field plugins_s3_object_version of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) PluginsS3ObjectVersion() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("plugins_s3_object_version"))
}

// PluginsS3Path returns a reference to field plugins_s3_path of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) PluginsS3Path() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("plugins_s3_path"))
}

// RequirementsS3ObjectVersion returns a reference to field requirements_s3_object_version of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) RequirementsS3ObjectVersion() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("requirements_s3_object_version"))
}

// RequirementsS3Path returns a reference to field requirements_s3_path of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) RequirementsS3Path() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("requirements_s3_path"))
}

// Schedulers returns a reference to field schedulers of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) Schedulers() terra.NumberValue {
	return terra.ReferenceAsNumber(me.ref.Append("schedulers"))
}

// ServiceRoleArn returns a reference to field service_role_arn of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) ServiceRoleArn() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("service_role_arn"))
}

// SourceBucketArn returns a reference to field source_bucket_arn of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) SourceBucketArn() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("source_bucket_arn"))
}

// StartupScriptS3ObjectVersion returns a reference to field startup_script_s3_object_version of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) StartupScriptS3ObjectVersion() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("startup_script_s3_object_version"))
}

// StartupScriptS3Path returns a reference to field startup_script_s3_path of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) StartupScriptS3Path() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("startup_script_s3_path"))
}

// Status returns a reference to field status of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](me.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](me.ref.Append("tags_all"))
}

// WebserverAccessMode returns a reference to field webserver_access_mode of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) WebserverAccessMode() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("webserver_access_mode"))
}

// WebserverUrl returns a reference to field webserver_url of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) WebserverUrl() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("webserver_url"))
}

// WeeklyMaintenanceWindowStart returns a reference to field weekly_maintenance_window_start of aws_mwaa_environment.
func (me mwaaEnvironmentAttributes) WeeklyMaintenanceWindowStart() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("weekly_maintenance_window_start"))
}

func (me mwaaEnvironmentAttributes) LastUpdated() terra.ListValue[mwaaenvironment.LastUpdatedAttributes] {
	return terra.ReferenceAsList[mwaaenvironment.LastUpdatedAttributes](me.ref.Append("last_updated"))
}

func (me mwaaEnvironmentAttributes) LoggingConfiguration() terra.ListValue[mwaaenvironment.LoggingConfigurationAttributes] {
	return terra.ReferenceAsList[mwaaenvironment.LoggingConfigurationAttributes](me.ref.Append("logging_configuration"))
}

func (me mwaaEnvironmentAttributes) NetworkConfiguration() terra.ListValue[mwaaenvironment.NetworkConfigurationAttributes] {
	return terra.ReferenceAsList[mwaaenvironment.NetworkConfigurationAttributes](me.ref.Append("network_configuration"))
}

func (me mwaaEnvironmentAttributes) Timeouts() mwaaenvironment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mwaaenvironment.TimeoutsAttributes](me.ref.Append("timeouts"))
}

type mwaaEnvironmentState struct {
	AirflowConfigurationOptions  map[string]string                           `json:"airflow_configuration_options"`
	AirflowVersion               string                                      `json:"airflow_version"`
	Arn                          string                                      `json:"arn"`
	CreatedAt                    string                                      `json:"created_at"`
	DagS3Path                    string                                      `json:"dag_s3_path"`
	EnvironmentClass             string                                      `json:"environment_class"`
	ExecutionRoleArn             string                                      `json:"execution_role_arn"`
	Id                           string                                      `json:"id"`
	KmsKey                       string                                      `json:"kms_key"`
	MaxWorkers                   float64                                     `json:"max_workers"`
	MinWorkers                   float64                                     `json:"min_workers"`
	Name                         string                                      `json:"name"`
	PluginsS3ObjectVersion       string                                      `json:"plugins_s3_object_version"`
	PluginsS3Path                string                                      `json:"plugins_s3_path"`
	RequirementsS3ObjectVersion  string                                      `json:"requirements_s3_object_version"`
	RequirementsS3Path           string                                      `json:"requirements_s3_path"`
	Schedulers                   float64                                     `json:"schedulers"`
	ServiceRoleArn               string                                      `json:"service_role_arn"`
	SourceBucketArn              string                                      `json:"source_bucket_arn"`
	StartupScriptS3ObjectVersion string                                      `json:"startup_script_s3_object_version"`
	StartupScriptS3Path          string                                      `json:"startup_script_s3_path"`
	Status                       string                                      `json:"status"`
	Tags                         map[string]string                           `json:"tags"`
	TagsAll                      map[string]string                           `json:"tags_all"`
	WebserverAccessMode          string                                      `json:"webserver_access_mode"`
	WebserverUrl                 string                                      `json:"webserver_url"`
	WeeklyMaintenanceWindowStart string                                      `json:"weekly_maintenance_window_start"`
	LastUpdated                  []mwaaenvironment.LastUpdatedState          `json:"last_updated"`
	LoggingConfiguration         []mwaaenvironment.LoggingConfigurationState `json:"logging_configuration"`
	NetworkConfiguration         []mwaaenvironment.NetworkConfigurationState `json:"network_configuration"`
	Timeouts                     *mwaaenvironment.TimeoutsState              `json:"timeouts"`
}

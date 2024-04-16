// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_mwaa_environment

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_mwaa_environment.
type Resource struct {
	Name      string
	Args      Args
	state     *awsMwaaEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ame *Resource) Type() string {
	return "aws_mwaa_environment"
}

// LocalName returns the local name for [Resource].
func (ame *Resource) LocalName() string {
	return ame.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ame *Resource) Configuration() interface{} {
	return ame.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ame *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ame)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ame *Resource) Dependencies() terra.Dependencies {
	return ame.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ame *Resource) LifecycleManagement() *terra.Lifecycle {
	return ame.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ame *Resource) Attributes() awsMwaaEnvironmentAttributes {
	return awsMwaaEnvironmentAttributes{ref: terra.ReferenceResource(ame)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ame *Resource) ImportState(state io.Reader) error {
	ame.state = &awsMwaaEnvironmentState{}
	if err := json.NewDecoder(state).Decode(ame.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ame.Type(), ame.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ame *Resource) State() (*awsMwaaEnvironmentState, bool) {
	return ame.state, ame.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ame *Resource) StateMust() *awsMwaaEnvironmentState {
	if ame.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ame.Type(), ame.LocalName()))
	}
	return ame.state
}

// Args contains the configurations for aws_mwaa_environment.
type Args struct {
	// AirflowConfigurationOptions: map of string, optional
	AirflowConfigurationOptions terra.MapValue[terra.StringValue] `hcl:"airflow_configuration_options,attr"`
	// AirflowVersion: string, optional
	AirflowVersion terra.StringValue `hcl:"airflow_version,attr"`
	// DagS3Path: string, required
	DagS3Path terra.StringValue `hcl:"dag_s3_path,attr" validate:"required"`
	// EndpointManagement: string, optional
	EndpointManagement terra.StringValue `hcl:"endpoint_management,attr"`
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
	// LoggingConfiguration: optional
	LoggingConfiguration *LoggingConfiguration `hcl:"logging_configuration,block"`
	// NetworkConfiguration: required
	NetworkConfiguration *NetworkConfiguration `hcl:"network_configuration,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsMwaaEnvironmentAttributes struct {
	ref terra.Reference
}

// AirflowConfigurationOptions returns a reference to field airflow_configuration_options of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) AirflowConfigurationOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ame.ref.Append("airflow_configuration_options"))
}

// AirflowVersion returns a reference to field airflow_version of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) AirflowVersion() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("airflow_version"))
}

// Arn returns a reference to field arn of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("created_at"))
}

// DagS3Path returns a reference to field dag_s3_path of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) DagS3Path() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("dag_s3_path"))
}

// EndpointManagement returns a reference to field endpoint_management of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) EndpointManagement() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("endpoint_management"))
}

// EnvironmentClass returns a reference to field environment_class of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) EnvironmentClass() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("environment_class"))
}

// ExecutionRoleArn returns a reference to field execution_role_arn of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) ExecutionRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("execution_role_arn"))
}

// Id returns a reference to field id of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("id"))
}

// KmsKey returns a reference to field kms_key of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("kms_key"))
}

// MaxWorkers returns a reference to field max_workers of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) MaxWorkers() terra.NumberValue {
	return terra.ReferenceAsNumber(ame.ref.Append("max_workers"))
}

// MinWorkers returns a reference to field min_workers of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) MinWorkers() terra.NumberValue {
	return terra.ReferenceAsNumber(ame.ref.Append("min_workers"))
}

// Name returns a reference to field name of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("name"))
}

// PluginsS3ObjectVersion returns a reference to field plugins_s3_object_version of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) PluginsS3ObjectVersion() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("plugins_s3_object_version"))
}

// PluginsS3Path returns a reference to field plugins_s3_path of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) PluginsS3Path() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("plugins_s3_path"))
}

// RequirementsS3ObjectVersion returns a reference to field requirements_s3_object_version of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) RequirementsS3ObjectVersion() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("requirements_s3_object_version"))
}

// RequirementsS3Path returns a reference to field requirements_s3_path of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) RequirementsS3Path() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("requirements_s3_path"))
}

// Schedulers returns a reference to field schedulers of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) Schedulers() terra.NumberValue {
	return terra.ReferenceAsNumber(ame.ref.Append("schedulers"))
}

// ServiceRoleArn returns a reference to field service_role_arn of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) ServiceRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("service_role_arn"))
}

// SourceBucketArn returns a reference to field source_bucket_arn of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) SourceBucketArn() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("source_bucket_arn"))
}

// StartupScriptS3ObjectVersion returns a reference to field startup_script_s3_object_version of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) StartupScriptS3ObjectVersion() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("startup_script_s3_object_version"))
}

// StartupScriptS3Path returns a reference to field startup_script_s3_path of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) StartupScriptS3Path() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("startup_script_s3_path"))
}

// Status returns a reference to field status of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ame.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ame.ref.Append("tags_all"))
}

// WebserverAccessMode returns a reference to field webserver_access_mode of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) WebserverAccessMode() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("webserver_access_mode"))
}

// WebserverUrl returns a reference to field webserver_url of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) WebserverUrl() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("webserver_url"))
}

// WeeklyMaintenanceWindowStart returns a reference to field weekly_maintenance_window_start of aws_mwaa_environment.
func (ame awsMwaaEnvironmentAttributes) WeeklyMaintenanceWindowStart() terra.StringValue {
	return terra.ReferenceAsString(ame.ref.Append("weekly_maintenance_window_start"))
}

func (ame awsMwaaEnvironmentAttributes) LastUpdated() terra.ListValue[LastUpdatedAttributes] {
	return terra.ReferenceAsList[LastUpdatedAttributes](ame.ref.Append("last_updated"))
}

func (ame awsMwaaEnvironmentAttributes) LoggingConfiguration() terra.ListValue[LoggingConfigurationAttributes] {
	return terra.ReferenceAsList[LoggingConfigurationAttributes](ame.ref.Append("logging_configuration"))
}

func (ame awsMwaaEnvironmentAttributes) NetworkConfiguration() terra.ListValue[NetworkConfigurationAttributes] {
	return terra.ReferenceAsList[NetworkConfigurationAttributes](ame.ref.Append("network_configuration"))
}

func (ame awsMwaaEnvironmentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ame.ref.Append("timeouts"))
}

type awsMwaaEnvironmentState struct {
	AirflowConfigurationOptions  map[string]string           `json:"airflow_configuration_options"`
	AirflowVersion               string                      `json:"airflow_version"`
	Arn                          string                      `json:"arn"`
	CreatedAt                    string                      `json:"created_at"`
	DagS3Path                    string                      `json:"dag_s3_path"`
	EndpointManagement           string                      `json:"endpoint_management"`
	EnvironmentClass             string                      `json:"environment_class"`
	ExecutionRoleArn             string                      `json:"execution_role_arn"`
	Id                           string                      `json:"id"`
	KmsKey                       string                      `json:"kms_key"`
	MaxWorkers                   float64                     `json:"max_workers"`
	MinWorkers                   float64                     `json:"min_workers"`
	Name                         string                      `json:"name"`
	PluginsS3ObjectVersion       string                      `json:"plugins_s3_object_version"`
	PluginsS3Path                string                      `json:"plugins_s3_path"`
	RequirementsS3ObjectVersion  string                      `json:"requirements_s3_object_version"`
	RequirementsS3Path           string                      `json:"requirements_s3_path"`
	Schedulers                   float64                     `json:"schedulers"`
	ServiceRoleArn               string                      `json:"service_role_arn"`
	SourceBucketArn              string                      `json:"source_bucket_arn"`
	StartupScriptS3ObjectVersion string                      `json:"startup_script_s3_object_version"`
	StartupScriptS3Path          string                      `json:"startup_script_s3_path"`
	Status                       string                      `json:"status"`
	Tags                         map[string]string           `json:"tags"`
	TagsAll                      map[string]string           `json:"tags_all"`
	WebserverAccessMode          string                      `json:"webserver_access_mode"`
	WebserverUrl                 string                      `json:"webserver_url"`
	WeeklyMaintenanceWindowStart string                      `json:"weekly_maintenance_window_start"`
	LastUpdated                  []LastUpdatedState          `json:"last_updated"`
	LoggingConfiguration         []LoggingConfigurationState `json:"logging_configuration"`
	NetworkConfiguration         []NetworkConfigurationState `json:"network_configuration"`
	Timeouts                     *TimeoutsState              `json:"timeouts"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudfunctionsfunction "github.com/golingon/terraproviders/google/4.63.1/cloudfunctionsfunction"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfunctionsFunction creates a new instance of [CloudfunctionsFunction].
func NewCloudfunctionsFunction(name string, args CloudfunctionsFunctionArgs) *CloudfunctionsFunction {
	return &CloudfunctionsFunction{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfunctionsFunction)(nil)

// CloudfunctionsFunction represents the Terraform resource google_cloudfunctions_function.
type CloudfunctionsFunction struct {
	Name      string
	Args      CloudfunctionsFunctionArgs
	state     *cloudfunctionsFunctionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfunctionsFunction].
func (cf *CloudfunctionsFunction) Type() string {
	return "google_cloudfunctions_function"
}

// LocalName returns the local name for [CloudfunctionsFunction].
func (cf *CloudfunctionsFunction) LocalName() string {
	return cf.Name
}

// Configuration returns the configuration (args) for [CloudfunctionsFunction].
func (cf *CloudfunctionsFunction) Configuration() interface{} {
	return cf.Args
}

// DependOn is used for other resources to depend on [CloudfunctionsFunction].
func (cf *CloudfunctionsFunction) DependOn() terra.Reference {
	return terra.ReferenceResource(cf)
}

// Dependencies returns the list of resources [CloudfunctionsFunction] depends_on.
func (cf *CloudfunctionsFunction) Dependencies() terra.Dependencies {
	return cf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfunctionsFunction].
func (cf *CloudfunctionsFunction) LifecycleManagement() *terra.Lifecycle {
	return cf.Lifecycle
}

// Attributes returns the attributes for [CloudfunctionsFunction].
func (cf *CloudfunctionsFunction) Attributes() cloudfunctionsFunctionAttributes {
	return cloudfunctionsFunctionAttributes{ref: terra.ReferenceResource(cf)}
}

// ImportState imports the given attribute values into [CloudfunctionsFunction]'s state.
func (cf *CloudfunctionsFunction) ImportState(av io.Reader) error {
	cf.state = &cloudfunctionsFunctionState{}
	if err := json.NewDecoder(av).Decode(cf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cf.Type(), cf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfunctionsFunction] has state.
func (cf *CloudfunctionsFunction) State() (*cloudfunctionsFunctionState, bool) {
	return cf.state, cf.state != nil
}

// StateMust returns the state for [CloudfunctionsFunction]. Panics if the state is nil.
func (cf *CloudfunctionsFunction) StateMust() *cloudfunctionsFunctionState {
	if cf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cf.Type(), cf.LocalName()))
	}
	return cf.state
}

// CloudfunctionsFunctionArgs contains the configurations for google_cloudfunctions_function.
type CloudfunctionsFunctionArgs struct {
	// AvailableMemoryMb: number, optional
	AvailableMemoryMb terra.NumberValue `hcl:"available_memory_mb,attr"`
	// BuildEnvironmentVariables: map of string, optional
	BuildEnvironmentVariables terra.MapValue[terra.StringValue] `hcl:"build_environment_variables,attr"`
	// BuildWorkerPool: string, optional
	BuildWorkerPool terra.StringValue `hcl:"build_worker_pool,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DockerRegistry: string, optional
	DockerRegistry terra.StringValue `hcl:"docker_registry,attr"`
	// DockerRepository: string, optional
	DockerRepository terra.StringValue `hcl:"docker_repository,attr"`
	// EntryPoint: string, optional
	EntryPoint terra.StringValue `hcl:"entry_point,attr"`
	// EnvironmentVariables: map of string, optional
	EnvironmentVariables terra.MapValue[terra.StringValue] `hcl:"environment_variables,attr"`
	// HttpsTriggerSecurityLevel: string, optional
	HttpsTriggerSecurityLevel terra.StringValue `hcl:"https_trigger_security_level,attr"`
	// HttpsTriggerUrl: string, optional
	HttpsTriggerUrl terra.StringValue `hcl:"https_trigger_url,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IngressSettings: string, optional
	IngressSettings terra.StringValue `hcl:"ingress_settings,attr"`
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MaxInstances: number, optional
	MaxInstances terra.NumberValue `hcl:"max_instances,attr"`
	// MinInstances: number, optional
	MinInstances terra.NumberValue `hcl:"min_instances,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Runtime: string, required
	Runtime terra.StringValue `hcl:"runtime,attr" validate:"required"`
	// ServiceAccountEmail: string, optional
	ServiceAccountEmail terra.StringValue `hcl:"service_account_email,attr"`
	// SourceArchiveBucket: string, optional
	SourceArchiveBucket terra.StringValue `hcl:"source_archive_bucket,attr"`
	// SourceArchiveObject: string, optional
	SourceArchiveObject terra.StringValue `hcl:"source_archive_object,attr"`
	// Timeout: number, optional
	Timeout terra.NumberValue `hcl:"timeout,attr"`
	// TriggerHttp: bool, optional
	TriggerHttp terra.BoolValue `hcl:"trigger_http,attr"`
	// VpcConnector: string, optional
	VpcConnector terra.StringValue `hcl:"vpc_connector,attr"`
	// VpcConnectorEgressSettings: string, optional
	VpcConnectorEgressSettings terra.StringValue `hcl:"vpc_connector_egress_settings,attr"`
	// EventTrigger: optional
	EventTrigger *cloudfunctionsfunction.EventTrigger `hcl:"event_trigger,block"`
	// SecretEnvironmentVariables: min=0
	SecretEnvironmentVariables []cloudfunctionsfunction.SecretEnvironmentVariables `hcl:"secret_environment_variables,block" validate:"min=0"`
	// SecretVolumes: min=0
	SecretVolumes []cloudfunctionsfunction.SecretVolumes `hcl:"secret_volumes,block" validate:"min=0"`
	// SourceRepository: optional
	SourceRepository *cloudfunctionsfunction.SourceRepository `hcl:"source_repository,block"`
	// Timeouts: optional
	Timeouts *cloudfunctionsfunction.Timeouts `hcl:"timeouts,block"`
}
type cloudfunctionsFunctionAttributes struct {
	ref terra.Reference
}

// AvailableMemoryMb returns a reference to field available_memory_mb of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) AvailableMemoryMb() terra.NumberValue {
	return terra.ReferenceAsNumber(cf.ref.Append("available_memory_mb"))
}

// BuildEnvironmentVariables returns a reference to field build_environment_variables of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) BuildEnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cf.ref.Append("build_environment_variables"))
}

// BuildWorkerPool returns a reference to field build_worker_pool of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) BuildWorkerPool() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("build_worker_pool"))
}

// Description returns a reference to field description of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("description"))
}

// DockerRegistry returns a reference to field docker_registry of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) DockerRegistry() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("docker_registry"))
}

// DockerRepository returns a reference to field docker_repository of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) DockerRepository() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("docker_repository"))
}

// EntryPoint returns a reference to field entry_point of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) EntryPoint() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("entry_point"))
}

// EnvironmentVariables returns a reference to field environment_variables of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) EnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cf.ref.Append("environment_variables"))
}

// HttpsTriggerSecurityLevel returns a reference to field https_trigger_security_level of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) HttpsTriggerSecurityLevel() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("https_trigger_security_level"))
}

// HttpsTriggerUrl returns a reference to field https_trigger_url of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) HttpsTriggerUrl() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("https_trigger_url"))
}

// Id returns a reference to field id of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("id"))
}

// IngressSettings returns a reference to field ingress_settings of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) IngressSettings() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("ingress_settings"))
}

// KmsKeyName returns a reference to field kms_key_name of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("kms_key_name"))
}

// Labels returns a reference to field labels of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cf.ref.Append("labels"))
}

// MaxInstances returns a reference to field max_instances of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) MaxInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(cf.ref.Append("max_instances"))
}

// MinInstances returns a reference to field min_instances of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) MinInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(cf.ref.Append("min_instances"))
}

// Name returns a reference to field name of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("project"))
}

// Region returns a reference to field region of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("region"))
}

// Runtime returns a reference to field runtime of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) Runtime() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("runtime"))
}

// ServiceAccountEmail returns a reference to field service_account_email of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("service_account_email"))
}

// SourceArchiveBucket returns a reference to field source_archive_bucket of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) SourceArchiveBucket() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("source_archive_bucket"))
}

// SourceArchiveObject returns a reference to field source_archive_object of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) SourceArchiveObject() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("source_archive_object"))
}

// Timeout returns a reference to field timeout of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceAsNumber(cf.ref.Append("timeout"))
}

// TriggerHttp returns a reference to field trigger_http of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) TriggerHttp() terra.BoolValue {
	return terra.ReferenceAsBool(cf.ref.Append("trigger_http"))
}

// VpcConnector returns a reference to field vpc_connector of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) VpcConnector() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("vpc_connector"))
}

// VpcConnectorEgressSettings returns a reference to field vpc_connector_egress_settings of google_cloudfunctions_function.
func (cf cloudfunctionsFunctionAttributes) VpcConnectorEgressSettings() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("vpc_connector_egress_settings"))
}

func (cf cloudfunctionsFunctionAttributes) EventTrigger() terra.ListValue[cloudfunctionsfunction.EventTriggerAttributes] {
	return terra.ReferenceAsList[cloudfunctionsfunction.EventTriggerAttributes](cf.ref.Append("event_trigger"))
}

func (cf cloudfunctionsFunctionAttributes) SecretEnvironmentVariables() terra.ListValue[cloudfunctionsfunction.SecretEnvironmentVariablesAttributes] {
	return terra.ReferenceAsList[cloudfunctionsfunction.SecretEnvironmentVariablesAttributes](cf.ref.Append("secret_environment_variables"))
}

func (cf cloudfunctionsFunctionAttributes) SecretVolumes() terra.ListValue[cloudfunctionsfunction.SecretVolumesAttributes] {
	return terra.ReferenceAsList[cloudfunctionsfunction.SecretVolumesAttributes](cf.ref.Append("secret_volumes"))
}

func (cf cloudfunctionsFunctionAttributes) SourceRepository() terra.ListValue[cloudfunctionsfunction.SourceRepositoryAttributes] {
	return terra.ReferenceAsList[cloudfunctionsfunction.SourceRepositoryAttributes](cf.ref.Append("source_repository"))
}

func (cf cloudfunctionsFunctionAttributes) Timeouts() cloudfunctionsfunction.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudfunctionsfunction.TimeoutsAttributes](cf.ref.Append("timeouts"))
}

type cloudfunctionsFunctionState struct {
	AvailableMemoryMb          float64                                                  `json:"available_memory_mb"`
	BuildEnvironmentVariables  map[string]string                                        `json:"build_environment_variables"`
	BuildWorkerPool            string                                                   `json:"build_worker_pool"`
	Description                string                                                   `json:"description"`
	DockerRegistry             string                                                   `json:"docker_registry"`
	DockerRepository           string                                                   `json:"docker_repository"`
	EntryPoint                 string                                                   `json:"entry_point"`
	EnvironmentVariables       map[string]string                                        `json:"environment_variables"`
	HttpsTriggerSecurityLevel  string                                                   `json:"https_trigger_security_level"`
	HttpsTriggerUrl            string                                                   `json:"https_trigger_url"`
	Id                         string                                                   `json:"id"`
	IngressSettings            string                                                   `json:"ingress_settings"`
	KmsKeyName                 string                                                   `json:"kms_key_name"`
	Labels                     map[string]string                                        `json:"labels"`
	MaxInstances               float64                                                  `json:"max_instances"`
	MinInstances               float64                                                  `json:"min_instances"`
	Name                       string                                                   `json:"name"`
	Project                    string                                                   `json:"project"`
	Region                     string                                                   `json:"region"`
	Runtime                    string                                                   `json:"runtime"`
	ServiceAccountEmail        string                                                   `json:"service_account_email"`
	SourceArchiveBucket        string                                                   `json:"source_archive_bucket"`
	SourceArchiveObject        string                                                   `json:"source_archive_object"`
	Timeout                    float64                                                  `json:"timeout"`
	TriggerHttp                bool                                                     `json:"trigger_http"`
	VpcConnector               string                                                   `json:"vpc_connector"`
	VpcConnectorEgressSettings string                                                   `json:"vpc_connector_egress_settings"`
	EventTrigger               []cloudfunctionsfunction.EventTriggerState               `json:"event_trigger"`
	SecretEnvironmentVariables []cloudfunctionsfunction.SecretEnvironmentVariablesState `json:"secret_environment_variables"`
	SecretVolumes              []cloudfunctionsfunction.SecretVolumesState              `json:"secret_volumes"`
	SourceRepository           []cloudfunctionsfunction.SourceRepositoryState           `json:"source_repository"`
	Timeouts                   *cloudfunctionsfunction.TimeoutsState                    `json:"timeouts"`
}

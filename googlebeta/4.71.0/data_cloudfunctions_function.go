// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacloudfunctionsfunction "github.com/golingon/terraproviders/googlebeta/4.71.0/datacloudfunctionsfunction"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudfunctionsFunction creates a new instance of [DataCloudfunctionsFunction].
func NewDataCloudfunctionsFunction(name string, args DataCloudfunctionsFunctionArgs) *DataCloudfunctionsFunction {
	return &DataCloudfunctionsFunction{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudfunctionsFunction)(nil)

// DataCloudfunctionsFunction represents the Terraform data resource google_cloudfunctions_function.
type DataCloudfunctionsFunction struct {
	Name string
	Args DataCloudfunctionsFunctionArgs
}

// DataSource returns the Terraform object type for [DataCloudfunctionsFunction].
func (cf *DataCloudfunctionsFunction) DataSource() string {
	return "google_cloudfunctions_function"
}

// LocalName returns the local name for [DataCloudfunctionsFunction].
func (cf *DataCloudfunctionsFunction) LocalName() string {
	return cf.Name
}

// Configuration returns the configuration (args) for [DataCloudfunctionsFunction].
func (cf *DataCloudfunctionsFunction) Configuration() interface{} {
	return cf.Args
}

// Attributes returns the attributes for [DataCloudfunctionsFunction].
func (cf *DataCloudfunctionsFunction) Attributes() dataCloudfunctionsFunctionAttributes {
	return dataCloudfunctionsFunctionAttributes{ref: terra.ReferenceDataResource(cf)}
}

// DataCloudfunctionsFunctionArgs contains the configurations for google_cloudfunctions_function.
type DataCloudfunctionsFunctionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// EventTrigger: min=0
	EventTrigger []datacloudfunctionsfunction.EventTrigger `hcl:"event_trigger,block" validate:"min=0"`
	// SecretEnvironmentVariables: min=0
	SecretEnvironmentVariables []datacloudfunctionsfunction.SecretEnvironmentVariables `hcl:"secret_environment_variables,block" validate:"min=0"`
	// SecretVolumes: min=0
	SecretVolumes []datacloudfunctionsfunction.SecretVolumes `hcl:"secret_volumes,block" validate:"min=0"`
	// SourceRepository: min=0
	SourceRepository []datacloudfunctionsfunction.SourceRepository `hcl:"source_repository,block" validate:"min=0"`
}
type dataCloudfunctionsFunctionAttributes struct {
	ref terra.Reference
}

// AvailableMemoryMb returns a reference to field available_memory_mb of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) AvailableMemoryMb() terra.NumberValue {
	return terra.ReferenceAsNumber(cf.ref.Append("available_memory_mb"))
}

// BuildEnvironmentVariables returns a reference to field build_environment_variables of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) BuildEnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cf.ref.Append("build_environment_variables"))
}

// BuildWorkerPool returns a reference to field build_worker_pool of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) BuildWorkerPool() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("build_worker_pool"))
}

// Description returns a reference to field description of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("description"))
}

// DockerRegistry returns a reference to field docker_registry of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) DockerRegistry() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("docker_registry"))
}

// DockerRepository returns a reference to field docker_repository of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) DockerRepository() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("docker_repository"))
}

// EntryPoint returns a reference to field entry_point of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) EntryPoint() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("entry_point"))
}

// EnvironmentVariables returns a reference to field environment_variables of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) EnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cf.ref.Append("environment_variables"))
}

// HttpsTriggerSecurityLevel returns a reference to field https_trigger_security_level of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) HttpsTriggerSecurityLevel() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("https_trigger_security_level"))
}

// HttpsTriggerUrl returns a reference to field https_trigger_url of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) HttpsTriggerUrl() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("https_trigger_url"))
}

// Id returns a reference to field id of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("id"))
}

// IngressSettings returns a reference to field ingress_settings of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) IngressSettings() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("ingress_settings"))
}

// KmsKeyName returns a reference to field kms_key_name of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("kms_key_name"))
}

// Labels returns a reference to field labels of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cf.ref.Append("labels"))
}

// MaxInstances returns a reference to field max_instances of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) MaxInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(cf.ref.Append("max_instances"))
}

// MinInstances returns a reference to field min_instances of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) MinInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(cf.ref.Append("min_instances"))
}

// Name returns a reference to field name of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("project"))
}

// Region returns a reference to field region of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("region"))
}

// Runtime returns a reference to field runtime of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) Runtime() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("runtime"))
}

// ServiceAccountEmail returns a reference to field service_account_email of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("service_account_email"))
}

// SourceArchiveBucket returns a reference to field source_archive_bucket of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) SourceArchiveBucket() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("source_archive_bucket"))
}

// SourceArchiveObject returns a reference to field source_archive_object of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) SourceArchiveObject() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("source_archive_object"))
}

// Status returns a reference to field status of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("status"))
}

// Timeout returns a reference to field timeout of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceAsNumber(cf.ref.Append("timeout"))
}

// TriggerHttp returns a reference to field trigger_http of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) TriggerHttp() terra.BoolValue {
	return terra.ReferenceAsBool(cf.ref.Append("trigger_http"))
}

// VpcConnector returns a reference to field vpc_connector of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) VpcConnector() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("vpc_connector"))
}

// VpcConnectorEgressSettings returns a reference to field vpc_connector_egress_settings of google_cloudfunctions_function.
func (cf dataCloudfunctionsFunctionAttributes) VpcConnectorEgressSettings() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("vpc_connector_egress_settings"))
}

func (cf dataCloudfunctionsFunctionAttributes) EventTrigger() terra.ListValue[datacloudfunctionsfunction.EventTriggerAttributes] {
	return terra.ReferenceAsList[datacloudfunctionsfunction.EventTriggerAttributes](cf.ref.Append("event_trigger"))
}

func (cf dataCloudfunctionsFunctionAttributes) SecretEnvironmentVariables() terra.ListValue[datacloudfunctionsfunction.SecretEnvironmentVariablesAttributes] {
	return terra.ReferenceAsList[datacloudfunctionsfunction.SecretEnvironmentVariablesAttributes](cf.ref.Append("secret_environment_variables"))
}

func (cf dataCloudfunctionsFunctionAttributes) SecretVolumes() terra.ListValue[datacloudfunctionsfunction.SecretVolumesAttributes] {
	return terra.ReferenceAsList[datacloudfunctionsfunction.SecretVolumesAttributes](cf.ref.Append("secret_volumes"))
}

func (cf dataCloudfunctionsFunctionAttributes) SourceRepository() terra.ListValue[datacloudfunctionsfunction.SourceRepositoryAttributes] {
	return terra.ReferenceAsList[datacloudfunctionsfunction.SourceRepositoryAttributes](cf.ref.Append("source_repository"))
}

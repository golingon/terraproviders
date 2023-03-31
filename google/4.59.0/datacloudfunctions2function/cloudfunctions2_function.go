// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacloudfunctions2function

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BuildConfig struct {
	// Source: min=0
	Source []Source `hcl:"source,block" validate:"min=0"`
}

type Source struct {
	// RepoSource: min=0
	RepoSource []RepoSource `hcl:"repo_source,block" validate:"min=0"`
	// StorageSource: min=0
	StorageSource []StorageSource `hcl:"storage_source,block" validate:"min=0"`
}

type RepoSource struct{}

type StorageSource struct{}

type EventTrigger struct {
	// EventFilters: min=0
	EventFilters []EventFilters `hcl:"event_filters,block" validate:"min=0"`
}

type EventFilters struct{}

type ServiceConfig struct {
	// SecretEnvironmentVariables: min=0
	SecretEnvironmentVariables []SecretEnvironmentVariables `hcl:"secret_environment_variables,block" validate:"min=0"`
	// SecretVolumes: min=0
	SecretVolumes []SecretVolumes `hcl:"secret_volumes,block" validate:"min=0"`
}

type SecretEnvironmentVariables struct{}

type SecretVolumes struct {
	// Versions: min=0
	Versions []Versions `hcl:"versions,block" validate:"min=0"`
}

type Versions struct{}

type BuildConfigAttributes struct {
	ref terra.Reference
}

func (bc BuildConfigAttributes) InternalRef() terra.Reference {
	return bc.ref
}

func (bc BuildConfigAttributes) InternalWithRef(ref terra.Reference) BuildConfigAttributes {
	return BuildConfigAttributes{ref: ref}
}

func (bc BuildConfigAttributes) InternalTokens() hclwrite.Tokens {
	return bc.ref.InternalTokens()
}

func (bc BuildConfigAttributes) Build() terra.StringValue {
	return terra.ReferenceString(bc.ref.Append("build"))
}

func (bc BuildConfigAttributes) DockerRepository() terra.StringValue {
	return terra.ReferenceString(bc.ref.Append("docker_repository"))
}

func (bc BuildConfigAttributes) EntryPoint() terra.StringValue {
	return terra.ReferenceString(bc.ref.Append("entry_point"))
}

func (bc BuildConfigAttributes) EnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](bc.ref.Append("environment_variables"))
}

func (bc BuildConfigAttributes) Runtime() terra.StringValue {
	return terra.ReferenceString(bc.ref.Append("runtime"))
}

func (bc BuildConfigAttributes) WorkerPool() terra.StringValue {
	return terra.ReferenceString(bc.ref.Append("worker_pool"))
}

func (bc BuildConfigAttributes) Source() terra.ListValue[SourceAttributes] {
	return terra.ReferenceList[SourceAttributes](bc.ref.Append("source"))
}

type SourceAttributes struct {
	ref terra.Reference
}

func (s SourceAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SourceAttributes) InternalWithRef(ref terra.Reference) SourceAttributes {
	return SourceAttributes{ref: ref}
}

func (s SourceAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SourceAttributes) RepoSource() terra.ListValue[RepoSourceAttributes] {
	return terra.ReferenceList[RepoSourceAttributes](s.ref.Append("repo_source"))
}

func (s SourceAttributes) StorageSource() terra.ListValue[StorageSourceAttributes] {
	return terra.ReferenceList[StorageSourceAttributes](s.ref.Append("storage_source"))
}

type RepoSourceAttributes struct {
	ref terra.Reference
}

func (rs RepoSourceAttributes) InternalRef() terra.Reference {
	return rs.ref
}

func (rs RepoSourceAttributes) InternalWithRef(ref terra.Reference) RepoSourceAttributes {
	return RepoSourceAttributes{ref: ref}
}

func (rs RepoSourceAttributes) InternalTokens() hclwrite.Tokens {
	return rs.ref.InternalTokens()
}

func (rs RepoSourceAttributes) BranchName() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("branch_name"))
}

func (rs RepoSourceAttributes) CommitSha() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("commit_sha"))
}

func (rs RepoSourceAttributes) Dir() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("dir"))
}

func (rs RepoSourceAttributes) InvertRegex() terra.BoolValue {
	return terra.ReferenceBool(rs.ref.Append("invert_regex"))
}

func (rs RepoSourceAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("project_id"))
}

func (rs RepoSourceAttributes) RepoName() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("repo_name"))
}

func (rs RepoSourceAttributes) TagName() terra.StringValue {
	return terra.ReferenceString(rs.ref.Append("tag_name"))
}

type StorageSourceAttributes struct {
	ref terra.Reference
}

func (ss StorageSourceAttributes) InternalRef() terra.Reference {
	return ss.ref
}

func (ss StorageSourceAttributes) InternalWithRef(ref terra.Reference) StorageSourceAttributes {
	return StorageSourceAttributes{ref: ref}
}

func (ss StorageSourceAttributes) InternalTokens() hclwrite.Tokens {
	return ss.ref.InternalTokens()
}

func (ss StorageSourceAttributes) Bucket() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("bucket"))
}

func (ss StorageSourceAttributes) Generation() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("generation"))
}

func (ss StorageSourceAttributes) Object() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("object"))
}

type EventTriggerAttributes struct {
	ref terra.Reference
}

func (et EventTriggerAttributes) InternalRef() terra.Reference {
	return et.ref
}

func (et EventTriggerAttributes) InternalWithRef(ref terra.Reference) EventTriggerAttributes {
	return EventTriggerAttributes{ref: ref}
}

func (et EventTriggerAttributes) InternalTokens() hclwrite.Tokens {
	return et.ref.InternalTokens()
}

func (et EventTriggerAttributes) EventType() terra.StringValue {
	return terra.ReferenceString(et.ref.Append("event_type"))
}

func (et EventTriggerAttributes) PubsubTopic() terra.StringValue {
	return terra.ReferenceString(et.ref.Append("pubsub_topic"))
}

func (et EventTriggerAttributes) RetryPolicy() terra.StringValue {
	return terra.ReferenceString(et.ref.Append("retry_policy"))
}

func (et EventTriggerAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceString(et.ref.Append("service_account_email"))
}

func (et EventTriggerAttributes) Trigger() terra.StringValue {
	return terra.ReferenceString(et.ref.Append("trigger"))
}

func (et EventTriggerAttributes) TriggerRegion() terra.StringValue {
	return terra.ReferenceString(et.ref.Append("trigger_region"))
}

func (et EventTriggerAttributes) EventFilters() terra.SetValue[EventFiltersAttributes] {
	return terra.ReferenceSet[EventFiltersAttributes](et.ref.Append("event_filters"))
}

type EventFiltersAttributes struct {
	ref terra.Reference
}

func (ef EventFiltersAttributes) InternalRef() terra.Reference {
	return ef.ref
}

func (ef EventFiltersAttributes) InternalWithRef(ref terra.Reference) EventFiltersAttributes {
	return EventFiltersAttributes{ref: ref}
}

func (ef EventFiltersAttributes) InternalTokens() hclwrite.Tokens {
	return ef.ref.InternalTokens()
}

func (ef EventFiltersAttributes) Attribute() terra.StringValue {
	return terra.ReferenceString(ef.ref.Append("attribute"))
}

func (ef EventFiltersAttributes) Operator() terra.StringValue {
	return terra.ReferenceString(ef.ref.Append("operator"))
}

func (ef EventFiltersAttributes) Value() terra.StringValue {
	return terra.ReferenceString(ef.ref.Append("value"))
}

type ServiceConfigAttributes struct {
	ref terra.Reference
}

func (sc ServiceConfigAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc ServiceConfigAttributes) InternalWithRef(ref terra.Reference) ServiceConfigAttributes {
	return ServiceConfigAttributes{ref: ref}
}

func (sc ServiceConfigAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc ServiceConfigAttributes) AllTrafficOnLatestRevision() terra.BoolValue {
	return terra.ReferenceBool(sc.ref.Append("all_traffic_on_latest_revision"))
}

func (sc ServiceConfigAttributes) AvailableCpu() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("available_cpu"))
}

func (sc ServiceConfigAttributes) AvailableMemory() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("available_memory"))
}

func (sc ServiceConfigAttributes) EnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sc.ref.Append("environment_variables"))
}

func (sc ServiceConfigAttributes) GcfUri() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("gcf_uri"))
}

func (sc ServiceConfigAttributes) IngressSettings() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("ingress_settings"))
}

func (sc ServiceConfigAttributes) MaxInstanceCount() terra.NumberValue {
	return terra.ReferenceNumber(sc.ref.Append("max_instance_count"))
}

func (sc ServiceConfigAttributes) MaxInstanceRequestConcurrency() terra.NumberValue {
	return terra.ReferenceNumber(sc.ref.Append("max_instance_request_concurrency"))
}

func (sc ServiceConfigAttributes) MinInstanceCount() terra.NumberValue {
	return terra.ReferenceNumber(sc.ref.Append("min_instance_count"))
}

func (sc ServiceConfigAttributes) Service() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("service"))
}

func (sc ServiceConfigAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("service_account_email"))
}

func (sc ServiceConfigAttributes) TimeoutSeconds() terra.NumberValue {
	return terra.ReferenceNumber(sc.ref.Append("timeout_seconds"))
}

func (sc ServiceConfigAttributes) Uri() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("uri"))
}

func (sc ServiceConfigAttributes) VpcConnector() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("vpc_connector"))
}

func (sc ServiceConfigAttributes) VpcConnectorEgressSettings() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("vpc_connector_egress_settings"))
}

func (sc ServiceConfigAttributes) SecretEnvironmentVariables() terra.ListValue[SecretEnvironmentVariablesAttributes] {
	return terra.ReferenceList[SecretEnvironmentVariablesAttributes](sc.ref.Append("secret_environment_variables"))
}

func (sc ServiceConfigAttributes) SecretVolumes() terra.ListValue[SecretVolumesAttributes] {
	return terra.ReferenceList[SecretVolumesAttributes](sc.ref.Append("secret_volumes"))
}

type SecretEnvironmentVariablesAttributes struct {
	ref terra.Reference
}

func (sev SecretEnvironmentVariablesAttributes) InternalRef() terra.Reference {
	return sev.ref
}

func (sev SecretEnvironmentVariablesAttributes) InternalWithRef(ref terra.Reference) SecretEnvironmentVariablesAttributes {
	return SecretEnvironmentVariablesAttributes{ref: ref}
}

func (sev SecretEnvironmentVariablesAttributes) InternalTokens() hclwrite.Tokens {
	return sev.ref.InternalTokens()
}

func (sev SecretEnvironmentVariablesAttributes) Key() terra.StringValue {
	return terra.ReferenceString(sev.ref.Append("key"))
}

func (sev SecretEnvironmentVariablesAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceString(sev.ref.Append("project_id"))
}

func (sev SecretEnvironmentVariablesAttributes) Secret() terra.StringValue {
	return terra.ReferenceString(sev.ref.Append("secret"))
}

func (sev SecretEnvironmentVariablesAttributes) Version() terra.StringValue {
	return terra.ReferenceString(sev.ref.Append("version"))
}

type SecretVolumesAttributes struct {
	ref terra.Reference
}

func (sv SecretVolumesAttributes) InternalRef() terra.Reference {
	return sv.ref
}

func (sv SecretVolumesAttributes) InternalWithRef(ref terra.Reference) SecretVolumesAttributes {
	return SecretVolumesAttributes{ref: ref}
}

func (sv SecretVolumesAttributes) InternalTokens() hclwrite.Tokens {
	return sv.ref.InternalTokens()
}

func (sv SecretVolumesAttributes) MountPath() terra.StringValue {
	return terra.ReferenceString(sv.ref.Append("mount_path"))
}

func (sv SecretVolumesAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceString(sv.ref.Append("project_id"))
}

func (sv SecretVolumesAttributes) Secret() terra.StringValue {
	return terra.ReferenceString(sv.ref.Append("secret"))
}

func (sv SecretVolumesAttributes) Versions() terra.ListValue[VersionsAttributes] {
	return terra.ReferenceList[VersionsAttributes](sv.ref.Append("versions"))
}

type VersionsAttributes struct {
	ref terra.Reference
}

func (v VersionsAttributes) InternalRef() terra.Reference {
	return v.ref
}

func (v VersionsAttributes) InternalWithRef(ref terra.Reference) VersionsAttributes {
	return VersionsAttributes{ref: ref}
}

func (v VersionsAttributes) InternalTokens() hclwrite.Tokens {
	return v.ref.InternalTokens()
}

func (v VersionsAttributes) Path() terra.StringValue {
	return terra.ReferenceString(v.ref.Append("path"))
}

func (v VersionsAttributes) Version() terra.StringValue {
	return terra.ReferenceString(v.ref.Append("version"))
}

type BuildConfigState struct {
	Build                string            `json:"build"`
	DockerRepository     string            `json:"docker_repository"`
	EntryPoint           string            `json:"entry_point"`
	EnvironmentVariables map[string]string `json:"environment_variables"`
	Runtime              string            `json:"runtime"`
	WorkerPool           string            `json:"worker_pool"`
	Source               []SourceState     `json:"source"`
}

type SourceState struct {
	RepoSource    []RepoSourceState    `json:"repo_source"`
	StorageSource []StorageSourceState `json:"storage_source"`
}

type RepoSourceState struct {
	BranchName  string `json:"branch_name"`
	CommitSha   string `json:"commit_sha"`
	Dir         string `json:"dir"`
	InvertRegex bool   `json:"invert_regex"`
	ProjectId   string `json:"project_id"`
	RepoName    string `json:"repo_name"`
	TagName     string `json:"tag_name"`
}

type StorageSourceState struct {
	Bucket     string  `json:"bucket"`
	Generation float64 `json:"generation"`
	Object     string  `json:"object"`
}

type EventTriggerState struct {
	EventType           string              `json:"event_type"`
	PubsubTopic         string              `json:"pubsub_topic"`
	RetryPolicy         string              `json:"retry_policy"`
	ServiceAccountEmail string              `json:"service_account_email"`
	Trigger             string              `json:"trigger"`
	TriggerRegion       string              `json:"trigger_region"`
	EventFilters        []EventFiltersState `json:"event_filters"`
}

type EventFiltersState struct {
	Attribute string `json:"attribute"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
}

type ServiceConfigState struct {
	AllTrafficOnLatestRevision    bool                              `json:"all_traffic_on_latest_revision"`
	AvailableCpu                  string                            `json:"available_cpu"`
	AvailableMemory               string                            `json:"available_memory"`
	EnvironmentVariables          map[string]string                 `json:"environment_variables"`
	GcfUri                        string                            `json:"gcf_uri"`
	IngressSettings               string                            `json:"ingress_settings"`
	MaxInstanceCount              float64                           `json:"max_instance_count"`
	MaxInstanceRequestConcurrency float64                           `json:"max_instance_request_concurrency"`
	MinInstanceCount              float64                           `json:"min_instance_count"`
	Service                       string                            `json:"service"`
	ServiceAccountEmail           string                            `json:"service_account_email"`
	TimeoutSeconds                float64                           `json:"timeout_seconds"`
	Uri                           string                            `json:"uri"`
	VpcConnector                  string                            `json:"vpc_connector"`
	VpcConnectorEgressSettings    string                            `json:"vpc_connector_egress_settings"`
	SecretEnvironmentVariables    []SecretEnvironmentVariablesState `json:"secret_environment_variables"`
	SecretVolumes                 []SecretVolumesState              `json:"secret_volumes"`
}

type SecretEnvironmentVariablesState struct {
	Key       string `json:"key"`
	ProjectId string `json:"project_id"`
	Secret    string `json:"secret"`
	Version   string `json:"version"`
}

type SecretVolumesState struct {
	MountPath string          `json:"mount_path"`
	ProjectId string          `json:"project_id"`
	Secret    string          `json:"secret"`
	Versions  []VersionsState `json:"versions"`
}

type VersionsState struct {
	Path    string `json:"path"`
	Version string `json:"version"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudfunctionsfunction

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EventTrigger struct {
	// EventType: string, required
	EventType terra.StringValue `hcl:"event_type,attr" validate:"required"`
	// Resource: string, required
	Resource terra.StringValue `hcl:"resource,attr" validate:"required"`
	// FailurePolicy: optional
	FailurePolicy *FailurePolicy `hcl:"failure_policy,block"`
}

type FailurePolicy struct {
	// Retry: bool, required
	Retry terra.BoolValue `hcl:"retry,attr" validate:"required"`
}

type SecretEnvironmentVariables struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// ProjectId: string, optional
	ProjectId terra.StringValue `hcl:"project_id,attr"`
	// Secret: string, required
	Secret terra.StringValue `hcl:"secret,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
}

type SecretVolumes struct {
	// MountPath: string, required
	MountPath terra.StringValue `hcl:"mount_path,attr" validate:"required"`
	// ProjectId: string, optional
	ProjectId terra.StringValue `hcl:"project_id,attr"`
	// Secret: string, required
	Secret terra.StringValue `hcl:"secret,attr" validate:"required"`
	// Versions: min=0
	Versions []Versions `hcl:"versions,block" validate:"min=0"`
}

type Versions struct {
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
}

type SourceRepository struct {
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type EventTriggerAttributes struct {
	ref terra.Reference
}

func (et EventTriggerAttributes) InternalRef() (terra.Reference, error) {
	return et.ref, nil
}

func (et EventTriggerAttributes) InternalWithRef(ref terra.Reference) EventTriggerAttributes {
	return EventTriggerAttributes{ref: ref}
}

func (et EventTriggerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return et.ref.InternalTokens()
}

func (et EventTriggerAttributes) EventType() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("event_type"))
}

func (et EventTriggerAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("resource"))
}

func (et EventTriggerAttributes) FailurePolicy() terra.ListValue[FailurePolicyAttributes] {
	return terra.ReferenceAsList[FailurePolicyAttributes](et.ref.Append("failure_policy"))
}

type FailurePolicyAttributes struct {
	ref terra.Reference
}

func (fp FailurePolicyAttributes) InternalRef() (terra.Reference, error) {
	return fp.ref, nil
}

func (fp FailurePolicyAttributes) InternalWithRef(ref terra.Reference) FailurePolicyAttributes {
	return FailurePolicyAttributes{ref: ref}
}

func (fp FailurePolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fp.ref.InternalTokens()
}

func (fp FailurePolicyAttributes) Retry() terra.BoolValue {
	return terra.ReferenceAsBool(fp.ref.Append("retry"))
}

type SecretEnvironmentVariablesAttributes struct {
	ref terra.Reference
}

func (sev SecretEnvironmentVariablesAttributes) InternalRef() (terra.Reference, error) {
	return sev.ref, nil
}

func (sev SecretEnvironmentVariablesAttributes) InternalWithRef(ref terra.Reference) SecretEnvironmentVariablesAttributes {
	return SecretEnvironmentVariablesAttributes{ref: ref}
}

func (sev SecretEnvironmentVariablesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sev.ref.InternalTokens()
}

func (sev SecretEnvironmentVariablesAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sev.ref.Append("key"))
}

func (sev SecretEnvironmentVariablesAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(sev.ref.Append("project_id"))
}

func (sev SecretEnvironmentVariablesAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(sev.ref.Append("secret"))
}

func (sev SecretEnvironmentVariablesAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(sev.ref.Append("version"))
}

type SecretVolumesAttributes struct {
	ref terra.Reference
}

func (sv SecretVolumesAttributes) InternalRef() (terra.Reference, error) {
	return sv.ref, nil
}

func (sv SecretVolumesAttributes) InternalWithRef(ref terra.Reference) SecretVolumesAttributes {
	return SecretVolumesAttributes{ref: ref}
}

func (sv SecretVolumesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sv.ref.InternalTokens()
}

func (sv SecretVolumesAttributes) MountPath() terra.StringValue {
	return terra.ReferenceAsString(sv.ref.Append("mount_path"))
}

func (sv SecretVolumesAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(sv.ref.Append("project_id"))
}

func (sv SecretVolumesAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(sv.ref.Append("secret"))
}

func (sv SecretVolumesAttributes) Versions() terra.ListValue[VersionsAttributes] {
	return terra.ReferenceAsList[VersionsAttributes](sv.ref.Append("versions"))
}

type VersionsAttributes struct {
	ref terra.Reference
}

func (v VersionsAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VersionsAttributes) InternalWithRef(ref terra.Reference) VersionsAttributes {
	return VersionsAttributes{ref: ref}
}

func (v VersionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v VersionsAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("path"))
}

func (v VersionsAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("version"))
}

type SourceRepositoryAttributes struct {
	ref terra.Reference
}

func (sr SourceRepositoryAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr SourceRepositoryAttributes) InternalWithRef(ref terra.Reference) SourceRepositoryAttributes {
	return SourceRepositoryAttributes{ref: ref}
}

func (sr SourceRepositoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr SourceRepositoryAttributes) DeployedUrl() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("deployed_url"))
}

func (sr SourceRepositoryAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("url"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type EventTriggerState struct {
	EventType     string               `json:"event_type"`
	Resource      string               `json:"resource"`
	FailurePolicy []FailurePolicyState `json:"failure_policy"`
}

type FailurePolicyState struct {
	Retry bool `json:"retry"`
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

type SourceRepositoryState struct {
	DeployedUrl string `json:"deployed_url"`
	Url         string `json:"url"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
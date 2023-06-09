// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacloudrunservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Metadata struct{}

type Status struct {
	// Conditions: min=0
	Conditions []Conditions `hcl:"conditions,block" validate:"min=0"`
}

type Conditions struct{}

type Template struct {
	// TemplateMetadata: min=0
	Metadata []TemplateMetadata `hcl:"metadata,block" validate:"min=0"`
	// Spec: min=0
	Spec []Spec `hcl:"spec,block" validate:"min=0"`
}

type TemplateMetadata struct{}

type Spec struct {
	// Containers: min=0
	Containers []Containers `hcl:"containers,block" validate:"min=0"`
	// Volumes: min=0
	Volumes []Volumes `hcl:"volumes,block" validate:"min=0"`
}

type Containers struct {
	// Env: min=0
	Env []Env `hcl:"env,block" validate:"min=0"`
	// EnvFrom: min=0
	EnvFrom []EnvFrom `hcl:"env_from,block" validate:"min=0"`
	// Ports: min=0
	Ports []Ports `hcl:"ports,block" validate:"min=0"`
	// Resources: min=0
	Resources []Resources `hcl:"resources,block" validate:"min=0"`
	// VolumeMounts: min=0
	VolumeMounts []VolumeMounts `hcl:"volume_mounts,block" validate:"min=0"`
}

type Env struct {
	// ValueFrom: min=0
	ValueFrom []ValueFrom `hcl:"value_from,block" validate:"min=0"`
}

type ValueFrom struct {
	// SecretKeyRef: min=0
	SecretKeyRef []SecretKeyRef `hcl:"secret_key_ref,block" validate:"min=0"`
}

type SecretKeyRef struct{}

type EnvFrom struct {
	// ConfigMapRef: min=0
	ConfigMapRef []ConfigMapRef `hcl:"config_map_ref,block" validate:"min=0"`
	// SecretRef: min=0
	SecretRef []SecretRef `hcl:"secret_ref,block" validate:"min=0"`
}

type ConfigMapRef struct {
	// ConfigMapRefLocalObjectReference: min=0
	LocalObjectReference []ConfigMapRefLocalObjectReference `hcl:"local_object_reference,block" validate:"min=0"`
}

type ConfigMapRefLocalObjectReference struct{}

type SecretRef struct {
	// SecretRefLocalObjectReference: min=0
	LocalObjectReference []SecretRefLocalObjectReference `hcl:"local_object_reference,block" validate:"min=0"`
}

type SecretRefLocalObjectReference struct{}

type Ports struct{}

type Resources struct{}

type VolumeMounts struct{}

type Volumes struct {
	// Secret: min=0
	Secret []Secret `hcl:"secret,block" validate:"min=0"`
}

type Secret struct {
	// Items: min=0
	Items []Items `hcl:"items,block" validate:"min=0"`
}

type Items struct{}

type Traffic struct{}

type MetadataAttributes struct {
	ref terra.Reference
}

func (m MetadataAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MetadataAttributes) InternalWithRef(ref terra.Reference) MetadataAttributes {
	return MetadataAttributes{ref: ref}
}

func (m MetadataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MetadataAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](m.ref.Append("annotations"))
}

func (m MetadataAttributes) Generation() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("generation"))
}

func (m MetadataAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](m.ref.Append("labels"))
}

func (m MetadataAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("namespace"))
}

func (m MetadataAttributes) ResourceVersion() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("resource_version"))
}

func (m MetadataAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("self_link"))
}

func (m MetadataAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("uid"))
}

type StatusAttributes struct {
	ref terra.Reference
}

func (s StatusAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StatusAttributes) InternalWithRef(ref terra.Reference) StatusAttributes {
	return StatusAttributes{ref: ref}
}

func (s StatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StatusAttributes) LatestCreatedRevisionName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("latest_created_revision_name"))
}

func (s StatusAttributes) LatestReadyRevisionName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("latest_ready_revision_name"))
}

func (s StatusAttributes) ObservedGeneration() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("observed_generation"))
}

func (s StatusAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("url"))
}

func (s StatusAttributes) Conditions() terra.ListValue[ConditionsAttributes] {
	return terra.ReferenceAsList[ConditionsAttributes](s.ref.Append("conditions"))
}

type ConditionsAttributes struct {
	ref terra.Reference
}

func (c ConditionsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionsAttributes) InternalWithRef(ref terra.Reference) ConditionsAttributes {
	return ConditionsAttributes{ref: ref}
}

func (c ConditionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionsAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("message"))
}

func (c ConditionsAttributes) Reason() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("reason"))
}

func (c ConditionsAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("status"))
}

func (c ConditionsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
}

type TemplateAttributes struct {
	ref terra.Reference
}

func (t TemplateAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TemplateAttributes) InternalWithRef(ref terra.Reference) TemplateAttributes {
	return TemplateAttributes{ref: ref}
}

func (t TemplateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TemplateAttributes) Metadata() terra.ListValue[TemplateMetadataAttributes] {
	return terra.ReferenceAsList[TemplateMetadataAttributes](t.ref.Append("metadata"))
}

func (t TemplateAttributes) Spec() terra.ListValue[SpecAttributes] {
	return terra.ReferenceAsList[SpecAttributes](t.ref.Append("spec"))
}

type TemplateMetadataAttributes struct {
	ref terra.Reference
}

func (m TemplateMetadataAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m TemplateMetadataAttributes) InternalWithRef(ref terra.Reference) TemplateMetadataAttributes {
	return TemplateMetadataAttributes{ref: ref}
}

func (m TemplateMetadataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m TemplateMetadataAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](m.ref.Append("annotations"))
}

func (m TemplateMetadataAttributes) Generation() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("generation"))
}

func (m TemplateMetadataAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](m.ref.Append("labels"))
}

func (m TemplateMetadataAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("name"))
}

func (m TemplateMetadataAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("namespace"))
}

func (m TemplateMetadataAttributes) ResourceVersion() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("resource_version"))
}

func (m TemplateMetadataAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("self_link"))
}

func (m TemplateMetadataAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("uid"))
}

type SpecAttributes struct {
	ref terra.Reference
}

func (s SpecAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SpecAttributes) InternalWithRef(ref terra.Reference) SpecAttributes {
	return SpecAttributes{ref: ref}
}

func (s SpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SpecAttributes) ContainerConcurrency() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("container_concurrency"))
}

func (s SpecAttributes) ServiceAccountName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("service_account_name"))
}

func (s SpecAttributes) ServingState() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("serving_state"))
}

func (s SpecAttributes) TimeoutSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("timeout_seconds"))
}

func (s SpecAttributes) Containers() terra.ListValue[ContainersAttributes] {
	return terra.ReferenceAsList[ContainersAttributes](s.ref.Append("containers"))
}

func (s SpecAttributes) Volumes() terra.ListValue[VolumesAttributes] {
	return terra.ReferenceAsList[VolumesAttributes](s.ref.Append("volumes"))
}

type ContainersAttributes struct {
	ref terra.Reference
}

func (c ContainersAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ContainersAttributes) InternalWithRef(ref terra.Reference) ContainersAttributes {
	return ContainersAttributes{ref: ref}
}

func (c ContainersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ContainersAttributes) Args() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("args"))
}

func (c ContainersAttributes) Command() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("command"))
}

func (c ContainersAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("image"))
}

func (c ContainersAttributes) WorkingDir() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("working_dir"))
}

func (c ContainersAttributes) Env() terra.SetValue[EnvAttributes] {
	return terra.ReferenceAsSet[EnvAttributes](c.ref.Append("env"))
}

func (c ContainersAttributes) EnvFrom() terra.ListValue[EnvFromAttributes] {
	return terra.ReferenceAsList[EnvFromAttributes](c.ref.Append("env_from"))
}

func (c ContainersAttributes) Ports() terra.ListValue[PortsAttributes] {
	return terra.ReferenceAsList[PortsAttributes](c.ref.Append("ports"))
}

func (c ContainersAttributes) Resources() terra.ListValue[ResourcesAttributes] {
	return terra.ReferenceAsList[ResourcesAttributes](c.ref.Append("resources"))
}

func (c ContainersAttributes) VolumeMounts() terra.ListValue[VolumeMountsAttributes] {
	return terra.ReferenceAsList[VolumeMountsAttributes](c.ref.Append("volume_mounts"))
}

type EnvAttributes struct {
	ref terra.Reference
}

func (e EnvAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EnvAttributes) InternalWithRef(ref terra.Reference) EnvAttributes {
	return EnvAttributes{ref: ref}
}

func (e EnvAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EnvAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("name"))
}

func (e EnvAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("value"))
}

func (e EnvAttributes) ValueFrom() terra.ListValue[ValueFromAttributes] {
	return terra.ReferenceAsList[ValueFromAttributes](e.ref.Append("value_from"))
}

type ValueFromAttributes struct {
	ref terra.Reference
}

func (vf ValueFromAttributes) InternalRef() (terra.Reference, error) {
	return vf.ref, nil
}

func (vf ValueFromAttributes) InternalWithRef(ref terra.Reference) ValueFromAttributes {
	return ValueFromAttributes{ref: ref}
}

func (vf ValueFromAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vf.ref.InternalTokens()
}

func (vf ValueFromAttributes) SecretKeyRef() terra.ListValue[SecretKeyRefAttributes] {
	return terra.ReferenceAsList[SecretKeyRefAttributes](vf.ref.Append("secret_key_ref"))
}

type SecretKeyRefAttributes struct {
	ref terra.Reference
}

func (skr SecretKeyRefAttributes) InternalRef() (terra.Reference, error) {
	return skr.ref, nil
}

func (skr SecretKeyRefAttributes) InternalWithRef(ref terra.Reference) SecretKeyRefAttributes {
	return SecretKeyRefAttributes{ref: ref}
}

func (skr SecretKeyRefAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return skr.ref.InternalTokens()
}

func (skr SecretKeyRefAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(skr.ref.Append("key"))
}

func (skr SecretKeyRefAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(skr.ref.Append("name"))
}

type EnvFromAttributes struct {
	ref terra.Reference
}

func (ef EnvFromAttributes) InternalRef() (terra.Reference, error) {
	return ef.ref, nil
}

func (ef EnvFromAttributes) InternalWithRef(ref terra.Reference) EnvFromAttributes {
	return EnvFromAttributes{ref: ref}
}

func (ef EnvFromAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ef.ref.InternalTokens()
}

func (ef EnvFromAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("prefix"))
}

func (ef EnvFromAttributes) ConfigMapRef() terra.ListValue[ConfigMapRefAttributes] {
	return terra.ReferenceAsList[ConfigMapRefAttributes](ef.ref.Append("config_map_ref"))
}

func (ef EnvFromAttributes) SecretRef() terra.ListValue[SecretRefAttributes] {
	return terra.ReferenceAsList[SecretRefAttributes](ef.ref.Append("secret_ref"))
}

type ConfigMapRefAttributes struct {
	ref terra.Reference
}

func (cmr ConfigMapRefAttributes) InternalRef() (terra.Reference, error) {
	return cmr.ref, nil
}

func (cmr ConfigMapRefAttributes) InternalWithRef(ref terra.Reference) ConfigMapRefAttributes {
	return ConfigMapRefAttributes{ref: ref}
}

func (cmr ConfigMapRefAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cmr.ref.InternalTokens()
}

func (cmr ConfigMapRefAttributes) Optional() terra.BoolValue {
	return terra.ReferenceAsBool(cmr.ref.Append("optional"))
}

func (cmr ConfigMapRefAttributes) LocalObjectReference() terra.ListValue[ConfigMapRefLocalObjectReferenceAttributes] {
	return terra.ReferenceAsList[ConfigMapRefLocalObjectReferenceAttributes](cmr.ref.Append("local_object_reference"))
}

type ConfigMapRefLocalObjectReferenceAttributes struct {
	ref terra.Reference
}

func (lor ConfigMapRefLocalObjectReferenceAttributes) InternalRef() (terra.Reference, error) {
	return lor.ref, nil
}

func (lor ConfigMapRefLocalObjectReferenceAttributes) InternalWithRef(ref terra.Reference) ConfigMapRefLocalObjectReferenceAttributes {
	return ConfigMapRefLocalObjectReferenceAttributes{ref: ref}
}

func (lor ConfigMapRefLocalObjectReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lor.ref.InternalTokens()
}

func (lor ConfigMapRefLocalObjectReferenceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lor.ref.Append("name"))
}

type SecretRefAttributes struct {
	ref terra.Reference
}

func (sr SecretRefAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr SecretRefAttributes) InternalWithRef(ref terra.Reference) SecretRefAttributes {
	return SecretRefAttributes{ref: ref}
}

func (sr SecretRefAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr SecretRefAttributes) Optional() terra.BoolValue {
	return terra.ReferenceAsBool(sr.ref.Append("optional"))
}

func (sr SecretRefAttributes) LocalObjectReference() terra.ListValue[SecretRefLocalObjectReferenceAttributes] {
	return terra.ReferenceAsList[SecretRefLocalObjectReferenceAttributes](sr.ref.Append("local_object_reference"))
}

type SecretRefLocalObjectReferenceAttributes struct {
	ref terra.Reference
}

func (lor SecretRefLocalObjectReferenceAttributes) InternalRef() (terra.Reference, error) {
	return lor.ref, nil
}

func (lor SecretRefLocalObjectReferenceAttributes) InternalWithRef(ref terra.Reference) SecretRefLocalObjectReferenceAttributes {
	return SecretRefLocalObjectReferenceAttributes{ref: ref}
}

func (lor SecretRefLocalObjectReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lor.ref.InternalTokens()
}

func (lor SecretRefLocalObjectReferenceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lor.ref.Append("name"))
}

type PortsAttributes struct {
	ref terra.Reference
}

func (p PortsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PortsAttributes) InternalWithRef(ref terra.Reference) PortsAttributes {
	return PortsAttributes{ref: ref}
}

func (p PortsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PortsAttributes) ContainerPort() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("container_port"))
}

func (p PortsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p PortsAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("protocol"))
}

type ResourcesAttributes struct {
	ref terra.Reference
}

func (r ResourcesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ResourcesAttributes) InternalWithRef(ref terra.Reference) ResourcesAttributes {
	return ResourcesAttributes{ref: ref}
}

func (r ResourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ResourcesAttributes) Limits() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](r.ref.Append("limits"))
}

func (r ResourcesAttributes) Requests() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](r.ref.Append("requests"))
}

type VolumeMountsAttributes struct {
	ref terra.Reference
}

func (vm VolumeMountsAttributes) InternalRef() (terra.Reference, error) {
	return vm.ref, nil
}

func (vm VolumeMountsAttributes) InternalWithRef(ref terra.Reference) VolumeMountsAttributes {
	return VolumeMountsAttributes{ref: ref}
}

func (vm VolumeMountsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vm.ref.InternalTokens()
}

func (vm VolumeMountsAttributes) MountPath() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("mount_path"))
}

func (vm VolumeMountsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("name"))
}

type VolumesAttributes struct {
	ref terra.Reference
}

func (v VolumesAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VolumesAttributes) InternalWithRef(ref terra.Reference) VolumesAttributes {
	return VolumesAttributes{ref: ref}
}

func (v VolumesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v VolumesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("name"))
}

func (v VolumesAttributes) Secret() terra.ListValue[SecretAttributes] {
	return terra.ReferenceAsList[SecretAttributes](v.ref.Append("secret"))
}

type SecretAttributes struct {
	ref terra.Reference
}

func (s SecretAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SecretAttributes) InternalWithRef(ref terra.Reference) SecretAttributes {
	return SecretAttributes{ref: ref}
}

func (s SecretAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SecretAttributes) DefaultMode() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("default_mode"))
}

func (s SecretAttributes) SecretName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("secret_name"))
}

func (s SecretAttributes) Items() terra.ListValue[ItemsAttributes] {
	return terra.ReferenceAsList[ItemsAttributes](s.ref.Append("items"))
}

type ItemsAttributes struct {
	ref terra.Reference
}

func (i ItemsAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i ItemsAttributes) InternalWithRef(ref terra.Reference) ItemsAttributes {
	return ItemsAttributes{ref: ref}
}

func (i ItemsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i ItemsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("key"))
}

func (i ItemsAttributes) Mode() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("mode"))
}

func (i ItemsAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("path"))
}

type TrafficAttributes struct {
	ref terra.Reference
}

func (t TrafficAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TrafficAttributes) InternalWithRef(ref terra.Reference) TrafficAttributes {
	return TrafficAttributes{ref: ref}
}

func (t TrafficAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TrafficAttributes) LatestRevision() terra.BoolValue {
	return terra.ReferenceAsBool(t.ref.Append("latest_revision"))
}

func (t TrafficAttributes) Percent() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("percent"))
}

func (t TrafficAttributes) RevisionName() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("revision_name"))
}

func (t TrafficAttributes) Tag() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("tag"))
}

func (t TrafficAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("url"))
}

type MetadataState struct {
	Annotations     map[string]string `json:"annotations"`
	Generation      float64           `json:"generation"`
	Labels          map[string]string `json:"labels"`
	Namespace       string            `json:"namespace"`
	ResourceVersion string            `json:"resource_version"`
	SelfLink        string            `json:"self_link"`
	Uid             string            `json:"uid"`
}

type StatusState struct {
	LatestCreatedRevisionName string            `json:"latest_created_revision_name"`
	LatestReadyRevisionName   string            `json:"latest_ready_revision_name"`
	ObservedGeneration        float64           `json:"observed_generation"`
	Url                       string            `json:"url"`
	Conditions                []ConditionsState `json:"conditions"`
}

type ConditionsState struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}

type TemplateState struct {
	Metadata []TemplateMetadataState `json:"metadata"`
	Spec     []SpecState             `json:"spec"`
}

type TemplateMetadataState struct {
	Annotations     map[string]string `json:"annotations"`
	Generation      float64           `json:"generation"`
	Labels          map[string]string `json:"labels"`
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace"`
	ResourceVersion string            `json:"resource_version"`
	SelfLink        string            `json:"self_link"`
	Uid             string            `json:"uid"`
}

type SpecState struct {
	ContainerConcurrency float64           `json:"container_concurrency"`
	ServiceAccountName   string            `json:"service_account_name"`
	ServingState         string            `json:"serving_state"`
	TimeoutSeconds       float64           `json:"timeout_seconds"`
	Containers           []ContainersState `json:"containers"`
	Volumes              []VolumesState    `json:"volumes"`
}

type ContainersState struct {
	Args         []string            `json:"args"`
	Command      []string            `json:"command"`
	Image        string              `json:"image"`
	WorkingDir   string              `json:"working_dir"`
	Env          []EnvState          `json:"env"`
	EnvFrom      []EnvFromState      `json:"env_from"`
	Ports        []PortsState        `json:"ports"`
	Resources    []ResourcesState    `json:"resources"`
	VolumeMounts []VolumeMountsState `json:"volume_mounts"`
}

type EnvState struct {
	Name      string           `json:"name"`
	Value     string           `json:"value"`
	ValueFrom []ValueFromState `json:"value_from"`
}

type ValueFromState struct {
	SecretKeyRef []SecretKeyRefState `json:"secret_key_ref"`
}

type SecretKeyRefState struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type EnvFromState struct {
	Prefix       string              `json:"prefix"`
	ConfigMapRef []ConfigMapRefState `json:"config_map_ref"`
	SecretRef    []SecretRefState    `json:"secret_ref"`
}

type ConfigMapRefState struct {
	Optional             bool                                    `json:"optional"`
	LocalObjectReference []ConfigMapRefLocalObjectReferenceState `json:"local_object_reference"`
}

type ConfigMapRefLocalObjectReferenceState struct {
	Name string `json:"name"`
}

type SecretRefState struct {
	Optional             bool                                 `json:"optional"`
	LocalObjectReference []SecretRefLocalObjectReferenceState `json:"local_object_reference"`
}

type SecretRefLocalObjectReferenceState struct {
	Name string `json:"name"`
}

type PortsState struct {
	ContainerPort float64 `json:"container_port"`
	Name          string  `json:"name"`
	Protocol      string  `json:"protocol"`
}

type ResourcesState struct {
	Limits   map[string]string `json:"limits"`
	Requests map[string]string `json:"requests"`
}

type VolumeMountsState struct {
	MountPath string `json:"mount_path"`
	Name      string `json:"name"`
}

type VolumesState struct {
	Name   string        `json:"name"`
	Secret []SecretState `json:"secret"`
}

type SecretState struct {
	DefaultMode float64      `json:"default_mode"`
	SecretName  string       `json:"secret_name"`
	Items       []ItemsState `json:"items"`
}

type ItemsState struct {
	Key  string  `json:"key"`
	Mode float64 `json:"mode"`
	Path string  `json:"path"`
}

type TrafficState struct {
	LatestRevision bool    `json:"latest_revision"`
	Percent        float64 `json:"percent"`
	RevisionName   string  `json:"revision_name"`
	Tag            string  `json:"tag"`
	Url            string  `json:"url"`
}

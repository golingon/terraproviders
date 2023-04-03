// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package containerapp

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Dapr struct {
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// AppPort: number, required
	AppPort terra.NumberValue `hcl:"app_port,attr" validate:"required"`
	// AppProtocol: string, optional
	AppProtocol terra.StringValue `hcl:"app_protocol,attr"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Ingress struct {
	// AllowInsecureConnections: bool, optional
	AllowInsecureConnections terra.BoolValue `hcl:"allow_insecure_connections,attr"`
	// ExternalEnabled: bool, optional
	ExternalEnabled terra.BoolValue `hcl:"external_enabled,attr"`
	// TargetPort: number, required
	TargetPort terra.NumberValue `hcl:"target_port,attr" validate:"required"`
	// Transport: string, optional
	Transport terra.StringValue `hcl:"transport,attr"`
	// CustomDomain: optional
	CustomDomain *CustomDomain `hcl:"custom_domain,block"`
	// TrafficWeight: min=1
	TrafficWeight []TrafficWeight `hcl:"traffic_weight,block" validate:"min=1"`
}

type CustomDomain struct {
	// CertificateBindingType: string, optional
	CertificateBindingType terra.StringValue `hcl:"certificate_binding_type,attr"`
	// CertificateId: string, required
	CertificateId terra.StringValue `hcl:"certificate_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type TrafficWeight struct {
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// LatestRevision: bool, optional
	LatestRevision terra.BoolValue `hcl:"latest_revision,attr"`
	// Percentage: number, required
	Percentage terra.NumberValue `hcl:"percentage,attr" validate:"required"`
	// RevisionSuffix: string, optional
	RevisionSuffix terra.StringValue `hcl:"revision_suffix,attr"`
}

type Registry struct {
	// Identity: string, optional
	Identity terra.StringValue `hcl:"identity,attr"`
	// PasswordSecretName: string, optional
	PasswordSecretName terra.StringValue `hcl:"password_secret_name,attr"`
	// Server: string, required
	Server terra.StringValue `hcl:"server,attr" validate:"required"`
	// Username: string, optional
	Username terra.StringValue `hcl:"username,attr"`
}

type Secret struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type Template struct {
	// MaxReplicas: number, optional
	MaxReplicas terra.NumberValue `hcl:"max_replicas,attr"`
	// MinReplicas: number, optional
	MinReplicas terra.NumberValue `hcl:"min_replicas,attr"`
	// RevisionSuffix: string, optional
	RevisionSuffix terra.StringValue `hcl:"revision_suffix,attr"`
	// Container: min=1
	Container []Container `hcl:"container,block" validate:"min=1"`
	// Volume: min=0
	Volume []Volume `hcl:"volume,block" validate:"min=0"`
}

type Container struct {
	// Args: list of string, optional
	Args terra.ListValue[terra.StringValue] `hcl:"args,attr"`
	// Command: list of string, optional
	Command terra.ListValue[terra.StringValue] `hcl:"command,attr"`
	// Cpu: number, required
	Cpu terra.NumberValue `hcl:"cpu,attr" validate:"required"`
	// Image: string, required
	Image terra.StringValue `hcl:"image,attr" validate:"required"`
	// Memory: string, required
	Memory terra.StringValue `hcl:"memory,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Env: min=0
	Env []Env `hcl:"env,block" validate:"min=0"`
	// LivenessProbe: min=0
	LivenessProbe []LivenessProbe `hcl:"liveness_probe,block" validate:"min=0"`
	// ReadinessProbe: min=0
	ReadinessProbe []ReadinessProbe `hcl:"readiness_probe,block" validate:"min=0"`
	// StartupProbe: min=0
	StartupProbe []StartupProbe `hcl:"startup_probe,block" validate:"min=0"`
	// VolumeMounts: min=0
	VolumeMounts []VolumeMounts `hcl:"volume_mounts,block" validate:"min=0"`
}

type Env struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SecretName: string, optional
	SecretName terra.StringValue `hcl:"secret_name,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type LivenessProbe struct {
	// FailureCountThreshold: number, optional
	FailureCountThreshold terra.NumberValue `hcl:"failure_count_threshold,attr"`
	// Host: string, optional
	Host terra.StringValue `hcl:"host,attr"`
	// InitialDelay: number, optional
	InitialDelay terra.NumberValue `hcl:"initial_delay,attr"`
	// IntervalSeconds: number, optional
	IntervalSeconds terra.NumberValue `hcl:"interval_seconds,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// Timeout: number, optional
	Timeout terra.NumberValue `hcl:"timeout,attr"`
	// Transport: string, required
	Transport terra.StringValue `hcl:"transport,attr" validate:"required"`
	// LivenessProbeHeader: min=0
	Header []LivenessProbeHeader `hcl:"header,block" validate:"min=0"`
}

type LivenessProbeHeader struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type ReadinessProbe struct {
	// FailureCountThreshold: number, optional
	FailureCountThreshold terra.NumberValue `hcl:"failure_count_threshold,attr"`
	// Host: string, optional
	Host terra.StringValue `hcl:"host,attr"`
	// IntervalSeconds: number, optional
	IntervalSeconds terra.NumberValue `hcl:"interval_seconds,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// SuccessCountThreshold: number, optional
	SuccessCountThreshold terra.NumberValue `hcl:"success_count_threshold,attr"`
	// Timeout: number, optional
	Timeout terra.NumberValue `hcl:"timeout,attr"`
	// Transport: string, required
	Transport terra.StringValue `hcl:"transport,attr" validate:"required"`
	// ReadinessProbeHeader: min=0
	Header []ReadinessProbeHeader `hcl:"header,block" validate:"min=0"`
}

type ReadinessProbeHeader struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type StartupProbe struct {
	// FailureCountThreshold: number, optional
	FailureCountThreshold terra.NumberValue `hcl:"failure_count_threshold,attr"`
	// Host: string, optional
	Host terra.StringValue `hcl:"host,attr"`
	// IntervalSeconds: number, optional
	IntervalSeconds terra.NumberValue `hcl:"interval_seconds,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// Timeout: number, optional
	Timeout terra.NumberValue `hcl:"timeout,attr"`
	// Transport: string, required
	Transport terra.StringValue `hcl:"transport,attr" validate:"required"`
	// StartupProbeHeader: min=0
	Header []StartupProbeHeader `hcl:"header,block" validate:"min=0"`
}

type StartupProbeHeader struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type VolumeMounts struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
}

type Volume struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageName: string, optional
	StorageName terra.StringValue `hcl:"storage_name,attr"`
	// StorageType: string, optional
	StorageType terra.StringValue `hcl:"storage_type,attr"`
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

type DaprAttributes struct {
	ref terra.Reference
}

func (d DaprAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DaprAttributes) InternalWithRef(ref terra.Reference) DaprAttributes {
	return DaprAttributes{ref: ref}
}

func (d DaprAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DaprAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("app_id"))
}

func (d DaprAttributes) AppPort() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("app_port"))
}

func (d DaprAttributes) AppProtocol() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("app_protocol"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type IngressAttributes struct {
	ref terra.Reference
}

func (i IngressAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IngressAttributes) InternalWithRef(ref terra.Reference) IngressAttributes {
	return IngressAttributes{ref: ref}
}

func (i IngressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IngressAttributes) AllowInsecureConnections() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("allow_insecure_connections"))
}

func (i IngressAttributes) ExternalEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("external_enabled"))
}

func (i IngressAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("fqdn"))
}

func (i IngressAttributes) TargetPort() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("target_port"))
}

func (i IngressAttributes) Transport() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("transport"))
}

func (i IngressAttributes) CustomDomain() terra.ListValue[CustomDomainAttributes] {
	return terra.ReferenceAsList[CustomDomainAttributes](i.ref.Append("custom_domain"))
}

func (i IngressAttributes) TrafficWeight() terra.ListValue[TrafficWeightAttributes] {
	return terra.ReferenceAsList[TrafficWeightAttributes](i.ref.Append("traffic_weight"))
}

type CustomDomainAttributes struct {
	ref terra.Reference
}

func (cd CustomDomainAttributes) InternalRef() (terra.Reference, error) {
	return cd.ref, nil
}

func (cd CustomDomainAttributes) InternalWithRef(ref terra.Reference) CustomDomainAttributes {
	return CustomDomainAttributes{ref: ref}
}

func (cd CustomDomainAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cd.ref.InternalTokens()
}

func (cd CustomDomainAttributes) CertificateBindingType() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("certificate_binding_type"))
}

func (cd CustomDomainAttributes) CertificateId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("certificate_id"))
}

func (cd CustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("name"))
}

type TrafficWeightAttributes struct {
	ref terra.Reference
}

func (tw TrafficWeightAttributes) InternalRef() (terra.Reference, error) {
	return tw.ref, nil
}

func (tw TrafficWeightAttributes) InternalWithRef(ref terra.Reference) TrafficWeightAttributes {
	return TrafficWeightAttributes{ref: ref}
}

func (tw TrafficWeightAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tw.ref.InternalTokens()
}

func (tw TrafficWeightAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(tw.ref.Append("label"))
}

func (tw TrafficWeightAttributes) LatestRevision() terra.BoolValue {
	return terra.ReferenceAsBool(tw.ref.Append("latest_revision"))
}

func (tw TrafficWeightAttributes) Percentage() terra.NumberValue {
	return terra.ReferenceAsNumber(tw.ref.Append("percentage"))
}

func (tw TrafficWeightAttributes) RevisionSuffix() terra.StringValue {
	return terra.ReferenceAsString(tw.ref.Append("revision_suffix"))
}

type RegistryAttributes struct {
	ref terra.Reference
}

func (r RegistryAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RegistryAttributes) InternalWithRef(ref terra.Reference) RegistryAttributes {
	return RegistryAttributes{ref: ref}
}

func (r RegistryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RegistryAttributes) Identity() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("identity"))
}

func (r RegistryAttributes) PasswordSecretName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("password_secret_name"))
}

func (r RegistryAttributes) Server() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("server"))
}

func (r RegistryAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("username"))
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

func (s SecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SecretAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("value"))
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

func (t TemplateAttributes) MaxReplicas() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("max_replicas"))
}

func (t TemplateAttributes) MinReplicas() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("min_replicas"))
}

func (t TemplateAttributes) RevisionSuffix() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("revision_suffix"))
}

func (t TemplateAttributes) Container() terra.ListValue[ContainerAttributes] {
	return terra.ReferenceAsList[ContainerAttributes](t.ref.Append("container"))
}

func (t TemplateAttributes) Volume() terra.ListValue[VolumeAttributes] {
	return terra.ReferenceAsList[VolumeAttributes](t.ref.Append("volume"))
}

type ContainerAttributes struct {
	ref terra.Reference
}

func (c ContainerAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ContainerAttributes) InternalWithRef(ref terra.Reference) ContainerAttributes {
	return ContainerAttributes{ref: ref}
}

func (c ContainerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ContainerAttributes) Args() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("args"))
}

func (c ContainerAttributes) Command() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("command"))
}

func (c ContainerAttributes) Cpu() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("cpu"))
}

func (c ContainerAttributes) EphemeralStorage() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("ephemeral_storage"))
}

func (c ContainerAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("image"))
}

func (c ContainerAttributes) Memory() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("memory"))
}

func (c ContainerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ContainerAttributes) Env() terra.ListValue[EnvAttributes] {
	return terra.ReferenceAsList[EnvAttributes](c.ref.Append("env"))
}

func (c ContainerAttributes) LivenessProbe() terra.ListValue[LivenessProbeAttributes] {
	return terra.ReferenceAsList[LivenessProbeAttributes](c.ref.Append("liveness_probe"))
}

func (c ContainerAttributes) ReadinessProbe() terra.ListValue[ReadinessProbeAttributes] {
	return terra.ReferenceAsList[ReadinessProbeAttributes](c.ref.Append("readiness_probe"))
}

func (c ContainerAttributes) StartupProbe() terra.ListValue[StartupProbeAttributes] {
	return terra.ReferenceAsList[StartupProbeAttributes](c.ref.Append("startup_probe"))
}

func (c ContainerAttributes) VolumeMounts() terra.ListValue[VolumeMountsAttributes] {
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

func (e EnvAttributes) SecretName() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("secret_name"))
}

func (e EnvAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("value"))
}

type LivenessProbeAttributes struct {
	ref terra.Reference
}

func (lp LivenessProbeAttributes) InternalRef() (terra.Reference, error) {
	return lp.ref, nil
}

func (lp LivenessProbeAttributes) InternalWithRef(ref terra.Reference) LivenessProbeAttributes {
	return LivenessProbeAttributes{ref: ref}
}

func (lp LivenessProbeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lp.ref.InternalTokens()
}

func (lp LivenessProbeAttributes) FailureCountThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(lp.ref.Append("failure_count_threshold"))
}

func (lp LivenessProbeAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("host"))
}

func (lp LivenessProbeAttributes) InitialDelay() terra.NumberValue {
	return terra.ReferenceAsNumber(lp.ref.Append("initial_delay"))
}

func (lp LivenessProbeAttributes) IntervalSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(lp.ref.Append("interval_seconds"))
}

func (lp LivenessProbeAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("path"))
}

func (lp LivenessProbeAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(lp.ref.Append("port"))
}

func (lp LivenessProbeAttributes) TerminationGracePeriodSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(lp.ref.Append("termination_grace_period_seconds"))
}

func (lp LivenessProbeAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceAsNumber(lp.ref.Append("timeout"))
}

func (lp LivenessProbeAttributes) Transport() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("transport"))
}

func (lp LivenessProbeAttributes) Header() terra.ListValue[LivenessProbeHeaderAttributes] {
	return terra.ReferenceAsList[LivenessProbeHeaderAttributes](lp.ref.Append("header"))
}

type LivenessProbeHeaderAttributes struct {
	ref terra.Reference
}

func (h LivenessProbeHeaderAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h LivenessProbeHeaderAttributes) InternalWithRef(ref terra.Reference) LivenessProbeHeaderAttributes {
	return LivenessProbeHeaderAttributes{ref: ref}
}

func (h LivenessProbeHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h LivenessProbeHeaderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("name"))
}

func (h LivenessProbeHeaderAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
}

type ReadinessProbeAttributes struct {
	ref terra.Reference
}

func (rp ReadinessProbeAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp ReadinessProbeAttributes) InternalWithRef(ref terra.Reference) ReadinessProbeAttributes {
	return ReadinessProbeAttributes{ref: ref}
}

func (rp ReadinessProbeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp ReadinessProbeAttributes) FailureCountThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("failure_count_threshold"))
}

func (rp ReadinessProbeAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("host"))
}

func (rp ReadinessProbeAttributes) IntervalSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("interval_seconds"))
}

func (rp ReadinessProbeAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("path"))
}

func (rp ReadinessProbeAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("port"))
}

func (rp ReadinessProbeAttributes) SuccessCountThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("success_count_threshold"))
}

func (rp ReadinessProbeAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("timeout"))
}

func (rp ReadinessProbeAttributes) Transport() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("transport"))
}

func (rp ReadinessProbeAttributes) Header() terra.ListValue[ReadinessProbeHeaderAttributes] {
	return terra.ReferenceAsList[ReadinessProbeHeaderAttributes](rp.ref.Append("header"))
}

type ReadinessProbeHeaderAttributes struct {
	ref terra.Reference
}

func (h ReadinessProbeHeaderAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h ReadinessProbeHeaderAttributes) InternalWithRef(ref terra.Reference) ReadinessProbeHeaderAttributes {
	return ReadinessProbeHeaderAttributes{ref: ref}
}

func (h ReadinessProbeHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h ReadinessProbeHeaderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("name"))
}

func (h ReadinessProbeHeaderAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
}

type StartupProbeAttributes struct {
	ref terra.Reference
}

func (sp StartupProbeAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp StartupProbeAttributes) InternalWithRef(ref terra.Reference) StartupProbeAttributes {
	return StartupProbeAttributes{ref: ref}
}

func (sp StartupProbeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp StartupProbeAttributes) FailureCountThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("failure_count_threshold"))
}

func (sp StartupProbeAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("host"))
}

func (sp StartupProbeAttributes) IntervalSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("interval_seconds"))
}

func (sp StartupProbeAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("path"))
}

func (sp StartupProbeAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("port"))
}

func (sp StartupProbeAttributes) TerminationGracePeriodSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("termination_grace_period_seconds"))
}

func (sp StartupProbeAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("timeout"))
}

func (sp StartupProbeAttributes) Transport() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("transport"))
}

func (sp StartupProbeAttributes) Header() terra.ListValue[StartupProbeHeaderAttributes] {
	return terra.ReferenceAsList[StartupProbeHeaderAttributes](sp.ref.Append("header"))
}

type StartupProbeHeaderAttributes struct {
	ref terra.Reference
}

func (h StartupProbeHeaderAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h StartupProbeHeaderAttributes) InternalWithRef(ref terra.Reference) StartupProbeHeaderAttributes {
	return StartupProbeHeaderAttributes{ref: ref}
}

func (h StartupProbeHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h StartupProbeHeaderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("name"))
}

func (h StartupProbeHeaderAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
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

func (vm VolumeMountsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("name"))
}

func (vm VolumeMountsAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(vm.ref.Append("path"))
}

type VolumeAttributes struct {
	ref terra.Reference
}

func (v VolumeAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VolumeAttributes) InternalWithRef(ref terra.Reference) VolumeAttributes {
	return VolumeAttributes{ref: ref}
}

func (v VolumeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v VolumeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("name"))
}

func (v VolumeAttributes) StorageName() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("storage_name"))
}

func (v VolumeAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("storage_type"))
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

type DaprState struct {
	AppId       string  `json:"app_id"`
	AppPort     float64 `json:"app_port"`
	AppProtocol string  `json:"app_protocol"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type IngressState struct {
	AllowInsecureConnections bool                 `json:"allow_insecure_connections"`
	ExternalEnabled          bool                 `json:"external_enabled"`
	Fqdn                     string               `json:"fqdn"`
	TargetPort               float64              `json:"target_port"`
	Transport                string               `json:"transport"`
	CustomDomain             []CustomDomainState  `json:"custom_domain"`
	TrafficWeight            []TrafficWeightState `json:"traffic_weight"`
}

type CustomDomainState struct {
	CertificateBindingType string `json:"certificate_binding_type"`
	CertificateId          string `json:"certificate_id"`
	Name                   string `json:"name"`
}

type TrafficWeightState struct {
	Label          string  `json:"label"`
	LatestRevision bool    `json:"latest_revision"`
	Percentage     float64 `json:"percentage"`
	RevisionSuffix string  `json:"revision_suffix"`
}

type RegistryState struct {
	Identity           string `json:"identity"`
	PasswordSecretName string `json:"password_secret_name"`
	Server             string `json:"server"`
	Username           string `json:"username"`
}

type SecretState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TemplateState struct {
	MaxReplicas    float64          `json:"max_replicas"`
	MinReplicas    float64          `json:"min_replicas"`
	RevisionSuffix string           `json:"revision_suffix"`
	Container      []ContainerState `json:"container"`
	Volume         []VolumeState    `json:"volume"`
}

type ContainerState struct {
	Args             []string              `json:"args"`
	Command          []string              `json:"command"`
	Cpu              float64               `json:"cpu"`
	EphemeralStorage string                `json:"ephemeral_storage"`
	Image            string                `json:"image"`
	Memory           string                `json:"memory"`
	Name             string                `json:"name"`
	Env              []EnvState            `json:"env"`
	LivenessProbe    []LivenessProbeState  `json:"liveness_probe"`
	ReadinessProbe   []ReadinessProbeState `json:"readiness_probe"`
	StartupProbe     []StartupProbeState   `json:"startup_probe"`
	VolumeMounts     []VolumeMountsState   `json:"volume_mounts"`
}

type EnvState struct {
	Name       string `json:"name"`
	SecretName string `json:"secret_name"`
	Value      string `json:"value"`
}

type LivenessProbeState struct {
	FailureCountThreshold         float64                    `json:"failure_count_threshold"`
	Host                          string                     `json:"host"`
	InitialDelay                  float64                    `json:"initial_delay"`
	IntervalSeconds               float64                    `json:"interval_seconds"`
	Path                          string                     `json:"path"`
	Port                          float64                    `json:"port"`
	TerminationGracePeriodSeconds float64                    `json:"termination_grace_period_seconds"`
	Timeout                       float64                    `json:"timeout"`
	Transport                     string                     `json:"transport"`
	Header                        []LivenessProbeHeaderState `json:"header"`
}

type LivenessProbeHeaderState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ReadinessProbeState struct {
	FailureCountThreshold float64                     `json:"failure_count_threshold"`
	Host                  string                      `json:"host"`
	IntervalSeconds       float64                     `json:"interval_seconds"`
	Path                  string                      `json:"path"`
	Port                  float64                     `json:"port"`
	SuccessCountThreshold float64                     `json:"success_count_threshold"`
	Timeout               float64                     `json:"timeout"`
	Transport             string                      `json:"transport"`
	Header                []ReadinessProbeHeaderState `json:"header"`
}

type ReadinessProbeHeaderState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type StartupProbeState struct {
	FailureCountThreshold         float64                   `json:"failure_count_threshold"`
	Host                          string                    `json:"host"`
	IntervalSeconds               float64                   `json:"interval_seconds"`
	Path                          string                    `json:"path"`
	Port                          float64                   `json:"port"`
	TerminationGracePeriodSeconds float64                   `json:"termination_grace_period_seconds"`
	Timeout                       float64                   `json:"timeout"`
	Transport                     string                    `json:"transport"`
	Header                        []StartupProbeHeaderState `json:"header"`
}

type StartupProbeHeaderState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type VolumeMountsState struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type VolumeState struct {
	Name        string `json:"name"`
	StorageName string `json:"storage_name"`
	StorageType string `json:"storage_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
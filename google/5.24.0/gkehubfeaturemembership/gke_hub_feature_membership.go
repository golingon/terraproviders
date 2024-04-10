// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package gkehubfeaturemembership

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Configmanagement struct {
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// Binauthz: optional
	Binauthz *Binauthz `hcl:"binauthz,block"`
	// ConfigSync: optional
	ConfigSync *ConfigSync `hcl:"config_sync,block"`
	// HierarchyController: optional
	HierarchyController *HierarchyController `hcl:"hierarchy_controller,block"`
	// PolicyController: optional
	PolicyController *PolicyController `hcl:"policy_controller,block"`
}

type Binauthz struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type ConfigSync struct {
	// MetricsGcpServiceAccountEmail: string, optional
	MetricsGcpServiceAccountEmail terra.StringValue `hcl:"metrics_gcp_service_account_email,attr"`
	// PreventDrift: bool, optional
	PreventDrift terra.BoolValue `hcl:"prevent_drift,attr"`
	// SourceFormat: string, optional
	SourceFormat terra.StringValue `hcl:"source_format,attr"`
	// Git: optional
	Git *Git `hcl:"git,block"`
	// Oci: optional
	Oci *Oci `hcl:"oci,block"`
}

type Git struct {
	// GcpServiceAccountEmail: string, optional
	GcpServiceAccountEmail terra.StringValue `hcl:"gcp_service_account_email,attr"`
	// HttpsProxy: string, optional
	HttpsProxy terra.StringValue `hcl:"https_proxy,attr"`
	// PolicyDir: string, optional
	PolicyDir terra.StringValue `hcl:"policy_dir,attr"`
	// SecretType: string, optional
	SecretType terra.StringValue `hcl:"secret_type,attr"`
	// SyncBranch: string, optional
	SyncBranch terra.StringValue `hcl:"sync_branch,attr"`
	// SyncRepo: string, optional
	SyncRepo terra.StringValue `hcl:"sync_repo,attr"`
	// SyncRev: string, optional
	SyncRev terra.StringValue `hcl:"sync_rev,attr"`
	// SyncWaitSecs: string, optional
	SyncWaitSecs terra.StringValue `hcl:"sync_wait_secs,attr"`
}

type Oci struct {
	// GcpServiceAccountEmail: string, optional
	GcpServiceAccountEmail terra.StringValue `hcl:"gcp_service_account_email,attr"`
	// PolicyDir: string, optional
	PolicyDir terra.StringValue `hcl:"policy_dir,attr"`
	// SecretType: string, optional
	SecretType terra.StringValue `hcl:"secret_type,attr"`
	// SyncRepo: string, optional
	SyncRepo terra.StringValue `hcl:"sync_repo,attr"`
	// SyncWaitSecs: string, optional
	SyncWaitSecs terra.StringValue `hcl:"sync_wait_secs,attr"`
}

type HierarchyController struct {
	// EnableHierarchicalResourceQuota: bool, optional
	EnableHierarchicalResourceQuota terra.BoolValue `hcl:"enable_hierarchical_resource_quota,attr"`
	// EnablePodTreeLabels: bool, optional
	EnablePodTreeLabels terra.BoolValue `hcl:"enable_pod_tree_labels,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type PolicyController struct {
	// AuditIntervalSeconds: string, optional
	AuditIntervalSeconds terra.StringValue `hcl:"audit_interval_seconds,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// ExemptableNamespaces: list of string, optional
	ExemptableNamespaces terra.ListValue[terra.StringValue] `hcl:"exemptable_namespaces,attr"`
	// LogDeniesEnabled: bool, optional
	LogDeniesEnabled terra.BoolValue `hcl:"log_denies_enabled,attr"`
	// MutationEnabled: bool, optional
	MutationEnabled terra.BoolValue `hcl:"mutation_enabled,attr"`
	// ReferentialRulesEnabled: bool, optional
	ReferentialRulesEnabled terra.BoolValue `hcl:"referential_rules_enabled,attr"`
	// TemplateLibraryInstalled: bool, optional
	TemplateLibraryInstalled terra.BoolValue `hcl:"template_library_installed,attr"`
	// PolicyControllerMonitoring: optional
	Monitoring *PolicyControllerMonitoring `hcl:"monitoring,block"`
}

type PolicyControllerMonitoring struct {
	// Backends: list of string, optional
	Backends terra.ListValue[terra.StringValue] `hcl:"backends,attr"`
}

type Mesh struct {
	// ControlPlane: string, optional
	ControlPlane terra.StringValue `hcl:"control_plane,attr"`
	// Management: string, optional
	Management terra.StringValue `hcl:"management,attr"`
}

type Policycontroller struct {
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// PolicyControllerHubConfig: required
	PolicyControllerHubConfig *PolicyControllerHubConfig `hcl:"policy_controller_hub_config,block" validate:"required"`
}

type PolicyControllerHubConfig struct {
	// AuditIntervalSeconds: number, optional
	AuditIntervalSeconds terra.NumberValue `hcl:"audit_interval_seconds,attr"`
	// ConstraintViolationLimit: number, optional
	ConstraintViolationLimit terra.NumberValue `hcl:"constraint_violation_limit,attr"`
	// ExemptableNamespaces: list of string, optional
	ExemptableNamespaces terra.ListValue[terra.StringValue] `hcl:"exemptable_namespaces,attr"`
	// InstallSpec: string, optional
	InstallSpec terra.StringValue `hcl:"install_spec,attr"`
	// LogDeniesEnabled: bool, optional
	LogDeniesEnabled terra.BoolValue `hcl:"log_denies_enabled,attr"`
	// MutationEnabled: bool, optional
	MutationEnabled terra.BoolValue `hcl:"mutation_enabled,attr"`
	// ReferentialRulesEnabled: bool, optional
	ReferentialRulesEnabled terra.BoolValue `hcl:"referential_rules_enabled,attr"`
	// DeploymentConfigs: min=0
	DeploymentConfigs []DeploymentConfigs `hcl:"deployment_configs,block" validate:"min=0"`
	// PolicyControllerHubConfigMonitoring: optional
	Monitoring *PolicyControllerHubConfigMonitoring `hcl:"monitoring,block"`
	// PolicyContent: optional
	PolicyContent *PolicyContent `hcl:"policy_content,block"`
}

type DeploymentConfigs struct {
	// ComponentName: string, required
	ComponentName terra.StringValue `hcl:"component_name,attr" validate:"required"`
	// PodAffinity: string, optional
	PodAffinity terra.StringValue `hcl:"pod_affinity,attr"`
	// ReplicaCount: number, optional
	ReplicaCount terra.NumberValue `hcl:"replica_count,attr"`
	// ContainerResources: optional
	ContainerResources *ContainerResources `hcl:"container_resources,block"`
	// PodTolerations: min=0
	PodTolerations []PodTolerations `hcl:"pod_tolerations,block" validate:"min=0"`
}

type ContainerResources struct {
	// Limits: optional
	Limits *Limits `hcl:"limits,block"`
	// Requests: optional
	Requests *Requests `hcl:"requests,block"`
}

type Limits struct {
	// Cpu: string, optional
	Cpu terra.StringValue `hcl:"cpu,attr"`
	// Memory: string, optional
	Memory terra.StringValue `hcl:"memory,attr"`
}

type Requests struct {
	// Cpu: string, optional
	Cpu terra.StringValue `hcl:"cpu,attr"`
	// Memory: string, optional
	Memory terra.StringValue `hcl:"memory,attr"`
}

type PodTolerations struct {
	// Effect: string, optional
	Effect terra.StringValue `hcl:"effect,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Operator: string, optional
	Operator terra.StringValue `hcl:"operator,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type PolicyControllerHubConfigMonitoring struct {
	// Backends: list of string, optional
	Backends terra.ListValue[terra.StringValue] `hcl:"backends,attr"`
}

type PolicyContent struct {
	// Bundles: min=0
	Bundles []Bundles `hcl:"bundles,block" validate:"min=0"`
	// TemplateLibrary: optional
	TemplateLibrary *TemplateLibrary `hcl:"template_library,block"`
}

type Bundles struct {
	// BundleName: string, required
	BundleName terra.StringValue `hcl:"bundle_name,attr" validate:"required"`
	// ExemptedNamespaces: list of string, optional
	ExemptedNamespaces terra.ListValue[terra.StringValue] `hcl:"exempted_namespaces,attr"`
}

type TemplateLibrary struct {
	// Installation: string, optional
	Installation terra.StringValue `hcl:"installation,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ConfigmanagementAttributes struct {
	ref terra.Reference
}

func (c ConfigmanagementAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigmanagementAttributes) InternalWithRef(ref terra.Reference) ConfigmanagementAttributes {
	return ConfigmanagementAttributes{ref: ref}
}

func (c ConfigmanagementAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigmanagementAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("version"))
}

func (c ConfigmanagementAttributes) Binauthz() terra.ListValue[BinauthzAttributes] {
	return terra.ReferenceAsList[BinauthzAttributes](c.ref.Append("binauthz"))
}

func (c ConfigmanagementAttributes) ConfigSync() terra.ListValue[ConfigSyncAttributes] {
	return terra.ReferenceAsList[ConfigSyncAttributes](c.ref.Append("config_sync"))
}

func (c ConfigmanagementAttributes) HierarchyController() terra.ListValue[HierarchyControllerAttributes] {
	return terra.ReferenceAsList[HierarchyControllerAttributes](c.ref.Append("hierarchy_controller"))
}

func (c ConfigmanagementAttributes) PolicyController() terra.ListValue[PolicyControllerAttributes] {
	return terra.ReferenceAsList[PolicyControllerAttributes](c.ref.Append("policy_controller"))
}

type BinauthzAttributes struct {
	ref terra.Reference
}

func (b BinauthzAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BinauthzAttributes) InternalWithRef(ref terra.Reference) BinauthzAttributes {
	return BinauthzAttributes{ref: ref}
}

func (b BinauthzAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BinauthzAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(b.ref.Append("enabled"))
}

type ConfigSyncAttributes struct {
	ref terra.Reference
}

func (cs ConfigSyncAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs ConfigSyncAttributes) InternalWithRef(ref terra.Reference) ConfigSyncAttributes {
	return ConfigSyncAttributes{ref: ref}
}

func (cs ConfigSyncAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs ConfigSyncAttributes) MetricsGcpServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("metrics_gcp_service_account_email"))
}

func (cs ConfigSyncAttributes) PreventDrift() terra.BoolValue {
	return terra.ReferenceAsBool(cs.ref.Append("prevent_drift"))
}

func (cs ConfigSyncAttributes) SourceFormat() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("source_format"))
}

func (cs ConfigSyncAttributes) Git() terra.ListValue[GitAttributes] {
	return terra.ReferenceAsList[GitAttributes](cs.ref.Append("git"))
}

func (cs ConfigSyncAttributes) Oci() terra.ListValue[OciAttributes] {
	return terra.ReferenceAsList[OciAttributes](cs.ref.Append("oci"))
}

type GitAttributes struct {
	ref terra.Reference
}

func (g GitAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GitAttributes) InternalWithRef(ref terra.Reference) GitAttributes {
	return GitAttributes{ref: ref}
}

func (g GitAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GitAttributes) GcpServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("gcp_service_account_email"))
}

func (g GitAttributes) HttpsProxy() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("https_proxy"))
}

func (g GitAttributes) PolicyDir() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("policy_dir"))
}

func (g GitAttributes) SecretType() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("secret_type"))
}

func (g GitAttributes) SyncBranch() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("sync_branch"))
}

func (g GitAttributes) SyncRepo() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("sync_repo"))
}

func (g GitAttributes) SyncRev() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("sync_rev"))
}

func (g GitAttributes) SyncWaitSecs() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("sync_wait_secs"))
}

type OciAttributes struct {
	ref terra.Reference
}

func (o OciAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OciAttributes) InternalWithRef(ref terra.Reference) OciAttributes {
	return OciAttributes{ref: ref}
}

func (o OciAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OciAttributes) GcpServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("gcp_service_account_email"))
}

func (o OciAttributes) PolicyDir() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("policy_dir"))
}

func (o OciAttributes) SecretType() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("secret_type"))
}

func (o OciAttributes) SyncRepo() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("sync_repo"))
}

func (o OciAttributes) SyncWaitSecs() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("sync_wait_secs"))
}

type HierarchyControllerAttributes struct {
	ref terra.Reference
}

func (hc HierarchyControllerAttributes) InternalRef() (terra.Reference, error) {
	return hc.ref, nil
}

func (hc HierarchyControllerAttributes) InternalWithRef(ref terra.Reference) HierarchyControllerAttributes {
	return HierarchyControllerAttributes{ref: ref}
}

func (hc HierarchyControllerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hc.ref.InternalTokens()
}

func (hc HierarchyControllerAttributes) EnableHierarchicalResourceQuota() terra.BoolValue {
	return terra.ReferenceAsBool(hc.ref.Append("enable_hierarchical_resource_quota"))
}

func (hc HierarchyControllerAttributes) EnablePodTreeLabels() terra.BoolValue {
	return terra.ReferenceAsBool(hc.ref.Append("enable_pod_tree_labels"))
}

func (hc HierarchyControllerAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(hc.ref.Append("enabled"))
}

type PolicyControllerAttributes struct {
	ref terra.Reference
}

func (pc PolicyControllerAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PolicyControllerAttributes) InternalWithRef(ref terra.Reference) PolicyControllerAttributes {
	return PolicyControllerAttributes{ref: ref}
}

func (pc PolicyControllerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PolicyControllerAttributes) AuditIntervalSeconds() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("audit_interval_seconds"))
}

func (pc PolicyControllerAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(pc.ref.Append("enabled"))
}

func (pc PolicyControllerAttributes) ExemptableNamespaces() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("exemptable_namespaces"))
}

func (pc PolicyControllerAttributes) LogDeniesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pc.ref.Append("log_denies_enabled"))
}

func (pc PolicyControllerAttributes) MutationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pc.ref.Append("mutation_enabled"))
}

func (pc PolicyControllerAttributes) ReferentialRulesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pc.ref.Append("referential_rules_enabled"))
}

func (pc PolicyControllerAttributes) TemplateLibraryInstalled() terra.BoolValue {
	return terra.ReferenceAsBool(pc.ref.Append("template_library_installed"))
}

func (pc PolicyControllerAttributes) Monitoring() terra.ListValue[PolicyControllerMonitoringAttributes] {
	return terra.ReferenceAsList[PolicyControllerMonitoringAttributes](pc.ref.Append("monitoring"))
}

type PolicyControllerMonitoringAttributes struct {
	ref terra.Reference
}

func (m PolicyControllerMonitoringAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m PolicyControllerMonitoringAttributes) InternalWithRef(ref terra.Reference) PolicyControllerMonitoringAttributes {
	return PolicyControllerMonitoringAttributes{ref: ref}
}

func (m PolicyControllerMonitoringAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m PolicyControllerMonitoringAttributes) Backends() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](m.ref.Append("backends"))
}

type MeshAttributes struct {
	ref terra.Reference
}

func (m MeshAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MeshAttributes) InternalWithRef(ref terra.Reference) MeshAttributes {
	return MeshAttributes{ref: ref}
}

func (m MeshAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MeshAttributes) ControlPlane() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("control_plane"))
}

func (m MeshAttributes) Management() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("management"))
}

type PolicycontrollerAttributes struct {
	ref terra.Reference
}

func (p PolicycontrollerAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PolicycontrollerAttributes) InternalWithRef(ref terra.Reference) PolicycontrollerAttributes {
	return PolicycontrollerAttributes{ref: ref}
}

func (p PolicycontrollerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PolicycontrollerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("version"))
}

func (p PolicycontrollerAttributes) PolicyControllerHubConfig() terra.ListValue[PolicyControllerHubConfigAttributes] {
	return terra.ReferenceAsList[PolicyControllerHubConfigAttributes](p.ref.Append("policy_controller_hub_config"))
}

type PolicyControllerHubConfigAttributes struct {
	ref terra.Reference
}

func (pchc PolicyControllerHubConfigAttributes) InternalRef() (terra.Reference, error) {
	return pchc.ref, nil
}

func (pchc PolicyControllerHubConfigAttributes) InternalWithRef(ref terra.Reference) PolicyControllerHubConfigAttributes {
	return PolicyControllerHubConfigAttributes{ref: ref}
}

func (pchc PolicyControllerHubConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pchc.ref.InternalTokens()
}

func (pchc PolicyControllerHubConfigAttributes) AuditIntervalSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(pchc.ref.Append("audit_interval_seconds"))
}

func (pchc PolicyControllerHubConfigAttributes) ConstraintViolationLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(pchc.ref.Append("constraint_violation_limit"))
}

func (pchc PolicyControllerHubConfigAttributes) ExemptableNamespaces() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pchc.ref.Append("exemptable_namespaces"))
}

func (pchc PolicyControllerHubConfigAttributes) InstallSpec() terra.StringValue {
	return terra.ReferenceAsString(pchc.ref.Append("install_spec"))
}

func (pchc PolicyControllerHubConfigAttributes) LogDeniesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pchc.ref.Append("log_denies_enabled"))
}

func (pchc PolicyControllerHubConfigAttributes) MutationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pchc.ref.Append("mutation_enabled"))
}

func (pchc PolicyControllerHubConfigAttributes) ReferentialRulesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pchc.ref.Append("referential_rules_enabled"))
}

func (pchc PolicyControllerHubConfigAttributes) DeploymentConfigs() terra.SetValue[DeploymentConfigsAttributes] {
	return terra.ReferenceAsSet[DeploymentConfigsAttributes](pchc.ref.Append("deployment_configs"))
}

func (pchc PolicyControllerHubConfigAttributes) Monitoring() terra.ListValue[PolicyControllerHubConfigMonitoringAttributes] {
	return terra.ReferenceAsList[PolicyControllerHubConfigMonitoringAttributes](pchc.ref.Append("monitoring"))
}

func (pchc PolicyControllerHubConfigAttributes) PolicyContent() terra.ListValue[PolicyContentAttributes] {
	return terra.ReferenceAsList[PolicyContentAttributes](pchc.ref.Append("policy_content"))
}

type DeploymentConfigsAttributes struct {
	ref terra.Reference
}

func (dc DeploymentConfigsAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DeploymentConfigsAttributes) InternalWithRef(ref terra.Reference) DeploymentConfigsAttributes {
	return DeploymentConfigsAttributes{ref: ref}
}

func (dc DeploymentConfigsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DeploymentConfigsAttributes) ComponentName() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("component_name"))
}

func (dc DeploymentConfigsAttributes) PodAffinity() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("pod_affinity"))
}

func (dc DeploymentConfigsAttributes) ReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(dc.ref.Append("replica_count"))
}

func (dc DeploymentConfigsAttributes) ContainerResources() terra.ListValue[ContainerResourcesAttributes] {
	return terra.ReferenceAsList[ContainerResourcesAttributes](dc.ref.Append("container_resources"))
}

func (dc DeploymentConfigsAttributes) PodTolerations() terra.ListValue[PodTolerationsAttributes] {
	return terra.ReferenceAsList[PodTolerationsAttributes](dc.ref.Append("pod_tolerations"))
}

type ContainerResourcesAttributes struct {
	ref terra.Reference
}

func (cr ContainerResourcesAttributes) InternalRef() (terra.Reference, error) {
	return cr.ref, nil
}

func (cr ContainerResourcesAttributes) InternalWithRef(ref terra.Reference) ContainerResourcesAttributes {
	return ContainerResourcesAttributes{ref: ref}
}

func (cr ContainerResourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cr.ref.InternalTokens()
}

func (cr ContainerResourcesAttributes) Limits() terra.ListValue[LimitsAttributes] {
	return terra.ReferenceAsList[LimitsAttributes](cr.ref.Append("limits"))
}

func (cr ContainerResourcesAttributes) Requests() terra.ListValue[RequestsAttributes] {
	return terra.ReferenceAsList[RequestsAttributes](cr.ref.Append("requests"))
}

type LimitsAttributes struct {
	ref terra.Reference
}

func (l LimitsAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LimitsAttributes) InternalWithRef(ref terra.Reference) LimitsAttributes {
	return LimitsAttributes{ref: ref}
}

func (l LimitsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LimitsAttributes) Cpu() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("cpu"))
}

func (l LimitsAttributes) Memory() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("memory"))
}

type RequestsAttributes struct {
	ref terra.Reference
}

func (r RequestsAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RequestsAttributes) InternalWithRef(ref terra.Reference) RequestsAttributes {
	return RequestsAttributes{ref: ref}
}

func (r RequestsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RequestsAttributes) Cpu() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("cpu"))
}

func (r RequestsAttributes) Memory() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("memory"))
}

type PodTolerationsAttributes struct {
	ref terra.Reference
}

func (pt PodTolerationsAttributes) InternalRef() (terra.Reference, error) {
	return pt.ref, nil
}

func (pt PodTolerationsAttributes) InternalWithRef(ref terra.Reference) PodTolerationsAttributes {
	return PodTolerationsAttributes{ref: ref}
}

func (pt PodTolerationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pt.ref.InternalTokens()
}

func (pt PodTolerationsAttributes) Effect() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("effect"))
}

func (pt PodTolerationsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("key"))
}

func (pt PodTolerationsAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("operator"))
}

func (pt PodTolerationsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("value"))
}

type PolicyControllerHubConfigMonitoringAttributes struct {
	ref terra.Reference
}

func (m PolicyControllerHubConfigMonitoringAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m PolicyControllerHubConfigMonitoringAttributes) InternalWithRef(ref terra.Reference) PolicyControllerHubConfigMonitoringAttributes {
	return PolicyControllerHubConfigMonitoringAttributes{ref: ref}
}

func (m PolicyControllerHubConfigMonitoringAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m PolicyControllerHubConfigMonitoringAttributes) Backends() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](m.ref.Append("backends"))
}

type PolicyContentAttributes struct {
	ref terra.Reference
}

func (pc PolicyContentAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PolicyContentAttributes) InternalWithRef(ref terra.Reference) PolicyContentAttributes {
	return PolicyContentAttributes{ref: ref}
}

func (pc PolicyContentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PolicyContentAttributes) Bundles() terra.SetValue[BundlesAttributes] {
	return terra.ReferenceAsSet[BundlesAttributes](pc.ref.Append("bundles"))
}

func (pc PolicyContentAttributes) TemplateLibrary() terra.ListValue[TemplateLibraryAttributes] {
	return terra.ReferenceAsList[TemplateLibraryAttributes](pc.ref.Append("template_library"))
}

type BundlesAttributes struct {
	ref terra.Reference
}

func (b BundlesAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BundlesAttributes) InternalWithRef(ref terra.Reference) BundlesAttributes {
	return BundlesAttributes{ref: ref}
}

func (b BundlesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BundlesAttributes) BundleName() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("bundle_name"))
}

func (b BundlesAttributes) ExemptedNamespaces() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](b.ref.Append("exempted_namespaces"))
}

type TemplateLibraryAttributes struct {
	ref terra.Reference
}

func (tl TemplateLibraryAttributes) InternalRef() (terra.Reference, error) {
	return tl.ref, nil
}

func (tl TemplateLibraryAttributes) InternalWithRef(ref terra.Reference) TemplateLibraryAttributes {
	return TemplateLibraryAttributes{ref: ref}
}

func (tl TemplateLibraryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tl.ref.InternalTokens()
}

func (tl TemplateLibraryAttributes) Installation() terra.StringValue {
	return terra.ReferenceAsString(tl.ref.Append("installation"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ConfigmanagementState struct {
	Version             string                     `json:"version"`
	Binauthz            []BinauthzState            `json:"binauthz"`
	ConfigSync          []ConfigSyncState          `json:"config_sync"`
	HierarchyController []HierarchyControllerState `json:"hierarchy_controller"`
	PolicyController    []PolicyControllerState    `json:"policy_controller"`
}

type BinauthzState struct {
	Enabled bool `json:"enabled"`
}

type ConfigSyncState struct {
	MetricsGcpServiceAccountEmail string     `json:"metrics_gcp_service_account_email"`
	PreventDrift                  bool       `json:"prevent_drift"`
	SourceFormat                  string     `json:"source_format"`
	Git                           []GitState `json:"git"`
	Oci                           []OciState `json:"oci"`
}

type GitState struct {
	GcpServiceAccountEmail string `json:"gcp_service_account_email"`
	HttpsProxy             string `json:"https_proxy"`
	PolicyDir              string `json:"policy_dir"`
	SecretType             string `json:"secret_type"`
	SyncBranch             string `json:"sync_branch"`
	SyncRepo               string `json:"sync_repo"`
	SyncRev                string `json:"sync_rev"`
	SyncWaitSecs           string `json:"sync_wait_secs"`
}

type OciState struct {
	GcpServiceAccountEmail string `json:"gcp_service_account_email"`
	PolicyDir              string `json:"policy_dir"`
	SecretType             string `json:"secret_type"`
	SyncRepo               string `json:"sync_repo"`
	SyncWaitSecs           string `json:"sync_wait_secs"`
}

type HierarchyControllerState struct {
	EnableHierarchicalResourceQuota bool `json:"enable_hierarchical_resource_quota"`
	EnablePodTreeLabels             bool `json:"enable_pod_tree_labels"`
	Enabled                         bool `json:"enabled"`
}

type PolicyControllerState struct {
	AuditIntervalSeconds     string                            `json:"audit_interval_seconds"`
	Enabled                  bool                              `json:"enabled"`
	ExemptableNamespaces     []string                          `json:"exemptable_namespaces"`
	LogDeniesEnabled         bool                              `json:"log_denies_enabled"`
	MutationEnabled          bool                              `json:"mutation_enabled"`
	ReferentialRulesEnabled  bool                              `json:"referential_rules_enabled"`
	TemplateLibraryInstalled bool                              `json:"template_library_installed"`
	Monitoring               []PolicyControllerMonitoringState `json:"monitoring"`
}

type PolicyControllerMonitoringState struct {
	Backends []string `json:"backends"`
}

type MeshState struct {
	ControlPlane string `json:"control_plane"`
	Management   string `json:"management"`
}

type PolicycontrollerState struct {
	Version                   string                           `json:"version"`
	PolicyControllerHubConfig []PolicyControllerHubConfigState `json:"policy_controller_hub_config"`
}

type PolicyControllerHubConfigState struct {
	AuditIntervalSeconds     float64                                    `json:"audit_interval_seconds"`
	ConstraintViolationLimit float64                                    `json:"constraint_violation_limit"`
	ExemptableNamespaces     []string                                   `json:"exemptable_namespaces"`
	InstallSpec              string                                     `json:"install_spec"`
	LogDeniesEnabled         bool                                       `json:"log_denies_enabled"`
	MutationEnabled          bool                                       `json:"mutation_enabled"`
	ReferentialRulesEnabled  bool                                       `json:"referential_rules_enabled"`
	DeploymentConfigs        []DeploymentConfigsState                   `json:"deployment_configs"`
	Monitoring               []PolicyControllerHubConfigMonitoringState `json:"monitoring"`
	PolicyContent            []PolicyContentState                       `json:"policy_content"`
}

type DeploymentConfigsState struct {
	ComponentName      string                    `json:"component_name"`
	PodAffinity        string                    `json:"pod_affinity"`
	ReplicaCount       float64                   `json:"replica_count"`
	ContainerResources []ContainerResourcesState `json:"container_resources"`
	PodTolerations     []PodTolerationsState     `json:"pod_tolerations"`
}

type ContainerResourcesState struct {
	Limits   []LimitsState   `json:"limits"`
	Requests []RequestsState `json:"requests"`
}

type LimitsState struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}

type RequestsState struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}

type PodTolerationsState struct {
	Effect   string `json:"effect"`
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type PolicyControllerHubConfigMonitoringState struct {
	Backends []string `json:"backends"`
}

type PolicyContentState struct {
	Bundles         []BundlesState         `json:"bundles"`
	TemplateLibrary []TemplateLibraryState `json:"template_library"`
}

type BundlesState struct {
	BundleName         string   `json:"bundle_name"`
	ExemptedNamespaces []string `json:"exempted_namespaces"`
}

type TemplateLibraryState struct {
	Installation string `json:"installation"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

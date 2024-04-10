// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package gkebackuprestoreplan

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type RestoreConfig struct {
	// AllNamespaces: bool, optional
	AllNamespaces terra.BoolValue `hcl:"all_namespaces,attr"`
	// ClusterResourceConflictPolicy: string, optional
	ClusterResourceConflictPolicy terra.StringValue `hcl:"cluster_resource_conflict_policy,attr"`
	// NamespacedResourceRestoreMode: string, optional
	NamespacedResourceRestoreMode terra.StringValue `hcl:"namespaced_resource_restore_mode,attr"`
	// NoNamespaces: bool, optional
	NoNamespaces terra.BoolValue `hcl:"no_namespaces,attr"`
	// VolumeDataRestorePolicy: string, optional
	VolumeDataRestorePolicy terra.StringValue `hcl:"volume_data_restore_policy,attr"`
	// ClusterResourceRestoreScope: optional
	ClusterResourceRestoreScope *ClusterResourceRestoreScope `hcl:"cluster_resource_restore_scope,block"`
	// ExcludedNamespaces: optional
	ExcludedNamespaces *ExcludedNamespaces `hcl:"excluded_namespaces,block"`
	// SelectedApplications: optional
	SelectedApplications *SelectedApplications `hcl:"selected_applications,block"`
	// SelectedNamespaces: optional
	SelectedNamespaces *SelectedNamespaces `hcl:"selected_namespaces,block"`
	// TransformationRules: min=0
	TransformationRules []TransformationRules `hcl:"transformation_rules,block" validate:"min=0"`
}

type ClusterResourceRestoreScope struct {
	// AllGroupKinds: bool, optional
	AllGroupKinds terra.BoolValue `hcl:"all_group_kinds,attr"`
	// NoGroupKinds: bool, optional
	NoGroupKinds terra.BoolValue `hcl:"no_group_kinds,attr"`
	// ExcludedGroupKinds: min=0
	ExcludedGroupKinds []ExcludedGroupKinds `hcl:"excluded_group_kinds,block" validate:"min=0"`
	// SelectedGroupKinds: min=0
	SelectedGroupKinds []SelectedGroupKinds `hcl:"selected_group_kinds,block" validate:"min=0"`
}

type ExcludedGroupKinds struct {
	// ResourceGroup: string, optional
	ResourceGroup terra.StringValue `hcl:"resource_group,attr"`
	// ResourceKind: string, optional
	ResourceKind terra.StringValue `hcl:"resource_kind,attr"`
}

type SelectedGroupKinds struct {
	// ResourceGroup: string, optional
	ResourceGroup terra.StringValue `hcl:"resource_group,attr"`
	// ResourceKind: string, optional
	ResourceKind terra.StringValue `hcl:"resource_kind,attr"`
}

type ExcludedNamespaces struct {
	// Namespaces: list of string, required
	Namespaces terra.ListValue[terra.StringValue] `hcl:"namespaces,attr" validate:"required"`
}

type SelectedApplications struct {
	// NamespacedNames: min=1
	NamespacedNames []NamespacedNames `hcl:"namespaced_names,block" validate:"min=1"`
}

type NamespacedNames struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
}

type SelectedNamespaces struct {
	// Namespaces: list of string, required
	Namespaces terra.ListValue[terra.StringValue] `hcl:"namespaces,attr" validate:"required"`
}

type TransformationRules struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FieldActions: min=1
	FieldActions []FieldActions `hcl:"field_actions,block" validate:"min=1"`
	// ResourceFilter: optional
	ResourceFilter *ResourceFilter `hcl:"resource_filter,block"`
}

type FieldActions struct {
	// FromPath: string, optional
	FromPath terra.StringValue `hcl:"from_path,attr"`
	// Op: string, required
	Op terra.StringValue `hcl:"op,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type ResourceFilter struct {
	// JsonPath: string, optional
	JsonPath terra.StringValue `hcl:"json_path,attr"`
	// Namespaces: list of string, optional
	Namespaces terra.ListValue[terra.StringValue] `hcl:"namespaces,attr"`
	// GroupKinds: min=0
	GroupKinds []GroupKinds `hcl:"group_kinds,block" validate:"min=0"`
}

type GroupKinds struct {
	// ResourceGroup: string, optional
	ResourceGroup terra.StringValue `hcl:"resource_group,attr"`
	// ResourceKind: string, optional
	ResourceKind terra.StringValue `hcl:"resource_kind,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type RestoreConfigAttributes struct {
	ref terra.Reference
}

func (rc RestoreConfigAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RestoreConfigAttributes) InternalWithRef(ref terra.Reference) RestoreConfigAttributes {
	return RestoreConfigAttributes{ref: ref}
}

func (rc RestoreConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RestoreConfigAttributes) AllNamespaces() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("all_namespaces"))
}

func (rc RestoreConfigAttributes) ClusterResourceConflictPolicy() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_resource_conflict_policy"))
}

func (rc RestoreConfigAttributes) NamespacedResourceRestoreMode() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("namespaced_resource_restore_mode"))
}

func (rc RestoreConfigAttributes) NoNamespaces() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("no_namespaces"))
}

func (rc RestoreConfigAttributes) VolumeDataRestorePolicy() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("volume_data_restore_policy"))
}

func (rc RestoreConfigAttributes) ClusterResourceRestoreScope() terra.ListValue[ClusterResourceRestoreScopeAttributes] {
	return terra.ReferenceAsList[ClusterResourceRestoreScopeAttributes](rc.ref.Append("cluster_resource_restore_scope"))
}

func (rc RestoreConfigAttributes) ExcludedNamespaces() terra.ListValue[ExcludedNamespacesAttributes] {
	return terra.ReferenceAsList[ExcludedNamespacesAttributes](rc.ref.Append("excluded_namespaces"))
}

func (rc RestoreConfigAttributes) SelectedApplications() terra.ListValue[SelectedApplicationsAttributes] {
	return terra.ReferenceAsList[SelectedApplicationsAttributes](rc.ref.Append("selected_applications"))
}

func (rc RestoreConfigAttributes) SelectedNamespaces() terra.ListValue[SelectedNamespacesAttributes] {
	return terra.ReferenceAsList[SelectedNamespacesAttributes](rc.ref.Append("selected_namespaces"))
}

func (rc RestoreConfigAttributes) TransformationRules() terra.ListValue[TransformationRulesAttributes] {
	return terra.ReferenceAsList[TransformationRulesAttributes](rc.ref.Append("transformation_rules"))
}

type ClusterResourceRestoreScopeAttributes struct {
	ref terra.Reference
}

func (crrs ClusterResourceRestoreScopeAttributes) InternalRef() (terra.Reference, error) {
	return crrs.ref, nil
}

func (crrs ClusterResourceRestoreScopeAttributes) InternalWithRef(ref terra.Reference) ClusterResourceRestoreScopeAttributes {
	return ClusterResourceRestoreScopeAttributes{ref: ref}
}

func (crrs ClusterResourceRestoreScopeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return crrs.ref.InternalTokens()
}

func (crrs ClusterResourceRestoreScopeAttributes) AllGroupKinds() terra.BoolValue {
	return terra.ReferenceAsBool(crrs.ref.Append("all_group_kinds"))
}

func (crrs ClusterResourceRestoreScopeAttributes) NoGroupKinds() terra.BoolValue {
	return terra.ReferenceAsBool(crrs.ref.Append("no_group_kinds"))
}

func (crrs ClusterResourceRestoreScopeAttributes) ExcludedGroupKinds() terra.ListValue[ExcludedGroupKindsAttributes] {
	return terra.ReferenceAsList[ExcludedGroupKindsAttributes](crrs.ref.Append("excluded_group_kinds"))
}

func (crrs ClusterResourceRestoreScopeAttributes) SelectedGroupKinds() terra.ListValue[SelectedGroupKindsAttributes] {
	return terra.ReferenceAsList[SelectedGroupKindsAttributes](crrs.ref.Append("selected_group_kinds"))
}

type ExcludedGroupKindsAttributes struct {
	ref terra.Reference
}

func (egk ExcludedGroupKindsAttributes) InternalRef() (terra.Reference, error) {
	return egk.ref, nil
}

func (egk ExcludedGroupKindsAttributes) InternalWithRef(ref terra.Reference) ExcludedGroupKindsAttributes {
	return ExcludedGroupKindsAttributes{ref: ref}
}

func (egk ExcludedGroupKindsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return egk.ref.InternalTokens()
}

func (egk ExcludedGroupKindsAttributes) ResourceGroup() terra.StringValue {
	return terra.ReferenceAsString(egk.ref.Append("resource_group"))
}

func (egk ExcludedGroupKindsAttributes) ResourceKind() terra.StringValue {
	return terra.ReferenceAsString(egk.ref.Append("resource_kind"))
}

type SelectedGroupKindsAttributes struct {
	ref terra.Reference
}

func (sgk SelectedGroupKindsAttributes) InternalRef() (terra.Reference, error) {
	return sgk.ref, nil
}

func (sgk SelectedGroupKindsAttributes) InternalWithRef(ref terra.Reference) SelectedGroupKindsAttributes {
	return SelectedGroupKindsAttributes{ref: ref}
}

func (sgk SelectedGroupKindsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sgk.ref.InternalTokens()
}

func (sgk SelectedGroupKindsAttributes) ResourceGroup() terra.StringValue {
	return terra.ReferenceAsString(sgk.ref.Append("resource_group"))
}

func (sgk SelectedGroupKindsAttributes) ResourceKind() terra.StringValue {
	return terra.ReferenceAsString(sgk.ref.Append("resource_kind"))
}

type ExcludedNamespacesAttributes struct {
	ref terra.Reference
}

func (en ExcludedNamespacesAttributes) InternalRef() (terra.Reference, error) {
	return en.ref, nil
}

func (en ExcludedNamespacesAttributes) InternalWithRef(ref terra.Reference) ExcludedNamespacesAttributes {
	return ExcludedNamespacesAttributes{ref: ref}
}

func (en ExcludedNamespacesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return en.ref.InternalTokens()
}

func (en ExcludedNamespacesAttributes) Namespaces() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](en.ref.Append("namespaces"))
}

type SelectedApplicationsAttributes struct {
	ref terra.Reference
}

func (sa SelectedApplicationsAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa SelectedApplicationsAttributes) InternalWithRef(ref terra.Reference) SelectedApplicationsAttributes {
	return SelectedApplicationsAttributes{ref: ref}
}

func (sa SelectedApplicationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa SelectedApplicationsAttributes) NamespacedNames() terra.ListValue[NamespacedNamesAttributes] {
	return terra.ReferenceAsList[NamespacedNamesAttributes](sa.ref.Append("namespaced_names"))
}

type NamespacedNamesAttributes struct {
	ref terra.Reference
}

func (nn NamespacedNamesAttributes) InternalRef() (terra.Reference, error) {
	return nn.ref, nil
}

func (nn NamespacedNamesAttributes) InternalWithRef(ref terra.Reference) NamespacedNamesAttributes {
	return NamespacedNamesAttributes{ref: ref}
}

func (nn NamespacedNamesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nn.ref.InternalTokens()
}

func (nn NamespacedNamesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nn.ref.Append("name"))
}

func (nn NamespacedNamesAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(nn.ref.Append("namespace"))
}

type SelectedNamespacesAttributes struct {
	ref terra.Reference
}

func (sn SelectedNamespacesAttributes) InternalRef() (terra.Reference, error) {
	return sn.ref, nil
}

func (sn SelectedNamespacesAttributes) InternalWithRef(ref terra.Reference) SelectedNamespacesAttributes {
	return SelectedNamespacesAttributes{ref: ref}
}

func (sn SelectedNamespacesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sn.ref.InternalTokens()
}

func (sn SelectedNamespacesAttributes) Namespaces() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sn.ref.Append("namespaces"))
}

type TransformationRulesAttributes struct {
	ref terra.Reference
}

func (tr TransformationRulesAttributes) InternalRef() (terra.Reference, error) {
	return tr.ref, nil
}

func (tr TransformationRulesAttributes) InternalWithRef(ref terra.Reference) TransformationRulesAttributes {
	return TransformationRulesAttributes{ref: ref}
}

func (tr TransformationRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tr.ref.InternalTokens()
}

func (tr TransformationRulesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("description"))
}

func (tr TransformationRulesAttributes) FieldActions() terra.ListValue[FieldActionsAttributes] {
	return terra.ReferenceAsList[FieldActionsAttributes](tr.ref.Append("field_actions"))
}

func (tr TransformationRulesAttributes) ResourceFilter() terra.ListValue[ResourceFilterAttributes] {
	return terra.ReferenceAsList[ResourceFilterAttributes](tr.ref.Append("resource_filter"))
}

type FieldActionsAttributes struct {
	ref terra.Reference
}

func (fa FieldActionsAttributes) InternalRef() (terra.Reference, error) {
	return fa.ref, nil
}

func (fa FieldActionsAttributes) InternalWithRef(ref terra.Reference) FieldActionsAttributes {
	return FieldActionsAttributes{ref: ref}
}

func (fa FieldActionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fa.ref.InternalTokens()
}

func (fa FieldActionsAttributes) FromPath() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("from_path"))
}

func (fa FieldActionsAttributes) Op() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("op"))
}

func (fa FieldActionsAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("path"))
}

func (fa FieldActionsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("value"))
}

type ResourceFilterAttributes struct {
	ref terra.Reference
}

func (rf ResourceFilterAttributes) InternalRef() (terra.Reference, error) {
	return rf.ref, nil
}

func (rf ResourceFilterAttributes) InternalWithRef(ref terra.Reference) ResourceFilterAttributes {
	return ResourceFilterAttributes{ref: ref}
}

func (rf ResourceFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rf.ref.InternalTokens()
}

func (rf ResourceFilterAttributes) JsonPath() terra.StringValue {
	return terra.ReferenceAsString(rf.ref.Append("json_path"))
}

func (rf ResourceFilterAttributes) Namespaces() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rf.ref.Append("namespaces"))
}

func (rf ResourceFilterAttributes) GroupKinds() terra.ListValue[GroupKindsAttributes] {
	return terra.ReferenceAsList[GroupKindsAttributes](rf.ref.Append("group_kinds"))
}

type GroupKindsAttributes struct {
	ref terra.Reference
}

func (gk GroupKindsAttributes) InternalRef() (terra.Reference, error) {
	return gk.ref, nil
}

func (gk GroupKindsAttributes) InternalWithRef(ref terra.Reference) GroupKindsAttributes {
	return GroupKindsAttributes{ref: ref}
}

func (gk GroupKindsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gk.ref.InternalTokens()
}

func (gk GroupKindsAttributes) ResourceGroup() terra.StringValue {
	return terra.ReferenceAsString(gk.ref.Append("resource_group"))
}

func (gk GroupKindsAttributes) ResourceKind() terra.StringValue {
	return terra.ReferenceAsString(gk.ref.Append("resource_kind"))
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

type RestoreConfigState struct {
	AllNamespaces                 bool                               `json:"all_namespaces"`
	ClusterResourceConflictPolicy string                             `json:"cluster_resource_conflict_policy"`
	NamespacedResourceRestoreMode string                             `json:"namespaced_resource_restore_mode"`
	NoNamespaces                  bool                               `json:"no_namespaces"`
	VolumeDataRestorePolicy       string                             `json:"volume_data_restore_policy"`
	ClusterResourceRestoreScope   []ClusterResourceRestoreScopeState `json:"cluster_resource_restore_scope"`
	ExcludedNamespaces            []ExcludedNamespacesState          `json:"excluded_namespaces"`
	SelectedApplications          []SelectedApplicationsState        `json:"selected_applications"`
	SelectedNamespaces            []SelectedNamespacesState          `json:"selected_namespaces"`
	TransformationRules           []TransformationRulesState         `json:"transformation_rules"`
}

type ClusterResourceRestoreScopeState struct {
	AllGroupKinds      bool                      `json:"all_group_kinds"`
	NoGroupKinds       bool                      `json:"no_group_kinds"`
	ExcludedGroupKinds []ExcludedGroupKindsState `json:"excluded_group_kinds"`
	SelectedGroupKinds []SelectedGroupKindsState `json:"selected_group_kinds"`
}

type ExcludedGroupKindsState struct {
	ResourceGroup string `json:"resource_group"`
	ResourceKind  string `json:"resource_kind"`
}

type SelectedGroupKindsState struct {
	ResourceGroup string `json:"resource_group"`
	ResourceKind  string `json:"resource_kind"`
}

type ExcludedNamespacesState struct {
	Namespaces []string `json:"namespaces"`
}

type SelectedApplicationsState struct {
	NamespacedNames []NamespacedNamesState `json:"namespaced_names"`
}

type NamespacedNamesState struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type SelectedNamespacesState struct {
	Namespaces []string `json:"namespaces"`
}

type TransformationRulesState struct {
	Description    string                `json:"description"`
	FieldActions   []FieldActionsState   `json:"field_actions"`
	ResourceFilter []ResourceFilterState `json:"resource_filter"`
}

type FieldActionsState struct {
	FromPath string `json:"from_path"`
	Op       string `json:"op"`
	Path     string `json:"path"`
	Value    string `json:"value"`
}

type ResourceFilterState struct {
	JsonPath   string            `json:"json_path"`
	Namespaces []string          `json:"namespaces"`
	GroupKinds []GroupKindsState `json:"group_kinds"`
}

type GroupKindsState struct {
	ResourceGroup string `json:"resource_group"`
	ResourceKind  string `json:"resource_kind"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

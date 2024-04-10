// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vertexaifeatureonlinestorefeatureview

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BigQuerySource struct {
	// EntityIdColumns: list of string, required
	EntityIdColumns terra.ListValue[terra.StringValue] `hcl:"entity_id_columns,attr" validate:"required"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
}

type FeatureRegistrySource struct {
	// FeatureGroups: min=1
	FeatureGroups []FeatureGroups `hcl:"feature_groups,block" validate:"min=1"`
}

type FeatureGroups struct {
	// FeatureGroupId: string, required
	FeatureGroupId terra.StringValue `hcl:"feature_group_id,attr" validate:"required"`
	// FeatureIds: list of string, required
	FeatureIds terra.ListValue[terra.StringValue] `hcl:"feature_ids,attr" validate:"required"`
}

type SyncConfig struct {
	// Cron: string, optional
	Cron terra.StringValue `hcl:"cron,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type BigQuerySourceAttributes struct {
	ref terra.Reference
}

func (bqs BigQuerySourceAttributes) InternalRef() (terra.Reference, error) {
	return bqs.ref, nil
}

func (bqs BigQuerySourceAttributes) InternalWithRef(ref terra.Reference) BigQuerySourceAttributes {
	return BigQuerySourceAttributes{ref: ref}
}

func (bqs BigQuerySourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bqs.ref.InternalTokens()
}

func (bqs BigQuerySourceAttributes) EntityIdColumns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bqs.ref.Append("entity_id_columns"))
}

func (bqs BigQuerySourceAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(bqs.ref.Append("uri"))
}

type FeatureRegistrySourceAttributes struct {
	ref terra.Reference
}

func (frs FeatureRegistrySourceAttributes) InternalRef() (terra.Reference, error) {
	return frs.ref, nil
}

func (frs FeatureRegistrySourceAttributes) InternalWithRef(ref terra.Reference) FeatureRegistrySourceAttributes {
	return FeatureRegistrySourceAttributes{ref: ref}
}

func (frs FeatureRegistrySourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return frs.ref.InternalTokens()
}

func (frs FeatureRegistrySourceAttributes) FeatureGroups() terra.ListValue[FeatureGroupsAttributes] {
	return terra.ReferenceAsList[FeatureGroupsAttributes](frs.ref.Append("feature_groups"))
}

type FeatureGroupsAttributes struct {
	ref terra.Reference
}

func (fg FeatureGroupsAttributes) InternalRef() (terra.Reference, error) {
	return fg.ref, nil
}

func (fg FeatureGroupsAttributes) InternalWithRef(ref terra.Reference) FeatureGroupsAttributes {
	return FeatureGroupsAttributes{ref: ref}
}

func (fg FeatureGroupsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fg.ref.InternalTokens()
}

func (fg FeatureGroupsAttributes) FeatureGroupId() terra.StringValue {
	return terra.ReferenceAsString(fg.ref.Append("feature_group_id"))
}

func (fg FeatureGroupsAttributes) FeatureIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fg.ref.Append("feature_ids"))
}

type SyncConfigAttributes struct {
	ref terra.Reference
}

func (sc SyncConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SyncConfigAttributes) InternalWithRef(ref terra.Reference) SyncConfigAttributes {
	return SyncConfigAttributes{ref: ref}
}

func (sc SyncConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SyncConfigAttributes) Cron() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("cron"))
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

type BigQuerySourceState struct {
	EntityIdColumns []string `json:"entity_id_columns"`
	Uri             string   `json:"uri"`
}

type FeatureRegistrySourceState struct {
	FeatureGroups []FeatureGroupsState `json:"feature_groups"`
}

type FeatureGroupsState struct {
	FeatureGroupId string   `json:"feature_group_id"`
	FeatureIds     []string `json:"feature_ids"`
}

type SyncConfigState struct {
	Cron string `json:"cron"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

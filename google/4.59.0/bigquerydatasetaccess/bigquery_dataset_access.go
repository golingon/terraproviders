// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package bigquerydatasetaccess

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Dataset struct {
	// TargetTypes: list of string, required
	TargetTypes terra.ListValue[terra.StringValue] `hcl:"target_types,attr" validate:"required"`
	// DatasetDataset: required
	Dataset *DatasetDataset `hcl:"dataset,block" validate:"required"`
}

type DatasetDataset struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// ProjectId: string, required
	ProjectId terra.StringValue `hcl:"project_id,attr" validate:"required"`
}

type Routine struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// ProjectId: string, required
	ProjectId terra.StringValue `hcl:"project_id,attr" validate:"required"`
	// RoutineId: string, required
	RoutineId terra.StringValue `hcl:"routine_id,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type View struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// ProjectId: string, required
	ProjectId terra.StringValue `hcl:"project_id,attr" validate:"required"`
	// TableId: string, required
	TableId terra.StringValue `hcl:"table_id,attr" validate:"required"`
}

type DatasetAttributes struct {
	ref terra.Reference
}

func (d DatasetAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d DatasetAttributes) InternalWithRef(ref terra.Reference) DatasetAttributes {
	return DatasetAttributes{ref: ref}
}

func (d DatasetAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d DatasetAttributes) TargetTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](d.ref.Append("target_types"))
}

func (d DatasetAttributes) Dataset() terra.ListValue[DatasetDatasetAttributes] {
	return terra.ReferenceList[DatasetDatasetAttributes](d.ref.Append("dataset"))
}

type DatasetDatasetAttributes struct {
	ref terra.Reference
}

func (d DatasetDatasetAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d DatasetDatasetAttributes) InternalWithRef(ref terra.Reference) DatasetDatasetAttributes {
	return DatasetDatasetAttributes{ref: ref}
}

func (d DatasetDatasetAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d DatasetDatasetAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("dataset_id"))
}

func (d DatasetDatasetAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("project_id"))
}

type RoutineAttributes struct {
	ref terra.Reference
}

func (r RoutineAttributes) InternalRef() terra.Reference {
	return r.ref
}

func (r RoutineAttributes) InternalWithRef(ref terra.Reference) RoutineAttributes {
	return RoutineAttributes{ref: ref}
}

func (r RoutineAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r RoutineAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("dataset_id"))
}

func (r RoutineAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("project_id"))
}

func (r RoutineAttributes) RoutineId() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("routine_id"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

type ViewAttributes struct {
	ref terra.Reference
}

func (v ViewAttributes) InternalRef() terra.Reference {
	return v.ref
}

func (v ViewAttributes) InternalWithRef(ref terra.Reference) ViewAttributes {
	return ViewAttributes{ref: ref}
}

func (v ViewAttributes) InternalTokens() hclwrite.Tokens {
	return v.ref.InternalTokens()
}

func (v ViewAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceString(v.ref.Append("dataset_id"))
}

func (v ViewAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceString(v.ref.Append("project_id"))
}

func (v ViewAttributes) TableId() terra.StringValue {
	return terra.ReferenceString(v.ref.Append("table_id"))
}

type DatasetState struct {
	TargetTypes []string              `json:"target_types"`
	Dataset     []DatasetDatasetState `json:"dataset"`
}

type DatasetDatasetState struct {
	DatasetId string `json:"dataset_id"`
	ProjectId string `json:"project_id"`
}

type RoutineState struct {
	DatasetId string `json:"dataset_id"`
	ProjectId string `json:"project_id"`
	RoutineId string `json:"routine_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

type ViewState struct {
	DatasetId string `json:"dataset_id"`
	ProjectId string `json:"project_id"`
	TableId   string `json:"table_id"`
}

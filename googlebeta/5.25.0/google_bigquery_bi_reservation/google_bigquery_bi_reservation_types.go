// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_bigquery_bi_reservation

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type PreferredTables struct {
	// DatasetId: string, optional
	DatasetId terra.StringValue `hcl:"dataset_id,attr"`
	// ProjectId: string, optional
	ProjectId terra.StringValue `hcl:"project_id,attr"`
	// TableId: string, optional
	TableId terra.StringValue `hcl:"table_id,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type PreferredTablesAttributes struct {
	ref terra.Reference
}

func (pt PreferredTablesAttributes) InternalRef() (terra.Reference, error) {
	return pt.ref, nil
}

func (pt PreferredTablesAttributes) InternalWithRef(ref terra.Reference) PreferredTablesAttributes {
	return PreferredTablesAttributes{ref: ref}
}

func (pt PreferredTablesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pt.ref.InternalTokens()
}

func (pt PreferredTablesAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("dataset_id"))
}

func (pt PreferredTablesAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("project_id"))
}

func (pt PreferredTablesAttributes) TableId() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("table_id"))
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

type PreferredTablesState struct {
	DatasetId string `json:"dataset_id"`
	ProjectId string `json:"project_id"`
	TableId   string `json:"table_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

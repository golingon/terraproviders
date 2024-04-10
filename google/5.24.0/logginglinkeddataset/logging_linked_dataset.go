// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package logginglinkeddataset

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BigqueryDataset struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type BigqueryDatasetAttributes struct {
	ref terra.Reference
}

func (bd BigqueryDatasetAttributes) InternalRef() (terra.Reference, error) {
	return bd.ref, nil
}

func (bd BigqueryDatasetAttributes) InternalWithRef(ref terra.Reference) BigqueryDatasetAttributes {
	return BigqueryDatasetAttributes{ref: ref}
}

func (bd BigqueryDatasetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bd.ref.InternalTokens()
}

func (bd BigqueryDatasetAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("dataset_id"))
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

type BigqueryDatasetState struct {
	DatasetId string `json:"dataset_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

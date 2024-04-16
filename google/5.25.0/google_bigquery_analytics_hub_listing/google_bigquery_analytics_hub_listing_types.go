// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_bigquery_analytics_hub_listing

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BigqueryDataset struct {
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
}

type DataProvider struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrimaryContact: string, optional
	PrimaryContact terra.StringValue `hcl:"primary_contact,attr"`
}

type Publisher struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrimaryContact: string, optional
	PrimaryContact terra.StringValue `hcl:"primary_contact,attr"`
}

type RestrictedExportConfig struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// RestrictQueryResult: bool, optional
	RestrictQueryResult terra.BoolValue `hcl:"restrict_query_result,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
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

func (bd BigqueryDatasetAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("dataset"))
}

type DataProviderAttributes struct {
	ref terra.Reference
}

func (dp DataProviderAttributes) InternalRef() (terra.Reference, error) {
	return dp.ref, nil
}

func (dp DataProviderAttributes) InternalWithRef(ref terra.Reference) DataProviderAttributes {
	return DataProviderAttributes{ref: ref}
}

func (dp DataProviderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dp.ref.InternalTokens()
}

func (dp DataProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("name"))
}

func (dp DataProviderAttributes) PrimaryContact() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("primary_contact"))
}

type PublisherAttributes struct {
	ref terra.Reference
}

func (p PublisherAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PublisherAttributes) InternalWithRef(ref terra.Reference) PublisherAttributes {
	return PublisherAttributes{ref: ref}
}

func (p PublisherAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PublisherAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p PublisherAttributes) PrimaryContact() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("primary_contact"))
}

type RestrictedExportConfigAttributes struct {
	ref terra.Reference
}

func (rec RestrictedExportConfigAttributes) InternalRef() (terra.Reference, error) {
	return rec.ref, nil
}

func (rec RestrictedExportConfigAttributes) InternalWithRef(ref terra.Reference) RestrictedExportConfigAttributes {
	return RestrictedExportConfigAttributes{ref: ref}
}

func (rec RestrictedExportConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rec.ref.InternalTokens()
}

func (rec RestrictedExportConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(rec.ref.Append("enabled"))
}

func (rec RestrictedExportConfigAttributes) RestrictQueryResult() terra.BoolValue {
	return terra.ReferenceAsBool(rec.ref.Append("restrict_query_result"))
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

type BigqueryDatasetState struct {
	Dataset string `json:"dataset"`
}

type DataProviderState struct {
	Name           string `json:"name"`
	PrimaryContact string `json:"primary_contact"`
}

type PublisherState struct {
	Name           string `json:"name"`
	PrimaryContact string `json:"primary_contact"`
}

type RestrictedExportConfigState struct {
	Enabled             bool `json:"enabled"`
	RestrictQueryResult bool `json:"restrict_query_result"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

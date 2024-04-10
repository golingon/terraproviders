// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datafactorydataflow

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Sink struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SinkDataset: optional
	Dataset *SinkDataset `hcl:"dataset,block"`
	// SinkFlowlet: optional
	Flowlet *SinkFlowlet `hcl:"flowlet,block"`
	// SinkLinkedService: optional
	LinkedService *SinkLinkedService `hcl:"linked_service,block"`
	// SinkRejectedLinkedService: optional
	RejectedLinkedService *SinkRejectedLinkedService `hcl:"rejected_linked_service,block"`
	// SinkSchemaLinkedService: optional
	SchemaLinkedService *SinkSchemaLinkedService `hcl:"schema_linked_service,block"`
}

type SinkDataset struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type SinkFlowlet struct {
	// DatasetParameters: string, optional
	DatasetParameters terra.StringValue `hcl:"dataset_parameters,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type SinkLinkedService struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type SinkRejectedLinkedService struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type SinkSchemaLinkedService struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type Source struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SourceDataset: optional
	Dataset *SourceDataset `hcl:"dataset,block"`
	// SourceFlowlet: optional
	Flowlet *SourceFlowlet `hcl:"flowlet,block"`
	// SourceLinkedService: optional
	LinkedService *SourceLinkedService `hcl:"linked_service,block"`
	// SourceRejectedLinkedService: optional
	RejectedLinkedService *SourceRejectedLinkedService `hcl:"rejected_linked_service,block"`
	// SourceSchemaLinkedService: optional
	SchemaLinkedService *SourceSchemaLinkedService `hcl:"schema_linked_service,block"`
}

type SourceDataset struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type SourceFlowlet struct {
	// DatasetParameters: string, optional
	DatasetParameters terra.StringValue `hcl:"dataset_parameters,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type SourceLinkedService struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type SourceRejectedLinkedService struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type SourceSchemaLinkedService struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
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

type Transformation struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TransformationDataset: optional
	Dataset *TransformationDataset `hcl:"dataset,block"`
	// TransformationFlowlet: optional
	Flowlet *TransformationFlowlet `hcl:"flowlet,block"`
	// TransformationLinkedService: optional
	LinkedService *TransformationLinkedService `hcl:"linked_service,block"`
}

type TransformationDataset struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type TransformationFlowlet struct {
	// DatasetParameters: string, optional
	DatasetParameters terra.StringValue `hcl:"dataset_parameters,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type TransformationLinkedService struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type SinkAttributes struct {
	ref terra.Reference
}

func (s SinkAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SinkAttributes) InternalWithRef(ref terra.Reference) SinkAttributes {
	return SinkAttributes{ref: ref}
}

func (s SinkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SinkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("description"))
}

func (s SinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SinkAttributes) Dataset() terra.ListValue[SinkDatasetAttributes] {
	return terra.ReferenceAsList[SinkDatasetAttributes](s.ref.Append("dataset"))
}

func (s SinkAttributes) Flowlet() terra.ListValue[SinkFlowletAttributes] {
	return terra.ReferenceAsList[SinkFlowletAttributes](s.ref.Append("flowlet"))
}

func (s SinkAttributes) LinkedService() terra.ListValue[SinkLinkedServiceAttributes] {
	return terra.ReferenceAsList[SinkLinkedServiceAttributes](s.ref.Append("linked_service"))
}

func (s SinkAttributes) RejectedLinkedService() terra.ListValue[SinkRejectedLinkedServiceAttributes] {
	return terra.ReferenceAsList[SinkRejectedLinkedServiceAttributes](s.ref.Append("rejected_linked_service"))
}

func (s SinkAttributes) SchemaLinkedService() terra.ListValue[SinkSchemaLinkedServiceAttributes] {
	return terra.ReferenceAsList[SinkSchemaLinkedServiceAttributes](s.ref.Append("schema_linked_service"))
}

type SinkDatasetAttributes struct {
	ref terra.Reference
}

func (d SinkDatasetAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d SinkDatasetAttributes) InternalWithRef(ref terra.Reference) SinkDatasetAttributes {
	return SinkDatasetAttributes{ref: ref}
}

func (d SinkDatasetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d SinkDatasetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("name"))
}

func (d SinkDatasetAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](d.ref.Append("parameters"))
}

type SinkFlowletAttributes struct {
	ref terra.Reference
}

func (f SinkFlowletAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f SinkFlowletAttributes) InternalWithRef(ref terra.Reference) SinkFlowletAttributes {
	return SinkFlowletAttributes{ref: ref}
}

func (f SinkFlowletAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f SinkFlowletAttributes) DatasetParameters() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("dataset_parameters"))
}

func (f SinkFlowletAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f SinkFlowletAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](f.ref.Append("parameters"))
}

type SinkLinkedServiceAttributes struct {
	ref terra.Reference
}

func (ls SinkLinkedServiceAttributes) InternalRef() (terra.Reference, error) {
	return ls.ref, nil
}

func (ls SinkLinkedServiceAttributes) InternalWithRef(ref terra.Reference) SinkLinkedServiceAttributes {
	return SinkLinkedServiceAttributes{ref: ref}
}

func (ls SinkLinkedServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ls.ref.InternalTokens()
}

func (ls SinkLinkedServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("name"))
}

func (ls SinkLinkedServiceAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ls.ref.Append("parameters"))
}

type SinkRejectedLinkedServiceAttributes struct {
	ref terra.Reference
}

func (rls SinkRejectedLinkedServiceAttributes) InternalRef() (terra.Reference, error) {
	return rls.ref, nil
}

func (rls SinkRejectedLinkedServiceAttributes) InternalWithRef(ref terra.Reference) SinkRejectedLinkedServiceAttributes {
	return SinkRejectedLinkedServiceAttributes{ref: ref}
}

func (rls SinkRejectedLinkedServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rls.ref.InternalTokens()
}

func (rls SinkRejectedLinkedServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rls.ref.Append("name"))
}

func (rls SinkRejectedLinkedServiceAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rls.ref.Append("parameters"))
}

type SinkSchemaLinkedServiceAttributes struct {
	ref terra.Reference
}

func (sls SinkSchemaLinkedServiceAttributes) InternalRef() (terra.Reference, error) {
	return sls.ref, nil
}

func (sls SinkSchemaLinkedServiceAttributes) InternalWithRef(ref terra.Reference) SinkSchemaLinkedServiceAttributes {
	return SinkSchemaLinkedServiceAttributes{ref: ref}
}

func (sls SinkSchemaLinkedServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sls.ref.InternalTokens()
}

func (sls SinkSchemaLinkedServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sls.ref.Append("name"))
}

func (sls SinkSchemaLinkedServiceAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sls.ref.Append("parameters"))
}

type SourceAttributes struct {
	ref terra.Reference
}

func (s SourceAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SourceAttributes) InternalWithRef(ref terra.Reference) SourceAttributes {
	return SourceAttributes{ref: ref}
}

func (s SourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SourceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("description"))
}

func (s SourceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SourceAttributes) Dataset() terra.ListValue[SourceDatasetAttributes] {
	return terra.ReferenceAsList[SourceDatasetAttributes](s.ref.Append("dataset"))
}

func (s SourceAttributes) Flowlet() terra.ListValue[SourceFlowletAttributes] {
	return terra.ReferenceAsList[SourceFlowletAttributes](s.ref.Append("flowlet"))
}

func (s SourceAttributes) LinkedService() terra.ListValue[SourceLinkedServiceAttributes] {
	return terra.ReferenceAsList[SourceLinkedServiceAttributes](s.ref.Append("linked_service"))
}

func (s SourceAttributes) RejectedLinkedService() terra.ListValue[SourceRejectedLinkedServiceAttributes] {
	return terra.ReferenceAsList[SourceRejectedLinkedServiceAttributes](s.ref.Append("rejected_linked_service"))
}

func (s SourceAttributes) SchemaLinkedService() terra.ListValue[SourceSchemaLinkedServiceAttributes] {
	return terra.ReferenceAsList[SourceSchemaLinkedServiceAttributes](s.ref.Append("schema_linked_service"))
}

type SourceDatasetAttributes struct {
	ref terra.Reference
}

func (d SourceDatasetAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d SourceDatasetAttributes) InternalWithRef(ref terra.Reference) SourceDatasetAttributes {
	return SourceDatasetAttributes{ref: ref}
}

func (d SourceDatasetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d SourceDatasetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("name"))
}

func (d SourceDatasetAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](d.ref.Append("parameters"))
}

type SourceFlowletAttributes struct {
	ref terra.Reference
}

func (f SourceFlowletAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f SourceFlowletAttributes) InternalWithRef(ref terra.Reference) SourceFlowletAttributes {
	return SourceFlowletAttributes{ref: ref}
}

func (f SourceFlowletAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f SourceFlowletAttributes) DatasetParameters() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("dataset_parameters"))
}

func (f SourceFlowletAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f SourceFlowletAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](f.ref.Append("parameters"))
}

type SourceLinkedServiceAttributes struct {
	ref terra.Reference
}

func (ls SourceLinkedServiceAttributes) InternalRef() (terra.Reference, error) {
	return ls.ref, nil
}

func (ls SourceLinkedServiceAttributes) InternalWithRef(ref terra.Reference) SourceLinkedServiceAttributes {
	return SourceLinkedServiceAttributes{ref: ref}
}

func (ls SourceLinkedServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ls.ref.InternalTokens()
}

func (ls SourceLinkedServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("name"))
}

func (ls SourceLinkedServiceAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ls.ref.Append("parameters"))
}

type SourceRejectedLinkedServiceAttributes struct {
	ref terra.Reference
}

func (rls SourceRejectedLinkedServiceAttributes) InternalRef() (terra.Reference, error) {
	return rls.ref, nil
}

func (rls SourceRejectedLinkedServiceAttributes) InternalWithRef(ref terra.Reference) SourceRejectedLinkedServiceAttributes {
	return SourceRejectedLinkedServiceAttributes{ref: ref}
}

func (rls SourceRejectedLinkedServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rls.ref.InternalTokens()
}

func (rls SourceRejectedLinkedServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rls.ref.Append("name"))
}

func (rls SourceRejectedLinkedServiceAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rls.ref.Append("parameters"))
}

type SourceSchemaLinkedServiceAttributes struct {
	ref terra.Reference
}

func (sls SourceSchemaLinkedServiceAttributes) InternalRef() (terra.Reference, error) {
	return sls.ref, nil
}

func (sls SourceSchemaLinkedServiceAttributes) InternalWithRef(ref terra.Reference) SourceSchemaLinkedServiceAttributes {
	return SourceSchemaLinkedServiceAttributes{ref: ref}
}

func (sls SourceSchemaLinkedServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sls.ref.InternalTokens()
}

func (sls SourceSchemaLinkedServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sls.ref.Append("name"))
}

func (sls SourceSchemaLinkedServiceAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sls.ref.Append("parameters"))
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

type TransformationAttributes struct {
	ref terra.Reference
}

func (t TransformationAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TransformationAttributes) InternalWithRef(ref terra.Reference) TransformationAttributes {
	return TransformationAttributes{ref: ref}
}

func (t TransformationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TransformationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("description"))
}

func (t TransformationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("name"))
}

func (t TransformationAttributes) Dataset() terra.ListValue[TransformationDatasetAttributes] {
	return terra.ReferenceAsList[TransformationDatasetAttributes](t.ref.Append("dataset"))
}

func (t TransformationAttributes) Flowlet() terra.ListValue[TransformationFlowletAttributes] {
	return terra.ReferenceAsList[TransformationFlowletAttributes](t.ref.Append("flowlet"))
}

func (t TransformationAttributes) LinkedService() terra.ListValue[TransformationLinkedServiceAttributes] {
	return terra.ReferenceAsList[TransformationLinkedServiceAttributes](t.ref.Append("linked_service"))
}

type TransformationDatasetAttributes struct {
	ref terra.Reference
}

func (d TransformationDatasetAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d TransformationDatasetAttributes) InternalWithRef(ref terra.Reference) TransformationDatasetAttributes {
	return TransformationDatasetAttributes{ref: ref}
}

func (d TransformationDatasetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d TransformationDatasetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("name"))
}

func (d TransformationDatasetAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](d.ref.Append("parameters"))
}

type TransformationFlowletAttributes struct {
	ref terra.Reference
}

func (f TransformationFlowletAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f TransformationFlowletAttributes) InternalWithRef(ref terra.Reference) TransformationFlowletAttributes {
	return TransformationFlowletAttributes{ref: ref}
}

func (f TransformationFlowletAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f TransformationFlowletAttributes) DatasetParameters() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("dataset_parameters"))
}

func (f TransformationFlowletAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f TransformationFlowletAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](f.ref.Append("parameters"))
}

type TransformationLinkedServiceAttributes struct {
	ref terra.Reference
}

func (ls TransformationLinkedServiceAttributes) InternalRef() (terra.Reference, error) {
	return ls.ref, nil
}

func (ls TransformationLinkedServiceAttributes) InternalWithRef(ref terra.Reference) TransformationLinkedServiceAttributes {
	return TransformationLinkedServiceAttributes{ref: ref}
}

func (ls TransformationLinkedServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ls.ref.InternalTokens()
}

func (ls TransformationLinkedServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("name"))
}

func (ls TransformationLinkedServiceAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ls.ref.Append("parameters"))
}

type SinkState struct {
	Description           string                           `json:"description"`
	Name                  string                           `json:"name"`
	Dataset               []SinkDatasetState               `json:"dataset"`
	Flowlet               []SinkFlowletState               `json:"flowlet"`
	LinkedService         []SinkLinkedServiceState         `json:"linked_service"`
	RejectedLinkedService []SinkRejectedLinkedServiceState `json:"rejected_linked_service"`
	SchemaLinkedService   []SinkSchemaLinkedServiceState   `json:"schema_linked_service"`
}

type SinkDatasetState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type SinkFlowletState struct {
	DatasetParameters string            `json:"dataset_parameters"`
	Name              string            `json:"name"`
	Parameters        map[string]string `json:"parameters"`
}

type SinkLinkedServiceState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type SinkRejectedLinkedServiceState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type SinkSchemaLinkedServiceState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type SourceState struct {
	Description           string                             `json:"description"`
	Name                  string                             `json:"name"`
	Dataset               []SourceDatasetState               `json:"dataset"`
	Flowlet               []SourceFlowletState               `json:"flowlet"`
	LinkedService         []SourceLinkedServiceState         `json:"linked_service"`
	RejectedLinkedService []SourceRejectedLinkedServiceState `json:"rejected_linked_service"`
	SchemaLinkedService   []SourceSchemaLinkedServiceState   `json:"schema_linked_service"`
}

type SourceDatasetState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type SourceFlowletState struct {
	DatasetParameters string            `json:"dataset_parameters"`
	Name              string            `json:"name"`
	Parameters        map[string]string `json:"parameters"`
}

type SourceLinkedServiceState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type SourceRejectedLinkedServiceState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type SourceSchemaLinkedServiceState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type TransformationState struct {
	Description   string                             `json:"description"`
	Name          string                             `json:"name"`
	Dataset       []TransformationDatasetState       `json:"dataset"`
	Flowlet       []TransformationFlowletState       `json:"flowlet"`
	LinkedService []TransformationLinkedServiceState `json:"linked_service"`
}

type TransformationDatasetState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

type TransformationFlowletState struct {
	DatasetParameters string            `json:"dataset_parameters"`
	Name              string            `json:"name"`
	Parameters        map[string]string `json:"parameters"`
}

type TransformationLinkedServiceState struct {
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

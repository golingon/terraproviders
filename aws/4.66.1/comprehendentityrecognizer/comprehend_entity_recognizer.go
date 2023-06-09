// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package comprehendentityrecognizer

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type InputDataConfig struct {
	// DataFormat: string, optional
	DataFormat terra.StringValue `hcl:"data_format,attr"`
	// Annotations: optional
	Annotations *Annotations `hcl:"annotations,block"`
	// AugmentedManifests: min=0
	AugmentedManifests []AugmentedManifests `hcl:"augmented_manifests,block" validate:"min=0"`
	// Documents: optional
	Documents *Documents `hcl:"documents,block"`
	// EntityList: optional
	EntityList *EntityList `hcl:"entity_list,block"`
	// EntityTypes: min=1,max=25
	EntityTypes []EntityTypes `hcl:"entity_types,block" validate:"min=1,max=25"`
}

type Annotations struct {
	// S3Uri: string, required
	S3Uri terra.StringValue `hcl:"s3_uri,attr" validate:"required"`
	// TestS3Uri: string, optional
	TestS3Uri terra.StringValue `hcl:"test_s3_uri,attr"`
}

type AugmentedManifests struct {
	// AnnotationDataS3Uri: string, optional
	AnnotationDataS3Uri terra.StringValue `hcl:"annotation_data_s3_uri,attr"`
	// AttributeNames: list of string, required
	AttributeNames terra.ListValue[terra.StringValue] `hcl:"attribute_names,attr" validate:"required"`
	// DocumentType: string, optional
	DocumentType terra.StringValue `hcl:"document_type,attr"`
	// S3Uri: string, required
	S3Uri terra.StringValue `hcl:"s3_uri,attr" validate:"required"`
	// SourceDocumentsS3Uri: string, optional
	SourceDocumentsS3Uri terra.StringValue `hcl:"source_documents_s3_uri,attr"`
	// Split: string, optional
	Split terra.StringValue `hcl:"split,attr"`
}

type Documents struct {
	// InputFormat: string, optional
	InputFormat terra.StringValue `hcl:"input_format,attr"`
	// S3Uri: string, required
	S3Uri terra.StringValue `hcl:"s3_uri,attr" validate:"required"`
	// TestS3Uri: string, optional
	TestS3Uri terra.StringValue `hcl:"test_s3_uri,attr"`
}

type EntityList struct {
	// S3Uri: string, required
	S3Uri terra.StringValue `hcl:"s3_uri,attr" validate:"required"`
}

type EntityTypes struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type VpcConfig struct {
	// SecurityGroupIds: set of string, required
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr" validate:"required"`
	// Subnets: set of string, required
	Subnets terra.SetValue[terra.StringValue] `hcl:"subnets,attr" validate:"required"`
}

type InputDataConfigAttributes struct {
	ref terra.Reference
}

func (idc InputDataConfigAttributes) InternalRef() (terra.Reference, error) {
	return idc.ref, nil
}

func (idc InputDataConfigAttributes) InternalWithRef(ref terra.Reference) InputDataConfigAttributes {
	return InputDataConfigAttributes{ref: ref}
}

func (idc InputDataConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return idc.ref.InternalTokens()
}

func (idc InputDataConfigAttributes) DataFormat() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("data_format"))
}

func (idc InputDataConfigAttributes) Annotations() terra.ListValue[AnnotationsAttributes] {
	return terra.ReferenceAsList[AnnotationsAttributes](idc.ref.Append("annotations"))
}

func (idc InputDataConfigAttributes) AugmentedManifests() terra.SetValue[AugmentedManifestsAttributes] {
	return terra.ReferenceAsSet[AugmentedManifestsAttributes](idc.ref.Append("augmented_manifests"))
}

func (idc InputDataConfigAttributes) Documents() terra.ListValue[DocumentsAttributes] {
	return terra.ReferenceAsList[DocumentsAttributes](idc.ref.Append("documents"))
}

func (idc InputDataConfigAttributes) EntityList() terra.ListValue[EntityListAttributes] {
	return terra.ReferenceAsList[EntityListAttributes](idc.ref.Append("entity_list"))
}

func (idc InputDataConfigAttributes) EntityTypes() terra.SetValue[EntityTypesAttributes] {
	return terra.ReferenceAsSet[EntityTypesAttributes](idc.ref.Append("entity_types"))
}

type AnnotationsAttributes struct {
	ref terra.Reference
}

func (a AnnotationsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AnnotationsAttributes) InternalWithRef(ref terra.Reference) AnnotationsAttributes {
	return AnnotationsAttributes{ref: ref}
}

func (a AnnotationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AnnotationsAttributes) S3Uri() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("s3_uri"))
}

func (a AnnotationsAttributes) TestS3Uri() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("test_s3_uri"))
}

type AugmentedManifestsAttributes struct {
	ref terra.Reference
}

func (am AugmentedManifestsAttributes) InternalRef() (terra.Reference, error) {
	return am.ref, nil
}

func (am AugmentedManifestsAttributes) InternalWithRef(ref terra.Reference) AugmentedManifestsAttributes {
	return AugmentedManifestsAttributes{ref: ref}
}

func (am AugmentedManifestsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return am.ref.InternalTokens()
}

func (am AugmentedManifestsAttributes) AnnotationDataS3Uri() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("annotation_data_s3_uri"))
}

func (am AugmentedManifestsAttributes) AttributeNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](am.ref.Append("attribute_names"))
}

func (am AugmentedManifestsAttributes) DocumentType() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("document_type"))
}

func (am AugmentedManifestsAttributes) S3Uri() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("s3_uri"))
}

func (am AugmentedManifestsAttributes) SourceDocumentsS3Uri() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("source_documents_s3_uri"))
}

func (am AugmentedManifestsAttributes) Split() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("split"))
}

type DocumentsAttributes struct {
	ref terra.Reference
}

func (d DocumentsAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DocumentsAttributes) InternalWithRef(ref terra.Reference) DocumentsAttributes {
	return DocumentsAttributes{ref: ref}
}

func (d DocumentsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DocumentsAttributes) InputFormat() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("input_format"))
}

func (d DocumentsAttributes) S3Uri() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("s3_uri"))
}

func (d DocumentsAttributes) TestS3Uri() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("test_s3_uri"))
}

type EntityListAttributes struct {
	ref terra.Reference
}

func (el EntityListAttributes) InternalRef() (terra.Reference, error) {
	return el.ref, nil
}

func (el EntityListAttributes) InternalWithRef(ref terra.Reference) EntityListAttributes {
	return EntityListAttributes{ref: ref}
}

func (el EntityListAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return el.ref.InternalTokens()
}

func (el EntityListAttributes) S3Uri() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("s3_uri"))
}

type EntityTypesAttributes struct {
	ref terra.Reference
}

func (et EntityTypesAttributes) InternalRef() (terra.Reference, error) {
	return et.ref, nil
}

func (et EntityTypesAttributes) InternalWithRef(ref terra.Reference) EntityTypesAttributes {
	return EntityTypesAttributes{ref: ref}
}

func (et EntityTypesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return et.ref.InternalTokens()
}

func (et EntityTypesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("type"))
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

type VpcConfigAttributes struct {
	ref terra.Reference
}

func (vc VpcConfigAttributes) InternalRef() (terra.Reference, error) {
	return vc.ref, nil
}

func (vc VpcConfigAttributes) InternalWithRef(ref terra.Reference) VpcConfigAttributes {
	return VpcConfigAttributes{ref: ref}
}

func (vc VpcConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vc.ref.InternalTokens()
}

func (vc VpcConfigAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("security_group_ids"))
}

func (vc VpcConfigAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("subnets"))
}

type InputDataConfigState struct {
	DataFormat         string                    `json:"data_format"`
	Annotations        []AnnotationsState        `json:"annotations"`
	AugmentedManifests []AugmentedManifestsState `json:"augmented_manifests"`
	Documents          []DocumentsState          `json:"documents"`
	EntityList         []EntityListState         `json:"entity_list"`
	EntityTypes        []EntityTypesState        `json:"entity_types"`
}

type AnnotationsState struct {
	S3Uri     string `json:"s3_uri"`
	TestS3Uri string `json:"test_s3_uri"`
}

type AugmentedManifestsState struct {
	AnnotationDataS3Uri  string   `json:"annotation_data_s3_uri"`
	AttributeNames       []string `json:"attribute_names"`
	DocumentType         string   `json:"document_type"`
	S3Uri                string   `json:"s3_uri"`
	SourceDocumentsS3Uri string   `json:"source_documents_s3_uri"`
	Split                string   `json:"split"`
}

type DocumentsState struct {
	InputFormat string `json:"input_format"`
	S3Uri       string `json:"s3_uri"`
	TestS3Uri   string `json:"test_s3_uri"`
}

type EntityListState struct {
	S3Uri string `json:"s3_uri"`
}

type EntityTypesState struct {
	Type string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type VpcConfigState struct {
	SecurityGroupIds []string `json:"security_group_ids"`
	Subnets          []string `json:"subnets"`
}

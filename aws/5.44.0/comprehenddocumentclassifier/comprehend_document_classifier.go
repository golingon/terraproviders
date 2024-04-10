// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package comprehenddocumentclassifier

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type InputDataConfig struct {
	// DataFormat: string, optional
	DataFormat terra.StringValue `hcl:"data_format,attr"`
	// LabelDelimiter: string, optional
	LabelDelimiter terra.StringValue `hcl:"label_delimiter,attr"`
	// S3Uri: string, optional
	S3Uri terra.StringValue `hcl:"s3_uri,attr"`
	// TestS3Uri: string, optional
	TestS3Uri terra.StringValue `hcl:"test_s3_uri,attr"`
	// AugmentedManifests: min=0
	AugmentedManifests []AugmentedManifests `hcl:"augmented_manifests,block" validate:"min=0"`
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

type OutputDataConfig struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// S3Uri: string, required
	S3Uri terra.StringValue `hcl:"s3_uri,attr" validate:"required"`
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

func (idc InputDataConfigAttributes) LabelDelimiter() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("label_delimiter"))
}

func (idc InputDataConfigAttributes) S3Uri() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("s3_uri"))
}

func (idc InputDataConfigAttributes) TestS3Uri() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("test_s3_uri"))
}

func (idc InputDataConfigAttributes) AugmentedManifests() terra.SetValue[AugmentedManifestsAttributes] {
	return terra.ReferenceAsSet[AugmentedManifestsAttributes](idc.ref.Append("augmented_manifests"))
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

type OutputDataConfigAttributes struct {
	ref terra.Reference
}

func (odc OutputDataConfigAttributes) InternalRef() (terra.Reference, error) {
	return odc.ref, nil
}

func (odc OutputDataConfigAttributes) InternalWithRef(ref terra.Reference) OutputDataConfigAttributes {
	return OutputDataConfigAttributes{ref: ref}
}

func (odc OutputDataConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return odc.ref.InternalTokens()
}

func (odc OutputDataConfigAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(odc.ref.Append("kms_key_id"))
}

func (odc OutputDataConfigAttributes) OutputS3Uri() terra.StringValue {
	return terra.ReferenceAsString(odc.ref.Append("output_s3_uri"))
}

func (odc OutputDataConfigAttributes) S3Uri() terra.StringValue {
	return terra.ReferenceAsString(odc.ref.Append("s3_uri"))
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
	LabelDelimiter     string                    `json:"label_delimiter"`
	S3Uri              string                    `json:"s3_uri"`
	TestS3Uri          string                    `json:"test_s3_uri"`
	AugmentedManifests []AugmentedManifestsState `json:"augmented_manifests"`
}

type AugmentedManifestsState struct {
	AnnotationDataS3Uri  string   `json:"annotation_data_s3_uri"`
	AttributeNames       []string `json:"attribute_names"`
	DocumentType         string   `json:"document_type"`
	S3Uri                string   `json:"s3_uri"`
	SourceDocumentsS3Uri string   `json:"source_documents_s3_uri"`
	Split                string   `json:"split"`
}

type OutputDataConfigState struct {
	KmsKeyId    string `json:"kms_key_id"`
	OutputS3Uri string `json:"output_s3_uri"`
	S3Uri       string `json:"s3_uri"`
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

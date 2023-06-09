// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datakendrafaq "github.com/golingon/terraproviders/aws/4.66.1/datakendrafaq"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKendraFaq creates a new instance of [DataKendraFaq].
func NewDataKendraFaq(name string, args DataKendraFaqArgs) *DataKendraFaq {
	return &DataKendraFaq{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKendraFaq)(nil)

// DataKendraFaq represents the Terraform data resource aws_kendra_faq.
type DataKendraFaq struct {
	Name string
	Args DataKendraFaqArgs
}

// DataSource returns the Terraform object type for [DataKendraFaq].
func (kf *DataKendraFaq) DataSource() string {
	return "aws_kendra_faq"
}

// LocalName returns the local name for [DataKendraFaq].
func (kf *DataKendraFaq) LocalName() string {
	return kf.Name
}

// Configuration returns the configuration (args) for [DataKendraFaq].
func (kf *DataKendraFaq) Configuration() interface{} {
	return kf.Args
}

// Attributes returns the attributes for [DataKendraFaq].
func (kf *DataKendraFaq) Attributes() dataKendraFaqAttributes {
	return dataKendraFaqAttributes{ref: terra.ReferenceDataResource(kf)}
}

// DataKendraFaqArgs contains the configurations for aws_kendra_faq.
type DataKendraFaqArgs struct {
	// FaqId: string, required
	FaqId terra.StringValue `hcl:"faq_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndexId: string, required
	IndexId terra.StringValue `hcl:"index_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// S3Path: min=0
	S3Path []datakendrafaq.S3Path `hcl:"s3_path,block" validate:"min=0"`
}
type dataKendraFaqAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kendra_faq.
func (kf dataKendraFaqAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_kendra_faq.
func (kf dataKendraFaqAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("created_at"))
}

// Description returns a reference to field description of aws_kendra_faq.
func (kf dataKendraFaqAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("description"))
}

// ErrorMessage returns a reference to field error_message of aws_kendra_faq.
func (kf dataKendraFaqAttributes) ErrorMessage() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("error_message"))
}

// FaqId returns a reference to field faq_id of aws_kendra_faq.
func (kf dataKendraFaqAttributes) FaqId() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("faq_id"))
}

// FileFormat returns a reference to field file_format of aws_kendra_faq.
func (kf dataKendraFaqAttributes) FileFormat() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("file_format"))
}

// Id returns a reference to field id of aws_kendra_faq.
func (kf dataKendraFaqAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("id"))
}

// IndexId returns a reference to field index_id of aws_kendra_faq.
func (kf dataKendraFaqAttributes) IndexId() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("index_id"))
}

// LanguageCode returns a reference to field language_code of aws_kendra_faq.
func (kf dataKendraFaqAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("language_code"))
}

// Name returns a reference to field name of aws_kendra_faq.
func (kf dataKendraFaqAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_kendra_faq.
func (kf dataKendraFaqAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_kendra_faq.
func (kf dataKendraFaqAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_kendra_faq.
func (kf dataKendraFaqAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kf.ref.Append("tags"))
}

// UpdatedAt returns a reference to field updated_at of aws_kendra_faq.
func (kf dataKendraFaqAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("updated_at"))
}

func (kf dataKendraFaqAttributes) S3Path() terra.ListValue[datakendrafaq.S3PathAttributes] {
	return terra.ReferenceAsList[datakendrafaq.S3PathAttributes](kf.ref.Append("s3_path"))
}

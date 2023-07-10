// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kendrafaq "github.com/golingon/terraproviders/aws/5.7.0/kendrafaq"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKendraFaq creates a new instance of [KendraFaq].
func NewKendraFaq(name string, args KendraFaqArgs) *KendraFaq {
	return &KendraFaq{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KendraFaq)(nil)

// KendraFaq represents the Terraform resource aws_kendra_faq.
type KendraFaq struct {
	Name      string
	Args      KendraFaqArgs
	state     *kendraFaqState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KendraFaq].
func (kf *KendraFaq) Type() string {
	return "aws_kendra_faq"
}

// LocalName returns the local name for [KendraFaq].
func (kf *KendraFaq) LocalName() string {
	return kf.Name
}

// Configuration returns the configuration (args) for [KendraFaq].
func (kf *KendraFaq) Configuration() interface{} {
	return kf.Args
}

// DependOn is used for other resources to depend on [KendraFaq].
func (kf *KendraFaq) DependOn() terra.Reference {
	return terra.ReferenceResource(kf)
}

// Dependencies returns the list of resources [KendraFaq] depends_on.
func (kf *KendraFaq) Dependencies() terra.Dependencies {
	return kf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KendraFaq].
func (kf *KendraFaq) LifecycleManagement() *terra.Lifecycle {
	return kf.Lifecycle
}

// Attributes returns the attributes for [KendraFaq].
func (kf *KendraFaq) Attributes() kendraFaqAttributes {
	return kendraFaqAttributes{ref: terra.ReferenceResource(kf)}
}

// ImportState imports the given attribute values into [KendraFaq]'s state.
func (kf *KendraFaq) ImportState(av io.Reader) error {
	kf.state = &kendraFaqState{}
	if err := json.NewDecoder(av).Decode(kf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kf.Type(), kf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KendraFaq] has state.
func (kf *KendraFaq) State() (*kendraFaqState, bool) {
	return kf.state, kf.state != nil
}

// StateMust returns the state for [KendraFaq]. Panics if the state is nil.
func (kf *KendraFaq) StateMust() *kendraFaqState {
	if kf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kf.Type(), kf.LocalName()))
	}
	return kf.state
}

// KendraFaqArgs contains the configurations for aws_kendra_faq.
type KendraFaqArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FileFormat: string, optional
	FileFormat terra.StringValue `hcl:"file_format,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndexId: string, required
	IndexId terra.StringValue `hcl:"index_id,attr" validate:"required"`
	// LanguageCode: string, optional
	LanguageCode terra.StringValue `hcl:"language_code,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// S3Path: required
	S3Path *kendrafaq.S3Path `hcl:"s3_path,block" validate:"required"`
	// Timeouts: optional
	Timeouts *kendrafaq.Timeouts `hcl:"timeouts,block"`
}
type kendraFaqAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kendra_faq.
func (kf kendraFaqAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_kendra_faq.
func (kf kendraFaqAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("created_at"))
}

// Description returns a reference to field description of aws_kendra_faq.
func (kf kendraFaqAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("description"))
}

// ErrorMessage returns a reference to field error_message of aws_kendra_faq.
func (kf kendraFaqAttributes) ErrorMessage() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("error_message"))
}

// FaqId returns a reference to field faq_id of aws_kendra_faq.
func (kf kendraFaqAttributes) FaqId() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("faq_id"))
}

// FileFormat returns a reference to field file_format of aws_kendra_faq.
func (kf kendraFaqAttributes) FileFormat() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("file_format"))
}

// Id returns a reference to field id of aws_kendra_faq.
func (kf kendraFaqAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("id"))
}

// IndexId returns a reference to field index_id of aws_kendra_faq.
func (kf kendraFaqAttributes) IndexId() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("index_id"))
}

// LanguageCode returns a reference to field language_code of aws_kendra_faq.
func (kf kendraFaqAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("language_code"))
}

// Name returns a reference to field name of aws_kendra_faq.
func (kf kendraFaqAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_kendra_faq.
func (kf kendraFaqAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_kendra_faq.
func (kf kendraFaqAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_kendra_faq.
func (kf kendraFaqAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kf.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kendra_faq.
func (kf kendraFaqAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kf.ref.Append("tags_all"))
}

// UpdatedAt returns a reference to field updated_at of aws_kendra_faq.
func (kf kendraFaqAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(kf.ref.Append("updated_at"))
}

func (kf kendraFaqAttributes) S3Path() terra.ListValue[kendrafaq.S3PathAttributes] {
	return terra.ReferenceAsList[kendrafaq.S3PathAttributes](kf.ref.Append("s3_path"))
}

func (kf kendraFaqAttributes) Timeouts() kendrafaq.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kendrafaq.TimeoutsAttributes](kf.ref.Append("timeouts"))
}

type kendraFaqState struct {
	Arn          string                   `json:"arn"`
	CreatedAt    string                   `json:"created_at"`
	Description  string                   `json:"description"`
	ErrorMessage string                   `json:"error_message"`
	FaqId        string                   `json:"faq_id"`
	FileFormat   string                   `json:"file_format"`
	Id           string                   `json:"id"`
	IndexId      string                   `json:"index_id"`
	LanguageCode string                   `json:"language_code"`
	Name         string                   `json:"name"`
	RoleArn      string                   `json:"role_arn"`
	Status       string                   `json:"status"`
	Tags         map[string]string        `json:"tags"`
	TagsAll      map[string]string        `json:"tags_all"`
	UpdatedAt    string                   `json:"updated_at"`
	S3Path       []kendrafaq.S3PathState  `json:"s3_path"`
	Timeouts     *kendrafaq.TimeoutsState `json:"timeouts"`
}

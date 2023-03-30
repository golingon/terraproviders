// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kendrafaq "github.com/golingon/terraproviders/aws/4.60.0/kendrafaq"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewKendraFaq(name string, args KendraFaqArgs) *KendraFaq {
	return &KendraFaq{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KendraFaq)(nil)

type KendraFaq struct {
	Name  string
	Args  KendraFaqArgs
	state *kendraFaqState
}

func (kf *KendraFaq) Type() string {
	return "aws_kendra_faq"
}

func (kf *KendraFaq) LocalName() string {
	return kf.Name
}

func (kf *KendraFaq) Configuration() interface{} {
	return kf.Args
}

func (kf *KendraFaq) Attributes() kendraFaqAttributes {
	return kendraFaqAttributes{ref: terra.ReferenceResource(kf)}
}

func (kf *KendraFaq) ImportState(av io.Reader) error {
	kf.state = &kendraFaqState{}
	if err := json.NewDecoder(av).Decode(kf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kf.Type(), kf.LocalName(), err)
	}
	return nil
}

func (kf *KendraFaq) State() (*kendraFaqState, bool) {
	return kf.state, kf.state != nil
}

func (kf *KendraFaq) StateMust() *kendraFaqState {
	if kf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kf.Type(), kf.LocalName()))
	}
	return kf.state
}

func (kf *KendraFaq) DependOn() terra.Reference {
	return terra.ReferenceResource(kf)
}

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
	// DependsOn contains resources that KendraFaq depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type kendraFaqAttributes struct {
	ref terra.Reference
}

func (kf kendraFaqAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("arn"))
}

func (kf kendraFaqAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("created_at"))
}

func (kf kendraFaqAttributes) Description() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("description"))
}

func (kf kendraFaqAttributes) ErrorMessage() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("error_message"))
}

func (kf kendraFaqAttributes) FaqId() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("faq_id"))
}

func (kf kendraFaqAttributes) FileFormat() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("file_format"))
}

func (kf kendraFaqAttributes) Id() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("id"))
}

func (kf kendraFaqAttributes) IndexId() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("index_id"))
}

func (kf kendraFaqAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("language_code"))
}

func (kf kendraFaqAttributes) Name() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("name"))
}

func (kf kendraFaqAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("role_arn"))
}

func (kf kendraFaqAttributes) Status() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("status"))
}

func (kf kendraFaqAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](kf.ref.Append("tags"))
}

func (kf kendraFaqAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](kf.ref.Append("tags_all"))
}

func (kf kendraFaqAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceString(kf.ref.Append("updated_at"))
}

func (kf kendraFaqAttributes) S3Path() terra.ListValue[kendrafaq.S3PathAttributes] {
	return terra.ReferenceList[kendrafaq.S3PathAttributes](kf.ref.Append("s3_path"))
}

func (kf kendraFaqAttributes) Timeouts() kendrafaq.TimeoutsAttributes {
	return terra.ReferenceSingle[kendrafaq.TimeoutsAttributes](kf.ref.Append("timeouts"))
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

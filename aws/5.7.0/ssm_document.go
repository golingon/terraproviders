// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssmdocument "github.com/golingon/terraproviders/aws/5.7.0/ssmdocument"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsmDocument creates a new instance of [SsmDocument].
func NewSsmDocument(name string, args SsmDocumentArgs) *SsmDocument {
	return &SsmDocument{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsmDocument)(nil)

// SsmDocument represents the Terraform resource aws_ssm_document.
type SsmDocument struct {
	Name      string
	Args      SsmDocumentArgs
	state     *ssmDocumentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsmDocument].
func (sd *SsmDocument) Type() string {
	return "aws_ssm_document"
}

// LocalName returns the local name for [SsmDocument].
func (sd *SsmDocument) LocalName() string {
	return sd.Name
}

// Configuration returns the configuration (args) for [SsmDocument].
func (sd *SsmDocument) Configuration() interface{} {
	return sd.Args
}

// DependOn is used for other resources to depend on [SsmDocument].
func (sd *SsmDocument) DependOn() terra.Reference {
	return terra.ReferenceResource(sd)
}

// Dependencies returns the list of resources [SsmDocument] depends_on.
func (sd *SsmDocument) Dependencies() terra.Dependencies {
	return sd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsmDocument].
func (sd *SsmDocument) LifecycleManagement() *terra.Lifecycle {
	return sd.Lifecycle
}

// Attributes returns the attributes for [SsmDocument].
func (sd *SsmDocument) Attributes() ssmDocumentAttributes {
	return ssmDocumentAttributes{ref: terra.ReferenceResource(sd)}
}

// ImportState imports the given attribute values into [SsmDocument]'s state.
func (sd *SsmDocument) ImportState(av io.Reader) error {
	sd.state = &ssmDocumentState{}
	if err := json.NewDecoder(av).Decode(sd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sd.Type(), sd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsmDocument] has state.
func (sd *SsmDocument) State() (*ssmDocumentState, bool) {
	return sd.state, sd.state != nil
}

// StateMust returns the state for [SsmDocument]. Panics if the state is nil.
func (sd *SsmDocument) StateMust() *ssmDocumentState {
	if sd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sd.Type(), sd.LocalName()))
	}
	return sd.state
}

// SsmDocumentArgs contains the configurations for aws_ssm_document.
type SsmDocumentArgs struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// DocumentFormat: string, optional
	DocumentFormat terra.StringValue `hcl:"document_format,attr"`
	// DocumentType: string, required
	DocumentType terra.StringValue `hcl:"document_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Permissions: map of string, optional
	Permissions terra.MapValue[terra.StringValue] `hcl:"permissions,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetType: string, optional
	TargetType terra.StringValue `hcl:"target_type,attr"`
	// VersionName: string, optional
	VersionName terra.StringValue `hcl:"version_name,attr"`
	// Parameter: min=0
	Parameter []ssmdocument.Parameter `hcl:"parameter,block" validate:"min=0"`
	// AttachmentsSource: min=0,max=20
	AttachmentsSource []ssmdocument.AttachmentsSource `hcl:"attachments_source,block" validate:"min=0,max=20"`
}
type ssmDocumentAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssm_document.
func (sd ssmDocumentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("arn"))
}

// Content returns a reference to field content of aws_ssm_document.
func (sd ssmDocumentAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("content"))
}

// CreatedDate returns a reference to field created_date of aws_ssm_document.
func (sd ssmDocumentAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("created_date"))
}

// DefaultVersion returns a reference to field default_version of aws_ssm_document.
func (sd ssmDocumentAttributes) DefaultVersion() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("default_version"))
}

// Description returns a reference to field description of aws_ssm_document.
func (sd ssmDocumentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("description"))
}

// DocumentFormat returns a reference to field document_format of aws_ssm_document.
func (sd ssmDocumentAttributes) DocumentFormat() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("document_format"))
}

// DocumentType returns a reference to field document_type of aws_ssm_document.
func (sd ssmDocumentAttributes) DocumentType() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("document_type"))
}

// DocumentVersion returns a reference to field document_version of aws_ssm_document.
func (sd ssmDocumentAttributes) DocumentVersion() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("document_version"))
}

// Hash returns a reference to field hash of aws_ssm_document.
func (sd ssmDocumentAttributes) Hash() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("hash"))
}

// HashType returns a reference to field hash_type of aws_ssm_document.
func (sd ssmDocumentAttributes) HashType() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("hash_type"))
}

// Id returns a reference to field id of aws_ssm_document.
func (sd ssmDocumentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("id"))
}

// LatestVersion returns a reference to field latest_version of aws_ssm_document.
func (sd ssmDocumentAttributes) LatestVersion() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("latest_version"))
}

// Name returns a reference to field name of aws_ssm_document.
func (sd ssmDocumentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_ssm_document.
func (sd ssmDocumentAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("owner"))
}

// Permissions returns a reference to field permissions of aws_ssm_document.
func (sd ssmDocumentAttributes) Permissions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sd.ref.Append("permissions"))
}

// PlatformTypes returns a reference to field platform_types of aws_ssm_document.
func (sd ssmDocumentAttributes) PlatformTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sd.ref.Append("platform_types"))
}

// SchemaVersion returns a reference to field schema_version of aws_ssm_document.
func (sd ssmDocumentAttributes) SchemaVersion() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("schema_version"))
}

// Status returns a reference to field status of aws_ssm_document.
func (sd ssmDocumentAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_ssm_document.
func (sd ssmDocumentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ssm_document.
func (sd ssmDocumentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sd.ref.Append("tags_all"))
}

// TargetType returns a reference to field target_type of aws_ssm_document.
func (sd ssmDocumentAttributes) TargetType() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("target_type"))
}

// VersionName returns a reference to field version_name of aws_ssm_document.
func (sd ssmDocumentAttributes) VersionName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("version_name"))
}

func (sd ssmDocumentAttributes) Parameter() terra.ListValue[ssmdocument.ParameterAttributes] {
	return terra.ReferenceAsList[ssmdocument.ParameterAttributes](sd.ref.Append("parameter"))
}

func (sd ssmDocumentAttributes) AttachmentsSource() terra.ListValue[ssmdocument.AttachmentsSourceAttributes] {
	return terra.ReferenceAsList[ssmdocument.AttachmentsSourceAttributes](sd.ref.Append("attachments_source"))
}

type ssmDocumentState struct {
	Arn               string                               `json:"arn"`
	Content           string                               `json:"content"`
	CreatedDate       string                               `json:"created_date"`
	DefaultVersion    string                               `json:"default_version"`
	Description       string                               `json:"description"`
	DocumentFormat    string                               `json:"document_format"`
	DocumentType      string                               `json:"document_type"`
	DocumentVersion   string                               `json:"document_version"`
	Hash              string                               `json:"hash"`
	HashType          string                               `json:"hash_type"`
	Id                string                               `json:"id"`
	LatestVersion     string                               `json:"latest_version"`
	Name              string                               `json:"name"`
	Owner             string                               `json:"owner"`
	Permissions       map[string]string                    `json:"permissions"`
	PlatformTypes     []string                             `json:"platform_types"`
	SchemaVersion     string                               `json:"schema_version"`
	Status            string                               `json:"status"`
	Tags              map[string]string                    `json:"tags"`
	TagsAll           map[string]string                    `json:"tags_all"`
	TargetType        string                               `json:"target_type"`
	VersionName       string                               `json:"version_name"`
	Parameter         []ssmdocument.ParameterState         `json:"parameter"`
	AttachmentsSource []ssmdocument.AttachmentsSourceState `json:"attachments_source"`
}

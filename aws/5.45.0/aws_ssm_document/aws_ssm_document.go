// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ssm_document

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_ssm_document.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSsmDocumentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asd *Resource) Type() string {
	return "aws_ssm_document"
}

// LocalName returns the local name for [Resource].
func (asd *Resource) LocalName() string {
	return asd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asd *Resource) Configuration() interface{} {
	return asd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asd *Resource) Dependencies() terra.Dependencies {
	return asd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asd *Resource) LifecycleManagement() *terra.Lifecycle {
	return asd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asd *Resource) Attributes() awsSsmDocumentAttributes {
	return awsSsmDocumentAttributes{ref: terra.ReferenceResource(asd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asd *Resource) ImportState(state io.Reader) error {
	asd.state = &awsSsmDocumentState{}
	if err := json.NewDecoder(state).Decode(asd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asd.Type(), asd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asd *Resource) State() (*awsSsmDocumentState, bool) {
	return asd.state, asd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asd *Resource) StateMust() *awsSsmDocumentState {
	if asd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asd.Type(), asd.LocalName()))
	}
	return asd.state
}

// Args contains the configurations for aws_ssm_document.
type Args struct {
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
	// AttachmentsSource: min=0,max=20
	AttachmentsSource []AttachmentsSource `hcl:"attachments_source,block" validate:"min=0,max=20"`
}

type awsSsmDocumentAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssm_document.
func (asd awsSsmDocumentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("arn"))
}

// Content returns a reference to field content of aws_ssm_document.
func (asd awsSsmDocumentAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("content"))
}

// CreatedDate returns a reference to field created_date of aws_ssm_document.
func (asd awsSsmDocumentAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("created_date"))
}

// DefaultVersion returns a reference to field default_version of aws_ssm_document.
func (asd awsSsmDocumentAttributes) DefaultVersion() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("default_version"))
}

// Description returns a reference to field description of aws_ssm_document.
func (asd awsSsmDocumentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("description"))
}

// DocumentFormat returns a reference to field document_format of aws_ssm_document.
func (asd awsSsmDocumentAttributes) DocumentFormat() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("document_format"))
}

// DocumentType returns a reference to field document_type of aws_ssm_document.
func (asd awsSsmDocumentAttributes) DocumentType() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("document_type"))
}

// DocumentVersion returns a reference to field document_version of aws_ssm_document.
func (asd awsSsmDocumentAttributes) DocumentVersion() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("document_version"))
}

// Hash returns a reference to field hash of aws_ssm_document.
func (asd awsSsmDocumentAttributes) Hash() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("hash"))
}

// HashType returns a reference to field hash_type of aws_ssm_document.
func (asd awsSsmDocumentAttributes) HashType() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("hash_type"))
}

// Id returns a reference to field id of aws_ssm_document.
func (asd awsSsmDocumentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("id"))
}

// LatestVersion returns a reference to field latest_version of aws_ssm_document.
func (asd awsSsmDocumentAttributes) LatestVersion() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("latest_version"))
}

// Name returns a reference to field name of aws_ssm_document.
func (asd awsSsmDocumentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_ssm_document.
func (asd awsSsmDocumentAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("owner"))
}

// Permissions returns a reference to field permissions of aws_ssm_document.
func (asd awsSsmDocumentAttributes) Permissions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asd.ref.Append("permissions"))
}

// PlatformTypes returns a reference to field platform_types of aws_ssm_document.
func (asd awsSsmDocumentAttributes) PlatformTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asd.ref.Append("platform_types"))
}

// SchemaVersion returns a reference to field schema_version of aws_ssm_document.
func (asd awsSsmDocumentAttributes) SchemaVersion() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("schema_version"))
}

// Status returns a reference to field status of aws_ssm_document.
func (asd awsSsmDocumentAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_ssm_document.
func (asd awsSsmDocumentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ssm_document.
func (asd awsSsmDocumentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asd.ref.Append("tags_all"))
}

// TargetType returns a reference to field target_type of aws_ssm_document.
func (asd awsSsmDocumentAttributes) TargetType() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("target_type"))
}

// VersionName returns a reference to field version_name of aws_ssm_document.
func (asd awsSsmDocumentAttributes) VersionName() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("version_name"))
}

func (asd awsSsmDocumentAttributes) Parameter() terra.ListValue[ParameterAttributes] {
	return terra.ReferenceAsList[ParameterAttributes](asd.ref.Append("parameter"))
}

func (asd awsSsmDocumentAttributes) AttachmentsSource() terra.ListValue[AttachmentsSourceAttributes] {
	return terra.ReferenceAsList[AttachmentsSourceAttributes](asd.ref.Append("attachments_source"))
}

type awsSsmDocumentState struct {
	Arn               string                   `json:"arn"`
	Content           string                   `json:"content"`
	CreatedDate       string                   `json:"created_date"`
	DefaultVersion    string                   `json:"default_version"`
	Description       string                   `json:"description"`
	DocumentFormat    string                   `json:"document_format"`
	DocumentType      string                   `json:"document_type"`
	DocumentVersion   string                   `json:"document_version"`
	Hash              string                   `json:"hash"`
	HashType          string                   `json:"hash_type"`
	Id                string                   `json:"id"`
	LatestVersion     string                   `json:"latest_version"`
	Name              string                   `json:"name"`
	Owner             string                   `json:"owner"`
	Permissions       map[string]string        `json:"permissions"`
	PlatformTypes     []string                 `json:"platform_types"`
	SchemaVersion     string                   `json:"schema_version"`
	Status            string                   `json:"status"`
	Tags              map[string]string        `json:"tags"`
	TagsAll           map[string]string        `json:"tags_all"`
	TargetType        string                   `json:"target_type"`
	VersionName       string                   `json:"version_name"`
	Parameter         []ParameterState         `json:"parameter"`
	AttachmentsSource []AttachmentsSourceState `json:"attachments_source"`
}

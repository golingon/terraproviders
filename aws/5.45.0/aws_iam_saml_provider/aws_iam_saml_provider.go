// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iam_saml_provider

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

// Resource represents the Terraform resource aws_iam_saml_provider.
type Resource struct {
	Name      string
	Args      Args
	state     *awsIamSamlProviderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aisp *Resource) Type() string {
	return "aws_iam_saml_provider"
}

// LocalName returns the local name for [Resource].
func (aisp *Resource) LocalName() string {
	return aisp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aisp *Resource) Configuration() interface{} {
	return aisp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aisp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aisp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aisp *Resource) Dependencies() terra.Dependencies {
	return aisp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aisp *Resource) LifecycleManagement() *terra.Lifecycle {
	return aisp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aisp *Resource) Attributes() awsIamSamlProviderAttributes {
	return awsIamSamlProviderAttributes{ref: terra.ReferenceResource(aisp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aisp *Resource) ImportState(state io.Reader) error {
	aisp.state = &awsIamSamlProviderState{}
	if err := json.NewDecoder(state).Decode(aisp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aisp.Type(), aisp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aisp *Resource) State() (*awsIamSamlProviderState, bool) {
	return aisp.state, aisp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aisp *Resource) StateMust() *awsIamSamlProviderState {
	if aisp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aisp.Type(), aisp.LocalName()))
	}
	return aisp.state
}

// Args contains the configurations for aws_iam_saml_provider.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SamlMetadataDocument: string, required
	SamlMetadataDocument terra.StringValue `hcl:"saml_metadata_document,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsIamSamlProviderAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_saml_provider.
func (aisp awsIamSamlProviderAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aisp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_iam_saml_provider.
func (aisp awsIamSamlProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aisp.ref.Append("id"))
}

// Name returns a reference to field name of aws_iam_saml_provider.
func (aisp awsIamSamlProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aisp.ref.Append("name"))
}

// SamlMetadataDocument returns a reference to field saml_metadata_document of aws_iam_saml_provider.
func (aisp awsIamSamlProviderAttributes) SamlMetadataDocument() terra.StringValue {
	return terra.ReferenceAsString(aisp.ref.Append("saml_metadata_document"))
}

// Tags returns a reference to field tags of aws_iam_saml_provider.
func (aisp awsIamSamlProviderAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aisp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_iam_saml_provider.
func (aisp awsIamSamlProviderAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aisp.ref.Append("tags_all"))
}

// ValidUntil returns a reference to field valid_until of aws_iam_saml_provider.
func (aisp awsIamSamlProviderAttributes) ValidUntil() terra.StringValue {
	return terra.ReferenceAsString(aisp.ref.Append("valid_until"))
}

type awsIamSamlProviderState struct {
	Arn                  string            `json:"arn"`
	Id                   string            `json:"id"`
	Name                 string            `json:"name"`
	SamlMetadataDocument string            `json:"saml_metadata_document"`
	Tags                 map[string]string `json:"tags"`
	TagsAll              map[string]string `json:"tags_all"`
	ValidUntil           string            `json:"valid_until"`
}

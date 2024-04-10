// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIamSamlProvider creates a new instance of [IamSamlProvider].
func NewIamSamlProvider(name string, args IamSamlProviderArgs) *IamSamlProvider {
	return &IamSamlProvider{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamSamlProvider)(nil)

// IamSamlProvider represents the Terraform resource aws_iam_saml_provider.
type IamSamlProvider struct {
	Name      string
	Args      IamSamlProviderArgs
	state     *iamSamlProviderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamSamlProvider].
func (isp *IamSamlProvider) Type() string {
	return "aws_iam_saml_provider"
}

// LocalName returns the local name for [IamSamlProvider].
func (isp *IamSamlProvider) LocalName() string {
	return isp.Name
}

// Configuration returns the configuration (args) for [IamSamlProvider].
func (isp *IamSamlProvider) Configuration() interface{} {
	return isp.Args
}

// DependOn is used for other resources to depend on [IamSamlProvider].
func (isp *IamSamlProvider) DependOn() terra.Reference {
	return terra.ReferenceResource(isp)
}

// Dependencies returns the list of resources [IamSamlProvider] depends_on.
func (isp *IamSamlProvider) Dependencies() terra.Dependencies {
	return isp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamSamlProvider].
func (isp *IamSamlProvider) LifecycleManagement() *terra.Lifecycle {
	return isp.Lifecycle
}

// Attributes returns the attributes for [IamSamlProvider].
func (isp *IamSamlProvider) Attributes() iamSamlProviderAttributes {
	return iamSamlProviderAttributes{ref: terra.ReferenceResource(isp)}
}

// ImportState imports the given attribute values into [IamSamlProvider]'s state.
func (isp *IamSamlProvider) ImportState(av io.Reader) error {
	isp.state = &iamSamlProviderState{}
	if err := json.NewDecoder(av).Decode(isp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", isp.Type(), isp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamSamlProvider] has state.
func (isp *IamSamlProvider) State() (*iamSamlProviderState, bool) {
	return isp.state, isp.state != nil
}

// StateMust returns the state for [IamSamlProvider]. Panics if the state is nil.
func (isp *IamSamlProvider) StateMust() *iamSamlProviderState {
	if isp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", isp.Type(), isp.LocalName()))
	}
	return isp.state
}

// IamSamlProviderArgs contains the configurations for aws_iam_saml_provider.
type IamSamlProviderArgs struct {
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
type iamSamlProviderAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_saml_provider.
func (isp iamSamlProviderAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_iam_saml_provider.
func (isp iamSamlProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("id"))
}

// Name returns a reference to field name of aws_iam_saml_provider.
func (isp iamSamlProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("name"))
}

// SamlMetadataDocument returns a reference to field saml_metadata_document of aws_iam_saml_provider.
func (isp iamSamlProviderAttributes) SamlMetadataDocument() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("saml_metadata_document"))
}

// Tags returns a reference to field tags of aws_iam_saml_provider.
func (isp iamSamlProviderAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](isp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_iam_saml_provider.
func (isp iamSamlProviderAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](isp.ref.Append("tags_all"))
}

// ValidUntil returns a reference to field valid_until of aws_iam_saml_provider.
func (isp iamSamlProviderAttributes) ValidUntil() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("valid_until"))
}

type iamSamlProviderState struct {
	Arn                  string            `json:"arn"`
	Id                   string            `json:"id"`
	Name                 string            `json:"name"`
	SamlMetadataDocument string            `json:"saml_metadata_document"`
	Tags                 map[string]string `json:"tags"`
	TagsAll              map[string]string `json:"tags_all"`
	ValidUntil           string            `json:"valid_until"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCognitoIdentityPoolProviderPrincipalTag creates a new instance of [CognitoIdentityPoolProviderPrincipalTag].
func NewCognitoIdentityPoolProviderPrincipalTag(name string, args CognitoIdentityPoolProviderPrincipalTagArgs) *CognitoIdentityPoolProviderPrincipalTag {
	return &CognitoIdentityPoolProviderPrincipalTag{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoIdentityPoolProviderPrincipalTag)(nil)

// CognitoIdentityPoolProviderPrincipalTag represents the Terraform resource aws_cognito_identity_pool_provider_principal_tag.
type CognitoIdentityPoolProviderPrincipalTag struct {
	Name      string
	Args      CognitoIdentityPoolProviderPrincipalTagArgs
	state     *cognitoIdentityPoolProviderPrincipalTagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitoIdentityPoolProviderPrincipalTag].
func (cipppt *CognitoIdentityPoolProviderPrincipalTag) Type() string {
	return "aws_cognito_identity_pool_provider_principal_tag"
}

// LocalName returns the local name for [CognitoIdentityPoolProviderPrincipalTag].
func (cipppt *CognitoIdentityPoolProviderPrincipalTag) LocalName() string {
	return cipppt.Name
}

// Configuration returns the configuration (args) for [CognitoIdentityPoolProviderPrincipalTag].
func (cipppt *CognitoIdentityPoolProviderPrincipalTag) Configuration() interface{} {
	return cipppt.Args
}

// DependOn is used for other resources to depend on [CognitoIdentityPoolProviderPrincipalTag].
func (cipppt *CognitoIdentityPoolProviderPrincipalTag) DependOn() terra.Reference {
	return terra.ReferenceResource(cipppt)
}

// Dependencies returns the list of resources [CognitoIdentityPoolProviderPrincipalTag] depends_on.
func (cipppt *CognitoIdentityPoolProviderPrincipalTag) Dependencies() terra.Dependencies {
	return cipppt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitoIdentityPoolProviderPrincipalTag].
func (cipppt *CognitoIdentityPoolProviderPrincipalTag) LifecycleManagement() *terra.Lifecycle {
	return cipppt.Lifecycle
}

// Attributes returns the attributes for [CognitoIdentityPoolProviderPrincipalTag].
func (cipppt *CognitoIdentityPoolProviderPrincipalTag) Attributes() cognitoIdentityPoolProviderPrincipalTagAttributes {
	return cognitoIdentityPoolProviderPrincipalTagAttributes{ref: terra.ReferenceResource(cipppt)}
}

// ImportState imports the given attribute values into [CognitoIdentityPoolProviderPrincipalTag]'s state.
func (cipppt *CognitoIdentityPoolProviderPrincipalTag) ImportState(av io.Reader) error {
	cipppt.state = &cognitoIdentityPoolProviderPrincipalTagState{}
	if err := json.NewDecoder(av).Decode(cipppt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cipppt.Type(), cipppt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitoIdentityPoolProviderPrincipalTag] has state.
func (cipppt *CognitoIdentityPoolProviderPrincipalTag) State() (*cognitoIdentityPoolProviderPrincipalTagState, bool) {
	return cipppt.state, cipppt.state != nil
}

// StateMust returns the state for [CognitoIdentityPoolProviderPrincipalTag]. Panics if the state is nil.
func (cipppt *CognitoIdentityPoolProviderPrincipalTag) StateMust() *cognitoIdentityPoolProviderPrincipalTagState {
	if cipppt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cipppt.Type(), cipppt.LocalName()))
	}
	return cipppt.state
}

// CognitoIdentityPoolProviderPrincipalTagArgs contains the configurations for aws_cognito_identity_pool_provider_principal_tag.
type CognitoIdentityPoolProviderPrincipalTagArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityPoolId: string, required
	IdentityPoolId terra.StringValue `hcl:"identity_pool_id,attr" validate:"required"`
	// IdentityProviderName: string, required
	IdentityProviderName terra.StringValue `hcl:"identity_provider_name,attr" validate:"required"`
	// PrincipalTags: map of string, optional
	PrincipalTags terra.MapValue[terra.StringValue] `hcl:"principal_tags,attr"`
	// UseDefaults: bool, optional
	UseDefaults terra.BoolValue `hcl:"use_defaults,attr"`
}
type cognitoIdentityPoolProviderPrincipalTagAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_cognito_identity_pool_provider_principal_tag.
func (cipppt cognitoIdentityPoolProviderPrincipalTagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cipppt.ref.Append("id"))
}

// IdentityPoolId returns a reference to field identity_pool_id of aws_cognito_identity_pool_provider_principal_tag.
func (cipppt cognitoIdentityPoolProviderPrincipalTagAttributes) IdentityPoolId() terra.StringValue {
	return terra.ReferenceAsString(cipppt.ref.Append("identity_pool_id"))
}

// IdentityProviderName returns a reference to field identity_provider_name of aws_cognito_identity_pool_provider_principal_tag.
func (cipppt cognitoIdentityPoolProviderPrincipalTagAttributes) IdentityProviderName() terra.StringValue {
	return terra.ReferenceAsString(cipppt.ref.Append("identity_provider_name"))
}

// PrincipalTags returns a reference to field principal_tags of aws_cognito_identity_pool_provider_principal_tag.
func (cipppt cognitoIdentityPoolProviderPrincipalTagAttributes) PrincipalTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cipppt.ref.Append("principal_tags"))
}

// UseDefaults returns a reference to field use_defaults of aws_cognito_identity_pool_provider_principal_tag.
func (cipppt cognitoIdentityPoolProviderPrincipalTagAttributes) UseDefaults() terra.BoolValue {
	return terra.ReferenceAsBool(cipppt.ref.Append("use_defaults"))
}

type cognitoIdentityPoolProviderPrincipalTagState struct {
	Id                   string            `json:"id"`
	IdentityPoolId       string            `json:"identity_pool_id"`
	IdentityProviderName string            `json:"identity_provider_name"`
	PrincipalTags        map[string]string `json:"principal_tags"`
	UseDefaults          bool              `json:"use_defaults"`
}

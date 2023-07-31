// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cognitoidentitypoolrolesattachment "github.com/golingon/terraproviders/aws/5.10.0/cognitoidentitypoolrolesattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCognitoIdentityPoolRolesAttachment creates a new instance of [CognitoIdentityPoolRolesAttachment].
func NewCognitoIdentityPoolRolesAttachment(name string, args CognitoIdentityPoolRolesAttachmentArgs) *CognitoIdentityPoolRolesAttachment {
	return &CognitoIdentityPoolRolesAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoIdentityPoolRolesAttachment)(nil)

// CognitoIdentityPoolRolesAttachment represents the Terraform resource aws_cognito_identity_pool_roles_attachment.
type CognitoIdentityPoolRolesAttachment struct {
	Name      string
	Args      CognitoIdentityPoolRolesAttachmentArgs
	state     *cognitoIdentityPoolRolesAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitoIdentityPoolRolesAttachment].
func (cipra *CognitoIdentityPoolRolesAttachment) Type() string {
	return "aws_cognito_identity_pool_roles_attachment"
}

// LocalName returns the local name for [CognitoIdentityPoolRolesAttachment].
func (cipra *CognitoIdentityPoolRolesAttachment) LocalName() string {
	return cipra.Name
}

// Configuration returns the configuration (args) for [CognitoIdentityPoolRolesAttachment].
func (cipra *CognitoIdentityPoolRolesAttachment) Configuration() interface{} {
	return cipra.Args
}

// DependOn is used for other resources to depend on [CognitoIdentityPoolRolesAttachment].
func (cipra *CognitoIdentityPoolRolesAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(cipra)
}

// Dependencies returns the list of resources [CognitoIdentityPoolRolesAttachment] depends_on.
func (cipra *CognitoIdentityPoolRolesAttachment) Dependencies() terra.Dependencies {
	return cipra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitoIdentityPoolRolesAttachment].
func (cipra *CognitoIdentityPoolRolesAttachment) LifecycleManagement() *terra.Lifecycle {
	return cipra.Lifecycle
}

// Attributes returns the attributes for [CognitoIdentityPoolRolesAttachment].
func (cipra *CognitoIdentityPoolRolesAttachment) Attributes() cognitoIdentityPoolRolesAttachmentAttributes {
	return cognitoIdentityPoolRolesAttachmentAttributes{ref: terra.ReferenceResource(cipra)}
}

// ImportState imports the given attribute values into [CognitoIdentityPoolRolesAttachment]'s state.
func (cipra *CognitoIdentityPoolRolesAttachment) ImportState(av io.Reader) error {
	cipra.state = &cognitoIdentityPoolRolesAttachmentState{}
	if err := json.NewDecoder(av).Decode(cipra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cipra.Type(), cipra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitoIdentityPoolRolesAttachment] has state.
func (cipra *CognitoIdentityPoolRolesAttachment) State() (*cognitoIdentityPoolRolesAttachmentState, bool) {
	return cipra.state, cipra.state != nil
}

// StateMust returns the state for [CognitoIdentityPoolRolesAttachment]. Panics if the state is nil.
func (cipra *CognitoIdentityPoolRolesAttachment) StateMust() *cognitoIdentityPoolRolesAttachmentState {
	if cipra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cipra.Type(), cipra.LocalName()))
	}
	return cipra.state
}

// CognitoIdentityPoolRolesAttachmentArgs contains the configurations for aws_cognito_identity_pool_roles_attachment.
type CognitoIdentityPoolRolesAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityPoolId: string, required
	IdentityPoolId terra.StringValue `hcl:"identity_pool_id,attr" validate:"required"`
	// Roles: map of string, required
	Roles terra.MapValue[terra.StringValue] `hcl:"roles,attr" validate:"required"`
	// RoleMapping: min=0
	RoleMapping []cognitoidentitypoolrolesattachment.RoleMapping `hcl:"role_mapping,block" validate:"min=0"`
}
type cognitoIdentityPoolRolesAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_cognito_identity_pool_roles_attachment.
func (cipra cognitoIdentityPoolRolesAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cipra.ref.Append("id"))
}

// IdentityPoolId returns a reference to field identity_pool_id of aws_cognito_identity_pool_roles_attachment.
func (cipra cognitoIdentityPoolRolesAttachmentAttributes) IdentityPoolId() terra.StringValue {
	return terra.ReferenceAsString(cipra.ref.Append("identity_pool_id"))
}

// Roles returns a reference to field roles of aws_cognito_identity_pool_roles_attachment.
func (cipra cognitoIdentityPoolRolesAttachmentAttributes) Roles() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cipra.ref.Append("roles"))
}

func (cipra cognitoIdentityPoolRolesAttachmentAttributes) RoleMapping() terra.SetValue[cognitoidentitypoolrolesattachment.RoleMappingAttributes] {
	return terra.ReferenceAsSet[cognitoidentitypoolrolesattachment.RoleMappingAttributes](cipra.ref.Append("role_mapping"))
}

type cognitoIdentityPoolRolesAttachmentState struct {
	Id             string                                                `json:"id"`
	IdentityPoolId string                                                `json:"identity_pool_id"`
	Roles          map[string]string                                     `json:"roles"`
	RoleMapping    []cognitoidentitypoolrolesattachment.RoleMappingState `json:"role_mapping"`
}

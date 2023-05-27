// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodebuildSourceCredential creates a new instance of [CodebuildSourceCredential].
func NewCodebuildSourceCredential(name string, args CodebuildSourceCredentialArgs) *CodebuildSourceCredential {
	return &CodebuildSourceCredential{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodebuildSourceCredential)(nil)

// CodebuildSourceCredential represents the Terraform resource aws_codebuild_source_credential.
type CodebuildSourceCredential struct {
	Name      string
	Args      CodebuildSourceCredentialArgs
	state     *codebuildSourceCredentialState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodebuildSourceCredential].
func (csc *CodebuildSourceCredential) Type() string {
	return "aws_codebuild_source_credential"
}

// LocalName returns the local name for [CodebuildSourceCredential].
func (csc *CodebuildSourceCredential) LocalName() string {
	return csc.Name
}

// Configuration returns the configuration (args) for [CodebuildSourceCredential].
func (csc *CodebuildSourceCredential) Configuration() interface{} {
	return csc.Args
}

// DependOn is used for other resources to depend on [CodebuildSourceCredential].
func (csc *CodebuildSourceCredential) DependOn() terra.Reference {
	return terra.ReferenceResource(csc)
}

// Dependencies returns the list of resources [CodebuildSourceCredential] depends_on.
func (csc *CodebuildSourceCredential) Dependencies() terra.Dependencies {
	return csc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodebuildSourceCredential].
func (csc *CodebuildSourceCredential) LifecycleManagement() *terra.Lifecycle {
	return csc.Lifecycle
}

// Attributes returns the attributes for [CodebuildSourceCredential].
func (csc *CodebuildSourceCredential) Attributes() codebuildSourceCredentialAttributes {
	return codebuildSourceCredentialAttributes{ref: terra.ReferenceResource(csc)}
}

// ImportState imports the given attribute values into [CodebuildSourceCredential]'s state.
func (csc *CodebuildSourceCredential) ImportState(av io.Reader) error {
	csc.state = &codebuildSourceCredentialState{}
	if err := json.NewDecoder(av).Decode(csc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csc.Type(), csc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodebuildSourceCredential] has state.
func (csc *CodebuildSourceCredential) State() (*codebuildSourceCredentialState, bool) {
	return csc.state, csc.state != nil
}

// StateMust returns the state for [CodebuildSourceCredential]. Panics if the state is nil.
func (csc *CodebuildSourceCredential) StateMust() *codebuildSourceCredentialState {
	if csc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csc.Type(), csc.LocalName()))
	}
	return csc.state
}

// CodebuildSourceCredentialArgs contains the configurations for aws_codebuild_source_credential.
type CodebuildSourceCredentialArgs struct {
	// AuthType: string, required
	AuthType terra.StringValue `hcl:"auth_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServerType: string, required
	ServerType terra.StringValue `hcl:"server_type,attr" validate:"required"`
	// Token: string, required
	Token terra.StringValue `hcl:"token,attr" validate:"required"`
	// UserName: string, optional
	UserName terra.StringValue `hcl:"user_name,attr"`
}
type codebuildSourceCredentialAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codebuild_source_credential.
func (csc codebuildSourceCredentialAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("arn"))
}

// AuthType returns a reference to field auth_type of aws_codebuild_source_credential.
func (csc codebuildSourceCredentialAttributes) AuthType() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("auth_type"))
}

// Id returns a reference to field id of aws_codebuild_source_credential.
func (csc codebuildSourceCredentialAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("id"))
}

// ServerType returns a reference to field server_type of aws_codebuild_source_credential.
func (csc codebuildSourceCredentialAttributes) ServerType() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("server_type"))
}

// Token returns a reference to field token of aws_codebuild_source_credential.
func (csc codebuildSourceCredentialAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("token"))
}

// UserName returns a reference to field user_name of aws_codebuild_source_credential.
func (csc codebuildSourceCredentialAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("user_name"))
}

type codebuildSourceCredentialState struct {
	Arn        string `json:"arn"`
	AuthType   string `json:"auth_type"`
	Id         string `json:"id"`
	ServerType string `json:"server_type"`
	Token      string `json:"token"`
	UserName   string `json:"user_name"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewIamServiceSpecificCredential(name string, args IamServiceSpecificCredentialArgs) *IamServiceSpecificCredential {
	return &IamServiceSpecificCredential{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamServiceSpecificCredential)(nil)

type IamServiceSpecificCredential struct {
	Name  string
	Args  IamServiceSpecificCredentialArgs
	state *iamServiceSpecificCredentialState
}

func (issc *IamServiceSpecificCredential) Type() string {
	return "aws_iam_service_specific_credential"
}

func (issc *IamServiceSpecificCredential) LocalName() string {
	return issc.Name
}

func (issc *IamServiceSpecificCredential) Configuration() interface{} {
	return issc.Args
}

func (issc *IamServiceSpecificCredential) Attributes() iamServiceSpecificCredentialAttributes {
	return iamServiceSpecificCredentialAttributes{ref: terra.ReferenceResource(issc)}
}

func (issc *IamServiceSpecificCredential) ImportState(av io.Reader) error {
	issc.state = &iamServiceSpecificCredentialState{}
	if err := json.NewDecoder(av).Decode(issc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", issc.Type(), issc.LocalName(), err)
	}
	return nil
}

func (issc *IamServiceSpecificCredential) State() (*iamServiceSpecificCredentialState, bool) {
	return issc.state, issc.state != nil
}

func (issc *IamServiceSpecificCredential) StateMust() *iamServiceSpecificCredentialState {
	if issc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", issc.Type(), issc.LocalName()))
	}
	return issc.state
}

func (issc *IamServiceSpecificCredential) DependOn() terra.Reference {
	return terra.ReferenceResource(issc)
}

type IamServiceSpecificCredentialArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// UserName: string, required
	UserName terra.StringValue `hcl:"user_name,attr" validate:"required"`
	// DependsOn contains resources that IamServiceSpecificCredential depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type iamServiceSpecificCredentialAttributes struct {
	ref terra.Reference
}

func (issc iamServiceSpecificCredentialAttributes) Id() terra.StringValue {
	return terra.ReferenceString(issc.ref.Append("id"))
}

func (issc iamServiceSpecificCredentialAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceString(issc.ref.Append("service_name"))
}

func (issc iamServiceSpecificCredentialAttributes) ServicePassword() terra.StringValue {
	return terra.ReferenceString(issc.ref.Append("service_password"))
}

func (issc iamServiceSpecificCredentialAttributes) ServiceSpecificCredentialId() terra.StringValue {
	return terra.ReferenceString(issc.ref.Append("service_specific_credential_id"))
}

func (issc iamServiceSpecificCredentialAttributes) ServiceUserName() terra.StringValue {
	return terra.ReferenceString(issc.ref.Append("service_user_name"))
}

func (issc iamServiceSpecificCredentialAttributes) Status() terra.StringValue {
	return terra.ReferenceString(issc.ref.Append("status"))
}

func (issc iamServiceSpecificCredentialAttributes) UserName() terra.StringValue {
	return terra.ReferenceString(issc.ref.Append("user_name"))
}

type iamServiceSpecificCredentialState struct {
	Id                          string `json:"id"`
	ServiceName                 string `json:"service_name"`
	ServicePassword             string `json:"service_password"`
	ServiceSpecificCredentialId string `json:"service_specific_credential_id"`
	ServiceUserName             string `json:"service_user_name"`
	Status                      string `json:"status"`
	UserName                    string `json:"user_name"`
}

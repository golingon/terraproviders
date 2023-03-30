// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSesEmailIdentity(name string, args SesEmailIdentityArgs) *SesEmailIdentity {
	return &SesEmailIdentity{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SesEmailIdentity)(nil)

type SesEmailIdentity struct {
	Name  string
	Args  SesEmailIdentityArgs
	state *sesEmailIdentityState
}

func (sei *SesEmailIdentity) Type() string {
	return "aws_ses_email_identity"
}

func (sei *SesEmailIdentity) LocalName() string {
	return sei.Name
}

func (sei *SesEmailIdentity) Configuration() interface{} {
	return sei.Args
}

func (sei *SesEmailIdentity) Attributes() sesEmailIdentityAttributes {
	return sesEmailIdentityAttributes{ref: terra.ReferenceResource(sei)}
}

func (sei *SesEmailIdentity) ImportState(av io.Reader) error {
	sei.state = &sesEmailIdentityState{}
	if err := json.NewDecoder(av).Decode(sei.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sei.Type(), sei.LocalName(), err)
	}
	return nil
}

func (sei *SesEmailIdentity) State() (*sesEmailIdentityState, bool) {
	return sei.state, sei.state != nil
}

func (sei *SesEmailIdentity) StateMust() *sesEmailIdentityState {
	if sei.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sei.Type(), sei.LocalName()))
	}
	return sei.state
}

func (sei *SesEmailIdentity) DependOn() terra.Reference {
	return terra.ReferenceResource(sei)
}

type SesEmailIdentityArgs struct {
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DependsOn contains resources that SesEmailIdentity depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sesEmailIdentityAttributes struct {
	ref terra.Reference
}

func (sei sesEmailIdentityAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(sei.ref.Append("arn"))
}

func (sei sesEmailIdentityAttributes) Email() terra.StringValue {
	return terra.ReferenceString(sei.ref.Append("email"))
}

func (sei sesEmailIdentityAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sei.ref.Append("id"))
}

type sesEmailIdentityState struct {
	Arn   string `json:"arn"`
	Email string `json:"email"`
	Id    string `json:"id"`
}

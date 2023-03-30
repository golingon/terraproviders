// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewQuicksightUser(name string, args QuicksightUserArgs) *QuicksightUser {
	return &QuicksightUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*QuicksightUser)(nil)

type QuicksightUser struct {
	Name  string
	Args  QuicksightUserArgs
	state *quicksightUserState
}

func (qu *QuicksightUser) Type() string {
	return "aws_quicksight_user"
}

func (qu *QuicksightUser) LocalName() string {
	return qu.Name
}

func (qu *QuicksightUser) Configuration() interface{} {
	return qu.Args
}

func (qu *QuicksightUser) Attributes() quicksightUserAttributes {
	return quicksightUserAttributes{ref: terra.ReferenceResource(qu)}
}

func (qu *QuicksightUser) ImportState(av io.Reader) error {
	qu.state = &quicksightUserState{}
	if err := json.NewDecoder(av).Decode(qu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", qu.Type(), qu.LocalName(), err)
	}
	return nil
}

func (qu *QuicksightUser) State() (*quicksightUserState, bool) {
	return qu.state, qu.state != nil
}

func (qu *QuicksightUser) StateMust() *quicksightUserState {
	if qu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", qu.Type(), qu.LocalName()))
	}
	return qu.state
}

func (qu *QuicksightUser) DependOn() terra.Reference {
	return terra.ReferenceResource(qu)
}

type QuicksightUserArgs struct {
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// IamArn: string, optional
	IamArn terra.StringValue `hcl:"iam_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityType: string, required
	IdentityType terra.StringValue `hcl:"identity_type,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// SessionName: string, optional
	SessionName terra.StringValue `hcl:"session_name,attr"`
	// UserName: string, optional
	UserName terra.StringValue `hcl:"user_name,attr"`
	// UserRole: string, required
	UserRole terra.StringValue `hcl:"user_role,attr" validate:"required"`
	// DependsOn contains resources that QuicksightUser depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type quicksightUserAttributes struct {
	ref terra.Reference
}

func (qu quicksightUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(qu.ref.Append("arn"))
}

func (qu quicksightUserAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceString(qu.ref.Append("aws_account_id"))
}

func (qu quicksightUserAttributes) Email() terra.StringValue {
	return terra.ReferenceString(qu.ref.Append("email"))
}

func (qu quicksightUserAttributes) IamArn() terra.StringValue {
	return terra.ReferenceString(qu.ref.Append("iam_arn"))
}

func (qu quicksightUserAttributes) Id() terra.StringValue {
	return terra.ReferenceString(qu.ref.Append("id"))
}

func (qu quicksightUserAttributes) IdentityType() terra.StringValue {
	return terra.ReferenceString(qu.ref.Append("identity_type"))
}

func (qu quicksightUserAttributes) Namespace() terra.StringValue {
	return terra.ReferenceString(qu.ref.Append("namespace"))
}

func (qu quicksightUserAttributes) SessionName() terra.StringValue {
	return terra.ReferenceString(qu.ref.Append("session_name"))
}

func (qu quicksightUserAttributes) UserName() terra.StringValue {
	return terra.ReferenceString(qu.ref.Append("user_name"))
}

func (qu quicksightUserAttributes) UserRole() terra.StringValue {
	return terra.ReferenceString(qu.ref.Append("user_role"))
}

type quicksightUserState struct {
	Arn          string `json:"arn"`
	AwsAccountId string `json:"aws_account_id"`
	Email        string `json:"email"`
	IamArn       string `json:"iam_arn"`
	Id           string `json:"id"`
	IdentityType string `json:"identity_type"`
	Namespace    string `json:"namespace"`
	SessionName  string `json:"session_name"`
	UserName     string `json:"user_name"`
	UserRole     string `json:"user_role"`
}

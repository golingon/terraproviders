// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAppstreamUser(name string, args AppstreamUserArgs) *AppstreamUser {
	return &AppstreamUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppstreamUser)(nil)

type AppstreamUser struct {
	Name  string
	Args  AppstreamUserArgs
	state *appstreamUserState
}

func (au *AppstreamUser) Type() string {
	return "aws_appstream_user"
}

func (au *AppstreamUser) LocalName() string {
	return au.Name
}

func (au *AppstreamUser) Configuration() interface{} {
	return au.Args
}

func (au *AppstreamUser) Attributes() appstreamUserAttributes {
	return appstreamUserAttributes{ref: terra.ReferenceResource(au)}
}

func (au *AppstreamUser) ImportState(av io.Reader) error {
	au.state = &appstreamUserState{}
	if err := json.NewDecoder(av).Decode(au.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", au.Type(), au.LocalName(), err)
	}
	return nil
}

func (au *AppstreamUser) State() (*appstreamUserState, bool) {
	return au.state, au.state != nil
}

func (au *AppstreamUser) StateMust() *appstreamUserState {
	if au.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", au.Type(), au.LocalName()))
	}
	return au.state
}

func (au *AppstreamUser) DependOn() terra.Reference {
	return terra.ReferenceResource(au)
}

type AppstreamUserArgs struct {
	// AuthenticationType: string, required
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// FirstName: string, optional
	FirstName terra.StringValue `hcl:"first_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LastName: string, optional
	LastName terra.StringValue `hcl:"last_name,attr"`
	// SendEmailNotification: bool, optional
	SendEmailNotification terra.BoolValue `hcl:"send_email_notification,attr"`
	// UserName: string, required
	UserName terra.StringValue `hcl:"user_name,attr" validate:"required"`
	// DependsOn contains resources that AppstreamUser depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type appstreamUserAttributes struct {
	ref terra.Reference
}

func (au appstreamUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(au.ref.Append("arn"))
}

func (au appstreamUserAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceString(au.ref.Append("authentication_type"))
}

func (au appstreamUserAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceString(au.ref.Append("created_time"))
}

func (au appstreamUserAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(au.ref.Append("enabled"))
}

func (au appstreamUserAttributes) FirstName() terra.StringValue {
	return terra.ReferenceString(au.ref.Append("first_name"))
}

func (au appstreamUserAttributes) Id() terra.StringValue {
	return terra.ReferenceString(au.ref.Append("id"))
}

func (au appstreamUserAttributes) LastName() terra.StringValue {
	return terra.ReferenceString(au.ref.Append("last_name"))
}

func (au appstreamUserAttributes) SendEmailNotification() terra.BoolValue {
	return terra.ReferenceBool(au.ref.Append("send_email_notification"))
}

func (au appstreamUserAttributes) UserName() terra.StringValue {
	return terra.ReferenceString(au.ref.Append("user_name"))
}

type appstreamUserState struct {
	Arn                   string `json:"arn"`
	AuthenticationType    string `json:"authentication_type"`
	CreatedTime           string `json:"created_time"`
	Enabled               bool   `json:"enabled"`
	FirstName             string `json:"first_name"`
	Id                    string `json:"id"`
	LastName              string `json:"last_name"`
	SendEmailNotification bool   `json:"send_email_notification"`
	UserName              string `json:"user_name"`
}

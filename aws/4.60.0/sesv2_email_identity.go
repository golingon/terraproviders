// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sesv2emailidentity "github.com/golingon/terraproviders/aws/4.60.0/sesv2emailidentity"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSesv2EmailIdentity(name string, args Sesv2EmailIdentityArgs) *Sesv2EmailIdentity {
	return &Sesv2EmailIdentity{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Sesv2EmailIdentity)(nil)

type Sesv2EmailIdentity struct {
	Name  string
	Args  Sesv2EmailIdentityArgs
	state *sesv2EmailIdentityState
}

func (sei *Sesv2EmailIdentity) Type() string {
	return "aws_sesv2_email_identity"
}

func (sei *Sesv2EmailIdentity) LocalName() string {
	return sei.Name
}

func (sei *Sesv2EmailIdentity) Configuration() interface{} {
	return sei.Args
}

func (sei *Sesv2EmailIdentity) Attributes() sesv2EmailIdentityAttributes {
	return sesv2EmailIdentityAttributes{ref: terra.ReferenceResource(sei)}
}

func (sei *Sesv2EmailIdentity) ImportState(av io.Reader) error {
	sei.state = &sesv2EmailIdentityState{}
	if err := json.NewDecoder(av).Decode(sei.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sei.Type(), sei.LocalName(), err)
	}
	return nil
}

func (sei *Sesv2EmailIdentity) State() (*sesv2EmailIdentityState, bool) {
	return sei.state, sei.state != nil
}

func (sei *Sesv2EmailIdentity) StateMust() *sesv2EmailIdentityState {
	if sei.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sei.Type(), sei.LocalName()))
	}
	return sei.state
}

func (sei *Sesv2EmailIdentity) DependOn() terra.Reference {
	return terra.ReferenceResource(sei)
}

type Sesv2EmailIdentityArgs struct {
	// ConfigurationSetName: string, optional
	ConfigurationSetName terra.StringValue `hcl:"configuration_set_name,attr"`
	// EmailIdentity: string, required
	EmailIdentity terra.StringValue `hcl:"email_identity,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DkimSigningAttributes: optional
	DkimSigningAttributes *sesv2emailidentity.DkimSigningAttributes `hcl:"dkim_signing_attributes,block"`
	// DependsOn contains resources that Sesv2EmailIdentity depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sesv2EmailIdentityAttributes struct {
	ref terra.Reference
}

func (sei sesv2EmailIdentityAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(sei.ref.Append("arn"))
}

func (sei sesv2EmailIdentityAttributes) ConfigurationSetName() terra.StringValue {
	return terra.ReferenceString(sei.ref.Append("configuration_set_name"))
}

func (sei sesv2EmailIdentityAttributes) EmailIdentity() terra.StringValue {
	return terra.ReferenceString(sei.ref.Append("email_identity"))
}

func (sei sesv2EmailIdentityAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sei.ref.Append("id"))
}

func (sei sesv2EmailIdentityAttributes) IdentityType() terra.StringValue {
	return terra.ReferenceString(sei.ref.Append("identity_type"))
}

func (sei sesv2EmailIdentityAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sei.ref.Append("tags"))
}

func (sei sesv2EmailIdentityAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sei.ref.Append("tags_all"))
}

func (sei sesv2EmailIdentityAttributes) VerifiedForSendingStatus() terra.BoolValue {
	return terra.ReferenceBool(sei.ref.Append("verified_for_sending_status"))
}

func (sei sesv2EmailIdentityAttributes) DkimSigningAttributes() terra.ListValue[sesv2emailidentity.DkimSigningAttributesAttributes] {
	return terra.ReferenceList[sesv2emailidentity.DkimSigningAttributesAttributes](sei.ref.Append("dkim_signing_attributes"))
}

type sesv2EmailIdentityState struct {
	Arn                      string                                          `json:"arn"`
	ConfigurationSetName     string                                          `json:"configuration_set_name"`
	EmailIdentity            string                                          `json:"email_identity"`
	Id                       string                                          `json:"id"`
	IdentityType             string                                          `json:"identity_type"`
	Tags                     map[string]string                               `json:"tags"`
	TagsAll                  map[string]string                               `json:"tags_all"`
	VerifiedForSendingStatus bool                                            `json:"verified_for_sending_status"`
	DkimSigningAttributes    []sesv2emailidentity.DkimSigningAttributesState `json:"dkim_signing_attributes"`
}

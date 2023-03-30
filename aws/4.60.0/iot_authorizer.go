// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewIotAuthorizer(name string, args IotAuthorizerArgs) *IotAuthorizer {
	return &IotAuthorizer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotAuthorizer)(nil)

type IotAuthorizer struct {
	Name  string
	Args  IotAuthorizerArgs
	state *iotAuthorizerState
}

func (ia *IotAuthorizer) Type() string {
	return "aws_iot_authorizer"
}

func (ia *IotAuthorizer) LocalName() string {
	return ia.Name
}

func (ia *IotAuthorizer) Configuration() interface{} {
	return ia.Args
}

func (ia *IotAuthorizer) Attributes() iotAuthorizerAttributes {
	return iotAuthorizerAttributes{ref: terra.ReferenceResource(ia)}
}

func (ia *IotAuthorizer) ImportState(av io.Reader) error {
	ia.state = &iotAuthorizerState{}
	if err := json.NewDecoder(av).Decode(ia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ia.Type(), ia.LocalName(), err)
	}
	return nil
}

func (ia *IotAuthorizer) State() (*iotAuthorizerState, bool) {
	return ia.state, ia.state != nil
}

func (ia *IotAuthorizer) StateMust() *iotAuthorizerState {
	if ia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ia.Type(), ia.LocalName()))
	}
	return ia.state
}

func (ia *IotAuthorizer) DependOn() terra.Reference {
	return terra.ReferenceResource(ia)
}

type IotAuthorizerArgs struct {
	// AuthorizerFunctionArn: string, required
	AuthorizerFunctionArn terra.StringValue `hcl:"authorizer_function_arn,attr" validate:"required"`
	// EnableCachingForHttp: bool, optional
	EnableCachingForHttp terra.BoolValue `hcl:"enable_caching_for_http,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SigningDisabled: bool, optional
	SigningDisabled terra.BoolValue `hcl:"signing_disabled,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// TokenKeyName: string, optional
	TokenKeyName terra.StringValue `hcl:"token_key_name,attr"`
	// TokenSigningPublicKeys: map of string, optional
	TokenSigningPublicKeys terra.MapValue[terra.StringValue] `hcl:"token_signing_public_keys,attr"`
	// DependsOn contains resources that IotAuthorizer depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type iotAuthorizerAttributes struct {
	ref terra.Reference
}

func (ia iotAuthorizerAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ia.ref.Append("arn"))
}

func (ia iotAuthorizerAttributes) AuthorizerFunctionArn() terra.StringValue {
	return terra.ReferenceString(ia.ref.Append("authorizer_function_arn"))
}

func (ia iotAuthorizerAttributes) EnableCachingForHttp() terra.BoolValue {
	return terra.ReferenceBool(ia.ref.Append("enable_caching_for_http"))
}

func (ia iotAuthorizerAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ia.ref.Append("id"))
}

func (ia iotAuthorizerAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ia.ref.Append("name"))
}

func (ia iotAuthorizerAttributes) SigningDisabled() terra.BoolValue {
	return terra.ReferenceBool(ia.ref.Append("signing_disabled"))
}

func (ia iotAuthorizerAttributes) Status() terra.StringValue {
	return terra.ReferenceString(ia.ref.Append("status"))
}

func (ia iotAuthorizerAttributes) TokenKeyName() terra.StringValue {
	return terra.ReferenceString(ia.ref.Append("token_key_name"))
}

func (ia iotAuthorizerAttributes) TokenSigningPublicKeys() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ia.ref.Append("token_signing_public_keys"))
}

type iotAuthorizerState struct {
	Arn                    string            `json:"arn"`
	AuthorizerFunctionArn  string            `json:"authorizer_function_arn"`
	EnableCachingForHttp   bool              `json:"enable_caching_for_http"`
	Id                     string            `json:"id"`
	Name                   string            `json:"name"`
	SigningDisabled        bool              `json:"signing_disabled"`
	Status                 string            `json:"status"`
	TokenKeyName           string            `json:"token_key_name"`
	TokenSigningPublicKeys map[string]string `json:"token_signing_public_keys"`
}

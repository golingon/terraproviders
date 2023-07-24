// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotAuthorizer creates a new instance of [IotAuthorizer].
func NewIotAuthorizer(name string, args IotAuthorizerArgs) *IotAuthorizer {
	return &IotAuthorizer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotAuthorizer)(nil)

// IotAuthorizer represents the Terraform resource aws_iot_authorizer.
type IotAuthorizer struct {
	Name      string
	Args      IotAuthorizerArgs
	state     *iotAuthorizerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotAuthorizer].
func (ia *IotAuthorizer) Type() string {
	return "aws_iot_authorizer"
}

// LocalName returns the local name for [IotAuthorizer].
func (ia *IotAuthorizer) LocalName() string {
	return ia.Name
}

// Configuration returns the configuration (args) for [IotAuthorizer].
func (ia *IotAuthorizer) Configuration() interface{} {
	return ia.Args
}

// DependOn is used for other resources to depend on [IotAuthorizer].
func (ia *IotAuthorizer) DependOn() terra.Reference {
	return terra.ReferenceResource(ia)
}

// Dependencies returns the list of resources [IotAuthorizer] depends_on.
func (ia *IotAuthorizer) Dependencies() terra.Dependencies {
	return ia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotAuthorizer].
func (ia *IotAuthorizer) LifecycleManagement() *terra.Lifecycle {
	return ia.Lifecycle
}

// Attributes returns the attributes for [IotAuthorizer].
func (ia *IotAuthorizer) Attributes() iotAuthorizerAttributes {
	return iotAuthorizerAttributes{ref: terra.ReferenceResource(ia)}
}

// ImportState imports the given attribute values into [IotAuthorizer]'s state.
func (ia *IotAuthorizer) ImportState(av io.Reader) error {
	ia.state = &iotAuthorizerState{}
	if err := json.NewDecoder(av).Decode(ia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ia.Type(), ia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotAuthorizer] has state.
func (ia *IotAuthorizer) State() (*iotAuthorizerState, bool) {
	return ia.state, ia.state != nil
}

// StateMust returns the state for [IotAuthorizer]. Panics if the state is nil.
func (ia *IotAuthorizer) StateMust() *iotAuthorizerState {
	if ia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ia.Type(), ia.LocalName()))
	}
	return ia.state
}

// IotAuthorizerArgs contains the configurations for aws_iot_authorizer.
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
}
type iotAuthorizerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iot_authorizer.
func (ia iotAuthorizerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("arn"))
}

// AuthorizerFunctionArn returns a reference to field authorizer_function_arn of aws_iot_authorizer.
func (ia iotAuthorizerAttributes) AuthorizerFunctionArn() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("authorizer_function_arn"))
}

// EnableCachingForHttp returns a reference to field enable_caching_for_http of aws_iot_authorizer.
func (ia iotAuthorizerAttributes) EnableCachingForHttp() terra.BoolValue {
	return terra.ReferenceAsBool(ia.ref.Append("enable_caching_for_http"))
}

// Id returns a reference to field id of aws_iot_authorizer.
func (ia iotAuthorizerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("id"))
}

// Name returns a reference to field name of aws_iot_authorizer.
func (ia iotAuthorizerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("name"))
}

// SigningDisabled returns a reference to field signing_disabled of aws_iot_authorizer.
func (ia iotAuthorizerAttributes) SigningDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(ia.ref.Append("signing_disabled"))
}

// Status returns a reference to field status of aws_iot_authorizer.
func (ia iotAuthorizerAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("status"))
}

// TokenKeyName returns a reference to field token_key_name of aws_iot_authorizer.
func (ia iotAuthorizerAttributes) TokenKeyName() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("token_key_name"))
}

// TokenSigningPublicKeys returns a reference to field token_signing_public_keys of aws_iot_authorizer.
func (ia iotAuthorizerAttributes) TokenSigningPublicKeys() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ia.ref.Append("token_signing_public_keys"))
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

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewCloudfrontPublicKey creates a new instance of [CloudfrontPublicKey].
func NewCloudfrontPublicKey(name string, args CloudfrontPublicKeyArgs) *CloudfrontPublicKey {
	return &CloudfrontPublicKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontPublicKey)(nil)

// CloudfrontPublicKey represents the Terraform resource aws_cloudfront_public_key.
type CloudfrontPublicKey struct {
	Name      string
	Args      CloudfrontPublicKeyArgs
	state     *cloudfrontPublicKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfrontPublicKey].
func (cpk *CloudfrontPublicKey) Type() string {
	return "aws_cloudfront_public_key"
}

// LocalName returns the local name for [CloudfrontPublicKey].
func (cpk *CloudfrontPublicKey) LocalName() string {
	return cpk.Name
}

// Configuration returns the configuration (args) for [CloudfrontPublicKey].
func (cpk *CloudfrontPublicKey) Configuration() interface{} {
	return cpk.Args
}

// DependOn is used for other resources to depend on [CloudfrontPublicKey].
func (cpk *CloudfrontPublicKey) DependOn() terra.Reference {
	return terra.ReferenceResource(cpk)
}

// Dependencies returns the list of resources [CloudfrontPublicKey] depends_on.
func (cpk *CloudfrontPublicKey) Dependencies() terra.Dependencies {
	return cpk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfrontPublicKey].
func (cpk *CloudfrontPublicKey) LifecycleManagement() *terra.Lifecycle {
	return cpk.Lifecycle
}

// Attributes returns the attributes for [CloudfrontPublicKey].
func (cpk *CloudfrontPublicKey) Attributes() cloudfrontPublicKeyAttributes {
	return cloudfrontPublicKeyAttributes{ref: terra.ReferenceResource(cpk)}
}

// ImportState imports the given attribute values into [CloudfrontPublicKey]'s state.
func (cpk *CloudfrontPublicKey) ImportState(av io.Reader) error {
	cpk.state = &cloudfrontPublicKeyState{}
	if err := json.NewDecoder(av).Decode(cpk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpk.Type(), cpk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfrontPublicKey] has state.
func (cpk *CloudfrontPublicKey) State() (*cloudfrontPublicKeyState, bool) {
	return cpk.state, cpk.state != nil
}

// StateMust returns the state for [CloudfrontPublicKey]. Panics if the state is nil.
func (cpk *CloudfrontPublicKey) StateMust() *cloudfrontPublicKeyState {
	if cpk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpk.Type(), cpk.LocalName()))
	}
	return cpk.state
}

// CloudfrontPublicKeyArgs contains the configurations for aws_cloudfront_public_key.
type CloudfrontPublicKeyArgs struct {
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// EncodedKey: string, required
	EncodedKey terra.StringValue `hcl:"encoded_key,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
}
type cloudfrontPublicKeyAttributes struct {
	ref terra.Reference
}

// CallerReference returns a reference to field caller_reference of aws_cloudfront_public_key.
func (cpk cloudfrontPublicKeyAttributes) CallerReference() terra.StringValue {
	return terra.ReferenceAsString(cpk.ref.Append("caller_reference"))
}

// Comment returns a reference to field comment of aws_cloudfront_public_key.
func (cpk cloudfrontPublicKeyAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(cpk.ref.Append("comment"))
}

// EncodedKey returns a reference to field encoded_key of aws_cloudfront_public_key.
func (cpk cloudfrontPublicKeyAttributes) EncodedKey() terra.StringValue {
	return terra.ReferenceAsString(cpk.ref.Append("encoded_key"))
}

// Etag returns a reference to field etag of aws_cloudfront_public_key.
func (cpk cloudfrontPublicKeyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cpk.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_public_key.
func (cpk cloudfrontPublicKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpk.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudfront_public_key.
func (cpk cloudfrontPublicKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cpk.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_cloudfront_public_key.
func (cpk cloudfrontPublicKeyAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(cpk.ref.Append("name_prefix"))
}

type cloudfrontPublicKeyState struct {
	CallerReference string `json:"caller_reference"`
	Comment         string `json:"comment"`
	EncodedKey      string `json:"encoded_key"`
	Etag            string `json:"etag"`
	Id              string `json:"id"`
	Name            string `json:"name"`
	NamePrefix      string `json:"name_prefix"`
}

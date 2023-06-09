// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailBucketAccessKey creates a new instance of [LightsailBucketAccessKey].
func NewLightsailBucketAccessKey(name string, args LightsailBucketAccessKeyArgs) *LightsailBucketAccessKey {
	return &LightsailBucketAccessKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailBucketAccessKey)(nil)

// LightsailBucketAccessKey represents the Terraform resource aws_lightsail_bucket_access_key.
type LightsailBucketAccessKey struct {
	Name      string
	Args      LightsailBucketAccessKeyArgs
	state     *lightsailBucketAccessKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailBucketAccessKey].
func (lbak *LightsailBucketAccessKey) Type() string {
	return "aws_lightsail_bucket_access_key"
}

// LocalName returns the local name for [LightsailBucketAccessKey].
func (lbak *LightsailBucketAccessKey) LocalName() string {
	return lbak.Name
}

// Configuration returns the configuration (args) for [LightsailBucketAccessKey].
func (lbak *LightsailBucketAccessKey) Configuration() interface{} {
	return lbak.Args
}

// DependOn is used for other resources to depend on [LightsailBucketAccessKey].
func (lbak *LightsailBucketAccessKey) DependOn() terra.Reference {
	return terra.ReferenceResource(lbak)
}

// Dependencies returns the list of resources [LightsailBucketAccessKey] depends_on.
func (lbak *LightsailBucketAccessKey) Dependencies() terra.Dependencies {
	return lbak.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailBucketAccessKey].
func (lbak *LightsailBucketAccessKey) LifecycleManagement() *terra.Lifecycle {
	return lbak.Lifecycle
}

// Attributes returns the attributes for [LightsailBucketAccessKey].
func (lbak *LightsailBucketAccessKey) Attributes() lightsailBucketAccessKeyAttributes {
	return lightsailBucketAccessKeyAttributes{ref: terra.ReferenceResource(lbak)}
}

// ImportState imports the given attribute values into [LightsailBucketAccessKey]'s state.
func (lbak *LightsailBucketAccessKey) ImportState(av io.Reader) error {
	lbak.state = &lightsailBucketAccessKeyState{}
	if err := json.NewDecoder(av).Decode(lbak.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lbak.Type(), lbak.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailBucketAccessKey] has state.
func (lbak *LightsailBucketAccessKey) State() (*lightsailBucketAccessKeyState, bool) {
	return lbak.state, lbak.state != nil
}

// StateMust returns the state for [LightsailBucketAccessKey]. Panics if the state is nil.
func (lbak *LightsailBucketAccessKey) StateMust() *lightsailBucketAccessKeyState {
	if lbak.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lbak.Type(), lbak.LocalName()))
	}
	return lbak.state
}

// LightsailBucketAccessKeyArgs contains the configurations for aws_lightsail_bucket_access_key.
type LightsailBucketAccessKeyArgs struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type lightsailBucketAccessKeyAttributes struct {
	ref terra.Reference
}

// AccessKeyId returns a reference to field access_key_id of aws_lightsail_bucket_access_key.
func (lbak lightsailBucketAccessKeyAttributes) AccessKeyId() terra.StringValue {
	return terra.ReferenceAsString(lbak.ref.Append("access_key_id"))
}

// BucketName returns a reference to field bucket_name of aws_lightsail_bucket_access_key.
func (lbak lightsailBucketAccessKeyAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(lbak.ref.Append("bucket_name"))
}

// CreatedAt returns a reference to field created_at of aws_lightsail_bucket_access_key.
func (lbak lightsailBucketAccessKeyAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(lbak.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_lightsail_bucket_access_key.
func (lbak lightsailBucketAccessKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lbak.ref.Append("id"))
}

// SecretAccessKey returns a reference to field secret_access_key of aws_lightsail_bucket_access_key.
func (lbak lightsailBucketAccessKeyAttributes) SecretAccessKey() terra.StringValue {
	return terra.ReferenceAsString(lbak.ref.Append("secret_access_key"))
}

// Status returns a reference to field status of aws_lightsail_bucket_access_key.
func (lbak lightsailBucketAccessKeyAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(lbak.ref.Append("status"))
}

type lightsailBucketAccessKeyState struct {
	AccessKeyId     string `json:"access_key_id"`
	BucketName      string `json:"bucket_name"`
	CreatedAt       string `json:"created_at"`
	Id              string `json:"id"`
	SecretAccessKey string `json:"secret_access_key"`
	Status          string `json:"status"`
}

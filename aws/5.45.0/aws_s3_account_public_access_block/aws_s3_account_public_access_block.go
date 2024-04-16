// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_s3_account_public_access_block

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_s3_account_public_access_block.
type Resource struct {
	Name      string
	Args      Args
	state     *awsS3AccountPublicAccessBlockState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asapab *Resource) Type() string {
	return "aws_s3_account_public_access_block"
}

// LocalName returns the local name for [Resource].
func (asapab *Resource) LocalName() string {
	return asapab.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asapab *Resource) Configuration() interface{} {
	return asapab.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asapab *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asapab)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asapab *Resource) Dependencies() terra.Dependencies {
	return asapab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asapab *Resource) LifecycleManagement() *terra.Lifecycle {
	return asapab.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asapab *Resource) Attributes() awsS3AccountPublicAccessBlockAttributes {
	return awsS3AccountPublicAccessBlockAttributes{ref: terra.ReferenceResource(asapab)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asapab *Resource) ImportState(state io.Reader) error {
	asapab.state = &awsS3AccountPublicAccessBlockState{}
	if err := json.NewDecoder(state).Decode(asapab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asapab.Type(), asapab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asapab *Resource) State() (*awsS3AccountPublicAccessBlockState, bool) {
	return asapab.state, asapab.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asapab *Resource) StateMust() *awsS3AccountPublicAccessBlockState {
	if asapab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asapab.Type(), asapab.LocalName()))
	}
	return asapab.state
}

// Args contains the configurations for aws_s3_account_public_access_block.
type Args struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// BlockPublicAcls: bool, optional
	BlockPublicAcls terra.BoolValue `hcl:"block_public_acls,attr"`
	// BlockPublicPolicy: bool, optional
	BlockPublicPolicy terra.BoolValue `hcl:"block_public_policy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnorePublicAcls: bool, optional
	IgnorePublicAcls terra.BoolValue `hcl:"ignore_public_acls,attr"`
	// RestrictPublicBuckets: bool, optional
	RestrictPublicBuckets terra.BoolValue `hcl:"restrict_public_buckets,attr"`
}

type awsS3AccountPublicAccessBlockAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_s3_account_public_access_block.
func (asapab awsS3AccountPublicAccessBlockAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(asapab.ref.Append("account_id"))
}

// BlockPublicAcls returns a reference to field block_public_acls of aws_s3_account_public_access_block.
func (asapab awsS3AccountPublicAccessBlockAttributes) BlockPublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(asapab.ref.Append("block_public_acls"))
}

// BlockPublicPolicy returns a reference to field block_public_policy of aws_s3_account_public_access_block.
func (asapab awsS3AccountPublicAccessBlockAttributes) BlockPublicPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(asapab.ref.Append("block_public_policy"))
}

// Id returns a reference to field id of aws_s3_account_public_access_block.
func (asapab awsS3AccountPublicAccessBlockAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asapab.ref.Append("id"))
}

// IgnorePublicAcls returns a reference to field ignore_public_acls of aws_s3_account_public_access_block.
func (asapab awsS3AccountPublicAccessBlockAttributes) IgnorePublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(asapab.ref.Append("ignore_public_acls"))
}

// RestrictPublicBuckets returns a reference to field restrict_public_buckets of aws_s3_account_public_access_block.
func (asapab awsS3AccountPublicAccessBlockAttributes) RestrictPublicBuckets() terra.BoolValue {
	return terra.ReferenceAsBool(asapab.ref.Append("restrict_public_buckets"))
}

type awsS3AccountPublicAccessBlockState struct {
	AccountId             string `json:"account_id"`
	BlockPublicAcls       bool   `json:"block_public_acls"`
	BlockPublicPolicy     bool   `json:"block_public_policy"`
	Id                    string `json:"id"`
	IgnorePublicAcls      bool   `json:"ignore_public_acls"`
	RestrictPublicBuckets bool   `json:"restrict_public_buckets"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3AccountPublicAccessBlock creates a new instance of [S3AccountPublicAccessBlock].
func NewS3AccountPublicAccessBlock(name string, args S3AccountPublicAccessBlockArgs) *S3AccountPublicAccessBlock {
	return &S3AccountPublicAccessBlock{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3AccountPublicAccessBlock)(nil)

// S3AccountPublicAccessBlock represents the Terraform resource aws_s3_account_public_access_block.
type S3AccountPublicAccessBlock struct {
	Name      string
	Args      S3AccountPublicAccessBlockArgs
	state     *s3AccountPublicAccessBlockState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3AccountPublicAccessBlock].
func (sapab *S3AccountPublicAccessBlock) Type() string {
	return "aws_s3_account_public_access_block"
}

// LocalName returns the local name for [S3AccountPublicAccessBlock].
func (sapab *S3AccountPublicAccessBlock) LocalName() string {
	return sapab.Name
}

// Configuration returns the configuration (args) for [S3AccountPublicAccessBlock].
func (sapab *S3AccountPublicAccessBlock) Configuration() interface{} {
	return sapab.Args
}

// DependOn is used for other resources to depend on [S3AccountPublicAccessBlock].
func (sapab *S3AccountPublicAccessBlock) DependOn() terra.Reference {
	return terra.ReferenceResource(sapab)
}

// Dependencies returns the list of resources [S3AccountPublicAccessBlock] depends_on.
func (sapab *S3AccountPublicAccessBlock) Dependencies() terra.Dependencies {
	return sapab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3AccountPublicAccessBlock].
func (sapab *S3AccountPublicAccessBlock) LifecycleManagement() *terra.Lifecycle {
	return sapab.Lifecycle
}

// Attributes returns the attributes for [S3AccountPublicAccessBlock].
func (sapab *S3AccountPublicAccessBlock) Attributes() s3AccountPublicAccessBlockAttributes {
	return s3AccountPublicAccessBlockAttributes{ref: terra.ReferenceResource(sapab)}
}

// ImportState imports the given attribute values into [S3AccountPublicAccessBlock]'s state.
func (sapab *S3AccountPublicAccessBlock) ImportState(av io.Reader) error {
	sapab.state = &s3AccountPublicAccessBlockState{}
	if err := json.NewDecoder(av).Decode(sapab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sapab.Type(), sapab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3AccountPublicAccessBlock] has state.
func (sapab *S3AccountPublicAccessBlock) State() (*s3AccountPublicAccessBlockState, bool) {
	return sapab.state, sapab.state != nil
}

// StateMust returns the state for [S3AccountPublicAccessBlock]. Panics if the state is nil.
func (sapab *S3AccountPublicAccessBlock) StateMust() *s3AccountPublicAccessBlockState {
	if sapab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sapab.Type(), sapab.LocalName()))
	}
	return sapab.state
}

// S3AccountPublicAccessBlockArgs contains the configurations for aws_s3_account_public_access_block.
type S3AccountPublicAccessBlockArgs struct {
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
type s3AccountPublicAccessBlockAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_s3_account_public_access_block.
func (sapab s3AccountPublicAccessBlockAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(sapab.ref.Append("account_id"))
}

// BlockPublicAcls returns a reference to field block_public_acls of aws_s3_account_public_access_block.
func (sapab s3AccountPublicAccessBlockAttributes) BlockPublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(sapab.ref.Append("block_public_acls"))
}

// BlockPublicPolicy returns a reference to field block_public_policy of aws_s3_account_public_access_block.
func (sapab s3AccountPublicAccessBlockAttributes) BlockPublicPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(sapab.ref.Append("block_public_policy"))
}

// Id returns a reference to field id of aws_s3_account_public_access_block.
func (sapab s3AccountPublicAccessBlockAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sapab.ref.Append("id"))
}

// IgnorePublicAcls returns a reference to field ignore_public_acls of aws_s3_account_public_access_block.
func (sapab s3AccountPublicAccessBlockAttributes) IgnorePublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(sapab.ref.Append("ignore_public_acls"))
}

// RestrictPublicBuckets returns a reference to field restrict_public_buckets of aws_s3_account_public_access_block.
func (sapab s3AccountPublicAccessBlockAttributes) RestrictPublicBuckets() terra.BoolValue {
	return terra.ReferenceAsBool(sapab.ref.Append("restrict_public_buckets"))
}

type s3AccountPublicAccessBlockState struct {
	AccountId             string `json:"account_id"`
	BlockPublicAcls       bool   `json:"block_public_acls"`
	BlockPublicPolicy     bool   `json:"block_public_policy"`
	Id                    string `json:"id"`
	IgnorePublicAcls      bool   `json:"ignore_public_acls"`
	RestrictPublicBuckets bool   `json:"restrict_public_buckets"`
}
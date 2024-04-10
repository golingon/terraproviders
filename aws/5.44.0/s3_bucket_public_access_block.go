// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewS3BucketPublicAccessBlock creates a new instance of [S3BucketPublicAccessBlock].
func NewS3BucketPublicAccessBlock(name string, args S3BucketPublicAccessBlockArgs) *S3BucketPublicAccessBlock {
	return &S3BucketPublicAccessBlock{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketPublicAccessBlock)(nil)

// S3BucketPublicAccessBlock represents the Terraform resource aws_s3_bucket_public_access_block.
type S3BucketPublicAccessBlock struct {
	Name      string
	Args      S3BucketPublicAccessBlockArgs
	state     *s3BucketPublicAccessBlockState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketPublicAccessBlock].
func (sbpab *S3BucketPublicAccessBlock) Type() string {
	return "aws_s3_bucket_public_access_block"
}

// LocalName returns the local name for [S3BucketPublicAccessBlock].
func (sbpab *S3BucketPublicAccessBlock) LocalName() string {
	return sbpab.Name
}

// Configuration returns the configuration (args) for [S3BucketPublicAccessBlock].
func (sbpab *S3BucketPublicAccessBlock) Configuration() interface{} {
	return sbpab.Args
}

// DependOn is used for other resources to depend on [S3BucketPublicAccessBlock].
func (sbpab *S3BucketPublicAccessBlock) DependOn() terra.Reference {
	return terra.ReferenceResource(sbpab)
}

// Dependencies returns the list of resources [S3BucketPublicAccessBlock] depends_on.
func (sbpab *S3BucketPublicAccessBlock) Dependencies() terra.Dependencies {
	return sbpab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketPublicAccessBlock].
func (sbpab *S3BucketPublicAccessBlock) LifecycleManagement() *terra.Lifecycle {
	return sbpab.Lifecycle
}

// Attributes returns the attributes for [S3BucketPublicAccessBlock].
func (sbpab *S3BucketPublicAccessBlock) Attributes() s3BucketPublicAccessBlockAttributes {
	return s3BucketPublicAccessBlockAttributes{ref: terra.ReferenceResource(sbpab)}
}

// ImportState imports the given attribute values into [S3BucketPublicAccessBlock]'s state.
func (sbpab *S3BucketPublicAccessBlock) ImportState(av io.Reader) error {
	sbpab.state = &s3BucketPublicAccessBlockState{}
	if err := json.NewDecoder(av).Decode(sbpab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbpab.Type(), sbpab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketPublicAccessBlock] has state.
func (sbpab *S3BucketPublicAccessBlock) State() (*s3BucketPublicAccessBlockState, bool) {
	return sbpab.state, sbpab.state != nil
}

// StateMust returns the state for [S3BucketPublicAccessBlock]. Panics if the state is nil.
func (sbpab *S3BucketPublicAccessBlock) StateMust() *s3BucketPublicAccessBlockState {
	if sbpab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbpab.Type(), sbpab.LocalName()))
	}
	return sbpab.state
}

// S3BucketPublicAccessBlockArgs contains the configurations for aws_s3_bucket_public_access_block.
type S3BucketPublicAccessBlockArgs struct {
	// BlockPublicAcls: bool, optional
	BlockPublicAcls terra.BoolValue `hcl:"block_public_acls,attr"`
	// BlockPublicPolicy: bool, optional
	BlockPublicPolicy terra.BoolValue `hcl:"block_public_policy,attr"`
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnorePublicAcls: bool, optional
	IgnorePublicAcls terra.BoolValue `hcl:"ignore_public_acls,attr"`
	// RestrictPublicBuckets: bool, optional
	RestrictPublicBuckets terra.BoolValue `hcl:"restrict_public_buckets,attr"`
}
type s3BucketPublicAccessBlockAttributes struct {
	ref terra.Reference
}

// BlockPublicAcls returns a reference to field block_public_acls of aws_s3_bucket_public_access_block.
func (sbpab s3BucketPublicAccessBlockAttributes) BlockPublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(sbpab.ref.Append("block_public_acls"))
}

// BlockPublicPolicy returns a reference to field block_public_policy of aws_s3_bucket_public_access_block.
func (sbpab s3BucketPublicAccessBlockAttributes) BlockPublicPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(sbpab.ref.Append("block_public_policy"))
}

// Bucket returns a reference to field bucket of aws_s3_bucket_public_access_block.
func (sbpab s3BucketPublicAccessBlockAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbpab.ref.Append("bucket"))
}

// Id returns a reference to field id of aws_s3_bucket_public_access_block.
func (sbpab s3BucketPublicAccessBlockAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbpab.ref.Append("id"))
}

// IgnorePublicAcls returns a reference to field ignore_public_acls of aws_s3_bucket_public_access_block.
func (sbpab s3BucketPublicAccessBlockAttributes) IgnorePublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(sbpab.ref.Append("ignore_public_acls"))
}

// RestrictPublicBuckets returns a reference to field restrict_public_buckets of aws_s3_bucket_public_access_block.
func (sbpab s3BucketPublicAccessBlockAttributes) RestrictPublicBuckets() terra.BoolValue {
	return terra.ReferenceAsBool(sbpab.ref.Append("restrict_public_buckets"))
}

type s3BucketPublicAccessBlockState struct {
	BlockPublicAcls       bool   `json:"block_public_acls"`
	BlockPublicPolicy     bool   `json:"block_public_policy"`
	Bucket                string `json:"bucket"`
	Id                    string `json:"id"`
	IgnorePublicAcls      bool   `json:"ignore_public_acls"`
	RestrictPublicBuckets bool   `json:"restrict_public_buckets"`
}

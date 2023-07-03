// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucketacl "github.com/golingon/terraproviders/aws/5.6.2/s3bucketacl"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3BucketAcl creates a new instance of [S3BucketAcl].
func NewS3BucketAcl(name string, args S3BucketAclArgs) *S3BucketAcl {
	return &S3BucketAcl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketAcl)(nil)

// S3BucketAcl represents the Terraform resource aws_s3_bucket_acl.
type S3BucketAcl struct {
	Name      string
	Args      S3BucketAclArgs
	state     *s3BucketAclState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketAcl].
func (sba *S3BucketAcl) Type() string {
	return "aws_s3_bucket_acl"
}

// LocalName returns the local name for [S3BucketAcl].
func (sba *S3BucketAcl) LocalName() string {
	return sba.Name
}

// Configuration returns the configuration (args) for [S3BucketAcl].
func (sba *S3BucketAcl) Configuration() interface{} {
	return sba.Args
}

// DependOn is used for other resources to depend on [S3BucketAcl].
func (sba *S3BucketAcl) DependOn() terra.Reference {
	return terra.ReferenceResource(sba)
}

// Dependencies returns the list of resources [S3BucketAcl] depends_on.
func (sba *S3BucketAcl) Dependencies() terra.Dependencies {
	return sba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketAcl].
func (sba *S3BucketAcl) LifecycleManagement() *terra.Lifecycle {
	return sba.Lifecycle
}

// Attributes returns the attributes for [S3BucketAcl].
func (sba *S3BucketAcl) Attributes() s3BucketAclAttributes {
	return s3BucketAclAttributes{ref: terra.ReferenceResource(sba)}
}

// ImportState imports the given attribute values into [S3BucketAcl]'s state.
func (sba *S3BucketAcl) ImportState(av io.Reader) error {
	sba.state = &s3BucketAclState{}
	if err := json.NewDecoder(av).Decode(sba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sba.Type(), sba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketAcl] has state.
func (sba *S3BucketAcl) State() (*s3BucketAclState, bool) {
	return sba.state, sba.state != nil
}

// StateMust returns the state for [S3BucketAcl]. Panics if the state is nil.
func (sba *S3BucketAcl) StateMust() *s3BucketAclState {
	if sba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sba.Type(), sba.LocalName()))
	}
	return sba.state
}

// S3BucketAclArgs contains the configurations for aws_s3_bucket_acl.
type S3BucketAclArgs struct {
	// Acl: string, optional
	Acl terra.StringValue `hcl:"acl,attr"`
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// ExpectedBucketOwner: string, optional
	ExpectedBucketOwner terra.StringValue `hcl:"expected_bucket_owner,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// AccessControlPolicy: optional
	AccessControlPolicy *s3bucketacl.AccessControlPolicy `hcl:"access_control_policy,block"`
}
type s3BucketAclAttributes struct {
	ref terra.Reference
}

// Acl returns a reference to field acl of aws_s3_bucket_acl.
func (sba s3BucketAclAttributes) Acl() terra.StringValue {
	return terra.ReferenceAsString(sba.ref.Append("acl"))
}

// Bucket returns a reference to field bucket of aws_s3_bucket_acl.
func (sba s3BucketAclAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sba.ref.Append("bucket"))
}

// ExpectedBucketOwner returns a reference to field expected_bucket_owner of aws_s3_bucket_acl.
func (sba s3BucketAclAttributes) ExpectedBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(sba.ref.Append("expected_bucket_owner"))
}

// Id returns a reference to field id of aws_s3_bucket_acl.
func (sba s3BucketAclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sba.ref.Append("id"))
}

func (sba s3BucketAclAttributes) AccessControlPolicy() terra.ListValue[s3bucketacl.AccessControlPolicyAttributes] {
	return terra.ReferenceAsList[s3bucketacl.AccessControlPolicyAttributes](sba.ref.Append("access_control_policy"))
}

type s3BucketAclState struct {
	Acl                 string                                 `json:"acl"`
	Bucket              string                                 `json:"bucket"`
	ExpectedBucketOwner string                                 `json:"expected_bucket_owner"`
	Id                  string                                 `json:"id"`
	AccessControlPolicy []s3bucketacl.AccessControlPolicyState `json:"access_control_policy"`
}

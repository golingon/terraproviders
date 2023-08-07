// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucketintelligenttieringconfiguration "github.com/golingon/terraproviders/aws/5.11.0/s3bucketintelligenttieringconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3BucketIntelligentTieringConfiguration creates a new instance of [S3BucketIntelligentTieringConfiguration].
func NewS3BucketIntelligentTieringConfiguration(name string, args S3BucketIntelligentTieringConfigurationArgs) *S3BucketIntelligentTieringConfiguration {
	return &S3BucketIntelligentTieringConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketIntelligentTieringConfiguration)(nil)

// S3BucketIntelligentTieringConfiguration represents the Terraform resource aws_s3_bucket_intelligent_tiering_configuration.
type S3BucketIntelligentTieringConfiguration struct {
	Name      string
	Args      S3BucketIntelligentTieringConfigurationArgs
	state     *s3BucketIntelligentTieringConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketIntelligentTieringConfiguration].
func (sbitc *S3BucketIntelligentTieringConfiguration) Type() string {
	return "aws_s3_bucket_intelligent_tiering_configuration"
}

// LocalName returns the local name for [S3BucketIntelligentTieringConfiguration].
func (sbitc *S3BucketIntelligentTieringConfiguration) LocalName() string {
	return sbitc.Name
}

// Configuration returns the configuration (args) for [S3BucketIntelligentTieringConfiguration].
func (sbitc *S3BucketIntelligentTieringConfiguration) Configuration() interface{} {
	return sbitc.Args
}

// DependOn is used for other resources to depend on [S3BucketIntelligentTieringConfiguration].
func (sbitc *S3BucketIntelligentTieringConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(sbitc)
}

// Dependencies returns the list of resources [S3BucketIntelligentTieringConfiguration] depends_on.
func (sbitc *S3BucketIntelligentTieringConfiguration) Dependencies() terra.Dependencies {
	return sbitc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketIntelligentTieringConfiguration].
func (sbitc *S3BucketIntelligentTieringConfiguration) LifecycleManagement() *terra.Lifecycle {
	return sbitc.Lifecycle
}

// Attributes returns the attributes for [S3BucketIntelligentTieringConfiguration].
func (sbitc *S3BucketIntelligentTieringConfiguration) Attributes() s3BucketIntelligentTieringConfigurationAttributes {
	return s3BucketIntelligentTieringConfigurationAttributes{ref: terra.ReferenceResource(sbitc)}
}

// ImportState imports the given attribute values into [S3BucketIntelligentTieringConfiguration]'s state.
func (sbitc *S3BucketIntelligentTieringConfiguration) ImportState(av io.Reader) error {
	sbitc.state = &s3BucketIntelligentTieringConfigurationState{}
	if err := json.NewDecoder(av).Decode(sbitc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbitc.Type(), sbitc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketIntelligentTieringConfiguration] has state.
func (sbitc *S3BucketIntelligentTieringConfiguration) State() (*s3BucketIntelligentTieringConfigurationState, bool) {
	return sbitc.state, sbitc.state != nil
}

// StateMust returns the state for [S3BucketIntelligentTieringConfiguration]. Panics if the state is nil.
func (sbitc *S3BucketIntelligentTieringConfiguration) StateMust() *s3BucketIntelligentTieringConfigurationState {
	if sbitc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbitc.Type(), sbitc.LocalName()))
	}
	return sbitc.state
}

// S3BucketIntelligentTieringConfigurationArgs contains the configurations for aws_s3_bucket_intelligent_tiering_configuration.
type S3BucketIntelligentTieringConfigurationArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// Filter: optional
	Filter *s3bucketintelligenttieringconfiguration.Filter `hcl:"filter,block"`
	// Tiering: min=1
	Tiering []s3bucketintelligenttieringconfiguration.Tiering `hcl:"tiering,block" validate:"min=1"`
}
type s3BucketIntelligentTieringConfigurationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_intelligent_tiering_configuration.
func (sbitc s3BucketIntelligentTieringConfigurationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbitc.ref.Append("bucket"))
}

// Id returns a reference to field id of aws_s3_bucket_intelligent_tiering_configuration.
func (sbitc s3BucketIntelligentTieringConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbitc.ref.Append("id"))
}

// Name returns a reference to field name of aws_s3_bucket_intelligent_tiering_configuration.
func (sbitc s3BucketIntelligentTieringConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sbitc.ref.Append("name"))
}

// Status returns a reference to field status of aws_s3_bucket_intelligent_tiering_configuration.
func (sbitc s3BucketIntelligentTieringConfigurationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(sbitc.ref.Append("status"))
}

func (sbitc s3BucketIntelligentTieringConfigurationAttributes) Filter() terra.ListValue[s3bucketintelligenttieringconfiguration.FilterAttributes] {
	return terra.ReferenceAsList[s3bucketintelligenttieringconfiguration.FilterAttributes](sbitc.ref.Append("filter"))
}

func (sbitc s3BucketIntelligentTieringConfigurationAttributes) Tiering() terra.SetValue[s3bucketintelligenttieringconfiguration.TieringAttributes] {
	return terra.ReferenceAsSet[s3bucketintelligenttieringconfiguration.TieringAttributes](sbitc.ref.Append("tiering"))
}

type s3BucketIntelligentTieringConfigurationState struct {
	Bucket  string                                                 `json:"bucket"`
	Id      string                                                 `json:"id"`
	Name    string                                                 `json:"name"`
	Status  string                                                 `json:"status"`
	Filter  []s3bucketintelligenttieringconfiguration.FilterState  `json:"filter"`
	Tiering []s3bucketintelligenttieringconfiguration.TieringState `json:"tiering"`
}

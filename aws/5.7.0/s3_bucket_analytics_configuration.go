// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucketanalyticsconfiguration "github.com/golingon/terraproviders/aws/5.7.0/s3bucketanalyticsconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3BucketAnalyticsConfiguration creates a new instance of [S3BucketAnalyticsConfiguration].
func NewS3BucketAnalyticsConfiguration(name string, args S3BucketAnalyticsConfigurationArgs) *S3BucketAnalyticsConfiguration {
	return &S3BucketAnalyticsConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketAnalyticsConfiguration)(nil)

// S3BucketAnalyticsConfiguration represents the Terraform resource aws_s3_bucket_analytics_configuration.
type S3BucketAnalyticsConfiguration struct {
	Name      string
	Args      S3BucketAnalyticsConfigurationArgs
	state     *s3BucketAnalyticsConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketAnalyticsConfiguration].
func (sbac *S3BucketAnalyticsConfiguration) Type() string {
	return "aws_s3_bucket_analytics_configuration"
}

// LocalName returns the local name for [S3BucketAnalyticsConfiguration].
func (sbac *S3BucketAnalyticsConfiguration) LocalName() string {
	return sbac.Name
}

// Configuration returns the configuration (args) for [S3BucketAnalyticsConfiguration].
func (sbac *S3BucketAnalyticsConfiguration) Configuration() interface{} {
	return sbac.Args
}

// DependOn is used for other resources to depend on [S3BucketAnalyticsConfiguration].
func (sbac *S3BucketAnalyticsConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(sbac)
}

// Dependencies returns the list of resources [S3BucketAnalyticsConfiguration] depends_on.
func (sbac *S3BucketAnalyticsConfiguration) Dependencies() terra.Dependencies {
	return sbac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketAnalyticsConfiguration].
func (sbac *S3BucketAnalyticsConfiguration) LifecycleManagement() *terra.Lifecycle {
	return sbac.Lifecycle
}

// Attributes returns the attributes for [S3BucketAnalyticsConfiguration].
func (sbac *S3BucketAnalyticsConfiguration) Attributes() s3BucketAnalyticsConfigurationAttributes {
	return s3BucketAnalyticsConfigurationAttributes{ref: terra.ReferenceResource(sbac)}
}

// ImportState imports the given attribute values into [S3BucketAnalyticsConfiguration]'s state.
func (sbac *S3BucketAnalyticsConfiguration) ImportState(av io.Reader) error {
	sbac.state = &s3BucketAnalyticsConfigurationState{}
	if err := json.NewDecoder(av).Decode(sbac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbac.Type(), sbac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketAnalyticsConfiguration] has state.
func (sbac *S3BucketAnalyticsConfiguration) State() (*s3BucketAnalyticsConfigurationState, bool) {
	return sbac.state, sbac.state != nil
}

// StateMust returns the state for [S3BucketAnalyticsConfiguration]. Panics if the state is nil.
func (sbac *S3BucketAnalyticsConfiguration) StateMust() *s3BucketAnalyticsConfigurationState {
	if sbac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbac.Type(), sbac.LocalName()))
	}
	return sbac.state
}

// S3BucketAnalyticsConfigurationArgs contains the configurations for aws_s3_bucket_analytics_configuration.
type S3BucketAnalyticsConfigurationArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Filter: optional
	Filter *s3bucketanalyticsconfiguration.Filter `hcl:"filter,block"`
	// StorageClassAnalysis: optional
	StorageClassAnalysis *s3bucketanalyticsconfiguration.StorageClassAnalysis `hcl:"storage_class_analysis,block"`
}
type s3BucketAnalyticsConfigurationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_analytics_configuration.
func (sbac s3BucketAnalyticsConfigurationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbac.ref.Append("bucket"))
}

// Id returns a reference to field id of aws_s3_bucket_analytics_configuration.
func (sbac s3BucketAnalyticsConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbac.ref.Append("id"))
}

// Name returns a reference to field name of aws_s3_bucket_analytics_configuration.
func (sbac s3BucketAnalyticsConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sbac.ref.Append("name"))
}

func (sbac s3BucketAnalyticsConfigurationAttributes) Filter() terra.ListValue[s3bucketanalyticsconfiguration.FilterAttributes] {
	return terra.ReferenceAsList[s3bucketanalyticsconfiguration.FilterAttributes](sbac.ref.Append("filter"))
}

func (sbac s3BucketAnalyticsConfigurationAttributes) StorageClassAnalysis() terra.ListValue[s3bucketanalyticsconfiguration.StorageClassAnalysisAttributes] {
	return terra.ReferenceAsList[s3bucketanalyticsconfiguration.StorageClassAnalysisAttributes](sbac.ref.Append("storage_class_analysis"))
}

type s3BucketAnalyticsConfigurationState struct {
	Bucket               string                                                     `json:"bucket"`
	Id                   string                                                     `json:"id"`
	Name                 string                                                     `json:"name"`
	Filter               []s3bucketanalyticsconfiguration.FilterState               `json:"filter"`
	StorageClassAnalysis []s3bucketanalyticsconfiguration.StorageClassAnalysisState `json:"storage_class_analysis"`
}

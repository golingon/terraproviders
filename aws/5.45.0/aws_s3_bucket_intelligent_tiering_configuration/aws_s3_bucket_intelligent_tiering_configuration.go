// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_s3_bucket_intelligent_tiering_configuration

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

// Resource represents the Terraform resource aws_s3_bucket_intelligent_tiering_configuration.
type Resource struct {
	Name      string
	Args      Args
	state     *awsS3BucketIntelligentTieringConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asbitc *Resource) Type() string {
	return "aws_s3_bucket_intelligent_tiering_configuration"
}

// LocalName returns the local name for [Resource].
func (asbitc *Resource) LocalName() string {
	return asbitc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asbitc *Resource) Configuration() interface{} {
	return asbitc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asbitc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asbitc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asbitc *Resource) Dependencies() terra.Dependencies {
	return asbitc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asbitc *Resource) LifecycleManagement() *terra.Lifecycle {
	return asbitc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asbitc *Resource) Attributes() awsS3BucketIntelligentTieringConfigurationAttributes {
	return awsS3BucketIntelligentTieringConfigurationAttributes{ref: terra.ReferenceResource(asbitc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asbitc *Resource) ImportState(state io.Reader) error {
	asbitc.state = &awsS3BucketIntelligentTieringConfigurationState{}
	if err := json.NewDecoder(state).Decode(asbitc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asbitc.Type(), asbitc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asbitc *Resource) State() (*awsS3BucketIntelligentTieringConfigurationState, bool) {
	return asbitc.state, asbitc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asbitc *Resource) StateMust() *awsS3BucketIntelligentTieringConfigurationState {
	if asbitc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asbitc.Type(), asbitc.LocalName()))
	}
	return asbitc.state
}

// Args contains the configurations for aws_s3_bucket_intelligent_tiering_configuration.
type Args struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// Filter: optional
	Filter *Filter `hcl:"filter,block"`
	// Tiering: min=1
	Tiering []Tiering `hcl:"tiering,block" validate:"min=1"`
}

type awsS3BucketIntelligentTieringConfigurationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_intelligent_tiering_configuration.
func (asbitc awsS3BucketIntelligentTieringConfigurationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(asbitc.ref.Append("bucket"))
}

// Id returns a reference to field id of aws_s3_bucket_intelligent_tiering_configuration.
func (asbitc awsS3BucketIntelligentTieringConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asbitc.ref.Append("id"))
}

// Name returns a reference to field name of aws_s3_bucket_intelligent_tiering_configuration.
func (asbitc awsS3BucketIntelligentTieringConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asbitc.ref.Append("name"))
}

// Status returns a reference to field status of aws_s3_bucket_intelligent_tiering_configuration.
func (asbitc awsS3BucketIntelligentTieringConfigurationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(asbitc.ref.Append("status"))
}

func (asbitc awsS3BucketIntelligentTieringConfigurationAttributes) Filter() terra.ListValue[FilterAttributes] {
	return terra.ReferenceAsList[FilterAttributes](asbitc.ref.Append("filter"))
}

func (asbitc awsS3BucketIntelligentTieringConfigurationAttributes) Tiering() terra.SetValue[TieringAttributes] {
	return terra.ReferenceAsSet[TieringAttributes](asbitc.ref.Append("tiering"))
}

type awsS3BucketIntelligentTieringConfigurationState struct {
	Bucket  string         `json:"bucket"`
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Status  string         `json:"status"`
	Filter  []FilterState  `json:"filter"`
	Tiering []TieringState `json:"tiering"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	loggingmetric "github.com/golingon/terraproviders/googlebeta/4.71.0/loggingmetric"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingMetric creates a new instance of [LoggingMetric].
func NewLoggingMetric(name string, args LoggingMetricArgs) *LoggingMetric {
	return &LoggingMetric{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingMetric)(nil)

// LoggingMetric represents the Terraform resource google_logging_metric.
type LoggingMetric struct {
	Name      string
	Args      LoggingMetricArgs
	state     *loggingMetricState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingMetric].
func (lm *LoggingMetric) Type() string {
	return "google_logging_metric"
}

// LocalName returns the local name for [LoggingMetric].
func (lm *LoggingMetric) LocalName() string {
	return lm.Name
}

// Configuration returns the configuration (args) for [LoggingMetric].
func (lm *LoggingMetric) Configuration() interface{} {
	return lm.Args
}

// DependOn is used for other resources to depend on [LoggingMetric].
func (lm *LoggingMetric) DependOn() terra.Reference {
	return terra.ReferenceResource(lm)
}

// Dependencies returns the list of resources [LoggingMetric] depends_on.
func (lm *LoggingMetric) Dependencies() terra.Dependencies {
	return lm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingMetric].
func (lm *LoggingMetric) LifecycleManagement() *terra.Lifecycle {
	return lm.Lifecycle
}

// Attributes returns the attributes for [LoggingMetric].
func (lm *LoggingMetric) Attributes() loggingMetricAttributes {
	return loggingMetricAttributes{ref: terra.ReferenceResource(lm)}
}

// ImportState imports the given attribute values into [LoggingMetric]'s state.
func (lm *LoggingMetric) ImportState(av io.Reader) error {
	lm.state = &loggingMetricState{}
	if err := json.NewDecoder(av).Decode(lm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lm.Type(), lm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingMetric] has state.
func (lm *LoggingMetric) State() (*loggingMetricState, bool) {
	return lm.state, lm.state != nil
}

// StateMust returns the state for [LoggingMetric]. Panics if the state is nil.
func (lm *LoggingMetric) StateMust() *loggingMetricState {
	if lm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lm.Type(), lm.LocalName()))
	}
	return lm.state
}

// LoggingMetricArgs contains the configurations for google_logging_metric.
type LoggingMetricArgs struct {
	// BucketName: string, optional
	BucketName terra.StringValue `hcl:"bucket_name,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Filter: string, required
	Filter terra.StringValue `hcl:"filter,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LabelExtractors: map of string, optional
	LabelExtractors terra.MapValue[terra.StringValue] `hcl:"label_extractors,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ValueExtractor: string, optional
	ValueExtractor terra.StringValue `hcl:"value_extractor,attr"`
	// BucketOptions: optional
	BucketOptions *loggingmetric.BucketOptions `hcl:"bucket_options,block"`
	// MetricDescriptor: optional
	MetricDescriptor *loggingmetric.MetricDescriptor `hcl:"metric_descriptor,block"`
	// Timeouts: optional
	Timeouts *loggingmetric.Timeouts `hcl:"timeouts,block"`
}
type loggingMetricAttributes struct {
	ref terra.Reference
}

// BucketName returns a reference to field bucket_name of google_logging_metric.
func (lm loggingMetricAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("bucket_name"))
}

// Description returns a reference to field description of google_logging_metric.
func (lm loggingMetricAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_logging_metric.
func (lm loggingMetricAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(lm.ref.Append("disabled"))
}

// Filter returns a reference to field filter of google_logging_metric.
func (lm loggingMetricAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("filter"))
}

// Id returns a reference to field id of google_logging_metric.
func (lm loggingMetricAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("id"))
}

// LabelExtractors returns a reference to field label_extractors of google_logging_metric.
func (lm loggingMetricAttributes) LabelExtractors() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lm.ref.Append("label_extractors"))
}

// Name returns a reference to field name of google_logging_metric.
func (lm loggingMetricAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("name"))
}

// Project returns a reference to field project of google_logging_metric.
func (lm loggingMetricAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("project"))
}

// ValueExtractor returns a reference to field value_extractor of google_logging_metric.
func (lm loggingMetricAttributes) ValueExtractor() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("value_extractor"))
}

func (lm loggingMetricAttributes) BucketOptions() terra.ListValue[loggingmetric.BucketOptionsAttributes] {
	return terra.ReferenceAsList[loggingmetric.BucketOptionsAttributes](lm.ref.Append("bucket_options"))
}

func (lm loggingMetricAttributes) MetricDescriptor() terra.ListValue[loggingmetric.MetricDescriptorAttributes] {
	return terra.ReferenceAsList[loggingmetric.MetricDescriptorAttributes](lm.ref.Append("metric_descriptor"))
}

func (lm loggingMetricAttributes) Timeouts() loggingmetric.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loggingmetric.TimeoutsAttributes](lm.ref.Append("timeouts"))
}

type loggingMetricState struct {
	BucketName       string                                `json:"bucket_name"`
	Description      string                                `json:"description"`
	Disabled         bool                                  `json:"disabled"`
	Filter           string                                `json:"filter"`
	Id               string                                `json:"id"`
	LabelExtractors  map[string]string                     `json:"label_extractors"`
	Name             string                                `json:"name"`
	Project          string                                `json:"project"`
	ValueExtractor   string                                `json:"value_extractor"`
	BucketOptions    []loggingmetric.BucketOptionsState    `json:"bucket_options"`
	MetricDescriptor []loggingmetric.MetricDescriptorState `json:"metric_descriptor"`
	Timeouts         *loggingmetric.TimeoutsState          `json:"timeouts"`
}

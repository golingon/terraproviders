// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudwatchlogmetricfilter "github.com/golingon/terraproviders/aws/5.12.0/cloudwatchlogmetricfilter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudwatchLogMetricFilter creates a new instance of [CloudwatchLogMetricFilter].
func NewCloudwatchLogMetricFilter(name string, args CloudwatchLogMetricFilterArgs) *CloudwatchLogMetricFilter {
	return &CloudwatchLogMetricFilter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudwatchLogMetricFilter)(nil)

// CloudwatchLogMetricFilter represents the Terraform resource aws_cloudwatch_log_metric_filter.
type CloudwatchLogMetricFilter struct {
	Name      string
	Args      CloudwatchLogMetricFilterArgs
	state     *cloudwatchLogMetricFilterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudwatchLogMetricFilter].
func (clmf *CloudwatchLogMetricFilter) Type() string {
	return "aws_cloudwatch_log_metric_filter"
}

// LocalName returns the local name for [CloudwatchLogMetricFilter].
func (clmf *CloudwatchLogMetricFilter) LocalName() string {
	return clmf.Name
}

// Configuration returns the configuration (args) for [CloudwatchLogMetricFilter].
func (clmf *CloudwatchLogMetricFilter) Configuration() interface{} {
	return clmf.Args
}

// DependOn is used for other resources to depend on [CloudwatchLogMetricFilter].
func (clmf *CloudwatchLogMetricFilter) DependOn() terra.Reference {
	return terra.ReferenceResource(clmf)
}

// Dependencies returns the list of resources [CloudwatchLogMetricFilter] depends_on.
func (clmf *CloudwatchLogMetricFilter) Dependencies() terra.Dependencies {
	return clmf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudwatchLogMetricFilter].
func (clmf *CloudwatchLogMetricFilter) LifecycleManagement() *terra.Lifecycle {
	return clmf.Lifecycle
}

// Attributes returns the attributes for [CloudwatchLogMetricFilter].
func (clmf *CloudwatchLogMetricFilter) Attributes() cloudwatchLogMetricFilterAttributes {
	return cloudwatchLogMetricFilterAttributes{ref: terra.ReferenceResource(clmf)}
}

// ImportState imports the given attribute values into [CloudwatchLogMetricFilter]'s state.
func (clmf *CloudwatchLogMetricFilter) ImportState(av io.Reader) error {
	clmf.state = &cloudwatchLogMetricFilterState{}
	if err := json.NewDecoder(av).Decode(clmf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", clmf.Type(), clmf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudwatchLogMetricFilter] has state.
func (clmf *CloudwatchLogMetricFilter) State() (*cloudwatchLogMetricFilterState, bool) {
	return clmf.state, clmf.state != nil
}

// StateMust returns the state for [CloudwatchLogMetricFilter]. Panics if the state is nil.
func (clmf *CloudwatchLogMetricFilter) StateMust() *cloudwatchLogMetricFilterState {
	if clmf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", clmf.Type(), clmf.LocalName()))
	}
	return clmf.state
}

// CloudwatchLogMetricFilterArgs contains the configurations for aws_cloudwatch_log_metric_filter.
type CloudwatchLogMetricFilterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogGroupName: string, required
	LogGroupName terra.StringValue `hcl:"log_group_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Pattern: string, required
	Pattern terra.StringValue `hcl:"pattern,attr" validate:"required"`
	// MetricTransformation: required
	MetricTransformation *cloudwatchlogmetricfilter.MetricTransformation `hcl:"metric_transformation,block" validate:"required"`
}
type cloudwatchLogMetricFilterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_cloudwatch_log_metric_filter.
func (clmf cloudwatchLogMetricFilterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(clmf.ref.Append("id"))
}

// LogGroupName returns a reference to field log_group_name of aws_cloudwatch_log_metric_filter.
func (clmf cloudwatchLogMetricFilterAttributes) LogGroupName() terra.StringValue {
	return terra.ReferenceAsString(clmf.ref.Append("log_group_name"))
}

// Name returns a reference to field name of aws_cloudwatch_log_metric_filter.
func (clmf cloudwatchLogMetricFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(clmf.ref.Append("name"))
}

// Pattern returns a reference to field pattern of aws_cloudwatch_log_metric_filter.
func (clmf cloudwatchLogMetricFilterAttributes) Pattern() terra.StringValue {
	return terra.ReferenceAsString(clmf.ref.Append("pattern"))
}

func (clmf cloudwatchLogMetricFilterAttributes) MetricTransformation() terra.ListValue[cloudwatchlogmetricfilter.MetricTransformationAttributes] {
	return terra.ReferenceAsList[cloudwatchlogmetricfilter.MetricTransformationAttributes](clmf.ref.Append("metric_transformation"))
}

type cloudwatchLogMetricFilterState struct {
	Id                   string                                                `json:"id"`
	LogGroupName         string                                                `json:"log_group_name"`
	Name                 string                                                `json:"name"`
	Pattern              string                                                `json:"pattern"`
	MetricTransformation []cloudwatchlogmetricfilter.MetricTransformationState `json:"metric_transformation"`
}

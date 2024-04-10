// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	cloudwatchmetricstream "github.com/golingon/terraproviders/aws/5.44.0/cloudwatchmetricstream"
	"io"
)

// NewCloudwatchMetricStream creates a new instance of [CloudwatchMetricStream].
func NewCloudwatchMetricStream(name string, args CloudwatchMetricStreamArgs) *CloudwatchMetricStream {
	return &CloudwatchMetricStream{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudwatchMetricStream)(nil)

// CloudwatchMetricStream represents the Terraform resource aws_cloudwatch_metric_stream.
type CloudwatchMetricStream struct {
	Name      string
	Args      CloudwatchMetricStreamArgs
	state     *cloudwatchMetricStreamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudwatchMetricStream].
func (cms *CloudwatchMetricStream) Type() string {
	return "aws_cloudwatch_metric_stream"
}

// LocalName returns the local name for [CloudwatchMetricStream].
func (cms *CloudwatchMetricStream) LocalName() string {
	return cms.Name
}

// Configuration returns the configuration (args) for [CloudwatchMetricStream].
func (cms *CloudwatchMetricStream) Configuration() interface{} {
	return cms.Args
}

// DependOn is used for other resources to depend on [CloudwatchMetricStream].
func (cms *CloudwatchMetricStream) DependOn() terra.Reference {
	return terra.ReferenceResource(cms)
}

// Dependencies returns the list of resources [CloudwatchMetricStream] depends_on.
func (cms *CloudwatchMetricStream) Dependencies() terra.Dependencies {
	return cms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudwatchMetricStream].
func (cms *CloudwatchMetricStream) LifecycleManagement() *terra.Lifecycle {
	return cms.Lifecycle
}

// Attributes returns the attributes for [CloudwatchMetricStream].
func (cms *CloudwatchMetricStream) Attributes() cloudwatchMetricStreamAttributes {
	return cloudwatchMetricStreamAttributes{ref: terra.ReferenceResource(cms)}
}

// ImportState imports the given attribute values into [CloudwatchMetricStream]'s state.
func (cms *CloudwatchMetricStream) ImportState(av io.Reader) error {
	cms.state = &cloudwatchMetricStreamState{}
	if err := json.NewDecoder(av).Decode(cms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cms.Type(), cms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudwatchMetricStream] has state.
func (cms *CloudwatchMetricStream) State() (*cloudwatchMetricStreamState, bool) {
	return cms.state, cms.state != nil
}

// StateMust returns the state for [CloudwatchMetricStream]. Panics if the state is nil.
func (cms *CloudwatchMetricStream) StateMust() *cloudwatchMetricStreamState {
	if cms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cms.Type(), cms.LocalName()))
	}
	return cms.state
}

// CloudwatchMetricStreamArgs contains the configurations for aws_cloudwatch_metric_stream.
type CloudwatchMetricStreamArgs struct {
	// FirehoseArn: string, required
	FirehoseArn terra.StringValue `hcl:"firehose_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludeLinkedAccountsMetrics: bool, optional
	IncludeLinkedAccountsMetrics terra.BoolValue `hcl:"include_linked_accounts_metrics,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// OutputFormat: string, required
	OutputFormat terra.StringValue `hcl:"output_format,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ExcludeFilter: min=0
	ExcludeFilter []cloudwatchmetricstream.ExcludeFilter `hcl:"exclude_filter,block" validate:"min=0"`
	// IncludeFilter: min=0
	IncludeFilter []cloudwatchmetricstream.IncludeFilter `hcl:"include_filter,block" validate:"min=0"`
	// StatisticsConfiguration: min=0
	StatisticsConfiguration []cloudwatchmetricstream.StatisticsConfiguration `hcl:"statistics_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *cloudwatchmetricstream.Timeouts `hcl:"timeouts,block"`
}
type cloudwatchMetricStreamAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("arn"))
}

// CreationDate returns a reference to field creation_date of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("creation_date"))
}

// FirehoseArn returns a reference to field firehose_arn of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) FirehoseArn() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("firehose_arn"))
}

// Id returns a reference to field id of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("id"))
}

// IncludeLinkedAccountsMetrics returns a reference to field include_linked_accounts_metrics of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) IncludeLinkedAccountsMetrics() terra.BoolValue {
	return terra.ReferenceAsBool(cms.ref.Append("include_linked_accounts_metrics"))
}

// LastUpdateDate returns a reference to field last_update_date of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) LastUpdateDate() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("last_update_date"))
}

// Name returns a reference to field name of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("name_prefix"))
}

// OutputFormat returns a reference to field output_format of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) OutputFormat() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("output_format"))
}

// RoleArn returns a reference to field role_arn of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("role_arn"))
}

// State returns a reference to field state of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cms.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_cloudwatch_metric_stream.
func (cms cloudwatchMetricStreamAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cms.ref.Append("tags_all"))
}

func (cms cloudwatchMetricStreamAttributes) ExcludeFilter() terra.SetValue[cloudwatchmetricstream.ExcludeFilterAttributes] {
	return terra.ReferenceAsSet[cloudwatchmetricstream.ExcludeFilterAttributes](cms.ref.Append("exclude_filter"))
}

func (cms cloudwatchMetricStreamAttributes) IncludeFilter() terra.SetValue[cloudwatchmetricstream.IncludeFilterAttributes] {
	return terra.ReferenceAsSet[cloudwatchmetricstream.IncludeFilterAttributes](cms.ref.Append("include_filter"))
}

func (cms cloudwatchMetricStreamAttributes) StatisticsConfiguration() terra.SetValue[cloudwatchmetricstream.StatisticsConfigurationAttributes] {
	return terra.ReferenceAsSet[cloudwatchmetricstream.StatisticsConfigurationAttributes](cms.ref.Append("statistics_configuration"))
}

func (cms cloudwatchMetricStreamAttributes) Timeouts() cloudwatchmetricstream.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudwatchmetricstream.TimeoutsAttributes](cms.ref.Append("timeouts"))
}

type cloudwatchMetricStreamState struct {
	Arn                          string                                                `json:"arn"`
	CreationDate                 string                                                `json:"creation_date"`
	FirehoseArn                  string                                                `json:"firehose_arn"`
	Id                           string                                                `json:"id"`
	IncludeLinkedAccountsMetrics bool                                                  `json:"include_linked_accounts_metrics"`
	LastUpdateDate               string                                                `json:"last_update_date"`
	Name                         string                                                `json:"name"`
	NamePrefix                   string                                                `json:"name_prefix"`
	OutputFormat                 string                                                `json:"output_format"`
	RoleArn                      string                                                `json:"role_arn"`
	State                        string                                                `json:"state"`
	Tags                         map[string]string                                     `json:"tags"`
	TagsAll                      map[string]string                                     `json:"tags_all"`
	ExcludeFilter                []cloudwatchmetricstream.ExcludeFilterState           `json:"exclude_filter"`
	IncludeFilter                []cloudwatchmetricstream.IncludeFilterState           `json:"include_filter"`
	StatisticsConfiguration      []cloudwatchmetricstream.StatisticsConfigurationState `json:"statistics_configuration"`
	Timeouts                     *cloudwatchmetricstream.TimeoutsState                 `json:"timeouts"`
}

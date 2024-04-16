// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudwatch_metric_stream

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

// Resource represents the Terraform resource aws_cloudwatch_metric_stream.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCloudwatchMetricStreamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acms *Resource) Type() string {
	return "aws_cloudwatch_metric_stream"
}

// LocalName returns the local name for [Resource].
func (acms *Resource) LocalName() string {
	return acms.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acms *Resource) Configuration() interface{} {
	return acms.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acms *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acms)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acms *Resource) Dependencies() terra.Dependencies {
	return acms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acms *Resource) LifecycleManagement() *terra.Lifecycle {
	return acms.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acms *Resource) Attributes() awsCloudwatchMetricStreamAttributes {
	return awsCloudwatchMetricStreamAttributes{ref: terra.ReferenceResource(acms)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acms *Resource) ImportState(state io.Reader) error {
	acms.state = &awsCloudwatchMetricStreamState{}
	if err := json.NewDecoder(state).Decode(acms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acms.Type(), acms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acms *Resource) State() (*awsCloudwatchMetricStreamState, bool) {
	return acms.state, acms.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acms *Resource) StateMust() *awsCloudwatchMetricStreamState {
	if acms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acms.Type(), acms.LocalName()))
	}
	return acms.state
}

// Args contains the configurations for aws_cloudwatch_metric_stream.
type Args struct {
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
	ExcludeFilter []ExcludeFilter `hcl:"exclude_filter,block" validate:"min=0"`
	// IncludeFilter: min=0
	IncludeFilter []IncludeFilter `hcl:"include_filter,block" validate:"min=0"`
	// StatisticsConfiguration: min=0
	StatisticsConfiguration []StatisticsConfiguration `hcl:"statistics_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsCloudwatchMetricStreamAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acms.ref.Append("arn"))
}

// CreationDate returns a reference to field creation_date of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(acms.ref.Append("creation_date"))
}

// FirehoseArn returns a reference to field firehose_arn of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) FirehoseArn() terra.StringValue {
	return terra.ReferenceAsString(acms.ref.Append("firehose_arn"))
}

// Id returns a reference to field id of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acms.ref.Append("id"))
}

// IncludeLinkedAccountsMetrics returns a reference to field include_linked_accounts_metrics of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) IncludeLinkedAccountsMetrics() terra.BoolValue {
	return terra.ReferenceAsBool(acms.ref.Append("include_linked_accounts_metrics"))
}

// LastUpdateDate returns a reference to field last_update_date of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) LastUpdateDate() terra.StringValue {
	return terra.ReferenceAsString(acms.ref.Append("last_update_date"))
}

// Name returns a reference to field name of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acms.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(acms.ref.Append("name_prefix"))
}

// OutputFormat returns a reference to field output_format of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) OutputFormat() terra.StringValue {
	return terra.ReferenceAsString(acms.ref.Append("output_format"))
}

// RoleArn returns a reference to field role_arn of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(acms.ref.Append("role_arn"))
}

// State returns a reference to field state of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(acms.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acms.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_cloudwatch_metric_stream.
func (acms awsCloudwatchMetricStreamAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acms.ref.Append("tags_all"))
}

func (acms awsCloudwatchMetricStreamAttributes) ExcludeFilter() terra.SetValue[ExcludeFilterAttributes] {
	return terra.ReferenceAsSet[ExcludeFilterAttributes](acms.ref.Append("exclude_filter"))
}

func (acms awsCloudwatchMetricStreamAttributes) IncludeFilter() terra.SetValue[IncludeFilterAttributes] {
	return terra.ReferenceAsSet[IncludeFilterAttributes](acms.ref.Append("include_filter"))
}

func (acms awsCloudwatchMetricStreamAttributes) StatisticsConfiguration() terra.SetValue[StatisticsConfigurationAttributes] {
	return terra.ReferenceAsSet[StatisticsConfigurationAttributes](acms.ref.Append("statistics_configuration"))
}

func (acms awsCloudwatchMetricStreamAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acms.ref.Append("timeouts"))
}

type awsCloudwatchMetricStreamState struct {
	Arn                          string                         `json:"arn"`
	CreationDate                 string                         `json:"creation_date"`
	FirehoseArn                  string                         `json:"firehose_arn"`
	Id                           string                         `json:"id"`
	IncludeLinkedAccountsMetrics bool                           `json:"include_linked_accounts_metrics"`
	LastUpdateDate               string                         `json:"last_update_date"`
	Name                         string                         `json:"name"`
	NamePrefix                   string                         `json:"name_prefix"`
	OutputFormat                 string                         `json:"output_format"`
	RoleArn                      string                         `json:"role_arn"`
	State                        string                         `json:"state"`
	Tags                         map[string]string              `json:"tags"`
	TagsAll                      map[string]string              `json:"tags_all"`
	ExcludeFilter                []ExcludeFilterState           `json:"exclude_filter"`
	IncludeFilter                []IncludeFilterState           `json:"include_filter"`
	StatisticsConfiguration      []StatisticsConfigurationState `json:"statistics_configuration"`
	Timeouts                     *TimeoutsState                 `json:"timeouts"`
}

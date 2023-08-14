// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigqueryanalyticshubdataexchangeiambinding "github.com/golingon/terraproviders/google/4.77.0/bigqueryanalyticshubdataexchangeiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryAnalyticsHubDataExchangeIamBinding creates a new instance of [BigqueryAnalyticsHubDataExchangeIamBinding].
func NewBigqueryAnalyticsHubDataExchangeIamBinding(name string, args BigqueryAnalyticsHubDataExchangeIamBindingArgs) *BigqueryAnalyticsHubDataExchangeIamBinding {
	return &BigqueryAnalyticsHubDataExchangeIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryAnalyticsHubDataExchangeIamBinding)(nil)

// BigqueryAnalyticsHubDataExchangeIamBinding represents the Terraform resource google_bigquery_analytics_hub_data_exchange_iam_binding.
type BigqueryAnalyticsHubDataExchangeIamBinding struct {
	Name      string
	Args      BigqueryAnalyticsHubDataExchangeIamBindingArgs
	state     *bigqueryAnalyticsHubDataExchangeIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryAnalyticsHubDataExchangeIamBinding].
func (bahdeib *BigqueryAnalyticsHubDataExchangeIamBinding) Type() string {
	return "google_bigquery_analytics_hub_data_exchange_iam_binding"
}

// LocalName returns the local name for [BigqueryAnalyticsHubDataExchangeIamBinding].
func (bahdeib *BigqueryAnalyticsHubDataExchangeIamBinding) LocalName() string {
	return bahdeib.Name
}

// Configuration returns the configuration (args) for [BigqueryAnalyticsHubDataExchangeIamBinding].
func (bahdeib *BigqueryAnalyticsHubDataExchangeIamBinding) Configuration() interface{} {
	return bahdeib.Args
}

// DependOn is used for other resources to depend on [BigqueryAnalyticsHubDataExchangeIamBinding].
func (bahdeib *BigqueryAnalyticsHubDataExchangeIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(bahdeib)
}

// Dependencies returns the list of resources [BigqueryAnalyticsHubDataExchangeIamBinding] depends_on.
func (bahdeib *BigqueryAnalyticsHubDataExchangeIamBinding) Dependencies() terra.Dependencies {
	return bahdeib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryAnalyticsHubDataExchangeIamBinding].
func (bahdeib *BigqueryAnalyticsHubDataExchangeIamBinding) LifecycleManagement() *terra.Lifecycle {
	return bahdeib.Lifecycle
}

// Attributes returns the attributes for [BigqueryAnalyticsHubDataExchangeIamBinding].
func (bahdeib *BigqueryAnalyticsHubDataExchangeIamBinding) Attributes() bigqueryAnalyticsHubDataExchangeIamBindingAttributes {
	return bigqueryAnalyticsHubDataExchangeIamBindingAttributes{ref: terra.ReferenceResource(bahdeib)}
}

// ImportState imports the given attribute values into [BigqueryAnalyticsHubDataExchangeIamBinding]'s state.
func (bahdeib *BigqueryAnalyticsHubDataExchangeIamBinding) ImportState(av io.Reader) error {
	bahdeib.state = &bigqueryAnalyticsHubDataExchangeIamBindingState{}
	if err := json.NewDecoder(av).Decode(bahdeib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bahdeib.Type(), bahdeib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryAnalyticsHubDataExchangeIamBinding] has state.
func (bahdeib *BigqueryAnalyticsHubDataExchangeIamBinding) State() (*bigqueryAnalyticsHubDataExchangeIamBindingState, bool) {
	return bahdeib.state, bahdeib.state != nil
}

// StateMust returns the state for [BigqueryAnalyticsHubDataExchangeIamBinding]. Panics if the state is nil.
func (bahdeib *BigqueryAnalyticsHubDataExchangeIamBinding) StateMust() *bigqueryAnalyticsHubDataExchangeIamBindingState {
	if bahdeib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bahdeib.Type(), bahdeib.LocalName()))
	}
	return bahdeib.state
}

// BigqueryAnalyticsHubDataExchangeIamBindingArgs contains the configurations for google_bigquery_analytics_hub_data_exchange_iam_binding.
type BigqueryAnalyticsHubDataExchangeIamBindingArgs struct {
	// DataExchangeId: string, required
	DataExchangeId terra.StringValue `hcl:"data_exchange_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *bigqueryanalyticshubdataexchangeiambinding.Condition `hcl:"condition,block"`
}
type bigqueryAnalyticsHubDataExchangeIamBindingAttributes struct {
	ref terra.Reference
}

// DataExchangeId returns a reference to field data_exchange_id of google_bigquery_analytics_hub_data_exchange_iam_binding.
func (bahdeib bigqueryAnalyticsHubDataExchangeIamBindingAttributes) DataExchangeId() terra.StringValue {
	return terra.ReferenceAsString(bahdeib.ref.Append("data_exchange_id"))
}

// Etag returns a reference to field etag of google_bigquery_analytics_hub_data_exchange_iam_binding.
func (bahdeib bigqueryAnalyticsHubDataExchangeIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bahdeib.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_analytics_hub_data_exchange_iam_binding.
func (bahdeib bigqueryAnalyticsHubDataExchangeIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bahdeib.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_analytics_hub_data_exchange_iam_binding.
func (bahdeib bigqueryAnalyticsHubDataExchangeIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bahdeib.ref.Append("location"))
}

// Members returns a reference to field members of google_bigquery_analytics_hub_data_exchange_iam_binding.
func (bahdeib bigqueryAnalyticsHubDataExchangeIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bahdeib.ref.Append("members"))
}

// Project returns a reference to field project of google_bigquery_analytics_hub_data_exchange_iam_binding.
func (bahdeib bigqueryAnalyticsHubDataExchangeIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bahdeib.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_analytics_hub_data_exchange_iam_binding.
func (bahdeib bigqueryAnalyticsHubDataExchangeIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bahdeib.ref.Append("role"))
}

func (bahdeib bigqueryAnalyticsHubDataExchangeIamBindingAttributes) Condition() terra.ListValue[bigqueryanalyticshubdataexchangeiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[bigqueryanalyticshubdataexchangeiambinding.ConditionAttributes](bahdeib.ref.Append("condition"))
}

type bigqueryAnalyticsHubDataExchangeIamBindingState struct {
	DataExchangeId string                                                      `json:"data_exchange_id"`
	Etag           string                                                      `json:"etag"`
	Id             string                                                      `json:"id"`
	Location       string                                                      `json:"location"`
	Members        []string                                                    `json:"members"`
	Project        string                                                      `json:"project"`
	Role           string                                                      `json:"role"`
	Condition      []bigqueryanalyticshubdataexchangeiambinding.ConditionState `json:"condition"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigqueryanalyticshubdataexchangeiammember "github.com/golingon/terraproviders/googlebeta/4.62.0/bigqueryanalyticshubdataexchangeiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryAnalyticsHubDataExchangeIamMember creates a new instance of [BigqueryAnalyticsHubDataExchangeIamMember].
func NewBigqueryAnalyticsHubDataExchangeIamMember(name string, args BigqueryAnalyticsHubDataExchangeIamMemberArgs) *BigqueryAnalyticsHubDataExchangeIamMember {
	return &BigqueryAnalyticsHubDataExchangeIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryAnalyticsHubDataExchangeIamMember)(nil)

// BigqueryAnalyticsHubDataExchangeIamMember represents the Terraform resource google_bigquery_analytics_hub_data_exchange_iam_member.
type BigqueryAnalyticsHubDataExchangeIamMember struct {
	Name      string
	Args      BigqueryAnalyticsHubDataExchangeIamMemberArgs
	state     *bigqueryAnalyticsHubDataExchangeIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryAnalyticsHubDataExchangeIamMember].
func (bahdeim *BigqueryAnalyticsHubDataExchangeIamMember) Type() string {
	return "google_bigquery_analytics_hub_data_exchange_iam_member"
}

// LocalName returns the local name for [BigqueryAnalyticsHubDataExchangeIamMember].
func (bahdeim *BigqueryAnalyticsHubDataExchangeIamMember) LocalName() string {
	return bahdeim.Name
}

// Configuration returns the configuration (args) for [BigqueryAnalyticsHubDataExchangeIamMember].
func (bahdeim *BigqueryAnalyticsHubDataExchangeIamMember) Configuration() interface{} {
	return bahdeim.Args
}

// DependOn is used for other resources to depend on [BigqueryAnalyticsHubDataExchangeIamMember].
func (bahdeim *BigqueryAnalyticsHubDataExchangeIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(bahdeim)
}

// Dependencies returns the list of resources [BigqueryAnalyticsHubDataExchangeIamMember] depends_on.
func (bahdeim *BigqueryAnalyticsHubDataExchangeIamMember) Dependencies() terra.Dependencies {
	return bahdeim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryAnalyticsHubDataExchangeIamMember].
func (bahdeim *BigqueryAnalyticsHubDataExchangeIamMember) LifecycleManagement() *terra.Lifecycle {
	return bahdeim.Lifecycle
}

// Attributes returns the attributes for [BigqueryAnalyticsHubDataExchangeIamMember].
func (bahdeim *BigqueryAnalyticsHubDataExchangeIamMember) Attributes() bigqueryAnalyticsHubDataExchangeIamMemberAttributes {
	return bigqueryAnalyticsHubDataExchangeIamMemberAttributes{ref: terra.ReferenceResource(bahdeim)}
}

// ImportState imports the given attribute values into [BigqueryAnalyticsHubDataExchangeIamMember]'s state.
func (bahdeim *BigqueryAnalyticsHubDataExchangeIamMember) ImportState(av io.Reader) error {
	bahdeim.state = &bigqueryAnalyticsHubDataExchangeIamMemberState{}
	if err := json.NewDecoder(av).Decode(bahdeim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bahdeim.Type(), bahdeim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryAnalyticsHubDataExchangeIamMember] has state.
func (bahdeim *BigqueryAnalyticsHubDataExchangeIamMember) State() (*bigqueryAnalyticsHubDataExchangeIamMemberState, bool) {
	return bahdeim.state, bahdeim.state != nil
}

// StateMust returns the state for [BigqueryAnalyticsHubDataExchangeIamMember]. Panics if the state is nil.
func (bahdeim *BigqueryAnalyticsHubDataExchangeIamMember) StateMust() *bigqueryAnalyticsHubDataExchangeIamMemberState {
	if bahdeim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bahdeim.Type(), bahdeim.LocalName()))
	}
	return bahdeim.state
}

// BigqueryAnalyticsHubDataExchangeIamMemberArgs contains the configurations for google_bigquery_analytics_hub_data_exchange_iam_member.
type BigqueryAnalyticsHubDataExchangeIamMemberArgs struct {
	// DataExchangeId: string, required
	DataExchangeId terra.StringValue `hcl:"data_exchange_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *bigqueryanalyticshubdataexchangeiammember.Condition `hcl:"condition,block"`
}
type bigqueryAnalyticsHubDataExchangeIamMemberAttributes struct {
	ref terra.Reference
}

// DataExchangeId returns a reference to field data_exchange_id of google_bigquery_analytics_hub_data_exchange_iam_member.
func (bahdeim bigqueryAnalyticsHubDataExchangeIamMemberAttributes) DataExchangeId() terra.StringValue {
	return terra.ReferenceAsString(bahdeim.ref.Append("data_exchange_id"))
}

// Etag returns a reference to field etag of google_bigquery_analytics_hub_data_exchange_iam_member.
func (bahdeim bigqueryAnalyticsHubDataExchangeIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bahdeim.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_analytics_hub_data_exchange_iam_member.
func (bahdeim bigqueryAnalyticsHubDataExchangeIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bahdeim.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_analytics_hub_data_exchange_iam_member.
func (bahdeim bigqueryAnalyticsHubDataExchangeIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bahdeim.ref.Append("location"))
}

// Member returns a reference to field member of google_bigquery_analytics_hub_data_exchange_iam_member.
func (bahdeim bigqueryAnalyticsHubDataExchangeIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(bahdeim.ref.Append("member"))
}

// Project returns a reference to field project of google_bigquery_analytics_hub_data_exchange_iam_member.
func (bahdeim bigqueryAnalyticsHubDataExchangeIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bahdeim.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_analytics_hub_data_exchange_iam_member.
func (bahdeim bigqueryAnalyticsHubDataExchangeIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bahdeim.ref.Append("role"))
}

func (bahdeim bigqueryAnalyticsHubDataExchangeIamMemberAttributes) Condition() terra.ListValue[bigqueryanalyticshubdataexchangeiammember.ConditionAttributes] {
	return terra.ReferenceAsList[bigqueryanalyticshubdataexchangeiammember.ConditionAttributes](bahdeim.ref.Append("condition"))
}

type bigqueryAnalyticsHubDataExchangeIamMemberState struct {
	DataExchangeId string                                                     `json:"data_exchange_id"`
	Etag           string                                                     `json:"etag"`
	Id             string                                                     `json:"id"`
	Location       string                                                     `json:"location"`
	Member         string                                                     `json:"member"`
	Project        string                                                     `json:"project"`
	Role           string                                                     `json:"role"`
	Condition      []bigqueryanalyticshubdataexchangeiammember.ConditionState `json:"condition"`
}

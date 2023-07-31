// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigqueryanalyticshublistingiammember "github.com/golingon/terraproviders/google/4.75.1/bigqueryanalyticshublistingiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryAnalyticsHubListingIamMember creates a new instance of [BigqueryAnalyticsHubListingIamMember].
func NewBigqueryAnalyticsHubListingIamMember(name string, args BigqueryAnalyticsHubListingIamMemberArgs) *BigqueryAnalyticsHubListingIamMember {
	return &BigqueryAnalyticsHubListingIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryAnalyticsHubListingIamMember)(nil)

// BigqueryAnalyticsHubListingIamMember represents the Terraform resource google_bigquery_analytics_hub_listing_iam_member.
type BigqueryAnalyticsHubListingIamMember struct {
	Name      string
	Args      BigqueryAnalyticsHubListingIamMemberArgs
	state     *bigqueryAnalyticsHubListingIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryAnalyticsHubListingIamMember].
func (bahlim *BigqueryAnalyticsHubListingIamMember) Type() string {
	return "google_bigquery_analytics_hub_listing_iam_member"
}

// LocalName returns the local name for [BigqueryAnalyticsHubListingIamMember].
func (bahlim *BigqueryAnalyticsHubListingIamMember) LocalName() string {
	return bahlim.Name
}

// Configuration returns the configuration (args) for [BigqueryAnalyticsHubListingIamMember].
func (bahlim *BigqueryAnalyticsHubListingIamMember) Configuration() interface{} {
	return bahlim.Args
}

// DependOn is used for other resources to depend on [BigqueryAnalyticsHubListingIamMember].
func (bahlim *BigqueryAnalyticsHubListingIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(bahlim)
}

// Dependencies returns the list of resources [BigqueryAnalyticsHubListingIamMember] depends_on.
func (bahlim *BigqueryAnalyticsHubListingIamMember) Dependencies() terra.Dependencies {
	return bahlim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryAnalyticsHubListingIamMember].
func (bahlim *BigqueryAnalyticsHubListingIamMember) LifecycleManagement() *terra.Lifecycle {
	return bahlim.Lifecycle
}

// Attributes returns the attributes for [BigqueryAnalyticsHubListingIamMember].
func (bahlim *BigqueryAnalyticsHubListingIamMember) Attributes() bigqueryAnalyticsHubListingIamMemberAttributes {
	return bigqueryAnalyticsHubListingIamMemberAttributes{ref: terra.ReferenceResource(bahlim)}
}

// ImportState imports the given attribute values into [BigqueryAnalyticsHubListingIamMember]'s state.
func (bahlim *BigqueryAnalyticsHubListingIamMember) ImportState(av io.Reader) error {
	bahlim.state = &bigqueryAnalyticsHubListingIamMemberState{}
	if err := json.NewDecoder(av).Decode(bahlim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bahlim.Type(), bahlim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryAnalyticsHubListingIamMember] has state.
func (bahlim *BigqueryAnalyticsHubListingIamMember) State() (*bigqueryAnalyticsHubListingIamMemberState, bool) {
	return bahlim.state, bahlim.state != nil
}

// StateMust returns the state for [BigqueryAnalyticsHubListingIamMember]. Panics if the state is nil.
func (bahlim *BigqueryAnalyticsHubListingIamMember) StateMust() *bigqueryAnalyticsHubListingIamMemberState {
	if bahlim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bahlim.Type(), bahlim.LocalName()))
	}
	return bahlim.state
}

// BigqueryAnalyticsHubListingIamMemberArgs contains the configurations for google_bigquery_analytics_hub_listing_iam_member.
type BigqueryAnalyticsHubListingIamMemberArgs struct {
	// DataExchangeId: string, required
	DataExchangeId terra.StringValue `hcl:"data_exchange_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ListingId: string, required
	ListingId terra.StringValue `hcl:"listing_id,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *bigqueryanalyticshublistingiammember.Condition `hcl:"condition,block"`
}
type bigqueryAnalyticsHubListingIamMemberAttributes struct {
	ref terra.Reference
}

// DataExchangeId returns a reference to field data_exchange_id of google_bigquery_analytics_hub_listing_iam_member.
func (bahlim bigqueryAnalyticsHubListingIamMemberAttributes) DataExchangeId() terra.StringValue {
	return terra.ReferenceAsString(bahlim.ref.Append("data_exchange_id"))
}

// Etag returns a reference to field etag of google_bigquery_analytics_hub_listing_iam_member.
func (bahlim bigqueryAnalyticsHubListingIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bahlim.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_analytics_hub_listing_iam_member.
func (bahlim bigqueryAnalyticsHubListingIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bahlim.ref.Append("id"))
}

// ListingId returns a reference to field listing_id of google_bigquery_analytics_hub_listing_iam_member.
func (bahlim bigqueryAnalyticsHubListingIamMemberAttributes) ListingId() terra.StringValue {
	return terra.ReferenceAsString(bahlim.ref.Append("listing_id"))
}

// Location returns a reference to field location of google_bigquery_analytics_hub_listing_iam_member.
func (bahlim bigqueryAnalyticsHubListingIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bahlim.ref.Append("location"))
}

// Member returns a reference to field member of google_bigquery_analytics_hub_listing_iam_member.
func (bahlim bigqueryAnalyticsHubListingIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(bahlim.ref.Append("member"))
}

// Project returns a reference to field project of google_bigquery_analytics_hub_listing_iam_member.
func (bahlim bigqueryAnalyticsHubListingIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bahlim.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_analytics_hub_listing_iam_member.
func (bahlim bigqueryAnalyticsHubListingIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bahlim.ref.Append("role"))
}

func (bahlim bigqueryAnalyticsHubListingIamMemberAttributes) Condition() terra.ListValue[bigqueryanalyticshublistingiammember.ConditionAttributes] {
	return terra.ReferenceAsList[bigqueryanalyticshublistingiammember.ConditionAttributes](bahlim.ref.Append("condition"))
}

type bigqueryAnalyticsHubListingIamMemberState struct {
	DataExchangeId string                                                `json:"data_exchange_id"`
	Etag           string                                                `json:"etag"`
	Id             string                                                `json:"id"`
	ListingId      string                                                `json:"listing_id"`
	Location       string                                                `json:"location"`
	Member         string                                                `json:"member"`
	Project        string                                                `json:"project"`
	Role           string                                                `json:"role"`
	Condition      []bigqueryanalyticshublistingiammember.ConditionState `json:"condition"`
}

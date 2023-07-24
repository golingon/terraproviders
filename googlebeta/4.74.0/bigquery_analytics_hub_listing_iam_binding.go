// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigqueryanalyticshublistingiambinding "github.com/golingon/terraproviders/googlebeta/4.74.0/bigqueryanalyticshublistingiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryAnalyticsHubListingIamBinding creates a new instance of [BigqueryAnalyticsHubListingIamBinding].
func NewBigqueryAnalyticsHubListingIamBinding(name string, args BigqueryAnalyticsHubListingIamBindingArgs) *BigqueryAnalyticsHubListingIamBinding {
	return &BigqueryAnalyticsHubListingIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryAnalyticsHubListingIamBinding)(nil)

// BigqueryAnalyticsHubListingIamBinding represents the Terraform resource google_bigquery_analytics_hub_listing_iam_binding.
type BigqueryAnalyticsHubListingIamBinding struct {
	Name      string
	Args      BigqueryAnalyticsHubListingIamBindingArgs
	state     *bigqueryAnalyticsHubListingIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryAnalyticsHubListingIamBinding].
func (bahlib *BigqueryAnalyticsHubListingIamBinding) Type() string {
	return "google_bigquery_analytics_hub_listing_iam_binding"
}

// LocalName returns the local name for [BigqueryAnalyticsHubListingIamBinding].
func (bahlib *BigqueryAnalyticsHubListingIamBinding) LocalName() string {
	return bahlib.Name
}

// Configuration returns the configuration (args) for [BigqueryAnalyticsHubListingIamBinding].
func (bahlib *BigqueryAnalyticsHubListingIamBinding) Configuration() interface{} {
	return bahlib.Args
}

// DependOn is used for other resources to depend on [BigqueryAnalyticsHubListingIamBinding].
func (bahlib *BigqueryAnalyticsHubListingIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(bahlib)
}

// Dependencies returns the list of resources [BigqueryAnalyticsHubListingIamBinding] depends_on.
func (bahlib *BigqueryAnalyticsHubListingIamBinding) Dependencies() terra.Dependencies {
	return bahlib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryAnalyticsHubListingIamBinding].
func (bahlib *BigqueryAnalyticsHubListingIamBinding) LifecycleManagement() *terra.Lifecycle {
	return bahlib.Lifecycle
}

// Attributes returns the attributes for [BigqueryAnalyticsHubListingIamBinding].
func (bahlib *BigqueryAnalyticsHubListingIamBinding) Attributes() bigqueryAnalyticsHubListingIamBindingAttributes {
	return bigqueryAnalyticsHubListingIamBindingAttributes{ref: terra.ReferenceResource(bahlib)}
}

// ImportState imports the given attribute values into [BigqueryAnalyticsHubListingIamBinding]'s state.
func (bahlib *BigqueryAnalyticsHubListingIamBinding) ImportState(av io.Reader) error {
	bahlib.state = &bigqueryAnalyticsHubListingIamBindingState{}
	if err := json.NewDecoder(av).Decode(bahlib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bahlib.Type(), bahlib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryAnalyticsHubListingIamBinding] has state.
func (bahlib *BigqueryAnalyticsHubListingIamBinding) State() (*bigqueryAnalyticsHubListingIamBindingState, bool) {
	return bahlib.state, bahlib.state != nil
}

// StateMust returns the state for [BigqueryAnalyticsHubListingIamBinding]. Panics if the state is nil.
func (bahlib *BigqueryAnalyticsHubListingIamBinding) StateMust() *bigqueryAnalyticsHubListingIamBindingState {
	if bahlib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bahlib.Type(), bahlib.LocalName()))
	}
	return bahlib.state
}

// BigqueryAnalyticsHubListingIamBindingArgs contains the configurations for google_bigquery_analytics_hub_listing_iam_binding.
type BigqueryAnalyticsHubListingIamBindingArgs struct {
	// DataExchangeId: string, required
	DataExchangeId terra.StringValue `hcl:"data_exchange_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ListingId: string, required
	ListingId terra.StringValue `hcl:"listing_id,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *bigqueryanalyticshublistingiambinding.Condition `hcl:"condition,block"`
}
type bigqueryAnalyticsHubListingIamBindingAttributes struct {
	ref terra.Reference
}

// DataExchangeId returns a reference to field data_exchange_id of google_bigquery_analytics_hub_listing_iam_binding.
func (bahlib bigqueryAnalyticsHubListingIamBindingAttributes) DataExchangeId() terra.StringValue {
	return terra.ReferenceAsString(bahlib.ref.Append("data_exchange_id"))
}

// Etag returns a reference to field etag of google_bigquery_analytics_hub_listing_iam_binding.
func (bahlib bigqueryAnalyticsHubListingIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bahlib.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_analytics_hub_listing_iam_binding.
func (bahlib bigqueryAnalyticsHubListingIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bahlib.ref.Append("id"))
}

// ListingId returns a reference to field listing_id of google_bigquery_analytics_hub_listing_iam_binding.
func (bahlib bigqueryAnalyticsHubListingIamBindingAttributes) ListingId() terra.StringValue {
	return terra.ReferenceAsString(bahlib.ref.Append("listing_id"))
}

// Location returns a reference to field location of google_bigquery_analytics_hub_listing_iam_binding.
func (bahlib bigqueryAnalyticsHubListingIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bahlib.ref.Append("location"))
}

// Members returns a reference to field members of google_bigquery_analytics_hub_listing_iam_binding.
func (bahlib bigqueryAnalyticsHubListingIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bahlib.ref.Append("members"))
}

// Project returns a reference to field project of google_bigquery_analytics_hub_listing_iam_binding.
func (bahlib bigqueryAnalyticsHubListingIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bahlib.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_analytics_hub_listing_iam_binding.
func (bahlib bigqueryAnalyticsHubListingIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bahlib.ref.Append("role"))
}

func (bahlib bigqueryAnalyticsHubListingIamBindingAttributes) Condition() terra.ListValue[bigqueryanalyticshublistingiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[bigqueryanalyticshublistingiambinding.ConditionAttributes](bahlib.ref.Append("condition"))
}

type bigqueryAnalyticsHubListingIamBindingState struct {
	DataExchangeId string                                                 `json:"data_exchange_id"`
	Etag           string                                                 `json:"etag"`
	Id             string                                                 `json:"id"`
	ListingId      string                                                 `json:"listing_id"`
	Location       string                                                 `json:"location"`
	Members        []string                                               `json:"members"`
	Project        string                                                 `json:"project"`
	Role           string                                                 `json:"role"`
	Condition      []bigqueryanalyticshublistingiambinding.ConditionState `json:"condition"`
}

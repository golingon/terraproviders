// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigqueryanalyticshubdataexchange "github.com/golingon/terraproviders/google/4.71.0/bigqueryanalyticshubdataexchange"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryAnalyticsHubDataExchange creates a new instance of [BigqueryAnalyticsHubDataExchange].
func NewBigqueryAnalyticsHubDataExchange(name string, args BigqueryAnalyticsHubDataExchangeArgs) *BigqueryAnalyticsHubDataExchange {
	return &BigqueryAnalyticsHubDataExchange{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryAnalyticsHubDataExchange)(nil)

// BigqueryAnalyticsHubDataExchange represents the Terraform resource google_bigquery_analytics_hub_data_exchange.
type BigqueryAnalyticsHubDataExchange struct {
	Name      string
	Args      BigqueryAnalyticsHubDataExchangeArgs
	state     *bigqueryAnalyticsHubDataExchangeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryAnalyticsHubDataExchange].
func (bahde *BigqueryAnalyticsHubDataExchange) Type() string {
	return "google_bigquery_analytics_hub_data_exchange"
}

// LocalName returns the local name for [BigqueryAnalyticsHubDataExchange].
func (bahde *BigqueryAnalyticsHubDataExchange) LocalName() string {
	return bahde.Name
}

// Configuration returns the configuration (args) for [BigqueryAnalyticsHubDataExchange].
func (bahde *BigqueryAnalyticsHubDataExchange) Configuration() interface{} {
	return bahde.Args
}

// DependOn is used for other resources to depend on [BigqueryAnalyticsHubDataExchange].
func (bahde *BigqueryAnalyticsHubDataExchange) DependOn() terra.Reference {
	return terra.ReferenceResource(bahde)
}

// Dependencies returns the list of resources [BigqueryAnalyticsHubDataExchange] depends_on.
func (bahde *BigqueryAnalyticsHubDataExchange) Dependencies() terra.Dependencies {
	return bahde.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryAnalyticsHubDataExchange].
func (bahde *BigqueryAnalyticsHubDataExchange) LifecycleManagement() *terra.Lifecycle {
	return bahde.Lifecycle
}

// Attributes returns the attributes for [BigqueryAnalyticsHubDataExchange].
func (bahde *BigqueryAnalyticsHubDataExchange) Attributes() bigqueryAnalyticsHubDataExchangeAttributes {
	return bigqueryAnalyticsHubDataExchangeAttributes{ref: terra.ReferenceResource(bahde)}
}

// ImportState imports the given attribute values into [BigqueryAnalyticsHubDataExchange]'s state.
func (bahde *BigqueryAnalyticsHubDataExchange) ImportState(av io.Reader) error {
	bahde.state = &bigqueryAnalyticsHubDataExchangeState{}
	if err := json.NewDecoder(av).Decode(bahde.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bahde.Type(), bahde.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryAnalyticsHubDataExchange] has state.
func (bahde *BigqueryAnalyticsHubDataExchange) State() (*bigqueryAnalyticsHubDataExchangeState, bool) {
	return bahde.state, bahde.state != nil
}

// StateMust returns the state for [BigqueryAnalyticsHubDataExchange]. Panics if the state is nil.
func (bahde *BigqueryAnalyticsHubDataExchange) StateMust() *bigqueryAnalyticsHubDataExchangeState {
	if bahde.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bahde.Type(), bahde.LocalName()))
	}
	return bahde.state
}

// BigqueryAnalyticsHubDataExchangeArgs contains the configurations for google_bigquery_analytics_hub_data_exchange.
type BigqueryAnalyticsHubDataExchangeArgs struct {
	// DataExchangeId: string, required
	DataExchangeId terra.StringValue `hcl:"data_exchange_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Documentation: string, optional
	Documentation terra.StringValue `hcl:"documentation,attr"`
	// Icon: string, optional
	Icon terra.StringValue `hcl:"icon,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// PrimaryContact: string, optional
	PrimaryContact terra.StringValue `hcl:"primary_contact,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *bigqueryanalyticshubdataexchange.Timeouts `hcl:"timeouts,block"`
}
type bigqueryAnalyticsHubDataExchangeAttributes struct {
	ref terra.Reference
}

// DataExchangeId returns a reference to field data_exchange_id of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) DataExchangeId() terra.StringValue {
	return terra.ReferenceAsString(bahde.ref.Append("data_exchange_id"))
}

// Description returns a reference to field description of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bahde.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bahde.ref.Append("display_name"))
}

// Documentation returns a reference to field documentation of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) Documentation() terra.StringValue {
	return terra.ReferenceAsString(bahde.ref.Append("documentation"))
}

// Icon returns a reference to field icon of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) Icon() terra.StringValue {
	return terra.ReferenceAsString(bahde.ref.Append("icon"))
}

// Id returns a reference to field id of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bahde.ref.Append("id"))
}

// ListingCount returns a reference to field listing_count of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) ListingCount() terra.NumberValue {
	return terra.ReferenceAsNumber(bahde.ref.Append("listing_count"))
}

// Location returns a reference to field location of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bahde.ref.Append("location"))
}

// Name returns a reference to field name of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bahde.ref.Append("name"))
}

// PrimaryContact returns a reference to field primary_contact of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) PrimaryContact() terra.StringValue {
	return terra.ReferenceAsString(bahde.ref.Append("primary_contact"))
}

// Project returns a reference to field project of google_bigquery_analytics_hub_data_exchange.
func (bahde bigqueryAnalyticsHubDataExchangeAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bahde.ref.Append("project"))
}

func (bahde bigqueryAnalyticsHubDataExchangeAttributes) Timeouts() bigqueryanalyticshubdataexchange.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigqueryanalyticshubdataexchange.TimeoutsAttributes](bahde.ref.Append("timeouts"))
}

type bigqueryAnalyticsHubDataExchangeState struct {
	DataExchangeId string                                          `json:"data_exchange_id"`
	Description    string                                          `json:"description"`
	DisplayName    string                                          `json:"display_name"`
	Documentation  string                                          `json:"documentation"`
	Icon           string                                          `json:"icon"`
	Id             string                                          `json:"id"`
	ListingCount   float64                                         `json:"listing_count"`
	Location       string                                          `json:"location"`
	Name           string                                          `json:"name"`
	PrimaryContact string                                          `json:"primary_contact"`
	Project        string                                          `json:"project"`
	Timeouts       *bigqueryanalyticshubdataexchange.TimeoutsState `json:"timeouts"`
}

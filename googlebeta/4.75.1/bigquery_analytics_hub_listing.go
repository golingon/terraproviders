// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigqueryanalyticshublisting "github.com/golingon/terraproviders/googlebeta/4.75.1/bigqueryanalyticshublisting"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryAnalyticsHubListing creates a new instance of [BigqueryAnalyticsHubListing].
func NewBigqueryAnalyticsHubListing(name string, args BigqueryAnalyticsHubListingArgs) *BigqueryAnalyticsHubListing {
	return &BigqueryAnalyticsHubListing{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryAnalyticsHubListing)(nil)

// BigqueryAnalyticsHubListing represents the Terraform resource google_bigquery_analytics_hub_listing.
type BigqueryAnalyticsHubListing struct {
	Name      string
	Args      BigqueryAnalyticsHubListingArgs
	state     *bigqueryAnalyticsHubListingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryAnalyticsHubListing].
func (bahl *BigqueryAnalyticsHubListing) Type() string {
	return "google_bigquery_analytics_hub_listing"
}

// LocalName returns the local name for [BigqueryAnalyticsHubListing].
func (bahl *BigqueryAnalyticsHubListing) LocalName() string {
	return bahl.Name
}

// Configuration returns the configuration (args) for [BigqueryAnalyticsHubListing].
func (bahl *BigqueryAnalyticsHubListing) Configuration() interface{} {
	return bahl.Args
}

// DependOn is used for other resources to depend on [BigqueryAnalyticsHubListing].
func (bahl *BigqueryAnalyticsHubListing) DependOn() terra.Reference {
	return terra.ReferenceResource(bahl)
}

// Dependencies returns the list of resources [BigqueryAnalyticsHubListing] depends_on.
func (bahl *BigqueryAnalyticsHubListing) Dependencies() terra.Dependencies {
	return bahl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryAnalyticsHubListing].
func (bahl *BigqueryAnalyticsHubListing) LifecycleManagement() *terra.Lifecycle {
	return bahl.Lifecycle
}

// Attributes returns the attributes for [BigqueryAnalyticsHubListing].
func (bahl *BigqueryAnalyticsHubListing) Attributes() bigqueryAnalyticsHubListingAttributes {
	return bigqueryAnalyticsHubListingAttributes{ref: terra.ReferenceResource(bahl)}
}

// ImportState imports the given attribute values into [BigqueryAnalyticsHubListing]'s state.
func (bahl *BigqueryAnalyticsHubListing) ImportState(av io.Reader) error {
	bahl.state = &bigqueryAnalyticsHubListingState{}
	if err := json.NewDecoder(av).Decode(bahl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bahl.Type(), bahl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryAnalyticsHubListing] has state.
func (bahl *BigqueryAnalyticsHubListing) State() (*bigqueryAnalyticsHubListingState, bool) {
	return bahl.state, bahl.state != nil
}

// StateMust returns the state for [BigqueryAnalyticsHubListing]. Panics if the state is nil.
func (bahl *BigqueryAnalyticsHubListing) StateMust() *bigqueryAnalyticsHubListingState {
	if bahl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bahl.Type(), bahl.LocalName()))
	}
	return bahl.state
}

// BigqueryAnalyticsHubListingArgs contains the configurations for google_bigquery_analytics_hub_listing.
type BigqueryAnalyticsHubListingArgs struct {
	// Categories: list of string, optional
	Categories terra.ListValue[terra.StringValue] `hcl:"categories,attr"`
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
	// ListingId: string, required
	ListingId terra.StringValue `hcl:"listing_id,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// PrimaryContact: string, optional
	PrimaryContact terra.StringValue `hcl:"primary_contact,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RequestAccess: string, optional
	RequestAccess terra.StringValue `hcl:"request_access,attr"`
	// BigqueryDataset: required
	BigqueryDataset *bigqueryanalyticshublisting.BigqueryDataset `hcl:"bigquery_dataset,block" validate:"required"`
	// DataProvider: optional
	DataProvider *bigqueryanalyticshublisting.DataProvider `hcl:"data_provider,block"`
	// Publisher: optional
	Publisher *bigqueryanalyticshublisting.Publisher `hcl:"publisher,block"`
	// Timeouts: optional
	Timeouts *bigqueryanalyticshublisting.Timeouts `hcl:"timeouts,block"`
}
type bigqueryAnalyticsHubListingAttributes struct {
	ref terra.Reference
}

// Categories returns a reference to field categories of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) Categories() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bahl.ref.Append("categories"))
}

// DataExchangeId returns a reference to field data_exchange_id of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) DataExchangeId() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("data_exchange_id"))
}

// Description returns a reference to field description of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("display_name"))
}

// Documentation returns a reference to field documentation of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) Documentation() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("documentation"))
}

// Icon returns a reference to field icon of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) Icon() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("icon"))
}

// Id returns a reference to field id of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("id"))
}

// ListingId returns a reference to field listing_id of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) ListingId() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("listing_id"))
}

// Location returns a reference to field location of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("location"))
}

// Name returns a reference to field name of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("name"))
}

// PrimaryContact returns a reference to field primary_contact of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) PrimaryContact() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("primary_contact"))
}

// Project returns a reference to field project of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("project"))
}

// RequestAccess returns a reference to field request_access of google_bigquery_analytics_hub_listing.
func (bahl bigqueryAnalyticsHubListingAttributes) RequestAccess() terra.StringValue {
	return terra.ReferenceAsString(bahl.ref.Append("request_access"))
}

func (bahl bigqueryAnalyticsHubListingAttributes) BigqueryDataset() terra.ListValue[bigqueryanalyticshublisting.BigqueryDatasetAttributes] {
	return terra.ReferenceAsList[bigqueryanalyticshublisting.BigqueryDatasetAttributes](bahl.ref.Append("bigquery_dataset"))
}

func (bahl bigqueryAnalyticsHubListingAttributes) DataProvider() terra.ListValue[bigqueryanalyticshublisting.DataProviderAttributes] {
	return terra.ReferenceAsList[bigqueryanalyticshublisting.DataProviderAttributes](bahl.ref.Append("data_provider"))
}

func (bahl bigqueryAnalyticsHubListingAttributes) Publisher() terra.ListValue[bigqueryanalyticshublisting.PublisherAttributes] {
	return terra.ReferenceAsList[bigqueryanalyticshublisting.PublisherAttributes](bahl.ref.Append("publisher"))
}

func (bahl bigqueryAnalyticsHubListingAttributes) Timeouts() bigqueryanalyticshublisting.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigqueryanalyticshublisting.TimeoutsAttributes](bahl.ref.Append("timeouts"))
}

type bigqueryAnalyticsHubListingState struct {
	Categories      []string                                           `json:"categories"`
	DataExchangeId  string                                             `json:"data_exchange_id"`
	Description     string                                             `json:"description"`
	DisplayName     string                                             `json:"display_name"`
	Documentation   string                                             `json:"documentation"`
	Icon            string                                             `json:"icon"`
	Id              string                                             `json:"id"`
	ListingId       string                                             `json:"listing_id"`
	Location        string                                             `json:"location"`
	Name            string                                             `json:"name"`
	PrimaryContact  string                                             `json:"primary_contact"`
	Project         string                                             `json:"project"`
	RequestAccess   string                                             `json:"request_access"`
	BigqueryDataset []bigqueryanalyticshublisting.BigqueryDatasetState `json:"bigquery_dataset"`
	DataProvider    []bigqueryanalyticshublisting.DataProviderState    `json:"data_provider"`
	Publisher       []bigqueryanalyticshublisting.PublisherState       `json:"publisher"`
	Timeouts        *bigqueryanalyticshublisting.TimeoutsState         `json:"timeouts"`
}

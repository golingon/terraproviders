// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_bigquery_analytics_hub_listing_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_bigquery_analytics_hub_listing_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gbahlip *DataSource) DataSource() string {
	return "google_bigquery_analytics_hub_listing_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gbahlip *DataSource) LocalName() string {
	return gbahlip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gbahlip *DataSource) Configuration() interface{} {
	return gbahlip.Args
}

// Attributes returns the attributes for [DataSource].
func (gbahlip *DataSource) Attributes() dataGoogleBigqueryAnalyticsHubListingIamPolicyAttributes {
	return dataGoogleBigqueryAnalyticsHubListingIamPolicyAttributes{ref: terra.ReferenceDataSource(gbahlip)}
}

// DataArgs contains the configurations for google_bigquery_analytics_hub_listing_iam_policy.
type DataArgs struct {
	// DataExchangeId: string, required
	DataExchangeId terra.StringValue `hcl:"data_exchange_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ListingId: string, required
	ListingId terra.StringValue `hcl:"listing_id,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type dataGoogleBigqueryAnalyticsHubListingIamPolicyAttributes struct {
	ref terra.Reference
}

// DataExchangeId returns a reference to field data_exchange_id of google_bigquery_analytics_hub_listing_iam_policy.
func (gbahlip dataGoogleBigqueryAnalyticsHubListingIamPolicyAttributes) DataExchangeId() terra.StringValue {
	return terra.ReferenceAsString(gbahlip.ref.Append("data_exchange_id"))
}

// Etag returns a reference to field etag of google_bigquery_analytics_hub_listing_iam_policy.
func (gbahlip dataGoogleBigqueryAnalyticsHubListingIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gbahlip.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_analytics_hub_listing_iam_policy.
func (gbahlip dataGoogleBigqueryAnalyticsHubListingIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbahlip.ref.Append("id"))
}

// ListingId returns a reference to field listing_id of google_bigquery_analytics_hub_listing_iam_policy.
func (gbahlip dataGoogleBigqueryAnalyticsHubListingIamPolicyAttributes) ListingId() terra.StringValue {
	return terra.ReferenceAsString(gbahlip.ref.Append("listing_id"))
}

// Location returns a reference to field location of google_bigquery_analytics_hub_listing_iam_policy.
func (gbahlip dataGoogleBigqueryAnalyticsHubListingIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gbahlip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_bigquery_analytics_hub_listing_iam_policy.
func (gbahlip dataGoogleBigqueryAnalyticsHubListingIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gbahlip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_bigquery_analytics_hub_listing_iam_policy.
func (gbahlip dataGoogleBigqueryAnalyticsHubListingIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbahlip.ref.Append("project"))
}

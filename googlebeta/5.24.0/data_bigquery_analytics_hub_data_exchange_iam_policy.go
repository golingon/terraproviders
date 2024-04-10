// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataBigqueryAnalyticsHubDataExchangeIamPolicy creates a new instance of [DataBigqueryAnalyticsHubDataExchangeIamPolicy].
func NewDataBigqueryAnalyticsHubDataExchangeIamPolicy(name string, args DataBigqueryAnalyticsHubDataExchangeIamPolicyArgs) *DataBigqueryAnalyticsHubDataExchangeIamPolicy {
	return &DataBigqueryAnalyticsHubDataExchangeIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBigqueryAnalyticsHubDataExchangeIamPolicy)(nil)

// DataBigqueryAnalyticsHubDataExchangeIamPolicy represents the Terraform data resource google_bigquery_analytics_hub_data_exchange_iam_policy.
type DataBigqueryAnalyticsHubDataExchangeIamPolicy struct {
	Name string
	Args DataBigqueryAnalyticsHubDataExchangeIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataBigqueryAnalyticsHubDataExchangeIamPolicy].
func (bahdeip *DataBigqueryAnalyticsHubDataExchangeIamPolicy) DataSource() string {
	return "google_bigquery_analytics_hub_data_exchange_iam_policy"
}

// LocalName returns the local name for [DataBigqueryAnalyticsHubDataExchangeIamPolicy].
func (bahdeip *DataBigqueryAnalyticsHubDataExchangeIamPolicy) LocalName() string {
	return bahdeip.Name
}

// Configuration returns the configuration (args) for [DataBigqueryAnalyticsHubDataExchangeIamPolicy].
func (bahdeip *DataBigqueryAnalyticsHubDataExchangeIamPolicy) Configuration() interface{} {
	return bahdeip.Args
}

// Attributes returns the attributes for [DataBigqueryAnalyticsHubDataExchangeIamPolicy].
func (bahdeip *DataBigqueryAnalyticsHubDataExchangeIamPolicy) Attributes() dataBigqueryAnalyticsHubDataExchangeIamPolicyAttributes {
	return dataBigqueryAnalyticsHubDataExchangeIamPolicyAttributes{ref: terra.ReferenceDataResource(bahdeip)}
}

// DataBigqueryAnalyticsHubDataExchangeIamPolicyArgs contains the configurations for google_bigquery_analytics_hub_data_exchange_iam_policy.
type DataBigqueryAnalyticsHubDataExchangeIamPolicyArgs struct {
	// DataExchangeId: string, required
	DataExchangeId terra.StringValue `hcl:"data_exchange_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataBigqueryAnalyticsHubDataExchangeIamPolicyAttributes struct {
	ref terra.Reference
}

// DataExchangeId returns a reference to field data_exchange_id of google_bigquery_analytics_hub_data_exchange_iam_policy.
func (bahdeip dataBigqueryAnalyticsHubDataExchangeIamPolicyAttributes) DataExchangeId() terra.StringValue {
	return terra.ReferenceAsString(bahdeip.ref.Append("data_exchange_id"))
}

// Etag returns a reference to field etag of google_bigquery_analytics_hub_data_exchange_iam_policy.
func (bahdeip dataBigqueryAnalyticsHubDataExchangeIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bahdeip.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_analytics_hub_data_exchange_iam_policy.
func (bahdeip dataBigqueryAnalyticsHubDataExchangeIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bahdeip.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_analytics_hub_data_exchange_iam_policy.
func (bahdeip dataBigqueryAnalyticsHubDataExchangeIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bahdeip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_bigquery_analytics_hub_data_exchange_iam_policy.
func (bahdeip dataBigqueryAnalyticsHubDataExchangeIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(bahdeip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_bigquery_analytics_hub_data_exchange_iam_policy.
func (bahdeip dataBigqueryAnalyticsHubDataExchangeIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bahdeip.ref.Append("project"))
}

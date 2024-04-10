// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataBigqueryDefaultServiceAccount creates a new instance of [DataBigqueryDefaultServiceAccount].
func NewDataBigqueryDefaultServiceAccount(name string, args DataBigqueryDefaultServiceAccountArgs) *DataBigqueryDefaultServiceAccount {
	return &DataBigqueryDefaultServiceAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBigqueryDefaultServiceAccount)(nil)

// DataBigqueryDefaultServiceAccount represents the Terraform data resource google_bigquery_default_service_account.
type DataBigqueryDefaultServiceAccount struct {
	Name string
	Args DataBigqueryDefaultServiceAccountArgs
}

// DataSource returns the Terraform object type for [DataBigqueryDefaultServiceAccount].
func (bdsa *DataBigqueryDefaultServiceAccount) DataSource() string {
	return "google_bigquery_default_service_account"
}

// LocalName returns the local name for [DataBigqueryDefaultServiceAccount].
func (bdsa *DataBigqueryDefaultServiceAccount) LocalName() string {
	return bdsa.Name
}

// Configuration returns the configuration (args) for [DataBigqueryDefaultServiceAccount].
func (bdsa *DataBigqueryDefaultServiceAccount) Configuration() interface{} {
	return bdsa.Args
}

// Attributes returns the attributes for [DataBigqueryDefaultServiceAccount].
func (bdsa *DataBigqueryDefaultServiceAccount) Attributes() dataBigqueryDefaultServiceAccountAttributes {
	return dataBigqueryDefaultServiceAccountAttributes{ref: terra.ReferenceDataResource(bdsa)}
}

// DataBigqueryDefaultServiceAccountArgs contains the configurations for google_bigquery_default_service_account.
type DataBigqueryDefaultServiceAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataBigqueryDefaultServiceAccountAttributes struct {
	ref terra.Reference
}

// Email returns a reference to field email of google_bigquery_default_service_account.
func (bdsa dataBigqueryDefaultServiceAccountAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(bdsa.ref.Append("email"))
}

// Id returns a reference to field id of google_bigquery_default_service_account.
func (bdsa dataBigqueryDefaultServiceAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bdsa.ref.Append("id"))
}

// Member returns a reference to field member of google_bigquery_default_service_account.
func (bdsa dataBigqueryDefaultServiceAccountAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(bdsa.ref.Append("member"))
}

// Project returns a reference to field project of google_bigquery_default_service_account.
func (bdsa dataBigqueryDefaultServiceAccountAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bdsa.ref.Append("project"))
}

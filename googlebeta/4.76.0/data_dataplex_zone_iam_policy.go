// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDataplexZoneIamPolicy creates a new instance of [DataDataplexZoneIamPolicy].
func NewDataDataplexZoneIamPolicy(name string, args DataDataplexZoneIamPolicyArgs) *DataDataplexZoneIamPolicy {
	return &DataDataplexZoneIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataplexZoneIamPolicy)(nil)

// DataDataplexZoneIamPolicy represents the Terraform data resource google_dataplex_zone_iam_policy.
type DataDataplexZoneIamPolicy struct {
	Name string
	Args DataDataplexZoneIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataDataplexZoneIamPolicy].
func (dzip *DataDataplexZoneIamPolicy) DataSource() string {
	return "google_dataplex_zone_iam_policy"
}

// LocalName returns the local name for [DataDataplexZoneIamPolicy].
func (dzip *DataDataplexZoneIamPolicy) LocalName() string {
	return dzip.Name
}

// Configuration returns the configuration (args) for [DataDataplexZoneIamPolicy].
func (dzip *DataDataplexZoneIamPolicy) Configuration() interface{} {
	return dzip.Args
}

// Attributes returns the attributes for [DataDataplexZoneIamPolicy].
func (dzip *DataDataplexZoneIamPolicy) Attributes() dataDataplexZoneIamPolicyAttributes {
	return dataDataplexZoneIamPolicyAttributes{ref: terra.ReferenceDataResource(dzip)}
}

// DataDataplexZoneIamPolicyArgs contains the configurations for google_dataplex_zone_iam_policy.
type DataDataplexZoneIamPolicyArgs struct {
	// DataplexZone: string, required
	DataplexZone terra.StringValue `hcl:"dataplex_zone,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataDataplexZoneIamPolicyAttributes struct {
	ref terra.Reference
}

// DataplexZone returns a reference to field dataplex_zone of google_dataplex_zone_iam_policy.
func (dzip dataDataplexZoneIamPolicyAttributes) DataplexZone() terra.StringValue {
	return terra.ReferenceAsString(dzip.ref.Append("dataplex_zone"))
}

// Etag returns a reference to field etag of google_dataplex_zone_iam_policy.
func (dzip dataDataplexZoneIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dzip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_zone_iam_policy.
func (dzip dataDataplexZoneIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dzip.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_zone_iam_policy.
func (dzip dataDataplexZoneIamPolicyAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dzip.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_zone_iam_policy.
func (dzip dataDataplexZoneIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dzip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataplex_zone_iam_policy.
func (dzip dataDataplexZoneIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dzip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataplex_zone_iam_policy.
func (dzip dataDataplexZoneIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dzip.ref.Append("project"))
}
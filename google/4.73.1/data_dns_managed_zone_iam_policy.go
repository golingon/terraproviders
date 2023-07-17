// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDnsManagedZoneIamPolicy creates a new instance of [DataDnsManagedZoneIamPolicy].
func NewDataDnsManagedZoneIamPolicy(name string, args DataDnsManagedZoneIamPolicyArgs) *DataDnsManagedZoneIamPolicy {
	return &DataDnsManagedZoneIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsManagedZoneIamPolicy)(nil)

// DataDnsManagedZoneIamPolicy represents the Terraform data resource google_dns_managed_zone_iam_policy.
type DataDnsManagedZoneIamPolicy struct {
	Name string
	Args DataDnsManagedZoneIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataDnsManagedZoneIamPolicy].
func (dmzip *DataDnsManagedZoneIamPolicy) DataSource() string {
	return "google_dns_managed_zone_iam_policy"
}

// LocalName returns the local name for [DataDnsManagedZoneIamPolicy].
func (dmzip *DataDnsManagedZoneIamPolicy) LocalName() string {
	return dmzip.Name
}

// Configuration returns the configuration (args) for [DataDnsManagedZoneIamPolicy].
func (dmzip *DataDnsManagedZoneIamPolicy) Configuration() interface{} {
	return dmzip.Args
}

// Attributes returns the attributes for [DataDnsManagedZoneIamPolicy].
func (dmzip *DataDnsManagedZoneIamPolicy) Attributes() dataDnsManagedZoneIamPolicyAttributes {
	return dataDnsManagedZoneIamPolicyAttributes{ref: terra.ReferenceDataResource(dmzip)}
}

// DataDnsManagedZoneIamPolicyArgs contains the configurations for google_dns_managed_zone_iam_policy.
type DataDnsManagedZoneIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedZone: string, required
	ManagedZone terra.StringValue `hcl:"managed_zone,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataDnsManagedZoneIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dns_managed_zone_iam_policy.
func (dmzip dataDnsManagedZoneIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dmzip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dns_managed_zone_iam_policy.
func (dmzip dataDnsManagedZoneIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmzip.ref.Append("id"))
}

// ManagedZone returns a reference to field managed_zone of google_dns_managed_zone_iam_policy.
func (dmzip dataDnsManagedZoneIamPolicyAttributes) ManagedZone() terra.StringValue {
	return terra.ReferenceAsString(dmzip.ref.Append("managed_zone"))
}

// PolicyData returns a reference to field policy_data of google_dns_managed_zone_iam_policy.
func (dmzip dataDnsManagedZoneIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dmzip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dns_managed_zone_iam_policy.
func (dmzip dataDnsManagedZoneIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmzip.ref.Append("project"))
}

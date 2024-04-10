// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"github.com/golingon/lingon/pkg/terra"
	datacloudidentitygrouplookup "github.com/golingon/terraproviders/google/5.24.0/datacloudidentitygrouplookup"
)

// NewDataCloudIdentityGroupLookup creates a new instance of [DataCloudIdentityGroupLookup].
func NewDataCloudIdentityGroupLookup(name string, args DataCloudIdentityGroupLookupArgs) *DataCloudIdentityGroupLookup {
	return &DataCloudIdentityGroupLookup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudIdentityGroupLookup)(nil)

// DataCloudIdentityGroupLookup represents the Terraform data resource google_cloud_identity_group_lookup.
type DataCloudIdentityGroupLookup struct {
	Name string
	Args DataCloudIdentityGroupLookupArgs
}

// DataSource returns the Terraform object type for [DataCloudIdentityGroupLookup].
func (cigl *DataCloudIdentityGroupLookup) DataSource() string {
	return "google_cloud_identity_group_lookup"
}

// LocalName returns the local name for [DataCloudIdentityGroupLookup].
func (cigl *DataCloudIdentityGroupLookup) LocalName() string {
	return cigl.Name
}

// Configuration returns the configuration (args) for [DataCloudIdentityGroupLookup].
func (cigl *DataCloudIdentityGroupLookup) Configuration() interface{} {
	return cigl.Args
}

// Attributes returns the attributes for [DataCloudIdentityGroupLookup].
func (cigl *DataCloudIdentityGroupLookup) Attributes() dataCloudIdentityGroupLookupAttributes {
	return dataCloudIdentityGroupLookupAttributes{ref: terra.ReferenceDataResource(cigl)}
}

// DataCloudIdentityGroupLookupArgs contains the configurations for google_cloud_identity_group_lookup.
type DataCloudIdentityGroupLookupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// GroupKey: required
	GroupKey *datacloudidentitygrouplookup.GroupKey `hcl:"group_key,block" validate:"required"`
}
type dataCloudIdentityGroupLookupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_cloud_identity_group_lookup.
func (cigl dataCloudIdentityGroupLookupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cigl.ref.Append("id"))
}

// Name returns a reference to field name of google_cloud_identity_group_lookup.
func (cigl dataCloudIdentityGroupLookupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cigl.ref.Append("name"))
}

func (cigl dataCloudIdentityGroupLookupAttributes) GroupKey() terra.ListValue[datacloudidentitygrouplookup.GroupKeyAttributes] {
	return terra.ReferenceAsList[datacloudidentitygrouplookup.GroupKeyAttributes](cigl.ref.Append("group_key"))
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacloudidentitygroups "github.com/golingon/terraproviders/googlebeta/4.74.0/datacloudidentitygroups"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudIdentityGroups creates a new instance of [DataCloudIdentityGroups].
func NewDataCloudIdentityGroups(name string, args DataCloudIdentityGroupsArgs) *DataCloudIdentityGroups {
	return &DataCloudIdentityGroups{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudIdentityGroups)(nil)

// DataCloudIdentityGroups represents the Terraform data resource google_cloud_identity_groups.
type DataCloudIdentityGroups struct {
	Name string
	Args DataCloudIdentityGroupsArgs
}

// DataSource returns the Terraform object type for [DataCloudIdentityGroups].
func (cig *DataCloudIdentityGroups) DataSource() string {
	return "google_cloud_identity_groups"
}

// LocalName returns the local name for [DataCloudIdentityGroups].
func (cig *DataCloudIdentityGroups) LocalName() string {
	return cig.Name
}

// Configuration returns the configuration (args) for [DataCloudIdentityGroups].
func (cig *DataCloudIdentityGroups) Configuration() interface{} {
	return cig.Args
}

// Attributes returns the attributes for [DataCloudIdentityGroups].
func (cig *DataCloudIdentityGroups) Attributes() dataCloudIdentityGroupsAttributes {
	return dataCloudIdentityGroupsAttributes{ref: terra.ReferenceDataResource(cig)}
}

// DataCloudIdentityGroupsArgs contains the configurations for google_cloud_identity_groups.
type DataCloudIdentityGroupsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Groups: min=0
	Groups []datacloudidentitygroups.Groups `hcl:"groups,block" validate:"min=0"`
}
type dataCloudIdentityGroupsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_cloud_identity_groups.
func (cig dataCloudIdentityGroupsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("id"))
}

// Parent returns a reference to field parent of google_cloud_identity_groups.
func (cig dataCloudIdentityGroupsAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("parent"))
}

func (cig dataCloudIdentityGroupsAttributes) Groups() terra.ListValue[datacloudidentitygroups.GroupsAttributes] {
	return terra.ReferenceAsList[datacloudidentitygroups.GroupsAttributes](cig.ref.Append("groups"))
}

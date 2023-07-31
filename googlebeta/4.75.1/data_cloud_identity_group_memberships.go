// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacloudidentitygroupmemberships "github.com/golingon/terraproviders/googlebeta/4.75.1/datacloudidentitygroupmemberships"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudIdentityGroupMemberships creates a new instance of [DataCloudIdentityGroupMemberships].
func NewDataCloudIdentityGroupMemberships(name string, args DataCloudIdentityGroupMembershipsArgs) *DataCloudIdentityGroupMemberships {
	return &DataCloudIdentityGroupMemberships{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudIdentityGroupMemberships)(nil)

// DataCloudIdentityGroupMemberships represents the Terraform data resource google_cloud_identity_group_memberships.
type DataCloudIdentityGroupMemberships struct {
	Name string
	Args DataCloudIdentityGroupMembershipsArgs
}

// DataSource returns the Terraform object type for [DataCloudIdentityGroupMemberships].
func (cigm *DataCloudIdentityGroupMemberships) DataSource() string {
	return "google_cloud_identity_group_memberships"
}

// LocalName returns the local name for [DataCloudIdentityGroupMemberships].
func (cigm *DataCloudIdentityGroupMemberships) LocalName() string {
	return cigm.Name
}

// Configuration returns the configuration (args) for [DataCloudIdentityGroupMemberships].
func (cigm *DataCloudIdentityGroupMemberships) Configuration() interface{} {
	return cigm.Args
}

// Attributes returns the attributes for [DataCloudIdentityGroupMemberships].
func (cigm *DataCloudIdentityGroupMemberships) Attributes() dataCloudIdentityGroupMembershipsAttributes {
	return dataCloudIdentityGroupMembershipsAttributes{ref: terra.ReferenceDataResource(cigm)}
}

// DataCloudIdentityGroupMembershipsArgs contains the configurations for google_cloud_identity_group_memberships.
type DataCloudIdentityGroupMembershipsArgs struct {
	// Group: string, required
	Group terra.StringValue `hcl:"group,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Memberships: min=0
	Memberships []datacloudidentitygroupmemberships.Memberships `hcl:"memberships,block" validate:"min=0"`
}
type dataCloudIdentityGroupMembershipsAttributes struct {
	ref terra.Reference
}

// Group returns a reference to field group of google_cloud_identity_group_memberships.
func (cigm dataCloudIdentityGroupMembershipsAttributes) Group() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("group"))
}

// Id returns a reference to field id of google_cloud_identity_group_memberships.
func (cigm dataCloudIdentityGroupMembershipsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cigm.ref.Append("id"))
}

func (cigm dataCloudIdentityGroupMembershipsAttributes) Memberships() terra.ListValue[datacloudidentitygroupmemberships.MembershipsAttributes] {
	return terra.ReferenceAsList[datacloudidentitygroupmemberships.MembershipsAttributes](cigm.ref.Append("memberships"))
}

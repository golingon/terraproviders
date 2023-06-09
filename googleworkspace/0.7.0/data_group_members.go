// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googleworkspace

import (
	datagroupmembers "github.com/golingon/terraproviders/googleworkspace/0.7.0/datagroupmembers"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataGroupMembers creates a new instance of [DataGroupMembers].
func NewDataGroupMembers(name string, args DataGroupMembersArgs) *DataGroupMembers {
	return &DataGroupMembers{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGroupMembers)(nil)

// DataGroupMembers represents the Terraform data resource googleworkspace_group_members.
type DataGroupMembers struct {
	Name string
	Args DataGroupMembersArgs
}

// DataSource returns the Terraform object type for [DataGroupMembers].
func (gm *DataGroupMembers) DataSource() string {
	return "googleworkspace_group_members"
}

// LocalName returns the local name for [DataGroupMembers].
func (gm *DataGroupMembers) LocalName() string {
	return gm.Name
}

// Configuration returns the configuration (args) for [DataGroupMembers].
func (gm *DataGroupMembers) Configuration() interface{} {
	return gm.Args
}

// Attributes returns the attributes for [DataGroupMembers].
func (gm *DataGroupMembers) Attributes() dataGroupMembersAttributes {
	return dataGroupMembersAttributes{ref: terra.ReferenceDataResource(gm)}
}

// DataGroupMembersArgs contains the configurations for googleworkspace_group_members.
type DataGroupMembersArgs struct {
	// GroupId: string, required
	GroupId terra.StringValue `hcl:"group_id,attr" validate:"required"`
	// IncludeDerivedMembership: bool, optional
	IncludeDerivedMembership terra.BoolValue `hcl:"include_derived_membership,attr"`
	// Members: min=0
	Members []datagroupmembers.Members `hcl:"members,block" validate:"min=0"`
}
type dataGroupMembersAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of googleworkspace_group_members.
func (gm dataGroupMembersAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("etag"))
}

// GroupId returns a reference to field group_id of googleworkspace_group_members.
func (gm dataGroupMembersAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("group_id"))
}

// Id returns a reference to field id of googleworkspace_group_members.
func (gm dataGroupMembersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("id"))
}

// IncludeDerivedMembership returns a reference to field include_derived_membership of googleworkspace_group_members.
func (gm dataGroupMembersAttributes) IncludeDerivedMembership() terra.BoolValue {
	return terra.ReferenceAsBool(gm.ref.Append("include_derived_membership"))
}

func (gm dataGroupMembersAttributes) Members() terra.SetValue[datagroupmembers.MembersAttributes] {
	return terra.ReferenceAsSet[datagroupmembers.MembersAttributes](gm.ref.Append("members"))
}

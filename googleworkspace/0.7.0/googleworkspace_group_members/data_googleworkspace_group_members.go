// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googleworkspace_group_members

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource googleworkspace_group_members.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (ggm *DataSource) DataSource() string {
	return "googleworkspace_group_members"
}

// LocalName returns the local name for [DataSource].
func (ggm *DataSource) LocalName() string {
	return ggm.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (ggm *DataSource) Configuration() interface{} {
	return ggm.Args
}

// Attributes returns the attributes for [DataSource].
func (ggm *DataSource) Attributes() dataGoogleworkspaceGroupMembersAttributes {
	return dataGoogleworkspaceGroupMembersAttributes{ref: terra.ReferenceDataSource(ggm)}
}

// DataArgs contains the configurations for googleworkspace_group_members.
type DataArgs struct {
	// GroupId: string, required
	GroupId terra.StringValue `hcl:"group_id,attr" validate:"required"`
	// IncludeDerivedMembership: bool, optional
	IncludeDerivedMembership terra.BoolValue `hcl:"include_derived_membership,attr"`
}

type dataGoogleworkspaceGroupMembersAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of googleworkspace_group_members.
func (ggm dataGoogleworkspaceGroupMembersAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ggm.ref.Append("etag"))
}

// GroupId returns a reference to field group_id of googleworkspace_group_members.
func (ggm dataGoogleworkspaceGroupMembersAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(ggm.ref.Append("group_id"))
}

// Id returns a reference to field id of googleworkspace_group_members.
func (ggm dataGoogleworkspaceGroupMembersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ggm.ref.Append("id"))
}

// IncludeDerivedMembership returns a reference to field include_derived_membership of googleworkspace_group_members.
func (ggm dataGoogleworkspaceGroupMembersAttributes) IncludeDerivedMembership() terra.BoolValue {
	return terra.ReferenceAsBool(ggm.ref.Append("include_derived_membership"))
}

func (ggm dataGoogleworkspaceGroupMembersAttributes) Members() terra.SetValue[DataMembersAttributes] {
	return terra.ReferenceAsSet[DataMembersAttributes](ggm.ref.Append("members"))
}

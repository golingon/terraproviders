// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googleworkspace

import "github.com/golingon/lingon/pkg/terra"

// NewDataGroup creates a new instance of [DataGroup].
func NewDataGroup(name string, args DataGroupArgs) *DataGroup {
	return &DataGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGroup)(nil)

// DataGroup represents the Terraform data resource googleworkspace_group.
type DataGroup struct {
	Name string
	Args DataGroupArgs
}

// DataSource returns the Terraform object type for [DataGroup].
func (g *DataGroup) DataSource() string {
	return "googleworkspace_group"
}

// LocalName returns the local name for [DataGroup].
func (g *DataGroup) LocalName() string {
	return g.Name
}

// Configuration returns the configuration (args) for [DataGroup].
func (g *DataGroup) Configuration() interface{} {
	return g.Args
}

// Attributes returns the attributes for [DataGroup].
func (g *DataGroup) Attributes() dataGroupAttributes {
	return dataGroupAttributes{ref: terra.ReferenceDataResource(g)}
}

// DataGroupArgs contains the configurations for googleworkspace_group.
type DataGroupArgs struct {
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataGroupAttributes struct {
	ref terra.Reference
}

// AdminCreated returns a reference to field admin_created of googleworkspace_group.
func (g dataGroupAttributes) AdminCreated() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("admin_created"))
}

// Aliases returns a reference to field aliases of googleworkspace_group.
func (g dataGroupAttributes) Aliases() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("aliases"))
}

// Description returns a reference to field description of googleworkspace_group.
func (g dataGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("description"))
}

// DirectMembersCount returns a reference to field direct_members_count of googleworkspace_group.
func (g dataGroupAttributes) DirectMembersCount() terra.NumberValue {
	return terra.ReferenceAsNumber(g.ref.Append("direct_members_count"))
}

// Email returns a reference to field email of googleworkspace_group.
func (g dataGroupAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("email"))
}

// Etag returns a reference to field etag of googleworkspace_group.
func (g dataGroupAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("etag"))
}

// Id returns a reference to field id of googleworkspace_group.
func (g dataGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("id"))
}

// Name returns a reference to field name of googleworkspace_group.
func (g dataGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("name"))
}

// NonEditableAliases returns a reference to field non_editable_aliases of googleworkspace_group.
func (g dataGroupAttributes) NonEditableAliases() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("non_editable_aliases"))
}

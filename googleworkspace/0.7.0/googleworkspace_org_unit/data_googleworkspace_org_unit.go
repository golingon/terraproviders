// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googleworkspace_org_unit

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource googleworkspace_org_unit.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gou *DataSource) DataSource() string {
	return "googleworkspace_org_unit"
}

// LocalName returns the local name for [DataSource].
func (gou *DataSource) LocalName() string {
	return gou.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gou *DataSource) Configuration() interface{} {
	return gou.Args
}

// Attributes returns the attributes for [DataSource].
func (gou *DataSource) Attributes() dataGoogleworkspaceOrgUnitAttributes {
	return dataGoogleworkspaceOrgUnitAttributes{ref: terra.ReferenceDataSource(gou)}
}

// DataArgs contains the configurations for googleworkspace_org_unit.
type DataArgs struct {
	// OrgUnitId: string, optional
	OrgUnitId terra.StringValue `hcl:"org_unit_id,attr"`
	// OrgUnitPath: string, optional
	OrgUnitPath terra.StringValue `hcl:"org_unit_path,attr"`
}

type dataGoogleworkspaceOrgUnitAttributes struct {
	ref terra.Reference
}

// BlockInheritance returns a reference to field block_inheritance of googleworkspace_org_unit.
func (gou dataGoogleworkspaceOrgUnitAttributes) BlockInheritance() terra.BoolValue {
	return terra.ReferenceAsBool(gou.ref.Append("block_inheritance"))
}

// Description returns a reference to field description of googleworkspace_org_unit.
func (gou dataGoogleworkspaceOrgUnitAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gou.ref.Append("description"))
}

// Etag returns a reference to field etag of googleworkspace_org_unit.
func (gou dataGoogleworkspaceOrgUnitAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gou.ref.Append("etag"))
}

// Id returns a reference to field id of googleworkspace_org_unit.
func (gou dataGoogleworkspaceOrgUnitAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gou.ref.Append("id"))
}

// Name returns a reference to field name of googleworkspace_org_unit.
func (gou dataGoogleworkspaceOrgUnitAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gou.ref.Append("name"))
}

// OrgUnitId returns a reference to field org_unit_id of googleworkspace_org_unit.
func (gou dataGoogleworkspaceOrgUnitAttributes) OrgUnitId() terra.StringValue {
	return terra.ReferenceAsString(gou.ref.Append("org_unit_id"))
}

// OrgUnitPath returns a reference to field org_unit_path of googleworkspace_org_unit.
func (gou dataGoogleworkspaceOrgUnitAttributes) OrgUnitPath() terra.StringValue {
	return terra.ReferenceAsString(gou.ref.Append("org_unit_path"))
}

// ParentOrgUnitId returns a reference to field parent_org_unit_id of googleworkspace_org_unit.
func (gou dataGoogleworkspaceOrgUnitAttributes) ParentOrgUnitId() terra.StringValue {
	return terra.ReferenceAsString(gou.ref.Append("parent_org_unit_id"))
}

// ParentOrgUnitPath returns a reference to field parent_org_unit_path of googleworkspace_org_unit.
func (gou dataGoogleworkspaceOrgUnitAttributes) ParentOrgUnitPath() terra.StringValue {
	return terra.ReferenceAsString(gou.ref.Append("parent_org_unit_path"))
}

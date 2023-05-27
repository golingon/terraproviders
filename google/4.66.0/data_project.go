// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataProject creates a new instance of [DataProject].
func NewDataProject(name string, args DataProjectArgs) *DataProject {
	return &DataProject{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataProject)(nil)

// DataProject represents the Terraform data resource google_project.
type DataProject struct {
	Name string
	Args DataProjectArgs
}

// DataSource returns the Terraform object type for [DataProject].
func (p *DataProject) DataSource() string {
	return "google_project"
}

// LocalName returns the local name for [DataProject].
func (p *DataProject) LocalName() string {
	return p.Name
}

// Configuration returns the configuration (args) for [DataProject].
func (p *DataProject) Configuration() interface{} {
	return p.Args
}

// Attributes returns the attributes for [DataProject].
func (p *DataProject) Attributes() dataProjectAttributes {
	return dataProjectAttributes{ref: terra.ReferenceDataResource(p)}
}

// DataProjectArgs contains the configurations for google_project.
type DataProjectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProjectId: string, optional
	ProjectId terra.StringValue `hcl:"project_id,attr"`
}
type dataProjectAttributes struct {
	ref terra.Reference
}

// AutoCreateNetwork returns a reference to field auto_create_network of google_project.
func (p dataProjectAttributes) AutoCreateNetwork() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("auto_create_network"))
}

// BillingAccount returns a reference to field billing_account of google_project.
func (p dataProjectAttributes) BillingAccount() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("billing_account"))
}

// FolderId returns a reference to field folder_id of google_project.
func (p dataProjectAttributes) FolderId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("folder_id"))
}

// Id returns a reference to field id of google_project.
func (p dataProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("id"))
}

// Labels returns a reference to field labels of google_project.
func (p dataProjectAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](p.ref.Append("labels"))
}

// Name returns a reference to field name of google_project.
func (p dataProjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

// Number returns a reference to field number of google_project.
func (p dataProjectAttributes) Number() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("number"))
}

// OrgId returns a reference to field org_id of google_project.
func (p dataProjectAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("org_id"))
}

// ProjectId returns a reference to field project_id of google_project.
func (p dataProjectAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("project_id"))
}

// SkipDelete returns a reference to field skip_delete of google_project.
func (p dataProjectAttributes) SkipDelete() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("skip_delete"))
}

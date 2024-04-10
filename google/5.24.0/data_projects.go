// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"github.com/golingon/lingon/pkg/terra"
	dataprojects "github.com/golingon/terraproviders/google/5.24.0/dataprojects"
)

// NewDataProjects creates a new instance of [DataProjects].
func NewDataProjects(name string, args DataProjectsArgs) *DataProjects {
	return &DataProjects{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataProjects)(nil)

// DataProjects represents the Terraform data resource google_projects.
type DataProjects struct {
	Name string
	Args DataProjectsArgs
}

// DataSource returns the Terraform object type for [DataProjects].
func (p *DataProjects) DataSource() string {
	return "google_projects"
}

// LocalName returns the local name for [DataProjects].
func (p *DataProjects) LocalName() string {
	return p.Name
}

// Configuration returns the configuration (args) for [DataProjects].
func (p *DataProjects) Configuration() interface{} {
	return p.Args
}

// Attributes returns the attributes for [DataProjects].
func (p *DataProjects) Attributes() dataProjectsAttributes {
	return dataProjectsAttributes{ref: terra.ReferenceDataResource(p)}
}

// DataProjectsArgs contains the configurations for google_projects.
type DataProjectsArgs struct {
	// Filter: string, required
	Filter terra.StringValue `hcl:"filter,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Projects: min=0
	Projects []dataprojects.Projects `hcl:"projects,block" validate:"min=0"`
}
type dataProjectsAttributes struct {
	ref terra.Reference
}

// Filter returns a reference to field filter of google_projects.
func (p dataProjectsAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("filter"))
}

// Id returns a reference to field id of google_projects.
func (p dataProjectsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("id"))
}

func (p dataProjectsAttributes) Projects() terra.ListValue[dataprojects.ProjectsAttributes] {
	return terra.ReferenceAsList[dataprojects.ProjectsAttributes](p.ref.Append("projects"))
}

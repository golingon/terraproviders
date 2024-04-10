// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataProjectService creates a new instance of [DataProjectService].
func NewDataProjectService(name string, args DataProjectServiceArgs) *DataProjectService {
	return &DataProjectService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataProjectService)(nil)

// DataProjectService represents the Terraform data resource google_project_service.
type DataProjectService struct {
	Name string
	Args DataProjectServiceArgs
}

// DataSource returns the Terraform object type for [DataProjectService].
func (ps *DataProjectService) DataSource() string {
	return "google_project_service"
}

// LocalName returns the local name for [DataProjectService].
func (ps *DataProjectService) LocalName() string {
	return ps.Name
}

// Configuration returns the configuration (args) for [DataProjectService].
func (ps *DataProjectService) Configuration() interface{} {
	return ps.Args
}

// Attributes returns the attributes for [DataProjectService].
func (ps *DataProjectService) Attributes() dataProjectServiceAttributes {
	return dataProjectServiceAttributes{ref: terra.ReferenceDataResource(ps)}
}

// DataProjectServiceArgs contains the configurations for google_project_service.
type DataProjectServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
}
type dataProjectServiceAttributes struct {
	ref terra.Reference
}

// DisableDependentServices returns a reference to field disable_dependent_services of google_project_service.
func (ps dataProjectServiceAttributes) DisableDependentServices() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("disable_dependent_services"))
}

// DisableOnDestroy returns a reference to field disable_on_destroy of google_project_service.
func (ps dataProjectServiceAttributes) DisableOnDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("disable_on_destroy"))
}

// Id returns a reference to field id of google_project_service.
func (ps dataProjectServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("id"))
}

// Project returns a reference to field project of google_project_service.
func (ps dataProjectServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("project"))
}

// Service returns a reference to field service of google_project_service.
func (ps dataProjectServiceAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("service"))
}

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_runtimeconfig_variable

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_runtimeconfig_variable.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (grv *DataSource) DataSource() string {
	return "google_runtimeconfig_variable"
}

// LocalName returns the local name for [DataSource].
func (grv *DataSource) LocalName() string {
	return grv.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (grv *DataSource) Configuration() interface{} {
	return grv.Args
}

// Attributes returns the attributes for [DataSource].
func (grv *DataSource) Attributes() dataGoogleRuntimeconfigVariableAttributes {
	return dataGoogleRuntimeconfigVariableAttributes{ref: terra.ReferenceDataSource(grv)}
}

// DataArgs contains the configurations for google_runtimeconfig_variable.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type dataGoogleRuntimeconfigVariableAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_runtimeconfig_variable.
func (grv dataGoogleRuntimeconfigVariableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(grv.ref.Append("id"))
}

// Name returns a reference to field name of google_runtimeconfig_variable.
func (grv dataGoogleRuntimeconfigVariableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(grv.ref.Append("name"))
}

// Parent returns a reference to field parent of google_runtimeconfig_variable.
func (grv dataGoogleRuntimeconfigVariableAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(grv.ref.Append("parent"))
}

// Project returns a reference to field project of google_runtimeconfig_variable.
func (grv dataGoogleRuntimeconfigVariableAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(grv.ref.Append("project"))
}

// Text returns a reference to field text of google_runtimeconfig_variable.
func (grv dataGoogleRuntimeconfigVariableAttributes) Text() terra.StringValue {
	return terra.ReferenceAsString(grv.ref.Append("text"))
}

// UpdateTime returns a reference to field update_time of google_runtimeconfig_variable.
func (grv dataGoogleRuntimeconfigVariableAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(grv.ref.Append("update_time"))
}

// Value returns a reference to field value of google_runtimeconfig_variable.
func (grv dataGoogleRuntimeconfigVariableAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(grv.ref.Append("value"))
}

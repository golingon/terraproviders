// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataRuntimeconfigConfig creates a new instance of [DataRuntimeconfigConfig].
func NewDataRuntimeconfigConfig(name string, args DataRuntimeconfigConfigArgs) *DataRuntimeconfigConfig {
	return &DataRuntimeconfigConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRuntimeconfigConfig)(nil)

// DataRuntimeconfigConfig represents the Terraform data resource google_runtimeconfig_config.
type DataRuntimeconfigConfig struct {
	Name string
	Args DataRuntimeconfigConfigArgs
}

// DataSource returns the Terraform object type for [DataRuntimeconfigConfig].
func (rc *DataRuntimeconfigConfig) DataSource() string {
	return "google_runtimeconfig_config"
}

// LocalName returns the local name for [DataRuntimeconfigConfig].
func (rc *DataRuntimeconfigConfig) LocalName() string {
	return rc.Name
}

// Configuration returns the configuration (args) for [DataRuntimeconfigConfig].
func (rc *DataRuntimeconfigConfig) Configuration() interface{} {
	return rc.Args
}

// Attributes returns the attributes for [DataRuntimeconfigConfig].
func (rc *DataRuntimeconfigConfig) Attributes() dataRuntimeconfigConfigAttributes {
	return dataRuntimeconfigConfigAttributes{ref: terra.ReferenceDataResource(rc)}
}

// DataRuntimeconfigConfigArgs contains the configurations for google_runtimeconfig_config.
type DataRuntimeconfigConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataRuntimeconfigConfigAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_runtimeconfig_config.
func (rc dataRuntimeconfigConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("description"))
}

// Id returns a reference to field id of google_runtimeconfig_config.
func (rc dataRuntimeconfigConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("id"))
}

// Name returns a reference to field name of google_runtimeconfig_config.
func (rc dataRuntimeconfigConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("name"))
}

// Project returns a reference to field project of google_runtimeconfig_config.
func (rc dataRuntimeconfigConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("project"))
}

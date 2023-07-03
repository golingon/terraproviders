// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomposerenvironment "github.com/golingon/terraproviders/google/4.71.0/datacomposerenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComposerEnvironment creates a new instance of [DataComposerEnvironment].
func NewDataComposerEnvironment(name string, args DataComposerEnvironmentArgs) *DataComposerEnvironment {
	return &DataComposerEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComposerEnvironment)(nil)

// DataComposerEnvironment represents the Terraform data resource google_composer_environment.
type DataComposerEnvironment struct {
	Name string
	Args DataComposerEnvironmentArgs
}

// DataSource returns the Terraform object type for [DataComposerEnvironment].
func (ce *DataComposerEnvironment) DataSource() string {
	return "google_composer_environment"
}

// LocalName returns the local name for [DataComposerEnvironment].
func (ce *DataComposerEnvironment) LocalName() string {
	return ce.Name
}

// Configuration returns the configuration (args) for [DataComposerEnvironment].
func (ce *DataComposerEnvironment) Configuration() interface{} {
	return ce.Args
}

// Attributes returns the attributes for [DataComposerEnvironment].
func (ce *DataComposerEnvironment) Attributes() dataComposerEnvironmentAttributes {
	return dataComposerEnvironmentAttributes{ref: terra.ReferenceDataResource(ce)}
}

// DataComposerEnvironmentArgs contains the configurations for google_composer_environment.
type DataComposerEnvironmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Config: min=0
	Config []datacomposerenvironment.Config `hcl:"config,block" validate:"min=0"`
}
type dataComposerEnvironmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_composer_environment.
func (ce dataComposerEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("id"))
}

// Labels returns a reference to field labels of google_composer_environment.
func (ce dataComposerEnvironmentAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ce.ref.Append("labels"))
}

// Name returns a reference to field name of google_composer_environment.
func (ce dataComposerEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("name"))
}

// Project returns a reference to field project of google_composer_environment.
func (ce dataComposerEnvironmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("project"))
}

// Region returns a reference to field region of google_composer_environment.
func (ce dataComposerEnvironmentAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("region"))
}

func (ce dataComposerEnvironmentAttributes) Config() terra.ListValue[datacomposerenvironment.ConfigAttributes] {
	return terra.ReferenceAsList[datacomposerenvironment.ConfigAttributes](ce.ref.Append("config"))
}

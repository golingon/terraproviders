// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomposerimageversions "github.com/golingon/terraproviders/google/4.59.0/datacomposerimageversions"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComposerImageVersions creates a new instance of [DataComposerImageVersions].
func NewDataComposerImageVersions(name string, args DataComposerImageVersionsArgs) *DataComposerImageVersions {
	return &DataComposerImageVersions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComposerImageVersions)(nil)

// DataComposerImageVersions represents the Terraform data resource google_composer_image_versions.
type DataComposerImageVersions struct {
	Name string
	Args DataComposerImageVersionsArgs
}

// DataSource returns the Terraform object type for [DataComposerImageVersions].
func (civ *DataComposerImageVersions) DataSource() string {
	return "google_composer_image_versions"
}

// LocalName returns the local name for [DataComposerImageVersions].
func (civ *DataComposerImageVersions) LocalName() string {
	return civ.Name
}

// Configuration returns the configuration (args) for [DataComposerImageVersions].
func (civ *DataComposerImageVersions) Configuration() interface{} {
	return civ.Args
}

// Attributes returns the attributes for [DataComposerImageVersions].
func (civ *DataComposerImageVersions) Attributes() dataComposerImageVersionsAttributes {
	return dataComposerImageVersionsAttributes{ref: terra.ReferenceDataResource(civ)}
}

// DataComposerImageVersionsArgs contains the configurations for google_composer_image_versions.
type DataComposerImageVersionsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ImageVersions: min=0
	ImageVersions []datacomposerimageversions.ImageVersions `hcl:"image_versions,block" validate:"min=0"`
}
type dataComposerImageVersionsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_composer_image_versions.
func (civ dataComposerImageVersionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(civ.ref.Append("id"))
}

// Project returns a reference to field project of google_composer_image_versions.
func (civ dataComposerImageVersionsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(civ.ref.Append("project"))
}

// Region returns a reference to field region of google_composer_image_versions.
func (civ dataComposerImageVersionsAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(civ.ref.Append("region"))
}

func (civ dataComposerImageVersionsAttributes) ImageVersions() terra.ListValue[datacomposerimageversions.ImageVersionsAttributes] {
	return terra.ReferenceAsList[datacomposerimageversions.ImageVersionsAttributes](civ.ref.Append("image_versions"))
}

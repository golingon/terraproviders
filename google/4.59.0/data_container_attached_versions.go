// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataContainerAttachedVersions(name string, args DataContainerAttachedVersionsArgs) *DataContainerAttachedVersions {
	return &DataContainerAttachedVersions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerAttachedVersions)(nil)

type DataContainerAttachedVersions struct {
	Name string
	Args DataContainerAttachedVersionsArgs
}

func (cav *DataContainerAttachedVersions) DataSource() string {
	return "google_container_attached_versions"
}

func (cav *DataContainerAttachedVersions) LocalName() string {
	return cav.Name
}

func (cav *DataContainerAttachedVersions) Configuration() interface{} {
	return cav.Args
}

func (cav *DataContainerAttachedVersions) Attributes() dataContainerAttachedVersionsAttributes {
	return dataContainerAttachedVersionsAttributes{ref: terra.ReferenceDataResource(cav)}
}

type DataContainerAttachedVersionsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
}
type dataContainerAttachedVersionsAttributes struct {
	ref terra.Reference
}

func (cav dataContainerAttachedVersionsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cav.ref.Append("id"))
}

func (cav dataContainerAttachedVersionsAttributes) Location() terra.StringValue {
	return terra.ReferenceString(cav.ref.Append("location"))
}

func (cav dataContainerAttachedVersionsAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cav.ref.Append("project"))
}

func (cav dataContainerAttachedVersionsAttributes) ValidVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](cav.ref.Append("valid_versions"))
}

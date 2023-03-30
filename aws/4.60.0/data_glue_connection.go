// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataglueconnection "github.com/golingon/terraproviders/aws/4.60.0/dataglueconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataGlueConnection(name string, args DataGlueConnectionArgs) *DataGlueConnection {
	return &DataGlueConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGlueConnection)(nil)

type DataGlueConnection struct {
	Name string
	Args DataGlueConnectionArgs
}

func (gc *DataGlueConnection) DataSource() string {
	return "aws_glue_connection"
}

func (gc *DataGlueConnection) LocalName() string {
	return gc.Name
}

func (gc *DataGlueConnection) Configuration() interface{} {
	return gc.Args
}

func (gc *DataGlueConnection) Attributes() dataGlueConnectionAttributes {
	return dataGlueConnectionAttributes{ref: terra.ReferenceDataResource(gc)}
}

type DataGlueConnectionArgs struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// PhysicalConnectionRequirements: min=0
	PhysicalConnectionRequirements []dataglueconnection.PhysicalConnectionRequirements `hcl:"physical_connection_requirements,block" validate:"min=0"`
}
type dataGlueConnectionAttributes struct {
	ref terra.Reference
}

func (gc dataGlueConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("arn"))
}

func (gc dataGlueConnectionAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("catalog_id"))
}

func (gc dataGlueConnectionAttributes) ConnectionProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gc.ref.Append("connection_properties"))
}

func (gc dataGlueConnectionAttributes) ConnectionType() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("connection_type"))
}

func (gc dataGlueConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("description"))
}

func (gc dataGlueConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("id"))
}

func (gc dataGlueConnectionAttributes) MatchCriteria() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](gc.ref.Append("match_criteria"))
}

func (gc dataGlueConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("name"))
}

func (gc dataGlueConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gc.ref.Append("tags"))
}

func (gc dataGlueConnectionAttributes) PhysicalConnectionRequirements() terra.ListValue[dataglueconnection.PhysicalConnectionRequirementsAttributes] {
	return terra.ReferenceList[dataglueconnection.PhysicalConnectionRequirementsAttributes](gc.ref.Append("physical_connection_requirements"))
}

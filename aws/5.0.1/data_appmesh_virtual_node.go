// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataappmeshvirtualnode "github.com/golingon/terraproviders/aws/5.0.1/dataappmeshvirtualnode"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppmeshVirtualNode creates a new instance of [DataAppmeshVirtualNode].
func NewDataAppmeshVirtualNode(name string, args DataAppmeshVirtualNodeArgs) *DataAppmeshVirtualNode {
	return &DataAppmeshVirtualNode{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppmeshVirtualNode)(nil)

// DataAppmeshVirtualNode represents the Terraform data resource aws_appmesh_virtual_node.
type DataAppmeshVirtualNode struct {
	Name string
	Args DataAppmeshVirtualNodeArgs
}

// DataSource returns the Terraform object type for [DataAppmeshVirtualNode].
func (avn *DataAppmeshVirtualNode) DataSource() string {
	return "aws_appmesh_virtual_node"
}

// LocalName returns the local name for [DataAppmeshVirtualNode].
func (avn *DataAppmeshVirtualNode) LocalName() string {
	return avn.Name
}

// Configuration returns the configuration (args) for [DataAppmeshVirtualNode].
func (avn *DataAppmeshVirtualNode) Configuration() interface{} {
	return avn.Args
}

// Attributes returns the attributes for [DataAppmeshVirtualNode].
func (avn *DataAppmeshVirtualNode) Attributes() dataAppmeshVirtualNodeAttributes {
	return dataAppmeshVirtualNodeAttributes{ref: terra.ReferenceDataResource(avn)}
}

// DataAppmeshVirtualNodeArgs contains the configurations for aws_appmesh_virtual_node.
type DataAppmeshVirtualNodeArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MeshName: string, required
	MeshName terra.StringValue `hcl:"mesh_name,attr" validate:"required"`
	// MeshOwner: string, optional
	MeshOwner terra.StringValue `hcl:"mesh_owner,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Spec: min=0
	Spec []dataappmeshvirtualnode.Spec `hcl:"spec,block" validate:"min=0"`
}
type dataAppmeshVirtualNodeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_virtual_node.
func (avn dataAppmeshVirtualNodeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_virtual_node.
func (avn dataAppmeshVirtualNodeAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_virtual_node.
func (avn dataAppmeshVirtualNodeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_virtual_node.
func (avn dataAppmeshVirtualNodeAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_virtual_node.
func (avn dataAppmeshVirtualNodeAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_virtual_node.
func (avn dataAppmeshVirtualNodeAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_virtual_node.
func (avn dataAppmeshVirtualNodeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_virtual_node.
func (avn dataAppmeshVirtualNodeAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_virtual_node.
func (avn dataAppmeshVirtualNodeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avn.ref.Append("tags"))
}

func (avn dataAppmeshVirtualNodeAttributes) Spec() terra.ListValue[dataappmeshvirtualnode.SpecAttributes] {
	return terra.ReferenceAsList[dataappmeshvirtualnode.SpecAttributes](avn.ref.Append("spec"))
}

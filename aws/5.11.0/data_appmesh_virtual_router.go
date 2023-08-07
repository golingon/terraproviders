// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataappmeshvirtualrouter "github.com/golingon/terraproviders/aws/5.11.0/dataappmeshvirtualrouter"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppmeshVirtualRouter creates a new instance of [DataAppmeshVirtualRouter].
func NewDataAppmeshVirtualRouter(name string, args DataAppmeshVirtualRouterArgs) *DataAppmeshVirtualRouter {
	return &DataAppmeshVirtualRouter{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppmeshVirtualRouter)(nil)

// DataAppmeshVirtualRouter represents the Terraform data resource aws_appmesh_virtual_router.
type DataAppmeshVirtualRouter struct {
	Name string
	Args DataAppmeshVirtualRouterArgs
}

// DataSource returns the Terraform object type for [DataAppmeshVirtualRouter].
func (avr *DataAppmeshVirtualRouter) DataSource() string {
	return "aws_appmesh_virtual_router"
}

// LocalName returns the local name for [DataAppmeshVirtualRouter].
func (avr *DataAppmeshVirtualRouter) LocalName() string {
	return avr.Name
}

// Configuration returns the configuration (args) for [DataAppmeshVirtualRouter].
func (avr *DataAppmeshVirtualRouter) Configuration() interface{} {
	return avr.Args
}

// Attributes returns the attributes for [DataAppmeshVirtualRouter].
func (avr *DataAppmeshVirtualRouter) Attributes() dataAppmeshVirtualRouterAttributes {
	return dataAppmeshVirtualRouterAttributes{ref: terra.ReferenceDataResource(avr)}
}

// DataAppmeshVirtualRouterArgs contains the configurations for aws_appmesh_virtual_router.
type DataAppmeshVirtualRouterArgs struct {
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
	Spec []dataappmeshvirtualrouter.Spec `hcl:"spec,block" validate:"min=0"`
}
type dataAppmeshVirtualRouterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_virtual_router.
func (avr dataAppmeshVirtualRouterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_virtual_router.
func (avr dataAppmeshVirtualRouterAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_virtual_router.
func (avr dataAppmeshVirtualRouterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_virtual_router.
func (avr dataAppmeshVirtualRouterAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_virtual_router.
func (avr dataAppmeshVirtualRouterAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_virtual_router.
func (avr dataAppmeshVirtualRouterAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_virtual_router.
func (avr dataAppmeshVirtualRouterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_virtual_router.
func (avr dataAppmeshVirtualRouterAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_virtual_router.
func (avr dataAppmeshVirtualRouterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avr.ref.Append("tags"))
}

func (avr dataAppmeshVirtualRouterAttributes) Spec() terra.ListValue[dataappmeshvirtualrouter.SpecAttributes] {
	return terra.ReferenceAsList[dataappmeshvirtualrouter.SpecAttributes](avr.ref.Append("spec"))
}

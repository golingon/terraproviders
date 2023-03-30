// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataappmeshroute "github.com/golingon/terraproviders/aws/4.60.0/dataappmeshroute"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataAppmeshRoute(name string, args DataAppmeshRouteArgs) *DataAppmeshRoute {
	return &DataAppmeshRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppmeshRoute)(nil)

type DataAppmeshRoute struct {
	Name string
	Args DataAppmeshRouteArgs
}

func (ar *DataAppmeshRoute) DataSource() string {
	return "aws_appmesh_route"
}

func (ar *DataAppmeshRoute) LocalName() string {
	return ar.Name
}

func (ar *DataAppmeshRoute) Configuration() interface{} {
	return ar.Args
}

func (ar *DataAppmeshRoute) Attributes() dataAppmeshRouteAttributes {
	return dataAppmeshRouteAttributes{ref: terra.ReferenceDataResource(ar)}
}

type DataAppmeshRouteArgs struct {
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
	// VirtualRouterName: string, required
	VirtualRouterName terra.StringValue `hcl:"virtual_router_name,attr" validate:"required"`
	// Spec: min=0
	Spec []dataappmeshroute.Spec `hcl:"spec,block" validate:"min=0"`
}
type dataAppmeshRouteAttributes struct {
	ref terra.Reference
}

func (ar dataAppmeshRouteAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("arn"))
}

func (ar dataAppmeshRouteAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("created_date"))
}

func (ar dataAppmeshRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("id"))
}

func (ar dataAppmeshRouteAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("last_updated_date"))
}

func (ar dataAppmeshRouteAttributes) MeshName() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("mesh_name"))
}

func (ar dataAppmeshRouteAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("mesh_owner"))
}

func (ar dataAppmeshRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("name"))
}

func (ar dataAppmeshRouteAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("resource_owner"))
}

func (ar dataAppmeshRouteAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ar.ref.Append("tags"))
}

func (ar dataAppmeshRouteAttributes) VirtualRouterName() terra.StringValue {
	return terra.ReferenceString(ar.ref.Append("virtual_router_name"))
}

func (ar dataAppmeshRouteAttributes) Spec() terra.ListValue[dataappmeshroute.SpecAttributes] {
	return terra.ReferenceList[dataappmeshroute.SpecAttributes](ar.ref.Append("spec"))
}

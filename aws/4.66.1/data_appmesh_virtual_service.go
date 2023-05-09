// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataappmeshvirtualservice "github.com/golingon/terraproviders/aws/4.66.1/dataappmeshvirtualservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppmeshVirtualService creates a new instance of [DataAppmeshVirtualService].
func NewDataAppmeshVirtualService(name string, args DataAppmeshVirtualServiceArgs) *DataAppmeshVirtualService {
	return &DataAppmeshVirtualService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppmeshVirtualService)(nil)

// DataAppmeshVirtualService represents the Terraform data resource aws_appmesh_virtual_service.
type DataAppmeshVirtualService struct {
	Name string
	Args DataAppmeshVirtualServiceArgs
}

// DataSource returns the Terraform object type for [DataAppmeshVirtualService].
func (avs *DataAppmeshVirtualService) DataSource() string {
	return "aws_appmesh_virtual_service"
}

// LocalName returns the local name for [DataAppmeshVirtualService].
func (avs *DataAppmeshVirtualService) LocalName() string {
	return avs.Name
}

// Configuration returns the configuration (args) for [DataAppmeshVirtualService].
func (avs *DataAppmeshVirtualService) Configuration() interface{} {
	return avs.Args
}

// Attributes returns the attributes for [DataAppmeshVirtualService].
func (avs *DataAppmeshVirtualService) Attributes() dataAppmeshVirtualServiceAttributes {
	return dataAppmeshVirtualServiceAttributes{ref: terra.ReferenceDataResource(avs)}
}

// DataAppmeshVirtualServiceArgs contains the configurations for aws_appmesh_virtual_service.
type DataAppmeshVirtualServiceArgs struct {
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
	Spec []dataappmeshvirtualservice.Spec `hcl:"spec,block" validate:"min=0"`
}
type dataAppmeshVirtualServiceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_virtual_service.
func (avs dataAppmeshVirtualServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_virtual_service.
func (avs dataAppmeshVirtualServiceAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_virtual_service.
func (avs dataAppmeshVirtualServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_virtual_service.
func (avs dataAppmeshVirtualServiceAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_virtual_service.
func (avs dataAppmeshVirtualServiceAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_virtual_service.
func (avs dataAppmeshVirtualServiceAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_virtual_service.
func (avs dataAppmeshVirtualServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_virtual_service.
func (avs dataAppmeshVirtualServiceAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_virtual_service.
func (avs dataAppmeshVirtualServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avs.ref.Append("tags"))
}

func (avs dataAppmeshVirtualServiceAttributes) Spec() terra.ListValue[dataappmeshvirtualservice.SpecAttributes] {
	return terra.ReferenceAsList[dataappmeshvirtualservice.SpecAttributes](avs.ref.Append("spec"))
}

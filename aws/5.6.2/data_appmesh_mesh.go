// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataappmeshmesh "github.com/golingon/terraproviders/aws/5.6.2/dataappmeshmesh"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppmeshMesh creates a new instance of [DataAppmeshMesh].
func NewDataAppmeshMesh(name string, args DataAppmeshMeshArgs) *DataAppmeshMesh {
	return &DataAppmeshMesh{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppmeshMesh)(nil)

// DataAppmeshMesh represents the Terraform data resource aws_appmesh_mesh.
type DataAppmeshMesh struct {
	Name string
	Args DataAppmeshMeshArgs
}

// DataSource returns the Terraform object type for [DataAppmeshMesh].
func (am *DataAppmeshMesh) DataSource() string {
	return "aws_appmesh_mesh"
}

// LocalName returns the local name for [DataAppmeshMesh].
func (am *DataAppmeshMesh) LocalName() string {
	return am.Name
}

// Configuration returns the configuration (args) for [DataAppmeshMesh].
func (am *DataAppmeshMesh) Configuration() interface{} {
	return am.Args
}

// Attributes returns the attributes for [DataAppmeshMesh].
func (am *DataAppmeshMesh) Attributes() dataAppmeshMeshAttributes {
	return dataAppmeshMeshAttributes{ref: terra.ReferenceDataResource(am)}
}

// DataAppmeshMeshArgs contains the configurations for aws_appmesh_mesh.
type DataAppmeshMeshArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MeshOwner: string, optional
	MeshOwner terra.StringValue `hcl:"mesh_owner,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Spec: min=0
	Spec []dataappmeshmesh.Spec `hcl:"spec,block" validate:"min=0"`
}
type dataAppmeshMeshAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_mesh.
func (am dataAppmeshMeshAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_mesh.
func (am dataAppmeshMeshAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_mesh.
func (am dataAppmeshMeshAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_mesh.
func (am dataAppmeshMeshAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("last_updated_date"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_mesh.
func (am dataAppmeshMeshAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_mesh.
func (am dataAppmeshMeshAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_mesh.
func (am dataAppmeshMeshAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_mesh.
func (am dataAppmeshMeshAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](am.ref.Append("tags"))
}

func (am dataAppmeshMeshAttributes) Spec() terra.ListValue[dataappmeshmesh.SpecAttributes] {
	return terra.ReferenceAsList[dataappmeshmesh.SpecAttributes](am.ref.Append("spec"))
}

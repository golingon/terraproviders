// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appmeshmesh "github.com/golingon/terraproviders/aws/4.60.0/appmeshmesh"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAppmeshMesh(name string, args AppmeshMeshArgs) *AppmeshMesh {
	return &AppmeshMesh{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppmeshMesh)(nil)

type AppmeshMesh struct {
	Name  string
	Args  AppmeshMeshArgs
	state *appmeshMeshState
}

func (am *AppmeshMesh) Type() string {
	return "aws_appmesh_mesh"
}

func (am *AppmeshMesh) LocalName() string {
	return am.Name
}

func (am *AppmeshMesh) Configuration() interface{} {
	return am.Args
}

func (am *AppmeshMesh) Attributes() appmeshMeshAttributes {
	return appmeshMeshAttributes{ref: terra.ReferenceResource(am)}
}

func (am *AppmeshMesh) ImportState(av io.Reader) error {
	am.state = &appmeshMeshState{}
	if err := json.NewDecoder(av).Decode(am.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", am.Type(), am.LocalName(), err)
	}
	return nil
}

func (am *AppmeshMesh) State() (*appmeshMeshState, bool) {
	return am.state, am.state != nil
}

func (am *AppmeshMesh) StateMust() *appmeshMeshState {
	if am.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", am.Type(), am.LocalName()))
	}
	return am.state
}

func (am *AppmeshMesh) DependOn() terra.Reference {
	return terra.ReferenceResource(am)
}

type AppmeshMeshArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Spec: optional
	Spec *appmeshmesh.Spec `hcl:"spec,block"`
	// DependsOn contains resources that AppmeshMesh depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type appmeshMeshAttributes struct {
	ref terra.Reference
}

func (am appmeshMeshAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("arn"))
}

func (am appmeshMeshAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("created_date"))
}

func (am appmeshMeshAttributes) Id() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("id"))
}

func (am appmeshMeshAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("last_updated_date"))
}

func (am appmeshMeshAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("mesh_owner"))
}

func (am appmeshMeshAttributes) Name() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("name"))
}

func (am appmeshMeshAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("resource_owner"))
}

func (am appmeshMeshAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](am.ref.Append("tags"))
}

func (am appmeshMeshAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](am.ref.Append("tags_all"))
}

func (am appmeshMeshAttributes) Spec() terra.ListValue[appmeshmesh.SpecAttributes] {
	return terra.ReferenceList[appmeshmesh.SpecAttributes](am.ref.Append("spec"))
}

type appmeshMeshState struct {
	Arn             string                  `json:"arn"`
	CreatedDate     string                  `json:"created_date"`
	Id              string                  `json:"id"`
	LastUpdatedDate string                  `json:"last_updated_date"`
	MeshOwner       string                  `json:"mesh_owner"`
	Name            string                  `json:"name"`
	ResourceOwner   string                  `json:"resource_owner"`
	Tags            map[string]string       `json:"tags"`
	TagsAll         map[string]string       `json:"tags_all"`
	Spec            []appmeshmesh.SpecState `json:"spec"`
}

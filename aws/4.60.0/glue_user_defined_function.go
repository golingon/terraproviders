// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	glueuserdefinedfunction "github.com/golingon/terraproviders/aws/4.60.0/glueuserdefinedfunction"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewGlueUserDefinedFunction(name string, args GlueUserDefinedFunctionArgs) *GlueUserDefinedFunction {
	return &GlueUserDefinedFunction{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueUserDefinedFunction)(nil)

type GlueUserDefinedFunction struct {
	Name  string
	Args  GlueUserDefinedFunctionArgs
	state *glueUserDefinedFunctionState
}

func (gudf *GlueUserDefinedFunction) Type() string {
	return "aws_glue_user_defined_function"
}

func (gudf *GlueUserDefinedFunction) LocalName() string {
	return gudf.Name
}

func (gudf *GlueUserDefinedFunction) Configuration() interface{} {
	return gudf.Args
}

func (gudf *GlueUserDefinedFunction) Attributes() glueUserDefinedFunctionAttributes {
	return glueUserDefinedFunctionAttributes{ref: terra.ReferenceResource(gudf)}
}

func (gudf *GlueUserDefinedFunction) ImportState(av io.Reader) error {
	gudf.state = &glueUserDefinedFunctionState{}
	if err := json.NewDecoder(av).Decode(gudf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gudf.Type(), gudf.LocalName(), err)
	}
	return nil
}

func (gudf *GlueUserDefinedFunction) State() (*glueUserDefinedFunctionState, bool) {
	return gudf.state, gudf.state != nil
}

func (gudf *GlueUserDefinedFunction) StateMust() *glueUserDefinedFunctionState {
	if gudf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gudf.Type(), gudf.LocalName()))
	}
	return gudf.state
}

func (gudf *GlueUserDefinedFunction) DependOn() terra.Reference {
	return terra.ReferenceResource(gudf)
}

type GlueUserDefinedFunctionArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// ClassName: string, required
	ClassName terra.StringValue `hcl:"class_name,attr" validate:"required"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OwnerName: string, required
	OwnerName terra.StringValue `hcl:"owner_name,attr" validate:"required"`
	// OwnerType: string, required
	OwnerType terra.StringValue `hcl:"owner_type,attr" validate:"required"`
	// ResourceUris: min=0,max=1000
	ResourceUris []glueuserdefinedfunction.ResourceUris `hcl:"resource_uris,block" validate:"min=0,max=1000"`
	// DependsOn contains resources that GlueUserDefinedFunction depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type glueUserDefinedFunctionAttributes struct {
	ref terra.Reference
}

func (gudf glueUserDefinedFunctionAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(gudf.ref.Append("arn"))
}

func (gudf glueUserDefinedFunctionAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceString(gudf.ref.Append("catalog_id"))
}

func (gudf glueUserDefinedFunctionAttributes) ClassName() terra.StringValue {
	return terra.ReferenceString(gudf.ref.Append("class_name"))
}

func (gudf glueUserDefinedFunctionAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceString(gudf.ref.Append("create_time"))
}

func (gudf glueUserDefinedFunctionAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceString(gudf.ref.Append("database_name"))
}

func (gudf glueUserDefinedFunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(gudf.ref.Append("id"))
}

func (gudf glueUserDefinedFunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(gudf.ref.Append("name"))
}

func (gudf glueUserDefinedFunctionAttributes) OwnerName() terra.StringValue {
	return terra.ReferenceString(gudf.ref.Append("owner_name"))
}

func (gudf glueUserDefinedFunctionAttributes) OwnerType() terra.StringValue {
	return terra.ReferenceString(gudf.ref.Append("owner_type"))
}

func (gudf glueUserDefinedFunctionAttributes) ResourceUris() terra.SetValue[glueuserdefinedfunction.ResourceUrisAttributes] {
	return terra.ReferenceSet[glueuserdefinedfunction.ResourceUrisAttributes](gudf.ref.Append("resource_uris"))
}

type glueUserDefinedFunctionState struct {
	Arn          string                                      `json:"arn"`
	CatalogId    string                                      `json:"catalog_id"`
	ClassName    string                                      `json:"class_name"`
	CreateTime   string                                      `json:"create_time"`
	DatabaseName string                                      `json:"database_name"`
	Id           string                                      `json:"id"`
	Name         string                                      `json:"name"`
	OwnerName    string                                      `json:"owner_name"`
	OwnerType    string                                      `json:"owner_type"`
	ResourceUris []glueuserdefinedfunction.ResourceUrisState `json:"resource_uris"`
}

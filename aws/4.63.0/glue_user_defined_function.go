// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	glueuserdefinedfunction "github.com/golingon/terraproviders/aws/4.63.0/glueuserdefinedfunction"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGlueUserDefinedFunction creates a new instance of [GlueUserDefinedFunction].
func NewGlueUserDefinedFunction(name string, args GlueUserDefinedFunctionArgs) *GlueUserDefinedFunction {
	return &GlueUserDefinedFunction{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueUserDefinedFunction)(nil)

// GlueUserDefinedFunction represents the Terraform resource aws_glue_user_defined_function.
type GlueUserDefinedFunction struct {
	Name      string
	Args      GlueUserDefinedFunctionArgs
	state     *glueUserDefinedFunctionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlueUserDefinedFunction].
func (gudf *GlueUserDefinedFunction) Type() string {
	return "aws_glue_user_defined_function"
}

// LocalName returns the local name for [GlueUserDefinedFunction].
func (gudf *GlueUserDefinedFunction) LocalName() string {
	return gudf.Name
}

// Configuration returns the configuration (args) for [GlueUserDefinedFunction].
func (gudf *GlueUserDefinedFunction) Configuration() interface{} {
	return gudf.Args
}

// DependOn is used for other resources to depend on [GlueUserDefinedFunction].
func (gudf *GlueUserDefinedFunction) DependOn() terra.Reference {
	return terra.ReferenceResource(gudf)
}

// Dependencies returns the list of resources [GlueUserDefinedFunction] depends_on.
func (gudf *GlueUserDefinedFunction) Dependencies() terra.Dependencies {
	return gudf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlueUserDefinedFunction].
func (gudf *GlueUserDefinedFunction) LifecycleManagement() *terra.Lifecycle {
	return gudf.Lifecycle
}

// Attributes returns the attributes for [GlueUserDefinedFunction].
func (gudf *GlueUserDefinedFunction) Attributes() glueUserDefinedFunctionAttributes {
	return glueUserDefinedFunctionAttributes{ref: terra.ReferenceResource(gudf)}
}

// ImportState imports the given attribute values into [GlueUserDefinedFunction]'s state.
func (gudf *GlueUserDefinedFunction) ImportState(av io.Reader) error {
	gudf.state = &glueUserDefinedFunctionState{}
	if err := json.NewDecoder(av).Decode(gudf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gudf.Type(), gudf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlueUserDefinedFunction] has state.
func (gudf *GlueUserDefinedFunction) State() (*glueUserDefinedFunctionState, bool) {
	return gudf.state, gudf.state != nil
}

// StateMust returns the state for [GlueUserDefinedFunction]. Panics if the state is nil.
func (gudf *GlueUserDefinedFunction) StateMust() *glueUserDefinedFunctionState {
	if gudf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gudf.Type(), gudf.LocalName()))
	}
	return gudf.state
}

// GlueUserDefinedFunctionArgs contains the configurations for aws_glue_user_defined_function.
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
}
type glueUserDefinedFunctionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_glue_user_defined_function.
func (gudf glueUserDefinedFunctionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gudf.ref.Append("arn"))
}

// CatalogId returns a reference to field catalog_id of aws_glue_user_defined_function.
func (gudf glueUserDefinedFunctionAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(gudf.ref.Append("catalog_id"))
}

// ClassName returns a reference to field class_name of aws_glue_user_defined_function.
func (gudf glueUserDefinedFunctionAttributes) ClassName() terra.StringValue {
	return terra.ReferenceAsString(gudf.ref.Append("class_name"))
}

// CreateTime returns a reference to field create_time of aws_glue_user_defined_function.
func (gudf glueUserDefinedFunctionAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gudf.ref.Append("create_time"))
}

// DatabaseName returns a reference to field database_name of aws_glue_user_defined_function.
func (gudf glueUserDefinedFunctionAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(gudf.ref.Append("database_name"))
}

// Id returns a reference to field id of aws_glue_user_defined_function.
func (gudf glueUserDefinedFunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gudf.ref.Append("id"))
}

// Name returns a reference to field name of aws_glue_user_defined_function.
func (gudf glueUserDefinedFunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gudf.ref.Append("name"))
}

// OwnerName returns a reference to field owner_name of aws_glue_user_defined_function.
func (gudf glueUserDefinedFunctionAttributes) OwnerName() terra.StringValue {
	return terra.ReferenceAsString(gudf.ref.Append("owner_name"))
}

// OwnerType returns a reference to field owner_type of aws_glue_user_defined_function.
func (gudf glueUserDefinedFunctionAttributes) OwnerType() terra.StringValue {
	return terra.ReferenceAsString(gudf.ref.Append("owner_type"))
}

func (gudf glueUserDefinedFunctionAttributes) ResourceUris() terra.SetValue[glueuserdefinedfunction.ResourceUrisAttributes] {
	return terra.ReferenceAsSet[glueuserdefinedfunction.ResourceUrisAttributes](gudf.ref.Append("resource_uris"))
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

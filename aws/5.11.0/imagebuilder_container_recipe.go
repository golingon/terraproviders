// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	imagebuildercontainerrecipe "github.com/golingon/terraproviders/aws/5.11.0/imagebuildercontainerrecipe"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewImagebuilderContainerRecipe creates a new instance of [ImagebuilderContainerRecipe].
func NewImagebuilderContainerRecipe(name string, args ImagebuilderContainerRecipeArgs) *ImagebuilderContainerRecipe {
	return &ImagebuilderContainerRecipe{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ImagebuilderContainerRecipe)(nil)

// ImagebuilderContainerRecipe represents the Terraform resource aws_imagebuilder_container_recipe.
type ImagebuilderContainerRecipe struct {
	Name      string
	Args      ImagebuilderContainerRecipeArgs
	state     *imagebuilderContainerRecipeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ImagebuilderContainerRecipe].
func (icr *ImagebuilderContainerRecipe) Type() string {
	return "aws_imagebuilder_container_recipe"
}

// LocalName returns the local name for [ImagebuilderContainerRecipe].
func (icr *ImagebuilderContainerRecipe) LocalName() string {
	return icr.Name
}

// Configuration returns the configuration (args) for [ImagebuilderContainerRecipe].
func (icr *ImagebuilderContainerRecipe) Configuration() interface{} {
	return icr.Args
}

// DependOn is used for other resources to depend on [ImagebuilderContainerRecipe].
func (icr *ImagebuilderContainerRecipe) DependOn() terra.Reference {
	return terra.ReferenceResource(icr)
}

// Dependencies returns the list of resources [ImagebuilderContainerRecipe] depends_on.
func (icr *ImagebuilderContainerRecipe) Dependencies() terra.Dependencies {
	return icr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ImagebuilderContainerRecipe].
func (icr *ImagebuilderContainerRecipe) LifecycleManagement() *terra.Lifecycle {
	return icr.Lifecycle
}

// Attributes returns the attributes for [ImagebuilderContainerRecipe].
func (icr *ImagebuilderContainerRecipe) Attributes() imagebuilderContainerRecipeAttributes {
	return imagebuilderContainerRecipeAttributes{ref: terra.ReferenceResource(icr)}
}

// ImportState imports the given attribute values into [ImagebuilderContainerRecipe]'s state.
func (icr *ImagebuilderContainerRecipe) ImportState(av io.Reader) error {
	icr.state = &imagebuilderContainerRecipeState{}
	if err := json.NewDecoder(av).Decode(icr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", icr.Type(), icr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ImagebuilderContainerRecipe] has state.
func (icr *ImagebuilderContainerRecipe) State() (*imagebuilderContainerRecipeState, bool) {
	return icr.state, icr.state != nil
}

// StateMust returns the state for [ImagebuilderContainerRecipe]. Panics if the state is nil.
func (icr *ImagebuilderContainerRecipe) StateMust() *imagebuilderContainerRecipeState {
	if icr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", icr.Type(), icr.LocalName()))
	}
	return icr.state
}

// ImagebuilderContainerRecipeArgs contains the configurations for aws_imagebuilder_container_recipe.
type ImagebuilderContainerRecipeArgs struct {
	// ContainerType: string, required
	ContainerType terra.StringValue `hcl:"container_type,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DockerfileTemplateData: string, optional
	DockerfileTemplateData terra.StringValue `hcl:"dockerfile_template_data,attr"`
	// DockerfileTemplateUri: string, optional
	DockerfileTemplateUri terra.StringValue `hcl:"dockerfile_template_uri,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentImage: string, required
	ParentImage terra.StringValue `hcl:"parent_image,attr" validate:"required"`
	// PlatformOverride: string, optional
	PlatformOverride terra.StringValue `hcl:"platform_override,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// WorkingDirectory: string, optional
	WorkingDirectory terra.StringValue `hcl:"working_directory,attr"`
	// Component: min=1
	Component []imagebuildercontainerrecipe.Component `hcl:"component,block" validate:"min=1"`
	// InstanceConfiguration: optional
	InstanceConfiguration *imagebuildercontainerrecipe.InstanceConfiguration `hcl:"instance_configuration,block"`
	// TargetRepository: required
	TargetRepository *imagebuildercontainerrecipe.TargetRepository `hcl:"target_repository,block" validate:"required"`
}
type imagebuilderContainerRecipeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("arn"))
}

// ContainerType returns a reference to field container_type of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) ContainerType() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("container_type"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("date_created"))
}

// Description returns a reference to field description of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("description"))
}

// DockerfileTemplateData returns a reference to field dockerfile_template_data of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) DockerfileTemplateData() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("dockerfile_template_data"))
}

// DockerfileTemplateUri returns a reference to field dockerfile_template_uri of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) DockerfileTemplateUri() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("dockerfile_template_uri"))
}

// Encrypted returns a reference to field encrypted of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(icr.ref.Append("encrypted"))
}

// Id returns a reference to field id of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("kms_key_id"))
}

// Name returns a reference to field name of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("owner"))
}

// ParentImage returns a reference to field parent_image of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) ParentImage() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("parent_image"))
}

// Platform returns a reference to field platform of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("platform"))
}

// PlatformOverride returns a reference to field platform_override of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) PlatformOverride() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("platform_override"))
}

// Tags returns a reference to field tags of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](icr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](icr.ref.Append("tags_all"))
}

// Version returns a reference to field version of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("version"))
}

// WorkingDirectory returns a reference to field working_directory of aws_imagebuilder_container_recipe.
func (icr imagebuilderContainerRecipeAttributes) WorkingDirectory() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("working_directory"))
}

func (icr imagebuilderContainerRecipeAttributes) Component() terra.ListValue[imagebuildercontainerrecipe.ComponentAttributes] {
	return terra.ReferenceAsList[imagebuildercontainerrecipe.ComponentAttributes](icr.ref.Append("component"))
}

func (icr imagebuilderContainerRecipeAttributes) InstanceConfiguration() terra.ListValue[imagebuildercontainerrecipe.InstanceConfigurationAttributes] {
	return terra.ReferenceAsList[imagebuildercontainerrecipe.InstanceConfigurationAttributes](icr.ref.Append("instance_configuration"))
}

func (icr imagebuilderContainerRecipeAttributes) TargetRepository() terra.ListValue[imagebuildercontainerrecipe.TargetRepositoryAttributes] {
	return terra.ReferenceAsList[imagebuildercontainerrecipe.TargetRepositoryAttributes](icr.ref.Append("target_repository"))
}

type imagebuilderContainerRecipeState struct {
	Arn                    string                                                   `json:"arn"`
	ContainerType          string                                                   `json:"container_type"`
	DateCreated            string                                                   `json:"date_created"`
	Description            string                                                   `json:"description"`
	DockerfileTemplateData string                                                   `json:"dockerfile_template_data"`
	DockerfileTemplateUri  string                                                   `json:"dockerfile_template_uri"`
	Encrypted              bool                                                     `json:"encrypted"`
	Id                     string                                                   `json:"id"`
	KmsKeyId               string                                                   `json:"kms_key_id"`
	Name                   string                                                   `json:"name"`
	Owner                  string                                                   `json:"owner"`
	ParentImage            string                                                   `json:"parent_image"`
	Platform               string                                                   `json:"platform"`
	PlatformOverride       string                                                   `json:"platform_override"`
	Tags                   map[string]string                                        `json:"tags"`
	TagsAll                map[string]string                                        `json:"tags_all"`
	Version                string                                                   `json:"version"`
	WorkingDirectory       string                                                   `json:"working_directory"`
	Component              []imagebuildercontainerrecipe.ComponentState             `json:"component"`
	InstanceConfiguration  []imagebuildercontainerrecipe.InstanceConfigurationState `json:"instance_configuration"`
	TargetRepository       []imagebuildercontainerrecipe.TargetRepositoryState      `json:"target_repository"`
}

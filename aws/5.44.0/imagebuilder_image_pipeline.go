// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	imagebuilderimagepipeline "github.com/golingon/terraproviders/aws/5.44.0/imagebuilderimagepipeline"
	"io"
)

// NewImagebuilderImagePipeline creates a new instance of [ImagebuilderImagePipeline].
func NewImagebuilderImagePipeline(name string, args ImagebuilderImagePipelineArgs) *ImagebuilderImagePipeline {
	return &ImagebuilderImagePipeline{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ImagebuilderImagePipeline)(nil)

// ImagebuilderImagePipeline represents the Terraform resource aws_imagebuilder_image_pipeline.
type ImagebuilderImagePipeline struct {
	Name      string
	Args      ImagebuilderImagePipelineArgs
	state     *imagebuilderImagePipelineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ImagebuilderImagePipeline].
func (iip *ImagebuilderImagePipeline) Type() string {
	return "aws_imagebuilder_image_pipeline"
}

// LocalName returns the local name for [ImagebuilderImagePipeline].
func (iip *ImagebuilderImagePipeline) LocalName() string {
	return iip.Name
}

// Configuration returns the configuration (args) for [ImagebuilderImagePipeline].
func (iip *ImagebuilderImagePipeline) Configuration() interface{} {
	return iip.Args
}

// DependOn is used for other resources to depend on [ImagebuilderImagePipeline].
func (iip *ImagebuilderImagePipeline) DependOn() terra.Reference {
	return terra.ReferenceResource(iip)
}

// Dependencies returns the list of resources [ImagebuilderImagePipeline] depends_on.
func (iip *ImagebuilderImagePipeline) Dependencies() terra.Dependencies {
	return iip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ImagebuilderImagePipeline].
func (iip *ImagebuilderImagePipeline) LifecycleManagement() *terra.Lifecycle {
	return iip.Lifecycle
}

// Attributes returns the attributes for [ImagebuilderImagePipeline].
func (iip *ImagebuilderImagePipeline) Attributes() imagebuilderImagePipelineAttributes {
	return imagebuilderImagePipelineAttributes{ref: terra.ReferenceResource(iip)}
}

// ImportState imports the given attribute values into [ImagebuilderImagePipeline]'s state.
func (iip *ImagebuilderImagePipeline) ImportState(av io.Reader) error {
	iip.state = &imagebuilderImagePipelineState{}
	if err := json.NewDecoder(av).Decode(iip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iip.Type(), iip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ImagebuilderImagePipeline] has state.
func (iip *ImagebuilderImagePipeline) State() (*imagebuilderImagePipelineState, bool) {
	return iip.state, iip.state != nil
}

// StateMust returns the state for [ImagebuilderImagePipeline]. Panics if the state is nil.
func (iip *ImagebuilderImagePipeline) StateMust() *imagebuilderImagePipelineState {
	if iip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iip.Type(), iip.LocalName()))
	}
	return iip.state
}

// ImagebuilderImagePipelineArgs contains the configurations for aws_imagebuilder_image_pipeline.
type ImagebuilderImagePipelineArgs struct {
	// ContainerRecipeArn: string, optional
	ContainerRecipeArn terra.StringValue `hcl:"container_recipe_arn,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DistributionConfigurationArn: string, optional
	DistributionConfigurationArn terra.StringValue `hcl:"distribution_configuration_arn,attr"`
	// EnhancedImageMetadataEnabled: bool, optional
	EnhancedImageMetadataEnabled terra.BoolValue `hcl:"enhanced_image_metadata_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageRecipeArn: string, optional
	ImageRecipeArn terra.StringValue `hcl:"image_recipe_arn,attr"`
	// InfrastructureConfigurationArn: string, required
	InfrastructureConfigurationArn terra.StringValue `hcl:"infrastructure_configuration_arn,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ImageScanningConfiguration: optional
	ImageScanningConfiguration *imagebuilderimagepipeline.ImageScanningConfiguration `hcl:"image_scanning_configuration,block"`
	// ImageTestsConfiguration: optional
	ImageTestsConfiguration *imagebuilderimagepipeline.ImageTestsConfiguration `hcl:"image_tests_configuration,block"`
	// Schedule: optional
	Schedule *imagebuilderimagepipeline.Schedule `hcl:"schedule,block"`
}
type imagebuilderImagePipelineAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("arn"))
}

// ContainerRecipeArn returns a reference to field container_recipe_arn of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) ContainerRecipeArn() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("container_recipe_arn"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("date_created"))
}

// DateLastRun returns a reference to field date_last_run of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) DateLastRun() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("date_last_run"))
}

// DateNextRun returns a reference to field date_next_run of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) DateNextRun() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("date_next_run"))
}

// DateUpdated returns a reference to field date_updated of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) DateUpdated() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("date_updated"))
}

// Description returns a reference to field description of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("description"))
}

// DistributionConfigurationArn returns a reference to field distribution_configuration_arn of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) DistributionConfigurationArn() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("distribution_configuration_arn"))
}

// EnhancedImageMetadataEnabled returns a reference to field enhanced_image_metadata_enabled of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) EnhancedImageMetadataEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(iip.ref.Append("enhanced_image_metadata_enabled"))
}

// Id returns a reference to field id of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("id"))
}

// ImageRecipeArn returns a reference to field image_recipe_arn of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) ImageRecipeArn() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("image_recipe_arn"))
}

// InfrastructureConfigurationArn returns a reference to field infrastructure_configuration_arn of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) InfrastructureConfigurationArn() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("infrastructure_configuration_arn"))
}

// Name returns a reference to field name of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("name"))
}

// Platform returns a reference to field platform of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("platform"))
}

// Status returns a reference to field status of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iip.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_imagebuilder_image_pipeline.
func (iip imagebuilderImagePipelineAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iip.ref.Append("tags_all"))
}

func (iip imagebuilderImagePipelineAttributes) ImageScanningConfiguration() terra.ListValue[imagebuilderimagepipeline.ImageScanningConfigurationAttributes] {
	return terra.ReferenceAsList[imagebuilderimagepipeline.ImageScanningConfigurationAttributes](iip.ref.Append("image_scanning_configuration"))
}

func (iip imagebuilderImagePipelineAttributes) ImageTestsConfiguration() terra.ListValue[imagebuilderimagepipeline.ImageTestsConfigurationAttributes] {
	return terra.ReferenceAsList[imagebuilderimagepipeline.ImageTestsConfigurationAttributes](iip.ref.Append("image_tests_configuration"))
}

func (iip imagebuilderImagePipelineAttributes) Schedule() terra.ListValue[imagebuilderimagepipeline.ScheduleAttributes] {
	return terra.ReferenceAsList[imagebuilderimagepipeline.ScheduleAttributes](iip.ref.Append("schedule"))
}

type imagebuilderImagePipelineState struct {
	Arn                            string                                                      `json:"arn"`
	ContainerRecipeArn             string                                                      `json:"container_recipe_arn"`
	DateCreated                    string                                                      `json:"date_created"`
	DateLastRun                    string                                                      `json:"date_last_run"`
	DateNextRun                    string                                                      `json:"date_next_run"`
	DateUpdated                    string                                                      `json:"date_updated"`
	Description                    string                                                      `json:"description"`
	DistributionConfigurationArn   string                                                      `json:"distribution_configuration_arn"`
	EnhancedImageMetadataEnabled   bool                                                        `json:"enhanced_image_metadata_enabled"`
	Id                             string                                                      `json:"id"`
	ImageRecipeArn                 string                                                      `json:"image_recipe_arn"`
	InfrastructureConfigurationArn string                                                      `json:"infrastructure_configuration_arn"`
	Name                           string                                                      `json:"name"`
	Platform                       string                                                      `json:"platform"`
	Status                         string                                                      `json:"status"`
	Tags                           map[string]string                                           `json:"tags"`
	TagsAll                        map[string]string                                           `json:"tags_all"`
	ImageScanningConfiguration     []imagebuilderimagepipeline.ImageScanningConfigurationState `json:"image_scanning_configuration"`
	ImageTestsConfiguration        []imagebuilderimagepipeline.ImageTestsConfigurationState    `json:"image_tests_configuration"`
	Schedule                       []imagebuilderimagepipeline.ScheduleState                   `json:"schedule"`
}

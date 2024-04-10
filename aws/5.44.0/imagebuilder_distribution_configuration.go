// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	imagebuilderdistributionconfiguration "github.com/golingon/terraproviders/aws/5.44.0/imagebuilderdistributionconfiguration"
	"io"
)

// NewImagebuilderDistributionConfiguration creates a new instance of [ImagebuilderDistributionConfiguration].
func NewImagebuilderDistributionConfiguration(name string, args ImagebuilderDistributionConfigurationArgs) *ImagebuilderDistributionConfiguration {
	return &ImagebuilderDistributionConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ImagebuilderDistributionConfiguration)(nil)

// ImagebuilderDistributionConfiguration represents the Terraform resource aws_imagebuilder_distribution_configuration.
type ImagebuilderDistributionConfiguration struct {
	Name      string
	Args      ImagebuilderDistributionConfigurationArgs
	state     *imagebuilderDistributionConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ImagebuilderDistributionConfiguration].
func (idc *ImagebuilderDistributionConfiguration) Type() string {
	return "aws_imagebuilder_distribution_configuration"
}

// LocalName returns the local name for [ImagebuilderDistributionConfiguration].
func (idc *ImagebuilderDistributionConfiguration) LocalName() string {
	return idc.Name
}

// Configuration returns the configuration (args) for [ImagebuilderDistributionConfiguration].
func (idc *ImagebuilderDistributionConfiguration) Configuration() interface{} {
	return idc.Args
}

// DependOn is used for other resources to depend on [ImagebuilderDistributionConfiguration].
func (idc *ImagebuilderDistributionConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(idc)
}

// Dependencies returns the list of resources [ImagebuilderDistributionConfiguration] depends_on.
func (idc *ImagebuilderDistributionConfiguration) Dependencies() terra.Dependencies {
	return idc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ImagebuilderDistributionConfiguration].
func (idc *ImagebuilderDistributionConfiguration) LifecycleManagement() *terra.Lifecycle {
	return idc.Lifecycle
}

// Attributes returns the attributes for [ImagebuilderDistributionConfiguration].
func (idc *ImagebuilderDistributionConfiguration) Attributes() imagebuilderDistributionConfigurationAttributes {
	return imagebuilderDistributionConfigurationAttributes{ref: terra.ReferenceResource(idc)}
}

// ImportState imports the given attribute values into [ImagebuilderDistributionConfiguration]'s state.
func (idc *ImagebuilderDistributionConfiguration) ImportState(av io.Reader) error {
	idc.state = &imagebuilderDistributionConfigurationState{}
	if err := json.NewDecoder(av).Decode(idc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", idc.Type(), idc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ImagebuilderDistributionConfiguration] has state.
func (idc *ImagebuilderDistributionConfiguration) State() (*imagebuilderDistributionConfigurationState, bool) {
	return idc.state, idc.state != nil
}

// StateMust returns the state for [ImagebuilderDistributionConfiguration]. Panics if the state is nil.
func (idc *ImagebuilderDistributionConfiguration) StateMust() *imagebuilderDistributionConfigurationState {
	if idc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", idc.Type(), idc.LocalName()))
	}
	return idc.state
}

// ImagebuilderDistributionConfigurationArgs contains the configurations for aws_imagebuilder_distribution_configuration.
type ImagebuilderDistributionConfigurationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Distribution: min=1
	Distribution []imagebuilderdistributionconfiguration.Distribution `hcl:"distribution,block" validate:"min=1"`
}
type imagebuilderDistributionConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_distribution_configuration.
func (idc imagebuilderDistributionConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("arn"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_distribution_configuration.
func (idc imagebuilderDistributionConfigurationAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("date_created"))
}

// DateUpdated returns a reference to field date_updated of aws_imagebuilder_distribution_configuration.
func (idc imagebuilderDistributionConfigurationAttributes) DateUpdated() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("date_updated"))
}

// Description returns a reference to field description of aws_imagebuilder_distribution_configuration.
func (idc imagebuilderDistributionConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("description"))
}

// Id returns a reference to field id of aws_imagebuilder_distribution_configuration.
func (idc imagebuilderDistributionConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("id"))
}

// Name returns a reference to field name of aws_imagebuilder_distribution_configuration.
func (idc imagebuilderDistributionConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_imagebuilder_distribution_configuration.
func (idc imagebuilderDistributionConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](idc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_imagebuilder_distribution_configuration.
func (idc imagebuilderDistributionConfigurationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](idc.ref.Append("tags_all"))
}

func (idc imagebuilderDistributionConfigurationAttributes) Distribution() terra.SetValue[imagebuilderdistributionconfiguration.DistributionAttributes] {
	return terra.ReferenceAsSet[imagebuilderdistributionconfiguration.DistributionAttributes](idc.ref.Append("distribution"))
}

type imagebuilderDistributionConfigurationState struct {
	Arn          string                                                    `json:"arn"`
	DateCreated  string                                                    `json:"date_created"`
	DateUpdated  string                                                    `json:"date_updated"`
	Description  string                                                    `json:"description"`
	Id           string                                                    `json:"id"`
	Name         string                                                    `json:"name"`
	Tags         map[string]string                                         `json:"tags"`
	TagsAll      map[string]string                                         `json:"tags_all"`
	Distribution []imagebuilderdistributionconfiguration.DistributionState `json:"distribution"`
}

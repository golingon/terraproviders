// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewImagebuilderWorkflow creates a new instance of [ImagebuilderWorkflow].
func NewImagebuilderWorkflow(name string, args ImagebuilderWorkflowArgs) *ImagebuilderWorkflow {
	return &ImagebuilderWorkflow{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ImagebuilderWorkflow)(nil)

// ImagebuilderWorkflow represents the Terraform resource aws_imagebuilder_workflow.
type ImagebuilderWorkflow struct {
	Name      string
	Args      ImagebuilderWorkflowArgs
	state     *imagebuilderWorkflowState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ImagebuilderWorkflow].
func (iw *ImagebuilderWorkflow) Type() string {
	return "aws_imagebuilder_workflow"
}

// LocalName returns the local name for [ImagebuilderWorkflow].
func (iw *ImagebuilderWorkflow) LocalName() string {
	return iw.Name
}

// Configuration returns the configuration (args) for [ImagebuilderWorkflow].
func (iw *ImagebuilderWorkflow) Configuration() interface{} {
	return iw.Args
}

// DependOn is used for other resources to depend on [ImagebuilderWorkflow].
func (iw *ImagebuilderWorkflow) DependOn() terra.Reference {
	return terra.ReferenceResource(iw)
}

// Dependencies returns the list of resources [ImagebuilderWorkflow] depends_on.
func (iw *ImagebuilderWorkflow) Dependencies() terra.Dependencies {
	return iw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ImagebuilderWorkflow].
func (iw *ImagebuilderWorkflow) LifecycleManagement() *terra.Lifecycle {
	return iw.Lifecycle
}

// Attributes returns the attributes for [ImagebuilderWorkflow].
func (iw *ImagebuilderWorkflow) Attributes() imagebuilderWorkflowAttributes {
	return imagebuilderWorkflowAttributes{ref: terra.ReferenceResource(iw)}
}

// ImportState imports the given attribute values into [ImagebuilderWorkflow]'s state.
func (iw *ImagebuilderWorkflow) ImportState(av io.Reader) error {
	iw.state = &imagebuilderWorkflowState{}
	if err := json.NewDecoder(av).Decode(iw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iw.Type(), iw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ImagebuilderWorkflow] has state.
func (iw *ImagebuilderWorkflow) State() (*imagebuilderWorkflowState, bool) {
	return iw.state, iw.state != nil
}

// StateMust returns the state for [ImagebuilderWorkflow]. Panics if the state is nil.
func (iw *ImagebuilderWorkflow) StateMust() *imagebuilderWorkflowState {
	if iw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iw.Type(), iw.LocalName()))
	}
	return iw.state
}

// ImagebuilderWorkflowArgs contains the configurations for aws_imagebuilder_workflow.
type ImagebuilderWorkflowArgs struct {
	// ChangeDescription: string, optional
	ChangeDescription terra.StringValue `hcl:"change_description,attr"`
	// Data: string, optional
	Data terra.StringValue `hcl:"data,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Uri: string, optional
	Uri terra.StringValue `hcl:"uri,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
}
type imagebuilderWorkflowAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("arn"))
}

// ChangeDescription returns a reference to field change_description of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) ChangeDescription() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("change_description"))
}

// Data returns a reference to field data of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("data"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("date_created"))
}

// Description returns a reference to field description of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("description"))
}

// Id returns a reference to field id of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("kms_key_id"))
}

// Name returns a reference to field name of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("owner"))
}

// Tags returns a reference to field tags of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iw.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iw.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("type"))
}

// Uri returns a reference to field uri of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("uri"))
}

// Version returns a reference to field version of aws_imagebuilder_workflow.
func (iw imagebuilderWorkflowAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(iw.ref.Append("version"))
}

type imagebuilderWorkflowState struct {
	Arn               string            `json:"arn"`
	ChangeDescription string            `json:"change_description"`
	Data              string            `json:"data"`
	DateCreated       string            `json:"date_created"`
	Description       string            `json:"description"`
	Id                string            `json:"id"`
	KmsKeyId          string            `json:"kms_key_id"`
	Name              string            `json:"name"`
	Owner             string            `json:"owner"`
	Tags              map[string]string `json:"tags"`
	TagsAll           map[string]string `json:"tags_all"`
	Type              string            `json:"type"`
	Uri               string            `json:"uri"`
	Version           string            `json:"version"`
}

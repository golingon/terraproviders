// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elastictranscoderpipeline "github.com/golingon/terraproviders/aws/5.7.0/elastictranscoderpipeline"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElastictranscoderPipeline creates a new instance of [ElastictranscoderPipeline].
func NewElastictranscoderPipeline(name string, args ElastictranscoderPipelineArgs) *ElastictranscoderPipeline {
	return &ElastictranscoderPipeline{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElastictranscoderPipeline)(nil)

// ElastictranscoderPipeline represents the Terraform resource aws_elastictranscoder_pipeline.
type ElastictranscoderPipeline struct {
	Name      string
	Args      ElastictranscoderPipelineArgs
	state     *elastictranscoderPipelineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElastictranscoderPipeline].
func (ep *ElastictranscoderPipeline) Type() string {
	return "aws_elastictranscoder_pipeline"
}

// LocalName returns the local name for [ElastictranscoderPipeline].
func (ep *ElastictranscoderPipeline) LocalName() string {
	return ep.Name
}

// Configuration returns the configuration (args) for [ElastictranscoderPipeline].
func (ep *ElastictranscoderPipeline) Configuration() interface{} {
	return ep.Args
}

// DependOn is used for other resources to depend on [ElastictranscoderPipeline].
func (ep *ElastictranscoderPipeline) DependOn() terra.Reference {
	return terra.ReferenceResource(ep)
}

// Dependencies returns the list of resources [ElastictranscoderPipeline] depends_on.
func (ep *ElastictranscoderPipeline) Dependencies() terra.Dependencies {
	return ep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElastictranscoderPipeline].
func (ep *ElastictranscoderPipeline) LifecycleManagement() *terra.Lifecycle {
	return ep.Lifecycle
}

// Attributes returns the attributes for [ElastictranscoderPipeline].
func (ep *ElastictranscoderPipeline) Attributes() elastictranscoderPipelineAttributes {
	return elastictranscoderPipelineAttributes{ref: terra.ReferenceResource(ep)}
}

// ImportState imports the given attribute values into [ElastictranscoderPipeline]'s state.
func (ep *ElastictranscoderPipeline) ImportState(av io.Reader) error {
	ep.state = &elastictranscoderPipelineState{}
	if err := json.NewDecoder(av).Decode(ep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ep.Type(), ep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElastictranscoderPipeline] has state.
func (ep *ElastictranscoderPipeline) State() (*elastictranscoderPipelineState, bool) {
	return ep.state, ep.state != nil
}

// StateMust returns the state for [ElastictranscoderPipeline]. Panics if the state is nil.
func (ep *ElastictranscoderPipeline) StateMust() *elastictranscoderPipelineState {
	if ep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ep.Type(), ep.LocalName()))
	}
	return ep.state
}

// ElastictranscoderPipelineArgs contains the configurations for aws_elastictranscoder_pipeline.
type ElastictranscoderPipelineArgs struct {
	// AwsKmsKeyArn: string, optional
	AwsKmsKeyArn terra.StringValue `hcl:"aws_kms_key_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InputBucket: string, required
	InputBucket terra.StringValue `hcl:"input_bucket,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// OutputBucket: string, optional
	OutputBucket terra.StringValue `hcl:"output_bucket,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ContentConfig: optional
	ContentConfig *elastictranscoderpipeline.ContentConfig `hcl:"content_config,block"`
	// ContentConfigPermissions: min=0
	ContentConfigPermissions []elastictranscoderpipeline.ContentConfigPermissions `hcl:"content_config_permissions,block" validate:"min=0"`
	// Notifications: optional
	Notifications *elastictranscoderpipeline.Notifications `hcl:"notifications,block"`
	// ThumbnailConfig: optional
	ThumbnailConfig *elastictranscoderpipeline.ThumbnailConfig `hcl:"thumbnail_config,block"`
	// ThumbnailConfigPermissions: min=0
	ThumbnailConfigPermissions []elastictranscoderpipeline.ThumbnailConfigPermissions `hcl:"thumbnail_config_permissions,block" validate:"min=0"`
}
type elastictranscoderPipelineAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_elastictranscoder_pipeline.
func (ep elastictranscoderPipelineAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("arn"))
}

// AwsKmsKeyArn returns a reference to field aws_kms_key_arn of aws_elastictranscoder_pipeline.
func (ep elastictranscoderPipelineAttributes) AwsKmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("aws_kms_key_arn"))
}

// Id returns a reference to field id of aws_elastictranscoder_pipeline.
func (ep elastictranscoderPipelineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("id"))
}

// InputBucket returns a reference to field input_bucket of aws_elastictranscoder_pipeline.
func (ep elastictranscoderPipelineAttributes) InputBucket() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("input_bucket"))
}

// Name returns a reference to field name of aws_elastictranscoder_pipeline.
func (ep elastictranscoderPipelineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("name"))
}

// OutputBucket returns a reference to field output_bucket of aws_elastictranscoder_pipeline.
func (ep elastictranscoderPipelineAttributes) OutputBucket() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("output_bucket"))
}

// Role returns a reference to field role of aws_elastictranscoder_pipeline.
func (ep elastictranscoderPipelineAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("role"))
}

func (ep elastictranscoderPipelineAttributes) ContentConfig() terra.ListValue[elastictranscoderpipeline.ContentConfigAttributes] {
	return terra.ReferenceAsList[elastictranscoderpipeline.ContentConfigAttributes](ep.ref.Append("content_config"))
}

func (ep elastictranscoderPipelineAttributes) ContentConfigPermissions() terra.SetValue[elastictranscoderpipeline.ContentConfigPermissionsAttributes] {
	return terra.ReferenceAsSet[elastictranscoderpipeline.ContentConfigPermissionsAttributes](ep.ref.Append("content_config_permissions"))
}

func (ep elastictranscoderPipelineAttributes) Notifications() terra.ListValue[elastictranscoderpipeline.NotificationsAttributes] {
	return terra.ReferenceAsList[elastictranscoderpipeline.NotificationsAttributes](ep.ref.Append("notifications"))
}

func (ep elastictranscoderPipelineAttributes) ThumbnailConfig() terra.ListValue[elastictranscoderpipeline.ThumbnailConfigAttributes] {
	return terra.ReferenceAsList[elastictranscoderpipeline.ThumbnailConfigAttributes](ep.ref.Append("thumbnail_config"))
}

func (ep elastictranscoderPipelineAttributes) ThumbnailConfigPermissions() terra.SetValue[elastictranscoderpipeline.ThumbnailConfigPermissionsAttributes] {
	return terra.ReferenceAsSet[elastictranscoderpipeline.ThumbnailConfigPermissionsAttributes](ep.ref.Append("thumbnail_config_permissions"))
}

type elastictranscoderPipelineState struct {
	Arn                        string                                                      `json:"arn"`
	AwsKmsKeyArn               string                                                      `json:"aws_kms_key_arn"`
	Id                         string                                                      `json:"id"`
	InputBucket                string                                                      `json:"input_bucket"`
	Name                       string                                                      `json:"name"`
	OutputBucket               string                                                      `json:"output_bucket"`
	Role                       string                                                      `json:"role"`
	ContentConfig              []elastictranscoderpipeline.ContentConfigState              `json:"content_config"`
	ContentConfigPermissions   []elastictranscoderpipeline.ContentConfigPermissionsState   `json:"content_config_permissions"`
	Notifications              []elastictranscoderpipeline.NotificationsState              `json:"notifications"`
	ThumbnailConfig            []elastictranscoderpipeline.ThumbnailConfigState            `json:"thumbnail_config"`
	ThumbnailConfigPermissions []elastictranscoderpipeline.ThumbnailConfigPermissionsState `json:"thumbnail_config_permissions"`
}
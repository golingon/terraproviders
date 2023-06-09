// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codepipelinewebhook "github.com/golingon/terraproviders/aws/5.0.1/codepipelinewebhook"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodepipelineWebhook creates a new instance of [CodepipelineWebhook].
func NewCodepipelineWebhook(name string, args CodepipelineWebhookArgs) *CodepipelineWebhook {
	return &CodepipelineWebhook{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodepipelineWebhook)(nil)

// CodepipelineWebhook represents the Terraform resource aws_codepipeline_webhook.
type CodepipelineWebhook struct {
	Name      string
	Args      CodepipelineWebhookArgs
	state     *codepipelineWebhookState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodepipelineWebhook].
func (cw *CodepipelineWebhook) Type() string {
	return "aws_codepipeline_webhook"
}

// LocalName returns the local name for [CodepipelineWebhook].
func (cw *CodepipelineWebhook) LocalName() string {
	return cw.Name
}

// Configuration returns the configuration (args) for [CodepipelineWebhook].
func (cw *CodepipelineWebhook) Configuration() interface{} {
	return cw.Args
}

// DependOn is used for other resources to depend on [CodepipelineWebhook].
func (cw *CodepipelineWebhook) DependOn() terra.Reference {
	return terra.ReferenceResource(cw)
}

// Dependencies returns the list of resources [CodepipelineWebhook] depends_on.
func (cw *CodepipelineWebhook) Dependencies() terra.Dependencies {
	return cw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodepipelineWebhook].
func (cw *CodepipelineWebhook) LifecycleManagement() *terra.Lifecycle {
	return cw.Lifecycle
}

// Attributes returns the attributes for [CodepipelineWebhook].
func (cw *CodepipelineWebhook) Attributes() codepipelineWebhookAttributes {
	return codepipelineWebhookAttributes{ref: terra.ReferenceResource(cw)}
}

// ImportState imports the given attribute values into [CodepipelineWebhook]'s state.
func (cw *CodepipelineWebhook) ImportState(av io.Reader) error {
	cw.state = &codepipelineWebhookState{}
	if err := json.NewDecoder(av).Decode(cw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cw.Type(), cw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodepipelineWebhook] has state.
func (cw *CodepipelineWebhook) State() (*codepipelineWebhookState, bool) {
	return cw.state, cw.state != nil
}

// StateMust returns the state for [CodepipelineWebhook]. Panics if the state is nil.
func (cw *CodepipelineWebhook) StateMust() *codepipelineWebhookState {
	if cw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cw.Type(), cw.LocalName()))
	}
	return cw.state
}

// CodepipelineWebhookArgs contains the configurations for aws_codepipeline_webhook.
type CodepipelineWebhookArgs struct {
	// Authentication: string, required
	Authentication terra.StringValue `hcl:"authentication,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetAction: string, required
	TargetAction terra.StringValue `hcl:"target_action,attr" validate:"required"`
	// TargetPipeline: string, required
	TargetPipeline terra.StringValue `hcl:"target_pipeline,attr" validate:"required"`
	// AuthenticationConfiguration: optional
	AuthenticationConfiguration *codepipelinewebhook.AuthenticationConfiguration `hcl:"authentication_configuration,block"`
	// Filter: min=1,max=5
	Filter []codepipelinewebhook.Filter `hcl:"filter,block" validate:"min=1,max=5"`
}
type codepipelineWebhookAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codepipeline_webhook.
func (cw codepipelineWebhookAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("arn"))
}

// Authentication returns a reference to field authentication of aws_codepipeline_webhook.
func (cw codepipelineWebhookAttributes) Authentication() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("authentication"))
}

// Id returns a reference to field id of aws_codepipeline_webhook.
func (cw codepipelineWebhookAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("id"))
}

// Name returns a reference to field name of aws_codepipeline_webhook.
func (cw codepipelineWebhookAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_codepipeline_webhook.
func (cw codepipelineWebhookAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cw.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codepipeline_webhook.
func (cw codepipelineWebhookAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cw.ref.Append("tags_all"))
}

// TargetAction returns a reference to field target_action of aws_codepipeline_webhook.
func (cw codepipelineWebhookAttributes) TargetAction() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("target_action"))
}

// TargetPipeline returns a reference to field target_pipeline of aws_codepipeline_webhook.
func (cw codepipelineWebhookAttributes) TargetPipeline() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("target_pipeline"))
}

// Url returns a reference to field url of aws_codepipeline_webhook.
func (cw codepipelineWebhookAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("url"))
}

func (cw codepipelineWebhookAttributes) AuthenticationConfiguration() terra.ListValue[codepipelinewebhook.AuthenticationConfigurationAttributes] {
	return terra.ReferenceAsList[codepipelinewebhook.AuthenticationConfigurationAttributes](cw.ref.Append("authentication_configuration"))
}

func (cw codepipelineWebhookAttributes) Filter() terra.SetValue[codepipelinewebhook.FilterAttributes] {
	return terra.ReferenceAsSet[codepipelinewebhook.FilterAttributes](cw.ref.Append("filter"))
}

type codepipelineWebhookState struct {
	Arn                         string                                                 `json:"arn"`
	Authentication              string                                                 `json:"authentication"`
	Id                          string                                                 `json:"id"`
	Name                        string                                                 `json:"name"`
	Tags                        map[string]string                                      `json:"tags"`
	TagsAll                     map[string]string                                      `json:"tags_all"`
	TargetAction                string                                                 `json:"target_action"`
	TargetPipeline              string                                                 `json:"target_pipeline"`
	Url                         string                                                 `json:"url"`
	AuthenticationConfiguration []codepipelinewebhook.AuthenticationConfigurationState `json:"authentication_configuration"`
	Filter                      []codepipelinewebhook.FilterState                      `json:"filter"`
}

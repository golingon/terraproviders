// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codebuildwebhook "github.com/golingon/terraproviders/aws/5.8.0/codebuildwebhook"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodebuildWebhook creates a new instance of [CodebuildWebhook].
func NewCodebuildWebhook(name string, args CodebuildWebhookArgs) *CodebuildWebhook {
	return &CodebuildWebhook{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodebuildWebhook)(nil)

// CodebuildWebhook represents the Terraform resource aws_codebuild_webhook.
type CodebuildWebhook struct {
	Name      string
	Args      CodebuildWebhookArgs
	state     *codebuildWebhookState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodebuildWebhook].
func (cw *CodebuildWebhook) Type() string {
	return "aws_codebuild_webhook"
}

// LocalName returns the local name for [CodebuildWebhook].
func (cw *CodebuildWebhook) LocalName() string {
	return cw.Name
}

// Configuration returns the configuration (args) for [CodebuildWebhook].
func (cw *CodebuildWebhook) Configuration() interface{} {
	return cw.Args
}

// DependOn is used for other resources to depend on [CodebuildWebhook].
func (cw *CodebuildWebhook) DependOn() terra.Reference {
	return terra.ReferenceResource(cw)
}

// Dependencies returns the list of resources [CodebuildWebhook] depends_on.
func (cw *CodebuildWebhook) Dependencies() terra.Dependencies {
	return cw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodebuildWebhook].
func (cw *CodebuildWebhook) LifecycleManagement() *terra.Lifecycle {
	return cw.Lifecycle
}

// Attributes returns the attributes for [CodebuildWebhook].
func (cw *CodebuildWebhook) Attributes() codebuildWebhookAttributes {
	return codebuildWebhookAttributes{ref: terra.ReferenceResource(cw)}
}

// ImportState imports the given attribute values into [CodebuildWebhook]'s state.
func (cw *CodebuildWebhook) ImportState(av io.Reader) error {
	cw.state = &codebuildWebhookState{}
	if err := json.NewDecoder(av).Decode(cw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cw.Type(), cw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodebuildWebhook] has state.
func (cw *CodebuildWebhook) State() (*codebuildWebhookState, bool) {
	return cw.state, cw.state != nil
}

// StateMust returns the state for [CodebuildWebhook]. Panics if the state is nil.
func (cw *CodebuildWebhook) StateMust() *codebuildWebhookState {
	if cw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cw.Type(), cw.LocalName()))
	}
	return cw.state
}

// CodebuildWebhookArgs contains the configurations for aws_codebuild_webhook.
type CodebuildWebhookArgs struct {
	// BranchFilter: string, optional
	BranchFilter terra.StringValue `hcl:"branch_filter,attr"`
	// BuildType: string, optional
	BuildType terra.StringValue `hcl:"build_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProjectName: string, required
	ProjectName terra.StringValue `hcl:"project_name,attr" validate:"required"`
	// FilterGroup: min=0
	FilterGroup []codebuildwebhook.FilterGroup `hcl:"filter_group,block" validate:"min=0"`
}
type codebuildWebhookAttributes struct {
	ref terra.Reference
}

// BranchFilter returns a reference to field branch_filter of aws_codebuild_webhook.
func (cw codebuildWebhookAttributes) BranchFilter() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("branch_filter"))
}

// BuildType returns a reference to field build_type of aws_codebuild_webhook.
func (cw codebuildWebhookAttributes) BuildType() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("build_type"))
}

// Id returns a reference to field id of aws_codebuild_webhook.
func (cw codebuildWebhookAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("id"))
}

// PayloadUrl returns a reference to field payload_url of aws_codebuild_webhook.
func (cw codebuildWebhookAttributes) PayloadUrl() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("payload_url"))
}

// ProjectName returns a reference to field project_name of aws_codebuild_webhook.
func (cw codebuildWebhookAttributes) ProjectName() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("project_name"))
}

// Secret returns a reference to field secret of aws_codebuild_webhook.
func (cw codebuildWebhookAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("secret"))
}

// Url returns a reference to field url of aws_codebuild_webhook.
func (cw codebuildWebhookAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(cw.ref.Append("url"))
}

func (cw codebuildWebhookAttributes) FilterGroup() terra.SetValue[codebuildwebhook.FilterGroupAttributes] {
	return terra.ReferenceAsSet[codebuildwebhook.FilterGroupAttributes](cw.ref.Append("filter_group"))
}

type codebuildWebhookState struct {
	BranchFilter string                              `json:"branch_filter"`
	BuildType    string                              `json:"build_type"`
	Id           string                              `json:"id"`
	PayloadUrl   string                              `json:"payload_url"`
	ProjectName  string                              `json:"project_name"`
	Secret       string                              `json:"secret"`
	Url          string                              `json:"url"`
	FilterGroup  []codebuildwebhook.FilterGroupState `json:"filter_group"`
}

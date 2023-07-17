// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudbuildtrigger "github.com/golingon/terraproviders/google/4.73.1/cloudbuildtrigger"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudbuildTrigger creates a new instance of [CloudbuildTrigger].
func NewCloudbuildTrigger(name string, args CloudbuildTriggerArgs) *CloudbuildTrigger {
	return &CloudbuildTrigger{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudbuildTrigger)(nil)

// CloudbuildTrigger represents the Terraform resource google_cloudbuild_trigger.
type CloudbuildTrigger struct {
	Name      string
	Args      CloudbuildTriggerArgs
	state     *cloudbuildTriggerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudbuildTrigger].
func (ct *CloudbuildTrigger) Type() string {
	return "google_cloudbuild_trigger"
}

// LocalName returns the local name for [CloudbuildTrigger].
func (ct *CloudbuildTrigger) LocalName() string {
	return ct.Name
}

// Configuration returns the configuration (args) for [CloudbuildTrigger].
func (ct *CloudbuildTrigger) Configuration() interface{} {
	return ct.Args
}

// DependOn is used for other resources to depend on [CloudbuildTrigger].
func (ct *CloudbuildTrigger) DependOn() terra.Reference {
	return terra.ReferenceResource(ct)
}

// Dependencies returns the list of resources [CloudbuildTrigger] depends_on.
func (ct *CloudbuildTrigger) Dependencies() terra.Dependencies {
	return ct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudbuildTrigger].
func (ct *CloudbuildTrigger) LifecycleManagement() *terra.Lifecycle {
	return ct.Lifecycle
}

// Attributes returns the attributes for [CloudbuildTrigger].
func (ct *CloudbuildTrigger) Attributes() cloudbuildTriggerAttributes {
	return cloudbuildTriggerAttributes{ref: terra.ReferenceResource(ct)}
}

// ImportState imports the given attribute values into [CloudbuildTrigger]'s state.
func (ct *CloudbuildTrigger) ImportState(av io.Reader) error {
	ct.state = &cloudbuildTriggerState{}
	if err := json.NewDecoder(av).Decode(ct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ct.Type(), ct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudbuildTrigger] has state.
func (ct *CloudbuildTrigger) State() (*cloudbuildTriggerState, bool) {
	return ct.state, ct.state != nil
}

// StateMust returns the state for [CloudbuildTrigger]. Panics if the state is nil.
func (ct *CloudbuildTrigger) StateMust() *cloudbuildTriggerState {
	if ct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ct.Type(), ct.LocalName()))
	}
	return ct.state
}

// CloudbuildTriggerArgs contains the configurations for google_cloudbuild_trigger.
type CloudbuildTriggerArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Filename: string, optional
	Filename terra.StringValue `hcl:"filename,attr"`
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoredFiles: list of string, optional
	IgnoredFiles terra.ListValue[terra.StringValue] `hcl:"ignored_files,attr"`
	// IncludeBuildLogs: string, optional
	IncludeBuildLogs terra.StringValue `hcl:"include_build_logs,attr"`
	// IncludedFiles: list of string, optional
	IncludedFiles terra.ListValue[terra.StringValue] `hcl:"included_files,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceAccount: string, optional
	ServiceAccount terra.StringValue `hcl:"service_account,attr"`
	// Substitutions: map of string, optional
	Substitutions terra.MapValue[terra.StringValue] `hcl:"substitutions,attr"`
	// Tags: list of string, optional
	Tags terra.ListValue[terra.StringValue] `hcl:"tags,attr"`
	// ApprovalConfig: optional
	ApprovalConfig *cloudbuildtrigger.ApprovalConfig `hcl:"approval_config,block"`
	// BitbucketServerTriggerConfig: optional
	BitbucketServerTriggerConfig *cloudbuildtrigger.BitbucketServerTriggerConfig `hcl:"bitbucket_server_trigger_config,block"`
	// Build: optional
	Build *cloudbuildtrigger.Build `hcl:"build,block"`
	// GitFileSource: optional
	GitFileSource *cloudbuildtrigger.GitFileSource `hcl:"git_file_source,block"`
	// Github: optional
	Github *cloudbuildtrigger.Github `hcl:"github,block"`
	// PubsubConfig: optional
	PubsubConfig *cloudbuildtrigger.PubsubConfig `hcl:"pubsub_config,block"`
	// SourceToBuild: optional
	SourceToBuild *cloudbuildtrigger.SourceToBuild `hcl:"source_to_build,block"`
	// Timeouts: optional
	Timeouts *cloudbuildtrigger.Timeouts `hcl:"timeouts,block"`
	// TriggerTemplate: optional
	TriggerTemplate *cloudbuildtrigger.TriggerTemplate `hcl:"trigger_template,block"`
	// WebhookConfig: optional
	WebhookConfig *cloudbuildtrigger.WebhookConfig `hcl:"webhook_config,block"`
}
type cloudbuildTriggerAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("create_time"))
}

// Description returns a reference to field description of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(ct.ref.Append("disabled"))
}

// Filename returns a reference to field filename of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) Filename() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("filename"))
}

// Filter returns a reference to field filter of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("filter"))
}

// Id returns a reference to field id of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("id"))
}

// IgnoredFiles returns a reference to field ignored_files of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) IgnoredFiles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ct.ref.Append("ignored_files"))
}

// IncludeBuildLogs returns a reference to field include_build_logs of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) IncludeBuildLogs() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("include_build_logs"))
}

// IncludedFiles returns a reference to field included_files of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) IncludedFiles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ct.ref.Append("included_files"))
}

// Location returns a reference to field location of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("location"))
}

// Name returns a reference to field name of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("project"))
}

// ServiceAccount returns a reference to field service_account of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("service_account"))
}

// Substitutions returns a reference to field substitutions of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) Substitutions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ct.ref.Append("substitutions"))
}

// Tags returns a reference to field tags of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) Tags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ct.ref.Append("tags"))
}

// TriggerId returns a reference to field trigger_id of google_cloudbuild_trigger.
func (ct cloudbuildTriggerAttributes) TriggerId() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("trigger_id"))
}

func (ct cloudbuildTriggerAttributes) ApprovalConfig() terra.ListValue[cloudbuildtrigger.ApprovalConfigAttributes] {
	return terra.ReferenceAsList[cloudbuildtrigger.ApprovalConfigAttributes](ct.ref.Append("approval_config"))
}

func (ct cloudbuildTriggerAttributes) BitbucketServerTriggerConfig() terra.ListValue[cloudbuildtrigger.BitbucketServerTriggerConfigAttributes] {
	return terra.ReferenceAsList[cloudbuildtrigger.BitbucketServerTriggerConfigAttributes](ct.ref.Append("bitbucket_server_trigger_config"))
}

func (ct cloudbuildTriggerAttributes) Build() terra.ListValue[cloudbuildtrigger.BuildAttributes] {
	return terra.ReferenceAsList[cloudbuildtrigger.BuildAttributes](ct.ref.Append("build"))
}

func (ct cloudbuildTriggerAttributes) GitFileSource() terra.ListValue[cloudbuildtrigger.GitFileSourceAttributes] {
	return terra.ReferenceAsList[cloudbuildtrigger.GitFileSourceAttributes](ct.ref.Append("git_file_source"))
}

func (ct cloudbuildTriggerAttributes) Github() terra.ListValue[cloudbuildtrigger.GithubAttributes] {
	return terra.ReferenceAsList[cloudbuildtrigger.GithubAttributes](ct.ref.Append("github"))
}

func (ct cloudbuildTriggerAttributes) PubsubConfig() terra.ListValue[cloudbuildtrigger.PubsubConfigAttributes] {
	return terra.ReferenceAsList[cloudbuildtrigger.PubsubConfigAttributes](ct.ref.Append("pubsub_config"))
}

func (ct cloudbuildTriggerAttributes) SourceToBuild() terra.ListValue[cloudbuildtrigger.SourceToBuildAttributes] {
	return terra.ReferenceAsList[cloudbuildtrigger.SourceToBuildAttributes](ct.ref.Append("source_to_build"))
}

func (ct cloudbuildTriggerAttributes) Timeouts() cloudbuildtrigger.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudbuildtrigger.TimeoutsAttributes](ct.ref.Append("timeouts"))
}

func (ct cloudbuildTriggerAttributes) TriggerTemplate() terra.ListValue[cloudbuildtrigger.TriggerTemplateAttributes] {
	return terra.ReferenceAsList[cloudbuildtrigger.TriggerTemplateAttributes](ct.ref.Append("trigger_template"))
}

func (ct cloudbuildTriggerAttributes) WebhookConfig() terra.ListValue[cloudbuildtrigger.WebhookConfigAttributes] {
	return terra.ReferenceAsList[cloudbuildtrigger.WebhookConfigAttributes](ct.ref.Append("webhook_config"))
}

type cloudbuildTriggerState struct {
	CreateTime                   string                                                `json:"create_time"`
	Description                  string                                                `json:"description"`
	Disabled                     bool                                                  `json:"disabled"`
	Filename                     string                                                `json:"filename"`
	Filter                       string                                                `json:"filter"`
	Id                           string                                                `json:"id"`
	IgnoredFiles                 []string                                              `json:"ignored_files"`
	IncludeBuildLogs             string                                                `json:"include_build_logs"`
	IncludedFiles                []string                                              `json:"included_files"`
	Location                     string                                                `json:"location"`
	Name                         string                                                `json:"name"`
	Project                      string                                                `json:"project"`
	ServiceAccount               string                                                `json:"service_account"`
	Substitutions                map[string]string                                     `json:"substitutions"`
	Tags                         []string                                              `json:"tags"`
	TriggerId                    string                                                `json:"trigger_id"`
	ApprovalConfig               []cloudbuildtrigger.ApprovalConfigState               `json:"approval_config"`
	BitbucketServerTriggerConfig []cloudbuildtrigger.BitbucketServerTriggerConfigState `json:"bitbucket_server_trigger_config"`
	Build                        []cloudbuildtrigger.BuildState                        `json:"build"`
	GitFileSource                []cloudbuildtrigger.GitFileSourceState                `json:"git_file_source"`
	Github                       []cloudbuildtrigger.GithubState                       `json:"github"`
	PubsubConfig                 []cloudbuildtrigger.PubsubConfigState                 `json:"pubsub_config"`
	SourceToBuild                []cloudbuildtrigger.SourceToBuildState                `json:"source_to_build"`
	Timeouts                     *cloudbuildtrigger.TimeoutsState                      `json:"timeouts"`
	TriggerTemplate              []cloudbuildtrigger.TriggerTemplateState              `json:"trigger_template"`
	WebhookConfig                []cloudbuildtrigger.WebhookConfigState                `json:"webhook_config"`
}

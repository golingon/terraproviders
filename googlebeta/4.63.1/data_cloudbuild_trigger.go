// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacloudbuildtrigger "github.com/golingon/terraproviders/googlebeta/4.63.1/datacloudbuildtrigger"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudbuildTrigger creates a new instance of [DataCloudbuildTrigger].
func NewDataCloudbuildTrigger(name string, args DataCloudbuildTriggerArgs) *DataCloudbuildTrigger {
	return &DataCloudbuildTrigger{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudbuildTrigger)(nil)

// DataCloudbuildTrigger represents the Terraform data resource google_cloudbuild_trigger.
type DataCloudbuildTrigger struct {
	Name string
	Args DataCloudbuildTriggerArgs
}

// DataSource returns the Terraform object type for [DataCloudbuildTrigger].
func (ct *DataCloudbuildTrigger) DataSource() string {
	return "google_cloudbuild_trigger"
}

// LocalName returns the local name for [DataCloudbuildTrigger].
func (ct *DataCloudbuildTrigger) LocalName() string {
	return ct.Name
}

// Configuration returns the configuration (args) for [DataCloudbuildTrigger].
func (ct *DataCloudbuildTrigger) Configuration() interface{} {
	return ct.Args
}

// Attributes returns the attributes for [DataCloudbuildTrigger].
func (ct *DataCloudbuildTrigger) Attributes() dataCloudbuildTriggerAttributes {
	return dataCloudbuildTriggerAttributes{ref: terra.ReferenceDataResource(ct)}
}

// DataCloudbuildTriggerArgs contains the configurations for google_cloudbuild_trigger.
type DataCloudbuildTriggerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TriggerId: string, required
	TriggerId terra.StringValue `hcl:"trigger_id,attr" validate:"required"`
	// ApprovalConfig: min=0
	ApprovalConfig []datacloudbuildtrigger.ApprovalConfig `hcl:"approval_config,block" validate:"min=0"`
	// BitbucketServerTriggerConfig: min=0
	BitbucketServerTriggerConfig []datacloudbuildtrigger.BitbucketServerTriggerConfig `hcl:"bitbucket_server_trigger_config,block" validate:"min=0"`
	// Build: min=0
	Build []datacloudbuildtrigger.Build `hcl:"build,block" validate:"min=0"`
	// GitFileSource: min=0
	GitFileSource []datacloudbuildtrigger.GitFileSource `hcl:"git_file_source,block" validate:"min=0"`
	// Github: min=0
	Github []datacloudbuildtrigger.Github `hcl:"github,block" validate:"min=0"`
	// PubsubConfig: min=0
	PubsubConfig []datacloudbuildtrigger.PubsubConfig `hcl:"pubsub_config,block" validate:"min=0"`
	// RepositoryEventConfig: min=0
	RepositoryEventConfig []datacloudbuildtrigger.RepositoryEventConfig `hcl:"repository_event_config,block" validate:"min=0"`
	// SourceToBuild: min=0
	SourceToBuild []datacloudbuildtrigger.SourceToBuild `hcl:"source_to_build,block" validate:"min=0"`
	// TriggerTemplate: min=0
	TriggerTemplate []datacloudbuildtrigger.TriggerTemplate `hcl:"trigger_template,block" validate:"min=0"`
	// WebhookConfig: min=0
	WebhookConfig []datacloudbuildtrigger.WebhookConfig `hcl:"webhook_config,block" validate:"min=0"`
}
type dataCloudbuildTriggerAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("create_time"))
}

// Description returns a reference to field description of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(ct.ref.Append("disabled"))
}

// Filename returns a reference to field filename of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) Filename() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("filename"))
}

// Filter returns a reference to field filter of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("filter"))
}

// Id returns a reference to field id of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("id"))
}

// IgnoredFiles returns a reference to field ignored_files of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) IgnoredFiles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ct.ref.Append("ignored_files"))
}

// IncludeBuildLogs returns a reference to field include_build_logs of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) IncludeBuildLogs() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("include_build_logs"))
}

// IncludedFiles returns a reference to field included_files of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) IncludedFiles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ct.ref.Append("included_files"))
}

// Location returns a reference to field location of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("location"))
}

// Name returns a reference to field name of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("project"))
}

// ServiceAccount returns a reference to field service_account of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("service_account"))
}

// Substitutions returns a reference to field substitutions of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) Substitutions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ct.ref.Append("substitutions"))
}

// Tags returns a reference to field tags of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) Tags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ct.ref.Append("tags"))
}

// TriggerId returns a reference to field trigger_id of google_cloudbuild_trigger.
func (ct dataCloudbuildTriggerAttributes) TriggerId() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("trigger_id"))
}

func (ct dataCloudbuildTriggerAttributes) ApprovalConfig() terra.ListValue[datacloudbuildtrigger.ApprovalConfigAttributes] {
	return terra.ReferenceAsList[datacloudbuildtrigger.ApprovalConfigAttributes](ct.ref.Append("approval_config"))
}

func (ct dataCloudbuildTriggerAttributes) BitbucketServerTriggerConfig() terra.ListValue[datacloudbuildtrigger.BitbucketServerTriggerConfigAttributes] {
	return terra.ReferenceAsList[datacloudbuildtrigger.BitbucketServerTriggerConfigAttributes](ct.ref.Append("bitbucket_server_trigger_config"))
}

func (ct dataCloudbuildTriggerAttributes) Build() terra.ListValue[datacloudbuildtrigger.BuildAttributes] {
	return terra.ReferenceAsList[datacloudbuildtrigger.BuildAttributes](ct.ref.Append("build"))
}

func (ct dataCloudbuildTriggerAttributes) GitFileSource() terra.ListValue[datacloudbuildtrigger.GitFileSourceAttributes] {
	return terra.ReferenceAsList[datacloudbuildtrigger.GitFileSourceAttributes](ct.ref.Append("git_file_source"))
}

func (ct dataCloudbuildTriggerAttributes) Github() terra.ListValue[datacloudbuildtrigger.GithubAttributes] {
	return terra.ReferenceAsList[datacloudbuildtrigger.GithubAttributes](ct.ref.Append("github"))
}

func (ct dataCloudbuildTriggerAttributes) PubsubConfig() terra.ListValue[datacloudbuildtrigger.PubsubConfigAttributes] {
	return terra.ReferenceAsList[datacloudbuildtrigger.PubsubConfigAttributes](ct.ref.Append("pubsub_config"))
}

func (ct dataCloudbuildTriggerAttributes) RepositoryEventConfig() terra.ListValue[datacloudbuildtrigger.RepositoryEventConfigAttributes] {
	return terra.ReferenceAsList[datacloudbuildtrigger.RepositoryEventConfigAttributes](ct.ref.Append("repository_event_config"))
}

func (ct dataCloudbuildTriggerAttributes) SourceToBuild() terra.ListValue[datacloudbuildtrigger.SourceToBuildAttributes] {
	return terra.ReferenceAsList[datacloudbuildtrigger.SourceToBuildAttributes](ct.ref.Append("source_to_build"))
}

func (ct dataCloudbuildTriggerAttributes) TriggerTemplate() terra.ListValue[datacloudbuildtrigger.TriggerTemplateAttributes] {
	return terra.ReferenceAsList[datacloudbuildtrigger.TriggerTemplateAttributes](ct.ref.Append("trigger_template"))
}

func (ct dataCloudbuildTriggerAttributes) WebhookConfig() terra.ListValue[datacloudbuildtrigger.WebhookConfigAttributes] {
	return terra.ReferenceAsList[datacloudbuildtrigger.WebhookConfigAttributes](ct.ref.Append("webhook_config"))
}

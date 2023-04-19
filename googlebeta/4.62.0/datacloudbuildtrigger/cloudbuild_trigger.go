// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacloudbuildtrigger

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ApprovalConfig struct{}

type BitbucketServerTriggerConfig struct {
	// BitbucketServerTriggerConfigPullRequest: min=0
	PullRequest []BitbucketServerTriggerConfigPullRequest `hcl:"pull_request,block" validate:"min=0"`
	// BitbucketServerTriggerConfigPush: min=0
	Push []BitbucketServerTriggerConfigPush `hcl:"push,block" validate:"min=0"`
}

type BitbucketServerTriggerConfigPullRequest struct{}

type BitbucketServerTriggerConfigPush struct{}

type Build struct {
	// Artifacts: min=0
	Artifacts []Artifacts `hcl:"artifacts,block" validate:"min=0"`
	// AvailableSecrets: min=0
	AvailableSecrets []AvailableSecrets `hcl:"available_secrets,block" validate:"min=0"`
	// Options: min=0
	Options []Options `hcl:"options,block" validate:"min=0"`
	// Secret: min=0
	Secret []Secret `hcl:"secret,block" validate:"min=0"`
	// Source: min=0
	Source []Source `hcl:"source,block" validate:"min=0"`
	// Step: min=0
	Step []Step `hcl:"step,block" validate:"min=0"`
}

type Artifacts struct {
	// Objects: min=0
	Objects []Objects `hcl:"objects,block" validate:"min=0"`
}

type Objects struct {
	// Timing: min=0
	Timing []Timing `hcl:"timing,block" validate:"min=0"`
}

type Timing struct{}

type AvailableSecrets struct {
	// SecretManager: min=0
	SecretManager []SecretManager `hcl:"secret_manager,block" validate:"min=0"`
}

type SecretManager struct{}

type Options struct {
	// OptionsVolumes: min=0
	Volumes []OptionsVolumes `hcl:"volumes,block" validate:"min=0"`
}

type OptionsVolumes struct{}

type Secret struct{}

type Source struct {
	// RepoSource: min=0
	RepoSource []RepoSource `hcl:"repo_source,block" validate:"min=0"`
	// StorageSource: min=0
	StorageSource []StorageSource `hcl:"storage_source,block" validate:"min=0"`
}

type RepoSource struct{}

type StorageSource struct{}

type Step struct {
	// StepVolumes: min=0
	Volumes []StepVolumes `hcl:"volumes,block" validate:"min=0"`
}

type StepVolumes struct{}

type GitFileSource struct{}

type Github struct {
	// GithubPullRequest: min=0
	PullRequest []GithubPullRequest `hcl:"pull_request,block" validate:"min=0"`
	// GithubPush: min=0
	Push []GithubPush `hcl:"push,block" validate:"min=0"`
}

type GithubPullRequest struct{}

type GithubPush struct{}

type PubsubConfig struct{}

type RepositoryEventConfig struct {
	// RepositoryEventConfigPullRequest: min=0
	PullRequest []RepositoryEventConfigPullRequest `hcl:"pull_request,block" validate:"min=0"`
	// RepositoryEventConfigPush: min=0
	Push []RepositoryEventConfigPush `hcl:"push,block" validate:"min=0"`
}

type RepositoryEventConfigPullRequest struct{}

type RepositoryEventConfigPush struct{}

type SourceToBuild struct{}

type TriggerTemplate struct{}

type WebhookConfig struct{}

type ApprovalConfigAttributes struct {
	ref terra.Reference
}

func (ac ApprovalConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac ApprovalConfigAttributes) InternalWithRef(ref terra.Reference) ApprovalConfigAttributes {
	return ApprovalConfigAttributes{ref: ref}
}

func (ac ApprovalConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac ApprovalConfigAttributes) ApprovalRequired() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("approval_required"))
}

type BitbucketServerTriggerConfigAttributes struct {
	ref terra.Reference
}

func (bstc BitbucketServerTriggerConfigAttributes) InternalRef() (terra.Reference, error) {
	return bstc.ref, nil
}

func (bstc BitbucketServerTriggerConfigAttributes) InternalWithRef(ref terra.Reference) BitbucketServerTriggerConfigAttributes {
	return BitbucketServerTriggerConfigAttributes{ref: ref}
}

func (bstc BitbucketServerTriggerConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bstc.ref.InternalTokens()
}

func (bstc BitbucketServerTriggerConfigAttributes) BitbucketServerConfigResource() terra.StringValue {
	return terra.ReferenceAsString(bstc.ref.Append("bitbucket_server_config_resource"))
}

func (bstc BitbucketServerTriggerConfigAttributes) ProjectKey() terra.StringValue {
	return terra.ReferenceAsString(bstc.ref.Append("project_key"))
}

func (bstc BitbucketServerTriggerConfigAttributes) RepoSlug() terra.StringValue {
	return terra.ReferenceAsString(bstc.ref.Append("repo_slug"))
}

func (bstc BitbucketServerTriggerConfigAttributes) PullRequest() terra.ListValue[BitbucketServerTriggerConfigPullRequestAttributes] {
	return terra.ReferenceAsList[BitbucketServerTriggerConfigPullRequestAttributes](bstc.ref.Append("pull_request"))
}

func (bstc BitbucketServerTriggerConfigAttributes) Push() terra.ListValue[BitbucketServerTriggerConfigPushAttributes] {
	return terra.ReferenceAsList[BitbucketServerTriggerConfigPushAttributes](bstc.ref.Append("push"))
}

type BitbucketServerTriggerConfigPullRequestAttributes struct {
	ref terra.Reference
}

func (pr BitbucketServerTriggerConfigPullRequestAttributes) InternalRef() (terra.Reference, error) {
	return pr.ref, nil
}

func (pr BitbucketServerTriggerConfigPullRequestAttributes) InternalWithRef(ref terra.Reference) BitbucketServerTriggerConfigPullRequestAttributes {
	return BitbucketServerTriggerConfigPullRequestAttributes{ref: ref}
}

func (pr BitbucketServerTriggerConfigPullRequestAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pr.ref.InternalTokens()
}

func (pr BitbucketServerTriggerConfigPullRequestAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(pr.ref.Append("branch"))
}

func (pr BitbucketServerTriggerConfigPullRequestAttributes) CommentControl() terra.StringValue {
	return terra.ReferenceAsString(pr.ref.Append("comment_control"))
}

func (pr BitbucketServerTriggerConfigPullRequestAttributes) InvertRegex() terra.BoolValue {
	return terra.ReferenceAsBool(pr.ref.Append("invert_regex"))
}

type BitbucketServerTriggerConfigPushAttributes struct {
	ref terra.Reference
}

func (p BitbucketServerTriggerConfigPushAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p BitbucketServerTriggerConfigPushAttributes) InternalWithRef(ref terra.Reference) BitbucketServerTriggerConfigPushAttributes {
	return BitbucketServerTriggerConfigPushAttributes{ref: ref}
}

func (p BitbucketServerTriggerConfigPushAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p BitbucketServerTriggerConfigPushAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("branch"))
}

func (p BitbucketServerTriggerConfigPushAttributes) InvertRegex() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("invert_regex"))
}

func (p BitbucketServerTriggerConfigPushAttributes) Tag() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("tag"))
}

type BuildAttributes struct {
	ref terra.Reference
}

func (b BuildAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BuildAttributes) InternalWithRef(ref terra.Reference) BuildAttributes {
	return BuildAttributes{ref: ref}
}

func (b BuildAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BuildAttributes) Images() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](b.ref.Append("images"))
}

func (b BuildAttributes) LogsBucket() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("logs_bucket"))
}

func (b BuildAttributes) QueueTtl() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("queue_ttl"))
}

func (b BuildAttributes) Substitutions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](b.ref.Append("substitutions"))
}

func (b BuildAttributes) Tags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](b.ref.Append("tags"))
}

func (b BuildAttributes) Timeout() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("timeout"))
}

func (b BuildAttributes) Artifacts() terra.ListValue[ArtifactsAttributes] {
	return terra.ReferenceAsList[ArtifactsAttributes](b.ref.Append("artifacts"))
}

func (b BuildAttributes) AvailableSecrets() terra.ListValue[AvailableSecretsAttributes] {
	return terra.ReferenceAsList[AvailableSecretsAttributes](b.ref.Append("available_secrets"))
}

func (b BuildAttributes) Options() terra.ListValue[OptionsAttributes] {
	return terra.ReferenceAsList[OptionsAttributes](b.ref.Append("options"))
}

func (b BuildAttributes) Secret() terra.ListValue[SecretAttributes] {
	return terra.ReferenceAsList[SecretAttributes](b.ref.Append("secret"))
}

func (b BuildAttributes) Source() terra.ListValue[SourceAttributes] {
	return terra.ReferenceAsList[SourceAttributes](b.ref.Append("source"))
}

func (b BuildAttributes) Step() terra.ListValue[StepAttributes] {
	return terra.ReferenceAsList[StepAttributes](b.ref.Append("step"))
}

type ArtifactsAttributes struct {
	ref terra.Reference
}

func (a ArtifactsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ArtifactsAttributes) InternalWithRef(ref terra.Reference) ArtifactsAttributes {
	return ArtifactsAttributes{ref: ref}
}

func (a ArtifactsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ArtifactsAttributes) Images() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](a.ref.Append("images"))
}

func (a ArtifactsAttributes) Objects() terra.ListValue[ObjectsAttributes] {
	return terra.ReferenceAsList[ObjectsAttributes](a.ref.Append("objects"))
}

type ObjectsAttributes struct {
	ref terra.Reference
}

func (o ObjectsAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o ObjectsAttributes) InternalWithRef(ref terra.Reference) ObjectsAttributes {
	return ObjectsAttributes{ref: ref}
}

func (o ObjectsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o ObjectsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("location"))
}

func (o ObjectsAttributes) Paths() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](o.ref.Append("paths"))
}

func (o ObjectsAttributes) Timing() terra.ListValue[TimingAttributes] {
	return terra.ReferenceAsList[TimingAttributes](o.ref.Append("timing"))
}

type TimingAttributes struct {
	ref terra.Reference
}

func (t TimingAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimingAttributes) InternalWithRef(ref terra.Reference) TimingAttributes {
	return TimingAttributes{ref: ref}
}

func (t TimingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimingAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("end_time"))
}

func (t TimingAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("start_time"))
}

type AvailableSecretsAttributes struct {
	ref terra.Reference
}

func (as AvailableSecretsAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as AvailableSecretsAttributes) InternalWithRef(ref terra.Reference) AvailableSecretsAttributes {
	return AvailableSecretsAttributes{ref: ref}
}

func (as AvailableSecretsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as AvailableSecretsAttributes) SecretManager() terra.ListValue[SecretManagerAttributes] {
	return terra.ReferenceAsList[SecretManagerAttributes](as.ref.Append("secret_manager"))
}

type SecretManagerAttributes struct {
	ref terra.Reference
}

func (sm SecretManagerAttributes) InternalRef() (terra.Reference, error) {
	return sm.ref, nil
}

func (sm SecretManagerAttributes) InternalWithRef(ref terra.Reference) SecretManagerAttributes {
	return SecretManagerAttributes{ref: ref}
}

func (sm SecretManagerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sm.ref.InternalTokens()
}

func (sm SecretManagerAttributes) Env() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("env"))
}

func (sm SecretManagerAttributes) VersionName() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("version_name"))
}

type OptionsAttributes struct {
	ref terra.Reference
}

func (o OptionsAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OptionsAttributes) InternalWithRef(ref terra.Reference) OptionsAttributes {
	return OptionsAttributes{ref: ref}
}

func (o OptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OptionsAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(o.ref.Append("disk_size_gb"))
}

func (o OptionsAttributes) DynamicSubstitutions() terra.BoolValue {
	return terra.ReferenceAsBool(o.ref.Append("dynamic_substitutions"))
}

func (o OptionsAttributes) Env() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](o.ref.Append("env"))
}

func (o OptionsAttributes) LogStreamingOption() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("log_streaming_option"))
}

func (o OptionsAttributes) Logging() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("logging"))
}

func (o OptionsAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("machine_type"))
}

func (o OptionsAttributes) RequestedVerifyOption() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("requested_verify_option"))
}

func (o OptionsAttributes) SecretEnv() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](o.ref.Append("secret_env"))
}

func (o OptionsAttributes) SourceProvenanceHash() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](o.ref.Append("source_provenance_hash"))
}

func (o OptionsAttributes) SubstitutionOption() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("substitution_option"))
}

func (o OptionsAttributes) WorkerPool() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("worker_pool"))
}

func (o OptionsAttributes) Volumes() terra.ListValue[OptionsVolumesAttributes] {
	return terra.ReferenceAsList[OptionsVolumesAttributes](o.ref.Append("volumes"))
}

type OptionsVolumesAttributes struct {
	ref terra.Reference
}

func (v OptionsVolumesAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v OptionsVolumesAttributes) InternalWithRef(ref terra.Reference) OptionsVolumesAttributes {
	return OptionsVolumesAttributes{ref: ref}
}

func (v OptionsVolumesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v OptionsVolumesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("name"))
}

func (v OptionsVolumesAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("path"))
}

type SecretAttributes struct {
	ref terra.Reference
}

func (s SecretAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SecretAttributes) InternalWithRef(ref terra.Reference) SecretAttributes {
	return SecretAttributes{ref: ref}
}

func (s SecretAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SecretAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("kms_key_name"))
}

func (s SecretAttributes) SecretEnv() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("secret_env"))
}

type SourceAttributes struct {
	ref terra.Reference
}

func (s SourceAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SourceAttributes) InternalWithRef(ref terra.Reference) SourceAttributes {
	return SourceAttributes{ref: ref}
}

func (s SourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SourceAttributes) RepoSource() terra.ListValue[RepoSourceAttributes] {
	return terra.ReferenceAsList[RepoSourceAttributes](s.ref.Append("repo_source"))
}

func (s SourceAttributes) StorageSource() terra.ListValue[StorageSourceAttributes] {
	return terra.ReferenceAsList[StorageSourceAttributes](s.ref.Append("storage_source"))
}

type RepoSourceAttributes struct {
	ref terra.Reference
}

func (rs RepoSourceAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs RepoSourceAttributes) InternalWithRef(ref terra.Reference) RepoSourceAttributes {
	return RepoSourceAttributes{ref: ref}
}

func (rs RepoSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs RepoSourceAttributes) BranchName() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("branch_name"))
}

func (rs RepoSourceAttributes) CommitSha() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("commit_sha"))
}

func (rs RepoSourceAttributes) Dir() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("dir"))
}

func (rs RepoSourceAttributes) InvertRegex() terra.BoolValue {
	return terra.ReferenceAsBool(rs.ref.Append("invert_regex"))
}

func (rs RepoSourceAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("project_id"))
}

func (rs RepoSourceAttributes) RepoName() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("repo_name"))
}

func (rs RepoSourceAttributes) Substitutions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rs.ref.Append("substitutions"))
}

func (rs RepoSourceAttributes) TagName() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("tag_name"))
}

type StorageSourceAttributes struct {
	ref terra.Reference
}

func (ss StorageSourceAttributes) InternalRef() (terra.Reference, error) {
	return ss.ref, nil
}

func (ss StorageSourceAttributes) InternalWithRef(ref terra.Reference) StorageSourceAttributes {
	return StorageSourceAttributes{ref: ref}
}

func (ss StorageSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ss.ref.InternalTokens()
}

func (ss StorageSourceAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("bucket"))
}

func (ss StorageSourceAttributes) Generation() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("generation"))
}

func (ss StorageSourceAttributes) Object() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("object"))
}

type StepAttributes struct {
	ref terra.Reference
}

func (s StepAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StepAttributes) InternalWithRef(ref terra.Reference) StepAttributes {
	return StepAttributes{ref: ref}
}

func (s StepAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StepAttributes) Args() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("args"))
}

func (s StepAttributes) Dir() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("dir"))
}

func (s StepAttributes) Entrypoint() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("entrypoint"))
}

func (s StepAttributes) Env() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("env"))
}

func (s StepAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

func (s StepAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s StepAttributes) Script() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("script"))
}

func (s StepAttributes) SecretEnv() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("secret_env"))
}

func (s StepAttributes) Timeout() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("timeout"))
}

func (s StepAttributes) Timing() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("timing"))
}

func (s StepAttributes) WaitFor() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("wait_for"))
}

func (s StepAttributes) Volumes() terra.ListValue[StepVolumesAttributes] {
	return terra.ReferenceAsList[StepVolumesAttributes](s.ref.Append("volumes"))
}

type StepVolumesAttributes struct {
	ref terra.Reference
}

func (v StepVolumesAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v StepVolumesAttributes) InternalWithRef(ref terra.Reference) StepVolumesAttributes {
	return StepVolumesAttributes{ref: ref}
}

func (v StepVolumesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v StepVolumesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("name"))
}

func (v StepVolumesAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("path"))
}

type GitFileSourceAttributes struct {
	ref terra.Reference
}

func (gfs GitFileSourceAttributes) InternalRef() (terra.Reference, error) {
	return gfs.ref, nil
}

func (gfs GitFileSourceAttributes) InternalWithRef(ref terra.Reference) GitFileSourceAttributes {
	return GitFileSourceAttributes{ref: ref}
}

func (gfs GitFileSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gfs.ref.InternalTokens()
}

func (gfs GitFileSourceAttributes) GithubEnterpriseConfig() terra.StringValue {
	return terra.ReferenceAsString(gfs.ref.Append("github_enterprise_config"))
}

func (gfs GitFileSourceAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(gfs.ref.Append("path"))
}

func (gfs GitFileSourceAttributes) RepoType() terra.StringValue {
	return terra.ReferenceAsString(gfs.ref.Append("repo_type"))
}

func (gfs GitFileSourceAttributes) Revision() terra.StringValue {
	return terra.ReferenceAsString(gfs.ref.Append("revision"))
}

func (gfs GitFileSourceAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(gfs.ref.Append("uri"))
}

type GithubAttributes struct {
	ref terra.Reference
}

func (g GithubAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GithubAttributes) InternalWithRef(ref terra.Reference) GithubAttributes {
	return GithubAttributes{ref: ref}
}

func (g GithubAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GithubAttributes) EnterpriseConfigResourceName() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("enterprise_config_resource_name"))
}

func (g GithubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("name"))
}

func (g GithubAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("owner"))
}

func (g GithubAttributes) PullRequest() terra.ListValue[GithubPullRequestAttributes] {
	return terra.ReferenceAsList[GithubPullRequestAttributes](g.ref.Append("pull_request"))
}

func (g GithubAttributes) Push() terra.ListValue[GithubPushAttributes] {
	return terra.ReferenceAsList[GithubPushAttributes](g.ref.Append("push"))
}

type GithubPullRequestAttributes struct {
	ref terra.Reference
}

func (pr GithubPullRequestAttributes) InternalRef() (terra.Reference, error) {
	return pr.ref, nil
}

func (pr GithubPullRequestAttributes) InternalWithRef(ref terra.Reference) GithubPullRequestAttributes {
	return GithubPullRequestAttributes{ref: ref}
}

func (pr GithubPullRequestAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pr.ref.InternalTokens()
}

func (pr GithubPullRequestAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(pr.ref.Append("branch"))
}

func (pr GithubPullRequestAttributes) CommentControl() terra.StringValue {
	return terra.ReferenceAsString(pr.ref.Append("comment_control"))
}

func (pr GithubPullRequestAttributes) InvertRegex() terra.BoolValue {
	return terra.ReferenceAsBool(pr.ref.Append("invert_regex"))
}

type GithubPushAttributes struct {
	ref terra.Reference
}

func (p GithubPushAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p GithubPushAttributes) InternalWithRef(ref terra.Reference) GithubPushAttributes {
	return GithubPushAttributes{ref: ref}
}

func (p GithubPushAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p GithubPushAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("branch"))
}

func (p GithubPushAttributes) InvertRegex() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("invert_regex"))
}

func (p GithubPushAttributes) Tag() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("tag"))
}

type PubsubConfigAttributes struct {
	ref terra.Reference
}

func (pc PubsubConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PubsubConfigAttributes) InternalWithRef(ref terra.Reference) PubsubConfigAttributes {
	return PubsubConfigAttributes{ref: ref}
}

func (pc PubsubConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PubsubConfigAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("service_account_email"))
}

func (pc PubsubConfigAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("state"))
}

func (pc PubsubConfigAttributes) Subscription() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("subscription"))
}

func (pc PubsubConfigAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("topic"))
}

type RepositoryEventConfigAttributes struct {
	ref terra.Reference
}

func (rec RepositoryEventConfigAttributes) InternalRef() (terra.Reference, error) {
	return rec.ref, nil
}

func (rec RepositoryEventConfigAttributes) InternalWithRef(ref terra.Reference) RepositoryEventConfigAttributes {
	return RepositoryEventConfigAttributes{ref: ref}
}

func (rec RepositoryEventConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rec.ref.InternalTokens()
}

func (rec RepositoryEventConfigAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(rec.ref.Append("repository"))
}

func (rec RepositoryEventConfigAttributes) PullRequest() terra.ListValue[RepositoryEventConfigPullRequestAttributes] {
	return terra.ReferenceAsList[RepositoryEventConfigPullRequestAttributes](rec.ref.Append("pull_request"))
}

func (rec RepositoryEventConfigAttributes) Push() terra.ListValue[RepositoryEventConfigPushAttributes] {
	return terra.ReferenceAsList[RepositoryEventConfigPushAttributes](rec.ref.Append("push"))
}

type RepositoryEventConfigPullRequestAttributes struct {
	ref terra.Reference
}

func (pr RepositoryEventConfigPullRequestAttributes) InternalRef() (terra.Reference, error) {
	return pr.ref, nil
}

func (pr RepositoryEventConfigPullRequestAttributes) InternalWithRef(ref terra.Reference) RepositoryEventConfigPullRequestAttributes {
	return RepositoryEventConfigPullRequestAttributes{ref: ref}
}

func (pr RepositoryEventConfigPullRequestAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pr.ref.InternalTokens()
}

func (pr RepositoryEventConfigPullRequestAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(pr.ref.Append("branch"))
}

func (pr RepositoryEventConfigPullRequestAttributes) CommentControl() terra.StringValue {
	return terra.ReferenceAsString(pr.ref.Append("comment_control"))
}

func (pr RepositoryEventConfigPullRequestAttributes) InvertRegex() terra.BoolValue {
	return terra.ReferenceAsBool(pr.ref.Append("invert_regex"))
}

type RepositoryEventConfigPushAttributes struct {
	ref terra.Reference
}

func (p RepositoryEventConfigPushAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p RepositoryEventConfigPushAttributes) InternalWithRef(ref terra.Reference) RepositoryEventConfigPushAttributes {
	return RepositoryEventConfigPushAttributes{ref: ref}
}

func (p RepositoryEventConfigPushAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p RepositoryEventConfigPushAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("branch"))
}

func (p RepositoryEventConfigPushAttributes) InvertRegex() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("invert_regex"))
}

func (p RepositoryEventConfigPushAttributes) Tag() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("tag"))
}

type SourceToBuildAttributes struct {
	ref terra.Reference
}

func (stb SourceToBuildAttributes) InternalRef() (terra.Reference, error) {
	return stb.ref, nil
}

func (stb SourceToBuildAttributes) InternalWithRef(ref terra.Reference) SourceToBuildAttributes {
	return SourceToBuildAttributes{ref: ref}
}

func (stb SourceToBuildAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return stb.ref.InternalTokens()
}

func (stb SourceToBuildAttributes) GithubEnterpriseConfig() terra.StringValue {
	return terra.ReferenceAsString(stb.ref.Append("github_enterprise_config"))
}

func (stb SourceToBuildAttributes) Ref() terra.StringValue {
	return terra.ReferenceAsString(stb.ref.Append("ref"))
}

func (stb SourceToBuildAttributes) RepoType() terra.StringValue {
	return terra.ReferenceAsString(stb.ref.Append("repo_type"))
}

func (stb SourceToBuildAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(stb.ref.Append("uri"))
}

type TriggerTemplateAttributes struct {
	ref terra.Reference
}

func (tt TriggerTemplateAttributes) InternalRef() (terra.Reference, error) {
	return tt.ref, nil
}

func (tt TriggerTemplateAttributes) InternalWithRef(ref terra.Reference) TriggerTemplateAttributes {
	return TriggerTemplateAttributes{ref: ref}
}

func (tt TriggerTemplateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tt.ref.InternalTokens()
}

func (tt TriggerTemplateAttributes) BranchName() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("branch_name"))
}

func (tt TriggerTemplateAttributes) CommitSha() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("commit_sha"))
}

func (tt TriggerTemplateAttributes) Dir() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("dir"))
}

func (tt TriggerTemplateAttributes) InvertRegex() terra.BoolValue {
	return terra.ReferenceAsBool(tt.ref.Append("invert_regex"))
}

func (tt TriggerTemplateAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("project_id"))
}

func (tt TriggerTemplateAttributes) RepoName() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("repo_name"))
}

func (tt TriggerTemplateAttributes) TagName() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("tag_name"))
}

type WebhookConfigAttributes struct {
	ref terra.Reference
}

func (wc WebhookConfigAttributes) InternalRef() (terra.Reference, error) {
	return wc.ref, nil
}

func (wc WebhookConfigAttributes) InternalWithRef(ref terra.Reference) WebhookConfigAttributes {
	return WebhookConfigAttributes{ref: ref}
}

func (wc WebhookConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wc.ref.InternalTokens()
}

func (wc WebhookConfigAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(wc.ref.Append("secret"))
}

func (wc WebhookConfigAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(wc.ref.Append("state"))
}

type ApprovalConfigState struct {
	ApprovalRequired bool `json:"approval_required"`
}

type BitbucketServerTriggerConfigState struct {
	BitbucketServerConfigResource string                                         `json:"bitbucket_server_config_resource"`
	ProjectKey                    string                                         `json:"project_key"`
	RepoSlug                      string                                         `json:"repo_slug"`
	PullRequest                   []BitbucketServerTriggerConfigPullRequestState `json:"pull_request"`
	Push                          []BitbucketServerTriggerConfigPushState        `json:"push"`
}

type BitbucketServerTriggerConfigPullRequestState struct {
	Branch         string `json:"branch"`
	CommentControl string `json:"comment_control"`
	InvertRegex    bool   `json:"invert_regex"`
}

type BitbucketServerTriggerConfigPushState struct {
	Branch      string `json:"branch"`
	InvertRegex bool   `json:"invert_regex"`
	Tag         string `json:"tag"`
}

type BuildState struct {
	Images           []string                `json:"images"`
	LogsBucket       string                  `json:"logs_bucket"`
	QueueTtl         string                  `json:"queue_ttl"`
	Substitutions    map[string]string       `json:"substitutions"`
	Tags             []string                `json:"tags"`
	Timeout          string                  `json:"timeout"`
	Artifacts        []ArtifactsState        `json:"artifacts"`
	AvailableSecrets []AvailableSecretsState `json:"available_secrets"`
	Options          []OptionsState          `json:"options"`
	Secret           []SecretState           `json:"secret"`
	Source           []SourceState           `json:"source"`
	Step             []StepState             `json:"step"`
}

type ArtifactsState struct {
	Images  []string       `json:"images"`
	Objects []ObjectsState `json:"objects"`
}

type ObjectsState struct {
	Location string        `json:"location"`
	Paths    []string      `json:"paths"`
	Timing   []TimingState `json:"timing"`
}

type TimingState struct {
	EndTime   string `json:"end_time"`
	StartTime string `json:"start_time"`
}

type AvailableSecretsState struct {
	SecretManager []SecretManagerState `json:"secret_manager"`
}

type SecretManagerState struct {
	Env         string `json:"env"`
	VersionName string `json:"version_name"`
}

type OptionsState struct {
	DiskSizeGb            float64               `json:"disk_size_gb"`
	DynamicSubstitutions  bool                  `json:"dynamic_substitutions"`
	Env                   []string              `json:"env"`
	LogStreamingOption    string                `json:"log_streaming_option"`
	Logging               string                `json:"logging"`
	MachineType           string                `json:"machine_type"`
	RequestedVerifyOption string                `json:"requested_verify_option"`
	SecretEnv             []string              `json:"secret_env"`
	SourceProvenanceHash  []string              `json:"source_provenance_hash"`
	SubstitutionOption    string                `json:"substitution_option"`
	WorkerPool            string                `json:"worker_pool"`
	Volumes               []OptionsVolumesState `json:"volumes"`
}

type OptionsVolumesState struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type SecretState struct {
	KmsKeyName string            `json:"kms_key_name"`
	SecretEnv  map[string]string `json:"secret_env"`
}

type SourceState struct {
	RepoSource    []RepoSourceState    `json:"repo_source"`
	StorageSource []StorageSourceState `json:"storage_source"`
}

type RepoSourceState struct {
	BranchName    string            `json:"branch_name"`
	CommitSha     string            `json:"commit_sha"`
	Dir           string            `json:"dir"`
	InvertRegex   bool              `json:"invert_regex"`
	ProjectId     string            `json:"project_id"`
	RepoName      string            `json:"repo_name"`
	Substitutions map[string]string `json:"substitutions"`
	TagName       string            `json:"tag_name"`
}

type StorageSourceState struct {
	Bucket     string `json:"bucket"`
	Generation string `json:"generation"`
	Object     string `json:"object"`
}

type StepState struct {
	Args       []string           `json:"args"`
	Dir        string             `json:"dir"`
	Entrypoint string             `json:"entrypoint"`
	Env        []string           `json:"env"`
	Id         string             `json:"id"`
	Name       string             `json:"name"`
	Script     string             `json:"script"`
	SecretEnv  []string           `json:"secret_env"`
	Timeout    string             `json:"timeout"`
	Timing     string             `json:"timing"`
	WaitFor    []string           `json:"wait_for"`
	Volumes    []StepVolumesState `json:"volumes"`
}

type StepVolumesState struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type GitFileSourceState struct {
	GithubEnterpriseConfig string `json:"github_enterprise_config"`
	Path                   string `json:"path"`
	RepoType               string `json:"repo_type"`
	Revision               string `json:"revision"`
	Uri                    string `json:"uri"`
}

type GithubState struct {
	EnterpriseConfigResourceName string                   `json:"enterprise_config_resource_name"`
	Name                         string                   `json:"name"`
	Owner                        string                   `json:"owner"`
	PullRequest                  []GithubPullRequestState `json:"pull_request"`
	Push                         []GithubPushState        `json:"push"`
}

type GithubPullRequestState struct {
	Branch         string `json:"branch"`
	CommentControl string `json:"comment_control"`
	InvertRegex    bool   `json:"invert_regex"`
}

type GithubPushState struct {
	Branch      string `json:"branch"`
	InvertRegex bool   `json:"invert_regex"`
	Tag         string `json:"tag"`
}

type PubsubConfigState struct {
	ServiceAccountEmail string `json:"service_account_email"`
	State               string `json:"state"`
	Subscription        string `json:"subscription"`
	Topic               string `json:"topic"`
}

type RepositoryEventConfigState struct {
	Repository  string                                  `json:"repository"`
	PullRequest []RepositoryEventConfigPullRequestState `json:"pull_request"`
	Push        []RepositoryEventConfigPushState        `json:"push"`
}

type RepositoryEventConfigPullRequestState struct {
	Branch         string `json:"branch"`
	CommentControl string `json:"comment_control"`
	InvertRegex    bool   `json:"invert_regex"`
}

type RepositoryEventConfigPushState struct {
	Branch      string `json:"branch"`
	InvertRegex bool   `json:"invert_regex"`
	Tag         string `json:"tag"`
}

type SourceToBuildState struct {
	GithubEnterpriseConfig string `json:"github_enterprise_config"`
	Ref                    string `json:"ref"`
	RepoType               string `json:"repo_type"`
	Uri                    string `json:"uri"`
}

type TriggerTemplateState struct {
	BranchName  string `json:"branch_name"`
	CommitSha   string `json:"commit_sha"`
	Dir         string `json:"dir"`
	InvertRegex bool   `json:"invert_regex"`
	ProjectId   string `json:"project_id"`
	RepoName    string `json:"repo_name"`
	TagName     string `json:"tag_name"`
}

type WebhookConfigState struct {
	Secret string `json:"secret"`
	State  string `json:"state"`
}

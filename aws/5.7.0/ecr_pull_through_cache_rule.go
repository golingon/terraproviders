// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEcrPullThroughCacheRule creates a new instance of [EcrPullThroughCacheRule].
func NewEcrPullThroughCacheRule(name string, args EcrPullThroughCacheRuleArgs) *EcrPullThroughCacheRule {
	return &EcrPullThroughCacheRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcrPullThroughCacheRule)(nil)

// EcrPullThroughCacheRule represents the Terraform resource aws_ecr_pull_through_cache_rule.
type EcrPullThroughCacheRule struct {
	Name      string
	Args      EcrPullThroughCacheRuleArgs
	state     *ecrPullThroughCacheRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EcrPullThroughCacheRule].
func (eptcr *EcrPullThroughCacheRule) Type() string {
	return "aws_ecr_pull_through_cache_rule"
}

// LocalName returns the local name for [EcrPullThroughCacheRule].
func (eptcr *EcrPullThroughCacheRule) LocalName() string {
	return eptcr.Name
}

// Configuration returns the configuration (args) for [EcrPullThroughCacheRule].
func (eptcr *EcrPullThroughCacheRule) Configuration() interface{} {
	return eptcr.Args
}

// DependOn is used for other resources to depend on [EcrPullThroughCacheRule].
func (eptcr *EcrPullThroughCacheRule) DependOn() terra.Reference {
	return terra.ReferenceResource(eptcr)
}

// Dependencies returns the list of resources [EcrPullThroughCacheRule] depends_on.
func (eptcr *EcrPullThroughCacheRule) Dependencies() terra.Dependencies {
	return eptcr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EcrPullThroughCacheRule].
func (eptcr *EcrPullThroughCacheRule) LifecycleManagement() *terra.Lifecycle {
	return eptcr.Lifecycle
}

// Attributes returns the attributes for [EcrPullThroughCacheRule].
func (eptcr *EcrPullThroughCacheRule) Attributes() ecrPullThroughCacheRuleAttributes {
	return ecrPullThroughCacheRuleAttributes{ref: terra.ReferenceResource(eptcr)}
}

// ImportState imports the given attribute values into [EcrPullThroughCacheRule]'s state.
func (eptcr *EcrPullThroughCacheRule) ImportState(av io.Reader) error {
	eptcr.state = &ecrPullThroughCacheRuleState{}
	if err := json.NewDecoder(av).Decode(eptcr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eptcr.Type(), eptcr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EcrPullThroughCacheRule] has state.
func (eptcr *EcrPullThroughCacheRule) State() (*ecrPullThroughCacheRuleState, bool) {
	return eptcr.state, eptcr.state != nil
}

// StateMust returns the state for [EcrPullThroughCacheRule]. Panics if the state is nil.
func (eptcr *EcrPullThroughCacheRule) StateMust() *ecrPullThroughCacheRuleState {
	if eptcr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eptcr.Type(), eptcr.LocalName()))
	}
	return eptcr.state
}

// EcrPullThroughCacheRuleArgs contains the configurations for aws_ecr_pull_through_cache_rule.
type EcrPullThroughCacheRuleArgs struct {
	// EcrRepositoryPrefix: string, required
	EcrRepositoryPrefix terra.StringValue `hcl:"ecr_repository_prefix,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// UpstreamRegistryUrl: string, required
	UpstreamRegistryUrl terra.StringValue `hcl:"upstream_registry_url,attr" validate:"required"`
}
type ecrPullThroughCacheRuleAttributes struct {
	ref terra.Reference
}

// EcrRepositoryPrefix returns a reference to field ecr_repository_prefix of aws_ecr_pull_through_cache_rule.
func (eptcr ecrPullThroughCacheRuleAttributes) EcrRepositoryPrefix() terra.StringValue {
	return terra.ReferenceAsString(eptcr.ref.Append("ecr_repository_prefix"))
}

// Id returns a reference to field id of aws_ecr_pull_through_cache_rule.
func (eptcr ecrPullThroughCacheRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eptcr.ref.Append("id"))
}

// RegistryId returns a reference to field registry_id of aws_ecr_pull_through_cache_rule.
func (eptcr ecrPullThroughCacheRuleAttributes) RegistryId() terra.StringValue {
	return terra.ReferenceAsString(eptcr.ref.Append("registry_id"))
}

// UpstreamRegistryUrl returns a reference to field upstream_registry_url of aws_ecr_pull_through_cache_rule.
func (eptcr ecrPullThroughCacheRuleAttributes) UpstreamRegistryUrl() terra.StringValue {
	return terra.ReferenceAsString(eptcr.ref.Append("upstream_registry_url"))
}

type ecrPullThroughCacheRuleState struct {
	EcrRepositoryPrefix string `json:"ecr_repository_prefix"`
	Id                  string `json:"id"`
	RegistryId          string `json:"registry_id"`
	UpstreamRegistryUrl string `json:"upstream_registry_url"`
}

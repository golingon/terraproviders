// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewRuntimeconfigConfigIamPolicy creates a new instance of [RuntimeconfigConfigIamPolicy].
func NewRuntimeconfigConfigIamPolicy(name string, args RuntimeconfigConfigIamPolicyArgs) *RuntimeconfigConfigIamPolicy {
	return &RuntimeconfigConfigIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RuntimeconfigConfigIamPolicy)(nil)

// RuntimeconfigConfigIamPolicy represents the Terraform resource google_runtimeconfig_config_iam_policy.
type RuntimeconfigConfigIamPolicy struct {
	Name      string
	Args      RuntimeconfigConfigIamPolicyArgs
	state     *runtimeconfigConfigIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RuntimeconfigConfigIamPolicy].
func (rcip *RuntimeconfigConfigIamPolicy) Type() string {
	return "google_runtimeconfig_config_iam_policy"
}

// LocalName returns the local name for [RuntimeconfigConfigIamPolicy].
func (rcip *RuntimeconfigConfigIamPolicy) LocalName() string {
	return rcip.Name
}

// Configuration returns the configuration (args) for [RuntimeconfigConfigIamPolicy].
func (rcip *RuntimeconfigConfigIamPolicy) Configuration() interface{} {
	return rcip.Args
}

// DependOn is used for other resources to depend on [RuntimeconfigConfigIamPolicy].
func (rcip *RuntimeconfigConfigIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(rcip)
}

// Dependencies returns the list of resources [RuntimeconfigConfigIamPolicy] depends_on.
func (rcip *RuntimeconfigConfigIamPolicy) Dependencies() terra.Dependencies {
	return rcip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RuntimeconfigConfigIamPolicy].
func (rcip *RuntimeconfigConfigIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return rcip.Lifecycle
}

// Attributes returns the attributes for [RuntimeconfigConfigIamPolicy].
func (rcip *RuntimeconfigConfigIamPolicy) Attributes() runtimeconfigConfigIamPolicyAttributes {
	return runtimeconfigConfigIamPolicyAttributes{ref: terra.ReferenceResource(rcip)}
}

// ImportState imports the given attribute values into [RuntimeconfigConfigIamPolicy]'s state.
func (rcip *RuntimeconfigConfigIamPolicy) ImportState(av io.Reader) error {
	rcip.state = &runtimeconfigConfigIamPolicyState{}
	if err := json.NewDecoder(av).Decode(rcip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rcip.Type(), rcip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RuntimeconfigConfigIamPolicy] has state.
func (rcip *RuntimeconfigConfigIamPolicy) State() (*runtimeconfigConfigIamPolicyState, bool) {
	return rcip.state, rcip.state != nil
}

// StateMust returns the state for [RuntimeconfigConfigIamPolicy]. Panics if the state is nil.
func (rcip *RuntimeconfigConfigIamPolicy) StateMust() *runtimeconfigConfigIamPolicyState {
	if rcip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rcip.Type(), rcip.LocalName()))
	}
	return rcip.state
}

// RuntimeconfigConfigIamPolicyArgs contains the configurations for google_runtimeconfig_config_iam_policy.
type RuntimeconfigConfigIamPolicyArgs struct {
	// Config: string, required
	Config terra.StringValue `hcl:"config,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type runtimeconfigConfigIamPolicyAttributes struct {
	ref terra.Reference
}

// Config returns a reference to field config of google_runtimeconfig_config_iam_policy.
func (rcip runtimeconfigConfigIamPolicyAttributes) Config() terra.StringValue {
	return terra.ReferenceAsString(rcip.ref.Append("config"))
}

// Etag returns a reference to field etag of google_runtimeconfig_config_iam_policy.
func (rcip runtimeconfigConfigIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(rcip.ref.Append("etag"))
}

// Id returns a reference to field id of google_runtimeconfig_config_iam_policy.
func (rcip runtimeconfigConfigIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rcip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_runtimeconfig_config_iam_policy.
func (rcip runtimeconfigConfigIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(rcip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_runtimeconfig_config_iam_policy.
func (rcip runtimeconfigConfigIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(rcip.ref.Append("project"))
}

type runtimeconfigConfigIamPolicyState struct {
	Config     string `json:"config"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_scc_source_iam_policy

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_scc_source_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleSccSourceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gssip *Resource) Type() string {
	return "google_scc_source_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gssip *Resource) LocalName() string {
	return gssip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gssip *Resource) Configuration() interface{} {
	return gssip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gssip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gssip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gssip *Resource) Dependencies() terra.Dependencies {
	return gssip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gssip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gssip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gssip *Resource) Attributes() googleSccSourceIamPolicyAttributes {
	return googleSccSourceIamPolicyAttributes{ref: terra.ReferenceResource(gssip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gssip *Resource) ImportState(state io.Reader) error {
	gssip.state = &googleSccSourceIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gssip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gssip.Type(), gssip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gssip *Resource) State() (*googleSccSourceIamPolicyState, bool) {
	return gssip.state, gssip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gssip *Resource) StateMust() *googleSccSourceIamPolicyState {
	if gssip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gssip.Type(), gssip.LocalName()))
	}
	return gssip.state
}

// Args contains the configurations for google_scc_source_iam_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
}

type googleSccSourceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_scc_source_iam_policy.
func (gssip googleSccSourceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gssip.ref.Append("etag"))
}

// Id returns a reference to field id of google_scc_source_iam_policy.
func (gssip googleSccSourceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gssip.ref.Append("id"))
}

// Organization returns a reference to field organization of google_scc_source_iam_policy.
func (gssip googleSccSourceIamPolicyAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(gssip.ref.Append("organization"))
}

// PolicyData returns a reference to field policy_data of google_scc_source_iam_policy.
func (gssip googleSccSourceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gssip.ref.Append("policy_data"))
}

// Source returns a reference to field source of google_scc_source_iam_policy.
func (gssip googleSccSourceIamPolicyAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(gssip.ref.Append("source"))
}

type googleSccSourceIamPolicyState struct {
	Etag         string `json:"etag"`
	Id           string `json:"id"`
	Organization string `json:"organization"`
	PolicyData   string `json:"policy_data"`
	Source       string `json:"source"`
}

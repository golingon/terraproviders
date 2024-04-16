// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_privateca_ca_pool_iam_policy

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

// Resource represents the Terraform resource google_privateca_ca_pool_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googlePrivatecaCaPoolIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gpcpip *Resource) Type() string {
	return "google_privateca_ca_pool_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gpcpip *Resource) LocalName() string {
	return gpcpip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gpcpip *Resource) Configuration() interface{} {
	return gpcpip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gpcpip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gpcpip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gpcpip *Resource) Dependencies() terra.Dependencies {
	return gpcpip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gpcpip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gpcpip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gpcpip *Resource) Attributes() googlePrivatecaCaPoolIamPolicyAttributes {
	return googlePrivatecaCaPoolIamPolicyAttributes{ref: terra.ReferenceResource(gpcpip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gpcpip *Resource) ImportState(state io.Reader) error {
	gpcpip.state = &googlePrivatecaCaPoolIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gpcpip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gpcpip.Type(), gpcpip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gpcpip *Resource) State() (*googlePrivatecaCaPoolIamPolicyState, bool) {
	return gpcpip.state, gpcpip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gpcpip *Resource) StateMust() *googlePrivatecaCaPoolIamPolicyState {
	if gpcpip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gpcpip.Type(), gpcpip.LocalName()))
	}
	return gpcpip.state
}

// Args contains the configurations for google_privateca_ca_pool_iam_policy.
type Args struct {
	// CaPool: string, required
	CaPool terra.StringValue `hcl:"ca_pool,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type googlePrivatecaCaPoolIamPolicyAttributes struct {
	ref terra.Reference
}

// CaPool returns a reference to field ca_pool of google_privateca_ca_pool_iam_policy.
func (gpcpip googlePrivatecaCaPoolIamPolicyAttributes) CaPool() terra.StringValue {
	return terra.ReferenceAsString(gpcpip.ref.Append("ca_pool"))
}

// Etag returns a reference to field etag of google_privateca_ca_pool_iam_policy.
func (gpcpip googlePrivatecaCaPoolIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gpcpip.ref.Append("etag"))
}

// Id returns a reference to field id of google_privateca_ca_pool_iam_policy.
func (gpcpip googlePrivatecaCaPoolIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gpcpip.ref.Append("id"))
}

// Location returns a reference to field location of google_privateca_ca_pool_iam_policy.
func (gpcpip googlePrivatecaCaPoolIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gpcpip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_privateca_ca_pool_iam_policy.
func (gpcpip googlePrivatecaCaPoolIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gpcpip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_privateca_ca_pool_iam_policy.
func (gpcpip googlePrivatecaCaPoolIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gpcpip.ref.Append("project"))
}

type googlePrivatecaCaPoolIamPolicyState struct {
	CaPool     string `json:"ca_pool"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}

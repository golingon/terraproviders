// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudbuildv2ConnectionIamPolicy creates a new instance of [Cloudbuildv2ConnectionIamPolicy].
func NewCloudbuildv2ConnectionIamPolicy(name string, args Cloudbuildv2ConnectionIamPolicyArgs) *Cloudbuildv2ConnectionIamPolicy {
	return &Cloudbuildv2ConnectionIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudbuildv2ConnectionIamPolicy)(nil)

// Cloudbuildv2ConnectionIamPolicy represents the Terraform resource google_cloudbuildv2_connection_iam_policy.
type Cloudbuildv2ConnectionIamPolicy struct {
	Name      string
	Args      Cloudbuildv2ConnectionIamPolicyArgs
	state     *cloudbuildv2ConnectionIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Cloudbuildv2ConnectionIamPolicy].
func (ccip *Cloudbuildv2ConnectionIamPolicy) Type() string {
	return "google_cloudbuildv2_connection_iam_policy"
}

// LocalName returns the local name for [Cloudbuildv2ConnectionIamPolicy].
func (ccip *Cloudbuildv2ConnectionIamPolicy) LocalName() string {
	return ccip.Name
}

// Configuration returns the configuration (args) for [Cloudbuildv2ConnectionIamPolicy].
func (ccip *Cloudbuildv2ConnectionIamPolicy) Configuration() interface{} {
	return ccip.Args
}

// DependOn is used for other resources to depend on [Cloudbuildv2ConnectionIamPolicy].
func (ccip *Cloudbuildv2ConnectionIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ccip)
}

// Dependencies returns the list of resources [Cloudbuildv2ConnectionIamPolicy] depends_on.
func (ccip *Cloudbuildv2ConnectionIamPolicy) Dependencies() terra.Dependencies {
	return ccip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Cloudbuildv2ConnectionIamPolicy].
func (ccip *Cloudbuildv2ConnectionIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return ccip.Lifecycle
}

// Attributes returns the attributes for [Cloudbuildv2ConnectionIamPolicy].
func (ccip *Cloudbuildv2ConnectionIamPolicy) Attributes() cloudbuildv2ConnectionIamPolicyAttributes {
	return cloudbuildv2ConnectionIamPolicyAttributes{ref: terra.ReferenceResource(ccip)}
}

// ImportState imports the given attribute values into [Cloudbuildv2ConnectionIamPolicy]'s state.
func (ccip *Cloudbuildv2ConnectionIamPolicy) ImportState(av io.Reader) error {
	ccip.state = &cloudbuildv2ConnectionIamPolicyState{}
	if err := json.NewDecoder(av).Decode(ccip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccip.Type(), ccip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Cloudbuildv2ConnectionIamPolicy] has state.
func (ccip *Cloudbuildv2ConnectionIamPolicy) State() (*cloudbuildv2ConnectionIamPolicyState, bool) {
	return ccip.state, ccip.state != nil
}

// StateMust returns the state for [Cloudbuildv2ConnectionIamPolicy]. Panics if the state is nil.
func (ccip *Cloudbuildv2ConnectionIamPolicy) StateMust() *cloudbuildv2ConnectionIamPolicyState {
	if ccip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccip.Type(), ccip.LocalName()))
	}
	return ccip.state
}

// Cloudbuildv2ConnectionIamPolicyArgs contains the configurations for google_cloudbuildv2_connection_iam_policy.
type Cloudbuildv2ConnectionIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type cloudbuildv2ConnectionIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloudbuildv2_connection_iam_policy.
func (ccip cloudbuildv2ConnectionIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ccip.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudbuildv2_connection_iam_policy.
func (ccip cloudbuildv2ConnectionIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccip.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudbuildv2_connection_iam_policy.
func (ccip cloudbuildv2ConnectionIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ccip.ref.Append("location"))
}

// Name returns a reference to field name of google_cloudbuildv2_connection_iam_policy.
func (ccip cloudbuildv2ConnectionIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_cloudbuildv2_connection_iam_policy.
func (ccip cloudbuildv2ConnectionIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ccip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_cloudbuildv2_connection_iam_policy.
func (ccip cloudbuildv2ConnectionIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ccip.ref.Append("project"))
}

type cloudbuildv2ConnectionIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
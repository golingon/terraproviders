// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewCloudfunctions2FunctionIamPolicy creates a new instance of [Cloudfunctions2FunctionIamPolicy].
func NewCloudfunctions2FunctionIamPolicy(name string, args Cloudfunctions2FunctionIamPolicyArgs) *Cloudfunctions2FunctionIamPolicy {
	return &Cloudfunctions2FunctionIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudfunctions2FunctionIamPolicy)(nil)

// Cloudfunctions2FunctionIamPolicy represents the Terraform resource google_cloudfunctions2_function_iam_policy.
type Cloudfunctions2FunctionIamPolicy struct {
	Name      string
	Args      Cloudfunctions2FunctionIamPolicyArgs
	state     *cloudfunctions2FunctionIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Cloudfunctions2FunctionIamPolicy].
func (cfip *Cloudfunctions2FunctionIamPolicy) Type() string {
	return "google_cloudfunctions2_function_iam_policy"
}

// LocalName returns the local name for [Cloudfunctions2FunctionIamPolicy].
func (cfip *Cloudfunctions2FunctionIamPolicy) LocalName() string {
	return cfip.Name
}

// Configuration returns the configuration (args) for [Cloudfunctions2FunctionIamPolicy].
func (cfip *Cloudfunctions2FunctionIamPolicy) Configuration() interface{} {
	return cfip.Args
}

// DependOn is used for other resources to depend on [Cloudfunctions2FunctionIamPolicy].
func (cfip *Cloudfunctions2FunctionIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cfip)
}

// Dependencies returns the list of resources [Cloudfunctions2FunctionIamPolicy] depends_on.
func (cfip *Cloudfunctions2FunctionIamPolicy) Dependencies() terra.Dependencies {
	return cfip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Cloudfunctions2FunctionIamPolicy].
func (cfip *Cloudfunctions2FunctionIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return cfip.Lifecycle
}

// Attributes returns the attributes for [Cloudfunctions2FunctionIamPolicy].
func (cfip *Cloudfunctions2FunctionIamPolicy) Attributes() cloudfunctions2FunctionIamPolicyAttributes {
	return cloudfunctions2FunctionIamPolicyAttributes{ref: terra.ReferenceResource(cfip)}
}

// ImportState imports the given attribute values into [Cloudfunctions2FunctionIamPolicy]'s state.
func (cfip *Cloudfunctions2FunctionIamPolicy) ImportState(av io.Reader) error {
	cfip.state = &cloudfunctions2FunctionIamPolicyState{}
	if err := json.NewDecoder(av).Decode(cfip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfip.Type(), cfip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Cloudfunctions2FunctionIamPolicy] has state.
func (cfip *Cloudfunctions2FunctionIamPolicy) State() (*cloudfunctions2FunctionIamPolicyState, bool) {
	return cfip.state, cfip.state != nil
}

// StateMust returns the state for [Cloudfunctions2FunctionIamPolicy]. Panics if the state is nil.
func (cfip *Cloudfunctions2FunctionIamPolicy) StateMust() *cloudfunctions2FunctionIamPolicyState {
	if cfip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfip.Type(), cfip.LocalName()))
	}
	return cfip.state
}

// Cloudfunctions2FunctionIamPolicyArgs contains the configurations for google_cloudfunctions2_function_iam_policy.
type Cloudfunctions2FunctionIamPolicyArgs struct {
	// CloudFunction: string, required
	CloudFunction terra.StringValue `hcl:"cloud_function,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type cloudfunctions2FunctionIamPolicyAttributes struct {
	ref terra.Reference
}

// CloudFunction returns a reference to field cloud_function of google_cloudfunctions2_function_iam_policy.
func (cfip cloudfunctions2FunctionIamPolicyAttributes) CloudFunction() terra.StringValue {
	return terra.ReferenceAsString(cfip.ref.Append("cloud_function"))
}

// Etag returns a reference to field etag of google_cloudfunctions2_function_iam_policy.
func (cfip cloudfunctions2FunctionIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cfip.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudfunctions2_function_iam_policy.
func (cfip cloudfunctions2FunctionIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfip.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudfunctions2_function_iam_policy.
func (cfip cloudfunctions2FunctionIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cfip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_cloudfunctions2_function_iam_policy.
func (cfip cloudfunctions2FunctionIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(cfip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_cloudfunctions2_function_iam_policy.
func (cfip cloudfunctions2FunctionIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cfip.ref.Append("project"))
}

type cloudfunctions2FunctionIamPolicyState struct {
	CloudFunction string `json:"cloud_function"`
	Etag          string `json:"etag"`
	Id            string `json:"id"`
	Location      string `json:"location"`
	PolicyData    string `json:"policy_data"`
	Project       string `json:"project"`
}

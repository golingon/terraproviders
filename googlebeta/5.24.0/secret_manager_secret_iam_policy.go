// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewSecretManagerSecretIamPolicy creates a new instance of [SecretManagerSecretIamPolicy].
func NewSecretManagerSecretIamPolicy(name string, args SecretManagerSecretIamPolicyArgs) *SecretManagerSecretIamPolicy {
	return &SecretManagerSecretIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecretManagerSecretIamPolicy)(nil)

// SecretManagerSecretIamPolicy represents the Terraform resource google_secret_manager_secret_iam_policy.
type SecretManagerSecretIamPolicy struct {
	Name      string
	Args      SecretManagerSecretIamPolicyArgs
	state     *secretManagerSecretIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecretManagerSecretIamPolicy].
func (smsip *SecretManagerSecretIamPolicy) Type() string {
	return "google_secret_manager_secret_iam_policy"
}

// LocalName returns the local name for [SecretManagerSecretIamPolicy].
func (smsip *SecretManagerSecretIamPolicy) LocalName() string {
	return smsip.Name
}

// Configuration returns the configuration (args) for [SecretManagerSecretIamPolicy].
func (smsip *SecretManagerSecretIamPolicy) Configuration() interface{} {
	return smsip.Args
}

// DependOn is used for other resources to depend on [SecretManagerSecretIamPolicy].
func (smsip *SecretManagerSecretIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(smsip)
}

// Dependencies returns the list of resources [SecretManagerSecretIamPolicy] depends_on.
func (smsip *SecretManagerSecretIamPolicy) Dependencies() terra.Dependencies {
	return smsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecretManagerSecretIamPolicy].
func (smsip *SecretManagerSecretIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return smsip.Lifecycle
}

// Attributes returns the attributes for [SecretManagerSecretIamPolicy].
func (smsip *SecretManagerSecretIamPolicy) Attributes() secretManagerSecretIamPolicyAttributes {
	return secretManagerSecretIamPolicyAttributes{ref: terra.ReferenceResource(smsip)}
}

// ImportState imports the given attribute values into [SecretManagerSecretIamPolicy]'s state.
func (smsip *SecretManagerSecretIamPolicy) ImportState(av io.Reader) error {
	smsip.state = &secretManagerSecretIamPolicyState{}
	if err := json.NewDecoder(av).Decode(smsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smsip.Type(), smsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecretManagerSecretIamPolicy] has state.
func (smsip *SecretManagerSecretIamPolicy) State() (*secretManagerSecretIamPolicyState, bool) {
	return smsip.state, smsip.state != nil
}

// StateMust returns the state for [SecretManagerSecretIamPolicy]. Panics if the state is nil.
func (smsip *SecretManagerSecretIamPolicy) StateMust() *secretManagerSecretIamPolicyState {
	if smsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smsip.Type(), smsip.LocalName()))
	}
	return smsip.state
}

// SecretManagerSecretIamPolicyArgs contains the configurations for google_secret_manager_secret_iam_policy.
type SecretManagerSecretIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
}
type secretManagerSecretIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_secret_manager_secret_iam_policy.
func (smsip secretManagerSecretIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(smsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_secret_manager_secret_iam_policy.
func (smsip secretManagerSecretIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_secret_manager_secret_iam_policy.
func (smsip secretManagerSecretIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(smsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_secret_manager_secret_iam_policy.
func (smsip secretManagerSecretIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(smsip.ref.Append("project"))
}

// SecretId returns a reference to field secret_id of google_secret_manager_secret_iam_policy.
func (smsip secretManagerSecretIamPolicyAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(smsip.ref.Append("secret_id"))
}

type secretManagerSecretIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	SecretId   string `json:"secret_id"`
}

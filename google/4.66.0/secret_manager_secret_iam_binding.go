// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	secretmanagersecretiambinding "github.com/golingon/terraproviders/google/4.66.0/secretmanagersecretiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecretManagerSecretIamBinding creates a new instance of [SecretManagerSecretIamBinding].
func NewSecretManagerSecretIamBinding(name string, args SecretManagerSecretIamBindingArgs) *SecretManagerSecretIamBinding {
	return &SecretManagerSecretIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecretManagerSecretIamBinding)(nil)

// SecretManagerSecretIamBinding represents the Terraform resource google_secret_manager_secret_iam_binding.
type SecretManagerSecretIamBinding struct {
	Name      string
	Args      SecretManagerSecretIamBindingArgs
	state     *secretManagerSecretIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecretManagerSecretIamBinding].
func (smsib *SecretManagerSecretIamBinding) Type() string {
	return "google_secret_manager_secret_iam_binding"
}

// LocalName returns the local name for [SecretManagerSecretIamBinding].
func (smsib *SecretManagerSecretIamBinding) LocalName() string {
	return smsib.Name
}

// Configuration returns the configuration (args) for [SecretManagerSecretIamBinding].
func (smsib *SecretManagerSecretIamBinding) Configuration() interface{} {
	return smsib.Args
}

// DependOn is used for other resources to depend on [SecretManagerSecretIamBinding].
func (smsib *SecretManagerSecretIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(smsib)
}

// Dependencies returns the list of resources [SecretManagerSecretIamBinding] depends_on.
func (smsib *SecretManagerSecretIamBinding) Dependencies() terra.Dependencies {
	return smsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecretManagerSecretIamBinding].
func (smsib *SecretManagerSecretIamBinding) LifecycleManagement() *terra.Lifecycle {
	return smsib.Lifecycle
}

// Attributes returns the attributes for [SecretManagerSecretIamBinding].
func (smsib *SecretManagerSecretIamBinding) Attributes() secretManagerSecretIamBindingAttributes {
	return secretManagerSecretIamBindingAttributes{ref: terra.ReferenceResource(smsib)}
}

// ImportState imports the given attribute values into [SecretManagerSecretIamBinding]'s state.
func (smsib *SecretManagerSecretIamBinding) ImportState(av io.Reader) error {
	smsib.state = &secretManagerSecretIamBindingState{}
	if err := json.NewDecoder(av).Decode(smsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smsib.Type(), smsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecretManagerSecretIamBinding] has state.
func (smsib *SecretManagerSecretIamBinding) State() (*secretManagerSecretIamBindingState, bool) {
	return smsib.state, smsib.state != nil
}

// StateMust returns the state for [SecretManagerSecretIamBinding]. Panics if the state is nil.
func (smsib *SecretManagerSecretIamBinding) StateMust() *secretManagerSecretIamBindingState {
	if smsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smsib.Type(), smsib.LocalName()))
	}
	return smsib.state
}

// SecretManagerSecretIamBindingArgs contains the configurations for google_secret_manager_secret_iam_binding.
type SecretManagerSecretIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
	// Condition: optional
	Condition *secretmanagersecretiambinding.Condition `hcl:"condition,block"`
}
type secretManagerSecretIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_secret_manager_secret_iam_binding.
func (smsib secretManagerSecretIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(smsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_secret_manager_secret_iam_binding.
func (smsib secretManagerSecretIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smsib.ref.Append("id"))
}

// Members returns a reference to field members of google_secret_manager_secret_iam_binding.
func (smsib secretManagerSecretIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](smsib.ref.Append("members"))
}

// Project returns a reference to field project of google_secret_manager_secret_iam_binding.
func (smsib secretManagerSecretIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(smsib.ref.Append("project"))
}

// Role returns a reference to field role of google_secret_manager_secret_iam_binding.
func (smsib secretManagerSecretIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(smsib.ref.Append("role"))
}

// SecretId returns a reference to field secret_id of google_secret_manager_secret_iam_binding.
func (smsib secretManagerSecretIamBindingAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(smsib.ref.Append("secret_id"))
}

func (smsib secretManagerSecretIamBindingAttributes) Condition() terra.ListValue[secretmanagersecretiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[secretmanagersecretiambinding.ConditionAttributes](smsib.ref.Append("condition"))
}

type secretManagerSecretIamBindingState struct {
	Etag      string                                         `json:"etag"`
	Id        string                                         `json:"id"`
	Members   []string                                       `json:"members"`
	Project   string                                         `json:"project"`
	Role      string                                         `json:"role"`
	SecretId  string                                         `json:"secret_id"`
	Condition []secretmanagersecretiambinding.ConditionState `json:"condition"`
}

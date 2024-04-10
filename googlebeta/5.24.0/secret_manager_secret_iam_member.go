// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	secretmanagersecretiammember "github.com/golingon/terraproviders/googlebeta/5.24.0/secretmanagersecretiammember"
	"io"
)

// NewSecretManagerSecretIamMember creates a new instance of [SecretManagerSecretIamMember].
func NewSecretManagerSecretIamMember(name string, args SecretManagerSecretIamMemberArgs) *SecretManagerSecretIamMember {
	return &SecretManagerSecretIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecretManagerSecretIamMember)(nil)

// SecretManagerSecretIamMember represents the Terraform resource google_secret_manager_secret_iam_member.
type SecretManagerSecretIamMember struct {
	Name      string
	Args      SecretManagerSecretIamMemberArgs
	state     *secretManagerSecretIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecretManagerSecretIamMember].
func (smsim *SecretManagerSecretIamMember) Type() string {
	return "google_secret_manager_secret_iam_member"
}

// LocalName returns the local name for [SecretManagerSecretIamMember].
func (smsim *SecretManagerSecretIamMember) LocalName() string {
	return smsim.Name
}

// Configuration returns the configuration (args) for [SecretManagerSecretIamMember].
func (smsim *SecretManagerSecretIamMember) Configuration() interface{} {
	return smsim.Args
}

// DependOn is used for other resources to depend on [SecretManagerSecretIamMember].
func (smsim *SecretManagerSecretIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(smsim)
}

// Dependencies returns the list of resources [SecretManagerSecretIamMember] depends_on.
func (smsim *SecretManagerSecretIamMember) Dependencies() terra.Dependencies {
	return smsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecretManagerSecretIamMember].
func (smsim *SecretManagerSecretIamMember) LifecycleManagement() *terra.Lifecycle {
	return smsim.Lifecycle
}

// Attributes returns the attributes for [SecretManagerSecretIamMember].
func (smsim *SecretManagerSecretIamMember) Attributes() secretManagerSecretIamMemberAttributes {
	return secretManagerSecretIamMemberAttributes{ref: terra.ReferenceResource(smsim)}
}

// ImportState imports the given attribute values into [SecretManagerSecretIamMember]'s state.
func (smsim *SecretManagerSecretIamMember) ImportState(av io.Reader) error {
	smsim.state = &secretManagerSecretIamMemberState{}
	if err := json.NewDecoder(av).Decode(smsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smsim.Type(), smsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecretManagerSecretIamMember] has state.
func (smsim *SecretManagerSecretIamMember) State() (*secretManagerSecretIamMemberState, bool) {
	return smsim.state, smsim.state != nil
}

// StateMust returns the state for [SecretManagerSecretIamMember]. Panics if the state is nil.
func (smsim *SecretManagerSecretIamMember) StateMust() *secretManagerSecretIamMemberState {
	if smsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smsim.Type(), smsim.LocalName()))
	}
	return smsim.state
}

// SecretManagerSecretIamMemberArgs contains the configurations for google_secret_manager_secret_iam_member.
type SecretManagerSecretIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
	// Condition: optional
	Condition *secretmanagersecretiammember.Condition `hcl:"condition,block"`
}
type secretManagerSecretIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_secret_manager_secret_iam_member.
func (smsim secretManagerSecretIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(smsim.ref.Append("etag"))
}

// Id returns a reference to field id of google_secret_manager_secret_iam_member.
func (smsim secretManagerSecretIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smsim.ref.Append("id"))
}

// Member returns a reference to field member of google_secret_manager_secret_iam_member.
func (smsim secretManagerSecretIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(smsim.ref.Append("member"))
}

// Project returns a reference to field project of google_secret_manager_secret_iam_member.
func (smsim secretManagerSecretIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(smsim.ref.Append("project"))
}

// Role returns a reference to field role of google_secret_manager_secret_iam_member.
func (smsim secretManagerSecretIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(smsim.ref.Append("role"))
}

// SecretId returns a reference to field secret_id of google_secret_manager_secret_iam_member.
func (smsim secretManagerSecretIamMemberAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(smsim.ref.Append("secret_id"))
}

func (smsim secretManagerSecretIamMemberAttributes) Condition() terra.ListValue[secretmanagersecretiammember.ConditionAttributes] {
	return terra.ReferenceAsList[secretmanagersecretiammember.ConditionAttributes](smsim.ref.Append("condition"))
}

type secretManagerSecretIamMemberState struct {
	Etag      string                                        `json:"etag"`
	Id        string                                        `json:"id"`
	Member    string                                        `json:"member"`
	Project   string                                        `json:"project"`
	Role      string                                        `json:"role"`
	SecretId  string                                        `json:"secret_id"`
	Condition []secretmanagersecretiammember.ConditionState `json:"condition"`
}

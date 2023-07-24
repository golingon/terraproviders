// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kmsgrant "github.com/golingon/terraproviders/aws/5.9.0/kmsgrant"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmsGrant creates a new instance of [KmsGrant].
func NewKmsGrant(name string, args KmsGrantArgs) *KmsGrant {
	return &KmsGrant{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsGrant)(nil)

// KmsGrant represents the Terraform resource aws_kms_grant.
type KmsGrant struct {
	Name      string
	Args      KmsGrantArgs
	state     *kmsGrantState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsGrant].
func (kg *KmsGrant) Type() string {
	return "aws_kms_grant"
}

// LocalName returns the local name for [KmsGrant].
func (kg *KmsGrant) LocalName() string {
	return kg.Name
}

// Configuration returns the configuration (args) for [KmsGrant].
func (kg *KmsGrant) Configuration() interface{} {
	return kg.Args
}

// DependOn is used for other resources to depend on [KmsGrant].
func (kg *KmsGrant) DependOn() terra.Reference {
	return terra.ReferenceResource(kg)
}

// Dependencies returns the list of resources [KmsGrant] depends_on.
func (kg *KmsGrant) Dependencies() terra.Dependencies {
	return kg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsGrant].
func (kg *KmsGrant) LifecycleManagement() *terra.Lifecycle {
	return kg.Lifecycle
}

// Attributes returns the attributes for [KmsGrant].
func (kg *KmsGrant) Attributes() kmsGrantAttributes {
	return kmsGrantAttributes{ref: terra.ReferenceResource(kg)}
}

// ImportState imports the given attribute values into [KmsGrant]'s state.
func (kg *KmsGrant) ImportState(av io.Reader) error {
	kg.state = &kmsGrantState{}
	if err := json.NewDecoder(av).Decode(kg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kg.Type(), kg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsGrant] has state.
func (kg *KmsGrant) State() (*kmsGrantState, bool) {
	return kg.state, kg.state != nil
}

// StateMust returns the state for [KmsGrant]. Panics if the state is nil.
func (kg *KmsGrant) StateMust() *kmsGrantState {
	if kg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kg.Type(), kg.LocalName()))
	}
	return kg.state
}

// KmsGrantArgs contains the configurations for aws_kms_grant.
type KmsGrantArgs struct {
	// GrantCreationTokens: set of string, optional
	GrantCreationTokens terra.SetValue[terra.StringValue] `hcl:"grant_creation_tokens,attr"`
	// GranteePrincipal: string, required
	GranteePrincipal terra.StringValue `hcl:"grantee_principal,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyId: string, required
	KeyId terra.StringValue `hcl:"key_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Operations: set of string, required
	Operations terra.SetValue[terra.StringValue] `hcl:"operations,attr" validate:"required"`
	// RetireOnDelete: bool, optional
	RetireOnDelete terra.BoolValue `hcl:"retire_on_delete,attr"`
	// RetiringPrincipal: string, optional
	RetiringPrincipal terra.StringValue `hcl:"retiring_principal,attr"`
	// Constraints: min=0
	Constraints []kmsgrant.Constraints `hcl:"constraints,block" validate:"min=0"`
}
type kmsGrantAttributes struct {
	ref terra.Reference
}

// GrantCreationTokens returns a reference to field grant_creation_tokens of aws_kms_grant.
func (kg kmsGrantAttributes) GrantCreationTokens() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kg.ref.Append("grant_creation_tokens"))
}

// GrantId returns a reference to field grant_id of aws_kms_grant.
func (kg kmsGrantAttributes) GrantId() terra.StringValue {
	return terra.ReferenceAsString(kg.ref.Append("grant_id"))
}

// GrantToken returns a reference to field grant_token of aws_kms_grant.
func (kg kmsGrantAttributes) GrantToken() terra.StringValue {
	return terra.ReferenceAsString(kg.ref.Append("grant_token"))
}

// GranteePrincipal returns a reference to field grantee_principal of aws_kms_grant.
func (kg kmsGrantAttributes) GranteePrincipal() terra.StringValue {
	return terra.ReferenceAsString(kg.ref.Append("grantee_principal"))
}

// Id returns a reference to field id of aws_kms_grant.
func (kg kmsGrantAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kg.ref.Append("id"))
}

// KeyId returns a reference to field key_id of aws_kms_grant.
func (kg kmsGrantAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(kg.ref.Append("key_id"))
}

// Name returns a reference to field name of aws_kms_grant.
func (kg kmsGrantAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kg.ref.Append("name"))
}

// Operations returns a reference to field operations of aws_kms_grant.
func (kg kmsGrantAttributes) Operations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kg.ref.Append("operations"))
}

// RetireOnDelete returns a reference to field retire_on_delete of aws_kms_grant.
func (kg kmsGrantAttributes) RetireOnDelete() terra.BoolValue {
	return terra.ReferenceAsBool(kg.ref.Append("retire_on_delete"))
}

// RetiringPrincipal returns a reference to field retiring_principal of aws_kms_grant.
func (kg kmsGrantAttributes) RetiringPrincipal() terra.StringValue {
	return terra.ReferenceAsString(kg.ref.Append("retiring_principal"))
}

func (kg kmsGrantAttributes) Constraints() terra.SetValue[kmsgrant.ConstraintsAttributes] {
	return terra.ReferenceAsSet[kmsgrant.ConstraintsAttributes](kg.ref.Append("constraints"))
}

type kmsGrantState struct {
	GrantCreationTokens []string                    `json:"grant_creation_tokens"`
	GrantId             string                      `json:"grant_id"`
	GrantToken          string                      `json:"grant_token"`
	GranteePrincipal    string                      `json:"grantee_principal"`
	Id                  string                      `json:"id"`
	KeyId               string                      `json:"key_id"`
	Name                string                      `json:"name"`
	Operations          []string                    `json:"operations"`
	RetireOnDelete      bool                        `json:"retire_on_delete"`
	RetiringPrincipal   string                      `json:"retiring_principal"`
	Constraints         []kmsgrant.ConstraintsState `json:"constraints"`
}

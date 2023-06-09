// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	iamdenypolicy "github.com/golingon/terraproviders/googlebeta/4.64.0/iamdenypolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIamDenyPolicy creates a new instance of [IamDenyPolicy].
func NewIamDenyPolicy(name string, args IamDenyPolicyArgs) *IamDenyPolicy {
	return &IamDenyPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamDenyPolicy)(nil)

// IamDenyPolicy represents the Terraform resource google_iam_deny_policy.
type IamDenyPolicy struct {
	Name      string
	Args      IamDenyPolicyArgs
	state     *iamDenyPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamDenyPolicy].
func (idp *IamDenyPolicy) Type() string {
	return "google_iam_deny_policy"
}

// LocalName returns the local name for [IamDenyPolicy].
func (idp *IamDenyPolicy) LocalName() string {
	return idp.Name
}

// Configuration returns the configuration (args) for [IamDenyPolicy].
func (idp *IamDenyPolicy) Configuration() interface{} {
	return idp.Args
}

// DependOn is used for other resources to depend on [IamDenyPolicy].
func (idp *IamDenyPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(idp)
}

// Dependencies returns the list of resources [IamDenyPolicy] depends_on.
func (idp *IamDenyPolicy) Dependencies() terra.Dependencies {
	return idp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamDenyPolicy].
func (idp *IamDenyPolicy) LifecycleManagement() *terra.Lifecycle {
	return idp.Lifecycle
}

// Attributes returns the attributes for [IamDenyPolicy].
func (idp *IamDenyPolicy) Attributes() iamDenyPolicyAttributes {
	return iamDenyPolicyAttributes{ref: terra.ReferenceResource(idp)}
}

// ImportState imports the given attribute values into [IamDenyPolicy]'s state.
func (idp *IamDenyPolicy) ImportState(av io.Reader) error {
	idp.state = &iamDenyPolicyState{}
	if err := json.NewDecoder(av).Decode(idp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", idp.Type(), idp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamDenyPolicy] has state.
func (idp *IamDenyPolicy) State() (*iamDenyPolicyState, bool) {
	return idp.state, idp.state != nil
}

// StateMust returns the state for [IamDenyPolicy]. Panics if the state is nil.
func (idp *IamDenyPolicy) StateMust() *iamDenyPolicyState {
	if idp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", idp.Type(), idp.LocalName()))
	}
	return idp.state
}

// IamDenyPolicyArgs contains the configurations for google_iam_deny_policy.
type IamDenyPolicyArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Rules: min=1
	Rules []iamdenypolicy.Rules `hcl:"rules,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *iamdenypolicy.Timeouts `hcl:"timeouts,block"`
}
type iamDenyPolicyAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_iam_deny_policy.
func (idp iamDenyPolicyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(idp.ref.Append("display_name"))
}

// Etag returns a reference to field etag of google_iam_deny_policy.
func (idp iamDenyPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(idp.ref.Append("etag"))
}

// Id returns a reference to field id of google_iam_deny_policy.
func (idp iamDenyPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(idp.ref.Append("id"))
}

// Name returns a reference to field name of google_iam_deny_policy.
func (idp iamDenyPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(idp.ref.Append("name"))
}

// Parent returns a reference to field parent of google_iam_deny_policy.
func (idp iamDenyPolicyAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(idp.ref.Append("parent"))
}

func (idp iamDenyPolicyAttributes) Rules() terra.ListValue[iamdenypolicy.RulesAttributes] {
	return terra.ReferenceAsList[iamdenypolicy.RulesAttributes](idp.ref.Append("rules"))
}

func (idp iamDenyPolicyAttributes) Timeouts() iamdenypolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iamdenypolicy.TimeoutsAttributes](idp.ref.Append("timeouts"))
}

type iamDenyPolicyState struct {
	DisplayName string                       `json:"display_name"`
	Etag        string                       `json:"etag"`
	Id          string                       `json:"id"`
	Name        string                       `json:"name"`
	Parent      string                       `json:"parent"`
	Rules       []iamdenypolicy.RulesState   `json:"rules"`
	Timeouts    *iamdenypolicy.TimeoutsState `json:"timeouts"`
}

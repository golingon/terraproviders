// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	orgpolicypolicy "github.com/golingon/terraproviders/googlebeta/4.75.1/orgpolicypolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrgPolicyPolicy creates a new instance of [OrgPolicyPolicy].
func NewOrgPolicyPolicy(name string, args OrgPolicyPolicyArgs) *OrgPolicyPolicy {
	return &OrgPolicyPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrgPolicyPolicy)(nil)

// OrgPolicyPolicy represents the Terraform resource google_org_policy_policy.
type OrgPolicyPolicy struct {
	Name      string
	Args      OrgPolicyPolicyArgs
	state     *orgPolicyPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrgPolicyPolicy].
func (opp *OrgPolicyPolicy) Type() string {
	return "google_org_policy_policy"
}

// LocalName returns the local name for [OrgPolicyPolicy].
func (opp *OrgPolicyPolicy) LocalName() string {
	return opp.Name
}

// Configuration returns the configuration (args) for [OrgPolicyPolicy].
func (opp *OrgPolicyPolicy) Configuration() interface{} {
	return opp.Args
}

// DependOn is used for other resources to depend on [OrgPolicyPolicy].
func (opp *OrgPolicyPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(opp)
}

// Dependencies returns the list of resources [OrgPolicyPolicy] depends_on.
func (opp *OrgPolicyPolicy) Dependencies() terra.Dependencies {
	return opp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrgPolicyPolicy].
func (opp *OrgPolicyPolicy) LifecycleManagement() *terra.Lifecycle {
	return opp.Lifecycle
}

// Attributes returns the attributes for [OrgPolicyPolicy].
func (opp *OrgPolicyPolicy) Attributes() orgPolicyPolicyAttributes {
	return orgPolicyPolicyAttributes{ref: terra.ReferenceResource(opp)}
}

// ImportState imports the given attribute values into [OrgPolicyPolicy]'s state.
func (opp *OrgPolicyPolicy) ImportState(av io.Reader) error {
	opp.state = &orgPolicyPolicyState{}
	if err := json.NewDecoder(av).Decode(opp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", opp.Type(), opp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrgPolicyPolicy] has state.
func (opp *OrgPolicyPolicy) State() (*orgPolicyPolicyState, bool) {
	return opp.state, opp.state != nil
}

// StateMust returns the state for [OrgPolicyPolicy]. Panics if the state is nil.
func (opp *OrgPolicyPolicy) StateMust() *orgPolicyPolicyState {
	if opp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", opp.Type(), opp.LocalName()))
	}
	return opp.state
}

// OrgPolicyPolicyArgs contains the configurations for google_org_policy_policy.
type OrgPolicyPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Spec: optional
	Spec *orgpolicypolicy.Spec `hcl:"spec,block"`
	// Timeouts: optional
	Timeouts *orgpolicypolicy.Timeouts `hcl:"timeouts,block"`
}
type orgPolicyPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_org_policy_policy.
func (opp orgPolicyPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(opp.ref.Append("id"))
}

// Name returns a reference to field name of google_org_policy_policy.
func (opp orgPolicyPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(opp.ref.Append("name"))
}

// Parent returns a reference to field parent of google_org_policy_policy.
func (opp orgPolicyPolicyAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(opp.ref.Append("parent"))
}

func (opp orgPolicyPolicyAttributes) Spec() terra.ListValue[orgpolicypolicy.SpecAttributes] {
	return terra.ReferenceAsList[orgpolicypolicy.SpecAttributes](opp.ref.Append("spec"))
}

func (opp orgPolicyPolicyAttributes) Timeouts() orgpolicypolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[orgpolicypolicy.TimeoutsAttributes](opp.ref.Append("timeouts"))
}

type orgPolicyPolicyState struct {
	Id       string                         `json:"id"`
	Name     string                         `json:"name"`
	Parent   string                         `json:"parent"`
	Spec     []orgpolicypolicy.SpecState    `json:"spec"`
	Timeouts *orgpolicypolicy.TimeoutsState `json:"timeouts"`
}

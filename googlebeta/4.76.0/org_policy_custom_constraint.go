// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	orgpolicycustomconstraint "github.com/golingon/terraproviders/googlebeta/4.76.0/orgpolicycustomconstraint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrgPolicyCustomConstraint creates a new instance of [OrgPolicyCustomConstraint].
func NewOrgPolicyCustomConstraint(name string, args OrgPolicyCustomConstraintArgs) *OrgPolicyCustomConstraint {
	return &OrgPolicyCustomConstraint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrgPolicyCustomConstraint)(nil)

// OrgPolicyCustomConstraint represents the Terraform resource google_org_policy_custom_constraint.
type OrgPolicyCustomConstraint struct {
	Name      string
	Args      OrgPolicyCustomConstraintArgs
	state     *orgPolicyCustomConstraintState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrgPolicyCustomConstraint].
func (opcc *OrgPolicyCustomConstraint) Type() string {
	return "google_org_policy_custom_constraint"
}

// LocalName returns the local name for [OrgPolicyCustomConstraint].
func (opcc *OrgPolicyCustomConstraint) LocalName() string {
	return opcc.Name
}

// Configuration returns the configuration (args) for [OrgPolicyCustomConstraint].
func (opcc *OrgPolicyCustomConstraint) Configuration() interface{} {
	return opcc.Args
}

// DependOn is used for other resources to depend on [OrgPolicyCustomConstraint].
func (opcc *OrgPolicyCustomConstraint) DependOn() terra.Reference {
	return terra.ReferenceResource(opcc)
}

// Dependencies returns the list of resources [OrgPolicyCustomConstraint] depends_on.
func (opcc *OrgPolicyCustomConstraint) Dependencies() terra.Dependencies {
	return opcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrgPolicyCustomConstraint].
func (opcc *OrgPolicyCustomConstraint) LifecycleManagement() *terra.Lifecycle {
	return opcc.Lifecycle
}

// Attributes returns the attributes for [OrgPolicyCustomConstraint].
func (opcc *OrgPolicyCustomConstraint) Attributes() orgPolicyCustomConstraintAttributes {
	return orgPolicyCustomConstraintAttributes{ref: terra.ReferenceResource(opcc)}
}

// ImportState imports the given attribute values into [OrgPolicyCustomConstraint]'s state.
func (opcc *OrgPolicyCustomConstraint) ImportState(av io.Reader) error {
	opcc.state = &orgPolicyCustomConstraintState{}
	if err := json.NewDecoder(av).Decode(opcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", opcc.Type(), opcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrgPolicyCustomConstraint] has state.
func (opcc *OrgPolicyCustomConstraint) State() (*orgPolicyCustomConstraintState, bool) {
	return opcc.state, opcc.state != nil
}

// StateMust returns the state for [OrgPolicyCustomConstraint]. Panics if the state is nil.
func (opcc *OrgPolicyCustomConstraint) StateMust() *orgPolicyCustomConstraintState {
	if opcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", opcc.Type(), opcc.LocalName()))
	}
	return opcc.state
}

// OrgPolicyCustomConstraintArgs contains the configurations for google_org_policy_custom_constraint.
type OrgPolicyCustomConstraintArgs struct {
	// ActionType: string, required
	ActionType terra.StringValue `hcl:"action_type,attr" validate:"required"`
	// Condition: string, required
	Condition terra.StringValue `hcl:"condition,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MethodTypes: list of string, required
	MethodTypes terra.ListValue[terra.StringValue] `hcl:"method_types,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// ResourceTypes: list of string, required
	ResourceTypes terra.ListValue[terra.StringValue] `hcl:"resource_types,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *orgpolicycustomconstraint.Timeouts `hcl:"timeouts,block"`
}
type orgPolicyCustomConstraintAttributes struct {
	ref terra.Reference
}

// ActionType returns a reference to field action_type of google_org_policy_custom_constraint.
func (opcc orgPolicyCustomConstraintAttributes) ActionType() terra.StringValue {
	return terra.ReferenceAsString(opcc.ref.Append("action_type"))
}

// Condition returns a reference to field condition of google_org_policy_custom_constraint.
func (opcc orgPolicyCustomConstraintAttributes) Condition() terra.StringValue {
	return terra.ReferenceAsString(opcc.ref.Append("condition"))
}

// Description returns a reference to field description of google_org_policy_custom_constraint.
func (opcc orgPolicyCustomConstraintAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(opcc.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_org_policy_custom_constraint.
func (opcc orgPolicyCustomConstraintAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(opcc.ref.Append("display_name"))
}

// Id returns a reference to field id of google_org_policy_custom_constraint.
func (opcc orgPolicyCustomConstraintAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(opcc.ref.Append("id"))
}

// MethodTypes returns a reference to field method_types of google_org_policy_custom_constraint.
func (opcc orgPolicyCustomConstraintAttributes) MethodTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](opcc.ref.Append("method_types"))
}

// Name returns a reference to field name of google_org_policy_custom_constraint.
func (opcc orgPolicyCustomConstraintAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(opcc.ref.Append("name"))
}

// Parent returns a reference to field parent of google_org_policy_custom_constraint.
func (opcc orgPolicyCustomConstraintAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(opcc.ref.Append("parent"))
}

// ResourceTypes returns a reference to field resource_types of google_org_policy_custom_constraint.
func (opcc orgPolicyCustomConstraintAttributes) ResourceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](opcc.ref.Append("resource_types"))
}

// UpdateTime returns a reference to field update_time of google_org_policy_custom_constraint.
func (opcc orgPolicyCustomConstraintAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(opcc.ref.Append("update_time"))
}

func (opcc orgPolicyCustomConstraintAttributes) Timeouts() orgpolicycustomconstraint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[orgpolicycustomconstraint.TimeoutsAttributes](opcc.ref.Append("timeouts"))
}

type orgPolicyCustomConstraintState struct {
	ActionType    string                                   `json:"action_type"`
	Condition     string                                   `json:"condition"`
	Description   string                                   `json:"description"`
	DisplayName   string                                   `json:"display_name"`
	Id            string                                   `json:"id"`
	MethodTypes   []string                                 `json:"method_types"`
	Name          string                                   `json:"name"`
	Parent        string                                   `json:"parent"`
	ResourceTypes []string                                 `json:"resource_types"`
	UpdateTime    string                                   `json:"update_time"`
	Timeouts      *orgpolicycustomconstraint.TimeoutsState `json:"timeouts"`
}

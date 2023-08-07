// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	organizationpolicy "github.com/golingon/terraproviders/google/4.76.0/organizationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrganizationPolicy creates a new instance of [OrganizationPolicy].
func NewOrganizationPolicy(name string, args OrganizationPolicyArgs) *OrganizationPolicy {
	return &OrganizationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrganizationPolicy)(nil)

// OrganizationPolicy represents the Terraform resource google_organization_policy.
type OrganizationPolicy struct {
	Name      string
	Args      OrganizationPolicyArgs
	state     *organizationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrganizationPolicy].
func (op *OrganizationPolicy) Type() string {
	return "google_organization_policy"
}

// LocalName returns the local name for [OrganizationPolicy].
func (op *OrganizationPolicy) LocalName() string {
	return op.Name
}

// Configuration returns the configuration (args) for [OrganizationPolicy].
func (op *OrganizationPolicy) Configuration() interface{} {
	return op.Args
}

// DependOn is used for other resources to depend on [OrganizationPolicy].
func (op *OrganizationPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(op)
}

// Dependencies returns the list of resources [OrganizationPolicy] depends_on.
func (op *OrganizationPolicy) Dependencies() terra.Dependencies {
	return op.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrganizationPolicy].
func (op *OrganizationPolicy) LifecycleManagement() *terra.Lifecycle {
	return op.Lifecycle
}

// Attributes returns the attributes for [OrganizationPolicy].
func (op *OrganizationPolicy) Attributes() organizationPolicyAttributes {
	return organizationPolicyAttributes{ref: terra.ReferenceResource(op)}
}

// ImportState imports the given attribute values into [OrganizationPolicy]'s state.
func (op *OrganizationPolicy) ImportState(av io.Reader) error {
	op.state = &organizationPolicyState{}
	if err := json.NewDecoder(av).Decode(op.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", op.Type(), op.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrganizationPolicy] has state.
func (op *OrganizationPolicy) State() (*organizationPolicyState, bool) {
	return op.state, op.state != nil
}

// StateMust returns the state for [OrganizationPolicy]. Panics if the state is nil.
func (op *OrganizationPolicy) StateMust() *organizationPolicyState {
	if op.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", op.Type(), op.LocalName()))
	}
	return op.state
}

// OrganizationPolicyArgs contains the configurations for google_organization_policy.
type OrganizationPolicyArgs struct {
	// Constraint: string, required
	Constraint terra.StringValue `hcl:"constraint,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Version: number, optional
	Version terra.NumberValue `hcl:"version,attr"`
	// BooleanPolicy: optional
	BooleanPolicy *organizationpolicy.BooleanPolicy `hcl:"boolean_policy,block"`
	// ListPolicy: optional
	ListPolicy *organizationpolicy.ListPolicy `hcl:"list_policy,block"`
	// RestorePolicy: optional
	RestorePolicy *organizationpolicy.RestorePolicy `hcl:"restore_policy,block"`
	// Timeouts: optional
	Timeouts *organizationpolicy.Timeouts `hcl:"timeouts,block"`
}
type organizationPolicyAttributes struct {
	ref terra.Reference
}

// Constraint returns a reference to field constraint of google_organization_policy.
func (op organizationPolicyAttributes) Constraint() terra.StringValue {
	return terra.ReferenceAsString(op.ref.Append("constraint"))
}

// Etag returns a reference to field etag of google_organization_policy.
func (op organizationPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(op.ref.Append("etag"))
}

// Id returns a reference to field id of google_organization_policy.
func (op organizationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(op.ref.Append("id"))
}

// OrgId returns a reference to field org_id of google_organization_policy.
func (op organizationPolicyAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(op.ref.Append("org_id"))
}

// UpdateTime returns a reference to field update_time of google_organization_policy.
func (op organizationPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(op.ref.Append("update_time"))
}

// Version returns a reference to field version of google_organization_policy.
func (op organizationPolicyAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(op.ref.Append("version"))
}

func (op organizationPolicyAttributes) BooleanPolicy() terra.ListValue[organizationpolicy.BooleanPolicyAttributes] {
	return terra.ReferenceAsList[organizationpolicy.BooleanPolicyAttributes](op.ref.Append("boolean_policy"))
}

func (op organizationPolicyAttributes) ListPolicy() terra.ListValue[organizationpolicy.ListPolicyAttributes] {
	return terra.ReferenceAsList[organizationpolicy.ListPolicyAttributes](op.ref.Append("list_policy"))
}

func (op organizationPolicyAttributes) RestorePolicy() terra.ListValue[organizationpolicy.RestorePolicyAttributes] {
	return terra.ReferenceAsList[organizationpolicy.RestorePolicyAttributes](op.ref.Append("restore_policy"))
}

func (op organizationPolicyAttributes) Timeouts() organizationpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[organizationpolicy.TimeoutsAttributes](op.ref.Append("timeouts"))
}

type organizationPolicyState struct {
	Constraint    string                                  `json:"constraint"`
	Etag          string                                  `json:"etag"`
	Id            string                                  `json:"id"`
	OrgId         string                                  `json:"org_id"`
	UpdateTime    string                                  `json:"update_time"`
	Version       float64                                 `json:"version"`
	BooleanPolicy []organizationpolicy.BooleanPolicyState `json:"boolean_policy"`
	ListPolicy    []organizationpolicy.ListPolicyState    `json:"list_policy"`
	RestorePolicy []organizationpolicy.RestorePolicyState `json:"restore_policy"`
	Timeouts      *organizationpolicy.TimeoutsState       `json:"timeouts"`
}

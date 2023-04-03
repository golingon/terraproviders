// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iamaccessboundarypolicy "github.com/golingon/terraproviders/google/4.59.0/iamaccessboundarypolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIamAccessBoundaryPolicy creates a new instance of [IamAccessBoundaryPolicy].
func NewIamAccessBoundaryPolicy(name string, args IamAccessBoundaryPolicyArgs) *IamAccessBoundaryPolicy {
	return &IamAccessBoundaryPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamAccessBoundaryPolicy)(nil)

// IamAccessBoundaryPolicy represents the Terraform resource google_iam_access_boundary_policy.
type IamAccessBoundaryPolicy struct {
	Name      string
	Args      IamAccessBoundaryPolicyArgs
	state     *iamAccessBoundaryPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamAccessBoundaryPolicy].
func (iabp *IamAccessBoundaryPolicy) Type() string {
	return "google_iam_access_boundary_policy"
}

// LocalName returns the local name for [IamAccessBoundaryPolicy].
func (iabp *IamAccessBoundaryPolicy) LocalName() string {
	return iabp.Name
}

// Configuration returns the configuration (args) for [IamAccessBoundaryPolicy].
func (iabp *IamAccessBoundaryPolicy) Configuration() interface{} {
	return iabp.Args
}

// DependOn is used for other resources to depend on [IamAccessBoundaryPolicy].
func (iabp *IamAccessBoundaryPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(iabp)
}

// Dependencies returns the list of resources [IamAccessBoundaryPolicy] depends_on.
func (iabp *IamAccessBoundaryPolicy) Dependencies() terra.Dependencies {
	return iabp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamAccessBoundaryPolicy].
func (iabp *IamAccessBoundaryPolicy) LifecycleManagement() *terra.Lifecycle {
	return iabp.Lifecycle
}

// Attributes returns the attributes for [IamAccessBoundaryPolicy].
func (iabp *IamAccessBoundaryPolicy) Attributes() iamAccessBoundaryPolicyAttributes {
	return iamAccessBoundaryPolicyAttributes{ref: terra.ReferenceResource(iabp)}
}

// ImportState imports the given attribute values into [IamAccessBoundaryPolicy]'s state.
func (iabp *IamAccessBoundaryPolicy) ImportState(av io.Reader) error {
	iabp.state = &iamAccessBoundaryPolicyState{}
	if err := json.NewDecoder(av).Decode(iabp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iabp.Type(), iabp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamAccessBoundaryPolicy] has state.
func (iabp *IamAccessBoundaryPolicy) State() (*iamAccessBoundaryPolicyState, bool) {
	return iabp.state, iabp.state != nil
}

// StateMust returns the state for [IamAccessBoundaryPolicy]. Panics if the state is nil.
func (iabp *IamAccessBoundaryPolicy) StateMust() *iamAccessBoundaryPolicyState {
	if iabp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iabp.Type(), iabp.LocalName()))
	}
	return iabp.state
}

// IamAccessBoundaryPolicyArgs contains the configurations for google_iam_access_boundary_policy.
type IamAccessBoundaryPolicyArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Rules: min=1
	Rules []iamaccessboundarypolicy.Rules `hcl:"rules,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *iamaccessboundarypolicy.Timeouts `hcl:"timeouts,block"`
}
type iamAccessBoundaryPolicyAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_iam_access_boundary_policy.
func (iabp iamAccessBoundaryPolicyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(iabp.ref.Append("display_name"))
}

// Etag returns a reference to field etag of google_iam_access_boundary_policy.
func (iabp iamAccessBoundaryPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iabp.ref.Append("etag"))
}

// Id returns a reference to field id of google_iam_access_boundary_policy.
func (iabp iamAccessBoundaryPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iabp.ref.Append("id"))
}

// Name returns a reference to field name of google_iam_access_boundary_policy.
func (iabp iamAccessBoundaryPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iabp.ref.Append("name"))
}

// Parent returns a reference to field parent of google_iam_access_boundary_policy.
func (iabp iamAccessBoundaryPolicyAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(iabp.ref.Append("parent"))
}

func (iabp iamAccessBoundaryPolicyAttributes) Rules() terra.ListValue[iamaccessboundarypolicy.RulesAttributes] {
	return terra.ReferenceAsList[iamaccessboundarypolicy.RulesAttributes](iabp.ref.Append("rules"))
}

func (iabp iamAccessBoundaryPolicyAttributes) Timeouts() iamaccessboundarypolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iamaccessboundarypolicy.TimeoutsAttributes](iabp.ref.Append("timeouts"))
}

type iamAccessBoundaryPolicyState struct {
	DisplayName string                                 `json:"display_name"`
	Etag        string                                 `json:"etag"`
	Id          string                                 `json:"id"`
	Name        string                                 `json:"name"`
	Parent      string                                 `json:"parent"`
	Rules       []iamaccessboundarypolicy.RulesState   `json:"rules"`
	Timeouts    *iamaccessboundarypolicy.TimeoutsState `json:"timeouts"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googleworkspace

import (
	"encoding/json"
	"fmt"
	chromepolicy "github.com/golingon/terraproviders/googleworkspace/0.7.0/chromepolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChromePolicy creates a new instance of [ChromePolicy].
func NewChromePolicy(name string, args ChromePolicyArgs) *ChromePolicy {
	return &ChromePolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChromePolicy)(nil)

// ChromePolicy represents the Terraform resource googleworkspace_chrome_policy.
type ChromePolicy struct {
	Name      string
	Args      ChromePolicyArgs
	state     *chromePolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChromePolicy].
func (cp *ChromePolicy) Type() string {
	return "googleworkspace_chrome_policy"
}

// LocalName returns the local name for [ChromePolicy].
func (cp *ChromePolicy) LocalName() string {
	return cp.Name
}

// Configuration returns the configuration (args) for [ChromePolicy].
func (cp *ChromePolicy) Configuration() interface{} {
	return cp.Args
}

// DependOn is used for other resources to depend on [ChromePolicy].
func (cp *ChromePolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cp)
}

// Dependencies returns the list of resources [ChromePolicy] depends_on.
func (cp *ChromePolicy) Dependencies() terra.Dependencies {
	return cp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChromePolicy].
func (cp *ChromePolicy) LifecycleManagement() *terra.Lifecycle {
	return cp.Lifecycle
}

// Attributes returns the attributes for [ChromePolicy].
func (cp *ChromePolicy) Attributes() chromePolicyAttributes {
	return chromePolicyAttributes{ref: terra.ReferenceResource(cp)}
}

// ImportState imports the given attribute values into [ChromePolicy]'s state.
func (cp *ChromePolicy) ImportState(av io.Reader) error {
	cp.state = &chromePolicyState{}
	if err := json.NewDecoder(av).Decode(cp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cp.Type(), cp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChromePolicy] has state.
func (cp *ChromePolicy) State() (*chromePolicyState, bool) {
	return cp.state, cp.state != nil
}

// StateMust returns the state for [ChromePolicy]. Panics if the state is nil.
func (cp *ChromePolicy) StateMust() *chromePolicyState {
	if cp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cp.Type(), cp.LocalName()))
	}
	return cp.state
}

// ChromePolicyArgs contains the configurations for googleworkspace_chrome_policy.
type ChromePolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrgUnitId: string, required
	OrgUnitId terra.StringValue `hcl:"org_unit_id,attr" validate:"required"`
	// Policies: min=1
	Policies []chromepolicy.Policies `hcl:"policies,block" validate:"min=1"`
}
type chromePolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of googleworkspace_chrome_policy.
func (cp chromePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("id"))
}

// OrgUnitId returns a reference to field org_unit_id of googleworkspace_chrome_policy.
func (cp chromePolicyAttributes) OrgUnitId() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("org_unit_id"))
}

func (cp chromePolicyAttributes) Policies() terra.ListValue[chromepolicy.PoliciesAttributes] {
	return terra.ReferenceAsList[chromepolicy.PoliciesAttributes](cp.ref.Append("policies"))
}

type chromePolicyState struct {
	Id        string                       `json:"id"`
	OrgUnitId string                       `json:"org_unit_id"`
	Policies  []chromepolicy.PoliciesState `json:"policies"`
}

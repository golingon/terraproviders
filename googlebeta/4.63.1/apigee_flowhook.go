// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigeeflowhook "github.com/golingon/terraproviders/googlebeta/4.63.1/apigeeflowhook"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeFlowhook creates a new instance of [ApigeeFlowhook].
func NewApigeeFlowhook(name string, args ApigeeFlowhookArgs) *ApigeeFlowhook {
	return &ApigeeFlowhook{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeFlowhook)(nil)

// ApigeeFlowhook represents the Terraform resource google_apigee_flowhook.
type ApigeeFlowhook struct {
	Name      string
	Args      ApigeeFlowhookArgs
	state     *apigeeFlowhookState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeFlowhook].
func (af *ApigeeFlowhook) Type() string {
	return "google_apigee_flowhook"
}

// LocalName returns the local name for [ApigeeFlowhook].
func (af *ApigeeFlowhook) LocalName() string {
	return af.Name
}

// Configuration returns the configuration (args) for [ApigeeFlowhook].
func (af *ApigeeFlowhook) Configuration() interface{} {
	return af.Args
}

// DependOn is used for other resources to depend on [ApigeeFlowhook].
func (af *ApigeeFlowhook) DependOn() terra.Reference {
	return terra.ReferenceResource(af)
}

// Dependencies returns the list of resources [ApigeeFlowhook] depends_on.
func (af *ApigeeFlowhook) Dependencies() terra.Dependencies {
	return af.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeFlowhook].
func (af *ApigeeFlowhook) LifecycleManagement() *terra.Lifecycle {
	return af.Lifecycle
}

// Attributes returns the attributes for [ApigeeFlowhook].
func (af *ApigeeFlowhook) Attributes() apigeeFlowhookAttributes {
	return apigeeFlowhookAttributes{ref: terra.ReferenceResource(af)}
}

// ImportState imports the given attribute values into [ApigeeFlowhook]'s state.
func (af *ApigeeFlowhook) ImportState(av io.Reader) error {
	af.state = &apigeeFlowhookState{}
	if err := json.NewDecoder(av).Decode(af.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", af.Type(), af.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeFlowhook] has state.
func (af *ApigeeFlowhook) State() (*apigeeFlowhookState, bool) {
	return af.state, af.state != nil
}

// StateMust returns the state for [ApigeeFlowhook]. Panics if the state is nil.
func (af *ApigeeFlowhook) StateMust() *apigeeFlowhookState {
	if af.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", af.Type(), af.LocalName()))
	}
	return af.state
}

// ApigeeFlowhookArgs contains the configurations for google_apigee_flowhook.
type ApigeeFlowhookArgs struct {
	// ContinueOnError: bool, optional
	ContinueOnError terra.BoolValue `hcl:"continue_on_error,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Environment: string, required
	Environment terra.StringValue `hcl:"environment,attr" validate:"required"`
	// FlowHookPoint: string, required
	FlowHookPoint terra.StringValue `hcl:"flow_hook_point,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Sharedflow: string, required
	Sharedflow terra.StringValue `hcl:"sharedflow,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apigeeflowhook.Timeouts `hcl:"timeouts,block"`
}
type apigeeFlowhookAttributes struct {
	ref terra.Reference
}

// ContinueOnError returns a reference to field continue_on_error of google_apigee_flowhook.
func (af apigeeFlowhookAttributes) ContinueOnError() terra.BoolValue {
	return terra.ReferenceAsBool(af.ref.Append("continue_on_error"))
}

// Description returns a reference to field description of google_apigee_flowhook.
func (af apigeeFlowhookAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("description"))
}

// Environment returns a reference to field environment of google_apigee_flowhook.
func (af apigeeFlowhookAttributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("environment"))
}

// FlowHookPoint returns a reference to field flow_hook_point of google_apigee_flowhook.
func (af apigeeFlowhookAttributes) FlowHookPoint() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("flow_hook_point"))
}

// Id returns a reference to field id of google_apigee_flowhook.
func (af apigeeFlowhookAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("id"))
}

// OrgId returns a reference to field org_id of google_apigee_flowhook.
func (af apigeeFlowhookAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("org_id"))
}

// Sharedflow returns a reference to field sharedflow of google_apigee_flowhook.
func (af apigeeFlowhookAttributes) Sharedflow() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("sharedflow"))
}

func (af apigeeFlowhookAttributes) Timeouts() apigeeflowhook.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeeflowhook.TimeoutsAttributes](af.ref.Append("timeouts"))
}

type apigeeFlowhookState struct {
	ContinueOnError bool                          `json:"continue_on_error"`
	Description     string                        `json:"description"`
	Environment     string                        `json:"environment"`
	FlowHookPoint   string                        `json:"flow_hook_point"`
	Id              string                        `json:"id"`
	OrgId           string                        `json:"org_id"`
	Sharedflow      string                        `json:"sharedflow"`
	Timeouts        *apigeeflowhook.TimeoutsState `json:"timeouts"`
}

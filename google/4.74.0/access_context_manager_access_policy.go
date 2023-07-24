// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanageraccesspolicy "github.com/golingon/terraproviders/google/4.74.0/accesscontextmanageraccesspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerAccessPolicy creates a new instance of [AccessContextManagerAccessPolicy].
func NewAccessContextManagerAccessPolicy(name string, args AccessContextManagerAccessPolicyArgs) *AccessContextManagerAccessPolicy {
	return &AccessContextManagerAccessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerAccessPolicy)(nil)

// AccessContextManagerAccessPolicy represents the Terraform resource google_access_context_manager_access_policy.
type AccessContextManagerAccessPolicy struct {
	Name      string
	Args      AccessContextManagerAccessPolicyArgs
	state     *accessContextManagerAccessPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerAccessPolicy].
func (acmap *AccessContextManagerAccessPolicy) Type() string {
	return "google_access_context_manager_access_policy"
}

// LocalName returns the local name for [AccessContextManagerAccessPolicy].
func (acmap *AccessContextManagerAccessPolicy) LocalName() string {
	return acmap.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerAccessPolicy].
func (acmap *AccessContextManagerAccessPolicy) Configuration() interface{} {
	return acmap.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerAccessPolicy].
func (acmap *AccessContextManagerAccessPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(acmap)
}

// Dependencies returns the list of resources [AccessContextManagerAccessPolicy] depends_on.
func (acmap *AccessContextManagerAccessPolicy) Dependencies() terra.Dependencies {
	return acmap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerAccessPolicy].
func (acmap *AccessContextManagerAccessPolicy) LifecycleManagement() *terra.Lifecycle {
	return acmap.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerAccessPolicy].
func (acmap *AccessContextManagerAccessPolicy) Attributes() accessContextManagerAccessPolicyAttributes {
	return accessContextManagerAccessPolicyAttributes{ref: terra.ReferenceResource(acmap)}
}

// ImportState imports the given attribute values into [AccessContextManagerAccessPolicy]'s state.
func (acmap *AccessContextManagerAccessPolicy) ImportState(av io.Reader) error {
	acmap.state = &accessContextManagerAccessPolicyState{}
	if err := json.NewDecoder(av).Decode(acmap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmap.Type(), acmap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerAccessPolicy] has state.
func (acmap *AccessContextManagerAccessPolicy) State() (*accessContextManagerAccessPolicyState, bool) {
	return acmap.state, acmap.state != nil
}

// StateMust returns the state for [AccessContextManagerAccessPolicy]. Panics if the state is nil.
func (acmap *AccessContextManagerAccessPolicy) StateMust() *accessContextManagerAccessPolicyState {
	if acmap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmap.Type(), acmap.LocalName()))
	}
	return acmap.state
}

// AccessContextManagerAccessPolicyArgs contains the configurations for google_access_context_manager_access_policy.
type AccessContextManagerAccessPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Scopes: list of string, optional
	Scopes terra.ListValue[terra.StringValue] `hcl:"scopes,attr"`
	// Title: string, required
	Title terra.StringValue `hcl:"title,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *accesscontextmanageraccesspolicy.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerAccessPolicyAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_access_context_manager_access_policy.
func (acmap accessContextManagerAccessPolicyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(acmap.ref.Append("create_time"))
}

// Id returns a reference to field id of google_access_context_manager_access_policy.
func (acmap accessContextManagerAccessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmap.ref.Append("id"))
}

// Name returns a reference to field name of google_access_context_manager_access_policy.
func (acmap accessContextManagerAccessPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acmap.ref.Append("name"))
}

// Parent returns a reference to field parent of google_access_context_manager_access_policy.
func (acmap accessContextManagerAccessPolicyAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(acmap.ref.Append("parent"))
}

// Scopes returns a reference to field scopes of google_access_context_manager_access_policy.
func (acmap accessContextManagerAccessPolicyAttributes) Scopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](acmap.ref.Append("scopes"))
}

// Title returns a reference to field title of google_access_context_manager_access_policy.
func (acmap accessContextManagerAccessPolicyAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(acmap.ref.Append("title"))
}

// UpdateTime returns a reference to field update_time of google_access_context_manager_access_policy.
func (acmap accessContextManagerAccessPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(acmap.ref.Append("update_time"))
}

func (acmap accessContextManagerAccessPolicyAttributes) Timeouts() accesscontextmanageraccesspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanageraccesspolicy.TimeoutsAttributes](acmap.ref.Append("timeouts"))
}

type accessContextManagerAccessPolicyState struct {
	CreateTime string                                          `json:"create_time"`
	Id         string                                          `json:"id"`
	Name       string                                          `json:"name"`
	Parent     string                                          `json:"parent"`
	Scopes     []string                                        `json:"scopes"`
	Title      string                                          `json:"title"`
	UpdateTime string                                          `json:"update_time"`
	Timeouts   *accesscontextmanageraccesspolicy.TimeoutsState `json:"timeouts"`
}

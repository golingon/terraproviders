// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanageraccesslevels "github.com/golingon/terraproviders/google/4.71.0/accesscontextmanageraccesslevels"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerAccessLevels creates a new instance of [AccessContextManagerAccessLevels].
func NewAccessContextManagerAccessLevels(name string, args AccessContextManagerAccessLevelsArgs) *AccessContextManagerAccessLevels {
	return &AccessContextManagerAccessLevels{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerAccessLevels)(nil)

// AccessContextManagerAccessLevels represents the Terraform resource google_access_context_manager_access_levels.
type AccessContextManagerAccessLevels struct {
	Name      string
	Args      AccessContextManagerAccessLevelsArgs
	state     *accessContextManagerAccessLevelsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerAccessLevels].
func (acmal *AccessContextManagerAccessLevels) Type() string {
	return "google_access_context_manager_access_levels"
}

// LocalName returns the local name for [AccessContextManagerAccessLevels].
func (acmal *AccessContextManagerAccessLevels) LocalName() string {
	return acmal.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerAccessLevels].
func (acmal *AccessContextManagerAccessLevels) Configuration() interface{} {
	return acmal.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerAccessLevels].
func (acmal *AccessContextManagerAccessLevels) DependOn() terra.Reference {
	return terra.ReferenceResource(acmal)
}

// Dependencies returns the list of resources [AccessContextManagerAccessLevels] depends_on.
func (acmal *AccessContextManagerAccessLevels) Dependencies() terra.Dependencies {
	return acmal.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerAccessLevels].
func (acmal *AccessContextManagerAccessLevels) LifecycleManagement() *terra.Lifecycle {
	return acmal.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerAccessLevels].
func (acmal *AccessContextManagerAccessLevels) Attributes() accessContextManagerAccessLevelsAttributes {
	return accessContextManagerAccessLevelsAttributes{ref: terra.ReferenceResource(acmal)}
}

// ImportState imports the given attribute values into [AccessContextManagerAccessLevels]'s state.
func (acmal *AccessContextManagerAccessLevels) ImportState(av io.Reader) error {
	acmal.state = &accessContextManagerAccessLevelsState{}
	if err := json.NewDecoder(av).Decode(acmal.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmal.Type(), acmal.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerAccessLevels] has state.
func (acmal *AccessContextManagerAccessLevels) State() (*accessContextManagerAccessLevelsState, bool) {
	return acmal.state, acmal.state != nil
}

// StateMust returns the state for [AccessContextManagerAccessLevels]. Panics if the state is nil.
func (acmal *AccessContextManagerAccessLevels) StateMust() *accessContextManagerAccessLevelsState {
	if acmal.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmal.Type(), acmal.LocalName()))
	}
	return acmal.state
}

// AccessContextManagerAccessLevelsArgs contains the configurations for google_access_context_manager_access_levels.
type AccessContextManagerAccessLevelsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// AccessLevels: min=0
	AccessLevels []accesscontextmanageraccesslevels.AccessLevels `hcl:"access_levels,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *accesscontextmanageraccesslevels.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerAccessLevelsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_access_context_manager_access_levels.
func (acmal accessContextManagerAccessLevelsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmal.ref.Append("id"))
}

// Parent returns a reference to field parent of google_access_context_manager_access_levels.
func (acmal accessContextManagerAccessLevelsAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(acmal.ref.Append("parent"))
}

func (acmal accessContextManagerAccessLevelsAttributes) AccessLevels() terra.SetValue[accesscontextmanageraccesslevels.AccessLevelsAttributes] {
	return terra.ReferenceAsSet[accesscontextmanageraccesslevels.AccessLevelsAttributes](acmal.ref.Append("access_levels"))
}

func (acmal accessContextManagerAccessLevelsAttributes) Timeouts() accesscontextmanageraccesslevels.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanageraccesslevels.TimeoutsAttributes](acmal.ref.Append("timeouts"))
}

type accessContextManagerAccessLevelsState struct {
	Id           string                                               `json:"id"`
	Parent       string                                               `json:"parent"`
	AccessLevels []accesscontextmanageraccesslevels.AccessLevelsState `json:"access_levels"`
	Timeouts     *accesscontextmanageraccesslevels.TimeoutsState      `json:"timeouts"`
}

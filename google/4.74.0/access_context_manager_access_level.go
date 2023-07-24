// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanageraccesslevel "github.com/golingon/terraproviders/google/4.74.0/accesscontextmanageraccesslevel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerAccessLevel creates a new instance of [AccessContextManagerAccessLevel].
func NewAccessContextManagerAccessLevel(name string, args AccessContextManagerAccessLevelArgs) *AccessContextManagerAccessLevel {
	return &AccessContextManagerAccessLevel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerAccessLevel)(nil)

// AccessContextManagerAccessLevel represents the Terraform resource google_access_context_manager_access_level.
type AccessContextManagerAccessLevel struct {
	Name      string
	Args      AccessContextManagerAccessLevelArgs
	state     *accessContextManagerAccessLevelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerAccessLevel].
func (acmal *AccessContextManagerAccessLevel) Type() string {
	return "google_access_context_manager_access_level"
}

// LocalName returns the local name for [AccessContextManagerAccessLevel].
func (acmal *AccessContextManagerAccessLevel) LocalName() string {
	return acmal.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerAccessLevel].
func (acmal *AccessContextManagerAccessLevel) Configuration() interface{} {
	return acmal.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerAccessLevel].
func (acmal *AccessContextManagerAccessLevel) DependOn() terra.Reference {
	return terra.ReferenceResource(acmal)
}

// Dependencies returns the list of resources [AccessContextManagerAccessLevel] depends_on.
func (acmal *AccessContextManagerAccessLevel) Dependencies() terra.Dependencies {
	return acmal.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerAccessLevel].
func (acmal *AccessContextManagerAccessLevel) LifecycleManagement() *terra.Lifecycle {
	return acmal.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerAccessLevel].
func (acmal *AccessContextManagerAccessLevel) Attributes() accessContextManagerAccessLevelAttributes {
	return accessContextManagerAccessLevelAttributes{ref: terra.ReferenceResource(acmal)}
}

// ImportState imports the given attribute values into [AccessContextManagerAccessLevel]'s state.
func (acmal *AccessContextManagerAccessLevel) ImportState(av io.Reader) error {
	acmal.state = &accessContextManagerAccessLevelState{}
	if err := json.NewDecoder(av).Decode(acmal.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmal.Type(), acmal.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerAccessLevel] has state.
func (acmal *AccessContextManagerAccessLevel) State() (*accessContextManagerAccessLevelState, bool) {
	return acmal.state, acmal.state != nil
}

// StateMust returns the state for [AccessContextManagerAccessLevel]. Panics if the state is nil.
func (acmal *AccessContextManagerAccessLevel) StateMust() *accessContextManagerAccessLevelState {
	if acmal.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmal.Type(), acmal.LocalName()))
	}
	return acmal.state
}

// AccessContextManagerAccessLevelArgs contains the configurations for google_access_context_manager_access_level.
type AccessContextManagerAccessLevelArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Title: string, required
	Title terra.StringValue `hcl:"title,attr" validate:"required"`
	// Basic: optional
	Basic *accesscontextmanageraccesslevel.Basic `hcl:"basic,block"`
	// Custom: optional
	Custom *accesscontextmanageraccesslevel.Custom `hcl:"custom,block"`
	// Timeouts: optional
	Timeouts *accesscontextmanageraccesslevel.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerAccessLevelAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_access_context_manager_access_level.
func (acmal accessContextManagerAccessLevelAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(acmal.ref.Append("description"))
}

// Id returns a reference to field id of google_access_context_manager_access_level.
func (acmal accessContextManagerAccessLevelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmal.ref.Append("id"))
}

// Name returns a reference to field name of google_access_context_manager_access_level.
func (acmal accessContextManagerAccessLevelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acmal.ref.Append("name"))
}

// Parent returns a reference to field parent of google_access_context_manager_access_level.
func (acmal accessContextManagerAccessLevelAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(acmal.ref.Append("parent"))
}

// Title returns a reference to field title of google_access_context_manager_access_level.
func (acmal accessContextManagerAccessLevelAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(acmal.ref.Append("title"))
}

func (acmal accessContextManagerAccessLevelAttributes) Basic() terra.ListValue[accesscontextmanageraccesslevel.BasicAttributes] {
	return terra.ReferenceAsList[accesscontextmanageraccesslevel.BasicAttributes](acmal.ref.Append("basic"))
}

func (acmal accessContextManagerAccessLevelAttributes) Custom() terra.ListValue[accesscontextmanageraccesslevel.CustomAttributes] {
	return terra.ReferenceAsList[accesscontextmanageraccesslevel.CustomAttributes](acmal.ref.Append("custom"))
}

func (acmal accessContextManagerAccessLevelAttributes) Timeouts() accesscontextmanageraccesslevel.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanageraccesslevel.TimeoutsAttributes](acmal.ref.Append("timeouts"))
}

type accessContextManagerAccessLevelState struct {
	Description string                                         `json:"description"`
	Id          string                                         `json:"id"`
	Name        string                                         `json:"name"`
	Parent      string                                         `json:"parent"`
	Title       string                                         `json:"title"`
	Basic       []accesscontextmanageraccesslevel.BasicState   `json:"basic"`
	Custom      []accesscontextmanageraccesslevel.CustomState  `json:"custom"`
	Timeouts    *accesscontextmanageraccesslevel.TimeoutsState `json:"timeouts"`
}

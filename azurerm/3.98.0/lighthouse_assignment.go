// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	lighthouseassignment "github.com/golingon/terraproviders/azurerm/3.98.0/lighthouseassignment"
	"io"
)

// NewLighthouseAssignment creates a new instance of [LighthouseAssignment].
func NewLighthouseAssignment(name string, args LighthouseAssignmentArgs) *LighthouseAssignment {
	return &LighthouseAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LighthouseAssignment)(nil)

// LighthouseAssignment represents the Terraform resource azurerm_lighthouse_assignment.
type LighthouseAssignment struct {
	Name      string
	Args      LighthouseAssignmentArgs
	state     *lighthouseAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LighthouseAssignment].
func (la *LighthouseAssignment) Type() string {
	return "azurerm_lighthouse_assignment"
}

// LocalName returns the local name for [LighthouseAssignment].
func (la *LighthouseAssignment) LocalName() string {
	return la.Name
}

// Configuration returns the configuration (args) for [LighthouseAssignment].
func (la *LighthouseAssignment) Configuration() interface{} {
	return la.Args
}

// DependOn is used for other resources to depend on [LighthouseAssignment].
func (la *LighthouseAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(la)
}

// Dependencies returns the list of resources [LighthouseAssignment] depends_on.
func (la *LighthouseAssignment) Dependencies() terra.Dependencies {
	return la.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LighthouseAssignment].
func (la *LighthouseAssignment) LifecycleManagement() *terra.Lifecycle {
	return la.Lifecycle
}

// Attributes returns the attributes for [LighthouseAssignment].
func (la *LighthouseAssignment) Attributes() lighthouseAssignmentAttributes {
	return lighthouseAssignmentAttributes{ref: terra.ReferenceResource(la)}
}

// ImportState imports the given attribute values into [LighthouseAssignment]'s state.
func (la *LighthouseAssignment) ImportState(av io.Reader) error {
	la.state = &lighthouseAssignmentState{}
	if err := json.NewDecoder(av).Decode(la.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", la.Type(), la.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LighthouseAssignment] has state.
func (la *LighthouseAssignment) State() (*lighthouseAssignmentState, bool) {
	return la.state, la.state != nil
}

// StateMust returns the state for [LighthouseAssignment]. Panics if the state is nil.
func (la *LighthouseAssignment) StateMust() *lighthouseAssignmentState {
	if la.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", la.Type(), la.LocalName()))
	}
	return la.state
}

// LighthouseAssignmentArgs contains the configurations for azurerm_lighthouse_assignment.
type LighthouseAssignmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LighthouseDefinitionId: string, required
	LighthouseDefinitionId terra.StringValue `hcl:"lighthouse_definition_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *lighthouseassignment.Timeouts `hcl:"timeouts,block"`
}
type lighthouseAssignmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_lighthouse_assignment.
func (la lighthouseAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("id"))
}

// LighthouseDefinitionId returns a reference to field lighthouse_definition_id of azurerm_lighthouse_assignment.
func (la lighthouseAssignmentAttributes) LighthouseDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("lighthouse_definition_id"))
}

// Name returns a reference to field name of azurerm_lighthouse_assignment.
func (la lighthouseAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("name"))
}

// Scope returns a reference to field scope of azurerm_lighthouse_assignment.
func (la lighthouseAssignmentAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("scope"))
}

func (la lighthouseAssignmentAttributes) Timeouts() lighthouseassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lighthouseassignment.TimeoutsAttributes](la.ref.Append("timeouts"))
}

type lighthouseAssignmentState struct {
	Id                     string                              `json:"id"`
	LighthouseDefinitionId string                              `json:"lighthouse_definition_id"`
	Name                   string                              `json:"name"`
	Scope                  string                              `json:"scope"`
	Timeouts               *lighthouseassignment.TimeoutsState `json:"timeouts"`
}

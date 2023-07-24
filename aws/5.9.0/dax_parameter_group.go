// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	daxparametergroup "github.com/golingon/terraproviders/aws/5.9.0/daxparametergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDaxParameterGroup creates a new instance of [DaxParameterGroup].
func NewDaxParameterGroup(name string, args DaxParameterGroupArgs) *DaxParameterGroup {
	return &DaxParameterGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DaxParameterGroup)(nil)

// DaxParameterGroup represents the Terraform resource aws_dax_parameter_group.
type DaxParameterGroup struct {
	Name      string
	Args      DaxParameterGroupArgs
	state     *daxParameterGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DaxParameterGroup].
func (dpg *DaxParameterGroup) Type() string {
	return "aws_dax_parameter_group"
}

// LocalName returns the local name for [DaxParameterGroup].
func (dpg *DaxParameterGroup) LocalName() string {
	return dpg.Name
}

// Configuration returns the configuration (args) for [DaxParameterGroup].
func (dpg *DaxParameterGroup) Configuration() interface{} {
	return dpg.Args
}

// DependOn is used for other resources to depend on [DaxParameterGroup].
func (dpg *DaxParameterGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(dpg)
}

// Dependencies returns the list of resources [DaxParameterGroup] depends_on.
func (dpg *DaxParameterGroup) Dependencies() terra.Dependencies {
	return dpg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DaxParameterGroup].
func (dpg *DaxParameterGroup) LifecycleManagement() *terra.Lifecycle {
	return dpg.Lifecycle
}

// Attributes returns the attributes for [DaxParameterGroup].
func (dpg *DaxParameterGroup) Attributes() daxParameterGroupAttributes {
	return daxParameterGroupAttributes{ref: terra.ReferenceResource(dpg)}
}

// ImportState imports the given attribute values into [DaxParameterGroup]'s state.
func (dpg *DaxParameterGroup) ImportState(av io.Reader) error {
	dpg.state = &daxParameterGroupState{}
	if err := json.NewDecoder(av).Decode(dpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpg.Type(), dpg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DaxParameterGroup] has state.
func (dpg *DaxParameterGroup) State() (*daxParameterGroupState, bool) {
	return dpg.state, dpg.state != nil
}

// StateMust returns the state for [DaxParameterGroup]. Panics if the state is nil.
func (dpg *DaxParameterGroup) StateMust() *daxParameterGroupState {
	if dpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpg.Type(), dpg.LocalName()))
	}
	return dpg.state
}

// DaxParameterGroupArgs contains the configurations for aws_dax_parameter_group.
type DaxParameterGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: min=0
	Parameters []daxparametergroup.Parameters `hcl:"parameters,block" validate:"min=0"`
}
type daxParameterGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_dax_parameter_group.
func (dpg daxParameterGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dpg.ref.Append("description"))
}

// Id returns a reference to field id of aws_dax_parameter_group.
func (dpg daxParameterGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpg.ref.Append("id"))
}

// Name returns a reference to field name of aws_dax_parameter_group.
func (dpg daxParameterGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpg.ref.Append("name"))
}

func (dpg daxParameterGroupAttributes) Parameters() terra.SetValue[daxparametergroup.ParametersAttributes] {
	return terra.ReferenceAsSet[daxparametergroup.ParametersAttributes](dpg.ref.Append("parameters"))
}

type daxParameterGroupState struct {
	Description string                              `json:"description"`
	Id          string                              `json:"id"`
	Name        string                              `json:"name"`
	Parameters  []daxparametergroup.ParametersState `json:"parameters"`
}

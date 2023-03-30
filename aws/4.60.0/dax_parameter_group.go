// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	daxparametergroup "github.com/golingon/terraproviders/aws/4.60.0/daxparametergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDaxParameterGroup(name string, args DaxParameterGroupArgs) *DaxParameterGroup {
	return &DaxParameterGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DaxParameterGroup)(nil)

type DaxParameterGroup struct {
	Name  string
	Args  DaxParameterGroupArgs
	state *daxParameterGroupState
}

func (dpg *DaxParameterGroup) Type() string {
	return "aws_dax_parameter_group"
}

func (dpg *DaxParameterGroup) LocalName() string {
	return dpg.Name
}

func (dpg *DaxParameterGroup) Configuration() interface{} {
	return dpg.Args
}

func (dpg *DaxParameterGroup) Attributes() daxParameterGroupAttributes {
	return daxParameterGroupAttributes{ref: terra.ReferenceResource(dpg)}
}

func (dpg *DaxParameterGroup) ImportState(av io.Reader) error {
	dpg.state = &daxParameterGroupState{}
	if err := json.NewDecoder(av).Decode(dpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpg.Type(), dpg.LocalName(), err)
	}
	return nil
}

func (dpg *DaxParameterGroup) State() (*daxParameterGroupState, bool) {
	return dpg.state, dpg.state != nil
}

func (dpg *DaxParameterGroup) StateMust() *daxParameterGroupState {
	if dpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpg.Type(), dpg.LocalName()))
	}
	return dpg.state
}

func (dpg *DaxParameterGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(dpg)
}

type DaxParameterGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: min=0
	Parameters []daxparametergroup.Parameters `hcl:"parameters,block" validate:"min=0"`
	// DependsOn contains resources that DaxParameterGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type daxParameterGroupAttributes struct {
	ref terra.Reference
}

func (dpg daxParameterGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dpg.ref.Append("description"))
}

func (dpg daxParameterGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dpg.ref.Append("id"))
}

func (dpg daxParameterGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dpg.ref.Append("name"))
}

func (dpg daxParameterGroupAttributes) Parameters() terra.SetValue[daxparametergroup.ParametersAttributes] {
	return terra.ReferenceSet[daxparametergroup.ParametersAttributes](dpg.ref.Append("parameters"))
}

type daxParameterGroupState struct {
	Description string                              `json:"description"`
	Id          string                              `json:"id"`
	Name        string                              `json:"name"`
	Parameters  []daxparametergroup.ParametersState `json:"parameters"`
}

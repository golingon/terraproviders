// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	neptuneparametergroup "github.com/golingon/terraproviders/aws/4.60.0/neptuneparametergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewNeptuneParameterGroup(name string, args NeptuneParameterGroupArgs) *NeptuneParameterGroup {
	return &NeptuneParameterGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NeptuneParameterGroup)(nil)

type NeptuneParameterGroup struct {
	Name  string
	Args  NeptuneParameterGroupArgs
	state *neptuneParameterGroupState
}

func (npg *NeptuneParameterGroup) Type() string {
	return "aws_neptune_parameter_group"
}

func (npg *NeptuneParameterGroup) LocalName() string {
	return npg.Name
}

func (npg *NeptuneParameterGroup) Configuration() interface{} {
	return npg.Args
}

func (npg *NeptuneParameterGroup) Attributes() neptuneParameterGroupAttributes {
	return neptuneParameterGroupAttributes{ref: terra.ReferenceResource(npg)}
}

func (npg *NeptuneParameterGroup) ImportState(av io.Reader) error {
	npg.state = &neptuneParameterGroupState{}
	if err := json.NewDecoder(av).Decode(npg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", npg.Type(), npg.LocalName(), err)
	}
	return nil
}

func (npg *NeptuneParameterGroup) State() (*neptuneParameterGroupState, bool) {
	return npg.state, npg.state != nil
}

func (npg *NeptuneParameterGroup) StateMust() *neptuneParameterGroupState {
	if npg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", npg.Type(), npg.LocalName()))
	}
	return npg.state
}

func (npg *NeptuneParameterGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(npg)
}

type NeptuneParameterGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Family: string, required
	Family terra.StringValue `hcl:"family,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Parameter: min=0
	Parameter []neptuneparametergroup.Parameter `hcl:"parameter,block" validate:"min=0"`
	// DependsOn contains resources that NeptuneParameterGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type neptuneParameterGroupAttributes struct {
	ref terra.Reference
}

func (npg neptuneParameterGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(npg.ref.Append("arn"))
}

func (npg neptuneParameterGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceString(npg.ref.Append("description"))
}

func (npg neptuneParameterGroupAttributes) Family() terra.StringValue {
	return terra.ReferenceString(npg.ref.Append("family"))
}

func (npg neptuneParameterGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(npg.ref.Append("id"))
}

func (npg neptuneParameterGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(npg.ref.Append("name"))
}

func (npg neptuneParameterGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](npg.ref.Append("tags"))
}

func (npg neptuneParameterGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](npg.ref.Append("tags_all"))
}

func (npg neptuneParameterGroupAttributes) Parameter() terra.SetValue[neptuneparametergroup.ParameterAttributes] {
	return terra.ReferenceSet[neptuneparametergroup.ParameterAttributes](npg.ref.Append("parameter"))
}

type neptuneParameterGroupState struct {
	Arn         string                                 `json:"arn"`
	Description string                                 `json:"description"`
	Family      string                                 `json:"family"`
	Id          string                                 `json:"id"`
	Name        string                                 `json:"name"`
	Tags        map[string]string                      `json:"tags"`
	TagsAll     map[string]string                      `json:"tags_all"`
	Parameter   []neptuneparametergroup.ParameterState `json:"parameter"`
}

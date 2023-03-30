// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appconfigextension "github.com/golingon/terraproviders/aws/4.60.0/appconfigextension"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAppconfigExtension(name string, args AppconfigExtensionArgs) *AppconfigExtension {
	return &AppconfigExtension{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppconfigExtension)(nil)

type AppconfigExtension struct {
	Name  string
	Args  AppconfigExtensionArgs
	state *appconfigExtensionState
}

func (ae *AppconfigExtension) Type() string {
	return "aws_appconfig_extension"
}

func (ae *AppconfigExtension) LocalName() string {
	return ae.Name
}

func (ae *AppconfigExtension) Configuration() interface{} {
	return ae.Args
}

func (ae *AppconfigExtension) Attributes() appconfigExtensionAttributes {
	return appconfigExtensionAttributes{ref: terra.ReferenceResource(ae)}
}

func (ae *AppconfigExtension) ImportState(av io.Reader) error {
	ae.state = &appconfigExtensionState{}
	if err := json.NewDecoder(av).Decode(ae.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ae.Type(), ae.LocalName(), err)
	}
	return nil
}

func (ae *AppconfigExtension) State() (*appconfigExtensionState, bool) {
	return ae.state, ae.state != nil
}

func (ae *AppconfigExtension) StateMust() *appconfigExtensionState {
	if ae.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ae.Type(), ae.LocalName()))
	}
	return ae.state
}

func (ae *AppconfigExtension) DependOn() terra.Reference {
	return terra.ReferenceResource(ae)
}

type AppconfigExtensionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ActionPoint: min=1
	ActionPoint []appconfigextension.ActionPoint `hcl:"action_point,block" validate:"min=1"`
	// Parameter: min=0
	Parameter []appconfigextension.Parameter `hcl:"parameter,block" validate:"min=0"`
	// DependsOn contains resources that AppconfigExtension depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type appconfigExtensionAttributes struct {
	ref terra.Reference
}

func (ae appconfigExtensionAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ae.ref.Append("arn"))
}

func (ae appconfigExtensionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(ae.ref.Append("description"))
}

func (ae appconfigExtensionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ae.ref.Append("id"))
}

func (ae appconfigExtensionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ae.ref.Append("name"))
}

func (ae appconfigExtensionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ae.ref.Append("tags"))
}

func (ae appconfigExtensionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ae.ref.Append("tags_all"))
}

func (ae appconfigExtensionAttributes) Version() terra.NumberValue {
	return terra.ReferenceNumber(ae.ref.Append("version"))
}

func (ae appconfigExtensionAttributes) ActionPoint() terra.SetValue[appconfigextension.ActionPointAttributes] {
	return terra.ReferenceSet[appconfigextension.ActionPointAttributes](ae.ref.Append("action_point"))
}

func (ae appconfigExtensionAttributes) Parameter() terra.SetValue[appconfigextension.ParameterAttributes] {
	return terra.ReferenceSet[appconfigextension.ParameterAttributes](ae.ref.Append("parameter"))
}

type appconfigExtensionState struct {
	Arn         string                                `json:"arn"`
	Description string                                `json:"description"`
	Id          string                                `json:"id"`
	Name        string                                `json:"name"`
	Tags        map[string]string                     `json:"tags"`
	TagsAll     map[string]string                     `json:"tags_all"`
	Version     float64                               `json:"version"`
	ActionPoint []appconfigextension.ActionPointState `json:"action_point"`
	Parameter   []appconfigextension.ParameterState   `json:"parameter"`
}

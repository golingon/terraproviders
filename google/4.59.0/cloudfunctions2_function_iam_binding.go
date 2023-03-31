// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudfunctions2functioniambinding "github.com/golingon/terraproviders/google/4.59.0/cloudfunctions2functioniambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCloudfunctions2FunctionIamBinding(name string, args Cloudfunctions2FunctionIamBindingArgs) *Cloudfunctions2FunctionIamBinding {
	return &Cloudfunctions2FunctionIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudfunctions2FunctionIamBinding)(nil)

type Cloudfunctions2FunctionIamBinding struct {
	Name  string
	Args  Cloudfunctions2FunctionIamBindingArgs
	state *cloudfunctions2FunctionIamBindingState
}

func (cfib *Cloudfunctions2FunctionIamBinding) Type() string {
	return "google_cloudfunctions2_function_iam_binding"
}

func (cfib *Cloudfunctions2FunctionIamBinding) LocalName() string {
	return cfib.Name
}

func (cfib *Cloudfunctions2FunctionIamBinding) Configuration() interface{} {
	return cfib.Args
}

func (cfib *Cloudfunctions2FunctionIamBinding) Attributes() cloudfunctions2FunctionIamBindingAttributes {
	return cloudfunctions2FunctionIamBindingAttributes{ref: terra.ReferenceResource(cfib)}
}

func (cfib *Cloudfunctions2FunctionIamBinding) ImportState(av io.Reader) error {
	cfib.state = &cloudfunctions2FunctionIamBindingState{}
	if err := json.NewDecoder(av).Decode(cfib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfib.Type(), cfib.LocalName(), err)
	}
	return nil
}

func (cfib *Cloudfunctions2FunctionIamBinding) State() (*cloudfunctions2FunctionIamBindingState, bool) {
	return cfib.state, cfib.state != nil
}

func (cfib *Cloudfunctions2FunctionIamBinding) StateMust() *cloudfunctions2FunctionIamBindingState {
	if cfib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfib.Type(), cfib.LocalName()))
	}
	return cfib.state
}

func (cfib *Cloudfunctions2FunctionIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(cfib)
}

type Cloudfunctions2FunctionIamBindingArgs struct {
	// CloudFunction: string, required
	CloudFunction terra.StringValue `hcl:"cloud_function,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudfunctions2functioniambinding.Condition `hcl:"condition,block"`
	// DependsOn contains resources that Cloudfunctions2FunctionIamBinding depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type cloudfunctions2FunctionIamBindingAttributes struct {
	ref terra.Reference
}

func (cfib cloudfunctions2FunctionIamBindingAttributes) CloudFunction() terra.StringValue {
	return terra.ReferenceString(cfib.ref.Append("cloud_function"))
}

func (cfib cloudfunctions2FunctionIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(cfib.ref.Append("etag"))
}

func (cfib cloudfunctions2FunctionIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cfib.ref.Append("id"))
}

func (cfib cloudfunctions2FunctionIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceString(cfib.ref.Append("location"))
}

func (cfib cloudfunctions2FunctionIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cfib.ref.Append("members"))
}

func (cfib cloudfunctions2FunctionIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cfib.ref.Append("project"))
}

func (cfib cloudfunctions2FunctionIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceString(cfib.ref.Append("role"))
}

func (cfib cloudfunctions2FunctionIamBindingAttributes) Condition() terra.ListValue[cloudfunctions2functioniambinding.ConditionAttributes] {
	return terra.ReferenceList[cloudfunctions2functioniambinding.ConditionAttributes](cfib.ref.Append("condition"))
}

type cloudfunctions2FunctionIamBindingState struct {
	CloudFunction string                                             `json:"cloud_function"`
	Etag          string                                             `json:"etag"`
	Id            string                                             `json:"id"`
	Location      string                                             `json:"location"`
	Members       []string                                           `json:"members"`
	Project       string                                             `json:"project"`
	Role          string                                             `json:"role"`
	Condition     []cloudfunctions2functioniambinding.ConditionState `json:"condition"`
}

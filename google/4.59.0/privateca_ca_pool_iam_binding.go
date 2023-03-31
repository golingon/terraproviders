// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	privatecacapooliambinding "github.com/golingon/terraproviders/google/4.59.0/privatecacapooliambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewPrivatecaCaPoolIamBinding(name string, args PrivatecaCaPoolIamBindingArgs) *PrivatecaCaPoolIamBinding {
	return &PrivatecaCaPoolIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivatecaCaPoolIamBinding)(nil)

type PrivatecaCaPoolIamBinding struct {
	Name  string
	Args  PrivatecaCaPoolIamBindingArgs
	state *privatecaCaPoolIamBindingState
}

func (pcpib *PrivatecaCaPoolIamBinding) Type() string {
	return "google_privateca_ca_pool_iam_binding"
}

func (pcpib *PrivatecaCaPoolIamBinding) LocalName() string {
	return pcpib.Name
}

func (pcpib *PrivatecaCaPoolIamBinding) Configuration() interface{} {
	return pcpib.Args
}

func (pcpib *PrivatecaCaPoolIamBinding) Attributes() privatecaCaPoolIamBindingAttributes {
	return privatecaCaPoolIamBindingAttributes{ref: terra.ReferenceResource(pcpib)}
}

func (pcpib *PrivatecaCaPoolIamBinding) ImportState(av io.Reader) error {
	pcpib.state = &privatecaCaPoolIamBindingState{}
	if err := json.NewDecoder(av).Decode(pcpib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pcpib.Type(), pcpib.LocalName(), err)
	}
	return nil
}

func (pcpib *PrivatecaCaPoolIamBinding) State() (*privatecaCaPoolIamBindingState, bool) {
	return pcpib.state, pcpib.state != nil
}

func (pcpib *PrivatecaCaPoolIamBinding) StateMust() *privatecaCaPoolIamBindingState {
	if pcpib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pcpib.Type(), pcpib.LocalName()))
	}
	return pcpib.state
}

func (pcpib *PrivatecaCaPoolIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(pcpib)
}

type PrivatecaCaPoolIamBindingArgs struct {
	// CaPool: string, required
	CaPool terra.StringValue `hcl:"ca_pool,attr" validate:"required"`
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
	Condition *privatecacapooliambinding.Condition `hcl:"condition,block"`
	// DependsOn contains resources that PrivatecaCaPoolIamBinding depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type privatecaCaPoolIamBindingAttributes struct {
	ref terra.Reference
}

func (pcpib privatecaCaPoolIamBindingAttributes) CaPool() terra.StringValue {
	return terra.ReferenceString(pcpib.ref.Append("ca_pool"))
}

func (pcpib privatecaCaPoolIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(pcpib.ref.Append("etag"))
}

func (pcpib privatecaCaPoolIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceString(pcpib.ref.Append("id"))
}

func (pcpib privatecaCaPoolIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceString(pcpib.ref.Append("location"))
}

func (pcpib privatecaCaPoolIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](pcpib.ref.Append("members"))
}

func (pcpib privatecaCaPoolIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceString(pcpib.ref.Append("project"))
}

func (pcpib privatecaCaPoolIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceString(pcpib.ref.Append("role"))
}

func (pcpib privatecaCaPoolIamBindingAttributes) Condition() terra.ListValue[privatecacapooliambinding.ConditionAttributes] {
	return terra.ReferenceList[privatecacapooliambinding.ConditionAttributes](pcpib.ref.Append("condition"))
}

type privatecaCaPoolIamBindingState struct {
	CaPool    string                                     `json:"ca_pool"`
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Location  string                                     `json:"location"`
	Members   []string                                   `json:"members"`
	Project   string                                     `json:"project"`
	Role      string                                     `json:"role"`
	Condition []privatecacapooliambinding.ConditionState `json:"condition"`
}

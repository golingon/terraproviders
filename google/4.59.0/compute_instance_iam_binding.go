// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeinstanceiambinding "github.com/golingon/terraproviders/google/4.59.0/computeinstanceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeInstanceIamBinding(name string, args ComputeInstanceIamBindingArgs) *ComputeInstanceIamBinding {
	return &ComputeInstanceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstanceIamBinding)(nil)

type ComputeInstanceIamBinding struct {
	Name  string
	Args  ComputeInstanceIamBindingArgs
	state *computeInstanceIamBindingState
}

func (ciib *ComputeInstanceIamBinding) Type() string {
	return "google_compute_instance_iam_binding"
}

func (ciib *ComputeInstanceIamBinding) LocalName() string {
	return ciib.Name
}

func (ciib *ComputeInstanceIamBinding) Configuration() interface{} {
	return ciib.Args
}

func (ciib *ComputeInstanceIamBinding) Attributes() computeInstanceIamBindingAttributes {
	return computeInstanceIamBindingAttributes{ref: terra.ReferenceResource(ciib)}
}

func (ciib *ComputeInstanceIamBinding) ImportState(av io.Reader) error {
	ciib.state = &computeInstanceIamBindingState{}
	if err := json.NewDecoder(av).Decode(ciib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ciib.Type(), ciib.LocalName(), err)
	}
	return nil
}

func (ciib *ComputeInstanceIamBinding) State() (*computeInstanceIamBindingState, bool) {
	return ciib.state, ciib.state != nil
}

func (ciib *ComputeInstanceIamBinding) StateMust() *computeInstanceIamBindingState {
	if ciib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ciib.Type(), ciib.LocalName()))
	}
	return ciib.state
}

func (ciib *ComputeInstanceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ciib)
}

type ComputeInstanceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Condition: optional
	Condition *computeinstanceiambinding.Condition `hcl:"condition,block"`
	// DependsOn contains resources that ComputeInstanceIamBinding depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeInstanceIamBindingAttributes struct {
	ref terra.Reference
}

func (ciib computeInstanceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(ciib.ref.Append("etag"))
}

func (ciib computeInstanceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ciib.ref.Append("id"))
}

func (ciib computeInstanceIamBindingAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceString(ciib.ref.Append("instance_name"))
}

func (ciib computeInstanceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ciib.ref.Append("members"))
}

func (ciib computeInstanceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceString(ciib.ref.Append("project"))
}

func (ciib computeInstanceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceString(ciib.ref.Append("role"))
}

func (ciib computeInstanceIamBindingAttributes) Zone() terra.StringValue {
	return terra.ReferenceString(ciib.ref.Append("zone"))
}

func (ciib computeInstanceIamBindingAttributes) Condition() terra.ListValue[computeinstanceiambinding.ConditionAttributes] {
	return terra.ReferenceList[computeinstanceiambinding.ConditionAttributes](ciib.ref.Append("condition"))
}

type computeInstanceIamBindingState struct {
	Etag         string                                     `json:"etag"`
	Id           string                                     `json:"id"`
	InstanceName string                                     `json:"instance_name"`
	Members      []string                                   `json:"members"`
	Project      string                                     `json:"project"`
	Role         string                                     `json:"role"`
	Zone         string                                     `json:"zone"`
	Condition    []computeinstanceiambinding.ConditionState `json:"condition"`
}

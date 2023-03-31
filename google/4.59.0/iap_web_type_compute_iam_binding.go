// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapwebtypecomputeiambinding "github.com/golingon/terraproviders/google/4.59.0/iapwebtypecomputeiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewIapWebTypeComputeIamBinding(name string, args IapWebTypeComputeIamBindingArgs) *IapWebTypeComputeIamBinding {
	return &IapWebTypeComputeIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebTypeComputeIamBinding)(nil)

type IapWebTypeComputeIamBinding struct {
	Name  string
	Args  IapWebTypeComputeIamBindingArgs
	state *iapWebTypeComputeIamBindingState
}

func (iwtcib *IapWebTypeComputeIamBinding) Type() string {
	return "google_iap_web_type_compute_iam_binding"
}

func (iwtcib *IapWebTypeComputeIamBinding) LocalName() string {
	return iwtcib.Name
}

func (iwtcib *IapWebTypeComputeIamBinding) Configuration() interface{} {
	return iwtcib.Args
}

func (iwtcib *IapWebTypeComputeIamBinding) Attributes() iapWebTypeComputeIamBindingAttributes {
	return iapWebTypeComputeIamBindingAttributes{ref: terra.ReferenceResource(iwtcib)}
}

func (iwtcib *IapWebTypeComputeIamBinding) ImportState(av io.Reader) error {
	iwtcib.state = &iapWebTypeComputeIamBindingState{}
	if err := json.NewDecoder(av).Decode(iwtcib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwtcib.Type(), iwtcib.LocalName(), err)
	}
	return nil
}

func (iwtcib *IapWebTypeComputeIamBinding) State() (*iapWebTypeComputeIamBindingState, bool) {
	return iwtcib.state, iwtcib.state != nil
}

func (iwtcib *IapWebTypeComputeIamBinding) StateMust() *iapWebTypeComputeIamBindingState {
	if iwtcib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwtcib.Type(), iwtcib.LocalName()))
	}
	return iwtcib.state
}

func (iwtcib *IapWebTypeComputeIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(iwtcib)
}

type IapWebTypeComputeIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *iapwebtypecomputeiambinding.Condition `hcl:"condition,block"`
	// DependsOn contains resources that IapWebTypeComputeIamBinding depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type iapWebTypeComputeIamBindingAttributes struct {
	ref terra.Reference
}

func (iwtcib iapWebTypeComputeIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(iwtcib.ref.Append("etag"))
}

func (iwtcib iapWebTypeComputeIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceString(iwtcib.ref.Append("id"))
}

func (iwtcib iapWebTypeComputeIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](iwtcib.ref.Append("members"))
}

func (iwtcib iapWebTypeComputeIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceString(iwtcib.ref.Append("project"))
}

func (iwtcib iapWebTypeComputeIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceString(iwtcib.ref.Append("role"))
}

func (iwtcib iapWebTypeComputeIamBindingAttributes) Condition() terra.ListValue[iapwebtypecomputeiambinding.ConditionAttributes] {
	return terra.ReferenceList[iapwebtypecomputeiambinding.ConditionAttributes](iwtcib.ref.Append("condition"))
}

type iapWebTypeComputeIamBindingState struct {
	Etag      string                                       `json:"etag"`
	Id        string                                       `json:"id"`
	Members   []string                                     `json:"members"`
	Project   string                                       `json:"project"`
	Role      string                                       `json:"role"`
	Condition []iapwebtypecomputeiambinding.ConditionState `json:"condition"`
}

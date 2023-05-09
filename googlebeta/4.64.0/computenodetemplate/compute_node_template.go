// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computenodetemplate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NodeTypeFlexibility struct {
	// Cpus: string, optional
	Cpus terra.StringValue `hcl:"cpus,attr"`
	// Memory: string, optional
	Memory terra.StringValue `hcl:"memory,attr"`
}

type ServerBinding struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type NodeTypeFlexibilityAttributes struct {
	ref terra.Reference
}

func (ntf NodeTypeFlexibilityAttributes) InternalRef() (terra.Reference, error) {
	return ntf.ref, nil
}

func (ntf NodeTypeFlexibilityAttributes) InternalWithRef(ref terra.Reference) NodeTypeFlexibilityAttributes {
	return NodeTypeFlexibilityAttributes{ref: ref}
}

func (ntf NodeTypeFlexibilityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ntf.ref.InternalTokens()
}

func (ntf NodeTypeFlexibilityAttributes) Cpus() terra.StringValue {
	return terra.ReferenceAsString(ntf.ref.Append("cpus"))
}

func (ntf NodeTypeFlexibilityAttributes) LocalSsd() terra.StringValue {
	return terra.ReferenceAsString(ntf.ref.Append("local_ssd"))
}

func (ntf NodeTypeFlexibilityAttributes) Memory() terra.StringValue {
	return terra.ReferenceAsString(ntf.ref.Append("memory"))
}

type ServerBindingAttributes struct {
	ref terra.Reference
}

func (sb ServerBindingAttributes) InternalRef() (terra.Reference, error) {
	return sb.ref, nil
}

func (sb ServerBindingAttributes) InternalWithRef(ref terra.Reference) ServerBindingAttributes {
	return ServerBindingAttributes{ref: ref}
}

func (sb ServerBindingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sb.ref.InternalTokens()
}

func (sb ServerBindingAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("type"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type NodeTypeFlexibilityState struct {
	Cpus     string `json:"cpus"`
	LocalSsd string `json:"local_ssd"`
	Memory   string `json:"memory"`
}

type ServerBindingState struct {
	Type string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package springcloudbuilddeployment

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Quota struct {
	// Cpu: string, optional
	Cpu terra.StringValue `hcl:"cpu,attr"`
	// Memory: string, optional
	Memory terra.StringValue `hcl:"memory,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type QuotaAttributes struct {
	ref terra.Reference
}

func (q QuotaAttributes) InternalRef() (terra.Reference, error) {
	return q.ref, nil
}

func (q QuotaAttributes) InternalWithRef(ref terra.Reference) QuotaAttributes {
	return QuotaAttributes{ref: ref}
}

func (q QuotaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return q.ref.InternalTokens()
}

func (q QuotaAttributes) Cpu() terra.StringValue {
	return terra.ReferenceAsString(q.ref.Append("cpu"))
}

func (q QuotaAttributes) Memory() terra.StringValue {
	return terra.ReferenceAsString(q.ref.Append("memory"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type QuotaState struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apigeeorganization

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Properties struct {
	// Property: min=0
	Property []Property `hcl:"property,block" validate:"min=0"`
}

type Property struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type PropertiesAttributes struct {
	ref terra.Reference
}

func (p PropertiesAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PropertiesAttributes) InternalWithRef(ref terra.Reference) PropertiesAttributes {
	return PropertiesAttributes{ref: ref}
}

func (p PropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PropertiesAttributes) Property() terra.ListValue[PropertyAttributes] {
	return terra.ReferenceAsList[PropertyAttributes](p.ref.Append("property"))
}

type PropertyAttributes struct {
	ref terra.Reference
}

func (p PropertyAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PropertyAttributes) InternalWithRef(ref terra.Reference) PropertyAttributes {
	return PropertyAttributes{ref: ref}
}

func (p PropertyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PropertyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p PropertyAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("value"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type PropertiesState struct {
	Property []PropertyState `json:"property"`
}

type PropertyState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
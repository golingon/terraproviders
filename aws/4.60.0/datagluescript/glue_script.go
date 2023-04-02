// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datagluescript

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DagEdge struct {
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// Target: string, required
	Target terra.StringValue `hcl:"target,attr" validate:"required"`
	// TargetParameter: string, optional
	TargetParameter terra.StringValue `hcl:"target_parameter,attr"`
}

type DagNode struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// LineNumber: number, optional
	LineNumber terra.NumberValue `hcl:"line_number,attr"`
	// NodeType: string, required
	NodeType terra.StringValue `hcl:"node_type,attr" validate:"required"`
	// Args: min=1
	Args []Args `hcl:"args,block" validate:"min=1"`
}

type Args struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Param: bool, optional
	Param terra.BoolValue `hcl:"param,attr"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type DagEdgeAttributes struct {
	ref terra.Reference
}

func (de DagEdgeAttributes) InternalRef() (terra.Reference, error) {
	return de.ref, nil
}

func (de DagEdgeAttributes) InternalWithRef(ref terra.Reference) DagEdgeAttributes {
	return DagEdgeAttributes{ref: ref}
}

func (de DagEdgeAttributes) InternalTokens() hclwrite.Tokens {
	return de.ref.InternalTokens()
}

func (de DagEdgeAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("source"))
}

func (de DagEdgeAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("target"))
}

func (de DagEdgeAttributes) TargetParameter() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("target_parameter"))
}

type DagNodeAttributes struct {
	ref terra.Reference
}

func (dn DagNodeAttributes) InternalRef() (terra.Reference, error) {
	return dn.ref, nil
}

func (dn DagNodeAttributes) InternalWithRef(ref terra.Reference) DagNodeAttributes {
	return DagNodeAttributes{ref: ref}
}

func (dn DagNodeAttributes) InternalTokens() hclwrite.Tokens {
	return dn.ref.InternalTokens()
}

func (dn DagNodeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dn.ref.Append("id"))
}

func (dn DagNodeAttributes) LineNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(dn.ref.Append("line_number"))
}

func (dn DagNodeAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(dn.ref.Append("node_type"))
}

func (dn DagNodeAttributes) Args() terra.ListValue[ArgsAttributes] {
	return terra.ReferenceAsList[ArgsAttributes](dn.ref.Append("args"))
}

type ArgsAttributes struct {
	ref terra.Reference
}

func (a ArgsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ArgsAttributes) InternalWithRef(ref terra.Reference) ArgsAttributes {
	return ArgsAttributes{ref: ref}
}

func (a ArgsAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a ArgsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

func (a ArgsAttributes) Param() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("param"))
}

func (a ArgsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("value"))
}

type DagEdgeState struct {
	Source          string `json:"source"`
	Target          string `json:"target"`
	TargetParameter string `json:"target_parameter"`
}

type DagNodeState struct {
	Id         string      `json:"id"`
	LineNumber float64     `json:"line_number"`
	NodeType   string      `json:"node_type"`
	Args       []ArgsState `json:"args"`
}

type ArgsState struct {
	Name  string `json:"name"`
	Param bool   `json:"param"`
	Value string `json:"value"`
}

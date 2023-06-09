// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package batchjobdefinition

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RetryStrategy struct {
	// Attempts: number, optional
	Attempts terra.NumberValue `hcl:"attempts,attr"`
	// EvaluateOnExit: min=0,max=5
	EvaluateOnExit []EvaluateOnExit `hcl:"evaluate_on_exit,block" validate:"min=0,max=5"`
}

type EvaluateOnExit struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// OnExitCode: string, optional
	OnExitCode terra.StringValue `hcl:"on_exit_code,attr"`
	// OnReason: string, optional
	OnReason terra.StringValue `hcl:"on_reason,attr"`
	// OnStatusReason: string, optional
	OnStatusReason terra.StringValue `hcl:"on_status_reason,attr"`
}

type Timeout struct {
	// AttemptDurationSeconds: number, optional
	AttemptDurationSeconds terra.NumberValue `hcl:"attempt_duration_seconds,attr"`
}

type RetryStrategyAttributes struct {
	ref terra.Reference
}

func (rs RetryStrategyAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs RetryStrategyAttributes) InternalWithRef(ref terra.Reference) RetryStrategyAttributes {
	return RetryStrategyAttributes{ref: ref}
}

func (rs RetryStrategyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs RetryStrategyAttributes) Attempts() terra.NumberValue {
	return terra.ReferenceAsNumber(rs.ref.Append("attempts"))
}

func (rs RetryStrategyAttributes) EvaluateOnExit() terra.ListValue[EvaluateOnExitAttributes] {
	return terra.ReferenceAsList[EvaluateOnExitAttributes](rs.ref.Append("evaluate_on_exit"))
}

type EvaluateOnExitAttributes struct {
	ref terra.Reference
}

func (eoe EvaluateOnExitAttributes) InternalRef() (terra.Reference, error) {
	return eoe.ref, nil
}

func (eoe EvaluateOnExitAttributes) InternalWithRef(ref terra.Reference) EvaluateOnExitAttributes {
	return EvaluateOnExitAttributes{ref: ref}
}

func (eoe EvaluateOnExitAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eoe.ref.InternalTokens()
}

func (eoe EvaluateOnExitAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(eoe.ref.Append("action"))
}

func (eoe EvaluateOnExitAttributes) OnExitCode() terra.StringValue {
	return terra.ReferenceAsString(eoe.ref.Append("on_exit_code"))
}

func (eoe EvaluateOnExitAttributes) OnReason() terra.StringValue {
	return terra.ReferenceAsString(eoe.ref.Append("on_reason"))
}

func (eoe EvaluateOnExitAttributes) OnStatusReason() terra.StringValue {
	return terra.ReferenceAsString(eoe.ref.Append("on_status_reason"))
}

type TimeoutAttributes struct {
	ref terra.Reference
}

func (t TimeoutAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutAttributes) InternalWithRef(ref terra.Reference) TimeoutAttributes {
	return TimeoutAttributes{ref: ref}
}

func (t TimeoutAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutAttributes) AttemptDurationSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("attempt_duration_seconds"))
}

type RetryStrategyState struct {
	Attempts       float64               `json:"attempts"`
	EvaluateOnExit []EvaluateOnExitState `json:"evaluate_on_exit"`
}

type EvaluateOnExitState struct {
	Action         string `json:"action"`
	OnExitCode     string `json:"on_exit_code"`
	OnReason       string `json:"on_reason"`
	OnStatusReason string `json:"on_status_reason"`
}

type TimeoutState struct {
	AttemptDurationSeconds float64 `json:"attempt_duration_seconds"`
}

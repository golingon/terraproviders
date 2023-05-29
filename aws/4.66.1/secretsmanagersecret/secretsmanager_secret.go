// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package secretsmanagersecret

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Replica struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
}

type RotationRules struct {
	// AutomaticallyAfterDays: number, optional
	AutomaticallyAfterDays terra.NumberValue `hcl:"automatically_after_days,attr"`
	// Duration: string, optional
	Duration terra.StringValue `hcl:"duration,attr"`
	// ScheduleExpression: string, optional
	ScheduleExpression terra.StringValue `hcl:"schedule_expression,attr"`
}

type ReplicaAttributes struct {
	ref terra.Reference
}

func (r ReplicaAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ReplicaAttributes) InternalWithRef(ref terra.Reference) ReplicaAttributes {
	return ReplicaAttributes{ref: ref}
}

func (r ReplicaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ReplicaAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("kms_key_id"))
}

func (r ReplicaAttributes) LastAccessedDate() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("last_accessed_date"))
}

func (r ReplicaAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("region"))
}

func (r ReplicaAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("status"))
}

func (r ReplicaAttributes) StatusMessage() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("status_message"))
}

type RotationRulesAttributes struct {
	ref terra.Reference
}

func (rr RotationRulesAttributes) InternalRef() (terra.Reference, error) {
	return rr.ref, nil
}

func (rr RotationRulesAttributes) InternalWithRef(ref terra.Reference) RotationRulesAttributes {
	return RotationRulesAttributes{ref: ref}
}

func (rr RotationRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rr.ref.InternalTokens()
}

func (rr RotationRulesAttributes) AutomaticallyAfterDays() terra.NumberValue {
	return terra.ReferenceAsNumber(rr.ref.Append("automatically_after_days"))
}

func (rr RotationRulesAttributes) Duration() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("duration"))
}

func (rr RotationRulesAttributes) ScheduleExpression() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("schedule_expression"))
}

type ReplicaState struct {
	KmsKeyId         string `json:"kms_key_id"`
	LastAccessedDate string `json:"last_accessed_date"`
	Region           string `json:"region"`
	Status           string `json:"status"`
	StatusMessage    string `json:"status_message"`
}

type RotationRulesState struct {
	AutomaticallyAfterDays float64 `json:"automatically_after_days"`
	Duration               string  `json:"duration"`
	ScheduleExpression     string  `json:"schedule_expression"`
}
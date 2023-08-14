// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package route53recoveryreadinessresourceset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Resources struct {
	// ReadinessScopes: list of string, optional
	ReadinessScopes terra.ListValue[terra.StringValue] `hcl:"readiness_scopes,attr"`
	// ResourceArn: string, optional
	ResourceArn terra.StringValue `hcl:"resource_arn,attr"`
	// DnsTargetResource: optional
	DnsTargetResource *DnsTargetResource `hcl:"dns_target_resource,block"`
}

type DnsTargetResource struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// HostedZoneArn: string, optional
	HostedZoneArn terra.StringValue `hcl:"hosted_zone_arn,attr"`
	// RecordSetId: string, optional
	RecordSetId terra.StringValue `hcl:"record_set_id,attr"`
	// RecordType: string, optional
	RecordType terra.StringValue `hcl:"record_type,attr"`
	// TargetResource: optional
	TargetResource *TargetResource `hcl:"target_resource,block"`
}

type TargetResource struct {
	// NlbResource: optional
	NlbResource *NlbResource `hcl:"nlb_resource,block"`
	// R53Resource: optional
	R53Resource *R53Resource `hcl:"r53_resource,block"`
}

type NlbResource struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
}

type R53Resource struct {
	// DomainName: string, optional
	DomainName terra.StringValue `hcl:"domain_name,attr"`
	// RecordSetId: string, optional
	RecordSetId terra.StringValue `hcl:"record_set_id,attr"`
}

type Timeouts struct {
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type ResourcesAttributes struct {
	ref terra.Reference
}

func (r ResourcesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ResourcesAttributes) InternalWithRef(ref terra.Reference) ResourcesAttributes {
	return ResourcesAttributes{ref: ref}
}

func (r ResourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ResourcesAttributes) ComponentId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("component_id"))
}

func (r ResourcesAttributes) ReadinessScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("readiness_scopes"))
}

func (r ResourcesAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("resource_arn"))
}

func (r ResourcesAttributes) DnsTargetResource() terra.ListValue[DnsTargetResourceAttributes] {
	return terra.ReferenceAsList[DnsTargetResourceAttributes](r.ref.Append("dns_target_resource"))
}

type DnsTargetResourceAttributes struct {
	ref terra.Reference
}

func (dtr DnsTargetResourceAttributes) InternalRef() (terra.Reference, error) {
	return dtr.ref, nil
}

func (dtr DnsTargetResourceAttributes) InternalWithRef(ref terra.Reference) DnsTargetResourceAttributes {
	return DnsTargetResourceAttributes{ref: ref}
}

func (dtr DnsTargetResourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dtr.ref.InternalTokens()
}

func (dtr DnsTargetResourceAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("domain_name"))
}

func (dtr DnsTargetResourceAttributes) HostedZoneArn() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("hosted_zone_arn"))
}

func (dtr DnsTargetResourceAttributes) RecordSetId() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("record_set_id"))
}

func (dtr DnsTargetResourceAttributes) RecordType() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("record_type"))
}

func (dtr DnsTargetResourceAttributes) TargetResource() terra.ListValue[TargetResourceAttributes] {
	return terra.ReferenceAsList[TargetResourceAttributes](dtr.ref.Append("target_resource"))
}

type TargetResourceAttributes struct {
	ref terra.Reference
}

func (tr TargetResourceAttributes) InternalRef() (terra.Reference, error) {
	return tr.ref, nil
}

func (tr TargetResourceAttributes) InternalWithRef(ref terra.Reference) TargetResourceAttributes {
	return TargetResourceAttributes{ref: ref}
}

func (tr TargetResourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tr.ref.InternalTokens()
}

func (tr TargetResourceAttributes) NlbResource() terra.ListValue[NlbResourceAttributes] {
	return terra.ReferenceAsList[NlbResourceAttributes](tr.ref.Append("nlb_resource"))
}

func (tr TargetResourceAttributes) R53Resource() terra.ListValue[R53ResourceAttributes] {
	return terra.ReferenceAsList[R53ResourceAttributes](tr.ref.Append("r53_resource"))
}

type NlbResourceAttributes struct {
	ref terra.Reference
}

func (nr NlbResourceAttributes) InternalRef() (terra.Reference, error) {
	return nr.ref, nil
}

func (nr NlbResourceAttributes) InternalWithRef(ref terra.Reference) NlbResourceAttributes {
	return NlbResourceAttributes{ref: ref}
}

func (nr NlbResourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nr.ref.InternalTokens()
}

func (nr NlbResourceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nr.ref.Append("arn"))
}

type R53ResourceAttributes struct {
	ref terra.Reference
}

func (rr R53ResourceAttributes) InternalRef() (terra.Reference, error) {
	return rr.ref, nil
}

func (rr R53ResourceAttributes) InternalWithRef(ref terra.Reference) R53ResourceAttributes {
	return R53ResourceAttributes{ref: ref}
}

func (rr R53ResourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rr.ref.InternalTokens()
}

func (rr R53ResourceAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("domain_name"))
}

func (rr R53ResourceAttributes) RecordSetId() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("record_set_id"))
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

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type ResourcesState struct {
	ComponentId       string                   `json:"component_id"`
	ReadinessScopes   []string                 `json:"readiness_scopes"`
	ResourceArn       string                   `json:"resource_arn"`
	DnsTargetResource []DnsTargetResourceState `json:"dns_target_resource"`
}

type DnsTargetResourceState struct {
	DomainName     string                `json:"domain_name"`
	HostedZoneArn  string                `json:"hosted_zone_arn"`
	RecordSetId    string                `json:"record_set_id"`
	RecordType     string                `json:"record_type"`
	TargetResource []TargetResourceState `json:"target_resource"`
}

type TargetResourceState struct {
	NlbResource []NlbResourceState `json:"nlb_resource"`
	R53Resource []R53ResourceState `json:"r53_resource"`
}

type NlbResourceState struct {
	Arn string `json:"arn"`
}

type R53ResourceState struct {
	DomainName  string `json:"domain_name"`
	RecordSetId string `json:"record_set_id"`
}

type TimeoutsState struct {
	Delete string `json:"delete"`
}

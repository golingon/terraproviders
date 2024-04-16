// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_firebase_hosting_custom_domain

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type CertAttributes struct {
	ref terra.Reference
}

func (c CertAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CertAttributes) InternalWithRef(ref terra.Reference) CertAttributes {
	return CertAttributes{ref: ref}
}

func (c CertAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CertAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("state"))
}

func (c CertAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
}

func (c CertAttributes) Verification() terra.ListValue[CertVerificationAttributes] {
	return terra.ReferenceAsList[CertVerificationAttributes](c.ref.Append("verification"))
}

type CertVerificationAttributes struct {
	ref terra.Reference
}

func (v CertVerificationAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v CertVerificationAttributes) InternalWithRef(ref terra.Reference) CertVerificationAttributes {
	return CertVerificationAttributes{ref: ref}
}

func (v CertVerificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v CertVerificationAttributes) Dns() terra.ListValue[CertVerificationDnsAttributes] {
	return terra.ReferenceAsList[CertVerificationDnsAttributes](v.ref.Append("dns"))
}

func (v CertVerificationAttributes) Http() terra.ListValue[CertVerificationHttpAttributes] {
	return terra.ReferenceAsList[CertVerificationHttpAttributes](v.ref.Append("http"))
}

type CertVerificationDnsAttributes struct {
	ref terra.Reference
}

func (d CertVerificationDnsAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d CertVerificationDnsAttributes) InternalWithRef(ref terra.Reference) CertVerificationDnsAttributes {
	return CertVerificationDnsAttributes{ref: ref}
}

func (d CertVerificationDnsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d CertVerificationDnsAttributes) CheckTime() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("check_time"))
}

func (d CertVerificationDnsAttributes) Desired() terra.ListValue[CertVerificationDnsDesiredAttributes] {
	return terra.ReferenceAsList[CertVerificationDnsDesiredAttributes](d.ref.Append("desired"))
}

func (d CertVerificationDnsAttributes) Discovered() terra.ListValue[CertVerificationDnsDiscoveredAttributes] {
	return terra.ReferenceAsList[CertVerificationDnsDiscoveredAttributes](d.ref.Append("discovered"))
}

type CertVerificationDnsDesiredAttributes struct {
	ref terra.Reference
}

func (d CertVerificationDnsDesiredAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d CertVerificationDnsDesiredAttributes) InternalWithRef(ref terra.Reference) CertVerificationDnsDesiredAttributes {
	return CertVerificationDnsDesiredAttributes{ref: ref}
}

func (d CertVerificationDnsDesiredAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d CertVerificationDnsDesiredAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("domain_name"))
}

func (d CertVerificationDnsDesiredAttributes) Records() terra.ListValue[CertVerificationDecc26A4Attributes] {
	return terra.ReferenceAsList[CertVerificationDecc26A4Attributes](d.ref.Append("records"))
}

type CertVerificationDecc26A4Attributes struct {
	ref terra.Reference
}

func (r CertVerificationDecc26A4Attributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r CertVerificationDecc26A4Attributes) InternalWithRef(ref terra.Reference) CertVerificationDecc26A4Attributes {
	return CertVerificationDecc26A4Attributes{ref: ref}
}

func (r CertVerificationDecc26A4Attributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r CertVerificationDecc26A4Attributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("domain_name"))
}

func (r CertVerificationDecc26A4Attributes) Rdata() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("rdata"))
}

func (r CertVerificationDecc26A4Attributes) RequiredAction() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("required_action"))
}

func (r CertVerificationDecc26A4Attributes) Type() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("type"))
}

type CertVerificationDnsDiscoveredAttributes struct {
	ref terra.Reference
}

func (d CertVerificationDnsDiscoveredAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d CertVerificationDnsDiscoveredAttributes) InternalWithRef(ref terra.Reference) CertVerificationDnsDiscoveredAttributes {
	return CertVerificationDnsDiscoveredAttributes{ref: ref}
}

func (d CertVerificationDnsDiscoveredAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d CertVerificationDnsDiscoveredAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("domain_name"))
}

func (d CertVerificationDnsDiscoveredAttributes) Records() terra.ListValue[CertVerification6F7683B8Attributes] {
	return terra.ReferenceAsList[CertVerification6F7683B8Attributes](d.ref.Append("records"))
}

type CertVerification6F7683B8Attributes struct {
	ref terra.Reference
}

func (r CertVerification6F7683B8Attributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r CertVerification6F7683B8Attributes) InternalWithRef(ref terra.Reference) CertVerification6F7683B8Attributes {
	return CertVerification6F7683B8Attributes{ref: ref}
}

func (r CertVerification6F7683B8Attributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r CertVerification6F7683B8Attributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("domain_name"))
}

func (r CertVerification6F7683B8Attributes) Rdata() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("rdata"))
}

func (r CertVerification6F7683B8Attributes) RequiredAction() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("required_action"))
}

func (r CertVerification6F7683B8Attributes) Type() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("type"))
}

type CertVerificationHttpAttributes struct {
	ref terra.Reference
}

func (h CertVerificationHttpAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h CertVerificationHttpAttributes) InternalWithRef(ref terra.Reference) CertVerificationHttpAttributes {
	return CertVerificationHttpAttributes{ref: ref}
}

func (h CertVerificationHttpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h CertVerificationHttpAttributes) Desired() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("desired"))
}

func (h CertVerificationHttpAttributes) Discovered() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("discovered"))
}

func (h CertVerificationHttpAttributes) LastCheckTime() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("last_check_time"))
}

func (h CertVerificationHttpAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("path"))
}

type IssuesAttributes struct {
	ref terra.Reference
}

func (i IssuesAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IssuesAttributes) InternalWithRef(ref terra.Reference) IssuesAttributes {
	return IssuesAttributes{ref: ref}
}

func (i IssuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IssuesAttributes) Code() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("code"))
}

func (i IssuesAttributes) Details() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("details"))
}

func (i IssuesAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("message"))
}

type RequiredDnsUpdatesAttributes struct {
	ref terra.Reference
}

func (rdu RequiredDnsUpdatesAttributes) InternalRef() (terra.Reference, error) {
	return rdu.ref, nil
}

func (rdu RequiredDnsUpdatesAttributes) InternalWithRef(ref terra.Reference) RequiredDnsUpdatesAttributes {
	return RequiredDnsUpdatesAttributes{ref: ref}
}

func (rdu RequiredDnsUpdatesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rdu.ref.InternalTokens()
}

func (rdu RequiredDnsUpdatesAttributes) CheckTime() terra.StringValue {
	return terra.ReferenceAsString(rdu.ref.Append("check_time"))
}

func (rdu RequiredDnsUpdatesAttributes) Desired() terra.ListValue[RequiredDnsUpdatesDesiredAttributes] {
	return terra.ReferenceAsList[RequiredDnsUpdatesDesiredAttributes](rdu.ref.Append("desired"))
}

func (rdu RequiredDnsUpdatesAttributes) Discovered() terra.ListValue[RequiredDnsUpdatesDiscoveredAttributes] {
	return terra.ReferenceAsList[RequiredDnsUpdatesDiscoveredAttributes](rdu.ref.Append("discovered"))
}

type RequiredDnsUpdatesDesiredAttributes struct {
	ref terra.Reference
}

func (d RequiredDnsUpdatesDesiredAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d RequiredDnsUpdatesDesiredAttributes) InternalWithRef(ref terra.Reference) RequiredDnsUpdatesDesiredAttributes {
	return RequiredDnsUpdatesDesiredAttributes{ref: ref}
}

func (d RequiredDnsUpdatesDesiredAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d RequiredDnsUpdatesDesiredAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("domain_name"))
}

func (d RequiredDnsUpdatesDesiredAttributes) Records() terra.ListValue[RequiredDnsUpdatesDesiredRecordsAttributes] {
	return terra.ReferenceAsList[RequiredDnsUpdatesDesiredRecordsAttributes](d.ref.Append("records"))
}

type RequiredDnsUpdatesDesiredRecordsAttributes struct {
	ref terra.Reference
}

func (r RequiredDnsUpdatesDesiredRecordsAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RequiredDnsUpdatesDesiredRecordsAttributes) InternalWithRef(ref terra.Reference) RequiredDnsUpdatesDesiredRecordsAttributes {
	return RequiredDnsUpdatesDesiredRecordsAttributes{ref: ref}
}

func (r RequiredDnsUpdatesDesiredRecordsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RequiredDnsUpdatesDesiredRecordsAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("domain_name"))
}

func (r RequiredDnsUpdatesDesiredRecordsAttributes) Rdata() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("rdata"))
}

func (r RequiredDnsUpdatesDesiredRecordsAttributes) RequiredAction() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("required_action"))
}

func (r RequiredDnsUpdatesDesiredRecordsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("type"))
}

type RequiredDnsUpdatesDiscoveredAttributes struct {
	ref terra.Reference
}

func (d RequiredDnsUpdatesDiscoveredAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d RequiredDnsUpdatesDiscoveredAttributes) InternalWithRef(ref terra.Reference) RequiredDnsUpdatesDiscoveredAttributes {
	return RequiredDnsUpdatesDiscoveredAttributes{ref: ref}
}

func (d RequiredDnsUpdatesDiscoveredAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d RequiredDnsUpdatesDiscoveredAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("domain_name"))
}

func (d RequiredDnsUpdatesDiscoveredAttributes) Records() terra.ListValue[RequiredDnsUpdatesDiscoveredRecordsAttributes] {
	return terra.ReferenceAsList[RequiredDnsUpdatesDiscoveredRecordsAttributes](d.ref.Append("records"))
}

type RequiredDnsUpdatesDiscoveredRecordsAttributes struct {
	ref terra.Reference
}

func (r RequiredDnsUpdatesDiscoveredRecordsAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RequiredDnsUpdatesDiscoveredRecordsAttributes) InternalWithRef(ref terra.Reference) RequiredDnsUpdatesDiscoveredRecordsAttributes {
	return RequiredDnsUpdatesDiscoveredRecordsAttributes{ref: ref}
}

func (r RequiredDnsUpdatesDiscoveredRecordsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RequiredDnsUpdatesDiscoveredRecordsAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("domain_name"))
}

func (r RequiredDnsUpdatesDiscoveredRecordsAttributes) Rdata() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("rdata"))
}

func (r RequiredDnsUpdatesDiscoveredRecordsAttributes) RequiredAction() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("required_action"))
}

func (r RequiredDnsUpdatesDiscoveredRecordsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("type"))
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

type CertState struct {
	State        string                  `json:"state"`
	Type         string                  `json:"type"`
	Verification []CertVerificationState `json:"verification"`
}

type CertVerificationState struct {
	Dns  []CertVerificationDnsState  `json:"dns"`
	Http []CertVerificationHttpState `json:"http"`
}

type CertVerificationDnsState struct {
	CheckTime  string                               `json:"check_time"`
	Desired    []CertVerificationDnsDesiredState    `json:"desired"`
	Discovered []CertVerificationDnsDiscoveredState `json:"discovered"`
}

type CertVerificationDnsDesiredState struct {
	DomainName string                          `json:"domain_name"`
	Records    []CertVerificationDecc26A4State `json:"records"`
}

type CertVerificationDecc26A4State struct {
	DomainName     string `json:"domain_name"`
	Rdata          string `json:"rdata"`
	RequiredAction string `json:"required_action"`
	Type           string `json:"type"`
}

type CertVerificationDnsDiscoveredState struct {
	DomainName string                          `json:"domain_name"`
	Records    []CertVerification6F7683B8State `json:"records"`
}

type CertVerification6F7683B8State struct {
	DomainName     string `json:"domain_name"`
	Rdata          string `json:"rdata"`
	RequiredAction string `json:"required_action"`
	Type           string `json:"type"`
}

type CertVerificationHttpState struct {
	Desired       string `json:"desired"`
	Discovered    string `json:"discovered"`
	LastCheckTime string `json:"last_check_time"`
	Path          string `json:"path"`
}

type IssuesState struct {
	Code    float64 `json:"code"`
	Details string  `json:"details"`
	Message string  `json:"message"`
}

type RequiredDnsUpdatesState struct {
	CheckTime  string                              `json:"check_time"`
	Desired    []RequiredDnsUpdatesDesiredState    `json:"desired"`
	Discovered []RequiredDnsUpdatesDiscoveredState `json:"discovered"`
}

type RequiredDnsUpdatesDesiredState struct {
	DomainName string                                  `json:"domain_name"`
	Records    []RequiredDnsUpdatesDesiredRecordsState `json:"records"`
}

type RequiredDnsUpdatesDesiredRecordsState struct {
	DomainName     string `json:"domain_name"`
	Rdata          string `json:"rdata"`
	RequiredAction string `json:"required_action"`
	Type           string `json:"type"`
}

type RequiredDnsUpdatesDiscoveredState struct {
	DomainName string                                     `json:"domain_name"`
	Records    []RequiredDnsUpdatesDiscoveredRecordsState `json:"records"`
}

type RequiredDnsUpdatesDiscoveredRecordsState struct {
	DomainName     string `json:"domain_name"`
	Rdata          string `json:"rdata"`
	RequiredAction string `json:"required_action"`
	Type           string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

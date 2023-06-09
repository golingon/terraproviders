// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package opensearchoutboundconnection

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type LocalDomainInfo struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// OwnerId: string, required
	OwnerId terra.StringValue `hcl:"owner_id,attr" validate:"required"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
}

type RemoteDomainInfo struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// OwnerId: string, required
	OwnerId terra.StringValue `hcl:"owner_id,attr" validate:"required"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type LocalDomainInfoAttributes struct {
	ref terra.Reference
}

func (ldi LocalDomainInfoAttributes) InternalRef() (terra.Reference, error) {
	return ldi.ref, nil
}

func (ldi LocalDomainInfoAttributes) InternalWithRef(ref terra.Reference) LocalDomainInfoAttributes {
	return LocalDomainInfoAttributes{ref: ref}
}

func (ldi LocalDomainInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ldi.ref.InternalTokens()
}

func (ldi LocalDomainInfoAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(ldi.ref.Append("domain_name"))
}

func (ldi LocalDomainInfoAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(ldi.ref.Append("owner_id"))
}

func (ldi LocalDomainInfoAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ldi.ref.Append("region"))
}

type RemoteDomainInfoAttributes struct {
	ref terra.Reference
}

func (rdi RemoteDomainInfoAttributes) InternalRef() (terra.Reference, error) {
	return rdi.ref, nil
}

func (rdi RemoteDomainInfoAttributes) InternalWithRef(ref terra.Reference) RemoteDomainInfoAttributes {
	return RemoteDomainInfoAttributes{ref: ref}
}

func (rdi RemoteDomainInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rdi.ref.InternalTokens()
}

func (rdi RemoteDomainInfoAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(rdi.ref.Append("domain_name"))
}

func (rdi RemoteDomainInfoAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(rdi.ref.Append("owner_id"))
}

func (rdi RemoteDomainInfoAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(rdi.ref.Append("region"))
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

type LocalDomainInfoState struct {
	DomainName string `json:"domain_name"`
	OwnerId    string `json:"owner_id"`
	Region     string `json:"region"`
}

type RemoteDomainInfoState struct {
	DomainName string `json:"domain_name"`
	OwnerId    string `json:"owner_id"`
	Region     string `json:"region"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

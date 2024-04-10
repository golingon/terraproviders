// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package lightsailcertificate

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DomainValidationOptions struct{}

type DomainValidationOptionsAttributes struct {
	ref terra.Reference
}

func (dvo DomainValidationOptionsAttributes) InternalRef() (terra.Reference, error) {
	return dvo.ref, nil
}

func (dvo DomainValidationOptionsAttributes) InternalWithRef(ref terra.Reference) DomainValidationOptionsAttributes {
	return DomainValidationOptionsAttributes{ref: ref}
}

func (dvo DomainValidationOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dvo.ref.InternalTokens()
}

func (dvo DomainValidationOptionsAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(dvo.ref.Append("domain_name"))
}

func (dvo DomainValidationOptionsAttributes) ResourceRecordName() terra.StringValue {
	return terra.ReferenceAsString(dvo.ref.Append("resource_record_name"))
}

func (dvo DomainValidationOptionsAttributes) ResourceRecordType() terra.StringValue {
	return terra.ReferenceAsString(dvo.ref.Append("resource_record_type"))
}

func (dvo DomainValidationOptionsAttributes) ResourceRecordValue() terra.StringValue {
	return terra.ReferenceAsString(dvo.ref.Append("resource_record_value"))
}

type DomainValidationOptionsState struct {
	DomainName          string `json:"domain_name"`
	ResourceRecordName  string `json:"resource_record_name"`
	ResourceRecordType  string `json:"resource_record_type"`
	ResourceRecordValue string `json:"resource_record_value"`
}

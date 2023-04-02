// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apprunnercustomdomainassociation

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CertificateValidationRecords struct{}

type CertificateValidationRecordsAttributes struct {
	ref terra.Reference
}

func (cvr CertificateValidationRecordsAttributes) InternalRef() (terra.Reference, error) {
	return cvr.ref, nil
}

func (cvr CertificateValidationRecordsAttributes) InternalWithRef(ref terra.Reference) CertificateValidationRecordsAttributes {
	return CertificateValidationRecordsAttributes{ref: ref}
}

func (cvr CertificateValidationRecordsAttributes) InternalTokens() hclwrite.Tokens {
	return cvr.ref.InternalTokens()
}

func (cvr CertificateValidationRecordsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cvr.ref.Append("name"))
}

func (cvr CertificateValidationRecordsAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(cvr.ref.Append("status"))
}

func (cvr CertificateValidationRecordsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cvr.ref.Append("type"))
}

func (cvr CertificateValidationRecordsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(cvr.ref.Append("value"))
}

type CertificateValidationRecordsState struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datasignersigningprofile

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RevocationRecord struct{}

type SignatureValidityPeriod struct{}

type RevocationRecordAttributes struct {
	ref terra.Reference
}

func (rr RevocationRecordAttributes) InternalRef() terra.Reference {
	return rr.ref
}

func (rr RevocationRecordAttributes) InternalWithRef(ref terra.Reference) RevocationRecordAttributes {
	return RevocationRecordAttributes{ref: ref}
}

func (rr RevocationRecordAttributes) InternalTokens() hclwrite.Tokens {
	return rr.ref.InternalTokens()
}

func (rr RevocationRecordAttributes) RevocationEffectiveFrom() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("revocation_effective_from"))
}

func (rr RevocationRecordAttributes) RevokedAt() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("revoked_at"))
}

func (rr RevocationRecordAttributes) RevokedBy() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("revoked_by"))
}

type SignatureValidityPeriodAttributes struct {
	ref terra.Reference
}

func (svp SignatureValidityPeriodAttributes) InternalRef() terra.Reference {
	return svp.ref
}

func (svp SignatureValidityPeriodAttributes) InternalWithRef(ref terra.Reference) SignatureValidityPeriodAttributes {
	return SignatureValidityPeriodAttributes{ref: ref}
}

func (svp SignatureValidityPeriodAttributes) InternalTokens() hclwrite.Tokens {
	return svp.ref.InternalTokens()
}

func (svp SignatureValidityPeriodAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(svp.ref.Append("type"))
}

func (svp SignatureValidityPeriodAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(svp.ref.Append("value"))
}

type RevocationRecordState struct {
	RevocationEffectiveFrom string `json:"revocation_effective_from"`
	RevokedAt               string `json:"revoked_at"`
	RevokedBy               string `json:"revoked_by"`
}

type SignatureValidityPeriodState struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

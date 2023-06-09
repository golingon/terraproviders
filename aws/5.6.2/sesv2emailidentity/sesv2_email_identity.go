// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sesv2emailidentity

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DkimSigningAttributes struct {
	// DomainSigningPrivateKey: string, optional
	DomainSigningPrivateKey terra.StringValue `hcl:"domain_signing_private_key,attr"`
	// DomainSigningSelector: string, optional
	DomainSigningSelector terra.StringValue `hcl:"domain_signing_selector,attr"`
	// NextSigningKeyLength: string, optional
	NextSigningKeyLength terra.StringValue `hcl:"next_signing_key_length,attr"`
}

type DkimSigningAttributesAttributes struct {
	ref terra.Reference
}

func (dsa DkimSigningAttributesAttributes) InternalRef() (terra.Reference, error) {
	return dsa.ref, nil
}

func (dsa DkimSigningAttributesAttributes) InternalWithRef(ref terra.Reference) DkimSigningAttributesAttributes {
	return DkimSigningAttributesAttributes{ref: ref}
}

func (dsa DkimSigningAttributesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dsa.ref.InternalTokens()
}

func (dsa DkimSigningAttributesAttributes) CurrentSigningKeyLength() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("current_signing_key_length"))
}

func (dsa DkimSigningAttributesAttributes) DomainSigningPrivateKey() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("domain_signing_private_key"))
}

func (dsa DkimSigningAttributesAttributes) DomainSigningSelector() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("domain_signing_selector"))
}

func (dsa DkimSigningAttributesAttributes) LastKeyGenerationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("last_key_generation_timestamp"))
}

func (dsa DkimSigningAttributesAttributes) NextSigningKeyLength() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("next_signing_key_length"))
}

func (dsa DkimSigningAttributesAttributes) SigningAttributesOrigin() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("signing_attributes_origin"))
}

func (dsa DkimSigningAttributesAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("status"))
}

func (dsa DkimSigningAttributesAttributes) Tokens() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dsa.ref.Append("tokens"))
}

type DkimSigningAttributesState struct {
	CurrentSigningKeyLength    string   `json:"current_signing_key_length"`
	DomainSigningPrivateKey    string   `json:"domain_signing_private_key"`
	DomainSigningSelector      string   `json:"domain_signing_selector"`
	LastKeyGenerationTimestamp string   `json:"last_key_generation_timestamp"`
	NextSigningKeyLength       string   `json:"next_signing_key_length"`
	SigningAttributesOrigin    string   `json:"signing_attributes_origin"`
	Status                     string   `json:"status"`
	Tokens                     []string `json:"tokens"`
}

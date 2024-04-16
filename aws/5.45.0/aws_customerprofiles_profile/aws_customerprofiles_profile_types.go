// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_customerprofiles_profile

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Address struct {
	// Address1: string, optional
	Address1 terra.StringValue `hcl:"address_1,attr"`
	// Address2: string, optional
	Address2 terra.StringValue `hcl:"address_2,attr"`
	// Address3: string, optional
	Address3 terra.StringValue `hcl:"address_3,attr"`
	// Address4: string, optional
	Address4 terra.StringValue `hcl:"address_4,attr"`
	// City: string, optional
	City terra.StringValue `hcl:"city,attr"`
	// Country: string, optional
	Country terra.StringValue `hcl:"country,attr"`
	// County: string, optional
	County terra.StringValue `hcl:"county,attr"`
	// PostalCode: string, optional
	PostalCode terra.StringValue `hcl:"postal_code,attr"`
	// Province: string, optional
	Province terra.StringValue `hcl:"province,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
}

type BillingAddress struct {
	// Address1: string, optional
	Address1 terra.StringValue `hcl:"address_1,attr"`
	// Address2: string, optional
	Address2 terra.StringValue `hcl:"address_2,attr"`
	// Address3: string, optional
	Address3 terra.StringValue `hcl:"address_3,attr"`
	// Address4: string, optional
	Address4 terra.StringValue `hcl:"address_4,attr"`
	// City: string, optional
	City terra.StringValue `hcl:"city,attr"`
	// Country: string, optional
	Country terra.StringValue `hcl:"country,attr"`
	// County: string, optional
	County terra.StringValue `hcl:"county,attr"`
	// PostalCode: string, optional
	PostalCode terra.StringValue `hcl:"postal_code,attr"`
	// Province: string, optional
	Province terra.StringValue `hcl:"province,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
}

type MailingAddress struct {
	// Address1: string, optional
	Address1 terra.StringValue `hcl:"address_1,attr"`
	// Address2: string, optional
	Address2 terra.StringValue `hcl:"address_2,attr"`
	// Address3: string, optional
	Address3 terra.StringValue `hcl:"address_3,attr"`
	// Address4: string, optional
	Address4 terra.StringValue `hcl:"address_4,attr"`
	// City: string, optional
	City terra.StringValue `hcl:"city,attr"`
	// Country: string, optional
	Country terra.StringValue `hcl:"country,attr"`
	// County: string, optional
	County terra.StringValue `hcl:"county,attr"`
	// PostalCode: string, optional
	PostalCode terra.StringValue `hcl:"postal_code,attr"`
	// Province: string, optional
	Province terra.StringValue `hcl:"province,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
}

type ShippingAddress struct {
	// Address1: string, optional
	Address1 terra.StringValue `hcl:"address_1,attr"`
	// Address2: string, optional
	Address2 terra.StringValue `hcl:"address_2,attr"`
	// Address3: string, optional
	Address3 terra.StringValue `hcl:"address_3,attr"`
	// Address4: string, optional
	Address4 terra.StringValue `hcl:"address_4,attr"`
	// City: string, optional
	City terra.StringValue `hcl:"city,attr"`
	// Country: string, optional
	Country terra.StringValue `hcl:"country,attr"`
	// County: string, optional
	County terra.StringValue `hcl:"county,attr"`
	// PostalCode: string, optional
	PostalCode terra.StringValue `hcl:"postal_code,attr"`
	// Province: string, optional
	Province terra.StringValue `hcl:"province,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
}

type AddressAttributes struct {
	ref terra.Reference
}

func (a AddressAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AddressAttributes) InternalWithRef(ref terra.Reference) AddressAttributes {
	return AddressAttributes{ref: ref}
}

func (a AddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AddressAttributes) Address1() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("address_1"))
}

func (a AddressAttributes) Address2() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("address_2"))
}

func (a AddressAttributes) Address3() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("address_3"))
}

func (a AddressAttributes) Address4() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("address_4"))
}

func (a AddressAttributes) City() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("city"))
}

func (a AddressAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("country"))
}

func (a AddressAttributes) County() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("county"))
}

func (a AddressAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("postal_code"))
}

func (a AddressAttributes) Province() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("province"))
}

func (a AddressAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("state"))
}

type BillingAddressAttributes struct {
	ref terra.Reference
}

func (ba BillingAddressAttributes) InternalRef() (terra.Reference, error) {
	return ba.ref, nil
}

func (ba BillingAddressAttributes) InternalWithRef(ref terra.Reference) BillingAddressAttributes {
	return BillingAddressAttributes{ref: ref}
}

func (ba BillingAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ba.ref.InternalTokens()
}

func (ba BillingAddressAttributes) Address1() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("address_1"))
}

func (ba BillingAddressAttributes) Address2() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("address_2"))
}

func (ba BillingAddressAttributes) Address3() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("address_3"))
}

func (ba BillingAddressAttributes) Address4() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("address_4"))
}

func (ba BillingAddressAttributes) City() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("city"))
}

func (ba BillingAddressAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("country"))
}

func (ba BillingAddressAttributes) County() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("county"))
}

func (ba BillingAddressAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("postal_code"))
}

func (ba BillingAddressAttributes) Province() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("province"))
}

func (ba BillingAddressAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("state"))
}

type MailingAddressAttributes struct {
	ref terra.Reference
}

func (ma MailingAddressAttributes) InternalRef() (terra.Reference, error) {
	return ma.ref, nil
}

func (ma MailingAddressAttributes) InternalWithRef(ref terra.Reference) MailingAddressAttributes {
	return MailingAddressAttributes{ref: ref}
}

func (ma MailingAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ma.ref.InternalTokens()
}

func (ma MailingAddressAttributes) Address1() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("address_1"))
}

func (ma MailingAddressAttributes) Address2() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("address_2"))
}

func (ma MailingAddressAttributes) Address3() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("address_3"))
}

func (ma MailingAddressAttributes) Address4() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("address_4"))
}

func (ma MailingAddressAttributes) City() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("city"))
}

func (ma MailingAddressAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("country"))
}

func (ma MailingAddressAttributes) County() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("county"))
}

func (ma MailingAddressAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("postal_code"))
}

func (ma MailingAddressAttributes) Province() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("province"))
}

func (ma MailingAddressAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("state"))
}

type ShippingAddressAttributes struct {
	ref terra.Reference
}

func (sa ShippingAddressAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa ShippingAddressAttributes) InternalWithRef(ref terra.Reference) ShippingAddressAttributes {
	return ShippingAddressAttributes{ref: ref}
}

func (sa ShippingAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa ShippingAddressAttributes) Address1() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("address_1"))
}

func (sa ShippingAddressAttributes) Address2() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("address_2"))
}

func (sa ShippingAddressAttributes) Address3() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("address_3"))
}

func (sa ShippingAddressAttributes) Address4() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("address_4"))
}

func (sa ShippingAddressAttributes) City() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("city"))
}

func (sa ShippingAddressAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("country"))
}

func (sa ShippingAddressAttributes) County() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("county"))
}

func (sa ShippingAddressAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("postal_code"))
}

func (sa ShippingAddressAttributes) Province() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("province"))
}

func (sa ShippingAddressAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("state"))
}

type AddressState struct {
	Address1   string `json:"address_1"`
	Address2   string `json:"address_2"`
	Address3   string `json:"address_3"`
	Address4   string `json:"address_4"`
	City       string `json:"city"`
	Country    string `json:"country"`
	County     string `json:"county"`
	PostalCode string `json:"postal_code"`
	Province   string `json:"province"`
	State      string `json:"state"`
}

type BillingAddressState struct {
	Address1   string `json:"address_1"`
	Address2   string `json:"address_2"`
	Address3   string `json:"address_3"`
	Address4   string `json:"address_4"`
	City       string `json:"city"`
	Country    string `json:"country"`
	County     string `json:"county"`
	PostalCode string `json:"postal_code"`
	Province   string `json:"province"`
	State      string `json:"state"`
}

type MailingAddressState struct {
	Address1   string `json:"address_1"`
	Address2   string `json:"address_2"`
	Address3   string `json:"address_3"`
	Address4   string `json:"address_4"`
	City       string `json:"city"`
	Country    string `json:"country"`
	County     string `json:"county"`
	PostalCode string `json:"postal_code"`
	Province   string `json:"province"`
	State      string `json:"state"`
}

type ShippingAddressState struct {
	Address1   string `json:"address_1"`
	Address2   string `json:"address_2"`
	Address3   string `json:"address_3"`
	Address4   string `json:"address_4"`
	City       string `json:"city"`
	Country    string `json:"country"`
	County     string `json:"county"`
	PostalCode string `json:"postal_code"`
	Province   string `json:"province"`
	State      string `json:"state"`
}

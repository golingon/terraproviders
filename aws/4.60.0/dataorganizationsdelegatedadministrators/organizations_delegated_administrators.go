// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataorganizationsdelegatedadministrators

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DelegatedAdministrators struct{}

type DelegatedAdministratorsAttributes struct {
	ref terra.Reference
}

func (da DelegatedAdministratorsAttributes) InternalRef() terra.Reference {
	return da.ref
}

func (da DelegatedAdministratorsAttributes) InternalWithRef(ref terra.Reference) DelegatedAdministratorsAttributes {
	return DelegatedAdministratorsAttributes{ref: ref}
}

func (da DelegatedAdministratorsAttributes) InternalTokens() hclwrite.Tokens {
	return da.ref.InternalTokens()
}

func (da DelegatedAdministratorsAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("arn"))
}

func (da DelegatedAdministratorsAttributes) DelegationEnabledDate() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("delegation_enabled_date"))
}

func (da DelegatedAdministratorsAttributes) Email() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("email"))
}

func (da DelegatedAdministratorsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("id"))
}

func (da DelegatedAdministratorsAttributes) JoinedMethod() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("joined_method"))
}

func (da DelegatedAdministratorsAttributes) JoinedTimestamp() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("joined_timestamp"))
}

func (da DelegatedAdministratorsAttributes) Name() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("name"))
}

func (da DelegatedAdministratorsAttributes) Status() terra.StringValue {
	return terra.ReferenceString(da.ref.Append("status"))
}

type DelegatedAdministratorsState struct {
	Arn                   string `json:"arn"`
	DelegationEnabledDate string `json:"delegation_enabled_date"`
	Email                 string `json:"email"`
	Id                    string `json:"id"`
	JoinedMethod          string `json:"joined_method"`
	JoinedTimestamp       string `json:"joined_timestamp"`
	Name                  string `json:"name"`
	Status                string `json:"status"`
}

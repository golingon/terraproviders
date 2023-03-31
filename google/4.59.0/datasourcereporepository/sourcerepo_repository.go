// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datasourcereporepository

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PubsubConfigs struct{}

type PubsubConfigsAttributes struct {
	ref terra.Reference
}

func (pc PubsubConfigsAttributes) InternalRef() terra.Reference {
	return pc.ref
}

func (pc PubsubConfigsAttributes) InternalWithRef(ref terra.Reference) PubsubConfigsAttributes {
	return PubsubConfigsAttributes{ref: ref}
}

func (pc PubsubConfigsAttributes) InternalTokens() hclwrite.Tokens {
	return pc.ref.InternalTokens()
}

func (pc PubsubConfigsAttributes) MessageFormat() terra.StringValue {
	return terra.ReferenceString(pc.ref.Append("message_format"))
}

func (pc PubsubConfigsAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceString(pc.ref.Append("service_account_email"))
}

func (pc PubsubConfigsAttributes) Topic() terra.StringValue {
	return terra.ReferenceString(pc.ref.Append("topic"))
}

type PubsubConfigsState struct {
	MessageFormat       string `json:"message_format"`
	ServiceAccountEmail string `json:"service_account_email"`
	Topic               string `json:"topic"`
}

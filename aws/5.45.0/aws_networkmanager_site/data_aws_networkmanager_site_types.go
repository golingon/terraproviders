// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_networkmanager_site

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataLocationAttributes struct {
	ref terra.Reference
}

func (l DataLocationAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l DataLocationAttributes) InternalWithRef(ref terra.Reference) DataLocationAttributes {
	return DataLocationAttributes{ref: ref}
}

func (l DataLocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l DataLocationAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("address"))
}

func (l DataLocationAttributes) Latitude() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("latitude"))
}

func (l DataLocationAttributes) Longitude() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("longitude"))
}

type DataLocationState struct {
	Address   string `json:"address"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

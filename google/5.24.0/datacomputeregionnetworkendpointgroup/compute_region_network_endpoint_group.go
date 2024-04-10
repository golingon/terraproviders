// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datacomputeregionnetworkendpointgroup

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AppEngine struct{}

type CloudFunction struct{}

type CloudRun struct{}

type AppEngineAttributes struct {
	ref terra.Reference
}

func (ae AppEngineAttributes) InternalRef() (terra.Reference, error) {
	return ae.ref, nil
}

func (ae AppEngineAttributes) InternalWithRef(ref terra.Reference) AppEngineAttributes {
	return AppEngineAttributes{ref: ref}
}

func (ae AppEngineAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ae.ref.InternalTokens()
}

func (ae AppEngineAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("service"))
}

func (ae AppEngineAttributes) UrlMask() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("url_mask"))
}

func (ae AppEngineAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("version"))
}

type CloudFunctionAttributes struct {
	ref terra.Reference
}

func (cf CloudFunctionAttributes) InternalRef() (terra.Reference, error) {
	return cf.ref, nil
}

func (cf CloudFunctionAttributes) InternalWithRef(ref terra.Reference) CloudFunctionAttributes {
	return CloudFunctionAttributes{ref: ref}
}

func (cf CloudFunctionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cf.ref.InternalTokens()
}

func (cf CloudFunctionAttributes) Function() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("function"))
}

func (cf CloudFunctionAttributes) UrlMask() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("url_mask"))
}

type CloudRunAttributes struct {
	ref terra.Reference
}

func (cr CloudRunAttributes) InternalRef() (terra.Reference, error) {
	return cr.ref, nil
}

func (cr CloudRunAttributes) InternalWithRef(ref terra.Reference) CloudRunAttributes {
	return CloudRunAttributes{ref: ref}
}

func (cr CloudRunAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cr.ref.InternalTokens()
}

func (cr CloudRunAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("service"))
}

func (cr CloudRunAttributes) Tag() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("tag"))
}

func (cr CloudRunAttributes) UrlMask() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("url_mask"))
}

type AppEngineState struct {
	Service string `json:"service"`
	UrlMask string `json:"url_mask"`
	Version string `json:"version"`
}

type CloudFunctionState struct {
	Function string `json:"function"`
	UrlMask  string `json:"url_mask"`
}

type CloudRunState struct {
	Service string `json:"service"`
	Tag     string `json:"tag"`
	UrlMask string `json:"url_mask"`
}

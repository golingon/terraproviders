// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package databeyondcorpappconnection

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ApplicationEndpoint struct{}

type Gateway struct{}

type ApplicationEndpointAttributes struct {
	ref terra.Reference
}

func (ae ApplicationEndpointAttributes) InternalRef() (terra.Reference, error) {
	return ae.ref, nil
}

func (ae ApplicationEndpointAttributes) InternalWithRef(ref terra.Reference) ApplicationEndpointAttributes {
	return ApplicationEndpointAttributes{ref: ref}
}

func (ae ApplicationEndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ae.ref.InternalTokens()
}

func (ae ApplicationEndpointAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("host"))
}

func (ae ApplicationEndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ae.ref.Append("port"))
}

type GatewayAttributes struct {
	ref terra.Reference
}

func (g GatewayAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GatewayAttributes) InternalWithRef(ref terra.Reference) GatewayAttributes {
	return GatewayAttributes{ref: ref}
}

func (g GatewayAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GatewayAttributes) AppGateway() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("app_gateway"))
}

func (g GatewayAttributes) IngressPort() terra.NumberValue {
	return terra.ReferenceAsNumber(g.ref.Append("ingress_port"))
}

func (g GatewayAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("type"))
}

func (g GatewayAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("uri"))
}

type ApplicationEndpointState struct {
	Host string  `json:"host"`
	Port float64 `json:"port"`
}

type GatewayState struct {
	AppGateway  string  `json:"app_gateway"`
	IngressPort float64 `json:"ingress_port"`
	Type        string  `json:"type"`
	Uri         string  `json:"uri"`
}
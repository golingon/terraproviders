// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_health_check

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type GrpcHealthCheck struct {
	// GrpcServiceName: string, optional
	GrpcServiceName terra.StringValue `hcl:"grpc_service_name,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PortName: string, optional
	PortName terra.StringValue `hcl:"port_name,attr"`
	// PortSpecification: string, optional
	PortSpecification terra.StringValue `hcl:"port_specification,attr"`
}

type Http2HealthCheck struct {
	// Host: string, optional
	Host terra.StringValue `hcl:"host,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PortName: string, optional
	PortName terra.StringValue `hcl:"port_name,attr"`
	// PortSpecification: string, optional
	PortSpecification terra.StringValue `hcl:"port_specification,attr"`
	// ProxyHeader: string, optional
	ProxyHeader terra.StringValue `hcl:"proxy_header,attr"`
	// RequestPath: string, optional
	RequestPath terra.StringValue `hcl:"request_path,attr"`
	// Response: string, optional
	Response terra.StringValue `hcl:"response,attr"`
}

type HttpHealthCheck struct {
	// Host: string, optional
	Host terra.StringValue `hcl:"host,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PortName: string, optional
	PortName terra.StringValue `hcl:"port_name,attr"`
	// PortSpecification: string, optional
	PortSpecification terra.StringValue `hcl:"port_specification,attr"`
	// ProxyHeader: string, optional
	ProxyHeader terra.StringValue `hcl:"proxy_header,attr"`
	// RequestPath: string, optional
	RequestPath terra.StringValue `hcl:"request_path,attr"`
	// Response: string, optional
	Response terra.StringValue `hcl:"response,attr"`
}

type HttpsHealthCheck struct {
	// Host: string, optional
	Host terra.StringValue `hcl:"host,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PortName: string, optional
	PortName terra.StringValue `hcl:"port_name,attr"`
	// PortSpecification: string, optional
	PortSpecification terra.StringValue `hcl:"port_specification,attr"`
	// ProxyHeader: string, optional
	ProxyHeader terra.StringValue `hcl:"proxy_header,attr"`
	// RequestPath: string, optional
	RequestPath terra.StringValue `hcl:"request_path,attr"`
	// Response: string, optional
	Response terra.StringValue `hcl:"response,attr"`
}

type LogConfig struct {
	// Enable: bool, optional
	Enable terra.BoolValue `hcl:"enable,attr"`
}

type SslHealthCheck struct {
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PortName: string, optional
	PortName terra.StringValue `hcl:"port_name,attr"`
	// PortSpecification: string, optional
	PortSpecification terra.StringValue `hcl:"port_specification,attr"`
	// ProxyHeader: string, optional
	ProxyHeader terra.StringValue `hcl:"proxy_header,attr"`
	// Request: string, optional
	Request terra.StringValue `hcl:"request,attr"`
	// Response: string, optional
	Response terra.StringValue `hcl:"response,attr"`
}

type TcpHealthCheck struct {
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PortName: string, optional
	PortName terra.StringValue `hcl:"port_name,attr"`
	// PortSpecification: string, optional
	PortSpecification terra.StringValue `hcl:"port_specification,attr"`
	// ProxyHeader: string, optional
	ProxyHeader terra.StringValue `hcl:"proxy_header,attr"`
	// Request: string, optional
	Request terra.StringValue `hcl:"request,attr"`
	// Response: string, optional
	Response terra.StringValue `hcl:"response,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type GrpcHealthCheckAttributes struct {
	ref terra.Reference
}

func (ghc GrpcHealthCheckAttributes) InternalRef() (terra.Reference, error) {
	return ghc.ref, nil
}

func (ghc GrpcHealthCheckAttributes) InternalWithRef(ref terra.Reference) GrpcHealthCheckAttributes {
	return GrpcHealthCheckAttributes{ref: ref}
}

func (ghc GrpcHealthCheckAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ghc.ref.InternalTokens()
}

func (ghc GrpcHealthCheckAttributes) GrpcServiceName() terra.StringValue {
	return terra.ReferenceAsString(ghc.ref.Append("grpc_service_name"))
}

func (ghc GrpcHealthCheckAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ghc.ref.Append("port"))
}

func (ghc GrpcHealthCheckAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(ghc.ref.Append("port_name"))
}

func (ghc GrpcHealthCheckAttributes) PortSpecification() terra.StringValue {
	return terra.ReferenceAsString(ghc.ref.Append("port_specification"))
}

type Http2HealthCheckAttributes struct {
	ref terra.Reference
}

func (hhc Http2HealthCheckAttributes) InternalRef() (terra.Reference, error) {
	return hhc.ref, nil
}

func (hhc Http2HealthCheckAttributes) InternalWithRef(ref terra.Reference) Http2HealthCheckAttributes {
	return Http2HealthCheckAttributes{ref: ref}
}

func (hhc Http2HealthCheckAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hhc.ref.InternalTokens()
}

func (hhc Http2HealthCheckAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("host"))
}

func (hhc Http2HealthCheckAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(hhc.ref.Append("port"))
}

func (hhc Http2HealthCheckAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("port_name"))
}

func (hhc Http2HealthCheckAttributes) PortSpecification() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("port_specification"))
}

func (hhc Http2HealthCheckAttributes) ProxyHeader() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("proxy_header"))
}

func (hhc Http2HealthCheckAttributes) RequestPath() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("request_path"))
}

func (hhc Http2HealthCheckAttributes) Response() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("response"))
}

type HttpHealthCheckAttributes struct {
	ref terra.Reference
}

func (hhc HttpHealthCheckAttributes) InternalRef() (terra.Reference, error) {
	return hhc.ref, nil
}

func (hhc HttpHealthCheckAttributes) InternalWithRef(ref terra.Reference) HttpHealthCheckAttributes {
	return HttpHealthCheckAttributes{ref: ref}
}

func (hhc HttpHealthCheckAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hhc.ref.InternalTokens()
}

func (hhc HttpHealthCheckAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("host"))
}

func (hhc HttpHealthCheckAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(hhc.ref.Append("port"))
}

func (hhc HttpHealthCheckAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("port_name"))
}

func (hhc HttpHealthCheckAttributes) PortSpecification() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("port_specification"))
}

func (hhc HttpHealthCheckAttributes) ProxyHeader() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("proxy_header"))
}

func (hhc HttpHealthCheckAttributes) RequestPath() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("request_path"))
}

func (hhc HttpHealthCheckAttributes) Response() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("response"))
}

type HttpsHealthCheckAttributes struct {
	ref terra.Reference
}

func (hhc HttpsHealthCheckAttributes) InternalRef() (terra.Reference, error) {
	return hhc.ref, nil
}

func (hhc HttpsHealthCheckAttributes) InternalWithRef(ref terra.Reference) HttpsHealthCheckAttributes {
	return HttpsHealthCheckAttributes{ref: ref}
}

func (hhc HttpsHealthCheckAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hhc.ref.InternalTokens()
}

func (hhc HttpsHealthCheckAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("host"))
}

func (hhc HttpsHealthCheckAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(hhc.ref.Append("port"))
}

func (hhc HttpsHealthCheckAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("port_name"))
}

func (hhc HttpsHealthCheckAttributes) PortSpecification() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("port_specification"))
}

func (hhc HttpsHealthCheckAttributes) ProxyHeader() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("proxy_header"))
}

func (hhc HttpsHealthCheckAttributes) RequestPath() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("request_path"))
}

func (hhc HttpsHealthCheckAttributes) Response() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("response"))
}

type LogConfigAttributes struct {
	ref terra.Reference
}

func (lc LogConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc LogConfigAttributes) InternalWithRef(ref terra.Reference) LogConfigAttributes {
	return LogConfigAttributes{ref: ref}
}

func (lc LogConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc LogConfigAttributes) Enable() terra.BoolValue {
	return terra.ReferenceAsBool(lc.ref.Append("enable"))
}

type SslHealthCheckAttributes struct {
	ref terra.Reference
}

func (shc SslHealthCheckAttributes) InternalRef() (terra.Reference, error) {
	return shc.ref, nil
}

func (shc SslHealthCheckAttributes) InternalWithRef(ref terra.Reference) SslHealthCheckAttributes {
	return SslHealthCheckAttributes{ref: ref}
}

func (shc SslHealthCheckAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return shc.ref.InternalTokens()
}

func (shc SslHealthCheckAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(shc.ref.Append("port"))
}

func (shc SslHealthCheckAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("port_name"))
}

func (shc SslHealthCheckAttributes) PortSpecification() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("port_specification"))
}

func (shc SslHealthCheckAttributes) ProxyHeader() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("proxy_header"))
}

func (shc SslHealthCheckAttributes) Request() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("request"))
}

func (shc SslHealthCheckAttributes) Response() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("response"))
}

type TcpHealthCheckAttributes struct {
	ref terra.Reference
}

func (thc TcpHealthCheckAttributes) InternalRef() (terra.Reference, error) {
	return thc.ref, nil
}

func (thc TcpHealthCheckAttributes) InternalWithRef(ref terra.Reference) TcpHealthCheckAttributes {
	return TcpHealthCheckAttributes{ref: ref}
}

func (thc TcpHealthCheckAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return thc.ref.InternalTokens()
}

func (thc TcpHealthCheckAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(thc.ref.Append("port"))
}

func (thc TcpHealthCheckAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(thc.ref.Append("port_name"))
}

func (thc TcpHealthCheckAttributes) PortSpecification() terra.StringValue {
	return terra.ReferenceAsString(thc.ref.Append("port_specification"))
}

func (thc TcpHealthCheckAttributes) ProxyHeader() terra.StringValue {
	return terra.ReferenceAsString(thc.ref.Append("proxy_header"))
}

func (thc TcpHealthCheckAttributes) Request() terra.StringValue {
	return terra.ReferenceAsString(thc.ref.Append("request"))
}

func (thc TcpHealthCheckAttributes) Response() terra.StringValue {
	return terra.ReferenceAsString(thc.ref.Append("response"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type GrpcHealthCheckState struct {
	GrpcServiceName   string  `json:"grpc_service_name"`
	Port              float64 `json:"port"`
	PortName          string  `json:"port_name"`
	PortSpecification string  `json:"port_specification"`
}

type Http2HealthCheckState struct {
	Host              string  `json:"host"`
	Port              float64 `json:"port"`
	PortName          string  `json:"port_name"`
	PortSpecification string  `json:"port_specification"`
	ProxyHeader       string  `json:"proxy_header"`
	RequestPath       string  `json:"request_path"`
	Response          string  `json:"response"`
}

type HttpHealthCheckState struct {
	Host              string  `json:"host"`
	Port              float64 `json:"port"`
	PortName          string  `json:"port_name"`
	PortSpecification string  `json:"port_specification"`
	ProxyHeader       string  `json:"proxy_header"`
	RequestPath       string  `json:"request_path"`
	Response          string  `json:"response"`
}

type HttpsHealthCheckState struct {
	Host              string  `json:"host"`
	Port              float64 `json:"port"`
	PortName          string  `json:"port_name"`
	PortSpecification string  `json:"port_specification"`
	ProxyHeader       string  `json:"proxy_header"`
	RequestPath       string  `json:"request_path"`
	Response          string  `json:"response"`
}

type LogConfigState struct {
	Enable bool `json:"enable"`
}

type SslHealthCheckState struct {
	Port              float64 `json:"port"`
	PortName          string  `json:"port_name"`
	PortSpecification string  `json:"port_specification"`
	ProxyHeader       string  `json:"proxy_header"`
	Request           string  `json:"request"`
	Response          string  `json:"response"`
}

type TcpHealthCheckState struct {
	Port              float64 `json:"port"`
	PortName          string  `json:"port_name"`
	PortSpecification string  `json:"port_specification"`
	ProxyHeader       string  `json:"proxy_header"`
	Request           string  `json:"request"`
	Response          string  `json:"response"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

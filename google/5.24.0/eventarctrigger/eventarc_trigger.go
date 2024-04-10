// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package eventarctrigger

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Destination struct {
	// Workflow: string, optional
	Workflow terra.StringValue `hcl:"workflow,attr"`
	// CloudRunService: optional
	CloudRunService *CloudRunService `hcl:"cloud_run_service,block"`
	// Gke: optional
	Gke *Gke `hcl:"gke,block"`
	// HttpEndpoint: optional
	HttpEndpoint *HttpEndpoint `hcl:"http_endpoint,block"`
	// NetworkConfig: optional
	NetworkConfig *NetworkConfig `hcl:"network_config,block"`
}

type CloudRunService struct {
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
}

type Gke struct {
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
}

type HttpEndpoint struct {
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
}

type NetworkConfig struct {
	// NetworkAttachment: string, required
	NetworkAttachment terra.StringValue `hcl:"network_attachment,attr" validate:"required"`
}

type MatchingCriteria struct {
	// Attribute: string, required
	Attribute terra.StringValue `hcl:"attribute,attr" validate:"required"`
	// Operator: string, optional
	Operator terra.StringValue `hcl:"operator,attr"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type Transport struct {
	// Pubsub: optional
	Pubsub *Pubsub `hcl:"pubsub,block"`
}

type Pubsub struct {
	// Topic: string, optional
	Topic terra.StringValue `hcl:"topic,attr"`
}

type DestinationAttributes struct {
	ref terra.Reference
}

func (d DestinationAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DestinationAttributes) InternalWithRef(ref terra.Reference) DestinationAttributes {
	return DestinationAttributes{ref: ref}
}

func (d DestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DestinationAttributes) CloudFunction() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("cloud_function"))
}

func (d DestinationAttributes) Workflow() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("workflow"))
}

func (d DestinationAttributes) CloudRunService() terra.ListValue[CloudRunServiceAttributes] {
	return terra.ReferenceAsList[CloudRunServiceAttributes](d.ref.Append("cloud_run_service"))
}

func (d DestinationAttributes) Gke() terra.ListValue[GkeAttributes] {
	return terra.ReferenceAsList[GkeAttributes](d.ref.Append("gke"))
}

func (d DestinationAttributes) HttpEndpoint() terra.ListValue[HttpEndpointAttributes] {
	return terra.ReferenceAsList[HttpEndpointAttributes](d.ref.Append("http_endpoint"))
}

func (d DestinationAttributes) NetworkConfig() terra.ListValue[NetworkConfigAttributes] {
	return terra.ReferenceAsList[NetworkConfigAttributes](d.ref.Append("network_config"))
}

type CloudRunServiceAttributes struct {
	ref terra.Reference
}

func (crs CloudRunServiceAttributes) InternalRef() (terra.Reference, error) {
	return crs.ref, nil
}

func (crs CloudRunServiceAttributes) InternalWithRef(ref terra.Reference) CloudRunServiceAttributes {
	return CloudRunServiceAttributes{ref: ref}
}

func (crs CloudRunServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return crs.ref.InternalTokens()
}

func (crs CloudRunServiceAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("path"))
}

func (crs CloudRunServiceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("region"))
}

func (crs CloudRunServiceAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("service"))
}

type GkeAttributes struct {
	ref terra.Reference
}

func (g GkeAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GkeAttributes) InternalWithRef(ref terra.Reference) GkeAttributes {
	return GkeAttributes{ref: ref}
}

func (g GkeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GkeAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("cluster"))
}

func (g GkeAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("location"))
}

func (g GkeAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("namespace"))
}

func (g GkeAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("path"))
}

func (g GkeAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("service"))
}

type HttpEndpointAttributes struct {
	ref terra.Reference
}

func (he HttpEndpointAttributes) InternalRef() (terra.Reference, error) {
	return he.ref, nil
}

func (he HttpEndpointAttributes) InternalWithRef(ref terra.Reference) HttpEndpointAttributes {
	return HttpEndpointAttributes{ref: ref}
}

func (he HttpEndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return he.ref.InternalTokens()
}

func (he HttpEndpointAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(he.ref.Append("uri"))
}

type NetworkConfigAttributes struct {
	ref terra.Reference
}

func (nc NetworkConfigAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NetworkConfigAttributes) InternalWithRef(ref terra.Reference) NetworkConfigAttributes {
	return NetworkConfigAttributes{ref: ref}
}

func (nc NetworkConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NetworkConfigAttributes) NetworkAttachment() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("network_attachment"))
}

type MatchingCriteriaAttributes struct {
	ref terra.Reference
}

func (mc MatchingCriteriaAttributes) InternalRef() (terra.Reference, error) {
	return mc.ref, nil
}

func (mc MatchingCriteriaAttributes) InternalWithRef(ref terra.Reference) MatchingCriteriaAttributes {
	return MatchingCriteriaAttributes{ref: ref}
}

func (mc MatchingCriteriaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mc.ref.InternalTokens()
}

func (mc MatchingCriteriaAttributes) Attribute() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("attribute"))
}

func (mc MatchingCriteriaAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("operator"))
}

func (mc MatchingCriteriaAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("value"))
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

type TransportAttributes struct {
	ref terra.Reference
}

func (t TransportAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TransportAttributes) InternalWithRef(ref terra.Reference) TransportAttributes {
	return TransportAttributes{ref: ref}
}

func (t TransportAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TransportAttributes) Pubsub() terra.ListValue[PubsubAttributes] {
	return terra.ReferenceAsList[PubsubAttributes](t.ref.Append("pubsub"))
}

type PubsubAttributes struct {
	ref terra.Reference
}

func (p PubsubAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PubsubAttributes) InternalWithRef(ref terra.Reference) PubsubAttributes {
	return PubsubAttributes{ref: ref}
}

func (p PubsubAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PubsubAttributes) Subscription() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("subscription"))
}

func (p PubsubAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("topic"))
}

type DestinationState struct {
	CloudFunction   string                 `json:"cloud_function"`
	Workflow        string                 `json:"workflow"`
	CloudRunService []CloudRunServiceState `json:"cloud_run_service"`
	Gke             []GkeState             `json:"gke"`
	HttpEndpoint    []HttpEndpointState    `json:"http_endpoint"`
	NetworkConfig   []NetworkConfigState   `json:"network_config"`
}

type CloudRunServiceState struct {
	Path    string `json:"path"`
	Region  string `json:"region"`
	Service string `json:"service"`
}

type GkeState struct {
	Cluster   string `json:"cluster"`
	Location  string `json:"location"`
	Namespace string `json:"namespace"`
	Path      string `json:"path"`
	Service   string `json:"service"`
}

type HttpEndpointState struct {
	Uri string `json:"uri"`
}

type NetworkConfigState struct {
	NetworkAttachment string `json:"network_attachment"`
}

type MatchingCriteriaState struct {
	Attribute string `json:"attribute"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type TransportState struct {
	Pubsub []PubsubState `json:"pubsub"`
}

type PubsubState struct {
	Subscription string `json:"subscription"`
	Topic        string `json:"topic"`
}

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_blockchain_node_engine_blockchain_nodes

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type EthereumDetails struct {
	// ApiEnableAdmin: bool, optional
	ApiEnableAdmin terra.BoolValue `hcl:"api_enable_admin,attr"`
	// ApiEnableDebug: bool, optional
	ApiEnableDebug terra.BoolValue `hcl:"api_enable_debug,attr"`
	// ConsensusClient: string, optional
	ConsensusClient terra.StringValue `hcl:"consensus_client,attr"`
	// ExecutionClient: string, optional
	ExecutionClient terra.StringValue `hcl:"execution_client,attr"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NodeType: string, optional
	NodeType terra.StringValue `hcl:"node_type,attr"`
	// EthereumDetailsGethDetails: optional
	GethDetails *EthereumDetailsGethDetails `hcl:"geth_details,block"`
	// EthereumDetailsValidatorConfig: optional
	ValidatorConfig *EthereumDetailsValidatorConfig `hcl:"validator_config,block"`
}

type EthereumDetailsGethDetails struct {
	// GarbageCollectionMode: string, optional
	GarbageCollectionMode terra.StringValue `hcl:"garbage_collection_mode,attr"`
}

type EthereumDetailsValidatorConfig struct {
	// MevRelayUrls: list of string, optional
	MevRelayUrls terra.ListValue[terra.StringValue] `hcl:"mev_relay_urls,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ConnectionInfoAttributes struct {
	ref terra.Reference
}

func (ci ConnectionInfoAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci ConnectionInfoAttributes) InternalWithRef(ref terra.Reference) ConnectionInfoAttributes {
	return ConnectionInfoAttributes{ref: ref}
}

func (ci ConnectionInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ci.ref.InternalTokens()
}

func (ci ConnectionInfoAttributes) ServiceAttachment() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("service_attachment"))
}

func (ci ConnectionInfoAttributes) EndpointInfo() terra.ListValue[ConnectionInfoEndpointInfoAttributes] {
	return terra.ReferenceAsList[ConnectionInfoEndpointInfoAttributes](ci.ref.Append("endpoint_info"))
}

type ConnectionInfoEndpointInfoAttributes struct {
	ref terra.Reference
}

func (ei ConnectionInfoEndpointInfoAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei ConnectionInfoEndpointInfoAttributes) InternalWithRef(ref terra.Reference) ConnectionInfoEndpointInfoAttributes {
	return ConnectionInfoEndpointInfoAttributes{ref: ref}
}

func (ei ConnectionInfoEndpointInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei ConnectionInfoEndpointInfoAttributes) JsonRpcApiEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("json_rpc_api_endpoint"))
}

func (ei ConnectionInfoEndpointInfoAttributes) WebsocketsApiEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("websockets_api_endpoint"))
}

type EthereumDetailsAttributes struct {
	ref terra.Reference
}

func (ed EthereumDetailsAttributes) InternalRef() (terra.Reference, error) {
	return ed.ref, nil
}

func (ed EthereumDetailsAttributes) InternalWithRef(ref terra.Reference) EthereumDetailsAttributes {
	return EthereumDetailsAttributes{ref: ref}
}

func (ed EthereumDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ed.ref.InternalTokens()
}

func (ed EthereumDetailsAttributes) ApiEnableAdmin() terra.BoolValue {
	return terra.ReferenceAsBool(ed.ref.Append("api_enable_admin"))
}

func (ed EthereumDetailsAttributes) ApiEnableDebug() terra.BoolValue {
	return terra.ReferenceAsBool(ed.ref.Append("api_enable_debug"))
}

func (ed EthereumDetailsAttributes) ConsensusClient() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("consensus_client"))
}

func (ed EthereumDetailsAttributes) ExecutionClient() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("execution_client"))
}

func (ed EthereumDetailsAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("network"))
}

func (ed EthereumDetailsAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("node_type"))
}

func (ed EthereumDetailsAttributes) AdditionalEndpoints() terra.ListValue[EthereumDetailsAdditionalEndpointsAttributes] {
	return terra.ReferenceAsList[EthereumDetailsAdditionalEndpointsAttributes](ed.ref.Append("additional_endpoints"))
}

func (ed EthereumDetailsAttributes) GethDetails() terra.ListValue[EthereumDetailsGethDetailsAttributes] {
	return terra.ReferenceAsList[EthereumDetailsGethDetailsAttributes](ed.ref.Append("geth_details"))
}

func (ed EthereumDetailsAttributes) ValidatorConfig() terra.ListValue[EthereumDetailsValidatorConfigAttributes] {
	return terra.ReferenceAsList[EthereumDetailsValidatorConfigAttributes](ed.ref.Append("validator_config"))
}

type EthereumDetailsAdditionalEndpointsAttributes struct {
	ref terra.Reference
}

func (ae EthereumDetailsAdditionalEndpointsAttributes) InternalRef() (terra.Reference, error) {
	return ae.ref, nil
}

func (ae EthereumDetailsAdditionalEndpointsAttributes) InternalWithRef(ref terra.Reference) EthereumDetailsAdditionalEndpointsAttributes {
	return EthereumDetailsAdditionalEndpointsAttributes{ref: ref}
}

func (ae EthereumDetailsAdditionalEndpointsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ae.ref.InternalTokens()
}

func (ae EthereumDetailsAdditionalEndpointsAttributes) BeaconApiEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("beacon_api_endpoint"))
}

func (ae EthereumDetailsAdditionalEndpointsAttributes) BeaconPrometheusMetricsApiEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("beacon_prometheus_metrics_api_endpoint"))
}

func (ae EthereumDetailsAdditionalEndpointsAttributes) ExecutionClientPrometheusMetricsApiEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("execution_client_prometheus_metrics_api_endpoint"))
}

type EthereumDetailsGethDetailsAttributes struct {
	ref terra.Reference
}

func (gd EthereumDetailsGethDetailsAttributes) InternalRef() (terra.Reference, error) {
	return gd.ref, nil
}

func (gd EthereumDetailsGethDetailsAttributes) InternalWithRef(ref terra.Reference) EthereumDetailsGethDetailsAttributes {
	return EthereumDetailsGethDetailsAttributes{ref: ref}
}

func (gd EthereumDetailsGethDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gd.ref.InternalTokens()
}

func (gd EthereumDetailsGethDetailsAttributes) GarbageCollectionMode() terra.StringValue {
	return terra.ReferenceAsString(gd.ref.Append("garbage_collection_mode"))
}

type EthereumDetailsValidatorConfigAttributes struct {
	ref terra.Reference
}

func (vc EthereumDetailsValidatorConfigAttributes) InternalRef() (terra.Reference, error) {
	return vc.ref, nil
}

func (vc EthereumDetailsValidatorConfigAttributes) InternalWithRef(ref terra.Reference) EthereumDetailsValidatorConfigAttributes {
	return EthereumDetailsValidatorConfigAttributes{ref: ref}
}

func (vc EthereumDetailsValidatorConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vc.ref.InternalTokens()
}

func (vc EthereumDetailsValidatorConfigAttributes) MevRelayUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vc.ref.Append("mev_relay_urls"))
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

type ConnectionInfoState struct {
	ServiceAttachment string                            `json:"service_attachment"`
	EndpointInfo      []ConnectionInfoEndpointInfoState `json:"endpoint_info"`
}

type ConnectionInfoEndpointInfoState struct {
	JsonRpcApiEndpoint    string `json:"json_rpc_api_endpoint"`
	WebsocketsApiEndpoint string `json:"websockets_api_endpoint"`
}

type EthereumDetailsState struct {
	ApiEnableAdmin      bool                                      `json:"api_enable_admin"`
	ApiEnableDebug      bool                                      `json:"api_enable_debug"`
	ConsensusClient     string                                    `json:"consensus_client"`
	ExecutionClient     string                                    `json:"execution_client"`
	Network             string                                    `json:"network"`
	NodeType            string                                    `json:"node_type"`
	AdditionalEndpoints []EthereumDetailsAdditionalEndpointsState `json:"additional_endpoints"`
	GethDetails         []EthereumDetailsGethDetailsState         `json:"geth_details"`
	ValidatorConfig     []EthereumDetailsValidatorConfigState     `json:"validator_config"`
}

type EthereumDetailsAdditionalEndpointsState struct {
	BeaconApiEndpoint                           string `json:"beacon_api_endpoint"`
	BeaconPrometheusMetricsApiEndpoint          string `json:"beacon_prometheus_metrics_api_endpoint"`
	ExecutionClientPrometheusMetricsApiEndpoint string `json:"execution_client_prometheus_metrics_api_endpoint"`
}

type EthereumDetailsGethDetailsState struct {
	GarbageCollectionMode string `json:"garbage_collection_mode"`
}

type EthereumDetailsValidatorConfigState struct {
	MevRelayUrls []string `json:"mev_relay_urls"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

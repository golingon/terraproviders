// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataprivatelinkserviceendpointconnections

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PrivateEndpointConnections struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type PrivateEndpointConnectionsAttributes struct {
	ref terra.Reference
}

func (pec PrivateEndpointConnectionsAttributes) InternalRef() (terra.Reference, error) {
	return pec.ref, nil
}

func (pec PrivateEndpointConnectionsAttributes) InternalWithRef(ref terra.Reference) PrivateEndpointConnectionsAttributes {
	return PrivateEndpointConnectionsAttributes{ref: ref}
}

func (pec PrivateEndpointConnectionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pec.ref.InternalTokens()
}

func (pec PrivateEndpointConnectionsAttributes) ActionRequired() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("action_required"))
}

func (pec PrivateEndpointConnectionsAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("connection_id"))
}

func (pec PrivateEndpointConnectionsAttributes) ConnectionName() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("connection_name"))
}

func (pec PrivateEndpointConnectionsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("description"))
}

func (pec PrivateEndpointConnectionsAttributes) PrivateEndpointId() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("private_endpoint_id"))
}

func (pec PrivateEndpointConnectionsAttributes) PrivateEndpointName() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("private_endpoint_name"))
}

func (pec PrivateEndpointConnectionsAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("status"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type PrivateEndpointConnectionsState struct {
	ActionRequired      string `json:"action_required"`
	ConnectionId        string `json:"connection_id"`
	ConnectionName      string `json:"connection_name"`
	Description         string `json:"description"`
	PrivateEndpointId   string `json:"private_endpoint_id"`
	PrivateEndpointName string `json:"private_endpoint_name"`
	Status              string `json:"status"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}

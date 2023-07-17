// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datadatabricksworkspaceprivateendpointconnection

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Connections struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ConnectionsAttributes struct {
	ref terra.Reference
}

func (c ConnectionsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConnectionsAttributes) InternalWithRef(ref terra.Reference) ConnectionsAttributes {
	return ConnectionsAttributes{ref: ref}
}

func (c ConnectionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConnectionsAttributes) ActionRequired() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("action_required"))
}

func (c ConnectionsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("description"))
}

func (c ConnectionsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ConnectionsAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("status"))
}

func (c ConnectionsAttributes) WorkspacePrivateEndpointId() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("workspace_private_endpoint_id"))
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

type ConnectionsState struct {
	ActionRequired             string `json:"action_required"`
	Description                string `json:"description"`
	Name                       string `json:"name"`
	Status                     string `json:"status"`
	WorkspacePrivateEndpointId string `json:"workspace_private_endpoint_id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
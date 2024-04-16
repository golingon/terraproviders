// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_databricks_workspace_private_endpoint_connection

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataConnectionsAttributes struct {
	ref terra.Reference
}

func (c DataConnectionsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c DataConnectionsAttributes) InternalWithRef(ref terra.Reference) DataConnectionsAttributes {
	return DataConnectionsAttributes{ref: ref}
}

func (c DataConnectionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c DataConnectionsAttributes) ActionRequired() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("action_required"))
}

func (c DataConnectionsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("description"))
}

func (c DataConnectionsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c DataConnectionsAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("status"))
}

func (c DataConnectionsAttributes) WorkspacePrivateEndpointId() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("workspace_private_endpoint_id"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataConnectionsState struct {
	ActionRequired             string `json:"action_required"`
	Description                string `json:"description"`
	Name                       string `json:"name"`
	Status                     string `json:"status"`
	WorkspacePrivateEndpointId string `json:"workspace_private_endpoint_id"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}

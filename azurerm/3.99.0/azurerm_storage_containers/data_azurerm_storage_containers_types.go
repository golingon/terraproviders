// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_storage_containers

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataContainersAttributes struct {
	ref terra.Reference
}

func (c DataContainersAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c DataContainersAttributes) InternalWithRef(ref terra.Reference) DataContainersAttributes {
	return DataContainersAttributes{ref: ref}
}

func (c DataContainersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c DataContainersAttributes) DataPlaneId() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("data_plane_id"))
}

func (c DataContainersAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c DataContainersAttributes) ResourceManagerId() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("resource_manager_id"))
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

type DataContainersState struct {
	DataPlaneId       string `json:"data_plane_id"`
	Name              string `json:"name"`
	ResourceManagerId string `json:"resource_manager_id"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}

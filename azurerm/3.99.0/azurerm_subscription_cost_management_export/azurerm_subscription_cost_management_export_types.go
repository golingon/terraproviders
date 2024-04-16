// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_subscription_cost_management_export

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ExportDataOptions struct {
	// TimeFrame: string, required
	TimeFrame terra.StringValue `hcl:"time_frame,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type ExportDataStorageLocation struct {
	// ContainerId: string, required
	ContainerId terra.StringValue `hcl:"container_id,attr" validate:"required"`
	// RootFolderPath: string, required
	RootFolderPath terra.StringValue `hcl:"root_folder_path,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ExportDataOptionsAttributes struct {
	ref terra.Reference
}

func (edo ExportDataOptionsAttributes) InternalRef() (terra.Reference, error) {
	return edo.ref, nil
}

func (edo ExportDataOptionsAttributes) InternalWithRef(ref terra.Reference) ExportDataOptionsAttributes {
	return ExportDataOptionsAttributes{ref: ref}
}

func (edo ExportDataOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return edo.ref.InternalTokens()
}

func (edo ExportDataOptionsAttributes) TimeFrame() terra.StringValue {
	return terra.ReferenceAsString(edo.ref.Append("time_frame"))
}

func (edo ExportDataOptionsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(edo.ref.Append("type"))
}

type ExportDataStorageLocationAttributes struct {
	ref terra.Reference
}

func (edsl ExportDataStorageLocationAttributes) InternalRef() (terra.Reference, error) {
	return edsl.ref, nil
}

func (edsl ExportDataStorageLocationAttributes) InternalWithRef(ref terra.Reference) ExportDataStorageLocationAttributes {
	return ExportDataStorageLocationAttributes{ref: ref}
}

func (edsl ExportDataStorageLocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return edsl.ref.InternalTokens()
}

func (edsl ExportDataStorageLocationAttributes) ContainerId() terra.StringValue {
	return terra.ReferenceAsString(edsl.ref.Append("container_id"))
}

func (edsl ExportDataStorageLocationAttributes) RootFolderPath() terra.StringValue {
	return terra.ReferenceAsString(edsl.ref.Append("root_folder_path"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ExportDataOptionsState struct {
	TimeFrame string `json:"time_frame"`
	Type      string `json:"type"`
}

type ExportDataStorageLocationState struct {
	ContainerId    string `json:"container_id"`
	RootFolderPath string `json:"root_folder_path"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

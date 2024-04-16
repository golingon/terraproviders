// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_virtual_machine_packet_capture

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Filter struct {
	// LocalIpAddress: string, optional
	LocalIpAddress terra.StringValue `hcl:"local_ip_address,attr"`
	// LocalPort: string, optional
	LocalPort terra.StringValue `hcl:"local_port,attr"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// RemoteIpAddress: string, optional
	RemoteIpAddress terra.StringValue `hcl:"remote_ip_address,attr"`
	// RemotePort: string, optional
	RemotePort terra.StringValue `hcl:"remote_port,attr"`
}

type StorageLocation struct {
	// FilePath: string, optional
	FilePath terra.StringValue `hcl:"file_path,attr"`
	// StorageAccountId: string, optional
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) LocalIpAddress() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("local_ip_address"))
}

func (f FilterAttributes) LocalPort() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("local_port"))
}

func (f FilterAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("protocol"))
}

func (f FilterAttributes) RemoteIpAddress() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("remote_ip_address"))
}

func (f FilterAttributes) RemotePort() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("remote_port"))
}

type StorageLocationAttributes struct {
	ref terra.Reference
}

func (sl StorageLocationAttributes) InternalRef() (terra.Reference, error) {
	return sl.ref, nil
}

func (sl StorageLocationAttributes) InternalWithRef(ref terra.Reference) StorageLocationAttributes {
	return StorageLocationAttributes{ref: ref}
}

func (sl StorageLocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sl.ref.InternalTokens()
}

func (sl StorageLocationAttributes) FilePath() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("file_path"))
}

func (sl StorageLocationAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("storage_account_id"))
}

func (sl StorageLocationAttributes) StoragePath() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("storage_path"))
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

type FilterState struct {
	LocalIpAddress  string `json:"local_ip_address"`
	LocalPort       string `json:"local_port"`
	Protocol        string `json:"protocol"`
	RemoteIpAddress string `json:"remote_ip_address"`
	RemotePort      string `json:"remote_port"`
}

type StorageLocationState struct {
	FilePath         string `json:"file_path"`
	StorageAccountId string `json:"storage_account_id"`
	StoragePath      string `json:"storage_path"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
}

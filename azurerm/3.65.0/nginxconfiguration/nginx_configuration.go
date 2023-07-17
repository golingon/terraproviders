// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package nginxconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ConfigFile struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// VirtualPath: string, required
	VirtualPath terra.StringValue `hcl:"virtual_path,attr" validate:"required"`
}

type ProtectedFile struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// VirtualPath: string, required
	VirtualPath terra.StringValue `hcl:"virtual_path,attr" validate:"required"`
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

type ConfigFileAttributes struct {
	ref terra.Reference
}

func (cf ConfigFileAttributes) InternalRef() (terra.Reference, error) {
	return cf.ref, nil
}

func (cf ConfigFileAttributes) InternalWithRef(ref terra.Reference) ConfigFileAttributes {
	return ConfigFileAttributes{ref: ref}
}

func (cf ConfigFileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cf.ref.InternalTokens()
}

func (cf ConfigFileAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("content"))
}

func (cf ConfigFileAttributes) VirtualPath() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("virtual_path"))
}

type ProtectedFileAttributes struct {
	ref terra.Reference
}

func (pf ProtectedFileAttributes) InternalRef() (terra.Reference, error) {
	return pf.ref, nil
}

func (pf ProtectedFileAttributes) InternalWithRef(ref terra.Reference) ProtectedFileAttributes {
	return ProtectedFileAttributes{ref: ref}
}

func (pf ProtectedFileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pf.ref.InternalTokens()
}

func (pf ProtectedFileAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(pf.ref.Append("content"))
}

func (pf ProtectedFileAttributes) VirtualPath() terra.StringValue {
	return terra.ReferenceAsString(pf.ref.Append("virtual_path"))
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

type ConfigFileState struct {
	Content     string `json:"content"`
	VirtualPath string `json:"virtual_path"`
}

type ProtectedFileState struct {
	Content     string `json:"content"`
	VirtualPath string `json:"virtual_path"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
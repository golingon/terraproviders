// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package elastictranscoderpipeline

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ContentConfig struct {
	// Bucket: string, optional
	Bucket terra.StringValue `hcl:"bucket,attr"`
	// StorageClass: string, optional
	StorageClass terra.StringValue `hcl:"storage_class,attr"`
}

type ContentConfigPermissions struct {
	// Access: list of string, optional
	Access terra.ListValue[terra.StringValue] `hcl:"access,attr"`
	// Grantee: string, optional
	Grantee terra.StringValue `hcl:"grantee,attr"`
	// GranteeType: string, optional
	GranteeType terra.StringValue `hcl:"grantee_type,attr"`
}

type Notifications struct {
	// Completed: string, optional
	Completed terra.StringValue `hcl:"completed,attr"`
	// Error: string, optional
	Error terra.StringValue `hcl:"error,attr"`
	// Progressing: string, optional
	Progressing terra.StringValue `hcl:"progressing,attr"`
	// Warning: string, optional
	Warning terra.StringValue `hcl:"warning,attr"`
}

type ThumbnailConfig struct {
	// Bucket: string, optional
	Bucket terra.StringValue `hcl:"bucket,attr"`
	// StorageClass: string, optional
	StorageClass terra.StringValue `hcl:"storage_class,attr"`
}

type ThumbnailConfigPermissions struct {
	// Access: list of string, optional
	Access terra.ListValue[terra.StringValue] `hcl:"access,attr"`
	// Grantee: string, optional
	Grantee terra.StringValue `hcl:"grantee,attr"`
	// GranteeType: string, optional
	GranteeType terra.StringValue `hcl:"grantee_type,attr"`
}

type ContentConfigAttributes struct {
	ref terra.Reference
}

func (cc ContentConfigAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc ContentConfigAttributes) InternalWithRef(ref terra.Reference) ContentConfigAttributes {
	return ContentConfigAttributes{ref: ref}
}

func (cc ContentConfigAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc ContentConfigAttributes) Bucket() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("bucket"))
}

func (cc ContentConfigAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("storage_class"))
}

type ContentConfigPermissionsAttributes struct {
	ref terra.Reference
}

func (ccp ContentConfigPermissionsAttributes) InternalRef() terra.Reference {
	return ccp.ref
}

func (ccp ContentConfigPermissionsAttributes) InternalWithRef(ref terra.Reference) ContentConfigPermissionsAttributes {
	return ContentConfigPermissionsAttributes{ref: ref}
}

func (ccp ContentConfigPermissionsAttributes) InternalTokens() hclwrite.Tokens {
	return ccp.ref.InternalTokens()
}

func (ccp ContentConfigPermissionsAttributes) Access() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ccp.ref.Append("access"))
}

func (ccp ContentConfigPermissionsAttributes) Grantee() terra.StringValue {
	return terra.ReferenceString(ccp.ref.Append("grantee"))
}

func (ccp ContentConfigPermissionsAttributes) GranteeType() terra.StringValue {
	return terra.ReferenceString(ccp.ref.Append("grantee_type"))
}

type NotificationsAttributes struct {
	ref terra.Reference
}

func (n NotificationsAttributes) InternalRef() terra.Reference {
	return n.ref
}

func (n NotificationsAttributes) InternalWithRef(ref terra.Reference) NotificationsAttributes {
	return NotificationsAttributes{ref: ref}
}

func (n NotificationsAttributes) InternalTokens() hclwrite.Tokens {
	return n.ref.InternalTokens()
}

func (n NotificationsAttributes) Completed() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("completed"))
}

func (n NotificationsAttributes) Error() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("error"))
}

func (n NotificationsAttributes) Progressing() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("progressing"))
}

func (n NotificationsAttributes) Warning() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("warning"))
}

type ThumbnailConfigAttributes struct {
	ref terra.Reference
}

func (tc ThumbnailConfigAttributes) InternalRef() terra.Reference {
	return tc.ref
}

func (tc ThumbnailConfigAttributes) InternalWithRef(ref terra.Reference) ThumbnailConfigAttributes {
	return ThumbnailConfigAttributes{ref: ref}
}

func (tc ThumbnailConfigAttributes) InternalTokens() hclwrite.Tokens {
	return tc.ref.InternalTokens()
}

func (tc ThumbnailConfigAttributes) Bucket() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("bucket"))
}

func (tc ThumbnailConfigAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("storage_class"))
}

type ThumbnailConfigPermissionsAttributes struct {
	ref terra.Reference
}

func (tcp ThumbnailConfigPermissionsAttributes) InternalRef() terra.Reference {
	return tcp.ref
}

func (tcp ThumbnailConfigPermissionsAttributes) InternalWithRef(ref terra.Reference) ThumbnailConfigPermissionsAttributes {
	return ThumbnailConfigPermissionsAttributes{ref: ref}
}

func (tcp ThumbnailConfigPermissionsAttributes) InternalTokens() hclwrite.Tokens {
	return tcp.ref.InternalTokens()
}

func (tcp ThumbnailConfigPermissionsAttributes) Access() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](tcp.ref.Append("access"))
}

func (tcp ThumbnailConfigPermissionsAttributes) Grantee() terra.StringValue {
	return terra.ReferenceString(tcp.ref.Append("grantee"))
}

func (tcp ThumbnailConfigPermissionsAttributes) GranteeType() terra.StringValue {
	return terra.ReferenceString(tcp.ref.Append("grantee_type"))
}

type ContentConfigState struct {
	Bucket       string `json:"bucket"`
	StorageClass string `json:"storage_class"`
}

type ContentConfigPermissionsState struct {
	Access      []string `json:"access"`
	Grantee     string   `json:"grantee"`
	GranteeType string   `json:"grantee_type"`
}

type NotificationsState struct {
	Completed   string `json:"completed"`
	Error       string `json:"error"`
	Progressing string `json:"progressing"`
	Warning     string `json:"warning"`
}

type ThumbnailConfigState struct {
	Bucket       string `json:"bucket"`
	StorageClass string `json:"storage_class"`
}

type ThumbnailConfigPermissionsState struct {
	Access      []string `json:"access"`
	Grantee     string   `json:"grantee"`
	GranteeType string   `json:"grantee_type"`
}

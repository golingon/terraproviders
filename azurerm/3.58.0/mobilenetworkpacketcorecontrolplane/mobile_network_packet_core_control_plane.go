// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mobilenetworkpacketcorecontrolplane

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Identity struct {
	// IdentityIds: set of string, required
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type LocalDiagnosticsAccess struct {
	// AuthenticationType: string, required
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr" validate:"required"`
	// HttpsServerCertificateUrl: string, optional
	HttpsServerCertificateUrl terra.StringValue `hcl:"https_server_certificate_url,attr"`
}

type Platform struct {
	// ArcKubernetesClusterId: string, optional
	ArcKubernetesClusterId terra.StringValue `hcl:"arc_kubernetes_cluster_id,attr"`
	// CustomLocationId: string, optional
	CustomLocationId terra.StringValue `hcl:"custom_location_id,attr"`
	// EdgeDeviceId: string, optional
	EdgeDeviceId terra.StringValue `hcl:"edge_device_id,attr"`
	// StackHciClusterId: string, optional
	StackHciClusterId terra.StringValue `hcl:"stack_hci_cluster_id,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
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

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type LocalDiagnosticsAccessAttributes struct {
	ref terra.Reference
}

func (lda LocalDiagnosticsAccessAttributes) InternalRef() (terra.Reference, error) {
	return lda.ref, nil
}

func (lda LocalDiagnosticsAccessAttributes) InternalWithRef(ref terra.Reference) LocalDiagnosticsAccessAttributes {
	return LocalDiagnosticsAccessAttributes{ref: ref}
}

func (lda LocalDiagnosticsAccessAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lda.ref.InternalTokens()
}

func (lda LocalDiagnosticsAccessAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(lda.ref.Append("authentication_type"))
}

func (lda LocalDiagnosticsAccessAttributes) HttpsServerCertificateUrl() terra.StringValue {
	return terra.ReferenceAsString(lda.ref.Append("https_server_certificate_url"))
}

type PlatformAttributes struct {
	ref terra.Reference
}

func (p PlatformAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PlatformAttributes) InternalWithRef(ref terra.Reference) PlatformAttributes {
	return PlatformAttributes{ref: ref}
}

func (p PlatformAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PlatformAttributes) ArcKubernetesClusterId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("arc_kubernetes_cluster_id"))
}

func (p PlatformAttributes) CustomLocationId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("custom_location_id"))
}

func (p PlatformAttributes) EdgeDeviceId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("edge_device_id"))
}

func (p PlatformAttributes) StackHciClusterId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("stack_hci_cluster_id"))
}

func (p PlatformAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("type"))
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

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	Type        string   `json:"type"`
}

type LocalDiagnosticsAccessState struct {
	AuthenticationType        string `json:"authentication_type"`
	HttpsServerCertificateUrl string `json:"https_server_certificate_url"`
}

type PlatformState struct {
	ArcKubernetesClusterId string `json:"arc_kubernetes_cluster_id"`
	CustomLocationId       string `json:"custom_location_id"`
	EdgeDeviceId           string `json:"edge_device_id"`
	StackHciClusterId      string `json:"stack_hci_cluster_id"`
	Type                   string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

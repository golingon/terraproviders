// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package labserviceplan

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DefaultAutoShutdown struct {
	// DisconnectDelay: string, optional
	DisconnectDelay terra.StringValue `hcl:"disconnect_delay,attr"`
	// IdleDelay: string, optional
	IdleDelay terra.StringValue `hcl:"idle_delay,attr"`
	// NoConnectDelay: string, optional
	NoConnectDelay terra.StringValue `hcl:"no_connect_delay,attr"`
	// ShutdownOnIdle: string, optional
	ShutdownOnIdle terra.StringValue `hcl:"shutdown_on_idle,attr"`
}

type DefaultConnection struct {
	// ClientRdpAccess: string, optional
	ClientRdpAccess terra.StringValue `hcl:"client_rdp_access,attr"`
	// ClientSshAccess: string, optional
	ClientSshAccess terra.StringValue `hcl:"client_ssh_access,attr"`
	// WebRdpAccess: string, optional
	WebRdpAccess terra.StringValue `hcl:"web_rdp_access,attr"`
	// WebSshAccess: string, optional
	WebSshAccess terra.StringValue `hcl:"web_ssh_access,attr"`
}

type Support struct {
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// Instructions: string, optional
	Instructions terra.StringValue `hcl:"instructions,attr"`
	// Phone: string, optional
	Phone terra.StringValue `hcl:"phone,attr"`
	// Url: string, optional
	Url terra.StringValue `hcl:"url,attr"`
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

type DefaultAutoShutdownAttributes struct {
	ref terra.Reference
}

func (das DefaultAutoShutdownAttributes) InternalRef() (terra.Reference, error) {
	return das.ref, nil
}

func (das DefaultAutoShutdownAttributes) InternalWithRef(ref terra.Reference) DefaultAutoShutdownAttributes {
	return DefaultAutoShutdownAttributes{ref: ref}
}

func (das DefaultAutoShutdownAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return das.ref.InternalTokens()
}

func (das DefaultAutoShutdownAttributes) DisconnectDelay() terra.StringValue {
	return terra.ReferenceAsString(das.ref.Append("disconnect_delay"))
}

func (das DefaultAutoShutdownAttributes) IdleDelay() terra.StringValue {
	return terra.ReferenceAsString(das.ref.Append("idle_delay"))
}

func (das DefaultAutoShutdownAttributes) NoConnectDelay() terra.StringValue {
	return terra.ReferenceAsString(das.ref.Append("no_connect_delay"))
}

func (das DefaultAutoShutdownAttributes) ShutdownOnIdle() terra.StringValue {
	return terra.ReferenceAsString(das.ref.Append("shutdown_on_idle"))
}

type DefaultConnectionAttributes struct {
	ref terra.Reference
}

func (dc DefaultConnectionAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DefaultConnectionAttributes) InternalWithRef(ref terra.Reference) DefaultConnectionAttributes {
	return DefaultConnectionAttributes{ref: ref}
}

func (dc DefaultConnectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DefaultConnectionAttributes) ClientRdpAccess() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("client_rdp_access"))
}

func (dc DefaultConnectionAttributes) ClientSshAccess() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("client_ssh_access"))
}

func (dc DefaultConnectionAttributes) WebRdpAccess() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("web_rdp_access"))
}

func (dc DefaultConnectionAttributes) WebSshAccess() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("web_ssh_access"))
}

type SupportAttributes struct {
	ref terra.Reference
}

func (s SupportAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SupportAttributes) InternalWithRef(ref terra.Reference) SupportAttributes {
	return SupportAttributes{ref: ref}
}

func (s SupportAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SupportAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("email"))
}

func (s SupportAttributes) Instructions() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("instructions"))
}

func (s SupportAttributes) Phone() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("phone"))
}

func (s SupportAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("url"))
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

type DefaultAutoShutdownState struct {
	DisconnectDelay string `json:"disconnect_delay"`
	IdleDelay       string `json:"idle_delay"`
	NoConnectDelay  string `json:"no_connect_delay"`
	ShutdownOnIdle  string `json:"shutdown_on_idle"`
}

type DefaultConnectionState struct {
	ClientRdpAccess string `json:"client_rdp_access"`
	ClientSshAccess string `json:"client_ssh_access"`
	WebRdpAccess    string `json:"web_rdp_access"`
	WebSshAccess    string `json:"web_ssh_access"`
}

type SupportState struct {
	Email        string `json:"email"`
	Instructions string `json:"instructions"`
	Phone        string `json:"phone"`
	Url          string `json:"url"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

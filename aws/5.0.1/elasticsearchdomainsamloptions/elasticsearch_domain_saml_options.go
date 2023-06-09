// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package elasticsearchdomainsamloptions

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SamlOptions struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// MasterBackendRole: string, optional
	MasterBackendRole terra.StringValue `hcl:"master_backend_role,attr"`
	// MasterUserName: string, optional
	MasterUserName terra.StringValue `hcl:"master_user_name,attr"`
	// RolesKey: string, optional
	RolesKey terra.StringValue `hcl:"roles_key,attr"`
	// SessionTimeoutMinutes: number, optional
	SessionTimeoutMinutes terra.NumberValue `hcl:"session_timeout_minutes,attr"`
	// SubjectKey: string, optional
	SubjectKey terra.StringValue `hcl:"subject_key,attr"`
	// Idp: optional
	Idp *Idp `hcl:"idp,block"`
}

type Idp struct {
	// EntityId: string, required
	EntityId terra.StringValue `hcl:"entity_id,attr" validate:"required"`
	// MetadataContent: string, required
	MetadataContent terra.StringValue `hcl:"metadata_content,attr" validate:"required"`
}

type Timeouts struct {
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type SamlOptionsAttributes struct {
	ref terra.Reference
}

func (so SamlOptionsAttributes) InternalRef() (terra.Reference, error) {
	return so.ref, nil
}

func (so SamlOptionsAttributes) InternalWithRef(ref terra.Reference) SamlOptionsAttributes {
	return SamlOptionsAttributes{ref: ref}
}

func (so SamlOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return so.ref.InternalTokens()
}

func (so SamlOptionsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(so.ref.Append("enabled"))
}

func (so SamlOptionsAttributes) MasterBackendRole() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("master_backend_role"))
}

func (so SamlOptionsAttributes) MasterUserName() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("master_user_name"))
}

func (so SamlOptionsAttributes) RolesKey() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("roles_key"))
}

func (so SamlOptionsAttributes) SessionTimeoutMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(so.ref.Append("session_timeout_minutes"))
}

func (so SamlOptionsAttributes) SubjectKey() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("subject_key"))
}

func (so SamlOptionsAttributes) Idp() terra.ListValue[IdpAttributes] {
	return terra.ReferenceAsList[IdpAttributes](so.ref.Append("idp"))
}

type IdpAttributes struct {
	ref terra.Reference
}

func (i IdpAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdpAttributes) InternalWithRef(ref terra.Reference) IdpAttributes {
	return IdpAttributes{ref: ref}
}

func (i IdpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdpAttributes) EntityId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("entity_id"))
}

func (i IdpAttributes) MetadataContent() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("metadata_content"))
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

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type SamlOptionsState struct {
	Enabled               bool       `json:"enabled"`
	MasterBackendRole     string     `json:"master_backend_role"`
	MasterUserName        string     `json:"master_user_name"`
	RolesKey              string     `json:"roles_key"`
	SessionTimeoutMinutes float64    `json:"session_timeout_minutes"`
	SubjectKey            string     `json:"subject_key"`
	Idp                   []IdpState `json:"idp"`
}

type IdpState struct {
	EntityId        string `json:"entity_id"`
	MetadataContent string `json:"metadata_content"`
}

type TimeoutsState struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
}

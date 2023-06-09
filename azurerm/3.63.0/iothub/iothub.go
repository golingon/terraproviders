// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package iothub

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Endpoint struct {
	// AuthenticationType: string, optional
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr"`
	// BatchFrequencyInSeconds: number, optional
	BatchFrequencyInSeconds terra.NumberValue `hcl:"batch_frequency_in_seconds,attr"`
	// ConnectionString: string, optional
	ConnectionString terra.StringValue `hcl:"connection_string,attr"`
	// ContainerName: string, optional
	ContainerName terra.StringValue `hcl:"container_name,attr"`
	// Encoding: string, optional
	Encoding terra.StringValue `hcl:"encoding,attr"`
	// EndpointUri: string, optional
	EndpointUri terra.StringValue `hcl:"endpoint_uri,attr"`
	// EntityPath: string, optional
	EntityPath terra.StringValue `hcl:"entity_path,attr"`
	// FileNameFormat: string, optional
	FileNameFormat terra.StringValue `hcl:"file_name_format,attr"`
	// IdentityId: string, optional
	IdentityId terra.StringValue `hcl:"identity_id,attr"`
	// MaxChunkSizeInBytes: number, optional
	MaxChunkSizeInBytes terra.NumberValue `hcl:"max_chunk_size_in_bytes,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type Enrichment struct {
	// EndpointNames: list of string, optional
	EndpointNames terra.ListValue[terra.StringValue] `hcl:"endpoint_names,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Route struct {
	// Condition: string, optional
	Condition terra.StringValue `hcl:"condition,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EndpointNames: list of string, optional
	EndpointNames terra.ListValue[terra.StringValue] `hcl:"endpoint_names,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Source: string, optional
	Source terra.StringValue `hcl:"source,attr"`
}

type SharedAccessPolicy struct{}

type CloudToDevice struct {
	// DefaultTtl: string, optional
	DefaultTtl terra.StringValue `hcl:"default_ttl,attr"`
	// MaxDeliveryCount: number, optional
	MaxDeliveryCount terra.NumberValue `hcl:"max_delivery_count,attr"`
	// Feedback: min=0
	Feedback []Feedback `hcl:"feedback,block" validate:"min=0"`
}

type Feedback struct {
	// LockDuration: string, optional
	LockDuration terra.StringValue `hcl:"lock_duration,attr"`
	// MaxDeliveryCount: number, optional
	MaxDeliveryCount terra.NumberValue `hcl:"max_delivery_count,attr"`
	// TimeToLive: string, optional
	TimeToLive terra.StringValue `hcl:"time_to_live,attr"`
}

type FallbackRoute struct {
	// Condition: string, optional
	Condition terra.StringValue `hcl:"condition,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EndpointNames: list of string, optional
	EndpointNames terra.ListValue[terra.StringValue] `hcl:"endpoint_names,attr"`
	// Source: string, optional
	Source terra.StringValue `hcl:"source,attr"`
}

type FileUpload struct {
	// AuthenticationType: string, optional
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr"`
	// ConnectionString: string, required
	ConnectionString terra.StringValue `hcl:"connection_string,attr" validate:"required"`
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// DefaultTtl: string, optional
	DefaultTtl terra.StringValue `hcl:"default_ttl,attr"`
	// IdentityId: string, optional
	IdentityId terra.StringValue `hcl:"identity_id,attr"`
	// LockDuration: string, optional
	LockDuration terra.StringValue `hcl:"lock_duration,attr"`
	// MaxDeliveryCount: number, optional
	MaxDeliveryCount terra.NumberValue `hcl:"max_delivery_count,attr"`
	// Notifications: bool, optional
	Notifications terra.BoolValue `hcl:"notifications,attr"`
	// SasTtl: string, optional
	SasTtl terra.StringValue `hcl:"sas_ttl,attr"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type NetworkRuleSet struct {
	// ApplyToBuiltinEventhubEndpoint: bool, optional
	ApplyToBuiltinEventhubEndpoint terra.BoolValue `hcl:"apply_to_builtin_eventhub_endpoint,attr"`
	// DefaultAction: string, optional
	DefaultAction terra.StringValue `hcl:"default_action,attr"`
	// IpRule: min=0
	IpRule []IpRule `hcl:"ip_rule,block" validate:"min=0"`
}

type IpRule struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// IpMask: string, required
	IpMask terra.StringValue `hcl:"ip_mask,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type Sku struct {
	// Capacity: number, required
	Capacity terra.NumberValue `hcl:"capacity,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
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

type EndpointAttributes struct {
	ref terra.Reference
}

func (e EndpointAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EndpointAttributes) InternalWithRef(ref terra.Reference) EndpointAttributes {
	return EndpointAttributes{ref: ref}
}

func (e EndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EndpointAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("authentication_type"))
}

func (e EndpointAttributes) BatchFrequencyInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("batch_frequency_in_seconds"))
}

func (e EndpointAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("connection_string"))
}

func (e EndpointAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("container_name"))
}

func (e EndpointAttributes) Encoding() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("encoding"))
}

func (e EndpointAttributes) EndpointUri() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("endpoint_uri"))
}

func (e EndpointAttributes) EntityPath() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("entity_path"))
}

func (e EndpointAttributes) FileNameFormat() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("file_name_format"))
}

func (e EndpointAttributes) IdentityId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("identity_id"))
}

func (e EndpointAttributes) MaxChunkSizeInBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("max_chunk_size_in_bytes"))
}

func (e EndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("name"))
}

func (e EndpointAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("resource_group_name"))
}

func (e EndpointAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("type"))
}

type EnrichmentAttributes struct {
	ref terra.Reference
}

func (e EnrichmentAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EnrichmentAttributes) InternalWithRef(ref terra.Reference) EnrichmentAttributes {
	return EnrichmentAttributes{ref: ref}
}

func (e EnrichmentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EnrichmentAttributes) EndpointNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("endpoint_names"))
}

func (e EnrichmentAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("key"))
}

func (e EnrichmentAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("value"))
}

type RouteAttributes struct {
	ref terra.Reference
}

func (r RouteAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RouteAttributes) InternalWithRef(ref terra.Reference) RouteAttributes {
	return RouteAttributes{ref: ref}
}

func (r RouteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RouteAttributes) Condition() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("condition"))
}

func (r RouteAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("enabled"))
}

func (r RouteAttributes) EndpointNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("endpoint_names"))
}

func (r RouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r RouteAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("source"))
}

type SharedAccessPolicyAttributes struct {
	ref terra.Reference
}

func (sap SharedAccessPolicyAttributes) InternalRef() (terra.Reference, error) {
	return sap.ref, nil
}

func (sap SharedAccessPolicyAttributes) InternalWithRef(ref terra.Reference) SharedAccessPolicyAttributes {
	return SharedAccessPolicyAttributes{ref: ref}
}

func (sap SharedAccessPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sap.ref.InternalTokens()
}

func (sap SharedAccessPolicyAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(sap.ref.Append("key_name"))
}

func (sap SharedAccessPolicyAttributes) Permissions() terra.StringValue {
	return terra.ReferenceAsString(sap.ref.Append("permissions"))
}

func (sap SharedAccessPolicyAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(sap.ref.Append("primary_key"))
}

func (sap SharedAccessPolicyAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(sap.ref.Append("secondary_key"))
}

type CloudToDeviceAttributes struct {
	ref terra.Reference
}

func (ctd CloudToDeviceAttributes) InternalRef() (terra.Reference, error) {
	return ctd.ref, nil
}

func (ctd CloudToDeviceAttributes) InternalWithRef(ref terra.Reference) CloudToDeviceAttributes {
	return CloudToDeviceAttributes{ref: ref}
}

func (ctd CloudToDeviceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ctd.ref.InternalTokens()
}

func (ctd CloudToDeviceAttributes) DefaultTtl() terra.StringValue {
	return terra.ReferenceAsString(ctd.ref.Append("default_ttl"))
}

func (ctd CloudToDeviceAttributes) MaxDeliveryCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ctd.ref.Append("max_delivery_count"))
}

func (ctd CloudToDeviceAttributes) Feedback() terra.ListValue[FeedbackAttributes] {
	return terra.ReferenceAsList[FeedbackAttributes](ctd.ref.Append("feedback"))
}

type FeedbackAttributes struct {
	ref terra.Reference
}

func (f FeedbackAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FeedbackAttributes) InternalWithRef(ref terra.Reference) FeedbackAttributes {
	return FeedbackAttributes{ref: ref}
}

func (f FeedbackAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FeedbackAttributes) LockDuration() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("lock_duration"))
}

func (f FeedbackAttributes) MaxDeliveryCount() terra.NumberValue {
	return terra.ReferenceAsNumber(f.ref.Append("max_delivery_count"))
}

func (f FeedbackAttributes) TimeToLive() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("time_to_live"))
}

type FallbackRouteAttributes struct {
	ref terra.Reference
}

func (fr FallbackRouteAttributes) InternalRef() (terra.Reference, error) {
	return fr.ref, nil
}

func (fr FallbackRouteAttributes) InternalWithRef(ref terra.Reference) FallbackRouteAttributes {
	return FallbackRouteAttributes{ref: ref}
}

func (fr FallbackRouteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fr.ref.InternalTokens()
}

func (fr FallbackRouteAttributes) Condition() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("condition"))
}

func (fr FallbackRouteAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(fr.ref.Append("enabled"))
}

func (fr FallbackRouteAttributes) EndpointNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fr.ref.Append("endpoint_names"))
}

func (fr FallbackRouteAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("source"))
}

type FileUploadAttributes struct {
	ref terra.Reference
}

func (fu FileUploadAttributes) InternalRef() (terra.Reference, error) {
	return fu.ref, nil
}

func (fu FileUploadAttributes) InternalWithRef(ref terra.Reference) FileUploadAttributes {
	return FileUploadAttributes{ref: ref}
}

func (fu FileUploadAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fu.ref.InternalTokens()
}

func (fu FileUploadAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(fu.ref.Append("authentication_type"))
}

func (fu FileUploadAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(fu.ref.Append("connection_string"))
}

func (fu FileUploadAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(fu.ref.Append("container_name"))
}

func (fu FileUploadAttributes) DefaultTtl() terra.StringValue {
	return terra.ReferenceAsString(fu.ref.Append("default_ttl"))
}

func (fu FileUploadAttributes) IdentityId() terra.StringValue {
	return terra.ReferenceAsString(fu.ref.Append("identity_id"))
}

func (fu FileUploadAttributes) LockDuration() terra.StringValue {
	return terra.ReferenceAsString(fu.ref.Append("lock_duration"))
}

func (fu FileUploadAttributes) MaxDeliveryCount() terra.NumberValue {
	return terra.ReferenceAsNumber(fu.ref.Append("max_delivery_count"))
}

func (fu FileUploadAttributes) Notifications() terra.BoolValue {
	return terra.ReferenceAsBool(fu.ref.Append("notifications"))
}

func (fu FileUploadAttributes) SasTtl() terra.StringValue {
	return terra.ReferenceAsString(fu.ref.Append("sas_ttl"))
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

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type NetworkRuleSetAttributes struct {
	ref terra.Reference
}

func (nrs NetworkRuleSetAttributes) InternalRef() (terra.Reference, error) {
	return nrs.ref, nil
}

func (nrs NetworkRuleSetAttributes) InternalWithRef(ref terra.Reference) NetworkRuleSetAttributes {
	return NetworkRuleSetAttributes{ref: ref}
}

func (nrs NetworkRuleSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nrs.ref.InternalTokens()
}

func (nrs NetworkRuleSetAttributes) ApplyToBuiltinEventhubEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(nrs.ref.Append("apply_to_builtin_eventhub_endpoint"))
}

func (nrs NetworkRuleSetAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceAsString(nrs.ref.Append("default_action"))
}

func (nrs NetworkRuleSetAttributes) IpRule() terra.ListValue[IpRuleAttributes] {
	return terra.ReferenceAsList[IpRuleAttributes](nrs.ref.Append("ip_rule"))
}

type IpRuleAttributes struct {
	ref terra.Reference
}

func (ir IpRuleAttributes) InternalRef() (terra.Reference, error) {
	return ir.ref, nil
}

func (ir IpRuleAttributes) InternalWithRef(ref terra.Reference) IpRuleAttributes {
	return IpRuleAttributes{ref: ref}
}

func (ir IpRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ir.ref.InternalTokens()
}

func (ir IpRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("action"))
}

func (ir IpRuleAttributes) IpMask() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("ip_mask"))
}

func (ir IpRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("name"))
}

type SkuAttributes struct {
	ref terra.Reference
}

func (s SkuAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SkuAttributes) InternalWithRef(ref terra.Reference) SkuAttributes {
	return SkuAttributes{ref: ref}
}

func (s SkuAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SkuAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("capacity"))
}

func (s SkuAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
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

type EndpointState struct {
	AuthenticationType      string  `json:"authentication_type"`
	BatchFrequencyInSeconds float64 `json:"batch_frequency_in_seconds"`
	ConnectionString        string  `json:"connection_string"`
	ContainerName           string  `json:"container_name"`
	Encoding                string  `json:"encoding"`
	EndpointUri             string  `json:"endpoint_uri"`
	EntityPath              string  `json:"entity_path"`
	FileNameFormat          string  `json:"file_name_format"`
	IdentityId              string  `json:"identity_id"`
	MaxChunkSizeInBytes     float64 `json:"max_chunk_size_in_bytes"`
	Name                    string  `json:"name"`
	ResourceGroupName       string  `json:"resource_group_name"`
	Type                    string  `json:"type"`
}

type EnrichmentState struct {
	EndpointNames []string `json:"endpoint_names"`
	Key           string   `json:"key"`
	Value         string   `json:"value"`
}

type RouteState struct {
	Condition     string   `json:"condition"`
	Enabled       bool     `json:"enabled"`
	EndpointNames []string `json:"endpoint_names"`
	Name          string   `json:"name"`
	Source        string   `json:"source"`
}

type SharedAccessPolicyState struct {
	KeyName      string `json:"key_name"`
	Permissions  string `json:"permissions"`
	PrimaryKey   string `json:"primary_key"`
	SecondaryKey string `json:"secondary_key"`
}

type CloudToDeviceState struct {
	DefaultTtl       string          `json:"default_ttl"`
	MaxDeliveryCount float64         `json:"max_delivery_count"`
	Feedback         []FeedbackState `json:"feedback"`
}

type FeedbackState struct {
	LockDuration     string  `json:"lock_duration"`
	MaxDeliveryCount float64 `json:"max_delivery_count"`
	TimeToLive       string  `json:"time_to_live"`
}

type FallbackRouteState struct {
	Condition     string   `json:"condition"`
	Enabled       bool     `json:"enabled"`
	EndpointNames []string `json:"endpoint_names"`
	Source        string   `json:"source"`
}

type FileUploadState struct {
	AuthenticationType string  `json:"authentication_type"`
	ConnectionString   string  `json:"connection_string"`
	ContainerName      string  `json:"container_name"`
	DefaultTtl         string  `json:"default_ttl"`
	IdentityId         string  `json:"identity_id"`
	LockDuration       string  `json:"lock_duration"`
	MaxDeliveryCount   float64 `json:"max_delivery_count"`
	Notifications      bool    `json:"notifications"`
	SasTtl             string  `json:"sas_ttl"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type NetworkRuleSetState struct {
	ApplyToBuiltinEventhubEndpoint bool          `json:"apply_to_builtin_eventhub_endpoint"`
	DefaultAction                  string        `json:"default_action"`
	IpRule                         []IpRuleState `json:"ip_rule"`
}

type IpRuleState struct {
	Action string `json:"action"`
	IpMask string `json:"ip_mask"`
	Name   string `json:"name"`
}

type SkuState struct {
	Capacity float64 `json:"capacity"`
	Name     string  `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

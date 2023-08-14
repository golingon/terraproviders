// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	webpubsub "github.com/golingon/terraproviders/azurerm/3.67.0/webpubsub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWebPubsub creates a new instance of [WebPubsub].
func NewWebPubsub(name string, args WebPubsubArgs) *WebPubsub {
	return &WebPubsub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WebPubsub)(nil)

// WebPubsub represents the Terraform resource azurerm_web_pubsub.
type WebPubsub struct {
	Name      string
	Args      WebPubsubArgs
	state     *webPubsubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WebPubsub].
func (wp *WebPubsub) Type() string {
	return "azurerm_web_pubsub"
}

// LocalName returns the local name for [WebPubsub].
func (wp *WebPubsub) LocalName() string {
	return wp.Name
}

// Configuration returns the configuration (args) for [WebPubsub].
func (wp *WebPubsub) Configuration() interface{} {
	return wp.Args
}

// DependOn is used for other resources to depend on [WebPubsub].
func (wp *WebPubsub) DependOn() terra.Reference {
	return terra.ReferenceResource(wp)
}

// Dependencies returns the list of resources [WebPubsub] depends_on.
func (wp *WebPubsub) Dependencies() terra.Dependencies {
	return wp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WebPubsub].
func (wp *WebPubsub) LifecycleManagement() *terra.Lifecycle {
	return wp.Lifecycle
}

// Attributes returns the attributes for [WebPubsub].
func (wp *WebPubsub) Attributes() webPubsubAttributes {
	return webPubsubAttributes{ref: terra.ReferenceResource(wp)}
}

// ImportState imports the given attribute values into [WebPubsub]'s state.
func (wp *WebPubsub) ImportState(av io.Reader) error {
	wp.state = &webPubsubState{}
	if err := json.NewDecoder(av).Decode(wp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wp.Type(), wp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WebPubsub] has state.
func (wp *WebPubsub) State() (*webPubsubState, bool) {
	return wp.state, wp.state != nil
}

// StateMust returns the state for [WebPubsub]. Panics if the state is nil.
func (wp *WebPubsub) StateMust() *webPubsubState {
	if wp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wp.Type(), wp.LocalName()))
	}
	return wp.state
}

// WebPubsubArgs contains the configurations for azurerm_web_pubsub.
type WebPubsubArgs struct {
	// AadAuthEnabled: bool, optional
	AadAuthEnabled terra.BoolValue `hcl:"aad_auth_enabled,attr"`
	// Capacity: number, optional
	Capacity terra.NumberValue `hcl:"capacity,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalAuthEnabled: bool, optional
	LocalAuthEnabled terra.BoolValue `hcl:"local_auth_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TlsClientCertEnabled: bool, optional
	TlsClientCertEnabled terra.BoolValue `hcl:"tls_client_cert_enabled,attr"`
	// Identity: optional
	Identity *webpubsub.Identity `hcl:"identity,block"`
	// LiveTrace: optional
	LiveTrace *webpubsub.LiveTrace `hcl:"live_trace,block"`
	// Timeouts: optional
	Timeouts *webpubsub.Timeouts `hcl:"timeouts,block"`
}
type webPubsubAttributes struct {
	ref terra.Reference
}

// AadAuthEnabled returns a reference to field aad_auth_enabled of azurerm_web_pubsub.
func (wp webPubsubAttributes) AadAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wp.ref.Append("aad_auth_enabled"))
}

// Capacity returns a reference to field capacity of azurerm_web_pubsub.
func (wp webPubsubAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(wp.ref.Append("capacity"))
}

// ExternalIp returns a reference to field external_ip of azurerm_web_pubsub.
func (wp webPubsubAttributes) ExternalIp() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("external_ip"))
}

// Hostname returns a reference to field hostname of azurerm_web_pubsub.
func (wp webPubsubAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_web_pubsub.
func (wp webPubsubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("id"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_web_pubsub.
func (wp webPubsubAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wp.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_web_pubsub.
func (wp webPubsubAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_web_pubsub.
func (wp webPubsubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_web_pubsub.
func (wp webPubsubAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("primary_access_key"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_web_pubsub.
func (wp webPubsubAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("primary_connection_string"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_web_pubsub.
func (wp webPubsubAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wp.ref.Append("public_network_access_enabled"))
}

// PublicPort returns a reference to field public_port of azurerm_web_pubsub.
func (wp webPubsubAttributes) PublicPort() terra.NumberValue {
	return terra.ReferenceAsNumber(wp.ref.Append("public_port"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_web_pubsub.
func (wp webPubsubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_web_pubsub.
func (wp webPubsubAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("secondary_access_key"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_web_pubsub.
func (wp webPubsubAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("secondary_connection_string"))
}

// ServerPort returns a reference to field server_port of azurerm_web_pubsub.
func (wp webPubsubAttributes) ServerPort() terra.NumberValue {
	return terra.ReferenceAsNumber(wp.ref.Append("server_port"))
}

// Sku returns a reference to field sku of azurerm_web_pubsub.
func (wp webPubsubAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_web_pubsub.
func (wp webPubsubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wp.ref.Append("tags"))
}

// TlsClientCertEnabled returns a reference to field tls_client_cert_enabled of azurerm_web_pubsub.
func (wp webPubsubAttributes) TlsClientCertEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wp.ref.Append("tls_client_cert_enabled"))
}

// Version returns a reference to field version of azurerm_web_pubsub.
func (wp webPubsubAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("version"))
}

func (wp webPubsubAttributes) Identity() terra.ListValue[webpubsub.IdentityAttributes] {
	return terra.ReferenceAsList[webpubsub.IdentityAttributes](wp.ref.Append("identity"))
}

func (wp webPubsubAttributes) LiveTrace() terra.ListValue[webpubsub.LiveTraceAttributes] {
	return terra.ReferenceAsList[webpubsub.LiveTraceAttributes](wp.ref.Append("live_trace"))
}

func (wp webPubsubAttributes) Timeouts() webpubsub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[webpubsub.TimeoutsAttributes](wp.ref.Append("timeouts"))
}

type webPubsubState struct {
	AadAuthEnabled             bool                       `json:"aad_auth_enabled"`
	Capacity                   float64                    `json:"capacity"`
	ExternalIp                 string                     `json:"external_ip"`
	Hostname                   string                     `json:"hostname"`
	Id                         string                     `json:"id"`
	LocalAuthEnabled           bool                       `json:"local_auth_enabled"`
	Location                   string                     `json:"location"`
	Name                       string                     `json:"name"`
	PrimaryAccessKey           string                     `json:"primary_access_key"`
	PrimaryConnectionString    string                     `json:"primary_connection_string"`
	PublicNetworkAccessEnabled bool                       `json:"public_network_access_enabled"`
	PublicPort                 float64                    `json:"public_port"`
	ResourceGroupName          string                     `json:"resource_group_name"`
	SecondaryAccessKey         string                     `json:"secondary_access_key"`
	SecondaryConnectionString  string                     `json:"secondary_connection_string"`
	ServerPort                 float64                    `json:"server_port"`
	Sku                        string                     `json:"sku"`
	Tags                       map[string]string          `json:"tags"`
	TlsClientCertEnabled       bool                       `json:"tls_client_cert_enabled"`
	Version                    string                     `json:"version"`
	Identity                   []webpubsub.IdentityState  `json:"identity"`
	LiveTrace                  []webpubsub.LiveTraceState `json:"live_trace"`
	Timeouts                   *webpubsub.TimeoutsState   `json:"timeouts"`
}
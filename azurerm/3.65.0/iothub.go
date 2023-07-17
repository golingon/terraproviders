// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothub "github.com/golingon/terraproviders/azurerm/3.65.0/iothub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothub creates a new instance of [Iothub].
func NewIothub(name string, args IothubArgs) *Iothub {
	return &Iothub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Iothub)(nil)

// Iothub represents the Terraform resource azurerm_iothub.
type Iothub struct {
	Name      string
	Args      IothubArgs
	state     *iothubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Iothub].
func (i *Iothub) Type() string {
	return "azurerm_iothub"
}

// LocalName returns the local name for [Iothub].
func (i *Iothub) LocalName() string {
	return i.Name
}

// Configuration returns the configuration (args) for [Iothub].
func (i *Iothub) Configuration() interface{} {
	return i.Args
}

// DependOn is used for other resources to depend on [Iothub].
func (i *Iothub) DependOn() terra.Reference {
	return terra.ReferenceResource(i)
}

// Dependencies returns the list of resources [Iothub] depends_on.
func (i *Iothub) Dependencies() terra.Dependencies {
	return i.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Iothub].
func (i *Iothub) LifecycleManagement() *terra.Lifecycle {
	return i.Lifecycle
}

// Attributes returns the attributes for [Iothub].
func (i *Iothub) Attributes() iothubAttributes {
	return iothubAttributes{ref: terra.ReferenceResource(i)}
}

// ImportState imports the given attribute values into [Iothub]'s state.
func (i *Iothub) ImportState(av io.Reader) error {
	i.state = &iothubState{}
	if err := json.NewDecoder(av).Decode(i.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", i.Type(), i.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Iothub] has state.
func (i *Iothub) State() (*iothubState, bool) {
	return i.state, i.state != nil
}

// StateMust returns the state for [Iothub]. Panics if the state is nil.
func (i *Iothub) StateMust() *iothubState {
	if i.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", i.Type(), i.LocalName()))
	}
	return i.state
}

// IothubArgs contains the configurations for azurerm_iothub.
type IothubArgs struct {
	// EventHubPartitionCount: number, optional
	EventHubPartitionCount terra.NumberValue `hcl:"event_hub_partition_count,attr"`
	// EventHubRetentionInDays: number, optional
	EventHubRetentionInDays terra.NumberValue `hcl:"event_hub_retention_in_days,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MinTlsVersion: string, optional
	MinTlsVersion terra.StringValue `hcl:"min_tls_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Endpoint: min=0
	Endpoint []iothub.Endpoint `hcl:"endpoint,block" validate:"min=0"`
	// Enrichment: min=0
	Enrichment []iothub.Enrichment `hcl:"enrichment,block" validate:"min=0"`
	// Route: min=0
	Route []iothub.Route `hcl:"route,block" validate:"min=0"`
	// SharedAccessPolicy: min=0
	SharedAccessPolicy []iothub.SharedAccessPolicy `hcl:"shared_access_policy,block" validate:"min=0"`
	// CloudToDevice: optional
	CloudToDevice *iothub.CloudToDevice `hcl:"cloud_to_device,block"`
	// FallbackRoute: optional
	FallbackRoute *iothub.FallbackRoute `hcl:"fallback_route,block"`
	// FileUpload: optional
	FileUpload *iothub.FileUpload `hcl:"file_upload,block"`
	// Identity: optional
	Identity *iothub.Identity `hcl:"identity,block"`
	// NetworkRuleSet: min=0
	NetworkRuleSet []iothub.NetworkRuleSet `hcl:"network_rule_set,block" validate:"min=0"`
	// Sku: required
	Sku *iothub.Sku `hcl:"sku,block" validate:"required"`
	// Timeouts: optional
	Timeouts *iothub.Timeouts `hcl:"timeouts,block"`
}
type iothubAttributes struct {
	ref terra.Reference
}

// EventHubEventsEndpoint returns a reference to field event_hub_events_endpoint of azurerm_iothub.
func (i iothubAttributes) EventHubEventsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("event_hub_events_endpoint"))
}

// EventHubEventsNamespace returns a reference to field event_hub_events_namespace of azurerm_iothub.
func (i iothubAttributes) EventHubEventsNamespace() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("event_hub_events_namespace"))
}

// EventHubEventsPath returns a reference to field event_hub_events_path of azurerm_iothub.
func (i iothubAttributes) EventHubEventsPath() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("event_hub_events_path"))
}

// EventHubOperationsEndpoint returns a reference to field event_hub_operations_endpoint of azurerm_iothub.
func (i iothubAttributes) EventHubOperationsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("event_hub_operations_endpoint"))
}

// EventHubOperationsPath returns a reference to field event_hub_operations_path of azurerm_iothub.
func (i iothubAttributes) EventHubOperationsPath() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("event_hub_operations_path"))
}

// EventHubPartitionCount returns a reference to field event_hub_partition_count of azurerm_iothub.
func (i iothubAttributes) EventHubPartitionCount() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("event_hub_partition_count"))
}

// EventHubRetentionInDays returns a reference to field event_hub_retention_in_days of azurerm_iothub.
func (i iothubAttributes) EventHubRetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("event_hub_retention_in_days"))
}

// Hostname returns a reference to field hostname of azurerm_iothub.
func (i iothubAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_iothub.
func (i iothubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_iothub.
func (i iothubAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("location"))
}

// MinTlsVersion returns a reference to field min_tls_version of azurerm_iothub.
func (i iothubAttributes) MinTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("min_tls_version"))
}

// Name returns a reference to field name of azurerm_iothub.
func (i iothubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_iothub.
func (i iothubAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub.
func (i iothubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_iothub.
func (i iothubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("tags"))
}

// Type returns a reference to field type of azurerm_iothub.
func (i iothubAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

func (i iothubAttributes) Endpoint() terra.ListValue[iothub.EndpointAttributes] {
	return terra.ReferenceAsList[iothub.EndpointAttributes](i.ref.Append("endpoint"))
}

func (i iothubAttributes) Enrichment() terra.ListValue[iothub.EnrichmentAttributes] {
	return terra.ReferenceAsList[iothub.EnrichmentAttributes](i.ref.Append("enrichment"))
}

func (i iothubAttributes) Route() terra.ListValue[iothub.RouteAttributes] {
	return terra.ReferenceAsList[iothub.RouteAttributes](i.ref.Append("route"))
}

func (i iothubAttributes) SharedAccessPolicy() terra.ListValue[iothub.SharedAccessPolicyAttributes] {
	return terra.ReferenceAsList[iothub.SharedAccessPolicyAttributes](i.ref.Append("shared_access_policy"))
}

func (i iothubAttributes) CloudToDevice() terra.ListValue[iothub.CloudToDeviceAttributes] {
	return terra.ReferenceAsList[iothub.CloudToDeviceAttributes](i.ref.Append("cloud_to_device"))
}

func (i iothubAttributes) FallbackRoute() terra.ListValue[iothub.FallbackRouteAttributes] {
	return terra.ReferenceAsList[iothub.FallbackRouteAttributes](i.ref.Append("fallback_route"))
}

func (i iothubAttributes) FileUpload() terra.ListValue[iothub.FileUploadAttributes] {
	return terra.ReferenceAsList[iothub.FileUploadAttributes](i.ref.Append("file_upload"))
}

func (i iothubAttributes) Identity() terra.ListValue[iothub.IdentityAttributes] {
	return terra.ReferenceAsList[iothub.IdentityAttributes](i.ref.Append("identity"))
}

func (i iothubAttributes) NetworkRuleSet() terra.ListValue[iothub.NetworkRuleSetAttributes] {
	return terra.ReferenceAsList[iothub.NetworkRuleSetAttributes](i.ref.Append("network_rule_set"))
}

func (i iothubAttributes) Sku() terra.ListValue[iothub.SkuAttributes] {
	return terra.ReferenceAsList[iothub.SkuAttributes](i.ref.Append("sku"))
}

func (i iothubAttributes) Timeouts() iothub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothub.TimeoutsAttributes](i.ref.Append("timeouts"))
}

type iothubState struct {
	EventHubEventsEndpoint     string                           `json:"event_hub_events_endpoint"`
	EventHubEventsNamespace    string                           `json:"event_hub_events_namespace"`
	EventHubEventsPath         string                           `json:"event_hub_events_path"`
	EventHubOperationsEndpoint string                           `json:"event_hub_operations_endpoint"`
	EventHubOperationsPath     string                           `json:"event_hub_operations_path"`
	EventHubPartitionCount     float64                          `json:"event_hub_partition_count"`
	EventHubRetentionInDays    float64                          `json:"event_hub_retention_in_days"`
	Hostname                   string                           `json:"hostname"`
	Id                         string                           `json:"id"`
	Location                   string                           `json:"location"`
	MinTlsVersion              string                           `json:"min_tls_version"`
	Name                       string                           `json:"name"`
	PublicNetworkAccessEnabled bool                             `json:"public_network_access_enabled"`
	ResourceGroupName          string                           `json:"resource_group_name"`
	Tags                       map[string]string                `json:"tags"`
	Type                       string                           `json:"type"`
	Endpoint                   []iothub.EndpointState           `json:"endpoint"`
	Enrichment                 []iothub.EnrichmentState         `json:"enrichment"`
	Route                      []iothub.RouteState              `json:"route"`
	SharedAccessPolicy         []iothub.SharedAccessPolicyState `json:"shared_access_policy"`
	CloudToDevice              []iothub.CloudToDeviceState      `json:"cloud_to_device"`
	FallbackRoute              []iothub.FallbackRouteState      `json:"fallback_route"`
	FileUpload                 []iothub.FileUploadState         `json:"file_upload"`
	Identity                   []iothub.IdentityState           `json:"identity"`
	NetworkRuleSet             []iothub.NetworkRuleSetState     `json:"network_rule_set"`
	Sku                        []iothub.SkuState                `json:"sku"`
	Timeouts                   *iothub.TimeoutsState            `json:"timeouts"`
}

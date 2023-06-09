// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kustocluster "github.com/golingon/terraproviders/azurerm/3.63.0/kustocluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKustoCluster creates a new instance of [KustoCluster].
func NewKustoCluster(name string, args KustoClusterArgs) *KustoCluster {
	return &KustoCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoCluster)(nil)

// KustoCluster represents the Terraform resource azurerm_kusto_cluster.
type KustoCluster struct {
	Name      string
	Args      KustoClusterArgs
	state     *kustoClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoCluster].
func (kc *KustoCluster) Type() string {
	return "azurerm_kusto_cluster"
}

// LocalName returns the local name for [KustoCluster].
func (kc *KustoCluster) LocalName() string {
	return kc.Name
}

// Configuration returns the configuration (args) for [KustoCluster].
func (kc *KustoCluster) Configuration() interface{} {
	return kc.Args
}

// DependOn is used for other resources to depend on [KustoCluster].
func (kc *KustoCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(kc)
}

// Dependencies returns the list of resources [KustoCluster] depends_on.
func (kc *KustoCluster) Dependencies() terra.Dependencies {
	return kc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoCluster].
func (kc *KustoCluster) LifecycleManagement() *terra.Lifecycle {
	return kc.Lifecycle
}

// Attributes returns the attributes for [KustoCluster].
func (kc *KustoCluster) Attributes() kustoClusterAttributes {
	return kustoClusterAttributes{ref: terra.ReferenceResource(kc)}
}

// ImportState imports the given attribute values into [KustoCluster]'s state.
func (kc *KustoCluster) ImportState(av io.Reader) error {
	kc.state = &kustoClusterState{}
	if err := json.NewDecoder(av).Decode(kc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kc.Type(), kc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoCluster] has state.
func (kc *KustoCluster) State() (*kustoClusterState, bool) {
	return kc.state, kc.state != nil
}

// StateMust returns the state for [KustoCluster]. Panics if the state is nil.
func (kc *KustoCluster) StateMust() *kustoClusterState {
	if kc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kc.Type(), kc.LocalName()))
	}
	return kc.state
}

// KustoClusterArgs contains the configurations for azurerm_kusto_cluster.
type KustoClusterArgs struct {
	// AllowedFqdns: list of string, optional
	AllowedFqdns terra.ListValue[terra.StringValue] `hcl:"allowed_fqdns,attr"`
	// AllowedIpRanges: list of string, optional
	AllowedIpRanges terra.ListValue[terra.StringValue] `hcl:"allowed_ip_ranges,attr"`
	// AutoStopEnabled: bool, optional
	AutoStopEnabled terra.BoolValue `hcl:"auto_stop_enabled,attr"`
	// DiskEncryptionEnabled: bool, optional
	DiskEncryptionEnabled terra.BoolValue `hcl:"disk_encryption_enabled,attr"`
	// DoubleEncryptionEnabled: bool, optional
	DoubleEncryptionEnabled terra.BoolValue `hcl:"double_encryption_enabled,attr"`
	// Engine: string, optional
	Engine terra.StringValue `hcl:"engine,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LanguageExtensions: set of string, optional
	LanguageExtensions terra.SetValue[terra.StringValue] `hcl:"language_extensions,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OutboundNetworkAccessRestricted: bool, optional
	OutboundNetworkAccessRestricted terra.BoolValue `hcl:"outbound_network_access_restricted,attr"`
	// PublicIpType: string, optional
	PublicIpType terra.StringValue `hcl:"public_ip_type,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// PurgeEnabled: bool, optional
	PurgeEnabled terra.BoolValue `hcl:"purge_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StreamingIngestionEnabled: bool, optional
	StreamingIngestionEnabled terra.BoolValue `hcl:"streaming_ingestion_enabled,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TrustedExternalTenants: list of string, optional
	TrustedExternalTenants terra.ListValue[terra.StringValue] `hcl:"trusted_external_tenants,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// Identity: optional
	Identity *kustocluster.Identity `hcl:"identity,block"`
	// OptimizedAutoScale: optional
	OptimizedAutoScale *kustocluster.OptimizedAutoScale `hcl:"optimized_auto_scale,block"`
	// Sku: required
	Sku *kustocluster.Sku `hcl:"sku,block" validate:"required"`
	// Timeouts: optional
	Timeouts *kustocluster.Timeouts `hcl:"timeouts,block"`
	// VirtualNetworkConfiguration: optional
	VirtualNetworkConfiguration *kustocluster.VirtualNetworkConfiguration `hcl:"virtual_network_configuration,block"`
}
type kustoClusterAttributes struct {
	ref terra.Reference
}

// AllowedFqdns returns a reference to field allowed_fqdns of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) AllowedFqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kc.ref.Append("allowed_fqdns"))
}

// AllowedIpRanges returns a reference to field allowed_ip_ranges of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) AllowedIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kc.ref.Append("allowed_ip_ranges"))
}

// AutoStopEnabled returns a reference to field auto_stop_enabled of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) AutoStopEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("auto_stop_enabled"))
}

// DataIngestionUri returns a reference to field data_ingestion_uri of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) DataIngestionUri() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("data_ingestion_uri"))
}

// DiskEncryptionEnabled returns a reference to field disk_encryption_enabled of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) DiskEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("disk_encryption_enabled"))
}

// DoubleEncryptionEnabled returns a reference to field double_encryption_enabled of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) DoubleEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("double_encryption_enabled"))
}

// Engine returns a reference to field engine of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("engine"))
}

// Id returns a reference to field id of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("id"))
}

// LanguageExtensions returns a reference to field language_extensions of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) LanguageExtensions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kc.ref.Append("language_extensions"))
}

// Location returns a reference to field location of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("name"))
}

// OutboundNetworkAccessRestricted returns a reference to field outbound_network_access_restricted of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) OutboundNetworkAccessRestricted() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("outbound_network_access_restricted"))
}

// PublicIpType returns a reference to field public_ip_type of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) PublicIpType() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("public_ip_type"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("public_network_access_enabled"))
}

// PurgeEnabled returns a reference to field purge_enabled of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) PurgeEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("purge_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("resource_group_name"))
}

// StreamingIngestionEnabled returns a reference to field streaming_ingestion_enabled of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) StreamingIngestionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("streaming_ingestion_enabled"))
}

// Tags returns a reference to field tags of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kc.ref.Append("tags"))
}

// TrustedExternalTenants returns a reference to field trusted_external_tenants of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) TrustedExternalTenants() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kc.ref.Append("trusted_external_tenants"))
}

// Uri returns a reference to field uri of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("uri"))
}

// Zones returns a reference to field zones of azurerm_kusto_cluster.
func (kc kustoClusterAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kc.ref.Append("zones"))
}

func (kc kustoClusterAttributes) Identity() terra.ListValue[kustocluster.IdentityAttributes] {
	return terra.ReferenceAsList[kustocluster.IdentityAttributes](kc.ref.Append("identity"))
}

func (kc kustoClusterAttributes) OptimizedAutoScale() terra.ListValue[kustocluster.OptimizedAutoScaleAttributes] {
	return terra.ReferenceAsList[kustocluster.OptimizedAutoScaleAttributes](kc.ref.Append("optimized_auto_scale"))
}

func (kc kustoClusterAttributes) Sku() terra.ListValue[kustocluster.SkuAttributes] {
	return terra.ReferenceAsList[kustocluster.SkuAttributes](kc.ref.Append("sku"))
}

func (kc kustoClusterAttributes) Timeouts() kustocluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustocluster.TimeoutsAttributes](kc.ref.Append("timeouts"))
}

func (kc kustoClusterAttributes) VirtualNetworkConfiguration() terra.ListValue[kustocluster.VirtualNetworkConfigurationAttributes] {
	return terra.ReferenceAsList[kustocluster.VirtualNetworkConfigurationAttributes](kc.ref.Append("virtual_network_configuration"))
}

type kustoClusterState struct {
	AllowedFqdns                    []string                                        `json:"allowed_fqdns"`
	AllowedIpRanges                 []string                                        `json:"allowed_ip_ranges"`
	AutoStopEnabled                 bool                                            `json:"auto_stop_enabled"`
	DataIngestionUri                string                                          `json:"data_ingestion_uri"`
	DiskEncryptionEnabled           bool                                            `json:"disk_encryption_enabled"`
	DoubleEncryptionEnabled         bool                                            `json:"double_encryption_enabled"`
	Engine                          string                                          `json:"engine"`
	Id                              string                                          `json:"id"`
	LanguageExtensions              []string                                        `json:"language_extensions"`
	Location                        string                                          `json:"location"`
	Name                            string                                          `json:"name"`
	OutboundNetworkAccessRestricted bool                                            `json:"outbound_network_access_restricted"`
	PublicIpType                    string                                          `json:"public_ip_type"`
	PublicNetworkAccessEnabled      bool                                            `json:"public_network_access_enabled"`
	PurgeEnabled                    bool                                            `json:"purge_enabled"`
	ResourceGroupName               string                                          `json:"resource_group_name"`
	StreamingIngestionEnabled       bool                                            `json:"streaming_ingestion_enabled"`
	Tags                            map[string]string                               `json:"tags"`
	TrustedExternalTenants          []string                                        `json:"trusted_external_tenants"`
	Uri                             string                                          `json:"uri"`
	Zones                           []string                                        `json:"zones"`
	Identity                        []kustocluster.IdentityState                    `json:"identity"`
	OptimizedAutoScale              []kustocluster.OptimizedAutoScaleState          `json:"optimized_auto_scale"`
	Sku                             []kustocluster.SkuState                         `json:"sku"`
	Timeouts                        *kustocluster.TimeoutsState                     `json:"timeouts"`
	VirtualNetworkConfiguration     []kustocluster.VirtualNetworkConfigurationState `json:"virtual_network_configuration"`
}

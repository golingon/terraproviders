// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datahybridcomputemachine "github.com/golingon/terraproviders/azurerm/3.67.0/datahybridcomputemachine"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataHybridComputeMachine creates a new instance of [DataHybridComputeMachine].
func NewDataHybridComputeMachine(name string, args DataHybridComputeMachineArgs) *DataHybridComputeMachine {
	return &DataHybridComputeMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHybridComputeMachine)(nil)

// DataHybridComputeMachine represents the Terraform data resource azurerm_hybrid_compute_machine.
type DataHybridComputeMachine struct {
	Name string
	Args DataHybridComputeMachineArgs
}

// DataSource returns the Terraform object type for [DataHybridComputeMachine].
func (hcm *DataHybridComputeMachine) DataSource() string {
	return "azurerm_hybrid_compute_machine"
}

// LocalName returns the local name for [DataHybridComputeMachine].
func (hcm *DataHybridComputeMachine) LocalName() string {
	return hcm.Name
}

// Configuration returns the configuration (args) for [DataHybridComputeMachine].
func (hcm *DataHybridComputeMachine) Configuration() interface{} {
	return hcm.Args
}

// Attributes returns the attributes for [DataHybridComputeMachine].
func (hcm *DataHybridComputeMachine) Attributes() dataHybridComputeMachineAttributes {
	return dataHybridComputeMachineAttributes{ref: terra.ReferenceDataResource(hcm)}
}

// DataHybridComputeMachineArgs contains the configurations for azurerm_hybrid_compute_machine.
type DataHybridComputeMachineArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AgentConfiguration: min=0
	AgentConfiguration []datahybridcomputemachine.AgentConfiguration `hcl:"agent_configuration,block" validate:"min=0"`
	// CloudMetadata: min=0
	CloudMetadata []datahybridcomputemachine.CloudMetadata `hcl:"cloud_metadata,block" validate:"min=0"`
	// ErrorDetails: min=0
	ErrorDetails []datahybridcomputemachine.ErrorDetails `hcl:"error_details,block" validate:"min=0"`
	// Identity: min=0
	Identity []datahybridcomputemachine.Identity `hcl:"identity,block" validate:"min=0"`
	// LocationData: min=0
	LocationData []datahybridcomputemachine.LocationData `hcl:"location_data,block" validate:"min=0"`
	// OsProfile: min=0
	OsProfile []datahybridcomputemachine.OsProfile `hcl:"os_profile,block" validate:"min=0"`
	// ServiceStatus: min=0
	ServiceStatus []datahybridcomputemachine.ServiceStatus `hcl:"service_status,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datahybridcomputemachine.Timeouts `hcl:"timeouts,block"`
}
type dataHybridComputeMachineAttributes struct {
	ref terra.Reference
}

// AdFqdn returns a reference to field ad_fqdn of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) AdFqdn() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("ad_fqdn"))
}

// AgentVersion returns a reference to field agent_version of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) AgentVersion() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("agent_version"))
}

// ClientPublicKey returns a reference to field client_public_key of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) ClientPublicKey() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("client_public_key"))
}

// DetectedProperties returns a reference to field detected_properties of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) DetectedProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hcm.ref.Append("detected_properties"))
}

// DisplayName returns a reference to field display_name of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("display_name"))
}

// DnsFqdn returns a reference to field dns_fqdn of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) DnsFqdn() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("dns_fqdn"))
}

// DomainName returns a reference to field domain_name of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("domain_name"))
}

// Id returns a reference to field id of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("id"))
}

// LastStatusChange returns a reference to field last_status_change of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) LastStatusChange() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("last_status_change"))
}

// Location returns a reference to field location of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("location"))
}

// MachineFqdn returns a reference to field machine_fqdn of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) MachineFqdn() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("machine_fqdn"))
}

// MssqlDiscovered returns a reference to field mssql_discovered of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) MssqlDiscovered() terra.BoolValue {
	return terra.ReferenceAsBool(hcm.ref.Append("mssql_discovered"))
}

// Name returns a reference to field name of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("name"))
}

// OsName returns a reference to field os_name of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) OsName() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("os_name"))
}

// OsSku returns a reference to field os_sku of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) OsSku() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("os_sku"))
}

// OsType returns a reference to field os_type of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("os_type"))
}

// OsVersion returns a reference to field os_version of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) OsVersion() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("os_version"))
}

// ParentClusterResourceId returns a reference to field parent_cluster_resource_id of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) ParentClusterResourceId() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("parent_cluster_resource_id"))
}

// PrivateLinkScopeResourceId returns a reference to field private_link_scope_resource_id of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) PrivateLinkScopeResourceId() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("private_link_scope_resource_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("resource_group_name"))
}

// Status returns a reference to field status of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("status"))
}

// Tags returns a reference to field tags of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hcm.ref.Append("tags"))
}

// VmId returns a reference to field vm_id of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) VmId() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("vm_id"))
}

// VmUuid returns a reference to field vm_uuid of azurerm_hybrid_compute_machine.
func (hcm dataHybridComputeMachineAttributes) VmUuid() terra.StringValue {
	return terra.ReferenceAsString(hcm.ref.Append("vm_uuid"))
}

func (hcm dataHybridComputeMachineAttributes) AgentConfiguration() terra.ListValue[datahybridcomputemachine.AgentConfigurationAttributes] {
	return terra.ReferenceAsList[datahybridcomputemachine.AgentConfigurationAttributes](hcm.ref.Append("agent_configuration"))
}

func (hcm dataHybridComputeMachineAttributes) CloudMetadata() terra.ListValue[datahybridcomputemachine.CloudMetadataAttributes] {
	return terra.ReferenceAsList[datahybridcomputemachine.CloudMetadataAttributes](hcm.ref.Append("cloud_metadata"))
}

func (hcm dataHybridComputeMachineAttributes) ErrorDetails() terra.ListValue[datahybridcomputemachine.ErrorDetailsAttributes] {
	return terra.ReferenceAsList[datahybridcomputemachine.ErrorDetailsAttributes](hcm.ref.Append("error_details"))
}

func (hcm dataHybridComputeMachineAttributes) Identity() terra.ListValue[datahybridcomputemachine.IdentityAttributes] {
	return terra.ReferenceAsList[datahybridcomputemachine.IdentityAttributes](hcm.ref.Append("identity"))
}

func (hcm dataHybridComputeMachineAttributes) LocationData() terra.ListValue[datahybridcomputemachine.LocationDataAttributes] {
	return terra.ReferenceAsList[datahybridcomputemachine.LocationDataAttributes](hcm.ref.Append("location_data"))
}

func (hcm dataHybridComputeMachineAttributes) OsProfile() terra.ListValue[datahybridcomputemachine.OsProfileAttributes] {
	return terra.ReferenceAsList[datahybridcomputemachine.OsProfileAttributes](hcm.ref.Append("os_profile"))
}

func (hcm dataHybridComputeMachineAttributes) ServiceStatus() terra.ListValue[datahybridcomputemachine.ServiceStatusAttributes] {
	return terra.ReferenceAsList[datahybridcomputemachine.ServiceStatusAttributes](hcm.ref.Append("service_status"))
}

func (hcm dataHybridComputeMachineAttributes) Timeouts() datahybridcomputemachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datahybridcomputemachine.TimeoutsAttributes](hcm.ref.Append("timeouts"))
}
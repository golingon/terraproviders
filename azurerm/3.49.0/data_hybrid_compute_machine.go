// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datahybridcomputemachine "github.com/golingon/terraproviders/azurerm/3.49.0/datahybridcomputemachine"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataHybridComputeMachine(name string, args DataHybridComputeMachineArgs) *DataHybridComputeMachine {
	return &DataHybridComputeMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHybridComputeMachine)(nil)

type DataHybridComputeMachine struct {
	Name string
	Args DataHybridComputeMachineArgs
}

func (hcm *DataHybridComputeMachine) DataSource() string {
	return "azurerm_hybrid_compute_machine"
}

func (hcm *DataHybridComputeMachine) LocalName() string {
	return hcm.Name
}

func (hcm *DataHybridComputeMachine) Configuration() interface{} {
	return hcm.Args
}

func (hcm *DataHybridComputeMachine) Attributes() dataHybridComputeMachineAttributes {
	return dataHybridComputeMachineAttributes{ref: terra.ReferenceDataResource(hcm)}
}

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

func (hcm dataHybridComputeMachineAttributes) AdFqdn() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("ad_fqdn"))
}

func (hcm dataHybridComputeMachineAttributes) AgentVersion() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("agent_version"))
}

func (hcm dataHybridComputeMachineAttributes) ClientPublicKey() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("client_public_key"))
}

func (hcm dataHybridComputeMachineAttributes) DetectedProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](hcm.ref.Append("detected_properties"))
}

func (hcm dataHybridComputeMachineAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("display_name"))
}

func (hcm dataHybridComputeMachineAttributes) DnsFqdn() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("dns_fqdn"))
}

func (hcm dataHybridComputeMachineAttributes) DomainName() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("domain_name"))
}

func (hcm dataHybridComputeMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("id"))
}

func (hcm dataHybridComputeMachineAttributes) LastStatusChange() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("last_status_change"))
}

func (hcm dataHybridComputeMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("location"))
}

func (hcm dataHybridComputeMachineAttributes) MachineFqdn() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("machine_fqdn"))
}

func (hcm dataHybridComputeMachineAttributes) MssqlDiscovered() terra.BoolValue {
	return terra.ReferenceBool(hcm.ref.Append("mssql_discovered"))
}

func (hcm dataHybridComputeMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("name"))
}

func (hcm dataHybridComputeMachineAttributes) OsName() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("os_name"))
}

func (hcm dataHybridComputeMachineAttributes) OsSku() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("os_sku"))
}

func (hcm dataHybridComputeMachineAttributes) OsType() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("os_type"))
}

func (hcm dataHybridComputeMachineAttributes) OsVersion() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("os_version"))
}

func (hcm dataHybridComputeMachineAttributes) ParentClusterResourceId() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("parent_cluster_resource_id"))
}

func (hcm dataHybridComputeMachineAttributes) PrivateLinkScopeResourceId() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("private_link_scope_resource_id"))
}

func (hcm dataHybridComputeMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("resource_group_name"))
}

func (hcm dataHybridComputeMachineAttributes) Status() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("status"))
}

func (hcm dataHybridComputeMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](hcm.ref.Append("tags"))
}

func (hcm dataHybridComputeMachineAttributes) VmId() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("vm_id"))
}

func (hcm dataHybridComputeMachineAttributes) VmUuid() terra.StringValue {
	return terra.ReferenceString(hcm.ref.Append("vm_uuid"))
}

func (hcm dataHybridComputeMachineAttributes) AgentConfiguration() terra.ListValue[datahybridcomputemachine.AgentConfigurationAttributes] {
	return terra.ReferenceList[datahybridcomputemachine.AgentConfigurationAttributes](hcm.ref.Append("agent_configuration"))
}

func (hcm dataHybridComputeMachineAttributes) CloudMetadata() terra.ListValue[datahybridcomputemachine.CloudMetadataAttributes] {
	return terra.ReferenceList[datahybridcomputemachine.CloudMetadataAttributes](hcm.ref.Append("cloud_metadata"))
}

func (hcm dataHybridComputeMachineAttributes) ErrorDetails() terra.ListValue[datahybridcomputemachine.ErrorDetailsAttributes] {
	return terra.ReferenceList[datahybridcomputemachine.ErrorDetailsAttributes](hcm.ref.Append("error_details"))
}

func (hcm dataHybridComputeMachineAttributes) Identity() terra.ListValue[datahybridcomputemachine.IdentityAttributes] {
	return terra.ReferenceList[datahybridcomputemachine.IdentityAttributes](hcm.ref.Append("identity"))
}

func (hcm dataHybridComputeMachineAttributes) LocationData() terra.ListValue[datahybridcomputemachine.LocationDataAttributes] {
	return terra.ReferenceList[datahybridcomputemachine.LocationDataAttributes](hcm.ref.Append("location_data"))
}

func (hcm dataHybridComputeMachineAttributes) OsProfile() terra.ListValue[datahybridcomputemachine.OsProfileAttributes] {
	return terra.ReferenceList[datahybridcomputemachine.OsProfileAttributes](hcm.ref.Append("os_profile"))
}

func (hcm dataHybridComputeMachineAttributes) ServiceStatus() terra.ListValue[datahybridcomputemachine.ServiceStatusAttributes] {
	return terra.ReferenceList[datahybridcomputemachine.ServiceStatusAttributes](hcm.ref.Append("service_status"))
}

func (hcm dataHybridComputeMachineAttributes) Timeouts() datahybridcomputemachine.TimeoutsAttributes {
	return terra.ReferenceSingle[datahybridcomputemachine.TimeoutsAttributes](hcm.ref.Append("timeouts"))
}

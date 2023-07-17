// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataarcmachine "github.com/golingon/terraproviders/azurerm/3.65.0/dataarcmachine"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataArcMachine creates a new instance of [DataArcMachine].
func NewDataArcMachine(name string, args DataArcMachineArgs) *DataArcMachine {
	return &DataArcMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataArcMachine)(nil)

// DataArcMachine represents the Terraform data resource azurerm_arc_machine.
type DataArcMachine struct {
	Name string
	Args DataArcMachineArgs
}

// DataSource returns the Terraform object type for [DataArcMachine].
func (am *DataArcMachine) DataSource() string {
	return "azurerm_arc_machine"
}

// LocalName returns the local name for [DataArcMachine].
func (am *DataArcMachine) LocalName() string {
	return am.Name
}

// Configuration returns the configuration (args) for [DataArcMachine].
func (am *DataArcMachine) Configuration() interface{} {
	return am.Args
}

// Attributes returns the attributes for [DataArcMachine].
func (am *DataArcMachine) Attributes() dataArcMachineAttributes {
	return dataArcMachineAttributes{ref: terra.ReferenceDataResource(am)}
}

// DataArcMachineArgs contains the configurations for azurerm_arc_machine.
type DataArcMachineArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Agent: min=0
	Agent []dataarcmachine.Agent `hcl:"agent,block" validate:"min=0"`
	// CloudMetadata: min=0
	CloudMetadata []dataarcmachine.CloudMetadata `hcl:"cloud_metadata,block" validate:"min=0"`
	// Identity: min=0
	Identity []dataarcmachine.Identity `hcl:"identity,block" validate:"min=0"`
	// LocationData: min=0
	LocationData []dataarcmachine.LocationData `hcl:"location_data,block" validate:"min=0"`
	// OsProfile: min=0
	OsProfile []dataarcmachine.OsProfile `hcl:"os_profile,block" validate:"min=0"`
	// ServiceStatus: min=0
	ServiceStatus []dataarcmachine.ServiceStatus `hcl:"service_status,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataarcmachine.Timeouts `hcl:"timeouts,block"`
}
type dataArcMachineAttributes struct {
	ref terra.Reference
}

// ActiveDirectoryFqdn returns a reference to field active_directory_fqdn of azurerm_arc_machine.
func (am dataArcMachineAttributes) ActiveDirectoryFqdn() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("active_directory_fqdn"))
}

// AgentVersion returns a reference to field agent_version of azurerm_arc_machine.
func (am dataArcMachineAttributes) AgentVersion() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("agent_version"))
}

// ClientPublicKey returns a reference to field client_public_key of azurerm_arc_machine.
func (am dataArcMachineAttributes) ClientPublicKey() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("client_public_key"))
}

// DetectedProperties returns a reference to field detected_properties of azurerm_arc_machine.
func (am dataArcMachineAttributes) DetectedProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](am.ref.Append("detected_properties"))
}

// DisplayName returns a reference to field display_name of azurerm_arc_machine.
func (am dataArcMachineAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("display_name"))
}

// DnsFqdn returns a reference to field dns_fqdn of azurerm_arc_machine.
func (am dataArcMachineAttributes) DnsFqdn() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("dns_fqdn"))
}

// DomainName returns a reference to field domain_name of azurerm_arc_machine.
func (am dataArcMachineAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("domain_name"))
}

// Id returns a reference to field id of azurerm_arc_machine.
func (am dataArcMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("id"))
}

// LastStatusChangeTime returns a reference to field last_status_change_time of azurerm_arc_machine.
func (am dataArcMachineAttributes) LastStatusChangeTime() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("last_status_change_time"))
}

// Location returns a reference to field location of azurerm_arc_machine.
func (am dataArcMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("location"))
}

// MachineFqdn returns a reference to field machine_fqdn of azurerm_arc_machine.
func (am dataArcMachineAttributes) MachineFqdn() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("machine_fqdn"))
}

// MssqlDiscovered returns a reference to field mssql_discovered of azurerm_arc_machine.
func (am dataArcMachineAttributes) MssqlDiscovered() terra.BoolValue {
	return terra.ReferenceAsBool(am.ref.Append("mssql_discovered"))
}

// Name returns a reference to field name of azurerm_arc_machine.
func (am dataArcMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("name"))
}

// OsName returns a reference to field os_name of azurerm_arc_machine.
func (am dataArcMachineAttributes) OsName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("os_name"))
}

// OsSku returns a reference to field os_sku of azurerm_arc_machine.
func (am dataArcMachineAttributes) OsSku() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("os_sku"))
}

// OsType returns a reference to field os_type of azurerm_arc_machine.
func (am dataArcMachineAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("os_type"))
}

// OsVersion returns a reference to field os_version of azurerm_arc_machine.
func (am dataArcMachineAttributes) OsVersion() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("os_version"))
}

// ParentClusterResourceId returns a reference to field parent_cluster_resource_id of azurerm_arc_machine.
func (am dataArcMachineAttributes) ParentClusterResourceId() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("parent_cluster_resource_id"))
}

// PrivateLinkScopeResourceId returns a reference to field private_link_scope_resource_id of azurerm_arc_machine.
func (am dataArcMachineAttributes) PrivateLinkScopeResourceId() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("private_link_scope_resource_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_arc_machine.
func (am dataArcMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("resource_group_name"))
}

// Status returns a reference to field status of azurerm_arc_machine.
func (am dataArcMachineAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("status"))
}

// Tags returns a reference to field tags of azurerm_arc_machine.
func (am dataArcMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](am.ref.Append("tags"))
}

// VmId returns a reference to field vm_id of azurerm_arc_machine.
func (am dataArcMachineAttributes) VmId() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("vm_id"))
}

// VmUuid returns a reference to field vm_uuid of azurerm_arc_machine.
func (am dataArcMachineAttributes) VmUuid() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("vm_uuid"))
}

func (am dataArcMachineAttributes) Agent() terra.ListValue[dataarcmachine.AgentAttributes] {
	return terra.ReferenceAsList[dataarcmachine.AgentAttributes](am.ref.Append("agent"))
}

func (am dataArcMachineAttributes) CloudMetadata() terra.ListValue[dataarcmachine.CloudMetadataAttributes] {
	return terra.ReferenceAsList[dataarcmachine.CloudMetadataAttributes](am.ref.Append("cloud_metadata"))
}

func (am dataArcMachineAttributes) Identity() terra.ListValue[dataarcmachine.IdentityAttributes] {
	return terra.ReferenceAsList[dataarcmachine.IdentityAttributes](am.ref.Append("identity"))
}

func (am dataArcMachineAttributes) LocationData() terra.ListValue[dataarcmachine.LocationDataAttributes] {
	return terra.ReferenceAsList[dataarcmachine.LocationDataAttributes](am.ref.Append("location_data"))
}

func (am dataArcMachineAttributes) OsProfile() terra.ListValue[dataarcmachine.OsProfileAttributes] {
	return terra.ReferenceAsList[dataarcmachine.OsProfileAttributes](am.ref.Append("os_profile"))
}

func (am dataArcMachineAttributes) ServiceStatus() terra.ListValue[dataarcmachine.ServiceStatusAttributes] {
	return terra.ReferenceAsList[dataarcmachine.ServiceStatusAttributes](am.ref.Append("service_status"))
}

func (am dataArcMachineAttributes) Timeouts() dataarcmachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataarcmachine.TimeoutsAttributes](am.ref.Append("timeouts"))
}

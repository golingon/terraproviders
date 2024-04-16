// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_arc_machine

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_arc_machine.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aam *DataSource) DataSource() string {
	return "azurerm_arc_machine"
}

// LocalName returns the local name for [DataSource].
func (aam *DataSource) LocalName() string {
	return aam.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aam *DataSource) Configuration() interface{} {
	return aam.Args
}

// Attributes returns the attributes for [DataSource].
func (aam *DataSource) Attributes() dataAzurermArcMachineAttributes {
	return dataAzurermArcMachineAttributes{ref: terra.ReferenceDataSource(aam)}
}

// DataArgs contains the configurations for azurerm_arc_machine.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermArcMachineAttributes struct {
	ref terra.Reference
}

// ActiveDirectoryFqdn returns a reference to field active_directory_fqdn of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) ActiveDirectoryFqdn() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("active_directory_fqdn"))
}

// AgentVersion returns a reference to field agent_version of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) AgentVersion() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("agent_version"))
}

// ClientPublicKey returns a reference to field client_public_key of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) ClientPublicKey() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("client_public_key"))
}

// DetectedProperties returns a reference to field detected_properties of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) DetectedProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aam.ref.Append("detected_properties"))
}

// DisplayName returns a reference to field display_name of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("display_name"))
}

// DnsFqdn returns a reference to field dns_fqdn of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) DnsFqdn() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("dns_fqdn"))
}

// DomainName returns a reference to field domain_name of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("domain_name"))
}

// Id returns a reference to field id of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("id"))
}

// LastStatusChangeTime returns a reference to field last_status_change_time of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) LastStatusChangeTime() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("last_status_change_time"))
}

// Location returns a reference to field location of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("location"))
}

// MachineFqdn returns a reference to field machine_fqdn of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) MachineFqdn() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("machine_fqdn"))
}

// MssqlDiscovered returns a reference to field mssql_discovered of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) MssqlDiscovered() terra.BoolValue {
	return terra.ReferenceAsBool(aam.ref.Append("mssql_discovered"))
}

// Name returns a reference to field name of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("name"))
}

// OsName returns a reference to field os_name of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) OsName() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("os_name"))
}

// OsSku returns a reference to field os_sku of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) OsSku() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("os_sku"))
}

// OsType returns a reference to field os_type of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("os_type"))
}

// OsVersion returns a reference to field os_version of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) OsVersion() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("os_version"))
}

// ParentClusterResourceId returns a reference to field parent_cluster_resource_id of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) ParentClusterResourceId() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("parent_cluster_resource_id"))
}

// PrivateLinkScopeResourceId returns a reference to field private_link_scope_resource_id of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) PrivateLinkScopeResourceId() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("private_link_scope_resource_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("resource_group_name"))
}

// Status returns a reference to field status of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("status"))
}

// Tags returns a reference to field tags of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aam.ref.Append("tags"))
}

// VmId returns a reference to field vm_id of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) VmId() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("vm_id"))
}

// VmUuid returns a reference to field vm_uuid of azurerm_arc_machine.
func (aam dataAzurermArcMachineAttributes) VmUuid() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("vm_uuid"))
}

func (aam dataAzurermArcMachineAttributes) Agent() terra.ListValue[DataAgentAttributes] {
	return terra.ReferenceAsList[DataAgentAttributes](aam.ref.Append("agent"))
}

func (aam dataAzurermArcMachineAttributes) CloudMetadata() terra.ListValue[DataCloudMetadataAttributes] {
	return terra.ReferenceAsList[DataCloudMetadataAttributes](aam.ref.Append("cloud_metadata"))
}

func (aam dataAzurermArcMachineAttributes) Identity() terra.ListValue[DataIdentityAttributes] {
	return terra.ReferenceAsList[DataIdentityAttributes](aam.ref.Append("identity"))
}

func (aam dataAzurermArcMachineAttributes) LocationData() terra.ListValue[DataLocationDataAttributes] {
	return terra.ReferenceAsList[DataLocationDataAttributes](aam.ref.Append("location_data"))
}

func (aam dataAzurermArcMachineAttributes) OsProfile() terra.ListValue[DataOsProfileAttributes] {
	return terra.ReferenceAsList[DataOsProfileAttributes](aam.ref.Append("os_profile"))
}

func (aam dataAzurermArcMachineAttributes) ServiceStatus() terra.ListValue[DataServiceStatusAttributes] {
	return terra.ReferenceAsList[DataServiceStatusAttributes](aam.ref.Append("service_status"))
}

func (aam dataAzurermArcMachineAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](aam.ref.Append("timeouts"))
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataactivedirectorydomainservice "github.com/golingon/terraproviders/azurerm/3.55.0/dataactivedirectorydomainservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataActiveDirectoryDomainService creates a new instance of [DataActiveDirectoryDomainService].
func NewDataActiveDirectoryDomainService(name string, args DataActiveDirectoryDomainServiceArgs) *DataActiveDirectoryDomainService {
	return &DataActiveDirectoryDomainService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataActiveDirectoryDomainService)(nil)

// DataActiveDirectoryDomainService represents the Terraform data resource azurerm_active_directory_domain_service.
type DataActiveDirectoryDomainService struct {
	Name string
	Args DataActiveDirectoryDomainServiceArgs
}

// DataSource returns the Terraform object type for [DataActiveDirectoryDomainService].
func (adds *DataActiveDirectoryDomainService) DataSource() string {
	return "azurerm_active_directory_domain_service"
}

// LocalName returns the local name for [DataActiveDirectoryDomainService].
func (adds *DataActiveDirectoryDomainService) LocalName() string {
	return adds.Name
}

// Configuration returns the configuration (args) for [DataActiveDirectoryDomainService].
func (adds *DataActiveDirectoryDomainService) Configuration() interface{} {
	return adds.Args
}

// Attributes returns the attributes for [DataActiveDirectoryDomainService].
func (adds *DataActiveDirectoryDomainService) Attributes() dataActiveDirectoryDomainServiceAttributes {
	return dataActiveDirectoryDomainServiceAttributes{ref: terra.ReferenceDataResource(adds)}
}

// DataActiveDirectoryDomainServiceArgs contains the configurations for azurerm_active_directory_domain_service.
type DataActiveDirectoryDomainServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Notifications: min=0
	Notifications []dataactivedirectorydomainservice.Notifications `hcl:"notifications,block" validate:"min=0"`
	// ReplicaSets: min=0
	ReplicaSets []dataactivedirectorydomainservice.ReplicaSets `hcl:"replica_sets,block" validate:"min=0"`
	// SecureLdap: min=0
	SecureLdap []dataactivedirectorydomainservice.SecureLdap `hcl:"secure_ldap,block" validate:"min=0"`
	// Security: min=0
	Security []dataactivedirectorydomainservice.Security `hcl:"security,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataactivedirectorydomainservice.Timeouts `hcl:"timeouts,block"`
}
type dataActiveDirectoryDomainServiceAttributes struct {
	ref terra.Reference
}

// DeploymentId returns a reference to field deployment_id of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) DeploymentId() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("deployment_id"))
}

// DomainConfigurationType returns a reference to field domain_configuration_type of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) DomainConfigurationType() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("domain_configuration_type"))
}

// DomainName returns a reference to field domain_name of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("domain_name"))
}

// FilteredSyncEnabled returns a reference to field filtered_sync_enabled of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) FilteredSyncEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(adds.ref.Append("filtered_sync_enabled"))
}

// Id returns a reference to field id of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("resource_group_name"))
}

// ResourceId returns a reference to field resource_id of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("resource_id"))
}

// Sku returns a reference to field sku of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("sku"))
}

// SyncOwner returns a reference to field sync_owner of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) SyncOwner() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("sync_owner"))
}

// Tags returns a reference to field tags of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adds.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("tenant_id"))
}

// Version returns a reference to field version of azurerm_active_directory_domain_service.
func (adds dataActiveDirectoryDomainServiceAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(adds.ref.Append("version"))
}

func (adds dataActiveDirectoryDomainServiceAttributes) Notifications() terra.ListValue[dataactivedirectorydomainservice.NotificationsAttributes] {
	return terra.ReferenceAsList[dataactivedirectorydomainservice.NotificationsAttributes](adds.ref.Append("notifications"))
}

func (adds dataActiveDirectoryDomainServiceAttributes) ReplicaSets() terra.ListValue[dataactivedirectorydomainservice.ReplicaSetsAttributes] {
	return terra.ReferenceAsList[dataactivedirectorydomainservice.ReplicaSetsAttributes](adds.ref.Append("replica_sets"))
}

func (adds dataActiveDirectoryDomainServiceAttributes) SecureLdap() terra.ListValue[dataactivedirectorydomainservice.SecureLdapAttributes] {
	return terra.ReferenceAsList[dataactivedirectorydomainservice.SecureLdapAttributes](adds.ref.Append("secure_ldap"))
}

func (adds dataActiveDirectoryDomainServiceAttributes) Security() terra.ListValue[dataactivedirectorydomainservice.SecurityAttributes] {
	return terra.ReferenceAsList[dataactivedirectorydomainservice.SecurityAttributes](adds.ref.Append("security"))
}

func (adds dataActiveDirectoryDomainServiceAttributes) Timeouts() dataactivedirectorydomainservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataactivedirectorydomainservice.TimeoutsAttributes](adds.ref.Append("timeouts"))
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	activedirectorydomainservice "github.com/golingon/terraproviders/azurerm/3.69.0/activedirectorydomainservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewActiveDirectoryDomainService creates a new instance of [ActiveDirectoryDomainService].
func NewActiveDirectoryDomainService(name string, args ActiveDirectoryDomainServiceArgs) *ActiveDirectoryDomainService {
	return &ActiveDirectoryDomainService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ActiveDirectoryDomainService)(nil)

// ActiveDirectoryDomainService represents the Terraform resource azurerm_active_directory_domain_service.
type ActiveDirectoryDomainService struct {
	Name      string
	Args      ActiveDirectoryDomainServiceArgs
	state     *activeDirectoryDomainServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ActiveDirectoryDomainService].
func (adds *ActiveDirectoryDomainService) Type() string {
	return "azurerm_active_directory_domain_service"
}

// LocalName returns the local name for [ActiveDirectoryDomainService].
func (adds *ActiveDirectoryDomainService) LocalName() string {
	return adds.Name
}

// Configuration returns the configuration (args) for [ActiveDirectoryDomainService].
func (adds *ActiveDirectoryDomainService) Configuration() interface{} {
	return adds.Args
}

// DependOn is used for other resources to depend on [ActiveDirectoryDomainService].
func (adds *ActiveDirectoryDomainService) DependOn() terra.Reference {
	return terra.ReferenceResource(adds)
}

// Dependencies returns the list of resources [ActiveDirectoryDomainService] depends_on.
func (adds *ActiveDirectoryDomainService) Dependencies() terra.Dependencies {
	return adds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ActiveDirectoryDomainService].
func (adds *ActiveDirectoryDomainService) LifecycleManagement() *terra.Lifecycle {
	return adds.Lifecycle
}

// Attributes returns the attributes for [ActiveDirectoryDomainService].
func (adds *ActiveDirectoryDomainService) Attributes() activeDirectoryDomainServiceAttributes {
	return activeDirectoryDomainServiceAttributes{ref: terra.ReferenceResource(adds)}
}

// ImportState imports the given attribute values into [ActiveDirectoryDomainService]'s state.
func (adds *ActiveDirectoryDomainService) ImportState(av io.Reader) error {
	adds.state = &activeDirectoryDomainServiceState{}
	if err := json.NewDecoder(av).Decode(adds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adds.Type(), adds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ActiveDirectoryDomainService] has state.
func (adds *ActiveDirectoryDomainService) State() (*activeDirectoryDomainServiceState, bool) {
	return adds.state, adds.state != nil
}

// StateMust returns the state for [ActiveDirectoryDomainService]. Panics if the state is nil.
func (adds *ActiveDirectoryDomainService) StateMust() *activeDirectoryDomainServiceState {
	if adds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adds.Type(), adds.LocalName()))
	}
	return adds.state
}

// ActiveDirectoryDomainServiceArgs contains the configurations for azurerm_active_directory_domain_service.
type ActiveDirectoryDomainServiceArgs struct {
	// DomainConfigurationType: string, optional
	DomainConfigurationType terra.StringValue `hcl:"domain_configuration_type,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// FilteredSyncEnabled: bool, optional
	FilteredSyncEnabled terra.BoolValue `hcl:"filtered_sync_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// InitialReplicaSet: required
	InitialReplicaSet *activedirectorydomainservice.InitialReplicaSet `hcl:"initial_replica_set,block" validate:"required"`
	// Notifications: optional
	Notifications *activedirectorydomainservice.Notifications `hcl:"notifications,block"`
	// SecureLdap: optional
	SecureLdap *activedirectorydomainservice.SecureLdap `hcl:"secure_ldap,block"`
	// Security: optional
	Security *activedirectorydomainservice.Security `hcl:"security,block"`
	// Timeouts: optional
	Timeouts *activedirectorydomainservice.Timeouts `hcl:"timeouts,block"`
}
type activeDirectoryDomainServiceAttributes struct {
	ref terra.Reference
}

// DeploymentId returns a reference to field deployment_id of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) DeploymentId() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("deployment_id"))
}

// DomainConfigurationType returns a reference to field domain_configuration_type of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) DomainConfigurationType() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("domain_configuration_type"))
}

// DomainName returns a reference to field domain_name of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("domain_name"))
}

// FilteredSyncEnabled returns a reference to field filtered_sync_enabled of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) FilteredSyncEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(adds.ref.Append("filtered_sync_enabled"))
}

// Id returns a reference to field id of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("resource_group_name"))
}

// ResourceId returns a reference to field resource_id of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("resource_id"))
}

// Sku returns a reference to field sku of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("sku"))
}

// SyncOwner returns a reference to field sync_owner of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) SyncOwner() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("sync_owner"))
}

// Tags returns a reference to field tags of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adds.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(adds.ref.Append("tenant_id"))
}

// Version returns a reference to field version of azurerm_active_directory_domain_service.
func (adds activeDirectoryDomainServiceAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(adds.ref.Append("version"))
}

func (adds activeDirectoryDomainServiceAttributes) InitialReplicaSet() terra.ListValue[activedirectorydomainservice.InitialReplicaSetAttributes] {
	return terra.ReferenceAsList[activedirectorydomainservice.InitialReplicaSetAttributes](adds.ref.Append("initial_replica_set"))
}

func (adds activeDirectoryDomainServiceAttributes) Notifications() terra.ListValue[activedirectorydomainservice.NotificationsAttributes] {
	return terra.ReferenceAsList[activedirectorydomainservice.NotificationsAttributes](adds.ref.Append("notifications"))
}

func (adds activeDirectoryDomainServiceAttributes) SecureLdap() terra.ListValue[activedirectorydomainservice.SecureLdapAttributes] {
	return terra.ReferenceAsList[activedirectorydomainservice.SecureLdapAttributes](adds.ref.Append("secure_ldap"))
}

func (adds activeDirectoryDomainServiceAttributes) Security() terra.ListValue[activedirectorydomainservice.SecurityAttributes] {
	return terra.ReferenceAsList[activedirectorydomainservice.SecurityAttributes](adds.ref.Append("security"))
}

func (adds activeDirectoryDomainServiceAttributes) Timeouts() activedirectorydomainservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[activedirectorydomainservice.TimeoutsAttributes](adds.ref.Append("timeouts"))
}

type activeDirectoryDomainServiceState struct {
	DeploymentId            string                                                `json:"deployment_id"`
	DomainConfigurationType string                                                `json:"domain_configuration_type"`
	DomainName              string                                                `json:"domain_name"`
	FilteredSyncEnabled     bool                                                  `json:"filtered_sync_enabled"`
	Id                      string                                                `json:"id"`
	Location                string                                                `json:"location"`
	Name                    string                                                `json:"name"`
	ResourceGroupName       string                                                `json:"resource_group_name"`
	ResourceId              string                                                `json:"resource_id"`
	Sku                     string                                                `json:"sku"`
	SyncOwner               string                                                `json:"sync_owner"`
	Tags                    map[string]string                                     `json:"tags"`
	TenantId                string                                                `json:"tenant_id"`
	Version                 float64                                               `json:"version"`
	InitialReplicaSet       []activedirectorydomainservice.InitialReplicaSetState `json:"initial_replica_set"`
	Notifications           []activedirectorydomainservice.NotificationsState     `json:"notifications"`
	SecureLdap              []activedirectorydomainservice.SecureLdapState        `json:"secure_ldap"`
	Security                []activedirectorydomainservice.SecurityState          `json:"security"`
	Timeouts                *activedirectorydomainservice.TimeoutsState           `json:"timeouts"`
}

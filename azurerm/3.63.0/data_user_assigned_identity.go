// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datauserassignedidentity "github.com/golingon/terraproviders/azurerm/3.63.0/datauserassignedidentity"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataUserAssignedIdentity creates a new instance of [DataUserAssignedIdentity].
func NewDataUserAssignedIdentity(name string, args DataUserAssignedIdentityArgs) *DataUserAssignedIdentity {
	return &DataUserAssignedIdentity{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataUserAssignedIdentity)(nil)

// DataUserAssignedIdentity represents the Terraform data resource azurerm_user_assigned_identity.
type DataUserAssignedIdentity struct {
	Name string
	Args DataUserAssignedIdentityArgs
}

// DataSource returns the Terraform object type for [DataUserAssignedIdentity].
func (uai *DataUserAssignedIdentity) DataSource() string {
	return "azurerm_user_assigned_identity"
}

// LocalName returns the local name for [DataUserAssignedIdentity].
func (uai *DataUserAssignedIdentity) LocalName() string {
	return uai.Name
}

// Configuration returns the configuration (args) for [DataUserAssignedIdentity].
func (uai *DataUserAssignedIdentity) Configuration() interface{} {
	return uai.Args
}

// Attributes returns the attributes for [DataUserAssignedIdentity].
func (uai *DataUserAssignedIdentity) Attributes() dataUserAssignedIdentityAttributes {
	return dataUserAssignedIdentityAttributes{ref: terra.ReferenceDataResource(uai)}
}

// DataUserAssignedIdentityArgs contains the configurations for azurerm_user_assigned_identity.
type DataUserAssignedIdentityArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datauserassignedidentity.Timeouts `hcl:"timeouts,block"`
}
type dataUserAssignedIdentityAttributes struct {
	ref terra.Reference
}

// ClientId returns a reference to field client_id of azurerm_user_assigned_identity.
func (uai dataUserAssignedIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("client_id"))
}

// Id returns a reference to field id of azurerm_user_assigned_identity.
func (uai dataUserAssignedIdentityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_user_assigned_identity.
func (uai dataUserAssignedIdentityAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_user_assigned_identity.
func (uai dataUserAssignedIdentityAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("name"))
}

// PrincipalId returns a reference to field principal_id of azurerm_user_assigned_identity.
func (uai dataUserAssignedIdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("principal_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_user_assigned_identity.
func (uai dataUserAssignedIdentityAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_user_assigned_identity.
func (uai dataUserAssignedIdentityAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](uai.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_user_assigned_identity.
func (uai dataUserAssignedIdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("tenant_id"))
}

func (uai dataUserAssignedIdentityAttributes) Timeouts() datauserassignedidentity.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datauserassignedidentity.TimeoutsAttributes](uai.ref.Append("timeouts"))
}

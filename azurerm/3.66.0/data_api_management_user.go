// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapimanagementuser "github.com/golingon/terraproviders/azurerm/3.66.0/dataapimanagementuser"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApiManagementUser creates a new instance of [DataApiManagementUser].
func NewDataApiManagementUser(name string, args DataApiManagementUserArgs) *DataApiManagementUser {
	return &DataApiManagementUser{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiManagementUser)(nil)

// DataApiManagementUser represents the Terraform data resource azurerm_api_management_user.
type DataApiManagementUser struct {
	Name string
	Args DataApiManagementUserArgs
}

// DataSource returns the Terraform object type for [DataApiManagementUser].
func (amu *DataApiManagementUser) DataSource() string {
	return "azurerm_api_management_user"
}

// LocalName returns the local name for [DataApiManagementUser].
func (amu *DataApiManagementUser) LocalName() string {
	return amu.Name
}

// Configuration returns the configuration (args) for [DataApiManagementUser].
func (amu *DataApiManagementUser) Configuration() interface{} {
	return amu.Args
}

// Attributes returns the attributes for [DataApiManagementUser].
func (amu *DataApiManagementUser) Attributes() dataApiManagementUserAttributes {
	return dataApiManagementUserAttributes{ref: terra.ReferenceDataResource(amu)}
}

// DataApiManagementUserArgs contains the configurations for azurerm_api_management_user.
type DataApiManagementUserArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// UserId: string, required
	UserId terra.StringValue `hcl:"user_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataapimanagementuser.Timeouts `hcl:"timeouts,block"`
}
type dataApiManagementUserAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_user.
func (amu dataApiManagementUserAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("api_management_name"))
}

// Email returns a reference to field email of azurerm_api_management_user.
func (amu dataApiManagementUserAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("email"))
}

// FirstName returns a reference to field first_name of azurerm_api_management_user.
func (amu dataApiManagementUserAttributes) FirstName() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("first_name"))
}

// Id returns a reference to field id of azurerm_api_management_user.
func (amu dataApiManagementUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("id"))
}

// LastName returns a reference to field last_name of azurerm_api_management_user.
func (amu dataApiManagementUserAttributes) LastName() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("last_name"))
}

// Note returns a reference to field note of azurerm_api_management_user.
func (amu dataApiManagementUserAttributes) Note() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("note"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_user.
func (amu dataApiManagementUserAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("resource_group_name"))
}

// State returns a reference to field state of azurerm_api_management_user.
func (amu dataApiManagementUserAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("state"))
}

// UserId returns a reference to field user_id of azurerm_api_management_user.
func (amu dataApiManagementUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("user_id"))
}

func (amu dataApiManagementUserAttributes) Timeouts() dataapimanagementuser.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapimanagementuser.TimeoutsAttributes](amu.ref.Append("timeouts"))
}

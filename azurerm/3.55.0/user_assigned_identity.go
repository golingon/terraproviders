// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	userassignedidentity "github.com/golingon/terraproviders/azurerm/3.55.0/userassignedidentity"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewUserAssignedIdentity creates a new instance of [UserAssignedIdentity].
func NewUserAssignedIdentity(name string, args UserAssignedIdentityArgs) *UserAssignedIdentity {
	return &UserAssignedIdentity{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*UserAssignedIdentity)(nil)

// UserAssignedIdentity represents the Terraform resource azurerm_user_assigned_identity.
type UserAssignedIdentity struct {
	Name      string
	Args      UserAssignedIdentityArgs
	state     *userAssignedIdentityState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [UserAssignedIdentity].
func (uai *UserAssignedIdentity) Type() string {
	return "azurerm_user_assigned_identity"
}

// LocalName returns the local name for [UserAssignedIdentity].
func (uai *UserAssignedIdentity) LocalName() string {
	return uai.Name
}

// Configuration returns the configuration (args) for [UserAssignedIdentity].
func (uai *UserAssignedIdentity) Configuration() interface{} {
	return uai.Args
}

// DependOn is used for other resources to depend on [UserAssignedIdentity].
func (uai *UserAssignedIdentity) DependOn() terra.Reference {
	return terra.ReferenceResource(uai)
}

// Dependencies returns the list of resources [UserAssignedIdentity] depends_on.
func (uai *UserAssignedIdentity) Dependencies() terra.Dependencies {
	return uai.DependsOn
}

// LifecycleManagement returns the lifecycle block for [UserAssignedIdentity].
func (uai *UserAssignedIdentity) LifecycleManagement() *terra.Lifecycle {
	return uai.Lifecycle
}

// Attributes returns the attributes for [UserAssignedIdentity].
func (uai *UserAssignedIdentity) Attributes() userAssignedIdentityAttributes {
	return userAssignedIdentityAttributes{ref: terra.ReferenceResource(uai)}
}

// ImportState imports the given attribute values into [UserAssignedIdentity]'s state.
func (uai *UserAssignedIdentity) ImportState(av io.Reader) error {
	uai.state = &userAssignedIdentityState{}
	if err := json.NewDecoder(av).Decode(uai.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", uai.Type(), uai.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [UserAssignedIdentity] has state.
func (uai *UserAssignedIdentity) State() (*userAssignedIdentityState, bool) {
	return uai.state, uai.state != nil
}

// StateMust returns the state for [UserAssignedIdentity]. Panics if the state is nil.
func (uai *UserAssignedIdentity) StateMust() *userAssignedIdentityState {
	if uai.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", uai.Type(), uai.LocalName()))
	}
	return uai.state
}

// UserAssignedIdentityArgs contains the configurations for azurerm_user_assigned_identity.
type UserAssignedIdentityArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *userassignedidentity.Timeouts `hcl:"timeouts,block"`
}
type userAssignedIdentityAttributes struct {
	ref terra.Reference
}

// ClientId returns a reference to field client_id of azurerm_user_assigned_identity.
func (uai userAssignedIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("client_id"))
}

// Id returns a reference to field id of azurerm_user_assigned_identity.
func (uai userAssignedIdentityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_user_assigned_identity.
func (uai userAssignedIdentityAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_user_assigned_identity.
func (uai userAssignedIdentityAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("name"))
}

// PrincipalId returns a reference to field principal_id of azurerm_user_assigned_identity.
func (uai userAssignedIdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("principal_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_user_assigned_identity.
func (uai userAssignedIdentityAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_user_assigned_identity.
func (uai userAssignedIdentityAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](uai.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_user_assigned_identity.
func (uai userAssignedIdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(uai.ref.Append("tenant_id"))
}

func (uai userAssignedIdentityAttributes) Timeouts() userassignedidentity.TimeoutsAttributes {
	return terra.ReferenceAsSingle[userassignedidentity.TimeoutsAttributes](uai.ref.Append("timeouts"))
}

type userAssignedIdentityState struct {
	ClientId          string                              `json:"client_id"`
	Id                string                              `json:"id"`
	Location          string                              `json:"location"`
	Name              string                              `json:"name"`
	PrincipalId       string                              `json:"principal_id"`
	ResourceGroupName string                              `json:"resource_group_name"`
	Tags              map[string]string                   `json:"tags"`
	TenantId          string                              `json:"tenant_id"`
	Timeouts          *userassignedidentity.TimeoutsState `json:"timeouts"`
}
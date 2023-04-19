// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	serviceprincipaldelegatedpermissiongrant "github.com/golingon/terraproviders/azuread/2.37.1/serviceprincipaldelegatedpermissiongrant"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicePrincipalDelegatedPermissionGrant creates a new instance of [ServicePrincipalDelegatedPermissionGrant].
func NewServicePrincipalDelegatedPermissionGrant(name string, args ServicePrincipalDelegatedPermissionGrantArgs) *ServicePrincipalDelegatedPermissionGrant {
	return &ServicePrincipalDelegatedPermissionGrant{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicePrincipalDelegatedPermissionGrant)(nil)

// ServicePrincipalDelegatedPermissionGrant represents the Terraform resource azuread_service_principal_delegated_permission_grant.
type ServicePrincipalDelegatedPermissionGrant struct {
	Name      string
	Args      ServicePrincipalDelegatedPermissionGrantArgs
	state     *servicePrincipalDelegatedPermissionGrantState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicePrincipalDelegatedPermissionGrant].
func (spdpg *ServicePrincipalDelegatedPermissionGrant) Type() string {
	return "azuread_service_principal_delegated_permission_grant"
}

// LocalName returns the local name for [ServicePrincipalDelegatedPermissionGrant].
func (spdpg *ServicePrincipalDelegatedPermissionGrant) LocalName() string {
	return spdpg.Name
}

// Configuration returns the configuration (args) for [ServicePrincipalDelegatedPermissionGrant].
func (spdpg *ServicePrincipalDelegatedPermissionGrant) Configuration() interface{} {
	return spdpg.Args
}

// DependOn is used for other resources to depend on [ServicePrincipalDelegatedPermissionGrant].
func (spdpg *ServicePrincipalDelegatedPermissionGrant) DependOn() terra.Reference {
	return terra.ReferenceResource(spdpg)
}

// Dependencies returns the list of resources [ServicePrincipalDelegatedPermissionGrant] depends_on.
func (spdpg *ServicePrincipalDelegatedPermissionGrant) Dependencies() terra.Dependencies {
	return spdpg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicePrincipalDelegatedPermissionGrant].
func (spdpg *ServicePrincipalDelegatedPermissionGrant) LifecycleManagement() *terra.Lifecycle {
	return spdpg.Lifecycle
}

// Attributes returns the attributes for [ServicePrincipalDelegatedPermissionGrant].
func (spdpg *ServicePrincipalDelegatedPermissionGrant) Attributes() servicePrincipalDelegatedPermissionGrantAttributes {
	return servicePrincipalDelegatedPermissionGrantAttributes{ref: terra.ReferenceResource(spdpg)}
}

// ImportState imports the given attribute values into [ServicePrincipalDelegatedPermissionGrant]'s state.
func (spdpg *ServicePrincipalDelegatedPermissionGrant) ImportState(av io.Reader) error {
	spdpg.state = &servicePrincipalDelegatedPermissionGrantState{}
	if err := json.NewDecoder(av).Decode(spdpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spdpg.Type(), spdpg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicePrincipalDelegatedPermissionGrant] has state.
func (spdpg *ServicePrincipalDelegatedPermissionGrant) State() (*servicePrincipalDelegatedPermissionGrantState, bool) {
	return spdpg.state, spdpg.state != nil
}

// StateMust returns the state for [ServicePrincipalDelegatedPermissionGrant]. Panics if the state is nil.
func (spdpg *ServicePrincipalDelegatedPermissionGrant) StateMust() *servicePrincipalDelegatedPermissionGrantState {
	if spdpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spdpg.Type(), spdpg.LocalName()))
	}
	return spdpg.state
}

// ServicePrincipalDelegatedPermissionGrantArgs contains the configurations for azuread_service_principal_delegated_permission_grant.
type ServicePrincipalDelegatedPermissionGrantArgs struct {
	// ClaimValues: set of string, required
	ClaimValues terra.SetValue[terra.StringValue] `hcl:"claim_values,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceServicePrincipalObjectId: string, required
	ResourceServicePrincipalObjectId terra.StringValue `hcl:"resource_service_principal_object_id,attr" validate:"required"`
	// ServicePrincipalObjectId: string, required
	ServicePrincipalObjectId terra.StringValue `hcl:"service_principal_object_id,attr" validate:"required"`
	// UserObjectId: string, optional
	UserObjectId terra.StringValue `hcl:"user_object_id,attr"`
	// Timeouts: optional
	Timeouts *serviceprincipaldelegatedpermissiongrant.Timeouts `hcl:"timeouts,block"`
}
type servicePrincipalDelegatedPermissionGrantAttributes struct {
	ref terra.Reference
}

// ClaimValues returns a reference to field claim_values of azuread_service_principal_delegated_permission_grant.
func (spdpg servicePrincipalDelegatedPermissionGrantAttributes) ClaimValues() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](spdpg.ref.Append("claim_values"))
}

// Id returns a reference to field id of azuread_service_principal_delegated_permission_grant.
func (spdpg servicePrincipalDelegatedPermissionGrantAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spdpg.ref.Append("id"))
}

// ResourceServicePrincipalObjectId returns a reference to field resource_service_principal_object_id of azuread_service_principal_delegated_permission_grant.
func (spdpg servicePrincipalDelegatedPermissionGrantAttributes) ResourceServicePrincipalObjectId() terra.StringValue {
	return terra.ReferenceAsString(spdpg.ref.Append("resource_service_principal_object_id"))
}

// ServicePrincipalObjectId returns a reference to field service_principal_object_id of azuread_service_principal_delegated_permission_grant.
func (spdpg servicePrincipalDelegatedPermissionGrantAttributes) ServicePrincipalObjectId() terra.StringValue {
	return terra.ReferenceAsString(spdpg.ref.Append("service_principal_object_id"))
}

// UserObjectId returns a reference to field user_object_id of azuread_service_principal_delegated_permission_grant.
func (spdpg servicePrincipalDelegatedPermissionGrantAttributes) UserObjectId() terra.StringValue {
	return terra.ReferenceAsString(spdpg.ref.Append("user_object_id"))
}

func (spdpg servicePrincipalDelegatedPermissionGrantAttributes) Timeouts() serviceprincipaldelegatedpermissiongrant.TimeoutsAttributes {
	return terra.ReferenceAsSingle[serviceprincipaldelegatedpermissiongrant.TimeoutsAttributes](spdpg.ref.Append("timeouts"))
}

type servicePrincipalDelegatedPermissionGrantState struct {
	ClaimValues                      []string                                                `json:"claim_values"`
	Id                               string                                                  `json:"id"`
	ResourceServicePrincipalObjectId string                                                  `json:"resource_service_principal_object_id"`
	ServicePrincipalObjectId         string                                                  `json:"service_principal_object_id"`
	UserObjectId                     string                                                  `json:"user_object_id"`
	Timeouts                         *serviceprincipaldelegatedpermissiongrant.TimeoutsState `json:"timeouts"`
}

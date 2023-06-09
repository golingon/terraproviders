// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	accesspackagecatalogroleassignment "github.com/golingon/terraproviders/azuread/2.38.0/accesspackagecatalogroleassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessPackageCatalogRoleAssignment creates a new instance of [AccessPackageCatalogRoleAssignment].
func NewAccessPackageCatalogRoleAssignment(name string, args AccessPackageCatalogRoleAssignmentArgs) *AccessPackageCatalogRoleAssignment {
	return &AccessPackageCatalogRoleAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessPackageCatalogRoleAssignment)(nil)

// AccessPackageCatalogRoleAssignment represents the Terraform resource azuread_access_package_catalog_role_assignment.
type AccessPackageCatalogRoleAssignment struct {
	Name      string
	Args      AccessPackageCatalogRoleAssignmentArgs
	state     *accessPackageCatalogRoleAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessPackageCatalogRoleAssignment].
func (apcra *AccessPackageCatalogRoleAssignment) Type() string {
	return "azuread_access_package_catalog_role_assignment"
}

// LocalName returns the local name for [AccessPackageCatalogRoleAssignment].
func (apcra *AccessPackageCatalogRoleAssignment) LocalName() string {
	return apcra.Name
}

// Configuration returns the configuration (args) for [AccessPackageCatalogRoleAssignment].
func (apcra *AccessPackageCatalogRoleAssignment) Configuration() interface{} {
	return apcra.Args
}

// DependOn is used for other resources to depend on [AccessPackageCatalogRoleAssignment].
func (apcra *AccessPackageCatalogRoleAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(apcra)
}

// Dependencies returns the list of resources [AccessPackageCatalogRoleAssignment] depends_on.
func (apcra *AccessPackageCatalogRoleAssignment) Dependencies() terra.Dependencies {
	return apcra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessPackageCatalogRoleAssignment].
func (apcra *AccessPackageCatalogRoleAssignment) LifecycleManagement() *terra.Lifecycle {
	return apcra.Lifecycle
}

// Attributes returns the attributes for [AccessPackageCatalogRoleAssignment].
func (apcra *AccessPackageCatalogRoleAssignment) Attributes() accessPackageCatalogRoleAssignmentAttributes {
	return accessPackageCatalogRoleAssignmentAttributes{ref: terra.ReferenceResource(apcra)}
}

// ImportState imports the given attribute values into [AccessPackageCatalogRoleAssignment]'s state.
func (apcra *AccessPackageCatalogRoleAssignment) ImportState(av io.Reader) error {
	apcra.state = &accessPackageCatalogRoleAssignmentState{}
	if err := json.NewDecoder(av).Decode(apcra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", apcra.Type(), apcra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessPackageCatalogRoleAssignment] has state.
func (apcra *AccessPackageCatalogRoleAssignment) State() (*accessPackageCatalogRoleAssignmentState, bool) {
	return apcra.state, apcra.state != nil
}

// StateMust returns the state for [AccessPackageCatalogRoleAssignment]. Panics if the state is nil.
func (apcra *AccessPackageCatalogRoleAssignment) StateMust() *accessPackageCatalogRoleAssignmentState {
	if apcra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", apcra.Type(), apcra.LocalName()))
	}
	return apcra.state
}

// AccessPackageCatalogRoleAssignmentArgs contains the configurations for azuread_access_package_catalog_role_assignment.
type AccessPackageCatalogRoleAssignmentArgs struct {
	// CatalogId: string, required
	CatalogId terra.StringValue `hcl:"catalog_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrincipalObjectId: string, required
	PrincipalObjectId terra.StringValue `hcl:"principal_object_id,attr" validate:"required"`
	// RoleId: string, required
	RoleId terra.StringValue `hcl:"role_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *accesspackagecatalogroleassignment.Timeouts `hcl:"timeouts,block"`
}
type accessPackageCatalogRoleAssignmentAttributes struct {
	ref terra.Reference
}

// CatalogId returns a reference to field catalog_id of azuread_access_package_catalog_role_assignment.
func (apcra accessPackageCatalogRoleAssignmentAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(apcra.ref.Append("catalog_id"))
}

// Id returns a reference to field id of azuread_access_package_catalog_role_assignment.
func (apcra accessPackageCatalogRoleAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apcra.ref.Append("id"))
}

// PrincipalObjectId returns a reference to field principal_object_id of azuread_access_package_catalog_role_assignment.
func (apcra accessPackageCatalogRoleAssignmentAttributes) PrincipalObjectId() terra.StringValue {
	return terra.ReferenceAsString(apcra.ref.Append("principal_object_id"))
}

// RoleId returns a reference to field role_id of azuread_access_package_catalog_role_assignment.
func (apcra accessPackageCatalogRoleAssignmentAttributes) RoleId() terra.StringValue {
	return terra.ReferenceAsString(apcra.ref.Append("role_id"))
}

func (apcra accessPackageCatalogRoleAssignmentAttributes) Timeouts() accesspackagecatalogroleassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesspackagecatalogroleassignment.TimeoutsAttributes](apcra.ref.Append("timeouts"))
}

type accessPackageCatalogRoleAssignmentState struct {
	CatalogId         string                                            `json:"catalog_id"`
	Id                string                                            `json:"id"`
	PrincipalObjectId string                                            `json:"principal_object_id"`
	RoleId            string                                            `json:"role_id"`
	Timeouts          *accesspackagecatalogroleassignment.TimeoutsState `json:"timeouts"`
}

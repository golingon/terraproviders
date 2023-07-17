// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	accesspackageresourcepackageassociation "github.com/golingon/terraproviders/azuread/2.40.0/accesspackageresourcepackageassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessPackageResourcePackageAssociation creates a new instance of [AccessPackageResourcePackageAssociation].
func NewAccessPackageResourcePackageAssociation(name string, args AccessPackageResourcePackageAssociationArgs) *AccessPackageResourcePackageAssociation {
	return &AccessPackageResourcePackageAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessPackageResourcePackageAssociation)(nil)

// AccessPackageResourcePackageAssociation represents the Terraform resource azuread_access_package_resource_package_association.
type AccessPackageResourcePackageAssociation struct {
	Name      string
	Args      AccessPackageResourcePackageAssociationArgs
	state     *accessPackageResourcePackageAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessPackageResourcePackageAssociation].
func (aprpa *AccessPackageResourcePackageAssociation) Type() string {
	return "azuread_access_package_resource_package_association"
}

// LocalName returns the local name for [AccessPackageResourcePackageAssociation].
func (aprpa *AccessPackageResourcePackageAssociation) LocalName() string {
	return aprpa.Name
}

// Configuration returns the configuration (args) for [AccessPackageResourcePackageAssociation].
func (aprpa *AccessPackageResourcePackageAssociation) Configuration() interface{} {
	return aprpa.Args
}

// DependOn is used for other resources to depend on [AccessPackageResourcePackageAssociation].
func (aprpa *AccessPackageResourcePackageAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(aprpa)
}

// Dependencies returns the list of resources [AccessPackageResourcePackageAssociation] depends_on.
func (aprpa *AccessPackageResourcePackageAssociation) Dependencies() terra.Dependencies {
	return aprpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessPackageResourcePackageAssociation].
func (aprpa *AccessPackageResourcePackageAssociation) LifecycleManagement() *terra.Lifecycle {
	return aprpa.Lifecycle
}

// Attributes returns the attributes for [AccessPackageResourcePackageAssociation].
func (aprpa *AccessPackageResourcePackageAssociation) Attributes() accessPackageResourcePackageAssociationAttributes {
	return accessPackageResourcePackageAssociationAttributes{ref: terra.ReferenceResource(aprpa)}
}

// ImportState imports the given attribute values into [AccessPackageResourcePackageAssociation]'s state.
func (aprpa *AccessPackageResourcePackageAssociation) ImportState(av io.Reader) error {
	aprpa.state = &accessPackageResourcePackageAssociationState{}
	if err := json.NewDecoder(av).Decode(aprpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aprpa.Type(), aprpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessPackageResourcePackageAssociation] has state.
func (aprpa *AccessPackageResourcePackageAssociation) State() (*accessPackageResourcePackageAssociationState, bool) {
	return aprpa.state, aprpa.state != nil
}

// StateMust returns the state for [AccessPackageResourcePackageAssociation]. Panics if the state is nil.
func (aprpa *AccessPackageResourcePackageAssociation) StateMust() *accessPackageResourcePackageAssociationState {
	if aprpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aprpa.Type(), aprpa.LocalName()))
	}
	return aprpa.state
}

// AccessPackageResourcePackageAssociationArgs contains the configurations for azuread_access_package_resource_package_association.
type AccessPackageResourcePackageAssociationArgs struct {
	// AccessPackageId: string, required
	AccessPackageId terra.StringValue `hcl:"access_package_id,attr" validate:"required"`
	// AccessType: string, optional
	AccessType terra.StringValue `hcl:"access_type,attr"`
	// CatalogResourceAssociationId: string, required
	CatalogResourceAssociationId terra.StringValue `hcl:"catalog_resource_association_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *accesspackageresourcepackageassociation.Timeouts `hcl:"timeouts,block"`
}
type accessPackageResourcePackageAssociationAttributes struct {
	ref terra.Reference
}

// AccessPackageId returns a reference to field access_package_id of azuread_access_package_resource_package_association.
func (aprpa accessPackageResourcePackageAssociationAttributes) AccessPackageId() terra.StringValue {
	return terra.ReferenceAsString(aprpa.ref.Append("access_package_id"))
}

// AccessType returns a reference to field access_type of azuread_access_package_resource_package_association.
func (aprpa accessPackageResourcePackageAssociationAttributes) AccessType() terra.StringValue {
	return terra.ReferenceAsString(aprpa.ref.Append("access_type"))
}

// CatalogResourceAssociationId returns a reference to field catalog_resource_association_id of azuread_access_package_resource_package_association.
func (aprpa accessPackageResourcePackageAssociationAttributes) CatalogResourceAssociationId() terra.StringValue {
	return terra.ReferenceAsString(aprpa.ref.Append("catalog_resource_association_id"))
}

// Id returns a reference to field id of azuread_access_package_resource_package_association.
func (aprpa accessPackageResourcePackageAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aprpa.ref.Append("id"))
}

func (aprpa accessPackageResourcePackageAssociationAttributes) Timeouts() accesspackageresourcepackageassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesspackageresourcepackageassociation.TimeoutsAttributes](aprpa.ref.Append("timeouts"))
}

type accessPackageResourcePackageAssociationState struct {
	AccessPackageId              string                                                 `json:"access_package_id"`
	AccessType                   string                                                 `json:"access_type"`
	CatalogResourceAssociationId string                                                 `json:"catalog_resource_association_id"`
	Id                           string                                                 `json:"id"`
	Timeouts                     *accesspackageresourcepackageassociation.TimeoutsState `json:"timeouts"`
}

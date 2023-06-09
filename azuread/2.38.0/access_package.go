// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	accesspackage "github.com/golingon/terraproviders/azuread/2.38.0/accesspackage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessPackage creates a new instance of [AccessPackage].
func NewAccessPackage(name string, args AccessPackageArgs) *AccessPackage {
	return &AccessPackage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessPackage)(nil)

// AccessPackage represents the Terraform resource azuread_access_package.
type AccessPackage struct {
	Name      string
	Args      AccessPackageArgs
	state     *accessPackageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessPackage].
func (ap *AccessPackage) Type() string {
	return "azuread_access_package"
}

// LocalName returns the local name for [AccessPackage].
func (ap *AccessPackage) LocalName() string {
	return ap.Name
}

// Configuration returns the configuration (args) for [AccessPackage].
func (ap *AccessPackage) Configuration() interface{} {
	return ap.Args
}

// DependOn is used for other resources to depend on [AccessPackage].
func (ap *AccessPackage) DependOn() terra.Reference {
	return terra.ReferenceResource(ap)
}

// Dependencies returns the list of resources [AccessPackage] depends_on.
func (ap *AccessPackage) Dependencies() terra.Dependencies {
	return ap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessPackage].
func (ap *AccessPackage) LifecycleManagement() *terra.Lifecycle {
	return ap.Lifecycle
}

// Attributes returns the attributes for [AccessPackage].
func (ap *AccessPackage) Attributes() accessPackageAttributes {
	return accessPackageAttributes{ref: terra.ReferenceResource(ap)}
}

// ImportState imports the given attribute values into [AccessPackage]'s state.
func (ap *AccessPackage) ImportState(av io.Reader) error {
	ap.state = &accessPackageState{}
	if err := json.NewDecoder(av).Decode(ap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ap.Type(), ap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessPackage] has state.
func (ap *AccessPackage) State() (*accessPackageState, bool) {
	return ap.state, ap.state != nil
}

// StateMust returns the state for [AccessPackage]. Panics if the state is nil.
func (ap *AccessPackage) StateMust() *accessPackageState {
	if ap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ap.Type(), ap.LocalName()))
	}
	return ap.state
}

// AccessPackageArgs contains the configurations for azuread_access_package.
type AccessPackageArgs struct {
	// CatalogId: string, required
	CatalogId terra.StringValue `hcl:"catalog_id,attr" validate:"required"`
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Hidden: bool, optional
	Hidden terra.BoolValue `hcl:"hidden,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *accesspackage.Timeouts `hcl:"timeouts,block"`
}
type accessPackageAttributes struct {
	ref terra.Reference
}

// CatalogId returns a reference to field catalog_id of azuread_access_package.
func (ap accessPackageAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("catalog_id"))
}

// Description returns a reference to field description of azuread_access_package.
func (ap accessPackageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_access_package.
func (ap accessPackageAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("display_name"))
}

// Hidden returns a reference to field hidden of azuread_access_package.
func (ap accessPackageAttributes) Hidden() terra.BoolValue {
	return terra.ReferenceAsBool(ap.ref.Append("hidden"))
}

// Id returns a reference to field id of azuread_access_package.
func (ap accessPackageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("id"))
}

func (ap accessPackageAttributes) Timeouts() accesspackage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesspackage.TimeoutsAttributes](ap.ref.Append("timeouts"))
}

type accessPackageState struct {
	CatalogId   string                       `json:"catalog_id"`
	Description string                       `json:"description"`
	DisplayName string                       `json:"display_name"`
	Hidden      bool                         `json:"hidden"`
	Id          string                       `json:"id"`
	Timeouts    *accesspackage.TimeoutsState `json:"timeouts"`
}

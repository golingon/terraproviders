// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	dataaccesspackage "github.com/golingon/terraproviders/azuread/2.41.0/dataaccesspackage"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAccessPackage creates a new instance of [DataAccessPackage].
func NewDataAccessPackage(name string, args DataAccessPackageArgs) *DataAccessPackage {
	return &DataAccessPackage{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAccessPackage)(nil)

// DataAccessPackage represents the Terraform data resource azuread_access_package.
type DataAccessPackage struct {
	Name string
	Args DataAccessPackageArgs
}

// DataSource returns the Terraform object type for [DataAccessPackage].
func (ap *DataAccessPackage) DataSource() string {
	return "azuread_access_package"
}

// LocalName returns the local name for [DataAccessPackage].
func (ap *DataAccessPackage) LocalName() string {
	return ap.Name
}

// Configuration returns the configuration (args) for [DataAccessPackage].
func (ap *DataAccessPackage) Configuration() interface{} {
	return ap.Args
}

// Attributes returns the attributes for [DataAccessPackage].
func (ap *DataAccessPackage) Attributes() dataAccessPackageAttributes {
	return dataAccessPackageAttributes{ref: terra.ReferenceDataResource(ap)}
}

// DataAccessPackageArgs contains the configurations for azuread_access_package.
type DataAccessPackageArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// Timeouts: optional
	Timeouts *dataaccesspackage.Timeouts `hcl:"timeouts,block"`
}
type dataAccessPackageAttributes struct {
	ref terra.Reference
}

// CatalogId returns a reference to field catalog_id of azuread_access_package.
func (ap dataAccessPackageAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("catalog_id"))
}

// Description returns a reference to field description of azuread_access_package.
func (ap dataAccessPackageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_access_package.
func (ap dataAccessPackageAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("display_name"))
}

// Hidden returns a reference to field hidden of azuread_access_package.
func (ap dataAccessPackageAttributes) Hidden() terra.BoolValue {
	return terra.ReferenceAsBool(ap.ref.Append("hidden"))
}

// Id returns a reference to field id of azuread_access_package.
func (ap dataAccessPackageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("id"))
}

// ObjectId returns a reference to field object_id of azuread_access_package.
func (ap dataAccessPackageAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("object_id"))
}

func (ap dataAccessPackageAttributes) Timeouts() dataaccesspackage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataaccesspackage.TimeoutsAttributes](ap.ref.Append("timeouts"))
}

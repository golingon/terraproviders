// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	dataaccesspackagecatalog "github.com/golingon/terraproviders/azuread/2.38.0/dataaccesspackagecatalog"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAccessPackageCatalog creates a new instance of [DataAccessPackageCatalog].
func NewDataAccessPackageCatalog(name string, args DataAccessPackageCatalogArgs) *DataAccessPackageCatalog {
	return &DataAccessPackageCatalog{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAccessPackageCatalog)(nil)

// DataAccessPackageCatalog represents the Terraform data resource azuread_access_package_catalog.
type DataAccessPackageCatalog struct {
	Name string
	Args DataAccessPackageCatalogArgs
}

// DataSource returns the Terraform object type for [DataAccessPackageCatalog].
func (apc *DataAccessPackageCatalog) DataSource() string {
	return "azuread_access_package_catalog"
}

// LocalName returns the local name for [DataAccessPackageCatalog].
func (apc *DataAccessPackageCatalog) LocalName() string {
	return apc.Name
}

// Configuration returns the configuration (args) for [DataAccessPackageCatalog].
func (apc *DataAccessPackageCatalog) Configuration() interface{} {
	return apc.Args
}

// Attributes returns the attributes for [DataAccessPackageCatalog].
func (apc *DataAccessPackageCatalog) Attributes() dataAccessPackageCatalogAttributes {
	return dataAccessPackageCatalogAttributes{ref: terra.ReferenceDataResource(apc)}
}

// DataAccessPackageCatalogArgs contains the configurations for azuread_access_package_catalog.
type DataAccessPackageCatalogArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// Timeouts: optional
	Timeouts *dataaccesspackagecatalog.Timeouts `hcl:"timeouts,block"`
}
type dataAccessPackageCatalogAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azuread_access_package_catalog.
func (apc dataAccessPackageCatalogAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(apc.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_access_package_catalog.
func (apc dataAccessPackageCatalogAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(apc.ref.Append("display_name"))
}

// ExternallyVisible returns a reference to field externally_visible of azuread_access_package_catalog.
func (apc dataAccessPackageCatalogAttributes) ExternallyVisible() terra.BoolValue {
	return terra.ReferenceAsBool(apc.ref.Append("externally_visible"))
}

// Id returns a reference to field id of azuread_access_package_catalog.
func (apc dataAccessPackageCatalogAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apc.ref.Append("id"))
}

// ObjectId returns a reference to field object_id of azuread_access_package_catalog.
func (apc dataAccessPackageCatalogAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(apc.ref.Append("object_id"))
}

// Published returns a reference to field published of azuread_access_package_catalog.
func (apc dataAccessPackageCatalogAttributes) Published() terra.BoolValue {
	return terra.ReferenceAsBool(apc.ref.Append("published"))
}

func (apc dataAccessPackageCatalogAttributes) Timeouts() dataaccesspackagecatalog.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataaccesspackagecatalog.TimeoutsAttributes](apc.ref.Append("timeouts"))
}

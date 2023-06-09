// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	dataadministrativeunit "github.com/golingon/terraproviders/azuread/2.39.0/dataadministrativeunit"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAdministrativeUnit creates a new instance of [DataAdministrativeUnit].
func NewDataAdministrativeUnit(name string, args DataAdministrativeUnitArgs) *DataAdministrativeUnit {
	return &DataAdministrativeUnit{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAdministrativeUnit)(nil)

// DataAdministrativeUnit represents the Terraform data resource azuread_administrative_unit.
type DataAdministrativeUnit struct {
	Name string
	Args DataAdministrativeUnitArgs
}

// DataSource returns the Terraform object type for [DataAdministrativeUnit].
func (au *DataAdministrativeUnit) DataSource() string {
	return "azuread_administrative_unit"
}

// LocalName returns the local name for [DataAdministrativeUnit].
func (au *DataAdministrativeUnit) LocalName() string {
	return au.Name
}

// Configuration returns the configuration (args) for [DataAdministrativeUnit].
func (au *DataAdministrativeUnit) Configuration() interface{} {
	return au.Args
}

// Attributes returns the attributes for [DataAdministrativeUnit].
func (au *DataAdministrativeUnit) Attributes() dataAdministrativeUnitAttributes {
	return dataAdministrativeUnitAttributes{ref: terra.ReferenceDataResource(au)}
}

// DataAdministrativeUnitArgs contains the configurations for azuread_administrative_unit.
type DataAdministrativeUnitArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// Timeouts: optional
	Timeouts *dataadministrativeunit.Timeouts `hcl:"timeouts,block"`
}
type dataAdministrativeUnitAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azuread_administrative_unit.
func (au dataAdministrativeUnitAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_administrative_unit.
func (au dataAdministrativeUnitAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("display_name"))
}

// Id returns a reference to field id of azuread_administrative_unit.
func (au dataAdministrativeUnitAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("id"))
}

// Members returns a reference to field members of azuread_administrative_unit.
func (au dataAdministrativeUnitAttributes) Members() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](au.ref.Append("members"))
}

// ObjectId returns a reference to field object_id of azuread_administrative_unit.
func (au dataAdministrativeUnitAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("object_id"))
}

// Visibility returns a reference to field visibility of azuread_administrative_unit.
func (au dataAdministrativeUnitAttributes) Visibility() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("visibility"))
}

func (au dataAdministrativeUnitAttributes) Timeouts() dataadministrativeunit.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataadministrativeunit.TimeoutsAttributes](au.ref.Append("timeouts"))
}

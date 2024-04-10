// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread

import (
	"github.com/golingon/lingon/pkg/terra"
	datanamedlocation "github.com/golingon/terraproviders/azuread/2.47.0/datanamedlocation"
)

// NewDataNamedLocation creates a new instance of [DataNamedLocation].
func NewDataNamedLocation(name string, args DataNamedLocationArgs) *DataNamedLocation {
	return &DataNamedLocation{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNamedLocation)(nil)

// DataNamedLocation represents the Terraform data resource azuread_named_location.
type DataNamedLocation struct {
	Name string
	Args DataNamedLocationArgs
}

// DataSource returns the Terraform object type for [DataNamedLocation].
func (nl *DataNamedLocation) DataSource() string {
	return "azuread_named_location"
}

// LocalName returns the local name for [DataNamedLocation].
func (nl *DataNamedLocation) LocalName() string {
	return nl.Name
}

// Configuration returns the configuration (args) for [DataNamedLocation].
func (nl *DataNamedLocation) Configuration() interface{} {
	return nl.Args
}

// Attributes returns the attributes for [DataNamedLocation].
func (nl *DataNamedLocation) Attributes() dataNamedLocationAttributes {
	return dataNamedLocationAttributes{ref: terra.ReferenceDataResource(nl)}
}

// DataNamedLocationArgs contains the configurations for azuread_named_location.
type DataNamedLocationArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Country: min=0
	Country []datanamedlocation.Country `hcl:"country,block" validate:"min=0"`
	// Ip: min=0
	Ip []datanamedlocation.Ip `hcl:"ip,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanamedlocation.Timeouts `hcl:"timeouts,block"`
}
type dataNamedLocationAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azuread_named_location.
func (nl dataNamedLocationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("display_name"))
}

// Id returns a reference to field id of azuread_named_location.
func (nl dataNamedLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("id"))
}

func (nl dataNamedLocationAttributes) Country() terra.ListValue[datanamedlocation.CountryAttributes] {
	return terra.ReferenceAsList[datanamedlocation.CountryAttributes](nl.ref.Append("country"))
}

func (nl dataNamedLocationAttributes) Ip() terra.ListValue[datanamedlocation.IpAttributes] {
	return terra.ReferenceAsList[datanamedlocation.IpAttributes](nl.ref.Append("ip"))
}

func (nl dataNamedLocationAttributes) Timeouts() datanamedlocation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanamedlocation.TimeoutsAttributes](nl.ref.Append("timeouts"))
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamanagedapi "github.com/golingon/terraproviders/azurerm/3.67.0/datamanagedapi"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataManagedApi creates a new instance of [DataManagedApi].
func NewDataManagedApi(name string, args DataManagedApiArgs) *DataManagedApi {
	return &DataManagedApi{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataManagedApi)(nil)

// DataManagedApi represents the Terraform data resource azurerm_managed_api.
type DataManagedApi struct {
	Name string
	Args DataManagedApiArgs
}

// DataSource returns the Terraform object type for [DataManagedApi].
func (ma *DataManagedApi) DataSource() string {
	return "azurerm_managed_api"
}

// LocalName returns the local name for [DataManagedApi].
func (ma *DataManagedApi) LocalName() string {
	return ma.Name
}

// Configuration returns the configuration (args) for [DataManagedApi].
func (ma *DataManagedApi) Configuration() interface{} {
	return ma.Args
}

// Attributes returns the attributes for [DataManagedApi].
func (ma *DataManagedApi) Attributes() dataManagedApiAttributes {
	return dataManagedApiAttributes{ref: terra.ReferenceDataResource(ma)}
}

// DataManagedApiArgs contains the configurations for azurerm_managed_api.
type DataManagedApiArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datamanagedapi.Timeouts `hcl:"timeouts,block"`
}
type dataManagedApiAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_managed_api.
func (ma dataManagedApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_managed_api.
func (ma dataManagedApiAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_managed_api.
func (ma dataManagedApiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_managed_api.
func (ma dataManagedApiAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ma.ref.Append("tags"))
}

func (ma dataManagedApiAttributes) Timeouts() datamanagedapi.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamanagedapi.TimeoutsAttributes](ma.ref.Append("timeouts"))
}

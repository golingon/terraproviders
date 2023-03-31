// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamanagedapi "github.com/golingon/terraproviders/azurerm/3.49.0/datamanagedapi"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataManagedApi(name string, args DataManagedApiArgs) *DataManagedApi {
	return &DataManagedApi{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataManagedApi)(nil)

type DataManagedApi struct {
	Name string
	Args DataManagedApiArgs
}

func (ma *DataManagedApi) DataSource() string {
	return "azurerm_managed_api"
}

func (ma *DataManagedApi) LocalName() string {
	return ma.Name
}

func (ma *DataManagedApi) Configuration() interface{} {
	return ma.Args
}

func (ma *DataManagedApi) Attributes() dataManagedApiAttributes {
	return dataManagedApiAttributes{ref: terra.ReferenceDataResource(ma)}
}

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

func (ma dataManagedApiAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ma.ref.Append("id"))
}

func (ma dataManagedApiAttributes) Location() terra.StringValue {
	return terra.ReferenceString(ma.ref.Append("location"))
}

func (ma dataManagedApiAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ma.ref.Append("name"))
}

func (ma dataManagedApiAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ma.ref.Append("tags"))
}

func (ma dataManagedApiAttributes) Timeouts() datamanagedapi.TimeoutsAttributes {
	return terra.ReferenceSingle[datamanagedapi.TimeoutsAttributes](ma.ref.Append("timeouts"))
}

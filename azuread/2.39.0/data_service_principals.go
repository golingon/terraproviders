// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	dataserviceprincipals "github.com/golingon/terraproviders/azuread/2.39.0/dataserviceprincipals"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicePrincipals creates a new instance of [DataServicePrincipals].
func NewDataServicePrincipals(name string, args DataServicePrincipalsArgs) *DataServicePrincipals {
	return &DataServicePrincipals{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicePrincipals)(nil)

// DataServicePrincipals represents the Terraform data resource azuread_service_principals.
type DataServicePrincipals struct {
	Name string
	Args DataServicePrincipalsArgs
}

// DataSource returns the Terraform object type for [DataServicePrincipals].
func (sp *DataServicePrincipals) DataSource() string {
	return "azuread_service_principals"
}

// LocalName returns the local name for [DataServicePrincipals].
func (sp *DataServicePrincipals) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [DataServicePrincipals].
func (sp *DataServicePrincipals) Configuration() interface{} {
	return sp.Args
}

// Attributes returns the attributes for [DataServicePrincipals].
func (sp *DataServicePrincipals) Attributes() dataServicePrincipalsAttributes {
	return dataServicePrincipalsAttributes{ref: terra.ReferenceDataResource(sp)}
}

// DataServicePrincipalsArgs contains the configurations for azuread_service_principals.
type DataServicePrincipalsArgs struct {
	// ApplicationIds: list of string, optional
	ApplicationIds terra.ListValue[terra.StringValue] `hcl:"application_ids,attr"`
	// DisplayNames: list of string, optional
	DisplayNames terra.ListValue[terra.StringValue] `hcl:"display_names,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreMissing: bool, optional
	IgnoreMissing terra.BoolValue `hcl:"ignore_missing,attr"`
	// ObjectIds: list of string, optional
	ObjectIds terra.ListValue[terra.StringValue] `hcl:"object_ids,attr"`
	// ReturnAll: bool, optional
	ReturnAll terra.BoolValue `hcl:"return_all,attr"`
	// ServicePrincipals: min=0
	ServicePrincipals []dataserviceprincipals.ServicePrincipals `hcl:"service_principals,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataserviceprincipals.Timeouts `hcl:"timeouts,block"`
}
type dataServicePrincipalsAttributes struct {
	ref terra.Reference
}

// ApplicationIds returns a reference to field application_ids of azuread_service_principals.
func (sp dataServicePrincipalsAttributes) ApplicationIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("application_ids"))
}

// DisplayNames returns a reference to field display_names of azuread_service_principals.
func (sp dataServicePrincipalsAttributes) DisplayNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("display_names"))
}

// Id returns a reference to field id of azuread_service_principals.
func (sp dataServicePrincipalsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// IgnoreMissing returns a reference to field ignore_missing of azuread_service_principals.
func (sp dataServicePrincipalsAttributes) IgnoreMissing() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("ignore_missing"))
}

// ObjectIds returns a reference to field object_ids of azuread_service_principals.
func (sp dataServicePrincipalsAttributes) ObjectIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("object_ids"))
}

// ReturnAll returns a reference to field return_all of azuread_service_principals.
func (sp dataServicePrincipalsAttributes) ReturnAll() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("return_all"))
}

func (sp dataServicePrincipalsAttributes) ServicePrincipals() terra.ListValue[dataserviceprincipals.ServicePrincipalsAttributes] {
	return terra.ReferenceAsList[dataserviceprincipals.ServicePrincipalsAttributes](sp.ref.Append("service_principals"))
}

func (sp dataServicePrincipalsAttributes) Timeouts() dataserviceprincipals.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataserviceprincipals.TimeoutsAttributes](sp.ref.Append("timeouts"))
}

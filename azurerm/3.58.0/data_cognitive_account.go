// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacognitiveaccount "github.com/golingon/terraproviders/azurerm/3.58.0/datacognitiveaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCognitiveAccount creates a new instance of [DataCognitiveAccount].
func NewDataCognitiveAccount(name string, args DataCognitiveAccountArgs) *DataCognitiveAccount {
	return &DataCognitiveAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCognitiveAccount)(nil)

// DataCognitiveAccount represents the Terraform data resource azurerm_cognitive_account.
type DataCognitiveAccount struct {
	Name string
	Args DataCognitiveAccountArgs
}

// DataSource returns the Terraform object type for [DataCognitiveAccount].
func (ca *DataCognitiveAccount) DataSource() string {
	return "azurerm_cognitive_account"
}

// LocalName returns the local name for [DataCognitiveAccount].
func (ca *DataCognitiveAccount) LocalName() string {
	return ca.Name
}

// Configuration returns the configuration (args) for [DataCognitiveAccount].
func (ca *DataCognitiveAccount) Configuration() interface{} {
	return ca.Args
}

// Attributes returns the attributes for [DataCognitiveAccount].
func (ca *DataCognitiveAccount) Attributes() dataCognitiveAccountAttributes {
	return dataCognitiveAccountAttributes{ref: terra.ReferenceDataResource(ca)}
}

// DataCognitiveAccountArgs contains the configurations for azurerm_cognitive_account.
type DataCognitiveAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *datacognitiveaccount.Timeouts `hcl:"timeouts,block"`
}
type dataCognitiveAccountAttributes struct {
	ref terra.Reference
}

// Endpoint returns a reference to field endpoint of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("endpoint"))
}

// Id returns a reference to field id of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("primary_access_key"))
}

// QnaRuntimeEndpoint returns a reference to field qna_runtime_endpoint of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) QnaRuntimeEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("qna_runtime_endpoint"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("secondary_access_key"))
}

// SkuName returns a reference to field sku_name of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_cognitive_account.
func (ca dataCognitiveAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ca.ref.Append("tags"))
}

func (ca dataCognitiveAccountAttributes) Timeouts() datacognitiveaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacognitiveaccount.TimeoutsAttributes](ca.ref.Append("timeouts"))
}

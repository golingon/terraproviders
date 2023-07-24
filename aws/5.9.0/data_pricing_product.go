// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datapricingproduct "github.com/golingon/terraproviders/aws/5.9.0/datapricingproduct"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPricingProduct creates a new instance of [DataPricingProduct].
func NewDataPricingProduct(name string, args DataPricingProductArgs) *DataPricingProduct {
	return &DataPricingProduct{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPricingProduct)(nil)

// DataPricingProduct represents the Terraform data resource aws_pricing_product.
type DataPricingProduct struct {
	Name string
	Args DataPricingProductArgs
}

// DataSource returns the Terraform object type for [DataPricingProduct].
func (pp *DataPricingProduct) DataSource() string {
	return "aws_pricing_product"
}

// LocalName returns the local name for [DataPricingProduct].
func (pp *DataPricingProduct) LocalName() string {
	return pp.Name
}

// Configuration returns the configuration (args) for [DataPricingProduct].
func (pp *DataPricingProduct) Configuration() interface{} {
	return pp.Args
}

// Attributes returns the attributes for [DataPricingProduct].
func (pp *DataPricingProduct) Attributes() dataPricingProductAttributes {
	return dataPricingProductAttributes{ref: terra.ReferenceDataResource(pp)}
}

// DataPricingProductArgs contains the configurations for aws_pricing_product.
type DataPricingProductArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceCode: string, required
	ServiceCode terra.StringValue `hcl:"service_code,attr" validate:"required"`
	// Filters: min=1
	Filters []datapricingproduct.Filters `hcl:"filters,block" validate:"min=1"`
}
type dataPricingProductAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_pricing_product.
func (pp dataPricingProductAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("id"))
}

// Result returns a reference to field result of aws_pricing_product.
func (pp dataPricingProductAttributes) Result() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("result"))
}

// ServiceCode returns a reference to field service_code of aws_pricing_product.
func (pp dataPricingProductAttributes) ServiceCode() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("service_code"))
}

func (pp dataPricingProductAttributes) Filters() terra.ListValue[datapricingproduct.FiltersAttributes] {
	return terra.ReferenceAsList[datapricingproduct.FiltersAttributes](pp.ref.Append("filters"))
}

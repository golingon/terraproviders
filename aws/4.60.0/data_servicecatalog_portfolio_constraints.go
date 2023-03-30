// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataservicecatalogportfolioconstraints "github.com/golingon/terraproviders/aws/4.60.0/dataservicecatalogportfolioconstraints"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataServicecatalogPortfolioConstraints(name string, args DataServicecatalogPortfolioConstraintsArgs) *DataServicecatalogPortfolioConstraints {
	return &DataServicecatalogPortfolioConstraints{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicecatalogPortfolioConstraints)(nil)

type DataServicecatalogPortfolioConstraints struct {
	Name string
	Args DataServicecatalogPortfolioConstraintsArgs
}

func (spc *DataServicecatalogPortfolioConstraints) DataSource() string {
	return "aws_servicecatalog_portfolio_constraints"
}

func (spc *DataServicecatalogPortfolioConstraints) LocalName() string {
	return spc.Name
}

func (spc *DataServicecatalogPortfolioConstraints) Configuration() interface{} {
	return spc.Args
}

func (spc *DataServicecatalogPortfolioConstraints) Attributes() dataServicecatalogPortfolioConstraintsAttributes {
	return dataServicecatalogPortfolioConstraintsAttributes{ref: terra.ReferenceDataResource(spc)}
}

type DataServicecatalogPortfolioConstraintsArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PortfolioId: string, required
	PortfolioId terra.StringValue `hcl:"portfolio_id,attr" validate:"required"`
	// ProductId: string, optional
	ProductId terra.StringValue `hcl:"product_id,attr"`
	// Details: min=0
	Details []dataservicecatalogportfolioconstraints.Details `hcl:"details,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataservicecatalogportfolioconstraints.Timeouts `hcl:"timeouts,block"`
}
type dataServicecatalogPortfolioConstraintsAttributes struct {
	ref terra.Reference
}

func (spc dataServicecatalogPortfolioConstraintsAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceString(spc.ref.Append("accept_language"))
}

func (spc dataServicecatalogPortfolioConstraintsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(spc.ref.Append("id"))
}

func (spc dataServicecatalogPortfolioConstraintsAttributes) PortfolioId() terra.StringValue {
	return terra.ReferenceString(spc.ref.Append("portfolio_id"))
}

func (spc dataServicecatalogPortfolioConstraintsAttributes) ProductId() terra.StringValue {
	return terra.ReferenceString(spc.ref.Append("product_id"))
}

func (spc dataServicecatalogPortfolioConstraintsAttributes) Details() terra.ListValue[dataservicecatalogportfolioconstraints.DetailsAttributes] {
	return terra.ReferenceList[dataservicecatalogportfolioconstraints.DetailsAttributes](spc.ref.Append("details"))
}

func (spc dataServicecatalogPortfolioConstraintsAttributes) Timeouts() dataservicecatalogportfolioconstraints.TimeoutsAttributes {
	return terra.ReferenceSingle[dataservicecatalogportfolioconstraints.TimeoutsAttributes](spc.ref.Append("timeouts"))
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataservicecatalogportfolio "github.com/golingon/terraproviders/aws/4.63.0/dataservicecatalogportfolio"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicecatalogPortfolio creates a new instance of [DataServicecatalogPortfolio].
func NewDataServicecatalogPortfolio(name string, args DataServicecatalogPortfolioArgs) *DataServicecatalogPortfolio {
	return &DataServicecatalogPortfolio{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicecatalogPortfolio)(nil)

// DataServicecatalogPortfolio represents the Terraform data resource aws_servicecatalog_portfolio.
type DataServicecatalogPortfolio struct {
	Name string
	Args DataServicecatalogPortfolioArgs
}

// DataSource returns the Terraform object type for [DataServicecatalogPortfolio].
func (sp *DataServicecatalogPortfolio) DataSource() string {
	return "aws_servicecatalog_portfolio"
}

// LocalName returns the local name for [DataServicecatalogPortfolio].
func (sp *DataServicecatalogPortfolio) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [DataServicecatalogPortfolio].
func (sp *DataServicecatalogPortfolio) Configuration() interface{} {
	return sp.Args
}

// Attributes returns the attributes for [DataServicecatalogPortfolio].
func (sp *DataServicecatalogPortfolio) Attributes() dataServicecatalogPortfolioAttributes {
	return dataServicecatalogPortfolioAttributes{ref: terra.ReferenceDataResource(sp)}
}

// DataServicecatalogPortfolioArgs contains the configurations for aws_servicecatalog_portfolio.
type DataServicecatalogPortfolioArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *dataservicecatalogportfolio.Timeouts `hcl:"timeouts,block"`
}
type dataServicecatalogPortfolioAttributes struct {
	ref terra.Reference
}

// AcceptLanguage returns a reference to field accept_language of aws_servicecatalog_portfolio.
func (sp dataServicecatalogPortfolioAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("accept_language"))
}

// Arn returns a reference to field arn of aws_servicecatalog_portfolio.
func (sp dataServicecatalogPortfolioAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_servicecatalog_portfolio.
func (sp dataServicecatalogPortfolioAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_servicecatalog_portfolio.
func (sp dataServicecatalogPortfolioAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("description"))
}

// Id returns a reference to field id of aws_servicecatalog_portfolio.
func (sp dataServicecatalogPortfolioAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// Name returns a reference to field name of aws_servicecatalog_portfolio.
func (sp dataServicecatalogPortfolioAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("name"))
}

// ProviderName returns a reference to field provider_name of aws_servicecatalog_portfolio.
func (sp dataServicecatalogPortfolioAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("provider_name"))
}

// Tags returns a reference to field tags of aws_servicecatalog_portfolio.
func (sp dataServicecatalogPortfolioAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags"))
}

func (sp dataServicecatalogPortfolioAttributes) Timeouts() dataservicecatalogportfolio.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataservicecatalogportfolio.TimeoutsAttributes](sp.ref.Append("timeouts"))
}

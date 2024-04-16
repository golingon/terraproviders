// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_servicecatalog_portfolio_constraints

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_servicecatalog_portfolio_constraints.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aspc *DataSource) DataSource() string {
	return "aws_servicecatalog_portfolio_constraints"
}

// LocalName returns the local name for [DataSource].
func (aspc *DataSource) LocalName() string {
	return aspc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aspc *DataSource) Configuration() interface{} {
	return aspc.Args
}

// Attributes returns the attributes for [DataSource].
func (aspc *DataSource) Attributes() dataAwsServicecatalogPortfolioConstraintsAttributes {
	return dataAwsServicecatalogPortfolioConstraintsAttributes{ref: terra.ReferenceDataSource(aspc)}
}

// DataArgs contains the configurations for aws_servicecatalog_portfolio_constraints.
type DataArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PortfolioId: string, required
	PortfolioId terra.StringValue `hcl:"portfolio_id,attr" validate:"required"`
	// ProductId: string, optional
	ProductId terra.StringValue `hcl:"product_id,attr"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAwsServicecatalogPortfolioConstraintsAttributes struct {
	ref terra.Reference
}

// AcceptLanguage returns a reference to field accept_language of aws_servicecatalog_portfolio_constraints.
func (aspc dataAwsServicecatalogPortfolioConstraintsAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceAsString(aspc.ref.Append("accept_language"))
}

// Id returns a reference to field id of aws_servicecatalog_portfolio_constraints.
func (aspc dataAwsServicecatalogPortfolioConstraintsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aspc.ref.Append("id"))
}

// PortfolioId returns a reference to field portfolio_id of aws_servicecatalog_portfolio_constraints.
func (aspc dataAwsServicecatalogPortfolioConstraintsAttributes) PortfolioId() terra.StringValue {
	return terra.ReferenceAsString(aspc.ref.Append("portfolio_id"))
}

// ProductId returns a reference to field product_id of aws_servicecatalog_portfolio_constraints.
func (aspc dataAwsServicecatalogPortfolioConstraintsAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(aspc.ref.Append("product_id"))
}

func (aspc dataAwsServicecatalogPortfolioConstraintsAttributes) Details() terra.ListValue[DataDetailsAttributes] {
	return terra.ReferenceAsList[DataDetailsAttributes](aspc.ref.Append("details"))
}

func (aspc dataAwsServicecatalogPortfolioConstraintsAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](aspc.ref.Append("timeouts"))
}

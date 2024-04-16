// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_servicecatalog_portfolio

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_servicecatalog_portfolio.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (asp *DataSource) DataSource() string {
	return "aws_servicecatalog_portfolio"
}

// LocalName returns the local name for [DataSource].
func (asp *DataSource) LocalName() string {
	return asp.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (asp *DataSource) Configuration() interface{} {
	return asp.Args
}

// Attributes returns the attributes for [DataSource].
func (asp *DataSource) Attributes() dataAwsServicecatalogPortfolioAttributes {
	return dataAwsServicecatalogPortfolioAttributes{ref: terra.ReferenceDataSource(asp)}
}

// DataArgs contains the configurations for aws_servicecatalog_portfolio.
type DataArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAwsServicecatalogPortfolioAttributes struct {
	ref terra.Reference
}

// AcceptLanguage returns a reference to field accept_language of aws_servicecatalog_portfolio.
func (asp dataAwsServicecatalogPortfolioAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("accept_language"))
}

// Arn returns a reference to field arn of aws_servicecatalog_portfolio.
func (asp dataAwsServicecatalogPortfolioAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_servicecatalog_portfolio.
func (asp dataAwsServicecatalogPortfolioAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_servicecatalog_portfolio.
func (asp dataAwsServicecatalogPortfolioAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("description"))
}

// Id returns a reference to field id of aws_servicecatalog_portfolio.
func (asp dataAwsServicecatalogPortfolioAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("id"))
}

// Name returns a reference to field name of aws_servicecatalog_portfolio.
func (asp dataAwsServicecatalogPortfolioAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("name"))
}

// ProviderName returns a reference to field provider_name of aws_servicecatalog_portfolio.
func (asp dataAwsServicecatalogPortfolioAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("provider_name"))
}

// Tags returns a reference to field tags of aws_servicecatalog_portfolio.
func (asp dataAwsServicecatalogPortfolioAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asp.ref.Append("tags"))
}

func (asp dataAwsServicecatalogPortfolioAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](asp.ref.Append("timeouts"))
}

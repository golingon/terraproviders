// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogportfolio "github.com/golingon/terraproviders/aws/5.0.1/servicecatalogportfolio"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicecatalogPortfolio creates a new instance of [ServicecatalogPortfolio].
func NewServicecatalogPortfolio(name string, args ServicecatalogPortfolioArgs) *ServicecatalogPortfolio {
	return &ServicecatalogPortfolio{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogPortfolio)(nil)

// ServicecatalogPortfolio represents the Terraform resource aws_servicecatalog_portfolio.
type ServicecatalogPortfolio struct {
	Name      string
	Args      ServicecatalogPortfolioArgs
	state     *servicecatalogPortfolioState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicecatalogPortfolio].
func (sp *ServicecatalogPortfolio) Type() string {
	return "aws_servicecatalog_portfolio"
}

// LocalName returns the local name for [ServicecatalogPortfolio].
func (sp *ServicecatalogPortfolio) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [ServicecatalogPortfolio].
func (sp *ServicecatalogPortfolio) Configuration() interface{} {
	return sp.Args
}

// DependOn is used for other resources to depend on [ServicecatalogPortfolio].
func (sp *ServicecatalogPortfolio) DependOn() terra.Reference {
	return terra.ReferenceResource(sp)
}

// Dependencies returns the list of resources [ServicecatalogPortfolio] depends_on.
func (sp *ServicecatalogPortfolio) Dependencies() terra.Dependencies {
	return sp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicecatalogPortfolio].
func (sp *ServicecatalogPortfolio) LifecycleManagement() *terra.Lifecycle {
	return sp.Lifecycle
}

// Attributes returns the attributes for [ServicecatalogPortfolio].
func (sp *ServicecatalogPortfolio) Attributes() servicecatalogPortfolioAttributes {
	return servicecatalogPortfolioAttributes{ref: terra.ReferenceResource(sp)}
}

// ImportState imports the given attribute values into [ServicecatalogPortfolio]'s state.
func (sp *ServicecatalogPortfolio) ImportState(av io.Reader) error {
	sp.state = &servicecatalogPortfolioState{}
	if err := json.NewDecoder(av).Decode(sp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sp.Type(), sp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicecatalogPortfolio] has state.
func (sp *ServicecatalogPortfolio) State() (*servicecatalogPortfolioState, bool) {
	return sp.state, sp.state != nil
}

// StateMust returns the state for [ServicecatalogPortfolio]. Panics if the state is nil.
func (sp *ServicecatalogPortfolio) StateMust() *servicecatalogPortfolioState {
	if sp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sp.Type(), sp.LocalName()))
	}
	return sp.state
}

// ServicecatalogPortfolioArgs contains the configurations for aws_servicecatalog_portfolio.
type ServicecatalogPortfolioArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProviderName: string, required
	ProviderName terra.StringValue `hcl:"provider_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *servicecatalogportfolio.Timeouts `hcl:"timeouts,block"`
}
type servicecatalogPortfolioAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_servicecatalog_portfolio.
func (sp servicecatalogPortfolioAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_servicecatalog_portfolio.
func (sp servicecatalogPortfolioAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_servicecatalog_portfolio.
func (sp servicecatalogPortfolioAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("description"))
}

// Id returns a reference to field id of aws_servicecatalog_portfolio.
func (sp servicecatalogPortfolioAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// Name returns a reference to field name of aws_servicecatalog_portfolio.
func (sp servicecatalogPortfolioAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("name"))
}

// ProviderName returns a reference to field provider_name of aws_servicecatalog_portfolio.
func (sp servicecatalogPortfolioAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("provider_name"))
}

// Tags returns a reference to field tags of aws_servicecatalog_portfolio.
func (sp servicecatalogPortfolioAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_servicecatalog_portfolio.
func (sp servicecatalogPortfolioAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags_all"))
}

func (sp servicecatalogPortfolioAttributes) Timeouts() servicecatalogportfolio.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicecatalogportfolio.TimeoutsAttributes](sp.ref.Append("timeouts"))
}

type servicecatalogPortfolioState struct {
	Arn          string                                 `json:"arn"`
	CreatedTime  string                                 `json:"created_time"`
	Description  string                                 `json:"description"`
	Id           string                                 `json:"id"`
	Name         string                                 `json:"name"`
	ProviderName string                                 `json:"provider_name"`
	Tags         map[string]string                      `json:"tags"`
	TagsAll      map[string]string                      `json:"tags_all"`
	Timeouts     *servicecatalogportfolio.TimeoutsState `json:"timeouts"`
}
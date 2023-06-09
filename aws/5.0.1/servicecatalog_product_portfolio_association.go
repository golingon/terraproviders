// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogproductportfolioassociation "github.com/golingon/terraproviders/aws/5.0.1/servicecatalogproductportfolioassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicecatalogProductPortfolioAssociation creates a new instance of [ServicecatalogProductPortfolioAssociation].
func NewServicecatalogProductPortfolioAssociation(name string, args ServicecatalogProductPortfolioAssociationArgs) *ServicecatalogProductPortfolioAssociation {
	return &ServicecatalogProductPortfolioAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogProductPortfolioAssociation)(nil)

// ServicecatalogProductPortfolioAssociation represents the Terraform resource aws_servicecatalog_product_portfolio_association.
type ServicecatalogProductPortfolioAssociation struct {
	Name      string
	Args      ServicecatalogProductPortfolioAssociationArgs
	state     *servicecatalogProductPortfolioAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicecatalogProductPortfolioAssociation].
func (sppa *ServicecatalogProductPortfolioAssociation) Type() string {
	return "aws_servicecatalog_product_portfolio_association"
}

// LocalName returns the local name for [ServicecatalogProductPortfolioAssociation].
func (sppa *ServicecatalogProductPortfolioAssociation) LocalName() string {
	return sppa.Name
}

// Configuration returns the configuration (args) for [ServicecatalogProductPortfolioAssociation].
func (sppa *ServicecatalogProductPortfolioAssociation) Configuration() interface{} {
	return sppa.Args
}

// DependOn is used for other resources to depend on [ServicecatalogProductPortfolioAssociation].
func (sppa *ServicecatalogProductPortfolioAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(sppa)
}

// Dependencies returns the list of resources [ServicecatalogProductPortfolioAssociation] depends_on.
func (sppa *ServicecatalogProductPortfolioAssociation) Dependencies() terra.Dependencies {
	return sppa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicecatalogProductPortfolioAssociation].
func (sppa *ServicecatalogProductPortfolioAssociation) LifecycleManagement() *terra.Lifecycle {
	return sppa.Lifecycle
}

// Attributes returns the attributes for [ServicecatalogProductPortfolioAssociation].
func (sppa *ServicecatalogProductPortfolioAssociation) Attributes() servicecatalogProductPortfolioAssociationAttributes {
	return servicecatalogProductPortfolioAssociationAttributes{ref: terra.ReferenceResource(sppa)}
}

// ImportState imports the given attribute values into [ServicecatalogProductPortfolioAssociation]'s state.
func (sppa *ServicecatalogProductPortfolioAssociation) ImportState(av io.Reader) error {
	sppa.state = &servicecatalogProductPortfolioAssociationState{}
	if err := json.NewDecoder(av).Decode(sppa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sppa.Type(), sppa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicecatalogProductPortfolioAssociation] has state.
func (sppa *ServicecatalogProductPortfolioAssociation) State() (*servicecatalogProductPortfolioAssociationState, bool) {
	return sppa.state, sppa.state != nil
}

// StateMust returns the state for [ServicecatalogProductPortfolioAssociation]. Panics if the state is nil.
func (sppa *ServicecatalogProductPortfolioAssociation) StateMust() *servicecatalogProductPortfolioAssociationState {
	if sppa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sppa.Type(), sppa.LocalName()))
	}
	return sppa.state
}

// ServicecatalogProductPortfolioAssociationArgs contains the configurations for aws_servicecatalog_product_portfolio_association.
type ServicecatalogProductPortfolioAssociationArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PortfolioId: string, required
	PortfolioId terra.StringValue `hcl:"portfolio_id,attr" validate:"required"`
	// ProductId: string, required
	ProductId terra.StringValue `hcl:"product_id,attr" validate:"required"`
	// SourcePortfolioId: string, optional
	SourcePortfolioId terra.StringValue `hcl:"source_portfolio_id,attr"`
	// Timeouts: optional
	Timeouts *servicecatalogproductportfolioassociation.Timeouts `hcl:"timeouts,block"`
}
type servicecatalogProductPortfolioAssociationAttributes struct {
	ref terra.Reference
}

// AcceptLanguage returns a reference to field accept_language of aws_servicecatalog_product_portfolio_association.
func (sppa servicecatalogProductPortfolioAssociationAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceAsString(sppa.ref.Append("accept_language"))
}

// Id returns a reference to field id of aws_servicecatalog_product_portfolio_association.
func (sppa servicecatalogProductPortfolioAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sppa.ref.Append("id"))
}

// PortfolioId returns a reference to field portfolio_id of aws_servicecatalog_product_portfolio_association.
func (sppa servicecatalogProductPortfolioAssociationAttributes) PortfolioId() terra.StringValue {
	return terra.ReferenceAsString(sppa.ref.Append("portfolio_id"))
}

// ProductId returns a reference to field product_id of aws_servicecatalog_product_portfolio_association.
func (sppa servicecatalogProductPortfolioAssociationAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(sppa.ref.Append("product_id"))
}

// SourcePortfolioId returns a reference to field source_portfolio_id of aws_servicecatalog_product_portfolio_association.
func (sppa servicecatalogProductPortfolioAssociationAttributes) SourcePortfolioId() terra.StringValue {
	return terra.ReferenceAsString(sppa.ref.Append("source_portfolio_id"))
}

func (sppa servicecatalogProductPortfolioAssociationAttributes) Timeouts() servicecatalogproductportfolioassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicecatalogproductportfolioassociation.TimeoutsAttributes](sppa.ref.Append("timeouts"))
}

type servicecatalogProductPortfolioAssociationState struct {
	AcceptLanguage    string                                                   `json:"accept_language"`
	Id                string                                                   `json:"id"`
	PortfolioId       string                                                   `json:"portfolio_id"`
	ProductId         string                                                   `json:"product_id"`
	SourcePortfolioId string                                                   `json:"source_portfolio_id"`
	Timeouts          *servicecatalogproductportfolioassociation.TimeoutsState `json:"timeouts"`
}

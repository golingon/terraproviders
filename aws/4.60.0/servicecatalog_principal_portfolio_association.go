// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogprincipalportfolioassociation "github.com/golingon/terraproviders/aws/4.60.0/servicecatalogprincipalportfolioassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewServicecatalogPrincipalPortfolioAssociation(name string, args ServicecatalogPrincipalPortfolioAssociationArgs) *ServicecatalogPrincipalPortfolioAssociation {
	return &ServicecatalogPrincipalPortfolioAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogPrincipalPortfolioAssociation)(nil)

type ServicecatalogPrincipalPortfolioAssociation struct {
	Name  string
	Args  ServicecatalogPrincipalPortfolioAssociationArgs
	state *servicecatalogPrincipalPortfolioAssociationState
}

func (sppa *ServicecatalogPrincipalPortfolioAssociation) Type() string {
	return "aws_servicecatalog_principal_portfolio_association"
}

func (sppa *ServicecatalogPrincipalPortfolioAssociation) LocalName() string {
	return sppa.Name
}

func (sppa *ServicecatalogPrincipalPortfolioAssociation) Configuration() interface{} {
	return sppa.Args
}

func (sppa *ServicecatalogPrincipalPortfolioAssociation) Attributes() servicecatalogPrincipalPortfolioAssociationAttributes {
	return servicecatalogPrincipalPortfolioAssociationAttributes{ref: terra.ReferenceResource(sppa)}
}

func (sppa *ServicecatalogPrincipalPortfolioAssociation) ImportState(av io.Reader) error {
	sppa.state = &servicecatalogPrincipalPortfolioAssociationState{}
	if err := json.NewDecoder(av).Decode(sppa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sppa.Type(), sppa.LocalName(), err)
	}
	return nil
}

func (sppa *ServicecatalogPrincipalPortfolioAssociation) State() (*servicecatalogPrincipalPortfolioAssociationState, bool) {
	return sppa.state, sppa.state != nil
}

func (sppa *ServicecatalogPrincipalPortfolioAssociation) StateMust() *servicecatalogPrincipalPortfolioAssociationState {
	if sppa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sppa.Type(), sppa.LocalName()))
	}
	return sppa.state
}

func (sppa *ServicecatalogPrincipalPortfolioAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(sppa)
}

type ServicecatalogPrincipalPortfolioAssociationArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PortfolioId: string, required
	PortfolioId terra.StringValue `hcl:"portfolio_id,attr" validate:"required"`
	// PrincipalArn: string, required
	PrincipalArn terra.StringValue `hcl:"principal_arn,attr" validate:"required"`
	// PrincipalType: string, optional
	PrincipalType terra.StringValue `hcl:"principal_type,attr"`
	// Timeouts: optional
	Timeouts *servicecatalogprincipalportfolioassociation.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ServicecatalogPrincipalPortfolioAssociation depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type servicecatalogPrincipalPortfolioAssociationAttributes struct {
	ref terra.Reference
}

func (sppa servicecatalogPrincipalPortfolioAssociationAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceString(sppa.ref.Append("accept_language"))
}

func (sppa servicecatalogPrincipalPortfolioAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sppa.ref.Append("id"))
}

func (sppa servicecatalogPrincipalPortfolioAssociationAttributes) PortfolioId() terra.StringValue {
	return terra.ReferenceString(sppa.ref.Append("portfolio_id"))
}

func (sppa servicecatalogPrincipalPortfolioAssociationAttributes) PrincipalArn() terra.StringValue {
	return terra.ReferenceString(sppa.ref.Append("principal_arn"))
}

func (sppa servicecatalogPrincipalPortfolioAssociationAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceString(sppa.ref.Append("principal_type"))
}

func (sppa servicecatalogPrincipalPortfolioAssociationAttributes) Timeouts() servicecatalogprincipalportfolioassociation.TimeoutsAttributes {
	return terra.ReferenceSingle[servicecatalogprincipalportfolioassociation.TimeoutsAttributes](sppa.ref.Append("timeouts"))
}

type servicecatalogPrincipalPortfolioAssociationState struct {
	AcceptLanguage string                                                     `json:"accept_language"`
	Id             string                                                     `json:"id"`
	PortfolioId    string                                                     `json:"portfolio_id"`
	PrincipalArn   string                                                     `json:"principal_arn"`
	PrincipalType  string                                                     `json:"principal_type"`
	Timeouts       *servicecatalogprincipalportfolioassociation.TimeoutsState `json:"timeouts"`
}

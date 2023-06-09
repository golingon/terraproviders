// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogportfolioshare "github.com/golingon/terraproviders/aws/4.63.0/servicecatalogportfolioshare"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicecatalogPortfolioShare creates a new instance of [ServicecatalogPortfolioShare].
func NewServicecatalogPortfolioShare(name string, args ServicecatalogPortfolioShareArgs) *ServicecatalogPortfolioShare {
	return &ServicecatalogPortfolioShare{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogPortfolioShare)(nil)

// ServicecatalogPortfolioShare represents the Terraform resource aws_servicecatalog_portfolio_share.
type ServicecatalogPortfolioShare struct {
	Name      string
	Args      ServicecatalogPortfolioShareArgs
	state     *servicecatalogPortfolioShareState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicecatalogPortfolioShare].
func (sps *ServicecatalogPortfolioShare) Type() string {
	return "aws_servicecatalog_portfolio_share"
}

// LocalName returns the local name for [ServicecatalogPortfolioShare].
func (sps *ServicecatalogPortfolioShare) LocalName() string {
	return sps.Name
}

// Configuration returns the configuration (args) for [ServicecatalogPortfolioShare].
func (sps *ServicecatalogPortfolioShare) Configuration() interface{} {
	return sps.Args
}

// DependOn is used for other resources to depend on [ServicecatalogPortfolioShare].
func (sps *ServicecatalogPortfolioShare) DependOn() terra.Reference {
	return terra.ReferenceResource(sps)
}

// Dependencies returns the list of resources [ServicecatalogPortfolioShare] depends_on.
func (sps *ServicecatalogPortfolioShare) Dependencies() terra.Dependencies {
	return sps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicecatalogPortfolioShare].
func (sps *ServicecatalogPortfolioShare) LifecycleManagement() *terra.Lifecycle {
	return sps.Lifecycle
}

// Attributes returns the attributes for [ServicecatalogPortfolioShare].
func (sps *ServicecatalogPortfolioShare) Attributes() servicecatalogPortfolioShareAttributes {
	return servicecatalogPortfolioShareAttributes{ref: terra.ReferenceResource(sps)}
}

// ImportState imports the given attribute values into [ServicecatalogPortfolioShare]'s state.
func (sps *ServicecatalogPortfolioShare) ImportState(av io.Reader) error {
	sps.state = &servicecatalogPortfolioShareState{}
	if err := json.NewDecoder(av).Decode(sps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sps.Type(), sps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicecatalogPortfolioShare] has state.
func (sps *ServicecatalogPortfolioShare) State() (*servicecatalogPortfolioShareState, bool) {
	return sps.state, sps.state != nil
}

// StateMust returns the state for [ServicecatalogPortfolioShare]. Panics if the state is nil.
func (sps *ServicecatalogPortfolioShare) StateMust() *servicecatalogPortfolioShareState {
	if sps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sps.Type(), sps.LocalName()))
	}
	return sps.state
}

// ServicecatalogPortfolioShareArgs contains the configurations for aws_servicecatalog_portfolio_share.
type ServicecatalogPortfolioShareArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PortfolioId: string, required
	PortfolioId terra.StringValue `hcl:"portfolio_id,attr" validate:"required"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// SharePrincipals: bool, optional
	SharePrincipals terra.BoolValue `hcl:"share_principals,attr"`
	// ShareTagOptions: bool, optional
	ShareTagOptions terra.BoolValue `hcl:"share_tag_options,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// WaitForAcceptance: bool, optional
	WaitForAcceptance terra.BoolValue `hcl:"wait_for_acceptance,attr"`
	// Timeouts: optional
	Timeouts *servicecatalogportfolioshare.Timeouts `hcl:"timeouts,block"`
}
type servicecatalogPortfolioShareAttributes struct {
	ref terra.Reference
}

// AcceptLanguage returns a reference to field accept_language of aws_servicecatalog_portfolio_share.
func (sps servicecatalogPortfolioShareAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("accept_language"))
}

// Accepted returns a reference to field accepted of aws_servicecatalog_portfolio_share.
func (sps servicecatalogPortfolioShareAttributes) Accepted() terra.BoolValue {
	return terra.ReferenceAsBool(sps.ref.Append("accepted"))
}

// Id returns a reference to field id of aws_servicecatalog_portfolio_share.
func (sps servicecatalogPortfolioShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("id"))
}

// PortfolioId returns a reference to field portfolio_id of aws_servicecatalog_portfolio_share.
func (sps servicecatalogPortfolioShareAttributes) PortfolioId() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("portfolio_id"))
}

// PrincipalId returns a reference to field principal_id of aws_servicecatalog_portfolio_share.
func (sps servicecatalogPortfolioShareAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("principal_id"))
}

// SharePrincipals returns a reference to field share_principals of aws_servicecatalog_portfolio_share.
func (sps servicecatalogPortfolioShareAttributes) SharePrincipals() terra.BoolValue {
	return terra.ReferenceAsBool(sps.ref.Append("share_principals"))
}

// ShareTagOptions returns a reference to field share_tag_options of aws_servicecatalog_portfolio_share.
func (sps servicecatalogPortfolioShareAttributes) ShareTagOptions() terra.BoolValue {
	return terra.ReferenceAsBool(sps.ref.Append("share_tag_options"))
}

// Type returns a reference to field type of aws_servicecatalog_portfolio_share.
func (sps servicecatalogPortfolioShareAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("type"))
}

// WaitForAcceptance returns a reference to field wait_for_acceptance of aws_servicecatalog_portfolio_share.
func (sps servicecatalogPortfolioShareAttributes) WaitForAcceptance() terra.BoolValue {
	return terra.ReferenceAsBool(sps.ref.Append("wait_for_acceptance"))
}

func (sps servicecatalogPortfolioShareAttributes) Timeouts() servicecatalogportfolioshare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicecatalogportfolioshare.TimeoutsAttributes](sps.ref.Append("timeouts"))
}

type servicecatalogPortfolioShareState struct {
	AcceptLanguage    string                                      `json:"accept_language"`
	Accepted          bool                                        `json:"accepted"`
	Id                string                                      `json:"id"`
	PortfolioId       string                                      `json:"portfolio_id"`
	PrincipalId       string                                      `json:"principal_id"`
	SharePrincipals   bool                                        `json:"share_principals"`
	ShareTagOptions   bool                                        `json:"share_tag_options"`
	Type              string                                      `json:"type"`
	WaitForAcceptance bool                                        `json:"wait_for_acceptance"`
	Timeouts          *servicecatalogportfolioshare.TimeoutsState `json:"timeouts"`
}

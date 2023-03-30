// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogconstraint "github.com/golingon/terraproviders/aws/4.60.0/servicecatalogconstraint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewServicecatalogConstraint(name string, args ServicecatalogConstraintArgs) *ServicecatalogConstraint {
	return &ServicecatalogConstraint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogConstraint)(nil)

type ServicecatalogConstraint struct {
	Name  string
	Args  ServicecatalogConstraintArgs
	state *servicecatalogConstraintState
}

func (sc *ServicecatalogConstraint) Type() string {
	return "aws_servicecatalog_constraint"
}

func (sc *ServicecatalogConstraint) LocalName() string {
	return sc.Name
}

func (sc *ServicecatalogConstraint) Configuration() interface{} {
	return sc.Args
}

func (sc *ServicecatalogConstraint) Attributes() servicecatalogConstraintAttributes {
	return servicecatalogConstraintAttributes{ref: terra.ReferenceResource(sc)}
}

func (sc *ServicecatalogConstraint) ImportState(av io.Reader) error {
	sc.state = &servicecatalogConstraintState{}
	if err := json.NewDecoder(av).Decode(sc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sc.Type(), sc.LocalName(), err)
	}
	return nil
}

func (sc *ServicecatalogConstraint) State() (*servicecatalogConstraintState, bool) {
	return sc.state, sc.state != nil
}

func (sc *ServicecatalogConstraint) StateMust() *servicecatalogConstraintState {
	if sc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sc.Type(), sc.LocalName()))
	}
	return sc.state
}

func (sc *ServicecatalogConstraint) DependOn() terra.Reference {
	return terra.ReferenceResource(sc)
}

type ServicecatalogConstraintArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parameters: string, required
	Parameters terra.StringValue `hcl:"parameters,attr" validate:"required"`
	// PortfolioId: string, required
	PortfolioId terra.StringValue `hcl:"portfolio_id,attr" validate:"required"`
	// ProductId: string, required
	ProductId terra.StringValue `hcl:"product_id,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *servicecatalogconstraint.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ServicecatalogConstraint depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type servicecatalogConstraintAttributes struct {
	ref terra.Reference
}

func (sc servicecatalogConstraintAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("accept_language"))
}

func (sc servicecatalogConstraintAttributes) Description() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("description"))
}

func (sc servicecatalogConstraintAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("id"))
}

func (sc servicecatalogConstraintAttributes) Owner() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("owner"))
}

func (sc servicecatalogConstraintAttributes) Parameters() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("parameters"))
}

func (sc servicecatalogConstraintAttributes) PortfolioId() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("portfolio_id"))
}

func (sc servicecatalogConstraintAttributes) ProductId() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("product_id"))
}

func (sc servicecatalogConstraintAttributes) Status() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("status"))
}

func (sc servicecatalogConstraintAttributes) Type() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("type"))
}

func (sc servicecatalogConstraintAttributes) Timeouts() servicecatalogconstraint.TimeoutsAttributes {
	return terra.ReferenceSingle[servicecatalogconstraint.TimeoutsAttributes](sc.ref.Append("timeouts"))
}

type servicecatalogConstraintState struct {
	AcceptLanguage string                                  `json:"accept_language"`
	Description    string                                  `json:"description"`
	Id             string                                  `json:"id"`
	Owner          string                                  `json:"owner"`
	Parameters     string                                  `json:"parameters"`
	PortfolioId    string                                  `json:"portfolio_id"`
	ProductId      string                                  `json:"product_id"`
	Status         string                                  `json:"status"`
	Type           string                                  `json:"type"`
	Timeouts       *servicecatalogconstraint.TimeoutsState `json:"timeouts"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogbudgetresourceassociation "github.com/golingon/terraproviders/aws/4.60.0/servicecatalogbudgetresourceassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicecatalogBudgetResourceAssociation creates a new instance of [ServicecatalogBudgetResourceAssociation].
func NewServicecatalogBudgetResourceAssociation(name string, args ServicecatalogBudgetResourceAssociationArgs) *ServicecatalogBudgetResourceAssociation {
	return &ServicecatalogBudgetResourceAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogBudgetResourceAssociation)(nil)

// ServicecatalogBudgetResourceAssociation represents the Terraform resource aws_servicecatalog_budget_resource_association.
type ServicecatalogBudgetResourceAssociation struct {
	Name      string
	Args      ServicecatalogBudgetResourceAssociationArgs
	state     *servicecatalogBudgetResourceAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicecatalogBudgetResourceAssociation].
func (sbra *ServicecatalogBudgetResourceAssociation) Type() string {
	return "aws_servicecatalog_budget_resource_association"
}

// LocalName returns the local name for [ServicecatalogBudgetResourceAssociation].
func (sbra *ServicecatalogBudgetResourceAssociation) LocalName() string {
	return sbra.Name
}

// Configuration returns the configuration (args) for [ServicecatalogBudgetResourceAssociation].
func (sbra *ServicecatalogBudgetResourceAssociation) Configuration() interface{} {
	return sbra.Args
}

// DependOn is used for other resources to depend on [ServicecatalogBudgetResourceAssociation].
func (sbra *ServicecatalogBudgetResourceAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(sbra)
}

// Dependencies returns the list of resources [ServicecatalogBudgetResourceAssociation] depends_on.
func (sbra *ServicecatalogBudgetResourceAssociation) Dependencies() terra.Dependencies {
	return sbra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicecatalogBudgetResourceAssociation].
func (sbra *ServicecatalogBudgetResourceAssociation) LifecycleManagement() *terra.Lifecycle {
	return sbra.Lifecycle
}

// Attributes returns the attributes for [ServicecatalogBudgetResourceAssociation].
func (sbra *ServicecatalogBudgetResourceAssociation) Attributes() servicecatalogBudgetResourceAssociationAttributes {
	return servicecatalogBudgetResourceAssociationAttributes{ref: terra.ReferenceResource(sbra)}
}

// ImportState imports the given attribute values into [ServicecatalogBudgetResourceAssociation]'s state.
func (sbra *ServicecatalogBudgetResourceAssociation) ImportState(av io.Reader) error {
	sbra.state = &servicecatalogBudgetResourceAssociationState{}
	if err := json.NewDecoder(av).Decode(sbra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbra.Type(), sbra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicecatalogBudgetResourceAssociation] has state.
func (sbra *ServicecatalogBudgetResourceAssociation) State() (*servicecatalogBudgetResourceAssociationState, bool) {
	return sbra.state, sbra.state != nil
}

// StateMust returns the state for [ServicecatalogBudgetResourceAssociation]. Panics if the state is nil.
func (sbra *ServicecatalogBudgetResourceAssociation) StateMust() *servicecatalogBudgetResourceAssociationState {
	if sbra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbra.Type(), sbra.LocalName()))
	}
	return sbra.state
}

// ServicecatalogBudgetResourceAssociationArgs contains the configurations for aws_servicecatalog_budget_resource_association.
type ServicecatalogBudgetResourceAssociationArgs struct {
	// BudgetName: string, required
	BudgetName terra.StringValue `hcl:"budget_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *servicecatalogbudgetresourceassociation.Timeouts `hcl:"timeouts,block"`
}
type servicecatalogBudgetResourceAssociationAttributes struct {
	ref terra.Reference
}

// BudgetName returns a reference to field budget_name of aws_servicecatalog_budget_resource_association.
func (sbra servicecatalogBudgetResourceAssociationAttributes) BudgetName() terra.StringValue {
	return terra.ReferenceAsString(sbra.ref.Append("budget_name"))
}

// Id returns a reference to field id of aws_servicecatalog_budget_resource_association.
func (sbra servicecatalogBudgetResourceAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbra.ref.Append("id"))
}

// ResourceId returns a reference to field resource_id of aws_servicecatalog_budget_resource_association.
func (sbra servicecatalogBudgetResourceAssociationAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(sbra.ref.Append("resource_id"))
}

func (sbra servicecatalogBudgetResourceAssociationAttributes) Timeouts() servicecatalogbudgetresourceassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicecatalogbudgetresourceassociation.TimeoutsAttributes](sbra.ref.Append("timeouts"))
}

type servicecatalogBudgetResourceAssociationState struct {
	BudgetName string                                                 `json:"budget_name"`
	Id         string                                                 `json:"id"`
	ResourceId string                                                 `json:"resource_id"`
	Timeouts   *servicecatalogbudgetresourceassociation.TimeoutsState `json:"timeouts"`
}

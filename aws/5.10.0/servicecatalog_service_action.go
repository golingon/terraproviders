// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogserviceaction "github.com/golingon/terraproviders/aws/5.10.0/servicecatalogserviceaction"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicecatalogServiceAction creates a new instance of [ServicecatalogServiceAction].
func NewServicecatalogServiceAction(name string, args ServicecatalogServiceActionArgs) *ServicecatalogServiceAction {
	return &ServicecatalogServiceAction{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogServiceAction)(nil)

// ServicecatalogServiceAction represents the Terraform resource aws_servicecatalog_service_action.
type ServicecatalogServiceAction struct {
	Name      string
	Args      ServicecatalogServiceActionArgs
	state     *servicecatalogServiceActionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicecatalogServiceAction].
func (ssa *ServicecatalogServiceAction) Type() string {
	return "aws_servicecatalog_service_action"
}

// LocalName returns the local name for [ServicecatalogServiceAction].
func (ssa *ServicecatalogServiceAction) LocalName() string {
	return ssa.Name
}

// Configuration returns the configuration (args) for [ServicecatalogServiceAction].
func (ssa *ServicecatalogServiceAction) Configuration() interface{} {
	return ssa.Args
}

// DependOn is used for other resources to depend on [ServicecatalogServiceAction].
func (ssa *ServicecatalogServiceAction) DependOn() terra.Reference {
	return terra.ReferenceResource(ssa)
}

// Dependencies returns the list of resources [ServicecatalogServiceAction] depends_on.
func (ssa *ServicecatalogServiceAction) Dependencies() terra.Dependencies {
	return ssa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicecatalogServiceAction].
func (ssa *ServicecatalogServiceAction) LifecycleManagement() *terra.Lifecycle {
	return ssa.Lifecycle
}

// Attributes returns the attributes for [ServicecatalogServiceAction].
func (ssa *ServicecatalogServiceAction) Attributes() servicecatalogServiceActionAttributes {
	return servicecatalogServiceActionAttributes{ref: terra.ReferenceResource(ssa)}
}

// ImportState imports the given attribute values into [ServicecatalogServiceAction]'s state.
func (ssa *ServicecatalogServiceAction) ImportState(av io.Reader) error {
	ssa.state = &servicecatalogServiceActionState{}
	if err := json.NewDecoder(av).Decode(ssa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssa.Type(), ssa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicecatalogServiceAction] has state.
func (ssa *ServicecatalogServiceAction) State() (*servicecatalogServiceActionState, bool) {
	return ssa.state, ssa.state != nil
}

// StateMust returns the state for [ServicecatalogServiceAction]. Panics if the state is nil.
func (ssa *ServicecatalogServiceAction) StateMust() *servicecatalogServiceActionState {
	if ssa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssa.Type(), ssa.LocalName()))
	}
	return ssa.state
}

// ServicecatalogServiceActionArgs contains the configurations for aws_servicecatalog_service_action.
type ServicecatalogServiceActionArgs struct {
	// AcceptLanguage: string, optional
	AcceptLanguage terra.StringValue `hcl:"accept_language,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Definition: required
	Definition *servicecatalogserviceaction.Definition `hcl:"definition,block" validate:"required"`
	// Timeouts: optional
	Timeouts *servicecatalogserviceaction.Timeouts `hcl:"timeouts,block"`
}
type servicecatalogServiceActionAttributes struct {
	ref terra.Reference
}

// AcceptLanguage returns a reference to field accept_language of aws_servicecatalog_service_action.
func (ssa servicecatalogServiceActionAttributes) AcceptLanguage() terra.StringValue {
	return terra.ReferenceAsString(ssa.ref.Append("accept_language"))
}

// Description returns a reference to field description of aws_servicecatalog_service_action.
func (ssa servicecatalogServiceActionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ssa.ref.Append("description"))
}

// Id returns a reference to field id of aws_servicecatalog_service_action.
func (ssa servicecatalogServiceActionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssa.ref.Append("id"))
}

// Name returns a reference to field name of aws_servicecatalog_service_action.
func (ssa servicecatalogServiceActionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssa.ref.Append("name"))
}

func (ssa servicecatalogServiceActionAttributes) Definition() terra.ListValue[servicecatalogserviceaction.DefinitionAttributes] {
	return terra.ReferenceAsList[servicecatalogserviceaction.DefinitionAttributes](ssa.ref.Append("definition"))
}

func (ssa servicecatalogServiceActionAttributes) Timeouts() servicecatalogserviceaction.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicecatalogserviceaction.TimeoutsAttributes](ssa.ref.Append("timeouts"))
}

type servicecatalogServiceActionState struct {
	AcceptLanguage string                                        `json:"accept_language"`
	Description    string                                        `json:"description"`
	Id             string                                        `json:"id"`
	Name           string                                        `json:"name"`
	Definition     []servicecatalogserviceaction.DefinitionState `json:"definition"`
	Timeouts       *servicecatalogserviceaction.TimeoutsState    `json:"timeouts"`
}

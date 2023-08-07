// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubenrichment "github.com/golingon/terraproviders/azurerm/3.68.0/iothubenrichment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothubEnrichment creates a new instance of [IothubEnrichment].
func NewIothubEnrichment(name string, args IothubEnrichmentArgs) *IothubEnrichment {
	return &IothubEnrichment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubEnrichment)(nil)

// IothubEnrichment represents the Terraform resource azurerm_iothub_enrichment.
type IothubEnrichment struct {
	Name      string
	Args      IothubEnrichmentArgs
	state     *iothubEnrichmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IothubEnrichment].
func (ie *IothubEnrichment) Type() string {
	return "azurerm_iothub_enrichment"
}

// LocalName returns the local name for [IothubEnrichment].
func (ie *IothubEnrichment) LocalName() string {
	return ie.Name
}

// Configuration returns the configuration (args) for [IothubEnrichment].
func (ie *IothubEnrichment) Configuration() interface{} {
	return ie.Args
}

// DependOn is used for other resources to depend on [IothubEnrichment].
func (ie *IothubEnrichment) DependOn() terra.Reference {
	return terra.ReferenceResource(ie)
}

// Dependencies returns the list of resources [IothubEnrichment] depends_on.
func (ie *IothubEnrichment) Dependencies() terra.Dependencies {
	return ie.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IothubEnrichment].
func (ie *IothubEnrichment) LifecycleManagement() *terra.Lifecycle {
	return ie.Lifecycle
}

// Attributes returns the attributes for [IothubEnrichment].
func (ie *IothubEnrichment) Attributes() iothubEnrichmentAttributes {
	return iothubEnrichmentAttributes{ref: terra.ReferenceResource(ie)}
}

// ImportState imports the given attribute values into [IothubEnrichment]'s state.
func (ie *IothubEnrichment) ImportState(av io.Reader) error {
	ie.state = &iothubEnrichmentState{}
	if err := json.NewDecoder(av).Decode(ie.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ie.Type(), ie.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IothubEnrichment] has state.
func (ie *IothubEnrichment) State() (*iothubEnrichmentState, bool) {
	return ie.state, ie.state != nil
}

// StateMust returns the state for [IothubEnrichment]. Panics if the state is nil.
func (ie *IothubEnrichment) StateMust() *iothubEnrichmentState {
	if ie.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ie.Type(), ie.LocalName()))
	}
	return ie.state
}

// IothubEnrichmentArgs contains the configurations for azurerm_iothub_enrichment.
type IothubEnrichmentArgs struct {
	// EndpointNames: list of string, required
	EndpointNames terra.ListValue[terra.StringValue] `hcl:"endpoint_names,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubName: string, required
	IothubName terra.StringValue `hcl:"iothub_name,attr" validate:"required"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *iothubenrichment.Timeouts `hcl:"timeouts,block"`
}
type iothubEnrichmentAttributes struct {
	ref terra.Reference
}

// EndpointNames returns a reference to field endpoint_names of azurerm_iothub_enrichment.
func (ie iothubEnrichmentAttributes) EndpointNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ie.ref.Append("endpoint_names"))
}

// Id returns a reference to field id of azurerm_iothub_enrichment.
func (ie iothubEnrichmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ie.ref.Append("id"))
}

// IothubName returns a reference to field iothub_name of azurerm_iothub_enrichment.
func (ie iothubEnrichmentAttributes) IothubName() terra.StringValue {
	return terra.ReferenceAsString(ie.ref.Append("iothub_name"))
}

// Key returns a reference to field key of azurerm_iothub_enrichment.
func (ie iothubEnrichmentAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(ie.ref.Append("key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_enrichment.
func (ie iothubEnrichmentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ie.ref.Append("resource_group_name"))
}

// Value returns a reference to field value of azurerm_iothub_enrichment.
func (ie iothubEnrichmentAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ie.ref.Append("value"))
}

func (ie iothubEnrichmentAttributes) Timeouts() iothubenrichment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothubenrichment.TimeoutsAttributes](ie.ref.Append("timeouts"))
}

type iothubEnrichmentState struct {
	EndpointNames     []string                        `json:"endpoint_names"`
	Id                string                          `json:"id"`
	IothubName        string                          `json:"iothub_name"`
	Key               string                          `json:"key"`
	ResourceGroupName string                          `json:"resource_group_name"`
	Value             string                          `json:"value"`
	Timeouts          *iothubenrichment.TimeoutsState `json:"timeouts"`
}

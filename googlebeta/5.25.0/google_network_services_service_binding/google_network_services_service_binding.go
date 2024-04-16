// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_network_services_service_binding

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_network_services_service_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleNetworkServicesServiceBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gnssb *Resource) Type() string {
	return "google_network_services_service_binding"
}

// LocalName returns the local name for [Resource].
func (gnssb *Resource) LocalName() string {
	return gnssb.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gnssb *Resource) Configuration() interface{} {
	return gnssb.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gnssb *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gnssb)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gnssb *Resource) Dependencies() terra.Dependencies {
	return gnssb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gnssb *Resource) LifecycleManagement() *terra.Lifecycle {
	return gnssb.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gnssb *Resource) Attributes() googleNetworkServicesServiceBindingAttributes {
	return googleNetworkServicesServiceBindingAttributes{ref: terra.ReferenceResource(gnssb)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gnssb *Resource) ImportState(state io.Reader) error {
	gnssb.state = &googleNetworkServicesServiceBindingState{}
	if err := json.NewDecoder(state).Decode(gnssb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gnssb.Type(), gnssb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gnssb *Resource) State() (*googleNetworkServicesServiceBindingState, bool) {
	return gnssb.state, gnssb.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gnssb *Resource) StateMust() *googleNetworkServicesServiceBindingState {
	if gnssb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gnssb.Type(), gnssb.LocalName()))
	}
	return gnssb.state
}

// Args contains the configurations for google_network_services_service_binding.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleNetworkServicesServiceBindingAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_services_service_binding.
func (gnssb googleNetworkServicesServiceBindingAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gnssb.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_services_service_binding.
func (gnssb googleNetworkServicesServiceBindingAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gnssb.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_network_services_service_binding.
func (gnssb googleNetworkServicesServiceBindingAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gnssb.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_network_services_service_binding.
func (gnssb googleNetworkServicesServiceBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gnssb.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_services_service_binding.
func (gnssb googleNetworkServicesServiceBindingAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gnssb.ref.Append("labels"))
}

// Name returns a reference to field name of google_network_services_service_binding.
func (gnssb googleNetworkServicesServiceBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gnssb.ref.Append("name"))
}

// Project returns a reference to field project of google_network_services_service_binding.
func (gnssb googleNetworkServicesServiceBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gnssb.ref.Append("project"))
}

// Service returns a reference to field service of google_network_services_service_binding.
func (gnssb googleNetworkServicesServiceBindingAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(gnssb.ref.Append("service"))
}

// TerraformLabels returns a reference to field terraform_labels of google_network_services_service_binding.
func (gnssb googleNetworkServicesServiceBindingAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gnssb.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_network_services_service_binding.
func (gnssb googleNetworkServicesServiceBindingAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gnssb.ref.Append("update_time"))
}

func (gnssb googleNetworkServicesServiceBindingAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gnssb.ref.Append("timeouts"))
}

type googleNetworkServicesServiceBindingState struct {
	CreateTime      string            `json:"create_time"`
	Description     string            `json:"description"`
	EffectiveLabels map[string]string `json:"effective_labels"`
	Id              string            `json:"id"`
	Labels          map[string]string `json:"labels"`
	Name            string            `json:"name"`
	Project         string            `json:"project"`
	Service         string            `json:"service"`
	TerraformLabels map[string]string `json:"terraform_labels"`
	UpdateTime      string            `json:"update_time"`
	Timeouts        *TimeoutsState    `json:"timeouts"`
}

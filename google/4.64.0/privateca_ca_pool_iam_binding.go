// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	privatecacapooliambinding "github.com/golingon/terraproviders/google/4.64.0/privatecacapooliambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivatecaCaPoolIamBinding creates a new instance of [PrivatecaCaPoolIamBinding].
func NewPrivatecaCaPoolIamBinding(name string, args PrivatecaCaPoolIamBindingArgs) *PrivatecaCaPoolIamBinding {
	return &PrivatecaCaPoolIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivatecaCaPoolIamBinding)(nil)

// PrivatecaCaPoolIamBinding represents the Terraform resource google_privateca_ca_pool_iam_binding.
type PrivatecaCaPoolIamBinding struct {
	Name      string
	Args      PrivatecaCaPoolIamBindingArgs
	state     *privatecaCaPoolIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivatecaCaPoolIamBinding].
func (pcpib *PrivatecaCaPoolIamBinding) Type() string {
	return "google_privateca_ca_pool_iam_binding"
}

// LocalName returns the local name for [PrivatecaCaPoolIamBinding].
func (pcpib *PrivatecaCaPoolIamBinding) LocalName() string {
	return pcpib.Name
}

// Configuration returns the configuration (args) for [PrivatecaCaPoolIamBinding].
func (pcpib *PrivatecaCaPoolIamBinding) Configuration() interface{} {
	return pcpib.Args
}

// DependOn is used for other resources to depend on [PrivatecaCaPoolIamBinding].
func (pcpib *PrivatecaCaPoolIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(pcpib)
}

// Dependencies returns the list of resources [PrivatecaCaPoolIamBinding] depends_on.
func (pcpib *PrivatecaCaPoolIamBinding) Dependencies() terra.Dependencies {
	return pcpib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivatecaCaPoolIamBinding].
func (pcpib *PrivatecaCaPoolIamBinding) LifecycleManagement() *terra.Lifecycle {
	return pcpib.Lifecycle
}

// Attributes returns the attributes for [PrivatecaCaPoolIamBinding].
func (pcpib *PrivatecaCaPoolIamBinding) Attributes() privatecaCaPoolIamBindingAttributes {
	return privatecaCaPoolIamBindingAttributes{ref: terra.ReferenceResource(pcpib)}
}

// ImportState imports the given attribute values into [PrivatecaCaPoolIamBinding]'s state.
func (pcpib *PrivatecaCaPoolIamBinding) ImportState(av io.Reader) error {
	pcpib.state = &privatecaCaPoolIamBindingState{}
	if err := json.NewDecoder(av).Decode(pcpib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pcpib.Type(), pcpib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivatecaCaPoolIamBinding] has state.
func (pcpib *PrivatecaCaPoolIamBinding) State() (*privatecaCaPoolIamBindingState, bool) {
	return pcpib.state, pcpib.state != nil
}

// StateMust returns the state for [PrivatecaCaPoolIamBinding]. Panics if the state is nil.
func (pcpib *PrivatecaCaPoolIamBinding) StateMust() *privatecaCaPoolIamBindingState {
	if pcpib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pcpib.Type(), pcpib.LocalName()))
	}
	return pcpib.state
}

// PrivatecaCaPoolIamBindingArgs contains the configurations for google_privateca_ca_pool_iam_binding.
type PrivatecaCaPoolIamBindingArgs struct {
	// CaPool: string, required
	CaPool terra.StringValue `hcl:"ca_pool,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *privatecacapooliambinding.Condition `hcl:"condition,block"`
}
type privatecaCaPoolIamBindingAttributes struct {
	ref terra.Reference
}

// CaPool returns a reference to field ca_pool of google_privateca_ca_pool_iam_binding.
func (pcpib privatecaCaPoolIamBindingAttributes) CaPool() terra.StringValue {
	return terra.ReferenceAsString(pcpib.ref.Append("ca_pool"))
}

// Etag returns a reference to field etag of google_privateca_ca_pool_iam_binding.
func (pcpib privatecaCaPoolIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(pcpib.ref.Append("etag"))
}

// Id returns a reference to field id of google_privateca_ca_pool_iam_binding.
func (pcpib privatecaCaPoolIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pcpib.ref.Append("id"))
}

// Location returns a reference to field location of google_privateca_ca_pool_iam_binding.
func (pcpib privatecaCaPoolIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pcpib.ref.Append("location"))
}

// Members returns a reference to field members of google_privateca_ca_pool_iam_binding.
func (pcpib privatecaCaPoolIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pcpib.ref.Append("members"))
}

// Project returns a reference to field project of google_privateca_ca_pool_iam_binding.
func (pcpib privatecaCaPoolIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pcpib.ref.Append("project"))
}

// Role returns a reference to field role of google_privateca_ca_pool_iam_binding.
func (pcpib privatecaCaPoolIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(pcpib.ref.Append("role"))
}

func (pcpib privatecaCaPoolIamBindingAttributes) Condition() terra.ListValue[privatecacapooliambinding.ConditionAttributes] {
	return terra.ReferenceAsList[privatecacapooliambinding.ConditionAttributes](pcpib.ref.Append("condition"))
}

type privatecaCaPoolIamBindingState struct {
	CaPool    string                                     `json:"ca_pool"`
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Location  string                                     `json:"location"`
	Members   []string                                   `json:"members"`
	Project   string                                     `json:"project"`
	Role      string                                     `json:"role"`
	Condition []privatecacapooliambinding.ConditionState `json:"condition"`
}

// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_network_security_address_group_iam_binding

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

// Resource represents the Terraform resource google_network_security_address_group_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleNetworkSecurityAddressGroupIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gnsagib *Resource) Type() string {
	return "google_network_security_address_group_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gnsagib *Resource) LocalName() string {
	return gnsagib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gnsagib *Resource) Configuration() interface{} {
	return gnsagib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gnsagib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gnsagib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gnsagib *Resource) Dependencies() terra.Dependencies {
	return gnsagib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gnsagib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gnsagib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gnsagib *Resource) Attributes() googleNetworkSecurityAddressGroupIamBindingAttributes {
	return googleNetworkSecurityAddressGroupIamBindingAttributes{ref: terra.ReferenceResource(gnsagib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gnsagib *Resource) ImportState(state io.Reader) error {
	gnsagib.state = &googleNetworkSecurityAddressGroupIamBindingState{}
	if err := json.NewDecoder(state).Decode(gnsagib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gnsagib.Type(), gnsagib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gnsagib *Resource) State() (*googleNetworkSecurityAddressGroupIamBindingState, bool) {
	return gnsagib.state, gnsagib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gnsagib *Resource) StateMust() *googleNetworkSecurityAddressGroupIamBindingState {
	if gnsagib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gnsagib.Type(), gnsagib.LocalName()))
	}
	return gnsagib.state
}

// Args contains the configurations for google_network_security_address_group_iam_binding.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleNetworkSecurityAddressGroupIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_network_security_address_group_iam_binding.
func (gnsagib googleNetworkSecurityAddressGroupIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gnsagib.ref.Append("etag"))
}

// Id returns a reference to field id of google_network_security_address_group_iam_binding.
func (gnsagib googleNetworkSecurityAddressGroupIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gnsagib.ref.Append("id"))
}

// Location returns a reference to field location of google_network_security_address_group_iam_binding.
func (gnsagib googleNetworkSecurityAddressGroupIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gnsagib.ref.Append("location"))
}

// Members returns a reference to field members of google_network_security_address_group_iam_binding.
func (gnsagib googleNetworkSecurityAddressGroupIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gnsagib.ref.Append("members"))
}

// Name returns a reference to field name of google_network_security_address_group_iam_binding.
func (gnsagib googleNetworkSecurityAddressGroupIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gnsagib.ref.Append("name"))
}

// Project returns a reference to field project of google_network_security_address_group_iam_binding.
func (gnsagib googleNetworkSecurityAddressGroupIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gnsagib.ref.Append("project"))
}

// Role returns a reference to field role of google_network_security_address_group_iam_binding.
func (gnsagib googleNetworkSecurityAddressGroupIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gnsagib.ref.Append("role"))
}

func (gnsagib googleNetworkSecurityAddressGroupIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gnsagib.ref.Append("condition"))
}

type googleNetworkSecurityAddressGroupIamBindingState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Location  string           `json:"location"`
	Members   []string         `json:"members"`
	Name      string           `json:"name"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}

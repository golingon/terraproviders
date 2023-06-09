// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dnsmanagedzoneiambinding "github.com/golingon/terraproviders/google/4.72.1/dnsmanagedzoneiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsManagedZoneIamBinding creates a new instance of [DnsManagedZoneIamBinding].
func NewDnsManagedZoneIamBinding(name string, args DnsManagedZoneIamBindingArgs) *DnsManagedZoneIamBinding {
	return &DnsManagedZoneIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsManagedZoneIamBinding)(nil)

// DnsManagedZoneIamBinding represents the Terraform resource google_dns_managed_zone_iam_binding.
type DnsManagedZoneIamBinding struct {
	Name      string
	Args      DnsManagedZoneIamBindingArgs
	state     *dnsManagedZoneIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsManagedZoneIamBinding].
func (dmzib *DnsManagedZoneIamBinding) Type() string {
	return "google_dns_managed_zone_iam_binding"
}

// LocalName returns the local name for [DnsManagedZoneIamBinding].
func (dmzib *DnsManagedZoneIamBinding) LocalName() string {
	return dmzib.Name
}

// Configuration returns the configuration (args) for [DnsManagedZoneIamBinding].
func (dmzib *DnsManagedZoneIamBinding) Configuration() interface{} {
	return dmzib.Args
}

// DependOn is used for other resources to depend on [DnsManagedZoneIamBinding].
func (dmzib *DnsManagedZoneIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(dmzib)
}

// Dependencies returns the list of resources [DnsManagedZoneIamBinding] depends_on.
func (dmzib *DnsManagedZoneIamBinding) Dependencies() terra.Dependencies {
	return dmzib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsManagedZoneIamBinding].
func (dmzib *DnsManagedZoneIamBinding) LifecycleManagement() *terra.Lifecycle {
	return dmzib.Lifecycle
}

// Attributes returns the attributes for [DnsManagedZoneIamBinding].
func (dmzib *DnsManagedZoneIamBinding) Attributes() dnsManagedZoneIamBindingAttributes {
	return dnsManagedZoneIamBindingAttributes{ref: terra.ReferenceResource(dmzib)}
}

// ImportState imports the given attribute values into [DnsManagedZoneIamBinding]'s state.
func (dmzib *DnsManagedZoneIamBinding) ImportState(av io.Reader) error {
	dmzib.state = &dnsManagedZoneIamBindingState{}
	if err := json.NewDecoder(av).Decode(dmzib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmzib.Type(), dmzib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsManagedZoneIamBinding] has state.
func (dmzib *DnsManagedZoneIamBinding) State() (*dnsManagedZoneIamBindingState, bool) {
	return dmzib.state, dmzib.state != nil
}

// StateMust returns the state for [DnsManagedZoneIamBinding]. Panics if the state is nil.
func (dmzib *DnsManagedZoneIamBinding) StateMust() *dnsManagedZoneIamBindingState {
	if dmzib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmzib.Type(), dmzib.LocalName()))
	}
	return dmzib.state
}

// DnsManagedZoneIamBindingArgs contains the configurations for google_dns_managed_zone_iam_binding.
type DnsManagedZoneIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedZone: string, required
	ManagedZone terra.StringValue `hcl:"managed_zone,attr" validate:"required"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *dnsmanagedzoneiambinding.Condition `hcl:"condition,block"`
}
type dnsManagedZoneIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dns_managed_zone_iam_binding.
func (dmzib dnsManagedZoneIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dmzib.ref.Append("etag"))
}

// Id returns a reference to field id of google_dns_managed_zone_iam_binding.
func (dmzib dnsManagedZoneIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmzib.ref.Append("id"))
}

// ManagedZone returns a reference to field managed_zone of google_dns_managed_zone_iam_binding.
func (dmzib dnsManagedZoneIamBindingAttributes) ManagedZone() terra.StringValue {
	return terra.ReferenceAsString(dmzib.ref.Append("managed_zone"))
}

// Members returns a reference to field members of google_dns_managed_zone_iam_binding.
func (dmzib dnsManagedZoneIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dmzib.ref.Append("members"))
}

// Project returns a reference to field project of google_dns_managed_zone_iam_binding.
func (dmzib dnsManagedZoneIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmzib.ref.Append("project"))
}

// Role returns a reference to field role of google_dns_managed_zone_iam_binding.
func (dmzib dnsManagedZoneIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dmzib.ref.Append("role"))
}

func (dmzib dnsManagedZoneIamBindingAttributes) Condition() terra.ListValue[dnsmanagedzoneiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[dnsmanagedzoneiambinding.ConditionAttributes](dmzib.ref.Append("condition"))
}

type dnsManagedZoneIamBindingState struct {
	Etag        string                                    `json:"etag"`
	Id          string                                    `json:"id"`
	ManagedZone string                                    `json:"managed_zone"`
	Members     []string                                  `json:"members"`
	Project     string                                    `json:"project"`
	Role        string                                    `json:"role"`
	Condition   []dnsmanagedzoneiambinding.ConditionState `json:"condition"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sccsourceiambinding "github.com/golingon/terraproviders/google/4.74.0/sccsourceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSccSourceIamBinding creates a new instance of [SccSourceIamBinding].
func NewSccSourceIamBinding(name string, args SccSourceIamBindingArgs) *SccSourceIamBinding {
	return &SccSourceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SccSourceIamBinding)(nil)

// SccSourceIamBinding represents the Terraform resource google_scc_source_iam_binding.
type SccSourceIamBinding struct {
	Name      string
	Args      SccSourceIamBindingArgs
	state     *sccSourceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SccSourceIamBinding].
func (ssib *SccSourceIamBinding) Type() string {
	return "google_scc_source_iam_binding"
}

// LocalName returns the local name for [SccSourceIamBinding].
func (ssib *SccSourceIamBinding) LocalName() string {
	return ssib.Name
}

// Configuration returns the configuration (args) for [SccSourceIamBinding].
func (ssib *SccSourceIamBinding) Configuration() interface{} {
	return ssib.Args
}

// DependOn is used for other resources to depend on [SccSourceIamBinding].
func (ssib *SccSourceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ssib)
}

// Dependencies returns the list of resources [SccSourceIamBinding] depends_on.
func (ssib *SccSourceIamBinding) Dependencies() terra.Dependencies {
	return ssib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SccSourceIamBinding].
func (ssib *SccSourceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ssib.Lifecycle
}

// Attributes returns the attributes for [SccSourceIamBinding].
func (ssib *SccSourceIamBinding) Attributes() sccSourceIamBindingAttributes {
	return sccSourceIamBindingAttributes{ref: terra.ReferenceResource(ssib)}
}

// ImportState imports the given attribute values into [SccSourceIamBinding]'s state.
func (ssib *SccSourceIamBinding) ImportState(av io.Reader) error {
	ssib.state = &sccSourceIamBindingState{}
	if err := json.NewDecoder(av).Decode(ssib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssib.Type(), ssib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SccSourceIamBinding] has state.
func (ssib *SccSourceIamBinding) State() (*sccSourceIamBindingState, bool) {
	return ssib.state, ssib.state != nil
}

// StateMust returns the state for [SccSourceIamBinding]. Panics if the state is nil.
func (ssib *SccSourceIamBinding) StateMust() *sccSourceIamBindingState {
	if ssib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssib.Type(), ssib.LocalName()))
	}
	return ssib.state
}

// SccSourceIamBindingArgs contains the configurations for google_scc_source_iam_binding.
type SccSourceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// Condition: optional
	Condition *sccsourceiambinding.Condition `hcl:"condition,block"`
}
type sccSourceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_scc_source_iam_binding.
func (ssib sccSourceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ssib.ref.Append("etag"))
}

// Id returns a reference to field id of google_scc_source_iam_binding.
func (ssib sccSourceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssib.ref.Append("id"))
}

// Members returns a reference to field members of google_scc_source_iam_binding.
func (ssib sccSourceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ssib.ref.Append("members"))
}

// Organization returns a reference to field organization of google_scc_source_iam_binding.
func (ssib sccSourceIamBindingAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(ssib.ref.Append("organization"))
}

// Role returns a reference to field role of google_scc_source_iam_binding.
func (ssib sccSourceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ssib.ref.Append("role"))
}

// Source returns a reference to field source of google_scc_source_iam_binding.
func (ssib sccSourceIamBindingAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(ssib.ref.Append("source"))
}

func (ssib sccSourceIamBindingAttributes) Condition() terra.ListValue[sccsourceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[sccsourceiambinding.ConditionAttributes](ssib.ref.Append("condition"))
}

type sccSourceIamBindingState struct {
	Etag         string                               `json:"etag"`
	Id           string                               `json:"id"`
	Members      []string                             `json:"members"`
	Organization string                               `json:"organization"`
	Role         string                               `json:"role"`
	Source       string                               `json:"source"`
	Condition    []sccsourceiambinding.ConditionState `json:"condition"`
}
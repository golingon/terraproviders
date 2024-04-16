// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_macie2_account

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

// Resource represents the Terraform resource aws_macie2_account.
type Resource struct {
	Name      string
	Args      Args
	state     *awsMacie2AccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ama *Resource) Type() string {
	return "aws_macie2_account"
}

// LocalName returns the local name for [Resource].
func (ama *Resource) LocalName() string {
	return ama.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ama *Resource) Configuration() interface{} {
	return ama.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ama *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ama)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ama *Resource) Dependencies() terra.Dependencies {
	return ama.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ama *Resource) LifecycleManagement() *terra.Lifecycle {
	return ama.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ama *Resource) Attributes() awsMacie2AccountAttributes {
	return awsMacie2AccountAttributes{ref: terra.ReferenceResource(ama)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ama *Resource) ImportState(state io.Reader) error {
	ama.state = &awsMacie2AccountState{}
	if err := json.NewDecoder(state).Decode(ama.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ama.Type(), ama.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ama *Resource) State() (*awsMacie2AccountState, bool) {
	return ama.state, ama.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ama *Resource) StateMust() *awsMacie2AccountState {
	if ama.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ama.Type(), ama.LocalName()))
	}
	return ama.state
}

// Args contains the configurations for aws_macie2_account.
type Args struct {
	// FindingPublishingFrequency: string, optional
	FindingPublishingFrequency terra.StringValue `hcl:"finding_publishing_frequency,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
}

type awsMacie2AccountAttributes struct {
	ref terra.Reference
}

// CreatedAt returns a reference to field created_at of aws_macie2_account.
func (ama awsMacie2AccountAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("created_at"))
}

// FindingPublishingFrequency returns a reference to field finding_publishing_frequency of aws_macie2_account.
func (ama awsMacie2AccountAttributes) FindingPublishingFrequency() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("finding_publishing_frequency"))
}

// Id returns a reference to field id of aws_macie2_account.
func (ama awsMacie2AccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("id"))
}

// ServiceRole returns a reference to field service_role of aws_macie2_account.
func (ama awsMacie2AccountAttributes) ServiceRole() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("service_role"))
}

// Status returns a reference to field status of aws_macie2_account.
func (ama awsMacie2AccountAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("status"))
}

// UpdatedAt returns a reference to field updated_at of aws_macie2_account.
func (ama awsMacie2AccountAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("updated_at"))
}

type awsMacie2AccountState struct {
	CreatedAt                  string `json:"created_at"`
	FindingPublishingFrequency string `json:"finding_publishing_frequency"`
	Id                         string `json:"id"`
	ServiceRole                string `json:"service_role"`
	Status                     string `json:"status"`
	UpdatedAt                  string `json:"updated_at"`
}

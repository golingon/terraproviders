// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudbuildv2connectioniambinding "github.com/golingon/terraproviders/googlebeta/4.72.1/cloudbuildv2connectioniambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudbuildv2ConnectionIamBinding creates a new instance of [Cloudbuildv2ConnectionIamBinding].
func NewCloudbuildv2ConnectionIamBinding(name string, args Cloudbuildv2ConnectionIamBindingArgs) *Cloudbuildv2ConnectionIamBinding {
	return &Cloudbuildv2ConnectionIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudbuildv2ConnectionIamBinding)(nil)

// Cloudbuildv2ConnectionIamBinding represents the Terraform resource google_cloudbuildv2_connection_iam_binding.
type Cloudbuildv2ConnectionIamBinding struct {
	Name      string
	Args      Cloudbuildv2ConnectionIamBindingArgs
	state     *cloudbuildv2ConnectionIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Cloudbuildv2ConnectionIamBinding].
func (ccib *Cloudbuildv2ConnectionIamBinding) Type() string {
	return "google_cloudbuildv2_connection_iam_binding"
}

// LocalName returns the local name for [Cloudbuildv2ConnectionIamBinding].
func (ccib *Cloudbuildv2ConnectionIamBinding) LocalName() string {
	return ccib.Name
}

// Configuration returns the configuration (args) for [Cloudbuildv2ConnectionIamBinding].
func (ccib *Cloudbuildv2ConnectionIamBinding) Configuration() interface{} {
	return ccib.Args
}

// DependOn is used for other resources to depend on [Cloudbuildv2ConnectionIamBinding].
func (ccib *Cloudbuildv2ConnectionIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ccib)
}

// Dependencies returns the list of resources [Cloudbuildv2ConnectionIamBinding] depends_on.
func (ccib *Cloudbuildv2ConnectionIamBinding) Dependencies() terra.Dependencies {
	return ccib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Cloudbuildv2ConnectionIamBinding].
func (ccib *Cloudbuildv2ConnectionIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ccib.Lifecycle
}

// Attributes returns the attributes for [Cloudbuildv2ConnectionIamBinding].
func (ccib *Cloudbuildv2ConnectionIamBinding) Attributes() cloudbuildv2ConnectionIamBindingAttributes {
	return cloudbuildv2ConnectionIamBindingAttributes{ref: terra.ReferenceResource(ccib)}
}

// ImportState imports the given attribute values into [Cloudbuildv2ConnectionIamBinding]'s state.
func (ccib *Cloudbuildv2ConnectionIamBinding) ImportState(av io.Reader) error {
	ccib.state = &cloudbuildv2ConnectionIamBindingState{}
	if err := json.NewDecoder(av).Decode(ccib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccib.Type(), ccib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Cloudbuildv2ConnectionIamBinding] has state.
func (ccib *Cloudbuildv2ConnectionIamBinding) State() (*cloudbuildv2ConnectionIamBindingState, bool) {
	return ccib.state, ccib.state != nil
}

// StateMust returns the state for [Cloudbuildv2ConnectionIamBinding]. Panics if the state is nil.
func (ccib *Cloudbuildv2ConnectionIamBinding) StateMust() *cloudbuildv2ConnectionIamBindingState {
	if ccib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccib.Type(), ccib.LocalName()))
	}
	return ccib.state
}

// Cloudbuildv2ConnectionIamBindingArgs contains the configurations for google_cloudbuildv2_connection_iam_binding.
type Cloudbuildv2ConnectionIamBindingArgs struct {
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
	Condition *cloudbuildv2connectioniambinding.Condition `hcl:"condition,block"`
}
type cloudbuildv2ConnectionIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloudbuildv2_connection_iam_binding.
func (ccib cloudbuildv2ConnectionIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ccib.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudbuildv2_connection_iam_binding.
func (ccib cloudbuildv2ConnectionIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccib.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudbuildv2_connection_iam_binding.
func (ccib cloudbuildv2ConnectionIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ccib.ref.Append("location"))
}

// Members returns a reference to field members of google_cloudbuildv2_connection_iam_binding.
func (ccib cloudbuildv2ConnectionIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ccib.ref.Append("members"))
}

// Name returns a reference to field name of google_cloudbuildv2_connection_iam_binding.
func (ccib cloudbuildv2ConnectionIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccib.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudbuildv2_connection_iam_binding.
func (ccib cloudbuildv2ConnectionIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ccib.ref.Append("project"))
}

// Role returns a reference to field role of google_cloudbuildv2_connection_iam_binding.
func (ccib cloudbuildv2ConnectionIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ccib.ref.Append("role"))
}

func (ccib cloudbuildv2ConnectionIamBindingAttributes) Condition() terra.ListValue[cloudbuildv2connectioniambinding.ConditionAttributes] {
	return terra.ReferenceAsList[cloudbuildv2connectioniambinding.ConditionAttributes](ccib.ref.Append("condition"))
}

type cloudbuildv2ConnectionIamBindingState struct {
	Etag      string                                            `json:"etag"`
	Id        string                                            `json:"id"`
	Location  string                                            `json:"location"`
	Members   []string                                          `json:"members"`
	Name      string                                            `json:"name"`
	Project   string                                            `json:"project"`
	Role      string                                            `json:"role"`
	Condition []cloudbuildv2connectioniambinding.ConditionState `json:"condition"`
}

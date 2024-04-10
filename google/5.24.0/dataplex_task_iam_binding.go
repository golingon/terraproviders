// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	dataplextaskiambinding "github.com/golingon/terraproviders/google/5.24.0/dataplextaskiambinding"
	"io"
)

// NewDataplexTaskIamBinding creates a new instance of [DataplexTaskIamBinding].
func NewDataplexTaskIamBinding(name string, args DataplexTaskIamBindingArgs) *DataplexTaskIamBinding {
	return &DataplexTaskIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexTaskIamBinding)(nil)

// DataplexTaskIamBinding represents the Terraform resource google_dataplex_task_iam_binding.
type DataplexTaskIamBinding struct {
	Name      string
	Args      DataplexTaskIamBindingArgs
	state     *dataplexTaskIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexTaskIamBinding].
func (dtib *DataplexTaskIamBinding) Type() string {
	return "google_dataplex_task_iam_binding"
}

// LocalName returns the local name for [DataplexTaskIamBinding].
func (dtib *DataplexTaskIamBinding) LocalName() string {
	return dtib.Name
}

// Configuration returns the configuration (args) for [DataplexTaskIamBinding].
func (dtib *DataplexTaskIamBinding) Configuration() interface{} {
	return dtib.Args
}

// DependOn is used for other resources to depend on [DataplexTaskIamBinding].
func (dtib *DataplexTaskIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(dtib)
}

// Dependencies returns the list of resources [DataplexTaskIamBinding] depends_on.
func (dtib *DataplexTaskIamBinding) Dependencies() terra.Dependencies {
	return dtib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexTaskIamBinding].
func (dtib *DataplexTaskIamBinding) LifecycleManagement() *terra.Lifecycle {
	return dtib.Lifecycle
}

// Attributes returns the attributes for [DataplexTaskIamBinding].
func (dtib *DataplexTaskIamBinding) Attributes() dataplexTaskIamBindingAttributes {
	return dataplexTaskIamBindingAttributes{ref: terra.ReferenceResource(dtib)}
}

// ImportState imports the given attribute values into [DataplexTaskIamBinding]'s state.
func (dtib *DataplexTaskIamBinding) ImportState(av io.Reader) error {
	dtib.state = &dataplexTaskIamBindingState{}
	if err := json.NewDecoder(av).Decode(dtib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtib.Type(), dtib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexTaskIamBinding] has state.
func (dtib *DataplexTaskIamBinding) State() (*dataplexTaskIamBindingState, bool) {
	return dtib.state, dtib.state != nil
}

// StateMust returns the state for [DataplexTaskIamBinding]. Panics if the state is nil.
func (dtib *DataplexTaskIamBinding) StateMust() *dataplexTaskIamBindingState {
	if dtib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtib.Type(), dtib.LocalName()))
	}
	return dtib.state
}

// DataplexTaskIamBindingArgs contains the configurations for google_dataplex_task_iam_binding.
type DataplexTaskIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TaskId: string, required
	TaskId terra.StringValue `hcl:"task_id,attr" validate:"required"`
	// Condition: optional
	Condition *dataplextaskiambinding.Condition `hcl:"condition,block"`
}
type dataplexTaskIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataplex_task_iam_binding.
func (dtib dataplexTaskIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dtib.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_task_iam_binding.
func (dtib dataplexTaskIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtib.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_task_iam_binding.
func (dtib dataplexTaskIamBindingAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dtib.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_task_iam_binding.
func (dtib dataplexTaskIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dtib.ref.Append("location"))
}

// Members returns a reference to field members of google_dataplex_task_iam_binding.
func (dtib dataplexTaskIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dtib.ref.Append("members"))
}

// Project returns a reference to field project of google_dataplex_task_iam_binding.
func (dtib dataplexTaskIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dtib.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_task_iam_binding.
func (dtib dataplexTaskIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dtib.ref.Append("role"))
}

// TaskId returns a reference to field task_id of google_dataplex_task_iam_binding.
func (dtib dataplexTaskIamBindingAttributes) TaskId() terra.StringValue {
	return terra.ReferenceAsString(dtib.ref.Append("task_id"))
}

func (dtib dataplexTaskIamBindingAttributes) Condition() terra.ListValue[dataplextaskiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[dataplextaskiambinding.ConditionAttributes](dtib.ref.Append("condition"))
}

type dataplexTaskIamBindingState struct {
	Etag      string                                  `json:"etag"`
	Id        string                                  `json:"id"`
	Lake      string                                  `json:"lake"`
	Location  string                                  `json:"location"`
	Members   []string                                `json:"members"`
	Project   string                                  `json:"project"`
	Role      string                                  `json:"role"`
	TaskId    string                                  `json:"task_id"`
	Condition []dataplextaskiambinding.ConditionState `json:"condition"`
}

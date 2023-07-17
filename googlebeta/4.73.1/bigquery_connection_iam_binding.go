// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigqueryconnectioniambinding "github.com/golingon/terraproviders/googlebeta/4.73.1/bigqueryconnectioniambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryConnectionIamBinding creates a new instance of [BigqueryConnectionIamBinding].
func NewBigqueryConnectionIamBinding(name string, args BigqueryConnectionIamBindingArgs) *BigqueryConnectionIamBinding {
	return &BigqueryConnectionIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryConnectionIamBinding)(nil)

// BigqueryConnectionIamBinding represents the Terraform resource google_bigquery_connection_iam_binding.
type BigqueryConnectionIamBinding struct {
	Name      string
	Args      BigqueryConnectionIamBindingArgs
	state     *bigqueryConnectionIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryConnectionIamBinding].
func (bcib *BigqueryConnectionIamBinding) Type() string {
	return "google_bigquery_connection_iam_binding"
}

// LocalName returns the local name for [BigqueryConnectionIamBinding].
func (bcib *BigqueryConnectionIamBinding) LocalName() string {
	return bcib.Name
}

// Configuration returns the configuration (args) for [BigqueryConnectionIamBinding].
func (bcib *BigqueryConnectionIamBinding) Configuration() interface{} {
	return bcib.Args
}

// DependOn is used for other resources to depend on [BigqueryConnectionIamBinding].
func (bcib *BigqueryConnectionIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(bcib)
}

// Dependencies returns the list of resources [BigqueryConnectionIamBinding] depends_on.
func (bcib *BigqueryConnectionIamBinding) Dependencies() terra.Dependencies {
	return bcib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryConnectionIamBinding].
func (bcib *BigqueryConnectionIamBinding) LifecycleManagement() *terra.Lifecycle {
	return bcib.Lifecycle
}

// Attributes returns the attributes for [BigqueryConnectionIamBinding].
func (bcib *BigqueryConnectionIamBinding) Attributes() bigqueryConnectionIamBindingAttributes {
	return bigqueryConnectionIamBindingAttributes{ref: terra.ReferenceResource(bcib)}
}

// ImportState imports the given attribute values into [BigqueryConnectionIamBinding]'s state.
func (bcib *BigqueryConnectionIamBinding) ImportState(av io.Reader) error {
	bcib.state = &bigqueryConnectionIamBindingState{}
	if err := json.NewDecoder(av).Decode(bcib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcib.Type(), bcib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryConnectionIamBinding] has state.
func (bcib *BigqueryConnectionIamBinding) State() (*bigqueryConnectionIamBindingState, bool) {
	return bcib.state, bcib.state != nil
}

// StateMust returns the state for [BigqueryConnectionIamBinding]. Panics if the state is nil.
func (bcib *BigqueryConnectionIamBinding) StateMust() *bigqueryConnectionIamBindingState {
	if bcib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcib.Type(), bcib.LocalName()))
	}
	return bcib.state
}

// BigqueryConnectionIamBindingArgs contains the configurations for google_bigquery_connection_iam_binding.
type BigqueryConnectionIamBindingArgs struct {
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
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
	Condition *bigqueryconnectioniambinding.Condition `hcl:"condition,block"`
}
type bigqueryConnectionIamBindingAttributes struct {
	ref terra.Reference
}

// ConnectionId returns a reference to field connection_id of google_bigquery_connection_iam_binding.
func (bcib bigqueryConnectionIamBindingAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(bcib.ref.Append("connection_id"))
}

// Etag returns a reference to field etag of google_bigquery_connection_iam_binding.
func (bcib bigqueryConnectionIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bcib.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_connection_iam_binding.
func (bcib bigqueryConnectionIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcib.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_connection_iam_binding.
func (bcib bigqueryConnectionIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcib.ref.Append("location"))
}

// Members returns a reference to field members of google_bigquery_connection_iam_binding.
func (bcib bigqueryConnectionIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bcib.ref.Append("members"))
}

// Project returns a reference to field project of google_bigquery_connection_iam_binding.
func (bcib bigqueryConnectionIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bcib.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_connection_iam_binding.
func (bcib bigqueryConnectionIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bcib.ref.Append("role"))
}

func (bcib bigqueryConnectionIamBindingAttributes) Condition() terra.ListValue[bigqueryconnectioniambinding.ConditionAttributes] {
	return terra.ReferenceAsList[bigqueryconnectioniambinding.ConditionAttributes](bcib.ref.Append("condition"))
}

type bigqueryConnectionIamBindingState struct {
	ConnectionId string                                        `json:"connection_id"`
	Etag         string                                        `json:"etag"`
	Id           string                                        `json:"id"`
	Location     string                                        `json:"location"`
	Members      []string                                      `json:"members"`
	Project      string                                        `json:"project"`
	Role         string                                        `json:"role"`
	Condition    []bigqueryconnectioniambinding.ConditionState `json:"condition"`
}

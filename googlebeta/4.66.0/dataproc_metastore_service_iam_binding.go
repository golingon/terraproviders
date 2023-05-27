// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataprocmetastoreserviceiambinding "github.com/golingon/terraproviders/googlebeta/4.66.0/dataprocmetastoreserviceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocMetastoreServiceIamBinding creates a new instance of [DataprocMetastoreServiceIamBinding].
func NewDataprocMetastoreServiceIamBinding(name string, args DataprocMetastoreServiceIamBindingArgs) *DataprocMetastoreServiceIamBinding {
	return &DataprocMetastoreServiceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocMetastoreServiceIamBinding)(nil)

// DataprocMetastoreServiceIamBinding represents the Terraform resource google_dataproc_metastore_service_iam_binding.
type DataprocMetastoreServiceIamBinding struct {
	Name      string
	Args      DataprocMetastoreServiceIamBindingArgs
	state     *dataprocMetastoreServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocMetastoreServiceIamBinding].
func (dmsib *DataprocMetastoreServiceIamBinding) Type() string {
	return "google_dataproc_metastore_service_iam_binding"
}

// LocalName returns the local name for [DataprocMetastoreServiceIamBinding].
func (dmsib *DataprocMetastoreServiceIamBinding) LocalName() string {
	return dmsib.Name
}

// Configuration returns the configuration (args) for [DataprocMetastoreServiceIamBinding].
func (dmsib *DataprocMetastoreServiceIamBinding) Configuration() interface{} {
	return dmsib.Args
}

// DependOn is used for other resources to depend on [DataprocMetastoreServiceIamBinding].
func (dmsib *DataprocMetastoreServiceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(dmsib)
}

// Dependencies returns the list of resources [DataprocMetastoreServiceIamBinding] depends_on.
func (dmsib *DataprocMetastoreServiceIamBinding) Dependencies() terra.Dependencies {
	return dmsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocMetastoreServiceIamBinding].
func (dmsib *DataprocMetastoreServiceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return dmsib.Lifecycle
}

// Attributes returns the attributes for [DataprocMetastoreServiceIamBinding].
func (dmsib *DataprocMetastoreServiceIamBinding) Attributes() dataprocMetastoreServiceIamBindingAttributes {
	return dataprocMetastoreServiceIamBindingAttributes{ref: terra.ReferenceResource(dmsib)}
}

// ImportState imports the given attribute values into [DataprocMetastoreServiceIamBinding]'s state.
func (dmsib *DataprocMetastoreServiceIamBinding) ImportState(av io.Reader) error {
	dmsib.state = &dataprocMetastoreServiceIamBindingState{}
	if err := json.NewDecoder(av).Decode(dmsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmsib.Type(), dmsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocMetastoreServiceIamBinding] has state.
func (dmsib *DataprocMetastoreServiceIamBinding) State() (*dataprocMetastoreServiceIamBindingState, bool) {
	return dmsib.state, dmsib.state != nil
}

// StateMust returns the state for [DataprocMetastoreServiceIamBinding]. Panics if the state is nil.
func (dmsib *DataprocMetastoreServiceIamBinding) StateMust() *dataprocMetastoreServiceIamBindingState {
	if dmsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmsib.Type(), dmsib.LocalName()))
	}
	return dmsib.state
}

// DataprocMetastoreServiceIamBindingArgs contains the configurations for google_dataproc_metastore_service_iam_binding.
type DataprocMetastoreServiceIamBindingArgs struct {
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
	// ServiceId: string, required
	ServiceId terra.StringValue `hcl:"service_id,attr" validate:"required"`
	// Condition: optional
	Condition *dataprocmetastoreserviceiambinding.Condition `hcl:"condition,block"`
}
type dataprocMetastoreServiceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_metastore_service_iam_binding.
func (dmsib dataprocMetastoreServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dmsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_metastore_service_iam_binding.
func (dmsib dataprocMetastoreServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmsib.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_metastore_service_iam_binding.
func (dmsib dataprocMetastoreServiceIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dmsib.ref.Append("location"))
}

// Members returns a reference to field members of google_dataproc_metastore_service_iam_binding.
func (dmsib dataprocMetastoreServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dmsib.ref.Append("members"))
}

// Project returns a reference to field project of google_dataproc_metastore_service_iam_binding.
func (dmsib dataprocMetastoreServiceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmsib.ref.Append("project"))
}

// Role returns a reference to field role of google_dataproc_metastore_service_iam_binding.
func (dmsib dataprocMetastoreServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dmsib.ref.Append("role"))
}

// ServiceId returns a reference to field service_id of google_dataproc_metastore_service_iam_binding.
func (dmsib dataprocMetastoreServiceIamBindingAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(dmsib.ref.Append("service_id"))
}

func (dmsib dataprocMetastoreServiceIamBindingAttributes) Condition() terra.ListValue[dataprocmetastoreserviceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[dataprocmetastoreserviceiambinding.ConditionAttributes](dmsib.ref.Append("condition"))
}

type dataprocMetastoreServiceIamBindingState struct {
	Etag      string                                              `json:"etag"`
	Id        string                                              `json:"id"`
	Location  string                                              `json:"location"`
	Members   []string                                            `json:"members"`
	Project   string                                              `json:"project"`
	Role      string                                              `json:"role"`
	ServiceId string                                              `json:"service_id"`
	Condition []dataprocmetastoreserviceiambinding.ConditionState `json:"condition"`
}

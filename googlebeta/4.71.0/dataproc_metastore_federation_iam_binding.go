// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataprocmetastorefederationiambinding "github.com/golingon/terraproviders/googlebeta/4.71.0/dataprocmetastorefederationiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocMetastoreFederationIamBinding creates a new instance of [DataprocMetastoreFederationIamBinding].
func NewDataprocMetastoreFederationIamBinding(name string, args DataprocMetastoreFederationIamBindingArgs) *DataprocMetastoreFederationIamBinding {
	return &DataprocMetastoreFederationIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocMetastoreFederationIamBinding)(nil)

// DataprocMetastoreFederationIamBinding represents the Terraform resource google_dataproc_metastore_federation_iam_binding.
type DataprocMetastoreFederationIamBinding struct {
	Name      string
	Args      DataprocMetastoreFederationIamBindingArgs
	state     *dataprocMetastoreFederationIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocMetastoreFederationIamBinding].
func (dmfib *DataprocMetastoreFederationIamBinding) Type() string {
	return "google_dataproc_metastore_federation_iam_binding"
}

// LocalName returns the local name for [DataprocMetastoreFederationIamBinding].
func (dmfib *DataprocMetastoreFederationIamBinding) LocalName() string {
	return dmfib.Name
}

// Configuration returns the configuration (args) for [DataprocMetastoreFederationIamBinding].
func (dmfib *DataprocMetastoreFederationIamBinding) Configuration() interface{} {
	return dmfib.Args
}

// DependOn is used for other resources to depend on [DataprocMetastoreFederationIamBinding].
func (dmfib *DataprocMetastoreFederationIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(dmfib)
}

// Dependencies returns the list of resources [DataprocMetastoreFederationIamBinding] depends_on.
func (dmfib *DataprocMetastoreFederationIamBinding) Dependencies() terra.Dependencies {
	return dmfib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocMetastoreFederationIamBinding].
func (dmfib *DataprocMetastoreFederationIamBinding) LifecycleManagement() *terra.Lifecycle {
	return dmfib.Lifecycle
}

// Attributes returns the attributes for [DataprocMetastoreFederationIamBinding].
func (dmfib *DataprocMetastoreFederationIamBinding) Attributes() dataprocMetastoreFederationIamBindingAttributes {
	return dataprocMetastoreFederationIamBindingAttributes{ref: terra.ReferenceResource(dmfib)}
}

// ImportState imports the given attribute values into [DataprocMetastoreFederationIamBinding]'s state.
func (dmfib *DataprocMetastoreFederationIamBinding) ImportState(av io.Reader) error {
	dmfib.state = &dataprocMetastoreFederationIamBindingState{}
	if err := json.NewDecoder(av).Decode(dmfib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmfib.Type(), dmfib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocMetastoreFederationIamBinding] has state.
func (dmfib *DataprocMetastoreFederationIamBinding) State() (*dataprocMetastoreFederationIamBindingState, bool) {
	return dmfib.state, dmfib.state != nil
}

// StateMust returns the state for [DataprocMetastoreFederationIamBinding]. Panics if the state is nil.
func (dmfib *DataprocMetastoreFederationIamBinding) StateMust() *dataprocMetastoreFederationIamBindingState {
	if dmfib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmfib.Type(), dmfib.LocalName()))
	}
	return dmfib.state
}

// DataprocMetastoreFederationIamBindingArgs contains the configurations for google_dataproc_metastore_federation_iam_binding.
type DataprocMetastoreFederationIamBindingArgs struct {
	// FederationId: string, required
	FederationId terra.StringValue `hcl:"federation_id,attr" validate:"required"`
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
	Condition *dataprocmetastorefederationiambinding.Condition `hcl:"condition,block"`
}
type dataprocMetastoreFederationIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_metastore_federation_iam_binding.
func (dmfib dataprocMetastoreFederationIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dmfib.ref.Append("etag"))
}

// FederationId returns a reference to field federation_id of google_dataproc_metastore_federation_iam_binding.
func (dmfib dataprocMetastoreFederationIamBindingAttributes) FederationId() terra.StringValue {
	return terra.ReferenceAsString(dmfib.ref.Append("federation_id"))
}

// Id returns a reference to field id of google_dataproc_metastore_federation_iam_binding.
func (dmfib dataprocMetastoreFederationIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmfib.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_metastore_federation_iam_binding.
func (dmfib dataprocMetastoreFederationIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dmfib.ref.Append("location"))
}

// Members returns a reference to field members of google_dataproc_metastore_federation_iam_binding.
func (dmfib dataprocMetastoreFederationIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dmfib.ref.Append("members"))
}

// Project returns a reference to field project of google_dataproc_metastore_federation_iam_binding.
func (dmfib dataprocMetastoreFederationIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmfib.ref.Append("project"))
}

// Role returns a reference to field role of google_dataproc_metastore_federation_iam_binding.
func (dmfib dataprocMetastoreFederationIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dmfib.ref.Append("role"))
}

func (dmfib dataprocMetastoreFederationIamBindingAttributes) Condition() terra.ListValue[dataprocmetastorefederationiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[dataprocmetastorefederationiambinding.ConditionAttributes](dmfib.ref.Append("condition"))
}

type dataprocMetastoreFederationIamBindingState struct {
	Etag         string                                                 `json:"etag"`
	FederationId string                                                 `json:"federation_id"`
	Id           string                                                 `json:"id"`
	Location     string                                                 `json:"location"`
	Members      []string                                               `json:"members"`
	Project      string                                                 `json:"project"`
	Role         string                                                 `json:"role"`
	Condition    []dataprocmetastorefederationiambinding.ConditionState `json:"condition"`
}

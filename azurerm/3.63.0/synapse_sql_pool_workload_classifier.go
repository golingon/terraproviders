// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapsesqlpoolworkloadclassifier "github.com/golingon/terraproviders/azurerm/3.63.0/synapsesqlpoolworkloadclassifier"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapseSqlPoolWorkloadClassifier creates a new instance of [SynapseSqlPoolWorkloadClassifier].
func NewSynapseSqlPoolWorkloadClassifier(name string, args SynapseSqlPoolWorkloadClassifierArgs) *SynapseSqlPoolWorkloadClassifier {
	return &SynapseSqlPoolWorkloadClassifier{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseSqlPoolWorkloadClassifier)(nil)

// SynapseSqlPoolWorkloadClassifier represents the Terraform resource azurerm_synapse_sql_pool_workload_classifier.
type SynapseSqlPoolWorkloadClassifier struct {
	Name      string
	Args      SynapseSqlPoolWorkloadClassifierArgs
	state     *synapseSqlPoolWorkloadClassifierState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseSqlPoolWorkloadClassifier].
func (sspwc *SynapseSqlPoolWorkloadClassifier) Type() string {
	return "azurerm_synapse_sql_pool_workload_classifier"
}

// LocalName returns the local name for [SynapseSqlPoolWorkloadClassifier].
func (sspwc *SynapseSqlPoolWorkloadClassifier) LocalName() string {
	return sspwc.Name
}

// Configuration returns the configuration (args) for [SynapseSqlPoolWorkloadClassifier].
func (sspwc *SynapseSqlPoolWorkloadClassifier) Configuration() interface{} {
	return sspwc.Args
}

// DependOn is used for other resources to depend on [SynapseSqlPoolWorkloadClassifier].
func (sspwc *SynapseSqlPoolWorkloadClassifier) DependOn() terra.Reference {
	return terra.ReferenceResource(sspwc)
}

// Dependencies returns the list of resources [SynapseSqlPoolWorkloadClassifier] depends_on.
func (sspwc *SynapseSqlPoolWorkloadClassifier) Dependencies() terra.Dependencies {
	return sspwc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseSqlPoolWorkloadClassifier].
func (sspwc *SynapseSqlPoolWorkloadClassifier) LifecycleManagement() *terra.Lifecycle {
	return sspwc.Lifecycle
}

// Attributes returns the attributes for [SynapseSqlPoolWorkloadClassifier].
func (sspwc *SynapseSqlPoolWorkloadClassifier) Attributes() synapseSqlPoolWorkloadClassifierAttributes {
	return synapseSqlPoolWorkloadClassifierAttributes{ref: terra.ReferenceResource(sspwc)}
}

// ImportState imports the given attribute values into [SynapseSqlPoolWorkloadClassifier]'s state.
func (sspwc *SynapseSqlPoolWorkloadClassifier) ImportState(av io.Reader) error {
	sspwc.state = &synapseSqlPoolWorkloadClassifierState{}
	if err := json.NewDecoder(av).Decode(sspwc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sspwc.Type(), sspwc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseSqlPoolWorkloadClassifier] has state.
func (sspwc *SynapseSqlPoolWorkloadClassifier) State() (*synapseSqlPoolWorkloadClassifierState, bool) {
	return sspwc.state, sspwc.state != nil
}

// StateMust returns the state for [SynapseSqlPoolWorkloadClassifier]. Panics if the state is nil.
func (sspwc *SynapseSqlPoolWorkloadClassifier) StateMust() *synapseSqlPoolWorkloadClassifierState {
	if sspwc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sspwc.Type(), sspwc.LocalName()))
	}
	return sspwc.state
}

// SynapseSqlPoolWorkloadClassifierArgs contains the configurations for azurerm_synapse_sql_pool_workload_classifier.
type SynapseSqlPoolWorkloadClassifierArgs struct {
	// Context: string, optional
	Context terra.StringValue `hcl:"context,attr"`
	// EndTime: string, optional
	EndTime terra.StringValue `hcl:"end_time,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Importance: string, optional
	Importance terra.StringValue `hcl:"importance,attr"`
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// MemberName: string, required
	MemberName terra.StringValue `hcl:"member_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
	// WorkloadGroupId: string, required
	WorkloadGroupId terra.StringValue `hcl:"workload_group_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *synapsesqlpoolworkloadclassifier.Timeouts `hcl:"timeouts,block"`
}
type synapseSqlPoolWorkloadClassifierAttributes struct {
	ref terra.Reference
}

// Context returns a reference to field context of azurerm_synapse_sql_pool_workload_classifier.
func (sspwc synapseSqlPoolWorkloadClassifierAttributes) Context() terra.StringValue {
	return terra.ReferenceAsString(sspwc.ref.Append("context"))
}

// EndTime returns a reference to field end_time of azurerm_synapse_sql_pool_workload_classifier.
func (sspwc synapseSqlPoolWorkloadClassifierAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(sspwc.ref.Append("end_time"))
}

// Id returns a reference to field id of azurerm_synapse_sql_pool_workload_classifier.
func (sspwc synapseSqlPoolWorkloadClassifierAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sspwc.ref.Append("id"))
}

// Importance returns a reference to field importance of azurerm_synapse_sql_pool_workload_classifier.
func (sspwc synapseSqlPoolWorkloadClassifierAttributes) Importance() terra.StringValue {
	return terra.ReferenceAsString(sspwc.ref.Append("importance"))
}

// Label returns a reference to field label of azurerm_synapse_sql_pool_workload_classifier.
func (sspwc synapseSqlPoolWorkloadClassifierAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(sspwc.ref.Append("label"))
}

// MemberName returns a reference to field member_name of azurerm_synapse_sql_pool_workload_classifier.
func (sspwc synapseSqlPoolWorkloadClassifierAttributes) MemberName() terra.StringValue {
	return terra.ReferenceAsString(sspwc.ref.Append("member_name"))
}

// Name returns a reference to field name of azurerm_synapse_sql_pool_workload_classifier.
func (sspwc synapseSqlPoolWorkloadClassifierAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sspwc.ref.Append("name"))
}

// StartTime returns a reference to field start_time of azurerm_synapse_sql_pool_workload_classifier.
func (sspwc synapseSqlPoolWorkloadClassifierAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(sspwc.ref.Append("start_time"))
}

// WorkloadGroupId returns a reference to field workload_group_id of azurerm_synapse_sql_pool_workload_classifier.
func (sspwc synapseSqlPoolWorkloadClassifierAttributes) WorkloadGroupId() terra.StringValue {
	return terra.ReferenceAsString(sspwc.ref.Append("workload_group_id"))
}

func (sspwc synapseSqlPoolWorkloadClassifierAttributes) Timeouts() synapsesqlpoolworkloadclassifier.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapsesqlpoolworkloadclassifier.TimeoutsAttributes](sspwc.ref.Append("timeouts"))
}

type synapseSqlPoolWorkloadClassifierState struct {
	Context         string                                          `json:"context"`
	EndTime         string                                          `json:"end_time"`
	Id              string                                          `json:"id"`
	Importance      string                                          `json:"importance"`
	Label           string                                          `json:"label"`
	MemberName      string                                          `json:"member_name"`
	Name            string                                          `json:"name"`
	StartTime       string                                          `json:"start_time"`
	WorkloadGroupId string                                          `json:"workload_group_id"`
	Timeouts        *synapsesqlpoolworkloadclassifier.TimeoutsState `json:"timeouts"`
}

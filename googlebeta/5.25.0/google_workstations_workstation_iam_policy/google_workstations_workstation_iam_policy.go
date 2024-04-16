// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_workstations_workstation_iam_policy

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

// Resource represents the Terraform resource google_workstations_workstation_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleWorkstationsWorkstationIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gwwip *Resource) Type() string {
	return "google_workstations_workstation_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gwwip *Resource) LocalName() string {
	return gwwip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gwwip *Resource) Configuration() interface{} {
	return gwwip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gwwip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gwwip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gwwip *Resource) Dependencies() terra.Dependencies {
	return gwwip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gwwip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gwwip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gwwip *Resource) Attributes() googleWorkstationsWorkstationIamPolicyAttributes {
	return googleWorkstationsWorkstationIamPolicyAttributes{ref: terra.ReferenceResource(gwwip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gwwip *Resource) ImportState(state io.Reader) error {
	gwwip.state = &googleWorkstationsWorkstationIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gwwip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gwwip.Type(), gwwip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gwwip *Resource) State() (*googleWorkstationsWorkstationIamPolicyState, bool) {
	return gwwip.state, gwwip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gwwip *Resource) StateMust() *googleWorkstationsWorkstationIamPolicyState {
	if gwwip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gwwip.Type(), gwwip.LocalName()))
	}
	return gwwip.state
}

// Args contains the configurations for google_workstations_workstation_iam_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// WorkstationClusterId: string, required
	WorkstationClusterId terra.StringValue `hcl:"workstation_cluster_id,attr" validate:"required"`
	// WorkstationConfigId: string, required
	WorkstationConfigId terra.StringValue `hcl:"workstation_config_id,attr" validate:"required"`
	// WorkstationId: string, required
	WorkstationId terra.StringValue `hcl:"workstation_id,attr" validate:"required"`
}

type googleWorkstationsWorkstationIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_workstations_workstation_iam_policy.
func (gwwip googleWorkstationsWorkstationIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gwwip.ref.Append("etag"))
}

// Id returns a reference to field id of google_workstations_workstation_iam_policy.
func (gwwip googleWorkstationsWorkstationIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gwwip.ref.Append("id"))
}

// Location returns a reference to field location of google_workstations_workstation_iam_policy.
func (gwwip googleWorkstationsWorkstationIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gwwip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_workstations_workstation_iam_policy.
func (gwwip googleWorkstationsWorkstationIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gwwip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_workstations_workstation_iam_policy.
func (gwwip googleWorkstationsWorkstationIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gwwip.ref.Append("project"))
}

// WorkstationClusterId returns a reference to field workstation_cluster_id of google_workstations_workstation_iam_policy.
func (gwwip googleWorkstationsWorkstationIamPolicyAttributes) WorkstationClusterId() terra.StringValue {
	return terra.ReferenceAsString(gwwip.ref.Append("workstation_cluster_id"))
}

// WorkstationConfigId returns a reference to field workstation_config_id of google_workstations_workstation_iam_policy.
func (gwwip googleWorkstationsWorkstationIamPolicyAttributes) WorkstationConfigId() terra.StringValue {
	return terra.ReferenceAsString(gwwip.ref.Append("workstation_config_id"))
}

// WorkstationId returns a reference to field workstation_id of google_workstations_workstation_iam_policy.
func (gwwip googleWorkstationsWorkstationIamPolicyAttributes) WorkstationId() terra.StringValue {
	return terra.ReferenceAsString(gwwip.ref.Append("workstation_id"))
}

type googleWorkstationsWorkstationIamPolicyState struct {
	Etag                 string `json:"etag"`
	Id                   string `json:"id"`
	Location             string `json:"location"`
	PolicyData           string `json:"policy_data"`
	Project              string `json:"project"`
	WorkstationClusterId string `json:"workstation_cluster_id"`
	WorkstationConfigId  string `json:"workstation_config_id"`
	WorkstationId        string `json:"workstation_id"`
}

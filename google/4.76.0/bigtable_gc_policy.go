// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigtablegcpolicy "github.com/golingon/terraproviders/google/4.76.0/bigtablegcpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigtableGcPolicy creates a new instance of [BigtableGcPolicy].
func NewBigtableGcPolicy(name string, args BigtableGcPolicyArgs) *BigtableGcPolicy {
	return &BigtableGcPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableGcPolicy)(nil)

// BigtableGcPolicy represents the Terraform resource google_bigtable_gc_policy.
type BigtableGcPolicy struct {
	Name      string
	Args      BigtableGcPolicyArgs
	state     *bigtableGcPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigtableGcPolicy].
func (bgp *BigtableGcPolicy) Type() string {
	return "google_bigtable_gc_policy"
}

// LocalName returns the local name for [BigtableGcPolicy].
func (bgp *BigtableGcPolicy) LocalName() string {
	return bgp.Name
}

// Configuration returns the configuration (args) for [BigtableGcPolicy].
func (bgp *BigtableGcPolicy) Configuration() interface{} {
	return bgp.Args
}

// DependOn is used for other resources to depend on [BigtableGcPolicy].
func (bgp *BigtableGcPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(bgp)
}

// Dependencies returns the list of resources [BigtableGcPolicy] depends_on.
func (bgp *BigtableGcPolicy) Dependencies() terra.Dependencies {
	return bgp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigtableGcPolicy].
func (bgp *BigtableGcPolicy) LifecycleManagement() *terra.Lifecycle {
	return bgp.Lifecycle
}

// Attributes returns the attributes for [BigtableGcPolicy].
func (bgp *BigtableGcPolicy) Attributes() bigtableGcPolicyAttributes {
	return bigtableGcPolicyAttributes{ref: terra.ReferenceResource(bgp)}
}

// ImportState imports the given attribute values into [BigtableGcPolicy]'s state.
func (bgp *BigtableGcPolicy) ImportState(av io.Reader) error {
	bgp.state = &bigtableGcPolicyState{}
	if err := json.NewDecoder(av).Decode(bgp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bgp.Type(), bgp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigtableGcPolicy] has state.
func (bgp *BigtableGcPolicy) State() (*bigtableGcPolicyState, bool) {
	return bgp.state, bgp.state != nil
}

// StateMust returns the state for [BigtableGcPolicy]. Panics if the state is nil.
func (bgp *BigtableGcPolicy) StateMust() *bigtableGcPolicyState {
	if bgp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bgp.Type(), bgp.LocalName()))
	}
	return bgp.state
}

// BigtableGcPolicyArgs contains the configurations for google_bigtable_gc_policy.
type BigtableGcPolicyArgs struct {
	// ColumnFamily: string, required
	ColumnFamily terra.StringValue `hcl:"column_family,attr" validate:"required"`
	// DeletionPolicy: string, optional
	DeletionPolicy terra.StringValue `hcl:"deletion_policy,attr"`
	// GcRules: string, optional
	GcRules terra.StringValue `hcl:"gc_rules,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Table: string, required
	Table terra.StringValue `hcl:"table,attr" validate:"required"`
	// MaxAge: optional
	MaxAge *bigtablegcpolicy.MaxAge `hcl:"max_age,block"`
	// MaxVersion: min=0
	MaxVersion []bigtablegcpolicy.MaxVersion `hcl:"max_version,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *bigtablegcpolicy.Timeouts `hcl:"timeouts,block"`
}
type bigtableGcPolicyAttributes struct {
	ref terra.Reference
}

// ColumnFamily returns a reference to field column_family of google_bigtable_gc_policy.
func (bgp bigtableGcPolicyAttributes) ColumnFamily() terra.StringValue {
	return terra.ReferenceAsString(bgp.ref.Append("column_family"))
}

// DeletionPolicy returns a reference to field deletion_policy of google_bigtable_gc_policy.
func (bgp bigtableGcPolicyAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(bgp.ref.Append("deletion_policy"))
}

// GcRules returns a reference to field gc_rules of google_bigtable_gc_policy.
func (bgp bigtableGcPolicyAttributes) GcRules() terra.StringValue {
	return terra.ReferenceAsString(bgp.ref.Append("gc_rules"))
}

// Id returns a reference to field id of google_bigtable_gc_policy.
func (bgp bigtableGcPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bgp.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of google_bigtable_gc_policy.
func (bgp bigtableGcPolicyAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(bgp.ref.Append("instance_name"))
}

// Mode returns a reference to field mode of google_bigtable_gc_policy.
func (bgp bigtableGcPolicyAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(bgp.ref.Append("mode"))
}

// Project returns a reference to field project of google_bigtable_gc_policy.
func (bgp bigtableGcPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bgp.ref.Append("project"))
}

// Table returns a reference to field table of google_bigtable_gc_policy.
func (bgp bigtableGcPolicyAttributes) Table() terra.StringValue {
	return terra.ReferenceAsString(bgp.ref.Append("table"))
}

func (bgp bigtableGcPolicyAttributes) MaxAge() terra.ListValue[bigtablegcpolicy.MaxAgeAttributes] {
	return terra.ReferenceAsList[bigtablegcpolicy.MaxAgeAttributes](bgp.ref.Append("max_age"))
}

func (bgp bigtableGcPolicyAttributes) MaxVersion() terra.ListValue[bigtablegcpolicy.MaxVersionAttributes] {
	return terra.ReferenceAsList[bigtablegcpolicy.MaxVersionAttributes](bgp.ref.Append("max_version"))
}

func (bgp bigtableGcPolicyAttributes) Timeouts() bigtablegcpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigtablegcpolicy.TimeoutsAttributes](bgp.ref.Append("timeouts"))
}

type bigtableGcPolicyState struct {
	ColumnFamily   string                             `json:"column_family"`
	DeletionPolicy string                             `json:"deletion_policy"`
	GcRules        string                             `json:"gc_rules"`
	Id             string                             `json:"id"`
	InstanceName   string                             `json:"instance_name"`
	Mode           string                             `json:"mode"`
	Project        string                             `json:"project"`
	Table          string                             `json:"table"`
	MaxAge         []bigtablegcpolicy.MaxAgeState     `json:"max_age"`
	MaxVersion     []bigtablegcpolicy.MaxVersionState `json:"max_version"`
	Timeouts       *bigtablegcpolicy.TimeoutsState    `json:"timeouts"`
}

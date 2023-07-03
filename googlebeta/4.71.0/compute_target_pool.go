// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computetargetpool "github.com/golingon/terraproviders/googlebeta/4.71.0/computetargetpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeTargetPool creates a new instance of [ComputeTargetPool].
func NewComputeTargetPool(name string, args ComputeTargetPoolArgs) *ComputeTargetPool {
	return &ComputeTargetPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeTargetPool)(nil)

// ComputeTargetPool represents the Terraform resource google_compute_target_pool.
type ComputeTargetPool struct {
	Name      string
	Args      ComputeTargetPoolArgs
	state     *computeTargetPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeTargetPool].
func (ctp *ComputeTargetPool) Type() string {
	return "google_compute_target_pool"
}

// LocalName returns the local name for [ComputeTargetPool].
func (ctp *ComputeTargetPool) LocalName() string {
	return ctp.Name
}

// Configuration returns the configuration (args) for [ComputeTargetPool].
func (ctp *ComputeTargetPool) Configuration() interface{} {
	return ctp.Args
}

// DependOn is used for other resources to depend on [ComputeTargetPool].
func (ctp *ComputeTargetPool) DependOn() terra.Reference {
	return terra.ReferenceResource(ctp)
}

// Dependencies returns the list of resources [ComputeTargetPool] depends_on.
func (ctp *ComputeTargetPool) Dependencies() terra.Dependencies {
	return ctp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeTargetPool].
func (ctp *ComputeTargetPool) LifecycleManagement() *terra.Lifecycle {
	return ctp.Lifecycle
}

// Attributes returns the attributes for [ComputeTargetPool].
func (ctp *ComputeTargetPool) Attributes() computeTargetPoolAttributes {
	return computeTargetPoolAttributes{ref: terra.ReferenceResource(ctp)}
}

// ImportState imports the given attribute values into [ComputeTargetPool]'s state.
func (ctp *ComputeTargetPool) ImportState(av io.Reader) error {
	ctp.state = &computeTargetPoolState{}
	if err := json.NewDecoder(av).Decode(ctp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ctp.Type(), ctp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeTargetPool] has state.
func (ctp *ComputeTargetPool) State() (*computeTargetPoolState, bool) {
	return ctp.state, ctp.state != nil
}

// StateMust returns the state for [ComputeTargetPool]. Panics if the state is nil.
func (ctp *ComputeTargetPool) StateMust() *computeTargetPoolState {
	if ctp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ctp.Type(), ctp.LocalName()))
	}
	return ctp.state
}

// ComputeTargetPoolArgs contains the configurations for google_compute_target_pool.
type ComputeTargetPoolArgs struct {
	// BackupPool: string, optional
	BackupPool terra.StringValue `hcl:"backup_pool,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FailoverRatio: number, optional
	FailoverRatio terra.NumberValue `hcl:"failover_ratio,attr"`
	// HealthChecks: list of string, optional
	HealthChecks terra.ListValue[terra.StringValue] `hcl:"health_checks,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instances: set of string, optional
	Instances terra.SetValue[terra.StringValue] `hcl:"instances,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// SessionAffinity: string, optional
	SessionAffinity terra.StringValue `hcl:"session_affinity,attr"`
	// Timeouts: optional
	Timeouts *computetargetpool.Timeouts `hcl:"timeouts,block"`
}
type computeTargetPoolAttributes struct {
	ref terra.Reference
}

// BackupPool returns a reference to field backup_pool of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) BackupPool() terra.StringValue {
	return terra.ReferenceAsString(ctp.ref.Append("backup_pool"))
}

// Description returns a reference to field description of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ctp.ref.Append("description"))
}

// FailoverRatio returns a reference to field failover_ratio of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) FailoverRatio() terra.NumberValue {
	return terra.ReferenceAsNumber(ctp.ref.Append("failover_ratio"))
}

// HealthChecks returns a reference to field health_checks of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) HealthChecks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ctp.ref.Append("health_checks"))
}

// Id returns a reference to field id of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ctp.ref.Append("id"))
}

// Instances returns a reference to field instances of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) Instances() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ctp.ref.Append("instances"))
}

// Name returns a reference to field name of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ctp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ctp.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ctp.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(ctp.ref.Append("self_link"))
}

// SessionAffinity returns a reference to field session_affinity of google_compute_target_pool.
func (ctp computeTargetPoolAttributes) SessionAffinity() terra.StringValue {
	return terra.ReferenceAsString(ctp.ref.Append("session_affinity"))
}

func (ctp computeTargetPoolAttributes) Timeouts() computetargetpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computetargetpool.TimeoutsAttributes](ctp.ref.Append("timeouts"))
}

type computeTargetPoolState struct {
	BackupPool      string                           `json:"backup_pool"`
	Description     string                           `json:"description"`
	FailoverRatio   float64                          `json:"failover_ratio"`
	HealthChecks    []string                         `json:"health_checks"`
	Id              string                           `json:"id"`
	Instances       []string                         `json:"instances"`
	Name            string                           `json:"name"`
	Project         string                           `json:"project"`
	Region          string                           `json:"region"`
	SelfLink        string                           `json:"self_link"`
	SessionAffinity string                           `json:"session_affinity"`
	Timeouts        *computetargetpool.TimeoutsState `json:"timeouts"`
}

// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregioncommitment "github.com/golingon/terraproviders/googlebeta/4.74.0/computeregioncommitment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionCommitment creates a new instance of [ComputeRegionCommitment].
func NewComputeRegionCommitment(name string, args ComputeRegionCommitmentArgs) *ComputeRegionCommitment {
	return &ComputeRegionCommitment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionCommitment)(nil)

// ComputeRegionCommitment represents the Terraform resource google_compute_region_commitment.
type ComputeRegionCommitment struct {
	Name      string
	Args      ComputeRegionCommitmentArgs
	state     *computeRegionCommitmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionCommitment].
func (crc *ComputeRegionCommitment) Type() string {
	return "google_compute_region_commitment"
}

// LocalName returns the local name for [ComputeRegionCommitment].
func (crc *ComputeRegionCommitment) LocalName() string {
	return crc.Name
}

// Configuration returns the configuration (args) for [ComputeRegionCommitment].
func (crc *ComputeRegionCommitment) Configuration() interface{} {
	return crc.Args
}

// DependOn is used for other resources to depend on [ComputeRegionCommitment].
func (crc *ComputeRegionCommitment) DependOn() terra.Reference {
	return terra.ReferenceResource(crc)
}

// Dependencies returns the list of resources [ComputeRegionCommitment] depends_on.
func (crc *ComputeRegionCommitment) Dependencies() terra.Dependencies {
	return crc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionCommitment].
func (crc *ComputeRegionCommitment) LifecycleManagement() *terra.Lifecycle {
	return crc.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionCommitment].
func (crc *ComputeRegionCommitment) Attributes() computeRegionCommitmentAttributes {
	return computeRegionCommitmentAttributes{ref: terra.ReferenceResource(crc)}
}

// ImportState imports the given attribute values into [ComputeRegionCommitment]'s state.
func (crc *ComputeRegionCommitment) ImportState(av io.Reader) error {
	crc.state = &computeRegionCommitmentState{}
	if err := json.NewDecoder(av).Decode(crc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crc.Type(), crc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionCommitment] has state.
func (crc *ComputeRegionCommitment) State() (*computeRegionCommitmentState, bool) {
	return crc.state, crc.state != nil
}

// StateMust returns the state for [ComputeRegionCommitment]. Panics if the state is nil.
func (crc *ComputeRegionCommitment) StateMust() *computeRegionCommitmentState {
	if crc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crc.Type(), crc.LocalName()))
	}
	return crc.state
}

// ComputeRegionCommitmentArgs contains the configurations for google_compute_region_commitment.
type ComputeRegionCommitmentArgs struct {
	// AutoRenew: bool, optional
	AutoRenew terra.BoolValue `hcl:"auto_renew,attr"`
	// Category: string, optional
	Category terra.StringValue `hcl:"category,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Plan: string, required
	Plan terra.StringValue `hcl:"plan,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// LicenseResource: optional
	LicenseResource *computeregioncommitment.LicenseResource `hcl:"license_resource,block"`
	// Resources: min=0
	Resources []computeregioncommitment.Resources `hcl:"resources,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeregioncommitment.Timeouts `hcl:"timeouts,block"`
}
type computeRegionCommitmentAttributes struct {
	ref terra.Reference
}

// AutoRenew returns a reference to field auto_renew of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) AutoRenew() terra.BoolValue {
	return terra.ReferenceAsBool(crc.ref.Append("auto_renew"))
}

// Category returns a reference to field category of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) Category() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("category"))
}

// CommitmentId returns a reference to field commitment_id of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) CommitmentId() terra.NumberValue {
	return terra.ReferenceAsNumber(crc.ref.Append("commitment_id"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("description"))
}

// EndTimestamp returns a reference to field end_timestamp of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) EndTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("end_timestamp"))
}

// Id returns a reference to field id of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("name"))
}

// Plan returns a reference to field plan of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) Plan() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("plan"))
}

// Project returns a reference to field project of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("self_link"))
}

// StartTimestamp returns a reference to field start_timestamp of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) StartTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("start_timestamp"))
}

// Status returns a reference to field status of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("status"))
}

// StatusMessage returns a reference to field status_message of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) StatusMessage() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("status_message"))
}

// Type returns a reference to field type of google_compute_region_commitment.
func (crc computeRegionCommitmentAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("type"))
}

func (crc computeRegionCommitmentAttributes) LicenseResource() terra.ListValue[computeregioncommitment.LicenseResourceAttributes] {
	return terra.ReferenceAsList[computeregioncommitment.LicenseResourceAttributes](crc.ref.Append("license_resource"))
}

func (crc computeRegionCommitmentAttributes) Resources() terra.ListValue[computeregioncommitment.ResourcesAttributes] {
	return terra.ReferenceAsList[computeregioncommitment.ResourcesAttributes](crc.ref.Append("resources"))
}

func (crc computeRegionCommitmentAttributes) Timeouts() computeregioncommitment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregioncommitment.TimeoutsAttributes](crc.ref.Append("timeouts"))
}

type computeRegionCommitmentState struct {
	AutoRenew         bool                                           `json:"auto_renew"`
	Category          string                                         `json:"category"`
	CommitmentId      float64                                        `json:"commitment_id"`
	CreationTimestamp string                                         `json:"creation_timestamp"`
	Description       string                                         `json:"description"`
	EndTimestamp      string                                         `json:"end_timestamp"`
	Id                string                                         `json:"id"`
	Name              string                                         `json:"name"`
	Plan              string                                         `json:"plan"`
	Project           string                                         `json:"project"`
	Region            string                                         `json:"region"`
	SelfLink          string                                         `json:"self_link"`
	StartTimestamp    string                                         `json:"start_timestamp"`
	Status            string                                         `json:"status"`
	StatusMessage     string                                         `json:"status_message"`
	Type              string                                         `json:"type"`
	LicenseResource   []computeregioncommitment.LicenseResourceState `json:"license_resource"`
	Resources         []computeregioncommitment.ResourcesState       `json:"resources"`
	Timeouts          *computeregioncommitment.TimeoutsState         `json:"timeouts"`
}

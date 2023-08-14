// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionDiskIamPolicy creates a new instance of [ComputeRegionDiskIamPolicy].
func NewComputeRegionDiskIamPolicy(name string, args ComputeRegionDiskIamPolicyArgs) *ComputeRegionDiskIamPolicy {
	return &ComputeRegionDiskIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionDiskIamPolicy)(nil)

// ComputeRegionDiskIamPolicy represents the Terraform resource google_compute_region_disk_iam_policy.
type ComputeRegionDiskIamPolicy struct {
	Name      string
	Args      ComputeRegionDiskIamPolicyArgs
	state     *computeRegionDiskIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionDiskIamPolicy].
func (crdip *ComputeRegionDiskIamPolicy) Type() string {
	return "google_compute_region_disk_iam_policy"
}

// LocalName returns the local name for [ComputeRegionDiskIamPolicy].
func (crdip *ComputeRegionDiskIamPolicy) LocalName() string {
	return crdip.Name
}

// Configuration returns the configuration (args) for [ComputeRegionDiskIamPolicy].
func (crdip *ComputeRegionDiskIamPolicy) Configuration() interface{} {
	return crdip.Args
}

// DependOn is used for other resources to depend on [ComputeRegionDiskIamPolicy].
func (crdip *ComputeRegionDiskIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(crdip)
}

// Dependencies returns the list of resources [ComputeRegionDiskIamPolicy] depends_on.
func (crdip *ComputeRegionDiskIamPolicy) Dependencies() terra.Dependencies {
	return crdip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionDiskIamPolicy].
func (crdip *ComputeRegionDiskIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return crdip.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionDiskIamPolicy].
func (crdip *ComputeRegionDiskIamPolicy) Attributes() computeRegionDiskIamPolicyAttributes {
	return computeRegionDiskIamPolicyAttributes{ref: terra.ReferenceResource(crdip)}
}

// ImportState imports the given attribute values into [ComputeRegionDiskIamPolicy]'s state.
func (crdip *ComputeRegionDiskIamPolicy) ImportState(av io.Reader) error {
	crdip.state = &computeRegionDiskIamPolicyState{}
	if err := json.NewDecoder(av).Decode(crdip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crdip.Type(), crdip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionDiskIamPolicy] has state.
func (crdip *ComputeRegionDiskIamPolicy) State() (*computeRegionDiskIamPolicyState, bool) {
	return crdip.state, crdip.state != nil
}

// StateMust returns the state for [ComputeRegionDiskIamPolicy]. Panics if the state is nil.
func (crdip *ComputeRegionDiskIamPolicy) StateMust() *computeRegionDiskIamPolicyState {
	if crdip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crdip.Type(), crdip.LocalName()))
	}
	return crdip.state
}

// ComputeRegionDiskIamPolicyArgs contains the configurations for google_compute_region_disk_iam_policy.
type ComputeRegionDiskIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type computeRegionDiskIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_region_disk_iam_policy.
func (crdip computeRegionDiskIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crdip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_region_disk_iam_policy.
func (crdip computeRegionDiskIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crdip.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_disk_iam_policy.
func (crdip computeRegionDiskIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crdip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_compute_region_disk_iam_policy.
func (crdip computeRegionDiskIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(crdip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_region_disk_iam_policy.
func (crdip computeRegionDiskIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crdip.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_disk_iam_policy.
func (crdip computeRegionDiskIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crdip.ref.Append("region"))
}

type computeRegionDiskIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Region     string `json:"region"`
}
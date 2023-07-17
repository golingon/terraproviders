// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudhsmv2hsm "github.com/golingon/terraproviders/aws/5.8.0/cloudhsmv2hsm"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudhsmV2Hsm creates a new instance of [CloudhsmV2Hsm].
func NewCloudhsmV2Hsm(name string, args CloudhsmV2HsmArgs) *CloudhsmV2Hsm {
	return &CloudhsmV2Hsm{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudhsmV2Hsm)(nil)

// CloudhsmV2Hsm represents the Terraform resource aws_cloudhsm_v2_hsm.
type CloudhsmV2Hsm struct {
	Name      string
	Args      CloudhsmV2HsmArgs
	state     *cloudhsmV2HsmState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudhsmV2Hsm].
func (cvh *CloudhsmV2Hsm) Type() string {
	return "aws_cloudhsm_v2_hsm"
}

// LocalName returns the local name for [CloudhsmV2Hsm].
func (cvh *CloudhsmV2Hsm) LocalName() string {
	return cvh.Name
}

// Configuration returns the configuration (args) for [CloudhsmV2Hsm].
func (cvh *CloudhsmV2Hsm) Configuration() interface{} {
	return cvh.Args
}

// DependOn is used for other resources to depend on [CloudhsmV2Hsm].
func (cvh *CloudhsmV2Hsm) DependOn() terra.Reference {
	return terra.ReferenceResource(cvh)
}

// Dependencies returns the list of resources [CloudhsmV2Hsm] depends_on.
func (cvh *CloudhsmV2Hsm) Dependencies() terra.Dependencies {
	return cvh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudhsmV2Hsm].
func (cvh *CloudhsmV2Hsm) LifecycleManagement() *terra.Lifecycle {
	return cvh.Lifecycle
}

// Attributes returns the attributes for [CloudhsmV2Hsm].
func (cvh *CloudhsmV2Hsm) Attributes() cloudhsmV2HsmAttributes {
	return cloudhsmV2HsmAttributes{ref: terra.ReferenceResource(cvh)}
}

// ImportState imports the given attribute values into [CloudhsmV2Hsm]'s state.
func (cvh *CloudhsmV2Hsm) ImportState(av io.Reader) error {
	cvh.state = &cloudhsmV2HsmState{}
	if err := json.NewDecoder(av).Decode(cvh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cvh.Type(), cvh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudhsmV2Hsm] has state.
func (cvh *CloudhsmV2Hsm) State() (*cloudhsmV2HsmState, bool) {
	return cvh.state, cvh.state != nil
}

// StateMust returns the state for [CloudhsmV2Hsm]. Panics if the state is nil.
func (cvh *CloudhsmV2Hsm) StateMust() *cloudhsmV2HsmState {
	if cvh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cvh.Type(), cvh.LocalName()))
	}
	return cvh.state
}

// CloudhsmV2HsmArgs contains the configurations for aws_cloudhsm_v2_hsm.
type CloudhsmV2HsmArgs struct {
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Timeouts: optional
	Timeouts *cloudhsmv2hsm.Timeouts `hcl:"timeouts,block"`
}
type cloudhsmV2HsmAttributes struct {
	ref terra.Reference
}

// AvailabilityZone returns a reference to field availability_zone of aws_cloudhsm_v2_hsm.
func (cvh cloudhsmV2HsmAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(cvh.ref.Append("availability_zone"))
}

// ClusterId returns a reference to field cluster_id of aws_cloudhsm_v2_hsm.
func (cvh cloudhsmV2HsmAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(cvh.ref.Append("cluster_id"))
}

// HsmEniId returns a reference to field hsm_eni_id of aws_cloudhsm_v2_hsm.
func (cvh cloudhsmV2HsmAttributes) HsmEniId() terra.StringValue {
	return terra.ReferenceAsString(cvh.ref.Append("hsm_eni_id"))
}

// HsmId returns a reference to field hsm_id of aws_cloudhsm_v2_hsm.
func (cvh cloudhsmV2HsmAttributes) HsmId() terra.StringValue {
	return terra.ReferenceAsString(cvh.ref.Append("hsm_id"))
}

// HsmState returns a reference to field hsm_state of aws_cloudhsm_v2_hsm.
func (cvh cloudhsmV2HsmAttributes) HsmState() terra.StringValue {
	return terra.ReferenceAsString(cvh.ref.Append("hsm_state"))
}

// Id returns a reference to field id of aws_cloudhsm_v2_hsm.
func (cvh cloudhsmV2HsmAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvh.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of aws_cloudhsm_v2_hsm.
func (cvh cloudhsmV2HsmAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(cvh.ref.Append("ip_address"))
}

// SubnetId returns a reference to field subnet_id of aws_cloudhsm_v2_hsm.
func (cvh cloudhsmV2HsmAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(cvh.ref.Append("subnet_id"))
}

func (cvh cloudhsmV2HsmAttributes) Timeouts() cloudhsmv2hsm.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudhsmv2hsm.TimeoutsAttributes](cvh.ref.Append("timeouts"))
}

type cloudhsmV2HsmState struct {
	AvailabilityZone string                       `json:"availability_zone"`
	ClusterId        string                       `json:"cluster_id"`
	HsmEniId         string                       `json:"hsm_eni_id"`
	HsmId            string                       `json:"hsm_id"`
	HsmState         string                       `json:"hsm_state"`
	Id               string                       `json:"id"`
	IpAddress        string                       `json:"ip_address"`
	SubnetId         string                       `json:"subnet_id"`
	Timeouts         *cloudhsmv2hsm.TimeoutsState `json:"timeouts"`
}

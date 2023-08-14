// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	emrinstancegroup "github.com/golingon/terraproviders/aws/5.12.0/emrinstancegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEmrInstanceGroup creates a new instance of [EmrInstanceGroup].
func NewEmrInstanceGroup(name string, args EmrInstanceGroupArgs) *EmrInstanceGroup {
	return &EmrInstanceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EmrInstanceGroup)(nil)

// EmrInstanceGroup represents the Terraform resource aws_emr_instance_group.
type EmrInstanceGroup struct {
	Name      string
	Args      EmrInstanceGroupArgs
	state     *emrInstanceGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EmrInstanceGroup].
func (eig *EmrInstanceGroup) Type() string {
	return "aws_emr_instance_group"
}

// LocalName returns the local name for [EmrInstanceGroup].
func (eig *EmrInstanceGroup) LocalName() string {
	return eig.Name
}

// Configuration returns the configuration (args) for [EmrInstanceGroup].
func (eig *EmrInstanceGroup) Configuration() interface{} {
	return eig.Args
}

// DependOn is used for other resources to depend on [EmrInstanceGroup].
func (eig *EmrInstanceGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(eig)
}

// Dependencies returns the list of resources [EmrInstanceGroup] depends_on.
func (eig *EmrInstanceGroup) Dependencies() terra.Dependencies {
	return eig.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EmrInstanceGroup].
func (eig *EmrInstanceGroup) LifecycleManagement() *terra.Lifecycle {
	return eig.Lifecycle
}

// Attributes returns the attributes for [EmrInstanceGroup].
func (eig *EmrInstanceGroup) Attributes() emrInstanceGroupAttributes {
	return emrInstanceGroupAttributes{ref: terra.ReferenceResource(eig)}
}

// ImportState imports the given attribute values into [EmrInstanceGroup]'s state.
func (eig *EmrInstanceGroup) ImportState(av io.Reader) error {
	eig.state = &emrInstanceGroupState{}
	if err := json.NewDecoder(av).Decode(eig.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eig.Type(), eig.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EmrInstanceGroup] has state.
func (eig *EmrInstanceGroup) State() (*emrInstanceGroupState, bool) {
	return eig.state, eig.state != nil
}

// StateMust returns the state for [EmrInstanceGroup]. Panics if the state is nil.
func (eig *EmrInstanceGroup) StateMust() *emrInstanceGroupState {
	if eig.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eig.Type(), eig.LocalName()))
	}
	return eig.state
}

// EmrInstanceGroupArgs contains the configurations for aws_emr_instance_group.
type EmrInstanceGroupArgs struct {
	// AutoscalingPolicy: string, optional
	AutoscalingPolicy terra.StringValue `hcl:"autoscaling_policy,attr"`
	// BidPrice: string, optional
	BidPrice terra.StringValue `hcl:"bid_price,attr"`
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// ConfigurationsJson: string, optional
	ConfigurationsJson terra.StringValue `hcl:"configurations_json,attr"`
	// EbsOptimized: bool, optional
	EbsOptimized terra.BoolValue `hcl:"ebs_optimized,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceCount: number, optional
	InstanceCount terra.NumberValue `hcl:"instance_count,attr"`
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// EbsConfig: min=0
	EbsConfig []emrinstancegroup.EbsConfig `hcl:"ebs_config,block" validate:"min=0"`
}
type emrInstanceGroupAttributes struct {
	ref terra.Reference
}

// AutoscalingPolicy returns a reference to field autoscaling_policy of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) AutoscalingPolicy() terra.StringValue {
	return terra.ReferenceAsString(eig.ref.Append("autoscaling_policy"))
}

// BidPrice returns a reference to field bid_price of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) BidPrice() terra.StringValue {
	return terra.ReferenceAsString(eig.ref.Append("bid_price"))
}

// ClusterId returns a reference to field cluster_id of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(eig.ref.Append("cluster_id"))
}

// ConfigurationsJson returns a reference to field configurations_json of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) ConfigurationsJson() terra.StringValue {
	return terra.ReferenceAsString(eig.ref.Append("configurations_json"))
}

// EbsOptimized returns a reference to field ebs_optimized of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) EbsOptimized() terra.BoolValue {
	return terra.ReferenceAsBool(eig.ref.Append("ebs_optimized"))
}

// Id returns a reference to field id of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eig.ref.Append("id"))
}

// InstanceCount returns a reference to field instance_count of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(eig.ref.Append("instance_count"))
}

// InstanceType returns a reference to field instance_type of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(eig.ref.Append("instance_type"))
}

// Name returns a reference to field name of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(eig.ref.Append("name"))
}

// RunningInstanceCount returns a reference to field running_instance_count of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) RunningInstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(eig.ref.Append("running_instance_count"))
}

// Status returns a reference to field status of aws_emr_instance_group.
func (eig emrInstanceGroupAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(eig.ref.Append("status"))
}

func (eig emrInstanceGroupAttributes) EbsConfig() terra.SetValue[emrinstancegroup.EbsConfigAttributes] {
	return terra.ReferenceAsSet[emrinstancegroup.EbsConfigAttributes](eig.ref.Append("ebs_config"))
}

type emrInstanceGroupState struct {
	AutoscalingPolicy    string                            `json:"autoscaling_policy"`
	BidPrice             string                            `json:"bid_price"`
	ClusterId            string                            `json:"cluster_id"`
	ConfigurationsJson   string                            `json:"configurations_json"`
	EbsOptimized         bool                              `json:"ebs_optimized"`
	Id                   string                            `json:"id"`
	InstanceCount        float64                           `json:"instance_count"`
	InstanceType         string                            `json:"instance_type"`
	Name                 string                            `json:"name"`
	RunningInstanceCount float64                           `json:"running_instance_count"`
	Status               string                            `json:"status"`
	EbsConfig            []emrinstancegroup.EbsConfigState `json:"ebs_config"`
}

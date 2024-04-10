// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	clouddeploytarget "github.com/golingon/terraproviders/google/5.24.0/clouddeploytarget"
	"io"
)

// NewClouddeployTarget creates a new instance of [ClouddeployTarget].
func NewClouddeployTarget(name string, args ClouddeployTargetArgs) *ClouddeployTarget {
	return &ClouddeployTarget{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ClouddeployTarget)(nil)

// ClouddeployTarget represents the Terraform resource google_clouddeploy_target.
type ClouddeployTarget struct {
	Name      string
	Args      ClouddeployTargetArgs
	state     *clouddeployTargetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ClouddeployTarget].
func (ct *ClouddeployTarget) Type() string {
	return "google_clouddeploy_target"
}

// LocalName returns the local name for [ClouddeployTarget].
func (ct *ClouddeployTarget) LocalName() string {
	return ct.Name
}

// Configuration returns the configuration (args) for [ClouddeployTarget].
func (ct *ClouddeployTarget) Configuration() interface{} {
	return ct.Args
}

// DependOn is used for other resources to depend on [ClouddeployTarget].
func (ct *ClouddeployTarget) DependOn() terra.Reference {
	return terra.ReferenceResource(ct)
}

// Dependencies returns the list of resources [ClouddeployTarget] depends_on.
func (ct *ClouddeployTarget) Dependencies() terra.Dependencies {
	return ct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ClouddeployTarget].
func (ct *ClouddeployTarget) LifecycleManagement() *terra.Lifecycle {
	return ct.Lifecycle
}

// Attributes returns the attributes for [ClouddeployTarget].
func (ct *ClouddeployTarget) Attributes() clouddeployTargetAttributes {
	return clouddeployTargetAttributes{ref: terra.ReferenceResource(ct)}
}

// ImportState imports the given attribute values into [ClouddeployTarget]'s state.
func (ct *ClouddeployTarget) ImportState(av io.Reader) error {
	ct.state = &clouddeployTargetState{}
	if err := json.NewDecoder(av).Decode(ct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ct.Type(), ct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ClouddeployTarget] has state.
func (ct *ClouddeployTarget) State() (*clouddeployTargetState, bool) {
	return ct.state, ct.state != nil
}

// StateMust returns the state for [ClouddeployTarget]. Panics if the state is nil.
func (ct *ClouddeployTarget) StateMust() *clouddeployTargetState {
	if ct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ct.Type(), ct.LocalName()))
	}
	return ct.state
}

// ClouddeployTargetArgs contains the configurations for google_clouddeploy_target.
type ClouddeployTargetArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// DeployParameters: map of string, optional
	DeployParameters terra.MapValue[terra.StringValue] `hcl:"deploy_parameters,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RequireApproval: bool, optional
	RequireApproval terra.BoolValue `hcl:"require_approval,attr"`
	// AnthosCluster: optional
	AnthosCluster *clouddeploytarget.AnthosCluster `hcl:"anthos_cluster,block"`
	// ExecutionConfigs: min=0
	ExecutionConfigs []clouddeploytarget.ExecutionConfigs `hcl:"execution_configs,block" validate:"min=0"`
	// Gke: optional
	Gke *clouddeploytarget.Gke `hcl:"gke,block"`
	// MultiTarget: optional
	MultiTarget *clouddeploytarget.MultiTarget `hcl:"multi_target,block"`
	// Run: optional
	Run *clouddeploytarget.Run `hcl:"run,block"`
	// Timeouts: optional
	Timeouts *clouddeploytarget.Timeouts `hcl:"timeouts,block"`
}
type clouddeployTargetAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ct.ref.Append("annotations"))
}

// CreateTime returns a reference to field create_time of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("create_time"))
}

// DeployParameters returns a reference to field deploy_parameters of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) DeployParameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ct.ref.Append("deploy_parameters"))
}

// Description returns a reference to field description of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("description"))
}

// EffectiveAnnotations returns a reference to field effective_annotations of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) EffectiveAnnotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ct.ref.Append("effective_annotations"))
}

// EffectiveLabels returns a reference to field effective_labels of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ct.ref.Append("effective_labels"))
}

// Etag returns a reference to field etag of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("etag"))
}

// Id returns a reference to field id of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("id"))
}

// Labels returns a reference to field labels of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ct.ref.Append("labels"))
}

// Location returns a reference to field location of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("location"))
}

// Name returns a reference to field name of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("name"))
}

// Project returns a reference to field project of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("project"))
}

// RequireApproval returns a reference to field require_approval of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) RequireApproval() terra.BoolValue {
	return terra.ReferenceAsBool(ct.ref.Append("require_approval"))
}

// TargetId returns a reference to field target_id of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) TargetId() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("target_id"))
}

// TerraformLabels returns a reference to field terraform_labels of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ct.ref.Append("terraform_labels"))
}

// Uid returns a reference to field uid of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_clouddeploy_target.
func (ct clouddeployTargetAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("update_time"))
}

func (ct clouddeployTargetAttributes) AnthosCluster() terra.ListValue[clouddeploytarget.AnthosClusterAttributes] {
	return terra.ReferenceAsList[clouddeploytarget.AnthosClusterAttributes](ct.ref.Append("anthos_cluster"))
}

func (ct clouddeployTargetAttributes) ExecutionConfigs() terra.ListValue[clouddeploytarget.ExecutionConfigsAttributes] {
	return terra.ReferenceAsList[clouddeploytarget.ExecutionConfigsAttributes](ct.ref.Append("execution_configs"))
}

func (ct clouddeployTargetAttributes) Gke() terra.ListValue[clouddeploytarget.GkeAttributes] {
	return terra.ReferenceAsList[clouddeploytarget.GkeAttributes](ct.ref.Append("gke"))
}

func (ct clouddeployTargetAttributes) MultiTarget() terra.ListValue[clouddeploytarget.MultiTargetAttributes] {
	return terra.ReferenceAsList[clouddeploytarget.MultiTargetAttributes](ct.ref.Append("multi_target"))
}

func (ct clouddeployTargetAttributes) Run() terra.ListValue[clouddeploytarget.RunAttributes] {
	return terra.ReferenceAsList[clouddeploytarget.RunAttributes](ct.ref.Append("run"))
}

func (ct clouddeployTargetAttributes) Timeouts() clouddeploytarget.TimeoutsAttributes {
	return terra.ReferenceAsSingle[clouddeploytarget.TimeoutsAttributes](ct.ref.Append("timeouts"))
}

type clouddeployTargetState struct {
	Annotations          map[string]string                         `json:"annotations"`
	CreateTime           string                                    `json:"create_time"`
	DeployParameters     map[string]string                         `json:"deploy_parameters"`
	Description          string                                    `json:"description"`
	EffectiveAnnotations map[string]string                         `json:"effective_annotations"`
	EffectiveLabels      map[string]string                         `json:"effective_labels"`
	Etag                 string                                    `json:"etag"`
	Id                   string                                    `json:"id"`
	Labels               map[string]string                         `json:"labels"`
	Location             string                                    `json:"location"`
	Name                 string                                    `json:"name"`
	Project              string                                    `json:"project"`
	RequireApproval      bool                                      `json:"require_approval"`
	TargetId             string                                    `json:"target_id"`
	TerraformLabels      map[string]string                         `json:"terraform_labels"`
	Uid                  string                                    `json:"uid"`
	UpdateTime           string                                    `json:"update_time"`
	AnthosCluster        []clouddeploytarget.AnthosClusterState    `json:"anthos_cluster"`
	ExecutionConfigs     []clouddeploytarget.ExecutionConfigsState `json:"execution_configs"`
	Gke                  []clouddeploytarget.GkeState              `json:"gke"`
	MultiTarget          []clouddeploytarget.MultiTargetState      `json:"multi_target"`
	Run                  []clouddeploytarget.RunState              `json:"run"`
	Timeouts             *clouddeploytarget.TimeoutsState          `json:"timeouts"`
}

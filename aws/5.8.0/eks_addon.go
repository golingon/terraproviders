// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	eksaddon "github.com/golingon/terraproviders/aws/5.8.0/eksaddon"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEksAddon creates a new instance of [EksAddon].
func NewEksAddon(name string, args EksAddonArgs) *EksAddon {
	return &EksAddon{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EksAddon)(nil)

// EksAddon represents the Terraform resource aws_eks_addon.
type EksAddon struct {
	Name      string
	Args      EksAddonArgs
	state     *eksAddonState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EksAddon].
func (ea *EksAddon) Type() string {
	return "aws_eks_addon"
}

// LocalName returns the local name for [EksAddon].
func (ea *EksAddon) LocalName() string {
	return ea.Name
}

// Configuration returns the configuration (args) for [EksAddon].
func (ea *EksAddon) Configuration() interface{} {
	return ea.Args
}

// DependOn is used for other resources to depend on [EksAddon].
func (ea *EksAddon) DependOn() terra.Reference {
	return terra.ReferenceResource(ea)
}

// Dependencies returns the list of resources [EksAddon] depends_on.
func (ea *EksAddon) Dependencies() terra.Dependencies {
	return ea.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EksAddon].
func (ea *EksAddon) LifecycleManagement() *terra.Lifecycle {
	return ea.Lifecycle
}

// Attributes returns the attributes for [EksAddon].
func (ea *EksAddon) Attributes() eksAddonAttributes {
	return eksAddonAttributes{ref: terra.ReferenceResource(ea)}
}

// ImportState imports the given attribute values into [EksAddon]'s state.
func (ea *EksAddon) ImportState(av io.Reader) error {
	ea.state = &eksAddonState{}
	if err := json.NewDecoder(av).Decode(ea.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ea.Type(), ea.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EksAddon] has state.
func (ea *EksAddon) State() (*eksAddonState, bool) {
	return ea.state, ea.state != nil
}

// StateMust returns the state for [EksAddon]. Panics if the state is nil.
func (ea *EksAddon) StateMust() *eksAddonState {
	if ea.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ea.Type(), ea.LocalName()))
	}
	return ea.state
}

// EksAddonArgs contains the configurations for aws_eks_addon.
type EksAddonArgs struct {
	// AddonName: string, required
	AddonName terra.StringValue `hcl:"addon_name,attr" validate:"required"`
	// AddonVersion: string, optional
	AddonVersion terra.StringValue `hcl:"addon_version,attr"`
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// ConfigurationValues: string, optional
	ConfigurationValues terra.StringValue `hcl:"configuration_values,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Preserve: bool, optional
	Preserve terra.BoolValue `hcl:"preserve,attr"`
	// ResolveConflicts: string, optional
	ResolveConflicts terra.StringValue `hcl:"resolve_conflicts,attr"`
	// ResolveConflictsOnCreate: string, optional
	ResolveConflictsOnCreate terra.StringValue `hcl:"resolve_conflicts_on_create,attr"`
	// ResolveConflictsOnUpdate: string, optional
	ResolveConflictsOnUpdate terra.StringValue `hcl:"resolve_conflicts_on_update,attr"`
	// ServiceAccountRoleArn: string, optional
	ServiceAccountRoleArn terra.StringValue `hcl:"service_account_role_arn,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *eksaddon.Timeouts `hcl:"timeouts,block"`
}
type eksAddonAttributes struct {
	ref terra.Reference
}

// AddonName returns a reference to field addon_name of aws_eks_addon.
func (ea eksAddonAttributes) AddonName() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("addon_name"))
}

// AddonVersion returns a reference to field addon_version of aws_eks_addon.
func (ea eksAddonAttributes) AddonVersion() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("addon_version"))
}

// Arn returns a reference to field arn of aws_eks_addon.
func (ea eksAddonAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("arn"))
}

// ClusterName returns a reference to field cluster_name of aws_eks_addon.
func (ea eksAddonAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("cluster_name"))
}

// ConfigurationValues returns a reference to field configuration_values of aws_eks_addon.
func (ea eksAddonAttributes) ConfigurationValues() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("configuration_values"))
}

// CreatedAt returns a reference to field created_at of aws_eks_addon.
func (ea eksAddonAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_eks_addon.
func (ea eksAddonAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("id"))
}

// ModifiedAt returns a reference to field modified_at of aws_eks_addon.
func (ea eksAddonAttributes) ModifiedAt() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("modified_at"))
}

// Preserve returns a reference to field preserve of aws_eks_addon.
func (ea eksAddonAttributes) Preserve() terra.BoolValue {
	return terra.ReferenceAsBool(ea.ref.Append("preserve"))
}

// ResolveConflicts returns a reference to field resolve_conflicts of aws_eks_addon.
func (ea eksAddonAttributes) ResolveConflicts() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("resolve_conflicts"))
}

// ResolveConflictsOnCreate returns a reference to field resolve_conflicts_on_create of aws_eks_addon.
func (ea eksAddonAttributes) ResolveConflictsOnCreate() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("resolve_conflicts_on_create"))
}

// ResolveConflictsOnUpdate returns a reference to field resolve_conflicts_on_update of aws_eks_addon.
func (ea eksAddonAttributes) ResolveConflictsOnUpdate() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("resolve_conflicts_on_update"))
}

// ServiceAccountRoleArn returns a reference to field service_account_role_arn of aws_eks_addon.
func (ea eksAddonAttributes) ServiceAccountRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("service_account_role_arn"))
}

// Tags returns a reference to field tags of aws_eks_addon.
func (ea eksAddonAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ea.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_eks_addon.
func (ea eksAddonAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ea.ref.Append("tags_all"))
}

func (ea eksAddonAttributes) Timeouts() eksaddon.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eksaddon.TimeoutsAttributes](ea.ref.Append("timeouts"))
}

type eksAddonState struct {
	AddonName                string                  `json:"addon_name"`
	AddonVersion             string                  `json:"addon_version"`
	Arn                      string                  `json:"arn"`
	ClusterName              string                  `json:"cluster_name"`
	ConfigurationValues      string                  `json:"configuration_values"`
	CreatedAt                string                  `json:"created_at"`
	Id                       string                  `json:"id"`
	ModifiedAt               string                  `json:"modified_at"`
	Preserve                 bool                    `json:"preserve"`
	ResolveConflicts         string                  `json:"resolve_conflicts"`
	ResolveConflictsOnCreate string                  `json:"resolve_conflicts_on_create"`
	ResolveConflictsOnUpdate string                  `json:"resolve_conflicts_on_update"`
	ServiceAccountRoleArn    string                  `json:"service_account_role_arn"`
	Tags                     map[string]string       `json:"tags"`
	TagsAll                  map[string]string       `json:"tags_all"`
	Timeouts                 *eksaddon.TimeoutsState `json:"timeouts"`
}

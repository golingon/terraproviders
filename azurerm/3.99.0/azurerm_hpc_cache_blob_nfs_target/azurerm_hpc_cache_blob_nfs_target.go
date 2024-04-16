// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_hpc_cache_blob_nfs_target

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

// Resource represents the Terraform resource azurerm_hpc_cache_blob_nfs_target.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermHpcCacheBlobNfsTargetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ahcbnt *Resource) Type() string {
	return "azurerm_hpc_cache_blob_nfs_target"
}

// LocalName returns the local name for [Resource].
func (ahcbnt *Resource) LocalName() string {
	return ahcbnt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ahcbnt *Resource) Configuration() interface{} {
	return ahcbnt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ahcbnt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ahcbnt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ahcbnt *Resource) Dependencies() terra.Dependencies {
	return ahcbnt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ahcbnt *Resource) LifecycleManagement() *terra.Lifecycle {
	return ahcbnt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ahcbnt *Resource) Attributes() azurermHpcCacheBlobNfsTargetAttributes {
	return azurermHpcCacheBlobNfsTargetAttributes{ref: terra.ReferenceResource(ahcbnt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ahcbnt *Resource) ImportState(state io.Reader) error {
	ahcbnt.state = &azurermHpcCacheBlobNfsTargetState{}
	if err := json.NewDecoder(state).Decode(ahcbnt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ahcbnt.Type(), ahcbnt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ahcbnt *Resource) State() (*azurermHpcCacheBlobNfsTargetState, bool) {
	return ahcbnt.state, ahcbnt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ahcbnt *Resource) StateMust() *azurermHpcCacheBlobNfsTargetState {
	if ahcbnt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ahcbnt.Type(), ahcbnt.LocalName()))
	}
	return ahcbnt.state
}

// Args contains the configurations for azurerm_hpc_cache_blob_nfs_target.
type Args struct {
	// AccessPolicyName: string, optional
	AccessPolicyName terra.StringValue `hcl:"access_policy_name,attr"`
	// CacheName: string, required
	CacheName terra.StringValue `hcl:"cache_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespacePath: string, required
	NamespacePath terra.StringValue `hcl:"namespace_path,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageContainerId: string, required
	StorageContainerId terra.StringValue `hcl:"storage_container_id,attr" validate:"required"`
	// UsageModel: string, required
	UsageModel terra.StringValue `hcl:"usage_model,attr" validate:"required"`
	// VerificationTimerInSeconds: number, optional
	VerificationTimerInSeconds terra.NumberValue `hcl:"verification_timer_in_seconds,attr"`
	// WriteBackTimerInSeconds: number, optional
	WriteBackTimerInSeconds terra.NumberValue `hcl:"write_back_timer_in_seconds,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermHpcCacheBlobNfsTargetAttributes struct {
	ref terra.Reference
}

// AccessPolicyName returns a reference to field access_policy_name of azurerm_hpc_cache_blob_nfs_target.
func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) AccessPolicyName() terra.StringValue {
	return terra.ReferenceAsString(ahcbnt.ref.Append("access_policy_name"))
}

// CacheName returns a reference to field cache_name of azurerm_hpc_cache_blob_nfs_target.
func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) CacheName() terra.StringValue {
	return terra.ReferenceAsString(ahcbnt.ref.Append("cache_name"))
}

// Id returns a reference to field id of azurerm_hpc_cache_blob_nfs_target.
func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ahcbnt.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_hpc_cache_blob_nfs_target.
func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ahcbnt.ref.Append("name"))
}

// NamespacePath returns a reference to field namespace_path of azurerm_hpc_cache_blob_nfs_target.
func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) NamespacePath() terra.StringValue {
	return terra.ReferenceAsString(ahcbnt.ref.Append("namespace_path"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_hpc_cache_blob_nfs_target.
func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ahcbnt.ref.Append("resource_group_name"))
}

// StorageContainerId returns a reference to field storage_container_id of azurerm_hpc_cache_blob_nfs_target.
func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) StorageContainerId() terra.StringValue {
	return terra.ReferenceAsString(ahcbnt.ref.Append("storage_container_id"))
}

// UsageModel returns a reference to field usage_model of azurerm_hpc_cache_blob_nfs_target.
func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) UsageModel() terra.StringValue {
	return terra.ReferenceAsString(ahcbnt.ref.Append("usage_model"))
}

// VerificationTimerInSeconds returns a reference to field verification_timer_in_seconds of azurerm_hpc_cache_blob_nfs_target.
func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) VerificationTimerInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ahcbnt.ref.Append("verification_timer_in_seconds"))
}

// WriteBackTimerInSeconds returns a reference to field write_back_timer_in_seconds of azurerm_hpc_cache_blob_nfs_target.
func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) WriteBackTimerInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ahcbnt.ref.Append("write_back_timer_in_seconds"))
}

func (ahcbnt azurermHpcCacheBlobNfsTargetAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ahcbnt.ref.Append("timeouts"))
}

type azurermHpcCacheBlobNfsTargetState struct {
	AccessPolicyName           string         `json:"access_policy_name"`
	CacheName                  string         `json:"cache_name"`
	Id                         string         `json:"id"`
	Name                       string         `json:"name"`
	NamespacePath              string         `json:"namespace_path"`
	ResourceGroupName          string         `json:"resource_group_name"`
	StorageContainerId         string         `json:"storage_container_id"`
	UsageModel                 string         `json:"usage_model"`
	VerificationTimerInSeconds float64        `json:"verification_timer_in_seconds"`
	WriteBackTimerInSeconds    float64        `json:"write_back_timer_in_seconds"`
	Timeouts                   *TimeoutsState `json:"timeouts"`
}

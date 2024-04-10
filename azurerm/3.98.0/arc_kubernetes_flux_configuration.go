// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	arckubernetesfluxconfiguration "github.com/golingon/terraproviders/azurerm/3.98.0/arckubernetesfluxconfiguration"
	"io"
)

// NewArcKubernetesFluxConfiguration creates a new instance of [ArcKubernetesFluxConfiguration].
func NewArcKubernetesFluxConfiguration(name string, args ArcKubernetesFluxConfigurationArgs) *ArcKubernetesFluxConfiguration {
	return &ArcKubernetesFluxConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ArcKubernetesFluxConfiguration)(nil)

// ArcKubernetesFluxConfiguration represents the Terraform resource azurerm_arc_kubernetes_flux_configuration.
type ArcKubernetesFluxConfiguration struct {
	Name      string
	Args      ArcKubernetesFluxConfigurationArgs
	state     *arcKubernetesFluxConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ArcKubernetesFluxConfiguration].
func (akfc *ArcKubernetesFluxConfiguration) Type() string {
	return "azurerm_arc_kubernetes_flux_configuration"
}

// LocalName returns the local name for [ArcKubernetesFluxConfiguration].
func (akfc *ArcKubernetesFluxConfiguration) LocalName() string {
	return akfc.Name
}

// Configuration returns the configuration (args) for [ArcKubernetesFluxConfiguration].
func (akfc *ArcKubernetesFluxConfiguration) Configuration() interface{} {
	return akfc.Args
}

// DependOn is used for other resources to depend on [ArcKubernetesFluxConfiguration].
func (akfc *ArcKubernetesFluxConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(akfc)
}

// Dependencies returns the list of resources [ArcKubernetesFluxConfiguration] depends_on.
func (akfc *ArcKubernetesFluxConfiguration) Dependencies() terra.Dependencies {
	return akfc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ArcKubernetesFluxConfiguration].
func (akfc *ArcKubernetesFluxConfiguration) LifecycleManagement() *terra.Lifecycle {
	return akfc.Lifecycle
}

// Attributes returns the attributes for [ArcKubernetesFluxConfiguration].
func (akfc *ArcKubernetesFluxConfiguration) Attributes() arcKubernetesFluxConfigurationAttributes {
	return arcKubernetesFluxConfigurationAttributes{ref: terra.ReferenceResource(akfc)}
}

// ImportState imports the given attribute values into [ArcKubernetesFluxConfiguration]'s state.
func (akfc *ArcKubernetesFluxConfiguration) ImportState(av io.Reader) error {
	akfc.state = &arcKubernetesFluxConfigurationState{}
	if err := json.NewDecoder(av).Decode(akfc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akfc.Type(), akfc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ArcKubernetesFluxConfiguration] has state.
func (akfc *ArcKubernetesFluxConfiguration) State() (*arcKubernetesFluxConfigurationState, bool) {
	return akfc.state, akfc.state != nil
}

// StateMust returns the state for [ArcKubernetesFluxConfiguration]. Panics if the state is nil.
func (akfc *ArcKubernetesFluxConfiguration) StateMust() *arcKubernetesFluxConfigurationState {
	if akfc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akfc.Type(), akfc.LocalName()))
	}
	return akfc.state
}

// ArcKubernetesFluxConfigurationArgs contains the configurations for azurerm_arc_kubernetes_flux_configuration.
type ArcKubernetesFluxConfigurationArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// ContinuousReconciliationEnabled: bool, optional
	ContinuousReconciliationEnabled terra.BoolValue `hcl:"continuous_reconciliation_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
	// Scope: string, optional
	Scope terra.StringValue `hcl:"scope,attr"`
	// BlobStorage: optional
	BlobStorage *arckubernetesfluxconfiguration.BlobStorage `hcl:"blob_storage,block"`
	// Bucket: optional
	Bucket *arckubernetesfluxconfiguration.Bucket `hcl:"bucket,block"`
	// GitRepository: optional
	GitRepository *arckubernetesfluxconfiguration.GitRepository `hcl:"git_repository,block"`
	// Kustomizations: min=1
	Kustomizations []arckubernetesfluxconfiguration.Kustomizations `hcl:"kustomizations,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *arckubernetesfluxconfiguration.Timeouts `hcl:"timeouts,block"`
}
type arcKubernetesFluxConfigurationAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of azurerm_arc_kubernetes_flux_configuration.
func (akfc arcKubernetesFluxConfigurationAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(akfc.ref.Append("cluster_id"))
}

// ContinuousReconciliationEnabled returns a reference to field continuous_reconciliation_enabled of azurerm_arc_kubernetes_flux_configuration.
func (akfc arcKubernetesFluxConfigurationAttributes) ContinuousReconciliationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(akfc.ref.Append("continuous_reconciliation_enabled"))
}

// Id returns a reference to field id of azurerm_arc_kubernetes_flux_configuration.
func (akfc arcKubernetesFluxConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akfc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_arc_kubernetes_flux_configuration.
func (akfc arcKubernetesFluxConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akfc.ref.Append("name"))
}

// Namespace returns a reference to field namespace of azurerm_arc_kubernetes_flux_configuration.
func (akfc arcKubernetesFluxConfigurationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(akfc.ref.Append("namespace"))
}

// Scope returns a reference to field scope of azurerm_arc_kubernetes_flux_configuration.
func (akfc arcKubernetesFluxConfigurationAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(akfc.ref.Append("scope"))
}

func (akfc arcKubernetesFluxConfigurationAttributes) BlobStorage() terra.ListValue[arckubernetesfluxconfiguration.BlobStorageAttributes] {
	return terra.ReferenceAsList[arckubernetesfluxconfiguration.BlobStorageAttributes](akfc.ref.Append("blob_storage"))
}

func (akfc arcKubernetesFluxConfigurationAttributes) Bucket() terra.ListValue[arckubernetesfluxconfiguration.BucketAttributes] {
	return terra.ReferenceAsList[arckubernetesfluxconfiguration.BucketAttributes](akfc.ref.Append("bucket"))
}

func (akfc arcKubernetesFluxConfigurationAttributes) GitRepository() terra.ListValue[arckubernetesfluxconfiguration.GitRepositoryAttributes] {
	return terra.ReferenceAsList[arckubernetesfluxconfiguration.GitRepositoryAttributes](akfc.ref.Append("git_repository"))
}

func (akfc arcKubernetesFluxConfigurationAttributes) Kustomizations() terra.SetValue[arckubernetesfluxconfiguration.KustomizationsAttributes] {
	return terra.ReferenceAsSet[arckubernetesfluxconfiguration.KustomizationsAttributes](akfc.ref.Append("kustomizations"))
}

func (akfc arcKubernetesFluxConfigurationAttributes) Timeouts() arckubernetesfluxconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[arckubernetesfluxconfiguration.TimeoutsAttributes](akfc.ref.Append("timeouts"))
}

type arcKubernetesFluxConfigurationState struct {
	ClusterId                       string                                               `json:"cluster_id"`
	ContinuousReconciliationEnabled bool                                                 `json:"continuous_reconciliation_enabled"`
	Id                              string                                               `json:"id"`
	Name                            string                                               `json:"name"`
	Namespace                       string                                               `json:"namespace"`
	Scope                           string                                               `json:"scope"`
	BlobStorage                     []arckubernetesfluxconfiguration.BlobStorageState    `json:"blob_storage"`
	Bucket                          []arckubernetesfluxconfiguration.BucketState         `json:"bucket"`
	GitRepository                   []arckubernetesfluxconfiguration.GitRepositoryState  `json:"git_repository"`
	Kustomizations                  []arckubernetesfluxconfiguration.KustomizationsState `json:"kustomizations"`
	Timeouts                        *arckubernetesfluxconfiguration.TimeoutsState        `json:"timeouts"`
}

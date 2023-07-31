// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kustoattacheddatabaseconfiguration "github.com/golingon/terraproviders/azurerm/3.67.0/kustoattacheddatabaseconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKustoAttachedDatabaseConfiguration creates a new instance of [KustoAttachedDatabaseConfiguration].
func NewKustoAttachedDatabaseConfiguration(name string, args KustoAttachedDatabaseConfigurationArgs) *KustoAttachedDatabaseConfiguration {
	return &KustoAttachedDatabaseConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoAttachedDatabaseConfiguration)(nil)

// KustoAttachedDatabaseConfiguration represents the Terraform resource azurerm_kusto_attached_database_configuration.
type KustoAttachedDatabaseConfiguration struct {
	Name      string
	Args      KustoAttachedDatabaseConfigurationArgs
	state     *kustoAttachedDatabaseConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoAttachedDatabaseConfiguration].
func (kadc *KustoAttachedDatabaseConfiguration) Type() string {
	return "azurerm_kusto_attached_database_configuration"
}

// LocalName returns the local name for [KustoAttachedDatabaseConfiguration].
func (kadc *KustoAttachedDatabaseConfiguration) LocalName() string {
	return kadc.Name
}

// Configuration returns the configuration (args) for [KustoAttachedDatabaseConfiguration].
func (kadc *KustoAttachedDatabaseConfiguration) Configuration() interface{} {
	return kadc.Args
}

// DependOn is used for other resources to depend on [KustoAttachedDatabaseConfiguration].
func (kadc *KustoAttachedDatabaseConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(kadc)
}

// Dependencies returns the list of resources [KustoAttachedDatabaseConfiguration] depends_on.
func (kadc *KustoAttachedDatabaseConfiguration) Dependencies() terra.Dependencies {
	return kadc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoAttachedDatabaseConfiguration].
func (kadc *KustoAttachedDatabaseConfiguration) LifecycleManagement() *terra.Lifecycle {
	return kadc.Lifecycle
}

// Attributes returns the attributes for [KustoAttachedDatabaseConfiguration].
func (kadc *KustoAttachedDatabaseConfiguration) Attributes() kustoAttachedDatabaseConfigurationAttributes {
	return kustoAttachedDatabaseConfigurationAttributes{ref: terra.ReferenceResource(kadc)}
}

// ImportState imports the given attribute values into [KustoAttachedDatabaseConfiguration]'s state.
func (kadc *KustoAttachedDatabaseConfiguration) ImportState(av io.Reader) error {
	kadc.state = &kustoAttachedDatabaseConfigurationState{}
	if err := json.NewDecoder(av).Decode(kadc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kadc.Type(), kadc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoAttachedDatabaseConfiguration] has state.
func (kadc *KustoAttachedDatabaseConfiguration) State() (*kustoAttachedDatabaseConfigurationState, bool) {
	return kadc.state, kadc.state != nil
}

// StateMust returns the state for [KustoAttachedDatabaseConfiguration]. Panics if the state is nil.
func (kadc *KustoAttachedDatabaseConfiguration) StateMust() *kustoAttachedDatabaseConfigurationState {
	if kadc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kadc.Type(), kadc.LocalName()))
	}
	return kadc.state
}

// KustoAttachedDatabaseConfigurationArgs contains the configurations for azurerm_kusto_attached_database_configuration.
type KustoAttachedDatabaseConfigurationArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// ClusterResourceId: string, required
	ClusterResourceId terra.StringValue `hcl:"cluster_resource_id,attr" validate:"required"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// DefaultPrincipalModificationKind: string, optional
	DefaultPrincipalModificationKind terra.StringValue `hcl:"default_principal_modification_kind,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sharing: optional
	Sharing *kustoattacheddatabaseconfiguration.Sharing `hcl:"sharing,block"`
	// Timeouts: optional
	Timeouts *kustoattacheddatabaseconfiguration.Timeouts `hcl:"timeouts,block"`
}
type kustoAttachedDatabaseConfigurationAttributes struct {
	ref terra.Reference
}

// AttachedDatabaseNames returns a reference to field attached_database_names of azurerm_kusto_attached_database_configuration.
func (kadc kustoAttachedDatabaseConfigurationAttributes) AttachedDatabaseNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kadc.ref.Append("attached_database_names"))
}

// ClusterName returns a reference to field cluster_name of azurerm_kusto_attached_database_configuration.
func (kadc kustoAttachedDatabaseConfigurationAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(kadc.ref.Append("cluster_name"))
}

// ClusterResourceId returns a reference to field cluster_resource_id of azurerm_kusto_attached_database_configuration.
func (kadc kustoAttachedDatabaseConfigurationAttributes) ClusterResourceId() terra.StringValue {
	return terra.ReferenceAsString(kadc.ref.Append("cluster_resource_id"))
}

// DatabaseName returns a reference to field database_name of azurerm_kusto_attached_database_configuration.
func (kadc kustoAttachedDatabaseConfigurationAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(kadc.ref.Append("database_name"))
}

// DefaultPrincipalModificationKind returns a reference to field default_principal_modification_kind of azurerm_kusto_attached_database_configuration.
func (kadc kustoAttachedDatabaseConfigurationAttributes) DefaultPrincipalModificationKind() terra.StringValue {
	return terra.ReferenceAsString(kadc.ref.Append("default_principal_modification_kind"))
}

// Id returns a reference to field id of azurerm_kusto_attached_database_configuration.
func (kadc kustoAttachedDatabaseConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kadc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_kusto_attached_database_configuration.
func (kadc kustoAttachedDatabaseConfigurationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kadc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_kusto_attached_database_configuration.
func (kadc kustoAttachedDatabaseConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kadc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kusto_attached_database_configuration.
func (kadc kustoAttachedDatabaseConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kadc.ref.Append("resource_group_name"))
}

func (kadc kustoAttachedDatabaseConfigurationAttributes) Sharing() terra.ListValue[kustoattacheddatabaseconfiguration.SharingAttributes] {
	return terra.ReferenceAsList[kustoattacheddatabaseconfiguration.SharingAttributes](kadc.ref.Append("sharing"))
}

func (kadc kustoAttachedDatabaseConfigurationAttributes) Timeouts() kustoattacheddatabaseconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustoattacheddatabaseconfiguration.TimeoutsAttributes](kadc.ref.Append("timeouts"))
}

type kustoAttachedDatabaseConfigurationState struct {
	AttachedDatabaseNames            []string                                          `json:"attached_database_names"`
	ClusterName                      string                                            `json:"cluster_name"`
	ClusterResourceId                string                                            `json:"cluster_resource_id"`
	DatabaseName                     string                                            `json:"database_name"`
	DefaultPrincipalModificationKind string                                            `json:"default_principal_modification_kind"`
	Id                               string                                            `json:"id"`
	Location                         string                                            `json:"location"`
	Name                             string                                            `json:"name"`
	ResourceGroupName                string                                            `json:"resource_group_name"`
	Sharing                          []kustoattacheddatabaseconfiguration.SharingState `json:"sharing"`
	Timeouts                         *kustoattacheddatabaseconfiguration.TimeoutsState `json:"timeouts"`
}

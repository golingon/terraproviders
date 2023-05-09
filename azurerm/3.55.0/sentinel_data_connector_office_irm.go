// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectorofficeirm "github.com/golingon/terraproviders/azurerm/3.55.0/sentineldataconnectorofficeirm"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorOfficeIrm creates a new instance of [SentinelDataConnectorOfficeIrm].
func NewSentinelDataConnectorOfficeIrm(name string, args SentinelDataConnectorOfficeIrmArgs) *SentinelDataConnectorOfficeIrm {
	return &SentinelDataConnectorOfficeIrm{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorOfficeIrm)(nil)

// SentinelDataConnectorOfficeIrm represents the Terraform resource azurerm_sentinel_data_connector_office_irm.
type SentinelDataConnectorOfficeIrm struct {
	Name      string
	Args      SentinelDataConnectorOfficeIrmArgs
	state     *sentinelDataConnectorOfficeIrmState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorOfficeIrm].
func (sdcoi *SentinelDataConnectorOfficeIrm) Type() string {
	return "azurerm_sentinel_data_connector_office_irm"
}

// LocalName returns the local name for [SentinelDataConnectorOfficeIrm].
func (sdcoi *SentinelDataConnectorOfficeIrm) LocalName() string {
	return sdcoi.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorOfficeIrm].
func (sdcoi *SentinelDataConnectorOfficeIrm) Configuration() interface{} {
	return sdcoi.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorOfficeIrm].
func (sdcoi *SentinelDataConnectorOfficeIrm) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcoi)
}

// Dependencies returns the list of resources [SentinelDataConnectorOfficeIrm] depends_on.
func (sdcoi *SentinelDataConnectorOfficeIrm) Dependencies() terra.Dependencies {
	return sdcoi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorOfficeIrm].
func (sdcoi *SentinelDataConnectorOfficeIrm) LifecycleManagement() *terra.Lifecycle {
	return sdcoi.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorOfficeIrm].
func (sdcoi *SentinelDataConnectorOfficeIrm) Attributes() sentinelDataConnectorOfficeIrmAttributes {
	return sentinelDataConnectorOfficeIrmAttributes{ref: terra.ReferenceResource(sdcoi)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorOfficeIrm]'s state.
func (sdcoi *SentinelDataConnectorOfficeIrm) ImportState(av io.Reader) error {
	sdcoi.state = &sentinelDataConnectorOfficeIrmState{}
	if err := json.NewDecoder(av).Decode(sdcoi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcoi.Type(), sdcoi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorOfficeIrm] has state.
func (sdcoi *SentinelDataConnectorOfficeIrm) State() (*sentinelDataConnectorOfficeIrmState, bool) {
	return sdcoi.state, sdcoi.state != nil
}

// StateMust returns the state for [SentinelDataConnectorOfficeIrm]. Panics if the state is nil.
func (sdcoi *SentinelDataConnectorOfficeIrm) StateMust() *sentinelDataConnectorOfficeIrmState {
	if sdcoi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcoi.Type(), sdcoi.LocalName()))
	}
	return sdcoi.state
}

// SentinelDataConnectorOfficeIrmArgs contains the configurations for azurerm_sentinel_data_connector_office_irm.
type SentinelDataConnectorOfficeIrmArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectorofficeirm.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorOfficeIrmAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_office_irm.
func (sdcoi sentinelDataConnectorOfficeIrmAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcoi.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_office_irm.
func (sdcoi sentinelDataConnectorOfficeIrmAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcoi.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_office_irm.
func (sdcoi sentinelDataConnectorOfficeIrmAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcoi.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_office_irm.
func (sdcoi sentinelDataConnectorOfficeIrmAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdcoi.ref.Append("tenant_id"))
}

func (sdcoi sentinelDataConnectorOfficeIrmAttributes) Timeouts() sentineldataconnectorofficeirm.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectorofficeirm.TimeoutsAttributes](sdcoi.ref.Append("timeouts"))
}

type sentinelDataConnectorOfficeIrmState struct {
	Id                      string                                        `json:"id"`
	LogAnalyticsWorkspaceId string                                        `json:"log_analytics_workspace_id"`
	Name                    string                                        `json:"name"`
	TenantId                string                                        `json:"tenant_id"`
	Timeouts                *sentineldataconnectorofficeirm.TimeoutsState `json:"timeouts"`
}

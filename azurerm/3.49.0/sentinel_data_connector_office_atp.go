// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectorofficeatp "github.com/golingon/terraproviders/azurerm/3.49.0/sentineldataconnectorofficeatp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorOfficeAtp creates a new instance of [SentinelDataConnectorOfficeAtp].
func NewSentinelDataConnectorOfficeAtp(name string, args SentinelDataConnectorOfficeAtpArgs) *SentinelDataConnectorOfficeAtp {
	return &SentinelDataConnectorOfficeAtp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorOfficeAtp)(nil)

// SentinelDataConnectorOfficeAtp represents the Terraform resource azurerm_sentinel_data_connector_office_atp.
type SentinelDataConnectorOfficeAtp struct {
	Name      string
	Args      SentinelDataConnectorOfficeAtpArgs
	state     *sentinelDataConnectorOfficeAtpState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorOfficeAtp].
func (sdcoa *SentinelDataConnectorOfficeAtp) Type() string {
	return "azurerm_sentinel_data_connector_office_atp"
}

// LocalName returns the local name for [SentinelDataConnectorOfficeAtp].
func (sdcoa *SentinelDataConnectorOfficeAtp) LocalName() string {
	return sdcoa.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorOfficeAtp].
func (sdcoa *SentinelDataConnectorOfficeAtp) Configuration() interface{} {
	return sdcoa.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorOfficeAtp].
func (sdcoa *SentinelDataConnectorOfficeAtp) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcoa)
}

// Dependencies returns the list of resources [SentinelDataConnectorOfficeAtp] depends_on.
func (sdcoa *SentinelDataConnectorOfficeAtp) Dependencies() terra.Dependencies {
	return sdcoa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorOfficeAtp].
func (sdcoa *SentinelDataConnectorOfficeAtp) LifecycleManagement() *terra.Lifecycle {
	return sdcoa.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorOfficeAtp].
func (sdcoa *SentinelDataConnectorOfficeAtp) Attributes() sentinelDataConnectorOfficeAtpAttributes {
	return sentinelDataConnectorOfficeAtpAttributes{ref: terra.ReferenceResource(sdcoa)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorOfficeAtp]'s state.
func (sdcoa *SentinelDataConnectorOfficeAtp) ImportState(av io.Reader) error {
	sdcoa.state = &sentinelDataConnectorOfficeAtpState{}
	if err := json.NewDecoder(av).Decode(sdcoa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcoa.Type(), sdcoa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorOfficeAtp] has state.
func (sdcoa *SentinelDataConnectorOfficeAtp) State() (*sentinelDataConnectorOfficeAtpState, bool) {
	return sdcoa.state, sdcoa.state != nil
}

// StateMust returns the state for [SentinelDataConnectorOfficeAtp]. Panics if the state is nil.
func (sdcoa *SentinelDataConnectorOfficeAtp) StateMust() *sentinelDataConnectorOfficeAtpState {
	if sdcoa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcoa.Type(), sdcoa.LocalName()))
	}
	return sdcoa.state
}

// SentinelDataConnectorOfficeAtpArgs contains the configurations for azurerm_sentinel_data_connector_office_atp.
type SentinelDataConnectorOfficeAtpArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectorofficeatp.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorOfficeAtpAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_office_atp.
func (sdcoa sentinelDataConnectorOfficeAtpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcoa.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_office_atp.
func (sdcoa sentinelDataConnectorOfficeAtpAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcoa.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_office_atp.
func (sdcoa sentinelDataConnectorOfficeAtpAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcoa.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_office_atp.
func (sdcoa sentinelDataConnectorOfficeAtpAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdcoa.ref.Append("tenant_id"))
}

func (sdcoa sentinelDataConnectorOfficeAtpAttributes) Timeouts() sentineldataconnectorofficeatp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectorofficeatp.TimeoutsAttributes](sdcoa.ref.Append("timeouts"))
}

type sentinelDataConnectorOfficeAtpState struct {
	Id                      string                                        `json:"id"`
	LogAnalyticsWorkspaceId string                                        `json:"log_analytics_workspace_id"`
	Name                    string                                        `json:"name"`
	TenantId                string                                        `json:"tenant_id"`
	Timeouts                *sentineldataconnectorofficeatp.TimeoutsState `json:"timeouts"`
}

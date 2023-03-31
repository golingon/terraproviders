// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectoroffice365project "github.com/golingon/terraproviders/azurerm/3.49.0/sentineldataconnectoroffice365project"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSentinelDataConnectorOffice365Project(name string, args SentinelDataConnectorOffice365ProjectArgs) *SentinelDataConnectorOffice365Project {
	return &SentinelDataConnectorOffice365Project{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorOffice365Project)(nil)

type SentinelDataConnectorOffice365Project struct {
	Name  string
	Args  SentinelDataConnectorOffice365ProjectArgs
	state *sentinelDataConnectorOffice365ProjectState
}

func (sdco3p *SentinelDataConnectorOffice365Project) Type() string {
	return "azurerm_sentinel_data_connector_office_365_project"
}

func (sdco3p *SentinelDataConnectorOffice365Project) LocalName() string {
	return sdco3p.Name
}

func (sdco3p *SentinelDataConnectorOffice365Project) Configuration() interface{} {
	return sdco3p.Args
}

func (sdco3p *SentinelDataConnectorOffice365Project) Attributes() sentinelDataConnectorOffice365ProjectAttributes {
	return sentinelDataConnectorOffice365ProjectAttributes{ref: terra.ReferenceResource(sdco3p)}
}

func (sdco3p *SentinelDataConnectorOffice365Project) ImportState(av io.Reader) error {
	sdco3p.state = &sentinelDataConnectorOffice365ProjectState{}
	if err := json.NewDecoder(av).Decode(sdco3p.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdco3p.Type(), sdco3p.LocalName(), err)
	}
	return nil
}

func (sdco3p *SentinelDataConnectorOffice365Project) State() (*sentinelDataConnectorOffice365ProjectState, bool) {
	return sdco3p.state, sdco3p.state != nil
}

func (sdco3p *SentinelDataConnectorOffice365Project) StateMust() *sentinelDataConnectorOffice365ProjectState {
	if sdco3p.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdco3p.Type(), sdco3p.LocalName()))
	}
	return sdco3p.state
}

func (sdco3p *SentinelDataConnectorOffice365Project) DependOn() terra.Reference {
	return terra.ReferenceResource(sdco3p)
}

type SentinelDataConnectorOffice365ProjectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectoroffice365project.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that SentinelDataConnectorOffice365Project depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sentinelDataConnectorOffice365ProjectAttributes struct {
	ref terra.Reference
}

func (sdco3p sentinelDataConnectorOffice365ProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sdco3p.ref.Append("id"))
}

func (sdco3p sentinelDataConnectorOffice365ProjectAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceString(sdco3p.ref.Append("log_analytics_workspace_id"))
}

func (sdco3p sentinelDataConnectorOffice365ProjectAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sdco3p.ref.Append("name"))
}

func (sdco3p sentinelDataConnectorOffice365ProjectAttributes) TenantId() terra.StringValue {
	return terra.ReferenceString(sdco3p.ref.Append("tenant_id"))
}

func (sdco3p sentinelDataConnectorOffice365ProjectAttributes) Timeouts() sentineldataconnectoroffice365project.TimeoutsAttributes {
	return terra.ReferenceSingle[sentineldataconnectoroffice365project.TimeoutsAttributes](sdco3p.ref.Append("timeouts"))
}

type sentinelDataConnectorOffice365ProjectState struct {
	Id                      string                                               `json:"id"`
	LogAnalyticsWorkspaceId string                                               `json:"log_analytics_workspace_id"`
	Name                    string                                               `json:"name"`
	TenantId                string                                               `json:"tenant_id"`
	Timeouts                *sentineldataconnectoroffice365project.TimeoutsState `json:"timeouts"`
}

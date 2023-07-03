// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datamonitoringuptimecheckips "github.com/golingon/terraproviders/googlebeta/4.71.0/datamonitoringuptimecheckips"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMonitoringUptimeCheckIps creates a new instance of [DataMonitoringUptimeCheckIps].
func NewDataMonitoringUptimeCheckIps(name string, args DataMonitoringUptimeCheckIpsArgs) *DataMonitoringUptimeCheckIps {
	return &DataMonitoringUptimeCheckIps{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitoringUptimeCheckIps)(nil)

// DataMonitoringUptimeCheckIps represents the Terraform data resource google_monitoring_uptime_check_ips.
type DataMonitoringUptimeCheckIps struct {
	Name string
	Args DataMonitoringUptimeCheckIpsArgs
}

// DataSource returns the Terraform object type for [DataMonitoringUptimeCheckIps].
func (muci *DataMonitoringUptimeCheckIps) DataSource() string {
	return "google_monitoring_uptime_check_ips"
}

// LocalName returns the local name for [DataMonitoringUptimeCheckIps].
func (muci *DataMonitoringUptimeCheckIps) LocalName() string {
	return muci.Name
}

// Configuration returns the configuration (args) for [DataMonitoringUptimeCheckIps].
func (muci *DataMonitoringUptimeCheckIps) Configuration() interface{} {
	return muci.Args
}

// Attributes returns the attributes for [DataMonitoringUptimeCheckIps].
func (muci *DataMonitoringUptimeCheckIps) Attributes() dataMonitoringUptimeCheckIpsAttributes {
	return dataMonitoringUptimeCheckIpsAttributes{ref: terra.ReferenceDataResource(muci)}
}

// DataMonitoringUptimeCheckIpsArgs contains the configurations for google_monitoring_uptime_check_ips.
type DataMonitoringUptimeCheckIpsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// UptimeCheckIps: min=0
	UptimeCheckIps []datamonitoringuptimecheckips.UptimeCheckIps `hcl:"uptime_check_ips,block" validate:"min=0"`
}
type dataMonitoringUptimeCheckIpsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_monitoring_uptime_check_ips.
func (muci dataMonitoringUptimeCheckIpsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(muci.ref.Append("id"))
}

func (muci dataMonitoringUptimeCheckIpsAttributes) UptimeCheckIps() terra.ListValue[datamonitoringuptimecheckips.UptimeCheckIpsAttributes] {
	return terra.ReferenceAsList[datamonitoringuptimecheckips.UptimeCheckIpsAttributes](muci.ref.Append("uptime_check_ips"))
}

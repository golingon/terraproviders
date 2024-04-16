// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_monitoring_uptime_check_ips

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_monitoring_uptime_check_ips.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gmuci *DataSource) DataSource() string {
	return "google_monitoring_uptime_check_ips"
}

// LocalName returns the local name for [DataSource].
func (gmuci *DataSource) LocalName() string {
	return gmuci.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gmuci *DataSource) Configuration() interface{} {
	return gmuci.Args
}

// Attributes returns the attributes for [DataSource].
func (gmuci *DataSource) Attributes() dataGoogleMonitoringUptimeCheckIpsAttributes {
	return dataGoogleMonitoringUptimeCheckIpsAttributes{ref: terra.ReferenceDataSource(gmuci)}
}

// DataArgs contains the configurations for google_monitoring_uptime_check_ips.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type dataGoogleMonitoringUptimeCheckIpsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_monitoring_uptime_check_ips.
func (gmuci dataGoogleMonitoringUptimeCheckIpsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gmuci.ref.Append("id"))
}

func (gmuci dataGoogleMonitoringUptimeCheckIpsAttributes) UptimeCheckIps() terra.ListValue[DataUptimeCheckIpsAttributes] {
	return terra.ReferenceAsList[DataUptimeCheckIpsAttributes](gmuci.ref.Append("uptime_check_ips"))
}

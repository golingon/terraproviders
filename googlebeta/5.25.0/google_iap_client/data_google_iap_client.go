// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_iap_client

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_iap_client.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gic *DataSource) DataSource() string {
	return "google_iap_client"
}

// LocalName returns the local name for [DataSource].
func (gic *DataSource) LocalName() string {
	return gic.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gic *DataSource) Configuration() interface{} {
	return gic.Args
}

// Attributes returns the attributes for [DataSource].
func (gic *DataSource) Attributes() dataGoogleIapClientAttributes {
	return dataGoogleIapClientAttributes{ref: terra.ReferenceDataSource(gic)}
}

// DataArgs contains the configurations for google_iap_client.
type DataArgs struct {
	// Brand: string, required
	Brand terra.StringValue `hcl:"brand,attr" validate:"required"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type dataGoogleIapClientAttributes struct {
	ref terra.Reference
}

// Brand returns a reference to field brand of google_iap_client.
func (gic dataGoogleIapClientAttributes) Brand() terra.StringValue {
	return terra.ReferenceAsString(gic.ref.Append("brand"))
}

// ClientId returns a reference to field client_id of google_iap_client.
func (gic dataGoogleIapClientAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(gic.ref.Append("client_id"))
}

// DisplayName returns a reference to field display_name of google_iap_client.
func (gic dataGoogleIapClientAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gic.ref.Append("display_name"))
}

// Id returns a reference to field id of google_iap_client.
func (gic dataGoogleIapClientAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gic.ref.Append("id"))
}

// Secret returns a reference to field secret of google_iap_client.
func (gic dataGoogleIapClientAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(gic.ref.Append("secret"))
}

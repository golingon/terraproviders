// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_router

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_compute_router.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gcr *DataSource) DataSource() string {
	return "google_compute_router"
}

// LocalName returns the local name for [DataSource].
func (gcr *DataSource) LocalName() string {
	return gcr.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gcr *DataSource) Configuration() interface{} {
	return gcr.Args
}

// Attributes returns the attributes for [DataSource].
func (gcr *DataSource) Attributes() dataGoogleComputeRouterAttributes {
	return dataGoogleComputeRouterAttributes{ref: terra.ReferenceDataSource(gcr)}
}

// DataArgs contains the configurations for google_compute_router.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}

type dataGoogleComputeRouterAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_router.
func (gcr dataGoogleComputeRouterAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_router.
func (gcr dataGoogleComputeRouterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("description"))
}

// EncryptedInterconnectRouter returns a reference to field encrypted_interconnect_router of google_compute_router.
func (gcr dataGoogleComputeRouterAttributes) EncryptedInterconnectRouter() terra.BoolValue {
	return terra.ReferenceAsBool(gcr.ref.Append("encrypted_interconnect_router"))
}

// Id returns a reference to field id of google_compute_router.
func (gcr dataGoogleComputeRouterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_router.
func (gcr dataGoogleComputeRouterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_router.
func (gcr dataGoogleComputeRouterAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_router.
func (gcr dataGoogleComputeRouterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_router.
func (gcr dataGoogleComputeRouterAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_router.
func (gcr dataGoogleComputeRouterAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("self_link"))
}

func (gcr dataGoogleComputeRouterAttributes) Bgp() terra.ListValue[DataBgpAttributes] {
	return terra.ReferenceAsList[DataBgpAttributes](gcr.ref.Append("bgp"))
}

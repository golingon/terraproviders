// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_firebase_hosting_channel

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_firebase_hosting_channel.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gfhc *DataSource) DataSource() string {
	return "google_firebase_hosting_channel"
}

// LocalName returns the local name for [DataSource].
func (gfhc *DataSource) LocalName() string {
	return gfhc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gfhc *DataSource) Configuration() interface{} {
	return gfhc.Args
}

// Attributes returns the attributes for [DataSource].
func (gfhc *DataSource) Attributes() dataGoogleFirebaseHostingChannelAttributes {
	return dataGoogleFirebaseHostingChannelAttributes{ref: terra.ReferenceDataSource(gfhc)}
}

// DataArgs contains the configurations for google_firebase_hosting_channel.
type DataArgs struct {
	// ChannelId: string, required
	ChannelId terra.StringValue `hcl:"channel_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SiteId: string, required
	SiteId terra.StringValue `hcl:"site_id,attr" validate:"required"`
}

type dataGoogleFirebaseHostingChannelAttributes struct {
	ref terra.Reference
}

// ChannelId returns a reference to field channel_id of google_firebase_hosting_channel.
func (gfhc dataGoogleFirebaseHostingChannelAttributes) ChannelId() terra.StringValue {
	return terra.ReferenceAsString(gfhc.ref.Append("channel_id"))
}

// EffectiveLabels returns a reference to field effective_labels of google_firebase_hosting_channel.
func (gfhc dataGoogleFirebaseHostingChannelAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gfhc.ref.Append("effective_labels"))
}

// ExpireTime returns a reference to field expire_time of google_firebase_hosting_channel.
func (gfhc dataGoogleFirebaseHostingChannelAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(gfhc.ref.Append("expire_time"))
}

// Id returns a reference to field id of google_firebase_hosting_channel.
func (gfhc dataGoogleFirebaseHostingChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gfhc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_firebase_hosting_channel.
func (gfhc dataGoogleFirebaseHostingChannelAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gfhc.ref.Append("labels"))
}

// Name returns a reference to field name of google_firebase_hosting_channel.
func (gfhc dataGoogleFirebaseHostingChannelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gfhc.ref.Append("name"))
}

// RetainedReleaseCount returns a reference to field retained_release_count of google_firebase_hosting_channel.
func (gfhc dataGoogleFirebaseHostingChannelAttributes) RetainedReleaseCount() terra.NumberValue {
	return terra.ReferenceAsNumber(gfhc.ref.Append("retained_release_count"))
}

// SiteId returns a reference to field site_id of google_firebase_hosting_channel.
func (gfhc dataGoogleFirebaseHostingChannelAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(gfhc.ref.Append("site_id"))
}

// TerraformLabels returns a reference to field terraform_labels of google_firebase_hosting_channel.
func (gfhc dataGoogleFirebaseHostingChannelAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gfhc.ref.Append("terraform_labels"))
}

// Ttl returns a reference to field ttl of google_firebase_hosting_channel.
func (gfhc dataGoogleFirebaseHostingChannelAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(gfhc.ref.Append("ttl"))
}

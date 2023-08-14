// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package kubernetesclusternodepool

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type KubeletConfig struct {
	// AllowedUnsafeSysctls: set of string, optional
	AllowedUnsafeSysctls terra.SetValue[terra.StringValue] `hcl:"allowed_unsafe_sysctls,attr"`
	// ContainerLogMaxLine: number, optional
	ContainerLogMaxLine terra.NumberValue `hcl:"container_log_max_line,attr"`
	// ContainerLogMaxSizeMb: number, optional
	ContainerLogMaxSizeMb terra.NumberValue `hcl:"container_log_max_size_mb,attr"`
	// CpuCfsQuotaEnabled: bool, optional
	CpuCfsQuotaEnabled terra.BoolValue `hcl:"cpu_cfs_quota_enabled,attr"`
	// CpuCfsQuotaPeriod: string, optional
	CpuCfsQuotaPeriod terra.StringValue `hcl:"cpu_cfs_quota_period,attr"`
	// CpuManagerPolicy: string, optional
	CpuManagerPolicy terra.StringValue `hcl:"cpu_manager_policy,attr"`
	// ImageGcHighThreshold: number, optional
	ImageGcHighThreshold terra.NumberValue `hcl:"image_gc_high_threshold,attr"`
	// ImageGcLowThreshold: number, optional
	ImageGcLowThreshold terra.NumberValue `hcl:"image_gc_low_threshold,attr"`
	// PodMaxPid: number, optional
	PodMaxPid terra.NumberValue `hcl:"pod_max_pid,attr"`
	// TopologyManagerPolicy: string, optional
	TopologyManagerPolicy terra.StringValue `hcl:"topology_manager_policy,attr"`
}

type LinuxOsConfig struct {
	// SwapFileSizeMb: number, optional
	SwapFileSizeMb terra.NumberValue `hcl:"swap_file_size_mb,attr"`
	// TransparentHugePageDefrag: string, optional
	TransparentHugePageDefrag terra.StringValue `hcl:"transparent_huge_page_defrag,attr"`
	// TransparentHugePageEnabled: string, optional
	TransparentHugePageEnabled terra.StringValue `hcl:"transparent_huge_page_enabled,attr"`
	// SysctlConfig: optional
	SysctlConfig *SysctlConfig `hcl:"sysctl_config,block"`
}

type SysctlConfig struct {
	// FsAioMaxNr: number, optional
	FsAioMaxNr terra.NumberValue `hcl:"fs_aio_max_nr,attr"`
	// FsFileMax: number, optional
	FsFileMax terra.NumberValue `hcl:"fs_file_max,attr"`
	// FsInotifyMaxUserWatches: number, optional
	FsInotifyMaxUserWatches terra.NumberValue `hcl:"fs_inotify_max_user_watches,attr"`
	// FsNrOpen: number, optional
	FsNrOpen terra.NumberValue `hcl:"fs_nr_open,attr"`
	// KernelThreadsMax: number, optional
	KernelThreadsMax terra.NumberValue `hcl:"kernel_threads_max,attr"`
	// NetCoreNetdevMaxBacklog: number, optional
	NetCoreNetdevMaxBacklog terra.NumberValue `hcl:"net_core_netdev_max_backlog,attr"`
	// NetCoreOptmemMax: number, optional
	NetCoreOptmemMax terra.NumberValue `hcl:"net_core_optmem_max,attr"`
	// NetCoreRmemDefault: number, optional
	NetCoreRmemDefault terra.NumberValue `hcl:"net_core_rmem_default,attr"`
	// NetCoreRmemMax: number, optional
	NetCoreRmemMax terra.NumberValue `hcl:"net_core_rmem_max,attr"`
	// NetCoreSomaxconn: number, optional
	NetCoreSomaxconn terra.NumberValue `hcl:"net_core_somaxconn,attr"`
	// NetCoreWmemDefault: number, optional
	NetCoreWmemDefault terra.NumberValue `hcl:"net_core_wmem_default,attr"`
	// NetCoreWmemMax: number, optional
	NetCoreWmemMax terra.NumberValue `hcl:"net_core_wmem_max,attr"`
	// NetIpv4IpLocalPortRangeMax: number, optional
	NetIpv4IpLocalPortRangeMax terra.NumberValue `hcl:"net_ipv4_ip_local_port_range_max,attr"`
	// NetIpv4IpLocalPortRangeMin: number, optional
	NetIpv4IpLocalPortRangeMin terra.NumberValue `hcl:"net_ipv4_ip_local_port_range_min,attr"`
	// NetIpv4NeighDefaultGcThresh1: number, optional
	NetIpv4NeighDefaultGcThresh1 terra.NumberValue `hcl:"net_ipv4_neigh_default_gc_thresh1,attr"`
	// NetIpv4NeighDefaultGcThresh2: number, optional
	NetIpv4NeighDefaultGcThresh2 terra.NumberValue `hcl:"net_ipv4_neigh_default_gc_thresh2,attr"`
	// NetIpv4NeighDefaultGcThresh3: number, optional
	NetIpv4NeighDefaultGcThresh3 terra.NumberValue `hcl:"net_ipv4_neigh_default_gc_thresh3,attr"`
	// NetIpv4TcpFinTimeout: number, optional
	NetIpv4TcpFinTimeout terra.NumberValue `hcl:"net_ipv4_tcp_fin_timeout,attr"`
	// NetIpv4TcpKeepaliveIntvl: number, optional
	NetIpv4TcpKeepaliveIntvl terra.NumberValue `hcl:"net_ipv4_tcp_keepalive_intvl,attr"`
	// NetIpv4TcpKeepaliveProbes: number, optional
	NetIpv4TcpKeepaliveProbes terra.NumberValue `hcl:"net_ipv4_tcp_keepalive_probes,attr"`
	// NetIpv4TcpKeepaliveTime: number, optional
	NetIpv4TcpKeepaliveTime terra.NumberValue `hcl:"net_ipv4_tcp_keepalive_time,attr"`
	// NetIpv4TcpMaxSynBacklog: number, optional
	NetIpv4TcpMaxSynBacklog terra.NumberValue `hcl:"net_ipv4_tcp_max_syn_backlog,attr"`
	// NetIpv4TcpMaxTwBuckets: number, optional
	NetIpv4TcpMaxTwBuckets terra.NumberValue `hcl:"net_ipv4_tcp_max_tw_buckets,attr"`
	// NetIpv4TcpTwReuse: bool, optional
	NetIpv4TcpTwReuse terra.BoolValue `hcl:"net_ipv4_tcp_tw_reuse,attr"`
	// NetNetfilterNfConntrackBuckets: number, optional
	NetNetfilterNfConntrackBuckets terra.NumberValue `hcl:"net_netfilter_nf_conntrack_buckets,attr"`
	// NetNetfilterNfConntrackMax: number, optional
	NetNetfilterNfConntrackMax terra.NumberValue `hcl:"net_netfilter_nf_conntrack_max,attr"`
	// VmMaxMapCount: number, optional
	VmMaxMapCount terra.NumberValue `hcl:"vm_max_map_count,attr"`
	// VmSwappiness: number, optional
	VmSwappiness terra.NumberValue `hcl:"vm_swappiness,attr"`
	// VmVfsCachePressure: number, optional
	VmVfsCachePressure terra.NumberValue `hcl:"vm_vfs_cache_pressure,attr"`
}

type NodeNetworkProfile struct {
	// NodePublicIpTags: map of string, optional
	NodePublicIpTags terra.MapValue[terra.StringValue] `hcl:"node_public_ip_tags,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type UpgradeSettings struct {
	// MaxSurge: string, required
	MaxSurge terra.StringValue `hcl:"max_surge,attr" validate:"required"`
}

type WindowsProfile struct {
	// OutboundNatEnabled: bool, optional
	OutboundNatEnabled terra.BoolValue `hcl:"outbound_nat_enabled,attr"`
}

type KubeletConfigAttributes struct {
	ref terra.Reference
}

func (kc KubeletConfigAttributes) InternalRef() (terra.Reference, error) {
	return kc.ref, nil
}

func (kc KubeletConfigAttributes) InternalWithRef(ref terra.Reference) KubeletConfigAttributes {
	return KubeletConfigAttributes{ref: ref}
}

func (kc KubeletConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kc.ref.InternalTokens()
}

func (kc KubeletConfigAttributes) AllowedUnsafeSysctls() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kc.ref.Append("allowed_unsafe_sysctls"))
}

func (kc KubeletConfigAttributes) ContainerLogMaxLine() terra.NumberValue {
	return terra.ReferenceAsNumber(kc.ref.Append("container_log_max_line"))
}

func (kc KubeletConfigAttributes) ContainerLogMaxSizeMb() terra.NumberValue {
	return terra.ReferenceAsNumber(kc.ref.Append("container_log_max_size_mb"))
}

func (kc KubeletConfigAttributes) CpuCfsQuotaEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("cpu_cfs_quota_enabled"))
}

func (kc KubeletConfigAttributes) CpuCfsQuotaPeriod() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("cpu_cfs_quota_period"))
}

func (kc KubeletConfigAttributes) CpuManagerPolicy() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("cpu_manager_policy"))
}

func (kc KubeletConfigAttributes) ImageGcHighThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(kc.ref.Append("image_gc_high_threshold"))
}

func (kc KubeletConfigAttributes) ImageGcLowThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(kc.ref.Append("image_gc_low_threshold"))
}

func (kc KubeletConfigAttributes) PodMaxPid() terra.NumberValue {
	return terra.ReferenceAsNumber(kc.ref.Append("pod_max_pid"))
}

func (kc KubeletConfigAttributes) TopologyManagerPolicy() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("topology_manager_policy"))
}

type LinuxOsConfigAttributes struct {
	ref terra.Reference
}

func (loc LinuxOsConfigAttributes) InternalRef() (terra.Reference, error) {
	return loc.ref, nil
}

func (loc LinuxOsConfigAttributes) InternalWithRef(ref terra.Reference) LinuxOsConfigAttributes {
	return LinuxOsConfigAttributes{ref: ref}
}

func (loc LinuxOsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return loc.ref.InternalTokens()
}

func (loc LinuxOsConfigAttributes) SwapFileSizeMb() terra.NumberValue {
	return terra.ReferenceAsNumber(loc.ref.Append("swap_file_size_mb"))
}

func (loc LinuxOsConfigAttributes) TransparentHugePageDefrag() terra.StringValue {
	return terra.ReferenceAsString(loc.ref.Append("transparent_huge_page_defrag"))
}

func (loc LinuxOsConfigAttributes) TransparentHugePageEnabled() terra.StringValue {
	return terra.ReferenceAsString(loc.ref.Append("transparent_huge_page_enabled"))
}

func (loc LinuxOsConfigAttributes) SysctlConfig() terra.ListValue[SysctlConfigAttributes] {
	return terra.ReferenceAsList[SysctlConfigAttributes](loc.ref.Append("sysctl_config"))
}

type SysctlConfigAttributes struct {
	ref terra.Reference
}

func (sc SysctlConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SysctlConfigAttributes) InternalWithRef(ref terra.Reference) SysctlConfigAttributes {
	return SysctlConfigAttributes{ref: ref}
}

func (sc SysctlConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SysctlConfigAttributes) FsAioMaxNr() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("fs_aio_max_nr"))
}

func (sc SysctlConfigAttributes) FsFileMax() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("fs_file_max"))
}

func (sc SysctlConfigAttributes) FsInotifyMaxUserWatches() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("fs_inotify_max_user_watches"))
}

func (sc SysctlConfigAttributes) FsNrOpen() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("fs_nr_open"))
}

func (sc SysctlConfigAttributes) KernelThreadsMax() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("kernel_threads_max"))
}

func (sc SysctlConfigAttributes) NetCoreNetdevMaxBacklog() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_core_netdev_max_backlog"))
}

func (sc SysctlConfigAttributes) NetCoreOptmemMax() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_core_optmem_max"))
}

func (sc SysctlConfigAttributes) NetCoreRmemDefault() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_core_rmem_default"))
}

func (sc SysctlConfigAttributes) NetCoreRmemMax() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_core_rmem_max"))
}

func (sc SysctlConfigAttributes) NetCoreSomaxconn() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_core_somaxconn"))
}

func (sc SysctlConfigAttributes) NetCoreWmemDefault() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_core_wmem_default"))
}

func (sc SysctlConfigAttributes) NetCoreWmemMax() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_core_wmem_max"))
}

func (sc SysctlConfigAttributes) NetIpv4IpLocalPortRangeMax() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_ip_local_port_range_max"))
}

func (sc SysctlConfigAttributes) NetIpv4IpLocalPortRangeMin() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_ip_local_port_range_min"))
}

func (sc SysctlConfigAttributes) NetIpv4NeighDefaultGcThresh1() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_neigh_default_gc_thresh1"))
}

func (sc SysctlConfigAttributes) NetIpv4NeighDefaultGcThresh2() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_neigh_default_gc_thresh2"))
}

func (sc SysctlConfigAttributes) NetIpv4NeighDefaultGcThresh3() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_neigh_default_gc_thresh3"))
}

func (sc SysctlConfigAttributes) NetIpv4TcpFinTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_tcp_fin_timeout"))
}

func (sc SysctlConfigAttributes) NetIpv4TcpKeepaliveIntvl() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_tcp_keepalive_intvl"))
}

func (sc SysctlConfigAttributes) NetIpv4TcpKeepaliveProbes() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_tcp_keepalive_probes"))
}

func (sc SysctlConfigAttributes) NetIpv4TcpKeepaliveTime() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_tcp_keepalive_time"))
}

func (sc SysctlConfigAttributes) NetIpv4TcpMaxSynBacklog() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_tcp_max_syn_backlog"))
}

func (sc SysctlConfigAttributes) NetIpv4TcpMaxTwBuckets() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_ipv4_tcp_max_tw_buckets"))
}

func (sc SysctlConfigAttributes) NetIpv4TcpTwReuse() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("net_ipv4_tcp_tw_reuse"))
}

func (sc SysctlConfigAttributes) NetNetfilterNfConntrackBuckets() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_netfilter_nf_conntrack_buckets"))
}

func (sc SysctlConfigAttributes) NetNetfilterNfConntrackMax() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("net_netfilter_nf_conntrack_max"))
}

func (sc SysctlConfigAttributes) VmMaxMapCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("vm_max_map_count"))
}

func (sc SysctlConfigAttributes) VmSwappiness() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("vm_swappiness"))
}

func (sc SysctlConfigAttributes) VmVfsCachePressure() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("vm_vfs_cache_pressure"))
}

type NodeNetworkProfileAttributes struct {
	ref terra.Reference
}

func (nnp NodeNetworkProfileAttributes) InternalRef() (terra.Reference, error) {
	return nnp.ref, nil
}

func (nnp NodeNetworkProfileAttributes) InternalWithRef(ref terra.Reference) NodeNetworkProfileAttributes {
	return NodeNetworkProfileAttributes{ref: ref}
}

func (nnp NodeNetworkProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nnp.ref.InternalTokens()
}

func (nnp NodeNetworkProfileAttributes) NodePublicIpTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nnp.ref.Append("node_public_ip_tags"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type UpgradeSettingsAttributes struct {
	ref terra.Reference
}

func (us UpgradeSettingsAttributes) InternalRef() (terra.Reference, error) {
	return us.ref, nil
}

func (us UpgradeSettingsAttributes) InternalWithRef(ref terra.Reference) UpgradeSettingsAttributes {
	return UpgradeSettingsAttributes{ref: ref}
}

func (us UpgradeSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return us.ref.InternalTokens()
}

func (us UpgradeSettingsAttributes) MaxSurge() terra.StringValue {
	return terra.ReferenceAsString(us.ref.Append("max_surge"))
}

type WindowsProfileAttributes struct {
	ref terra.Reference
}

func (wp WindowsProfileAttributes) InternalRef() (terra.Reference, error) {
	return wp.ref, nil
}

func (wp WindowsProfileAttributes) InternalWithRef(ref terra.Reference) WindowsProfileAttributes {
	return WindowsProfileAttributes{ref: ref}
}

func (wp WindowsProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wp.ref.InternalTokens()
}

func (wp WindowsProfileAttributes) OutboundNatEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wp.ref.Append("outbound_nat_enabled"))
}

type KubeletConfigState struct {
	AllowedUnsafeSysctls  []string `json:"allowed_unsafe_sysctls"`
	ContainerLogMaxLine   float64  `json:"container_log_max_line"`
	ContainerLogMaxSizeMb float64  `json:"container_log_max_size_mb"`
	CpuCfsQuotaEnabled    bool     `json:"cpu_cfs_quota_enabled"`
	CpuCfsQuotaPeriod     string   `json:"cpu_cfs_quota_period"`
	CpuManagerPolicy      string   `json:"cpu_manager_policy"`
	ImageGcHighThreshold  float64  `json:"image_gc_high_threshold"`
	ImageGcLowThreshold   float64  `json:"image_gc_low_threshold"`
	PodMaxPid             float64  `json:"pod_max_pid"`
	TopologyManagerPolicy string   `json:"topology_manager_policy"`
}

type LinuxOsConfigState struct {
	SwapFileSizeMb             float64             `json:"swap_file_size_mb"`
	TransparentHugePageDefrag  string              `json:"transparent_huge_page_defrag"`
	TransparentHugePageEnabled string              `json:"transparent_huge_page_enabled"`
	SysctlConfig               []SysctlConfigState `json:"sysctl_config"`
}

type SysctlConfigState struct {
	FsAioMaxNr                     float64 `json:"fs_aio_max_nr"`
	FsFileMax                      float64 `json:"fs_file_max"`
	FsInotifyMaxUserWatches        float64 `json:"fs_inotify_max_user_watches"`
	FsNrOpen                       float64 `json:"fs_nr_open"`
	KernelThreadsMax               float64 `json:"kernel_threads_max"`
	NetCoreNetdevMaxBacklog        float64 `json:"net_core_netdev_max_backlog"`
	NetCoreOptmemMax               float64 `json:"net_core_optmem_max"`
	NetCoreRmemDefault             float64 `json:"net_core_rmem_default"`
	NetCoreRmemMax                 float64 `json:"net_core_rmem_max"`
	NetCoreSomaxconn               float64 `json:"net_core_somaxconn"`
	NetCoreWmemDefault             float64 `json:"net_core_wmem_default"`
	NetCoreWmemMax                 float64 `json:"net_core_wmem_max"`
	NetIpv4IpLocalPortRangeMax     float64 `json:"net_ipv4_ip_local_port_range_max"`
	NetIpv4IpLocalPortRangeMin     float64 `json:"net_ipv4_ip_local_port_range_min"`
	NetIpv4NeighDefaultGcThresh1   float64 `json:"net_ipv4_neigh_default_gc_thresh1"`
	NetIpv4NeighDefaultGcThresh2   float64 `json:"net_ipv4_neigh_default_gc_thresh2"`
	NetIpv4NeighDefaultGcThresh3   float64 `json:"net_ipv4_neigh_default_gc_thresh3"`
	NetIpv4TcpFinTimeout           float64 `json:"net_ipv4_tcp_fin_timeout"`
	NetIpv4TcpKeepaliveIntvl       float64 `json:"net_ipv4_tcp_keepalive_intvl"`
	NetIpv4TcpKeepaliveProbes      float64 `json:"net_ipv4_tcp_keepalive_probes"`
	NetIpv4TcpKeepaliveTime        float64 `json:"net_ipv4_tcp_keepalive_time"`
	NetIpv4TcpMaxSynBacklog        float64 `json:"net_ipv4_tcp_max_syn_backlog"`
	NetIpv4TcpMaxTwBuckets         float64 `json:"net_ipv4_tcp_max_tw_buckets"`
	NetIpv4TcpTwReuse              bool    `json:"net_ipv4_tcp_tw_reuse"`
	NetNetfilterNfConntrackBuckets float64 `json:"net_netfilter_nf_conntrack_buckets"`
	NetNetfilterNfConntrackMax     float64 `json:"net_netfilter_nf_conntrack_max"`
	VmMaxMapCount                  float64 `json:"vm_max_map_count"`
	VmSwappiness                   float64 `json:"vm_swappiness"`
	VmVfsCachePressure             float64 `json:"vm_vfs_cache_pressure"`
}

type NodeNetworkProfileState struct {
	NodePublicIpTags map[string]string `json:"node_public_ip_tags"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type UpgradeSettingsState struct {
	MaxSurge string `json:"max_surge"`
}

type WindowsProfileState struct {
	OutboundNatEnabled bool `json:"outbound_nat_enabled"`
}
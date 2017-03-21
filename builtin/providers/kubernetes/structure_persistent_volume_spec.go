package kubernetes

import (
	"k8s.io/kubernetes/pkg/api/v1"

	"github.com/hashicorp/terraform/helper/schema"
)

// Flatteners

func flattenAWSElasticBlockStoreVolumeSource(in *v1.AWSElasticBlockStoreVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.VolumeID != "" {
		att["volume_id"] = in.VolumeID
	}
	if in.FSType != "" {
		att["fs_type"] = in.FSType
	}
	if in.Partition != 0 {
		att["partition"] = in.Partition
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	return []interface{}{att}
}

func flattenAzureDiskVolumeSource(in *v1.AzureDiskVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.DiskName != "" {
		att["disk_name"] = in.DiskName
	}
	if in.DataDiskURI != "" {
		att["data_disk_uri"] = in.DataDiskURI
	}
	if in.CachingMode != nil {
		att["caching_mode"] = *in.CachingMode
	}
	if in.FSType != nil {
		att["fs_type"] = *in.FSType
	}
	if in.ReadOnly != nil {
		att["read_only"] = *in.ReadOnly
	}
	return []interface{}{att}
}

func flattenAzureFileVolumeSource(in *v1.AzureFileVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.SecretName != "" {
		att["secret_name"] = in.SecretName
	}
	if in.ShareName != "" {
		att["share_name"] = in.ShareName
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	return []interface{}{att}
}

func flattenCephFSVolumeSource(in *v1.CephFSVolumeSource) []interface{} {
	att := make(map[string]interface{})
	att["monitors"] = in.Monitors
	if in.Path != "" {
		att["path"] = in.Path
	}
	if in.User != "" {
		att["user"] = in.User
	}
	if in.SecretFile != "" {
		att["secret_file"] = in.SecretFile
	}
	if in.SecretRef != nil {
		att["secret_ref"] = flattenLocalObjectReference(in.SecretRef)
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	return []interface{}{att}
}

func flattenCinderVolumeSource(in *v1.CinderVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.VolumeID != "" {
		att["volume_id"] = in.VolumeID
	}
	if in.FSType != "" {
		att["fs_type"] = in.FSType
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	return []interface{}{att}
}

func flattenFCVolumeSource(in *v1.FCVolumeSource) []interface{} {
	att := make(map[string]interface{})
	att["target_ww_ns"] = in.TargetWWNs
	att["lun"] = *in.Lun
	if in.FSType != "" {
		att["fs_type"] = in.FSType
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	return []interface{}{att}
}

func flattenFlexVolumeSource(in *v1.FlexVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.Driver != "" {
		att["driver"] = in.Driver
	}
	if in.FSType != "" {
		att["fs_type"] = in.FSType
	}
	if in.SecretRef != nil {
		att["secret_ref"] = flattenLocalObjectReference(in.SecretRef)
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	if len(in.Options) > 0 {
		att["options"] = in.Options
	}
	return []interface{}{att}
}

func flattenFlockerVolumeSource(in *v1.FlockerVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.DatasetName != "" {
		att["dataset_name"] = in.DatasetName
	}
	if in.DatasetUUID != "" {
		att["dataset_uuid"] = in.DatasetUUID
	}
	return []interface{}{att}
}

func flattenGCEPersistentDiskVolumeSource(in *v1.GCEPersistentDiskVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.PDName != "" {
		att["pd_name"] = in.PDName
	}
	if in.FSType != "" {
		att["fs_type"] = in.FSType
	}
	if in.Partition != 0 {
		att["partition"] = in.Partition
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	return []interface{}{att}
}

func flattenGlusterfsVolumeSource(in *v1.GlusterfsVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.EndpointsName != "" {
		att["endpoints_name"] = in.EndpointsName
	}
	if in.Path != "" {
		att["path"] = in.Path
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	return []interface{}{att}
}

func flattenHostPathVolumeSource(in *v1.HostPathVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.Path != "" {
		att["path"] = in.Path
	}
	return []interface{}{att}
}

func flattenISCSIVolumeSource(in *v1.ISCSIVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.TargetPortal != "" {
		att["target_portal"] = in.TargetPortal
	}
	if in.IQN != "" {
		att["iqn"] = in.IQN
	}
	if in.Lun != 0 {
		att["lun"] = in.Lun
	}
	if in.ISCSIInterface != "" {
		att["iscsi_interface"] = in.ISCSIInterface
	}
	if in.FSType != "" {
		att["fs_type"] = in.FSType
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	return []interface{}{att}
}

func flattenLocalObjectReference(in *v1.LocalObjectReference) []interface{} {
	att := make(map[string]interface{})
	if in.Name != "" {
		att["name"] = in.Name
	}
	return []interface{}{att}
}

func flattenNFSVolumeSource(in *v1.NFSVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.Server != "" {
		att["server"] = in.Server
	}
	if in.Path != "" {
		att["path"] = in.Path
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	return []interface{}{att}
}

func flattenPersistentVolumeSource(in v1.PersistentVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.GCEPersistentDisk != nil {
		att["gce_persistent_disk"] = flattenGCEPersistentDiskVolumeSource(in.GCEPersistentDisk)
	}
	if in.AWSElasticBlockStore != nil {
		att["aws_elastic_block_store"] = flattenAWSElasticBlockStoreVolumeSource(in.AWSElasticBlockStore)
	}
	if in.HostPath != nil {
		att["host_path"] = flattenHostPathVolumeSource(in.HostPath)
	}
	if in.Glusterfs != nil {
		att["glusterfs"] = flattenGlusterfsVolumeSource(in.Glusterfs)
	}
	if in.NFS != nil {
		att["nfs"] = flattenNFSVolumeSource(in.NFS)
	}
	if in.RBD != nil {
		att["rbd"] = flattenRBDVolumeSource(in.RBD)
	}
	if in.ISCSI != nil {
		att["iscsi"] = flattenISCSIVolumeSource(in.ISCSI)
	}
	if in.Cinder != nil {
		att["cinder"] = flattenCinderVolumeSource(in.Cinder)
	}
	if in.CephFS != nil {
		att["ceph_fs"] = flattenCephFSVolumeSource(in.CephFS)
	}
	if in.FC != nil {
		att["fc"] = flattenFCVolumeSource(in.FC)
	}
	if in.Flocker != nil {
		att["flocker"] = flattenFlockerVolumeSource(in.Flocker)
	}
	if in.FlexVolume != nil {
		att["flex_volume"] = flattenFlexVolumeSource(in.FlexVolume)
	}
	if in.AzureFile != nil {
		att["azure_file"] = flattenAzureFileVolumeSource(in.AzureFile)
	}
	if in.VsphereVolume != nil {
		att["vsphere_volume"] = flattenVsphereVirtualDiskVolumeSource(in.VsphereVolume)
	}
	if in.Quobyte != nil {
		att["quobyte"] = flattenQuobyteVolumeSource(in.Quobyte)
	}
	if in.AzureDisk != nil {
		att["azure_disk"] = flattenAzureDiskVolumeSource(in.AzureDisk)
	}
	if in.PhotonPersistentDisk != nil {
		att["photon_persistent_disk"] = flattenPhotonPersistentDiskVolumeSource(in.PhotonPersistentDisk)
	}
	return []interface{}{att}
}

func flattenPersistentVolumeSpec(in v1.PersistentVolumeSpec) []interface{} {
	att := make(map[string]interface{})
	if len(in.Capacity) > 0 {
		att["capacity"] = flattenResourceList(in.Capacity)
	}

	att["persistent_volume_source"] = flattenPersistentVolumeSource(in.PersistentVolumeSource)
	if len(in.AccessModes) > 0 {
		att["access_modes"] = flattenPersistentVolumeAccessModes(in.AccessModes)
	}
	if in.PersistentVolumeReclaimPolicy != "" {
		att["persistent_volume_reclaim_policy"] = in.PersistentVolumeReclaimPolicy
	}
	return []interface{}{att}
}

func flattenPhotonPersistentDiskVolumeSource(in *v1.PhotonPersistentDiskVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.PdID != "" {
		att["pd_id"] = in.PdID
	}
	if in.FSType != "" {
		att["fs_type"] = in.FSType
	}
	return []interface{}{att}
}

func flattenQuobyteVolumeSource(in *v1.QuobyteVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.Registry != "" {
		att["registry"] = in.Registry
	}
	if in.Volume != "" {
		att["volume"] = in.Volume
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	if in.User != "" {
		att["user"] = in.User
	}
	if in.Group != "" {
		att["group"] = in.Group
	}
	return []interface{}{att}
}

func flattenRBDVolumeSource(in *v1.RBDVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if len(in.CephMonitors) > 0 {
		att["ceph_monitors"] = in.CephMonitors
	}
	if in.RBDImage != "" {
		att["rbd_image"] = in.RBDImage
	}
	if in.FSType != "" {
		att["fs_type"] = in.FSType
	}
	if in.RBDPool != "" {
		att["rbd_pool"] = in.RBDPool
	}
	if in.RadosUser != "" {
		att["rados_user"] = in.RadosUser
	}
	if in.Keyring != "" {
		att["keyring"] = in.Keyring
	}
	if in.SecretRef != nil {
		att["secret_ref"] = flattenLocalObjectReference(in.SecretRef)
	}
	if in.ReadOnly != false {
		att["read_only"] = in.ReadOnly
	}
	return []interface{}{att}
}

func flattenVsphereVirtualDiskVolumeSource(in *v1.VsphereVirtualDiskVolumeSource) []interface{} {
	att := make(map[string]interface{})
	if in.VolumePath != "" {
		att["volume_path"] = in.VolumePath
	}
	if in.FSType != "" {
		att["fs_type"] = in.FSType
	}
	return []interface{}{att}
}

// Expanders

func expandAWSElasticBlockStoreVolumeSource(l []interface{}) *v1.AWSElasticBlockStoreVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.AWSElasticBlockStoreVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.AWSElasticBlockStoreVolumeSource{}
	if v, ok := in["volume_id"].(string); ok {
		obj.VolumeID = v
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = v
	}
	if v, ok := in["partition"].(int32); ok {
		obj.Partition = v
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return obj
}

func expandAzureDiskVolumeSource(l []interface{}) *v1.AzureDiskVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.AzureDiskVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.AzureDiskVolumeSource{}
	if v, ok := in["disk_name"].(string); ok {
		obj.DiskName = v
	}
	if v, ok := in["data_disk_uri"].(string); ok {
		obj.DataDiskURI = v
	}
	if v, ok := in["caching_mode"].(v1.AzureDataDiskCachingMode); ok {
		obj.CachingMode = &v
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = ptrToString(v)
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = ptrToBool(v)
	}
	return obj
}

func expandAzureFileVolumeSource(l []interface{}) *v1.AzureFileVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.AzureFileVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.AzureFileVolumeSource{}
	if v, ok := in["secret_name"].(string); ok {
		obj.SecretName = v
	}
	if v, ok := in["share_name"].(string); ok {
		obj.ShareName = v
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return obj
}

func expandCephFSVolumeSource(l []interface{}) *v1.CephFSVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.CephFSVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.CephFSVolumeSource{
		Monitors: sliceOfString(in["monitors"].([]interface{})),
	}
	if v, ok := in["path"].(string); ok {
		obj.Path = v
	}
	if v, ok := in["user"].(string); ok {
		obj.User = v
	}
	if v, ok := in["secret_file"].(string); ok {
		obj.SecretFile = v
	}
	if v, ok := in["secret_ref"].([]interface{}); ok && len(v) > 0 {
		obj.SecretRef = expandLocalObjectReference(v)
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return obj
}

func expandCinderVolumeSource(l []interface{}) *v1.CinderVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.CinderVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.CinderVolumeSource{}
	if v, ok := in["volume_id"].(string); ok {
		obj.VolumeID = v
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = v
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return obj
}

func expandFCVolumeSource(l []interface{}) *v1.FCVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.FCVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.FCVolumeSource{
		TargetWWNs: sliceOfString(in["target_ww_ns"].([]interface{})),
		Lun:        ptrToInt32(in["lun"].(int32)),
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = v
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return obj
}

func expandFlexVolumeSource(l []interface{}) *v1.FlexVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.FlexVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.FlexVolumeSource{}
	if v, ok := in["driver"].(string); ok {
		obj.Driver = v
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = v
	}
	if v, ok := in["secret_ref"].([]interface{}); ok && len(v) > 0 {
		obj.SecretRef = expandLocalObjectReference(v)
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	if v, ok := in["options"].(map[string]interface{}); ok && len(v) > 0 {
		obj.Options = expandStringMap(v)
	}
	return obj
}

func expandFlockerVolumeSource(l []interface{}) *v1.FlockerVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.FlockerVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.FlockerVolumeSource{}
	if v, ok := in["dataset_name"].(string); ok {
		obj.DatasetName = v
	}
	if v, ok := in["dataset_uuid"].(string); ok {
		obj.DatasetUUID = v
	}
	return obj
}

func expandGCEPersistentDiskVolumeSource(l []interface{}) *v1.GCEPersistentDiskVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.GCEPersistentDiskVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.GCEPersistentDiskVolumeSource{}
	if v, ok := in["pd_name"].(string); ok {
		obj.PDName = v
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = v
	}
	if v, ok := in["partition"].(int32); ok {
		obj.Partition = v
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return obj
}

func expandGlusterfsVolumeSource(l []interface{}) *v1.GlusterfsVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.GlusterfsVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.GlusterfsVolumeSource{}
	if v, ok := in["endpoints_name"].(string); ok {
		obj.EndpointsName = v
	}
	if v, ok := in["path"].(string); ok {
		obj.Path = v
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return obj
}

func expandHostPathVolumeSource(l []interface{}) *v1.HostPathVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.HostPathVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.HostPathVolumeSource{}
	if v, ok := in["path"].(string); ok {
		obj.Path = v
	}
	return obj
}

func expandISCSIVolumeSource(l []interface{}) *v1.ISCSIVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.ISCSIVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.ISCSIVolumeSource{}
	if v, ok := in["target_portal"].(string); ok {
		obj.TargetPortal = v
	}
	if v, ok := in["iqn"].(string); ok {
		obj.IQN = v
	}
	if v, ok := in["lun"].(int32); ok {
		obj.Lun = v
	}
	if v, ok := in["iscsi_interface"].(string); ok {
		obj.ISCSIInterface = v
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = v
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return obj
}

func expandLocalObjectReference(l []interface{}) *v1.LocalObjectReference {
	if len(l) == 0 || l[0] == nil {
		return &v1.LocalObjectReference{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.LocalObjectReference{}
	if v, ok := in["name"].(string); ok {
		obj.Name = v
	}
	return obj
}

func expandNFSVolumeSource(l []interface{}) *v1.NFSVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.NFSVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.NFSVolumeSource{}
	if v, ok := in["server"].(string); ok {
		obj.Server = v
	}
	if v, ok := in["path"].(string); ok {
		obj.Path = v
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return obj
}

func expandPersistentVolumeSource(l []interface{}) v1.PersistentVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return v1.PersistentVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.PersistentVolumeSource{}
	if v, ok := in["gce_persistent_disk"].([]interface{}); ok && len(v) > 0 {
		obj.GCEPersistentDisk = expandGCEPersistentDiskVolumeSource(v)
	}
	if v, ok := in["aws_elastic_block_store"].([]interface{}); ok && len(v) > 0 {
		obj.AWSElasticBlockStore = expandAWSElasticBlockStoreVolumeSource(v)
	}
	if v, ok := in["host_path"].([]interface{}); ok && len(v) > 0 {
		obj.HostPath = expandHostPathVolumeSource(v)
	}
	if v, ok := in["glusterfs"].([]interface{}); ok && len(v) > 0 {
		obj.Glusterfs = expandGlusterfsVolumeSource(v)
	}
	if v, ok := in["nfs"].([]interface{}); ok && len(v) > 0 {
		obj.NFS = expandNFSVolumeSource(v)
	}
	if v, ok := in["rbd"].([]interface{}); ok && len(v) > 0 {
		obj.RBD = expandRBDVolumeSource(v)
	}
	if v, ok := in["iscsi"].([]interface{}); ok && len(v) > 0 {
		obj.ISCSI = expandISCSIVolumeSource(v)
	}
	if v, ok := in["cinder"].([]interface{}); ok && len(v) > 0 {
		obj.Cinder = expandCinderVolumeSource(v)
	}
	if v, ok := in["ceph_fs"].([]interface{}); ok && len(v) > 0 {
		obj.CephFS = expandCephFSVolumeSource(v)
	}
	if v, ok := in["fc"].([]interface{}); ok && len(v) > 0 {
		obj.FC = expandFCVolumeSource(v)
	}
	if v, ok := in["flocker"].([]interface{}); ok && len(v) > 0 {
		obj.Flocker = expandFlockerVolumeSource(v)
	}
	if v, ok := in["flex_volume"].([]interface{}); ok && len(v) > 0 {
		obj.FlexVolume = expandFlexVolumeSource(v)
	}
	if v, ok := in["azure_file"].([]interface{}); ok && len(v) > 0 {
		obj.AzureFile = expandAzureFileVolumeSource(v)
	}
	if v, ok := in["vsphere_volume"].([]interface{}); ok && len(v) > 0 {
		obj.VsphereVolume = expandVsphereVirtualDiskVolumeSource(v)
	}
	if v, ok := in["quobyte"].([]interface{}); ok && len(v) > 0 {
		obj.Quobyte = expandQuobyteVolumeSource(v)
	}
	if v, ok := in["azure_disk"].([]interface{}); ok && len(v) > 0 {
		obj.AzureDisk = expandAzureDiskVolumeSource(v)
	}
	if v, ok := in["photon_persistent_disk"].([]interface{}); ok && len(v) > 0 {
		obj.PhotonPersistentDisk = expandPhotonPersistentDiskVolumeSource(v)
	}
	return obj
}

func expandPersistentVolumeSpec(l []interface{}) (v1.PersistentVolumeSpec, error) {
	if len(l) == 0 || l[0] == nil {
		return v1.PersistentVolumeSpec{}, nil
	}
	in := l[0].(map[string]interface{})
	obj := v1.PersistentVolumeSpec{}
	if v, ok := in["capacity"].(map[string]interface{}); ok && len(v) > 0 {
		var err error
		obj.Capacity, err = expandMapToResourceList(v)
		if err != nil {
			return obj, err
		}
	}
	if v, ok := in["persistent_volume_source"].([]interface{}); ok && len(v) > 0 {
		obj.PersistentVolumeSource = expandPersistentVolumeSource(v)
	}
	if v, ok := in["access_modes"].(*schema.Set); ok && v.Len() > 0 {
		obj.AccessModes = expandPersistentVolumeAccessModes(v.List())
	}
	if v, ok := in["persistent_volume_reclaim_policy"].(v1.PersistentVolumeReclaimPolicy); ok {
		obj.PersistentVolumeReclaimPolicy = v
	}
	return obj, nil
}

func expandPhotonPersistentDiskVolumeSource(l []interface{}) *v1.PhotonPersistentDiskVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.PhotonPersistentDiskVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.PhotonPersistentDiskVolumeSource{}
	if v, ok := in["pd_id"].(string); ok {
		obj.PdID = v
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = v
	}
	return obj
}

func expandQuobyteVolumeSource(l []interface{}) *v1.QuobyteVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.QuobyteVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.QuobyteVolumeSource{}
	if v, ok := in["registry"].(string); ok {
		obj.Registry = v
	}
	if v, ok := in["volume"].(string); ok {
		obj.Volume = v
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	if v, ok := in["user"].(string); ok {
		obj.User = v
	}
	if v, ok := in["group"].(string); ok {
		obj.Group = v
	}
	return obj
}

func expandRBDVolumeSource(l []interface{}) *v1.RBDVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.RBDVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.RBDVolumeSource{}
	if v, ok := in["ceph_monitors"].([]interface{}); ok && len(v) > 0 {
		obj.CephMonitors = sliceOfString(v)
	}
	if v, ok := in["rbd_image"].(string); ok {
		obj.RBDImage = v
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = v
	}
	if v, ok := in["rbd_pool"].(string); ok {
		obj.RBDPool = v
	}
	if v, ok := in["rados_user"].(string); ok {
		obj.RadosUser = v
	}
	if v, ok := in["keyring"].(string); ok {
		obj.Keyring = v
	}
	if v, ok := in["secret_ref"].([]interface{}); ok && len(v) > 0 {
		obj.SecretRef = expandLocalObjectReference(v)
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return obj
}

func expandVsphereVirtualDiskVolumeSource(l []interface{}) *v1.VsphereVirtualDiskVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.VsphereVirtualDiskVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := &v1.VsphereVirtualDiskVolumeSource{}
	if v, ok := in["volume_path"].(string); ok {
		obj.VolumePath = v
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = v
	}
	return obj
}

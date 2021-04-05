// Code generated by github.com/argoproj/argo-cd/hack/known_types. DO NOT EDIT.
package normalizers

import corev1 "k8s.io/api/core/v1"

func init() {
	knownTypes["core/v1/AWSElasticBlockStoreVolumeSource"] = func() interface{} {
		return &corev1.AWSElasticBlockStoreVolumeSource{}
	}
	knownTypes["core/v1/Affinity"] = func() interface{} {
		return &corev1.Affinity{}
	}
	knownTypes["core/v1/AttachedVolume"] = func() interface{} {
		return &corev1.AttachedVolume{}
	}
	knownTypes["core/v1/AvoidPods"] = func() interface{} {
		return &corev1.AvoidPods{}
	}
	knownTypes["core/v1/AzureDiskVolumeSource"] = func() interface{} {
		return &corev1.AzureDiskVolumeSource{}
	}
	knownTypes["core/v1/AzureFilePersistentVolumeSource"] = func() interface{} {
		return &corev1.AzureFilePersistentVolumeSource{}
	}
	knownTypes["core/v1/AzureFileVolumeSource"] = func() interface{} {
		return &corev1.AzureFileVolumeSource{}
	}
	knownTypes["core/v1/Binding"] = func() interface{} {
		return &corev1.Binding{}
	}
	knownTypes["core/v1/CSIPersistentVolumeSource"] = func() interface{} {
		return &corev1.CSIPersistentVolumeSource{}
	}
	knownTypes["core/v1/CSIVolumeSource"] = func() interface{} {
		return &corev1.CSIVolumeSource{}
	}
	knownTypes["core/v1/Capabilities"] = func() interface{} {
		return &corev1.Capabilities{}
	}
	knownTypes["core/v1/CephFSPersistentVolumeSource"] = func() interface{} {
		return &corev1.CephFSPersistentVolumeSource{}
	}
	knownTypes["core/v1/CephFSVolumeSource"] = func() interface{} {
		return &corev1.CephFSVolumeSource{}
	}
	knownTypes["core/v1/CinderPersistentVolumeSource"] = func() interface{} {
		return &corev1.CinderPersistentVolumeSource{}
	}
	knownTypes["core/v1/CinderVolumeSource"] = func() interface{} {
		return &corev1.CinderVolumeSource{}
	}
	knownTypes["core/v1/ClientIPConfig"] = func() interface{} {
		return &corev1.ClientIPConfig{}
	}
	knownTypes["core/v1/ComponentCondition"] = func() interface{} {
		return &corev1.ComponentCondition{}
	}
	knownTypes["core/v1/ComponentStatus"] = func() interface{} {
		return &corev1.ComponentStatus{}
	}
	knownTypes["core/v1/ComponentStatusList"] = func() interface{} {
		return &corev1.ComponentStatusList{}
	}
	knownTypes["core/v1/ConfigMap"] = func() interface{} {
		return &corev1.ConfigMap{}
	}
	knownTypes["core/v1/ConfigMapEnvSource"] = func() interface{} {
		return &corev1.ConfigMapEnvSource{}
	}
	knownTypes["core/v1/ConfigMapKeySelector"] = func() interface{} {
		return &corev1.ConfigMapKeySelector{}
	}
	knownTypes["core/v1/ConfigMapList"] = func() interface{} {
		return &corev1.ConfigMapList{}
	}
	knownTypes["core/v1/ConfigMapNodeConfigSource"] = func() interface{} {
		return &corev1.ConfigMapNodeConfigSource{}
	}
	knownTypes["core/v1/ConfigMapProjection"] = func() interface{} {
		return &corev1.ConfigMapProjection{}
	}
	knownTypes["core/v1/ConfigMapVolumeSource"] = func() interface{} {
		return &corev1.ConfigMapVolumeSource{}
	}
	knownTypes["core/v1/Container"] = func() interface{} {
		return &corev1.Container{}
	}
	knownTypes["core/v1/ContainerImage"] = func() interface{} {
		return &corev1.ContainerImage{}
	}
	knownTypes["core/v1/ContainerPort"] = func() interface{} {
		return &corev1.ContainerPort{}
	}
	knownTypes["core/v1/ContainerState"] = func() interface{} {
		return &corev1.ContainerState{}
	}
	knownTypes["core/v1/ContainerStateRunning"] = func() interface{} {
		return &corev1.ContainerStateRunning{}
	}
	knownTypes["core/v1/ContainerStateTerminated"] = func() interface{} {
		return &corev1.ContainerStateTerminated{}
	}
	knownTypes["core/v1/ContainerStateWaiting"] = func() interface{} {
		return &corev1.ContainerStateWaiting{}
	}
	knownTypes["core/v1/ContainerStatus"] = func() interface{} {
		return &corev1.ContainerStatus{}
	}
	knownTypes["core/v1/DaemonEndpoint"] = func() interface{} {
		return &corev1.DaemonEndpoint{}
	}
	knownTypes["core/v1/DownwardAPIProjection"] = func() interface{} {
		return &corev1.DownwardAPIProjection{}
	}
	knownTypes["core/v1/DownwardAPIVolumeFile"] = func() interface{} {
		return &corev1.DownwardAPIVolumeFile{}
	}
	knownTypes["core/v1/DownwardAPIVolumeSource"] = func() interface{} {
		return &corev1.DownwardAPIVolumeSource{}
	}
	knownTypes["core/v1/EmptyDirVolumeSource"] = func() interface{} {
		return &corev1.EmptyDirVolumeSource{}
	}
	knownTypes["core/v1/EndpointAddress"] = func() interface{} {
		return &corev1.EndpointAddress{}
	}
	knownTypes["core/v1/EndpointPort"] = func() interface{} {
		return &corev1.EndpointPort{}
	}
	knownTypes["core/v1/EndpointSubset"] = func() interface{} {
		return &corev1.EndpointSubset{}
	}
	knownTypes["core/v1/Endpoints"] = func() interface{} {
		return &corev1.Endpoints{}
	}
	knownTypes["core/v1/EndpointsList"] = func() interface{} {
		return &corev1.EndpointsList{}
	}
	knownTypes["core/v1/EnvFromSource"] = func() interface{} {
		return &corev1.EnvFromSource{}
	}
	knownTypes["core/v1/EnvVar"] = func() interface{} {
		return &corev1.EnvVar{}
	}
	knownTypes["core/v1/EnvVarSource"] = func() interface{} {
		return &corev1.EnvVarSource{}
	}
	knownTypes["core/v1/EphemeralContainer"] = func() interface{} {
		return &corev1.EphemeralContainer{}
	}
	knownTypes["core/v1/EphemeralContainerCommon"] = func() interface{} {
		return &corev1.EphemeralContainerCommon{}
	}
	knownTypes["core/v1/EphemeralContainers"] = func() interface{} {
		return &corev1.EphemeralContainers{}
	}
	knownTypes["core/v1/Event"] = func() interface{} {
		return &corev1.Event{}
	}
	knownTypes["core/v1/EventList"] = func() interface{} {
		return &corev1.EventList{}
	}
	knownTypes["core/v1/EventSeries"] = func() interface{} {
		return &corev1.EventSeries{}
	}
	knownTypes["core/v1/EventSource"] = func() interface{} {
		return &corev1.EventSource{}
	}
	knownTypes["core/v1/ExecAction"] = func() interface{} {
		return &corev1.ExecAction{}
	}
	knownTypes["core/v1/FCVolumeSource"] = func() interface{} {
		return &corev1.FCVolumeSource{}
	}
	knownTypes["core/v1/FlexPersistentVolumeSource"] = func() interface{} {
		return &corev1.FlexPersistentVolumeSource{}
	}
	knownTypes["core/v1/FlexVolumeSource"] = func() interface{} {
		return &corev1.FlexVolumeSource{}
	}
	knownTypes["core/v1/FlockerVolumeSource"] = func() interface{} {
		return &corev1.FlockerVolumeSource{}
	}
	knownTypes["core/v1/GCEPersistentDiskVolumeSource"] = func() interface{} {
		return &corev1.GCEPersistentDiskVolumeSource{}
	}
	knownTypes["core/v1/GitRepoVolumeSource"] = func() interface{} {
		return &corev1.GitRepoVolumeSource{}
	}
	knownTypes["core/v1/GlusterfsPersistentVolumeSource"] = func() interface{} {
		return &corev1.GlusterfsPersistentVolumeSource{}
	}
	knownTypes["core/v1/GlusterfsVolumeSource"] = func() interface{} {
		return &corev1.GlusterfsVolumeSource{}
	}
	knownTypes["core/v1/HTTPGetAction"] = func() interface{} {
		return &corev1.HTTPGetAction{}
	}
	knownTypes["core/v1/HTTPHeader"] = func() interface{} {
		return &corev1.HTTPHeader{}
	}
	knownTypes["core/v1/Handler"] = func() interface{} {
		return &corev1.Handler{}
	}
	knownTypes["core/v1/HostAlias"] = func() interface{} {
		return &corev1.HostAlias{}
	}
	knownTypes["core/v1/HostPathVolumeSource"] = func() interface{} {
		return &corev1.HostPathVolumeSource{}
	}
	knownTypes["core/v1/ISCSIPersistentVolumeSource"] = func() interface{} {
		return &corev1.ISCSIPersistentVolumeSource{}
	}
	knownTypes["core/v1/ISCSIVolumeSource"] = func() interface{} {
		return &corev1.ISCSIVolumeSource{}
	}
	knownTypes["core/v1/KeyToPath"] = func() interface{} {
		return &corev1.KeyToPath{}
	}
	knownTypes["core/v1/Lifecycle"] = func() interface{} {
		return &corev1.Lifecycle{}
	}
	knownTypes["core/v1/LimitRange"] = func() interface{} {
		return &corev1.LimitRange{}
	}
	knownTypes["core/v1/LimitRangeItem"] = func() interface{} {
		return &corev1.LimitRangeItem{}
	}
	knownTypes["core/v1/LimitRangeList"] = func() interface{} {
		return &corev1.LimitRangeList{}
	}
	knownTypes["core/v1/LimitRangeSpec"] = func() interface{} {
		return &corev1.LimitRangeSpec{}
	}
	knownTypes["core/v1/List"] = func() interface{} {
		return &corev1.List{}
	}
	knownTypes["core/v1/LoadBalancerIngress"] = func() interface{} {
		return &corev1.LoadBalancerIngress{}
	}
	knownTypes["core/v1/LoadBalancerStatus"] = func() interface{} {
		return &corev1.LoadBalancerStatus{}
	}
	knownTypes["core/v1/LocalObjectReference"] = func() interface{} {
		return &corev1.LocalObjectReference{}
	}
	knownTypes["core/v1/LocalVolumeSource"] = func() interface{} {
		return &corev1.LocalVolumeSource{}
	}
	knownTypes["core/v1/NFSVolumeSource"] = func() interface{} {
		return &corev1.NFSVolumeSource{}
	}
	knownTypes["core/v1/Namespace"] = func() interface{} {
		return &corev1.Namespace{}
	}
	knownTypes["core/v1/NamespaceCondition"] = func() interface{} {
		return &corev1.NamespaceCondition{}
	}
	knownTypes["core/v1/NamespaceList"] = func() interface{} {
		return &corev1.NamespaceList{}
	}
	knownTypes["core/v1/NamespaceSpec"] = func() interface{} {
		return &corev1.NamespaceSpec{}
	}
	knownTypes["core/v1/NamespaceStatus"] = func() interface{} {
		return &corev1.NamespaceStatus{}
	}
	knownTypes["core/v1/Node"] = func() interface{} {
		return &corev1.Node{}
	}
	knownTypes["core/v1/NodeAddress"] = func() interface{} {
		return &corev1.NodeAddress{}
	}
	knownTypes["core/v1/NodeAffinity"] = func() interface{} {
		return &corev1.NodeAffinity{}
	}
	knownTypes["core/v1/NodeCondition"] = func() interface{} {
		return &corev1.NodeCondition{}
	}
	knownTypes["core/v1/NodeConfigSource"] = func() interface{} {
		return &corev1.NodeConfigSource{}
	}
	knownTypes["core/v1/NodeConfigStatus"] = func() interface{} {
		return &corev1.NodeConfigStatus{}
	}
	knownTypes["core/v1/NodeDaemonEndpoints"] = func() interface{} {
		return &corev1.NodeDaemonEndpoints{}
	}
	knownTypes["core/v1/NodeList"] = func() interface{} {
		return &corev1.NodeList{}
	}
	knownTypes["core/v1/NodeProxyOptions"] = func() interface{} {
		return &corev1.NodeProxyOptions{}
	}
	knownTypes["core/v1/NodeResources"] = func() interface{} {
		return &corev1.NodeResources{}
	}
	knownTypes["core/v1/NodeSelector"] = func() interface{} {
		return &corev1.NodeSelector{}
	}
	knownTypes["core/v1/NodeSelectorRequirement"] = func() interface{} {
		return &corev1.NodeSelectorRequirement{}
	}
	knownTypes["core/v1/NodeSelectorTerm"] = func() interface{} {
		return &corev1.NodeSelectorTerm{}
	}
	knownTypes["core/v1/NodeSpec"] = func() interface{} {
		return &corev1.NodeSpec{}
	}
	knownTypes["core/v1/NodeStatus"] = func() interface{} {
		return &corev1.NodeStatus{}
	}
	knownTypes["core/v1/NodeSystemInfo"] = func() interface{} {
		return &corev1.NodeSystemInfo{}
	}
	knownTypes["core/v1/ObjectFieldSelector"] = func() interface{} {
		return &corev1.ObjectFieldSelector{}
	}
	knownTypes["core/v1/ObjectReference"] = func() interface{} {
		return &corev1.ObjectReference{}
	}
	knownTypes["core/v1/PersistentVolume"] = func() interface{} {
		return &corev1.PersistentVolume{}
	}
	knownTypes["core/v1/PersistentVolumeClaim"] = func() interface{} {
		return &corev1.PersistentVolumeClaim{}
	}
	knownTypes["core/v1/PersistentVolumeClaimCondition"] = func() interface{} {
		return &corev1.PersistentVolumeClaimCondition{}
	}
	knownTypes["core/v1/PersistentVolumeClaimList"] = func() interface{} {
		return &corev1.PersistentVolumeClaimList{}
	}
	knownTypes["core/v1/PersistentVolumeClaimSpec"] = func() interface{} {
		return &corev1.PersistentVolumeClaimSpec{}
	}
	knownTypes["core/v1/PersistentVolumeClaimStatus"] = func() interface{} {
		return &corev1.PersistentVolumeClaimStatus{}
	}
	knownTypes["core/v1/PersistentVolumeClaimVolumeSource"] = func() interface{} {
		return &corev1.PersistentVolumeClaimVolumeSource{}
	}
	knownTypes["core/v1/PersistentVolumeList"] = func() interface{} {
		return &corev1.PersistentVolumeList{}
	}
	knownTypes["core/v1/PersistentVolumeSource"] = func() interface{} {
		return &corev1.PersistentVolumeSource{}
	}
	knownTypes["core/v1/PersistentVolumeSpec"] = func() interface{} {
		return &corev1.PersistentVolumeSpec{}
	}
	knownTypes["core/v1/PersistentVolumeStatus"] = func() interface{} {
		return &corev1.PersistentVolumeStatus{}
	}
	knownTypes["core/v1/PhotonPersistentDiskVolumeSource"] = func() interface{} {
		return &corev1.PhotonPersistentDiskVolumeSource{}
	}
	knownTypes["core/v1/Pod"] = func() interface{} {
		return &corev1.Pod{}
	}
	knownTypes["core/v1/PodAffinity"] = func() interface{} {
		return &corev1.PodAffinity{}
	}
	knownTypes["core/v1/PodAffinityTerm"] = func() interface{} {
		return &corev1.PodAffinityTerm{}
	}
	knownTypes["core/v1/PodAntiAffinity"] = func() interface{} {
		return &corev1.PodAntiAffinity{}
	}
	knownTypes["core/v1/PodAttachOptions"] = func() interface{} {
		return &corev1.PodAttachOptions{}
	}
	knownTypes["core/v1/PodCondition"] = func() interface{} {
		return &corev1.PodCondition{}
	}
	knownTypes["core/v1/PodDNSConfig"] = func() interface{} {
		return &corev1.PodDNSConfig{}
	}
	knownTypes["core/v1/PodDNSConfigOption"] = func() interface{} {
		return &corev1.PodDNSConfigOption{}
	}
	knownTypes["core/v1/PodExecOptions"] = func() interface{} {
		return &corev1.PodExecOptions{}
	}
	knownTypes["core/v1/PodIP"] = func() interface{} {
		return &corev1.PodIP{}
	}
	knownTypes["core/v1/PodList"] = func() interface{} {
		return &corev1.PodList{}
	}
	knownTypes["core/v1/PodLogOptions"] = func() interface{} {
		return &corev1.PodLogOptions{}
	}
	knownTypes["core/v1/PodPortForwardOptions"] = func() interface{} {
		return &corev1.PodPortForwardOptions{}
	}
	knownTypes["core/v1/PodProxyOptions"] = func() interface{} {
		return &corev1.PodProxyOptions{}
	}
	knownTypes["core/v1/PodReadinessGate"] = func() interface{} {
		return &corev1.PodReadinessGate{}
	}
	knownTypes["core/v1/PodSecurityContext"] = func() interface{} {
		return &corev1.PodSecurityContext{}
	}
	knownTypes["core/v1/PodSignature"] = func() interface{} {
		return &corev1.PodSignature{}
	}
	knownTypes["core/v1/PodSpec"] = func() interface{} {
		return &corev1.PodSpec{}
	}
	knownTypes["core/v1/PodStatus"] = func() interface{} {
		return &corev1.PodStatus{}
	}
	knownTypes["core/v1/PodStatusResult"] = func() interface{} {
		return &corev1.PodStatusResult{}
	}
	knownTypes["core/v1/PodTemplate"] = func() interface{} {
		return &corev1.PodTemplate{}
	}
	knownTypes["core/v1/PodTemplateList"] = func() interface{} {
		return &corev1.PodTemplateList{}
	}
	knownTypes["core/v1/PodTemplateSpec"] = func() interface{} {
		return &corev1.PodTemplateSpec{}
	}
	knownTypes["core/v1/PortworxVolumeSource"] = func() interface{} {
		return &corev1.PortworxVolumeSource{}
	}
	knownTypes["core/v1/Preconditions"] = func() interface{} {
		return &corev1.Preconditions{}
	}
	knownTypes["core/v1/PreferAvoidPodsEntry"] = func() interface{} {
		return &corev1.PreferAvoidPodsEntry{}
	}
	knownTypes["core/v1/PreferredSchedulingTerm"] = func() interface{} {
		return &corev1.PreferredSchedulingTerm{}
	}
	knownTypes["core/v1/Probe"] = func() interface{} {
		return &corev1.Probe{}
	}
	knownTypes["core/v1/ProjectedVolumeSource"] = func() interface{} {
		return &corev1.ProjectedVolumeSource{}
	}
	knownTypes["core/v1/QuobyteVolumeSource"] = func() interface{} {
		return &corev1.QuobyteVolumeSource{}
	}
	knownTypes["core/v1/RBDPersistentVolumeSource"] = func() interface{} {
		return &corev1.RBDPersistentVolumeSource{}
	}
	knownTypes["core/v1/RBDVolumeSource"] = func() interface{} {
		return &corev1.RBDVolumeSource{}
	}
	knownTypes["core/v1/RangeAllocation"] = func() interface{} {
		return &corev1.RangeAllocation{}
	}
	knownTypes["core/v1/ReplicationController"] = func() interface{} {
		return &corev1.ReplicationController{}
	}
	knownTypes["core/v1/ReplicationControllerCondition"] = func() interface{} {
		return &corev1.ReplicationControllerCondition{}
	}
	knownTypes["core/v1/ReplicationControllerList"] = func() interface{} {
		return &corev1.ReplicationControllerList{}
	}
	knownTypes["core/v1/ReplicationControllerSpec"] = func() interface{} {
		return &corev1.ReplicationControllerSpec{}
	}
	knownTypes["core/v1/ReplicationControllerStatus"] = func() interface{} {
		return &corev1.ReplicationControllerStatus{}
	}
	knownTypes["core/v1/ResourceFieldSelector"] = func() interface{} {
		return &corev1.ResourceFieldSelector{}
	}
	knownTypes["core/v1/ResourceList"] = func() interface{} {
		return &corev1.ResourceList{}
	}
	knownTypes["core/v1/ResourceQuota"] = func() interface{} {
		return &corev1.ResourceQuota{}
	}
	knownTypes["core/v1/ResourceQuotaList"] = func() interface{} {
		return &corev1.ResourceQuotaList{}
	}
	knownTypes["core/v1/ResourceQuotaSpec"] = func() interface{} {
		return &corev1.ResourceQuotaSpec{}
	}
	knownTypes["core/v1/ResourceQuotaStatus"] = func() interface{} {
		return &corev1.ResourceQuotaStatus{}
	}
	knownTypes["core/v1/ResourceRequirements"] = func() interface{} {
		return &corev1.ResourceRequirements{}
	}
	knownTypes["core/v1/SELinuxOptions"] = func() interface{} {
		return &corev1.SELinuxOptions{}
	}
	knownTypes["core/v1/ScaleIOPersistentVolumeSource"] = func() interface{} {
		return &corev1.ScaleIOPersistentVolumeSource{}
	}
	knownTypes["core/v1/ScaleIOVolumeSource"] = func() interface{} {
		return &corev1.ScaleIOVolumeSource{}
	}
	knownTypes["core/v1/ScopeSelector"] = func() interface{} {
		return &corev1.ScopeSelector{}
	}
	knownTypes["core/v1/ScopedResourceSelectorRequirement"] = func() interface{} {
		return &corev1.ScopedResourceSelectorRequirement{}
	}
	knownTypes["core/v1/Secret"] = func() interface{} {
		return &corev1.Secret{}
	}
	knownTypes["core/v1/SecretEnvSource"] = func() interface{} {
		return &corev1.SecretEnvSource{}
	}
	knownTypes["core/v1/SecretKeySelector"] = func() interface{} {
		return &corev1.SecretKeySelector{}
	}
	knownTypes["core/v1/SecretList"] = func() interface{} {
		return &corev1.SecretList{}
	}
	knownTypes["core/v1/SecretProjection"] = func() interface{} {
		return &corev1.SecretProjection{}
	}
	knownTypes["core/v1/SecretReference"] = func() interface{} {
		return &corev1.SecretReference{}
	}
	knownTypes["core/v1/SecretVolumeSource"] = func() interface{} {
		return &corev1.SecretVolumeSource{}
	}
	knownTypes["core/v1/SecurityContext"] = func() interface{} {
		return &corev1.SecurityContext{}
	}
	knownTypes["core/v1/SerializedReference"] = func() interface{} {
		return &corev1.SerializedReference{}
	}
	knownTypes["core/v1/Service"] = func() interface{} {
		return &corev1.Service{}
	}
	knownTypes["core/v1/ServiceAccount"] = func() interface{} {
		return &corev1.ServiceAccount{}
	}
	knownTypes["core/v1/ServiceAccountList"] = func() interface{} {
		return &corev1.ServiceAccountList{}
	}
	knownTypes["core/v1/ServiceAccountTokenProjection"] = func() interface{} {
		return &corev1.ServiceAccountTokenProjection{}
	}
	knownTypes["core/v1/ServiceList"] = func() interface{} {
		return &corev1.ServiceList{}
	}
	knownTypes["core/v1/ServicePort"] = func() interface{} {
		return &corev1.ServicePort{}
	}
	knownTypes["core/v1/ServiceProxyOptions"] = func() interface{} {
		return &corev1.ServiceProxyOptions{}
	}
	knownTypes["core/v1/ServiceSpec"] = func() interface{} {
		return &corev1.ServiceSpec{}
	}
	knownTypes["core/v1/ServiceStatus"] = func() interface{} {
		return &corev1.ServiceStatus{}
	}
	knownTypes["core/v1/SessionAffinityConfig"] = func() interface{} {
		return &corev1.SessionAffinityConfig{}
	}
	knownTypes["core/v1/StorageOSPersistentVolumeSource"] = func() interface{} {
		return &corev1.StorageOSPersistentVolumeSource{}
	}
	knownTypes["core/v1/StorageOSVolumeSource"] = func() interface{} {
		return &corev1.StorageOSVolumeSource{}
	}
	knownTypes["core/v1/Sysctl"] = func() interface{} {
		return &corev1.Sysctl{}
	}
	knownTypes["core/v1/TCPSocketAction"] = func() interface{} {
		return &corev1.TCPSocketAction{}
	}
	knownTypes["core/v1/Taint"] = func() interface{} {
		return &corev1.Taint{}
	}
	knownTypes["core/v1/Toleration"] = func() interface{} {
		return &corev1.Toleration{}
	}
	knownTypes["core/v1/TopologySelectorLabelRequirement"] = func() interface{} {
		return &corev1.TopologySelectorLabelRequirement{}
	}
	knownTypes["core/v1/TopologySelectorTerm"] = func() interface{} {
		return &corev1.TopologySelectorTerm{}
	}
	knownTypes["core/v1/TopologySpreadConstraint"] = func() interface{} {
		return &corev1.TopologySpreadConstraint{}
	}
	knownTypes["core/v1/TypedLocalObjectReference"] = func() interface{} {
		return &corev1.TypedLocalObjectReference{}
	}
	knownTypes["core/v1/Volume"] = func() interface{} {
		return &corev1.Volume{}
	}
	knownTypes["core/v1/VolumeDevice"] = func() interface{} {
		return &corev1.VolumeDevice{}
	}
	knownTypes["core/v1/VolumeMount"] = func() interface{} {
		return &corev1.VolumeMount{}
	}
	knownTypes["core/v1/VolumeNodeAffinity"] = func() interface{} {
		return &corev1.VolumeNodeAffinity{}
	}
	knownTypes["core/v1/VolumeProjection"] = func() interface{} {
		return &corev1.VolumeProjection{}
	}
	knownTypes["core/v1/VolumeSource"] = func() interface{} {
		return &corev1.VolumeSource{}
	}
	knownTypes["core/v1/VsphereVirtualDiskVolumeSource"] = func() interface{} {
		return &corev1.VsphereVirtualDiskVolumeSource{}
	}
	knownTypes["core/v1/WeightedPodAffinityTerm"] = func() interface{} {
		return &corev1.WeightedPodAffinityTerm{}
	}
	knownTypes["core/v1/WindowsSecurityContextOptions"] = func() interface{} {
		return &corev1.WindowsSecurityContextOptions{}
	}
}

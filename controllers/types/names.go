/* Copyright © 2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: Apache-2.0 */

package types

const (
	AntreaClusterOperatorName      = "antrea"
	AntreaAgentImageRenderKey      = "AntreaAgentImage"
	AntreaControllerImageRenderKey = "AntreaControllerImage"
	ReleaseVersion                 = "ReleaseVersion"

	AntreaAgentConfigOption    = "antrea-agent.conf"
	AntreaAgentConfigRenderKey = "AntreaAgentConfig"

	AntreaCNIConfigOption    = "antrea-cni.conflist"
	AntreaCNIConfigRenderKey = "AntreaCNIConfig"

	AntreaControllerConfigOption    = "antrea-controller.conf"
	AntreaControllerConfigRenderKey = "AntreaControllerConfig"

	ServiceCIDROption = "serviceCIDR"
	DefaultMTUOption  = "defaultMTU"

	OperatorNameSpace          = "antrea-operator"
	ClusterConfigName          = "cluster"
	OperatorConfigName         = "antrea-install"
	ClusterOperatorNetworkName = "cluster"

	AntreaNamespace                = "kube-system"
	AntreaAgentDaemonSetName       = "antrea-agent"
	AntreaControllerDeploymentName = "antrea-controller"
	AntreaConfigMapName            = "antrea-config"

	CNIConfDirRenderKey = "CNIConfDir"
	CNIBinDirRenderKey  = "CNIBinDir"
)

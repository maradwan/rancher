package scalinginput

import (
	"github.com/rancher/rancher/tests/framework/extensions/clusters/aks"
	"github.com/rancher/rancher/tests/framework/extensions/clusters/eks"
	"github.com/rancher/rancher/tests/framework/extensions/clusters/gke"
	"github.com/rancher/rancher/tests/framework/extensions/machinepools"
	nodepools "github.com/rancher/rancher/tests/framework/extensions/rke1/nodepools"
	corev1 "k8s.io/api/core/v1"
)

const (
	ConfigurationFileKey = "scalingInput"
)

type Pools struct {
	NodeLabels             map[string]string `json:"nodeLabels" yaml:"nodeLabels"`
	NodeTaints             []corev1.Taint    `json:"nodeTaints" yaml:"nodeTaints"`
	SpecifyCustomPrivateIP bool              `json:"specifyPrivateIP" yaml:"specifyPrivateIP"`
	SpecifyCustomPublicIP  bool              `json:"specifyPublicIP" yaml:"specifyPublicIP" default:"true"`
	CustomNodeNameSuffix   string            `json:"nodeNameSuffix" yaml:"nodeNameSuffix"`
}

type MachinePools struct {
	Pools
	NodeRoles *machinepools.NodeRoles `json:"nodeRoles" yaml:"nodeRoles"`
	IsSecure  bool                    `json:"isSecure" yaml:"isSecure" default:"false"`
}

type NodePools struct {
	Pools
	NodeRoles *nodepools.NodeRoles `json:"nodeRoles" yaml:"nodeRoles"`
}

type Config struct {
	NodePools    NodePools            `json:"nodePools" yaml:"nodePools"`
	MachinePools MachinePools         `json:"machinePools" yaml:"machinePools"`
	AKSNodePool  *aks.NodePool        `json:"aksNodePool" yaml:"aksNodePool"`
	EKSNodePool  *eks.NodeGroupConfig `json:"eksNodePool" yaml:"eksNodePool"`
	GKENodePool  *gke.NodePool        `json:"gkeNodePool" yaml:"gkeNodePool"`
}

package node

import (
	osdnv1 "github.com/openshift/api/network/v1"
	"github.com/openshift/library-go/pkg/network/networkutils"
	corev1 "k8s.io/api/core/v1"
)

type singleTenantPlugin struct{}

func NewSingleTenantPlugin() osdnPolicy {
	return &singleTenantPlugin{}
}

func (sp *singleTenantPlugin) Name() string {
	return networkutils.SingleTenantPluginName
}

func (sp *singleTenantPlugin) SupportsVNIDs() bool {
	return false
}

func (np *singleTenantPlugin) AllowDuplicateNetID() bool {
	return false
}

func (sp *singleTenantPlugin) Start(node *OsdnNode) error {
	otx := node.oc.NewTransaction()
	otx.AddFlow("table=27, priority=500, actions=goto_table:30")
	otx.AddFlow("table=80, priority=200, actions=output:NXM_NX_REG2[]")
	return otx.Commit()
}

func (sp *singleTenantPlugin) AddNetNamespace(netns *osdnv1.NetNamespace) {
}

func (sp *singleTenantPlugin) UpdateNetNamespace(netns *osdnv1.NetNamespace, oldNetID uint32) {
}

func (sp *singleTenantPlugin) DeleteNetNamespace(netns *osdnv1.NetNamespace) {
}

func (sp *singleTenantPlugin) SetUpPod(pod *corev1.Pod, podIP string) error {
	return nil
}

func (sp *singleTenantPlugin) GetVNID(namespace string) (uint32, error) {
	return 0, nil
}

func (sp *singleTenantPlugin) GetNamespaces(vnid uint32) []string {
	return nil
}

func (sp *singleTenantPlugin) GetMulticastEnabled(vnid uint32) bool {
	return false
}

func (sp *singleTenantPlugin) EnsureVNIDRules(vnid uint32) {
}

func (sp *singleTenantPlugin) SyncVNIDRules() {
}

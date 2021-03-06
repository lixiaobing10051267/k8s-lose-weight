/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resourcequota

import (
	"testing"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/resource"
	apitesting "k8s.io/kubernetes/pkg/api/testing"
	"k8s.io/kubernetes/pkg/apimachinery/registered"
)

func TestResourceQuotaStrategy(t *testing.T) {
	if !Strategy.NamespaceScoped() {
		t.Errorf("ResourceQuota should be namespace scoped")
	}
	if Strategy.AllowCreateOnUpdate() {
		t.Errorf("ResourceQuota should not allow create on update")
	}
	resourceQuota := &api.ResourceQuota{
		ObjectMeta: api.ObjectMeta{Name: "foo"},
		Status: api.ResourceQuotaStatus{
			Used: api.ResourceList{
				api.ResourceCPU:                    resource.MustParse("1"),
				api.ResourceMemory:                 resource.MustParse("1Gi"),
				api.ResourcePods:                   resource.MustParse("1"),
				api.ResourceServices:               resource.MustParse("1"),
				api.ResourceReplicationControllers: resource.MustParse("1"),
				api.ResourceQuotas:                 resource.MustParse("1"),
			},
			Hard: api.ResourceList{
				api.ResourceCPU:                    resource.MustParse("100"),
				api.ResourceMemory:                 resource.MustParse("4Gi"),
				api.ResourcePods:                   resource.MustParse("10"),
				api.ResourceServices:               resource.MustParse("10"),
				api.ResourceReplicationControllers: resource.MustParse("10"),
				api.ResourceQuotas:                 resource.MustParse("1"),
			},
		},
	}
	Strategy.PrepareForCreate(api.NewContext(), resourceQuota)
	if resourceQuota.Status.Used != nil {
		t.Errorf("ResourceQuota does not allow setting status on create")
	}
}

func TestSelectableFieldLabelConversions(t *testing.T) {
	apitesting.TestSelectableFieldLabelConversionsOfKind(t,
		registered.GroupOrDie(api.GroupName).GroupVersion.String(),
		"ResourceQuota",
		ResourceQuotaToSelectableFields(&api.ResourceQuota{}),
		nil,
	)
}

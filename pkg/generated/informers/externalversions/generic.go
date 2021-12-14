/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha "github.com/EdgeNet-project/edgenet/pkg/apis/apps/v1alpha"
	corev1alpha "github.com/EdgeNet-project/edgenet/pkg/apis/core/v1alpha"
	networkingv1alpha "github.com/EdgeNet-project/edgenet/pkg/apis/networking/v1alpha"
	registrationv1alpha "github.com/EdgeNet-project/edgenet/pkg/apis/registration/v1alpha"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=apps.edgenet.io, Version=v1alpha
	case v1alpha.SchemeGroupVersion.WithResource("selectivedeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha().SelectiveDeployments().Informer()}, nil

		// Group=core.edgenet.io, Version=v1alpha
	case corev1alpha.SchemeGroupVersion.WithResource("acceptableusepolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha().AcceptableUsePolicies().Informer()}, nil
	case corev1alpha.SchemeGroupVersion.WithResource("nodecontributions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha().NodeContributions().Informer()}, nil
	case corev1alpha.SchemeGroupVersion.WithResource("subnamespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha().SubNamespaces().Informer()}, nil
	case corev1alpha.SchemeGroupVersion.WithResource("tenants"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha().Tenants().Informer()}, nil
	case corev1alpha.SchemeGroupVersion.WithResource("tenantresourcequotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha().TenantResourceQuotas().Informer()}, nil

		// Group=networking.edgenet.io, Version=v1alpha
	case networkingv1alpha.SchemeGroupVersion.WithResource("vpnpeers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha().VPNPeers().Informer()}, nil

		// Group=registration.edgenet.io, Version=v1alpha
	case registrationv1alpha.SchemeGroupVersion.WithResource("emailverifications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Registration().V1alpha().EmailVerifications().Informer()}, nil
	case registrationv1alpha.SchemeGroupVersion.WithResource("rolerequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Registration().V1alpha().RoleRequests().Informer()}, nil
	case registrationv1alpha.SchemeGroupVersion.WithResource("tenantrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Registration().V1alpha().TenantRequests().Informer()}, nil
	case registrationv1alpha.SchemeGroupVersion.WithResource("userrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Registration().V1alpha().UserRequests().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

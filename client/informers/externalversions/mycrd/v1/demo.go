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

package v1

import (
	"context"
	time "time"

	mycrdv1 "github.com/wwq-2020/crd-codegen-demo/apis/mycrd/v1"
	versioned "github.com/wwq-2020/crd-codegen-demo/client/clientset/versioned"
	internalinterfaces "github.com/wwq-2020/crd-codegen-demo/client/informers/externalversions/internalinterfaces"
	v1 "github.com/wwq-2020/crd-codegen-demo/client/listers/mycrd/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DemoInformer provides access to a shared informer and lister for
// Demos.
type DemoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DemoLister
}

type demoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDemoInformer constructs a new informer for Demo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDemoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDemoInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDemoInformer constructs a new informer for Demo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDemoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MycrdV1().Demos(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MycrdV1().Demos(namespace).Watch(context.TODO(), options)
			},
		},
		&mycrdv1.Demo{},
		resyncPeriod,
		indexers,
	)
}

func (f *demoInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDemoInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *demoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mycrdv1.Demo{}, f.defaultInformer)
}

func (f *demoInformer) Lister() v1.DemoLister {
	return v1.NewDemoLister(f.Informer().GetIndexer())
}

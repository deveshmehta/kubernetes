/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package unversioned

import (
	"testing"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/testapi"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/extensions"
)

func getDSResourceName() string {
	return "daemonsets"
}

func TestListDaemonSets(t *testing.T) {
	ns := api.NamespaceAll
	c := &testClient{
		Request: testRequest{
			Method: "GET",
			Path:   testapi.Extensions.ResourcePath(getDSResourceName(), ns, ""),
		},
		Response: Response{StatusCode: 200,
			Body: &extensions.DaemonSetList{
				Items: []extensions.DaemonSet{
					{
						ObjectMeta: api.ObjectMeta{
							Name: "foo",
							Labels: map[string]string{
								"foo":  "bar",
								"name": "baz",
							},
						},
						Spec: extensions.DaemonSetSpec{
							Template: &api.PodTemplateSpec{},
						},
					},
				},
			},
		},
	}
	receivedDSs, err := c.Setup(t).Extensions().DaemonSets(ns).List(unversioned.ListOptions{})
	c.Validate(t, receivedDSs, err)

}

func TestGetDaemonSet(t *testing.T) {
	ns := api.NamespaceDefault
	c := &testClient{
		Request: testRequest{Method: "GET", Path: testapi.Extensions.ResourcePath(getDSResourceName(), ns, "foo"), Query: buildQueryValues(nil)},
		Response: Response{
			StatusCode: 200,
			Body: &extensions.DaemonSet{
				ObjectMeta: api.ObjectMeta{
					Name: "foo",
					Labels: map[string]string{
						"foo":  "bar",
						"name": "baz",
					},
				},
				Spec: extensions.DaemonSetSpec{
					Template: &api.PodTemplateSpec{},
				},
			},
		},
	}
	receivedDaemonSet, err := c.Setup(t).Extensions().DaemonSets(ns).Get("foo")
	c.Validate(t, receivedDaemonSet, err)
}

func TestGetDaemonSetWithNoName(t *testing.T) {
	ns := api.NamespaceDefault
	c := &testClient{Error: true}
	receivedPod, err := c.Setup(t).Extensions().DaemonSets(ns).Get("")
	if (err != nil) && (err.Error() != nameRequiredError) {
		t.Errorf("Expected error: %v, but got %v", nameRequiredError, err)
	}

	c.Validate(t, receivedPod, err)
}

func TestUpdateDaemonSet(t *testing.T) {
	ns := api.NamespaceDefault
	requestDaemonSet := &extensions.DaemonSet{
		ObjectMeta: api.ObjectMeta{Name: "foo", ResourceVersion: "1"},
	}
	c := &testClient{
		Request: testRequest{Method: "PUT", Path: testapi.Extensions.ResourcePath(getDSResourceName(), ns, "foo"), Query: buildQueryValues(nil)},
		Response: Response{
			StatusCode: 200,
			Body: &extensions.DaemonSet{
				ObjectMeta: api.ObjectMeta{
					Name: "foo",
					Labels: map[string]string{
						"foo":  "bar",
						"name": "baz",
					},
				},
				Spec: extensions.DaemonSetSpec{
					Template: &api.PodTemplateSpec{},
				},
			},
		},
	}
	receivedDaemonSet, err := c.Setup(t).Extensions().DaemonSets(ns).Update(requestDaemonSet)
	c.Validate(t, receivedDaemonSet, err)
}

func TestUpdateDaemonSetUpdateStatus(t *testing.T) {
	ns := api.NamespaceDefault
	requestDaemonSet := &extensions.DaemonSet{
		ObjectMeta: api.ObjectMeta{Name: "foo", ResourceVersion: "1"},
	}
	c := &testClient{
		Request: testRequest{Method: "PUT", Path: testapi.Extensions.ResourcePath(getDSResourceName(), ns, "foo") + "/status", Query: buildQueryValues(nil)},
		Response: Response{
			StatusCode: 200,
			Body: &extensions.DaemonSet{
				ObjectMeta: api.ObjectMeta{
					Name: "foo",
					Labels: map[string]string{
						"foo":  "bar",
						"name": "baz",
					},
				},
				Spec: extensions.DaemonSetSpec{
					Template: &api.PodTemplateSpec{},
				},
				Status: extensions.DaemonSetStatus{},
			},
		},
	}
	receivedDaemonSet, err := c.Setup(t).Extensions().DaemonSets(ns).UpdateStatus(requestDaemonSet)
	c.Validate(t, receivedDaemonSet, err)
}

func TestDeleteDaemon(t *testing.T) {
	ns := api.NamespaceDefault
	c := &testClient{
		Request:  testRequest{Method: "DELETE", Path: testapi.Extensions.ResourcePath(getDSResourceName(), ns, "foo"), Query: buildQueryValues(nil)},
		Response: Response{StatusCode: 200},
	}
	err := c.Setup(t).Extensions().DaemonSets(ns).Delete("foo")
	c.Validate(t, nil, err)
}

func TestCreateDaemonSet(t *testing.T) {
	ns := api.NamespaceDefault
	requestDaemonSet := &extensions.DaemonSet{
		ObjectMeta: api.ObjectMeta{Name: "foo"},
	}
	c := &testClient{
		Request: testRequest{Method: "POST", Path: testapi.Extensions.ResourcePath(getDSResourceName(), ns, ""), Body: requestDaemonSet, Query: buildQueryValues(nil)},
		Response: Response{
			StatusCode: 200,
			Body: &extensions.DaemonSet{
				ObjectMeta: api.ObjectMeta{
					Name: "foo",
					Labels: map[string]string{
						"foo":  "bar",
						"name": "baz",
					},
				},
				Spec: extensions.DaemonSetSpec{
					Template: &api.PodTemplateSpec{},
				},
			},
		},
	}
	receivedDaemonSet, err := c.Setup(t).Extensions().DaemonSets(ns).Create(requestDaemonSet)
	c.Validate(t, receivedDaemonSet, err)
}

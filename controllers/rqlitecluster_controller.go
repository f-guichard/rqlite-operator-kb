/*


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

package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pks/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	rqlitev1 "rqlite-kubebuilder/api/v1"
)

// RqliteClusterReconciler reconciles a RqliteCluster object
type RqliteClusterReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=rqlite.rqlite.fnetworks.tf,resources=rqliteclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rqlite.rqlite.fnetworks.tf,resources=rqliteclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get;update;patch

func (r *RqliteClusterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	log := r.Log.WithValues("Reconcile RqliteCluster ", req.Name, " in namespace ", req.NamespacedName)

	log.V(1).Info("Get Object Info")
	//objectInfo := new(rqlitev1.RqliteCluster{})
	objectInfo := &rqlitev1.RqliteCluster{}
	err := r.Get(context.TODO(), req.NamespacedName, objectInfo)

	if err != nil {
		log.Error(err, "Error during r.Get")
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	log.Info("Dump Object Info", "ClusterName", objectInfo.Spec.Name, "ClusterSize", objectInfo.Spec.ClusterSize)

	log.V(1).Info("Update Object Status")
	log.V(1).Info("Get Object Current Status", "NAme", objectInfo.Spec.Name, "Status", objectInfo.Status.CurrentStatus)
	if objectInfo.Status.CurrentStatus == "" {
		log.V(1).Info("Creating new RqliteCluster)
		pod := newRqliteCluster(objectInfo)
		objectInfo.Status.CurrentStatus = "OK"
	}

	log.V(1).Info("Set Object Target Status : ", "Name", objectInfo.Spec.Name, "Status ", objectInfo.Status.CurrentStatus)

	err = r.Status().Update(context.TODO(), objectInfo)
	if err != nil {
		log.Error(err, "Error during r.Status")
		return reconcile.Result{}, err
	}
	//if anything else happens
	return ctrl.Result{}, nil
}

func (r *RqliteClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rqlitev1.RqliteCluster{}).
		Complete(r)
}

func newRqliteCluster(rqliteCR *rqlitev1.RqliteCluster) $corev1.Pod {
	labels := map[string]string {
		"app": rqliteCR.Name,
	}
	return &corev1.Pod {
		ObjectMeta: metav1.ObjectMeta {
			Name: rqliteCR.Name+"-newpod",
			Namespace: rqliteCR.Namespace,
			Labels: labels,
		},
		Spec: corev1.PodSpec {
			Containers: []corev1.Container {{
				Name: "busybox",
				Image: "busybox",
				Command: strings.Split(rqliteCR.Spec.Command, ""),
			}},
			RestartPolicy: corev1.RestartPolicyOnFailure,
		},
	}
}

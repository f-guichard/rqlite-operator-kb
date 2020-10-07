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

	"github.com/go-logr/logr"
	//	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	//	"sigs.k8s.io/controller-runtime/pkg/reconcile"

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

	log.V(1).Info("Get Object Status")

	//if status doesnt match
	log.V(1).Info("Update Object Status")

	//if anything else happens
	//log.Error("unable to reconcile object")

	return ctrl.Result{}, nil
}

func (r *RqliteClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rqlitev1.RqliteCluster{}).
		Complete(r)
}

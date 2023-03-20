/*
Copyright 2022.

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

package main

import (
	"flag"
	"os"
	"time"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	securityv1alpha1 "github.com/nirmata/kyverno-aws-adapter/api/v1alpha1"
	"github.com/nirmata/kyverno-aws-adapter/controllers"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(securityv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	var syncPeriod int64
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Int64Var(&syncPeriod, "sync-period", 30, "The time interval for syncing the configuration in minutes ")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "46ad53e1.nirmata.io",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	r := &controllers.AWSAdapterConfigReconciler{
		Client:          getClient(),
		Scheme:          mgr.GetScheme(),
		RequeueInterval: time.Duration(syncPeriod) * time.Minute,
	}
	if err = r.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AWSAdapterConfig")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	createAWSAdapterConfigIfNotPresent(r)

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func getClient() client.Client {
	cl, err := client.New(ctrl.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		setupLog.Error(err, "unable to create client")
		os.Exit(1)
	}
	return cl
}

type requiredParams struct {
	clusterName      string
	clusterRegion    string
	adapterName      string
	adapterNamespace string
}

func (rp *requiredParams) areAllPresent() bool {
	return rp.clusterName != "" && rp.clusterRegion != "" && rp.adapterName != "" && rp.adapterNamespace != ""
}

func createAWSAdapterConfigIfNotPresent(r *controllers.AWSAdapterConfigReconciler) {
	rp := requiredParams{
		clusterName:      getClusterName(),
		clusterRegion:    getClusterRegion(),
		adapterName:      getAdapterName(),
		adapterNamespace: getAdapterNamespace(),
	}

	if !rp.areAllPresent() {
		setupLog.Info("One or more of the required parameters could not be found: clusterName='%s' clusterRegion='%s' adapterName='%s' adapterNamespace='%s'", rp.clusterName, rp.clusterRegion, rp.adapterName, rp.adapterNamespace)
		return
	}

	if isAWSAdapterConfigPresent, err := r.IsAWSAdapterConfigPresent(rp.adapterName, rp.adapterNamespace); err != nil {
		setupLog.Error(err, "problem checking if AWS Adapter config exists")
		os.Exit(1)
	} else if isAWSAdapterConfigPresent {
		setupLog.Info("AWS Adapter config already exists. Skipping resource creation.")
	} else {
		setupLog.Info("creating AWS Adapter config")
		if err := r.CreateAWSAdapterConfig(rp.clusterName, rp.clusterRegion, rp.adapterName, rp.adapterNamespace); err != nil {
			setupLog.Error(err, "unable to create AWS Adapter config")
			os.Exit(1)
		}
		setupLog.Info("AWS Adapter config created successfully")
	}
}

const (
	ADAPTER_NAME_ENV_VAR      = "ADAPTER_NAME"
	ADAPTER_NAMESPACE_ENV_VAR = "ADAPTER_NAMESPACE"
	CLUSTER_NAME_ENV_VAR      = "CLUSTER_NAME"
	CLUSTER_REGION_ENV_VAR    = "CLUSTER_REGION"
)

func getAdapterName() string {
	return os.Getenv(ADAPTER_NAME_ENV_VAR)
}

func getAdapterNamespace() string {
	return os.Getenv(ADAPTER_NAMESPACE_ENV_VAR)
}

func getClusterName() string {
	return os.Getenv(CLUSTER_NAME_ENV_VAR)
}

func getClusterRegion() string {
	return os.Getenv(CLUSTER_REGION_ENV_VAR)
}

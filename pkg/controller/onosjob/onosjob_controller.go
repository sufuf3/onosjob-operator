package onosjob

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	onosjobv1alpha1 "github.com/sufuf3/onosjob-operator/pkg/apis/onosjob/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_onosjob")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ONOSJob Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileONOSJob{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("onosjob-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ONOSJob
	err = c.Watch(&source.Kind{Type: &onosjobv1alpha1.ONOSJob{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ONOSJob
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &onosjobv1alpha1.ONOSJob{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileONOSJob implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileONOSJob{}

// ReconcileONOSJob reconciles a ONOSJob object
type ReconcileONOSJob struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a ONOSJob object and makes changes based on the state read
// and what is in the ONOSJob.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileONOSJob) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ONOSJob")

	// Fetch the ONOSJob instance
	instance := &onosjobv1alpha1.ONOSJob{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Pod object
	job := newJobForCR(instance)

	// Set ONOSJob instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, job, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &batchv1.Job{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: job.Name, Namespace: job.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Job", "Job.Namespace", job.Namespace, "Job.Name", job.Name)
		err = r.client.Create(context.TODO(), job)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Job already exists - don't requeue
	reqLogger.Info("Skip reconcile: Job already exists", "Job.Namespace", found.Namespace, "Job.Name", found.Name)
	return reconcile.Result{}, nil
}

// newJobForCR returns a busybox pod with the same name/namespace as the cr
func newJobForCR(cr *onosjobv1alpha1.ONOSJob) *batchv1.Job {
	labels := map[string]string{
		"app": cr.Name,
	}
	commandString := "echo starting; "
	commands := []string{}

	if cr.Spec.Hosts != nil {
		hostApiUrl := "http://" + cr.Spec.ControllerIp + ":" + cr.Spec.ControllerPort + "/onos/v1/hosts"
		fmt.Println(hostApiUrl)
		body := onosjobv1alpha1.Hosts{
			Mac:         cr.Spec.Hosts[0].Mac,
			Vlan:        cr.Spec.Hosts[0].Vlan,
			IpAddresses: []string{cr.Spec.Hosts[0].IpAddresses[0]},
			Locations: []onosjobv1alpha1.HostLocations{{
				ElementId: cr.Spec.Hosts[0].Locations[0].ElementId,
				Port:      cr.Spec.Hosts[0].Locations[0].Port,
			}},
		}
		bodyJson, _ := json.Marshal(body)
		fmt.Println(string(bodyJson))
		commandString = commandString + "curl -v POST -H Content-Type: application/json -H Accept: application/json --user onos:rocks " + hostApiUrl + " -d '" + string(bodyJson) + "'; "
	}

	if cr.Spec.FlowsDevice != nil {
		fdApiUrl := "http://" + cr.Spec.ControllerIp + ":" + cr.Spec.ControllerPort + "/onos/v1/flows/" + cr.Spec.FlowsDevice[0].Deviceid + "?appId=" + cr.Spec.FlowsDevice[0].Appid

		body := "{\"priority\": " + strconv.FormatInt(cr.Spec.FlowsDevice[0].Priority, 10) + ", \"timeout\": " + strconv.FormatInt(cr.Spec.FlowsDevice[0].Timeout, 10) + ", \"isPermanent\": " + strconv.FormatBool(cr.Spec.FlowsDevice[0].IsPermanent) + ", \"deviceId\": \"" + cr.Spec.FlowsDevice[0].Deviceid + "\", \"treatment\": { \"instructions\": [{\"type\": \"" + cr.Spec.FlowsDevice[0].Instructions[0].Type + "\", \"port\": \"" + cr.Spec.FlowsDevice[0].Instructions[0].Port + "\"}]}, \"selector\": { \"criteria\": [{ \"type\": \"" + cr.Spec.FlowsDevice[0].Criteria[0].Type + "\", \"ethType\": \"" + cr.Spec.FlowsDevice[0].Criteria[0].EthType + "\"}]}}"
		bodyJson := body

		fmt.Println(string(bodyJson))
		commandString = commandString + "curl -v POST -H Content-Type: application/json -H Accept: application/json --user onos:rocks " + fdApiUrl + " -d '" + string(bodyJson) + "'; "
	}
	if cr.Spec.Hosts == nil && cr.Spec.FlowsDevice == nil {
		commandString = commandString + "date; "
	}
	commandString = commandString + "echo done;"
	fmt.Println(commandString)
	commands = append(commands, commandString)

	return &batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-job",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "busybox",
							Image:   "everpeace/curl-jq",
							Command: []string{"/bin/sh", "-c"},
							Args:    commands,
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
				},
			},
		},
	}
}

package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/open-policy-agent/opa/util"
)

func main() {
	ctx := context.Background()

	data := `
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: "2019-05-28T17:15:53Z"
  generateName: opa-747d5d784c-
  labels:
    app: opa
    pod-template-hash: 747d5d784c
  name: opa-747d5d784c-mlkkr
  namespace: opa
  ownerReferences:
  - apiVersion: apps/v1
    blockOwnerDeletion: true
    controller: true
    kind: ReplicaSet
    name: opa-747d5d784c
    uid: 3f537d08-816c-11e9-bd63-025000000001
  resourceVersion: "221698"
  selfLink: /api/v1/namespaces/opa/pods/opa-747d5d784c-mlkkr
  uid: 3f563c43-816c-11e9-bd63-025000000001
spec:
  containers:
  - args:
    - run
    - --server
    image: openpolicyagent/opa
    imagePullPolicy: Always
    name: opa
    ports:
    - containerPort: 8181
      name: http
      protocol: TCP
    resources: {}
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-88znl
      readOnly: true
  - image: openpolicyagent/kube-mgmt:0.8
    imagePullPolicy: IfNotPresent
    name: kube-mgmt
    resources: {}
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-88znl
      readOnly: true
  dnsPolicy: ClusterFirst
  enableServiceLinks: true
  nodeName: docker-desktop
  priority: 0
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
    tolerationSeconds: 300
  volumes:
  - name: default-token-88znl
    secret:
      defaultMode: 420
      secretName: default-token-88znl
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: "2019-05-28T17:15:53Z"
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: "2019-05-29T15:40:41Z"
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: "2019-05-29T15:40:41Z"
    status: "True"
    type: ContainersReady
  - lastProbeTime: null
    lastTransitionTime: "2019-05-28T17:15:53Z"
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://a68c8d549b90073adf389c062204a12d8f786e7c3be54269c2bd390c994d21c8
    image: openpolicyagent/kube-mgmt:0.8
    imageID: docker-pullable://openpolicyagent/kube-mgmt@sha256:a54693fd75f588af0d5f5a019be39f4a1f316cdce70d0a3f72833f543f754285
    lastState: {}
    name: kube-mgmt
    ready: true
    restartCount: 0
    state:
      running:
        startedAt: "2019-05-29T15:40:41Z"
  - containerID: docker://74f28f21138687168d456fe6faed7a8171621f8eb4dd9c2e91cc9afb0810a941
    image: openpolicyagent/opa:latest
    imageID: docker-pullable://openpolicyagent/opa@sha256:9d5e59018983cfd4dc3fdd1e181e52f6dcdcb64d7652653f579ebe14734247f9
    lastState: {}
    name: opa
    ready: true
    restartCount: 0
    state:
      running:
        startedAt: "2019-05-29T15:40:40Z"
  hostIP: 192.168.65.3
  phase: Running
  podIP: 10.1.0.37
  qosClass: BestEffort
  startTime: "2019-05-28T17:15:53Z"
`
	var json map[string]interface{}

	err := util.Unmarshal([]byte(data), &json)
	if err != nil {
		// Handle error.
	}

	// Manually create the storage layer. inmem.NewFromObject returns an
	// in-memory store containing the supplied data.
	store := inmem.NewFromObject(json)

	module := "bal"
	compiler, _ := ast.CompileModules(map[string]string{
		"bla": module,
		"ble": module,
	})
	// Create new query that returns the value
	rego := rego.New(
		rego.Query("data.metadata.labels"),
		rego.Compiler(compiler),
		rego.Store(store))

	// Run evaluation.
	rs, err := rego.Eval(ctx)
	if err != nil {
		// Handle error.
	}

	// Inspect the result.
	fmt.Println("value:", rs[0].Expressions[0].Value)
}

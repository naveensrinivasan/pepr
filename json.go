package main

func podInfoJson() string {
	return `
	{
		"apiVersion": "v1",
		"kind": "Pod",
		"metadata": {
			"annotations": {
				"prometheus.io/port": "9898",
				"prometheus.io/scrape": "true"
			},
			"creationTimestamp": "2023-03-07T07:48:13Z",
			"generateName": "cool-name-podinfo-66bbff7cf4-",
			"labels": {
				"app.kubernetes.io/name": "cool-name-podinfo",
				"pod-template-hash": "66bbff7cf4",
				"zarf-agent": "patched"
			},
			"name": "cool-name-podinfo-66bbff7cf4-fwhl2",
			"namespace": "helm-releasename",
			"ownerReferences": [
				{
					"apiVersion": "apps/v1",
					"blockOwnerDeletion": true,
					"controller": true,
					"kind": "ReplicaSet",
					"name": "cool-name-podinfo-66bbff7cf4",
					"uid": "41d30484-6ccd-462b-b96a-e166095e3681"
				}
			],
			"resourceVersion": "46512",
			"uid": "883303bc-e4b7-4fa8-a576-575cc4407201"
		},
		"spec": {
			"containers": [
				{
					"command": [
						"./podinfo",
						"--port=9898",
						"--cert-path=/data/cert",
						"--port-metrics=9797",
						"--grpc-port=9999",
						"--grpc-service-name=podinfo",
						"--level=info",
						"--random-delay=false",
						"--random-error=false"
					],
					"env": [
						{
							"name": "PODINFO_UI_COLOR",
							"value": "#34577c"
						}
					],
					"image": "127.0.0.1:31999/stefanprodan/podinfo-2985051089:6.1.6",
					"imagePullPolicy": "IfNotPresent",
					"livenessProbe": {
						"exec": {
							"command": [
								"podcli",
								"check",
								"http",
								"localhost:9898/healthz"
							]
						},
						"failureThreshold": 3,
						"initialDelaySeconds": 1,
						"periodSeconds": 10,
						"successThreshold": 1,
						"timeoutSeconds": 5
					},
					"name": "podinfo",
					"ports": [
						{
							"containerPort": 9898,
							"name": "http",
							"protocol": "TCP"
						},
						{
							"containerPort": 9797,
							"name": "http-metrics",
							"protocol": "TCP"
						},
						{
							"containerPort": 9999,
							"name": "grpc",
							"protocol": "TCP"
						}
					],
					"readinessProbe": {
						"exec": {
							"command": [
								"podcli",
								"check",
								"http",
								"localhost:9898/readyz"
							]
						},
						"failureThreshold": 3,
						"initialDelaySeconds": 1,
						"periodSeconds": 10,
						"successThreshold": 1,
						"timeoutSeconds": 5
					},
					"resources": {
						"requests": {
							"cpu": "1m",
							"memory": "16Mi"
						}
					},
					"terminationMessagePath": "/dev/termination-log",
					"terminationMessagePolicy": "File",
					"volumeMounts": [
						{
							"mountPath": "/data",
							"name": "data"
						},
						{
							"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
							"name": "kube-api-access-vn9dc",
							"readOnly": true
						}
					]
				}
			],
			"dnsPolicy": "ClusterFirst",
			"enableServiceLinks": true,
			"imagePullSecrets": [
				{
					"name": "private-registry"
				}
			],
			"nodeName": "k3d-k3s-default-server-0",
			"preemptionPolicy": "PreemptLowerPriority",
			"priority": 0,
			"restartPolicy": "Always",
			"schedulerName": "default-scheduler",
			"securityContext": {},
			"serviceAccount": "default",
			"serviceAccountName": "default",
			"terminationGracePeriodSeconds": 30,
			"tolerations": [
				{
					"effect": "NoExecute",
					"key": "node.kubernetes.io/not-ready",
					"operator": "Exists",
					"tolerationSeconds": 300
				},
				{
					"effect": "NoExecute",
					"key": "node.kubernetes.io/unreachable",
					"operator": "Exists",
					"tolerationSeconds": 300
				}
			],
			"volumes": [
				{
					"emptyDir": {},
					"name": "data"
				},
				{
					"name": "kube-api-access-vn9dc",
					"projected": {
						"defaultMode": 420,
						"sources": [
							{
								"serviceAccountToken": {
									"expirationSeconds": 3607,
									"path": "token"
								}
							},
							{
								"configMap": {
									"items": [
										{
											"key": "ca.crt",
											"path": "ca.crt"
										}
									],
									"name": "kube-root-ca.crt"
								}
							},
							{
								"downwardAPI": {
									"items": [
										{
											"fieldRef": {
												"apiVersion": "v1",
												"fieldPath": "metadata.namespace"
											},
											"path": "namespace"
										}
									]
								}
							}
						]
					}
				}
			]
		},
		"status": {
			"conditions": [
				{
					"lastProbeTime": null,
					"lastTransitionTime": "2023-03-07T07:48:13Z",
					"status": "True",
					"type": "Initialized"
				},
				{
					"lastProbeTime": null,
					"lastTransitionTime": "2023-03-07T07:48:16Z",
					"status": "True",
					"type": "Ready"
				},
				{
					"lastProbeTime": null,
					"lastTransitionTime": "2023-03-07T07:48:16Z",
					"status": "True",
					"type": "ContainersReady"
				},
				{
					"lastProbeTime": null,
					"lastTransitionTime": "2023-03-07T07:48:13Z",
					"status": "True",
					"type": "PodScheduled"
				}
			],
			"containerStatuses": [
				{
					"containerID": "containerd://81b20dafbf2498a52c8ab46fa1e95ce25c06271684e590a377ee4ac9fd2a85ee",
					"image": "127.0.0.1:31999/stefanprodan/podinfo-2985051089:6.1.6",
					"imageID": "127.0.0.1:31999/stefanprodan/podinfo-2985051089@sha256:1f1fb2d1bfadac8a487106ea62b2b582bedc83fe17b3332dffe9151dfc820742",
					"lastState": {},
					"name": "podinfo",
					"ready": true,
					"restartCount": 0,
					"started": true,
					"state": {
						"running": {
							"startedAt": "2023-03-07T07:48:14Z"
						}
					}
				}
			],
			"hostIP": "172.19.0.3",
			"phase": "Running",
			"podIP": "10.42.0.26",
			"podIPs": [
				{
					"ip": "10.42.0.26"
				}
			],
			"qosClass": "Burstable",
			"startTime": "2023-03-07T07:48:13Z"
		}
	}	
	`
}

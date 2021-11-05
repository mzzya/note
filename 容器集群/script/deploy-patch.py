

# kubectl patch {"spec": {"template": {"spec": {"containers": [{"livenessProbe": {"failureThreshold": 6, "httpGet": {"path": "/actuator/health/liveness", "port": 8888, "scheme": "HTTP"}, "initialDelaySeconds": 10, "periodSeconds": 10, "successThreshold": 1, "timeoutSeconds": 5}, "name": "tr-dev-sso-api", "readinessProbe": {"failureThreshold": 3, "httpGet": {"path": "/actuator/health/readiness", "port": 8888, "scheme": "HTTP"}, "initialDelaySeconds": 10, "periodSeconds": 10, "successThreshold": 1, "timeoutSeconds": 5}, "startupProbe": {"failureThreshold": 60, "httpGet": {"path": "/actuator/health/readiness", "port": 8888, "scheme": "HTTP"}, "initialDelaySeconds": 10, "periodSeconds": 10, "successThreshold": 1, "timeoutSeconds": 5}}]}}}}


import os
import subprocess
import sys
import re
from numpy import math

from requests import request

kubeConf = "kubectl --kubeconfig ~/.kube/test.yaml -n tr "

deployList = subprocess.getstatusoutput(
    kubeConf+"get deploy -o=custom-columns=ns:.metadata.namespace,name:.metadata.name,containerName:.spec.template.spec.containers[0].name,terminationGracePeriodSeconds:.spec.template.spec.terminationGracePeriodSeconds,lifecycle:.spec.template.spec.containers[0].lifecycle,startupProbe:.spec.template.spec.containers[0].startupProbe,readinessProbe:.spec.template.spec.containers[0].readinessProbe!=nil|grep actuator")

for deployInfo in deployList[1].split('\n'):
    deploy = re.findall(r"[^\s]\S+", deployInfo)
    # print(deployInfo)
    ns = deploy[0]
    deployName = deploy[1]
    podName = deploy[2]
    # print(ns, deployName, podName)
    print(kubeConf+' patch -n '+ns+' deploy '+deployName +
          ' -p "{\\\"spec\\\":{\\\"template\\\":{\\\"spec\\\":{\\\"containers\\\":[{\\\"name\\\":\\\"'+podName+'\\\",\\\"lifecycle\\\":{\\\"preStop\\\":{\\\"exec\\\":{\\\"command\\\":[\\\"sleep\\\",\\\"10\\\"]}}}}],\\\"terminationGracePeriodSeconds\\\":50}}}}"')

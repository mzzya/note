

import os
import subprocess
import sys
import re
from numpy import math

from requests import request

kubeConf = "kubectl --kubeconfig ~/.kube/test.yaml -n web "

deployList = subprocess.getstatusoutput(
    kubeConf+"get deploy -o=custom-columns=name:.metadata.name,containerName:.spec.template.spec.containers[0].name,readinessProbe:.spec.template.spec.containers[0].readinessProbe,livenessProbe:.spec.template.spec.containers[0].livenessProbe,startupProbe:.spec.template.spec.containers[0].startupProbe|grep actuator|grep 8888")

for deployInfo in deployList[1].split('\n'):
    deploy = re.findall(r"[^\s]\S+", deployInfo)
    # print(deploy[0], deploy[1])
    print(kubeConf+' patch deploy '+deploy[0]+' -p "{\\\"spec\\\":{\\\"template\\\":{\\\"spec\\\":{\\\"containers\\\":[{\\\"livenessProbe\\\":{\\\"failureThreshold\\\":6,\\\"httpGet\\\":{\\\"path\\\":\\\"/actuator/health/liveness\\\",\\\"port\\\":8888,\\\"scheme\\\":\\\"HTTP\\\"},\\\"initialDelaySeconds\\\":10,\\\"periodSeconds\\\":10,\\\"successThreshold\\\":1,\\\"timeoutSeconds\\\":5},\\\"name\\\":\\\"'+deploy[1] +
          '\\\",\\\"readinessProbe\\\":{\\\"failureThreshold\\\":3,\\\"httpGet\\\":{\\\"path\\\":\\\"/actuator/health/readiness\\\",\\\"port\\\":8888,\\\"scheme\\\":\\\"HTTP\\\"},\\\"initialDelaySeconds\\\":10,\\\"periodSeconds\\\":10,\\\"successThreshold\\\":1,\\\"timeoutSeconds\\\":5},\\\"startupProbe\\\":{\\\"failureThreshold\\\":60,\\\"httpGet\\\":{\\\"path\\\":\\\"/actuator/health/readiness\\\",\\\"port\\\":8888,\\\"scheme\\\":\\\"HTTP\\\"},\\\"initialDelaySeconds\\\":10,\\\"periodSeconds\\\":10,\\\"successThreshold\\\":1,\\\"timeoutSeconds\\\":5}}]}}}}"')

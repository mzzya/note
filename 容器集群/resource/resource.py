import os
import subprocess
import sys
import re
from numpy import math

from requests import request

kubeConf = "kubectl --kubeconfig ~/.kube/test.yaml "
kubeNs = " -n web "

deployList = []

allocInfo = subprocess.getstatusoutput(
    kubeConf+"get deploy -o=custom-columns=name:.metadata.name,ns:.metadata.namespace,replicas:.spec.replicas,request-cpu:.spec.template.spec.containers[0].resources.requests.cpu,request-memory:.spec.template.spec.containers[0].resources.requests.memory"+kubeNs)

# print(allocInfo)
for aInfo in allocInfo[1].split('\n'):
    a = re.findall(r"[^\s]\S+", aInfo)
    deployList.append((a[0], a[2], a[3]))
    # print(aInfo, len(a), a)

# sys.exit()
# print(deployList)


podsInfo = subprocess.getstatusoutput(
    kubeConf+"top pods --use-protocol-buffers --no-headers --sort-by=memory"+kubeNs)

# print(podsInfo)


def findReq(podInfo):
    for deploy in deployList:
        if podInfo.find(deploy[0]) == 0:
            return deploy
    return ()


def getNum(res: str):
    if res.find("m") != -1:
        return int(res.replace("m", ""))
    if res.find("Mi") != -1:
        return int(res.replace("Mi", ""))
    return int(res)


for podInfo in podsInfo[1].split('\n'):
    pod = re.findall(r"[^\s]\S+", podInfo)
    reqInfo = findReq(pod[0])
    reqCpu = getNum(reqInfo[1])
    reqMem = getNum(reqInfo[2])
    usedCpu = getNum(pod[1])
    usedMem = getNum(pod[2])
    # if usedMem > reqMem:
    newMem = (math.ceil(usedMem/100))*100
    if newMem <= 0:
        newMem = 100
    if usedMem <= 20:
        newMem = 20
    print("**mem****", newMem,
          usedMem, reqMem, pod[0], reqInfo[0])
    # print(
    #     'kubectl patch deploy '+reqInfo[0]+' -p "{\\\"spec\\\":{\\\"template\\\":{\\\"spec\\\":{\\\"containers\\\":[{\\\"name\\\":\\\"'+reqInfo[0] +
    #     '\\\",\\\"resources\\\":{\\\"requests\\\":{\\\"memory\\\":\\\"' +
    #     str(newMem)+'Mi\\\"},\\\"limits\\\":{\\\"memory\\\":\\\"2Gi\\\"}}}]}}}}"')
    # if usedCpu > reqCpu:
    # print("--cpu----", (math.ceil(usedCpu/10)) *
    #       10, usedCpu, reqCpu, pod[0], reqInfo[0])
    # print(pod[0], reqInfo[0], reqCpu, usedCpu, reqMem, usedMem)

import os
import subprocess
import sys
import re
from numpy import math

from requests import request


kubeConf = "kubectl --kubeconfig ~/.kube/test.yaml "
kubeNs = " -n tr "

nsInfo = subprocess.getstatusoutput(
    kubeConf+"get ns |grep -v NAME|awk '{print $1}'")

# print(nsInfo)

notExecNsList = ["ahas-sentinel-pilot",
                 "aliyun", "appcenter", "arms", "catalog", "default", "edas", "system", "kube"]


def exists(ary, key: str):
    for item in ary:
        if key.find(item) != -1:
            return True
    return False


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
    return int(res)*1000


print("**mem****", "推荐",
      "使用", "申请", "pod名称", "deployment名称")
for ns in nsInfo[1].split('\n'):
    if exists(notExecNsList, ns):
        continue
    cmd = kubeConf +\
        "get deploy -o=custom-columns=name:.metadata.name,ns:.metadata.namespace,replicas:.spec.replicas,request-cpu:.spec.template.spec.containers[0].resources.requests.cpu,request-memory:.spec.template.spec.containers[0].resources.requests.memory -n "+ns
    allocInfo = subprocess.getstatusoutput(cmd)

    deployList = []

    for aInfo in allocInfo[1].split('\n'):
        a = re.findall(r"[^\s]+", aInfo)
        deployList.append((a[0], a[2], a[3]))

    podsInfo = subprocess.getstatusoutput(
        kubeConf+"top pods --use-protocol-buffers --no-headers --sort-by=memory -n "+ns)
    for podInfo in podsInfo[1].split('\n'):
        if podInfo.find("No resources found") != -1:
            break
        pod = re.findall(r"[^\s]+", podInfo)
        # print("pod", pod, pod[0])
        reqInfo = findReq(pod[0])
        if reqInfo == ():
            break
        reqCpu = getNum(reqInfo[1])
        reqMem = getNum(reqInfo[2])
        usedCpu = getNum(pod[1])
        usedMem = getNum(pod[2])
        # if usedMem > reqMem:
        if usedMem > 100:
            newMem = (math.floor(usedMem/100))*100
        if newMem <= 0:
            newMem = 100
        if usedMem <= 20:
            newMem = 20
        print("**mem****", newMem,
              usedMem, reqMem, pod[0], reqInfo[0])
        # print(kubeConf+kubeNs +
        #       'patch deploy '+reqInfo[0]+' -p "{\\\"spec\\\":{\\\"template\\\":{\\\"spec\\\":{\\\"containers\\\":[{\\\"name\\\":\\\"'+reqInfo[0] +
        #       '\\\",\\\"resources\\\":{\\\"requests\\\":{\\\"memory\\\":\\\"' +
        #       str(newMem)+'Mi\\\"},\\\"limits\\\":{\\\"memory\\\":\\\"2Gi\\\"}}}]}}}}"')
        # if usedCpu > reqCpu:
        # print("--cpu----", (math.ceil(usedCpu/10)) *
        #       10, usedCpu, reqCpu, pod[0], reqInfo[0])
        # print(pod[0], reqInfo[0], reqCpu, usedCpu, reqMem, usedMem)

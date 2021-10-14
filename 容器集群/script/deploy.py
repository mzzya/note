# -*- coding: UTF-8 -*-

from email.policy import default
from platform import node
import subprocess
import yaml
import json
import os

os.makedirs(name="./temp", exist_ok=True)

path = "./"

# 指定k8s集群配置文件
kubeConf = "kubectl --kubeconfig ~/.kube/test.yaml "

# 指定命名空间
kubeNs = " -n tr "

# 指定PVC名称
pvcName = "test-k8s-podtemp"


def getDeployList():
    cmdStr = kubeConf + " get deploy"+kubeNs+" -o json"
    # print("cmdStr", cmdStr)
    result = subprocess.getstatusoutput(cmdStr)
    listObj = json.loads(result[1])
    list = listObj["items"]
    return list


def getPodsCmdResult(podName, ns, cmd):
    cmdStr = kubeConf + " exec "+podName+" -n " + ns + " -- " + cmd
    result = subprocess.getstatusoutput(cmdStr)
    return result


def apply(content):
    print(content)
    cmdStr = kubeConf + " apply -f "+filePath
    result = subprocess.getstatusoutput(cmdStr)
    return result


def getPodList():
    cmdStr = kubeConf + " get pods"+kubeNs+" -o json"
    # print("cmdStr", cmdStr)
    result = subprocess.getstatusoutput(cmdStr)
    listObj = json.loads(result[1])
    list = listObj["items"]
    return list


def printPods(list):
    for pod in list:
        print(pod["metadata"]["name"], pod["metadata"]["namespace"])


deployMap = {}

for deploy in getDeployList():
    deployMap[deploy["metadata"]["name"]] = deploy

noAPPNamePodList = []
noDeployPodList = []
noLogFolderPodList = []

# 想要合并的Env列表，根据name判重
podEnvs = [{
    "name": "POD_NAME",
    "valueFrom": {
        "fieldRef": {
            "apiVersion": "v1",
            "fieldPath": "metadata.name"
        }
    }
}]

# 想要合并的Volumns列表，根据name判重
pvcVolumns = [{
    "name": pvcName,
    "persistentVolumeClaim": {
        "claimName": pvcName
    }
}]

# 想要合并的volumeMount列表，根据name和mountPath判重
pvcVolumeMounts = [{
    "mountPath": "/app/log/",
    "name": pvcName,
    "subPathExpr": "$(POD_NAME)/log/"
}, {
    "mountPath": "/app/temp_file/",
    "name": pvcName,
    "subPathExpr": "$(POD_NAME)/temp_file/"
}, {
    "mountPath": "/app/xxljob-log/",
    "name": pvcName,
    "subPathExpr": "$(POD_NAME)/xxljob_log/"
}]


def mergeObjectList(oldList, newList, fields, replace):
    for newIdx, new in enumerate(newList):
        contain = False
        for oldIdx, old in enumerate(oldList):
            allFieldsEqual = True
            for field in fields:
                allFieldsEqual = allFieldsEqual and old[field] == new[field]
                if allFieldsEqual == False:
                    break
            contain = allFieldsEqual
            if contain:
                break
        if contain:
            if replace:
                oldList[oldIdx] = new
            continue
        oldList.append(new)

def containEmptyVolume(volume):
    return volume.get("emptyDir") == None


execPodMap = {}
for pod in getPodList():
    labels = pod["metadata"]["labels"]
    podName = pod["metadata"]["name"]
    ns = pod["metadata"]["namespace"]
    deployName = ""
    if labels.get('app'):
        deployName = labels['app']
    if labels.get('app.kubernetes.io/name'):
        deployName = labels['app.kubernetes.io/name']
    if deployName == "":
        noAPPNamePodList.append(pod)
        continue
    if deployMap.get(deployName) == None:
        noDeployPodList.append(pod)
        continue
    if execPodMap.get(deployName):
        continue
    execPodMap[deployName] = True
    cmdResult = getPodsCmdResult(podName, ns, "du -h /app/log")
    if cmdResult[0] == 1:
        noLogFolderPodList.append(pod)
        continue
    print("\n\n\n\n符合条件的容器", pod["metadata"]["name"], pod["metadata"]
          ["namespace"], deployName)
    podEnvList = deployMap[deployName]["spec"]["template"]["spec"]["containers"][0]["env"]
    # print(podEnvList)
    mergeObjectList(podEnvList, podEnvs, ["name"], True)
    # print(podEnvList)
    deployMap[deployName]["spec"]["template"]["spec"]["containers"][0]["env"] = podEnvList

    volumes = []
    if deployMap[deployName]["spec"]["template"]["spec"].get("volumes"):
        volumes = deployMap[deployName]["spec"]["template"]["spec"]["volumes"]
    print("老的", volumes)
    f = filter(containEmptyVolume, volumes)
    volumes = list(f)
    print("过滤", volumes)
    mergeObjectList(volumes, pvcVolumns, ["name"], True)
    print("新的", volumes, "\r\n")
    deployMap[deployName]["spec"]["template"]["spec"]["volumes"] = volumes

    volumeMounts = []
    if deployMap[deployName]["spec"]["template"]["spec"]["containers"][0].get("volumeMounts"):
        volumeMounts = deployMap[deployName]["spec"]["template"]["spec"]["containers"][0]["volumeMounts"]
    print("老的", volumeMounts)
    mergeObjectList(volumeMounts, pvcVolumeMounts, ["mountPath"], True)
    print("新的", volumeMounts, "\r\n")
    deployMap[deployName]["spec"]["template"]["spec"]["containers"][0]["volumeMounts"] = volumeMounts

    deployJsonStr = json.dumps(deployMap[deployName])

    filePath = "./temp/"+deployName+".json"
    # 将更新文件写入到临时目录
    newJsonFile = open("./temp/"+deployName+".json", "w", encoding="utf8")
    newJsonFile.write(deployJsonStr)
    newJsonFile.close()

    # 下面是直接更新脚本
    applyResult = apply(filePath)
    if applyResult[0] == 1:
        print("更新失败", deployName, applyResult)
        break
    print("更新成功", deployName)
    # break

# print("未找到appName")
# printPods(noAPPNamePodList)
# print("非deploy的Pods")
# printPods(noDeployPodList)
# print("工作空间下没有log文件夹的Pods")
# printPods(noLogFolderPodList)

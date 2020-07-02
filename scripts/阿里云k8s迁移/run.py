# -*- coding: UTF-8 -*-

import yaml
import json
import os

path = "./"


# 无用的属性节点
delItems = {
    "metadata": {
        "annotations": None,
        "creationTimestamp": None,
        "generation": None,
        "resourceVersion": None,
        "selfLink": None,
        "uid": None,
        "namespace": "smb",
    },
    "spec": {
        "template": {"metadata": {"annotations": None, "creationTimestamp": None, }, },
        "clusterIP": None,
        "templateGeneration": None,
    },
    "status": None,
}

# 循环删除没用的属性节点


def delProp(item, delItem):
    if item is None or isinstance(delItem, str):
        return
    keys = dict.keys(delItem)
    for key in keys:
        if isinstance(delItem[key], str) and not delItem[key] is None:
            item[key] = delItem[key]
        if not (delItem[key] is None):
            if dict.__contains__(item, key):
                delProp(item[key], delItem[key])
            continue
        # print(dict.keys(item), key)
        if dict.__contains__(item, key):
            del item[key]


def main():
    for item in os.listdir(path):
        if not str.endswith(item, ".json"):
            continue
        name = os.path.splitext(item)[0]
        suffix = os.path.splitext(item)[1]
        oldFilePath = path + name + suffix
        if name.endswith("-u"):
            os.remove(oldFilePath)
            continue
        newFileYamlPath = path + name + "-u.yaml"
        newFileJsonPath = path + name + "-u.json"
        # print(name, suffix, fullpath)
        file = open(oldFilePath, "r", encoding="utf8")
        content = file.read()
        file.close()
        # 测试用yaml格式校验
        # obj = yaml.full_load(content)
        # 运行用json格式
        obj = json.loads(content)
        # print(obj["metadata"]["name"])
        # 移除status字段
        delProp(obj, delItems)
        if dict.__contains__(obj, "kind") and (
            obj["kind"] == "Deployment" or obj["kind"] == "DaemonSet"
        ):
            obj["apiVersion"] = "apps/v1"
        print(dict.keys(obj["metadata"]))
        # # 测试用yaml格式校验
        # newContent = yaml.dump(obj)
        # newYamlFile = open(newFileYamlPath, "w", encoding="utf8")
        # newYamlFile.write(newContent)
        # newYamlFile.close()
        # 运行用json格式
        newContent = json.dumps(obj)
        newJsonFile = open(newFileJsonPath, "w", encoding="utf8")
        newJsonFile.write(newContent)
        newJsonFile.close()
        os.remove(oldFilePath)


main()

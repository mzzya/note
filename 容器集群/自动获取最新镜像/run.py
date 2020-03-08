# -*- coding: UTF-8 -*-

import requests
import re
import subprocess
import os

# 想要获取的镜像
imageNames = ["mysql", "redis", "mongo", "busybox", "alpine", "hellojqk/alpine", "prom/prometheus",
              "jaegertracing/all-in-one", "grafana/grafana", "prom/prometheus", "nginx", "gitlab/gitlab-ce",
              "jenkins/jenkins", "docker", "rabbitmq", "wurstmeister/kafka", "logstash", "kibana", "elasticsearch", "jenkins"]

# 要匹配的版本号正则 纯v数字版本号
patterns = ["^v\d+\.\d+\.\d+$", "^\d+\.\d+\.\d+$",
            "^v\d+\.\d+$", "^\d+\.\d+$", "^\d+$", "^v\d+$"]

# 获取镜像列表url
url = "https://hub.docker.com/v2/repositories/{}/tags/?page_size=50&page=1"

imagesMap = {}

# 遍历镜像列表
for imageName in imageNames:
    # 获取结果
    remoteName = imageName
    if not str.__contains__(imageName, "/"):
        remoteName = "library/"+remoteName
    resp = requests.get(url.format(remoteName))
    body = []
    # print(resp.text)
    try:
        body = resp.json()
    except BaseException:
        exit(0)

    results = body["results"]

    # 最大版本号
    maxTag = ""
    # 版本号列表
    tags = []
    for item in results:
        name = item["name"]
        # 循环匹配版本号
        for pattern in patterns:
            r = re.search(pattern, name, re.I)
            if r != None:
                tags.append(name)
                break

    if len(tags) == 0:
        maxTag = "latest"
        # print("{}没有匹配的版本号请检查正则".format(imageName))
        pass
    elif len(tags) == 1:
        maxTag = tags[0]
    else:
        # 第一位最大的版本号数字
        # 第二位最大的版本号数字
        # 第三位组大的版本号数字
        firstMaxNumber = -1
        secondMaxNumber = -1
        threeMaxNumber = -1
        for i in range(len(tags)):
            tagStrs = str.split(tags[i], ".")
            if tagStrs != None:
                # 先比较第一个版本号数字
                tagStrs[0] = tagStrs[0].replace("v", "")
                tagStrs[0] = tagStrs[0].replace("V", "")
                first = int(tagStrs[0])
                if first > 100:
                    continue
                second = -1
                three = -1
                if len(tagStrs) > 1:
                    second = int(tagStrs[1])
                if len(tagStrs) > 2:
                    three = int(tagStrs[2])

                if first < firstMaxNumber:
                    continue
                elif first == firstMaxNumber:
                    # 如果第一个版本号数字相同则开始比较第二个版本数字
                    if second < secondMaxNumber:
                        continue
                    elif second == secondMaxNumber:
                        if three < threeMaxNumber:
                            continue
                        else:
                            threeMaxNumber = three
                            maxTag = tags[i]
                    else:
                        secondMaxNumber = second
                        threeMaxNumber = three
                        maxTag = tags[i]
                else:
                    firstMaxNumber = first
                    secondMaxNumber = second
                    threeMaxNumber = three
                    maxTag = tags[i]
    if maxTag == "":
        print("{}未匹配到标签".format(imageName))
        continue
    imagesMap[imageName] = maxTag
    # 获取镜像
    getCmd = "docker pull {}:{}".format(imageName, maxTag)
    print(getCmd)
    print("docker images list")
    # s = subprocess.Popen(getCmd, shell=True)
    # print(s)
    # os.system(getCmd)

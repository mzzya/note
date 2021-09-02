# 获取项目列表
# curl --header "Authorization: Bearer ${GITLAB_PERSON_AK}" "${GITLAB_ORIGIN}/api/v4/projects?pagination=keyset&id_after=0&per_page=100&order_by=id&sort=asc"
# 获取文件内容
# curl --header "Authorization: Bearer ${GITLAB_PERSON_AK}" "${GITLAB_ORIGIN}/tr/sso-api/-/raw/uat/pom.xml"

import re
from regex import findLibraryInfo, getDependencyRegex, getParentRegex
import requests

import os

import numpy

import aiohttp
import asyncio

import json

from prettytable import PrettyTable


gitlabPersonAK = os.getenv("GITLAB_PERSON_AK")
gitlabOrigin = os.getenv("GITLAB_ORIGIN")

id_after = 0
table = PrettyTable(['id', 'name', 'project'])

projectList = []


projectFilePath = "./git-project.json"
if os.path.isfile(projectFilePath):
    fd = open(projectFilePath)
    projectListData = fd.read()
    projectList = json.loads(projectListData)
    fd.close()
else:
    while(True):
        response = requests.get(
            gitlabOrigin+"/api/v4/projects?pagination=keyset&id_after="+str(id_after)+"&per_page=100&order_by=id&sort=asc", headers={"Authorization": "Bearer "+gitlabPersonAK})
        list = response.json()
        projectList = projectList+list
        print("id_after", id_after, len(list))
        if len(list) != 100:
            break
        id_after = list[99]["id"]
        # for project in list:
        #     table.add_row([project["id"], project["name"], project["web_url"]])
    projectListData = json.dumps(projectList)
    fd = open(projectFilePath, "w+")
    fd.write(projectListData)
    fd.close()
    # print(table)


print(len(projectList))

projectData = []


async def fetch(url, sem, project):
    async with sem:
        async with aiohttp.ClientSession() as session:
            async with session.get(url, headers={"Authorization": "Bearer "+gitlabPersonAK}) as response:
                html = await response.text()
                if response.status != 200:
                    return
                projectData.append(
                    {"id": project["id"], "name": project["name"], "url": url, "html": html})

loop = asyncio.get_event_loop()
sem = asyncio.Semaphore(100)


async def fetchProject(project, sem):
    await fetch(project["web_url"]+"/-/raw/uat/pom.xml", sem, project)
    await fetch(project["web_url"]+"/-/raw/test/pom.xml", sem, project)
    await fetch(project["web_url"]+"/-/raw/master/pom.xml", sem, project)

tasks = [fetchProject(project, sem)
         for project in projectList]
loop.run_until_complete(asyncio.wait(tasks))
loop.close()

for project in projectData:
    try:
        packageInfo = findLibraryInfo(
            project['html'], 'com.microsoft.sqlserver', getDependencyRegex)
        if packageInfo == None:
            print(project["id"], project["url"], 'None')
            continue
        print(project["id"], project["url"], packageInfo[1:3])
    except:
        print("发生了异常", project)

print("====================分界线====================")
for project in projectData:
    try:
        packageInfo = findLibraryInfo(
            project['html'], 'org.springframework.boot', getParentRegex)
        if packageInfo == None:
            print(project["id"], project["url"], 'None')
            continue
        print(project["id"], project["url"], packageInfo[1:3])
    except:
        print("发生了异常", project)

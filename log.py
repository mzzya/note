# -*- coding: UTF-8 -*-
import json
import time


file = open("/Users/wangyang/Downloads/111353.txt")

timeArray = []

for line in file:
    j = json.loads(line)
    timestampStr = j["timestamp"]
    timeStamp = float(timestampStr)/1000
    timeArray.append(timeStamp)
    # timeArray = time.localtime(timeStamp)
    # otherStyleTime = time.strftime("%Y-%m-%d %H:%M:%S", timeArray)
    # ymd = otherStyleTime[0:10]
    # if(hasattr(timelogin, ymd)):
    #     continue
    # timelogin[ymd] = otherStyleTime

file.close()
timeArray.sort()

timelogin = {}
for timeStamp in timeArray:
    timeArray = time.localtime(timeStamp)
    otherStyleTime = time.strftime("%Y-%m-%d %H:%M:%S", timeArray)
    ymd = otherStyleTime[0:10]
    if(ymd in timelogin):
        continue
    timelogin[ymd] = otherStyleTime
    print(ymd, otherStyleTime)


print(json.dumps(timelogin))

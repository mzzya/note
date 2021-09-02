import re
from sys import version

fd = open("./testpom.xml")
pomInfo = fd.read()


def getDependencyRegex(libraryStr):
    return r'(<dependency><groupId>'+libraryStr+r'</groupId><artifactId>(.*?)</artifactId><version>(.*?)</version>(.*?)</dependency>)'


def getParentRegex(libraryStr):
    return r'(<parent><groupId>.*?</groupId><artifactId>(.*?)</artifactId><version>(.*?)</version>)'


def findLibraryInfo(pomStr, libraryStr, getRegexStr):
    pattern = re.compile(getRegexStr(libraryStr), re.S)
    pomStr = re.sub(r'[\r\n ]', '', pomStr)
    resTuple = pattern.findall(pomStr)
    if resTuple == None or len(resTuple) == 0 or len(resTuple[0]) < 3:
        return None
    if resTuple[0][2].find("$") < 0:
        return list(resTuple[0])
    versionPropName = re.sub(r'[\$\{\}]', '', resTuple[0][2])
    # print(versionPropName)
    versionPattern = re.compile(
        r'<'+versionPropName+'>(.*?)</'+versionPropName+'>', re.S)
    versionList = versionPattern.findall(pomStr)
    if versionList == None or len(versionList) == 0:
        return list(resTuple[0])
    # print(versionList)
    resList = list(resTuple[0])
    resList[2] = versionList[0]
    return resList

# print(pomInfo)


packageInfo = findLibraryInfo(
    pomInfo, 'org.springframework.boot', getParentRegex)

print(packageInfo)

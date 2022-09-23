const util = require('util');
const child_process = require('child_process');
const { json } = require('stream/consumers');
// 调用util.promisify方法，返回一个promise,如const { stdout, stderr } = await exec('rm -rf build')
const exec = util.promisify(child_process.exec);
// const appPath = join(__dirname, 'app');


const ns = ' -n erp '
const topPodCmd = `kubectl top pods --no-headers --sort-by memory ${ns}`
const getPodWideCmd = `kubectl get pods  --no-headers -o wide ${ns}`
const getDeployCmd = `kubectl get deploy --no-headers -o=custom-columns=name:.metadata.name,ns:.metadata.namespace,podName:.spec.template.spec.containers[0].name,replicas:.spec.replicas,request-cpu:.spec.template.spec.containers[0].resources.requests.cpu,limit-cpu:.spec.template.spec.containers[0].resources.limits.cpu,request-memory:.spec.template.spec.containers[0].resources.requests.memory,limit-memory:.spec.template.spec.containers[0].resources.limits.memory ${ns}`
const getPodCmd = `kubectl get pod --no-headers -o=custom-columns=name:.metadata.name,ns:.metadata.namespace,request-cpu:.spec.containers[0].resources.requests.cpu,limit-cpu:.spec.containers[0].resources.limits.cpu,request-memory:.spec.containers[0].resources.requests.memory,limit-memory:.spec.containers[0].resources.limits.memory,nodeName:.spec.nodeName ${ns}`
const convertMem = (mem) => {
  const m = mem?.replace('Mi', '').replace('Gi', '000')
  return parseInt(m)
}
const getExecRes = async function (cmd) {
  const shellRes = await exec(cmd)
  // console.log("res", cmd, shellRes.stdout)
  const content = shellRes.stdout
  const rows = content.split('\n')
  for (let i = 0; i < rows.length; i++) {
    // rows[i]=rows[i].replace(/ +/,' ').split(' ')
    rows[i] = rows[i].split(/\s+/)
    // console.log("rows[i]",rows[i].split(/\s+/))
    // break
  }
  return rows;
}
const fn = async function () {
  const topPods = await getExecRes(topPodCmd)

  const podWides = await getExecRes(getPodWideCmd)

  const pods = topPods.map(m => {
    // console.log("podWides.find(f => f[0] == m[0])", podWides.find(f => f[0] == m[0])[6])
    return {
      deployName: m[0].replace(/-[^-]+-[^-]+$/, ''),
      podName: m[0],
      node: podWides.find(f => f[0] == m[0])[6],
      cpu: m[1],
      mem: convertMem(m[2]),
    }
  }).filter(f => f.mem > 500)

  const getDeploys = await getExecRes(getDeployCmd)
  const deploys = getDeploys.map(m => {
    return {
      deployName: m[0],
      ns: m[1],
      podName: m[2],
      replicas: m[3],
      reqCpu: m[4],
      limitCpu: m[5],
      reqMem: convertMem(m[6]),
      limitMem: convertMem(m[7]),
    }
  })
  // console.log("deploys", deploys)

  for (let i = 0; i < pods.length; i++) {
    const pod = pods[i];
    const deploy = deploys.find(f => f.podName == pod.deployName && f.deployName == pod.deployName)
    if (!deploy) {
      console.log("未找到", pod)
      continue
    }
    // if (pod.mem < deploy.reqMem || deploy.reqMem >= 500) {
    //   continue
    // }
    const tmp = {
      "spec": {
        "template": {
          "spec": {
            "containers": [
              {
                "name": deploy.deployName,
                "resources": {
                  "requests": {
                    "cpu": "20m",
                    "memory": "500Mi"
                  },
                  "limits": {
                    "cpu": "2",
                    "memory": "2Gi"
                  }
                }
              }
            ]
          }
        }
      }
    }

    // if (pod.node != "cn-shanghai.10.101.103.70") {
    //   continue
    // }
    const patchCmd = `kubectl patch ${ns} deploy ${deploy.deployName} -p "${JSON.stringify(tmp).replace(/"/g, '\\"')}"`
    console.log(pod.mem, pod.node, "申请", deploy.reqMem, "超额", pod.mem - deploy.reqMem, `${deploy.reqMem}->500Mi`, patchCmd)
    // const patchRes = await getExecRes(patchCmd)
    // console.log("patchRes", patchRes)
    // console.log({ ...pod, ...deploy })
    // console.log(JSON.stringify({ ...pod, ...deploy }))
    // console.log(`pod:${pod.podName}使用了cpu:${pod.cpu} mem:${pod.mem} deploy reqCpu:${deploy?.reqCpu} reqMem:${deploy?.reqMem}`)
  }

  // const getPods = await getExecRes(getPodCmd)
  // console.log("getPods", getPods)
}
fn()



const util = require('util');
const fs = require('fs');
const child_process = require('child_process');
const { json } = require('stream/consumers');
// 调用util.promisify方法，返回一个promise,如const { stdout, stderr } = await exec('rm -rf build')
const exec = util.promisify(child_process.exec);
// const appPath = join(__dirname, 'app');


const ns = ' -n web '
const getDeployJsonCmd = `kubectl get deploy -o json ${ns}`


const getJsonRes = async function (cmd) {
  const shellRes = await exec(cmd)
  // console.log("res", cmd, shellRes.stdout)
  const content = shellRes.stdout
  return content
}

const getExecRes = async function (cmd) {
  const shellRes = await exec(cmd)
  // console.log("res", cmd, shellRes.stdout)
  const content = shellRes.stdout
  return content
}

const ENV_POD_NAME_TMP = {
  name: 'POD_NAME',
  valueFrom: { fieldRef: { apiVersion: 'v1', fieldPath: 'metadata.name' } }
}

const VOLUMN_TMP = { name: "k8s-podtemp", persistentVolumeClaim: { claimName: "k8s-podtemp" } }


const VOLUME_TEMP_FILE_TMP = {
  mountPath: '/app/temp_file/',
  name: 'k8s-podtemp',
  subPathExpr: 'temp_file/$(POD_NAME)/'
}

const VOLUME_XXLJOB_TMP = {
  mountPath: '/app/xxljob-log/',
  name: 'k8s-podtemp',
  subPathExpr: 'xxljob-log/'
}

const save = (list, key, name, value) => {
  let index = list.findIndex(f => f[key] == name)
  if (index == -1) {
    console.log(`添加${key} \t ${name}`)
    list.push(value)
  } else {
    console.log(`修改${key} \t ${name}`)
    list[index] = value
  }
}
const fn = async function () {
  const res = await getJsonRes(getDeployJsonCmd)

  let deployRes = JSON.parse(res)

  // fs.rmdirSync('temp/')
  // fs.mkdirSync('temp/', { recursive: true });

  for (let i = 0; i < deployRes.items.length; i++) {
    const deploy = deployRes.items[i];
    const { name, namespace, labels } = deploy.metadata;
    const { volumeMounts, env } = deploy.spec.template.spec.containers[0];
    //跳过非helm java-app 部署的应用
    if (!labels || !labels['helm.sh/chart'] || labels['helm.sh/chart'].indexOf("java-app") !== 0) {
      continue
    }

    console.log(`\n正在处理 ${namespace} \t ${name} `)

    save(env, "name", "POD_NAME", ENV_POD_NAME_TMP)
    // console.log("env end", deploy.spec.template.spec.containers[0].env)

    const { volumes } = deploy.spec.template.spec;
    save(volumes, "name", "k8s-podtemp", VOLUMN_TMP)
    // console.log("volumn", deploy.spec.template.spec.volumes)

    save(volumeMounts, "mountPath", "/app/temp_file/", VOLUME_TEMP_FILE_TMP)
    save(volumeMounts, "mountPath", "/app/xxljob-log/", VOLUME_XXLJOB_TMP)

    // console.log("volumeMounts", deploy.spec.template.spec.containers[0].volumeMounts)
    console.log(`volumes: ${volumes.length} \t volumeMounts: ${volumeMounts.length}`)

    fs.writeFileSync(`temp/${namespace}-${name}.json`, JSON.stringify(deploy))
    // getExecRes(`kubectl apply -f ${}`)
    // console.log(JSON.stringify(deploy))
  }
}

fn()
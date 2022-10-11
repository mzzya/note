const util = require('util');
const fs = require('fs');
const child_process = require('child_process');
const { json } = require('stream/consumers');
// 调用util.promisify方法，返回一个promise,如const { stdout, stderr } = await exec('rm -rf build')
const exec = util.promisify(child_process.exec);
// const appPath = join(__dirname, 'app');

// cip        
// cip-job    
// cip-project
// cip-vip    
// cip-web    

// 警告 必须指定命名空间 否则会移除系统组件的必要节点亲和性
const ns = ' -n cip '
const getDeployJsonCmd = `kubectl get deploy -o json ${ns} > deploy.json`


const getJsonRes = async function (cmd) {
  const shellRes = await exec(cmd)
  // console.log("res", cmd, shellRes.stdout)
  const content = shellRes.stdout

  // return content
  return fs.readFileSync("deploy.json").toString()
}

const fn = async function () {
  const res = await getJsonRes(getDeployJsonCmd)

  let deployRes = JSON.parse(res)

  // fs.rmdirSync('temp/')
  // fs.mkdirSync('temp/', { recursive: true });

  for (let i = 0; i < deployRes.items.length; i++) {
    const deploy = deployRes.items[i];
    const { name, namespace, labels } = deploy.metadata;

    const { affinity } = deploy.spec.template.spec;
    if (!affinity) {
      continue
    }

    console.log(name, namespace, JSON.stringify(affinity))
    // delete deploy.spec.template.spec.affinity
    deploy.spec.template.spec.affinity = null
    fs.writeFileSync(`temp/${namespace}-${name}.json`, JSON.stringify(deploy))
    // getExecRes(`kubectl apply -f ${}`)
    // console.log(JSON.stringify(deploy))
  }
}

fn()
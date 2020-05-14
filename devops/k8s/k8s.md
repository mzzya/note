# k8s

## 资源介绍

- 控制器
  - Deployment 部署又称无状态集 常用
    - [脚本说明含注解](./deployment.yaml)
    - [官方文档](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
  - StatefulSet 有状态集
    - [脚本说明含注解](./statefulset.yaml)
    - [官方文档](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/)
  - DaemonSet 守护程序集
    - [脚本说明含注解](./daemonset.yaml)
    - [官方文档](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)
  - Job 任务
    - [脚本说明含注解](./job.yaml)
    - [官方文档](https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/)
  - CronJob 定时任务 依赖Job
    - [脚本说明含注解](./cronjob.yaml)
    - [官方文档](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/)
  - StatefulSet 有状态集 很少用
    - [脚本说明含注解](./deployment.yaml)
    - [官方文档](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
  - ReplicaSet 副本集
    - [脚本说明含注解](./deployment.yaml)
    - [官方文档](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
  - ReplicationController 副本控制器
    - [脚本说明含注解](./deployment.yaml)
    - [官方文档](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

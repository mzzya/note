创建三台虚拟机 IP分别是 

192.168.56.3 主 
192.168.56.4 从
192.168.56.5 从

在192.168.56.3 上执行
vim /etc/hosts
192.168.56.4 node2
192.168.56.5 node3

配置免密登录
ssh-keygen -t rsa
ssh-copy-id -i ~/.ssh/id_rsa.pub root@192.168.56.3
ssh-copy-id -i ~/.ssh/id_rsa.pub root@192.168.56.4
ssh-copy-id -i ~/.ssh/id_rsa.pub root@192.168.56.5
ssh-copy-id -i ~/.ssh/id_rsa.pub root@192.168.56.6
ssh-copy-id -i ~/.ssh/id_rsa.pub root@192.168.56.7

kubeadm token create --print-join-command|sed 's/${192.168.56.4}/${192.168.56.7}/g'

查看主机名
hostname
设置主机名称
hostnamectl set-hostname node2  #格式node+编号

centos 安装时图形界面网络选择允许外部网络访问
master 网络 
网卡一 网络地址转换 NAT
网卡二 仅主机 host-only
注意网卡顺序及网络文件顺序
网卡一路径
/etc/sysconfig/network-scripts/ifcfg-enp0s3  
网卡二路径
/etc/sysconfig/network-scripts/ifcfg-enp0s8
编辑
BOOTPROTO=static
ONBOOT=yes
IPADDR=192.168.56.3

重启网络服务
service network restart
生效

更新软件环境
yum update

cat /etc/hosts

关闭防火墙 不用开端口
systemctl stop firewalld && systemctl disable firewalld


yum -y install yum-utils



scp ~/Downloads/k8s-v1.13.0-rpms.tgz root@192.168.56.3:2379:/path/to/downloaded/file/



cd ~/
 
# 创建集群信息文件
echo """
CP0_IP=192.168.56.3
CP1_IP=192.168.56.4
CP2_IP=192.168.56.5
VIP=192.168.56.7
NET_IF=eth0
CIDR=10.244.0.0/16
""" > ./cluster-info
 
bash -c "$(curl -fsSL https://raw.githubusercontent.com/Lentil1016/kubeadm-ha/1.13.0/kubeha-gen.sh)"
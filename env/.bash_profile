# macOS Catalina
export BASH_SILENCE_DEPRECATION_WARNING=1

# openjdk 11
# export PATH="/usr/local/opt/openjdk@11/bin:$PATH"
# export CPPFLAGS="-I/usr/local/opt/openjdk@11/include"
# openjdk 最新
export PATH="/usr/local/opt/openjdk/bin:$PATH"
export CPPFLAGS="-I/usr/local/opt/openjdk/include"
# go语言的包路径
export GOPATH=~/gopath
export GOBIN=~/gopath/bin
export GO=/usr/local/go/bin
# node 10 版本
export PATH="/usr/local/opt/node@10/bin:$PATH"
export LDFLAGS="-L/usr/local/opt/node@10/lib"
export CPPFLAGS="-I/usr/local/opt/node@10/include"
# go 独立安装需要启用这个
# export PATH=$GO:$PATH
export PATH=$PATH:$GOBIN
# # istio
# export PATH=~/istio/bin:$PATH
# npm
export PATH=~/.npm-global/bin:$PATH
# coreutils
export PATH="/usr/local/opt/coreutils/libexec/gnubin:$PATH"
export MANPATH="/usr/local/opt/coreutils/libexec/gnuman:$MANPATH"
# binutils
export PATH="/usr/local/opt/binutils/bin:$PATH"
export LDFLAGS="-L/usr/local/opt/binutils/lib"
export CPPFLAGS="-I/usr/local/opt/binutils/include"
# krew
export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"
# ======================gnu========================
export PATH="/usr/local/opt/grep/libexec/gnubin:$PATH"
export PATH="/usr/local/opt/gnu-sed/libexec/gnubin:$PATH"
export PATH="/usr/local/opt/gnu-tar/libexec/gnubin:$PATH"
export PATH="/usr/local/opt/gnu-time/libexec/gnubin:$PATH"
export PATH="/usr/local/opt/ed/libexec/gnubin:$PATH"
export PATH="/usr/local/opt/findutils/libexec/gnubin:$PATH"

export PATH="/usr/local/opt/curl/bin:$PATH"
export LDFLAGS="-L/usr/local/opt/curl/lib"
export CPPFLAGS="-I/usr/local/opt/curl/include"
export PKG_CONFIG_PATH="/usr/local/opt/curl/lib/pkgconfig"
# ======================gnu========================

export GO111MODULE=on
export MICRO_REGISTRY=consul
export PATH="/usr/local/sbin:$PATH"
#lua
export LUA_PATH="~/lua/?.lua;;"
export GO15VENDOREXPERIMENT=1
export ETCDCTL_API=3
export GOPROXY="goproxy.io"
export GONOPROXY="gitlab.****.com"
export GOSUMDB="sum.golang.google.cn"
export GOPRIVATE="gitlab.****.com"
export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles
export TIME_STYLE='+%Y-%m-%d %H:%M:%S'
alias ll='gls -lht --color'

#enables colorin the terminal bash shell export
export CLICOLOR=1
#sets up thecolor scheme for list export
export LSCOLORS=gxfxcxdxbxegedabagacad
#sets up theprompt color (currently a green similar to linux terminal)
export PS1='\[\033[01;32m\]👻😀😁🤣😂😄😅😆😊😋👻\[\033[00m\]:\[\033[01;31m\]\w\[\033[00m\]\$ '
#enables colorfor iTerm
export TERM=xterm-color

# export debug_proxy="http://127.0.0.1:8899/"
# export http_proxy="http://127.0.0.1:8899"
# export https_proxy="http://127.0.0.1:8899"

export CONFIG_SOURCE=file

alias grep="grep --color=auto"
alias code='/usr/local/bin/code-insiders'
alias smb='cd ~/git/****/'
alias codeb='code ~/.bash_profile'
alias sourceb='source ~/.bash_profile'
alias note='code ~/note'
alias kcuc='kubectl config use-context'
alias h='cd ~'
alias hc='cd ~ && clear'

alias sshtr0='ssh tr@10.10.112.50'
alias sshtr1='ssh tr@10.10.112.51'
alias sshci='ssh root@10.10.22.216'

# export PATH="/usr/local/opt/python@3.8/bin:$PATH"
# export LDFLAGS="-L/usr/local/opt/python@3.8/lib"
# export PKG_CONFIG_PATH="/usr/local/opt/python@3.8/lib/pkgconfig"
alias python='/usr/local/bin/python3'
alias py='/usr/local/bin/python3'
alias pip='/usr/local/bin/pip3'
alias gt="http_proxy='http://127.0.0.1:8899' go test -v -run"

export ACM_API_KEY_ID="******"
export ACM_API_KEY_SECRET="******"

export SEARCH_API_KEY_ID="******"
export SEARCH_API_KEY_SECRET="******"

export REGISTRY_USERNAME="******"
export REGISTRY_PASSWORD="******"
export REGISTRY_HOST="******"

# bash-completion 命令补全
[[ -r "/usr/local/etc/profile.d/bash_completion.sh" ]] && . "/usr/local/etc/profile.d/bash_completion.sh"

# Created by mirror-config-china
export IOJS_ORG_MIRROR=https://npm.taobao.org/mirrors/iojs
export NODIST_IOJS_MIRROR=https://npm.taobao.org/mirrors/iojs
export NVM_IOJS_ORG_MIRROR=https://npm.taobao.org/mirrors/iojs
export NVMW_IOJS_ORG_MIRROR=https://npm.taobao.org/mirrors/iojs
export NODEJS_ORG_MIRROR=https://npm.taobao.org/mirrors/node
export NODIST_NODE_MIRROR=https://npm.taobao.org/mirrors/node
export NVM_NODEJS_ORG_MIRROR=https://npm.taobao.org/mirrors/node
export NVMW_NODEJS_ORG_MIRROR=https://npm.taobao.org/mirrors/node
export NVMW_NPM_MIRROR=https://npm.taobao.org/mirrors/npm
# End of mirror-config-china

[[ -s "/Users/wangyang/.gvm/scripts/gvm" ]] && source "/Users/wangyang/.gvm/scripts/gvm"

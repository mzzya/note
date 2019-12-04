#!/usr/bin/env bash
#
# Copyright (c) .NET Foundation and contributors. All rights reserved.
# Licensed under the MIT license. See LICENSE file in the project root for full license information.
#

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# 判断有没有sudo权限
current_userid=$(id -u)
if [ $current_userid -ne 0 ]; then
    echo "$(basename "$0") uninstallation script requires superuser privileges to run" >&2
    exit 1
fi

# this is the common suffix for all the dotnet pkgs
# .NET Core pkg文件的安装前缀
dotnet_pkg_name_suffix="com.microsoft.dotnet"
# 安装目录与配置文件
dotnet_install_root="/usr/local/share/dotnet"
dotnet_path_file="/etc/paths.d/dotnet"
dotnet_tool_path_file="/etc/paths.d/dotnet-cli-tools"

remove_dotnet_pkgs(){
    # 使用pkgutil工具查询已安装的.NET Core pkg(通过com.microsoft.dotnet查询)
    installed_pkgs=($(pkgutil --pkgs | grep $dotnet_pkg_name_suffix))

    for i in "${installed_pkgs[@]}"
    do
        echo "Removing dotnet component - \"$i\"" >&2
        #使用 pkgutil 删除.NET Core 组件
        pkgutil --force --forget "$i"
    done
}
#调用删除函数
remove_dotnet_pkgs
[ "$?" -ne 0 ] && echo "Failed to remove dotnet packages." >&2 && exit 1

echo "Deleting install root - $dotnet_install_root" >&2
# 删除文件夹及配置
rm -rf "$dotnet_install_root"
rm -f "$dotnet_path_file"
rm -f "$dotnet_tool_path_file"

echo "dotnet packages removal succeeded." >&2
exit 0
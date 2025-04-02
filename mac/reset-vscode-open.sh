brew install duti

for ext in js ts tsx json yaml xml html css md py rb php go java kotlin swift c h cpp cs sh sql csv log; do
    # 创建一个示例文件
    touch "example.$ext"

    # 获取 UTI
    uti=$(mdls -name kMDItemContentType "example.$ext" | awk -F'"' '{print $2}')
    
    if [ -n "$uti" ]; then
        echo "Setting VS Code as default for *.$ext files ($uti)"
        duti -s com.microsoft.VSCode "$uti" all
    else
        echo "Could not determine UTI for *.$ext files"
    fi

    # 删除示例文件
    rm "example.$ext"
done

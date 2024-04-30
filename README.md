# GolangStudy

1. 关于go get URL
    在工作目录下使用,生成go.sum文件,在项目中就可以使用相应的库函数

2. 关于Go mod
    使用Go Modules初始化项目 
        开启Go Modules模块打开 go env -2 GO111MODULE=on 或者 export GO111MODULE=on(设置在用户启动脚本中,需要重新打开终端)
        初始化项目,创建任意文件夹(项目,不要求一定在$GOPATH/src) go mod init github.com/Krewoe/...
        生成一个go.mod文件
        手动down --- go get URL
        自动down --- 
        go.mod文件会添加新代码 -> 当前模块依赖的版本号 // indirect 表示间接依赖 
        生成go.sum文件 
            作用:罗列当年前项目直接或间接依赖的所有模块版本,保证今后项目依赖的版本不会被篡改
            h1.hash(哈希形式) 整体项目的zip打开之后的全部文件的校验和, 如果不存在,说明依赖的库用不上
            xxx/go.mod h1:hash(go mod形式) 对go.mod文件进行校验和hash
            
    修改模块之间版本的依赖关系
        go mod edit -replace=before=after // before -> after
        

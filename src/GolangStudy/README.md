1. 导包问题
    go env -w GO111MODULE=off 关闭GOMODULE
    go env -w GO111MODULE=on 打开

2. slice 动态数组
    底层编译器二倍扩容
    s := int[]{11,22,3,4}
    t := s[i:j] -> [) from i then get j values

3. 关于type
    static type: int string ...
    concrete type: interface所指向的具体数据类型,系统看得见的类型
    所谓pair 就是指<type, value>

4. 进程/线程的数量越多,切换成本就越大,也就越浪费

5. 进程占用内存      虚拟内存4GB(32big OS)
   线程            大概4MB

6. GMP G -> goroutine, P -> processor, M -> thread 

7. Golang项目依赖Go modules(为了淘汰GOPATH)
    1.Go语言长久以来的依赖管理问题
    2.淘汰现有的GOPATH的使用模式
    3.统一社区中的其他的依赖管理工具

8. GOPATH的弊端
    无版本控制概念 go get github.com/Krewoe/Lars -u 表示下载过程(无法传达任何版本概念)
    无法同步一致第三方版本号, 别人期待的依赖库和自己需要的版本不一致
    无法指定当前项目引用的第三方版本号








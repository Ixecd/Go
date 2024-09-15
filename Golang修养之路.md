## Golang 修养之路

# 如何分析程序的运行时间与CPU利用率情况
1. shell内置指令 `time`
    - time go run demo.go
    - time ./a.out
    - 使用time之后会显示`usr time`和`sys time`, `real time` 一般 >= user + sys,因为系统还有其他进程

2. /usr/bin/time 指令
    - cd /usr/bin/time
    - -v go run test_go.go
    - 这个指令比内置的time更加详细,使用的时候需要使用绝对路径,并且要加上参数-v
    - CPU占用率,内存使用情况,Page Fault情况,进程切换情况,文件系统IO,Socket使用情况...

3. top指令
    - top指令显示当前运行中的程序占用CPU和内存的情况

# 如何分析golang程序的内存使用情况
1. top指令

2. GODEBUG和gctrace
    - 在demo中test函数执行完后,切片容器所申请的堆空间都被垃圾回收器回收了
    - 使用top指令后发现依然占用内存很高,说明**垃圾回收器回收了应用层的内存后,(可能)并不会立即将内存归还给系统

3. runtime.ReadMemStats

4. pprof
    - pprof 支持在网页上查看内存的使用情况,需要在代码中添加一个协程即可
    ```go
        import (
            "net/http",
            _ "net/http/pprof"
        )
        
        go func() {
            log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
        }()

    ```

# 如何分析Golang程序的CPU性能情况

- 性能分析必须在一个可重复的/稳定的环境下进行
    1. 机器必须闲置
        - 不要在共享硬件上进行性能分析
        - 不要在性能分析期间,在同一个机器上去浏览网页
    2. 注意省电模式和过热保护
    3. **不要使用虚拟机/共享的云主机**
- 建议购买专用的性能测试分析的硬件设备
    - 关闭电源管理,过热管理
    - 决不要升级,保证测试的一致性,以及具有可比性

- pprof

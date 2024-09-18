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


## 协程
- 在操作系统中线程分为**用户态**线程和**内核态线程**,对于CPU而言其可见的是内核态线程,一个用户态线程必须绑定一个内核态线程
- 1 : 1关系
    - 一个协程绑定到一个内核态线程上
    - 协程的创建/删除和切换的代价都有CPU完成,代价很大

- N : 1关系
    - N个协程绑定到一个内核态线程上,**协程在用户态线程完成切换,不会进入到内核态,这种切换非常的轻量快速,但是一个进程的所有协程都绑定在一个线程上**
    - 由于只有一个线程,而这N个协程都跑在一个内核态线程上,所以无法利用硬件的多核加速能力
    - 一旦一个协程阻塞了,造成线程阻塞,本进程的其他协程都无法执行,所以就没有并发能力

- M : N关系
    - M个协程绑定到N个线程上,是N:1和1:1类型的结合,克服了以上2种模型的缺点,但是实现起来较为复杂

# Go种的Goroutine
- Go语言天生支持协程,使用goroutine和channel
- Goroutine来自协程的概念,让一组**可复用**的函数(M)运行在一组线程(N)上,这样即使有协程阻塞,该线程的其他协程也可以被调度,程序员看不到这些底层细节
- 一个goroutine只占几KB,可以在有限的内存空间中创建大量的协程实现更多的并发

## 调度器

# Go之前的调度器存在性能问题,仅仅使用4年之后就被抛弃
- 基于M:N和一个队列和一个互斥锁,线程从协程队列中取协程任务会对队列加锁和解锁,由于有锁的开销,所以性能不高
- 同时其还无法高效利用缓存,比如,当前线程绑定一个协程,然而这个在这个协程执行期间又创建了一个新的协程,这说明这俩协程是相关的,所以这个新的协程应该继续跑在当前线程上,然而新创建的协程会被调度到其他线程上
- 多个线程的切换会增加系统的开销

# Go目前的调度器GMP
- 新的调度器中除了之前的**Goroutine**,**thread**新增了**processor**
- **processor**->包含运行goroutine的资源,如果线程想运行goroutine,必须先获取processor,其还包含了可运行的Goroutine队列
- **其设计思路如下**
    1. 一个全局队列:存放等待运行的Goroutine
    2. Processor的本地队列:存放的也是等待运行的Goroutine,数量优先最多256个,**创建新的Goroutine**时,优先加入到本地队列,如果本地队列满了,会把本地队列中一半的Goroutine移动到全局队列中
    3. Processor列表:在程序启动时创建,并保存在数组中,最多有**GOMAXPROCS**个(可配置),有GOMAXPROCS个本地队列
    4. M:线程想运行任务就得获取P,从P的本地队列中先获取任务,如果有任务就执行,没有任务就从全局队列中取一批任务,如果全局队列中没有Goroutine就从其他本地队列中偷一半放到自己的本地队列中,M运行G,运行完之后取下一个G

- 关于线程的数量M
    - 如果线程在执行过程中阻塞了,会由Go语言本身再创建新的线程执行,线程的数量有上限位10000,内核很难支持这么多的线程数
    - thread和processor没有绝对的数量对应关系,如果当前有一个Processor和一个thread但是thread在执行过程中阻塞了,就会由Processor再去创建一个新的线程

# Processor和thread什么时候创建
- processor:在程序开始运行时,确定了processor的数量之后,就会创建相应数量的processor
- thread:没有足够的thread来关联processor并运行其中的可运行的Gorotuine,如果thread都被阻塞了,但是任务队列中依然有很多Goroutine等待被调度,这时候就会由Processor创建新的thread

# 调度器的设计策略
- 复用线程
    1. work stealing机制
        - 当本线程无可运行的Goroutine时,尝试从其他线程绑定的Processor偷取Goroutine而不是销毁线程
    2. hand off机制
        - 当本线程因为Goroutine进行系统调用时,线程释放绑定的Processor,把其转移给其他空闲的线程执行
- 利用并行
    - **GOMAXPROCS**设置processor的数量,最多有GPMAXPROCS个线程分布在多个CPU上同时运行,GOMAXPROCS同时也限制了并发的程度

- 抢占
    - 在coroutine中要等待一个协程主动让出CPU才可以执行下一个协程,而在Go中,一个Goroutine最多占用CPU10ms,防止其他协程被饿死(这时和Cpp中coroutine的一个很大的区别)

- 全局Goroutine队列
    - 新的调度器中依然有全局队列,当本地队列为空时,优先从全局队列中获取,如果全局队列为空则通过work stealing机制从其他P的本地队列偷取G

# 调度器的生命周期
- 调度器的生命周期几乎占满了以哦个Go程序的一生,`runtime.main`的goroutine执行之前都是为调度器做准备工作,`runtime.main`的goroutine(main goroutine)运行,才是调度器的真正开始,直到runtime.main结束而结束

# 可视化GMP编程
- go tool trace
    1. trace 记录了运行的信息,提供可视化Web页面
- debug(不推荐)

# Go调度器调度场景过程

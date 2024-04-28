1. 导包问题
    go env -w GO111MODULE=off

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
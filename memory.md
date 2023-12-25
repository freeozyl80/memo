# 从一次node内存泄漏，分析线上服务内存分配



* POD 分配内存:  


	案例	
	``` code
	NAME↑                                PF   READY     RESTARTS STATUS       CPU    MEM    %CPU/R    %CPU/L    %MEM/R    %MEM/L IP                NODE              AGE        
	│ cmsproxy-1a907ec5-d54f46df5-rt97j    ●    1/1              0 Running       34    570        13         6        57        57 10.133.87.160     172.30.11.228     6d23h
	```


	- POD 分配内存 为 kube_pod_container_resource_requests_memory_bytes， 案例中未 1000mb 

	- POD 占用内存 为 container_memory_working_set_bytes， 案例中 为 570 mb ， 包括 缓冲区， 已占用内存，内核消耗

* top 内存

	- VIRT 进程中所有虚拟内存： code, data, shared libraries, swap out 到交换区和映射了的map 但尚未使用的部分
	- RES 进程所占用的所有实体内存，不包括交换出去
	- SHR 进程可读的全部共享内存，并非所有部分都包含在 RES 中。它反映了可能被其他进程共享的内存部分。
	- COD: 进程所占用的实体内存中，可执行代码所占用的内存大小。此项亦称为驻存代码集合
	- DATA 进程所占用的实体内存中，除去可执行代码所占用部分之外的内存大小。此项亦称为驻存数据集合

  名词解释：
  1. 内存交换 => 内存交换是减缓内存使用压力常用的一种技术手段，KVM等hypervisor可以显式地通知宿主机系统是否允许进行内存交换，从而平衡运行时的效率和内存大小。本文的目的是通过概括内存交换的原理，介绍KVM中的内存交换机制，进而深化对内存管理的理解。
  2. 缓冲区：高速设备与低速设备的不匹配，势必会让高速设备花时间等待低速设备，我们可以在这两者之间设立一个缓冲区。()[https://zhuanlan.zhihu.com/p/563185831]



* Node 内存：

	- Rss (操作系统分配给进程的总的内存大小)

	- heapTotal (堆内存)

		- 已分配内润，用于对象的创建和存储，对应 heapUsed (可用于heapdump 取判断)
		- 未分配的 但可用于分配的内存
		- 未分配的不能分配的内存，内存碎片

* heapdump 内存分析 

	- (参考文章)[https://www.wenjiangs.com/doc/cb4pbwvjw]
## go 打开文件参数
	/*
		  os.O_CREATE|os.O_APPEND
		  或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY
		  os.O_RDONLY // 只读
		  os.O_WRONLY // 只写
		  os.O_RDWR // 读写
		  os.O_APPEND // 追加（Append）
		  os.O_CREATE // 如果文件不存在则先创建
		  os.O_TRUNC // 文件打开时裁剪文件
		  os.O_EXCL // 和O_CREATE一起使用，文件不能存在
		  os.O_SYNC // 以同步I/O的方式打开
		第三个参数：权限(rwx:0-7)
		  0：没有任何权限
		  1：执行权限
		  2：写权限
		  3：写权限和执行权限
		  4：读权限
		  5：读权限和执行权限
		  6：读权限和写权限
		  7：读权限，写权限，执行权限
	*/
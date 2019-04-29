# share
此项目用于分享工作学习过程遇到的问题，以及解决问题的方法
***

1.[image_to_file.py](image_to_file.py)  
　功能：通过docer的image反推出dockerfile,以便大致了解镜像的构建过程  
　使用方法：pip安装docker-py.执行 python image_to_file.py XXXXX(imageID)
  
2.[procheck.go](procheck.go)  
　功能：参考node_exporter，简单实现监测系统服务状态的exporter，供prometheus收集。  
　使用方法：建议将源码编译成二进制文件，加入systemd运行。 或 ./procheck --config=procheck.json  
　访问 procheck_exporter 效果如下： http://localhost:9101/metrics　
  ```
  # HELP process_status two status of the process, up or down.
  # TYPE process_status gauge
  process_status{name="rsyslog"} up
  process_status{name="sshd"} up
  ```

# 字节跳动青训营抖音商城项目
## 代码结构
<ol>
<li>biz：hertz生成代码，handler文件夹包括了各种服务从http协议到rpc调用的数据转换处理</li>
<li>default.etcd：etcd自动保存信息</li>
<li>idl：各种服务的protobuf接口文档</li>
<li>rpc：kitex的各种微服务定义及实现</li>
<li>test：一些测试函数</li>
</ol>

## proto生成代码命令
kitex -I ../../idl -service myauth -module github.com/nihonge/tiktok -use github.com/nihonge/tiktok/kitex_gen ../../idl/user.proto

-I ./idl 表示 kitex 会在 ./idl 目录下查找 auth.proto 和它所导入的任何其他 Proto 文件（如 common.proto）。

-module github.com/yourusername/yourproject 是指定 Go 模块的路径。

./idl/auth.proto 是需要生成代码的 Proto 文件。

注意:

1.在同一个位置分别根据不同的proto生成代码会部分覆盖，比如结构体只有一个

2.服务直接起名"auth"会和代码生成工具中的结构体重名导致错误，于是改名authservice

**需要注意proto和thrift生成的代码样子略有差异**
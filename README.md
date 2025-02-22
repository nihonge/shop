# proto生成代码： kitex -I ../../idl -service auth -module github.com/nihonge/tiktok ..\..\idl\user.proto
-I ./idl 表示 kitex 会在 ./idl 目录下查找 auth.proto 和它所导入的任何其他 Proto 文件（如 common.proto）。
-module github.com/yourusername/yourproject 是指定 Go 模块的路径。
./idl/auth.proto 是需要生成代码的 Proto 文件。

bug:在同一个位置分别根据不同的proto生成代码会部分覆盖，比如结构体只有一个
## 需要注意proto和thrift生成的代码样子略有差异
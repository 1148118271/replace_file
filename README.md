### 修改 config.conf 配置文件
#### project-path 配置项为(项目工程路径)
#### replace-path 配置项为(自己修改后需要替换的文件的文件夹)
#### package-path 包结构下的具体文件内容替换某个自己修改的文件内容, 
#### 替换格式为: com/gxk/a.xml(需要替换的项目内的文件): user/gxk/xx.xml(自定义要替换的文件) 原文件名称不变

### 配置文件修改之后, 根目录下 go build 编译文件, 运行 replace_file 二进制文件即可替换



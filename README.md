# gitAutoStar使用说明

## 介绍

配合gitStar网站（218.241.135.34:88）进行自动获取要Star的列表，并对这些仓库进行自动Star
更加稳定，选项配置更灵活，不许安装任何额外的包即可运行，跨平台（windows， Linux）

## 项目介绍

1. src文件夹的Go文件为源代码
2. release文件夹中为配置文件和两个平台的可运行release版本

## 使用说明

1. 填写配置文件

```
[gitStar]
gs_name=gitStar用户名
gs_pwd=gitStar密码
git_name=github用户名
git_pwd=github密码
addr=gitStar现在的网址，因为gitStar经常换URL
delay=点每个Star的间隔时间，太频繁会被github禁止访问API，一般设置为6-8即可
```

2. Windows平台：双击运行gitStar-win.exe

3. Linux平台：

```
chmod +x gitStar-linux
./gitStar-linux
```
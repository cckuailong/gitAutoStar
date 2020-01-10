# gitAutoStar使用说明

此版本修复了一些bug，如程序有问题，请联系我346813862@qq.com，持续更新

## 介绍

配合gitStar网站（47.95.194.180）进行自动获取要Star的列表，并对这些仓库进行自动Star

- 更加稳定
- 选项配置更灵活
- 不需安装任何额外的包即可运行
- 跨平台（Win, Linux, Mac）

## 项目介绍

- src文件夹的Go文件为源代码
- release文件夹中为配置文件和三个平台的可运行release版本

## 使用说明

P.S. personal access token注册 参见
[How to create personal access token for GitHub](https://www.jianshu.com/p/df002fc555ff)

### 填写配置文件

```
[gitStar]
gs_name=gitStar用户名
gs_pwd=gitStar密码
git_name=github用户名
git_pwd=github个人token (personal access token)
addr=gitStar现在的网址，因为gitStar经常换URL
delay=点每个Star的间隔时间，太频繁会被github禁止访问API，一般设置为6-8即可
```

### 运行

- Windows平台

```
双击运行gitStar-win.exe
```

- Linux平台

```
chmod +x gitStar-linux
./gitStar-linux
```

- Mac平台

```
chmod a+x gitStar-mac
./gitStar-mac
```

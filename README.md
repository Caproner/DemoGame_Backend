# DemoGame_Backend
Demo游戏后端

## 开发准备

### Go环境配置

go版本使用1.15.11，使用go mod管理包。
该文件夹需要放在这个目录下：

```
$GOPATH/src/github.com/Caproner/
```

每次引入新的包之后进行如下操作：

```
go mod tidy
go mod vendor
```

开发环境建议使用Unix类系统，并使用VSCode或其他具备可扩展组件的IDE进行开发。
建议安装Golint（VSCode下只需要安装Go扩展即可）用于自动规范代码。

### 开发规范

代码规范遵循Golint即可，并注意不要使用难以理解的变量/函数名。
git提交规范使用简易版的angular规范，commit的备注分如下几类：

```
feat: 特性提交，涉及新功能开发的用这个前缀
fix: bug修复，涉及到功能bug修复的用这个前缀
style: 不涉及功能的代码修改用这个前缀
refactor: 代码架构级别重构用这个前缀
chore: 非代码修改用这个前缀，包括.gitignore，配置文件，数据文件等
```

#### git工作流

新特性开发的工作流如下：
1. 在`main`分支上拉一个新分支，名称为`feature/xxx`
2. 开发自测完成之后合入`dev`分支，再进行自测
3. 当前版本所有特性开发完成之后在`dev`分支上验收所有功能，之后上线并合并到`main`分支，并打对应版本tag

紧急bug修复的工作流如下：
1. 在`main`分支上拉一个新分支，名称为`hotfix/xxx`
2. 修复完成自测后合入`main`分支
3. 将`main`分支合入`dev`分支，进行自测

#### 项目架构树

项目包括如下几个文件/目录：

+ main.go：主文件，一般情况下不用动它（除非加中间件/代码整体重构/缺陷修复）
+ routers：路由层，存放web.go，作为整体服务路由。一般情况下不会有多出来的文件
+ handlers：句柄层，直接对接接口，仅负责简单的工作和调用逻辑层函数
+ services：逻辑层，对接句柄层，提供业务逻辑函数
+ utils：工具层，对接句柄层和逻辑层，提供业务无关的通用函数（例如算sha256，数据库工具，日志工具等）
+ tasks：定时任务存放在这里，由main.go来启动，下接services和utils
+ middlewares：全局中间件放这里，由main.go引入，并由routers选择在哪些路由启用哪些中间件
+ vendor：由go mod vendor自行生成和管理
+ config：配置文件和加载配置文件的代码放这里，**关键配置文件不要上传到github**
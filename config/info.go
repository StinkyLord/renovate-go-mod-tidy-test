package config

var (
	Logo = `
     _   
    | | _ 
 _  | |/ )  _    _
[_] | | /  | |  | |
 _  | | \  | |  | |
| | | |\ \ | |__| |
|_| |_| \_) \____/      v0.1
`
	Statement   = `iKurum [cyan@ikurum.cn] Open Source with MIT License`
	VersionDesc = `
 Project    ikufile
 Version    v0.1
Released    2021.11.18
 Licence    MIT
  Author    iKurum [cyan@ikurum.cn]
    Blog    https://ikurum.cn

 LearnRefer
 Author     dengsgo [dengsgo@gmail.com]
 Website    https://github.com/dengsgo/fileboy
`
	FirstRunHelp = `第一次运行 ikufile ?
你可能需要先执行 ikufile init 生成配置。
更多信息使用 ikufile help 查看帮助
`
	HelpStr = `ikufile [--yaml confilePath] [option]
Global Options:
    --yaml [-y -Y]    加载指定路径的配置文件. loads the configuration file for the specified path.
                  配置文件配置的相对路径，是相对于运行 ikufile 所在的目录 

Usage of ikufile:
    无参数 
        读取 .ikufile.yaml 配置，开始监听并工作
    init 
        初始化 ikufile, (在当前目录或指定目录)生成 .ikufile.yaml 配置文件
    exec 
        尝试运行定义的 command 命令
    daemon 
        读取 .ikufile.yaml 配置，以守护进程的方式运行在后台
    stop 
        停止守护进程
    version 
        查看当前版本信息
`
	ExampleFileGirl string = `####################
## 配置文件说明
## “当前目录”是指运行 ikufile 所在的目录;
## 使用 --yaml 命令参数可以指定 ikufile 加载配置的路径，如 "ikufile --yaml /user/f/go.yml" 或者 "ikufile --yaml ../../f/go.yml";
####################

# 监控配置
monitor:
    # 要监听的目录
    # test1       监听当前目录下 test1 目录
    # test1/test2 监听当前目录下 test1/test2 目录
    # test1,*     监听当前目录下 test1 目录及其所有子目录（递归）
    # .,*         监听当前目录及其所有子目录（递归）
    includeDirs:
        - .,*

    # 不监听的目录
    # .idea   忽略.idea目录及其所有子目录的监听
    exceptDirs:
        - .idea
        - .git
        - .vscode
        - node_modules
        - vendor

    # 监听文件的格式，此类文件更改会执行 command 中的命令
    # .go   后缀为 .go 的文件更改，会执行 command 中的命令
    # .*    所有的文件更改都会执行 command 中的命令
    types:
        - .go

    # 监听的事件类型，发生此类事件才执行 command 中的命令
    # 没有该配置默认监听所有事件
    # write   写入文件事件
    # rename  重命名文件事件
    # remove  移除文件事件
    # create  创建文件事件
    # chmod   更新文件权限事件(类unix)
    events:
        - write
        - rename
        - remove
        - create
        - chmod

# 命令
command:
    # 监听的文件有更改会执行的命令
    # 可以有多条命令，会依次执行
    # 如有多条命令，每条命令都会等待上一条命令执行完毕后才会执行
    # 如遇交互式命令，允许外部获取输入
    # 支持变量占位符,运行命令时会替换成实际值：
    #    {{file}}    文件名(如 a.txt 、test/test2/a.go)
    #    {{ext}}     文件后缀(如 .go)
    #    {{event}}   事件(上面的events, 如 write)
    #    {{changed}} 文件更新的本地时间戳(纳秒,如 1537326690523046400)
    # 变量占位符使用示例：cp {{file}} /root/sync -rf  、 myCommand --{{ext}} {{changed}}
    exec:
        - go version

    # 文件变更后命令在xx毫秒后才会执行，单位为毫秒
    # 一个变更事件(A)如果在定义的延迟时间(t)内, 又有新的文件变更事件(B), 那么A会取消执行。
    # B及以后的事件均依次类推，直到事件Z在t内没有新事件产生，Z 会执行
    # 合理设置延迟时间，将有效减少冗余和重复任务的执行
    # 如果不需要该特性，设置为 0
    delayMillSecond: 2000

# 通知器
notifier:
    # 文件更改会向该 url 发送请求（POST 一段 json 文本数据）
    # 触发请求的时机和执行 command 命令是一致的
    # 请求超时 15 秒
    # POST 格式:
    #    Content-Type: application/json;charset=UTF-8
    #    User-Agent: ikufile Net Notifier v1.16
    #    Body: {"project_folder":"/project/path","file":"main.go","changed":1576567861913824940,"ext":".go","event":"write"}
    # 例: http://example.com/notifier/ikufile-listener
    # 不启用通知，请留空 ""
    callUrl: ""

# 特殊指令
instruction:
    # 可以通过特殊的指令选项来控制 command 的行为，指令可以有多个
    # 指令选项解释：
    #   exec-when-start    ikufile启动就绪后，自动执行一次 'exec' 定义的命令
    #   should-finish      触发执行 'exec' 时(C)，如果上一次的命令(L)未退出（还在执行），会等待 L 退出（而不是强制 kill ），直到 L 有明确 exit code 才会开始执行本次命令。
    #                      在等待 L 退出时，又有新事件触发了命令执行(N)，则 C 执行取消，只会保留最后一次的 N 执行
    #   ignore-stdout      执行 'exec' 产生的 stdout 会被丢弃
    #   ignore-warn        ikufile 自身的 warn 信息会被丢弃
    #   ignore-info        ikufile 自身的 info 信息会被丢弃
    #   ignore-exec-error  执行 'exec' 出错仍继续执行下面的命令而不退出 

    #- should-finish
    #- exec-when-start
    - ignore-warn
`
)

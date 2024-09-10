## 0.0.4

更新日志：

- [功能] 支持 `-t` 指定 `txt` 文件批量运行
- [功能] 允许配置 `check` 参数决定扫描前检测存活
- [BUG] 修复某些情况 `path` 开头包含多个 `/` 问题
- [优化] 使用 `interactsh` 反连随机选择可用的服务器
- [优化] 开启 `debug` 时仅详细打印来自 `poc` 的请求
- [优化] 开启 `debug` 时请求响应限制打印长度
- [其他] 首页添加 `asciinema` 展示动画
- [其他] 测试 `follow redirect` 是否有效

感谢以下用户的贡献:

- 4ra1n (https://github.com/4ra1n)

所有可供下载的文件都由 `Github Actions` 构建

## 0.0.3

支持了 `payloads` 语法和 `interactsh` 反连

更新日志：

- [功能] 支持 `xray v2` 规范的 `payloads` 语法
- [功能] 支持 `projectdiscovery/interactsh` 反连
- [BUG] 修复 `dnslog.cn` 不走代理设置的问题
- [BUG] 指定输出类型时不会打印漏洞信息问题
- [优化] 优化输入目标非标准 `url` 的处理
- [其他] 优化 `build action` 的触发方式

感谢以下用户的贡献:

- 4ra1n (https://github.com/4ra1n)

所有可供下载的文件都由 `Github Actions` 构建

## 0.0.2

支持 `txt/json/html` 等多种 `output` 类型

更新日志：

- [功能] 支持 `output` 参数指定输出类型
- [功能] 支持 `api` 方式调用
- [优化] 记录请求响应原始字节方便输出报告
- [其他] 编译添加 `-trimpath` 参数

感谢以下用户的贡献:

- 4ra1n (https://github.com/4ra1n)

所有可供下载的文件都由 `Github Actions` 构建

## 0.0.1

初始版本

更新日志：

- 暂无

感谢以下用户的贡献:

- 4ra1n (https://github.com/4ra1n)

所有可供下载的文件都由 `Github Actions` 构建
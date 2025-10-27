# DeepSeek CLI

[![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/deantook/xx.svg)](https://github.com/deantook/xx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/deantook/xx)](https://goreportcard.com/report/github.com/deantook/xx)
[![Build Status](https://github.com/deantook/xx/workflows/Go/badge.svg)](https://github.com/deantook/xx/actions)
[![Docker](https://img.shields.io/docker/v/deantook/xx?label=docker)](https://hub.docker.com/r/deantook/xx)

一个基于 Go 的命令行工具，用于与 DeepSeek API 进行交互式对话。

## ✨ 功能特性

- 🤖 **交互式对话** - 与 DeepSeek AI 进行自然对话
- 💬 **多轮对话** - 支持上下文记忆的连续对话
- 🌊 **流式输出** - 实时显示 AI 回复，提升用户体验
- 🔧 **灵活配置** - 支持自定义 API URL 和模型选择
- 🎯 **简洁界面** - 直观的命令行交互界面
- ⚙️ **系统提示词** - 自定义 AI 行为模式
- 📝 **历史管理** - 保存、加载和管理对话历史
- 🔍 **配置持久化** - 自动保存和加载配置信息

## 🚀 快速开始

### 安装

#### 使用 go install（推荐）

```bash
go install github.com/deantook/xx@latest
```

#### 从源码构建

```bash
git clone https://github.com/deantook/xx.git
cd xx
go mod tidy
go build -o xx
```

#### 使用 Docker

```bash
docker run -it --rm deantook/xx:latest
```

### 配置 API Key

首次使用时，程序会提示您输入 DeepSeek API Key：

```bash
xx
# 程序会提示输入 API Key
```

或者使用环境变量：

```bash
export DEEPSEEK_API_KEY="your-api-key-here"
xx
```

### 开始对话

```bash
xx
```

## 📖 使用指南

### 基本命令

```bash
# 启动交互式对话
xx

# 查看帮助
xx --help

# 配置管理
xx config show          # 显示当前配置
xx config set api-key "your-key"  # 设置 API Key
xx config clear         # 清空配置

# 历史记录管理
xx history list         # 列出所有历史记录
xx history show <file>   # 查看历史记录内容
xx history delete <file> # 删除历史记录
xx history clear        # 清空所有历史记录
```

### 交互式命令

在对话过程中，您可以使用以下命令：

- `exit` 或 `quit` - 退出程序
- `clear` - 清空当前对话历史
- `save <标题>` - 保存当前对话
- `load <文件名>` - 加载历史对话
- `list` - 查看所有历史记录

### 配置选项

```bash
xx --api-key "your-key"           # API Key
xx --base-url "https://api.deepseek.com"  # API 基础 URL
xx --model "deepseek-chat"        # 使用的模型
xx --system-prompt "你是一个专业的编程助手"  # 系统提示词
```

## 🔧 配置

配置文件位置：`~/.deepseek-cli/config.json`

支持的配置项：
- `api-key`: DeepSeek API Key
- `base-url`: API 基础 URL（默认：https://api.deepseek.com）
- `model`: 使用的模型（默认：deepseek-chat）
- `system-prompt`: 系统提示词

## 📁 项目结构

```
.
├── main.go              # 程序入口
├── go.mod               # Go 模块文件
├── cmd/
│   └── root.go         # CLI 命令定义
├── client/
│   └── deepseek.go     # DeepSeek API 客户端
├── chat/
│   └── session.go      # 对话会话管理
├── config/
│   └── config.go       # 配置管理
├── history/
│   └── history.go      # 历史记录管理
├── .github/workflows/  # GitHub Actions
├── Dockerfile          # Docker 配置
├── LICENSE             # MIT 许可证
├── CONTRIBUTING.md     # 贡献指南
├── CHANGELOG.md        # 更新日志
├── SECURITY.md         # 安全政策
├── FAQ.md              # 常见问题
└── README.md           # 项目说明
```

## 🤝 贡献

我们欢迎任何形式的贡献！请查看 [CONTRIBUTING.md](CONTRIBUTING.md) 了解如何参与项目开发。

### 开发环境设置

1. 安装 Go 1.25+
2. Fork 并克隆项目
3. 安装依赖：`go mod tidy`
4. 运行测试：`go test ./...`
5. 构建项目：`go build -o xx`

## 📄 许可证

本项目采用 [MIT 许可证](LICENSE)。

## 🔗 相关链接

- [DeepSeek API 文档](https://platform.deepseek.com/api-docs/)
- [Go 官方文档](https://golang.org/doc/)
- [Cobra CLI 框架](https://github.com/spf13/cobra)

## 📊 项目状态

- ✅ 基础对话功能
- ✅ 流式输出
- ✅ 配置管理
- ✅ 历史记录
- ✅ 多模型支持
- 🔄 持续改进中...

## ❓ 常见问题

遇到问题？请查看 [FAQ.md](FAQ.md) 获取常见问题的解答。

## 🔒 安全

请查看 [SECURITY.md](SECURITY.md) 了解安全政策和报告漏洞的方式。

---

如果这个项目对您有帮助，请给我们一个 ⭐️！
<!-- 342db409-d986-4c1e-9caf-61b86a6f931a 3a6b1688-ebe8-4467-9e29-fa04a371612f -->
# DeepSeek CLI 对话工具

## 概述

创建一个基于 Go 的命令行工具,用于与 DeepSeek API 进行交互式对话。

## 技术选型

- HTTP 客户端: 标准库 `net/http`
- CLI 框架: `github.com/spf13/cobra`
- 环境变量管理: 标准库 `os`
- JSON 处理: 标准库 `encoding/json`

## 核心功能

### 1. API 客户端封装

创建 `client/deepseek.go`,实现:

- DeepSeek API 的 HTTP 请求封装
- Bearer Token 认证
- 流式响应处理 (SSE)
- 错误处理

### 2. 对话管理

创建 `chat/session.go`,实现:

- 多轮对话历史管理
- 消息列表维护 (system/user/assistant)
- 对话上下文保持

### 3. CLI 命令

创建 `cmd/root.go` 和 `main.go`,实现:

- 交互式对话模式
- 从环境变量读取 API Key (DEEPSEEK_API_KEY)
- 命令行参数支持 (模型选择、API URL 等)
- 用户友好的输入输出界面

### 4. 关键文件结构

```
/Users/dean/code/xx/
├── main.go                 # 程序入口
├── go.mod                  # 依赖管理
├── cmd/
│   └── root.go            # CLI 命令定义
├── client/
│   └── deepseek.go        # API 客户端
└── chat/
    └── session.go         # 对话会话管理
```

## 实现要点

1. **API 调用**: 基于文档中的 `/chat/completions` 端点,使用 `stream: true` 实现实时响应
2. **认证**: 使用 Bearer Token 从 `DEEPSEEK_API_KEY` 环境变量获取
3. **默认模型**: `deepseek-chat`
4. **API 基础 URL**: `https://api.deepseek.com`
5. **交互体验**: 

   - 支持输入 `exit` 或 `quit` 退出
   - 支持 `clear` 清空对话历史
   - 流式显示 AI 回复

## 依赖包

- `github.com/spf13/cobra`: CLI 框架
- 其他使用 Go 标准库

### To-dos

- [ ] 初始化项目结构和依赖管理
- [ ] 实现 DeepSeek API 客户端 (client/deepseek.go)
- [ ] 实现对话会话管理 (chat/session.go)
- [ ] 实现 CLI 命令和交互界面 (cmd/root.go, main.go)
- [ ] 添加 README 文档说明使用方法
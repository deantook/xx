# DeepSeek CLI 对话工具

一个基于 Go 的命令行工具，用于与 DeepSeek API 进行交互式对话。

## 功能特性

- 🤖 与 DeepSeek AI 进行交互式对话
- 💬 支持多轮对话历史
- 🌊 流式输出，实时显示 AI 回复
- 🔧 支持自定义 API URL 和模型
- 🎯 简洁的命令行界面

## 安装

### 从源码构建

1. 克隆或下载项目代码
2. 安装依赖：
   ```bash
   go mod tidy
   ```
3. 构建程序：
   ```bash
   go build -o deepseek-cli
   ```

## 使用方法

### 1. 设置 API Key

有多种方式设置 API Key（按优先级排序）：

**方式一：配置文件（推荐）**
```bash
# 首次运行时会自动提示输入并保存
./deepseek-cli

# 或使用配置命令设置
./deepseek-cli config set api-key "your-api-key-here"
```

**方式二：环境变量**
```bash
export DEEPSEEK_API_KEY="your-api-key-here"
```

**方式三：命令行参数**
```bash
./deepseek-cli --api-key "your-api-key-here"
```

### 2. 运行程序

```bash
./deepseek-cli
```

### 3. 开始对话

程序启动后，你可以直接输入问题与 DeepSeek AI 对话：

```
🤖 DeepSeek CLI 对话工具
输入 'exit' 或 'quit' 退出，输入 'clear' 清空对话历史
================================================
👤 你: 你好，请介绍一下你自己
🤖 DeepSeek: 你好！我是 DeepSeek，一个由深度求索公司开发的 AI 助手...

👤 你: 你能帮我写一个 Python 函数吗？
🤖 DeepSeek: 当然可以！请告诉我你需要什么样的 Python 函数...
```

### 4. 特殊命令

- `exit` 或 `quit`: 退出程序
- `clear`: 清空当前对话历史

### 5. 配置管理

程序提供了完整的配置管理功能：

**查看当前配置**
```bash
./deepseek-cli config show
```

**设置配置项**
```bash
# 设置 API Key
./deepseek-cli config set api-key "your-api-key-here"

# 设置基础 URL
./deepseek-cli config set base-url "https://api.deepseek.com"

# 设置模型
./deepseek-cli config set model "deepseek-chat"
```

**清空配置**
```bash
./deepseek-cli config clear
```

配置文件位置：`~/.deepseek-cli/config.json`

## 命令行参数

```bash
./deepseek-cli [选项]

选项:
  --api-key string    DeepSeek API Key (也可以通过 DEEPSEEK_API_KEY 环境变量设置)
  --base-url string   DeepSeek API 基础 URL (默认: https://api.deepseek.com)
  --model string      使用的模型名称 (默认: deepseek-chat)
```

## 示例用法

### 使用自定义模型
```bash
./deepseek-cli --model deepseek-reasoner
```

### 使用自定义 API URL
```bash
./deepseek-cli --base-url https://your-custom-api.com
```

## 项目结构

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
└── README.md           # 项目说明
```

## 依赖

- [cobra](https://github.com/spf13/cobra): CLI 框架
- Go 标准库

## 注意事项

1. 请确保你有有效的 DeepSeek API Key
2. 程序使用流式输出，AI 回复会实时显示
3. 对话历史会保存在内存中，重启程序后会丢失
4. 支持多轮对话，AI 会记住之前的对话内容

## 故障排除

### API Key 错误
如果遇到 API Key 相关错误，请检查：
- 环境变量 `DEEPSEEK_API_KEY` 是否正确设置
- API Key 是否有效且未过期

### 网络连接问题
如果遇到网络连接问题，请检查：
- 网络连接是否正常
- API URL 是否正确
- 防火墙设置是否阻止了连接

## 许可证

本项目采用 MIT 许可证。

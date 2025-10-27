# 常见问题

## 安装问题

### Q: 如何安装最新版本？

A: 使用以下命令安装最新版本：

```bash
go install github.com/deantook/xx@latest
```

### Q: 安装后找不到命令？

A: 确保 `$GOPATH/bin` 或 `$GOBIN` 在您的 PATH 环境变量中：

```bash
# 检查 Go 环境
go env GOPATH
go env GOBIN

# 添加到 PATH（如果使用 bash）
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc

# 添加到 PATH（如果使用 zsh）
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

## 配置问题

### Q: 如何设置 API Key？

A: 有多种方式设置 API Key：

1. **首次运行时设置**：
   ```bash
   xx
   # 程序会提示输入 API Key
   ```

2. **使用配置命令**：
   ```bash
   xx config set api-key "your-api-key-here"
   ```

3. **使用环境变量**：
   ```bash
   export DEEPSEEK_API_KEY="your-api-key-here"
   xx
   ```

### Q: 配置文件在哪里？

A: 配置文件位置：`~/.deepseek-cli/config.json`

### Q: 如何重置配置？

A: 使用以下命令清空所有配置：

```bash
xx config clear
```

## 使用问题

### Q: 如何保存对话历史？

A: 在对话过程中使用 `save` 命令：

```
👤 你: save 我的编程问题
✅ 对话已保存: 我的编程问题 (2024-01-15_14-30-25.json)
```

### Q: 如何加载之前的对话？

A: 使用 `load` 命令加载历史对话：

```
👤 你: load 2024-01-15_14-30-25.json
✅ 历史记录已加载: 2024-01-15_14-30-25.json
💬 加载了 6 条消息
```

### Q: 如何查看所有历史记录？

A: 使用 `list` 命令或 `xx history list`：

```bash
xx history list
```

### Q: 如何清空当前对话？

A: 在对话过程中输入 `clear`：

```
👤 你: clear
🧹 对话历史已清空
```

## API 问题

### Q: API Key 无效怎么办？

A: 请检查：

1. API Key 是否正确复制
2. API Key 是否已过期
3. 账户是否有足够的额度
4. 网络连接是否正常

### Q: 连接超时怎么办？

A: 可以尝试：

1. 检查网络连接
2. 使用自定义 API URL：
   ```bash
   xx --base-url "https://api.deepseek.com"
   ```
3. 检查防火墙设置

### Q: 支持哪些模型？

A: 目前支持 DeepSeek 的所有可用模型，默认使用 `deepseek-chat`。您可以通过以下方式指定模型：

```bash
xx --model "deepseek-reasoner"
```

## 故障排除

### Q: 程序启动失败？

A: 请检查：

1. Go 版本是否 >= 1.21
2. 依赖是否正确安装：`go mod tidy`
3. 是否有足够的磁盘空间
4. 配置文件权限是否正确

### Q: 输出乱码？

A: 确保终端支持 UTF-8 编码：

```bash
# Linux/macOS
export LANG=en_US.UTF-8

# Windows (PowerShell)
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
```

### Q: 历史记录丢失？

A: 历史记录保存在 `~/.deepseek-cli/history/` 目录下，请检查：

1. 目录是否存在
2. 文件权限是否正确
3. 磁盘空间是否充足

## 性能问题

### Q: 响应速度慢？

A: 可能的原因：

1. 网络延迟
2. API 服务器负载
3. 模型复杂度
4. 本地网络问题

可以尝试使用更快的网络或更换 API 端点。

### Q: 内存使用过高？

A: 长时间对话会占用更多内存，可以：

1. 定期使用 `clear` 清空对话
2. 保存重要对话后重新启动程序
3. 使用 `save` 命令保存对话历史

## 其他问题

### Q: 如何贡献代码？

A: 请查看 [CONTRIBUTING.md](CONTRIBUTING.md) 了解详细的贡献指南。

### Q: 如何报告 Bug？

A: 请在 GitHub 上创建 Issue，详细描述：

1. 操作系统和版本
2. Go 版本
3. 重现步骤
4. 错误信息
5. 期望行为

### Q: 如何获取帮助？

A: 您可以通过以下方式获取帮助：

1. 查看本文档
2. 在 GitHub 上创建 Issue
3. 查看源代码注释
4. 联系维护者

---

如果这里没有您遇到的问题，请随时创建 Issue 或联系维护者！

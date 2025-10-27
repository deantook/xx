# 贡献指南

感谢您对 DeepSeek CLI 项目的关注！我们欢迎任何形式的贡献。

## 如何贡献

### 报告问题

如果您发现了 bug 或有功能建议，请：

1. 检查 [Issues](https://github.com/deantook/xx/issues) 中是否已有相关问题
2. 如果没有，请创建新的 Issue，详细描述问题或建议
3. 提供复现步骤（如果是 bug）

### 提交代码

1. **Fork 项目**
   ```bash
   # 在 GitHub 上 Fork 项目
   ```

2. **克隆您的 Fork**
   ```bash
   git clone https://github.com/your-username/xx.git
   cd xx
   ```

3. **创建功能分支**
   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **进行开发**
   - 编写代码
   - 添加测试（如果适用）
   - 确保代码通过所有测试

5. **提交更改**
   ```bash
   git add .
   git commit -m "feat: 添加新功能描述"
   ```

6. **推送分支**
   ```bash
   git push origin feature/your-feature-name
   ```

7. **创建 Pull Request**
   - 在 GitHub 上创建 Pull Request
   - 详细描述您的更改
   - 等待代码审查

## 代码规范

### Go 代码规范

- 遵循 [Go 官方代码规范](https://golang.org/doc/effective_go.html)
- 使用 `gofmt` 格式化代码
- 使用 `golint` 检查代码风格
- 添加适当的注释

### 提交信息规范

使用 [Conventional Commits](https://www.conventionalcommits.org/) 规范：

- `feat:` 新功能
- `fix:` 修复 bug
- `docs:` 文档更新
- `style:` 代码格式调整
- `refactor:` 代码重构
- `test:` 测试相关
- `chore:` 构建过程或辅助工具的变动

示例：
```
feat: 添加历史记录保存功能
fix: 修复 API 连接超时问题
docs: 更新安装说明
```

## 开发环境设置

1. **安装 Go**
   - 确保 Go 版本 >= 1.21

2. **安装依赖**
   ```bash
   go mod tidy
   ```

3. **运行测试**
   ```bash
   go test ./...
   ```

4. **构建项目**
   ```bash
   go build -o xx
   ```

## 测试

- 为新功能添加测试
- 确保所有测试通过
- 保持测试覆盖率

## 文档

- 更新相关文档
- 添加代码注释
- 更新 README.md（如果需要）

## 许可证

通过贡献代码，您同意您的贡献将在 MIT 许可证下发布。

## 联系方式

如果您有任何问题，请：

- 创建 Issue
- 发送邮件到 [your-email@example.com]

感谢您的贡献！🎉

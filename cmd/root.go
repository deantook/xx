package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"xx/chat"
	"xx/config"

	"github.com/spf13/cobra"
)

var (
	apiKey       string
	baseURL      string
	model        string
	systemPrompt string
)

// rootCmd 表示基础命令
var rootCmd = &cobra.Command{
	Use:   "xx",
	Short: "DeepSeek CLI 对话工具",
	Long:  `一个基于 Go 的命令行工具，用于与 DeepSeek API 进行交互式对话。`,
	Run:   runChat,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "DeepSeek API Key (也可以通过 DEEPSEEK_API_KEY 环境变量设置)")
	rootCmd.PersistentFlags().StringVar(&baseURL, "base-url", "https://api.deepseek.com", "DeepSeek API 基础 URL")
	rootCmd.PersistentFlags().StringVar(&model, "model", "deepseek-chat", "使用的模型名称")
	rootCmd.PersistentFlags().StringVar(&systemPrompt, "system-prompt", "", "系统提示词")

	// 添加配置管理子命令
	rootCmd.AddCommand(configCmd)

	// 添加历史管理子命令
	rootCmd.AddCommand(historyCmd)
}

// configCmd 配置管理命令
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置管理",
	Long:  `管理 DeepSeek CLI 的配置，包括 API Key、基础 URL 和模型设置。`,
}

// setCmd 设置配置命令
var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "设置配置项",
	Long:  `设置配置项，支持: api-key, base-url, model`,
	Args:  cobra.ExactArgs(2),
	Run:   runSetConfig,
}

// showCmd 显示配置命令
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "显示当前配置",
	Long:  `显示当前保存的配置信息`,
	Run:   runShowConfig,
}

// clearCmd 清空配置命令
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "清空配置",
	Long:  `清空所有保存的配置信息`,
	Run:   runClearConfig,
}

// historyCmd 历史管理命令
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "历史记录管理",
	Long:  `管理对话历史记录，包括保存、加载、查看和删除历史对话。`,
}

// historyListCmd 列出历史记录命令
var historyListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有历史记录",
	Long:  `列出所有保存的对话历史记录`,
	Run:   runHistoryList,
}

// historyShowCmd 显示历史记录命令
var historyShowCmd = &cobra.Command{
	Use:   "show [filename]",
	Short: "显示历史记录内容",
	Long:  `显示指定历史记录的详细内容`,
	Args:  cobra.ExactArgs(1),
	Run:   runHistoryShow,
}

// historySaveCmd 保存历史记录命令
var historySaveCmd = &cobra.Command{
	Use:   "save [title]",
	Short: "保存当前对话",
	Long:  `保存当前对话到历史记录`,
	Args:  cobra.ExactArgs(1),
	Run:   runHistorySave,
}

// historyLoadCmd 加载历史记录命令
var historyLoadCmd = &cobra.Command{
	Use:   "load [filename]",
	Short: "加载历史记录",
	Long:  `加载指定的历史记录并继续对话`,
	Args:  cobra.ExactArgs(1),
	Run:   runHistoryLoad,
}

// historyDeleteCmd 删除历史记录命令
var historyDeleteCmd = &cobra.Command{
	Use:   "delete [filename]",
	Short: "删除历史记录",
	Long:  `删除指定的历史记录`,
	Args:  cobra.ExactArgs(1),
	Run:   runHistoryDelete,
}

// historyClearCmd 清空历史记录命令
var historyClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "清空所有历史记录",
	Long:  `清空所有保存的历史记录`,
	Run:   runHistoryClear,
}

func init() {
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(showCmd)
	configCmd.AddCommand(clearCmd)

	// 历史管理命令
	historyCmd.AddCommand(historyListCmd)
	historyCmd.AddCommand(historyShowCmd)
	historyCmd.AddCommand(historySaveCmd)
	historyCmd.AddCommand(historyLoadCmd)
	historyCmd.AddCommand(historyDeleteCmd)
	historyCmd.AddCommand(historyClearCmd)
}

// runSetConfig 设置配置
func runSetConfig(cmd *cobra.Command, args []string) {
	key := args[0]
	value := args[1]

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("❌ 加载配置失败: %v\n", err)
		os.Exit(1)
	}

	switch key {
	case "api-key":
		cfg.APIKey = value
	case "base-url":
		cfg.BaseURL = value
	case "model":
		cfg.Model = value
	case "system-prompt":
		cfg.SystemPrompt = value
	default:
		fmt.Printf("❌ 不支持的配置项: %s\n", key)
		fmt.Println("支持的配置项: api-key, base-url, model, system-prompt")
		os.Exit(1)
	}

	if err := config.SaveConfig(cfg); err != nil {
		fmt.Printf("❌ 保存配置失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ 配置已更新: %s = %s\n", key, value)
}

// runShowConfig 显示配置
func runShowConfig(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("❌ 加载配置失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("📋 当前配置:")
	fmt.Printf("  API Key: %s\n", maskAPIKey(cfg.APIKey))
	fmt.Printf("  Base URL: %s\n", cfg.BaseURL)
	fmt.Printf("  Model: %s\n", cfg.Model)
	fmt.Printf("  System Prompt: %s\n", cfg.SystemPrompt)
}

// runClearConfig 清空配置
func runClearConfig(cmd *cobra.Command, args []string) {
	cfg := config.DefaultConfig()
	if err := config.SaveConfig(cfg); err != nil {
		fmt.Printf("❌ 清空配置失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ 配置已清空")
}

// maskAPIKey 遮蔽 API Key 显示
func maskAPIKey(apiKey string) string {
	if apiKey == "" {
		return "(未设置)"
	}
	if len(apiKey) <= 8 {
		return "***"
	}
	return apiKey[:4] + "***" + apiKey[len(apiKey)-4:]
}

// runHistoryList 列出历史记录
func runHistoryList(cmd *cobra.Command, args []string) {
	// 创建临时会话来访问历史管理器
	session := chat.NewSession("", "", "", "")

	records, err := session.ListHistory()
	if err != nil {
		fmt.Printf("❌ 获取历史记录失败: %v\n", err)
		os.Exit(1)
	}

	if len(records) == 0 {
		fmt.Println("📝 暂无历史记录")
		return
	}

	fmt.Printf("📚 历史记录 (%d 条):\n\n", len(records))
	for i, record := range records {
		fmt.Printf("%d. %s\n", i+1, record.Title)
		fmt.Printf("   📁 文件: %s\n", record.File)
		fmt.Printf("   🕒 时间: %s\n", record.Timestamp.Format("2006-01-02 15:04:05"))
		fmt.Printf("   🤖 模型: %s\n", record.Model)
		fmt.Printf("   💬 消息数: %d\n\n", len(record.Messages))
	}
}

// runHistoryShow 显示历史记录内容
func runHistoryShow(cmd *cobra.Command, args []string) {
	filename := args[0]

	// 创建临时会话来访问历史管理器
	session := chat.NewSession("", "", "", "")

	record, err := session.HistoryManager.LoadHistory(filename)
	if err != nil {
		fmt.Printf("❌ 加载历史记录失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("📖 历史记录: %s\n", record.Title)
	fmt.Printf("🕒 时间: %s\n", record.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("🤖 模型: %s\n", record.Model)
	fmt.Printf("💬 消息数: %d\n\n", len(record.Messages))

	for _, msg := range record.Messages {
		switch msg.Role {
		case "user":
			fmt.Printf("👤 用户: %s\n\n", msg.Content)
		case "assistant":
			fmt.Printf("🤖 DeepSeek: %s\n\n", msg.Content)
		case "system":
			fmt.Printf("⚙️ 系统: %s\n\n", msg.Content)
		}
	}
}

// runHistorySave 保存历史记录
func runHistorySave(cmd *cobra.Command, args []string) {
	title := args[0]

	// 这里需要从当前会话获取消息，暂时使用空消息
	// 在实际使用中，这应该从当前活跃的会话中获取
	fmt.Printf("⚠️  保存功能需要在对话会话中使用\n")
	fmt.Printf("💡 提示: 在对话中输入 'save %s' 来保存当前对话\n", title)
}

// runHistoryLoad 加载历史记录
func runHistoryLoad(cmd *cobra.Command, args []string) {
	filename := args[0]

	// 创建临时会话来访问历史管理器
	session := chat.NewSession("", "", "", "")

	err := session.LoadFromHistory(filename)
	if err != nil {
		fmt.Printf("❌ 加载历史记录失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ 历史记录已加载: %s\n", filename)
	fmt.Printf("💬 加载了 %d 条消息\n", len(session.Messages))
	fmt.Printf("💡 提示: 使用 './xx' 开始对话以继续此历史记录\n")
}

// runHistoryDelete 删除历史记录
func runHistoryDelete(cmd *cobra.Command, args []string) {
	filename := args[0]

	// 创建临时会话来访问历史管理器
	session := chat.NewSession("", "", "", "")

	err := session.DeleteHistory(filename)
	if err != nil {
		fmt.Printf("❌ 删除历史记录失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ 历史记录已删除: %s\n", filename)
}

// runHistoryClear 清空历史记录
func runHistoryClear(cmd *cobra.Command, args []string) {
	// 创建临时会话来访问历史管理器
	session := chat.NewSession("", "", "", "")

	err := session.ClearHistory()
	if err != nil {
		fmt.Printf("❌ 清空历史记录失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ 所有历史记录已清空")
}

// runChat 运行交互式对话
func runChat(cmd *cobra.Command, args []string) {
	// 获取配置
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("❌ 加载配置失败: %v\n", err)
		os.Exit(1)
	}

	// 获取 API Key (优先级: 命令行参数 > 环境变量 > 配置文件)
	if apiKey == "" {
		apiKey = os.Getenv("DEEPSEEK_API_KEY")
		if apiKey == "" {
			apiKey = cfg.APIKey
		}
	}

	// 如果仍然没有 API Key，提示用户设置
	if apiKey == "" {
		fmt.Println("🔑 首次使用需要设置 API Key")
		fmt.Print("请输入您的 DeepSeek API Key: ")

		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			fmt.Println("❌ 输入失败")
			os.Exit(1)
		}

		apiKey = strings.TrimSpace(scanner.Text())
		if apiKey == "" {
			fmt.Println("❌ API Key 不能为空")
			os.Exit(1)
		}

		// 保存 API Key 到配置文件
		if err := config.SetAPIKey(apiKey); err != nil {
			fmt.Printf("⚠️  保存配置失败: %v\n", err)
		} else {
			fmt.Println("✅ API Key 已保存到配置文件")
		}
	}

	// 使用命令行参数或配置文件中的其他设置
	if baseURL == "" {
		baseURL = cfg.BaseURL
	}
	if model == "" {
		model = cfg.Model
	}
	if systemPrompt == "" {
		systemPrompt = cfg.SystemPrompt
	}

	// 创建会话
	session := chat.NewSession(apiKey, baseURL, model, systemPrompt)

	fmt.Println("🤖 DeepSeek CLI 对话工具")
	fmt.Println("输入 'exit' 或 'quit' 退出，输入 'clear' 清空对话历史")
	fmt.Println("输入 'save 标题' 保存对话，输入 'load 文件名' 加载历史")
	fmt.Println("输入 'list' 查看历史记录")
	fmt.Println("================================================")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("👤 你: ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		// 处理特殊命令
		switch strings.ToLower(input) {
		case "exit", "quit":
			fmt.Println("👋 再见!")
			return
		case "clear":
			session.Clear()
			fmt.Println("🧹 对话历史已清空")
			continue
		}

		// 处理 save 命令
		if strings.HasPrefix(strings.ToLower(input), "save ") {
			title := strings.TrimSpace(strings.TrimPrefix(input, "save "))
			if title == "" {
				title = "未命名对话"
			}

			record, err := session.SaveToHistory(title)
			if err != nil {
				fmt.Printf("❌ 保存失败: %v\n", err)
			} else {
				fmt.Printf("✅ 对话已保存: %s (%s)\n", record.Title, record.File)
			}
			continue
		}

		// 处理 load 命令
		if strings.HasPrefix(strings.ToLower(input), "load ") {
			filename := strings.TrimSpace(strings.TrimPrefix(input, "load "))
			if filename == "" {
				fmt.Println("❌ 请指定要加载的文件名")
				continue
			}

			err := session.LoadFromHistory(filename)
			if err != nil {
				fmt.Printf("❌ 加载失败: %v\n", err)
			} else {
				fmt.Printf("✅ 历史记录已加载: %s\n", filename)
				fmt.Printf("💬 加载了 %d 条消息\n", len(session.Messages))
			}
			continue
		}

		// 处理 list 命令
		if strings.ToLower(input) == "list" {
			records, err := session.ListHistory()
			if err != nil {
				fmt.Printf("❌ 获取历史记录失败: %v\n", err)
			} else if len(records) == 0 {
				fmt.Println("📝 暂无历史记录")
			} else {
				fmt.Printf("📚 历史记录 (%d 条):\n", len(records))
				for i, record := range records {
					fmt.Printf("  %d. %s (%s)\n", i+1, record.Title, record.File)
				}
			}
			continue
		}

		// 添加用户消息到会话
		session.AddUserMessage(input)

		// 显示 AI 回复
		fmt.Print("🤖 DeepSeek: ")

		// 使用流式响应
		var fullResponse strings.Builder
		err := session.Client.ChatStream(session.GetMessagesWithSystem(), func(chunk string) error {
			fmt.Print(chunk)
			fullResponse.WriteString(chunk)
			return nil
		})

		if err != nil {
			fmt.Printf("\n❌ 错误: %v\n", err)
			continue
		}

		// 添加助手回复到会话
		session.AddAssistantMessage(fullResponse.String())
		fmt.Println() // 换行
		fmt.Println()
	}
}

// Execute 执行根命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

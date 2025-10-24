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
	apiKey  string
	baseURL string
	model   string
)

// rootCmd 表示基础命令
var rootCmd = &cobra.Command{
	Use:   "deepseek-cli",
	Short: "DeepSeek CLI 对话工具",
	Long:  `一个基于 Go 的命令行工具，用于与 DeepSeek API 进行交互式对话。`,
	Run:   runChat,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "DeepSeek API Key (也可以通过 DEEPSEEK_API_KEY 环境变量设置)")
	rootCmd.PersistentFlags().StringVar(&baseURL, "base-url", "https://api.deepseek.com", "DeepSeek API 基础 URL")
	rootCmd.PersistentFlags().StringVar(&model, "model", "deepseek-chat", "使用的模型名称")

	// 添加配置管理子命令
	rootCmd.AddCommand(configCmd)
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

func init() {
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(showCmd)
	configCmd.AddCommand(clearCmd)
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
	default:
		fmt.Printf("❌ 不支持的配置项: %s\n", key)
		fmt.Println("支持的配置项: api-key, base-url, model")
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

	// 创建会话
	session := chat.NewSession(apiKey, baseURL, model)

	fmt.Println("🤖 DeepSeek CLI 对话工具")
	fmt.Println("输入 'exit' 或 'quit' 退出，输入 'clear' 清空对话历史")
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

		// 添加用户消息到会话
		session.AddUserMessage(input)

		// 显示 AI 回复
		fmt.Print("🤖 DeepSeek: ")

		// 使用流式响应
		var fullResponse strings.Builder
		err := session.Client.ChatStream(session.GetMessages(), func(chunk string) error {
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

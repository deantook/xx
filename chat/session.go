package chat

import (
	"fmt"

	"github.com/deantook/xx/client"
	"github.com/deantook/xx/history"
)

// Session 表示一个对话会话
type Session struct {
	Messages       []client.Message
	Client         *client.DeepSeekClient
	HistoryManager *history.HistoryManager
	Model          string
	SystemPrompt   string
}

// NewSession 创建新的对话会话
func NewSession(apiKey, baseURL, model, systemPrompt string) *Session {
	historyManager, err := history.NewHistoryManager()
	if err != nil {
		fmt.Printf("⚠️  初始化历史管理器失败: %v\n", err)
	}

	return &Session{
		Messages:       make([]client.Message, 0),
		Client:         client.NewDeepSeekClient(apiKey, baseURL, model),
		HistoryManager: historyManager,
		Model:          model,
		SystemPrompt:   systemPrompt,
	}
}

// AddMessage 添加消息到会话中
func (s *Session) AddMessage(role, content string) {
	s.Messages = append(s.Messages, client.Message{
		Role:    role,
		Content: content,
	})
}

// AddUserMessage 添加用户消息
func (s *Session) AddUserMessage(content string) {
	s.AddMessage("user", content)
}

// AddAssistantMessage 添加助手消息
func (s *Session) AddAssistantMessage(content string) {
	s.AddMessage("assistant", content)
}

// AddSystemMessage 添加系统消息
func (s *Session) AddSystemMessage(content string) {
	s.AddMessage("system", content)
}

// Clear 清空对话历史
func (s *Session) Clear() {
	s.Messages = make([]client.Message, 0)
}

// GetMessages 获取所有消息
func (s *Session) GetMessages() []client.Message {
	return s.Messages
}

// GetMessagesWithSystem 获取包含系统提示词的消息列表
func (s *Session) GetMessagesWithSystem() []client.Message {
	messages := make([]client.Message, 0)

	// 添加系统提示词（如果存在且消息列表为空或第一条不是系统消息）
	if s.SystemPrompt != "" {
		hasSystemMessage := false
		if len(s.Messages) > 0 && s.Messages[0].Role == "system" {
			hasSystemMessage = true
		}

		if !hasSystemMessage {
			messages = append(messages, client.Message{
				Role:    "system",
				Content: s.SystemPrompt,
			})
		}
	}

	// 添加其他消息
	messages = append(messages, s.Messages...)

	return messages
}

// GetLastUserMessage 获取最后一条用户消息
func (s *Session) GetLastUserMessage() string {
	for i := len(s.Messages) - 1; i >= 0; i-- {
		if s.Messages[i].Role == "user" {
			return s.Messages[i].Content
		}
	}
	return ""
}

// GetLastAssistantMessage 获取最后一条助手消息
func (s *Session) GetLastAssistantMessage() string {
	for i := len(s.Messages) - 1; i >= 0; i-- {
		if s.Messages[i].Role == "assistant" {
			return s.Messages[i].Content
		}
	}
	return ""
}

// GetMessageCount 获取消息数量
func (s *Session) GetMessageCount() int {
	return len(s.Messages)
}

// HasMessages 检查是否有消息
func (s *Session) HasMessages() bool {
	return len(s.Messages) > 0
}

// SaveToHistory 保存当前对话到历史记录
func (s *Session) SaveToHistory(title string) (*history.HistoryRecord, error) {
	if s.HistoryManager == nil {
		return nil, fmt.Errorf("历史管理器未初始化")
	}

	if len(s.Messages) == 0 {
		return nil, fmt.Errorf("没有对话内容可保存")
	}

	return s.HistoryManager.SaveHistory(title, s.Model, s.Messages)
}

// LoadFromHistory 从历史记录加载对话
func (s *Session) LoadFromHistory(filename string) error {
	if s.HistoryManager == nil {
		return fmt.Errorf("历史管理器未初始化")
	}

	record, err := s.HistoryManager.LoadHistory(filename)
	if err != nil {
		return err
	}

	// 清空当前消息并加载历史消息
	s.Messages = record.Messages
	return nil
}

// ListHistory 列出所有历史记录
func (s *Session) ListHistory() ([]*history.HistoryRecord, error) {
	if s.HistoryManager == nil {
		return nil, fmt.Errorf("历史管理器未初始化")
	}

	return s.HistoryManager.ListHistory()
}

// DeleteHistory 删除历史记录
func (s *Session) DeleteHistory(filename string) error {
	if s.HistoryManager == nil {
		return fmt.Errorf("历史管理器未初始化")
	}

	return s.HistoryManager.DeleteHistory(filename)
}

// ClearHistory 清空所有历史记录
func (s *Session) ClearHistory() error {
	if s.HistoryManager == nil {
		return fmt.Errorf("历史管理器未初始化")
	}

	return s.HistoryManager.ClearHistory()
}

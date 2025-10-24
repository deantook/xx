package chat

import (
	"xx/client"
)

// Session 表示一个对话会话
type Session struct {
	Messages []client.Message
	Client   *client.DeepSeekClient
}

// NewSession 创建新的对话会话
func NewSession(apiKey, baseURL, model string) *Session {
	return &Session{
		Messages: make([]client.Message, 0),
		Client:   client.NewDeepSeekClient(apiKey, baseURL, model),
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

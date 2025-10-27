package history

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/deantook/xx/client"
)

// HistoryRecord 表示一条历史记录
type HistoryRecord struct {
	Title     string           `json:"title"`
	Timestamp time.Time        `json:"timestamp"`
	Model     string           `json:"model"`
	Messages  []client.Message `json:"messages"`
	File      string           `json:"file"`
}

// HistoryManager 历史记录管理器
type HistoryManager struct {
	HistoryDir string
}

// NewHistoryManager 创建新的历史记录管理器
func NewHistoryManager() (*HistoryManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("获取用户主目录失败: %v", err)
	}

	historyDir := filepath.Join(homeDir, ".xx", "history")
	if err := os.MkdirAll(historyDir, 0755); err != nil {
		return nil, fmt.Errorf("创建历史目录失败: %v", err)
	}

	return &HistoryManager{
		HistoryDir: historyDir,
	}, nil
}

// SaveHistory 保存对话历史
func (hm *HistoryManager) SaveHistory(title string, model string, messages []client.Message) (*HistoryRecord, error) {
	if title == "" {
		title = "未命名对话"
	}

	// 生成文件名
	timestamp := time.Now()
	filename := fmt.Sprintf("%s_%s.md",
		timestamp.Format("2006-01-02_15-04-05"),
		hm.sanitizeFilename(title))

	filepath := filepath.Join(hm.HistoryDir, filename)

	// 创建历史记录
	record := &HistoryRecord{
		Title:     title,
		Timestamp: timestamp,
		Model:     model,
		Messages:  messages,
		File:      filename,
	}

	// 保存为 Markdown 格式
	if err := hm.saveAsMarkdown(record, filepath); err != nil {
		return nil, fmt.Errorf("保存 Markdown 文件失败: %v", err)
	}

	// 保存元数据
	if err := hm.saveMetadata(record); err != nil {
		return nil, fmt.Errorf("保存元数据失败: %v", err)
	}

	return record, nil
}

// LoadHistory 加载对话历史
func (hm *HistoryManager) LoadHistory(filename string) (*HistoryRecord, error) {
	// 尝试加载元数据
	metadataFile := filepath.Join(hm.HistoryDir, strings.TrimSuffix(filename, ".md")+".json")
	if _, err := os.Stat(metadataFile); err == nil {
		return hm.loadFromMetadata(metadataFile)
	}

	// 如果没有元数据，从 Markdown 文件解析
	filepath := filepath.Join(hm.HistoryDir, filename)
	return hm.loadFromMarkdown(filepath)
}

// ListHistory 列出所有历史记录
func (hm *HistoryManager) ListHistory() ([]*HistoryRecord, error) {
	files, err := os.ReadDir(hm.HistoryDir)
	if err != nil {
		return nil, fmt.Errorf("读取历史目录失败: %v", err)
	}

	var records []*HistoryRecord
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			record, err := hm.loadFromMetadata(filepath.Join(hm.HistoryDir, file.Name()))
			if err != nil {
				continue // 跳过损坏的文件
			}
			records = append(records, record)
		}
	}

	// 按时间排序（最新的在前）
	sort.Slice(records, func(i, j int) bool {
		return records[i].Timestamp.After(records[j].Timestamp)
	})

	return records, nil
}

// DeleteHistory 删除历史记录
func (hm *HistoryManager) DeleteHistory(filename string) error {
	// 删除 Markdown 文件
	mdFile := filepath.Join(hm.HistoryDir, filename)
	if err := os.Remove(mdFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("删除 Markdown 文件失败: %v", err)
	}

	// 删除元数据文件
	jsonFile := filepath.Join(hm.HistoryDir, strings.TrimSuffix(filename, ".md")+".json")
	if err := os.Remove(jsonFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("删除元数据文件失败: %v", err)
	}

	return nil
}

// ClearHistory 清空所有历史记录
func (hm *HistoryManager) ClearHistory() error {
	files, err := os.ReadDir(hm.HistoryDir)
	if err != nil {
		return fmt.Errorf("读取历史目录失败: %v", err)
	}

	for _, file := range files {
		if err := os.Remove(filepath.Join(hm.HistoryDir, file.Name())); err != nil {
			return fmt.Errorf("删除文件 %s 失败: %v", file.Name(), err)
		}
	}

	return nil
}

// saveAsMarkdown 保存为 Markdown 格式
func (hm *HistoryManager) saveAsMarkdown(record *HistoryRecord, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// 写入标题和时间信息
	fmt.Fprintf(writer, "# %s\n\n", record.Title)
	fmt.Fprintf(writer, "**时间**: %s\n", record.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(writer, "**模型**: %s\n\n", record.Model)

	// 写入对话内容
	for _, msg := range record.Messages {
		switch msg.Role {
		case "user":
			fmt.Fprintf(writer, "## 👤 用户\n\n%s\n\n", msg.Content)
		case "assistant":
			fmt.Fprintf(writer, "## 🤖 DeepSeek\n\n%s\n\n", msg.Content)
		case "system":
			fmt.Fprintf(writer, "## ⚙️ 系统\n\n%s\n\n", msg.Content)
		}
	}

	return nil
}

// saveMetadata 保存元数据
func (hm *HistoryManager) saveMetadata(record *HistoryRecord) error {
	metadataFile := filepath.Join(hm.HistoryDir, strings.TrimSuffix(record.File, ".md")+".json")

	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(metadataFile, data, 0644)
}

// loadFromMetadata 从元数据加载
func (hm *HistoryManager) loadFromMetadata(metadataFile string) (*HistoryRecord, error) {
	data, err := os.ReadFile(metadataFile)
	if err != nil {
		return nil, err
	}

	var record HistoryRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return nil, err
	}

	return &record, nil
}

// loadFromMarkdown 从 Markdown 文件加载
func (hm *HistoryManager) loadFromMarkdown(filePath string) (*HistoryRecord, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var messages []client.Message
	var currentRole string
	var currentContent strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "## 👤 用户") {
			if currentRole != "" && currentContent.Len() > 0 {
				messages = append(messages, client.Message{
					Role:    currentRole,
					Content: strings.TrimSpace(currentContent.String()),
				})
			}
			currentRole = "user"
			currentContent.Reset()
		} else if strings.HasPrefix(line, "## 🤖 DeepSeek") {
			if currentRole != "" && currentContent.Len() > 0 {
				messages = append(messages, client.Message{
					Role:    currentRole,
					Content: strings.TrimSpace(currentContent.String()),
				})
			}
			currentRole = "assistant"
			currentContent.Reset()
		} else if strings.HasPrefix(line, "## ⚙️ 系统") {
			if currentRole != "" && currentContent.Len() > 0 {
				messages = append(messages, client.Message{
					Role:    currentRole,
					Content: strings.TrimSpace(currentContent.String()),
				})
			}
			currentRole = "system"
			currentContent.Reset()
		} else if currentRole != "" && !strings.HasPrefix(line, "#") {
			if currentContent.Len() > 0 {
				currentContent.WriteString("\n")
			}
			currentContent.WriteString(line)
		}
	}

	// 添加最后一条消息
	if currentRole != "" && currentContent.Len() > 0 {
		messages = append(messages, client.Message{
			Role:    currentRole,
			Content: strings.TrimSpace(currentContent.String()),
		})
	}

	// 从文件名提取信息
	baseFilename := filepath.Base(filePath)
	parts := strings.Split(baseFilename, "_")
	var title string
	if len(parts) >= 3 {
		title = strings.TrimSuffix(strings.Join(parts[2:], "_"), ".md")
	} else {
		title = "未命名对话"
	}

	return &HistoryRecord{
		Title:     title,
		Timestamp: time.Now(), // 无法从文件名精确解析时间
		Model:     "deepseek-chat",
		Messages:  messages,
		File:      baseFilename,
	}, nil
}

// sanitizeFilename 清理文件名
func (hm *HistoryManager) sanitizeFilename(filename string) string {
	// 替换不安全的字符
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
		" ", "_",
	)

	result := replacer.Replace(filename)
	if len(result) > 50 {
		result = result[:50]
	}

	return result
}

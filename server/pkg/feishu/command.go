package feishu

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"flec_blog/pkg/logger"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// CommandHandler 命令处理器函数类型
type CommandHandler func(ctx context.Context, args []string, openID string) error

var (
	commandHandlers = make(map[string]CommandHandler)
	commandsMu      sync.RWMutex
)

// RegisterCommand 注册命令处理器
func RegisterCommand(command string, handler CommandHandler) {
	commandsMu.Lock()
	defer commandsMu.Unlock()
	commandHandlers[command] = handler
}

// handleCommand 处理命令消息
func handleCommand(_ context.Context, event *larkim.P2MessageReceiveV1) error {
	if event.Event == nil || event.Event.Message == nil || event.Event.Message.Content == nil {
		return nil
	}

	// 验证消息来自管理员群
	globalMu.RLock()
	chatID := globalClient.chatID
	globalMu.RUnlock()

	if event.Event.Message.ChatId == nil || *event.Event.Message.ChatId != chatID {
		return nil
	}

	// 解析消息内容
	var content struct {
		Text string `json:"text"`
	}
	if err := json.Unmarshal([]byte(*event.Event.Message.Content), &content); err != nil {
		return nil
	}

	text := strings.TrimSpace(content.Text)

	// 查找命令（格式：/command）
	cmdIndex := strings.Index(text, "/")
	if cmdIndex == -1 {
		return nil
	}

	// 提取命令和参数
	cmdPart := strings.TrimSpace(text[cmdIndex:])
	parts := strings.Fields(cmdPart)
	if len(parts) == 0 {
		return nil
	}

	command := parts[0]
	args := parts[1:]

	// 获取发送者 OpenID
	var openID string
	if event.Event.Sender != nil && event.Event.Sender.SenderId != nil && event.Event.Sender.SenderId.OpenId != nil {
		openID = *event.Event.Sender.SenderId.OpenId
	}
	if openID == "" {
		return nil
	}

	// 查找并执行命令处理器
	commandsMu.RLock()
	handler, exists := commandHandlers[command]
	commandsMu.RUnlock()

	if !exists {
		return nil
	}

	return handler(context.Background(), args, openID)
}

// InitCommands 初始化命令处理器
func InitCommands(userService UserBinder) {
	userBinderInstance = userService

	// 注册 /bind 命令
	RegisterCommand("/bind", func(ctx context.Context, args []string, openID string) error {
		if len(args) == 0 {
			return sendBindErrorCard(ctx, "请提供邮箱地址\n\n用法：`/bind 邮箱`")
		}

		email := args[0]

		// 处理飞书Markdown链接格式
		if strings.HasPrefix(email, "[") && strings.Contains(email, "](mailto:") {
			if endBracket := strings.Index(email, "]"); endBracket > 1 {
				email = email[1:endBracket]
			}
		}

		if userBinderInstance != nil {
			if err := userBinderInstance.BindFeishuByEmail(email, openID); err != nil {
				logger.Error("[Feishu] 绑定失败: %v", err)
				return sendBindErrorCard(ctx, err.Error())
			}
			logger.Info("[Feishu] 绑定成功: %s", email)
			return sendBindSuccessCard(ctx, email)
		}
		return nil
	})

	// 注册 /help 命令
	RegisterCommand("/help", func(ctx context.Context, args []string, openID string) error {
		helpText := "**可用命令：**\n\n" +
			"• `/bind 邮箱` - 绑定飞书账号\n" +
			"• `/help` - 显示此帮助信息"

		globalMu.RLock()
		client := globalClient
		globalMu.RUnlock()

		if client != nil && client.IsEnabled() {
			card := buildHelpCard(helpText)
			return client.SendMessage(ctx, card)
		}
		return nil
	})
}

// buildHelpCard 构建帮助卡片
func buildHelpCard(helpText string) string {
	elements := []interface{}{
		&MarkdownElement{Tag: "markdown", Content: helpText},
	}
	card, _ := buildCard("📖 命令帮助", "blue", elements)
	return card
}

// sendBindSuccessCard 发送绑定成功卡片
func sendBindSuccessCard(ctx context.Context, email string) error {
	globalMu.RLock()
	client := globalClient
	globalMu.RUnlock()

	if client != nil && client.IsEnabled() {
		content := fmt.Sprintf("✅ 绑定成功\n\n邮箱：`%s`", email)
		elements := []interface{}{
			&MarkdownElement{Tag: "markdown", Content: content},
		}
		card, _ := buildCard("绑定成功", "green", elements)
		return client.SendMessage(ctx, card)
	}
	return nil
}

// sendBindErrorCard 发送绑定失败卡片
func sendBindErrorCard(ctx context.Context, errMsg string) error {
	globalMu.RLock()
	client := globalClient
	globalMu.RUnlock()

	if client != nil && client.IsEnabled() {
		content := fmt.Sprintf("❌ 绑定失败\n\n%s", errMsg)
		elements := []interface{}{
			&MarkdownElement{Tag: "markdown", Content: content},
		}
		card, _ := buildCard("绑定失败", "red", elements)
		return client.SendMessage(ctx, card)
	}
	return nil
}

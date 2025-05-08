//实现回调：当群内@机器人时，机器人回复对应链接

// package main

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	lark "github.com/larksuite/oapi-sdk-go/v3"
// 	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
// 	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
// 	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
// 	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
// )

// func main() {
// 	app_id := os.Getenv("APP_ID")
// 	app_secret := os.Getenv("APP_SECRET")

// 	/**
// 	 * 创建 LarkClient 对象，用于请求OpenAPI。
// 	 * Create LarkClient object for requesting OpenAPI
// 	 */
// 	client := lark.NewClient(app_id, app_secret)

// 	/**
// 	 * 注册事件处理器。
// 	 * Register event handler.
// 	 */
// 	eventHandler := dispatcher.NewEventDispatcher("", "").
// 		/**
// 		 * 注册接收消息事件，处理接收到的消息。
// 		 * Register event handler to handle received messages.
// 		 * https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive
// 		 */
// 		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
// 			fmt.Printf("[OnP2MessageReceiveV1 access], data: %s\n", larkcore.Prettify(event))
// 			/**
// 			 * 解析用户发送的消息。
// 			var respContent map[string]string
// 			err := json.Unmarshal([]byte(*event.Event.Message.Content), &respContent)
// 			/**
// 			 * 检查消息类型是否为文本
// 			 * Check if the message type is text
// 			*/
// 			if err != nil || *event.Event.Message.MessageType != "text" {
// 				respContent = map[string]string{
// 					"text": "解析消息失败，请发送文本消息\nparse message failed, please send text message",
// 				}
// 			}

// 			/**
// 			 * 构建回复消息
// 			 * Build reply message
// 			 */
// 			content := larkim.NewTextMsgBuilder().
// 				TextLine("https://li.feishu.cn/share/base/form/shrcn9HifpjawRtdXAsUPlVH4qb: ").
// 				//TextLine("收到你发送的消息: " + respContent["text"]).
// 				//TextLine("Received message: " + respContent["text"]).
// 				Build()

// 			if *event.Event.Message.ChatType == "p2p" {
// 				/**
// 				 * 使用SDK调用发送消息接口。 Use SDK to call send message interface.
// 				 * https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create
// 				 */
// 				resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
// 					ReceiveIdType(larkim.ReceiveIdTypeChatId). // 消息接收者的 ID 类型，设置为会话ID。 ID type of the message receiver, set to chat ID.
// 					Body(larkim.NewCreateMessageReqBodyBuilder().
// 						MsgType(larkim.MsgTypeText).            // 设置消息类型为文本消息。 Set message type to text message.
// 						ReceiveId(*event.Event.Message.ChatId). // 消息接收者的 ID 为消息发送的会话ID。 ID of the message receiver is the chat ID of the message sending.
// 						Content(content).
// 						Build()).
// 					Build())

// 				if err != nil || !resp.Success() {
// 					fmt.Println(err)
// 					fmt.Println(resp.Code, resp.Msg, resp.RequestId())
// 					return nil
// 				}

// 			} else {
// 				/**
// 				 * 使用SDK调用回复消息接口。 Use SDK to call send message interface.
// 				 * https://open.feishu.cn/document/server-docs/im-v1/message/reply
// 				 */
// 				resp, err := client.Im.Message.Reply(context.Background(), larkim.NewReplyMessageReqBuilder().
// 					MessageId(*event.Event.Message.MessageId).
// 					Body(larkim.NewReplyMessageReqBodyBuilder().
// 						MsgType(larkim.MsgTypeText). // 设置消息类型为文本消息。 Set message type to text message.
// 						Content(content).
// 						Build()).
// 					Build())
// 				if err != nil || !resp.Success() {
// 					fmt.Printf("logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
// 					return nil
// 				}
// 			}

// 			return nil
// 		})

// 	/**
// 	 * 启动长连接，并注册事件处理器。
// 	 * Start long connection and register event handler.
// 	 */
// 	cli := larkws.NewClient(app_id, app_secret,
// 		larkws.WithEventHandler(eventHandler),
// 		larkws.WithLogLevel(larkcore.LogLevelDebug),
// 	)
// 	err := cli.Start(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // ---------------------------------------------------Second Version ➡️缺少im：massage：receive_id--------------------------------------------------
// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strings"

// 	lark "github.com/larksuite/oapi-sdk-go/v3"
// 	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
// 	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
// 	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
// 	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
// )

// func main() {
// 	app_id := os.Getenv("APP_ID")
// 	app_secret := os.Getenv("APP_SECRET")

// 	// 创建 LarkClient 对象，用于请求OpenAPI
// 	client := lark.NewClient(app_id, app_secret)

// 	// 注册事件处理器

// 	eventHandler := dispatcher.NewEventDispatcher("", "").
// 		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
// 			fmt.Printf("[OnP2MessageReceiveV1 access], data: %s\n", larkcore.Prettify(event))

// 			// 检查消息类型
// 			if *event.Event.Message.ChatType != "group" && *event.Event.Message.ChatType != "p2p" {
// 				return nil
// 			}

// 			// 解析消息内容
// 			var respContent map[string]interface{}
// 			err := json.Unmarshal([]byte(*event.Event.Message.Content), &respContent)
// 			if err != nil {
// 				fmt.Println("Failed to parse message content:", err)
// 				return nil
// 			}
// 			fmt.Printf("Parsed content: %v\n", respContent)

// 			// 检查是否包含 @ 机器人
// 			mentions, ok := respContent["mentions"].([]interface{})
// 			if !ok || len(mentions) == 0 {
// 				return nil
// 			}

// 			// 遍历 mentions，检查是否 @ 了机器人
// 			isMentioned := false
// 			for _, mention := range mentions {
// 				mentionMap := mention.(map[string]interface{})
// 				if mentionMap["name"] == "自动回复机器人" { // 替换为你的机器人名称
// 					isMentioned = true
// 					break
// 				}
// 			}

// 			if !isMentioned {
// 				return nil
// 			}

// 			// 检查消息内容是否包含关键词
// 			text, ok := respContent["text"].(string)
// 			if !ok {
// 				text = ""
// 			}

// 			// 去掉 @ 的部分
// 			text = strings.TrimSpace(strings.ReplaceAll(text, "@_user_1", ""))

// 			var replyLink string
// 			if strings.Contains(text, "换车") || strings.Contains(text, "置换") {
// 				replyLink = "https://li.feishu.cn/share/base/form/shrcnhPseIh9AHP5fqUURPTGWbh"
// 			} else if strings.Contains(text, "修车") || strings.Contains(text, "出险") {
// 				replyLink = "https://li.feishu.cn/share/base/form/shrcn9HifpjawRtdXAsUPlVH4qb"
// 			} else {
// 				return nil
// 			}

// 			// 构建回复消息
// 			content := larkim.NewTextMsgBuilder().
// 				TextLine(fmt.Sprintf("这是你要的链接: %s", replyLink)).
// 				Build()

// 			// 发送消息
// 			resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
// 				ReceiveIdType(larkim.ReceiveIdTypeChatId).
// 				Body(larkim.NewCreateMessageReqBodyBuilder().
// 					MsgType(larkim.MsgTypeText).
// 					ReceiveId(*event.Event.Message.ChatId).
// 					Content(content).
// 					Build()).
// 				Build())
// 			if err != nil || !resp.Success() {
// 				fmt.Printf("logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
// 				return nil
// 			}

// 			return nil
// 		})

// 	// 启动长连接，并注册事件处理器
// 	cli := larkws.NewClient(app_id, app_secret,
// 		larkws.WithEventHandler(eventHandler),
// 		larkws.WithLogLevel(larkcore.LogLevelDebug),
// 	)
// 	err := cli.Start(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // ------------------------------------------------------------Third Version 显示权限不足--------------------------------------------------------
// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strings"

// 	lark "github.com/larksuite/oapi-sdk-go/v3"
// 	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
// )

// func main() {
// 	app_id := os.Getenv("APP_ID")
// 	app_secret := os.Getenv("APP_SECRET")

// 	// 创建 LarkClient 对象，用于请求OpenAPI
// 	client := lark.NewClient(app_id, app_secret)

// 	// 模拟轮询消息（替代事件订阅）
// 	for {
// 		// 获取最近的消息
// 		resp, err := client.Im.Message.List(context.Background(), larkim.NewListMessageReqBuilder().
// 			PageSize(10). // 设置分页大小
// 			Build())
// 		if err != nil || !resp.Success() {
// 			fmt.Printf("Failed to fetch messages: %v, response: %v\n", err, resp)
// 			continue
// 		}

// 		// 遍历消息
// 		for _, message := range resp.Data.Items {
// 			fmt.Printf("Received message: %v\n", message)

// 			// 检查是否包含 @ 机器人
// 			var content map[string]interface{}
// 			err := json.Unmarshal([]byte(*message.Body.Content), &content) // 使用 Body.Content 获取消息内容
// 			if err != nil {
// 				fmt.Printf("Failed to parse message content: %v\n", err)
// 				continue
// 			}

// 			mentions, ok := content["mentions"].([]interface{})
// 			if !ok || len(mentions) == 0 {
// 				continue
// 			}

// 			isMentioned := false
// 			for _, mention := range mentions {
// 				mentionMap := mention.(map[string]interface{})
// 				if mentionMap["name"] == "自动回复机器人" { // 替换为你的机器人名称
// 					isMentioned = true
// 					break
// 				}
// 			}

// 			if !isMentioned {
// 				continue
// 			}

// 			// 检查消息内容是否包含关键词
// 			text, ok := content["text"].(string)
// 			if !ok {
// 				text = ""
// 			}

// 			// 去掉 @ 的部分
// 			text = strings.TrimSpace(strings.ReplaceAll(text, "@_user_1", ""))

// 			var replyLink string
// 			if strings.Contains(text, "换车") || strings.Contains(text, "置换") {
// 				replyLink = "https://li.feishu.cn/share/base/form/shrcnhPseIh9AHP5fqUURPTGWbh"
// 			} else if strings.Contains(text, "修车") || strings.Contains(text, "出险") {
// 				replyLink = "https://li.feishu.cn/share/base/form/shrcn9HifpjawRtdXAsUPlVH4qb"
// 			} else {
// 				continue
// 			}

// 			// 构建回复消息内容
// 			replyContent := map[string]interface{}{
// 				"text": fmt.Sprintf("这是你要的链接: %s", replyLink),
// 			}

// 			// 将内容转换为 JSON 字符串
// 			contentJSON, err := json.Marshal(replyContent)
// 			if err != nil {
// 				fmt.Printf("Failed to marshal content: %v\n", err)
// 				continue
// 			}

// 			// 发送消息
// 			_, err = client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
// 				ReceiveIdType(larkim.ReceiveIdTypeChatId).
// 				Body(larkim.NewCreateMessageReqBodyBuilder().
// 					MsgType(larkim.MsgTypeText).
// 					ReceiveId(*message.ChatId).   // 解引用指针
// 					Content(string(contentJSON)). // 使用 JSON 字符串作为内容
// 					Build()).
// 				Build())
// 			if err != nil {
// 				fmt.Printf("Failed to send message: %v\n", err)
// 			}
// 		}
// 	}
// }
//----------------------------------✅ Forth Version------------------------------------

package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

func main() {
	app_id := os.Getenv("APP_ID")
	app_secret := os.Getenv("APP_SECRET")

	/**
	 * 创建 LarkClient 对象，用于请求OpenAPI。
	 * Create LarkClient object for requesting OpenAPI
	 */
	client := lark.NewClient(app_id, app_secret)

	/**
	 * 注册事件处理器。
	 * Register event handler.
	 */

	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			fmt.Printf("[OnP2MessageReceiveV1 access], data: %s\n", larkcore.Prettify(event))

			// 检查消息类型是否为文本
			if *event.Event.Message.MessageType != "text" {
				fmt.Println("消息类型不是文本，忽略处理")
				return nil
			}

			// 解析消息内容
			var respContent map[string]string
			err := json.Unmarshal([]byte(*event.Event.Message.Content), &respContent)
			if err != nil {
				fmt.Println("解析消息内容失败:", err)
				return nil
			}

			// 获取消息文本
			text := respContent["text"]

			// ...existing code...

			// ...existing code...

			// 检查是否包含关键词
			var replyLink string
			if strings.Contains(text, "买车") || strings.Contains(text, "置换") {
				replyLink = "https://li.feishu.cn/share/base/form/shrcn9HifpjawRtdXAsUPlVH4qb"
			} else if strings.Contains(text, "事故") || strings.Contains(text, "出险") {
				replyLink = "https://li.feishu.cn/share/base/form/shrcnhPseIh9AHP5fqUURPTGWbh"
			} else if strings.Contains(text, "图片") {
				// 读取本地图片文件
				imagePath := "/Users/raocong/Desktop/code/lark-samples-main/echo_bot/go/assets/123.jpeg"
				// 替换为你的图片路径
				imageFile, err := os.Open(imagePath)
				if err != nil {
					fmt.Printf("打开图片失败: %v\n", err)
					return nil
				}
				defer imageFile.Close()

				// 将图片文件读取为二进制数据
				imageData, err := io.ReadAll(imageFile)
				if err != nil {
					fmt.Printf("读取图片数据失败: %v\n", err)
					return nil
				}

				// 构建图片消息内容
				content := fmt.Sprintf(`{"file_name":"123.jpeg","file_data":"%s"}`, base64.StdEncoding.EncodeToString(imageData))

				// 发送图片消息
				resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
					ReceiveIdType(larkim.ReceiveIdTypeChatId).
					Body(larkim.NewCreateMessageReqBodyBuilder().
						MsgType(larkim.MsgTypeFile). // 使用文件类型发送图片
						ReceiveId(*event.Event.Message.ChatId).
						Content(content).
						Build()).
					Build())
				if err != nil || !resp.Success() {
					fmt.Printf("发送图片消息失败: logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
					return nil
				}

				fmt.Println("图片消息发送成功")
				return nil
			} else {
				// 如果不包含关键词，则不回复
				fmt.Println("消息中不包含关键词，忽略处理")
				return nil
			}

			// ...existing code...

			// ...existing code...
			// 构建回复消息
			content := larkim.NewTextMsgBuilder().
				TextLine(fmt.Sprintf("%s", replyLink)).
				Build()

			// 发送消息
			resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
				ReceiveIdType(larkim.ReceiveIdTypeChatId).
				Body(larkim.NewCreateMessageReqBodyBuilder().
					MsgType(larkim.MsgTypeText).
					ReceiveId(*event.Event.Message.ChatId).
					Content(content).
					Build()).
				Build())
			if err != nil || !resp.Success() {
				fmt.Printf("发送消息失败: logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
				return nil
			}

			return nil
		})

	/**
	 * 启动长连接，并注册事件处理器。
	 * Start long connection and register event handler.
	 */
	cli := larkws.NewClient(app_id, app_secret,
		larkws.WithEventHandler(eventHandler),
		larkws.WithLogLevel(larkcore.LogLevelDebug),
	)
	err := cli.Start(context.Background())
	if err != nil {
		panic(err)
	}
}

// ----------------------------------❌Fifth Version （testing）------------------------------------
// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strings"

// 	lark "github.com/larksuite/oapi-sdk-go/v3"
// 	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
// 	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
// 	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
// 	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
// )

// func main() {
// 	app_id := os.Getenv("APP_ID")
// 	app_secret := os.Getenv("APP_SECRET")

// 	/**
// 	 * 创建 LarkClient 对象，用于请求OpenAPI。
// 	 * Create LarkClient object for requesting OpenAPI
// 	 */
// 	client := lark.NewClient(app_id, app_secret)

// 	/**
// 	 * 注册事件处理器。
// 	 * Register event handler.
// 	 */
// 	eventHandler := dispatcher.NewEventDispatcher("", "").
// 		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
// 			fmt.Printf("[OnP2MessageReceiveV1 access], data: %s\n", larkcore.Prettify(event))

// 			// 检查消息类型是否为文本
// 			if *event.Event.Message.MessageType != "text" {
// 				fmt.Println("消息类型不是文本，忽略处理")
// 				return nil
// 			}

// 			// 解析消息内容
// 			var respContent map[string]string
// 			err := json.Unmarshal([]byte(*event.Event.Message.Content), &respContent)
// 			if err != nil {
// 				fmt.Println("解析消息内容失败:", err)
// 				return nil
// 			}

// 			// 获取消息文本
// 			text := respContent["text"]

// 			// 检查是否包含关键词
// 			var replyLink string
// 			if strings.Contains(text, "买车") || strings.Contains(text, "置换") {
// 				replyLink = "https://li.feishu.cn/share/base/form/shrcn9HifpjawRtdXAsUPlVH4qb"
// 			} else if strings.Contains(text, "事故") || strings.Contains(text, "出险") {
// 				replyLink = "https://li.feishu.cn/share/base/form/shrcnhPseIh9AHP5fqUURPTGWbh"
// 			} else {
// 				// 如果不包含关键词，则不回复
// 				fmt.Println("消息中不包含关键词，忽略处理")
// 				return nil
// 			}

// 			// 上传图片
// 			imagePath := "./assets/Title.jpeg" // 替换为你的图片路径
// 			imageFile, err := os.Open(imagePath)
// 			if err != nil {
// 				fmt.Printf("打开图片失败: %v\n", err)
// 				return nil
// 			}
// 			defer imageFile.Close()

// 			imageResp, err := client.Im.Image.Create(context.Background(), larkim.NewCreateImageReqBuilder().
// 				SetImageType(larkim.ImageTypeMessage). // 使用 SetImageType 方法
// 				Image(imageFile).
// 				Build())
// 			if err != nil || !imageResp.Success() {
// 				fmt.Printf("上传图片失败: %v, response: %v\n", err, imageResp)
// 				return nil
// 			}

// 			imageKey := *imageResp.Data.ImageKey
// 			fmt.Printf("图片上传成功，image_key: %s\n", imageKey)

// 			// 构建图片消息
// 			content := fmt.Sprintf(`{"image_key":"%s"}`, imageKey)

// 			// 发送图片消息
// 			resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
// 				ReceiveIdType(larkim.ReceiveIdTypeChatId).
// 				Body(larkim.NewCreateMessageReqBodyBuilder().
// 					MsgType(larkim.MsgTypeImage).
// 					ReceiveId(*event.Event.Message.ChatId).
// 					Content(content).
// 					Build()).
// 				Build())
// 			if err != nil || !resp.Success() {
// 				fmt.Printf("发送图片消息失败: logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
// 				return nil
// 			}

// 			fmt.Println("图片消息发送成功")

// 			// 发送文本消息（包含 replyLink）
// 			textContent := larkim.NewTextMsgBuilder().
// 				TextLine(fmt.Sprintf("这是你要的链接: %s", replyLink)).
// 				Build()

// 			_, err = client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
// 				ReceiveIdType(larkim.ReceiveIdTypeChatId).
// 				Body(larkim.NewCreateMessageReqBodyBuilder().
// 					MsgType(larkim.MsgTypeText).
// 					ReceiveId(*event.Event.Message.ChatId).
// 					Content(textContent).
// 					Build()).
// 				Build())
// 			if err != nil {
// 				fmt.Printf("发送文本消息失败: %v\n", err)
// 			}

// 			return nil
// 		})

// 	/**
// 	 * 启动长连接，并注册事件处理器。
// 	 * Start long connection and register event handler.
// 	 */
// 	cli := larkws.NewClient(app_id, app_secret,
// 		larkws.WithEventHandler(eventHandler),
// 		larkws.WithLogLevel(larkcore.LogLevelDebug),
// 	)
// 	err := cli.Start(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// }

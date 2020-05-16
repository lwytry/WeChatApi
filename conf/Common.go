package conf

type MessageDstType uint8
type MessageType uint8

const (
	MessageDstTypeUser MessageDstType = iota
	MessageDstTypeGroup
)

const (
	MessageTypeUnknown MessageType = iota
	MessageTypeText          // 文字
	MessageTypeImage         // 图片
	MessageTypeExpression    // 表情
	MessageTypeVoice         // 语音
	MessageTypeVideo         // 视频
	MessageTypeURL           // 链接
	MessageTypePosition      // 位置
	MessageTypeBusinessCard  // 名片
	MessageTypeSystem        // 系统
	MessageTypeOther
)

type Message struct {
	Id      	string  			`json:"id" form:"id"`           //消息ID
	UserId  	string  			`json:"userid" form:"userId"`   //谁发的
	DstType 	uint8  				`json:"DstType" form:"DstType"` //群聊还是私聊
	DstID  		string  			`json:"dstid" form:"dstId"`     //对端用户ID/群ID
	MsgType 	uint8  			`json:"msgType" form:"msgType"` //消息类型
	Content		string 			`json:"content"` 				// 消息内容
	Date 		int64 			`json:"date,omitempty" form:"dstId"`     	//时间
}


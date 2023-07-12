# 集合列表

## 用户集合
```
{
    "account":"账号",
    "password":"密码",
    "nickname":"昵称",
    "sex" : 1, // 0-未知 1-男 2-女
    "email": "邮箱",
    "avatar":"头像",
    "created_at": 1, // 创建时间
    "updated_at": 1, // 更新时间
}
```

## Friend 集合
```
{
    
	ID         uint64    自增主键
	UserID     uint64    用户id
	FriendID   uint64    好友id
	CreateTime time.Time 创建时间
	UpdateTime time.Time 更新时间

}
```


## Message 消息集合
```text
{
    ID          uint64    自增id
	UserID      uint64    接收者id
	SenderID    uint64    发送者用户id
	SessionType int8      聊天类型，群聊/单聊
	ReceiverId  uint64    接收者id，群聊id/用户id
	MessageType int8      消息类型,语言、文字、图片
	Content     []byte   消息内容
	Seq         uint64    消息序列号
	SendTime    消息发送时间
	CreateTime  创建时间
	UpdateTime  更新时间

}
```

## Group群集合
```text
{
    ID         uint64    自增主键
	Name       string  群组名称
	OwnerID    uint64  群主id
	CreateTime time.Time 创建时间
	UpdateTime time.Time 更新时间

}
```

## GroupUser 用户群组关联集合
```text
{
	ID         uint64    自增主键
	GroupID    uint64    组id
	UserID     uint64    用户id
	CreateTime time.Time 创建时间
	UpdateTime time.Time 更新时间

}
```

package constvar

const (
	// 机器人BotToken 键值对键名前缀 键名=REDIS_KEY_PREFIX_BOT_TOKEN+botId
	REDIS_KEY_PREFIX_BOT_TOKEN string = "bot_token_"
	// 单用户信息键值对键名前缀 键名=REDIS_KEY_PREFIX_USERINFO+userInfoHash
	REDIS_KEY_PREFIX_USERINFO string = "userinfo_hash_"
	// 所有用户集合键名
	REDIS_SET_KEY_ALL_USER string = "user_setkey_all"
	// 所有关键词集合键名
	REDIS_SET_KEY_ALL_KEYWORD string = "keyword_setkey_all"
	// 单用户关键词集合键名前缀 键名=REDIS_SET_KEY_USER_PREFIX_KEYWORD+userInfoHash
	REDIS_SET_KEY_USER_PREFIX_KEYWORD string = "keyword_setkey_user_"
	// 所有屏蔽词集合键名
	REDIS_SET_KEY_ALL_BLOCKKEYWORD string = "blockkeyword_setkey_all"
	// 所有屏蔽来源会话ID集合键名
	REDIS_SET_KEY_ALL_BLOCKFORMCHATID string = "blockformchatid_setkey_all"
	// 所有屏蔽来源发送者ID集合键名
	REDIS_SET_KEY_ALL_BLOCKFORMSENDERID string = "blockformsenderid_setkey_all"

	// 单用户屏蔽词集合键名 键名=REDIS_SET_KEY_USER_PREFIX_BLOCKKEYWORD+userInfoHash
	REDIS_SET_KEY_USER_PREFIX_BLOCKKEYWORD string = "blockkeyword_setkey_user_"
	// 单用户屏蔽来源会话ID集合键名 键名=REDIS_SET_KEY_USER_PREFIX_BLOCKFORMCHATID+userInfoHash
	REDIS_SET_KEY_USER_PREFIX_BLOCKFORMCHATID string = "blockformchatid_setkey_user_"
	// 单用户屏蔽来源发送者ID集合键名 键名=REDIS_SET_KEY_USER_PREFIX_BLOCKFORMSENDERID+userInfoHash
	REDIS_SET_KEY_USER_PREFIX_BLOCKFORMSENDERID string = "blockformsenderid_setkey_user_"
)

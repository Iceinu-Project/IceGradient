package gradient

// 字体格式转义
const (
	Reset     = "\033[0m" // 重置所有格式
	Bold      = "\033[1m" // 粗体
	Dim       = "\033[2m" // 暗淡
	Italic    = "\033[3m" // 斜体
	Underline = "\033[4m" // 下划线
	Blink     = "\033[5m" // 闪烁
	Reverse   = "\033[7m" // 反转
	Hidden    = "\033[8m" // 隐藏
)

// 字体颜色转义
const (
	Black   = "\033[30m" // 黑色
	Red     = "\033[31m" // 红色
	Green   = "\033[32m" // 绿色
	Yellow  = "\033[33m" // 黄色
	Blue    = "\033[34m" // 蓝色
	Magenta = "\033[35m" // 品红
	Cyan    = "\033[36m" // 青色
	White   = "\033[37m" // 白色
)

// 字体背景转义
const (
	BgBlack   = "\033[40m" // 黑色
	BgRed     = "\033[41m" // 红色
	BgGreen   = "\033[42m" // 绿色
	BgYellow  = "\033[43m" // 黄色
	BgBlue    = "\033[44m" // 蓝色
	BgMagenta = "\033[45m" // 品红
	BgCyan    = "\033[46m" // 青色
	BgWhite   = "\033[47m" // 白色
)

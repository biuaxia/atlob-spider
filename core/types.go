package core

// ParseResult 解析结果对象
type ParseResult struct {
	Requests []Request
}

// Request 发起请求的对象
type Request struct {
	// Url 链接
	Url  string
	Item Item
}

// Parse 解析方法的接口定义
type Parse interface {
	Parser(contents []byte, r Request) ParseResult
}

// Item 单个数据
type Item struct {
	// Parse 解析方法
	Parse

	// Title 标题
	Title string
}

// ParseFunc 解析方法
type ParseFunc func(bytes []byte, r Request) ParseResult

// NilParser 无解析方法
func NilParser(_ []byte, _ Request) ParseResult {
	return ParseResult{}
}

// FuncParser 方法
type FuncParser struct {
	parser ParseFunc
	r      Request
}

// Parser 解析方法
func (f *FuncParser) Parser(contents []byte, r Request) ParseResult {
	return f.parser(contents, r)
}

// CreateParse 创建解析
func CreateParse(parse ParseFunc, r Request) *FuncParser {
	return &FuncParser{
		parse,
		r,
	}
}

// ReadyNotifier 准备就绪通知接口
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Task 任务接口
type Task interface {
	ReadyNotifier
	Submit(Request)
	WorkChan() chan Request
	Run()
}

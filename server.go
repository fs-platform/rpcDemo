package main


type People struct {
}
// Say 1.方法必须包外可见 2.方法必须有个参数,都是导出类型,内建类型 3.第二个参数必须是指针 4. 方法只能返回一个error
func (This People) Say(messages string, reply *string) error {
	*reply = "Aron say:" + messages
	return nil
}

func main() {
	//注册服务名称
     RegisterService(new(People))
     //初始化server
	 InitServer()
}

package array

//用数组实现了链表的简单操作
type  List interface {
	Size() int  //函数大小，返回大小
	Get(index int)(interface{},error) //根据索引抓取数据
	Set(index int,newval interface{})error //设置
	Insert(index int,newval interface{})error //插入
	Append(newval interface{}) error//追加
	Remove(index int)error   //删除
	Clear() //清空
	String() string //返回字符串
}

//结构体
type ArrayList struct{
	dataStore []  interface{}
	theSize  int
}

//初始化
func NewArrayList() *ArrayList  {

	al := ArrayList{
		dataStore: make([]interface{}, 10),
		theSize:   0,
	}
	return &al
}

func (list *ArrayList)Size()int  {
	return list.theSize
}
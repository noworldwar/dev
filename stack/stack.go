package stack

import "fmt"

type Stack struct {
	list []interface{}
}

func New() *Stack {
	s := new(Stack)
	s.list = make([]interface{}, 0, 8)
	return s
}

func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}

// 取出上一次加入的元素
// 也就是最上層的元素 (切片最尾的元素)
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tmp := s.list[len(s.list)-1]

	// 取出時要更新 s.list
	// 這麼做的目的是為了讓 slice 的 len 做更新
	// 透過 len(s.list) 可以不必額外記憶最頂端的 index
	s.list = s.list[:len(s.list)-1]
	return tmp
}

// 將新元素堆入頂端 (切片最尾)
func (s *Stack) Push(element interface{}) {
	// 透過 golang 內建的 append() 函式
	// 自動將新元素加入在切片尾
	// 除此之外該函式會自動更新切片的 len 值
	// 另外若底層陣列不夠時也會自動擴充
	s.list = append(s.list, element)
}

// 方便主程式對整個堆疊做遍歷走仿
// 因為有控製讓 len 對齊最尾元素
// 所以可以直接印出或用 for + range 走訪
func (s *Stack) ToSlice() []interface{} {
	return s.list
}

func printStack(s *Stack) {
	list := s.ToSlice()
	fmt.Println(list)
}

func main() {
	myStack := New()
	fmt.Println("push(小櫻)")
	myStack.Push("小櫻")
	fmt.Println("push(小可)")
	myStack.Push("小可")
	fmt.Println("push(桃矢)")
	myStack.Push("桃矢")
	fmt.Println("push(知世)")
	myStack.Push("知世")
	fmt.Println("push(1)")
	myStack.Push(1)
	fmt.Println("push(2)")
	myStack.Push(2)
	fmt.Println("push(3)")
	myStack.Push(3)
	fmt.Println("push(4)")
	myStack.Push(4)
	printStack(myStack) // [小櫻 小可 桃矢 知世 1 2 3 4]
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // [小櫻 小可 桃矢 知世 1 2 3]
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // [小櫻 小可 桃矢 知世 1 2]
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // [小櫻 小可 桃矢 知世 1]
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // [小櫻 小可 桃矢 知世]
	fmt.Println("push(小郎)")
	myStack.Push("小郎")
	fmt.Println("push(莓鈴)")
	myStack.Push("莓鈴")
	printStack(myStack) // [小櫻 小可 桃矢 知世 小郎 莓鈴]
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // [小櫻 小可 桃矢 知世 小郎]
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // [小櫻 小可 桃矢 知世]
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // [小櫻 小可 桃矢]
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // [小櫻 小可]
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // [小櫻]
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // []
	fmt.Println("pop()", myStack.Pop())
	printStack(myStack) // []
}

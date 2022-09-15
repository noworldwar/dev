package queue

import "fmt"

type Queue struct {
	list   []interface{}
	top    int // 新元素從頭加入
	bottom int // 提取元素從尾端提取
	len    int // 記錄元素個數方便後續處理
}

func New() *Queue {
	q := new(Queue)
	// 預設長度為 8, 容量為 8
	// 在本範例中將slice的長度與容量設為一樣
	// 以方便實作
	q.list = make([]interface{}, 8, 8)
	q.top = -1
	//q.bottom = 0    new()的時候已預設為 0
	//q.len    = 0    new()的時候已預設為 0
	return q
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue) En(val interface{}) {
	// 如果 q.list 滿了就將 q.list 擴大原先的兩倍
	if q.len == cap(q.list) {
		oldCap := cap(q.list)
		newCap := oldCap << 1
		newList := make([]interface{}, newCap, newCap)
		// 一共有 len 個元素所以迴圈 len 次
		for i := 0; i < q.len; i++ {
			newList[i] = q.list[(q.bottom+i)%oldCap]
		}
		q.list = newList
		q.top = q.len - 1
		q.bottom = 0
	}

	// O: 空 X: 有值
	// 欲加入新元素至：
	// OOOOXXXX bottom = 4, top = 7, len = 4, cap(list) = 8
	// 此時將 (top + 1) % cap 則可得 top = 0
	// XOOOXXXX

	q.top = (q.top + 1) % cap(q.list)
	q.list[q.top] = val
	q.len++
}

func (q *Queue) De() interface{} {
	if q.IsEmpty() {
		return nil
	}
	temp := q.list[q.bottom]
	q.bottom = (q.bottom + 1) % cap(q.list)
	q.len--
	return temp
}

func (q *Queue) ToSlice() []interface{} {
	s := make([]interface{}, q.len, q.len)
	cap := cap(q.list)
	for i := 0; i < q.len; i++ {
		s[i] = q.list[(q.bottom+i)%cap]
	}
	return s
}

func main() {
	myQueue := New()
	myQueue.En("小櫻")
	fmt.Println(myQueue.ToSlice()) // [小櫻]
	myQueue.En("知世")
	fmt.Println(myQueue.ToSlice()) // [小櫻 知世]
	myQueue.En("桃矢")
	fmt.Println(myQueue.ToSlice()) // [小櫻 知世 桃矢]
	myQueue.En("小狼")
	fmt.Println(myQueue.ToSlice()) // [小櫻 知世 桃矢 小狼]
	myQueue.En("莓玲")
	fmt.Println(myQueue.ToSlice()) // [小櫻 知世 桃矢 小狼 莓玲]
	myQueue.En("千春")
	fmt.Println(myQueue.ToSlice()) // [小櫻 知世 桃矢 小狼 莓玲 千春]
	myQueue.En("唬爛王山琦")
	fmt.Println(myQueue.ToSlice()) // [小櫻 知世 桃矢 小狼 莓玲 千春 唬爛王山琦]
	myQueue.En("利佳")
	fmt.Println(myQueue.ToSlice()) // [小櫻 知世 桃矢 小狼 莓玲 千春 唬爛王山琦 利佳]
	myQueue.En("寺田老師")
	fmt.Println(myQueue.ToSlice()) // [小櫻 知世 桃矢 小狼 莓玲 千春 唬爛王山琦 利佳 寺田老師]
	myQueue.En("觀月老師")
	fmt.Println(myQueue.ToSlice()) // [小櫻 知世 桃矢 小狼 莓玲 千春 唬爛王山琦 利佳 寺田老師 觀月老師]
	myQueue.En("花店阿姨")
	fmt.Println(myQueue.ToSlice()) // [小櫻 知世 桃矢 小狼 莓玲 千春 唬爛王山琦 利佳 寺田老師 觀月老師 花店阿姨]
	fmt.Println(myQueue.De())      // 小櫻
	fmt.Println(myQueue.De())      // 知世
	fmt.Println(myQueue.De())      // 桃矢
	fmt.Println(myQueue.De())      // 小狼
	fmt.Println(myQueue.De())      // 莓玲
	fmt.Println(myQueue.De())      // 千春
	fmt.Println(myQueue.ToSlice()) // [唬爛王山琦 利佳 寺田老師 觀月老師 花店阿姨]
	myQueue.En("小可")
	fmt.Println(myQueue.ToSlice()) // [唬爛王山琦 利佳 寺田老師 觀月老師 花店阿姨 小可]
	myQueue.En("雪兔")
	fmt.Println(myQueue.ToSlice()) // [唬爛王山琦 利佳 寺田老師 觀月老師 花店阿姨 小可 雪兔]
	fmt.Println(myQueue.De())      // 唬爛王山琦
	fmt.Println(myQueue.De())      // 利佳
	fmt.Println(myQueue.De())      // 寺田老師
	fmt.Println(myQueue.De())      // 觀月老師
	fmt.Println(myQueue.De())      // 花店阿姨
	fmt.Println(myQueue.ToSlice()) // [小可 雪兔]
	fmt.Println(myQueue.De())      // 小可
	fmt.Println(myQueue.De())      // 雪兔
	fmt.Println(myQueue.De())      // <nil>
}

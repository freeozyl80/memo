# 如何用 golang 实现可选参数

# golang  方法集：

```golang
func (b Book) Pages() int {
  return Book.Pages(b)
}
// Book 的方法集为：Pages() int

func (b *Book) SetPages(pages int) {
  (*Book).SetPages(b, pages)
}

// *Book 的方法集为：Pages() int SetPages(pages int)
// 值接收器的方法隐式地同时被声明为指针类型的方法。反之不成立。(重点)

```

# 好好准备下 golang 的函数类型
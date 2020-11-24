package main

import "fmt"

/*
Go中的slice依赖于数组，它的底层就是数组，所以数组具有的优点，slice都有。
且slice支持可以通过append向slice中追加元素，长度不够时会动态扩展，通过再次slice切片，可以得到得到更小的slice结构，可以迭代、遍历等。

实际上slice是这样的结构：先创建一个有特定长度和数据类型的底层数组，然后从这个底层数组中选取一部分元素，返回这些元素组成的集合(或容器)，并将slice指向集合中的第一个元素。
换句话说，slice自身维护了一个指针属性，指向它底层数组中的某些元素的集合。
*/
func main() {
	a := []int{1, 2, 3}
	b := append(a, 1)
	c := a
	c[1] = 1
	d := a[:1]
	d[0] = 2
	fmt.Println(a, b, c, d)
	//和数组不一样，slice是引用 修改c会改变a
	e := []int{}
	fmt.Println(e == nil) //false    var slice []int这样定义的是nil
	//虽然nil slice和Empty slice的长度和容量都为0，输出时的结果都是[]，且都不存储任何数据，但它们是不同的。nil slice不会指向底层数组，而空slice会指向底层数组，只不过这个底层数组暂时是空数组。
	println(e) //[0/0]0xc00007fef0
	//务必记住slice的本质是[x/y]0xADDR，记住它将在很多地方有助于理解slice的特性。另外，个人建议，虽然slice的本质不是指针，但仍然可以将它看作是一种包含了另外两种属性的不纯粹的指针，也就是说，直接认为它是指针。其实不仅slice如此，map也如此。

	f := make([]int, 2, 3)
	f = append(f, 2)
	f = append(f, 2)
	f = append(f, 2) //append是可以突破slice的容量的,突破之后就是原来的两倍，在突破就是继续两倍3->6->12
	fmt.Println(f, cap(f))
	fmt.Printf("%T", f)
	//底层数组扩展时，会生成一个新的底层数组。所以旧底层数组仍然会被旧slice引用，新slice和旧slice不再共享同一个底层数组。

	g := [3]int{1, 2, 3}
	h := g[1:]
	h = append(h, 4)
	i := []int{1, 2, 3}
	fmt.Println(h, i)
	fmt.Printf("%T\n", h)

	j := []int{1, 2, 3}
	foo(j)
	fmt.Println(j)
	/*
		Go中函数的参数是按值传递的，所以调用函数时会复制一个参数的副本传递给函数。
		如果传递给函数的是slice，它将复制该slice副本给函数，这个副本实际上就是[3/5]0xc42003df10，所以传递给函数的副本仍然指向源slice的底层数组。
		换句话说，如果函数内部对slice进行了修改，有可能会直接影响函数外部的底层数组，从而影响其它slice。
		但并不总是如此，例如函数内部对slice进行扩容，扩容时生成了一个新的底层数组，函数后续的代码只对新的底层数组操作，这样就不会影响原始的底层数组。
	*/

	k := []int{1, 2, 3}
	l := 0
	for _, s := range k {
		if s != 2 {
			k[l] = s
			l++
		}
	}
	fmt.Println(k)
	fmt.Println(k[:l])
	println(k, k[:l])

	m := []int{1, 2, 3}
	n := make([]int, 3) //copy函数要根据len来进行，如果不够就只会赋值相应的长度，为0则不复制，
	copy(n, m)
	fmt.Println(n, m, cap(n))

	o := []int{1, 2, 3, 4, 5}
	copy(o[2:], o[3:])
	fmt.Println(o)
}
func foo(s []int) {
	for index, _ := range s {
		s[index] += 1
	}
}

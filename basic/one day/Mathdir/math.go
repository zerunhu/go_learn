package mathtest

/*package即这个文件属于哪个包，我们在外部导入的时候就要通过mathtest.Add来导入这个文件的函数
同时这个package的名字和目录夹名字或者文件的名字不需要一致，只要在需要导入的文件中import ./目录名即可
同一个文件夹下的文件只能有一个包名，否则编译报错 --> Multiple packages in directory: main, Mathdir
*/
func Add(x, y int) int {
	return x + y
}

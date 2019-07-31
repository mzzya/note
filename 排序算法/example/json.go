package main

type student struct{
	id int
	name string
	good bool
}
func main(){
	s:=&student{
		id:1,
		name:"张三",
		good:true
	}
	json.Marshal(s)
}
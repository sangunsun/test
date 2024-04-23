package main
import(
    "fmt"
)

type ss struct{
    name string
}

func (ass ss) print(){
    fmt.Println(ass.name)
}
func (ass *ss)printa(){
    fmt.Println(ass.name+"a")
}
func main(){
    fmt.Println("hello world")
    tss:=ss{
        name:"golang",
    }
    tss.print()
    tss.printa()
}

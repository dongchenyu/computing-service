package main

import "fmt"
func qs(test []int,left int,right int){
	var i int = left
	var j int = right
	var now int = test[left]
	var pos int = left
    for (i<=j){
        for (j>=pos && test[j]>=now){
            j--
        }
        if (j>=pos){
            test[pos]=test[j]
            pos=j
        }
        for (i<=pos && test[i]<=now){
            i++
        }
        if (i<=pos){
            test[pos]=test[i]
            pos=i
        }
    }
    test[pos]=now 
    if (pos-left>1){
        qs(test,left,pos-1)
    }
    if (right-pos>1){
        qs(test,pos+1,right)
    }
}
func main(){
	test := []int{3,5,1,8,9,2,6,7,4,10}
	qs(test,0,len(test)-1)
	fmt.Println(test)
}

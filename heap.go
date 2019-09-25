//结构体数组中参与建堆的为索引为1-10的元素
package main
import "fmt"
type Node struct {
	Value int
}
var n int =10
// 用于构建结构体数组为最小堆，需要调用down函数
func Init(nodes []Node) {
	for i:=n;i>=1;i-- {
		down(nodes,i,n)
    }
}

// 需要down（下沉）的元素在数组中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
    j:=i*2;
    for ;j<=n;{
		if(j+1<=n&&nodes[j+1].Value<nodes[j].Value){
            j+=1;
        }
        if(nodes[j].Value<nodes[i].Value){
			//fmt.Println(nodes[i].Value)
			nodes[j].Value,nodes[i].Value=nodes[i].Value,nodes[j].Value
			//fmt.Println(nodes[j].Value)
			i=j
            j=i*2
        }else{
            break
        }
    }
}

// 用于保证插入新元素(j为元素的索引，数组末尾插入，堆底插入)的结构体数组之后仍然是一个最小堆
func up(nodes []Node, j int) {
	//upAdjust(nodes,1,j)
	//i:=high;
    k:=j/2;
    for ;k>=1; {
        if(nodes[k].Value>nodes[j].Value){
			nodes[k].Value,nodes[j].Value=nodes[j].Value,nodes[k].Value
            j=k;
            k=j/2;
        }else{
            break;
        }
    }
}
// 弹出最小元素，并保证弹出后的结构体数组仍然是一个最小堆
func Pop(nodes []Node) Node {
	res:= nodes[1]
	nodes[1]=nodes[n]
	n-=1
	down(nodes,1,n)
	return res
}
// 保证插入新元素时，结构体数组仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) {
	temp:=node.Value
	n+=1;
    nodes[n].Value=temp;
    up(nodes,n);
}
// 移除数组中指定索引的元素，保证移除后结构体数组仍然是一个最小堆
func Remove(nodes []Node, node Node) {
	tar:=node.Value
	var index int
	for i:=1;i<=n;i+=1 {
		if(nodes[i].Value==tar){
			index=i
			break
		}
	}
	for i:=index;i<n;i++ {
		nodes[i].Value=nodes[i+1].Value
	}
	n--;
	Init(nodes)
}

func main() {
	nodes:=[]Node{
		Node{
			11,
		},
		Node{
			10,
		},
		Node{
			9,
		},
		Node{
			8,
		},
		Node{
			7,
		},
		Node{
			6,
		},
		Node{
			5,
		},
		Node{
			4,
		},
		Node{
			3,
		},
		Node{
			2,
		},
		Node{
			1,
		},
	}
	Init(nodes)
	fmt.Println(Pop(nodes))
	var node Node
	node.Value=12
	Push(node,nodes)
	for i:=1;i<=n;i++{
		fmt.Println(nodes[i].Value)
	}
	fmt.Println()
	var node2 Node
	node2.Value=5
	Remove(nodes,node2)
	for i:=1;i<=n;i++{
		fmt.Println(nodes[i].Value)
	}
}
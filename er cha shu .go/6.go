package main
func Depth(n*Node)int{
	var depleft,depright int
	if n==nil{
		return 0
	}else{
		depleft=Depth(n.Left)
		depright=Depth(n.Right)
		if depleft>depright{
			return depl'+1'
		}else{
			return depright+1
		}
	}
}


func LeafCount(n*Node)int{
	if n==nil{
		return 0
	}else if(n.Left==nil)&&(n.Right==nil){
		return 1
	}else{
		return (LeafCount(n.Left)+LeafCount(n.Right))
	}
}
package maze

import "errors"

type PointStack struct {
	_data []int64
}

func NewPointStack() *PointStack {
	return &PointStack{_data:make([]int64,0,1000)}
}


func (p *PointStack)Push(x,y int) {
	temp:=int64(x)<<32+int64(y)
	p._data=append(p._data,temp)
}

func (p *PointStack)Pop() (x,y int) {
	temp:=p._data[len(p._data)-1]
	p._data=p._data[:len(p._data)-1]
	tx:=temp>>32
	ty:=temp&0xffffffff

	return int(tx),int(ty)
}

func (p *PointStack)Last() (x,y int,ok bool){
	if len(p._data)<2 {
		return -10000,-10000,false
	} else {
		temp:=p._data[len(p._data)-1]
		tx:=temp>>32
		ty:=temp&0xffffffff
		return	int(tx),int(ty),true
	}

}

func (p *PointStack)HasPoint(x,y int) bool {
	temp:=int64(x)<<32+int64(y)
	for i:=0;i<len(p._data);i++ {

		if p._data[i]==temp {
			return true
		}
	}

	return false
}

func (p *PointStack)Index(i int) (x,y int) {
	if len(p._data)-1<i {
		panic(errors.New("i is too bigger"))
	}

	temp:=p._data[i]
	tx:=temp>>32
	ty:=temp&0xffffffff
	return	int(tx),int(ty)
}

func (p *PointStack)Count() int {
	return len(p._data)
}
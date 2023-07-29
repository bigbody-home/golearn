package main

import "path"

type Listener interface {
	Do()
	GetName() string
}

type Watcher interface {
	Add(l *Listener)
	Remove(l *Listener)
	Notify()
}
type HuangRong struct {
	name string
}

func (h *HuangRong) Do() {
	println(h.name, "收到江湖追杀令，黄蓉前往丐帮主持大局。。。")
}
func (h *HuangRong) GetName() string {
	return h.name
}

type XuZhu struct {
	name string
}

func (h *XuZhu) Do() {
	println(h.name, "收到江湖追杀令，虚竹前往少林寺营救方丈。。。")
}
func (x *XuZhu) GetName() string {
	return x.name
}

type BaiXiaoSheng struct {
	l []Listener
}

func (b *BaiXiaoSheng) Add(l Listener) {
	println("百晓生添加新的观察者", l.GetName())
	b.l = append(b.l, l)
}
func (b *BaiXiaoSheng) Remove(l Listener) {
	for i, v := range b.l {
		if v.GetName() == l.GetName() {
			b.l = append(b.l[:i], b.l[i+1:]...)
			println("百晓生删除观察者", l.GetName())
			break
		}
	}
}
func (b *BaiXiaoSheng) Notify() {
	println("江湖追杀令已经发出")
	for _, v := range b.l {
		println("请", v.GetName(), "进行相应行动！")
		v.Do()
	}
}

func main() {
	h := &HuangRong{name: "黄蓉"}
	x := &XuZhu{name: "虚竹"}
	b := &BaiXiaoSheng{l: make([]Listener, 0, 3)}
	b.Add(h)
	b.Add(x)
	b.Remove(x)
	println(len(b.l))
	b.Notify()
	s := path.Join("/aa/", "ab")
	println(s)
}

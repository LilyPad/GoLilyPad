package packet

import (
	"errors"
	"fmt"
	"io"
)

type PacketPipeline struct {
	childrenMap map[string]PacketPipelineChild
	children    []PacketPipelineChild
}

func NewPacketPipeline() (this *PacketPipeline) {
	this = new(PacketPipeline)
	this.childrenMap = make(map[string]PacketPipelineChild)
	this.children = make([]PacketPipelineChild, 0)
	return
}

func (this *PacketPipeline) Decode(reader io.Reader) (packet Packet, err error) {
	packet, err = this.children[0].Decode(reader)
	return
}

func (this *PacketPipeline) Encode(writer io.Writer, packet Packet) (err error) {
	err = this.children[0].Encode(writer, packet)
	return
}

func (this *PacketPipeline) AddFirst(name string, child PacketPipelineChild) (err error) {
	if _, ok := this.childrenMap[name]; ok {
		err = errors.New(fmt.Sprintf("PacketPipeline addFirst, Child %s duplicate", name))
		return
	}
	this.childrenMap[name] = child
	this.children = append(this.children, nil)
	copy(this.children, this.children[1:])
	this.children[0] = child
	this.rebuild()
	return
}

func (this *PacketPipeline) AddLast(name string, child PacketPipelineChild) (err error) {
	if _, ok := this.childrenMap[name]; ok {
		err = errors.New(fmt.Sprintf("PacketPipeline addLast, Child %s duplicate", name))
		return
	}
	this.childrenMap[name] = child
	this.children = append(this.children, child)
	this.rebuild()
	return
}

func (this *PacketPipeline) AddBefore(name string, subject string, child PacketPipelineChild) (err error) {
	var ok bool
	if _, ok = this.childrenMap[name]; ok {
		err = errors.New(fmt.Sprintf("PacketPipeline addBefore, Child %s duplicate", name))
		return
	}
	var subjectChild PacketPipelineChild
	if subjectChild, ok = this.childrenMap[subject]; !ok {
		err = errors.New(fmt.Sprintf("PacketPipeline addBefore, Child %s no such subject", subject))
		return
	}
	i := -1
	for j, match := range this.children {
		if match != subjectChild {
			continue
		}
		i = j
		break
	}
	this.childrenMap[name] = child
	this.children = append(this.children, nil)
	copy(this.children[i+1:], this.children[i:])
	this.children[i] = child
	this.rebuild()
	return
}

func (this *PacketPipeline) AddAfter(name string, subject string, child PacketPipelineChild) (err error) {
	var ok bool
	if _, ok = this.childrenMap[name]; ok {
		err = errors.New(fmt.Sprintf("PacketPipeline addAfter, Child %s duplicate", name))
		return
	}
	var subjectChild PacketPipelineChild
	if subjectChild, ok = this.childrenMap[subject]; !ok {
		err = errors.New(fmt.Sprintf("PacketPipeline addAfter, Child %s no such subject", subject))
		return
	}
	i := -1
	for j, match := range this.children {
		if match != subjectChild {
			continue
		}
		i = j
		break
	}
	this.childrenMap[name] = child
	if len(this.children) == 1 {
		this.children = append(this.children, child)
	} else {
		this.children = append(this.children, nil)
		copy(this.children[i+2:], this.children[i+1:])
		this.children[i+1] = child
	}
	this.rebuild()
	return
}

func (this *PacketPipeline) Replace(name string, child PacketPipelineChild) (err error) {
	var subjectChild PacketPipelineChild
	var ok bool
	if subjectChild, ok = this.childrenMap[name]; !ok {
		err = errors.New(fmt.Sprintf("PacketPipeline replace, Child %s no such child", name))
		return
	}
	this.childrenMap[name] = child
	for i, match := range this.children {
		if match != subjectChild {
			continue
		}
		this.children[i] = child
		break
	}
	this.rebuild()
	return
}

func (this *PacketPipeline) Remove(name string) (err error) {
	var child PacketPipelineChild
	var ok bool
	if child, ok = this.childrenMap[name]; !ok {
		err = errors.New(fmt.Sprintf("PacketPipeline remove, Child %s no such child", name))
		return
	}
	this.childrenMap[name] = nil
	i := -1
	for j, match := range this.children {
		if match != child {
			continue
		}
		i = j
		break
	}
	this.children = append(this.children[:i], this.children[i+1:]...)
	this.rebuild()
	return
}

func (this *PacketPipeline) Get(name string) (child PacketPipelineChild) {
	child = this.childrenMap[name]
	return
}

func (this *PacketPipeline) HasName(name string) (ok bool) {
	_, ok = this.childrenMap[name]
	return
}

func (this *PacketPipeline) rebuild() {
	for i := 0; i < len(this.children)-1; i++ {
		this.children[i].SetCodec(this.children[i+1])
	}
}

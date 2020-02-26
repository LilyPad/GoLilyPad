package minecraft

import (
	"fmt"
)

type VersionTable struct {
	latest      *Version
	versionById map[int]*Version
}

func NewVersionTable() (this *VersionTable) {
	this = new(VersionTable)
	this.versionById = make(map[int]*Version)
	return
}

func NewVersionTableFrom(versions ...*Version) (this *VersionTable) {
	this = new(VersionTable)
	this.versionById = make(map[int]*Version)
	for _, version := range versions {
		this.Register(version)
	}
	return
}

func (this *VersionTable) Register(register *Version) {
	if len(register.NameLatest) == 0 {
		register.NameLatest = register.Name
	}
	for _, id := range register.Id {
		if id > register.IdLatest {
			register.IdLatest = id
		}
		if duplicate, ok := this.versionById[id]; ok {
			panic(fmt.Sprintf("Duplicate version, id: %d register: %s duplicate: %s", id, register.Name, duplicate.Name))
		}
		this.versionById[id] = register
	}
	if this.latest == nil || this.latest.IdLatest < register.IdLatest {
		this.latest = register
	}
}

func (this *VersionTable) ById(id int) *Version {
	return this.versionById[id]
}

func (this *VersionTable) Latest() *Version {
	return this.latest
}

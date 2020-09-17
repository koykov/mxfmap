package mxf

type Identifier interface {
	GetId() int
}

type Imap struct {
	p  []Identifier
	lk uint32
}

func NewImap(size int) *Imap {
	m := Imap{}
	m.Grow(size)
	return &m
}

func (m *Imap) Len() int {
	return len(m.p)
}

func (m *Imap) Grow(size int) {
	if size < 0 {
		return
	}
	if m.p == nil {
		m.p = make([]Identifier, size)
	} else {
		grow := size - len(m.p)
		if grow <= 0 {
			return
		}
		m.p = append(m.p, make([]Identifier, grow)...)
	}
}

func (m *Imap) Set(idx int, x Identifier) {
	if idx >= len(m.p) {
		m.Grow(idx * 2)
	}
	m.p[idx] = x
}

func (m *Imap) BulkSet(l []Identifier) {
	m.p = append(m.p[:0], l...)
}

func (m *Imap) Get(idx int) Identifier {
	if idx < len(m.p) {
		return m.p[idx]
	}
	return nil
}

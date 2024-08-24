package server

import (
	"fmt"
	"sync"
)

// Registro
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

// Manejador de registros
type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

// Manejador de error por valor de offset inválido
var ErrOffsetNotFound = fmt.Errorf("offset not found")

// Agregar el registro al set
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// Leer el registro del set
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

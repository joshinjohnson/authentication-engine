package models

type Config struct {
	Mode Mode
}

type Mode int

const (
	_ Mode = iota
	Emulation
	Operation
)
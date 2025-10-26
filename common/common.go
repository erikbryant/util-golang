package common

type Integers interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Floats interface {
	float32 | float64
}

type Numbers interface {
	Integers | Floats
}

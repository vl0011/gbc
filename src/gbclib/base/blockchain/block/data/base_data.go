package data

type BaseData struct {
	data []byte
	size uint32
}

func NewData(max uint32) BaseData {
	return BaseData{
		data: make([]byte, max),
		size: max,
	}
}

func (data *BaseData) GetDataSize() uint32 {
	return data.size
}

func (data *BaseData) GetData() []byte {
	return data.data
}

func (data *BaseData) WriteData(d []byte) {
	copy(data.data, d)
}

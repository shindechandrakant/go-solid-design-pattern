package design_pattern

type DataSource interface {
	WriteData(data string)
	ReadData() string
}

type FileDataSource struct {
	FileName string
	Storage  string
}

func NewDataSource(fileName string) *FileDataSource {
	return &FileDataSource{
		FileName: fileName,
	}
}

func (f *FileDataSource) WriteData(data string) {
	f.Storage = data
}

func (f *FileDataSource) ReadData() string {
	return f.Storage
}

type EncryptionDecorator struct {
	Wrapped DataSource
}

func NewEncryptionDecorator(source DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{
		Wrapped: source,
	}
}

func (e *EncryptionDecorator) WriteData(data string) {
	encrypted := "encrypted: " + data
	e.Wrapped.WriteData(encrypted)
}

func TimePrefix(data, toTrim string) string {
	if len(data) >= len(toTrim) && data[:len(toTrim)] == toTrim {
		return data[len(toTrim):]
	}
	return data
}

func (e *EncryptionDecorator) ReadData() string {
	data := e.Wrapped.ReadData()
	return TimePrefix(data, "encrypted: ")
}

type CompressionDecorator struct {
	Wrapped DataSource
}

func NewCompressionDecorator(source DataSource) *CompressionDecorator {
	return &CompressionDecorator{
		Wrapped: source,
	}
}

func (d *CompressionDecorator) WriteData(data string) {
	compressed := "compresses: " + data
	d.Wrapped.WriteData(compressed)
}

func (d *CompressionDecorator) ReadData() string {
	data := d.Wrapped.ReadData()
	return TimePrefix(data, "Compressed:")
}

func main() {
	var source DataSource = NewDataSource("data.txt")
	source = NewEncryptionDecorator(source)
	source = NewCompressionDecorator(source)
	source.WriteData("Sensitive Info")
}

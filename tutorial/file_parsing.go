package main

import (
	"fmt"
)

//1. Validate the file — check file size (max 10MB), check file extension (only .csv, .json, .xml allowed),
//and check that the file isn't empty
//2. Parse the file — CSV files are parsed into rows, JSON files are unmarshalled into maps,
//XML files are decoded into structs. Each format has completely different parsing logic
//3. Transform the data — apply transformations like trimming whitespace, converting strings to lowercase,
//removing duplicates. Multiple transformations can be applied in sequence
//4. Export the result — the processed data can be exported as JSON to disk, sent to an S3 bucket,
//or POSTed to a webhook URL

// validation it will be common for all the type of files
// different parsing logic for all the formats
// transform pipeline
// export methods

func ReadDataFromDisk() (*File, error) {
	return &File{
		FileName:  "File.json",
		Size:      1000,
		Extension: "json",
		data:      "Hii this is data",
	}, nil
}

type FileData []map[string]string

type Exporter interface {
	Export(FileData) error
}

type JsonExporter struct{}
type S3Exporter struct{}
type WebhookExporter struct{}

func (j *JsonExporter) Export(data FileData) error {
	fmt.Println("Exporting to json")
	return nil
}

func (s3 *S3Exporter) Export(data FileData) error {
	fmt.Println("Exporting to S3")
	return nil
}

func (w *WebhookExporter) Export(data FileData) error {
	fmt.Println("Exporting to Webhook")
	return nil
}

type File struct {
	FileName  string
	Size      int64
	Extension string
	data      any
}

type Parser interface {
	Parse(*File) (FileData, error)
}

type JSONParser struct{}
type XMLParser struct{}
type CSVParser struct{}

func (f *JSONParser) Parse(file *File) (FileData, error) {
	return []map[string]string{}, nil
}

func (x *XMLParser) Parse(file *File) (FileData, error) {
	return []map[string]string{}, nil
}
func (c *CSVParser) Parse(file *File) (FileData, error) {
	return []map[string]string{}, nil
}

type Transformer interface {
	Transform([]map[string]string) ([]map[string]string, error)
}

type TrimSpace struct{}

func (t *TrimSpace) Transform([]map[string]string) ([]map[string]string, error) {
	return []map[string]string{}, nil
}

type FileProcessor struct {
	Parser      map[string]Parser
	Transformer []Transformer
	Exporter    Exporter
}

func NewFileProcessor(parser map[string]Parser, transformer []Transformer, exporter Exporter) *FileProcessor {
	return &FileProcessor{
		Parser:      parser,
		Transformer: transformer,
		Exporter:    exporter,
	}
}

func ValidateFile(file *File) error {

	if file.Size >= 1000 {
		return fmt.Errorf("file size is big")
	}
	return nil
}

func (p *FileProcessor) Process(file *File) (FileData, error) {

	if err := ValidateFile(file); err != nil {
		return nil, err
	}

	_, exists := p.Parser[file.Extension]
	if !exists {
		return nil, fmt.Errorf("unsupported file extension")
	}

	parser, _ := p.Parser[file.Extension]
	data, err := parser.Parse(file)
	if err != nil {
		return nil, err
	}

	for _, transform := range p.Transformer {
		data, err = transform.Transform(data)
		if err != nil {
			return nil, err
		}
	}

	err = p.Exporter.Export(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	data, _ := ReadDataFromDisk()
	processor := &FileProcessor{
		Parser:      map[string]Parser{"json": &JSONParser{}},
		Transformer: []Transformer{&TrimSpace{}},
		Exporter:    &S3Exporter{},
	}

	processedData, err := processor.Process(data)
	fmt.Println(err)
	fmt.Println(processedData)
}

// validation

// entities
//	1. Files
//		- fileName
//		- Size
//		- extension
//		- data
//		parse()

// FLUX V1 DEPRECATION NOTICE. https://github.com/weaveworks/eksctl/issues/2963
package fileprocessor

// File represents the contents and the path to a file
type File struct {
	Path string
	Data []byte
}

// FileProcessor can process a template file and produce an output file
type FileProcessor interface {
	ProcessFile(file File) (File, error)
}

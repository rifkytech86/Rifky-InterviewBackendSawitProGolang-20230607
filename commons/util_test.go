package commons

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestNewFileReader(t *testing.T) {
	tests := []struct {
		name string
		want IFileReader
	}{
		{
			name: "test file reader",
			want: NewFileReader(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewFileReader(), "NewFileReader()")
		})
	}
}

func Test_fileReader_ReadFile(t *testing.T) {
	// Create an instance of the fileReader
	fileReader := NewFileReader()

	// Define a temporary file for testing
	tempFile := "temp.txt"
	content := []byte("test content")
	err := ioutil.WriteFile(tempFile, content, 0644)
	assert.NoError(t, err)

	// Clean up the temporary file after the test
	defer func() {
		err := os.Remove(tempFile)
		assert.NoError(t, err)
	}()

	// Test reading the file
	readContent, err := fileReader.ReadFile(tempFile)
	assert.NoError(t, err)
	assert.Equal(t, content, readContent)
}

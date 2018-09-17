package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCopy(t *testing.T) {
	tempDir, err := ioutil.TempDir(".", "test-copy-")
	if err != nil {
		t.Fail()
	}
	tempFile, err := ioutil.TempFile(tempDir, "tmp-")
	if err != nil {
		t.Fail()
	}
	defer os.Remove(tempDir)
	defer os.Remove(tempFile.Name())

	tempDirAbsPath, _ := filepath.Abs(tempDir)
	fmt.Println("TempDirAbsPath: " + tempDirAbsPath)
	fmt.Println("TempFile: " + tempFile.Name())
	tempFileAbsPath, _ := filepath.Abs(tempFile.Name())
	fmt.Println("TempFileAbsPath: " + tempFileAbsPath)

	for i := 0; i < 10; i++ {
		tempFile.WriteString(fmt.Sprintf("index: %d\n", i))
	}
	err = tempFile.Close()
	if err != nil {
		t.Fail()
	}

	err = Copy(tempFileAbsPath, fmt.Sprint(tempDirAbsPath, "/", fmt.Sprintf("dst-%s", tempFile.Name())), false)
	if err != nil {
		t.Error(err)
	}

}

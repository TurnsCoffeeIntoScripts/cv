package utils

import (
	"cv/pkg/terminal"
	"net/http"
	"os"
)

func CloseFileHandler(f *os.File, filename string) {
	err := f.Close()
	if err != nil {
		terminal.ErrorMessage("Failed to close %s file handler", filename)
	}
}

func CloseHttpReadCloser(resp *http.Response) {
	if resp != nil {
		if resp.Body != nil {
			err := resp.Body.Close()
			if err != nil {
				terminal.ErrorMessage("Failed to close ReadCloser in an http context")
			}
		}
	}
}

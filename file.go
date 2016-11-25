// Copyright (c) 2016 Betalo AB
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"errors"
	"net/url"
	"os"
	"path/filepath"

	"context"
)

type fileResource struct {
	url.URL
}

func (r *fileResource) Await(context.Context) error {
	// Unify absolute and relative file paths
	filePath := filepath.Join(r.URL.Host, r.URL.Path)

	tags := parseTags(r.URL.Fragment)

	_, err := os.Stat(filePath)
	if _, ok := tags["absent"]; ok {
		if err == nil {
			return &unavailabilityError{errors.New("file exists")}
		} else if os.IsNotExist(err) {
			return nil
		}
	} else {
		if err == nil {
			return nil
		} else if os.IsNotExist(err) {
			return &unavailabilityError{err}
		}
	}

	return err
}

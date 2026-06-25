// Copyright (C) 2021 CGI France
//
// This file is part of PIMO.
//
// PIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// PIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with PIMO.  If not, see <http://www.gnu.org/licenses/>.

package uri

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/maskingdata"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestUriReaderShouldCreateListFromDoc(t *testing.T) {
	nameList, err := Read("file://../../test/names.txt")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	waitedList := []model.Entry{"Mickael", "Marc", "Benjamin"}
	assert.Equal(t, waitedList, nameList.Collect(), "Should return the right list")
}

func TestUriReaderShouldCreateListFromLink(t *testing.T) {
	link := "https://gist.githubusercontent.com/youencgi/68548750b266136db183ad4bdfd6436a/raw/c735e55a9553f3db7f649c16616ebf22b46d7bfb/gistfile1.txt"
	nameList, err := Read(link)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	waitedList := []model.Entry{"Mickael", "Marc", "Benjamin"}
	assert.Equal(t, waitedList, nameList.Collect(), "Should return the right list")
}

func TestUriReaderShouldCreateListFromInsideFiles(t *testing.T) {
	nameList, err := Read("pimo://nameFR")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	// nolint: gocritic
	waitedList := append(maskingdata.NameFRM, maskingdata.NameFRF...)
	for i := range waitedList {
		assert.Equal(t, waitedList[i], nameList.Get(i), "Should return the right list")
	}

	nameListES, err := Read("pimo://nameES")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	// nolint: gocritic
	waitedListES := append(maskingdata.NameESM, maskingdata.NameESF...)
	for i := range waitedListES {
		assert.Equal(t, waitedListES[i], nameListES.Get(i), "Should return the right list")
	}

	surnameListES, err := Read("pimo://surnameES")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	for i := range maskingdata.SurnameES {
		assert.Equal(t, maskingdata.SurnameES[i], surnameListES.Get(i), "Should return the right list")
	}
}

func TestUriReaderShouldCreateGermanNameListFromInsideFiles(t *testing.T) {
	nameList, err := Read("pimo://nameGE")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	// nolint: gocritic
	waitedList := append(maskingdata.NameGEM, maskingdata.NameGEF...)
	for i := range waitedList {
		assert.Equal(t, waitedList[i], nameList.Get(i), "Should return the right list")
	}
}

func TestUriReaderShouldCreateGermanSurnameListFromInsideFiles(t *testing.T) {
	nameList, err := Read("pimo://surnameGE")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	waitedList := maskingdata.SurnameGE
	for i := range waitedList {
		assert.Equal(t, waitedList[i], nameList.Get(i), "Should return the right list")
	}
}

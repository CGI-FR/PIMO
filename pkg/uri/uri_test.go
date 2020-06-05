package uri

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/maskingdata"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestUriReaderShouldCreateListFromDoc(t *testing.T) {
	nameList, err := Read("file://../../test/names.txt")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	waitedList := []model.Entry{"Mickael", "Marc", "Benjamin"}
	assert.Equal(t, waitedList, nameList, "Should return the right list")
}

func TestUriReaderShouldCreateListFromLink(t *testing.T) {
	link := "https://gist.githubusercontent.com/youencgi/68548750b266136db183ad4bdfd6436a/raw/c735e55a9553f3db7f649c16616ebf22b46d7bfb/gistfile1.txt"
	nameList, err := Read(link)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	waitedList := []model.Entry{"Mickael", "Marc", "Benjamin"}
	assert.Equal(t, waitedList, nameList, "Should return the right list")
}

func TestUriReaderShouldCreateListFromInsideFiles(t *testing.T) {
	nameList, err := Read("pimo://nameFR")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	waitedList := append(maskingdata.NameFRM, maskingdata.NameFRF...)
	for i := range waitedList {
		assert.Equal(t, waitedList[i], nameList[i], "Should return the right list")
	}
}

package uri

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

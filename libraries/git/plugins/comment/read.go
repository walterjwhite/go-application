package comment

import (
	"github.com/walterjwhite/go-application/libraries/git"
	"time"

	"github.com/walterjwhite/go-application/libraries/foreachfile"
	"github.com/walterjwhite/go-application/libraries/logging"
	"path/filepath"
	"strings"

	//"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"strconv"
)

var (
	// @see: activity package
	timeZone string
	Location *time.Location
)

type comments struct {
	Comments []*Comment
}

func init() {
	var err error
	Location, err = time.LoadLocation(timeZone)
	logging.Panic(err)
}

func Read(w *git.WorkTreeConfig) []*Comment {
	c := &comments{}

	foreachfile.Execute(filepath.Join(w.Path, commentPath), c.getComment)

	return c.Comments
}

func (c *comments) getComment(filePath string) {
	s := strings.Split(filePath, string(filepath.Separator))

	timeComponents := strings.Split(s[len(s)-1], ".")

	// TODO: convert string to integer
	nanoSecond, err := strconv.Atoi(timeComponents[3])
	logging.Panic(err)

	second, err := strconv.Atoi(timeComponents[2])
	logging.Panic(err)

	minute, err := strconv.Atoi(timeComponents[1])
	logging.Panic(err)

	hour, err := strconv.Atoi(timeComponents[0])
	logging.Panic(err)

	day, err := strconv.Atoi(s[len(s)-2])
	logging.Panic(err)

	month, err := strconv.Atoi(s[len(s)-3])
	logging.Panic(err)

	year, err := strconv.Atoi(s[len(s)-4])
	logging.Panic(err)

	date := time.Date(year, time.Month(month), day, hour, minute, second, nanoSecond, Location)

	// read file contents
	messageContents, err := ioutil.ReadFile(filePath)
	logging.Panic(err)

	comment := &Comment{Message: string(messageContents), DateTime: date}
	c.Comments = append(c.Comments, comment)
}
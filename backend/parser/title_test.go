package parser

import (
	"e-book-manager/epub/epubReader"
	"testing"
)

func TestGetTitle__titleFound_expect_returning_title(t *testing.T) {
	var titles []epubReader.Title
	titles = append(titles, epubReader.Title{Text: "Test Title"})
	title, err := GetTitle(epubReader.Metadata{Title: &titles})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if title != "Test Title" {
		t.Log("wrong title " + title)
		t.Fail()
	}
}

func TestGetTitle__noTitleFound_expect_error(t *testing.T) {
	var titles []epubReader.Title
	title, err := GetTitle(epubReader.Metadata{Title: &titles})
	if err == nil {
		t.Log("error was expected")
		t.Fail()
	} else if err.Error() != "no title found" {
		t.Log("wrong error " + err.Error())
		t.Fail()
	}
	if title != "" {
		t.Log("title must be empty but is " + title)
		t.Fail()
	}
}

func TestGetTitle__toManyTitleFound_expect_error(t *testing.T) {
	var titles []epubReader.Title
	titles = append(titles, epubReader.Title{Text: "Test Title 1"})
	titles = append(titles, epubReader.Title{Text: "Test Title 2"})
	title, err := GetTitle(epubReader.Metadata{Title: &titles})
	if err == nil {
		t.Log("error was expected")
		t.Fail()
	} else if err.Error() != "to many title found" {
		t.Log("wrong error " + err.Error())
		t.Fail()
	}
	if title != "" {
		t.Log("title must be empty but is " + title)
		t.Fail()
	}
}

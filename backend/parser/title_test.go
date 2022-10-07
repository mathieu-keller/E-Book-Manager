package parser

import (
	"archive/zip"
	"e-book-manager/epub/epubReader"
	"testing"
)

func TestGetTitle__titleFound_expect_returning_title(t *testing.T) {
	file, err := zip.OpenReader("../testFiles/test.epub")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	bookFile, err := epubReader.Open(&file.Reader)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	title, err := GetTitle(bookFile)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if title != "Test Title" {
		t.Log("wrong title " + title)
		t.Fail()
	}
}

func TestGetTitle__noTitleFound_expect_error(t *testing.T) {
	file, err := zip.OpenReader("../testFiles/test2.epub")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	bookFile, err := epubReader.Open(&file.Reader)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	title, err := GetTitle(bookFile)
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
	file, err := zip.OpenReader("../testFiles/test3.epub")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	bookFile, err := epubReader.Open(&file.Reader)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	title, err := GetTitle(bookFile)
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

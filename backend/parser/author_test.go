package parser

import (
	"e-book-manager/epub/epubReader"
	"testing"
)

func TestGetCreator__CreatorFound_expect_returning_Creator(t *testing.T) {
	var creators []epubReader.Creator
	creators = append(creators, epubReader.Creator{Text: "Test Creator"})
	creator, err := getCreator(epubReader.Metadata{Creator: &creators})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if creator != "Test Creator" {
		t.Log("wrong creator " + creator)
		t.Fail()
	}
}

func TestGetCreator__noCreatorFound_expect_error(t *testing.T) {
	var creators []epubReader.Creator
	creator, err := getCreator(epubReader.Metadata{Creator: &creators})
	if err == nil {
		t.Log("error was expected")
		t.Fail()
	} else if err.Error() != "no creator found" {
		t.Log("wrong error " + err.Error())
		t.Fail()
	}
	if creator != "" {
		t.Log("creator must be empty but is " + creator)
		t.Fail()
	}
}

func TestGetCreator__toManyCreatorsFound_expect_error(t *testing.T) {
	var creators []epubReader.Creator
	creators = append(creators, epubReader.Creator{Text: "Test Creator 1"})
	creators = append(creators, epubReader.Creator{Text: "Test Creator 2"})
	creator, err := getCreator(epubReader.Metadata{Creator: &creators})
	if err == nil {
		t.Log("error was expected")
		t.Fail()
	} else if err.Error() != "to many creator found" {
		t.Log("wrong error " + err.Error())
		t.Fail()
	}
	if creator != "" {
		t.Log("creator must be empty but is " + creator)
		t.Fail()
	}
}

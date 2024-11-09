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
	} else if len(creator) > 1 && len(creator) < 1 {
		t.Log("wrong creator count")
		t.Fail()
	} else if creator[0] != "Test Creator" {
		t.Log("wrong creator " + creator[0])
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
	if creator != nil {
		t.Log("creator must be nil")
		t.Fail()
	}
}

func TestGetCreator__toManyCreatorsFound_expect_error(t *testing.T) {
	var creators []epubReader.Creator
	creators = append(creators, epubReader.Creator{Text: "Test Creator 1"})
	creators = append(creators, epubReader.Creator{Text: "Test Creator 2"})
	creator, err := getCreator(epubReader.Metadata{Creator: &creators})
	if err != nil {
		t.Log("error was expected")
		t.Fail()
	} else if len(creator) != 2 {
		t.Log("creator must be a length of 2")
		t.Fail()
	}
}

func TestGetCreator__toManyCreatorsFound_but_one_of_them_has_role_attr_aut_expect_returning_Creator(t *testing.T) {
	var creators []epubReader.Creator
	creators = append(creators, epubReader.Creator{Text: "Test Creator 1", Role: "bkp"})
	creators = append(creators, epubReader.Creator{Text: "Test Creator 2", Role: "aut"})
	creator, err := getCreator(epubReader.Metadata{Creator: &creators})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if len(creator) != 1 {
		t.Log("creator must be a length of 1")
		t.Fail()
	} else if creator[0] != "Test Creator 2" {
		t.Log("wrong creator " + creator[0])
		t.Fail()
	}
}

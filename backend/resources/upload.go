package resources

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mathieu-keller/epub-parser"
)

func InitUploadApi(r fiber.Router) {
	r.Post("/upload", func(c *fiber.Ctx) error {
		files, err := c.MultipartForm()
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		fileErrors := ""
		for ix, fileHeader := range files.File["myFiles"] {
			println(ix, " / ", len(files.File["myFiles"]))
			if fileHeader.Header.Get("Content-Type") != "application/epub+zip" {
				fileErrors += "Error: Book " + fileHeader.Filename + ": is not in epub format\n"
				continue
			}
			file, err := fileHeader.Open()
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
			binaryFile := make([]byte, fileHeader.Size)
			fileLength, err := file.Read(binaryFile)
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
			zipReader, err := zip.NewReader(bytes.NewReader(binaryFile), int64(fileLength))
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
			book, err := epub.OpenBook(zipReader)
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}

			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
		}
		if len(fileErrors) > 0 {
			return c.Status(400).SendString(fileErrors)
		} else {
			return c.Status(200).SendString("Done")
		}
	})
}

type dbBook struct {
	Id    string
	Title string `gorm:"index:idx_book_title;not null"`
}

func save(book *epub.Book) {
	id := ""
	for _, identifier := range *book.Opf.Metadata.Identifier {
		if identifier.Id == book.Opf.UniqueIdentifier {
			id = identifier.Text
		}
	}
	if id == "" {
		return //error
	}

	title, _ := GetTitle(book)
	dbBook{Id: id, Title: title}
}

type Title struct {
	Name      string
	Seq       int
	TitleType string
	Lang      string
}

func GetTitle(book *epub.Book) (string, error) {
	if book.Opf.Metadata.Title == nil {
		return "", errors.New("no title found")
	}
	titles := *book.Opf.Metadata.Title
	if len(titles) == 0 {
		return "", errors.New("no title found")
	}
	if len(titles) > 1 {
		return "", errors.New("to many title found")
	}
	return titles[0].Text, nil
}

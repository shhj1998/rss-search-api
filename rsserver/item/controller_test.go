package item_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mmcdole/gofeed"
	"github.com/shhj1998/rss-search-api/rsserver"
	"github.com/shhj1998/rss-search-api/rsserver/item"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestController_GetItems(t *testing.T) {
	var actualItems, expectedItems []*gofeed.Item

	expectedItems = append(expectedItems, &gofeed.Item{GUID: "media:news:20190708154212517", Title: "I think it is a good day", Link: "link01", Description: "so do I!", Published: "2019-07-08 06:42:12"})
	expectedItems = append(expectedItems, &gofeed.Item{GUID: "media:news:20190708154523636", Title: "It's raining outside...", Link: "link02", Description: "I'm getting tired...", Published: "2019-07-09 07:42:12"})

	conn, mock, _ := sqlmock.New()
	defer conn.Close()

	mockRows := sqlmock.NewRows([]string{"guid", "title", "link", "description", "published", "creator", "url", "length", "type"}).
		AddRow("media:news:20190708154212517", "I think it is a good day", "link01", "so do I!", "2019-07-08 06:42:12", nil, nil, nil, nil).
		AddRow("media:news:20190708154523636", "It's raining outside...", "link02", "I'm getting tired...", "2019-07-09 07:42:12", nil, nil, nil, nil)

	mock.ExpectQuery("SELECT").WillReturnRows(mockRows)
	db := rsserver.DB{Connection: conn}
	controller := item.Controller{Table: &db}
	err := controller.GetItems(&actualItems)

	assert.Nil(t, err)
	assert.Equal(t, expectedItems, actualItems)
}

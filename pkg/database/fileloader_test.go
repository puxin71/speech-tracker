package database_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
	"reflect"

	"github.com/puxin71/talk-server/pkg"
	"github.com/puxin71/talk-server/pkg/database"

	"github.com/stretchr/testify/assert"
)

func parseTime(timeStr string) time.Time {
	var t time.Time
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		log.Println(err)
	}
	return t
}

func TestGetJSONFile(t *testing.T) {
	filePath := database.NewFileLoader(pkg.MockConfigProvider{}).Filename
	assert.Equal(t, "dataset.json", filepath.Base(filePath))
	_, err := os.Stat(filePath)
	assert.NoError(t, err)
}

func TestGetAllTalks(t *testing.T) {

		loader := database.NewFileLoader(pkg.MockConfigProvider{})
		talks, err := loader.GetAllTalks()
		assert.NoError(t, err)

		expected := []database.Talk{
			database.Talk{
				Title:    "CILLUM NON",
				Abstract: "Aliqua consequat mollit Lorem dolor nulla qui sunt tempor veniam eiusmod ullamco quis commodo.",
				Room:     873,
				Speaker: database.Attendant{
					Name:       "Hendrix Carroll",
					Company:    "Songlines",
					Email:      "hendrixcarroll@songlines.com",
					Registered: time.Time{},
					Role:       database.SPEAKER,
					Bio:        "Magna velit adipisicing ullamco sint duis nisi.",
				},
				Attendees: []database.Attendant{
					database.Attendant{
						Name:       "Sanders Riley",
						Company:    "Comtext",
						Email:      "sandersriley@comtext.com",
						Registered: parseTime("2015-05-24T02:15:00+07:00"),
						Role:       database.ATTENDEE,
						Bio:        "",
					},
					database.Attendant{
						Name:       "Bean Cline",
						Company:    "Utarian",
						Email:      "beancline@utarian.com",
						Registered: parseTime("2015-06-21T02:54:39+07:00"),
						Role:       database.ATTENDEE,
						Bio:        "",
					},
					database.Attendant{
						Name:       "Alfreda Mitchell",
						Company:    "Dreamia",
						Email:      "alfredamitchell@dreamia.com",
						Registered: parseTime("2015-09-22T06:35:29+07:00"),
						Role:       database.ATTENDEE,
						Bio:        "",
					},
				},
			},
			database.Talk{
				Title:    "AD IPSUM",
				Abstract: "Fugiat nisi et mollit consequat sint.",
				Room:     343,
				Speaker: database.Attendant{
					Name:       "Melody Juarez",
					Company:    "Zillatide",
					Email:      "melodyjuarez@zillatide.com",
					Registered: time.Time{},
					Role:       database.SPEAKER,
					Bio:        "Veniam do eu quis officia enim.",
				},
				Attendees: make([]database.Attendant, 0),
			},
		}

		assert.True(t, reflect.DeepEqual(expected, talks))
}

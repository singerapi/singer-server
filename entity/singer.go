package entity

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

// SingerType for bolt dababase key refers to bucket
var SingerType = "singer"

// Singer implements interface of Entity
type Singer struct {
	ID        int
	Name      string
	SeasonsID []int
	SongsID   []int
	AlbumsID  []int
}

// NewSinger create new Season entity
func NewSinger(id int, na string, seasonsID []int, songsID []int, albumsID []int) Singer {
	return Singer{
		id,
		na,
		seasonsID,
		songsID,
		albumsID,
	}
}

// Type return SingerType
func (s *Singer) Type() string {
	return SingerType
}

// Encode encode
func (s *Singer) Encode() ([]byte, error) {
	return json.Marshal(s)
}

// Decode decode
func (s *Singer) Decode(b []byte) error {
	if err := json.Unmarshal(b, s); err != nil {
		return err
	}
	return nil
}

// HasSong check song id list contain id
func (s *Singer) HasSong(id int) bool {
	return hasID(s.SongsID, id)
}

// HasAlbum check album id list contain id
func (s *Singer) HasAlbum(id int) bool {
	return hasID(s.AlbumsID, id)
}

// HasSeason check season id list contain id
func (s *Singer) HasSeason(id int) bool {
	return hasID(s.SeasonsID, id)
}

// AddSongID check duplicate singer id
func (s *Singer) AddSongID(id int) {
	if !s.HasSong(id) {
		s.SongsID = append(s.SongsID, id)
	}
}

// AddAlbumID check duplicate singer id
func (s *Singer) AddAlbumID(id int) {
	if !s.HasAlbum(id) {
		s.AlbumsID = append(s.AlbumsID, id)
	}
}

// AddSeasonID check duplicate singer id
func (s *Singer) AddSeasonID(id int) {
	if !s.HasSeason(id) {
		s.SeasonsID = append(s.SeasonsID, id)
	}
}

// APIFormatConstruct construct struct with api prefix
func (s *Singer) APIFormatConstruct(HostPrefix string) interface{} {
	var apiItem = struct {
		ID        int      `json:"id"`
		Name      string   `json:"name"`
		AlbumsID  []string `json:"albums"`
		SongsID   []string `json:"songs"`
		SeasonsID []string `json:"seasons"`
		URL       string   `json:"url"`
	}{
		s.ID,
		s.Name,
		[]string{},
		[]string{},
		[]string{},
		HostPrefix + "/" + s.Type() + "s/" + strconv.Itoa(s.ID) + "/",
	}
	for _, ID := range s.AlbumsID {
		apiItem.AlbumsID = append(apiItem.AlbumsID, HostPrefix+"/"+AlbumType+"s/"+strconv.Itoa(ID)+"/")
	}
	for _, ID := range s.SongsID {
		apiItem.SongsID = append(apiItem.SongsID, HostPrefix+"/"+SongType+"s/"+strconv.Itoa(ID)+"/")
	}
	for _, ID := range s.SeasonsID {
		apiItem.SeasonsID = append(apiItem.SeasonsID, HostPrefix+"/"+SeasonType+"s/"+strconv.Itoa(ID)+"/")
	}
	return apiItem
}

// WriteSingerToBucket write entity to bucket with key of entity ID
func WriteSingerToBucket(entityList []Singer, bucket *bolt.Bucket) {
	for i := range entityList {
		data, err := entityList[i].Encode()
		if err != nil {
			log.Fatal(err)
		}
		if err := bucket.Put([]byte(strconv.Itoa(entityList[i].ID)), data); err != nil {
			log.Fatal(err)
		}
	}
}

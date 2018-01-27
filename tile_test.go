package go2048

import (
	"encoding/json"
	"image"
	"testing"
)

func TestTileJSON(t *testing.T) {

	var p = newTile(image.Point{-13, 178}, 25)

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))

	var q = new(Tile)

	err = json.Unmarshal(data, q)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(q)
}

package go2048

import (
	"encoding/json"
)

type StorageManager struct {
	storage Storage
}

const (
	keyBestScore    = "bestScore"
	keyGameState    = "gameState"
	keyNoticeClosed = "noticeClosed"
)

func NewStorageManager(storage Storage) *StorageManager {
	return &StorageManager{storage: storage}
}

func (sm *StorageManager) getValue(key string, v interface{}) error {
	data, err := sm.storage.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (sm *StorageManager) setValue(key string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return sm.storage.Put(key, data)
}

func (sm *StorageManager) setGameState(state *gameState) error {
	return sm.setValue(keyGameState, state)
}

func (sm *StorageManager) getGameState() (*gameState, error) {
	var state gameState
	err := sm.getValue(keyGameState, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (sm *StorageManager) clearGameState() {
	sm.storage.Remove(keyGameState)
}

func (sm *StorageManager) setBestScore(bestScore int) error {
	return sm.setValue(keyBestScore, bestScore)
}

func (sm *StorageManager) getBestScore() int {
	var bestScore int
	err := sm.getValue(keyBestScore, &bestScore)
	if err != nil {
		bestScore = 0
	}
	return bestScore
}

func (sm *StorageManager) getNoticeClosed() bool {
	var noticeClosed bool
	err := sm.getValue(keyNoticeClosed, &noticeClosed)
	if err != nil {
		return false
	}
	return noticeClosed
}

func (sm *StorageManager) setNoticeClosed(noticeClosed bool) error {
	return sm.setValue(keyNoticeClosed, noticeClosed)
}

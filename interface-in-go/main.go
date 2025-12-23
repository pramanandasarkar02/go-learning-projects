package main

import (
	"errors"
	"fmt"
)

type Storage interface {
	Save(data string) error
}


type MemoryStorage struct {
	data []string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make([]string, 0),
	}
}

func (m *MemoryStorage) Save(data string) error {
	if data == "" {
		return errors.New("empty data")
	}
	m.data = append(m.data, data)
	fmt.Println("Saved:", data)
	return nil
}

// Optional helper to inspect stored data
func (m *MemoryStorage) All() []string {
	return m.data
}


type Service struct {
	store Storage
}

func NewService(s Storage) *Service {
	return &Service{store: s}
}

func (s *Service) Process(data string) error {
	if err := s.store.Save(data); err != nil {
		return fmt.Errorf("process failed: %w", err)
	}
	return nil
}


func main() {
	storage := NewMemoryStorage()
	service := NewService(storage)

	inputs := []string{"hello", "golang", ""}

	for _, v := range inputs {
		if err := service.Process(v); err != nil {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println("Stored values:", storage.All())
}

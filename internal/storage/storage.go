package storage

import (
	"encoding/json"
	"os"
	"sync"
)

// Storage, verileri diske JSON formatında kaydetmek için kullanılır.
type Storage struct {
	FilePath string
	mu       sync.Mutex
}

// NewStorage, yeni bir Storage nesnesi oluşturur.
func NewStorage(filePath string) *Storage {
	return &Storage{
		FilePath: filePath,
	}
}

// Save, verilen veriyi dosyaya kaydeder.
func (s *Storage) Save(data interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, bytes, 0644)
}

// Load, dosyadan veriyi yükler.
func (s *Storage) Load(data interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	bytes, err := os.ReadFile(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Dosya yoksa, hata yerine boş veri döndür.
		}
		return err
	}

	return json.Unmarshal(bytes, data)
}

package storage

import (
	"os"
	"reflect"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	// Test verisi ve geçici dosya yolu
	testData := map[string]string{"key1": "value1", "key2": "value2"}
	tempFile, err := os.CreateTemp("", "storage_test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tempFilePath := tempFile.Name()
	tempFile.Close()
	defer os.Remove(tempFilePath) // Test sonunda geçici dosyayı sil

	// Storage nesnesini oluştur
	storage := NewStorage(tempFilePath)

	// Veriyi kaydet
	if err := storage.Save(testData); err != nil {
		t.Errorf("Failed to save data: %v", err)
	}

	// Veriyi yükle
	var loadedData map[string]string
	if err := storage.Load(&loadedData); err != nil {
		t.Errorf("Failed to load data: %v", err)
	}

	// Kaydedilen ve yüklenen verinin aynı olup olmadığını kontrol et
	if !reflect.DeepEqual(testData, loadedData) {
		t.Errorf("Loaded data does not match saved data. Got %v, want %v", loadedData, testData)
	}
}

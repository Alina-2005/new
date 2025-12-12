package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
    "log"
    "os"
)

// Компьютерное оборудование
type HardwareAsset struct {
    AssetID      string
    Category     string // Desktop, Laptop, Server, Network
    SerialNumber string
    Manufacturer string
    Model        string
    PurchaseDate string
    Warranty     WarrantyInfo
    Location     Location
    Specifications map[string]interface{}
}

type WarrantyInfo struct {
    ValidUntil string
    Provider   string
    SupportID  string
}

type Location struct {
    Building string
    Room     string
    Rack     string
    Position string
}

func main() {
    // Регистрируем типы для gob
    gob.Register(map[string]interface{}{})
    
    // Создаем инвентарь
    assets := []HardwareAsset{
        {
            AssetID:      "ASSET-2023-001",
            Category:     "Сервер",
            SerialNumber: "SRV7890123",
            Manufacturer: "Dell",
            Model:        "PowerEdge R750",
            PurchaseDate: "2023-01-15",
            Warranty: WarrantyInfo{
                ValidUntil: "2026-01-15",
                Provider:   "Dell ProSupport",
                SupportID:  "PS-789456",
            },
            Location: Location{
                Building: "DC-1",
                Room:     "Server Room A",
                Rack:     "Rack-42",
                Position: "U15-U20",
            },
            Specifications: map[string]interface{}{
                "cpu":      "2x Intel Xeon Gold 6338",
                "ram_gb":   256,
                "storage_tb": 8,
                "network_ports": 4,
                "power_supply": "Dual 1400W",
            },
        },
        {
            AssetID:      "ASSET-2023-002",
            Category:     "Ноутбук",
            SerialNumber: "LT5678901",
            Manufacturer: "Lenovo",
            Model:        "ThinkPad X1 Carbon Gen 10",
            PurchaseDate: "2023-03-20",
            Warranty: WarrantyInfo{
                ValidUntil: "2025-03-20",
                Provider:   "Lenovo Premier Support",
                SupportID:  "LPS-123456",
            },
            Location: Location{
                Building: "Офисное здание",
                Room:     "3 этаж",
                Rack:     "Шкаф IT-1",
                Position: "Полка 2",
            },
            Specifications: map[string]interface{}{
                "cpu":       "Intel i7-1260P",
                "ram_gb":    32,
                "storage_gb": 1024,
                "display":   "14\" 4K",
                "weight_kg": 1.12,
            },
        },
    }

    // Бинарная сериализация
    var buf bytes.Buffer
    encoder := gob.NewEncoder(&buf)
    
    err := encoder.Encode(assets)
    if err != nil {
        log.Fatal("Ошибка кодирования:", err)
    }

    // Сохраняем в файл
    err = os.WriteFile("hardware_inventory.gob", buf.Bytes(), 0644)
    if err != nil {
        log.Fatal("Ошибка записи файла:", err)
    }
    fmt.Println("✅ Инвентарь сохранен в hardware_inventory.gob")

    // Чтение из файла
    fileData, err := os.ReadFile("hardware_inventory.gob")
    if err != nil {
        log.Fatal("Ошибка чтения файла:", err)
    }

    // Десериализация
    var loadedAssets []HardwareAsset
    decoder := gob.NewDecoder(bytes.NewReader(fileData))
    err = decoder.Decode(&loadedAssets)
    if err != nil {
        log.Fatal("Ошибка декодирования:", err)
    }

    fmt.Println("\n📋 Загруженный инвентарь:")
    for _, asset := range loadedAssets {
        fmt.Printf("ID: %s | %s %s | Локация: %s\n", 
            asset.AssetID, asset.Manufacturer, asset.Model, asset.Location.Room)
    }
}

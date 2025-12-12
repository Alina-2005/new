package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

// Компоненты компьютера
type ComputerSpec struct {
    Model       string    `json:"model"`
    Manufacturer string   `json:"manufacturer"`
    CPU         CPU       `json:"cpu"`
    RAM         RAM       `json:"ram"`
    Storage     []Storage `json:"storage"`
    GPU         *GPU      `json:"gpu,omitempty"`
    Price       float64   `json:"price"`
}

type CPU struct {
    Brand      string  `json:"brand"`
    Model      string  `json:"model"`
    Cores      int     `json:"cores"`
    Threads    int     `json:"threads"`
    BaseClock  float64 `json:"base_clock_ghz"`
    TurboClock float64 `json:"turbo_clock_ghz,omitempty"`
}

type RAM struct {
    CapacityGB int    `json:"capacity_gb"`
    Type       string `json:"type"`
    SpeedMHz   int    `json:"speed_mhz"`
}

type Storage struct {
    Type     string `json:"type"` // SSD, HDD, NVMe
    Capacity int    `json:"capacity_gb"`
    Interface string `json:"interface,omitempty"` // SATA, PCIe
}

type GPU struct {
    Brand    string `json:"brand"`
    Model    string `json:"model"`
    VRAMGB   int    `json:"vram_gb"`
    Chipset  string `json:"chipset,omitempty"`
}

func main() {
    // Создаем конфигурацию компьютера
    gamingPC := ComputerSpec{
        Model:        "Nova Gaming Pro",
        Manufacturer: "Quantum Systems",
        CPU: CPU{
            Brand:      "AMD",
            Model:      "Ryzen 9 7950X",
            Cores:      16,
            Threads:    32,
            BaseClock:  4.5,
            TurboClock: 5.7,
        },
        RAM: RAM{
            CapacityGB: 64,
            Type:       "DDR5",
            SpeedMHz:   6000,
        },
        Storage: []Storage{
            {
                Type:       "NVMe",
                Capacity:   2000,
                Interface:  "PCIe 4.0",
            },
            {
                Type:       "SSD",
                Capacity:   4000,
                Interface:  "SATA",
            },
        },
        GPU: &GPU{
            Brand:   "NVIDIA",
            Model:   "RTX 4090",
            VRAMGB:  24,
            Chipset: "Ada Lovelace",
        },
        Price: 2999.99,
    }

    // Сериализация в JSON
    jsonData, err := json.MarshalIndent(gamingPC, "", "  ")
    if err != nil {
        log.Fatal("Ошибка сериализации:", err)
    }

    // Сохраняем в файл
    err = os.WriteFile("gaming_pc.json", jsonData, 0644)
    if err != nil {
        log.Fatal("Ошибка записи файла:", err)
    }
    fmt.Println("✅ Конфигурация сохранена в gaming_pc.json")

    // Десериализация
    var loadedPC ComputerSpec
    err = json.Unmarshal(jsonData, &loadedPC)
    if err != nil {
        log.Fatal("Ошибка десериализации:", err)
    }
    
    fmt.Printf("\n📊 Загруженная конфигурация:\n")
    fmt.Printf("Модель: %s\n", loadedPC.Model)
    fmt.Printf("Процессор: %s %s (%d ядер)\n", 
        loadedPC.CPU.Brand, loadedPC.CPU.Model, loadedPC.CPU.Cores)
    fmt.Printf("Оперативная память: %d GB %s\n", 
        loadedPC.RAM.CapacityGB, loadedPC.RAM.Type)
    fmt.Printf("Видеокарта: %s %s\n", 
        loadedPC.GPU.Brand, loadedPC.GPU.Model)
    fmt.Printf("Цена: $%.2f\n", loadedPC.Price)
}

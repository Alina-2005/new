package main

import (
    "encoding/xml"
    "fmt"
    "os"
)

type NetworkInfrastructure struct {
    XMLName  xml.Name `xml:"network_infrastructure"`
    Name     string   `xml:"name,attr"`
    Devices  []Device `xml:"device"`
    Topology string   `xml:"topology"`
}

type Device struct {
    XMLName     xml.Name `xml:"device"`
    ID          int      `xml:"id,attr"`
    Type        string   `xml:"type"`
    Hostname    string   `xml:"hostname"`
    IPAddress   string   `xml:"ip_address"`
    MACAddress  string   `xml:"mac_address,omitempty"`
    OS          string   `xml:"operating_system"`
    CPUUsage    float64  `xml:"cpu_usage"`
    MemoryUsage float64  `xml:"memory_usage"`
    Status      string   `xml:"status"` // online, offline, maintenance
}

func main() {
    network := NetworkInfrastructure{
        Name:     "Дата-центр Alpha",
        Topology: "Звезда",
        Devices: []Device{
            {
                ID:          1,
                Type:        "Сервер",
                Hostname:    "web-server-01",
                IPAddress:   "192.168.1.10",
                MACAddress:  "00:1A:2B:3C:4D:5E",
                OS:          "Ubuntu Server 22.04",
                CPUUsage:    23.5,
                MemoryUsage: 67.8,
                Status:      "online",
            },
            {
                ID:          2,
                Type:        "Маршрутизатор",
                Hostname:    "core-router-01",
                IPAddress:   "192.168.1.1",
                OS:          "Cisco IOS",
                CPUUsage:    45.2,
                MemoryUsage: 32.1,
                Status:      "online",
            },
            {
                ID:          3,
                Type:        "Файловый сервер",
                Hostname:    "nas-01",
                IPAddress:   "192.168.1.20",
                MACAddress:  "00:1B:2C:3D:4E:5F",
                OS:          "FreeNAS",
                CPUUsage:    12.3,
                MemoryUsage: 45.6,
                Status:      "maintenance",
            },
        },
    }

    // Сериализация в XML
    xmlData, err := xml.MarshalIndent(network, "", "  ")
    if err != nil {
        panic(err)
    }

    // Добавляем XML заголовок
    xmlWithHeader := []byte(xml.Header + string(xmlData))
    
    // Сохраняем в файл
    err = os.WriteFile("network_config.xml", xmlWithHeader, 0644)
    if err != nil {
        panic(err)
    }
    
    fmt.Println("✅ Конфигурация сети сохранена в network_config.xml")
    fmt.Println("\n📡 Сетевая инфраструктура:")
    for _, device := range network.Devices {
        fmt.Printf("[%s] %s (%s) - %s\n", 
            device.Type, device.Hostname, device.IPAddress, device.Status)
    }
}

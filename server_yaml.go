// go get gopkg.in/yaml.v3
package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "log"
    "os"
)

type ServerConfig struct {
    Server struct {
        Name        string `yaml:"name"`
        Environment string `yaml:"environment"` // production, staging, development
        Host        string `yaml:"host"`
        Port        int    `yaml:"port"`
        SSL         struct {
            Enabled bool   `yaml:"enabled"`
            Cert    string `yaml:"certificate"`
            Key     string `yaml:"key"`
        } `yaml:"ssl"`
    } `yaml:"server"`
    
    Database struct {
        Primary struct {
            Host     string `yaml:"host"`
            Port     int    `yaml:"port"`
            Name     string `yaml:"name"`
            User     string `yaml:"user"`
            Password string `yaml:"password"`
            Pool     struct {
                MaxConnections int `yaml:"max_connections"`
                IdleTimeout    int `yaml:"idle_timeout_seconds"`
            } `yaml:"connection_pool"`
        } `yaml:"primary"`
        
        Replica struct {
            Host string `yaml:"host"`
            Port int    `yaml:"port"`
        } `yaml:"replica,omitempty"`
    } `yaml:"database"`
    
    Monitoring struct {
        Enabled      bool     `yaml:"enabled"`
        MetricsPort  int      `yaml:"metrics_port"`
        HealthChecks []string `yaml:"health_checks"`
        LogLevel     string   `yaml:"log_level"`
    } `yaml:"monitoring"`
    
    Resources struct {
        CPU struct {
            Cores     int `yaml:"cores"`
            Threads   int `yaml:"threads"`
        } `yaml:"cpu"`
        Memory struct {
            TotalGB   int `yaml:"total_gb"`
            JVMHeapGB int `yaml:"jvm_heap_gb,omitempty"`
        } `yaml:"memory"`
        Storage struct {
            Type      string `yaml:"type"`
            Capacity  int    `yaml:"capacity_gb"`
            IOPS      int    `yaml:"iops"`
        } `yaml:"storage"`
    } `yaml:"resources"`
}

func main() {
    config := ServerConfig{}
    
    config.Server.Name = "api-server-prod"
    config.Server.Environment = "production"
    config.Server.Host = "api.example.com"
    config.Server.Port = 443
    config.Server.SSL.Enabled = true
    config.Server.SSL.Cert = "/etc/ssl/certs/api.crt"
    config.Server.SSL.Key = "/etc/ssl/private/api.key"
    
    config.Database.Primary.Host = "db-primary.example.com"
    config.Database.Primary.Port = 5432
    config.Database.Primary.Name = "app_database"
    config.Database.Primary.User = "app_user"
    config.Database.Primary.Password = "secure_password_123"
    config.Database.Primary.Pool.MaxConnections = 100
    config.Database.Primary.Pool.IdleTimeout = 300
    
    config.Database.Replica.Host = "db-replica.example.com"
    config.Database.Replica.Port = 5432
    
    config.Monitoring.Enabled = true
    config.Monitoring.MetricsPort = 9090
    config.Monitoring.HealthChecks = []string{
        "/health",
        "/metrics",
        "/ready",
        "/live",
    }
    config.Monitoring.LogLevel = "info"
    
    config.Resources.CPU.Cores = 8
    config.Resources.CPU.Threads = 16
    config.Resources.Memory.TotalGB = 32
    config.Resources.Memory.JVMHeapGB = 16
    config.Resources.Storage.Type = "NVMe SSD"
    config.Resources.Storage.Capacity = 500
    config.Resources.Storage.IOPS = 50000
    
    // Сериализация в YAML
    yamlData, err := yaml.Marshal(&config)
    if err != nil {
        log.Fatal(err)
    }
    
    // Сохраняем в файл
    err = os.WriteFile("server_config.yaml", yamlData, 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("✅ Конфигурация сервера сохранена в server_config.yaml")
    fmt.Println("\n⚙️ Конфигурация сервера:")
    fmt.Printf("Имя: %s (%s)\n", config.Server.Name, config.Server.Environment)
    fmt.Printf("Хост: %s:%d\n", config.Server.Host, config.Server.Port)
    fmt.Printf("SSL: %v\n", config.Server.SSL.Enabled)
    fmt.Printf("Ресурсы: %d ядер, %d GB RAM\n", 
        config.Resources.CPU.Cores, config.Resources.Memory.TotalGB)
}

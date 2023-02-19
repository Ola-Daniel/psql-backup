package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "os/exec"
    "time"

    "github.com/Azure/azure-sdk-for-go/storage"
    "github.com/lib/pq"
)

func main() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    backupDir := os.Getenv("BACKUP_DIR")
    accountName := os.Getenv("ACCOUNT_NAME")
    accountKey := os.Getenv("ACCOUNT_KEY")
    containerName := os.Getenv("CONTAINER_NAME")
    backupTimeStr := os.Getenv("BACKUP_TIME")
    connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", dbUser, dbPassword, dbHost, dbPort, dbName)

    backupTime, err := time.Parse("15:04", backupTimeStr)
    if err != nil {
        fmt.Println("Error parsing backup time:", err)
        os.Exit(1)
    }

    for {
        now := time.Now()
        if now.After(backupTime) {
            backupFile := fmt.Sprintf("backup_%s.sql", now.Format("20060102_150405"))
            backupPath := filepath.Join(backupDir, backupFile)
            
            db, err := pq.Connect("postgres", connectionString)
            if err != nil {
                fmt.Println("Error connecting to database:", err)
                os.Exit(1)
            }
            defer db.Close()

            cmd := fmt.Sprintf("pg_dump %s > %s", connectionString, backupPath)
            if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
                fmt.Println("Error running pg_dump command:", err)
                os.Exit(1)
            }

            account := storage.NewAccount(
                accountName,
                accountKey,
                true,
            )
            client, err := account.GetBlobService()
            if err != nil {
                fmt.Println("Error creating Blob Storage client:", err)
                os.Exit(1)
            }

            backupBlobName := fmt.Sprintf("%s/%s", containerName, backupFile)
            backupBytes, err := ioutil.ReadFile(backupPath)
            if err != nil {
            -    fmt.Println("Error reading backup file:", err)
                os.Exit(1)
            }

            ctx := context.Background()
            _, err = client.CreateBlockBlobFromReader(ctx, containerName, backupBlobName, int64(len(backupBytes)), bytes.NewReader(backupBytes), nil)
            if err != nil {
                fmt.Println("Error uploading backup to Blob Storage:", err)
                os.Exit(1)
            }

            fmt.Println("Backup complete")
            os.Exit(0)
        }
        time.Sleep(1 * time.Minute)
    }
}
This is a Golang Program that backups an Azure Postgresql Database Server and uploads the backup files to an Azure Blob Storage Container.


Note: Required Parameters are passed to the program through environment variables, these are **required** to successfully run the program. 

Required Parameters
DB_PORT : Port where Database Server is running on
DB_USER : Database Username
DB_PASSWORD : Database Password
DB_NAME : Database Name
BACKUP_DIR : Directory to store backup data
ACCOUNT_NAME : Azure Storage Account Name
ACCOUNT_KEY : Azure Storage Account Key
CONTAINER_NAME : Azure Storage Account Container to store backup data files
BACKUP_TIME : Time to initiate backup of Database Server daily
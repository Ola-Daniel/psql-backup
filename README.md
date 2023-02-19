This is a Golang Program that backups an Azure Postgresql Database Server and uploads the backup files to an Azure Blob Storage Container. <br />


Note: Required Parameters are passed to the program through environment variables, these are **required** to successfully run the program. <br />


Required Parameters <br />
DB_PORT : Port where Database Server is running on <br />
DB_USER : Database Username <br />
DB_PASSWORD : Database Password <br />
DB_NAME : Database Name <br />
BACKUP_DIR : Directory to store backup data <br />
ACCOUNT_NAME : Azure Storage Account Name <br />
ACCOUNT_KEY : Azure Storage Account Key <br />
CONTAINER_NAME : Azure Storage Account Container to store backup data files <br />
BACKUP_TIME : Time to initiate backup of Database Server daily <br />
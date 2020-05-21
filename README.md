# DataApi.Go

## Deploy DataAPi.GO webservice
```
$ docker-compose up --build -d
```

## Change database
Because I don't want to maintain two env files. So if you want to change 
connecting database in each apis, please edit docker-compose file.
```
environment:
      - DB_CONFIG=<database_address>
      - LOG_FILE_NAME=<log_file_name>
```
# go-kafka

A simple go application to illustrate the producer-consumer architecture using kafka


# steps to setup the application

1. Run the below command to get the official docker-compose from bitnami/kafka. 

```
curl -sSL \https://raw.githubusercontent.com/bitnami/containers/main/bitnami/kafka/docker-compose.yml > docker-compose.yml
```

2. Above command will download the yaml that defines Docker service for running Kafka using the Bitnami Kafka Docker image. Replace below line in the yaml 

from
```
KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://:9092
```

to
```
KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
```

this will now run kafka locally on port 9092 on docker-compose up command

3. Run the below command which will start a docker container on port 9092

```
docker-compose up -d
```

4. Now kafka is up and running on ```localhost:9092```, to verify the same run ```lsof -i tcp:9092```

5. Next step is to setup the application and play with it, inorder to do so run

```
go get github.com/md-mudassir7/go-kafka
go mod tidy
go run main.go
```

6. Now both ```kafka[:9092]``` and the ```application[:3000]``` are up and running, test the complete flow with below command and visualise the logs from producer and consumer in the terminal

```
curl --location --request POST '0.0.0.0:3000/api/v1/publish' \
--header 'Content-Type: application/json' \--data-raw '{ "text":"hello go-kafka" }'
```

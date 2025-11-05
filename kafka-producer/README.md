# Kafka Producer Example (Go)

This is a simple Go example that demonstrates how to produce messages to a Kafka cluster.

## Prerequisites

- Go 1.21 or higher
- Kafka cluster running (via docker-compose)
- Kafka UI accessible at http://localhost:12345

## Setup

1. Make sure your Kafka cluster is running:
   ```bash
   docker compose up -d
   ```

2. Install dependencies:
   ```bash
   cd kafka-producer
   go mod download
   ```
   
   **Note:** This example uses [Sarama](https://github.com/IBM/sarama), a pure Go Kafka client library that doesn't require CGO or external C libraries (unlike confluent-kafka-go).

## Running the Producer

```bash
go run main.go
```

The producer will:
- Connect to the Kafka cluster at `localhost:9092,localhost:9094,localhost:9096`
- Create and send 10 messages to the `test-topic` topic
- Display delivery reports for each message

## Creating the Topic (if needed)

Before running the producer, you may want to create the topic first. You can do this via Kafka UI or using the Kafka CLI tools:

```bash
# Using docker exec
docker exec -it kafka-1 kafka-topics.sh --create \
  --bootstrap-server localhost:9092 \
  --topic test-topic \
  --partitions 3 \
  --replication-factor 3
```

Or use Kafka UI at http://localhost:12345 to create the topic.

## Configuration

The producer is configured with:
- **brokers**: All three Kafka broker addresses (localhost:9092, localhost:9094, localhost:9096)
- **RequiredAcks**: WaitForAll - waits for all replicas to acknowledge
- **Retry.Max**: 3 - retries failed messages up to 3 times
- **SyncProducer**: Uses synchronous producer for immediate delivery confirmation

## Verifying Messages

You can verify the messages were produced by:
1. Using Kafka UI at http://localhost:12345
2. Creating a consumer to read from `test-topic`
3. Using Kafka CLI tools to consume messages


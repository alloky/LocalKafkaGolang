# Local Kafka (3-node) with Kafdrop

Spin up a local Apache Kafka cluster (3 brokers) with Zookeeper and Kafdrop for inspection.

## Prerequisites
- Docker Desktop (or Docker Engine) with Docker Compose v2
- Ports used: Zookeeper `2181`, Kafka `9092`, `9093`, `9094`, Kafdrop `9000`

## Start the stack
```bash
# From project root
docker compose up -d
```
- This launches:
  - Zookeeper (`zookeeper:2181`)
  - Kafka brokers: `kafka-0` (`9092`), `kafka-1` (`9093`), `kafka-2` (`9094`)
  - Kafdrop UI on port `9000`

Check container health/logs (optional):
```bash
docker compose ps
docker compose logs -f
```

## Open Kafdrop
1. Wait until brokers are up (see logs/ps above).
2. Open your browser to:
   - http://localhost:9000
3. You should see the cluster and be able to browse topics, partitions, and messages.

Kafdrop is configured to connect to the internal broker endpoints: `kafka-0:29092,kafka-1:29092,kafka-2:29092` as defined in `docker-compose.yml`.

## Stop / Restart / Clean up
```bash
# Stop containers but keep data (if any volumes are defined)
docker compose stop

# Stop and remove containers
# (use this when you want a clean restart next time)
docker compose down

# Recreate containers after changes (rebuild if needed)
docker compose up -d --force-recreate
```

## Notes
- External client ports exposed on localhost:
  - `kafka-0`: `9092`
  - `kafka-1`: `9093`
  - `kafka-2`: `9094`
- The brokers advertise `localhost` for the external listeners so local clients can connect via those ports.
- There is a simple Go producer under `kafka-producer/`. Adjust its broker address to one of the exposed ports above (e.g., `localhost:9092`).

## Kafka producer (Go example)
The `kafka-producer/` folder contains a minimal Sarama-based producer.

Run it against this stack:
```bash
cd kafka-producer
go mod download
go run main.go
```

Defaults in code:
- Brokers: `localhost:9092, localhost:9093, localhost:9094`
- Topic: `test-topic`

If the topic doesn't exist yet, create it (example):
```bash
docker exec -it kafka-0 kafka-topics.sh \
  --bootstrap-server localhost:9092 \
  --create --topic test-topic \
  --partitions 3 --replication-factor 3
```


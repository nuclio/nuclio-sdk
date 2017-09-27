# Kafka Example

Example of using Kafka as event source. The example listens on `topic1` for messages.

You can run kafka locally with the following command:

    docker run \
	-p 2181:2181 \
	-p 9092:9092 \
	--env ADVERTISED_HOST=127.0.0.1 \
	--env ADVERTISED_PORT=9092 \
	spotify/kafka

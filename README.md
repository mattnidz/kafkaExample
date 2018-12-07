# Kafka Example

## example producer and consumer from confluent-kafka-go

1. Start docker registry
`docker run -d -p 5000:5000 --restart=always --name registry registry:2`

1. Build docker images and push to registry
`docker build --no-cache -t localhost:5000/consumer .\consumer`

`docker push localhost:5000/consumer`

`docker build --no-cache -t localhost:5000/producer .\producer`

`docker push localhost:5000/producer `

1. Zookeeper k8s deployment
`kubectl create -f zookeeper-cluster.yaml`

1. Kafka k8s deployment
`kubectl create -f kafka-cluster.yaml`

1. Start k8s kafkaExample
`kubectl create -f deployment.yaml`


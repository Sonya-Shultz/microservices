const { Kafka } = require("kafkajs")
const http = require("http");

const ports_br = ["kafka:9092"]
const userName = 'service1'
const topic = "lab3_messages"

const kafka = new Kafka({ clientId: 'service1', brokers: ports_br})
const consumer = kafka.consumer({groupId: 'service1'})

const consume = async () => {
    await consumer.connect()
    await consumer.subscribe({topic})
    await consumer.run({
        eachMessage: ({data}) => {
            console.log(`we get by service1: ${data?.value}`);
            alert(`we get by service1: ${data?.value}`);
        },
    })
}

const requestListener = function (req, res) {
    if (req.url === '/api/service1'){
        consume();
        res.writeHead(200);
        res.write('Hello frome 1');
    }else{
        res.writeHead(404)
    }
    res.end();
}
const server = http.createServer(requestListener);
server.listen(8080);
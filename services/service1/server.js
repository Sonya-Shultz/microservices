const { Kafka } = require("kafkajs")

const ports_br = ["kafka:9092"]
const userName = "service1"
const topic = "lab3_messages"

const kafka = new Kafka({ userName, ports_br})
const consumer = kafka.consumer({qroupId: userName})

const consume = async () => {
    await consumer.conect()
    await consumer.subscribe({topic})
    await consumer.run({
        eachMessage: ({data}) => {
            console.log(`we get by service1: ${data.value}`);
            alert(`we get by service1: ${data.value}`);
        },
    })
}

consume();
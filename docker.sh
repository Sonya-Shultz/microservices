minikube docker-env
--обери один з рядків 1ий для powershell 2ий - для cmd. Але я рекомендую взагалі ввести це вручну в терміналі (powershell краще) і плтім запускати 5-8 рядки. --
minikube -p minikube docker-env | Invoke-Expression
@FOR /f "tokens=*" %i IN ('minikube -p minikube docker-env') DO @%i
docker build -t client:0.1 -f client/Dockerfile .
docker build -t service1:0.2 -f services/service1/Dockerfile .
docker build -t service2:0.2 -f services/service2/Dockerfile .
docker build -t root-service:0.1 -f services/root-service/Dockerfile .
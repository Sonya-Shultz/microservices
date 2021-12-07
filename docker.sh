minikube docker-env
--обери один з рядків 1ий для powershell 2ий - для cmd. Але я рекомендую взагалі ввести це вручну в терміналі (powershell краще) і плтім запускати 5-8 рядки. --
minikube -p minikube docker-env | Invoke-Expression
@FOR /f "tokens=*" %i IN ('minikube -p minikube docker-env') DO @%i
docker build -t client:0.1 -f client/Dockerfile .
docker build -t service1:0.2 -f services/service1/Dockerfile .
docker build -t service2:0.2 -f services/service2/Dockerfile .
docker build -t root-service:0.1 -f services/root-service/Dockerfile .


1. minikube delete (почистить місце)
2. minikube start --driver=hyperv --cpus 3 --memory 4096 (мінять кількість пам'ять і тд, бо ми слабі)
3. minikube docker-env
4. minikube addons enable ingress
5. docker build -t service1:0.3 -f services/service1/Dockerfile . (одразу 3я версія, бо пропускаєм всі лаби і йдем до кінця)
6. docker build -t service2:0.3 -f services/service2/Dockerfile .
7. docker build -t root-service:0.3 -f services/root-service/Dockerfile .
8. kubectl apply -f k8s_v4 (білд без istio)
9. kubectl apply -f k8s_1/ingress.yaml (білд ingress для зручного рероуту)
10. kubectl apply -f k8s/postgres
11. kubectl apply -f k8s/storage
12. kubectl apply -f k8s/kafka

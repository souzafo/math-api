# Math API - Kubernetes & SRE Challenge

API em Go desenvolvida para operacoes matematicas basicas e validacoes em clusters Kubernetes.

## Requisitos

- Go 1.22+
- Docker
- Kubernetes (k3d, minikube ou similar)
- kubectl

## Execucao dos Testes Unitarios

Para executar os testes e verificar a cobertura de codigo:
```
go test -v -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```
## Build e Deploy Local

1. Gerar a imagem Docker:
```
docker build -t math-api:v1 .
```
2. Importar a imagem para o cluster k3d (se aplicavel):
```
k3d image import math-api:v1 -c math-cluster
```
3. Aplicar os manifestos do Kubernetes:
```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```
4. Verificar o status do deploy:
```
kubectl get pods -l app=math-api
kubectl get svc math-api-service
```
## Validacao dos Endpoints

Como o servico utiliza o tipo ClusterIP, execute um pod temporario para realizar os testes:
```
kubectl run curl-test --rm -i --tty --image=curlimages/curl -- sh
```
Dentro da sessao do container:
```
curl -s http://math-api-service:8000/healthcheck
curl -s "http://math-api-service:8000/api/sum?term_one=10&term_two=5"
curl -s "http://math-api-service:8000/api/sub?term_one=10&term_two=3"
curl -s "http://math-api-service:8000/api/mul?term_one=3&term_two=4"
curl -s "http://math-api-service:8000/api/div?term_one=20&term_two=5"
```
## Atualizacao de Versao (Rollout v2)

Para realizar o update da aplicacao sem downtime:

1. Gerar a nova imagem:
```
docker build -t math-api:v2 .
```
2. Importar para o cluster k3d:
```
k3d image import math-api:v2 -c math-cluster
```
3. Atualizar a imagem no Deployment:
```
kubectl set image deployment/math-api math-api=math-api:v2
```
4. Acompanhar o rollout:
```
kubectl rollout status deployment/math-api
```
## Limpeza

Para remover os recursos criados:
```
kubectl delete -f k8s/service.yaml
kubectl delete -f k8s/deployment.yaml
```
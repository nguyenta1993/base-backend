
### Deployment checklist

|Service Name| Content  | From Infra Teams  |
|--|--|--|
|  Routing| /users   | [ ] |
|Exposed Port| HTTP 5000 /api/{http route} <br> GRPC 5001 /{grpc route} | [ ] |
|Health Check| liveliness :6001/live <br> readiness :6001/ready | [ ] |
|Metrics| :8001/metrics | [ ] |
|Dockerfile| dev: deployment/dev/Dockerfile <br> production: deployment/prod/Dockerfile | [ ] |
|Dependency| [x] Database <br> [ ] Kafka <br> [ ] Redis | [ ] |
|ARG|GH_ACCESS_TOKEN: token to get the private go modules| [x] |
|ENV| OTEL_EXPORTER_OTLP_ENDPOINT : OTLP Endpint.  <br> VAULT_TOKEN: Token to access vault | [x] |
|Note | Some special note for service | [ ] |

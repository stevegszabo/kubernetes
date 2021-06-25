# docker images

```
curl -s http://notebook.local:5000/v2/_catalog | jq -r .
curl -s http://notebook.local:5000/v2/webapp/tags/list | jq -r .

docker build -t notebook.local:5000/webapp:1.0.17 -f Dockerfile.webapp .
docker push notebook.local:5000/webapp:1.0.17
```

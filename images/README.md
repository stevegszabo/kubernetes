# docker images

```
curl -s http://192.168.56.101:5000/v2/_catalog | jq -r .
curl -s http://192.168.56.101:5000/v2/webapp/tags/list | jq -r .

docker build -t 192.168.56.101:5000/webapp:1.0.9 -f Dockerfile.webapp .
docker push 192.168.56.101:5000/webapp:1.0.9
```

# pingdx-server

    buildah bud -f Dockerfile  -t pingdx-server .
    
### run
    podman run \
    --name pingdx-server \
    -p 38080:8080 \
    -v /etc/localtime:/etc/localtime:ro \
    localhost/pingdx-server
 
 ### run -d
    podman run -d \
    --name pingdx-server \
    -p 38080:8080 \
    -v /etc/localtime:/etc/localtime:ro \
    localhost/pingdx-server
# [Docker](https://www.docker.com/)

# Links

* [Get Started, Part 1: Orientation and setup | Docker Documentation](https://docs.docker.com/get-started/)
* [The easy way to set up Docker on a Raspberry Pi – freeCodeCamp.org](https://medium.freecodecamp.org/the-easy-way-to-set-up-docker-on-a-raspberry-pi-7d24ced073ef)

# Install

* centos  

* [Get Docker Engine - Community for CentOS | Docker Documentation](https://docs.docker.com/install/linux/docker-ce/centos/)

```
# 0
sudo yum install -y yum-utils \
  device-mapper-persistent-data \
  lvm2

# 1
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo

# 2
# Install a specific version
sudo yum install docker-ce-18.09.1 docker-ce-cli-18.09.1 containerd.io

# 3
sudo systemctl start docker
sudo systemctl enable docker

# 4
sudo usermod -aG docker x
```

* Raspberry or Ubuntu 18.04

```
curl -sSL https://get.docker.com | sh

sudo usermod -aG docker x
  -G, --groups GROUPS           new list of supplementary GROUPS
  -a, --append                  append the user to the supplemental GROUPS
                                mentioned by the -G option without removing
                                him/her from other groups

Remember that you will have to log out and back in for this to take effect!

docker run hello-world
```

test docker installation

```
docker version
docker info
docker run hello-world
```

# logs

```
docker logs [OPTIONS] CONTAINER
```

# save and load image

* [docker save | Docker Documentation](https://docs.docker.com/engine/reference/commandline/save/)
* [docker load | Docker Documentation](https://docs.docker.com/engine/reference/commandline/load/)

```
// list
docker images

// save
docker save -o busybox.tar busybox

// load
docker load -i busybox.tar
```

## save and load all images

* [How to save all Docker images and copy to another machine - Stack Overflow](https://stackoverflow.com/a/54669431)

```
// save
docker save $(docker images --format '{{.Repository}}:{{.Tag}}') -o allinone.tar

// load
docker load -i allinone.tar
```

# delete all containers and images

```
docker rm $(docker ps -a -q)
docker rmi $(docker images -q)
```

# Basic 

![Container](img/Container@2x.png)

![VM](img/VM@2x.png)


```docker
docker build -t friendlyhello .  # Create image using this directory's Dockerfile
docker run -p 4000:80 friendlyhello  # Run "friendlyname" mapping port 4000 to 80
docker run -d -p 4000:80 friendlyhello         # Same thing, but in detached mode

docker container ls                                # List all running containers
docker container ls -a             # List all containers, even those not running
docker container ls -aq					     # all in quiet mode
docker container stop <hash>           # Gracefully stop the specified container
docker container kill <hash>         # Force shutdown of the specified container
docker container rm <hash>        # Remove specified container from this machine
docker container rm $(docker container ls -a -q)         # Remove all containers

docker image ls -a                             # List all images on this machine
docker image rm <image id>            # Remove specified image from this machine
docker image rm $(docker image ls -a -q)   # Remove all images from this machine

docker login             # Log in this CLI session using your Docker credentials
docker tag <image> username/repository:tag  # Tag <image> for upload to registry
docker push username/repository:tag            # Upload tagged image to registry
docker run username/repository:tag                   # Run image from a registry
```

# Docker Compose

* [Releases · docker/compose · GitHub](https://github.com/docker/compose/releases)

```
curl -L https://github.com/docker/compose/releases/download/1.25.0-rc2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
```


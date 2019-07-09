# [PhotoPrism](https://docs.photoprism.org/en/latest/)

# Links

* [GitHub - photoprism/photoprism: Personal Photo Management powered by Go and Google TensorFlow](https://github.com/photoprism/photoprism)

# Install

Open a terminal and run this command after replacing `~/Pictures` with the
folder containing your photos:

```
docker run -p 2342:80 -d --name photoprism \
  -v ~/Pictures:/srv/photoprism/photos/originals photoprism/photoprism
```

Now open `http://localhost:2342/` in a Web browser to see the user interface.


Index photos

`docker exec -ti photoprism photoprism index`

# Demo

```
docker run -p 2342:80 -d --name demo photoprism/demo
docker rm -f demo
```

# Updating

```
docker pull photoprism/photoprism:latest
```

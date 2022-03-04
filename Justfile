_default:
  @just --list --unsorted


tool := "sonarkube-diff"
version := "latest"
docker_user_repo := "mihaigalos"
docker_image := docker_user_repo + "/" + tool + ":" + version
docker_image_version := docker_image +  version
docker_image_latest :=  docker_image + "latest"

dockerize:
    docker build -t {{docker_image_version}} .
    docker build -t {{docker_image_latest}} .

docker_run url_A url_B:
    docker run --rm -it {{ docker_image }} {{ url_A }} {{ url_B }}

docker_push:
    docker push {{ docker_image_version }}
    docker push {{ docker_image_latest }}

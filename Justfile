_default:
  @just --list --unsorted


tool := "sonarkube-diff"
version := "latest"
docker_user_repo := "mihaigalos"
docker_image := docker_user_repo + "/" + tool + ":" + version

dockerize:
    docker build -t {{docker_image}} .
    docker build -t {{docker_image}} .

docker_run url_A url_B:
    docker run --rm -it {{ docker_image }} {{ url_A }} {{ url_B }}

docker_push:
    docker push {{ docker_image }}

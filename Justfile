_default:
  @just --list --unsorted


tool := "sonarkube-diff"
version := "0.0.1"

dockerize:
    docker build -t {{ tool }}:{{ version }} .

run_docker url_A url_B:
    docker run --rm -it {{ tool }}:{{ version }} {{ url_A }} {{ url_B }}


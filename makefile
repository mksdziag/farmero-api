dev:
	docker compose up
serve:
	docker build -t golang_docker_multibuild_prod .
	docker run -p 5000:5000 -d --name farmero-api-prod golang_docker_multibuild_prod
build : enzod enzo

enzod :
	go build -o ./out/enzod ./app/enzod

enzo :
	go build -o ./out/enzo ./cmd/enzo

.PHONY : container
container :
	docker build -f Dockerfile -t enzod:latest .

.PHONY : container-run
container-run : container
	docker run -d -p 8080:8080 --name enzo enzod:latest

.PHONY : docker-clean
docker-clean : docker-stop
	docker rm enzo

.PHONY : docker-stop
docker-stop :
	docker stop enzo

.PHONY : build-clean
build-clean :
	rm -r ./out

.PHONY : clean
clean : docker-clean build-clean

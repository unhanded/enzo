build : enzod enzo

enzod :
	go build -o ./out/enzod ./app/enzod

enzo :
	go build -o ./out/enzo ./cmd/enzo

.PHONY : container
container :
	podman build -f dockerfile -t enzod:latest

.PHONY : container-run
container-run : container
	podman run -d -p 8080:8080 --name enzo enzod:latest

.PHONY : podman-clean
podman-clean : podman-stop
	podman rm enzo

.PHONY : podman-stop
podman-stop :
	podman stop enzo

.PHONY : build-clean
build-clean :
	rm -r ./out

.PHONY : clean
clean : podman-clean build-clean

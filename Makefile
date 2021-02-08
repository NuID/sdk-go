lang=go
image=nuid/sdk-$(lang)
container=nuid-sdk-$(lang)
work_dir=/nuid/sdk-$(lang)

build:
	docker build -t "$(image):latest" .

clean: stop rm rmi

rm:
	docker rm $(container)

rmi:
	docker rmi $(image)

run:
	docker run -v $$PWD:$(work_dir) -it -d --env-file .env --name $(container) $(image) /bin/sh

shell:
	docker exec -it $(container) /bin/sh

stop:
	docker stop $(container)

test:
	docker exec -it $(container) rake test

.PHONY: test

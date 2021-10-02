IMAGE = imander/secret-share

build: minified
	docker build -t $(IMAGE) .

push:
	docker push $(IMAGE)

minified:
	mkdir minified
	cp -r ./assets minified
	cp -r ./templates minified
	docker run -it --rm -u `id -u`:`id -g` -v $(PWD):/tmp -w /tmp tdewolff/minify minify --recursive ./assets ./templates -o minified

clean:
	rm -rf minified

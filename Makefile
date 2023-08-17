convert:
	cd docs && \
	rm index.html && \
	docker run --rm -i yousan/swagger-yaml-to-html < swagger.yaml > index.html
# If you use this beyond tinkering, make sure you change the access and secret key
.PHONY: dev
dev:
	docker run -d --rm -p 9000:9000 --name minio1 \
		-e "MINIO_ACCESS_KEY=minio" \
		-e "MINIO_SECRET_KEY=minio123" \
		minio/minio:latest server /data


.PHONY: dev-stop
dev-stop:
	docker stop minio1

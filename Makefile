default:
	echo hello

build_app:
	docker build -t urlshort .

run_app:
	docker stop urlshort || true && docker rm urlshort || true
	docker run -it --env-file urlshort.env -p 12321:12321 --link some-redis:redis --name urlshort urlshort #--network=host 

run_redis:
	docker stop some-redis || true && docker container rm some-redis || true
	docker run -p 6379:6379 -v redis_vol:/data --name some-redis -d redis
	# -v /usr/local/etc/redis.conf:/usr/local/etc/redis/redis.conf

create_volume:
	docker volume create --driver local \
      --opt type=none \
      --opt device=/usr/local/var/db/redis \
      --opt o=bind \
      redis_vol


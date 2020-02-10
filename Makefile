default:
	echo hello

prepare_static:
	cp -r static/html/* /var/www/urls.exler.xyz/html/
	cp -r static/css/* /var/www/urls.exler.xyz/css/
	cp -r static/js/* /var/www/urls.exler.xyz/js/
	cp -r static/img/* /var/www/urls.exler.xyz/img/

build_app:
	docker build -t urlshort .

run_app_no_db:
	docker stop urlshort || true && docker rm urlshort || true
	docker run -it -d --env-file urlshort.env -p 127.0.0.1:12321:12321 --name urlshort urlshort

stop:
	docker stop urlshort || true && docker rm urlshort || true
	docker stop some-redis || true && docker rm some-redis || true

run_app:
	docker stop urlshort || true && docker rm urlshort || true
	docker run -it -d --env-file urlshort.env -p 127.0.0.1:12321:12321 --link some-redis:redis --name urlshort urlshort #--network=host 

run_redis:
	docker stop some-redis || true && docker container rm some-redis || true
	docker run -p 127.0.0.1:6379:6379 -v redis_vol:/data --name some-redis -d redis 
	# -v /usr/local/etc/redis.conf:/usr/local/etc/redis/redis.conf

create_volume:
	docker volume create --driver local \
      --opt type=none \
      --opt device=/usr/local/var/db/redis \
      --opt o=bind \
      redis_vol

connect_to_redis:
	sudo docker run -it --name my-redis-cli --net=host --rm redis redis-cli -h 127.0.0.1 -p 6379

create_vol_dirs:
	# mkdir -p /var/log/urlshort
	# mkdir -p /var/log/redis
	mkdir -p /var/redis

user-rpc:
	@make -f deploy/mk/user-rpc.mk release

user-api:
	@make -f deploy/mk/user-api.mk release

social-rpc:
	@make -f deploy/mk/social-rpc.mk release

social-api:
	@make -f deploy/mk/social-api.mk release

im-ws:
	@make -f deploy/mk/im-ws.mk release

im-rpc:
	@make -f deploy/mk/im-rpc.mk release

im-api:
	@make -f deploy/mk/im-api.mk release

task-mq:
	@make -f deploy/mk/task-mq.mk release


release: user-rpc user-api social-api social-rpc im-ws im-rpc im-api task-mq

install-server:
	sed 's/\r//' -i  ./deploy/script/release.sh && cd ./deploy/script && chmod +x release.sh && ./release.sh

install-server-user-rpc:
	cd ./deploy/script && chmod +x user-rpc.sh && ./user-rpc.sh

install-server-user-api:
	cd ./deploy/script && chmod +x user-api.sh && ./user-api.sh

install-docker:
	chmod 777 -R ./components
	docker-compose up -d
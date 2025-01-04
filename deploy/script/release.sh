need_start_server_shell=(
  # rpc
  im-ws.sh
  im-rpc.sh
  user-rpc.sh
  social-rpc.sh

  # api
  im-api.sh
  user-api.sh
  social-api.sh

  # task
  task-mq.sh
)

for i in ${need_start_server_shell[*]} ; do
    chmod +x $i
    sed 's/\r//' -i  $i
    ./$i
done


docker ps
docker exec -it etcd etcdctl get --prefix ""
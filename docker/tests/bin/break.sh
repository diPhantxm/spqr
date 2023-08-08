#!/bin/sh

psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "REGISTER ROUTER r1 ADDRESS spqr_router_1_1:7000;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "REGISTER ROUTER r2 ADDRESS spqr_router_1_2:7000;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "REGISTER ROUTER r3 ADDRESS spqr_router_1_3:7000;"


psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "CREATE SHARDING RULE r1 COLUMN id;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "CREATE SHARDING RULE r2 COLUMN tid;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "CREATE SHARDING RULE r3 COLUMN bid;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "CREATE SHARDING RULE r4 COLUMN aid;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "CREATE KEY RANGE krid1 FROM 1 TO 1000001 ROUTE TO sh1;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "CREATE KEY RANGE krid2 FROM 1000001 TO 2000000 ROUTE TO sh2;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "CREATE KEY RANGE krid3 FROM 2000001 TO 3000000 ROUTE TO sh2;"

psql -h spqr_router_1_1 -p 6432 -U user1 -d db1 -c "CREATE TABLE test(id int, name text, age smallint, balance decimal)"
psql -h spqr_shard_1 -p 6432 -U user1 -d db1 -c "INSERT INTO test (id) VALUES (generate_series(1, 1000000));"
psql -h spqr_shard_2 -p 6432 -U user1 -d db1 -c "INSERT INTO test (id) VALUES (generate_series(1000001, 2000000));"
psql -h spqr_shard_2 -p 6432 -U user1 -d db1 -c "INSERT INTO test (id) VALUES (generate_series(2000001, 3000000));"

while true
do
    id=$(shuf -i 1-3000000 -n 1)
    router=$(($RANDOM % 3 + 1))
    query="SELECT id, name, age, balance FROM test WHERE id=$id;"
    host="spqr_router_1_$router"
    psql -h $host -p 6432 -U user1 -d db1 -c "$query"

    sleep 0.01s
done

#sysbench --threads=24 --table_size=3000000 --auto_inc=false --pgsql-host=spqr_router_1_1 --pgsql-port=6432 --pgsql-user=user1 --pgsql-db=db1 --auto_inc=false --tables=2 --db-driver=pgsql --db-ps-mode=disable /usr/local/bin/oltp_read_write.lua prepare
#sysbench --threads=24 --table_size=3000000 --auto_inc=false --pgsql-host=spqr_router_1_1 --pgsql-port=6432 --pgsql-user=user1 --pgsql-db=db1 --auto_inc=false --tables=2 --db-driver=pgsql --db-ps-mode=disable /usr/local/bin/oltp_read_write.lua run
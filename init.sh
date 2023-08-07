psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "REGISTER ROUTER r1 ADDRESS spqr_router_1_1:7000;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "REGISTER ROUTER r2 ADDRESS spqr_router_1_2:7000;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "REGISTER ROUTER r3 ADDRESS spqr_router_1_3:7000;"


psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "CREATE SHARDING RULE r1 COLUMN id;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "CREATE KEY RANGE krid1 FROM 1 TO 1000000 ROUTE TO sh1;"
psql -h spqr_coordinator -p 7002 -U user1 -d db1 -c "CREATE KEY RANGE krid2 FROM 1000000 TO 2000000 ROUTE TO sh2;"

psql -h spqr_router_1_1 -p 6432 -U user1 -d db1 -c "CREATE TABLE test(id int, name text, age smallint, balance decimal)"
psql -h spqr_router_1_1 -p 6432 -U user1 -d db1 -c "INSERT INTO test (id) VALUES (generate_series(1, 500000));"
psql -h spqr_router_1_2 -p 6432 -U user1 -d db1 -c "INSERT INTO test (id) VALUES (generate_series(500001, 1000000));"
psql -h spqr_router_1_3 -p 6432 -U user1 -d db1 -c "INSERT INTO test (id) VALUES (generate_series(1000001, 2000000));"
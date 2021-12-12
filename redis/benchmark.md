##### 1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
```
# 1 get 10
redis-benchmark -n 100000 -t get -d 10
====== GET ======
  100000 requests completed in 0.76 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
131752.31 requests per second

# 2 get 20
redis-benchmark -n 100000 -t get -d 20
====== GET ======
  100000 requests completed in 0.74 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
134952.77 requests per second

# 3 get 50
redis-benchmark -n 100000 -t get -d 50
====== GET ======
  100000 requests completed in 0.74 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
135135.14 requests per second

# 4 get 100
redis-benchmark -n 100000 -t get -d 100
====== GET ======
  100000 requests completed in 0.75 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

99.93% <= 1 milliseconds
100.00% <= 1 milliseconds
133689.83 requests per second

# 5 get 200
edis-benchmark -n 100000 -t get -d 200
====== GET ======
  100000 requests completed in 0.75 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

99.88% <= 1 milliseconds
100.00% <= 1 milliseconds
133868.81 requests per second

# 6 get 500
redis-benchmark -n 100000 -t get -d 500
====== GET ======
  100000 requests completed in 0.75 seconds
  50 parallel clients
  500 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
133333.33 requests per second

# 7 get 1k
redis-benchmark -n 100000 -t get -d 1000
====== GET ======
  100000 requests completed in 0.80 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

99.62% <= 1 milliseconds
99.82% <= 2 milliseconds
100.00% <= 2 milliseconds
125313.29 requests per second

# 8 get 5k
redis-benchmark -n 100000 -t get -d 5000
====== GET ======
  100000 requests completed in 0.74 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

99.77% <= 1 milliseconds
100.00% <= 1 milliseconds
134408.59 requests per second

# 9 set 10
redis-benchmark -n 100000 -t set -d 10
====== SET ======
  100000 requests completed in 0.80 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

99.79% <= 1 milliseconds
100.00% <= 1 milliseconds
125628.14 requests per second

# 10 set 20
redis-benchmark -n 100000 -t set -d 20
====== SET ======
  100000 requests completed in 0.78 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

99.68% <= 1 milliseconds
100.00% <= 2 milliseconds
128205.13 requests per second

# 11 set 50
redis-benchmark -n 100000 -t set -d 50
====== SET ======
  100000 requests completed in 0.79 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

99.71% <= 1 milliseconds
99.96% <= 2 milliseconds
100.00% <= 2 milliseconds
126422.25 requests per second

# 12 set 100
redis-benchmark -n 100000 -t set -d 100
====== SET ======
  100000 requests completed in 0.81 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

99.67% <= 1 milliseconds
99.92% <= 2 milliseconds
99.95% <= 3 milliseconds
100.00% <= 4 milliseconds
100.00% <= 4 milliseconds
122850.12 requests per second

# 13 set 200
redis-benchmark -n 100000 -t set -d 200
====== SET ======
  100000 requests completed in 0.80 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

99.49% <= 1 milliseconds
99.85% <= 2 milliseconds
99.90% <= 3 milliseconds
99.96% <= 4 milliseconds
100.00% <= 4 milliseconds
124533.01 requests per second

# 14 set 500
redis-benchmark -n 100000 -t set -d 500
====== SET ======
  100000 requests completed in 0.86 seconds
  50 parallel clients
  500 bytes payload
  keep alive: 1

99.12% <= 1 milliseconds
99.84% <= 2 milliseconds
100.00% <= 3 milliseconds
100.00% <= 3 milliseconds
116822.43 requests per second

# 15 set 1k
redis-benchmark -n 100000 -t set -d 1000
====== SET ======
  100000 requests completed in 0.80 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

99.60% <= 1 milliseconds
99.95% <= 2 milliseconds
100.00% <= 2 milliseconds
124688.28 requests per second

# 16 set 5k
redis-benchmark -n 100000 -t set -d 5000
====== SET ======
  100000 requests completed in 0.86 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

99.71% <= 1 milliseconds
99.95% <= 41 milliseconds
100.00% <= 41 milliseconds
115874.86 requests per second

```
- creating records 1 by 1 in memory is about twice as fast as by file
- however, create in batches is same speed on both

benchmarks - 1000 logs
Single: 44.88ms
Batches: 7.9ms

single log: 480us

I think latency between http servers is bigger than 480us, so db is not the limiting factor here

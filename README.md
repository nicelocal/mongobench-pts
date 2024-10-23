# MongoBench-pts
Parallelized MongoDB benchmarking suite for the Phoronix Test Suite based on the official Golang driver.

Manual install (as non-root):

```
git clone https://github.com/nicelocal/mongobench-pts

mkdir -p ~/.phoronix-test-suite/test-suites/pts
mkdir -p ~/.phoronix-test-suite/installed-tests/pts/
mkdir -p ~/.phoronix-test-suite/test-profiles/pts/

cp -a mongobench-pts ~/.phoronix-test-suite/test-suites/pts/mongobench
cp -a mongobench-pts ~/.phoronix-test-suite/installed-tests/pts/
cp -a mongobench-pts ~/.phoronix-test-suite/test-profiles/pts/
```

Manual install (as root):

```
git clone https://github.com/nicelocal/mongobench-pts

mkdir -p /var/lib/phoronix-test-suite/test-suites/pts
mkdir -p /var/lib/phoronix-test-suite/installed-tests/pts/
mkdir -p /var/lib/phoronix-test-suite/test-profiles/pts/

cp -a mongobench-pts /var/lib/phoronix-test-suite/test-suites/pts/mongobench
cp -a mongobench-pts /var/lib/phoronix-test-suite/installed-tests/pts/
cp -a mongobench-pts /var/lib/phoronix-test-suite/test-profiles/pts/
```
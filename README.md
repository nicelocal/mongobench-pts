# MongoBench-pts
Parallelized MongoDB benchmarking suite for the Phoronix Test Suite based on the official Golang driver.

Manual install (as non-root):

```
git clone https://github.com/nicelocal/mongobench-pts

cd mongobench-pts

mkdir -p ~/.phoronix-test-suite/test-suites/pts
mkdir -p ~/.phoronix-test-suite/installed-tests/pts/
mkdir -p ~/.phoronix-test-suite/test-profiles/pts/

cp -a . ~/.phoronix-test-suite/test-suites/pts/mongobench
cp -a . ~/.phoronix-test-suite/installed-tests/pts/mongobench
cp -a . ~/.phoronix-test-suite/test-profiles/pts/mongobench
```

Manual install (as root):

```
git clone https://github.com/nicelocal/mongobench-pts

cd mongobench-pts

mkdir -p /var/lib/phoronix-test-suite/test-suites/pts
mkdir -p /var/lib/phoronix-test-suite/installed-tests/pts/
mkdir -p /var/lib/phoronix-test-suite/test-profiles/pts/

cp -a . /var/lib/phoronix-test-suite/test-suites/pts/mongobench
cp -a . /var/lib/phoronix-test-suite/installed-tests/pts/mongobench
cp -a . /var/lib/phoronix-test-suite/test-profiles/pts/mongobench
```
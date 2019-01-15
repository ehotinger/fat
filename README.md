# fat

This is a demo repository that can be used to run a simple Go binary which reads a large file and discards its content.

For creating the file:

```bash
dd if=/dev/urandom of=file.txt bs=64M count=64 iflag=fullblock
```

Build with Docker:

```bash
docker build -t fat .
```

Run:

```
docker run --rm fat -read -file /usr/bin/file.txt
```

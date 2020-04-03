According to this [wiki](https://en.wikipedia.org/wiki/List_of_file_signatures), jpeg files are begin with these following bytes
x = `[FF D8 FF E0 00 10 4A 46 49 46 00 01]`
Using `hexdump -C my_magic_bytes.jpg.enc` on encrypted file show that encrypted file begin with
y = `[b9 14 06 45 71 e0 b5 f7 37 07 cb 85]`
Inverse operator of XOR is also XOR
=> key = x ^ y = `[46 cc f9 a5 71 f0 ff b1 7e 41 cb 84]`
using the above key to decrypt the file resulted in an image with flag

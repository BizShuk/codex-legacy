# Cryptographic

### Message-digest
驗證訊息完整性, output a fixed length hashtag( or called message digest) 且不可逆轉 irreversible
- MD5
- SHA1, SHA2, SHA3, hard to break through brute-force
- HMAC
是金鑰相關的雜湊運算訊息認證碼（Hash-based Message Authentication Code）,HMAC運算利用雜湊演算法，以一個金鑰和一個訊息為輸入，生成一個訊息摘要作為輸出。也就是說HMAC是需要一個金鑰的。所以，HMAC_SHA1也是需要一個金鑰的，而SHA1不需要。


### Symmetric encryption
雙方都知道加密與解密的金鑰
- AES（Advanced Encryption Standard), high-efficiency
- DES（Data Encryption Standard)
- 3DES（Triple DES, based on DES and encrypt triple times.


### Asymmetric encryption, Public-key crypography
一方用公開金鑰加密後,傳給另一方用私密金鑰解密
- RSA, ssh
- DSA（Digital Signature Algorithm）, 數字簽名演算法，是一種標準的 DSS（數字簽名標準），嚴格來說不算加密演算法。
- ECC（Elliptic Curves Cryptography）
橢圓曲線密碼編碼學。ECC和RSA相比，具有多方面的絕對優勢，主要有：抗攻擊性強。相同的金鑰長度，其抗攻擊性要強很多倍。計算量小，處理速度快。ECC總的速度比RSA、DSA要快得多。儲存空間佔用小。ECC的金鑰尺寸和系統引數與RSA、DSA相比要小得多，意味著它所佔的存貯空間要小得多。這對於加密演算法在IC卡上的應用具有特別重要的意義。頻寬要求低。當對長訊息進行加解密時，三類密碼系統有相同的頻寬要求，但應用於短訊息時ECC頻寬要求卻低得多。頻寬要求低使ECC在無線網路領域具有廣泛的應用前景。


### base64
轉換後的字串理論上將要比原來的長1/3
編碼後的資料是一個字串，其中包含的字元為：A-Z、a-z、0-9、+、/，共64個字元(26 + 26 + 10 + 1 + 1 = 64，其實是65個字元，“=”是填充字元。Base64要求把每三個8Bit的位元組轉換為四個6Bit的位元組(3*8 = 4*6 = 24)，然後把6Bit再添兩位高位0，組成四個8Bit的位元組

### SSL
for https(Hypertext Transfer Protocol over Secure Socket Layer)
SSL使用40 位關鍵字作為RC4流加密演算法

[Online hash function](http://www.fileformat.info/tool/hash.htm)
- hash
- md5
- crc32
- base64
- sha1
- sha2
- des3
- aes128
- aes256
- bf
- rc
- rsa
- x509
- 


- [ ] [A Few Thoughts on Cryptographic Engineering](http://blog.cryptographyengineering.com/2011/11/how-not-to-use-symmetric-encryption.html)
- [_] [An Internet Encyclopedia](http://www.freesoft.org/CIE/index.htm)
- [] [Cryptographic_hash_function](https://en.wikipedia.org/wiki/Cryptographic_hash_function)

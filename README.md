# Blockchain-go

## [BlockChain Tutorial](https://www.youtube.com/watch?v=mYlHT9bB6OE&list=PLpP5MQvVi4PGmNYGEsShrlvuE2B33xV1L) 
<br><br>
## Part 1 
 > #### [Definition](https://blockgeeks.com/guides/what-is-blockchain-technology/)
1. #### `go mod init github.com/Guo-lab/Blockchain-go` for compiling (pretty handy feature)
2. #### What is in Block:
   - The data we want to pass around inside our database.
   - The hash related to the Block itself.

Output: 
   ```
   ###################################
   Last Block Hash    : 
   Block Data         : Block Head
   Current Block Hash : 164047fe99c4a1f534ef73528533b11118de7bd77897751cf638dbc72824002c
   -----------------------------------
   ###################################
   Last Block Hash    : 164047fe99c4a1f534ef73528533b11118de7bd77897751cf638dbc72824002c
   Block Data         : First Block after Genesis
   Current Block Hash : 87167d6fc72aad7296baf65c0d430742c2e9a8cd371064eb372962635ab31680
   -----------------------------------
   ###################################
   Last Block Hash    : 87167d6fc72aad7296baf65c0d430742c2e9a8cd371064eb372962635ab31680
   Block Data         : Second Block after First
   Current Block Hash : b7d74545d0187d5ca92582021371ace01ede34f5821fca4cd98f208e299944b5
   -----------------------------------
   ###################################
   Last Block Hash    : b7d74545d0187d5ca92582021371ace01ede34f5821fca4cd98f208e299944b5
   Block Data         : Third Block after Second
   Current Block Hash : 18cb77f41cb3853014012c150ea262e52fd806cc45a578ffd964363b261a80a0
   -----------------------------------
   ```

<br><br>
## Part 2
> import: use the absolute path considering GOPATH
1. #### Refactoring
2. #### [Proof of Work](https://crypto-academy.org/what-is-proof-of-work/)   
   Simply speaking, find right sha256 corresponding to the combination of data and nonce 


<br><br>
## Part 2
1. #### Add a key value storage database BadgerDB (accepting only bytes or slices of bytes)
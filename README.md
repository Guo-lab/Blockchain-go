# Blockchain-go

## [BlockChain Tutorial](https://www.youtube.com/watch?v=mYlHT9bB6OE&list=PLpP5MQvVi4PGmNYGEsShrlvuE2B33xV1L) 
<br><br>
## Part 1 
 > ### [Definition](https://blockgeeks.com/guides/what-is-blockchain-technology/)
1. ### `go mod init github.com/Guo-lab/Blockchain-go` for compiling (pretty handy feature)
2. ### What is in Block:
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
1. ### Refactoring
2. ### [Proof of Work](https://crypto-academy.org/what-is-proof-of-work/)   
   Simply speaking, find right sha256 corresponding to the combination of data and nonce 


<br><br>
## Part 3
1. ### Add a key value storage database BadgerDB (accepting only bytes or slices of bytes)
2. ### Fetch or Store data in the Database
3. ### CommandLine > BlockChain > Block
4. ### Design Command Line (3 functions)
   - print help, 
   - add blocks, 
   - print the chain
5. ### Database Path, same level with main.go 
   ```
   badger 2022/11/23 02:16:55 INFO: All 1 tables opened in 0s
   badger 2022/11/23 02:16:55 INFO: Replaying file id: 0 at offset: 949
   badger 2022/11/23 02:16:55 INFO: Replay took: 7.313µs
   badger 2022/11/23 02:16:55 DEBUG: Value log discard stats empty
   00003b00bdb0d5d5f92f0ba4c937d4b6b482a4278281c1b35892e951d7e86da0
   Block added successfully.
   -------------- Block Information ----------
   Last Block Hash    : 00000869ccf72435b92f7e9e1adea7ff01366a17df92240f3fca4ac97a09ce17
   Block Data         : Send b 10BTS to a
   Current Block Hash : 00003b00bdb0d5d5f92f0ba4c937d4b6b482a4278281c1b35892e951d7e86da0
   Proof of Work      : true
   -------------------------------------------
   -------------- Block Information ----------
   Last Block Hash    : 000004d81ebadcc778a71cd908f5a6380215607b9d8efea5821e6179ed320c34
   Block Data         : Send a 10BTS to b
   Current Block Hash : 00000869ccf72435b92f7e9e1adea7ff01366a17df92240f3fca4ac97a09ce17
   Proof of Work      : true
   -------------------------------------------
   -------------- Block Information ----------
   Last Block Hash    : 000024f1345fbec1899505809d1d681698365b75e9742b0bf0d00606f0514b7a
   Block Data         : OK
   Current Block Hash : 000004d81ebadcc778a71cd908f5a6380215607b9d8efea5821e6179ed320c34
   Proof of Work      : true
   -------------------------------------------
   -------------- Block Information ----------
   Last Block Hash    : 
   Block Data         : Block Head
   Current Block Hash : 000024f1345fbec1899505809d1d681698365b75e9742b0bf0d00606f0514b7a
   Proof of Work      : true
   -------------------------------------------
   badger 2022/11/23 02:16:59 DEBUG: Storing value log head: {Fid:0 Len:42 Offset:1242}
   badger 2022/11/23 02:16:59 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}
   badger 2022/11/23 02:16:59 INFO: Running for level: 0
   badger 2022/11/23 02:16:59 DEBUG: LOG Compact. Added 7 keys. Skipped 3 keys. Iteration took: 432.363µs
   badger 2022/11/23 02:16:59 DEBUG: Discard stats: map[0:140]
   badger 2022/11/23 02:16:59 INFO: LOG Compact 0->1, del 2 tables, add 1 tables, took 42.100966ms
   badger 2022/11/23 02:16:59 INFO: Compaction for level: 0 DONE
   badger 2022/11/23 02:16:59 INFO: Force compaction on level 0 done
   ```

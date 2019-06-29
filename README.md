# ipfs-different-hash-same-data

# Get `dhsd`
```
git clone https://github.com/drbh/ipfs-different-hash-same-data.git
go build -o dhsd
```

# Same file different CIDs

**note: make sure you have a local go-ipfs running with a gateway at `:5001`**

Add with the default settings
```bash
ipfs add ./000_cutecat.png 
# added QmPEFkUZtjnRJPGC3k6Mz7a7rJ2KJ4HrHo4wP5dqGEvbYU 000_cutecat.png
# 1.05 MiB / 1.05 MiB [===================================] 100.00%(base)
 ```

Add it again but with a trickle DAG schemea (or change any other hashing param)
```bash
ipfs add --trickle ./000_cutecat.png 
# added QmXNQkk81cXFSD3hs4etEZrfEPLbC6N1MvbHaoNLAQ2iHr 000_cutecat.png
# 1.05 MiB / 1.05 MiB [===================================] 100.00%
```

We know that `QmPEFkUZtjnRJPGC3k6Mz7a7rJ2KJ4HrHo4wP5dqGEvbYU` and `QmXNQkk81cXFSD3hs4etEZrfEPLbC6N1MvbHaoNLAQ2iHr` both contain the same data but we can see that by changing the hashing config we get different CIDs

This is a side effect of the powerful IPFS CID concept but can cause unnecessary duplication of idenitcal files. 

To combat this unwanted redudency we resolve the files bytes, compare them and then store the similar files in an embedded SQLITE3 db. 

## The concept

1. Resolve CIDs
2. Compare CIDs
3. Add to datastore if new
4. Return all similar CIDs on request

# CLI interface

#### compare
```bash
./dhsd compare QmPEFkUZtjnRJPGC3k6Mz7a7rJ2KJ4HrHo4wP5dqGEvbYU QmXNQkk81cXFSD3hs4etEZrfEPLbC6N1MvbHaoNLAQ2iHr

# +1 CIDs return equal byte arrays, DB updated
```

#### eq
```bash
./dhsd eq QmPEFkUZtjnRJPGC3k6Mz7a7rJ2KJ4HrHo4wP5dqGEvbYU

# QmXNQkk81cXFSD3hs4etEZrfEPLbC6N1MvbHaoNLAQ2iHr
```

the mapping is bi-directional so...
```bash
./dhsd eq QmXNQkk81cXFSD3hs4etEZrfEPLbC6N1MvbHaoNLAQ2iHr

# QmPEFkUZtjnRJPGC3k6Mz7a7rJ2KJ4HrHo4wP5dqGEvbYU
```

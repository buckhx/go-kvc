# go-kvc

Simple Go Key Value Cache that is useful for mocking more complex stores.

Short TTLs (< 1 second) don't work


## Installation

```
go get github.com/buckhx/go-kvc
```

## Usage

```
ages := kvc.NewMem()
ages.Set("me", 27)
ages.Set("ash", 10)
ages.Set("oak", 47)

// Set key that expires in 1 second
ages.SetTTL("butterfree", 15, 1*time.second)
butterAge := ages.Get("butterfree")
time.Sleep(1*time.Second) // pause long enough for key to expire
exp := ages.Get("butterfree")
fmt.Println("Butterfree age: %d, Expired: %t", butterAge, exp)

// Only set Delia if Oak in cache
ages.CompareAndSet("delia", 29, func() bool {
	return ages.(*kvc.MemKVC).UnsafeHas("oak")
}

// Atomic increment
ages.GetAndSet("me", func(cur kvc.Value) kvc.Value) {
	return cur.(int) + 1
})
```

## TODO

- [ ] godoc badge
- [ ] thorough testing 
- [ ] fix short TTL 

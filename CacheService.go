package main




type CasheService interface {
	Post(x int, y int,z int,res float64)
	Get(x int, y int,z int) (res float64)
	Put(x int, y int,z int)
	Delete(x int, y int,z int)
}

func (ss *cCache) Post(x int, y int, z int,res float64) {

	ss.cashes[key{x, y, z}] = res
}

func (ss *cCache) Get(x int, y int, z int) (res float64) {
	res, exists := ss.cashes[key{x, y, z}]
	if !exists {
		return -1
	}
	return
}

func (ss *cCache) Delete(x int, y int, z int) {
	delete(ss.cashes, key{x, y, z})
}

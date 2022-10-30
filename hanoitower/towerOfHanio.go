package hanoitower

import "fmt"

func TowerOfHanio(nDisk int, src, dest, temp string) {
	if nDisk == 1 {
		fmt.Printf("Disk move from %s to %s\n", src, dest)
		return
	}

	TowerOfHanio(nDisk-1, src, temp, dest)
	fmt.Printf("Disk move from %s to %s\n", src, dest)

	TowerOfHanio(nDisk-1, dest, temp, src)
	// fmt.Printf("Disk move from %s to %s\n", dest, temp)

}
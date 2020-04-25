package cache

var imageCache = make(map[string]([]byte))

func CachingImage(name string, image []byte) {
	imageCache[name] = image
}

func CheckCachedImage(name string) bool {
	_, ok := imageCache[name]
	return ok
}

func GetCachedImage(name string) []byte {
	return imageCache[name]
}
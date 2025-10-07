package genrator

import "math/rand"

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func base62Encode(n int) string {
	if n == 0 {
		return string(charset[0])
	}
	res := ""
	for n > 0 {
		res = string(charset[n%62]) + res
		n /= 62
	}
	return res
}

func GenerateIdentifier(id int) string {
	encoded := base62Encode(id)
	unique := make(map[rune]bool)

	// Đảm bảo mỗi ký tự khác nhau
	for _, c := range encoded {
		unique[c] = true
	}

	// Thêm ký tự khác để đủ 8
	for len(unique) < 8 {
		r := rune(charset[rand.Intn(len(charset))])
		if !unique[r] {
			unique[r] = true
		}
	}

	// Chuyển sang slice để shuffle
	result := make([]rune, 0, 8)
	for k := range unique {
		result = append(result, k)
	}

	// Trộn để random hoá
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}

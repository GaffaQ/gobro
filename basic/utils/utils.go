package utils

//function yang awalnya huruf besar berarti bisa diakses diluar packagenya
//function yang awalnya huruf kecil berarti gabisa diakses diluar package walaupun sudah di import

func apaKabar() string {
	return "Baiik"
}

func ApaKabar() string {
	return "Baikkk"
}

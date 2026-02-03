package moviesapi

type Movies struct {
	ID int `json:id`;
	Title string `json:title`;
	Year int `json:year`;
	Rating float64 `json:rating`;
    Director *Director `json:director`;
}

type Director struct {
	Name string `json:name`;
}	

func main (){

}
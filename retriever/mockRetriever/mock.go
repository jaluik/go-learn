package mock

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return "retriever my to string"
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) NewGet(url string) string {
	return r.Contents
}

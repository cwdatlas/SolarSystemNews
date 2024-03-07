package main

/*
@Author Aidan Scott
Stores article data
*/
var localArticles = map[string][2]string{
	"updates":  {"Belt sector AP-089 daily updates", "Body"},
	"comets":   {"Mining asteroids? Try comets instead!", "Body"},
	"carriers": {"WARNING: local carriers will be going on break next solar week", "Body"},
}

// returns map of articles
func getLocalNews() map[string][2]string {
	return localArticles
}

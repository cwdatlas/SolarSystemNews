package main

/*
@Author Aidan Scott
Stores article data
*/
var systemArticles = map[string][2]string{
	"venus": {"New Life on Venus",
		"From: Venus immigration agency, 10/8/2224 Solar Standard Time\n" +
			"Come to lovely Venus with is great sky cities like Vaccos, and Tenebra!\n" +
			"50 kilometers over the surface of Venus is the best place to live in the System!\n" +
			"No better place to raise a family than the sunny, warm days in the Vaccos suburbs!\n" +
			"Come today and reserve a high in demand home! People come every day, space is running out!\n" +
			"Come today!",
	},
	"mirror": {"New Mirror Project in Ganymede orbit",
		"From: Juvian Progress and Peoples Council, 10/3/2224 Solar Standard Time.\n" +
			"A new mega project is being built in orbit of IO. This project is meant alter the local day length.\n" +
			"The mirror is also meant to warm the surface to a more barable tempurature in preparation for larger \n" +
			"terraforming projects. Our largest moon will soon be provided with adequate power to run large refineries \n" +
			"making Ganymede the core of the "},
	"work": {"Work on the Moon; What you need to know", "Body"}}

// returns map of articles
func getSystemNews() map[string][2]string {
	return systemArticles
}

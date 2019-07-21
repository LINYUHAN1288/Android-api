package models

type ChineseDietNutriFromBH struct {
	GroupID         string  `orm:"column(GroupID);size(30)"`
	DietID          string  `orm:"column(DietID);size(120)"`
	DietChineseName string  `orm:"column(DietChineseName);size(255);null"`
	DietName        string  `orm:"column(DietName);size(255);null"`
	DietType        string  `orm:"column(DietType);size(255);null"`
	DietWeight      string  `orm:"column(DietWeight);size(30);null"`
	Energy          float64 `orm:"column(Energy);null;digits(10);decimals(0)"`
	Protein         float64 `orm:"column(Protein);null;digits(10);decimals(0)"`
	Carbohydrate    float64 `orm:"column(Carbohydrate);null;digits(10);decimals(0)"`
	Fat             float64 `orm:"column(Fat);null;digits(10);decimals(0)"`
	Fiber           float64 `orm:"column(Fiber);null;digits(10);decimals(0)"`
	Sugar           float64 `orm:"column(Sugar);null;digits(10);decimals(0)"`
	VitaminC        float64 `orm:"column(VitaminC);null;digits(10);decimals(0)"`
	DietSteps       string  `orm:"column(DietSteps);size(8000);null"`
	Meterial        string  `orm:"column(Meterial);size(1024);null"`
	ImageSrc        string  `orm:"column(ImageSrc);size(255);null"`
	Source          int     `orm:"column(Source);null"`
	VitaminA        float64 `orm:"column(VitaminA);null;digits(10);decimals(0)"`
	VitaminE        float64 `orm:"column(VitaminE);null;digits(10);decimals(0)"`
	VitaminB1       float64 `orm:"column(VitaminB1);null;digits(10);decimals(0)"`
	VitaminB2       float64 `orm:"column(VitaminB2);null;digits(10);decimals(0)"`
	VitaminB3       float64 `orm:"column(VitaminB3);null;digits(10);decimals(0)"`
	Cholesterol     float64 `orm:"column(Cholesterol);null;digits(10);decimals(0)"`
	MagnesiumMg     float64 `orm:"column(MagnesiumMg);null;digits(10);decimals(0)"`
	CalciumCa       float64 `orm:"column(CalciumCa);null;digits(10);decimals(0)"`
	IronFe          float64 `orm:"column(IronFe);null;digits(10);decimals(0)"`
	ZincZn          float64 `orm:"column(ZincZn);null;digits(10);decimals(0)"`
	CopperCu        float64 `orm:"column(CopperCu);null;digits(10);decimals(0)"`
	ManganeseMn     float64 `orm:"column(ManganeseMn);null;digits(10);decimals(0)"`
	KaliumK         float64 `orm:"column(KaliumK);null;digits(10);decimals(0)"`
	PhosphorP       float64 `orm:"column(PhosphorP);null;digits(10);decimals(0)"`
	SodiumNa        float64 `orm:"column(SodiumNa);null;digits(10);decimals(0)"`
	SeleniumSe      float64 `orm:"column(SeleniumSe);null;digits(10);decimals(0)"`
	ReadingNum      int     `orm:"column(ReadingNum);null"`
}

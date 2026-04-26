package data

import "backend/models"

type ClassifiedData interface {
	GetAll() []models.Classified
}

type classifiedData struct {
	items []models.Classified
}

func InitClassifiedData() ClassifiedData {
	return &classifiedData{
		items: []models.Classified{
			{ID: 1, ClassTitle: "Villa for Rent near Phsar Deum Tkov 6 Bedroom= 1200USD/month", CategoryField: "HOUSES/FLATS - For Sale"},
			{ID: 2, ClassTitle: "Nice Villa for rent near Russain Market,4br,6bh,$2000/m", CategoryField: "GENERAL - For Sale"},
			{ID: 3, ClassTitle: "chbar ampov house for sale- urgent", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 4, ClassTitle: "1bedroom available doun penh", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 5, ClassTitle: "4bed =$600 1-2floor western styles flat free 1set at Russain market", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 6, ClassTitle: "appartement for rent in Phnom Penh / April-October 2015", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 7, ClassTitle: "2 flat houses Olympic area ID:WALLC-10596", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 8, ClassTitle: "ដោះស្រាយរាល់បញ្ហាក្នុងការវិនិយោគ Forex", CategoryField: "BUSINESS - Lease or Sale"},
			{ID: 9, ClassTitle: "===//=== Very good condiction apartment for rent Near riveside 400$", CategoryField: "LAND - For Sale"},
			{ID: 10, ClassTitle: "Italian restaurant Pizza for rent", CategoryField: "BUSINESS - Lease or Sale"},
			{ID: 11, ClassTitle: "Small Villa for sale in Sihanouk Ville", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 12, ClassTitle: "House for Rent, 4br,2bh, near Derm Tkov Market,350$/m", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 13, ClassTitle: "Villa for Rent, 3Br,4Bh, 800$/m, Near Russian Market", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 14, ClassTitle: "New Flat for rent", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 15, ClassTitle: "C - Car Rental Service", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 16, ClassTitle: "Flat for sale!!!!!!!!!!!! 115,000$ Negotiate", CategoryField: "HOUSES/FLATS - For Sale"},
			{ID: 17, ClassTitle: "095 999 168, 095 777 168,... FOR SALE", CategoryField: "TELEPHONES"},
			{ID: 18, ClassTitle: "Cambodian Plastic Cup, Paper Cup and Paper Bowl supplier", CategoryField: "GENERAL - For Sale"},
			{ID: 19, ClassTitle: "Apartment FOR RENT near Orussey market $550", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 20, ClassTitle: "newly unit of condo for sale toul kork area affordable Price only $ 68000", CategoryField: "HOUSES/FLATS - For Sale"},
			{ID: 21, ClassTitle: "Lexus Rx300 Year 2001 Gold Color for Rental 1Pong 600$/Month", CategoryField: "VEHICLES - For Rent"},
			{ID: 22, ClassTitle: "2Bedroom near independent monument, $500/month, Quiet house", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 23, ClassTitle: "កាត់កញ្ចក់ តុ ទូ ទ្វារ បង្អួច អំពីអាលុយមីញូម គ្រប់ប្រភេទ", CategoryField: "SHOPS & SERVICES"},
			{ID: 24, ClassTitle: "Very Nice Apartment , 2bed room, 2bathroom only $430 in bkk1", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 25, ClassTitle: "0889997991, 0887778772, 0889919993, 0887775779, 0974448885", CategoryField: "TELEPHONES"},
			{ID: 26, ClassTitle: "CREATIVE TRANSLATION SERVICE Co., Ltd (Best Quality)", CategoryField: "SHOPS & SERVICES"},
			{ID: 27, ClassTitle: "Fully Furnished Apartment in Russian Market for Rent,2BR=$300/m", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 28, ClassTitle: "New and safe apartment near river side with 1bedroom,1bathroom=300$***", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 29, ClassTitle: "E0 flat for sale $15,000", CategoryField: "HOUSES/FLATS - For Sale"},
			{ID: 30, ClassTitle: "Camry 03 Silver ABS auto seat plate 2Txxxx 13000$", CategoryField: "VEHICLES - For Rent"},
			{ID: 31, ClassTitle: "Nice big Balcony furnished 1BR Apartment free wifi,cable TV:$280/m", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 32, ClassTitle: "looking for new biz owner- transferring Sports Club business", CategoryField: "BUSINESS - Lease or Sale"},
			{ID: 33, ClassTitle: "Normal Apartment For Rent Near Toul Tom poung market $ 300", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 34, ClassTitle: "Land for sale or Rent near Borey Phnom Penh Thmey Sen Sok", CategoryField: "LAND - For Sale"},
			{ID: 35, ClassTitle: "2beds apat furished with big balcony, plants, flowers.$400/m Russian Market", CategoryField: "VEHICLES - For Rent"},
			{ID: 36, ClassTitle: "New furnished Apartment/$500per month/2BR/TTP Russian market", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 37, ClassTitle: "Bentley Watches For Sale - New", CategoryField: "GENERAL - For Sale"},
			{ID: 38, ClassTitle: "New Apartment for Rent Near CIA School In Sen Sok", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 39, ClassTitle: "Nice 3BR furnished Apartment near River side:$500/month free wifi See photo", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 40, ClassTitle: "+++++5bed=$350 nice flat for rent at Beong Toum Porn++++++++", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 41, ClassTitle: "New Flat 1 bedroom =230$ per month, Near Olympic Stadium", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 42, ClassTitle: "=====2 bed apartment WITH CAR PARK market with furnished 430$ per month", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 43, ClassTitle: "2 BEDROOMS serviceAPARTMENT=450$ PER MONTH, See PIC", CategoryField: "HOUSES/FLATS - For Rent"},
			{ID: 44, ClassTitle: "1room apartment for rent:$150/month near river side", CategoryField: "HOUSES/FLATS - For Rent"},
		},
	}
}

func (d *classifiedData) GetAll() []models.Classified{
	return d.items
}

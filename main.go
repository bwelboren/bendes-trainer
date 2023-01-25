package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

const (

	// Found free code
	foundCode = "VIND"

	// Health
	foodBought = "Product gekocht!"
	foodEaten  = "Weer wat energie erbij!"

	// Energy
	pillUsed = "Weer wat gezondheid erbij!"

	// Warnings
	energyWarning = "Je hebt niet genoeg energie"
	foodWarning   = "Niet gebruikt, eet en drink wat rustiger."

	// Hospital
	freeMedWarning = "Je hebt gratis medicijnen gekregen, je moet nu 30 seconden wachten voordat de medicijnen werken."
)

func buyFood(loginCookie []*http.Cookie) {

	foods := map[string]string{
		"menu":          "https://www.bendes.nl/?go=winkel&winkel=burgerbar&buy=512",
		"fries":         "https://www.bendes.nl/?go=winkel&winkel=burgerbar&buy=511",
		"donut":         "https://www.bendes.nl/?go=winkel&winkel=burgerbar&buy=514",
		"chickenNugget": "https://www.bendes.nl/?go=winkel&winkel=burgerbar&buy=510",
		"chickenBurger": "https://www.bendes.nl/?go=winkel&winkel=burgerbar&buy=508",
		"hamburger":     "https://www.bendes.nl/?go=winkel&winkel=burgerbar&buy=509",
		"milkshake":     "https://www.bendes.nl/?go=winkel&winkel=burgerbar&buy=513",
	}

	food := colly.NewCollector()

	food.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	food.OnResponse(func(r *colly.Response) {
		fmt.Println("Eten kopen.")
	})

	food.OnHTML("p", func(e *colly.HTMLElement) {
		// If product is bought succesfully, we can start eating
		foodId, _ := strconv.Atoi(foods["menu"][len(foods["menu"])-3:])
		message := e.Text
		if strings.Contains(message, foodBought) {
			eatFood(foodId, loginCookie)
		}
	})

	food.SetCookies(foods["menu"], loginCookie)
	food.Visit(foods["menu"])

}

func buyPill(loginCookie []*http.Cookie) {

	pills := map[string]string{
		"painkiller": "http://www.bendes.nl/?go=winkel&winkel=apotheek&buy=576",
		"coughsirup": "http://www.bendes.nl/?go=winkel&winkel=apotheek&buy=577",
	}

	pill := colly.NewCollector()

	pill.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	pill.OnResponse(func(r *colly.Response) {
		fmt.Println("Pijnstiller kopen.")
	})

	// If we bought the pill succesfully we can start using it.
	pill.OnHTML("p", func(e *colly.HTMLElement) {
		pillId, _ := strconv.Atoi(pills["painkiller"][len(pills["painkiller"])-3:])
		message := e.Text
		if strings.Contains(message, pillUsed) {
			eatPill(pillId, loginCookie)
		}
	})

	pill.SetCookies(pills["painkiller"], loginCookie)
	pill.Visit(pills["painkiller"])

}

func eatPill(pillId int, loginCookie []*http.Cookie) {

	eatLink := fmt.Sprintf("https://www.bendes.nl/?go=eigendom&gebruik=%d", pillId)
	eat := colly.NewCollector()

	eat.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	eat.OnResponse(func(r *colly.Response) {
		fmt.Println("Pijnstiller slikken.")
	})

	// If we used the pill succesfully return to the crime function.
	eat.OnHTML("p", func(e *colly.HTMLElement) {
		message := e.Text
		if strings.Contains(message, pillUsed) {

		}
	})

	eat.SetCookies(eatLink, loginCookie)
	eat.Visit(eatLink)
}

func eatFood(foodId int, loginCookie []*http.Cookie) {
	fmt.Println("eating food")

	eatLink := fmt.Sprintf("https://www.bendes.nl/?go=eigendom&gebruik=%d", foodId)
	eat := colly.NewCollector()

	eat.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	// If we ate the food succesfully return to the crime function.
	eat.OnHTML("p", func(e *colly.HTMLElement) {
		message := e.Text
		if strings.Contains(message, foodEaten) {
			fmt.Println("Gegeten, nu uitbuiken..")

		}
	})

	eat.SetCookies(eatLink, loginCookie)
	eat.Visit(eatLink)

	time.Sleep(5 * time.Second)

}

func hospital(loginCookie []*http.Cookie) {

	treatments := map[string]string{
		"free":     "http://www.bendes.nl/?go=ziekenhuis&a=gratis",
		"paidMeds": "http://www.bendes.nl/?go=ziekenhuis&a=betaald",
		"bang":     "http://www.bendes.nl/?go=ziekenhuis&a=neuk",
	}

	hospital := colly.NewCollector()

	hospital.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	hospital.SetCookies(treatments["paidMeds"], loginCookie)
	hospital.Visit(treatments["paidMeds"])

}

func buyGun(loginCookie []*http.Cookie) {

	weapons := map[string]string{
		"m16": "https://www.bendes.nl/?go=winkel&winkel=tonie&buy=108",
	}

	weapon := colly.NewCollector()

	weapon.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	weapon.OnHTML(".melding", func(e *colly.HTMLElement) {

		fmt.Println("m16 gekocht.")
		// crime(loginCookie)

	})

	weapon.SetCookies(weapons["m16"], loginCookie)
	weapon.Visit(weapons["m16"])

}

// Before this function can even run, it did some checks in the crime function to see
// if the player is in jail & has hunger etc.
// So no need to do checks in this function
// Because it is done in crime()
func work(loginCookie []*http.Cookie) {
	fmt.Println("work")

	jobs := map[string]string{
		"workForMafia":      "https://www.bendes.nl/?go=werk&type=&kies=20&w=a&d=e&m=a&v=a&r=p",
		"workForGovernment": "https://www.bendes.nl/?go=werk&type=&kies=21&w=a&d=e&m=a&v=a&r=p",
	}

	job := colly.NewCollector()

	job.OnRequest(func(r *colly.Request) {
		stats(loginCookie)
		checkJailTime(loginCookie)
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	job.OnHTML("p", func(e *colly.HTMLElement) {
		message := e.Text
		if strings.Contains(message, "€") {
			fmt.Println(message)
		} else if strings.Contains(message, energyWarning) {
			fmt.Println("[!]", energyWarning)
			buyFood(loginCookie)
			job.SetCookies(jobs["workForGovernment"], loginCookie)
			job.Visit(jobs["workForGovernment"])

		} else {
			fmt.Println(message + " ????????????? ")
		}
	})

	job.OnHTML(".melding", func(e *colly.HTMLElement) {
		message := e.Text
		if strings.Contains(message, "verlaten") {
			job.SetCookies(jobs["workForGovernment"], loginCookie)
			job.Visit(jobs["workForGovernment"])
		}
	})

	job.OnResponse(func(r *colly.Response) {
		message := string(r.Body)
		if strings.Contains(message, "gek") {
			fmt.Println("Niet te hard werken. 2 seconden straf. Foei.")
			time.Sleep(2 * time.Second)
			job.SetCookies(jobs["workForGovernment"], loginCookie)
			job.Visit(jobs["workForGovernment"])
		}
		if strings.Contains(string(r.Body), "niet-PRO") {
			fmt.Println("Paywall")
			time.Sleep(60 * time.Second)
		}

	})

	//je wordt gek als je achter elkaar werkt.
	job.SetCookies(jobs["workForGovernment"], loginCookie)
	job.Visit(jobs["workForGovernment"])

	crime(loginCookie)

}

func crime(loginCookie []*http.Cookie) {
	fmt.Println("crime")

	crimes := map[string]string{
		"stealFromKid":       "https://www.bendes.nl/?go=misdaad&type=sloeber&kies=1&w=a&d=e&m=a&v=a&r=p",
		"creditCardTheft":    "https://www.bendes.nl/?go=misdaad&type=echt&kies=8&w=a&d=e&m=a&v=a&r=p",
		"robToesant":         "https://www.bendes.nl/?go=misdaad&type=echt&kies=11&w=a&d=e&m=a&v=a&r=p",
		"fireworkRobbery":    "http://www.bendes.nl/?go=misdaad&type=echt&kies=19&w=a&d=e&m=a&v=a&r=p",
		"atmRobbery":         "https://www.bendes.nl/?go=misdaad&type=echt&kies=14&w=a&d=e&m=a&v=a&r=p",
		"carTheft":           "https://www.bendes.nl/?go=misdaad&type=echt&kies=12&w=a&d=e&m=a&v=a&r=p",
		"breakInVilla":       "https://www.bendes.nl/?go=misdaad&type=echt&kies=10&w=a&d=e&m=a&v=a&r=p",
		"fuelStationRobbery": "https://www.bendes.nl/?go=misdaad&type=echt&kies=15&w=a&d=e&m=a&v=a&r=p",
	}

	crime := colly.NewCollector()

	crime.OnResponse(func(r *colly.Response) {

		if strings.Contains(string(r.Body), "Je ligt in het ziekenhuis") {
			hospital(loginCookie)
		}
		if strings.Contains(string(r.Body), "niet-PRO") {
			fmt.Println("Paywall")
			time.Sleep(60 * time.Second)
		}
	})

	crime.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	crime.OnHTML(".error", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "Even wachten") {
			fmt.Println("Eem wachten")
			time.Sleep(2 * time.Second)
			crime.Visit(crimes["fuelStationRobbery"])

		}
	})

	crime.OnHTML("p", func(e *colly.HTMLElement) {

		message := e.Text
		if strings.Contains(message, energyWarning) {
			fmt.Println("[!]", energyWarning)
			buyFood(loginCookie)
		} else if strings.Contains(message, "€") || strings.Contains(message, "buit") || strings.Contains(message, "Buit") {
			fmt.Println(message)
			time.Sleep(20 * time.Second)
		} else if strings.Contains(message, "voor je iets kan doen.") {
			rxp := regexp.MustCompile(`\d+`)
			pr := rxp.FindString(message)
			fmt.Printf("[!] Kom eerst even tot rust, nog eens proberen in %s seconden.\n", pr)
			timeLeft, _ := strconv.Atoi(pr)
			time.Sleep(time.Second*time.Duration(timeLeft) + 3)
		} else if strings.Contains(message, "Oops") {
			fmt.Println("Oops")
			checkJailTime(loginCookie)
		} else {
			fmt.Println(message)
			fmt.Println("20 seconden break.")
			time.Sleep(20 * time.Second)
			crime.SetCookies(crimes["fuelStationRobbery"], loginCookie)
			crime.Visit(crimes["fuelStationRobbery"])
			stealScooter(loginCookie)

		}

	})

	// Default cooldown

	crime.SetCookies(crimes["fuelStationRobbery"], loginCookie)
	crime.Visit(crimes["fuelStationRobbery"])

	defer work(loginCookie)

}

func checkJailTime(loginCookie []*http.Cookie) {

	fmt.Println("checkJailTime")
	jailLink := "https://www.bendes.nl/?go=gevangenis"

	jail := colly.NewCollector()

	jail.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	jail.OnHTML("#dl", func(e *colly.HTMLElement) {
		message := e.Text
		if strings.Contains(message, "Je moet nog") {
			// inJail = true
			rxp := regexp.MustCompile(`[0-9]+`)
			pr := rxp.FindString(message)
			timeLeft, _ := strconv.Atoi(pr)
			fmt.Printf("%s.\n", pr)
			time.Sleep(time.Second*(time.Duration(timeLeft)) + 5)

		}
	})

	jail.SetCookies(jailLink, loginCookie)
	jail.Visit(jailLink)

}

// Doesnt work
func depositMoney(loginCookie []*http.Cookie) {

	log.Panic("Wat")

	bank := colly.NewCollector()

	//Het geld is op je bankrekening bijgeschreven.

	bank.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	bankLink := "https://bendes.nl/?go=bank&bedrag=18000&do=storten"

	err := bank.Post(bankLink, map[string]string{"bedrag": "18000", "do": "storten"})
	if err != nil {
		log.Fatal(err)
	}

	bank.OnResponse(func(r *colly.Response) {

	})

	bank.OnHTML("p", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "bijgeschreven") {
			fmt.Println(e.Text)

		}

	})

	bank.SetCookies("https://bendes.nl/?go=bank", loginCookie)
	bank.Visit("https://bendes.nl/?go=bank")
}

func login() []*http.Cookie {
	c := colly.NewCollector()

	user := ""
	pass := ""

	err := c.Post("https://bendes.nl/?go=inloggen", map[string]string{"gebruikersnaam": user, "wachtwoord": pass})
	if err != nil {
		log.Fatal(err)
	}

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	// On response take the session cookie and start doing crimes
	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(string(r.Body), user) {
			log.Println("Logged in.")
		}
	})

	c.Visit("https://bendes.nl")
	return c.Cookies("https://bendes.nl/?go=inloggen")
}

func stats(loginCookie []*http.Cookie) {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	c.OnHTML(".sc3", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "€") {
			fmt.Println("-------------------------")
			fmt.Println("Jouw geld:\t", e.Text)
		}

		rxp := regexp.MustCompile(`[0-9]+`)
		pr := rxp.FindString(e.Text)
		money, _ := strconv.Atoi(pr)

		if money >= 66000 {
			fmt.Println("Money >= 66000.")
			buyGun(c.Cookies("https://www.bendes.nl/?go=winkel&winkel=tonie&buy=108"))
		}

	})

	c.OnHTML(".sc4", func(e *colly.HTMLElement) {
		ervaring := strings.Split(e.Text, ":")[1][1:]
		rank := ""
		fmt.Println("Jouw ervaring:\t", ervaring)

		ervaringInt, _ := strconv.Atoi(ervaring)

		//https://info-bendes.webnode.nl/rangen-ervaring/
		switch true {
		case ervaringInt >= 0 && ervaringInt < 12:
			rank = "Zwerver"
		case ervaringInt >= 12 && ervaringInt < 25:
			rank = "Nietsnut"
		case ervaringInt >= 25 && ervaringInt < 40:
			rank = "Kruimeldief"
		case ervaringInt >= 40 && ervaringInt < 70:
			rank = "Winkeldief"
		case ervaringInt >= 70 && ervaringInt < 140:
			rank = "Loopjongen"
		case ervaringInt >= 140 && ervaringInt < 250:
			rank = "Crimi"
		case ervaringInt >= 250 && ervaringInt < 500:
			rank = "Gangster"
		case ervaringInt >= 500 && ervaringInt < 750:
			rank = "Mafioso"
		case ervaringInt >= 750 && ervaringInt < 1250:
			rank = "Master"
		case ervaringInt >= 1250 && ervaringInt < 2500:
			rank = "Capacino"
		case ervaringInt >= 2500 && ervaringInt < 5000:
			rank = "Don"
		case ervaringInt >= 5000 && ervaringInt < 8000:
			rank = "Godfather"
		case ervaringInt >= 8000 && ervaringInt < 15000:
			rank = "Ruler"
		case ervaringInt >= 15000 && ervaringInt < 30000:
			rank = "Coltelo"
		case ervaringInt >= 30000 && ervaringInt < 60000:
			rank = "Kingpin"
		case ervaringInt >= 60000 && ervaringInt < 120000:
			rank = "Dictator"
		case ervaringInt >= 480000:
			rank = "Master of B&D"
		}

		fmt.Println("Jouw rank:\t", rank)

		fmt.Println("-------------------------")

	})

	c.SetCookies("https://www.bendes.nl/?go=overzicht", loginCookie)
	c.Visit("https://www.bendes.nl/?go=overzicht")

}

func stealScooter(loginCookie []*http.Cookie) {
	c := colly.NewCollector()

	stealScooterLink := "https://www.bendes.nl/?go=scooter"

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	})

	// On response take the session cookie and start doing crimes
	c.OnResponse(func(r *colly.Response) {
		// 10 minute cooldown
		time.Sleep(600 * time.Second)
	})

	c.SetCookies(stealScooterLink, loginCookie)
	c.Visit(stealScooterLink)
}

func main() {
	session := login()
	work(session)

}

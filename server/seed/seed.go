package seed

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path"
	"renthome/boiler"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

type Seeder struct {
	Conn       *sql.DB
	SeedFolder string
}

// NewSeeder returns a new Seeder
func NewSeeder(conn *sql.DB, seedFolder string) *Seeder {
	s := &Seeder{conn, seedFolder}
	return s
}

// Run for database spinup
func (s *Seeder) Run(isProd bool) error {
	ctx := context.Background()

	fmt.Println("Seeding agencies")
	err := seedAgencies(ctx, s.Conn)
	if err != nil {
		fmt.Println("Failed seeding agencies")
	}

	fmt.Println("Seeding notifications")
	err = seedNotifications(ctx, s.Conn)
	if err != nil {
		fmt.Println("Failed seeding notifications")
		return err
	}

	fmt.Println("Seeding privacies")
	err = seedPrivacies(ctx, s.Conn)
	if err != nil {
		fmt.Println("Failed seeding privacies")
		return err
	}

	fmt.Println("Seeding users")
	err = seedUsers(ctx, s.Conn)
	if err != nil {
		fmt.Println("Failed seeding users")
		return err
	}

	fmt.Println("Seeding locations")
	err = seedLocations(ctx, s.Conn, s.SeedFolder)
	if err != nil {
		fmt.Println("Failed to seed locations")
		return err
	}

	fmt.Println("Seed complete")
	return nil
}

const memberID = "38b6df11-5abb-498e-94c8-3b765ff0db40"
const superadminID = "fb0da5e6-834d-4a88-963c-c8cdd3d92528"

func seedUsers(ctx context.Context, conn *sql.DB) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	password := "devdev"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return err
	}

	// Insert superadmin
	superAdmin := &boiler.User{
		ID:          superadminID,
		FirstName:   "Nancy",
		LastName:    "Smith",
		Email:       null.StringFrom("superadmin@example.com"),
		Title:       null.StringFrom("Developer"),
		Description: null.StringFrom("Developer of renthome.com"),
		IsVerified:  true,
		Role:        "ADMIN",
	}

	err = superAdmin.Insert(tx, boil.Infer())
	if err != nil {
		return err
	}

	superAdminPasswordHash := &boiler.PasswordHash{
		PasswordHash: string(hash),
		UserID:       superAdmin.ID,
	}

	err = superAdminPasswordHash.Insert(tx, boil.Infer())
	if err != nil {
		return err
	}

	// Insert member
	member := &boiler.User{
		ID:          memberID,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       null.StringFrom("member@example.com"),
		Title:       null.StringFrom("Registered Nurse"),
		Description: null.StringFrom("Member of renthome.com"),
		IsVerified:  true,
		Role:        "MEMBER",
	}

	err = member.Insert(tx, boil.Infer())
	if err != nil {
		return err
	}

	agencies, err := boiler.Agencies().All(tx)
	if err != nil {
		return err
	}

	// for each agency create one admin and one property manager
	for _, agency := range agencies {

		lowerCase := strings.ToLower(strings.ReplaceAll(agency.Name, " ", ""))

		agencyAdmin := &boiler.User{
			FirstName:   "Admin",
			LastName:    "Admin",
			Email:       null.StringFrom(lowerCase + "@example.com"),
			Title:       null.StringFrom("Property Manager"),
			Description: null.StringFrom("Manager of " + agency.Name),
			IsVerified:  true,
			Role:        "AGENCY",
			AgencyID:    null.StringFrom(agency.ID),
		}

		propertyManager := &boiler.User{
			FirstName:   "Property",
			LastName:    "Manager",
			Email:       null.StringFrom("pm" + lowerCase + "@example.com"),
			Title:       null.StringFrom("Property Manager"),
			Description: null.StringFrom("Property Manger of " + agency.Name),
			IsVerified:  true,
			Role:        "MANAGER",
			AgencyID:    null.StringFrom(agency.ID),
		}

		err = agencyAdmin.Insert(tx, boil.Infer())
		if err != nil {
			return err
		}

		err = propertyManager.Insert(tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	memberPasswordHash := &boiler.PasswordHash{
		PasswordHash: string(hash),
		UserID:       member.ID,
	}

	err = memberPasswordHash.Insert(tx, boil.Infer())
	if err != nil {
		return err
	}

	notifications, err := boiler.Notifications().All(tx)
	if err != nil {
		return err
	}

	privacies, err := boiler.Privacies().All(tx)
	if err != nil {
		return err
	}

	for _, notification := range notifications {
		userNotification := &boiler.UserNotification{
			UserID:         member.ID,
			NotificationID: notification.ID,
		}
		err = userNotification.Insert(tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	for _, privacy := range privacies {
		userPrivacy := &boiler.UserPrivacy{
			UserID:    member.ID,
			PrivacyID: privacy.ID,
		}

		err = userPrivacy.Insert(tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func seedNotifications(ctx context.Context, conn *sql.DB) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	notifications := []boiler.Notification{
		{
			Name:        "Your property journey",
			Slug:        "your-property-journey",
			Method:      "Email",
			Category:    "Property journey",
			Description: "Recommended property information and tools based on your searches and activity.",
		},
		{
			Name:        "Saved search alerts",
			Slug:        "saved-search-alerts-email",
			Method:      "Email",
			Category:    "Properties",
			Description: "Manage what alerts you get when you've saved a search.",
		},
		{
			Name:        "Saved search alerts",
			Slug:        "saved-search-alerts-push",
			Method:      "Push",
			Category:    "Properties",
			Description: "Manage what alerts you get when you've saved a search.",
		},
		{
			Name:        "Property updates",
			Slug:        "property-updates",
			Method:      "Push",
			Category:    "Properties",
			Description: "Notifications about properties you've shown interest in.",
		},
		{
			Name:        "Promoted residential properties",
			Slug:        "promoted-residential-properties",
			Method:      "Email",
			Category:    "Properties",
			Description: "Notifications about residential properties relevant to your search.",
		},
		{
			Name:        "Promoted new development",
			Slug:        "promoted-new-developement",
			Method:      "Email",
			Category:    "Properties",
			Description: "Recommended new developments and property projects based on your searches and activity.",
		},
		{
			Name:        "Market updates",
			Slug:        "market-updates",
			Method:      "Email",
			Category:    "Property market",
			Description: "Market data, recent sales, auction results and updates on properties you like.",
		},
		{
			Name:        "Sales and auction results",
			Slug:        "sales-and-auction-results",
			Method:      "Push",
			Category:    "Property market",
			Description: "Latest auction results and property sales.",
		},
		{
			Name:        "Property news and guides",
			Slug:        "property-news-and-guides",
			Method:      "Email",
			Category:    "Property market",
			Description: "The latest property news, guides and inspiration.",
		},
		{
			Name:        "Property finance",
			Slug:        "property-finance",
			Method:      "Email",
			Category:    "Finance",
			Description: "Finance updates and tools like calculators and guides.",
		},
	}

	for _, notification := range notifications {
		err = notification.Insert(tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func seedPrivacies(ctx context.Context, conn *sql.DB) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	privacies := []boiler.Privacy{
		{
			Name:        "Personalised advertising",
			Slug:        "personalised-advertising",
			Description: "Advertising tailored to you based on your activity and the information you've provided. If you opt out, you'll still get ads but they won't be tailored to you.",
		},
		{
			Name:        "Suggested properties",
			Slug:        "suggested-properties",
			Description: "Property suggestions that match your activity and searches.",
		},
		{
			Name:        "Property updates",
			Slug:        "property updates",
			Description: "Relates to the bell icon and notifications about your saved properties you've shown interest in.",
		},
	}

	for _, privacy := range privacies {
		err = privacy.Insert(tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func seedAgencies(ctx context.Context, conn *sql.DB) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	agencies := []boiler.Agency{
		{
			Name:  "Ray White",
			Color: "#f4f4f3",
		},
		{
			Name:  "LJ Hooker",
			Color: "#e3f6f1",
		},
		{
			Name:  "Harcourts",
			Color: "#bfdedd",
		},
	}

	for _, agency := range agencies {
		err = agency.Insert(tx, boil.Infer())
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

type LocationCSV struct {
	Suburb   string `csv:"suburb"`
	Postcode string `csv:"postcode"`
	State    string `csv:"state"`
}

func seedLocations(ctx context.Context, conn *sql.DB, seedFolder string) error {
	locations := []*LocationCSV{}
	f, err := os.Open(path.Join(seedFolder, "/suburbs.csv"))
	if err != nil {
		return nil
	}
	defer f.Close()

	if err = gocsv.UnmarshalFile(f, &locations); err != nil {
		return terror.Error(err, "unmarshal suburbs csv")
	}

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, loc := range locations {
		location := &boiler.Location{
			Suburb:      loc.Suburb,
			Postcode:    loc.Postcode,
			State:       loc.State,
			Description: fmt.Sprintf("%s, %s %s", loc.Suburb, loc.State, loc.Postcode),
		}

		location.Insert(tx, boil.Infer())
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

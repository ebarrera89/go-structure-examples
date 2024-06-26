package memory

import (
	"fmt"
	"github.com/katzien/go-structure-examples/new/storage"
	"log"
	"time"

	"github.com/katzien/go-structure-examples/new/adding"
	"github.com/katzien/go-structure-examples/new/listing"
	"github.com/katzien/go-structure-examples/new/reviewing"
)

// Memory storage keeps data in memory
type Storage struct {
	beers   []Beer
	reviews []Review
}

// Add saves the given beer to the repository
func (m *Storage) AddBeer(b adding.Beer) error {
	id, err := storage.GetID("beer")
	if err != nil {
		log.Fatal(err)
	}

	newB := Beer{
		ID:        id,
		Created:   time.Now(),
		Name:      b.Name,
		Brewery:   b.Brewery,
		Abv:       b.Abv,
		ShortDesc: b.ShortDesc,
	}
	m.beers = append(m.beers, newB)

	return nil
}

// Add saves the given review in the repository
func (m *Storage) AddReview(r reviewing.Review) error {
	found := false
	for b := range m.beers {
		if m.beers[b].ID == r.BeerID {
			found = true
		}
	}

	if found {
		created := time.Now()
		id := fmt.Sprintf("%d_%s_%s_%d", r.BeerID, r.FirstName, r.LastName, created.Unix())

		newR := Review{
			ID:        id,
			Created:   created,
			BeerID:    r.BeerID,
			FirstName: r.FirstName,
			LastName:  r.LastName,
			Score:     r.Score,
			Text:      r.Text,
		}

		m.reviews = append(m.reviews, newR)
	} else {
		return listing.ErrNotFound
	}

	return nil
}

// Get returns a beer with the specified ID
func (m *Storage) GetBeer(id string) (listing.Beer, error) {
	var beer listing.Beer

	for i := range m.beers {

		if m.beers[i].ID == id {
			beer.ID = m.beers[i].ID
			beer.Name = m.beers[i].Name
			beer.Brewery = m.beers[i].Brewery
			beer.Abv = m.beers[i].Abv
			beer.ShortDesc = m.beers[i].ShortDesc
			beer.Created = m.beers[i].Created

			return beer, nil
		}
	}

	return beer, listing.ErrNotFound
}

// GetAll return all beers
func (m *Storage) GetAllBeers() []listing.Beer {
	var beers []listing.Beer

	for i := range m.beers {

		beer := listing.Beer{
			ID:        m.beers[i].ID,
			Name:      m.beers[i].Name,
			Brewery:   m.beers[i].Brewery,
			Abv:       m.beers[i].Abv,
			ShortDesc: m.beers[i].ShortDesc,
			Created:   m.beers[i].Created,
		}

		beers = append(beers, beer)
	}

	return beers
}

// GetAllReviews returns all reviews for a given beer
func (m *Storage) GetAllReviews(beerID string) []listing.Review {
	var list []listing.Review

	for i := range m.reviews {
		if m.reviews[i].BeerID == beerID {
			r := listing.Review{
				ID:        m.reviews[i].ID,
				BeerID:    m.reviews[i].BeerID,
				FirstName: m.reviews[i].FirstName,
				LastName:  m.reviews[i].LastName,
				Score:     m.reviews[i].Score,
				Text:      m.reviews[i].Text,
				Created:   m.reviews[i].Created,
			}

			list = append(list, r)
		}
	}

	return list
}

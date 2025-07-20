package apiserver

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mateuszmidor/FlightFinder/pkg/application"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
)

func GetRoutes() Routes {
	return routes
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "map.html", nil)
}

func GetAirportByIATACode(c *gin.Context) {
	_airportsSVC, ok := c.Get("airports")
	if !ok {
		err := errors.New("airports svc is missing")
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorToJSON(err))
		return
	}

	airportsSVC, ok := _airportsSVC.(*application.AirportFinder)
	if !ok {
		err := fmt.Errorf("airports svc is of wrong type: %T", _airportsSVC)
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorToJSON(err))
		return
	}

	code := strings.ToUpper(c.Param("code"))
	airport, err := airportsSVC.ByIATACode(code)

	// FIND ERROR
	if errors.Is(err, application.NotFoundError) {
		log.Printf("ERROR: %v\n", err)
		c.AbortWithStatusJSON(http.StatusNotFound, errorToJSON(err))
		return
	}
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorToJSON(err))
		return
	}

	// FIND OK
	c.JSON(200, fromAirport(airport))
}

func GetAirports(c *gin.Context) {
	_airportsSVC, ok := c.Get("airports")
	if !ok {
		err := errors.New("airports svc is missing")
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorToJSON(err))
		return
	}

	airportsSVC, ok := _airportsSVC.(*application.AirportFinder)
	if !ok {
		err := fmt.Errorf("airports svc is of wrong type: %T", _airportsSVC)
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorToJSON(err))
		return
	}

	getAirports(airportsSVC, c)
}

// FindFromToConnection - Flight connections by from and to airport IATA codes
func FindFromToConnection(c *gin.Context) {
	_finder, ok := c.Get("finder")
	if !ok {
		err := errors.New("finder svc is missing")
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorToJSON(err))
		return
	}

	finder, ok := _finder.(*application.ConnectionFinder)
	if !ok {
		err := fmt.Errorf("finder svc is of wrong type: %T", _finder)
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorToJSON(err))
		return
	}

	find(finder, c)
}

func getAirports(svc *application.AirportFinder, c *gin.Context) {
	airports := []Airport{}
	for _, a := range svc.AllAirports() {
		airports = append(airports, fromAirport(a))
	}
	c.JSON(200, airports)
}

func fromAirport(a airports.Airport) Airport {
	return Airport{Code: a.Code(), Name: a.Name(), Nation: a.Nation(), NationFullName: a.Nation(), Lon: float32(a.Longitude()), Lat: float32(a.Latitude())}
}

func find(finder *application.ConnectionFinder, c *gin.Context) {
	from := getFromAirportCode(c.Request)
	to := getToAirportCode(c.Request)
	segmentLimit := getMaxSegmentsCount(c.Request)
	log.Printf("%s -> %s (segment limit: %d)...", from, to, segmentLimit)

	buff := &bytes.Buffer{}
	pathRenderer := NewPathRendererAsJSON(buff)
	err := finder.Find(from, to, segmentLimit, pathRenderer)

	// FIND ERROR
	if err != nil {
		log.Printf("%s -> %s: ERROR: %v", from, to, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorToJSON(err))
		return
	}

	// FIND OK
	numConnections := strings.Count(buff.String(), "from_airport")
	log.Printf("%s -> %s: %d connections found (segment limit: %d)", from, to, numConnections, segmentLimit)

	// return connections
	c.Data(200, "application/json", buff.Bytes())
}

func getFromAirportCode(r *http.Request) string {
	return strings.ToUpper(r.FormValue("from"))
}

func getToAirportCode(r *http.Request) string {
	return strings.ToUpper(r.FormValue("to"))
}

func getMaxSegmentsCount(r *http.Request) int {
	count, _ := strconv.Atoi(r.FormValue("maxsegmentcount"))
	return count
}

func errorToJSON(err error) interface{} {
	type ErrorJSON struct {
		Error string `json:"error"`
	}

	return ErrorJSON{Error: err.Error()}
}

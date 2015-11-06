package main

import (
	"fmt"
	"github.com/att-innovate/boomer/handler"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Boomer")
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},

	Route{"AddUser", "POST", "/user", handler.AddUser},
	Route{"GetUser", "GET", "/user/{Userid}", handler.GetUser},
	Route{"GetUsers", "GET", "/user", handler.GetUsers},
	Route{"UpdateUser", "POST", "/user/{Userid}", handler.UpdateUser},
	Route{"DeleteUser", "DELETE", "/user/{Userid}", handler.DeleteUser},
	Route{"GetUserCpe", "GET", "/user/{Userid}/cpe", handler.GetUserCpe},
	Route{"Login", "GET", "/login", handler.Login},

	Route{"AddCpe", "POST", "/cpe", handler.AddCpe},
	Route{"UpdateCpe", "POST", "/cpe/{Cpeid}", handler.UpdateCpe},
	Route{"DeleteUser", "DELETE", "/cpe/{Cpeid}", handler.DeleteCpe},
	Route{"GetCpe", "GET", "/cpe/{Cpeid}", handler.GetCpe},
	Route{"GetCpes", "GET", "/cpe", handler.GetCpes},

	Route{"AddTopo", "POST", "/topo", handler.AddTopo},
	Route{"GetTopo", "GET", "/topo/{Topoid}", handler.GetTopo},
	Route{"GetTopos", "GET", "/topo", handler.GetTopos},
	Route{"UpdateTopo", "POST", "/topo/{Topoid}", handler.UpdateTopo},
	Route{"DeleteTopo", "DELETE", "/topo/{Topoid}", handler.DeleteTopo},

	Route{"AddVnf", "POST", "/vnf", handler.AddVnf},
	Route{"GetVnf", "GET", "/vnf/{Vnfid}", handler.GetVnf},
	Route{"GetVnfs", "GET", "/vnf", handler.GetVnfs},
	Route{"UpdateVnf", "POST", "/vnf/{Vnfid}", handler.UpdateVnf},
	Route{"DeleteVnf", "DELETE", "/vnf/{Vnfid}", handler.DeleteVnf},

	// TODO:: Implement  Route
	Route{"GetCpeTopo", "GET", "/cpe/{Cpeid}/topo", handler.GetCpeTopo},

	// TODO: APIs for the VNFs we have
	//APIs ffor the Router and FW

}

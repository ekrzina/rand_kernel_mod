package main

import apiserve "rand_kernel_mod/pkg/api"

func main() {
	var apiHandler = apiserve.ApiHandler{
		Port:     "2315",
		FileName: "/proc/rngen",
	}
	apiHandler.HandleRequests()
}

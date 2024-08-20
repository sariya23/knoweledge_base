package main

import (
	"errors"
	"fmt"
)

func main() {
	noShadowVarWiyhCopy()
}

func noShadowVarWiyhCopy() {
	tracing := false
	client := "no client"
	if tracing {
		tmpClient, err := createClientWithTracing()
		if err != nil {
			return
		}
		client = tmpClient
		fmt.Println(client)
	} else {
		tmpClient, err := createDefaultClient()
		if err != nil {
			return
		}
		client = tmpClient
		fmt.Println(client)
	}
	fmt.Println(client)
}

func noShadowVar() {
	tracing := false
	client := "no client"
	var err error
	if tracing {
		client, err = createClientWithTracing()
		fmt.Println(client)
	} else {
		client, err = createDefaultClient()
		fmt.Println(client)
	}
	if err != nil {
		return
	}
	fmt.Println(client)
}

func shadowVar() {
	tracing := false
	client := "no client"

	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return
		}
		fmt.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return
		}
		fmt.Println(client)
	}
	fmt.Println(client)
}

func createClientWithTracing() (string, error) {
	return "client with tracing", nil
}

func createDefaultClient() (string, error) {
	return "default client", nil
}

func joinBad(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatanate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func joinGood(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}

	if s2 == "" {
		return "", errors.New("s2 is empty")
	}

	concat, err := concatanate(s1, s2)
	if err != nil {
		return "", err
	}

	if len(concat) > max {
		return concat[:max], err
	}

	return concat, nil
}

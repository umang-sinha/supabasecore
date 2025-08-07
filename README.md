# SupabaseCore for Go

<p align="center">
  <a href="https://github.com/umang-sinha/supabasecore/actions"><img src="https://github.com/umang-sinha/supabasecore/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
  <a href="https://pkg.go.dev/github.com/umang-sinha/supabasecore"><img src="https://pkg.go.dev/badge/github.com/umang-sinha/supabasecore.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/umang-sinha/supabasecore"><img src="https://goreportcard.com/badge/github.com/umang-sinha/supabasecore" alt="Go Report Card"></a>
  <a href="https://github.com/umang-sinha/supabasecore/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></a>
</p>

`supabasecore` is a community-driven, idiomatic Go client for Supabase. It is designed to be modular, typed, and a joy to use, providing a seamless developer experience for interacting with all of Supabase's features.

## Features

This SDK aims to provide complete coverage of the Supabase ecosystem:

- **Database**: A powerful and intuitive query builder for your Supabase database (PostgREST).
- **Auth**: Complete authentication flow including sign-up, sign-in (email/password, OTP, OAuth), session management, and JWT handling.
- **Storage**: Manage files with a simple API for uploads, downloads, and permission control.
- **Realtime**: Listen to database changes in real-time using WebSockets.
- **Edge Functions**: Easily invoke your custom serverless functions.

## Installation

To add the SDK to your project, use `go get`:

```sh
go get github.com/umang-sinha/supabasecore
```

## Quick Start

Here's a simple example of how to initialize the client and fetch data from a `countries` table.

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/umang-sinha/supabasecore"
)

func main() {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	// Initialize the client
	client := supabasecore.NewClient(supabaseURL, supabaseKey, nil)

	// Define a struct to hold your data
	type Country struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	var results []Country

	// Build and execute the query
	err := client.DB.From("countries").Select("*").Execute(&results)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	fmt.Println("Fetched countries:")
	for _, country := range results {
		fmt.Printf("- ID: %d, Name: %s\n", country.ID, country.Name)
	}
}
```

## Project Status & Roadmap

This project is currently under active development. The plan is to build the SDK in phases. Contributions are welcome at any stage!

- [ ] **Phase 1 – Foundation**
  - [ ] `supabase.NewClient(url, key)`
  - [ ] Modular layout: `auth`, `db`, `storage`, `functions`, `realtime`
  - [ ] Typed config and HTTP client abstraction

- [ ] **Phase 2 – Auth Module**
  - [ ] SignUp/SignIn with email + password
  - [ ] SignIn with OTP or OAuth
  - [ ] Session and JWT handling
  - [ ] Unit tests and example usage

- [ ] **Phase 3 – DB Query Builder**
  - [ ] `.From("table").Select().Eq().Limit().Execute()`
  - [ ] Build query string generation engine
  - [ ] Typed/untyped result handling

- [ ] **Phase 4 – Storage**
  - [ ] Upload/download file API
  - [ ] File listing
  - [ ] Permission management

- [ ] **Phase 5 – Realtime**
  - [ ] Phoenix channels with WebSocket
  - [ ] Subscribe to INSERT, UPDATE, DELETE
  - [ ] Deliver callbacks to clients

- [ ] **Phase 6 – Supabase Edge Functions**
  - [ ] Thin wrapper around POST-ing to custom edge function endpoints

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

Please see `CONTRIBUTING.md` for guidelines on how to get started.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

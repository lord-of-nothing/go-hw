package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand/v2"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Encoded struct {
	Encoded string `json:"encoded"`
}

type Decoded struct {
	Decoded string `json:"decoded"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", VersionHandle)
	mux.HandleFunc("/hard-op", HardOpHandle)
	mux.HandleFunc("/decode", DecodeHandle)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	shutdownTimeout := 20 * time.Second

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("failed to serve http server: %w", err)
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}

		return nil
	})

	err := group.Wait()
	if err != nil {
		return
	}
}

func VersionHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v1.0.0"))
}

func HardOpHandle(w http.ResponseWriter, r *http.Request) {
	sleepTime := rand.IntN(10) + 10
	time.Sleep(time.Duration(sleepTime) * time.Second)

	isOk := rand.Float32() < 0.5
	if isOk {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DecodeHandle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input Encoded
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rawData, err := base64.StdEncoding.DecodeString(input.Encoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output := Decoded{string(rawData)}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

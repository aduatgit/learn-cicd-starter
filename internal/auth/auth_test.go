package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Testfälle definieren
	tests := map[string]struct {
		authHeader string
		want       string
		wantErr    bool
	}{
		"empty":            {authHeader: "", want: "", wantErr: true},
		"invalid format":   {authHeader: "InvalidFormat", want: "", wantErr: true},
		"invalid prefix":   {authHeader: "WrongPrefix 12345", want: "", wantErr: true},
		"correct format":   {authHeader: "ApiKey 12345", want: "12345", wantErr: false},
		"lowercase prefix": {authHeader: "apikey 12345", want: "", wantErr: true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// Header erstellen
			headers := http.Header{}
			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			// Funktion aufrufen
			got, err := GetAPIKey(headers)

			// Überprüfen, ob das Ergebnis dem erwarteten Wert entspricht
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

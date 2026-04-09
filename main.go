// The toy-turn command runs a minimal TURN server using pion/turn.
package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/pion/turn/v5"
)

func main() {
	publicIP := os.Getenv("PUBLIC_IP")
	if publicIP == "" {
		log.Fatal("PUBLIC_IP environment variable is required")
	}

	realm := os.Getenv("REALM")
	if realm == "" {
		realm = "toy-turn"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3478"
	}

	username := os.Getenv("TURN_USER")
	password := os.Getenv("TURN_PASSWORD")
	if username == "" || password == "" {
		log.Fatal("TURN_USER and TURN_PASSWORD environment variables are required")
	}

	users := map[string][]byte{
		username: turn.GenerateAuthKey(username, realm, password),
	}

	udpConn, err := net.ListenPacket("udp4", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen on UDP: %v", err)
	}

	s, err := turn.NewServer(turn.ServerConfig{
		Realm: realm,
		AuthHandler: func(ra *turn.RequestAttributes) (string, []byte, bool) {
			key, ok := users[ra.Username]
			return ra.Username, key, ok
		},
		PacketConnConfigs: []turn.PacketConnConfig{
			{
				PacketConn: udpConn,
				RelayAddressGenerator: &turn.RelayAddressGeneratorStatic{
					RelayAddress: net.ParseIP(publicIP),
					Address:      "0.0.0.0",
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to create TURN server: %v", err)
	}

	log.Printf("TURN server listening on UDP :%s (realm=%s)", port, realm)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	if err := s.Close(); err != nil {
		log.Printf("failed to close server: %v", err)
	}
}
